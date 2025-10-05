package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	aviatrix "github.com/pulumi/pulumi-aviatrix/provider"
)

func main() {
	tfgen.Main("aviatrix", aviatrix.Version, aviatrix.TerraformBridgeProvider())
}