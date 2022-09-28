package metadata

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

const inBackgroundVideoPath = "./cmd/in/background.mp4"
const inSoundAudioPath = "./cmd/in/track.wav"

const outOverlayedNoBackgroundImgPath = "./cmd/out/overlayed-no-background.png"
const outTempVideoPath = "./cmd/out/video-no-sound.mp4"
const outFinalVideoPath = "./cmd/out/video-with-sound.mp4"

// GenerateAndSaveVideo Not used anywhere. Current approach is just pointing the wav, mp4 and jpeg image to the iframe. Leaving this for reference.
func GenerateAndSaveVideo(genes []string) {
	revGenes := reverseGenesOrder(genes)

	backgroundVideoUrl := fmt.Sprintf("./videos/%s.mp4", revGenes[0])
	trackUrl := fmt.Sprintf("./music/%s.wav", revGenes[len(revGenes)-1])
	readAndSaveMedia(backgroundVideoUrl, inBackgroundVideoPath)
	log.Infoln("Loaded background video in memory!")
	readAndSaveMedia(trackUrl, inSoundAudioPath)
	log.Infoln("Loaded track in memory!")
	addBackground()
	log.Println("Background video + Deviant picture generated!")
	addAudio()
	log.Println("Track added to animation!")
	videoNameBuilder := strings.Builder{}
	for _, gene := range revGenes {
		videoNameBuilder.WriteString(gene)
	}

	videoNameBuilder.WriteString(".mp4")

	saveVideoToGCloud(outFinalVideoPath, videoNameBuilder.String())
	log.Println("Saved final video to bucket!")
}

func readAndSaveMedia(mediaUrl string, savePath string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	mediaReader, err := bucket.Object(mediaUrl).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	defer mediaReader.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(dst, mediaReader)
	if err != nil {
		log.Fatal(err)
	}
}

func addBackground() {
	// ffmpeg -y -i ./in/background.mp4 -framerate 30 -i ./out/combined.png -filter_complex [0]overlay=x=0:y=0[out] -map [out] -map 0:a? -tag:v hvc1 -vcodec libx265 -pix_fmt yuv420p ./out/video-no-sound-compressed.mp4
	toVideo := exec.Command("ffmpeg", "-y",
		"-i", inBackgroundVideoPath,
		"-framerate", "30",
		"-i", outOverlayedNoBackgroundImgPath,
		"-filter_complex", "[0]overlay=x=0:y=0[out]",
		"-map", "[out]",
		"-map", "0:a?",
		"-tag:v", "hvc1",
		"-vcodec", "libx265",
		"-pix_fmt", "yuv420p",
		outTempVideoPath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	toVideo.Stdout = &out
	toVideo.Stderr = &stderr
	log.Println("Starting to generate dope video")
	err := toVideo.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

func addAudio() {
	// ffmpeg -y -i ./out/video-no-sound-compressed.mp4 -i ./in/track.wav -map 0:v -map 1:a -c:v copy -shortest ./out/final.mp4
	addAudio := exec.Command("ffmpeg", "-y",
		"-i", outTempVideoPath,
		"-i", inSoundAudioPath,
		"-map", "0:v",
		"-map", "1:a",
		"-c:v", "copy",
		"-shortest",
		outFinalVideoPath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	addAudio.Stdout = &out
	addAudio.Stderr = &stderr
	err := addAudio.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

func saveVideoToGCloud(filePath string, fileName string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	log.Println("Opened file for upload")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(fileName).NewWriter(ctx)

	if _, err := io.Copy(bucket, file); err != nil {
		log.Fatalf("io.Copy: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}
}
