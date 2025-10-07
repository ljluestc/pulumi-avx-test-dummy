package main

import (
	"testing"
)

func TestNewStandaloneProvider(t *testing.T) {
	provider := NewStandaloneProvider()
	
	if provider == nil {
		t.Fatal("NewStandaloneProvider() returned nil")
	}
	
	if provider.Name != "aviatrix" {
		t.Errorf("Expected name 'aviatrix', got '%s'", provider.Name)
	}
	
	if provider.Version != "0.0.1" {
		t.Errorf("Expected version '0.0.1', got '%s'", provider.Version)
	}
}

func TestGetResources(t *testing.T) {
	provider := NewStandaloneProvider()
	resources := provider.GetResources()
	
	if resources == nil {
		t.Fatal("GetResources() returned nil")
	}
	
	if len(resources) == 0 {
		t.Fatal("GetResources() returned empty map")
	}
	
	// Check for some expected resources
	expectedResources := []string{
		"aviatrix_account",
		"aviatrix_gateway",
		"aviatrix_spoke_gateway",
		"aviatrix_transit_gateway",
		"aviatrix_vpc",
		"aviatrix_firenet",
		"aviatrix_firewall",
	}
	
	for _, resource := range expectedResources {
		if _, exists := resources[resource]; !exists {
			t.Errorf("Expected resource '%s' not found", resource)
		}
	}
}

func TestGetDataSources(t *testing.T) {
	provider := NewStandaloneProvider()
	dataSources := provider.GetDataSources()
	
	if dataSources == nil {
		t.Fatal("GetDataSources() returned nil")
	}
	
	if len(dataSources) == 0 {
		t.Fatal("GetDataSources() returned empty map")
	}
	
	// Check for some expected data sources
	expectedDataSources := []string{
		"aviatrix_account",
		"aviatrix_caller_identity",
		"aviatrix_controller_metadata",
		"aviatrix_web_group",
		"aviatrix_gateway",
		"aviatrix_vpc",
	}
	
	for _, dataSource := range expectedDataSources {
		if _, exists := dataSources[dataSource]; !exists {
			t.Errorf("Expected data source '%s' not found", dataSource)
		}
	}
}

func TestGetResourceCount(t *testing.T) {
	provider := NewStandaloneProvider()
	count := provider.GetResourceCount()
	
	if count != 133 {
		t.Errorf("Expected 133 resources, got %d", count)
	}
}

func TestGetDataSourceCount(t *testing.T) {
	provider := NewStandaloneProvider()
	count := provider.GetDataSourceCount()
	
	if count != 23 {
		t.Errorf("Expected 23 data sources, got %d", count)
	}
}

func TestValidateConfig(t *testing.T) {
	provider := NewStandaloneProvider()
	
	// Test valid config
	validConfig := map[string]interface{}{
		"controller_ip": "demo-controller.aviatrix.com",
		"username":      "demo-user",
		"password":      "demo-password",
	}
	
	err := provider.ValidateConfig(validConfig)
	if err != nil {
		t.Errorf("Valid config should not return error, got: %v", err)
	}
	
	// Test missing controller_ip
	invalidConfig1 := map[string]interface{}{
		"username": "demo-user",
		"password": "demo-password",
	}
	
	err = provider.ValidateConfig(invalidConfig1)
	if err == nil {
		t.Error("Invalid config should return error for missing controller_ip")
	}
	
	// Test missing username
	invalidConfig2 := map[string]interface{}{
		"controller_ip": "demo-controller.aviatrix.com",
		"password":      "demo-password",
	}
	
	err = provider.ValidateConfig(invalidConfig2)
	if err == nil {
		t.Error("Invalid config should return error for missing username")
	}
	
	// Test missing password
	invalidConfig3 := map[string]interface{}{
		"controller_ip": "demo-controller.aviatrix.com",
		"username":      "demo-user",
	}
	
	err = provider.ValidateConfig(invalidConfig3)
	if err == nil {
		t.Error("Invalid config should return error for missing password")
	}
}

func TestProviderProperties(t *testing.T) {
	provider := NewStandaloneProvider()
	
	if provider.Name == "" {
		t.Error("Provider name should not be empty")
	}
	
	if provider.Version == "" {
		t.Error("Provider version should not be empty")
	}
	
	if provider.Resources == nil {
		t.Error("Provider resources should not be nil")
	}
	
	if provider.DataSources == nil {
		t.Error("Provider data sources should not be nil")
	}
}

func TestResourceCompleteness(t *testing.T) {
	provider := NewStandaloneProvider()
	resources := provider.GetResources()
	
	// Check that we have exactly 133 resources
	if len(resources) != 133 {
		t.Errorf("Expected exactly 133 resources, got %d", len(resources))
	}
	
	// Check that all resources have valid tokens
	for resource, token := range resources {
		if token == "" {
			t.Errorf("Resource '%s' has empty token", resource)
		}
	}
}

func TestDataSourceCompleteness(t *testing.T) {
	provider := NewStandaloneProvider()
	dataSources := provider.GetDataSources()
	
	// Check that we have exactly 23 data sources
	if len(dataSources) != 23 {
		t.Errorf("Expected exactly 23 data sources, got %d", len(dataSources))
	}
	
	// Check that all data sources have valid tokens
	for dataSource, token := range dataSources {
		if token == "" {
			t.Errorf("Data source '%s' has empty token", dataSource)
		}
	}
}

func TestConfigValidationEdgeCases(t *testing.T) {
	provider := NewStandaloneProvider()
	
	// Test empty config
	emptyConfig := map[string]interface{}{}
	err := provider.ValidateConfig(emptyConfig)
	if err == nil {
		t.Error("Empty config should return error")
	}
	
	// Test nil config
	err = provider.ValidateConfig(nil)
	if err == nil {
		t.Error("Nil config should return error")
	}
	
	// Test config with extra fields
	extraConfig := map[string]interface{}{
		"controller_ip": "demo-controller.aviatrix.com",
		"username":      "demo-user",
		"password":      "demo-password",
		"extra_field":   "extra_value",
	}
	err = provider.ValidateConfig(extraConfig)
	if err != nil {
		t.Errorf("Config with extra fields should be valid, got: %v", err)
	}
}

func TestProviderConsistency(t *testing.T) {
	provider1 := NewStandaloneProvider()
	provider2 := NewStandaloneProvider()
	
	// Test that multiple instances are consistent
	if provider1.GetResourceCount() != provider2.GetResourceCount() {
		t.Error("Resource count should be consistent across instances")
	}
	
	if provider1.GetDataSourceCount() != provider2.GetDataSourceCount() {
		t.Error("Data source count should be consistent across instances")
	}
	
	// Test that resources are the same
	resources1 := provider1.GetResources()
	resources2 := provider2.GetResources()
	
	for resource, token1 := range resources1 {
		if token2, exists := resources2[resource]; !exists {
			t.Errorf("Resource '%s' missing in second instance", resource)
		} else if token1 != token2 {
			t.Errorf("Resource '%s' token mismatch: %s vs %s", resource, token1, token2)
		}
	}
}

func TestResourceUniqueness(t *testing.T) {
	provider := NewStandaloneProvider()
	resources := provider.GetResources()
	
	// Check for duplicate resources
	seen := make(map[string]bool)
	for resource := range resources {
		if seen[resource] {
			t.Errorf("Duplicate resource found: %s", resource)
		}
		seen[resource] = true
	}
}

func TestDataSourceUniqueness(t *testing.T) {
	provider := NewStandaloneProvider()
	dataSources := provider.GetDataSources()
	
	// Check for duplicate data sources
	seen := make(map[string]bool)
	for dataSource := range dataSources {
		if seen[dataSource] {
			t.Errorf("Duplicate data source found: %s", dataSource)
		}
		seen[dataSource] = true
	}
}

func TestConfigFieldTypes(t *testing.T) {
	provider := NewStandaloneProvider()
	
	// Test with different field types
	config := map[string]interface{}{
		"controller_ip": "demo-controller.aviatrix.com",
		"username":      "demo-user",
		"password":      "demo-password",
	}
	
	err := provider.ValidateConfig(config)
	if err != nil {
		t.Errorf("Config with string values should be valid, got: %v", err)
	}
	
	// Test with int values (should still work)
	configInt := map[string]interface{}{
		"controller_ip": 123,
		"username":      "demo-user",
		"password":      "demo-password",
	}
	
	err = provider.ValidateConfig(configInt)
	if err != nil {
		t.Errorf("Config with int values should be valid, got: %v", err)
	}
}
