package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gavotte25/blockchain_lab1/server"
)

const syncIntervalInSecond = 5

type Wallet struct {
	blockchain *server.Blockchain
	repo       *Repo
	ticker     *time.Ticker
	done       chan bool
}

func (w *Wallet) init() {
	w.repo = new(Repo)
	w.repo.init()
	w.loadBlockchainDataFromFile()
	if w.blockchain == nil {
		w.fetchEntireBlockchain()
	}
	w.sync()
}

func (w *Wallet) loadBlockchainDataFromFile() {
	// TODO: assign blockchain value from file. If file does not exist, do nothing
	w.blockchain = nil
}

func (w *Wallet) fetchEntireBlockchain() {
	w.blockchain = w.repo.getEntireBlockchain()
}

// sync periodically checks latest length of fullnode blockchain. If it's longer than local blockchain,
// it will fetch the missing blocks and check the current local last block is valid,
// if not valid, fetch the whole blockchain
func (w *Wallet) sync() {
	fmt.Println("Blockchain sync is enabled")
	w.ticker = time.NewTicker(time.Second * syncIntervalInSecond)
	w.done = make(chan bool)
	go func() {
		for {
			select {
			case <-w.done:
				return
			case <-w.ticker.C:
				localVersion := w.blockchain.GetVersionNumber()
				fullNodeVersion := w.repo.getBlockchainVersion()
				fmt.Println("1")
				if localVersion < fullNodeVersion {
					if localVersion < 2 {
						w.fetchEntireBlockchain()
					} else {
						newBlocks := w.repo.getNewBlocks(localVersion - 1)
						if w.blockchain.BlockArr[len(w.blockchain.BlockArr)-1].GetHash() == newBlocks[0].GetHash() {
							w.blockchain.Append(newBlocks[1:])
						} else {
							w.fetchEntireBlockchain()
						}
					}
				} else if localVersion > fullNodeVersion {
					w.fetchEntireBlockchain()
				}
			}
		}
	}()
}

func (w *Wallet) makeTransaction(txDetail string) bool {
	return w.repo.makeTransaction(txDetail)
}

func (w *Wallet) finish() {
	if w.ticker != nil {
		w.ticker.Stop()
	}
	if w.done != nil {
		w.done <- true
	}
}

func Start() {
	log.Println("Client started")
	reader := bufio.NewReader(os.Stdin)
	wallet := new(Wallet)
	wallet.init()
	for {
		fmt.Println("Type info and press enter to make transaction, type 'exit' to close")
		info, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Something wrong", err)
		}
		if info == "exit\n" {
			wallet.finish()
			break
		}
		fmt.Printf("Is success: %t", wallet.makeTransaction(info))
		fmt.Println("\n##################")
	}
}
