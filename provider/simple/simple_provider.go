package aviatrix

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SimpleAviatrixProvider is a minimal Pulumi provider for Aviatrix
type SimpleAviatrixProvider struct {
	version string
}

// NewSimpleAviatrixProvider creates a new instance of SimpleAviatrixProvider
func NewSimpleAviatrixProvider(version string) *SimpleAviatrixProvider {
	return &SimpleAviatrixProvider{
		version: version,
	}
}

// GetVersion returns the provider's version
func (p *SimpleAviatrixProvider) GetVersion() string {
	return p.version
}

// GetDescription returns the provider's description
func (p *SimpleAviatrixProvider) GetDescription() string {
	return "A Pulumi package for creating and managing Aviatrix cloud resources."
}

// GetResources returns a list of supported resources
func (p *SimpleAviatrixProvider) GetResources() []string {
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

// GetDataSources returns a list of supported data sources
func (p *SimpleAviatrixProvider) GetDataSources() []string {
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

// GetConfig returns the provider's configuration schema
func (p *SimpleAviatrixProvider) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"controller_ip": map[string]interface{}{
			"type":        "string",
			"description": "Controller IP address",
			"required":    true,
		},
		"username": map[string]interface{}{
			"type":        "string", 
			"description": "Username for the controller",
			"required":    true,
		},
		"password": map[string]interface{}{
			"type":        "string",
			"description": "Password for the controller", 
			"required":    true,
		},
		"skip_version_validation": map[string]interface{}{
			"type":        "boolean",
			"description": "Skip controller version validation",
			"optional":    true,
			"default":     false,
		},
		"verify_ssl_certificate": map[string]interface{}{
			"type":        "boolean",
			"description": "Verify SSL certificate",
			"optional":    true,
			"default":     false,
		},
		"path_to_ca_certificate": map[string]interface{}{
			"type":        "string",
			"description": "Path to CA certificate",
			"optional":    true,
		},
	}
}

// ValidateConfig validates the provider configuration
func (p *SimpleAviatrixProvider) ValidateConfig(config map[string]interface{}) error {
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

// Version constant for the provider
const Version = "0.0.1"

// Provider returns a simple Aviatrix provider
func Provider() *SimpleAviatrixProvider {
	return NewSimpleAviatrixProvider(Version)
}

// Mock resource and data source implementations for the simple provider
type MyResource struct {
	pulumi.ResourceState
	Name pulumi.StringOutput `pulumi:"name"`
}

func NewMyResource(ctx *pulumi.Context, name string, args *MyResourceArgs, opts ...pulumi.ResourceOption) (*MyResource, error) {
	var resource MyResource
	err := ctx.RegisterResource("my:module:MyResource", name, pulumi.Map{
		"name": args.Name,
	}, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type MyResourceArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}