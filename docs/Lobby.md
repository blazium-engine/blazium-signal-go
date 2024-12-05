# Lobby

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the lobby. | [optional] 
**Name** | Pointer to **string** | Name of the lobby. | [optional] 
**HostId** | Pointer to **string** | ID of the lobby host. | [optional] 
**HostName** | Pointer to **string** | Name of the lobby host. | [optional] 
**Sealed** | Pointer to **bool** | Indicates if the lobby is sealed. | [optional] 
**MaxPlayers** | Pointer to **int32** | Maximum number of players allowed in the lobby. | [optional] 
**Players** | Pointer to **int32** | Current number of players in the lobby. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of when the lobby was created. | [optional] 
**HasPassword** | Pointer to **bool** | Indicates if the lobby has a password. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the lobby. | [optional] 

## Methods

### NewLobby

`func NewLobby() *Lobby`

NewLobby instantiates a new Lobby object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyWithDefaults

`func NewLobbyWithDefaults() *Lobby`

NewLobbyWithDefaults instantiates a new Lobby object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Lobby) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lobby) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lobby) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Lobby) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Lobby) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Lobby) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Lobby) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Lobby) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostId

`func (o *Lobby) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *Lobby) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *Lobby) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *Lobby) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHostName

`func (o *Lobby) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Lobby) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Lobby) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Lobby) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetSealed

`func (o *Lobby) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *Lobby) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *Lobby) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *Lobby) HasSealed() bool`

HasSealed returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *Lobby) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *Lobby) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *Lobby) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *Lobby) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetPlayers

`func (o *Lobby) GetPlayers() int32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *Lobby) GetPlayersOk() (*int32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *Lobby) SetPlayers(v int32)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *Lobby) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Lobby) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Lobby) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Lobby) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Lobby) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHasPassword

`func (o *Lobby) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *Lobby) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *Lobby) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *Lobby) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetTags

`func (o *Lobby) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Lobby) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Lobby) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Lobby) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


