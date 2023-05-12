/*
Nextlinux Engine API Server

This is the Nextlinux Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// PaginatedVulnerableImageListAllOf struct for PaginatedVulnerableImageListAllOf
type PaginatedVulnerableImageListAllOf struct {
	Images *[]VulnerableImage `json:"images,omitempty"`
}

// NewPaginatedVulnerableImageListAllOf instantiates a new PaginatedVulnerableImageListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedVulnerableImageListAllOf() *PaginatedVulnerableImageListAllOf {
	this := PaginatedVulnerableImageListAllOf{}
	return &this
}

// NewPaginatedVulnerableImageListAllOfWithDefaults instantiates a new PaginatedVulnerableImageListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedVulnerableImageListAllOfWithDefaults() *PaginatedVulnerableImageListAllOf {
	this := PaginatedVulnerableImageListAllOf{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *PaginatedVulnerableImageListAllOf) GetImages() []VulnerableImage {
	if o == nil || o.Images == nil {
		var ret []VulnerableImage
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedVulnerableImageListAllOf) GetImagesOk() (*[]VulnerableImage, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *PaginatedVulnerableImageListAllOf) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []VulnerableImage and assigns it to the Images field.
func (o *PaginatedVulnerableImageListAllOf) SetImages(v []VulnerableImage) {
	o.Images = &v
}

func (o PaginatedVulnerableImageListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedVulnerableImageListAllOf struct {
	value *PaginatedVulnerableImageListAllOf
	isSet bool
}

func (v NullablePaginatedVulnerableImageListAllOf) Get() *PaginatedVulnerableImageListAllOf {
	return v.value
}

func (v *NullablePaginatedVulnerableImageListAllOf) Set(val *PaginatedVulnerableImageListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedVulnerableImageListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedVulnerableImageListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedVulnerableImageListAllOf(val *PaginatedVulnerableImageListAllOf) *NullablePaginatedVulnerableImageListAllOf {
	return &NullablePaginatedVulnerableImageListAllOf{value: val, isSet: true}
}

func (v NullablePaginatedVulnerableImageListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedVulnerableImageListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

