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
	expect := new(MerklePath)
	steps := make([][]byte, 3)
	steps[2] = tree.Root.Left.Hash
	tmp := tree.Root.Right
	steps[1] = tmp.Right.Hash
	tmp = tmp.Left
	steps[0] = tmp.Left.Hash
	result := tree.GetMerklePath(index)
	expect.Steps = steps
	expect.Direction = []bool{false, true, false}
	if len(result.Steps) != 3 {
		t.Errorf("Path does not have enough length")
	}
	for i := 0; i < 3; i++ {
		if string(expect.Steps[i][:]) != string(result.Steps[i][:]) {
			t.Errorf("Steps mismatch!")
			break
		}
	}
	for i := 0; i < 3; i++ {
		if expect.Direction[i] != result.Direction[i] {
			t.Errorf("Directions mismatch!")
			break
		}
	}
}
