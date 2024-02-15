package algolio

import "fmt"

// Simple learning mechanism to adjust Algolio's peer discovery strategy
func (m *MockAlgolio) LearnFromInteraction(interaction Interaction) {
	// Example: Adjust discovery criteria based on successful interactions
	if interaction.GetType() == KnowledgeSharing {
		success, exists := interaction.GetDetails()["success"]
		if exists && success == "true" {
			fmt.Println("Learning from successful knowledge sharing: Adjusting peer discovery strategy.")
			// Adjust the discovery criteria, e.g., prioritize peers with matching interests
			m.DiscoveryCriteria["interest"] = interaction.GetDetails()["sharedInterest"]
		}
	}
}
