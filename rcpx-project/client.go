package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	clientNew := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer func(clientNew client.XClient) {
		err := clientNew.Close()
		if err != nil {

		}
	}(clientNew)

	args := &Args{A: 10, B: 20} // Используем структуру Args из server.go
	var reply int
	err := clientNew.Call(context.Background(), "Mul", args, &reply)
	if err != nil {
		fmt.Printf("failed to call: %v\n", err)
	}

	fmt.Printf("Result: %d\n", reply)
}
