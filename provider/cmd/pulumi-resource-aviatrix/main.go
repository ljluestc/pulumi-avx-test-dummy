package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	aviatrix "github.com/AviatrixSystems/pulumi-avx-test-dummy/provider"
	"github.com/AviatrixSystems/pulumi-avx-test-dummy/provider/pkg/version"
)

func main() {
	// Initialize the Pulumi provider with tfbridge.Main()
	// Parameters:
	// - "aviatrix": Provider name
	// - version.Version: Provider version from the version package
	// - aviatrix.TerraformBridgeProvider(): Complete provider configuration with all resources and data sources
	tfbridge.Main("aviatrix", version.Version, aviatrix.MinimalTerraformBridgeProvider())
}