package server

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

type Transaction struct {
	Data []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) SetHash() {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, b.Timestamp)
	if err != nil {
		log.Fatal("cast from int64 failed", err)
	}
	tsByte := buf.Bytes()
	hashInput := append(b.HashTransactions(), tsByte...)
	hashOutput := sha256.Sum256(hashInput)
	b.Hash = hashOutput[:]
}

func (b *Block) HashTransactions() []byte {
	var hashInput []byte
	for _, tx := range b.Transactions {
		hashTx := sha256.Sum256(tx.Data)
		hashInput = append(hashInput, hashTx[:]...)
	}
	hashOutput := sha256.Sum256(hashInput)
	return hashOutput[:]
}
