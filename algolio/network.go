package algolio

import (
	"github.com/genovatix/algoliocdm/cryptography"
	ped "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"go.dedis.ch/kyber/v3/suites"
)

type AlgoliosNetwork struct {
	Participants []*AlgolioImpl
	Suite        suites.Suite
}

func (net *AlgoliosNetwork) InitializeDKG(threshold int) error {
	for _, algolio := range net.Participants {
		km := cryptography.NewKeyManager()
		privKey, pubKey := km.GenerateKeys()
		algolio.PrivateKey = privKey
		algolio.PublicKey = pubKey

		dkg, err := ped.NewDistKeyGenerator(net.Suite, privKey, nil, THRESHOLD)
		if err != nil {
			return err
		}
		algolio.DKGInstance = dkg
	}
	return nil
}
