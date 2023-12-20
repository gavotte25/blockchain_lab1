package client

import (
	"log"
	"net/rpc"

	"github.com/gavotte25/blockchain_lab1/server"
)

const serverAddress = "localhost:1234"

type Repo struct {
	client *rpc.Client
}

func (r *Repo) init() {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	r.client = client
}

func (r *Repo) makeTransaction(txDetail string) bool {
	var succeed bool
	err := r.client.Call("Service.MakeTransaction", txDetail, &succeed)
	if err != nil {
		log.Fatal("Something wrong", err)
	}
	return succeed
}

// getBlockchainVersion returns the length of blockchain in fullnode / server
func (r *Repo) getBlockchainVersion() int {
	var version int
	err := r.client.Call("Service.GetBlockchainVersion", "", &version)
	if err != nil {
		log.Fatal("Something wrong", err)
	}
	return version
}

func (r *Repo) getNewBlocks(fromIndex int) []*server.Block {
	var blocks []*server.Block
	err := r.client.Call("Service.SyncBlockchain", fromIndex, &blocks)
	if err != nil {
		log.Fatal("Something wrong", err)
	}
	return blocks
}

func (r *Repo) getEntireBlockchain() *server.Blockchain {
	var bc server.Blockchain
	err := r.client.Call("Service.GetBlockchain", true, &bc)
	if err != nil {
		log.Fatal("Something wrong", err)
	}
	return &bc
}
