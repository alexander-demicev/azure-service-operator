/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"
)

type StorageConversionPropertyTestCase struct {
	name        string
	current     astmodel.TypeDefinition
	other       astmodel.TypeDefinition
	definitions astmodel.TypeDefinitionSet
}

// StorageConversionPropertyTestCaseFactory creates test cases for property conversion testing.
// We have a lot of test cases to cover, so we use a factory to create them.
type StorageConversionPropertyTestCaseFactory struct {
	cases          map[string]*StorageConversionPropertyTestCase
	currentPackage astmodel.LocalPackageReference
	nextPackage    astmodel.LocalPackageReference
}

// newStorageConversionPropertyTestCaseFactory creates a new test case factory for property conversion testing
func newStorageConversionPropertyTestCaseFactory() *StorageConversionPropertyTestCaseFactory {
	return &StorageConversionPropertyTestCaseFactory{
		cases:          make(map[string]*StorageConversionPropertyTestCase),
		currentPackage: test.MakeLocalPackageReference("verification", "vCurrent"),
		nextPackage:    test.MakeLocalPackageReference("verification", "vNext"),
	}
}

func (factory *StorageConversionPropertyTestCaseFactory) CreatePropertyAssignmentFunctionTestCases() map[string]*StorageConversionPropertyTestCase {
	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)

	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))
	requiredCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", currentEnum.Name())
	requiredHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", nextEnum.Name())
	optionalCurrentEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(currentEnum.Name()))
	optionalHubEnumProperty := astmodel.NewPropertyDefinition("Release", "release", astmodel.NewOptionalType(nextEnum.Name()))

	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Release"), roleType)
	hubRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Release"), roleType)

	referenceProperty := astmodel.NewPropertyDefinition("Reference", "reference", astmodel.ResourceReferenceType)
	knownReferenceProperty := astmodel.NewPropertyDefinition("KnownReference", "known-reference", astmodel.KnownResourceReferenceType)
	jsonProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.JSONType)
	optionalJSONProperty := astmodel.NewPropertyDefinition("JSONBlob", "jsonBlob", astmodel.NewOptionalType(astmodel.JSONType))
	jsonObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", jsonObjectType)
	optionalJSONObjectProperty := astmodel.NewPropertyDefinition("JSONObject", "jsonObject", astmodel.NewOptionalType(jsonObjectType))

	// Some pretty contrived cases to catch shadowing bugs in conversion code for arrays and maps.
	// Array of object
	indexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(currentRole.Name()))
	hubIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(hubRole.Name()))

	// Map of object
	keysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, currentRole.Name()))
	hubKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, hubRole.Name()))

	// Handcrafted impls
	sliceOfStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.StringType))
	mapOfStringToStringProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	bagProperty := astmodel.NewPropertyDefinition("propertyBag", "$propertyBag", astmodel.PropertyBagType)

	idFactory := astmodel.NewIdentifierFactory()
	ageFunction := test.NewFakeFunction("Age", idFactory)
	ageFunction.TypeReturned = astmodel.IntType

	nastyProperty := astmodel.NewPropertyDefinition(
		"nasty",
		"nasty",
		astmodel.NewMapType(
			astmodel.StringType,
			astmodel.NewArrayType(
				astmodel.NewMapType(astmodel.StringType, astmodel.BoolType))))

	factory.createPropertyAssignmentTest("NastyTest", nastyProperty, nastyProperty)

	factory.createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndRequiredEnum", requiredCurrentEnumProperty, requiredHubEnumProperty, currentEnum, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenRequiredEnumAndOptionalEnum", requiredCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalEnumAndOptionalEnum", optionalCurrentEnumProperty, optionalHubEnumProperty, currentEnum, nextEnum)

	factory.createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndRequiredObject", requiredCurrentRoleProperty, requiredHubRoleProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertBetweenRequiredObjectAndOptionalObject", requiredCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalObjectAndOptionalObject", optionalCurrentRoleProperty, optionalNextRoleProperty, currentRole, hubRole)
	// We don't handle/test Object Aliases here because they are always removed by the RemoveTypeAliases pipeline stage
	factory.createFunctionAssignmentTest("ReadFromFunctionIntoProperty", requiredIntProperty, ageFunction)
	factory.createFunctionAssignmentTest("ReadFromFunctionIntoOptionalProperty", optionalIntProperty, ageFunction)

	factory.createPropertyAssignmentTest("CopyReferenceProperty", referenceProperty, referenceProperty)
	factory.createPropertyAssignmentTest("CopyKnownReferenceProperty", knownReferenceProperty, knownReferenceProperty)

	factory.createPropertyAssignmentTest("CopyJSONProperty", jsonProperty, jsonProperty)
	factory.createPropertyAssignmentTest("CopyJSONObjectProperty", jsonObjectProperty, jsonObjectProperty)
	factory.createPropertyAssignmentTest("CopyOptionalJSONProperty", optionalJSONProperty, jsonProperty)
	factory.createPropertyAssignmentTest("CopyJSONObjectProperty", optionalJSONObjectProperty, jsonObjectProperty)

	factory.createPropertyAssignmentTest("CopyRequiredStringToPropertyBag", requiredStringProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyOptionalStringToPropertyBag", optionalStringProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyRequiredIntToPropertyBag", requiredIntProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyOptionalIntToPropertyBag", optionalIntProperty, bagProperty)

	factory.createPropertyAssignmentTest("CopyRequiredObjectToPropertyBag", requiredCurrentRoleProperty, bagProperty, currentRole)
	factory.createPropertyAssignmentTest("CopyOptionalObjectToPropertyBag", optionalCurrentRoleProperty, bagProperty, currentRole)

	factory.createPropertyAssignmentTest("CopySliceToPropertyBag", arrayOfRequiredIntProperty, bagProperty)
	factory.createPropertyAssignmentTest("CopyMapToPropertyBag", mapOfRequiredIntsProperty, bagProperty)

	factory.createPropertyAssignmentTest("ConvertSliceNamedIndexes", indexesProperty, hubIndexesProperty, currentRole, hubRole)
	factory.createPropertyAssignmentTest("ConvertMapNamedKeys", keysProperty, hubKeysProperty, currentRole, hubRole)

	factory.createPropertyAssignmentTest("SetSliceOfString", sliceOfStringProperty, sliceOfStringProperty)
	factory.createPropertyAssignmentTest("SetMapOfStringToString", mapOfStringToStringProperty, mapOfStringToStringProperty)

	factory.createPrimitiveTypeTestCases()
	factory.createCollectionTestCases()
	factory.createEnumTestCases()
	return factory.cases
}

// createPrimitiveTypeTestCases creates test cases for conversion of primitive types (aka string, int, etc),
// optional variations, and aliases of those types.
func (factory *StorageConversionPropertyTestCaseFactory) createPrimitiveTypeTestCases() {
	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	optionalStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType)
	requiredIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalIntProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)

	factory.createPropertyAssignmentTest("SetStringFromString", requiredStringProperty, requiredStringProperty)
	factory.createPropertyAssignmentTest("SetStringFromOptionalString", requiredStringProperty, optionalStringProperty)
	factory.createPropertyAssignmentTest("SetOptionalStringFromString", optionalStringProperty, requiredStringProperty)
	factory.createPropertyAssignmentTest("SetOptionalStringFromOptionalString", optionalStringProperty, optionalStringProperty)

	factory.createPropertyAssignmentTest("SetIntFromInt", requiredIntProperty, requiredIntProperty)
	factory.createPropertyAssignmentTest("SetIntFromOptionalInt", requiredIntProperty, optionalIntProperty)
	factory.createPropertyAssignmentTest("SetOptionalIntFromOptionalInt", optionalIntProperty, optionalIntProperty)

	// Aliases of primitive types
	currentAge := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Age"), astmodel.IntType)
	nextAge := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Age"), astmodel.IntType)

	requiredPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.IntType)
	optionalPrimitiveAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.OptionalIntType)
	requiredCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", currentAge.Name())
	requiredNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", nextAge.Name())
	optionalCurrentAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(currentAge.Name()))
	optionalNextAgeProperty := astmodel.NewPropertyDefinition("Age", "age", astmodel.NewOptionalType(nextAge.Name()))

	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndAliasType", requiredCurrentAgeProperty, requiredNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndOptionalAliasType", requiredCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalAliasType", optionalCurrentAgeProperty, optionalNextAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenAliasAndBaseType", requiredCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge, nextAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndBaseType", optionalCurrentAgeProperty, requiredPrimitiveAgeProperty, currentAge)
	factory.createPropertyAssignmentTest("ConvertBetweenOptionalAliasAndOptionalBaseType", optionalCurrentAgeProperty, optionalPrimitiveAgeProperty, currentAge)
}

// createCollectionTestCases creates test cases for conversion of collection types (aka array, map), and aliases of those types.
// We don't handle/test optional array/map aliases because they are always made non-optional by the FixOptionalCollectionAliases stage
func (factory *StorageConversionPropertyTestCaseFactory) createCollectionTestCases() {
	arrayOfRequiredIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.IntType))
	arrayOfOptionalIntProperty := astmodel.NewPropertyDefinition("Scores", "scores", astmodel.NewArrayType(astmodel.OptionalIntType))

	mapOfRequiredIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.IntType))
	mapOfOptionalIntsProperty := astmodel.NewPropertyDefinition("Ratings", "ratings", astmodel.NewMapType(astmodel.StringType, astmodel.OptionalIntType))

	factory.createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfRequired", arrayOfRequiredIntProperty, arrayOfRequiredIntProperty)
	factory.createPropertyAssignmentTest("SetArrayOfRequiredFromArrayOfOptional", arrayOfRequiredIntProperty, arrayOfOptionalIntProperty)
	factory.createPropertyAssignmentTest("SetArrayOfOptionalFromArrayOfRequired", arrayOfOptionalIntProperty, arrayOfRequiredIntProperty)

	factory.createPropertyAssignmentTest("SetMapOfRequiredFromMapOfRequired", mapOfRequiredIntsProperty, mapOfRequiredIntsProperty)
	factory.createPropertyAssignmentTest("SetMapOfRequiredFromMapOfOptional", mapOfRequiredIntsProperty, mapOfOptionalIntsProperty)
	factory.createPropertyAssignmentTest("SetMapOfOptionalFromMapOfRequired", mapOfOptionalIntsProperty, mapOfRequiredIntsProperty)

	requiredStringProperty := astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType)
	roleType := astmodel.NewObjectType().WithProperty(requiredStringProperty).WithProperty(arrayOfRequiredIntProperty)
	currentRole := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Release"), roleType)

	// Array of optional object
	optionalIndexesProperty := astmodel.NewPropertyDefinition("Indexes", "indexes", astmodel.NewArrayType(astmodel.NewOptionalType(currentRole.Name())))

	// Map of optional object
	optionalKeysProperty := astmodel.NewPropertyDefinition("Keys", "keys", astmodel.NewMapType(astmodel.StringType, astmodel.NewOptionalType(currentRole.Name())))

	// Array of array of int
	arrayofArraysProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewArrayType(astmodel.NewArrayType(astmodel.IntType)))

	// Map of map of int
	mapOfMapsProperty := astmodel.NewPropertyDefinition("Items", "items", astmodel.NewMapType(astmodel.StringType, astmodel.NewMapType(astmodel.StringType, astmodel.IntType)))

	factory.createPropertyAssignmentTest("ConvertSliceNamedIndexesOfOptionalObject", optionalIndexesProperty, optionalIndexesProperty, currentRole)
	factory.createPropertyAssignmentTest("ConvertMapNamedKeysOfOptionalObject", optionalKeysProperty, optionalKeysProperty, currentRole)
	factory.createPropertyAssignmentTest("SetSliceOfSlices", arrayofArraysProperty, arrayofArraysProperty)
	factory.createPropertyAssignmentTest("SetMapOfMaps", mapOfMapsProperty, mapOfMapsProperty)

	// Aliases of collection types
	currentPhoneNumbers := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "PhoneNumbers"), astmodel.NewValidatedType(astmodel.NewArrayType(astmodel.StringType), astmodel.ArrayValidations{}))
	nextPhoneNumbers := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "PhoneNumbers"), astmodel.NewValidatedType(astmodel.NewArrayType(astmodel.StringType), astmodel.ArrayValidations{}))
	currentAddresses := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.currentPackage, "Addresses"), astmodel.NewValidatedType(astmodel.NewMapType(astmodel.StringType, astmodel.StringType), nil))
	nextAddresses := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(factory.nextPackage, "Addresses"), astmodel.NewValidatedType(astmodel.NewMapType(astmodel.StringType, astmodel.StringType), nil))

	requiredCurrentPhoneNumbersProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", currentPhoneNumbers.Name())
	requiredNextPhoneNumbersProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", nextPhoneNumbers.Name())
	requiredCurrentPhoneNumbersBaseTypeProperty := astmodel.NewPropertyDefinition("PhoneNumbers", "phoneNumbers", astmodel.NewArrayType(astmodel.StringType))
	requiredCurrentAddressesProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", currentAddresses.Name())
	requiredNextAddressesProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", nextAddresses.Name())
	requiredCurrentAddressesBaseTypeProperty := astmodel.NewPropertyDefinition("Addresses", "addresses", astmodel.NewMapType(astmodel.StringType, astmodel.StringType))

	factory.createPropertyAssignmentTest("ConvertBetweenArrayAliasAndArrayAlias", requiredCurrentPhoneNumbersProperty, requiredNextPhoneNumbersProperty, currentPhoneNumbers, nextPhoneNumbers)
	factory.createPropertyAssignmentTest("ConvertBetweenArrayAliasAndBaseType", requiredCurrentPhoneNumbersBaseTypeProperty, requiredNextPhoneNumbersProperty, nextPhoneNumbers)
	factory.createPropertyAssignmentTest("ConvertBetweenMapAliasAndMapAlias", requiredCurrentAddressesProperty, requiredNextAddressesProperty, currentAddresses, nextAddresses)
	factory.createPropertyAssignmentTest("ConvertBetweenMapAliasAndBaseType", requiredCurrentAddressesBaseTypeProperty, requiredNextAddressesProperty, nextAddresses)
}

func (factory *StorageConversionPropertyTestCaseFactory) createPropertyAssignmentTest(
	name string,
	currentProperty *astmodel.PropertyDefinition,
	hubProperty *astmodel.PropertyDefinition,
	otherDefinitions ...astmodel.TypeDefinition,
) {
	currentType := astmodel.NewObjectType().WithProperty(currentProperty)
	currentDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.currentPackage, "Person"),
		currentType)

	hubType := astmodel.NewObjectType().WithProperty(hubProperty)
	hubDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.nextPackage, "Person"),
		hubType)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(currentDefinition)
	defs.Add(hubDefinition)
	defs.AddAll(otherDefinitions...)

	factory.addCase(&StorageConversionPropertyTestCase{
		name:        name,
		current:     currentDefinition,
		other:       hubDefinition,
		definitions: defs,
	})
}

func (factory *StorageConversionPropertyTestCaseFactory) createFunctionAssignmentTest(
	name string,
	property *astmodel.PropertyDefinition,
	function astmodel.Function,
) {
	currentType := astmodel.NewObjectType().WithFunction(function)
	currentDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.currentPackage, "Person"),
		currentType)

	hubType := astmodel.NewObjectType().WithProperty(property)
	hubDefinition := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(factory.nextPackage, "Person"),
		hubType)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(currentDefinition)
	defs.Add(hubDefinition)

	factory.addCase(&StorageConversionPropertyTestCase{
		name:        name,
		current:     currentDefinition,
		other:       hubDefinition,
		definitions: defs,
	})
}

func (factory *StorageConversionPropertyTestCaseFactory) addCase(c *StorageConversionPropertyTestCase) {
	factory.cases[c.name] = c
}

func TestGolden_PropertyAssignmentFunction_AsFunc(t *testing.T) {
	t.Parallel()

	factory := newStorageConversionPropertyTestCaseFactory()
	for n, c := range factory.CreatePropertyAssignmentFunctionTestCases() {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			runTestPropertyAssignmentFunction_AsFunc(c, t)
		})
	}
}

func runTestPropertyAssignmentFunction_AsFunc(c *StorageConversionPropertyTestCase, t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	currentType, ok := astmodel.AsObjectType(c.current.Type())
	g.Expect(ok).To(BeTrue())

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, c.definitions, idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(c.current, c.other, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(BeNil())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(c.current, c.other, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(BeNil())

	receiverDefinition := c.current.WithType(currentType.WithFunction(assignFrom).WithFunction(assignTo))

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, c.name, receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenPropertyBagPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	person2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.PropertyBagProperty)

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, make(astmodel.TypeDefinitionSet), idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "PropertyBags", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenTypeRenamed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	// We have a definition for Location
	location := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty,
		test.CityProperty)

	// ... which gets renamed to Venue in a later version
	venue := test.CreateObjectDefinition(
		test.Pkg2021,
		"Venue",
		test.FullAddressProperty,
		test.CityProperty)

	// The earlier version of Event has a property "Where" of type "Location" ...
	whereLocationProperty := astmodel.NewPropertyDefinition("Where", "where", location.Name())
	event2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Event",
		whereLocationProperty)

	// ... in the later version of Event, the property "Where" is now of type "Venue"
	whereVenueProperty := astmodel.NewPropertyDefinition("Where", "where", venue.Name())
	event2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Event",
		whereVenueProperty)

	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			location.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.NameInNextVersion.Set(venue.Name().Name())
				return nil
			})).
		To(Succeed())

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(location, venue)

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, defs, idFactory).
		WithConfiguration(omc)

	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(event2020, event2021, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(event2020, event2021, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(event2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	// The generated code should be using the type "Location" when referencing the earlier version of the property
	// "Where" and "Venue" for the later version. The types are visible in declarations of temporary variables,
	// and in the name of the Assign*() functions.
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "PropertyTypeRenamed", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenSharedObjectVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	location2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty)

	location2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Location",
		test.FullAddressProperty)

	location2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Location",
		test.FullAddressProperty)

	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2020.Name()))

	person2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2022.Name()))

	definitions := make(astmodel.TypeDefinitionSet)
	definitions.AddAll(location2020, location2021, location2022)
	definitions.AddAll(person2020, person2022)

	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(person2020.Name(), person2022.Name())
	builder.Add(location2020.Name(), location2021.Name(), location2022.Name())

	graph, err := builder.Build()
	g.Expect(err).To(BeNil())

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, definitions, idFactory).
		WithConversionGraph(graph)

	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "SharedObject", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenMultipleIntermediateSharedObjectVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	location2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location",
		test.FullAddressProperty)

	location202101 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210301"),
		"Location",
		test.FullAddressProperty)

	location202106 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210306"),
		"Location",
		test.FullAddressProperty)

	location202112 := test.CreateObjectDefinition(
		test.MakeLocalPackageReference(test.Group, "v20210312"),
		"Location",
		test.FullAddressProperty)

	location2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Location",
		test.FullAddressProperty)

	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2020.Name()))

	person2022 := test.CreateObjectDefinition(
		test.Pkg2022,
		"Person",
		astmodel.NewPropertyDefinition("Residence", "residence", location2022.Name()))

	definitions := make(astmodel.TypeDefinitionSet)
	definitions.AddAll(location2020, location202101, location202106, location202112, location2022)
	definitions.AddAll(person2020, person2022)

	cfg := config.NewObjectModelConfiguration()
	builder := storage.NewConversionGraphBuilder(cfg, "v")
	builder.Add(
		person2020.Name(),
		person2022.Name(),
		location2020.Name(),
		location202101.Name(),
		location202106.Name(),
		location202112.Name(),
		location2022.Name())

	graph, err := builder.Build()
	g.Expect(err).To(BeNil())

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, definitions, idFactory).
		WithConversionGraph(graph)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertFrom)
	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2022, conversions.ConvertTo)
	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "SharedObjectMultiple", receiverDefinition)
}

func TestGolden_PropertyAssignmentFunction_WhenOverrideInterfacePresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()
	injector := astmodel.NewFunctionInjector()

	person2020 := test.CreateObjectDefinition(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	person2021 := test.CreateObjectDefinition(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.PropertyBagProperty)

	overrideInterfaceName := astmodel.MakeInternalTypeName(test.Pkg2020, "personAssignable")

	conversionContext := conversions.NewPropertyConversionContext(conversions.AssignPropertiesMethodPrefix, make(astmodel.TypeDefinitionSet), idFactory)
	assignFromBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertFrom)
	assignFromBuilder.UseAugmentationInterface(overrideInterfaceName)

	assignFrom, err := assignFromBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	assignToBuilder := NewPropertyAssignmentFunctionBuilder(person2020, person2021, conversions.ConvertTo)
	assignToBuilder.UseAugmentationInterface(overrideInterfaceName)

	assignTo, err := assignToBuilder.Build(conversionContext)
	g.Expect(err).To(Succeed())

	receiverDefinition, err := injector.Inject(person2020, assignFrom, assignTo)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OverrideInterface", receiverDefinition)
}
