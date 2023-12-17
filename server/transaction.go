package server

import "crypto/sha256"

type Transaction struct {
	Data      []byte
	Timestamp int64
}

func (tx *Transaction) Hash() []byte {
	hash := sha256.Sum256(tx.Data)
	return hash[:]
}
