/*
Blazium API

API for managing lobbies and game sessions.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LobbyCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LobbyCreateRequest{}

// LobbyCreateRequest struct for LobbyCreateRequest
type LobbyCreateRequest struct {
	// The name of the lobby.
	Name string `json:"name"`
	// Optional password for the lobby.
	Password *string `json:"password,omitempty"`
	// Maximum number of players allowed in the lobby.
	MaxPlayers int32 `json:"max_players"`
	// Tags associated with the lobby.
	Tags []string `json:"tags,omitempty"`
}

type _LobbyCreateRequest LobbyCreateRequest

// NewLobbyCreateRequest instantiates a new LobbyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobbyCreateRequest(name string, maxPlayers int32) *LobbyCreateRequest {
	this := LobbyCreateRequest{}
	this.Name = name
	this.MaxPlayers = maxPlayers
	return &this
}

// NewLobbyCreateRequestWithDefaults instantiates a new LobbyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobbyCreateRequestWithDefaults() *LobbyCreateRequest {
	this := LobbyCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *LobbyCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LobbyCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LobbyCreateRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LobbyCreateRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyCreateRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LobbyCreateRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LobbyCreateRequest) SetPassword(v string) {
	o.Password = &v
}

// GetMaxPlayers returns the MaxPlayers field value
func (o *LobbyCreateRequest) GetMaxPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPlayers
}

// GetMaxPlayersOk returns a tuple with the MaxPlayers field value
// and a boolean to check if the value has been set.
func (o *LobbyCreateRequest) GetMaxPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPlayers, true
}

// SetMaxPlayers sets field value
func (o *LobbyCreateRequest) SetMaxPlayers(v int32) {
	o.MaxPlayers = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LobbyCreateRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyCreateRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LobbyCreateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LobbyCreateRequest) SetTags(v []string) {
	o.Tags = v
}

func (o LobbyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LobbyCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["max_players"] = o.MaxPlayers
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *LobbyCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"max_players",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLobbyCreateRequest := _LobbyCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLobbyCreateRequest)

	if err != nil {
		return err
	}

	*o = LobbyCreateRequest(varLobbyCreateRequest)

	return err
}

type NullableLobbyCreateRequest struct {
	value *LobbyCreateRequest
	isSet bool
}

func (v NullableLobbyCreateRequest) Get() *LobbyCreateRequest {
	return v.value
}

func (v *NullableLobbyCreateRequest) Set(val *LobbyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLobbyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLobbyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobbyCreateRequest(val *LobbyCreateRequest) *NullableLobbyCreateRequest {
	return &NullableLobbyCreateRequest{value: val, isSet: true}
}

func (v NullableLobbyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobbyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

