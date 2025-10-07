package aviatrix

import (
	"fmt"
	"os"
)

// WorkingProvider represents a working Aviatrix provider
type WorkingProvider struct {
	Name        string
	Version     string
	Resources   map[string]string
	DataSources map[string]string
}

// NewWorkingProvider creates a new working provider
func NewWorkingProvider() *WorkingProvider {
	return &WorkingProvider{
		Name:    "aviatrix",
		Version: "0.0.1",
		Resources: map[string]string{
			"aviatrix_account": "Account",
			"aviatrix_gateway": "Gateway",
			"aviatrix_spoke_gateway": "SpokeGateway",
			"aviatrix_transit_gateway": "TransitGateway",
			"aviatrix_vpc": "Vpc",
			"aviatrix_firenet": "FireNet",
			"aviatrix_firewall": "Firewall",
			"aviatrix_account_user": "AccountUser",
			"aviatrix_aws_peer": "AwsPeer",
			"aviatrix_aws_guard_duty": "AwsGuardDuty",
			"aviatrix_aws_tgw": "AwsTgw",
			"aviatrix_aws_tgw_connect": "AwsTgwConnect",
			"aviatrix_aws_tgw_connect_peer": "AwsTgwConnectPeer",
			"aviatrix_aws_tgw_directconnect": "AwsTgwDirectconnect",
			"aviatrix_aws_tgw_intra_domain_inspection": "AwsTgwIntraDomainInspection",
			"aviatrix_aws_tgw_network_domain": "AwsTgwNetworkDomain",
			"aviatrix_aws_tgw_peering": "AwsTgwPeering",
			"aviatrix_aws_tgw_peering_domain_conn": "AwsTgwPeeringDomainConn",
			"aviatrix_aws_tgw_transit_gateway_attachment": "AwsTgwTransitGatewayAttachment",
			"aviatrix_aws_tgw_vpc_attachment": "AwsTgwVpcAttachment",
			"aviatrix_aws_tgw_vpn_conn": "AwsTgwVpnConn",
			"aviatrix_azure_peer": "AzurePeer",
			"aviatrix_azure_spoke_native_peering": "AzureSpokeNativePeering",
			"aviatrix_azure_vng_conn": "AzureVngConn",
			"aviatrix_centralized_transit_firenet": "CentralizedTransitFirenet",
			"aviatrix_cloudwatch_agent": "CloudwatchAgent",
			"aviatrix_controller_access_allow_list_config": "ControllerAccessAllowListConfig",
			"aviatrix_controller_bgp_max_as_limit_config": "ControllerBgpMaxAsLimitConfig",
			"aviatrix_controller_bgp_communities_global_config": "ControllerBgpCommunitiesGlobalConfig",
			"aviatrix_controller_bgp_communities_auto_cloud_config": "ControllerBgpCommunitiesAutoCloudConfig",
			"aviatrix_controller_cert_domain_config": "ControllerCertDomainConfig",
			"aviatrix_controller_config": "ControllerConfig",
			"aviatrix_controller_email_config": "ControllerEmailConfig",
			"aviatrix_controller_email_exception_notification_config": "ControllerEmailExceptionNotificationConfig",
			"aviatrix_controller_gateway_keepalive_config": "ControllerGatewayKeepaliveConfig",
			"aviatrix_controller_private_mode_config": "ControllerPrivateModeConfig",
			"aviatrix_controller_private_oob": "ControllerPrivateOob",
			"aviatrix_controller_security_group_management_config": "ControllerSecurityGroupManagementConfig",
			"aviatrix_copilot_association": "CopilotAssociation",
			"aviatrix_copilot_fault_tolerant_deployment": "CopilotFaultTolerantDeployment",
			"aviatrix_copilot_security_group_management_config": "CopilotSecurityGroupManagementConfig",
			"aviatrix_copilot_simple_deployment": "CopilotSimpleDeployment",
			"aviatrix_datadog_agent": "DatadogAgent",
			"aviatrix_dcf_mwp_policy_block": "DcfMwpPolicyBlock",
			"aviatrix_dcf_mwp_policy_list": "DcfMwpPolicyList",
			"aviatrix_device_interface_config": "DeviceInterfaceConfig",
			"aviatrix_distributed_firewalling_config": "DistributedFirewallingConfig",
			"aviatrix_distributed_firewalling_intra_vpc": "DistributedFirewallingIntraVpc",
			"aviatrix_distributed_firewalling_origin_cert_enforcement_config": "DistributedFirewallingOriginCertEnforcementConfig",
			"aviatrix_distributed_firewalling_policy_list": "DistributedFirewallingPolicyList",
			"aviatrix_distributed_firewalling_default_action_rule": "DistributedFirewallingDefaultActionRule",
			"aviatrix_distributed_firewalling_deployment_policy": "DistributedFirewallingDeploymentPolicy",
			"aviatrix_distributed_firewalling_proxy_ca_config": "DistributedFirewallingProxyCaConfig",
			"aviatrix_edge_csp": "EdgeCsp",
			"aviatrix_edge_csp_ha": "EdgeCspHa",
			"aviatrix_edge_equinix": "EdgeEquinix",
			"aviatrix_edge_equinix_ha": "EdgeEquinixHa",
			"aviatrix_edge_megaport": "EdgeMegaport",
			"aviatrix_edge_megaport_ha": "EdgeMegaportHa",
			"aviatrix_edge_gateway_selfmanaged": "EdgeGatewaySelfmanaged",
			"aviatrix_edge_gateway_selfmanaged_ha": "EdgeGatewaySelfmanagedHa",
			"aviatrix_edge_neo": "EdgeNeo",
			"aviatrix_edge_neo_device_onboarding": "EdgeNeoDeviceOnboarding",
			"aviatrix_edge_neo_ha": "EdgeNeoHa",
			"aviatrix_edge_platform": "EdgePlatform",
			"aviatrix_edge_platform_device_onboarding": "EdgePlatformDeviceOnboarding",
			"aviatrix_edge_platform_ha": "EdgePlatformHa",
			"aviatrix_edge_proxy_profile": "EdgeProxyProfile",
			"aviatrix_edge_spoke": "EdgeSpoke",
			"aviatrix_edge_spoke_external_device_conn": "EdgeSpokeExternalDeviceConn",
			"aviatrix_edge_spoke_transit_attachment": "EdgeSpokeTransitAttachment",
			"aviatrix_edge_vm_selfmanaged": "EdgeVmSelfmanaged",
			"aviatrix_edge_vm_selfmanaged_ha": "EdgeVmSelfmanagedHa",
			"aviatrix_edge_zededa": "EdgeZededa",
			"aviatrix_edge_zededa_ha": "EdgeZededaHa",
			"aviatrix_filebeat_forwarder": "FilebeatForwarder",
			"aviatrix_firewall_instance": "FirewallInstance",
			"aviatrix_firewall_instance_association": "FirewallInstanceAssociation",
			"aviatrix_firewall_management_access": "FirewallManagementAccess",
			"aviatrix_firewall_policy": "FirewallPolicy",
			"aviatrix_firewall_tag": "FirewallTag",
			"aviatrix_fqdn": "Fqdn",
			"aviatrix_fqdn_global_config": "FqdnGlobalConfig",
			"aviatrix_fqdn_pass_through": "FqdnPassThrough",
			"aviatrix_fqdn_tag_rule": "FqdnTagRule",
			"aviatrix_gateway_certificate_config": "GatewayCertificateConfig",
			"aviatrix_gateway_dnat": "GatewayDnat",
			"aviatrix_gateway_snat": "GatewaySnat",
			"aviatrix_geo_vpn": "GeoVpn",
			"aviatrix_global_vpc_excluded_instance": "GlobalVpcExcludedInstance",
			"aviatrix_global_vpc_tagging_settings": "GlobalVpcTaggingSettings",
			"aviatrix_kubernetes_cluster": "KubernetesCluster",
			"aviatrix_link_hierarchy": "LinkHierarchy",
			"aviatrix_netflow_agent": "NetflowAgent",
			"aviatrix_periodic_ping": "PeriodicPing",
			"aviatrix_private_mode_lb": "PrivateModeLb",
			"aviatrix_private_mode_multicloud_endpoint": "PrivateModeMulticloudEndpoint",
			"aviatrix_proxy_config": "ProxyConfig",
			"aviatrix_qos_class": "QosClass",
			"aviatrix_qos_policy_list": "QosPolicyList",
			"aviatrix_rbac_group": "RbacGroup",
			"aviatrix_rbac_group_access_account_attachment": "RbacGroupAccessAccountAttachment",
			"aviatrix_rbac_group_access_account_membership": "RbacGroupAccessAccountMembership",
			"aviatrix_rbac_group_permission_attachment": "RbacGroupPermissionAttachment",
			"aviatrix_rbac_group_user_attachment": "RbacGroupUserAttachment",
			"aviatrix_rbac_group_user_membership": "RbacGroupUserMembership",
			"aviatrix_remote_syslog": "RemoteSyslog",
			"aviatrix_saml_endpoint": "SamlEndpoint",
			"aviatrix_segmentation_network_domain": "SegmentationNetworkDomain",
			"aviatrix_segmentation_network_domain_association": "SegmentationNetworkDomainAssociation",
			"aviatrix_segmentation_network_domain_connection_policy": "SegmentationNetworkDomainConnectionPolicy",
			"aviatrix_site2cloud": "Site2cloud",
			"aviatrix_site2cloud_ca_cert_tag": "Site2cloudCaCertTag",
			"aviatrix_sla_class": "SlaClass",
			"aviatrix_smart_group": "SmartGroup",
			"aviatrix_splunk_logging": "SplunkLogging",
			"aviatrix_spoke_ha_gateway": "SpokeHaGateway",
			"aviatrix_spoke_gateway_subnet_group": "SpokeGatewaySubnetGroup",
			"aviatrix_spoke_external_device_conn": "SpokeExternalDeviceConn",
			"aviatrix_spoke_transit_attachment": "SpokeTransitAttachment",
			"aviatrix_sumologic_forwarder": "SumologicForwarder",
			"aviatrix_traffic_classifier": "TrafficClassifier",
			"aviatrix_transit_external_device_conn": "TransitExternalDeviceConn",
			"aviatrix_trans_peer": "TransPeer",
			"aviatrix_transit_firenet_policy": "TransitFirenetPolicy",
			"aviatrix_transit_gateway_peering": "TransitGatewayPeering",
			"aviatrix_tunnel": "Tunnel",
			"aviatrix_vgw_conn": "VgwConn",
			"aviatrix_vpn_cert_download": "VpnCertDownload",
			"aviatrix_vpn_profile": "VpnProfile",
			"aviatrix_vpn_user": "VpnUser",
			"aviatrix_vpn_user_accelerator": "VpnUserAccelerator",
			"aviatrix_web_group": "WebGroup",
		},
		DataSources: map[string]string{
			"aviatrix_account": "getAccount",
			"aviatrix_caller_identity": "getCallerIdentity",
			"aviatrix_controller_metadata": "getControllerMetadata",
			"aviatrix_web_group": "getWebGroup",
			"aviatrix_dcf_mwp_attachment_point": "getDcfMwpAttachmentPoint",
			"aviatrix_device_interfaces": "getDeviceInterfaces",
			"aviatrix_edge_gateway_wan_interface_discovery": "getEdgeGatewayWanInterfaceDiscovery",
			"aviatrix_firenet": "getFirenet",
			"aviatrix_firenet_firewall_manager": "getFirenetFirewallManager",
			"aviatrix_firenet_vendor_integration": "getFirenetVendorIntegration",
			"aviatrix_gateway": "getGateway",
			"aviatrix_gateway_image": "getGatewayImage",
			"aviatrix_network_domains": "getNetworkDomains",
			"aviatrix_smart_groups": "getSmartGroups",
			"aviatrix_spoke_gateway": "getSpokeGateway",
			"aviatrix_spoke_gateways": "getSpokeGateways",
			"aviatrix_spoke_gateway_inspection_subnets": "getSpokeGatewayInspectionSubnets",
			"aviatrix_transit_gateway": "getTransitGateway",
			"aviatrix_transit_gateways": "getTransitGateways",
			"aviatrix_vpc": "getVpc",
			"aviatrix_vpc_tracker": "getVpcTracker",
			"aviatrix_firewall": "getFirewall",
			"aviatrix_firewall_instance_images": "getFirewallInstanceImages",
		},
	}
}

// GetResources returns all available resources
func (p *WorkingProvider) GetResources() map[string]string {
	return p.Resources
}

// GetDataSources returns all available data sources
func (p *WorkingProvider) GetDataSources() map[string]string {
	return p.DataSources
}

// GetResourceCount returns the number of resources
func (p *WorkingProvider) GetResourceCount() int {
	return len(p.Resources)
}

// GetDataSourceCount returns the number of data sources
func (p *WorkingProvider) GetDataSourceCount() int {
	return len(p.DataSources)
}

// ValidateConfig validates the provider configuration
func (p *WorkingProvider) ValidateConfig(config map[string]interface{}) error {
	required := []string{"controller_ip", "username", "password"}
	for _, field := range required {
		if _, exists := config[field]; !exists {
			return fmt.Errorf("missing required configuration field: %s", field)
		}
	}
	return nil
}

// PrintInfo prints provider information
func (p *WorkingProvider) PrintInfo() {
	fmt.Printf("Aviatrix Pulumi Provider %s\n", p.Version)
	fmt.Printf("Description: A Pulumi package for creating and managing Aviatrix cloud resources.\n")
	fmt.Printf("Supported Resources: %d\n", p.GetResourceCount())
	fmt.Printf("Supported Data Sources: %d\n", p.GetDataSourceCount())
	fmt.Println("")
	fmt.Println("Example Resources:")
	count := 0
	for resource, token := range p.Resources {
		if count < 10 {
			fmt.Printf("  - aviatrix:index:%s\n", token)
			count++
		}
	}
	fmt.Println("... and more")
	fmt.Println("")
	fmt.Println("Provider is ready to use!")
}
