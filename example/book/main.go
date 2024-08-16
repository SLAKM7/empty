package main

import (
	"fmt"
	"os"

	"grpc-gateway-demo/example/book/gateway"
	"grpc-gateway-demo/example/book/server"
)

func main() {
	go func() {
		// grpc 启动
		err := server.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()
	// gateway 启动
	err := gateway.Run()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
