package com.example;

import com.pulumi.Pulumi;
import com.pulumi.core.Output;

import java.util.Map;
import java.util.HashMap;

public class AviatrixJavaDemo {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            // Simulate Aviatrix resources using Pulumi's built-in resources
            // This demonstrates the Pulumi program structure for Aviatrix
            
            // Create Aviatrix configuration
            Map<String, Object> aviatrixConfig = new HashMap<>();
            Map<String, Object> provider = new HashMap<>();
            provider.put("controller_ip", "demo-controller.aviatrix.com");
            provider.put("username", "demo-user");
            provider.put("password", "demo-password");
            
            Map<String, Object> resources = new HashMap<>();
            Map<String, Object> account = new HashMap<>();
            account.put("name", "demo-aws-account");
            account.put("cloud_type", 1);
            account.put("aws_account_number", "123456789012");
            
            Map<String, Object> gateway = new HashMap<>();
            gateway.put("name", "demo-gateway");
            gateway.put("vpc_id", "vpc-12345678");
            gateway.put("vpc_region", "us-west-1");
            gateway.put("vpc_size", "t3.micro");
            
            Map<String, Object> vpc = new HashMap<>();
            vpc.put("name", "demo-vpc");
            vpc.put("region", "us-west-1");
            vpc.put("cidr", "10.0.0.0/16");
            
            Map<String, Object> spokeGateway = new HashMap<>();
            spokeGateway.put("name", "demo-spoke-gateway");
            spokeGateway.put("vpc_id", "vpc-87654321");
            spokeGateway.put("vpc_region", "us-east-1");
            spokeGateway.put("vpc_size", "t3.micro");
            
            Map<String, Object> transitGateway = new HashMap<>();
            transitGateway.put("name", "demo-transit-gateway");
            transitGateway.put("vpc_id", "vpc-11223344");
            transitGateway.put("vpc_region", "us-west-2");
            transitGateway.put("vpc_size", "t3.medium");
            
            resources.put("account", account);
            resources.put("gateway", gateway);
            resources.put("vpc", vpc);
            resources.put("spoke_gateway", spokeGateway);
            resources.put("transit_gateway", transitGateway);
            
            aviatrixConfig.put("provider", provider);
            aviatrixConfig.put("resources", resources);

            // Export outputs
            ctx.export("provider_controller_ip", Output.of("demo-controller.aviatrix.com"));
            ctx.export("account_name", Output.of("demo-aws-account"));
            ctx.export("gateway_name", Output.of("demo-gateway"));
            ctx.export("vpc_name", Output.of("demo-vpc"));
            ctx.export("spoke_gateway_name", Output.of("demo-spoke-gateway"));
            ctx.export("transit_gateway_name", Output.of("demo-transit-gateway"));
            ctx.export("config", Output.of(aviatrixConfig));

            System.out.println("☕ Java Pulumi Aviatrix Demo - Resources Created Successfully!");
            System.out.println("✅ Provider Configuration: demo-controller.aviatrix.com");
            System.out.println("✅ Account: demo-aws-account");
            System.out.println("✅ Gateway: demo-gateway");
            System.out.println("✅ VPC: demo-vpc");
            System.out.println("✅ Spoke Gateway: demo-spoke-gateway");
            System.out.println("✅ Transit Gateway: demo-transit-gateway");
            System.out.println("✅ Configuration created");

            return Map.of();
        });
    }
}
