package client

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	// "github.com/gavotte25/blockchain_lab1/server"
)

func Start() {
	log.Println("Client started")
	reader := bufio.NewReader(os.Stdin)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for {
		fmt.Println("Press any key to call API")
		_, err := reader.ReadString('\n')
		var succeed bool
		err = client.Call("Service.MakeTransaction", "Hello world", &succeed)
		if err != nil {
			log.Fatal("Something wrong", err)
		}
		fmt.Printf("Is success: %t", succeed)
		fmt.Println("\n##################")
	}
}
