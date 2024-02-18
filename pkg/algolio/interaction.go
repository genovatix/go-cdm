package algolio

// Interaction defines the interface for interactions between Algolios.
type Interaction interface {
	GetType() string               // Returns the type of interaction (e.g., "KnowledgeSharing", "TransactionProcessing")
	GetDetails() map[string]string // Returns details about the interaction
}

type MockInteraction struct{}

func (mi MockInteraction) GetType() string {
	return "Mock"
}

func (mi MockInteraction) GetDetails() map[string]string {
	return map[string]string{}
}
