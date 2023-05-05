// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_Topics_Subscription_Spec. Use v1api20210101preview.Namespaces_Topics_Subscription_Spec instead
type Namespaces_Topics_Subscription_Spec_ARM struct {
	Name       string                        `json:"name,omitempty"`
	Properties *SBSubscriptionProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Topics_Subscription_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (subscription Namespaces_Topics_Subscription_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (subscription *Namespaces_Topics_Subscription_Spec_ARM) GetName() string {
	return subscription.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions"
func (subscription *Namespaces_Topics_Subscription_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions"
}

// Deprecated version of SBSubscriptionProperties. Use v1api20210101preview.SBSubscriptionProperties instead
type SBSubscriptionProperties_ARM struct {
	AutoDeleteOnIdle                          *string `json:"autoDeleteOnIdle,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool   `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool   `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string `json:"forwardTo,omitempty"`
	LockDuration                              *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int    `json:"maxDeliveryCount,omitempty"`
	RequiresSession                           *bool   `json:"requiresSession,omitempty"`
}