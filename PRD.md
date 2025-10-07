# Product Requirements Document: Pulumi Aviatrix Provider

## Executive Summary

This document outlines the requirements for the Pulumi Aviatrix Provider, a Terraform bridge-based provider that enables users to manage Aviatrix cloud networking resources using Pulumi's infrastructure-as-code platform across multiple programming languages (Go, Python, TypeScript/JavaScript, C#, Java, and YAML).

## Project Overview

### Purpose
Create a fully-featured Pulumi provider that bridges the Terraform Aviatrix provider, enabling Pulumi users to programmatically manage Aviatrix Controller configurations, cloud accounts, gateways, networking policies, and security features.

### Goals
1. Provide 100% feature parity with the Terraform Aviatrix provider
2. Support all major Pulumi-supported languages (Go, Python, TypeScript, C#, Java)
3. Implement proper SDK generation and distribution
4. Enable type-safe resource management with proper IDE support
5. Maintain compatibility with Pulumi's ecosystem and best practices

### Non-Goals
- Direct REST API implementation (using Terraform bridge instead)
- Custom resource implementations beyond what Terraform provider offers
- Controller deployment or installation automation

## Architecture

### High-Level Architecture
```
┌─────────────────────────────────────────────────────────────┐
│                    Pulumi Programs                          │
│  (Go, Python, TypeScript, C#, Java, YAML)                  │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│              Language-Specific SDKs                         │
│  (Generated from Schema)                                    │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│         Pulumi Resource Provider Binary                     │
│      (pulumi-resource-aviatrix)                            │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│          Terraform Bridge (tfbridge)                        │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│     Terraform Aviatrix Provider                             │
│  (github.com/AviatrixSystems/terraform-provider-aviatrix)  │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│           Aviatrix Controller API                           │
└─────────────────────────────────────────────────────────────┘
```

### Components

#### 1. Provider Bridge Core
- **Location**: `provider/terraform_bridge_provider.go`
- **Purpose**: Bridge configuration mapping Terraform resources to Pulumi tokens
- **Key Elements**:
  - Provider metadata (name, version, publisher, description)
  - Resource mappings (133 resources)
  - Data source mappings (24 data sources)
  - Configuration schema with environment variable support

#### 2. Main Entry Point
- **Location**: `provider/cmd/pulumi-resource-aviatrix/main.go`
- **Purpose**: Provider binary main function
- **Functionality**: Initialize tfbridge.Main() with provider configuration

#### 3. Schema Generator
- **Location**: `provider/cmd/schema-generator/main.go`
- **Purpose**: Generate Pulumi schema JSON from bridge configuration
- **Output**: `schema.json` file for SDK generation

#### 4. Version Management
- **Location**: `provider/pkg/version/version.go`
- **Purpose**: Centralized version constant
- **Current Version**: 0.0.1

## Resource Coverage

### Account Management (2 resources)
- `aviatrix_account` - Cloud account onboarding
- `aviatrix_account_user` - User account management

### AWS Transit Gateway (13 resources)
- `aviatrix_aws_tgw` - AWS Transit Gateway
- `aviatrix_aws_tgw_connect` - TGW Connect
- `aviatrix_aws_tgw_connect_peer` - TGW Connect Peer
- `aviatrix_aws_tgw_directconnect` - TGW Direct Connect
- `aviatrix_aws_tgw_intra_domain_inspection` - Intra-domain inspection
- `aviatrix_aws_tgw_network_domain` - Network domain
- `aviatrix_aws_tgw_peering` - TGW peering
- `aviatrix_aws_tgw_peering_domain_conn` - Domain connection
- `aviatrix_aws_tgw_transit_gateway_attachment` - Transit attachment
- `aviatrix_aws_tgw_vpc_attachment` - VPC attachment
- `aviatrix_aws_tgw_vpn_conn` - VPN connection

### Gateway Management (8 resources)
- `aviatrix_gateway` - Generic gateway
- `aviatrix_spoke_gateway` - Spoke gateway
- `aviatrix_spoke_ha_gateway` - Spoke HA gateway
- `aviatrix_transit_gateway` - Transit gateway
- `aviatrix_gateway_certificate_config` - Gateway certificates
- `aviatrix_gateway_dnat` - DNAT configuration
- `aviatrix_gateway_snat` - SNAT configuration

### Edge Platform (24 resources)
- Edge CSP (Cloud Service Provider)
- Edge Equinix
- Edge Megaport
- Edge Self-managed
- Edge Neo
- Edge Platform
- Edge VM
- Edge Zededa
- Edge Spoke configurations and attachments

### Firewall & Security (15 resources)
- FireNet and firewall instances
- Firewall policies and tags
- Distributed firewalling (7 resources)
- Firewall management access

### Controller Configuration (12 resources)
- Access allow list
- BGP configuration (max AS limit, communities)
- Certificate domain configuration
- Email configuration and notifications
- Gateway keepalive
- Private mode configuration
- Security group management

### Network Segmentation (3 resources)
- Segmentation network domains
- Domain associations
- Connection policies

### VPN & Site2Cloud (7 resources)
- VPN user management
- VPN profiles
- VPN certificate download
- VPN user accelerator
- Site2Cloud connections
- Site2Cloud CA certificates
- Geo VPN

### Peering (4 resources)
- AWS peering
- Azure peering
- Azure spoke native peering
- Transit peering

### Monitoring & Logging (8 resources)
- CloudWatch agent
- Datadog agent
- Filebeat forwarder
- Netflow agent
- Remote syslog
- Splunk logging
- Sumologic forwarder

### Advanced Features (20+ resources)
- FQDN filtering and global config
- QoS classes and policies
- RBAC groups and permissions
- Smart groups
- Traffic classifier
- Kubernetes cluster integration
- Copilot deployment and management
- Link hierarchy
- SLA classes
- Web groups

### Data Sources (24 total)
- Account information
- Caller identity
- Controller metadata
- Device interfaces
- Edge gateway WAN interface discovery
- FireNet details
- Gateway information and images
- Network domains
- Smart groups
- Spoke gateways and inspection subnets
- Transit gateways
- VPC information and tracking
- Firewall details and images

## Configuration Requirements

### Provider Configuration
```yaml
Required:
  - controller_ip: Aviatrix Controller IP/hostname
  - username: Controller admin username
  - password: Controller admin password

Optional:
  - skip_version_validation: Skip version compatibility checks
  - verify_ssl_certificate: SSL certificate verification
  - path_to_ca_certificate: Custom CA certificate path
```

### Environment Variables
- `AVIATRIX_CONTROLLER_IP`
- `AVIATRIX_USERNAME`
- `AVIATRIX_PASSWORD`

## SDK Generation Requirements

### Supported Languages
1. **Go** - Native Pulumi language
2. **Python** - pip installable package
3. **TypeScript/JavaScript** - npm package
4. **C#** - NuGet package
5. **Java** - Maven package

### SDK Features
- Full type safety and IntelliSense support
- Comprehensive documentation from Terraform schema
- Input/output property mappings
- Proper error handling and validation
- Resource state management

## Build and Distribution

### Build Pipeline
1. **Schema Generation**: Generate `schema.json`
2. **SDK Generation**: Generate language-specific SDKs
3. **Provider Binary**: Build `pulumi-resource-aviatrix`
4. **Package Distribution**: Publish SDKs to respective registries

### Build Outputs
- Provider binary: `pulumi-resource-aviatrix` (Linux, macOS, Windows)
- SDK packages for each language
- Schema JSON file
- Documentation

### Installation Methods
- Pulumi plugin installation: `pulumi plugin install resource aviatrix v0.0.1`
- Language package managers (pip, npm, NuGet, Maven)
- Manual binary installation

## Testing Requirements

### Unit Tests
- Provider configuration validation
- Resource token generation
- Schema generation correctness

### Integration Tests
- Resource CRUD operations
- Data source queries
- Cross-resource dependencies
- State management

### Language SDK Tests
- Go: standard Go tests
- Python: pytest
- TypeScript: Jest/Mocha
- C#: xUnit/NUnit
- Java: JUnit

## Documentation Requirements

### Provider Documentation
- Getting started guide
- Installation instructions
- Configuration reference
- Authentication setup

### Resource Documentation
- Resource descriptions and usage
- Property reference
- Example code in all languages
- Import instructions

### Migration Guide
- Terraform to Pulumi migration
- Configuration conversion
- State import procedures

## Success Criteria

### Functional Requirements
- ✅ All 133 resources accessible via Pulumi
- ✅ All 24 data sources accessible via Pulumi
- ✅ SDKs generated for all 5 languages
- ✅ Provider binary builds successfully
- ✅ Schema generation produces valid output

### Quality Requirements
- 100% resource coverage from Terraform provider
- Proper type safety in all SDKs
- Comprehensive error messages
- Documentation for all resources

### Performance Requirements
- Provider startup time < 2 seconds
- Resource creation/update performance equivalent to Terraform
- Schema generation completes in < 10 seconds

## Dependencies

### External Dependencies
- Pulumi Terraform Bridge v3: `github.com/pulumi/pulumi-terraform-bridge/v3`
- Terraform Aviatrix Provider v3: `github.com/AviatrixSystems/terraform-provider-aviatrix/v3`
- Pulumi SDK: `github.com/pulumi/pulumi/sdk/v3`

### Build Dependencies
- Go 1.21+
- Node.js 18+ (for TypeScript SDK)
- Python 3.8+ (for Python SDK)
- .NET 6+ (for C# SDK)
- Java 11+ (for Java SDK)

## Timeline and Phases

### Phase 1: Core Infrastructure ✅
- Version management
- Provider bridge configuration
- Main entry point
- Schema generator

### Phase 2: SDK Generation
- Schema JSON generation
- Go SDK generation
- Python SDK generation
- TypeScript SDK generation
- C# SDK generation
- Java SDK generation

### Phase 3: Testing & Validation
- Unit test implementation
- Integration test setup
- Example programs in all languages
- Cross-language compatibility testing

### Phase 4: Documentation & Release
- API documentation
- Getting started guides
- Migration documentation
- Package publishing
- Release automation

## Risks and Mitigations

### Risk 1: Terraform Provider Updates
- **Risk**: Upstream provider changes breaking compatibility
- **Mitigation**: Version pinning, automated compatibility testing

### Risk 2: SDK Generation Issues
- **Risk**: Language-specific generation problems
- **Mitigation**: Comprehensive SDK tests, manual review of generated code

### Risk 3: Performance Issues
- **Risk**: Bridge overhead affecting performance
- **Mitigation**: Performance benchmarking, optimization of resource mappings

### Risk 4: Documentation Gaps
- **Risk**: Missing or incomplete documentation
- **Mitigation**: Documentation review process, example validation

## Maintenance and Support

### Version Management
- Semantic versioning (MAJOR.MINOR.PATCH)
- Compatibility matrix with Terraform provider versions
- Regular updates to track upstream provider

### Breaking Changes
- Clearly documented in changelog
- Migration guides provided
- Deprecated features marked with warnings

### Community Support
- GitHub issues for bug reports
- Example repository for common patterns
- Regular release cadence

## Appendix

### Resource Naming Convention
Terraform resource `aviatrix_<name>` maps to Pulumi token `aviatrix:index/<Name>:Type`

Example:
- Terraform: `aviatrix_transit_gateway`
- Pulumi: `aviatrix:index/transitGateway:TransitGateway`

### File Structure
```
pulumi-avx-test-dummy/
├── provider/
│   ├── cmd/
│   │   ├── pulumi-resource-aviatrix/    # Main provider binary
│   │   ├── pulumi-tfgen-aviatrix/       # SDK generator
│   │   └── schema-generator/            # Schema generator
│   ├── pkg/
│   │   └── version/                     # Version management
│   ├── terraform_bridge_provider.go     # Bridge configuration
│   ├── go.mod
│   └── go.sum
├── sdk/
│   ├── go/                              # Go SDK
│   ├── python/                          # Python SDK
│   ├── nodejs/                          # TypeScript/JavaScript SDK
│   ├── dotnet/                          # C# SDK
│   └── java/                            # Java SDK
├── examples/                            # Example programs
└── schema.json                          # Generated schema
```

### Configuration Examples

#### Go
```go
provider, err := aviatrix.NewProvider(ctx, "aviatrix", &aviatrix.ProviderArgs{
    ControllerIp: pulumi.String("controller.example.com"),
    Username:     pulumi.String("admin"),
    Password:     pulumi.String(os.Getenv("AVIATRIX_PASSWORD")),
})
```

#### Python
```python
provider = aviatrix.Provider("aviatrix",
    controller_ip="controller.example.com",
    username="admin",
    password=os.getenv("AVIATRIX_PASSWORD")
)
```

#### TypeScript
```typescript
const provider = new aviatrix.Provider("aviatrix", {
    controllerIp: "controller.example.com",
    username: "admin",
    password: process.env.AVIATRIX_PASSWORD,
});
```

---

**Document Version**: 1.0
**Last Updated**: 2025-10-06
**Status**: Draft
**Owner**: AviatrixSystems
