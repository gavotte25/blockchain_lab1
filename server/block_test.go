package server

import (
	"testing"
	"time"
)

func TestVerifyTransaction(t *testing.T) {
	tx1 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx1")}
	tx2 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx2")}
	tx3 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx3")}
	tx4 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx4")}
	tx5 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx5")}
	tx6 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx6")}
	tx7 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx7")}
	tx8 := Transaction{Timestamp: time.Now().Unix(), Data: []byte("tx8")}
	transactions := []*Transaction{&tx1, &tx2, &tx3, &tx4, &tx5, &tx6, &tx7, &tx8}
	block := NewBlock(transactions, nil)
	tree := CreateMerkleTree(transactions)
	if !block.VerifyTransaction(&tx6, tree.GetMerklePath(5)) {
		t.Errorf("Expect true but result is false")
	}
	if block.VerifyTransaction(&tx3, tree.GetMerklePath(1)) {
		t.Errorf("Expect false but result is true")
	}
}
