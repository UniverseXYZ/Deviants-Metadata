package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Service struct {
	Character    []string          `json:"character"`
	Footwear     []string          `json:"footwear"`
	Pants        []string          `json:"pants"`
	Torso        []string          `json:"torso"`
	Eyewear      []string          `json:"eyewear"`
	Headwear     []string          `json:"headwear"`
	Head         []string          `json:"head"`
	RightHand    []string          `json:"righthand"`
	LeftHand     []string          `json:"lefthand"`
	Body         []string          `json:"body"`
	Background   []string          `json:"background"`
	Music        []string          `json:"music"`
	Descriptions map[string]string `json:"descriptions"`
	Species      map[string]string `json:"species"`
}

func NewConfigService(configPath string) *Service {
	jsonFile, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service Service

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &service)
	if err != nil {
		return nil
	}

	return &service
}
