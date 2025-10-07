package aviatrix

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleAviatrixProvider(t *testing.T) {
	provider := NewSimpleAviatrixProvider("1.0.0")
	assert.NotNil(t, provider)
	assert.Equal(t, "1.0.0", provider.GetVersion())
}

func TestProvider(t *testing.T) {
	provider := Provider()
	assert.NotNil(t, provider)
	assert.Equal(t, "0.0.1", provider.GetVersion())
}

func TestGetResources(t *testing.T) {
	provider := Provider()
	resources := provider.GetResources()
	assert.NotNil(t, resources)
	assert.Greater(t, len(resources), 0)
	assert.Contains(t, resources, "aviatrix:index:Gateway")
	assert.Contains(t, resources, "aviatrix:index:SpokeGateway")
}

func TestGetDataSources(t *testing.T) {
	provider := Provider()
	dataSources := provider.GetDataSources()
	assert.NotNil(t, dataSources)
	assert.Greater(t, len(dataSources), 0)
	assert.Contains(t, dataSources, "aviatrix:index:getAccount")
	assert.Contains(t, dataSources, "aviatrix:index:getGateway")
}

func TestGetConfig(t *testing.T) {
	provider := Provider()
	config := provider.GetConfig()
	assert.NotNil(t, config)
	
	// Check required fields
	requiredFields := []string{"controller_ip", "username", "password"}
	for _, field := range requiredFields {
		assert.Contains(t, config, field, "Config should contain field %s", field)
		fieldConfig := config[field].(map[string]interface{})
		assert.Equal(t, "string", fieldConfig["type"])
		assert.Equal(t, true, fieldConfig["required"])
	}
	
	// Check optional fields
	optionalFields := []string{"skip_version_validation", "verify_ssl_certificate", "path_to_ca_certificate"}
	for _, field := range optionalFields {
		assert.Contains(t, config, field, "Config should contain field %s", field)
		fieldConfig := config[field].(map[string]interface{})
		if field == "path_to_ca_certificate" {
			assert.Equal(t, "string", fieldConfig["type"])
		} else {
			assert.Equal(t, "boolean", fieldConfig["type"])
		}
	}
}

func TestValidateConfig(t *testing.T) {
	provider := Provider()
	
	// Valid config
	validConfig := map[string]interface{}{
		"controller_ip": "1.2.3.4",
		"username": "admin",
		"password": "password123",
	}
	err := provider.ValidateConfig(validConfig)
	assert.NoError(t, err)
	
	// Valid config with optional fields
	fullConfig := map[string]interface{}{
		"controller_ip": "1.2.3.4",
		"username": "admin", 
		"password": "password123",
		"skip_version_validation": true,
		"verify_ssl_certificate": true,
		"path_to_ca_certificate": "/path/to/ca.pem",
	}
	err = provider.ValidateConfig(fullConfig)
	assert.NoError(t, err)
	
	// Missing required field
	missingFieldConfig := map[string]interface{}{
		"username": "admin",
		"password": "password123",
	}
	err = provider.ValidateConfig(missingFieldConfig)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required configuration field 'controller_ip' is missing")
	
	// Empty required field
	emptyFieldConfig := map[string]interface{}{
		"controller_ip": "",
		"username": "admin",
		"password": "password123",
	}
	err = provider.ValidateConfig(emptyFieldConfig)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required configuration field 'controller_ip' cannot be empty")
	
	// Nil required field
	nilFieldConfig := map[string]interface{}{
		"controller_ip": nil,
		"username": "admin",
		"password": "password123",
	}
	err = provider.ValidateConfig(nilFieldConfig)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required configuration field 'controller_ip' cannot be nil")
}

func TestProviderProperties(t *testing.T) {
	provider := Provider()
	assert.Equal(t, "0.0.1", provider.GetVersion())
	assert.Equal(t, "A Pulumi package for creating and managing Aviatrix cloud resources.", provider.GetDescription())
	assert.Greater(t, len(provider.GetResources()), 100)
	assert.Greater(t, len(provider.GetDataSources()), 20)
}

func TestResourceCompleteness(t *testing.T) {
	provider := Provider()
	resources := provider.GetResources()
	assert.Contains(t, resources, "aviatrix:index:Gateway")
	assert.Contains(t, resources, "aviatrix:index:SpokeGateway")
	assert.Contains(t, resources, "aviatrix:index:TransitGateway")
	assert.Contains(t, resources, "aviatrix:index:Vpc")
	assert.Contains(t, resources, "aviatrix:index:Account")
	assert.Contains(t, resources, "aviatrix:index:FireNet")
	assert.Contains(t, resources, "aviatrix:index:Firewall")
	assert.Contains(t, resources, "aviatrix:index:EdgeSpoke")
	assert.Contains(t, resources, "aviatrix:index:Site2Cloud")
	assert.Contains(t, resources, "aviatrix:index:VpnUser")
}

func TestDataSourceCompleteness(t *testing.T) {
	provider := Provider()
	dataSources := provider.GetDataSources()
	assert.Contains(t, dataSources, "aviatrix:index:getAccount")
	assert.Contains(t, dataSources, "aviatrix:index:getGateway")
	assert.Contains(t, dataSources, "aviatrix:index:getVpc")
	assert.Contains(t, dataSources, "aviatrix:index:getFireNet")
	assert.Contains(t, dataSources, "aviatrix:index:getFirewall")
	assert.Contains(t, dataSources, "aviatrix:index:getSpokeGateway")
	assert.Contains(t, dataSources, "aviatrix:index:getTransitGateway")
}

func TestConfigValidationEdgeCases(t *testing.T) {
	provider := Provider()
	
	// Test with nil values
	configWithNil := map[string]interface{}{
		"controller_ip": nil,
		"username": "admin",
		"password": "password123",
	}
	err := provider.ValidateConfig(configWithNil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required configuration field 'controller_ip' cannot be nil")
	
	// Test with empty strings
	configWithEmpty := map[string]interface{}{
		"controller_ip": "",
		"username": "",
		"password": "",
	}
	err = provider.ValidateConfig(configWithEmpty)
	assert.Error(t, err) // Empty strings should be invalid
	assert.Contains(t, err.Error(), "required configuration field 'controller_ip' cannot be empty")
	
	// Test with extra fields
	configWithExtra := map[string]interface{}{
		"controller_ip": "1.2.3.4",
		"username": "admin",
		"password": "password123",
		"extra_field": "some_value",
	}
	err = provider.ValidateConfig(configWithExtra)
	assert.NoError(t, err) // Extra fields should be ignored
}

func TestProviderConsistency(t *testing.T) {
	provider1 := Provider()
	provider2 := Provider()
	assert.Equal(t, provider1.GetVersion(), provider2.GetVersion())
	assert.Equal(t, provider1.GetDescription(), provider2.GetDescription())
	assert.Equal(t, provider1.GetResources(), provider2.GetResources())
	assert.Equal(t, provider1.GetDataSources(), provider2.GetDataSources())
	assert.Equal(t, provider1.GetConfig(), provider2.GetConfig())
}

func TestResourceUniqueness(t *testing.T) {
	provider := Provider()
	resources := provider.GetResources()
	resourceMap := make(map[string]bool)
	for _, resource := range resources {
		assert.False(t, resourceMap[resource], "Resource %s appears multiple times", resource)
		resourceMap[resource] = true
	}
}

func TestDataSourceUniqueness(t *testing.T) {
	provider := Provider()
	dataSources := provider.GetDataSources()
	dataSourceMap := make(map[string]bool)
	for _, dataSource := range dataSources {
		assert.False(t, dataSourceMap[dataSource], "Data source %s appears multiple times", dataSource)
		dataSourceMap[dataSource] = true
	}
}

func TestConfigFieldTypes(t *testing.T) {
	provider := Provider()
	config := provider.GetConfig()
	
	assert.Equal(t, "string", config["controller_ip"].(map[string]interface{})["type"])
	assert.Equal(t, "string", config["username"].(map[string]interface{})["type"])
	assert.Equal(t, "string", config["password"].(map[string]interface{})["type"])
	assert.Equal(t, "boolean", config["skip_version_validation"].(map[string]interface{})["type"])
	assert.Equal(t, "boolean", config["verify_ssl_certificate"].(map[string]interface{})["type"])
	assert.Equal(t, "string", config["path_to_ca_certificate"].(map[string]interface{})["type"])
}