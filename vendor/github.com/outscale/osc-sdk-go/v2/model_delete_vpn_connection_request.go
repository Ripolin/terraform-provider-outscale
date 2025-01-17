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

// DeleteVpnConnectionRequest struct for DeleteVpnConnectionRequest
type DeleteVpnConnectionRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the VPN connection you want to delete.
	VpnConnectionId string `json:"VpnConnectionId"`
}

// NewDeleteVpnConnectionRequest instantiates a new DeleteVpnConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVpnConnectionRequest(vpnConnectionId string) *DeleteVpnConnectionRequest {
	this := DeleteVpnConnectionRequest{}
	this.VpnConnectionId = vpnConnectionId
	return &this
}

// NewDeleteVpnConnectionRequestWithDefaults instantiates a new DeleteVpnConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVpnConnectionRequestWithDefaults() *DeleteVpnConnectionRequest {
	this := DeleteVpnConnectionRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteVpnConnectionRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVpnConnectionRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteVpnConnectionRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteVpnConnectionRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVpnConnectionId returns the VpnConnectionId field value
func (o *DeleteVpnConnectionRequest) GetVpnConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnConnectionId
}

// GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field value
// and a boolean to check if the value has been set.
func (o *DeleteVpnConnectionRequest) GetVpnConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnConnectionId, true
}

// SetVpnConnectionId sets field value
func (o *DeleteVpnConnectionRequest) SetVpnConnectionId(v string) {
	o.VpnConnectionId = v
}

func (o DeleteVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VpnConnectionId"] = o.VpnConnectionId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVpnConnectionRequest struct {
	value *DeleteVpnConnectionRequest
	isSet bool
}

func (v NullableDeleteVpnConnectionRequest) Get() *DeleteVpnConnectionRequest {
	return v.value
}

func (v *NullableDeleteVpnConnectionRequest) Set(val *DeleteVpnConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVpnConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVpnConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVpnConnectionRequest(val *DeleteVpnConnectionRequest) *NullableDeleteVpnConnectionRequest {
	return &NullableDeleteVpnConnectionRequest{value: val, isSet: true}
}

func (v NullableDeleteVpnConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVpnConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
