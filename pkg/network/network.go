package network

import (
	"github.com/genovatix/algoliocdm/algolio"
	"github.com/genovatix/algoliocdm/pkg/cryptography"
	ped "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"go.dedis.ch/kyber/v3/suites"
)

type AlgoliosNetwork struct {
	Participants []*algolio.AlgolioImpl
	Suite        suites.Suite
}

func (net *AlgoliosNetwork) InitializeDKG(threshold int) error {
	for _, algolio := range net.Participants {
		km := cryptography.NewKeyManager()
		privKey, pubKey := km.GenerateKeys()
		algolio.PrivateKey = privKey
		algolio.PublicKey = pubKey

		dkg, err := ped.NewDistKeyGenerator(net.Suite, privKey, nil, algolio.THRESHOLD)
		if err != nil {
			return err
		}
		algolio.DKGInstance = dkg
	}
	return nil
}

func (net *AlgoliosNetwork) StartDKG() error {
	// Example: Broadcast initial deals to all participants
	for _, algolio := range net.Participants {
		deals, err := algolio.DKGInstance.Deals()
		if err != nil {
			return err
		}
		// Distribute each deal to the respective participant
		for _, deal := range deals {
			// Assuming a method to send a deal to a participant by ID
			err := net.(deal.Receiver, deal)
			if err != nil {
				return err
			}
		}
	}
	// Further steps to process responses and finalize the DKG...
	return nil
}

func (net *AlgoliosNetwork) SendDealToParticipant() {




}