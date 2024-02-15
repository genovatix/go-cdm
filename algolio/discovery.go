package algolio

import "github.com/genovatix/algoliocdm/identity"

type Graph struct {
	nodes map[string]*Node // Keyed by address string
}

type Node struct {
	address  *identity.Address
	peers    map[string]*Node // Keyed by address string, representing connections
	gScore   int              // Cost from start to the node
	fScore   int              // Estimated cost from start to goal through this node
	previous *Node            // To reconstruct the path
}

func (g *Graph) AStar(startAddr, goalAddr string) (path []*Node) {
	openSet := make(map[string]*Node) // Nodes to be evaluated
	openSet[startAddr] = g.nodes[startAddr]

	// Initialize gScore and fScore for start node
	g.nodes[startAddr].gScore = 0
	g.nodes[startAddr].fScore = heuristic(g.nodes[startAddr], g.nodes[goalAddr])

	for len(openSet) > 0 {
		current := getNodeWithLowestFScore(openSet)
		if current.address.String() == goalAddr {
			return reconstructPath(current)
		}

		delete(openSet, current.address.String())
		for peerAddr, peerNode := range current.peers {
			var raw identity.RawAddress
			var key []byte
			var peerNodeAddr identity.RawAddress
			current.address.Unmarshal(key, current.address.Bytes(), &raw)
			peerNode.address.Unmarshal(key, peerNode.address.Bytes(), &peerNodeAddr)
			tentativeGScore := current.gScore + calculateCloseness(&raw, &peerNodeAddr)

			if tentativeGScore < peerNode.gScore {
				peerNode.previous = current
				peerNode.gScore = tentativeGScore
				peerNode.fScore = peerNode.gScore + heuristic(peerNode, g.nodes[goalAddr])
				if _, found := openSet[peerAddr]; !found {
					openSet[peerAddr] = peerNode
				}
			}
		}
	}

	return nil // Path not found
}

func getNodeWithLowestFScore(openSet map[string]*Node) *Node {
	var lowestNode *Node
	lowestFScore := int(^uint(0) >> 1) // Max int

	for _, node := range openSet {
		if node.fScore < lowestFScore {
			lowestNode = node
			lowestFScore = node.fScore
		}
	}

	return lowestNode
}

func reconstructPath(node *Node) (path []*Node) {
	for current := node; current != nil; current = current.previous {
		path = append([]*Node{current}, path...)
	}
	return path
}

func heuristic(node, goal *Node) int {
	// Example heuristic: calculate closeness based on cryptographic properties
	var raw identity.RawAddress
	var key []byte
	var peerNodeAddr identity.RawAddress
	node.address.Unmarshal(key, node.address.Bytes(), &raw)
	goal.address.Unmarshal(key, goal.address.Bytes(), &peerNodeAddr)
	return -calculateCloseness(&raw, &peerNodeAddr) // Inverse of closeness
}

func calculateCloseness(addr1, addr2 *identity.RawAddress) int {
	// Implement your logic based on cryptographic properties
	// This could involve comparing chain codes, public keys, etc.
	// For demonstration, a simple byte comparison is used.
	hash1 := identity.HashKeyAddress(addr1)
	hash2 := identity.HashKeyAddress(addr2)
	distance := 0
	for i := range hash1 {
		if hash1[i] != hash2[i] {
			distance++
		}
	}
	return distance
}
