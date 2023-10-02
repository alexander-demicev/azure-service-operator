// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401storage

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

func Test_TrafficManagerProfilesNestedEndpoint_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesNestedEndpoint via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint, TrafficManagerProfilesNestedEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint runs a test to see if a specific instance of TrafficManagerProfilesNestedEndpoint round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesNestedEndpoint(subject TrafficManagerProfilesNestedEndpoint) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesNestedEndpoint
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

// Generator of TrafficManagerProfilesNestedEndpoint instances for property testing - lazily instantiated by
// TrafficManagerProfilesNestedEndpointGenerator()
var trafficManagerProfilesNestedEndpointGenerator gopter.Gen

// TrafficManagerProfilesNestedEndpointGenerator returns a generator of TrafficManagerProfilesNestedEndpoint instances for property testing.
func TrafficManagerProfilesNestedEndpointGenerator() gopter.Gen {
	if trafficManagerProfilesNestedEndpointGenerator != nil {
		return trafficManagerProfilesNestedEndpointGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint(generators)
	trafficManagerProfilesNestedEndpointGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesNestedEndpoint{}), generators)

	return trafficManagerProfilesNestedEndpointGenerator
}

// AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficManagerProfilesNestedEndpoint(gens map[string]gopter.Gen) {
	gens["Spec"] = Trafficmanagerprofiles_NestedEndpoint_SpecGenerator()
	gens["Status"] = Trafficmanagerprofiles_NestedEndpoint_STATUSGenerator()
}

func Test_Trafficmanagerprofiles_NestedEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Trafficmanagerprofiles_NestedEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_Spec, Trafficmanagerprofiles_NestedEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_Spec runs a test to see if a specific instance of Trafficmanagerprofiles_NestedEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_Spec(subject Trafficmanagerprofiles_NestedEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Trafficmanagerprofiles_NestedEndpoint_Spec
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

// Generator of Trafficmanagerprofiles_NestedEndpoint_Spec instances for property testing - lazily instantiated by
// Trafficmanagerprofiles_NestedEndpoint_SpecGenerator()
var trafficmanagerprofiles_NestedEndpoint_SpecGenerator gopter.Gen

// Trafficmanagerprofiles_NestedEndpoint_SpecGenerator returns a generator of Trafficmanagerprofiles_NestedEndpoint_Spec instances for property testing.
// We first initialize trafficmanagerprofiles_NestedEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Trafficmanagerprofiles_NestedEndpoint_SpecGenerator() gopter.Gen {
	if trafficmanagerprofiles_NestedEndpoint_SpecGenerator != nil {
		return trafficmanagerprofiles_NestedEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec(generators)
	trafficmanagerprofiles_NestedEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_NestedEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec(generators)
	trafficmanagerprofiles_NestedEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_NestedEndpoint_Spec{}), generators)

	return trafficmanagerprofiles_NestedEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeadersGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_SubnetsGenerator())
}

func Test_Trafficmanagerprofiles_NestedEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Trafficmanagerprofiles_NestedEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_STATUS, Trafficmanagerprofiles_NestedEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_STATUS runs a test to see if a specific instance of Trafficmanagerprofiles_NestedEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficmanagerprofiles_NestedEndpoint_STATUS(subject Trafficmanagerprofiles_NestedEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Trafficmanagerprofiles_NestedEndpoint_STATUS
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

// Generator of Trafficmanagerprofiles_NestedEndpoint_STATUS instances for property testing - lazily instantiated by
// Trafficmanagerprofiles_NestedEndpoint_STATUSGenerator()
var trafficmanagerprofiles_NestedEndpoint_STATUSGenerator gopter.Gen

// Trafficmanagerprofiles_NestedEndpoint_STATUSGenerator returns a generator of Trafficmanagerprofiles_NestedEndpoint_STATUS instances for property testing.
// We first initialize trafficmanagerprofiles_NestedEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Trafficmanagerprofiles_NestedEndpoint_STATUSGenerator() gopter.Gen {
	if trafficmanagerprofiles_NestedEndpoint_STATUSGenerator != nil {
		return trafficmanagerprofiles_NestedEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS(generators)
	trafficmanagerprofiles_NestedEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_NestedEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS(generators)
	trafficmanagerprofiles_NestedEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_NestedEndpoint_STATUS{}), generators)

	return trafficmanagerprofiles_NestedEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["TargetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_NestedEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeaders_STATUSGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_Subnets_STATUSGenerator())
}