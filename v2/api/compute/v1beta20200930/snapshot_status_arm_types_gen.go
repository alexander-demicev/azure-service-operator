// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930

// Deprecated version of Snapshot_STATUS. Use v1api20200930.Snapshot_STATUS instead
type Snapshot_STATUS_ARM struct {
	ExtendedLocation *ExtendedLocation_STATUS_ARM   `json:"extendedLocation,omitempty"`
	Id               *string                        `json:"id,omitempty"`
	Location         *string                        `json:"location,omitempty"`
	ManagedBy        *string                        `json:"managedBy,omitempty"`
	Name             *string                        `json:"name,omitempty"`
	Properties       *SnapshotProperties_STATUS_ARM `json:"properties,omitempty"`
	Sku              *SnapshotSku_STATUS_ARM        `json:"sku,omitempty"`
	Tags             map[string]string              `json:"tags,omitempty"`
	Type             *string                        `json:"type,omitempty"`
}

// Deprecated version of SnapshotProperties_STATUS. Use v1api20200930.SnapshotProperties_STATUS instead
type SnapshotProperties_STATUS_ARM struct {
	CreationData                 *CreationData_STATUS_ARM                    `json:"creationData,omitempty"`
	DiskAccessId                 *string                                     `json:"diskAccessId,omitempty"`
	DiskSizeBytes                *int                                        `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                        `json:"diskSizeGB,omitempty"`
	DiskState                    *DiskState_STATUS                           `json:"diskState,omitempty"`
	Encryption                   *Encryption_STATUS_ARM                      `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUS_ARM    `json:"encryptionSettingsCollection,omitempty"`
	HyperVGeneration             *SnapshotProperties_HyperVGeneration_STATUS `json:"hyperVGeneration,omitempty"`
	Incremental                  *bool                                       `json:"incremental,omitempty"`
	NetworkAccessPolicy          *NetworkAccessPolicy_STATUS                 `json:"networkAccessPolicy,omitempty"`
	OsType                       *SnapshotProperties_OsType_STATUS           `json:"osType,omitempty"`
	ProvisioningState            *string                                     `json:"provisioningState,omitempty"`
	PurchasePlan                 *PurchasePlan_STATUS_ARM                    `json:"purchasePlan,omitempty"`
	TimeCreated                  *string                                     `json:"timeCreated,omitempty"`
	UniqueId                     *string                                     `json:"uniqueId,omitempty"`
}

// Deprecated version of SnapshotSku_STATUS. Use v1api20200930.SnapshotSku_STATUS instead
type SnapshotSku_STATUS_ARM struct {
	Name *SnapshotSku_Name_STATUS `json:"name,omitempty"`
	Tier *string                  `json:"tier,omitempty"`
}

// Deprecated version of SnapshotSku_Name_STATUS. Use v1api20200930.SnapshotSku_Name_STATUS instead
type SnapshotSku_Name_STATUS string

const (
	SnapshotSku_Name_STATUS_Premium_LRS  = SnapshotSku_Name_STATUS("Premium_LRS")
	SnapshotSku_Name_STATUS_Standard_LRS = SnapshotSku_Name_STATUS("Standard_LRS")
	SnapshotSku_Name_STATUS_Standard_ZRS = SnapshotSku_Name_STATUS("Standard_ZRS")
)