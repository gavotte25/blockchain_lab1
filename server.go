package main

// import (
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/rpc"
// )

// // Sample codes to create a server in goroutine, no need no start sever in other terminal
// type Args struct {
// 	A, B int
// }

// type Arith int

// func (t *Arith) Multiply(args *Args, reply *int) error {
// 	*reply = args.A * args.B
// 	return nil
// }

// func Start() {
// 	log.Println("Server started")
// 	arith := new(Arith)
// 	rpc.Register(arith)
// 	rpc.HandleHTTP()
// 	l, err := net.Listen("tcp", ":1234")
// 	if err != nil {
// 		log.Fatal("listen error:", err)
// 	}
// 	go http.Serve(l, nil)
// }
