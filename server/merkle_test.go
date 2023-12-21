package server

import (
	"strconv"
	"testing"
)

func TestGetMerklePath(t *testing.T) {
	transactions := make([]*Transaction, 8)
	for i := 0; i < 8; i++ {
		tx := new(Transaction)
		tx.Data = []byte("transaction " + strconv.Itoa(i))
		transactions[i] = tx
	}
	tree := CreateMerkleTree(transactions)
	index := 5
	expect := make([][]byte, 3)
	expect[2] = tree.Root.Left.Hash
	tmp := tree.Root.Right
	expect[1] = tmp.Right.Hash
	tmp = tmp.Left
	expect[0] = tmp.Left.Hash
	result := tree.GetMerklePath(index)
	if len(result) != 3 {
		t.Errorf("Path does not have enough length")
	}
	for i := 0; i < 3; i++ {
		if string(expect[i][:]) != string(result[i][:]) {
			t.Errorf("Path mismatch!")
			break
		}
	}
}
