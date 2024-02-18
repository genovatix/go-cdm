package algolio

type Peer interface {
	GetID() string
	GetMetadata() map[string]string
	CommunicateWithPeer(peer Peer) error
}
