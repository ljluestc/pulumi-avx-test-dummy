package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	aviatrix "github.com/pulumi/pulumi-aviatrix/provider"
)

func main() {
	tfbridge.Main("aviatrix", aviatrix.Version, aviatrix.TerraformBridgeProvider())
}