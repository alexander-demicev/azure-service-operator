// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_RedisEnterprise_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_STATUS_ARM, RedisEnterprise_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_STATUS_ARM runs a test to see if a specific instance of RedisEnterprise_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_STATUS_ARM(subject RedisEnterprise_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of RedisEnterprise_STATUS_ARM instances for property testing - lazily instantiated by
// RedisEnterprise_STATUS_ARMGenerator()
var redisEnterprise_STATUS_ARMGenerator gopter.Gen

// RedisEnterprise_STATUS_ARMGenerator returns a generator of RedisEnterprise_STATUS_ARM instances for property testing.
// We first initialize redisEnterprise_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_STATUS_ARMGenerator() gopter.Gen {
	if redisEnterprise_STATUS_ARMGenerator != nil {
		return redisEnterprise_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS_ARM(generators)
	redisEnterprise_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS_ARM(generators)
	redisEnterprise_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_STATUS_ARM{}), generators)

	return redisEnterprise_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ClusterProperties_STATUS_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUS_ARMGenerator())
}

func Test_ClusterProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterProperties_STATUS_ARM, ClusterProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterProperties_STATUS_ARM runs a test to see if a specific instance of ClusterProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterProperties_STATUS_ARM(subject ClusterProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ClusterProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ClusterProperties_STATUS_ARMGenerator()
var clusterProperties_STATUS_ARMGenerator gopter.Gen

// ClusterProperties_STATUS_ARMGenerator returns a generator of ClusterProperties_STATUS_ARM instances for property testing.
// We first initialize clusterProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ClusterProperties_STATUS_ARMGenerator() gopter.Gen {
	if clusterProperties_STATUS_ARMGenerator != nil {
		return clusterProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterProperties_STATUS_ARM(generators)
	clusterProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForClusterProperties_STATUS_ARM(generators)
	clusterProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ClusterProperties_STATUS_ARM{}), generators)

	return clusterProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForClusterProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(ClusterProperties_MinimumTlsVersion_STATUS_10, ClusterProperties_MinimumTlsVersion_STATUS_11, ClusterProperties_MinimumTlsVersion_STATUS_12))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ResourceState_STATUS_CreateFailed,
		ResourceState_STATUS_Creating,
		ResourceState_STATUS_DeleteFailed,
		ResourceState_STATUS_Deleting,
		ResourceState_STATUS_DisableFailed,
		ResourceState_STATUS_Disabled,
		ResourceState_STATUS_Disabling,
		ResourceState_STATUS_EnableFailed,
		ResourceState_STATUS_Enabling,
		ResourceState_STATUS_Running,
		ResourceState_STATUS_UpdateFailed,
		ResourceState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForClusterProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForClusterProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUS_ARMGenerator())
}

func Test_Sku_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS_ARM, Sku_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS_ARM runs a test to see if a specific instance of Sku_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS_ARM(subject Sku_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku_STATUS_ARM instances for property testing - lazily instantiated by Sku_STATUS_ARMGenerator()
var sku_STATUS_ARMGenerator gopter.Gen

// Sku_STATUS_ARMGenerator returns a generator of Sku_STATUS_ARM instances for property testing.
func Sku_STATUS_ARMGenerator() gopter.Gen {
	if sku_STATUS_ARMGenerator != nil {
		return sku_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS_ARM(generators)
	sku_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS_ARM{}), generators)

	return sku_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_STATUS_EnterpriseFlash_F1500,
		Sku_Name_STATUS_EnterpriseFlash_F300,
		Sku_Name_STATUS_EnterpriseFlash_F700,
		Sku_Name_STATUS_Enterprise_E10,
		Sku_Name_STATUS_Enterprise_E100,
		Sku_Name_STATUS_Enterprise_E20,
		Sku_Name_STATUS_Enterprise_E50))
}

func Test_PrivateEndpointConnection_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM, PrivateEndpointConnection_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS_ARM(subject PrivateEndpointConnection_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PrivateEndpointConnection_STATUS_ARM instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUS_ARMGenerator()
var privateEndpointConnection_STATUS_ARMGenerator gopter.Gen

// PrivateEndpointConnection_STATUS_ARMGenerator returns a generator of PrivateEndpointConnection_STATUS_ARM instances for property testing.
func PrivateEndpointConnection_STATUS_ARMGenerator() gopter.Gen {
	if privateEndpointConnection_STATUS_ARMGenerator != nil {
		return privateEndpointConnection_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(generators)
	privateEndpointConnection_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_ARM{}), generators)

	return privateEndpointConnection_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}