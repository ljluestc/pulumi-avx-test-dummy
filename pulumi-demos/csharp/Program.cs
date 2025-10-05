using System;
using System.Collections.Generic;
using Pulumi;

class Program
{
    static int Main(string[] args)
    {
        return Deployment.Run(() => {
            // Simulate Aviatrix resources using Pulumi's built-in resources
            // This demonstrates the Pulumi program structure for Aviatrix
            
            // Create Aviatrix configuration
            var aviatrixConfig = new Dictionary<string, object>
            {
                ["provider"] = new Dictionary<string, object>
                {
                    ["controller_ip"] = "demo-controller.aviatrix.com",
                    ["username"] = "demo-user",
                    ["password"] = "demo-password"
                },
                ["resources"] = new Dictionary<string, object>
                {
                    ["account"] = new Dictionary<string, object>
                    {
                        ["name"] = "demo-aws-account",
                        ["cloud_type"] = 1,
                        ["aws_account_number"] = "123456789012"
                    },
                    ["gateway"] = new Dictionary<string, object>
                    {
                        ["name"] = "demo-gateway",
                        ["vpc_id"] = "vpc-12345678",
                        ["vpc_region"] = "us-west-1",
                        ["vpc_size"] = "t3.micro"
                    },
                    ["vpc"] = new Dictionary<string, object>
                    {
                        ["name"] = "demo-vpc",
                        ["region"] = "us-west-1",
                        ["cidr"] = "10.0.0.0/16"
                    },
                    ["spoke_gateway"] = new Dictionary<string, object>
                    {
                        ["name"] = "demo-spoke-gateway",
                        ["vpc_id"] = "vpc-87654321",
                        ["vpc_region"] = "us-east-1",
                        ["vpc_size"] = "t3.micro"
                    },
                    ["transit_gateway"] = new Dictionary<string, object>
                    {
                        ["name"] = "demo-transit-gateway",
                        ["vpc_id"] = "vpc-11223344",
                        ["vpc_region"] = "us-west-2",
                        ["vpc_size"] = "t3.medium"
                    }
                }
            };

            // Export outputs
            return new Dictionary<string, object?>
            {
                ["provider_controller_ip"] = "demo-controller.aviatrix.com",
                ["account_name"] = "demo-aws-account",
                ["gateway_name"] = "demo-gateway",
                ["vpc_name"] = "demo-vpc",
                ["spoke_gateway_name"] = "demo-spoke-gateway",
                ["transit_gateway_name"] = "demo-transit-gateway",
                ["config"] = aviatrixConfig
            };
        });
    }
}
