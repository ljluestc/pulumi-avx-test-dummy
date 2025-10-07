package main

import (
	"fmt"
	"os"

	aviatrix "github.com/AviatrixSystems/pulumi-avx-test-dummy/provider"
)

func main() {
	// Create provider
	provider := aviatrix.Provider()
	
	// Print provider information
	fmt.Printf("Aviatrix Pulumi Schema Generator %s\n", provider.GetVersion())
	fmt.Printf("Description: %s\n", provider.GetDescription())
	fmt.Printf("Supported Resources: %d\n", len(provider.GetResources()))
	fmt.Printf("Supported Data Sources: %d\n", len(provider.GetDataSources()))
	
	// Print some example resources
	fmt.Println("\nExample Resources:")
	for i, resource := range provider.GetResources() {
		if i >= 10 { // Show only first 10
			fmt.Printf("... and %d more\n", len(provider.GetResources())-10)
			break
		}
		fmt.Printf("  - %s\n", resource)
	}
	
	fmt.Println("\nSchema generation completed successfully!")
	os.Exit(0)
}
