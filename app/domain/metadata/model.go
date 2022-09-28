package metadata

import (
	"fmt"
	"github.com/deviants-metadata/app/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
)

const DEVIANTS_IMAGE_URL string = "https://storage.googleapis.com/deviants-upload/"
const DEVIANTS_SOURCE_URL string = "https://storage.googleapis.com/deviants-source/"
const DEVIANTS_ANIMATION_URL string = "https://storage.googleapis.com/deviants-iframe/"
const EXTERNAL_URL string = "https://metaversia.universe.xyz/"
const GENES_COUNT = 10
const BACKGROUND_GENE_COUNT int = 16
const BASE_GENES_COUNT int = 44
const HEAD_GENES_COUNT = 44
const HEADWEAR_GENES_COUNT = 31
const MUSIC_GENES_COUNT = 7
const SHOES_GENES_COUNT int = 16
const PANTS_GENES_COUNT int = 16
const TORSO_GENES_COUNT int = 16
const EYEWEAR_GENES_COUNT int = 13
const WEAPON_RIGHT_GENES_COUNT int = 26
const WEAPON_LEFT_GENES_COUNT int = 26

type Genome string
type Gene int
type StringAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type IntegerAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type FloatAttribute struct {
	TraitType   string  `json:"trait_type"`
	Value       float64 `json:"value"`
	DisplayType string  `json:"display_type"`
}

type IframeEndpoints struct {
	CharacterEndpoint  string `json:"character_endpoint"`
	BackgroundEndpoint string `json:"background_endpoint"`
	MusicEndpoint      string `json:"music_endpoint"`
}

func (g Gene) toPath() string {
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end, count int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	return gene % count
}

func getLeftHandGene(g string) int {
	return getGeneInt(g, -8, -6, WEAPON_LEFT_GENES_COUNT)
}

func getLeftHandGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getLeftHandGene(g)
	return StringAttribute{
		TraitType: "Left Hand",
		Value:     configService.LeftHand[gene],
	}
}

func getLeftHandGenePath(g string) string {
	gene := getLeftHandGene(g)
	return Gene(gene).toPath()
}

func getRightHandGene(g string) int {
	return getGeneInt(g, -14, -12, WEAPON_RIGHT_GENES_COUNT)
}

func getRightHandGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getRightHandGene(g)
	return StringAttribute{
		TraitType: "Right Hand",
		Value:     configService.RightHand[gene],
	}
}

func getRightHandGenePath(g string) string {
	gene := getRightHandGene(g)
	return Gene(gene).toPath()
}

func getHeadGene(g string) int {
	return getGeneInt(g, -16, -14, HEAD_GENES_COUNT)
}
func getBodyGene(g string) int {
	return getGeneInt(g, -4, -2, BASE_GENES_COUNT)
}

func getHeadGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getHeadGene(g)
	return StringAttribute{
		TraitType: "Head",
		Value:     configService.Head[gene],
	}
}

func getHeadGenePath(g string) string {
	gene := getHeadGene(g)
	return Gene(gene).toPath()
}

func getEyewearGene(g string) int {
	return getGeneInt(g, -18, -16, EYEWEAR_GENES_COUNT)
}

func getHeadWearGene(g string) int {
	return getGeneInt(g, -20, -18, HEADWEAR_GENES_COUNT)
}

func getMusicGene(g string) int {
	return getGeneInt(g, -22, -20, MUSIC_GENES_COUNT)
}

func getEyewearGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getEyewearGene(g)
	return StringAttribute{
		TraitType: "Eyewear",
		Value:     configService.Eyewear[gene],
	}
}

func getHeadWearGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getHeadWearGene(g)
	return StringAttribute{
		TraitType: "Headwear",
		Value:     configService.Headwear[gene],
	}
}

func getSpeciesGeneAttribute(g string, configService *config.Service) StringAttribute {
	//gene := getBaseGene(g)
	baseGeneAttribute := getHeadGeneAttribute(g, configService)
	return StringAttribute{
		TraitType: "Species",
		Value:     configService.Species[baseGeneAttribute.Value],
	}
}

func getMusicGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getMusicGene(g)
	return StringAttribute{
		TraitType: "Music",
		Value:     configService.Music[gene],
	}
}

func getBodyGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getBodyGene(g)
	return StringAttribute{
		TraitType: "Body",
		Value:     configService.Body[gene],
	}
}

func getEyewearGenePath(g string) string {
	gene := getEyewearGene(g)
	return Gene(gene).toPath()
}

func getHeadwearGenePath(g string) string {
	gene := getHeadWearGene(g)
	return Gene(gene).toPath()
}

func getMuiscGenePath(g string) string {
	gene := getMusicGene(g)
	return Gene(gene).toPath()
}

func getShoesGene(g string) int {
	return getGeneInt(g, -6, -4, SHOES_GENES_COUNT)
}

func getShoesGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getShoesGene(g)
	return StringAttribute{
		TraitType: "Footwear",
		Value:     configService.Footwear[gene],
	}
}

func getShoesGenePath(g string) string {
	gene := getShoesGene(g)
	return Gene(gene).toPath()
}

func getTorsoGene(g string) int {
	return getGeneInt(g, -12, -10, TORSO_GENES_COUNT)
}

func getTorsoGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getTorsoGene(g)
	return StringAttribute{
		TraitType: "Torso",
		Value:     configService.Torso[gene],
	}
}

func getTorsoGenePath(g string) string {
	gene := getTorsoGene(g)
	return Gene(gene).toPath()
}

func getPantsGene(g string) int {
	return getGeneInt(g, -10, -8, PANTS_GENES_COUNT)
}

func getPantsGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getPantsGene(g)
	return StringAttribute{
		TraitType: "Pants",
		Value:     configService.Pants[gene],
	}
}

func getPantsGenePath(g string) string {
	gene := getPantsGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGene(g string) int {
	return getGeneInt(g, -2, 0, BACKGROUND_GENE_COUNT)
}

func getBackgroundGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getBackgroundGene(g)
	return StringAttribute{
		TraitType: "Background",
		Value:     configService.Background[gene],
	}
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func getBaseGene(g string) int {
	return getGeneInt(g, -4, -2, BASE_GENES_COUNT)
}

func getBaseGeneAttribute(g string, configService *config.Service) StringAttribute {
	gene := getHeadGene(g)
	return StringAttribute{
		TraitType: "Character",
		Value:     cases.Title(language.Und).String(configService.Body[gene]),
	}
}

func getBaseGenePath(g string) string {
	gene := getHeadGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *config.Service, tokenId string) string {
	gStr := string(*g)
	gene := getHeadGene(gStr)
	return fmt.Sprintf("%v #%v", cases.Title(language.Und).String(configService.Body[gene]), tokenId)
}

func (g *Genome) description(configService *config.Service, tokenId string) string {
	gStr := string(*g)
	baseGeneAttribute := getHeadGeneAttribute(gStr, configService)
	description := configService.Descriptions[baseGeneAttribute.Value]
	resultDescription := strings.Replace(description, "NUMBER", tokenId, -1)
	return resultDescription
}

func IsLeftHandAbove(leftHand string) bool {
	switch leftHand {
	case "Bag of Herbs", "Cheese", "Chocolate Bar", "Dark Matter Beam", "Ice Cream Cone", "Peanut Butter", "Lizard", "Cat":
		return true
	}
	return false
}

func IsEyewearAbove(eyeWear string) bool {
	return eyeWear == "Scuba Gear"
}

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getHeadwearGenePath(gStr))
	res = append(res, getEyewearGenePath(gStr))
	headGenePath := getHeadGenePath(gStr)
	res = append(res, headGenePath)

	res = append(res, getRightHandGenePath(gStr))

	res = append(res, getTorsoGenePath(gStr))
	res = append(res, getPantsGenePath(gStr))
	res = append(res, getLeftHandGenePath(gStr))
	res = append(res, getShoesGenePath(gStr))
	// body should be same as head
	res = append(res, getBaseGenePath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *config.Service) []interface{} {
	gStr := string(*g)

	res := []interface{}{}
	res = append(res, getMusicGeneAttribute(gStr, configService))
	res = append(res, getHeadWearGeneAttribute(gStr, configService))
	res = append(res, getEyewearGeneAttribute(gStr, configService))

	res = append(res, getSpeciesGeneAttribute(gStr, configService))

	res = append(res, getRightHandGeneAttribute(gStr, configService))

	res = append(res, getTorsoGeneAttribute(gStr, configService))
	res = append(res, getPantsGeneAttribute(gStr, configService))
	res = append(res, getLeftHandGeneAttribute(gStr, configService))
	res = append(res, getShoesGeneAttribute(gStr, configService))
	// body should be same as the head
	res = append(res, getBaseGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))

	return res
}

func (g *Genome) Metadata(tokenId string, configService *config.Service) Metadata {
	var m Metadata
	genes := g.genes()

	m.Attributes = g.attributes(configService)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
	m.ExternalUrl = EXTERNAL_URL

	b := strings.Builder{}

	noBackgroundImgUrl := strings.Builder{}

	iframeURL := strings.Builder{}

	b.WriteString(DEVIANTS_IMAGE_URL) // Start with base url
	noBackgroundImgUrl.WriteString(DEVIANTS_IMAGE_URL)

	for _, gene := range genes {
		b.WriteString(gene)
		noBackgroundImgUrl.WriteString(gene)
		iframeURL.WriteString(gene)
	}

	b.WriteString(".jpg") // Finish with jpg extension

	noBackgroundImgUrl.WriteString("-no.png")

	iframeURL.WriteString(".html")

	imageURL := b.String()

	imageExists2d := imageExists(imageURL)

	gStr := string(*g)

	leftHand := getLeftHandGeneAttribute(gStr, configService)
	eyeWear := getEyewearGeneAttribute(gStr, configService)

	if !imageExists2d {
		generateAndSaveImage(genes, leftHand.Value, eyeWear.Value)
	}

	musicGene := getMuiscGenePath(gStr)

	m.Image = imageURL

	videoURL := DEVIANTS_SOURCE_URL + "videos/" + genes[9] + ".mp4"
	musicURL := DEVIANTS_SOURCE_URL + "music/" + musicGene + ".wav"

	stillURL := DEVIANTS_SOURCE_URL + "images/0/" + genes[9] + ".png"

	cid := generateAndSaveToIpfs(iframeURL.String(), videoURL, musicURL, noBackgroundImgUrl.String(), stillURL)

	m.AnimationUrl = "ipfs://" + cid

	iframeEndpoints := IframeEndpoints{
		CharacterEndpoint:  noBackgroundImgUrl.String(),
		BackgroundEndpoint: videoURL,
		MusicEndpoint:      musicURL,
	}

	m.IframeEndpoints = iframeEndpoints

	return m
}

type Metadata struct {
	Description     string          `json:"description"`
	Name            string          `json:"name"`
	Image           string          `json:"image"`
	Attributes      interface{}     `json:"attributes"`
	ExternalUrl     string          `json:"external_url"`
	AnimationUrl    string          `json:"animation_url"`
	IframeEndpoints IframeEndpoints `json:"iframe_endpoints"`
}
