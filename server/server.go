package server

import (
	//"bytes"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/gavotte25/blockchain_lab1/utils"
)

// // Sample codes
// type Args struct {
// 	A, B int
// }

// type Arith int

// func (t *Arith) Multiply(args *Args, reply *int) error {
// 	*reply = args.A * args.B
// 	return nil
// }

type Args struct {
	MerkelPath *MerklePath
	Status     string
	BlockIndex int
}

const ServerAddress = "localhost:1234"

type Queue struct {
	transactions []*Transaction
	mu           sync.Mutex
}

func (q *Queue) push(tx *Transaction) {
	q.transactions = append(q.transactions, tx)
}

func (q *Queue) clear() {
	q.transactions = make([]*Transaction, 0)
}

type Service struct {
	blockchain     *Blockchain
	transactionIdx *TransactionIndex
	txStack        chan *Transaction
	done           chan bool
	queue          Queue
}

const MinBlockTransactions = 4 // the minimum transactions of block is 4 (for easily debug)
const cacheDir = "./server/database"

func (s *Service) init() {
	if s.blockchain == nil {
		s.blockchain = new(Blockchain)
		s.blockchain.Init()
	}
	err := s.loadBlockChainFromFile()
	if err != nil {
		log.Println(err.Error())
	}
	s.txStack = make(chan *Transaction)
	s.done = make(chan bool)
	s.queue = Queue{transactions: make([]*Transaction, 0)}
	go func() {
		//queue := Queue{transactions: make([]*Transaction, 0)}
		for {
			select {
			case <-s.done:
				return
			case transaction := <-s.txStack:
				s.queue.mu.Lock()
				s.queue.push(transaction)
				if len(s.queue.transactions) >= MinBlockTransactions {
					block := s.blockchain.AddBlock(s.queue.transactions)
					block.SaveBlockAsJSON(cacheDir)
					s.blockchain.SaveMetaDataFile(cacheDir)
					s.queue.clear()
				}
				s.queue.mu.Unlock()
			}
		}
	}()
}

func (s *Service) finish() {
	s.done <- true
}

func (s *Service) MakeTransaction(tx Transaction, result *bool) error {
	log.Println("MakeTransaction: ", string(tx.Data))
	s.txStack <- &tx
	*result = true
	return nil
}

func (s *Service) GetBlockchainVersion(_ string, version *int) error {
	//log.Println("GetBlockchainVersion")
	*version = len(s.blockchain.BlockArr)
	//s.blockchain.BlockArr[len(s.blockchain.BlockArr)-1].PrintTransaction()
	return nil
}

func (s *Service) SyncBlockchain(fromBlockIndex int, blocks *[]*Block) error {
	log.Println("SyncBlockchain: ", fromBlockIndex)
	*blocks = s.blockchain.BlockArr[fromBlockIndex:]
	return nil
}

func (s *Service) GetBlockchain(headerOnly bool, bc *Blockchain) error {
	log.Println("GetBlockchain")
	if headerOnly {
		*bc = *s.blockchain.getLightVersion()
	} else {
		*bc = *s.blockchain
	}
	return nil
}

// GetTransactionLocation returns array of block index and transaction index in that block. Return {-1, -1} if can't find
func (s *Service) GetTransactionLocation(tx *Transaction, location *[2]int) error {
	txHash := utils.GetStringEncode(tx.Hash())
	txIndex, exist := s.transactionIdx.Index[txHash]
	if exist {
		location = &[2]int{txIndex.BlockIndex, txIndex.TransactionIndex}
	} else {
		location = &[2]int{-1, -1}
	}
	return nil
}

func (s *Service) loadBlockChainFromFile() error {
	arr := utils.ReadFile("metadata.bc", cacheDir)
	if arr == nil {
		return errors.New("metadata can't be loaded or does not exist at " + cacheDir + "/metadata.bc")
	}

	txIndex := &TransactionIndex{Index: make(map[string]TransactionLocation)}
	s.transactionIdx = txIndex
	for blockIndex, blockFile := range arr {
		block, err := LoadBlockFromJSON(blockFile, cacheDir)
		if err != nil {
			return err
		}
		if block == nil {
			return errors.New("LoadBlockFromJSON failed to retrieve from: " + blockFile)
		}
		s.blockchain.BlockArr = append(s.blockchain.BlockArr, block)

		for txIndex, _ := range block.Transactions {
			tx := s.blockchain.BlockArr[blockIndex].Transactions[txIndex]
			stringTxHash := utils.GetStringEncode(tx.Hash())
			s.transactionIdx.Index[stringTxHash] = TransactionLocation{BlockIndex: blockIndex, TransactionIndex: txIndex}
		}

	}
	return nil
}

func (s *Service) GetBlock(blockIndex int, block *Block) error {
	log.Println("GetBlock")
	*block = *s.blockchain.GetBlock(blockIndex)
	return nil
}

func (s *Service) GetTransaction(txIndex [2]int, transaction *Transaction) error {
	log.Println("GetTransaction")
	blockIndex, transactionIndex := txIndex[0], txIndex[1]
	block := s.blockchain.GetBlock(blockIndex)
	*transaction = *block.GetTransaction(transactionIndex)
	return nil
}

func (s *Service) IsTransactionInChannel(tx *Transaction) bool {
	isFound := false
	for _, i := range s.queue.transactions {
		if bytes.Equal(i.Hash(), tx.Hash()) {
			isFound = true
			break
		}
	}
	return isFound
}

func (s *Service) VerifyTransaction(tx *Transaction, agrs *Args) error {
	log.Println("VerifyTransaction")
	// check in queue
	target_agrs := &Args{Status: "processing", BlockIndex: -1, MerkelPath: nil}
	isInQueue := s.IsTransactionInChannel(tx)
	if isInQueue {
		*agrs = *target_agrs
		return nil
	} else {
		hashCode := utils.GetStringEncode(tx.Hash())
		if value, ok := s.transactionIdx.Index[hashCode]; ok {
			blockIndex, txIndex := value.BlockIndex, value.TransactionIndex
			block := s.blockchain.GetBlock(blockIndex)
			target_agrs.Status = "found"
			target_agrs.MerkelPath = CreateMerkleTree(block.Transactions).GetMerklePath(txIndex)
			target_agrs.BlockIndex = blockIndex
			*agrs = *target_agrs
			fmt.Println(agrs.BlockIndex, agrs.MerkelPath, agrs.Status)
			return nil
		} else {
			target_agrs.Status = "not_found"
			*agrs = *target_agrs
			fmt.Println(agrs.Status)
			return nil
		}

	}
}

func Start() {
	//log.Println("Server started")
	service := new(Service)
	service.init()
	rpc.Register(service)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	//log.Println("Listening at: ", ServerAddress)
	go http.Serve(l, nil)
}
