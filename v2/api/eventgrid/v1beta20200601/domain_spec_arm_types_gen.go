// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Domain_Spec_ARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Domain.
	Properties *DomainProperties_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Domain_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domain Domain_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (domain *Domain_Spec_ARM) GetName() string {
	return domain.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains"
func (domain *Domain_Spec_ARM) GetType() string {
	return "Microsoft.EventGrid/domains"
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/DomainProperties
type DomainProperties_ARM struct {
	// InboundIpRules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered
	// only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule_ARM `json:"inboundIpRules,omitempty"`

	// InputSchema: This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *DomainProperties_InputSchema `json:"inputSchema,omitempty"`

	// InputSchemaMapping: By default, Event Grid expects events to be in the Event Grid event schema. Specifying an input
	// schema mapping enables publishing to Event Grid using a custom input schema. Currently, the only supported type of
	// InputSchemaMapping is 'JsonInputSchemaMapping'.
	InputSchemaMapping *JsonInputSchemaMapping_ARM `json:"inputSchemaMapping,omitempty"`

	// PublicNetworkAccess: This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso
	// cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />.
	PublicNetworkAccess *DomainProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/InboundIpRule
type InboundIpRule_ARM struct {
	// Action: Action to perform based on the match or no match of the IpMask.
	Action *InboundIpRule_Action `json:"action,omitempty"`

	// IpMask: IP Address in CIDR notation e.g., 10.0.0.0/8.
	IpMask *string `json:"ipMask,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMapping
type JsonInputSchemaMapping_ARM struct {
	InputSchemaMappingType *JsonInputSchemaMapping_InputSchemaMappingType `json:"inputSchemaMappingType,omitempty"`

	// Properties: This can be used to map properties of a source schema (or default values, for certain supported properties)
	// to properties of the EventGridEvent schema.
	Properties *JsonInputSchemaMappingProperties_ARM `json:"properties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMappingProperties
type JsonInputSchemaMappingProperties_ARM struct {
	// DataVersion: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	DataVersion *JsonFieldWithDefault_ARM `json:"dataVersion,omitempty"`

	// EventTime: This is used to express the source of an input schema mapping for a single target field in the Event Grid
	// Event schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a
	// field in the input event schema.
	EventTime *JsonField_ARM `json:"eventTime,omitempty"`

	// EventType: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	EventType *JsonFieldWithDefault_ARM `json:"eventType,omitempty"`

	// Id: This is used to express the source of an input schema mapping for a single target field in the Event Grid Event
	// schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field
	// in the input event schema.
	Id *JsonField_ARM `json:"id,omitempty"`

	// Subject: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	Subject *JsonFieldWithDefault_ARM `json:"subject,omitempty"`

	// Topic: This is used to express the source of an input schema mapping for a single target field in the Event Grid Event
	// schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field
	// in the input event schema.
	Topic *JsonField_ARM `json:"topic,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonField
type JsonField_ARM struct {
	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonFieldWithDefault
type JsonFieldWithDefault_ARM struct {
	// DefaultValue: The default value to be used for mapping when a SourceField is not provided or if there's no property with
	// the specified name in the published JSON event payload.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}