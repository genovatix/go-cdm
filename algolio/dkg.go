package algolio

import (
	"fmt"
	ped "go.dedis.ch/kyber/v3/share/dkg/pedersen"
)

var PARTICIPANTS int
var THRESHOLD int

func (a *AlgolioImpl) DistributeDeals(participants []*AlgolioImpl) {
	deals, _ := a.DKGInstance.Deals()
	for _, p := range participants {
		if p.InternalID != a.InternalID {
			a.Deals[p.InternalID] = deals[p.InternalID]
		}
	}
}

// ProcessDeal processes an incoming deal from another Algolio.
func (a *AlgolioImpl) ProcessDeal(deal *ped.Deal) *ped.Response {
	response, _ := a.DKGInstance.ProcessDeal(deal)
	return response
}

func SimulateDKGProcess(participants []*AlgolioImpl) {
	// Distribute deals
	for _, algolio := range participants {
		algolio.DistributeDeals(participants)
	}

	// Each Algolio processes incoming deals and stores responses
	for _, receiver := range participants {
		for _, sender := range participants {
			if sender.InternalID != receiver.InternalID {
				deal := sender.Deals[receiver.InternalID]
				response := receiver.ProcessDeal(deal)
				sender.Responses[receiver.InternalID] = append(sender.Responses[receiver.InternalID], response)
			}
		}
	}

	// Process responses
	for _, algolio := range participants {
		for _, responses := range algolio.Responses {
			for _, response := range responses {
				_, _ = algolio.DKGInstance.ProcessResponse(response)
			}
		}
	}
}

func FinalizeDKG(participants []*AlgolioImpl) {
	for _, algolio := range participants {
		// Finalizing the DKG to get the shared public key and private share
		dkShare, _ := algolio.DKGInstance.DistKeyShare()
		fmt.Printf("Algolio %s: Public Key: %s\n", algolio.InternalID, dkShare.Public())
	}
}
