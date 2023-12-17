package server

import (
	"crypto/sha256"
	"log"
)

type MerkleNode struct {
	Hash  []byte
	Left  *MerkleNode
	Right *MerkleNode
}

type MerkleTree struct {
	Root *MerkleNode
}

func CreateMerkleTree(transactions []*Transaction) *MerkleTree {
	if len(transactions) == 0 {
		return nil
	}
	layer := make([]*MerkleNode, len(transactions))
	for i, tx := range transactions {
		layer[i] = &MerkleNode{tx.Hash(), nil, nil}
	}
	for len(layer) > 1 {
		layer = createMerkleLayer(layer)
	}
	return &MerkleTree{layer[0]}
}

func createParentMerkleNode(left *MerkleNode, right *MerkleNode) *MerkleNode {
	if left == nil || right == nil {
		log.Fatal("Can't create MerkleNode due to null children")
	}
	newNode := new(MerkleNode)
	newNode.Left = left
	newNode.Right = right
	newHash := sha256.Sum256(append(left.Hash, right.Hash[:]...))
	newNode.Hash = newHash[:]
	return newNode
}

func createMerkleLayer(lowerLayer []*MerkleNode) []*MerkleNode {
	if len(lowerLayer) == 1 {
		return lowerLayer
	}
	if len(lowerLayer)%2 == 1 {
		lowerLayer = append(lowerLayer, lowerLayer[len(lowerLayer)-1])
	}
	result := make([]*MerkleNode, len(lowerLayer)/2)
	for i := 0; i < len(lowerLayer); i += 2 {
		result[i/2] = createParentMerkleNode(lowerLayer[i], lowerLayer[i+1])
	}
	return result
}

func (tree *MerkleTree) GetMerklePath() [][]byte {
	// TODO (Phuc) return hash values of required nodes in merkle tree to verify transaction, this method used by full node (server)
	return nil
}
