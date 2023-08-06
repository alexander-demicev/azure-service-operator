// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

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

func Test_NetworkSecurityGroups_SecurityRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SecurityRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec_ARM, NetworkSecurityGroups_SecurityRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec_ARM runs a test to see if a specific instance of NetworkSecurityGroups_SecurityRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_Spec_ARM(subject NetworkSecurityGroups_SecurityRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SecurityRule_Spec_ARM
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

// Generator of NetworkSecurityGroups_SecurityRule_Spec_ARM instances for property testing - lazily instantiated by
// NetworkSecurityGroups_SecurityRule_Spec_ARMGenerator()
var networkSecurityGroups_SecurityRule_Spec_ARMGenerator gopter.Gen

// NetworkSecurityGroups_SecurityRule_Spec_ARMGenerator returns a generator of NetworkSecurityGroups_SecurityRule_Spec_ARM instances for property testing.
// We first initialize networkSecurityGroups_SecurityRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroups_SecurityRule_Spec_ARMGenerator() gopter.Gen {
	if networkSecurityGroups_SecurityRule_Spec_ARMGenerator != nil {
		return networkSecurityGroups_SecurityRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM(generators)
	networkSecurityGroups_SecurityRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM(generators)
	networkSecurityGroups_SecurityRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_Spec_ARM{}), generators)

	return networkSecurityGroups_SecurityRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator())
}

func Test_SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM, SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM runs a test to see if a specific instance of SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(subject SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM
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

// Generator of SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for
// property testing - lazily instantiated by
// SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()
var securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator gopter.Gen

// SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator returns a generator of SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for property testing.
// We first initialize securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator != nil {
		return securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(generators)
	securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(generators)
	securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM{}), generators)

	return securityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_Allow, SecurityRuleAccess_Deny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_Inbound, SecurityRuleDirection_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah,
		SecurityRulePropertiesFormat_Protocol_Esp,
		SecurityRulePropertiesFormat_Protocol_Icmp,
		SecurityRulePropertiesFormat_Protocol_Star,
		SecurityRulePropertiesFormat_Protocol_Tcp,
		SecurityRulePropertiesFormat_Protocol_Udp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator())
}

func Test_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM, ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM runs a test to see if a specific instance of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(subject ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM
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

// Generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator()
var applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator gopter.Gen

// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator returns a generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM instances for property testing.
func ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator() gopter.Gen {
	if applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator != nil {
		return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(generators)
	applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM{}), generators)

	return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}