package cryptography

import (
	"github.com/consensys/gnark/frontend"
)

type AlgolioCircuit struct {
	Address frontend.Variable
	Data    frontend.Variable
	Weight  frontend.Variable
}

func (c *AlgolioCircuit) Define(api frontend.API) error {
	// Example constraint: Strength + Adaptability must be less than some value

	api.Println(c)

	return nil
}

/*
func Start(values *AlgolioCircuit) (groth16.Proof, error) {
	var algolioCircuit AlgolioCircuit

	r1cs1, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &algolioCircuit)
	if err != nil {
		log.Println(err)
	}
	pk, vk, err := groth16.Setup(r1cs1)
	if err != nil {
		log.Println(err)
	}

	var assignment AlgolioCircuit
	// assign message value
	assignment.Strength = values.Strength
	assignment.Adaptability = values.Adaptability
	assignment.BehaviorType = values.BehaviorType
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, err := witness.Public()
	// generate the proof
	proof, err := groth16.Prove(r1cs1, pk, witness)

	// verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		// invalid proof
		return nil, err
	}
	return proof, nil
}*/
