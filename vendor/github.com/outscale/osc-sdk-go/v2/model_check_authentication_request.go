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

// CheckAuthenticationRequest struct for CheckAuthenticationRequest
type CheckAuthenticationRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The email address of the account.
	Login *string `json:"Login,omitempty"`
	// The password of the account.
	Password *string `json:"Password,omitempty"`
}

// NewCheckAuthenticationRequest instantiates a new CheckAuthenticationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAuthenticationRequest() *CheckAuthenticationRequest {
	this := CheckAuthenticationRequest{}
	return &this
}

// NewCheckAuthenticationRequestWithDefaults instantiates a new CheckAuthenticationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAuthenticationRequestWithDefaults() *CheckAuthenticationRequest {
	this := CheckAuthenticationRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CheckAuthenticationRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAuthenticationRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CheckAuthenticationRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CheckAuthenticationRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *CheckAuthenticationRequest) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAuthenticationRequest) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *CheckAuthenticationRequest) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *CheckAuthenticationRequest) SetLogin(v string) {
	o.Login = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CheckAuthenticationRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAuthenticationRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CheckAuthenticationRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CheckAuthenticationRequest) SetPassword(v string) {
	o.Password = &v
}

func (o CheckAuthenticationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Login != nil {
		toSerialize["Login"] = o.Login
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCheckAuthenticationRequest struct {
	value *CheckAuthenticationRequest
	isSet bool
}

func (v NullableCheckAuthenticationRequest) Get() *CheckAuthenticationRequest {
	return v.value
}

func (v *NullableCheckAuthenticationRequest) Set(val *CheckAuthenticationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAuthenticationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAuthenticationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAuthenticationRequest(val *CheckAuthenticationRequest) *NullableCheckAuthenticationRequest {
	return &NullableCheckAuthenticationRequest{value: val, isSet: true}
}

func (v NullableCheckAuthenticationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAuthenticationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
