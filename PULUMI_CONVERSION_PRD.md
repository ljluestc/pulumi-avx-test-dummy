# Product Requirements Document (PRD)
## Aviatrix Terraform to Pulumi Provider Conversion

**Document Version:** 1.0
**Date:** October 4, 2025
**Status:** Draft for Implementation

---

## 1. Executive Summary

### 1.1 Overview
This document outlines the requirements and implementation strategy for converting the existing Aviatrix Terraform Provider (v3.1.10) into a fully functional Pulumi provider. The conversion will create a native Pulumi experience while maintaining 100% feature parity with the Terraform provider.

### 1.2 Current State Analysis
- **Source:** `/root/GolandProjects/pulumi-avx-test-dummy/terraform-provider-aviatrix/`
- **Provider Version:** 8.1.10
- **Terraform Provider SDK:** v2.34.0
- **Resources Count:** 148 resources
- **Data Sources Count:** 23 data sources
- **Total Files:** 328 Go files (resources + data sources + tests)
- **Go Version:** 1.23+ with toolchain 1.24.1
- **Main SDK Library:** `goaviatrix` (custom Aviatrix API client)

### 1.3 Goals
1. **100% Feature Parity:** All 148 resources and 23 data sources must be converted
2. **Production Ready:** Full testing coverage, error handling, and documentation
3. **Pulumi Native:** Leverage Pulumi's multi-language SDK generation (Go, Python, TypeScript, C#, Java)
4. **Maintainable:** Clean architecture that allows easy updates from upstream Terraform provider
5. **Type Safe:** Strong typing across all supported languages

---

## 2. Technical Architecture

### 2.1 Project Structure
```
/root/GolandProjects/pulumi-avx-test-dummy/pulumi-aviatrix/
├── provider/
│   ├── cmd/
│   │   └── pulumi-resource-aviatrix/
│   │       └── main.go                    # Provider entrypoint
│   ├── pkg/
│   │   ├── provider/
│   │   │   ├── provider.go                # Main provider logic
│   │   │   ├── config.go                  # Provider configuration
│   │   │   ├── resources/                 # Resource implementations
│   │   │   │   ├── account.go
│   │   │   │   ├── gateway.go
│   │   │   │   ├── vpc.go
│   │   │   │   └── ... (148 resources)
│   │   │   └── datasources/               # Data source implementations
│   │   │       ├── account.go
│   │   │       ├── vpc.go
│   │   │       └── ... (23 data sources)
│   │   └── version/
│   │       └── version.go
│   ├── resources.go                        # Resource registration
│   ├── go.mod
│   └── go.sum
├── sdk/
│   ├── go/
│   │   └── aviatrix/                      # Generated Go SDK
│   ├── python/
│   │   └── pulumi_aviatrix/               # Generated Python SDK
│   ├── nodejs/
│   │   └── bin/                           # Generated TypeScript SDK
│   ├── dotnet/
│   │   └── Pulumi.Aviatrix/               # Generated C# SDK
│   └── java/
│       └── src/                           # Generated Java SDK
├── examples/
│   ├── go/
│   ├── python/
│   ├── typescript/
│   └── csharp/
├── scripts/
│   ├── build.sh
│   ├── test.sh
│   └── generate-sdks.sh
├── Pulumi.yaml                            # Pulumi plugin metadata
├── Makefile
├── README.md
└── CHANGELOG.md
```

### 2.2 Technology Stack

#### Core Components
1. **Pulumi Provider Framework**
   - `github.com/pulumi/pulumi/sdk/v3` (v3.x)
   - Provider bridge architecture
   - Resource and data source lifecycle management

2. **Schema Bridge**
   - `github.com/pulumi/pulumi-terraform-bridge/v3` (v3.x)
   - Automated conversion from Terraform schemas to Pulumi schemas
   - Property mapping and type conversion

3. **Aviatrix SDK (Reuse)**
   - `github.com/AviatrixSystems/terraform-provider-aviatrix/v3/goaviatrix`
   - No changes needed - reuse existing API client
   - Handles all Aviatrix Controller REST API interactions

4. **Language SDKs**
   - Go: Native
   - Python: `pulumi-python` code generator
   - TypeScript/JavaScript: `pulumi-nodejs` code generator
   - C#/.NET: `pulumi-dotnet` code generator
   - Java: `pulumi-java` code generator

#### Build Tools
- **Go 1.23+** for provider implementation
- **Pulumi CLI 3.x** for SDK generation
- **goreleaser** for release automation
- **GitHub Actions** for CI/CD

---

## 3. Detailed Requirements

### 3.1 Provider Configuration

#### Authentication & Connection
The provider must support the following configuration options:

```go
type ProviderConfig struct {
    ControllerIP           string   // Required: Aviatrix Controller IP
    Username               string   // Required: Admin username
    Password               string   // Required: Admin password (sensitive)
    SkipVersionValidation  bool     // Optional: Skip controller version check
    VerifySSLCertificate   bool     // Optional: Verify SSL certs (default: false)
    PathToCACertificate    string   // Optional: Custom CA certificate path
    IgnoreTags             []IgnoreTag // Optional: Tags to ignore
}

type IgnoreTag struct {
    Keys        []string
    KeyPrefixes []string
}
```

#### Environment Variables
Support standard Pulumi configuration plus environment variables:
- `AVIATRIX_CONTROLLER_IP`
- `AVIATRIX_USERNAME`
- `AVIATRIX_PASSWORD`

#### Version Compatibility
- Support Aviatrix Controller 8.1.x (as per current TF provider)
- Implement version validation (can be skipped via config)

### 3.2 Resource Conversion Requirements

#### 3.2.1 Resource Categories (148 Total)

**Account Management (9 resources)**
- `aviatrix_account` - Cloud account integration
- `aviatrix_account_user` - User management
- `aviatrix_rbac_group` - Role-based access control
- `aviatrix_rbac_group_access_account_attachment`
- `aviatrix_rbac_group_permission_attachment`
- `aviatrix_rbac_group_user_attachment`
- `aviatrix_remote_syslog`
- `aviatrix_sumologic_forwarder`
- `aviatrix_splunk_logging`

**AWS Resources (27 resources)**
- AWS Transit Gateway (TGW) management
- VPC peering
- GuardDuty integration
- TGW attachments and peering
- Direct Connect integration
- VPN connections

**Azure Resources (8 resources)**
- Azure peering
- VNet integration
- Virtual Network Gateway connections
- Spoke native peering

**Gateway Resources (25 resources)**
- Transit gateways
- Spoke gateways
- Edge gateways (Equinix, CloudN)
- Gateway HA configurations
- Custom spoke/transit gateways

**Networking (35 resources)**
- VPC/VNet creation
- Site2Cloud connections
- Segmentation
- Network domains
- Distributed firewalling
- BGP configurations
- FQDN filtering

**Security & Firewall (18 resources)**
- FireNet
- Firewall instances
- Firewall policies
- Distributed firewalling
- Security policies
- Web groups and domains

**Controller Configuration (26 resources)**
- Email configuration
- Certificate management
- BGP settings
- Access control lists
- Security group management
- Logging configuration
- License management

#### 3.2.2 Resource Conversion Pattern

Each resource must implement:

```go
type Resource struct {
    // Schema definition (auto-generated from TF schema)
    Schema map[string]*schema.Schema

    // Lifecycle methods
    Create func(ctx context.Context, d *schema.ResourceData, meta interface{}) error
    Read   func(ctx context.Context, d *schema.ResourceData, meta interface{}) error
    Update func(ctx context.Context, d *schema.ResourceData, meta interface{}) error
    Delete func(ctx context.Context, d *schema.ResourceData, meta interface{}) error

    // Import support
    Importer *schema.ResourceImporter

    // Timeouts
    Timeouts *schema.ResourceTimeout
}
```

**Example: Account Resource**
```go
// Input (Terraform)
resource "aviatrix_account" "aws_account" {
  account_name       = "my-aws-account"
  cloud_type         = 1
  aws_account_number = "123456789012"
  aws_iam            = true
  aws_gateway_role_app = "arn:aws:iam::123456789012:role/aviatrix-role-app"
  aws_gateway_role_ec2 = "arn:aws:iam::123456789012:role/aviatrix-role-ec2"
}

// Output (Pulumi - Go)
account, err := aviatrix.NewAccount(ctx, "aws_account", &aviatrix.AccountArgs{
    AccountName:      pulumi.String("my-aws-account"),
    CloudType:        pulumi.Int(1),
    AwsAccountNumber: pulumi.String("123456789012"),
    AwsIam:           pulumi.Bool(true),
    AwsGatewayRoleApp: pulumi.String("arn:aws:iam::123456789012:role/aviatrix-role-app"),
    AwsGatewayRoleEc2: pulumi.String("arn:aws:iam::123456789012:role/aviatrix-role-ec2"),
})

// Output (Pulumi - Python)
account = aviatrix.Account("aws_account",
    account_name="my-aws-account",
    cloud_type=1,
    aws_account_number="123456789012",
    aws_iam=True,
    aws_gateway_role_app="arn:aws:iam::123456789012:role/aviatrix-role-app",
    aws_gateway_role_ec2="arn:aws:iam::123456789012:role/aviatrix-role-ec2"
)

// Output (Pulumi - TypeScript)
const account = new aviatrix.Account("aws_account", {
    accountName: "my-aws-account",
    cloudType: 1,
    awsAccountNumber: "123456789012",
    awsIam: true,
    awsGatewayRoleApp: "arn:aws:iam::123456789012:role/aviatrix-role-app",
    awsGatewayRoleEc2: "arn:aws:iam::123456789012:role/aviatrix-role-ec2",
});
```

### 3.3 Data Source Conversion (23 Total)

**Data Sources List:**
- `aviatrix_account`
- `aviatrix_caller_identity`
- `aviatrix_controller_metadata`
- `aviatrix_dcf_mwp_attachment_points`
- `aviatrix_dcf_webgroups`
- `aviatrix_device_interfaces`
- `aviatrix_edge_gateway_wan_interface_discovery`
- `aviatrix_firenet`
- `aviatrix_firenet_firewall_manager`
- `aviatrix_firenet_vendor_integration`
- `aviatrix_firewall`
- `aviatrix_firewall_instance_images`
- `aviatrix_gateway`
- `aviatrix_gateway_image`
- `aviatrix_network_domains`
- `aviatrix_smart_groups`
- `aviatrix_spoke_gateway`
- `aviatrix_spoke_gateway_inspection_subnets`
- `aviatrix_spoke_gateways`
- `aviatrix_transit_gateway`
- `aviatrix_transit_gateways`
- `aviatrix_vpc`
- `aviatrix_vpc_tracker`

All data sources must support:
- Read-only operations
- Proper output typing
- Error handling
- Documentation

### 3.4 Bridge Strategy

Use Pulumi Terraform Bridge for automated conversion:

```go
// provider/cmd/pulumi-resource-aviatrix/main.go
package main

import (
    "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
    aviatrix "github.com/AviatrixSystems/terraform-provider-aviatrix/v3/aviatrix"
)

func main() {
    // Create provider
    tfbridge.Main("aviatrix", version.Version,
        Provider(),
        tfbridge.ProviderMetadata{...})
}

func Provider() tfbridge.ProviderInfo {
    p := tfbridge.ProviderInfo{
        P:           aviatrix.Provider(),
        Name:        "aviatrix",
        DisplayName: "Aviatrix",
        Publisher:   "Aviatrix",
        Description: "A Pulumi package for creating and managing Aviatrix cloud resources.",
        Keywords:    []string{"pulumi", "aviatrix", "networking", "multicloud"},
        License:     "Apache-2.0",
        Homepage:    "https://aviatrix.com",
        Repository:  "https://github.com/AviatrixSystems/pulumi-aviatrix",
        GitHubOrg:   "AviatrixSystems",
        Config:      map[string]*tfbridge.SchemaInfo{...},
        Resources:   map[string]*tfbridge.ResourceInfo{...},
        DataSources: map[string]*tfbridge.DataSourceInfo{...},
        JavaScript: &tfbridge.JavaScriptInfo{
            Dependencies: map[string]string{
                "@pulumi/pulumi": "^3.0.0",
            },
            DevDependencies: map[string]string{
                "@types/node": "^14",
            },
        },
        Python: &tfbridge.PythonInfo{
            Requires: map[string]string{
                "pulumi": ">=3.0.0,<4.0.0",
            },
        },
        CSharp: &tfbridge.CSharpInfo{
            PackageReferences: map[string]string{
                "Pulumi": "3.*",
            },
            Namespaces: map[string]string{
                "aviatrix": "Aviatrix",
            },
        },
        Golang: &tfbridge.GolangInfo{
            ImportBasePath: "github.com/AviatrixSystems/pulumi-aviatrix/sdk/go/aviatrix",
        },
        Java: &tfbridge.JavaInfo{
            BasePackage: "com.pulumi.aviatrix",
        },
    }

    return p
}
```

---

## 4. Implementation Phases

### Phase 1: Project Scaffolding & Setup (Week 1)
**Duration:** 5 days
**Owner:** DevOps/Infrastructure Team

#### Tasks:
1. **Directory Structure Creation**
   - Create `pulumi-aviatrix/` root directory
   - Set up provider, SDK, and examples folders
   - Initialize Git repository

2. **Dependency Setup**
   - Create `go.mod` with Pulumi dependencies
   - Import Terraform provider as module
   - Configure Pulumi bridge framework

3. **Build System**
   - Create Makefile with targets:
     - `make build` - Build provider binary
     - `make install` - Install to local Pulumi plugins
     - `make generate` - Generate SDKs
     - `make test` - Run tests
     - `make clean` - Clean build artifacts
   - Set up goreleaser configuration
   - Create build scripts for CI/CD

4. **CI/CD Pipeline**
   - GitHub Actions workflow for builds
   - Automated testing on PR
   - Release automation
   - Multi-platform binary builds (Linux, macOS, Windows)

**Deliverables:**
- Working build system
- CI/CD pipeline configured
- Repository structure ready

**Success Criteria:**
- `make build` successfully compiles provider
- `make install` installs to Pulumi plugin directory
- CI pipeline passes on sample PR

---

### Phase 2: Core Provider Implementation (Week 2-3)
**Duration:** 10 days
**Owner:** Backend Engineering Team

#### Tasks:
1. **Provider Bridge Setup**
   ```go
   // Implement main.go
   // Configure provider metadata
   // Set up resource/datasource mapping
   ```

2. **Provider Configuration**
   - Map Terraform schema to Pulumi config
   - Implement authentication logic
   - Configure SSL verification options
   - Set up tag ignore logic

3. **Resource Schema Mapping**
   - Create resource info mappings for all 148 resources
   - Define property overrides where needed
   - Configure computed vs. input properties
   - Set up deprecation notices

4. **Data Source Schema Mapping**
   - Map all 23 data sources
   - Configure output properties
   - Set up filtering logic

5. **Client Reuse**
   - Integrate existing `goaviatrix` client
   - Ensure compatibility with Pulumi context
   - Implement proper error wrapping

**Deliverables:**
- Provider compiles and runs
- Basic resource registration working
- Configuration validated

**Success Criteria:**
- Provider starts successfully
- At least 5 sample resources work end-to-end
- Authentication succeeds against test controller

---

### Phase 3: Resource Implementation (Week 4-7)
**Duration:** 20 days (4 weeks)
**Owner:** Backend Engineering Team (2-3 engineers)

#### Strategy: Batch Processing
Divide resources into logical groups and implement in parallel:

**Batch 1: Core Resources (Week 4)**
- Account management (9 resources)
- Basic VPC/VNet (5 resources)
- Simple gateway resources (10 resources)

**Batch 2: AWS Integration (Week 5)**
- AWS TGW resources (15 resources)
- AWS peering (5 resources)
- AWS-specific features (7 resources)

**Batch 3: Advanced Networking (Week 6)**
- Site2Cloud (8 resources)
- BGP configuration (10 resources)
- Segmentation (12 resources)

**Batch 4: Security & Firewall (Week 7)**
- FireNet (10 resources)
- Distributed firewalling (8 resources)
- Security policies (10 resources)

**Batch 5: Controller & Misc (Week 7)**
- Controller configuration (26 resources)
- Edge gateways (8 resources)
- Remaining resources (20 resources)

#### Per-Resource Implementation Checklist:
- [ ] Schema mapping complete
- [ ] Create function implemented
- [ ] Read function implemented
- [ ] Update function implemented
- [ ] Delete function implemented
- [ ] Import support added
- [ ] Timeout configuration
- [ ] Error handling
- [ ] Unit tests written
- [ ] Integration test written
- [ ] Documentation generated
- [ ] Example code created

**Deliverables:**
- All 148 resources implemented
- Unit tests for each resource
- Integration test suite

**Success Criteria:**
- Each resource passes CRUD operations
- Error handling covers edge cases
- Tests achieve >80% code coverage

---

### Phase 4: Data Source Implementation (Week 8)
**Duration:** 5 days
**Owner:** Backend Engineering Team

#### Tasks:
1. Implement all 23 data sources
2. Add filtering capabilities
3. Implement output transformations
4. Write tests for each data source
5. Generate documentation

**Deliverables:**
- All 23 data sources functional
- Tests passing

**Success Criteria:**
- Data sources return accurate data
- Filtering works correctly
- Tests pass

---

### Phase 5: SDK Generation & Testing (Week 9-10)
**Duration:** 10 days
**Owner:** Full Team

#### Tasks:
1. **SDK Generation**
   - Generate Go SDK
   - Generate Python SDK
   - Generate TypeScript SDK
   - Generate C# SDK
   - Generate Java SDK

2. **SDK Testing**
   - Write example programs in each language
   - Test resource creation/update/deletion
   - Validate type safety
   - Check IntelliSense/autocomplete

3. **Documentation Generation**
   - Auto-generate API reference docs
   - Create getting started guides
   - Write migration guide from Terraform
   - Create language-specific examples

**Deliverables:**
- SDKs for all 5 languages
- Working examples in each language
- Complete documentation

**Success Criteria:**
- All SDKs compile without errors
- Examples run successfully
- Documentation is accurate and complete

---

### Phase 6: Integration Testing (Week 11)
**Duration:** 5 days
**Owner:** QA + Backend Team

#### Test Coverage Requirements:

1. **Unit Tests**
   - Schema validation
   - Property mapping
   - Error handling
   - Input validation

2. **Integration Tests**
   - Real Aviatrix Controller interaction
   - Complete CRUD operations
   - Import functionality
   - Update scenarios
   - Edge cases

3. **End-to-End Tests**
   - Multi-resource stacks
   - Dependency management
   - State management
   - Rollback scenarios

4. **Language SDK Tests**
   - Go example projects
   - Python example projects
   - TypeScript example projects
   - C# example projects
   - Java example projects

**Test Infrastructure:**
- Aviatrix test controller
- Cloud accounts (AWS, Azure, GCP, OCI)
- Automated test runner
- Nightly test suite

**Deliverables:**
- Test suite with >80% coverage
- Integration tests passing
- E2E test scenarios documented

**Success Criteria:**
- All tests pass consistently
- No critical bugs
- Performance benchmarks met

---

### Phase 7: Documentation & Examples (Week 12)
**Duration:** 5 days
**Owner:** Documentation Team + Engineering

#### Deliverables:

1. **API Documentation**
   - Auto-generated reference for all resources
   - Auto-generated reference for all data sources
   - Configuration options documented
   - Type definitions

2. **Getting Started Guides**
   - Installation instructions
   - Provider configuration
   - First resource deployment
   - Authentication setup

3. **Migration Guide**
   - Terraform to Pulumi conversion steps
   - Schema differences
   - Common patterns translation
   - State migration (if applicable)

4. **Examples**
   - Basic examples (5+ per language)
   - Advanced multi-resource examples
   - Real-world scenarios:
     - Multi-cloud transit network
     - FireNet deployment
     - Spoke-transit architecture
     - Site2Cloud VPN

5. **Troubleshooting Guide**
   - Common errors and solutions
   - Debugging tips
   - FAQ

**Success Criteria:**
- Documentation is comprehensive
- Examples run without modification
- Migration guide tested with real projects

---

### Phase 8: Release Preparation (Week 13)
**Duration:** 5 days
**Owner:** DevOps + Product Team

#### Tasks:

1. **Version Management**
   - Set initial version (1.0.0)
   - Create CHANGELOG
   - Tag release in Git

2. **Package Publishing**
   - Publish to Pulumi Registry
   - Publish SDKs to package managers:
     - Go: pkg.go.dev
     - Python: PyPI
     - TypeScript: npm
     - C#: NuGet
     - Java: Maven Central

3. **Binary Distribution**
   - Build multi-platform binaries
   - Create GitHub releases
   - Upload artifacts

4. **Registry Listing**
   - Create Pulumi Registry listing
   - Add descriptions, tags
   - Upload logo and assets

5. **Announcement**
   - Blog post
   - Release notes
   - Community notification

**Deliverables:**
- Published packages
- GitHub release
- Registry listing
- Announcement materials

**Success Criteria:**
- All packages installable
- Registry listing live
- No critical issues in release

---

## 5. Testing Strategy

### 5.1 Unit Testing
```go
// Example unit test structure
func TestAccountResource_Create(t *testing.T) {
    tests := []struct {
        name    string
        args    AccountArgs
        want    *Account
        wantErr bool
    }{
        {
            name: "valid AWS account",
            args: AccountArgs{
                AccountName:      pulumi.String("test"),
                CloudType:        pulumi.Int(1),
                AwsAccountNumber: pulumi.String("123456789012"),
            },
            want:    &Account{...},
            wantErr: false,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 5.2 Integration Testing
- Test against live Aviatrix controller
- Validate complete CRUD lifecycle
- Test inter-resource dependencies
- Verify state management

### 5.3 Acceptance Testing
```go
func TestAccAccount_basic(t *testing.T) {
    program := pulumi.RunFunc(func(ctx *pulumi.Context) error {
        account, err := aviatrix.NewAccount(ctx, "test", &aviatrix.AccountArgs{
            AccountName:      pulumi.String("test-account"),
            CloudType:        pulumi.Int(1),
            AwsAccountNumber: pulumi.String("123456789012"),
        })
        if err != nil {
            return err
        }

        ctx.Export("accountId", account.ID())
        return nil
    })

    // Run program and verify outputs
    result, err := auto.UpsertStackInlineSource(ctx, "test-stack", "test-project", program)
    // Assertions...
}
```

### 5.4 Test Coverage Goals
- Unit tests: >80% code coverage
- Integration tests: All resources tested
- E2E tests: Critical user journeys covered
- Language SDKs: At least 3 examples per language

---

## 6. Quality Assurance

### 6.1 Code Quality Standards
- **Linting:** golangci-lint with strict rules
- **Formatting:** gofmt, goimports
- **Static Analysis:** go vet, staticcheck
- **Security Scanning:** gosec
- **Dependency Audit:** go mod verify

### 6.2 Review Process
- All code must be peer-reviewed
- Automated checks must pass before merge
- Breaking changes require approval from maintainers
- Security changes require security team review

### 6.3 Performance Requirements
- Provider startup: <2 seconds
- Resource creation: <30 seconds (depends on Aviatrix API)
- SDK import time: <1 second
- Memory footprint: <100MB idle

### 6.4 Compatibility Matrix
| Component | Version |
|-----------|---------|
| Pulumi CLI | >=3.0.0 |
| Go | >=1.23 |
| Python | >=3.8 |
| Node.js | >=14.x |
| .NET | >=6.0 |
| Java | >=11 |
| Aviatrix Controller | 8.1.x |

---

## 7. Documentation Requirements

### 7.1 Required Documentation
1. **README.md**
   - Project overview
   - Installation instructions
   - Quick start example
   - Links to detailed docs

2. **API Reference**
   - Auto-generated from code
   - All resources documented
   - All data sources documented
   - All configuration options

3. **User Guides**
   - Getting started
   - Authentication setup
   - Common patterns
   - Best practices

4. **Developer Guide**
   - Contributing guidelines
   - Build instructions
   - Testing guide
   - Release process

5. **Migration Guide**
   - Terraform to Pulumi conversion
   - Schema mapping reference
   - State migration steps

### 7.2 Example Requirements
Each resource must have:
- Basic usage example
- At least one advanced example
- Examples in all 5 supported languages

---

## 8. Dependencies & Prerequisites

### 8.1 External Dependencies
```go
require (
    github.com/pulumi/pulumi/sdk/v3 v3.x.x
    github.com/pulumi/pulumi-terraform-bridge/v3 v3.x.x
    github.com/AviatrixSystems/terraform-provider-aviatrix/v3 v3.1.10
    // Additional dependencies from TF provider
)
```

### 8.2 Development Tools
- Go 1.23+
- Pulumi CLI 3.x
- Make
- goreleaser
- golangci-lint
- Python 3.8+ (for Python SDK testing)
- Node.js 14+ (for TypeScript SDK testing)
- .NET 6+ (for C# SDK testing)
- Java 11+ (for Java SDK testing)

### 8.3 Test Infrastructure
- Aviatrix Controller (test environment)
- Cloud provider accounts:
  - AWS account
  - Azure subscription
  - GCP project
  - OCI tenancy
- CI/CD environment (GitHub Actions)

---

## 9. Risk Assessment & Mitigation

### 9.1 Technical Risks

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Pulumi bridge incompatibility with complex TF schemas | High | Medium | Thorough testing; manual overrides where needed |
| API breaking changes in Aviatrix Controller | High | Low | Version pinning; comprehensive integration tests |
| SDK generation issues for specific languages | Medium | Medium | Manual adjustments; language-specific overrides |
| Performance degradation vs TF provider | Medium | Low | Performance benchmarking; optimization |
| State migration challenges | Medium | Medium | Clear migration docs; state import tooling |

### 9.2 Schedule Risks

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Resource implementation taking longer than estimated | High | Medium | Parallel development; increase team size |
| Testing infrastructure delays | Medium | Low | Early setup; backup test environment |
| SDK generation bugs | Medium | Medium | Start SDK work early; allocate buffer time |
| Documentation writing delays | Low | Medium | Start docs in parallel with dev; templates |

### 9.3 Resource Risks

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Insufficient engineering resources | High | Low | Prioritize core resources; staged rollout |
| Aviatrix controller access issues | High | Low | Backup controller; local dev environment |
| Knowledge gaps in Pulumi | Medium | Medium | Training sessions; Pulumi support engagement |

---

## 10. Success Metrics

### 10.1 Functional Metrics
- ✅ 100% of resources (148/148) implemented
- ✅ 100% of data sources (23/23) implemented
- ✅ All 5 language SDKs generated and functional
- ✅ >80% code coverage in tests
- ✅ 0 critical bugs in release

### 10.2 Quality Metrics
- ✅ All CI/CD checks passing
- ✅ Security scans clean
- ✅ Performance benchmarks met
- ✅ Documentation completeness >90%

### 10.3 Adoption Metrics (Post-Launch)
- Downloads/installs tracked
- Community feedback collected
- GitHub stars and forks
- Issue resolution time <7 days

---

## 11. Deployment & Release

### 11.1 Versioning Strategy
- Semantic versioning (SemVer 2.0)
- Initial release: v1.0.0
- Maintain compatibility with Terraform provider versions
- Clear upgrade paths

### 11.2 Release Checklist
- [ ] All tests passing
- [ ] Documentation complete
- [ ] Examples validated
- [ ] CHANGELOG updated
- [ ] Version bumped
- [ ] Git tag created
- [ ] Binaries built for all platforms
- [ ] Packages published to registries
- [ ] Pulumi Registry updated
- [ ] Release notes published
- [ ] Announcement sent

### 11.3 Distribution Channels
1. **Pulumi Registry** (primary)
2. **GitHub Releases** (binaries)
3. **Package Managers:**
   - Go: pkg.go.dev
   - Python: PyPI
   - Node: npmjs.com
   - .NET: nuget.org
   - Java: Maven Central

---

## 12. Maintenance & Support

### 12.1 Ongoing Maintenance
- Monitor Terraform provider updates
- Sync with Aviatrix controller releases
- Address community issues and PRs
- Security patches within 48 hours
- Regular dependency updates

### 12.2 Support Channels
- GitHub Issues (primary)
- Pulumi Community Slack
- Stack Overflow
- Email support (for enterprise)

### 12.3 Update Cadence
- Patch releases: As needed
- Minor releases: Monthly or as needed
- Major releases: Quarterly or for breaking changes

---

## 13. Acceptance Criteria

### 13.1 Must Have (P0)
- ✅ All 148 resources implemented and tested
- ✅ All 23 data sources implemented and tested
- ✅ Provider configuration working (auth, SSL, etc.)
- ✅ SDKs generated for Go, Python, TypeScript
- ✅ Basic examples for each language
- ✅ CI/CD pipeline functional
- ✅ Documentation published
- ✅ Published to Pulumi Registry

### 13.2 Should Have (P1)
- ✅ C# and Java SDKs
- ✅ Advanced examples and tutorials
- ✅ Migration guide from Terraform
- ✅ Performance benchmarks
- ✅ Integration test suite

### 13.3 Nice to Have (P2)
- ⭕ VS Code extension support
- ⭕ Terraform state import tool
- ⭕ Interactive documentation
- ⭕ Video tutorials

---

## 14. Implementation Commands & Scripts

### 14.1 Initial Setup Script
```bash
#!/bin/bash
# setup.sh - Initialize Pulumi Aviatrix provider

set -e

PROJECT_ROOT="/root/GolandProjects/pulumi-avx-test-dummy"
PULUMI_DIR="$PROJECT_ROOT/pulumi-aviatrix"

echo "Creating Pulumi Aviatrix provider structure..."

# Create directory structure
mkdir -p "$PULUMI_DIR/provider/cmd/pulumi-resource-aviatrix"
mkdir -p "$PULUMI_DIR/provider/pkg/provider/resources"
mkdir -p "$PULUMI_DIR/provider/pkg/provider/datasources"
mkdir -p "$PULUMI_DIR/provider/pkg/version"
mkdir -p "$PULUMI_DIR/sdk/"{go,python,nodejs,dotnet,java}
mkdir -p "$PULUMI_DIR/examples/"{go,python,typescript,csharp,java}
mkdir -p "$PULUMI_DIR/scripts"
mkdir -p "$PULUMI_DIR/tests"

# Create go.mod
cat > "$PULUMI_DIR/provider/go.mod" << 'EOF'
module github.com/AviatrixSystems/pulumi-aviatrix/provider

go 1.23

require (
    github.com/pulumi/pulumi/sdk/v3 v3.136.1
    github.com/pulumi/pulumi-terraform-bridge/v3 v3.91.0
    github.com/AviatrixSystems/terraform-provider-aviatrix/v3 v3.1.10
)
EOF

# Create version file
cat > "$PULUMI_DIR/provider/pkg/version/version.go" << 'EOF'
package version

var Version = "1.0.0"
EOF

echo "✅ Directory structure created"
echo "Next steps:"
echo "  1. cd $PULUMI_DIR/provider"
echo "  2. go mod tidy"
echo "  3. Implement main.go"
echo "  4. Run: make build"
```

### 14.2 Build Makefile
```makefile
# Makefile for Pulumi Aviatrix Provider

PROJECT_NAME := pulumi-aviatrix
PROVIDER_NAME := aviatrix
PROVIDER_VERSION ?= $(shell pulumictl get version)

WORKING_DIR := $(shell pwd)
PROVIDER_DIR := $(WORKING_DIR)/provider
SDK_DIR := $(WORKING_DIR)/sdk

.PHONY: build install generate test clean

build:
	cd $(PROVIDER_DIR) && go build -o $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME) ./cmd/pulumi-resource-$(PROVIDER_NAME)

install: build
	cp $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME) ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v$(PROVIDER_VERSION)/

generate: build install
	cd $(PROVIDER_DIR) && pulumi package gen-sdk --language go $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)
	cd $(PROVIDER_DIR) && pulumi package gen-sdk --language python $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)
	cd $(PROVIDER_DIR) && pulumi package gen-sdk --language nodejs $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)
	cd $(PROVIDER_DIR) && pulumi package gen-sdk --language dotnet $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)
	cd $(PROVIDER_DIR) && pulumi package gen-sdk --language java $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)

test:
	cd $(PROVIDER_DIR) && go test -v ./...

test-integration:
	cd $(PROVIDER_DIR) && go test -v -tags=integration ./...

clean:
	rm -rf $(WORKING_DIR)/bin
	rm -rf $(SDK_DIR)/go/$(PROVIDER_NAME)
	rm -rf $(SDK_DIR)/python/pulumi_$(PROVIDER_NAME)
	rm -rf $(SDK_DIR)/nodejs/bin
	rm -rf $(SDK_DIR)/dotnet/Pulumi.$(PROVIDER_NAME)
	rm -rf $(SDK_DIR)/java/src

lint:
	cd $(PROVIDER_DIR) && golangci-lint run

fmt:
	cd $(PROVIDER_DIR) && go fmt ./...

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build             - Build the provider binary"
	@echo "  install           - Install provider to Pulumi plugins directory"
	@echo "  generate          - Generate SDKs for all languages"
	@echo "  test              - Run unit tests"
	@echo "  test-integration  - Run integration tests"
	@echo "  clean             - Clean build artifacts"
	@echo "  lint              - Run linters"
	@echo "  fmt               - Format code"
```

### 14.3 Resource Implementation Template
```go
// provider/pkg/provider/resources/account.go
package resources

import (
    "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
    "github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func GetAccountResource() *tfbridge.ResourceInfo {
    return &tfbridge.ResourceInfo{
        Tok: tokens.NewTypeToken("aviatrix", "Account"),
        Fields: map[string]*tfbridge.SchemaInfo{
            "account_name": {
                Name: "accountName",
            },
            "cloud_type": {
                Name: "cloudType",
            },
            "aws_account_number": {
                Name: "awsAccountNumber",
            },
            "aws_iam": {
                Name: "awsIam",
            },
            // ... more field mappings
        },
        Docs: &tfbridge.DocInfo{
            Description: "Manages Aviatrix cloud account",
            Source:      "account.html.markdown",
        },
    }
}
```

---

## 15. Appendices

### Appendix A: Complete Resource List (148 Resources)

<details>
<summary>Click to expand full resource list</summary>

1. aviatrix_account
2. aviatrix_account_user
3. aviatrix_aws_guard_duty
4. aviatrix_aws_peer
5. aviatrix_aws_tgw
6. aviatrix_aws_tgw_connect
7. aviatrix_aws_tgw_connect_peer
8. aviatrix_aws_tgw_directconnect
9. aviatrix_aws_tgw_intra_domain_inspection
10. aviatrix_aws_tgw_network_domain
... (148 total - see provider.go for complete list)

</details>

### Appendix B: Complete Data Source List (23 Data Sources)

1. aviatrix_account
2. aviatrix_caller_identity
3. aviatrix_controller_metadata
4. aviatrix_dcf_mwp_attachment_points
5. aviatrix_dcf_webgroups
6. aviatrix_device_interfaces
7. aviatrix_edge_gateway_wan_interface_discovery
8. aviatrix_firenet
9. aviatrix_firenet_firewall_manager
10. aviatrix_firenet_vendor_integration
11. aviatrix_firewall
12. aviatrix_firewall_instance_images
13. aviatrix_gateway
14. aviatrix_gateway_image
15. aviatrix_network_domains
16. aviatrix_smart_groups
17. aviatrix_spoke_gateway
18. aviatrix_spoke_gateway_inspection_subnets
19. aviatrix_spoke_gateways
20. aviatrix_transit_gateway
21. aviatrix_transit_gateways
22. aviatrix_vpc
23. aviatrix_vpc_tracker

### Appendix C: Technology Stack Summary

**Core Framework:**
- Pulumi SDK v3
- Pulumi Terraform Bridge v3
- Go 1.23+

**Existing Dependencies (Reused):**
- github.com/AviatrixSystems/terraform-provider-aviatrix/v3
- github.com/AviatrixSystems/terraform-provider-aviatrix/v3/goaviatrix

**Language SDKs:**
- Go (native)
- Python 3.8+
- TypeScript/JavaScript (Node 14+)
- C# (.NET 6+)
- Java (11+)

**Build & CI/CD:**
- Make
- goreleaser
- GitHub Actions
- golangci-lint

### Appendix D: Reference Links

- Pulumi Documentation: https://www.pulumi.com/docs/
- Pulumi Bridge Guide: https://github.com/pulumi/pulumi-terraform-bridge
- Aviatrix Terraform Provider: https://github.com/AviatrixSystems/terraform-provider-aviatrix
- Aviatrix API Documentation: https://docs.aviatrix.com/
- Pulumi Registry: https://www.pulumi.com/registry/

---

## 16. Sign-off & Approvals

**Document Status:** Ready for Implementation
**Last Updated:** October 4, 2025

**Approval Required From:**
- [ ] Engineering Lead
- [ ] Product Manager
- [ ] DevOps Lead
- [ ] QA Lead
- [ ] Technical Writer
- [ ] Security Team

**Estimated Timeline:** 13 weeks (3.25 months)
**Estimated Effort:** ~800 engineering hours

---

**END OF PRD**
