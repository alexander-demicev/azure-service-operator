// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

type StorageAccounts_FileServices_Share_STATUS_ARM struct {
	// Etag: Resource Etag.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the file share.
	Properties *FileShareProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of the file share.
type FileShareProperties_STATUS_ARM struct {
	// AccessTier: Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and
	// Cool. FileStorage account can choose Premium.
	AccessTier *FileShareProperties_AccessTier_STATUS `json:"accessTier,omitempty"`

	// AccessTierChangeTime: Indicates the last modification time for share access tier.
	AccessTierChangeTime *string `json:"accessTierChangeTime,omitempty"`

	// AccessTierStatus: Indicates if there is a pending transition for access tier.
	AccessTierStatus *string `json:"accessTierStatus,omitempty"`

	// Deleted: Indicates whether the share was deleted.
	Deleted *bool `json:"deleted,omitempty"`

	// DeletedTime: The deleted time if the share was deleted.
	DeletedTime *string `json:"deletedTime,omitempty"`

	// EnabledProtocols: The authentication protocol that is used for the file share. Can only be specified when creating a
	// share.
	EnabledProtocols *FileShareProperties_EnabledProtocols_STATUS `json:"enabledProtocols,omitempty"`

	// LastModifiedTime: Returns the date and time the share was last modified.
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`

	// LeaseDuration: Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
	LeaseDuration *FileShareProperties_LeaseDuration_STATUS `json:"leaseDuration,omitempty"`

	// LeaseState: Lease state of the share.
	LeaseState *FileShareProperties_LeaseState_STATUS `json:"leaseState,omitempty"`

	// LeaseStatus: The lease status of the share.
	LeaseStatus *FileShareProperties_LeaseStatus_STATUS `json:"leaseStatus,omitempty"`

	// Metadata: A name-value pair to associate with the share as metadata.
	Metadata map[string]string `json:"metadata"`

	// RemainingRetentionDays: Remaining retention days for share that was soft deleted.
	RemainingRetentionDays *int `json:"remainingRetentionDays,omitempty"`

	// RootSquash: The property is for NFS share only. The default is NoRootSquash.
	RootSquash *FileShareProperties_RootSquash_STATUS `json:"rootSquash,omitempty"`

	// ShareQuota: The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
	// For Large File Shares, the maximum size is 102400.
	ShareQuota *int `json:"shareQuota,omitempty"`

	// ShareUsageBytes: The approximate size of the data stored on the share. Note that this value may not include all recently
	// created or recently resized files.
	ShareUsageBytes *int `json:"shareUsageBytes,omitempty"`

	// SignedIdentifiers: List of stored access policies specified on the share.
	SignedIdentifiers []SignedIdentifier_STATUS_ARM `json:"signedIdentifiers"`

	// SnapshotTime: Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
	SnapshotTime *string `json:"snapshotTime,omitempty"`

	// Version: The version of the share.
	Version *string `json:"version,omitempty"`
}

type SignedIdentifier_STATUS_ARM struct {
	// AccessPolicy: Access policy
	AccessPolicy *AccessPolicy_STATUS_ARM `json:"accessPolicy,omitempty"`

	// Id: An unique identifier of the stored access policy.
	Id *string `json:"id,omitempty"`
}

type AccessPolicy_STATUS_ARM struct {
	// ExpiryTime: Expiry time of the access policy
	ExpiryTime *string `json:"expiryTime,omitempty"`

	// Permission: List of abbreviated permissions.
	Permission *string `json:"permission,omitempty"`

	// StartTime: Start time of the access policy
	StartTime *string `json:"startTime,omitempty"`
}