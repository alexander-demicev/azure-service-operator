// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210101p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101preview"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NamespacesTopicsSubscriptionExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NamespacesTopicsSubscriptionExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210101p.NamespacesTopicsSubscription{},
		&v20210101ps.NamespacesTopicsSubscription{}}
}