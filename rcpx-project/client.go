package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
)

type Args struct {
	A int
	B int
}

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8989", "")
	clientNew := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer func(clientNew client.XClient) {
		err := clientNew.Close()
		if err != nil {
			fmt.Printf("Error closing client: %v\n", err)
		}
	}(clientNew)

	args := &Args{A: 10, B: 20}
	var reply int

	// Вызываем метод "Mul"
	err := clientNew.Call(context.Background(), "Mul", args, &reply)
	if err != nil {
		fmt.Printf("Failed to call: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", reply)
	}
}
