package metadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
)

const IpfsBaseUploadURL = `https://api.pinata.cloud/pinning/pinFileToIPFS`

type pinataBodyResponse struct {
	IPFSHash string `json:"IpfsHash"`
	PinSize  string `json:"PinSize"`
}

const IMG_SIZE = 3000
const GCLOUD_UPLOAD_BUCKET_NAME = "deviants-upload"
const GCLOUD_SOURCE_BUCKET_NAME = "deviants-source"
const IFRAME_HTMLS_BUCKET_NAME = "deviants-iframe"

type TemplateHTML struct {
	VideoURL     string
	MusicURL     string
	CharacterURL string
	StillURL     string
}

func imageExists(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.StatusCode != 404
}

func objectExists(url string, image bool) bool {

	if !image {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		iframeHtmlsBucketName := os.Getenv("IFRAME_HTMLS_BUCKET_NAME")

		if err != nil {
			log.Errorf("storage.NewClient: %v", err)
		}
		defer client.Close()

		bucket := client.Bucket(iframeHtmlsBucketName)
		exists, err := bucket.Object(url).If(storage.Conditions{DoesNotExist: true}).NewReader(ctx)
		if err != nil {
			log.Errorf("Couldn't open bucket object. URL: " + url)
		}
		if err := client.Close(); err != nil {
			fmt.Println("createFile: unable to close bucket %q, file %q: %v", err)
		}

		return exists != nil
	} else {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		deviantsUploadBucketName := os.Getenv("DEVIANTS_UPLOAD_BUCKET_NAME")

		if err != nil {
			log.Errorf("storage.NewClient: %v", err)
		}
		defer client.Close()

		bucket := client.Bucket(deviantsUploadBucketName)
		exists, err := bucket.Object(url).If(storage.Conditions{DoesNotExist: true}).NewReader(ctx)
		if err != nil {
			log.Errorf("Couldn't open bucket object. URL:" + url)
		}
		if err := client.Close(); err != nil {
			fmt.Println("createFile: unable to close bucket %q, file %q: %v", err)
		}

		return exists != nil
	}

}

//func objectExists(objectURL string) bool {
//	ctx := context.Background()
//	client, err := storage.NewClient(ctx)
//
//	if err != nil {
//		log.Errorf("storage.NewClient: %v", err)
//	}
//	defer client.Close()
//
//	bucket := client.Bucket(IFRAME_HTMLS_BUCKET_NAME)
//	stats, err := bucket.Object(objectURL).Attrs(ctx)
//
//	return stats != nil
//}

//// @generateIframeAnimation Queries iframe cloud function so that the animation_url is automatically generated from metadata function
//func generateIframeAnimation(tokenId *string) string {
//	getUrl := IFRAME_CLOUD_FUNCTION_URL + *tokenId
//	resp, err := http.Get(getUrl)
//	if err != nil {
//		log.Errorf("Couldn't query iframe-cloud-function. Original error: %v", err)
//	}
//
//	type iframeResponse struct {
//		AnimationUrl string `json:"animation_url"`
//	}
//
//	iframeResp := iframeResponse{}
//
//	decoder := json.NewDecoder(resp.Body)
//
//	err = decoder.Decode(&iframeResp)
//	return iframeResp.AnimationUrl
//}

func combineRemoteImages(bucket *storage.BucketHandle, basePath string, overlayPaths ...string) *image.NRGBA {

	ctx := context.Background()

	baseReader, err := bucket.Object(basePath).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer baseReader.Close()
	base, err := imaging.Decode(baseReader)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	dst := imaging.New(IMG_SIZE, IMG_SIZE, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	for _, op := range overlayPaths {
		r, err := bucket.Object(op).NewReader(ctx)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		defer r.Close()
		o, err := imaging.Decode(r)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		dst = imaging.Overlay(dst, o, image.Pt(0, 0), 1)
	}

	return dst
}

func reverseGenesOrder(genes []string) []string {
	res := make([]string, 0, len(genes))
	for i := len(genes) - 1; i >= 0; i-- {
		res = append(res, genes[i])
	}
	return res
}

func saveToGCloud(i *image.NRGBA, noBackground *image.NRGBA, name string, noBackgroundName string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}

	//f, err := os.OpenFile("./cmd/out/overlayed-no-background.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//err = imaging.Encode(f, noBackground, imaging.PNG)

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(name).NewWriter(ctx)
	// f, err := imaging.FormatFromFilename(name)
	// if err != nil {
	// 	log.Errorf("Format from filename: %v", err)
	// }
	err = imaging.Encode(bucket, i, imaging.JPEG, imaging.JPEGQuality(80))

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}

	bucketNoBackgroundWr := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(noBackgroundName).NewWriter(ctx)

	err = imaging.Encode(bucketNoBackgroundWr, noBackground, imaging.PNG)

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = bucketNoBackgroundWr.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}

	err = client.Close()
	if err != nil {
		log.Errorf("Client.Close: %v", err)
	}

}

func generateAndSaveImage(genes []string, leftHand string, eyeWear string) {
	// Reverse
	revGenes := reverseGenesOrder(genes)

	f := make([]string, len(genes)) // without the music gene

	specialRulesLeftHand := IsLeftHandAbove(leftHand)

	specialRulesEyewear := IsEyewearAbove(eyeWear)

	for i, gene := range revGenes {
		f[i] = fmt.Sprintf("./images/%v/%s.png", i, gene)
	}

	if specialRulesLeftHand {
		f[3], f[5] = f[5], f[3] // left hand above torso
		f[3], f[4] = f[4], f[3] // and also torso should always be above pants
	}

	if specialRulesEyewear {
		f[8], f[9] = f[9], f[8] // eyewear above headwear
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	i := combineRemoteImages(bucket, f[0], f[1:10]...)

	noBackgroundImg := combineRemoteImages(bucket, f[1], f[2:10]...)

	b := strings.Builder{}
	noBackgroundName := strings.Builder{}

	for _, gene := range genes {
		b.WriteString(gene)
		noBackgroundName.WriteString(gene)
	}

	noBackgroundName.WriteString("-no")

	b.WriteString(".jpg") // Finish with jpg extension
	noBackgroundName.WriteString(".png")

	saveToGCloud(i, noBackgroundImg, b.String(), noBackgroundName.String())
}

// generateAndSaveToIpfs Generates the deviant`s animation url and uploads it to IPFS
func generateAndSaveToIpfs(iframeURL string, videoURL string, musicURL string, characterURL string, stillURL string) (cid string) {

	iframeExistsGcloud := imageExists(DEVIANTS_ANIMATION_URL + iframeURL)

	tmpl := template.Must(template.ParseFiles("./serverless_function_source_code/index.html"))
	data := TemplateHTML{
		VideoURL:     videoURL,
		MusicURL:     musicURL,
		CharacterURL: characterURL,
		StillURL:     stillURL,
	}

	if !iframeExistsGcloud { // If false, write the animation html to GCloud
		ctx := context.Background()
		client, err := storage.NewClient(ctx)

		iframeHtmlsBucketName := os.Getenv("IFRAME_HTMLS_BUCKET_NAME")

		if err != nil {
			log.Errorf("storage.NewClient: %v", err)
		}
		defer client.Close()

		tpl := &bytes.Buffer{}

		bucket := client.Bucket(iframeHtmlsBucketName).Object(iframeURL).NewWriter(ctx)

		if err := tmpl.Execute(tpl, data); err != nil {
			fmt.Println(err)
		}

		bucket.Write(tpl.Bytes())
		if err := bucket.Close(); err != nil {
			fmt.Println("createFile: unable to close bucket %q, file %q: %v", err)
		}
	}

	f, err := os.Create("/tmp/iframe-go.html")
	if err != nil {
		log.Println("create file: ", err.Error())
	}
	err = tmpl.Execute(f, data)

	if err != nil {
		log.Print("Error executing html animation template: ", err.Error())
	}
	err = f.Close()
	if err != nil {
		log.Println("Error closing file %v %s. Original error ", f.Name(), err.Error())
	}

	PinataApiKey := os.Getenv("PINATA_API_KEY")
	PinataSecretKey := os.Getenv("PINATA_SECRET_KEY")

	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/tmp/iframe-go.html")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("/tmp/iframe-go.html"))

	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	httpClient := &http.Client{}

	req, err := http.NewRequest(method, IpfsBaseUploadURL, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("pinata_api_key", PinataApiKey)
	req.Header.Add("pinata_secret_api_key", PinataSecretKey)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send HTTP Post request to Pinata
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	pinataResponse := pinataBodyResponse{}

	err = decoder.Decode(&pinataResponse)

	httpClient.CloseIdleConnections()

	e := os.Remove("/tmp/iframe-go.html")
	if e != nil {
		log.Fatal(e)
	}

	return pinataResponse.IPFSHash
}
