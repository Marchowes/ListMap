package main

import (
	"fmt"
	"os"

	"github.com/Marchowes/ListMap/cmd"
)

func main() {
	// kick off the Cobra Command.
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
