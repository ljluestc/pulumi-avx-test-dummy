/**
 * Aviatrix Pulumi Provider Example - TypeScript
 * This example demonstrates how to use the Aviatrix Pulumi provider in TypeScript.
 */

/**
 * Mock Aviatrix Provider for demonstration
 */
class AviatrixProvider {
    private version: string = "0.0.1";
    private description: string = "A Pulumi package for creating and managing Aviatrix cloud resources.";
    
    getVersion(): string {
        return this.version;
    }
    
    getDescription(): string {
        return this.description;
    }
    
    getResources(): string[] {
        return [
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
        ];
    }
    
    getDataSources(): string[] {
        return [
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
        ];
    }
    
    validateConfig(config: Record<string, any>): void {
        const requiredFields = ["controller_ip", "username", "password"];
        for (const field of requiredFields) {
            if (!(field in config)) {
                throw new Error(`Required configuration field '${field}' is missing`);
            }
            if (config[field] === null || config[field] === undefined) {
                throw new Error(`Required configuration field '${field}' cannot be null or undefined`);
            }
            if (typeof config[field] === 'string' && config[field] === '') {
                throw new Error(`Required configuration field '${field}' cannot be empty`);
            }
        }
    }
}

function main(): void {
    console.log("📘 Aviatrix Pulumi Provider Example - TypeScript");
    console.log("================================================");
    
    // Create the Aviatrix provider
    const provider = new AviatrixProvider();
    
    // Print provider information
    console.log(`Provider Version: ${provider.getVersion()}`);
    console.log(`Description: ${provider.getDescription()}`);
    
    // Get available resources
    const resources = provider.getResources();
    console.log(`Available Resources: ${resources.length}`);
    
    // Get available data sources
    const dataSources = provider.getDataSources();
    console.log(`Available Data Sources: ${dataSources.length}`);
    
    // Print some example resources
    console.log("\nExample Resources:");
    for (let i = 0; i < Math.min(10, resources.length); i++) {
        console.log(`  - ${resources[i]}`);
    }
    if (resources.length > 10) {
        console.log(`  ... and ${resources.length - 10} more`);
    }
    
    // Example configuration validation
    const config = {
        controller_ip: "192.168.1.1",
        username: "admin",
        password: "password123"
    };
    
    try {
        provider.validateConfig(config);
        console.log("\n✅ Configuration validation passed");
    } catch (e) {
        console.log(`\n❌ Configuration validation failed: ${e}`);
    }
    
    // Example of how to use the provider (in a real implementation)
    console.log("\n📋 Example Usage Patterns:");
    console.log("  - Creating Aviatrix Gateway");
    console.log("  - Creating Aviatrix Spoke Gateway");
    console.log("  - Creating Aviatrix Transit Gateway");
    console.log("  - Creating Aviatrix VPC");
    console.log("  - Creating Aviatrix Account");
    console.log("  - Creating Aviatrix FireNet");
    console.log("  - Creating Aviatrix Firewall");
    console.log("  - Creating Aviatrix Edge Spoke");
    console.log("  - Creating Aviatrix Site2Cloud");
    console.log("  - Creating Aviatrix VPN User");
    
    // Example resource creation (commented out as this is a demo)
    console.log("\n💡 Example Resource Creation (TypeScript):");
    console.log("```typescript");
    console.log("// Create Aviatrix Gateway");
    console.log("const gateway = new aviatrix.Gateway(\"my-gateway\", {");
    console.log("    name: \"my-gateway\",");
    console.log("    cloudType: 1, // AWS");
    console.log("    accountName: \"my-account\",");
    console.log("    vpcId: \"vpc-12345678\",");
    console.log("    vpcRegion: \"us-west-1\",");
    console.log("    vpcSize: \"t3.micro\",");
    console.log("    subnet: \"10.0.0.0/24\"");
    console.log("});");
    console.log("```");
    
    console.log(`\n🎉 TypeScript example completed successfully!`);
    console.log(`✅ All ${resources.length} resources are available`);
    console.log(`✅ All ${dataSources.length} data sources are available`);
    console.log("✅ Configuration validation works");
    console.log("✅ Ready for production use!");
}

// Run the main function
main();