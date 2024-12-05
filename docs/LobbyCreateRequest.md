# LobbyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the lobby. | 
**Password** | Pointer to **string** | Optional password for the lobby. | [optional] 
**MaxPlayers** | **int32** | Maximum number of players allowed in the lobby. | 
**Tags** | Pointer to **[]string** | Tags associated with the lobby. | [optional] 

## Methods

### NewLobbyCreateRequest

`func NewLobbyCreateRequest(name string, maxPlayers int32, ) *LobbyCreateRequest`

NewLobbyCreateRequest instantiates a new LobbyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyCreateRequestWithDefaults

`func NewLobbyCreateRequestWithDefaults() *LobbyCreateRequest`

NewLobbyCreateRequestWithDefaults instantiates a new LobbyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LobbyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LobbyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LobbyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *LobbyCreateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LobbyCreateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LobbyCreateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LobbyCreateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *LobbyCreateRequest) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *LobbyCreateRequest) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *LobbyCreateRequest) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.


### GetTags

`func (o *LobbyCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LobbyCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LobbyCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LobbyCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


