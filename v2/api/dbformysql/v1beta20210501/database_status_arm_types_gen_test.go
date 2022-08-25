// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_Database_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_STATUSARM, Database_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_STATUSARM runs a test to see if a specific instance of Database_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_STATUSARM(subject Database_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUSARM
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

// Generator of Database_STATUSARM instances for property testing - lazily instantiated by Database_STATUSARMGenerator()
var database_STATUSARMGenerator gopter.Gen

// Database_STATUSARMGenerator returns a generator of Database_STATUSARM instances for property testing.
// We first initialize database_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Database_STATUSARMGenerator() gopter.Gen {
	if database_STATUSARMGenerator != nil {
		return database_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_STATUSARM(generators)
	database_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Database_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForDatabase_STATUSARM(generators)
	database_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Database_STATUSARM{}), generators)

	return database_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabase_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_DatabaseProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_STATUSARM, DatabaseProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_STATUSARM runs a test to see if a specific instance of DatabaseProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_STATUSARM(subject DatabaseProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_STATUSARM
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

// Generator of DatabaseProperties_STATUSARM instances for property testing - lazily instantiated by
// DatabaseProperties_STATUSARMGenerator()
var databaseProperties_STATUSARMGenerator gopter.Gen

// DatabaseProperties_STATUSARMGenerator returns a generator of DatabaseProperties_STATUSARM instances for property testing.
func DatabaseProperties_STATUSARMGenerator() gopter.Gen {
	if databaseProperties_STATUSARMGenerator != nil {
		return databaseProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM(generators)
	databaseProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUSARM{}), generators)

	return databaseProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUSARM, SystemData_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUSARM runs a test to see if a specific instance of SystemData_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUSARM(subject SystemData_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUSARM
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

// Generator of SystemData_STATUSARM instances for property testing - lazily instantiated by
// SystemData_STATUSARMGenerator()
var systemData_STATUSARMGenerator gopter.Gen

// SystemData_STATUSARMGenerator returns a generator of SystemData_STATUSARM instances for property testing.
func SystemData_STATUSARMGenerator() gopter.Gen {
	if systemData_STATUSARMGenerator != nil {
		return systemData_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUSARM(generators)
	systemData_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUSARM{}), generators)

	return systemData_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUSARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_STATUS_CreatedByType_Application,
		SystemData_STATUS_CreatedByType_Key,
		SystemData_STATUS_CreatedByType_ManagedIdentity,
		SystemData_STATUS_CreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_STATUS_LastModifiedByType_Application,
		SystemData_STATUS_LastModifiedByType_Key,
		SystemData_STATUS_LastModifiedByType_ManagedIdentity,
		SystemData_STATUS_LastModifiedByType_User))
}