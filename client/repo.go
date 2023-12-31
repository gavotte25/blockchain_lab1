package client

import (
	"log"
	"net/rpc"

	"github.com/gavotte25/blockchain_lab1/server"
)

const serverAddress = server.ServerAddress

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

func (r *Repo) makeTransaction(tx server.Transaction) (bool, error) {
	var processing bool
	err := r.client.Call("Service.MakeTransaction", tx, &processing)
	return processing, err
}

// getBlockchainVersion returns the length of blockchain in fullnode / server
func (r *Repo) getBlockchainVersion() (int, error) {
	var version int
	err := r.client.Call("Service.GetBlockchainVersion", "", &version)
	return version, err
}

func (r *Repo) getNewBlocks(fromIndex int) ([]*server.Block, error) {
	var blocks []*server.Block
	err := r.client.Call("Service.SyncBlockchain", fromIndex, &blocks)
	return blocks, err
}

func (r *Repo) getEntireBlockchain() (*server.Blockchain, error) {
	var bc server.Blockchain
	err := r.client.Call("Service.GetBlockchain", false, &bc) // Set headerOnly = false because not enough time to handle Transaction caching at Client side
	return &bc, err
}

func (r *Repo) saveBlockchainToDatabase(bc *server.Blockchain, dir string) error {
	err := bc.SaveMetaDataFile(dir)
	if err != nil {
		return err
	}
	for i := 0; i < len(bc.BlockArr); i++ {
		err = bc.BlockArr[i].SaveBlockAsJSON(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// getBlock
func (r *Repo) getBlock(blockIndex int) (*server.Block, error) {
	var block *server.Block
	err := r.client.Call("Service.GetBlock", blockIndex, &block)
	return block, err
}

// getTransaction
func (r *Repo) getTransaction(blockIndex int, transactionIndex int) (*server.Transaction, error) {
	var tx *server.Transaction
	txIndex := [2]int{blockIndex, transactionIndex}
	err := r.client.Call("Service.GetTransaction", txIndex, &tx)
	return tx, err
}

// verifyTransaction
func (r *Repo) verifyTransaction(transaction *server.Transaction) (*server.Args, error) {
	var args *server.Args
	err := r.client.Call("Service.VerifyTransaction", transaction, &args)
	return args, err
}
