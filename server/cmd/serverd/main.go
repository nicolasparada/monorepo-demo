package main

import (
	"fmt"
	"os"

	"github.com/nicolasparada/monorepo-demo/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	svc := &server.Service{}
	usr := svc.User("1")
	fmt.Println(usr)
	return nil
}
