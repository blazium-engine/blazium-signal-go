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

// checks if the LobbyCreatedResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LobbyCreatedResponseData{}

// LobbyCreatedResponseData struct for LobbyCreatedResponseData
type LobbyCreatedResponseData struct {
	Lobby *Lobby `json:"lobby,omitempty"`
	Peers []Peer `json:"peers,omitempty"`
}

// NewLobbyCreatedResponseData instantiates a new LobbyCreatedResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobbyCreatedResponseData() *LobbyCreatedResponseData {
	this := LobbyCreatedResponseData{}
	return &this
}

// NewLobbyCreatedResponseDataWithDefaults instantiates a new LobbyCreatedResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobbyCreatedResponseDataWithDefaults() *LobbyCreatedResponseData {
	this := LobbyCreatedResponseData{}
	return &this
}

// GetLobby returns the Lobby field value if set, zero value otherwise.
func (o *LobbyCreatedResponseData) GetLobby() Lobby {
	if o == nil || IsNil(o.Lobby) {
		var ret Lobby
		return ret
	}
	return *o.Lobby
}

// GetLobbyOk returns a tuple with the Lobby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyCreatedResponseData) GetLobbyOk() (*Lobby, bool) {
	if o == nil || IsNil(o.Lobby) {
		return nil, false
	}
	return o.Lobby, true
}

// HasLobby returns a boolean if a field has been set.
func (o *LobbyCreatedResponseData) HasLobby() bool {
	if o != nil && !IsNil(o.Lobby) {
		return true
	}

	return false
}

// SetLobby gets a reference to the given Lobby and assigns it to the Lobby field.
func (o *LobbyCreatedResponseData) SetLobby(v Lobby) {
	o.Lobby = &v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *LobbyCreatedResponseData) GetPeers() []Peer {
	if o == nil || IsNil(o.Peers) {
		var ret []Peer
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyCreatedResponseData) GetPeersOk() ([]Peer, bool) {
	if o == nil || IsNil(o.Peers) {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *LobbyCreatedResponseData) HasPeers() bool {
	if o != nil && !IsNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []Peer and assigns it to the Peers field.
func (o *LobbyCreatedResponseData) SetPeers(v []Peer) {
	o.Peers = v
}

func (o LobbyCreatedResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LobbyCreatedResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lobby) {
		toSerialize["lobby"] = o.Lobby
	}
	if !IsNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return toSerialize, nil
}

type NullableLobbyCreatedResponseData struct {
	value *LobbyCreatedResponseData
	isSet bool
}

func (v NullableLobbyCreatedResponseData) Get() *LobbyCreatedResponseData {
	return v.value
}

func (v *NullableLobbyCreatedResponseData) Set(val *LobbyCreatedResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableLobbyCreatedResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableLobbyCreatedResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobbyCreatedResponseData(val *LobbyCreatedResponseData) *NullableLobbyCreatedResponseData {
	return &NullableLobbyCreatedResponseData{value: val, isSet: true}
}

func (v NullableLobbyCreatedResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobbyCreatedResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

