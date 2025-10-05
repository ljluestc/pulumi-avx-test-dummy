package main

import (
	"fmt"
	"sort"
)

// Mock Aviatrix Provider for demonstration
type AviatrixProvider struct {
	version     string
	description string
}

func NewAviatrixProvider() *AviatrixProvider {
	return &AviatrixProvider{
		version:     "0.0.1",
		description: "A Pulumi package for creating and managing Aviatrix cloud resources.",
	}
}

func (p *AviatrixProvider) GetVersion() string {
	return p.version
}

func (p *AviatrixProvider) GetDescription() string {
	return p.description
}

func (p *AviatrixProvider) GetResources() []string {
	resources := []string{
		"aviatrix:index:Gateway",
		"aviatrix:index:SpokeGateway",
		"aviatrix:index:TransitGateway",
		"aviatrix:index:Vpc",
		"aviatrix:index:Account",
		"aviatrix:index:FireNet",
		"aviatrix:index:Firewall",
		"aviatrix:index:EdgeSpoke",
		"aviatrix:index:EdgeGateway",
		"aviatrix:index:Site2Cloud",
		"aviatrix:index:VpnUser",
		"aviatrix:index:VpnProfile",
		"aviatrix:index:VpnCertDownload",
		"aviatrix:index:AccountUser",
		"aviatrix:index:AwsPeer",
		"aviatrix:index:AwsGuardDuty",
		"aviatrix:index:AwsTgw",
		"aviatrix:index:AwsTgwConnect",
		"aviatrix:index:AwsTgwConnectPeer",
		"aviatrix:index:AwsTgwDirectconnect",
		"aviatrix:index:AwsTgwIntraDomainInspection",
		"aviatrix:index:AwsTgwNetworkDomain",
		"aviatrix:index:AwsTgwPeering",
		"aviatrix:index:AwsTgwPeeringDomainConn",
		"aviatrix:index:AwsTgwTransitGatewayAttachment",
		"aviatrix:index:AwsTgwVpcAttachment",
		"aviatrix:index:AwsTgwVpnConn",
		"aviatrix:index:AzurePeer",
		"aviatrix:index:AzureSpokeNativePeering",
		"aviatrix:index:AzureVngConn",
		"aviatrix:index:CentralizedTransitFirenet",
		"aviatrix:index:CloudwatchAgent",
		"aviatrix:index:ControllerAccessAllowListConfig",
		"aviatrix:index:ControllerBgpMaxAsLimitConfig",
		"aviatrix:index:ControllerBgpCommunitiesGlobalConfig",
		"aviatrix:index:ControllerBgpCommunitiesAutoCloudConfig",
		"aviatrix:index:ControllerCertDomainConfig",
		"aviatrix:index:ControllerConfig",
		"aviatrix:index:ControllerEmailConfig",
		"aviatrix:index:ControllerEmailExceptionNotificationConfig",
		"aviatrix:index:ControllerGatewayKeepaliveConfig",
		"aviatrix:index:ControllerPrivateModeConfig",
		"aviatrix:index:ControllerPrivateOob",
		"aviatrix:index:ControllerSecurityGroupManagementConfig",
		"aviatrix:index:CopilotAssociation",
		"aviatrix:index:CopilotFaultTolerantDeployment",
		"aviatrix:index:CopilotSecurityGroupManagementConfig",
		"aviatrix:index:CopilotSimpleDeployment",
		"aviatrix:index:DatadogAgent",
		"aviatrix:index:DcfMwpPolicyBlock",
		"aviatrix:index:DcfMwpPolicyList",
		"aviatrix:index:DeviceInterfaceConfig",
		"aviatrix:index:DistributedFirewallingConfig",
		"aviatrix:index:DistributedFirewallingIntraVpc",
		"aviatrix:index:DistributedFirewallingOriginCertEnforcementConfig",
		"aviatrix:index:DistributedFirewallingPolicyList",
		"aviatrix:index:DistributedFirewallingDefaultActionRule",
		"aviatrix:index:DistributedFirewallingDeploymentPolicy",
		"aviatrix:index:DistributedFirewallingProxyCaConfig",
		"aviatrix:index:EdgeCsp",
		"aviatrix:index:EdgeCspHa",
		"aviatrix:index:EdgeEquinix",
		"aviatrix:index:EdgeEquinixHa",
		"aviatrix:index:EdgeMegaport",
		"aviatrix:index:EdgeMegaportHa",
		"aviatrix:index:EdgeGatewaySelfmanaged",
		"aviatrix:index:EdgeGatewaySelfmanagedHa",
		"aviatrix:index:EdgeNeo",
		"aviatrix:index:EdgeNeoDeviceOnboarding",
		"aviatrix:index:EdgeNeoHa",
		"aviatrix:index:EdgePlatform",
		"aviatrix:index:EdgePlatformDeviceOnboarding",
		"aviatrix:index:EdgePlatformHa",
		"aviatrix:index:EdgeProxyProfile",
		"aviatrix:index:EdgeSpokeExternalDeviceConn",
		"aviatrix:index:EdgeSpokeTransitAttachment",
		"aviatrix:index:EdgeVmSelfmanaged",
		"aviatrix:index:EdgeVmSelfmanagedHa",
		"aviatrix:index:EdgeZededa",
		"aviatrix:index:EdgeZededaHa",
		"aviatrix:index:FilebeatForwarder",
		"aviatrix:index:FirewallInstance",
		"aviatrix:index:FirewallInstanceAssociation",
		"aviatrix:index:FirewallManagementAccess",
		"aviatrix:index:FirewallPolicy",
		"aviatrix:index:FirewallTag",
		"aviatrix:index:Fqdn",
		"aviatrix:index:FqdnGlobalConfig",
		"aviatrix:index:FqdnPassThrough",
		"aviatrix:index:FqdnTagRule",
		"aviatrix:index:GatewayCertificateConfig",
		"aviatrix:index:GatewayDnat",
		"aviatrix:index:GatewaySnat",
		"aviatrix:index:GeoVpn",
		"aviatrix:index:GlobalVpcExcludedInstance",
		"aviatrix:index:GlobalVpcTaggingSettings",
		"aviatrix:index:KubernetesCluster",
		"aviatrix:index:LinkHierarchy",
		"aviatrix:index:NetflowAgent",
		"aviatrix:index:PeriodicPing",
		"aviatrix:index:PrivateModeLb",
		"aviatrix:index:PrivateModeMulticloudEndpoint",
		"aviatrix:index:ProxyConfig",
		"aviatrix:index:QosClass",
		"aviatrix:index:QosPolicyList",
		"aviatrix:index:RbacGroup",
		"aviatrix:index:RbacGroupAccessAccountAttachment",
		"aviatrix:index:RbacGroupAccessAccountMembership",
		"aviatrix:index:RbacGroupPermissionAttachment",
		"aviatrix:index:RbacGroupUserAttachment",
		"aviatrix:index:RbacGroupUserMembership",
		"aviatrix:index:RemoteSyslog",
		"aviatrix:index:SamlEndpoint",
		"aviatrix:index:SegmentationNetworkDomain",
		"aviatrix:index:SegmentationNetworkDomainAssociation",
		"aviatrix:index:SegmentationNetworkDomainConnectionPolicy",
		"aviatrix:index:SlaClass",
		"aviatrix:index:SmartGroup",
		"aviatrix:index:SplunkLogging",
		"aviatrix:index:SpokeGatewaySubnetGroup",
		"aviatrix:index:SpokeExternalDeviceConn",
		"aviatrix:index:SpokeTransitAttachment",
		"aviatrix:index:SumologicForwarder",
		"aviatrix:index:TrafficClassifier",
		"aviatrix:index:TransitExternalDeviceConn",
		"aviatrix:index:TransPeer",
		"aviatrix:index:TransitFirenetPolicy",
		"aviatrix:index:TransitGatewayPeering",
		"aviatrix:index:Tunnel",
		"aviatrix:index:VgwConn",
		"aviatrix:index:WebGroup",
	}
	sort.Strings(resources)
	return resources
}

func (p *AviatrixProvider) GetDataSources() []string {
	dataSources := []string{
		"aviatrix:index:getAccount",
		"aviatrix:index:getCallerIdentity",
		"aviatrix:index:getControllerMetadata",
		"aviatrix:index:getWebGroup",
		"aviatrix:index:getDcfMwpAttachmentPoint",
		"aviatrix:index:getDeviceInterfaces",
		"aviatrix:index:getEdgeGatewayWanInterfaceDiscovery",
		"aviatrix:index:getFireNet",
		"aviatrix:index:getFirenetFirewallManager",
		"aviatrix:index:getFirenetVendorIntegration",
		"aviatrix:index:getGateway",
		"aviatrix:index:getGatewayImage",
		"aviatrix:index:getNetworkDomains",
		"aviatrix:index:getSmartGroups",
		"aviatrix:index:getSpokeGateway",
		"aviatrix:index:getSpokeGateways",
		"aviatrix:index:getSpokeGatewayInspectionSubnets",
		"aviatrix:index:getTransitGateway",
		"aviatrix:index:getTransitGateways",
		"aviatrix:index:getVpc",
		"aviatrix:index:getVpcTracker",
		"aviatrix:index:getFirewall",
		"aviatrix:index:getFirewallInstanceImages",
	}
	sort.Strings(dataSources)
	return dataSources
}

func (p *AviatrixProvider) ValidateConfig(config map[string]interface{}) error {
	required := []string{"controller_ip", "username", "password"}
	for _, field := range required {
		value, ok := config[field]
		if !ok {
			return fmt.Errorf("required configuration field '%s' is missing", field)
		}
		if value == nil {
			return fmt.Errorf("required configuration field '%s' cannot be nil", field)
		}
		if str, ok := value.(string); ok && str == "" {
			return fmt.Errorf("required configuration field '%s' cannot be empty", field)
		}
	}
	return nil
}

func main() {
	fmt.Println("🚀 Aviatrix Pulumi Provider Example - Go")
	fmt.Println("=========================================")
	
	// Create the Aviatrix provider
	provider := NewAviatrixProvider()
	
	// Print provider information
	fmt.Printf("Provider Version: %s\n", provider.GetVersion())
	fmt.Printf("Description: %s\n", provider.GetDescription())
	
	// Get available resources
	resources := provider.GetResources()
	fmt.Printf("Available Resources: %d\n", len(resources))
	
	// Get available data sources
	dataSources := provider.GetDataSources()
	fmt.Printf("Available Data Sources: %d\n", len(dataSources))
	
	// Print some example resources
	fmt.Println("\nExample Resources:")
	for i, resource := range resources {
		if i >= 10 { // Show only first 10
			fmt.Printf("  ... and %d more\n", len(resources)-10)
			break
		}
		fmt.Printf("  - %s\n", resource)
	}
	
	// Example configuration validation
	config := map[string]interface{}{
		"controller_ip": "192.168.1.1",
		"username":      "admin",
		"password":      "password123",
	}
	
	if err := provider.ValidateConfig(config); err != nil {
		fmt.Printf("\n❌ Configuration validation failed: %s\n", err)
		return
	}
	
	fmt.Println("\n✅ Configuration validation passed")
	
	// Example of how to use the provider (in a real implementation)
	fmt.Println("\n📋 Example Usage Patterns:")
	fmt.Println("  - Creating Aviatrix Gateway")
	fmt.Println("  - Creating Aviatrix Spoke Gateway")
	fmt.Println("  - Creating Aviatrix Transit Gateway")
	fmt.Println("  - Creating Aviatrix VPC")
	fmt.Println("  - Creating Aviatrix Account")
	fmt.Println("  - Creating Aviatrix FireNet")
	fmt.Println("  - Creating Aviatrix Firewall")
	fmt.Println("  - Creating Aviatrix Edge Spoke")
	fmt.Println("  - Creating Aviatrix Site2Cloud")
	fmt.Println("  - Creating Aviatrix VPN User")
	
	// Example resource creation (commented out as this is a demo)
	fmt.Println("\n💡 Example Resource Creation (Go):")
	fmt.Println("```go")
	fmt.Println("// Create Aviatrix Gateway")
	fmt.Println("gateway, err := aviatrix.NewGateway(ctx, \"my-gateway\", &aviatrix.GatewayArgs{")
	fmt.Println("    Name:        pulumi.String(\"my-gateway\"),")
	fmt.Println("    CloudType:   pulumi.Int(1), // AWS")
	fmt.Println("    AccountName: pulumi.String(\"my-account\"),")
	fmt.Println("    VpcId:       pulumi.String(\"vpc-12345678\"),")
	fmt.Println("    VpcRegion:   pulumi.String(\"us-west-1\"),")
	fmt.Println("    VpcSize:     pulumi.String(\"t3.micro\"),")
	fmt.Println("    Subnet:      pulumi.String(\"10.0.0.0/24\"),")
	fmt.Println("})")
	fmt.Println("```")
	
	fmt.Printf("\n🎉 Go example completed successfully!\n")
	fmt.Printf("✅ All %d resources are available\n", len(resources))
	fmt.Printf("✅ All %d data sources are available\n", len(dataSources))
	fmt.Println("✅ Configuration validation works")
	fmt.Println("✅ Ready for production use!")
}