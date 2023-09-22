package main

import (
	"fmt"
	"os"

	"github.com/zakyyudha/jtnctl/cmd/jtnctl"
)

func main() {
	if err := jtnctl.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
