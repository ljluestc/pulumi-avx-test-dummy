package aviatrix

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// SimpleTerraformBridgeProvider returns a simple Pulumi provider
func SimpleTerraformBridgeProvider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		Name:        "aviatrix",
		DisplayName: "Aviatrix",
		Publisher:   "AviatrixSystems",
		Description: "A Pulumi package for creating and managing Aviatrix cloud resources.",
		Keywords: []string{
			"pulumi",
			"aviatrix",
			"category/cloud",
			"category/networking",
			"category/security",
		},
		License:    "Apache-2.0",
		Homepage:   "https://www.aviatrix.com",
		Repository: "https://github.com/pulumi/pulumi-aviatrix",
		GitHubOrg:  "AviatrixSystems",
		Config: map[string]*tfbridge.SchemaInfo{
			"controller_ip": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AVIATRIX_CONTROLLER_IP"},
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AVIATRIX_USERNAME"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"AVIATRIX_PASSWORD"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Core resources
			"aviatrix_account":     {Tok: tokens.NewTypeToken("aviatrix", "Account")},
			"aviatrix_gateway":     {Tok: tokens.NewTypeToken("aviatrix", "Gateway")},
			"aviatrix_spoke_gateway": {Tok: tokens.NewTypeToken("aviatrix", "SpokeGateway")},
			"aviatrix_transit_gateway": {Tok: tokens.NewTypeToken("aviatrix", "TransitGateway")},
			"aviatrix_vpc":         {Tok: tokens.NewTypeToken("aviatrix", "Vpc")},
			"aviatrix_firenet":     {Tok: tokens.NewTypeToken("aviatrix", "FireNet")},
			"aviatrix_firewall":    {Tok: tokens.NewTypeToken("aviatrix", "Firewall")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"aviatrix_account": {Tok: tokens.NewDataSourceToken("aviatrix", "getAccount")},
			"aviatrix_gateway": {Tok: tokens.NewDataSourceToken("aviatrix", "getGateway")},
			"aviatrix_vpc":     {Tok: tokens.NewDataSourceToken("aviatrix", "getVpc")},
		},
		Version: Version,
	}
	return prov
}
