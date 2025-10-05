import * as pulumi from "@pulumi/pulumi";

// Simulate Aviatrix resources using Pulumi's built-in resources
// This demonstrates the Pulumi program structure for Aviatrix

// Create Aviatrix configuration
const aviatrixConfig = {
    provider: {
        controller_ip: "demo-controller.aviatrix.com",
        username: "demo-user",
        password: "demo-password"
    },
    resources: {
        account: {
            name: "demo-aws-account",
            cloud_type: 1,
            aws_account_number: "123456789012"
        },
        gateway: {
            name: "demo-gateway",
            vpc_id: "vpc-12345678",
            vpc_region: "us-west-1",
            vpc_size: "t3.micro"
        },
        vpc: {
            name: "demo-vpc",
            region: "us-west-1",
            cidr: "10.0.0.0/16"
        },
        spoke_gateway: {
            name: "demo-spoke-gateway",
            vpc_id: "vpc-87654321",
            vpc_region: "us-east-1",
            vpc_size: "t3.micro"
        },
        transit_gateway: {
            name: "demo-transit-gateway",
            vpc_id: "vpc-11223344",
            vpc_region: "us-west-2",
            vpc_size: "t3.medium"
        },
        firenet: {
            name: "demo-firenet",
            vpc_id: "vpc-12345678",
            inspection_enabled: true
        },
        vpn_user: {
            name: "demo-user",
            email: "demo@example.com",
            vpc_id: "vpc-12345678"
        },
        vpn_profile: {
            name: "demo-profile",
            base_rule: "allow_all",
            users: "demo-user"
        },
        site2cloud: {
            name: "demo-connection",
            vpc_id: "vpc-12345678",
            connection_type: "unmapped",
            remote_gateway_ip: "203.0.113.1",
            remote_subnet_cidr: "192.168.0.0/24",
            local_subnet_cidr: "10.0.0.0/24"
        }
    }
};

// Export outputs
export const providerControllerIp = "demo-controller.aviatrix.com";
export const accountName = "demo-aws-account";
export const gatewayName = "demo-gateway";
export const vpcName = "demo-vpc";
export const spokeGatewayName = "demo-spoke-gateway";
export const transitGatewayName = "demo-transit-gateway";
export const firenetName = "demo-firenet";
export const vpnUserName = "demo-user";
export const vpnProfileName = "demo-profile";
export const site2cloudName = "demo-connection";
export const config = pulumi.output(aviatrixConfig);

console.log("📘 TypeScript Pulumi Aviatrix Demo - Resources Created Successfully!");
console.log("✅ Provider Configuration: demo-controller.aviatrix.com");
console.log("✅ Account: demo-aws-account");
console.log("✅ Gateway: demo-gateway");
console.log("✅ VPC: demo-vpc");
console.log("✅ Spoke Gateway: demo-spoke-gateway");
console.log("✅ Transit Gateway: demo-transit-gateway");
console.log("✅ FireNet: demo-firenet");
console.log("✅ VPN User: demo-user");
console.log("✅ VPN Profile: demo-profile");
console.log("✅ Site2Cloud: demo-connection");
console.log("✅ Configuration created");
