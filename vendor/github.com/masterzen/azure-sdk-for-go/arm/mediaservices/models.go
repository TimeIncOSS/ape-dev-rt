package mediaservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// EntityNameUnavailabilityReason enumerates the values for entity name
// unavailability reason.
type EntityNameUnavailabilityReason string

const (
	// AlreadyExists specifies the already exists state for entity name
	// unavailability reason.
	AlreadyExists EntityNameUnavailabilityReason = "AlreadyExists"
	// Invalid specifies the invalid state for entity name unavailability
	// reason.
	Invalid EntityNameUnavailabilityReason = "Invalid"
	// None specifies the none state for entity name unavailability reason.
	None EntityNameUnavailabilityReason = "None"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary specifies the primary state for key type.
	Primary KeyType = "Primary"
	// Secondary specifies the secondary state for key type.
	Secondary KeyType = "Secondary"
)

// APIEndpoint is the properties for a Media Services REST API endpoint.
type APIEndpoint struct {
	Endpoint     *string `json:"endpoint,omitempty"`
	MajorVersion *string `json:"majorVersion,omitempty"`
}

// APIError is the error returned from a failed Media Services REST API call.
type APIError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// CheckNameAvailabilityInput is the request body for CheckNameAvailability
// API.
type CheckNameAvailabilityInput struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput is the response body for CheckNameAvailability
// API.
type CheckNameAvailabilityOutput struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool                          `json:"nameAvailable,omitempty"`
	Reason            EntityNameUnavailabilityReason `json:"reason,omitempty"`
	Message           *string                        `json:"message,omitempty"`
}

// MediaService is the properties of a Media Service resource.
type MediaService struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	Tags              *map[string]*string     `json:"tags,omitempty"`
	Properties        *MediaServiceProperties `json:"properties,omitempty"`
}

// MediaServiceCollection is the collection of Media Service resources.
type MediaServiceCollection struct {
	autorest.Response `json:"-"`
	Value             *[]MediaService `json:"value,omitempty"`
}

// MediaServiceProperties is the additional properties of a Media Service
// resource.
type MediaServiceProperties struct {
	APIEndpoints    *[]APIEndpoint    `json:"apiEndpoints,omitempty"`
	StorageAccounts *[]StorageAccount `json:"storageAccounts,omitempty"`
}

// RegenerateKeyInput is the request body for a RegenerateKey API.
type RegenerateKeyInput struct {
	KeyType KeyType `json:"keyType,omitempty"`
}

// RegenerateKeyOutput is the response body for a RegenerateKey API.
type RegenerateKeyOutput struct {
	autorest.Response `json:"-"`
	Key               *string `json:"key,omitempty"`
}

// Resource is
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ServiceKeys is the response body for a ListKeys API.
type ServiceKeys struct {
	autorest.Response     `json:"-"`
	PrimaryAuthEndpoint   *string `json:"primaryAuthEndpoint,omitempty"`
	SecondaryAuthEndpoint *string `json:"secondaryAuthEndpoint,omitempty"`
	PrimaryKey            *string `json:"primaryKey,omitempty"`
	SecondaryKey          *string `json:"secondaryKey,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
}

// StorageAccount is the properties of a storage account associated with this
// resource.
type StorageAccount struct {
	ID        *string `json:"id,omitempty"`
	IsPrimary *bool   `json:"isPrimary,omitempty"`
}

// SyncStorageKeysInput is the request  body for a SyncStorageKeys API.
type SyncStorageKeysInput struct {
	ID *string `json:"id,omitempty"`
}

// TrackedResource is aRM tracked resource
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}
