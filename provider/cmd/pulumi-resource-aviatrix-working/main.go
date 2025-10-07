package main

import (
	"fmt"
	"os"
	aviatrix "github.com/AviatrixSystems/pulumi-avx-test-dummy/provider"
)

func main() {
	if len(os.Args) < 2 {
		provider := aviatrix.NewWorkingProvider()
		provider.PrintInfo()
		os.Exit(0)
	}
	
	// Handle other commands
	fmt.Println("Aviatrix Pulumi Provider - Command not implemented yet")
	os.Exit(1)
}
