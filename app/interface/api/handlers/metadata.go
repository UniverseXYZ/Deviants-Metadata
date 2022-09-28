package handlers

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/deviants-metadata/app/config"
	"github.com/deviants-metadata/app/contracts"
	"github.com/deviants-metadata/app/domain/metadata"
	"github.com/deviants-metadata/app/interface/dlt/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

func HandleMetadataRequest(ethClient *ethereum.EthereumClient, address string, configService *config.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		deviantsInstance, err := contracts.NewDeviants(common.HexToAddress(address), ethClient.Client)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		tokenId := r.URL.Query().Get("id")

		iTokenId, err := strconv.Atoi(tokenId)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		ownerOf, err := deviantsInstance.OwnerOf(nil, big.NewInt(int64(iTokenId)))

		var genomeInt *big.Int

		if ownerOf == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			msg := "Query for non-existing token"
			render.Status(r, 404)
			render.JSON(w, r, msg)
			log.Errorln(msg)
			return
		} else {
			genomeInt, err = deviantsInstance.GeneOf(nil, big.NewInt(int64(iTokenId)))
			if err != nil {
				render.Status(r, 500)
				render.JSON(w, r, err)
				log.Errorln(err)
				return
			}
		}

		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId, configService))
	}
}
