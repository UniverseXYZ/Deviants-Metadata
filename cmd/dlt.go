package main

import (
	"os"

	"github.com/deviants-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
)

func connectToNodes() *ethereum.EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := ethereum.NewEthereumClient(nodeURL)

	if err != nil {
		log.Errorln("Error creating new Ethereum client: ", err)
	}

	log.Infoln("Successfully connected to ethereum client")

	return client
}
