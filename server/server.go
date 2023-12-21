package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
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

type Service struct {
	blockchain *Blockchain
}

func (s *Service) init() {
	s.loadBlockchainDataFromFile()
	if s.blockchain == nil {
		s.blockchain = InitBlockchain()
	}
}

func (s *Service) MakeTransaction(txDetail string, result *bool) error {
	resultBool := s.blockchain.AddBlock(txDetail)
	result = &resultBool
	return nil
}

func (s *Service) loadBlockchainDataFromFile() {
	// TODO: assign blockchain value from file. If file does not exist, do nothing
	s.blockchain = nil
}

func (s *Service) GetBlockchainVersion(_ string, version *int) error {
	*version = len(s.blockchain.BlockArr)
	return nil
}

func (s *Service) SyncBlockchain(fromBlockIndex int, blocks *[]Block) error {
	fmt.Println("SyncBlockchain called")
	return nil
}

func (s *Service) GetBlockchain(headerOnly bool, bc *Blockchain) error {
	if headerOnly {
		bc = s.blockchain.getLightVersion()
	} else {
		bc = s.blockchain
	}
	return nil
}

func Start() {
	log.Println("Server started")
	service := new(Service)
	service.init()
	rpc.Register(service)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)
}
