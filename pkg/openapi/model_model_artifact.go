/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelArtifact{}

// ModelArtifact An ML model artifact.
type ModelArtifact struct {
	ArtifactType string `json:"artifactType"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
	// The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact.
	Uri   *string        `json:"uri,omitempty"`
	State *ArtifactState `json:"state,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// Output only. The unique server generated id of the resource.
	Id *string `json:"id,omitempty"`
	// Output only. Create time of the resource in millisecond since epoch.
	CreateTimeSinceEpoch *string `json:"createTimeSinceEpoch,omitempty"`
	// Output only. Last update time of the resource since epoch in millisecond since epoch.
	LastUpdateTimeSinceEpoch *string `json:"lastUpdateTimeSinceEpoch,omitempty"`
	// Name of the model format.
	ModelFormatName *string `json:"modelFormatName,omitempty"`
	// Storage secret name.
	StorageKey *string `json:"storageKey,omitempty"`
	// Path for model in storage provided by `storageKey`.
	StoragePath *string `json:"storagePath,omitempty"`
	// Version of the model format.
	ModelFormatVersion *string `json:"modelFormatVersion,omitempty"`
	// Name of the service account with storage secret.
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`
}

// NewModelArtifact instantiates a new ModelArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelArtifact(artifactType string) *ModelArtifact {
	this := ModelArtifact{}
	var state ArtifactState = ARTIFACTSTATE_UNKNOWN
	this.State = &state
	return &this
}

// NewModelArtifactWithDefaults instantiates a new ModelArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelArtifactWithDefaults() *ModelArtifact {
	this := ModelArtifact{}
	var artifactType string = "model-artifact"
	this.ArtifactType = artifactType
	var state ArtifactState = ARTIFACTSTATE_UNKNOWN
	this.State = &state
	return &this
}

// GetArtifactType returns the ArtifactType field value
func (o *ModelArtifact) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *ModelArtifact) SetArtifactType(v string) {
	o.ArtifactType = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *ModelArtifact) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ModelArtifact) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *ModelArtifact) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelArtifact) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelArtifact) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelArtifact) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *ModelArtifact) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *ModelArtifact) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *ModelArtifact) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ModelArtifact) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ModelArtifact) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ModelArtifact) SetUri(v string) {
	o.Uri = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModelArtifact) GetState() ArtifactState {
	if o == nil || IsNil(o.State) {
		var ret ArtifactState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetStateOk() (*ArtifactState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModelArtifact) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ArtifactState and assigns it to the State field.
func (o *ModelArtifact) SetState(v ArtifactState) {
	o.State = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelArtifact) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelArtifact) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelArtifact) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelArtifact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelArtifact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelArtifact) SetId(v string) {
	o.Id = &v
}

// GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field value if set, zero value otherwise.
func (o *ModelArtifact) GetCreateTimeSinceEpoch() string {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.CreateTimeSinceEpoch
}

// GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetCreateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTimeSinceEpoch) {
		return nil, false
	}
	return o.CreateTimeSinceEpoch, true
}

// HasCreateTimeSinceEpoch returns a boolean if a field has been set.
func (o *ModelArtifact) HasCreateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.CreateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetCreateTimeSinceEpoch gets a reference to the given string and assigns it to the CreateTimeSinceEpoch field.
func (o *ModelArtifact) SetCreateTimeSinceEpoch(v string) {
	o.CreateTimeSinceEpoch = &v
}

// GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field value if set, zero value otherwise.
func (o *ModelArtifact) GetLastUpdateTimeSinceEpoch() string {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		var ret string
		return ret
	}
	return *o.LastUpdateTimeSinceEpoch
}

// GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetLastUpdateTimeSinceEpochOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateTimeSinceEpoch) {
		return nil, false
	}
	return o.LastUpdateTimeSinceEpoch, true
}

// HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.
func (o *ModelArtifact) HasLastUpdateTimeSinceEpoch() bool {
	if o != nil && !IsNil(o.LastUpdateTimeSinceEpoch) {
		return true
	}

	return false
}

// SetLastUpdateTimeSinceEpoch gets a reference to the given string and assigns it to the LastUpdateTimeSinceEpoch field.
func (o *ModelArtifact) SetLastUpdateTimeSinceEpoch(v string) {
	o.LastUpdateTimeSinceEpoch = &v
}

// GetModelFormatName returns the ModelFormatName field value if set, zero value otherwise.
func (o *ModelArtifact) GetModelFormatName() string {
	if o == nil || IsNil(o.ModelFormatName) {
		var ret string
		return ret
	}
	return *o.ModelFormatName
}

// GetModelFormatNameOk returns a tuple with the ModelFormatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetModelFormatNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelFormatName) {
		return nil, false
	}
	return o.ModelFormatName, true
}

// HasModelFormatName returns a boolean if a field has been set.
func (o *ModelArtifact) HasModelFormatName() bool {
	if o != nil && !IsNil(o.ModelFormatName) {
		return true
	}

	return false
}

// SetModelFormatName gets a reference to the given string and assigns it to the ModelFormatName field.
func (o *ModelArtifact) SetModelFormatName(v string) {
	o.ModelFormatName = &v
}

// GetStorageKey returns the StorageKey field value if set, zero value otherwise.
func (o *ModelArtifact) GetStorageKey() string {
	if o == nil || IsNil(o.StorageKey) {
		var ret string
		return ret
	}
	return *o.StorageKey
}

// GetStorageKeyOk returns a tuple with the StorageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetStorageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StorageKey) {
		return nil, false
	}
	return o.StorageKey, true
}

// HasStorageKey returns a boolean if a field has been set.
func (o *ModelArtifact) HasStorageKey() bool {
	if o != nil && !IsNil(o.StorageKey) {
		return true
	}

	return false
}

// SetStorageKey gets a reference to the given string and assigns it to the StorageKey field.
func (o *ModelArtifact) SetStorageKey(v string) {
	o.StorageKey = &v
}

// GetStoragePath returns the StoragePath field value if set, zero value otherwise.
func (o *ModelArtifact) GetStoragePath() string {
	if o == nil || IsNil(o.StoragePath) {
		var ret string
		return ret
	}
	return *o.StoragePath
}

// GetStoragePathOk returns a tuple with the StoragePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetStoragePathOk() (*string, bool) {
	if o == nil || IsNil(o.StoragePath) {
		return nil, false
	}
	return o.StoragePath, true
}

// HasStoragePath returns a boolean if a field has been set.
func (o *ModelArtifact) HasStoragePath() bool {
	if o != nil && !IsNil(o.StoragePath) {
		return true
	}

	return false
}

// SetStoragePath gets a reference to the given string and assigns it to the StoragePath field.
func (o *ModelArtifact) SetStoragePath(v string) {
	o.StoragePath = &v
}

// GetModelFormatVersion returns the ModelFormatVersion field value if set, zero value otherwise.
func (o *ModelArtifact) GetModelFormatVersion() string {
	if o == nil || IsNil(o.ModelFormatVersion) {
		var ret string
		return ret
	}
	return *o.ModelFormatVersion
}

// GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetModelFormatVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ModelFormatVersion) {
		return nil, false
	}
	return o.ModelFormatVersion, true
}

// HasModelFormatVersion returns a boolean if a field has been set.
func (o *ModelArtifact) HasModelFormatVersion() bool {
	if o != nil && !IsNil(o.ModelFormatVersion) {
		return true
	}

	return false
}

// SetModelFormatVersion gets a reference to the given string and assigns it to the ModelFormatVersion field.
func (o *ModelArtifact) SetModelFormatVersion(v string) {
	o.ModelFormatVersion = &v
}

// GetServiceAccountName returns the ServiceAccountName field value if set, zero value otherwise.
func (o *ModelArtifact) GetServiceAccountName() string {
	if o == nil || IsNil(o.ServiceAccountName) {
		var ret string
		return ret
	}
	return *o.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifact) GetServiceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountName) {
		return nil, false
	}
	return o.ServiceAccountName, true
}

// HasServiceAccountName returns a boolean if a field has been set.
func (o *ModelArtifact) HasServiceAccountName() bool {
	if o != nil && !IsNil(o.ServiceAccountName) {
		return true
	}

	return false
}

// SetServiceAccountName gets a reference to the given string and assigns it to the ServiceAccountName field.
func (o *ModelArtifact) SetServiceAccountName(v string) {
	o.ServiceAccountName = &v
}

func (o ModelArtifact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifactType"] = o.ArtifactType
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreateTimeSinceEpoch) {
		toSerialize["createTimeSinceEpoch"] = o.CreateTimeSinceEpoch
	}
	if !IsNil(o.LastUpdateTimeSinceEpoch) {
		toSerialize["lastUpdateTimeSinceEpoch"] = o.LastUpdateTimeSinceEpoch
	}
	if !IsNil(o.ModelFormatName) {
		toSerialize["modelFormatName"] = o.ModelFormatName
	}
	if !IsNil(o.StorageKey) {
		toSerialize["storageKey"] = o.StorageKey
	}
	if !IsNil(o.StoragePath) {
		toSerialize["storagePath"] = o.StoragePath
	}
	if !IsNil(o.ModelFormatVersion) {
		toSerialize["modelFormatVersion"] = o.ModelFormatVersion
	}
	if !IsNil(o.ServiceAccountName) {
		toSerialize["serviceAccountName"] = o.ServiceAccountName
	}
	return toSerialize, nil
}

type NullableModelArtifact struct {
	value *ModelArtifact
	isSet bool
}

func (v NullableModelArtifact) Get() *ModelArtifact {
	return v.value
}

func (v *NullableModelArtifact) Set(val *ModelArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableModelArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableModelArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelArtifact(val *ModelArtifact) *NullableModelArtifact {
	return &NullableModelArtifact{value: val, isSet: true}
}

func (v NullableModelArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
