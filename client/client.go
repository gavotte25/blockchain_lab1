package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gavotte25/blockchain_lab1/server"
	"github.com/gavotte25/blockchain_lab1/utils"
)

const syncIntervalInSecond = 5
const cacheDir = "./client/database"

type Wallet struct {
	blockchain *server.Blockchain
	repo       *Repo
	ticker     *time.Ticker
	done       chan bool
	logger     utils.Logger
}

func NewWallet() *Wallet {
	wallet := new(Wallet)
	wallet.init(false)
	return wallet
}

func (w *Wallet) init(loggingEnabled bool) {
	w.repo = new(Repo)
	w.logger = utils.Logger{Enable: loggingEnabled}
	w.repo.init()
	w.blockchain = new(server.Blockchain)
	w.blockchain.Init()
	w.loadBlockchainDataFromFile()
	if w.blockchain.GetVersionNumber() == 0 {
		w.fetchEntireBlockchain()
	}
	w.sync()
}

func (w *Wallet) loadBlockchainDataFromFile() {
	w.logger.Println("started loadBlockchainDataFromFile ", cacheDir+"/metadata.bc")
	arr := utils.ReadFile("metadata.bc", cacheDir)
	if arr == nil {
		w.logger.Panicln("loadBlockchainDataFromFile failed, reason: can't load metadata from path ", cacheDir+"/metadata.bc")
	} else {
		for _, blockFile := range arr {
			block, err := server.LoadBlockFromJSON(blockFile, cacheDir)
			if err != nil {
				w.logger.Panicln("loadBlockchainDataFromFile failed, reason: ", err.Error())
				return
			}
			w.blockchain.BlockArr = append(w.blockchain.BlockArr, block)
		}
	}
}

func (w *Wallet) fetchEntireBlockchain() {
	w.logger.Println("fetchEntireBlockchain started")
	utils.WipeFolder(cacheDir)
	var err error
	w.blockchain, err = w.repo.getEntireBlockchain()
	if err != nil {
		w.logger.Println("fetchEntireBlockchain failed: ", err.Error())
		return
	}
	err = w.repo.saveBlockchainToDatabase(w.blockchain, cacheDir)
	if err != nil {
		w.logger.Panicln("fetchEntireBlockchain failed: ", err.Error())
	} else {
		w.logger.Println("fetchEntireBlockchain succeed")
	}

}

// sync periodically checks latest length of fullnode blockchain. If it's longer than local blockchain,
// it will fetch the missing blocks and check the current local last block is valid,
// if not valid, fetch the whole blockchain
func (w *Wallet) sync() {
	w.logger.Println("sync started at interval ", syncIntervalInSecond, " seconds")
	w.ticker = time.NewTicker(time.Second * syncIntervalInSecond)
	w.done = make(chan bool)
	go func() {
		for {
			select {
			case <-w.done:
				return
			case <-w.ticker.C:
				localVersion := w.blockchain.GetVersionNumber()
				fullNodeVersion, err := w.repo.getBlockchainVersion()
				if err != nil {
					w.logger.Panicln("cannot fetch latest blockchain version from server, reason:  ", err.Error())
					continue
				}
				if localVersion < fullNodeVersion {
					if localVersion < 2 {
						w.fetchEntireBlockchain()
					} else {
						newBlocks, err := w.repo.getNewBlocks(localVersion - 1)
						if err != nil {
							w.logger.Panicln("cannot fetch new blocks from server, reason:  ", err.Error())
							continue
						}
						if len(newBlocks) == 0 {
							continue
						}
						if w.blockchain.BlockArr[len(w.blockchain.BlockArr)-1].GetHash() == newBlocks[0].GetHash() {
							w.blockchain.Append(newBlocks[1:])
							for _, block := range newBlocks[1:] {
								block.SaveBlockAsJSON(cacheDir)
							}
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

func (w *Wallet) MakeTransaction(txDetail string) bool {
	tx := server.Transaction{Timestamp: time.Now().Unix(), Data: []byte(txDetail)}

	// save into history.tx
	f, err := os.OpenFile("./client/database/history.tx", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		w.logger.Println("Error create history.tx")
		return false
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%d\t%s\n", tx.Timestamp, string(tx.Data))
	if err != nil {
		w.logger.Println("Error save transaction history.tx")
		return false
	}

	// call server to makeTransaction
	success, err := w.repo.makeTransaction(tx)
	if err != nil {
		w.logger.Println("makeTransaction failed: ", err.Error())
		return success
	}
	if success {
		w.logger.Println("Transaction is being queued for processing")
	} else {
		w.logger.Println("Transaction is not accepted")
	}
	// Test
	// flag := w.VerifyTransaction(&tx)
	// fmt.Println("test flag: ", flag)
	return success
}

func (w *Wallet) Finish() {
	if w.ticker != nil {
		w.ticker.Stop()
	}
	if w.done != nil {
		w.done <- true
	}
}

// GetBlock
func (w *Wallet) GetBlock(bIndex int) *server.Block {
	success, err := w.repo.getBlock(bIndex)
	if err != nil {
		w.logger.Println("GetBlock failed: ", err.Error())
		return nil
	}
	return success
}

// GetTransaction
func (w *Wallet) GetTransaction(bIndex int, txIndex int) *server.Transaction {
	success, err := w.repo.getTransaction(bIndex, txIndex)
	if err != nil {
		w.logger.Println("GetBlock failed: ", err.Error())
		return nil
	}
	return success
}

// VerifyTransaction
func (w *Wallet) VerifyTransaction(tx *server.Transaction) bool {
	tx = w.blockchain.BlockArr[0].GetTransaction(0)
	args, err := w.repo.verifyTransaction(tx)

	if err != nil {
		w.logger.Println("Verify failed: ", err.Error())
		return false
	}
	if args.Status == "not_found" {
		w.logger.Println("Transaction is not found in entire blockchain")
		return false
	} else if args.Status == "processing" {
		w.logger.Println("Transaction is being queued for processing")
		return false
	} else {
		block := w.blockchain.GetBlock(args.BlockIndex)

		if block == nil {
			w.logger.Println("Need to synchronize data !")
			return false
		} else {
			// check transaction by verify merkel path from server
			isVerified := block.VerifyTransaction(tx, args.MerkelPath)
			if isVerified {
				txDetail := fmt.Sprintf("%d", tx.Timestamp) + string(tx.Data)
				utils.DeleteVerifiedRowInFile(cacheDir, "history.tx", txDetail)
			}
			return isVerified
		}
	}
}

func (w *Wallet) PrintBlock(bIndex int) bool {
	block := w.GetBlock(bIndex)
	if block == nil {
		fmt.Println("Your block does not exist.")
		return false
	}
	block.PrintInfo()
	//block.PrintTransaction()
	return true
}

func (w *Wallet) ReadTransactionFile() ([]*server.Transaction, error) {
	filePath := "./client/database/history.tx"
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Contents of file:", string(file))

	var transactions []*server.Transaction
	file_str := strings.Split(string(file), "\n")
	leng := len(file_str)
	for _, trans := range file_str[0 : leng-1] {
		//fmt.Println(trans)
		parts := strings.Split(trans, "\t")
		str_id := parts[0]
		content := parts[1]
		id, _ := strconv.ParseInt(str_id, 10, 64)
		transaction := &server.Transaction{Data: []byte(content), Timestamp: id}
		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, nil
	}
	return transactions, nil
}

func Start(loggingEnabled bool) {
	log.Println("Client started")
	reader := bufio.NewReader(os.Stdin)
	wallet := new(Wallet)
	wallet.init(loggingEnabled)
	for {
		fmt.Println("Type info and press enter to make transaction, type 'exit' to close")
		info, err := reader.ReadString('\n')
		info = utils.TrimInputByOS(info)
		if err != nil {
			log.Fatal(err.Error())
		}
		if info == "exit" {
			fmt.Printf("error %s", info)
			wallet.Finish()
			break
		}
    fmt.Printf("Is success: %t\n", wallet.MakeTransaction(info))
		fmt.Println("##################")
	}
}
