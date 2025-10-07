package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	aviatrix "github.com/AviatrixSystems/pulumi-avx-test-dummy/provider"
)

func main() {
	tfgen.Main("aviatrix", aviatrix.Version, aviatrix.CompleteTerraformBridgeProvider())
}