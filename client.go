package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net/rpc"
// 	"os"

// 	//"github.com/gavotte25/blockchain_lab1/server"
// )

// func Start() {
// 	log.Println("Client started")
// 	reader := bufio.NewReader(os.Stdin)
// 	client, err := rpc.DialHTTP("tcp", "localhost:1234")
// 	if err != nil {
// 		log.Fatal("dialing:", err)
// 	}
// 	for {
// 		fmt.Println("Press any key to call API")
// 		_, err := reader.ReadString('\n')
// 		args := &server.Args{7, 8}
// 		var reply int
// 		err = client.Call("Arith.Multiply", args, &reply)
// 		if err != nil {
// 			log.Fatal("arith error:", err)
// 		}
// 		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
// 		fmt.Println("\n##################")
// 	}
// }
