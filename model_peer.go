/*
Blazium API

API for managing lobbies and game sessions.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Peer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Peer{}

// Peer struct for Peer
type Peer struct {
	// Unique identifier for the peer.
	Id *string `json:"id,omitempty"`
	// Indicates if the peer is ready.
	Ready *bool `json:"ready,omitempty"`
	// Name of the peer.
	Name *string `json:"name,omitempty"`
}

// NewPeer instantiates a new Peer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeer() *Peer {
	this := Peer{}
	return &this
}

// NewPeerWithDefaults instantiates a new Peer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeerWithDefaults() *Peer {
	this := Peer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Peer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Peer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Peer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Peer) SetId(v string) {
	o.Id = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *Peer) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Peer) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *Peer) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *Peer) SetReady(v bool) {
	o.Ready = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Peer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Peer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Peer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Peer) SetName(v string) {
	o.Name = &v
}

func (o Peer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Peer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePeer struct {
	value *Peer
	isSet bool
}

func (v NullablePeer) Get() *Peer {
	return v.value
}

func (v *NullablePeer) Set(val *Peer) {
	v.value = val
	v.isSet = true
}

func (v NullablePeer) IsSet() bool {
	return v.isSet
}

func (v *NullablePeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeer(val *Peer) *NullablePeer {
	return &NullablePeer{value: val, isSet: true}
}

func (v NullablePeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


