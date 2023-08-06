// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

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

func Test_Factory_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Factory_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactory_STATUS_ARM, Factory_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactory_STATUS_ARM runs a test to see if a specific instance of Factory_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactory_STATUS_ARM(subject Factory_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Factory_STATUS_ARM
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

// Generator of Factory_STATUS_ARM instances for property testing - lazily instantiated by Factory_STATUS_ARMGenerator()
var factory_STATUS_ARMGenerator gopter.Gen

// Factory_STATUS_ARMGenerator returns a generator of Factory_STATUS_ARM instances for property testing.
// We first initialize factory_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Factory_STATUS_ARMGenerator() gopter.Gen {
	if factory_STATUS_ARMGenerator != nil {
		return factory_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactory_STATUS_ARM(generators)
	factory_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Factory_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactory_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFactory_STATUS_ARM(generators)
	factory_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Factory_STATUS_ARM{}), generators)

	return factory_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactory_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactory_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFactory_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactory_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(FactoryIdentity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(FactoryProperties_STATUS_ARMGenerator())
}

func Test_FactoryIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryIdentity_STATUS_ARM, FactoryIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryIdentity_STATUS_ARM runs a test to see if a specific instance of FactoryIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryIdentity_STATUS_ARM(subject FactoryIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryIdentity_STATUS_ARM
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

// Generator of FactoryIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// FactoryIdentity_STATUS_ARMGenerator()
var factoryIdentity_STATUS_ARMGenerator gopter.Gen

// FactoryIdentity_STATUS_ARMGenerator returns a generator of FactoryIdentity_STATUS_ARM instances for property testing.
func FactoryIdentity_STATUS_ARMGenerator() gopter.Gen {
	if factoryIdentity_STATUS_ARMGenerator != nil {
		return factoryIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryIdentity_STATUS_ARM(generators)
	factoryIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryIdentity_STATUS_ARM{}), generators)

	return factoryIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(FactoryIdentity_Type_STATUS_SystemAssigned, FactoryIdentity_Type_STATUS_SystemAssignedUserAssigned, FactoryIdentity_Type_STATUS_UserAssigned))
}

func Test_FactoryProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryProperties_STATUS_ARM, FactoryProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryProperties_STATUS_ARM runs a test to see if a specific instance of FactoryProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryProperties_STATUS_ARM(subject FactoryProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryProperties_STATUS_ARM
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

// Generator of FactoryProperties_STATUS_ARM instances for property testing - lazily instantiated by
// FactoryProperties_STATUS_ARMGenerator()
var factoryProperties_STATUS_ARMGenerator gopter.Gen

// FactoryProperties_STATUS_ARMGenerator returns a generator of FactoryProperties_STATUS_ARM instances for property testing.
// We first initialize factoryProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FactoryProperties_STATUS_ARMGenerator() gopter.Gen {
	if factoryProperties_STATUS_ARMGenerator != nil {
		return factoryProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryProperties_STATUS_ARM(generators)
	factoryProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFactoryProperties_STATUS_ARM(generators)
	factoryProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryProperties_STATUS_ARM{}), generators)

	return factoryProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreateTime"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(FactoryProperties_PublicNetworkAccess_STATUS_Disabled, FactoryProperties_PublicNetworkAccess_STATUS_Enabled))
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFactoryProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionConfiguration_STATUS_ARMGenerator())
	gens["GlobalParameters"] = gen.MapOf(gen.AlphaString(), GlobalParameterSpecification_STATUS_ARMGenerator())
	gens["PurviewConfiguration"] = gen.PtrOf(PurviewConfiguration_STATUS_ARMGenerator())
	gens["RepoConfiguration"] = gen.PtrOf(FactoryRepoConfiguration_STATUS_ARMGenerator())
}

func Test_EncryptionConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionConfiguration_STATUS_ARM, EncryptionConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionConfiguration_STATUS_ARM runs a test to see if a specific instance of EncryptionConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionConfiguration_STATUS_ARM(subject EncryptionConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionConfiguration_STATUS_ARM
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

// Generator of EncryptionConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// EncryptionConfiguration_STATUS_ARMGenerator()
var encryptionConfiguration_STATUS_ARMGenerator gopter.Gen

// EncryptionConfiguration_STATUS_ARMGenerator returns a generator of EncryptionConfiguration_STATUS_ARM instances for property testing.
// We first initialize encryptionConfiguration_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if encryptionConfiguration_STATUS_ARMGenerator != nil {
		return encryptionConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM(generators)
	encryptionConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionConfiguration_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM(generators)
	encryptionConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionConfiguration_STATUS_ARM{}), generators)

	return encryptionConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
	gens["VaultBaseUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(CMKIdentityDefinition_STATUS_ARMGenerator())
}

func Test_FactoryRepoConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryRepoConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryRepoConfiguration_STATUS_ARM, FactoryRepoConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryRepoConfiguration_STATUS_ARM runs a test to see if a specific instance of FactoryRepoConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryRepoConfiguration_STATUS_ARM(subject FactoryRepoConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryRepoConfiguration_STATUS_ARM
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

// Generator of FactoryRepoConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// FactoryRepoConfiguration_STATUS_ARMGenerator()
var factoryRepoConfiguration_STATUS_ARMGenerator gopter.Gen

// FactoryRepoConfiguration_STATUS_ARMGenerator returns a generator of FactoryRepoConfiguration_STATUS_ARM instances for property testing.
func FactoryRepoConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if factoryRepoConfiguration_STATUS_ARMGenerator != nil {
		return factoryRepoConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_STATUS_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(FactoryRepoConfiguration_STATUS_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	factoryRepoConfiguration_STATUS_ARMGenerator = gen.OneGenOf(gens...)

	return factoryRepoConfiguration_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryRepoConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["FactoryGitHub"] = FactoryGitHubConfiguration_STATUS_ARMGenerator().Map(func(it FactoryGitHubConfiguration_STATUS_ARM) *FactoryGitHubConfiguration_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["FactoryVSTS"] = FactoryVSTSConfiguration_STATUS_ARMGenerator().Map(func(it FactoryVSTSConfiguration_STATUS_ARM) *FactoryVSTSConfiguration_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_GlobalParameterSpecification_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GlobalParameterSpecification_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGlobalParameterSpecification_STATUS_ARM, GlobalParameterSpecification_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGlobalParameterSpecification_STATUS_ARM runs a test to see if a specific instance of GlobalParameterSpecification_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForGlobalParameterSpecification_STATUS_ARM(subject GlobalParameterSpecification_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GlobalParameterSpecification_STATUS_ARM
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

// Generator of GlobalParameterSpecification_STATUS_ARM instances for property testing - lazily instantiated by
// GlobalParameterSpecification_STATUS_ARMGenerator()
var globalParameterSpecification_STATUS_ARMGenerator gopter.Gen

// GlobalParameterSpecification_STATUS_ARMGenerator returns a generator of GlobalParameterSpecification_STATUS_ARM instances for property testing.
func GlobalParameterSpecification_STATUS_ARMGenerator() gopter.Gen {
	if globalParameterSpecification_STATUS_ARMGenerator != nil {
		return globalParameterSpecification_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGlobalParameterSpecification_STATUS_ARM(generators)
	globalParameterSpecification_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(GlobalParameterSpecification_STATUS_ARM{}), generators)

	return globalParameterSpecification_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForGlobalParameterSpecification_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGlobalParameterSpecification_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		GlobalParameterSpecification_Type_STATUS_Array,
		GlobalParameterSpecification_Type_STATUS_Bool,
		GlobalParameterSpecification_Type_STATUS_Float,
		GlobalParameterSpecification_Type_STATUS_Int,
		GlobalParameterSpecification_Type_STATUS_Object,
		GlobalParameterSpecification_Type_STATUS_String))
}

func Test_PurviewConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PurviewConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPurviewConfiguration_STATUS_ARM, PurviewConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPurviewConfiguration_STATUS_ARM runs a test to see if a specific instance of PurviewConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPurviewConfiguration_STATUS_ARM(subject PurviewConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PurviewConfiguration_STATUS_ARM
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

// Generator of PurviewConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// PurviewConfiguration_STATUS_ARMGenerator()
var purviewConfiguration_STATUS_ARMGenerator gopter.Gen

// PurviewConfiguration_STATUS_ARMGenerator returns a generator of PurviewConfiguration_STATUS_ARM instances for property testing.
func PurviewConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if purviewConfiguration_STATUS_ARMGenerator != nil {
		return purviewConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPurviewConfiguration_STATUS_ARM(generators)
	purviewConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PurviewConfiguration_STATUS_ARM{}), generators)

	return purviewConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPurviewConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPurviewConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PurviewResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_CMKIdentityDefinition_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CMKIdentityDefinition_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCMKIdentityDefinition_STATUS_ARM, CMKIdentityDefinition_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCMKIdentityDefinition_STATUS_ARM runs a test to see if a specific instance of CMKIdentityDefinition_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCMKIdentityDefinition_STATUS_ARM(subject CMKIdentityDefinition_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CMKIdentityDefinition_STATUS_ARM
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

// Generator of CMKIdentityDefinition_STATUS_ARM instances for property testing - lazily instantiated by
// CMKIdentityDefinition_STATUS_ARMGenerator()
var cmkIdentityDefinition_STATUS_ARMGenerator gopter.Gen

// CMKIdentityDefinition_STATUS_ARMGenerator returns a generator of CMKIdentityDefinition_STATUS_ARM instances for property testing.
func CMKIdentityDefinition_STATUS_ARMGenerator() gopter.Gen {
	if cmkIdentityDefinition_STATUS_ARMGenerator != nil {
		return cmkIdentityDefinition_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCMKIdentityDefinition_STATUS_ARM(generators)
	cmkIdentityDefinition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CMKIdentityDefinition_STATUS_ARM{}), generators)

	return cmkIdentityDefinition_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCMKIdentityDefinition_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCMKIdentityDefinition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_FactoryGitHubConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryGitHubConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryGitHubConfiguration_STATUS_ARM, FactoryGitHubConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryGitHubConfiguration_STATUS_ARM runs a test to see if a specific instance of FactoryGitHubConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryGitHubConfiguration_STATUS_ARM(subject FactoryGitHubConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryGitHubConfiguration_STATUS_ARM
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

// Generator of FactoryGitHubConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// FactoryGitHubConfiguration_STATUS_ARMGenerator()
var factoryGitHubConfiguration_STATUS_ARMGenerator gopter.Gen

// FactoryGitHubConfiguration_STATUS_ARMGenerator returns a generator of FactoryGitHubConfiguration_STATUS_ARM instances for property testing.
// We first initialize factoryGitHubConfiguration_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FactoryGitHubConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if factoryGitHubConfiguration_STATUS_ARMGenerator != nil {
		return factoryGitHubConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM(generators)
	factoryGitHubConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryGitHubConfiguration_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM(generators)
	factoryGitHubConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryGitHubConfiguration_STATUS_ARM{}), generators)

	return factoryGitHubConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AccountName"] = gen.PtrOf(gen.AlphaString())
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["CollaborationBranch"] = gen.PtrOf(gen.AlphaString())
	gens["DisablePublish"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["LastCommitId"] = gen.PtrOf(gen.AlphaString())
	gens["RepositoryName"] = gen.PtrOf(gen.AlphaString())
	gens["RootFolder"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(FactoryGitHubConfiguration_Type_STATUS_FactoryGitHubConfiguration)
}

// AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFactoryGitHubConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientSecret"] = gen.PtrOf(GitHubClientSecret_STATUS_ARMGenerator())
}

func Test_FactoryVSTSConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FactoryVSTSConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFactoryVSTSConfiguration_STATUS_ARM, FactoryVSTSConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFactoryVSTSConfiguration_STATUS_ARM runs a test to see if a specific instance of FactoryVSTSConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFactoryVSTSConfiguration_STATUS_ARM(subject FactoryVSTSConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FactoryVSTSConfiguration_STATUS_ARM
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

// Generator of FactoryVSTSConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// FactoryVSTSConfiguration_STATUS_ARMGenerator()
var factoryVSTSConfiguration_STATUS_ARMGenerator gopter.Gen

// FactoryVSTSConfiguration_STATUS_ARMGenerator returns a generator of FactoryVSTSConfiguration_STATUS_ARM instances for property testing.
func FactoryVSTSConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if factoryVSTSConfiguration_STATUS_ARMGenerator != nil {
		return factoryVSTSConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_STATUS_ARM(generators)
	factoryVSTSConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FactoryVSTSConfiguration_STATUS_ARM{}), generators)

	return factoryVSTSConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFactoryVSTSConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AccountName"] = gen.PtrOf(gen.AlphaString())
	gens["CollaborationBranch"] = gen.PtrOf(gen.AlphaString())
	gens["DisablePublish"] = gen.PtrOf(gen.Bool())
	gens["LastCommitId"] = gen.PtrOf(gen.AlphaString())
	gens["ProjectName"] = gen.PtrOf(gen.AlphaString())
	gens["RepositoryName"] = gen.PtrOf(gen.AlphaString())
	gens["RootFolder"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(FactoryVSTSConfiguration_Type_STATUS_FactoryVSTSConfiguration)
}

func Test_GitHubClientSecret_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GitHubClientSecret_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGitHubClientSecret_STATUS_ARM, GitHubClientSecret_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGitHubClientSecret_STATUS_ARM runs a test to see if a specific instance of GitHubClientSecret_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForGitHubClientSecret_STATUS_ARM(subject GitHubClientSecret_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GitHubClientSecret_STATUS_ARM
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

// Generator of GitHubClientSecret_STATUS_ARM instances for property testing - lazily instantiated by
// GitHubClientSecret_STATUS_ARMGenerator()
var gitHubClientSecret_STATUS_ARMGenerator gopter.Gen

// GitHubClientSecret_STATUS_ARMGenerator returns a generator of GitHubClientSecret_STATUS_ARM instances for property testing.
func GitHubClientSecret_STATUS_ARMGenerator() gopter.Gen {
	if gitHubClientSecret_STATUS_ARMGenerator != nil {
		return gitHubClientSecret_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitHubClientSecret_STATUS_ARM(generators)
	gitHubClientSecret_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(GitHubClientSecret_STATUS_ARM{}), generators)

	return gitHubClientSecret_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForGitHubClientSecret_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGitHubClientSecret_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ByoaSecretAkvUrl"] = gen.PtrOf(gen.AlphaString())
	gens["ByoaSecretName"] = gen.PtrOf(gen.AlphaString())
}