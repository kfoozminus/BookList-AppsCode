package main

import (
	"log"

	"github.com/kfoozminus/BookList-AppsCode/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
