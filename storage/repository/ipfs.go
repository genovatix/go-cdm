package repository

import (
	"context"
	"github.com/genovatix/algoliocdm/cryptography"
	"github.com/genovatix/algoliocdm/log"
	"github.com/genovatix/algoliocdm/storage/drivers/ipfs"
	"go.uber.org/zap"
)

func StartIpfs() {
	km := cryptography.NewKeyManager()
	ipfsDriver := ipfs.NewIPFSDriver("localhost:5001", km)
	privateKey, publicKey := km.GenerateKeys()
	data := []byte("Hello, World!")
	cid, err := ipfsDriver.Create(context.Background(), publicKey, data)
	if err != nil {
		log.Logger.Error("Failed to create data in IPFS", zap.Error(err))
		return
	}
	log.Logger.Info("Data created in IPFS", zap.String("cid", cid))

	retrievedData, err := ipfsDriver.Read(context.Background(), privateKey, cid)
	if err != nil {
		log.Logger.Error("Failed to retrieve data from IPFS", zap.Error(err))
		return
	}
	log.Logger.Info("Data retrieved from IPFS", zap.ByteString("data", retrievedData))
}
