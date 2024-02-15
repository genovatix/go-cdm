package algolio

import (
	"bufio"
	"context"
	"fmt"
	"github.com/genovatix/algoliocdm/transaction"
	"go.dedis.ch/kyber/v3"
	ped "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"net"
	"strings"
)

type Algolio interface {
	ProcessTransaction(transaction.Transaction) error
	DiscoverPeers(targetAddr string)
	ShareKnowledge(transaction.Transaction) error

	ContributeToCommunityMemory(transaction.Transaction) error
	ProcessInteraction(interaction Interaction)
	LearnFromInteraction(interaction Interaction)
	Id() string
	Metadata() map[string]string
	Con(context.Context) net.Conn
	Log(level, message string)
	GetPeers() map[string]string
	StartServer()
	Peer
}

type AlgolioImpl struct {
	ID          string
	InternalID  int
	Conn        net.Conn
	Peers       map[string]string
	Addr        string
	DKGInstance *ped.DistKeyGenerator
	Deals       map[int]*ped.Deal
	Responses   map[int][]*ped.Response
	PrivateKey  kyber.Scalar
	PublicKey   kyber.Point
}

func (a *AlgolioImpl) GetPeers() map[string]string {
	return a.Peers
}

func (a *AlgolioImpl) ProcessTransaction(t transaction.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) DiscoverPeers(targetAddr string) {
	conn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", targetAddr, err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("DISCOVER\n"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		peerInfo := strings.Split(scanner.Text(), ":")
		if len(peerInfo) == 2 {
			a.GetPeers()[peerInfo[0]] = peerInfo[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading discovery response: %v\n", err)
	}

	fmt.Printf("Updated peers for %s: %v\n", a.ID, a.Peers)
}
func (a *AlgolioImpl) ShareKnowledge(t transaction.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) Communicate(peer interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) ContributeToCommunityMemory(t transaction.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) ProcessInteraction(interaction Interaction) {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) LearnFromInteraction(interaction Interaction) {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) Id() string {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) Metadata() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) Con(ctx context.Context) net.Conn {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) Log(level, message string) {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) GetID() string {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) GetMetadata() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (a *AlgolioImpl) CommunicateWithPeer(peer Peer) error {
	//TODO implement me
	panic("implement me")
}

func NewAlgolio(id, address string) Algolio {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := net.Dial("udp", address)
	if err != nil {
		GetLogger().Fatal(err.Error())
	}
	algolio := &AlgolioImpl{
		ID:        id,
		Conn:      conn,
		Peers:     make(map[string]string),
		Addr:      address,
		Deals:     make(map[int]*ped.Deal),
		Responses: make(map[int][]*ped.Response),
	}

	return algolio
}

func (a *AlgolioImpl) StartServer() {
	listener, err := net.Listen("tcp", a.Addr)
	if err != nil {
		fmt.Printf("Error starting server for %s: %v\n", a.ID, err)
		return
	}
	defer listener.Close()
	fmt.Printf("Algolio %s listening on %s\n", a.ID, a.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go a.handleConnection(conn)
	}
}

func (a *AlgolioImpl) handleConnection(conn net.Conn) {
	defer conn.Close()
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading message: %v\n", err)
		return
	}

	fmt.Printf("Received message at %s: %s", a.ID, message)
	if strings.TrimSpace(message) == "DISCOVER" {
		for peerID, peerAddr := range a.GetPeers() {
			conn.Write([]byte(fmt.Sprintf("%s:%s\n", peerID, peerAddr)))
		}
	}
}
