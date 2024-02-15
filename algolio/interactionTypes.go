package algolio

// Define specific interaction types as constants
const (
	KnowledgeSharing      = "KnowledgeSharing"
	TransactionProcessing = "TransactionProcessing"
	PeerDiscovery         = "PeerDiscovery"
)

// SpecificInteraction represents a concrete implementation of the Interaction interface,
// tailored to our defined interaction types.
type SpecificInteraction struct {
	Type    string
	Details map[string]string
}

func (si SpecificInteraction) GetType() string {
	return si.Type
}

func (si SpecificInteraction) GetDetails() map[string]string {
	return si.Details
}

// NewSpecificInteraction creates a new instance of a specific interaction.
func NewSpecificInteraction(interactionType string, details map[string]string) SpecificInteraction {
	return SpecificInteraction{
		Type:    interactionType,
		Details: details,
	}
}
