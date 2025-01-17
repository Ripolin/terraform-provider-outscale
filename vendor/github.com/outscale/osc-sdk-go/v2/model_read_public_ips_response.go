/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadPublicIpsResponse struct for ReadPublicIpsResponse
type ReadPublicIpsResponse struct {
	// Information about one or more EIPs.
	PublicIps       *[]PublicIp      `json:"PublicIps,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadPublicIpsResponse instantiates a new ReadPublicIpsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadPublicIpsResponse() *ReadPublicIpsResponse {
	this := ReadPublicIpsResponse{}
	return &this
}

// NewReadPublicIpsResponseWithDefaults instantiates a new ReadPublicIpsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadPublicIpsResponseWithDefaults() *ReadPublicIpsResponse {
	this := ReadPublicIpsResponse{}
	return &this
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *ReadPublicIpsResponse) GetPublicIps() []PublicIp {
	if o == nil || o.PublicIps == nil {
		var ret []PublicIp
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpsResponse) GetPublicIpsOk() (*[]PublicIp, bool) {
	if o == nil || o.PublicIps == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *ReadPublicIpsResponse) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []PublicIp and assigns it to the PublicIps field.
func (o *ReadPublicIpsResponse) SetPublicIps(v []PublicIp) {
	o.PublicIps = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadPublicIpsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPublicIpsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadPublicIpsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadPublicIpsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadPublicIpsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicIps != nil {
		toSerialize["PublicIps"] = o.PublicIps
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadPublicIpsResponse struct {
	value *ReadPublicIpsResponse
	isSet bool
}

func (v NullableReadPublicIpsResponse) Get() *ReadPublicIpsResponse {
	return v.value
}

func (v *NullableReadPublicIpsResponse) Set(val *ReadPublicIpsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadPublicIpsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadPublicIpsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadPublicIpsResponse(val *ReadPublicIpsResponse) *NullableReadPublicIpsResponse {
	return &NullableReadPublicIpsResponse{value: val, isSet: true}
}

func (v NullableReadPublicIpsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadPublicIpsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
