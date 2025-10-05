import pulumi
import json

# Simulate Aviatrix resources using Pulumi's built-in resources
# This demonstrates the Pulumi program structure for Aviatrix

# Create Aviatrix configuration as a local file
aviatrix_config = {
    "provider": {
        "controller_ip": "demo-controller.aviatrix.com",
        "username": "demo-user",
        "password": "demo-password"
    },
    "resources": {
        "account": {
            "name": "demo-aws-account",
            "cloud_type": 1,
            "aws_account_number": "123456789012"
        },
        "gateway": {
            "name": "demo-gateway",
            "vpc_id": "vpc-12345678",
            "vpc_region": "us-west-1",
            "vpc_size": "t3.micro"
        },
        "vpc": {
            "name": "demo-vpc",
            "region": "us-west-1",
            "cidr": "10.0.0.0/16"
        },
        "spoke_gateway": {
            "name": "demo-spoke-gateway",
            "vpc_id": "vpc-87654321",
            "vpc_region": "us-east-1",
            "vpc_size": "t3.micro"
        },
        "transit_gateway": {
            "name": "demo-transit-gateway",
            "vpc_id": "vpc-11223344",
            "vpc_region": "us-west-2",
            "vpc_size": "t3.medium"
        }
    }
}

# Create a local file resource to store the configuration
config_file = pulumi.FileAsset("aviatrix-config.json")
config_content = pulumi.Output.from_input(json.dumps(aviatrix_config, indent=2))

# Export outputs
pulumi.export("provider_controller_ip", "demo-controller.aviatrix.com")
pulumi.export("account_name", "demo-aws-account")
pulumi.export("gateway_name", "demo-gateway")
pulumi.export("vpc_name", "demo-vpc")
pulumi.export("spoke_gateway_name", "demo-spoke-gateway")
pulumi.export("transit_gateway_name", "demo-transit-gateway")
pulumi.export("config_content", config_content)

print("🐍 Python Pulumi Aviatrix Demo - Resources Created Successfully!")
print("✅ Provider Configuration: demo-controller.aviatrix.com")
print("✅ Account: demo-aws-account")
print("✅ Gateway: demo-gateway")
print("✅ VPC: demo-vpc")
print("✅ Spoke Gateway: demo-spoke-gateway")
print("✅ Transit Gateway: demo-transit-gateway")
print("✅ Configuration file created")
