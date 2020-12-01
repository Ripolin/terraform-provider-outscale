/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// HealthCheck Information about the health check configuration.
type HealthCheck struct {
	// The number of seconds between two pings (between `5` and `600` both included).
	CheckInterval int32 `json:"CheckInterval"`
	// The number of consecutive successful pings before considering the VM as healthy (between `2` and `10` both included).
	HealthyThreshold int32 `json:"HealthyThreshold"`
	// The path for HTTP or HTTPS requests.
	Path *string `json:"Path,omitempty"`
	// The port number (between `1` and `65535`, both included).
	Port int32 `json:"Port"`
	// The protocol for the URL of the VM (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	Protocol string `json:"Protocol"`
	// The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between `2` and `60` both included).
	Timeout int32 `json:"Timeout"`
	// The number of consecutive failed pings before considering the VM as unhealthy (between `2` and `10` both included).
	UnhealthyThreshold int32 `json:"UnhealthyThreshold"`
}

// NewHealthCheck instantiates a new HealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheck(checkInterval int32, healthyThreshold int32, port int32, protocol string, timeout int32, unhealthyThreshold int32, ) *HealthCheck {
	this := HealthCheck{}
	this.CheckInterval = checkInterval
	this.HealthyThreshold = healthyThreshold
	this.Port = port
	this.Protocol = protocol
	this.Timeout = timeout
	this.UnhealthyThreshold = unhealthyThreshold
	return &this
}

// NewHealthCheckWithDefaults instantiates a new HealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckWithDefaults() *HealthCheck {
	this := HealthCheck{}
	return &this
}

// GetCheckInterval returns the CheckInterval field value
func (o *HealthCheck) GetCheckInterval() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetCheckIntervalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckInterval, true
}

// SetCheckInterval sets field value
func (o *HealthCheck) SetCheckInterval(v int32) {
	o.CheckInterval = v
}

// GetHealthyThreshold returns the HealthyThreshold field value
func (o *HealthCheck) GetHealthyThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.HealthyThreshold
}

// GetHealthyThresholdOk returns a tuple with the HealthyThreshold field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetHealthyThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HealthyThreshold, true
}

// SetHealthyThreshold sets field value
func (o *HealthCheck) SetHealthyThreshold(v int32) {
	o.HealthyThreshold = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HealthCheck) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HealthCheck) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HealthCheck) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value
func (o *HealthCheck) GetPort() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *HealthCheck) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value
func (o *HealthCheck) GetProtocol() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *HealthCheck) SetProtocol(v string) {
	o.Protocol = v
}

// GetTimeout returns the Timeout field value
func (o *HealthCheck) GetTimeout() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *HealthCheck) SetTimeout(v int32) {
	o.Timeout = v
}

// GetUnhealthyThreshold returns the UnhealthyThreshold field value
func (o *HealthCheck) GetUnhealthyThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.UnhealthyThreshold
}

// GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetUnhealthyThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnhealthyThreshold, true
}

// SetUnhealthyThreshold sets field value
func (o *HealthCheck) SetUnhealthyThreshold(v int32) {
	o.UnhealthyThreshold = v
}

func (o HealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CheckInterval"] = o.CheckInterval
	}
	if true {
		toSerialize["HealthyThreshold"] = o.HealthyThreshold
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if true {
		toSerialize["Port"] = o.Port
	}
	if true {
		toSerialize["Protocol"] = o.Protocol
	}
	if true {
		toSerialize["Timeout"] = o.Timeout
	}
	if true {
		toSerialize["UnhealthyThreshold"] = o.UnhealthyThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableHealthCheck struct {
	value *HealthCheck
	isSet bool
}

func (v NullableHealthCheck) Get() *HealthCheck {
	return v.value
}

func (v *NullableHealthCheck) Set(val *HealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheck(val *HealthCheck) *NullableHealthCheck {
	return &NullableHealthCheck{value: val, isSet: true}
}

func (v NullableHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

