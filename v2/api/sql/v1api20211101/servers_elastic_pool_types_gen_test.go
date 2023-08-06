// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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

func Test_ServersElasticPool_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersElasticPool to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersElasticPool, ServersElasticPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersElasticPool tests if a specific instance of ServersElasticPool round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersElasticPool(subject ServersElasticPool) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersElasticPool
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersElasticPool
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersElasticPool_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersElasticPool to ServersElasticPool via AssignProperties_To_ServersElasticPool & AssignProperties_From_ServersElasticPool returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersElasticPool, ServersElasticPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersElasticPool tests if a specific instance of ServersElasticPool can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServersElasticPool(subject ServersElasticPool) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersElasticPool
	err := copied.AssignProperties_To_ServersElasticPool(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersElasticPool
	err = actual.AssignProperties_From_ServersElasticPool(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersElasticPool_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersElasticPool via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersElasticPool, ServersElasticPoolGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersElasticPool runs a test to see if a specific instance of ServersElasticPool round trips to JSON and back losslessly
func RunJSONSerializationTestForServersElasticPool(subject ServersElasticPool) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersElasticPool
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

// Generator of ServersElasticPool instances for property testing - lazily instantiated by ServersElasticPoolGenerator()
var serversElasticPoolGenerator gopter.Gen

// ServersElasticPoolGenerator returns a generator of ServersElasticPool instances for property testing.
func ServersElasticPoolGenerator() gopter.Gen {
	if serversElasticPoolGenerator != nil {
		return serversElasticPoolGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersElasticPool(generators)
	serversElasticPoolGenerator = gen.Struct(reflect.TypeOf(ServersElasticPool{}), generators)

	return serversElasticPoolGenerator
}

// AddRelatedPropertyGeneratorsForServersElasticPool is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersElasticPool(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_ElasticPool_SpecGenerator()
	gens["Status"] = Servers_ElasticPool_STATUSGenerator()
}

func Test_Servers_ElasticPool_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_ElasticPool_Spec to Servers_ElasticPool_Spec via AssignProperties_To_Servers_ElasticPool_Spec & AssignProperties_From_Servers_ElasticPool_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_ElasticPool_Spec, Servers_ElasticPool_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_ElasticPool_Spec tests if a specific instance of Servers_ElasticPool_Spec can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_ElasticPool_Spec(subject Servers_ElasticPool_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_ElasticPool_Spec
	err := copied.AssignProperties_To_Servers_ElasticPool_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_ElasticPool_Spec
	err = actual.AssignProperties_From_Servers_ElasticPool_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_ElasticPool_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ElasticPool_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ElasticPool_Spec, Servers_ElasticPool_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ElasticPool_Spec runs a test to see if a specific instance of Servers_ElasticPool_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ElasticPool_Spec(subject Servers_ElasticPool_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ElasticPool_Spec
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

// Generator of Servers_ElasticPool_Spec instances for property testing - lazily instantiated by
// Servers_ElasticPool_SpecGenerator()
var servers_ElasticPool_SpecGenerator gopter.Gen

// Servers_ElasticPool_SpecGenerator returns a generator of Servers_ElasticPool_Spec instances for property testing.
// We first initialize servers_ElasticPool_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_ElasticPool_SpecGenerator() gopter.Gen {
	if servers_ElasticPool_SpecGenerator != nil {
		return servers_ElasticPool_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	servers_ElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec(generators)
	servers_ElasticPool_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec{}), generators)

	return servers_ElasticPool_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["LicenseType"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_LicenseType_BasePrice, ElasticPoolProperties_LicenseType_LicenseIncluded))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettingsGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Servers_ElasticPool_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_ElasticPool_STATUS to Servers_ElasticPool_STATUS via AssignProperties_To_Servers_ElasticPool_STATUS & AssignProperties_From_Servers_ElasticPool_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_ElasticPool_STATUS, Servers_ElasticPool_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_ElasticPool_STATUS tests if a specific instance of Servers_ElasticPool_STATUS can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_ElasticPool_STATUS(subject Servers_ElasticPool_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_ElasticPool_STATUS
	err := copied.AssignProperties_To_Servers_ElasticPool_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_ElasticPool_STATUS
	err = actual.AssignProperties_From_Servers_ElasticPool_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_ElasticPool_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ElasticPool_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ElasticPool_STATUS, Servers_ElasticPool_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ElasticPool_STATUS runs a test to see if a specific instance of Servers_ElasticPool_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ElasticPool_STATUS(subject Servers_ElasticPool_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ElasticPool_STATUS
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

// Generator of Servers_ElasticPool_STATUS instances for property testing - lazily instantiated by
// Servers_ElasticPool_STATUSGenerator()
var servers_ElasticPool_STATUSGenerator gopter.Gen

// Servers_ElasticPool_STATUSGenerator returns a generator of Servers_ElasticPool_STATUS instances for property testing.
// We first initialize servers_ElasticPool_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_ElasticPool_STATUSGenerator() gopter.Gen {
	if servers_ElasticPool_STATUSGenerator != nil {
		return servers_ElasticPool_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	servers_ElasticPool_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS(generators)
	servers_ElasticPool_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_STATUS{}), generators)

	return servers_ElasticPool_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ElasticPool_STATUS(gens map[string]gopter.Gen) {
	gens["CreationDate"] = gen.PtrOf(gen.AlphaString())
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["LicenseType"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_LicenseType_STATUS_BasePrice, ElasticPoolProperties_LicenseType_STATUS_LicenseIncluded))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_State_STATUS_Creating, ElasticPoolProperties_State_STATUS_Disabled, ElasticPoolProperties_State_STATUS_Ready))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_ElasticPool_STATUS(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettings_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_ElasticPoolPerDatabaseSettings_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ElasticPoolPerDatabaseSettings to ElasticPoolPerDatabaseSettings via AssignProperties_To_ElasticPoolPerDatabaseSettings & AssignProperties_From_ElasticPoolPerDatabaseSettings returns original",
		prop.ForAll(RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings, ElasticPoolPerDatabaseSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings tests if a specific instance of ElasticPoolPerDatabaseSettings can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings(subject ElasticPoolPerDatabaseSettings) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ElasticPoolPerDatabaseSettings
	err := copied.AssignProperties_To_ElasticPoolPerDatabaseSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ElasticPoolPerDatabaseSettings
	err = actual.AssignProperties_From_ElasticPoolPerDatabaseSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ElasticPoolPerDatabaseSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings, ElasticPoolPerDatabaseSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings(subject ElasticPoolPerDatabaseSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings
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

// Generator of ElasticPoolPerDatabaseSettings instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettingsGenerator()
var elasticPoolPerDatabaseSettingsGenerator gopter.Gen

// ElasticPoolPerDatabaseSettingsGenerator returns a generator of ElasticPoolPerDatabaseSettings instances for property testing.
func ElasticPoolPerDatabaseSettingsGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettingsGenerator != nil {
		return elasticPoolPerDatabaseSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings(generators)
	elasticPoolPerDatabaseSettingsGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings{}), generators)

	return elasticPoolPerDatabaseSettingsGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}

func Test_ElasticPoolPerDatabaseSettings_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ElasticPoolPerDatabaseSettings_STATUS to ElasticPoolPerDatabaseSettings_STATUS via AssignProperties_To_ElasticPoolPerDatabaseSettings_STATUS & AssignProperties_From_ElasticPoolPerDatabaseSettings_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings_STATUS, ElasticPoolPerDatabaseSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings_STATUS tests if a specific instance of ElasticPoolPerDatabaseSettings_STATUS can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForElasticPoolPerDatabaseSettings_STATUS(subject ElasticPoolPerDatabaseSettings_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ElasticPoolPerDatabaseSettings_STATUS
	err := copied.AssignProperties_To_ElasticPoolPerDatabaseSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ElasticPoolPerDatabaseSettings_STATUS
	err = actual.AssignProperties_From_ElasticPoolPerDatabaseSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ElasticPoolPerDatabaseSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS, ElasticPoolPerDatabaseSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings_STATUS(subject ElasticPoolPerDatabaseSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings_STATUS
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

// Generator of ElasticPoolPerDatabaseSettings_STATUS instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettings_STATUSGenerator()
var elasticPoolPerDatabaseSettings_STATUSGenerator gopter.Gen

// ElasticPoolPerDatabaseSettings_STATUSGenerator returns a generator of ElasticPoolPerDatabaseSettings_STATUS instances for property testing.
func ElasticPoolPerDatabaseSettings_STATUSGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettings_STATUSGenerator != nil {
		return elasticPoolPerDatabaseSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS(generators)
	elasticPoolPerDatabaseSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings_STATUS{}), generators)

	return elasticPoolPerDatabaseSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}