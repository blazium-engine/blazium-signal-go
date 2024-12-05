# LobbyCreatedResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lobby** | Pointer to [**Lobby**](Lobby.md) |  | [optional] 
**Peers** | Pointer to [**[]Peer**](Peer.md) |  | [optional] 

## Methods

### NewLobbyCreatedResponseData

`func NewLobbyCreatedResponseData() *LobbyCreatedResponseData`

NewLobbyCreatedResponseData instantiates a new LobbyCreatedResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyCreatedResponseDataWithDefaults

`func NewLobbyCreatedResponseDataWithDefaults() *LobbyCreatedResponseData`

NewLobbyCreatedResponseDataWithDefaults instantiates a new LobbyCreatedResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLobby

`func (o *LobbyCreatedResponseData) GetLobby() Lobby`

GetLobby returns the Lobby field if non-nil, zero value otherwise.

### GetLobbyOk

`func (o *LobbyCreatedResponseData) GetLobbyOk() (*Lobby, bool)`

GetLobbyOk returns a tuple with the Lobby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobby

`func (o *LobbyCreatedResponseData) SetLobby(v Lobby)`

SetLobby sets Lobby field to given value.

### HasLobby

`func (o *LobbyCreatedResponseData) HasLobby() bool`

HasLobby returns a boolean if a field has been set.

### GetPeers

`func (o *LobbyCreatedResponseData) GetPeers() []Peer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *LobbyCreatedResponseData) GetPeersOk() (*[]Peer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *LobbyCreatedResponseData) SetPeers(v []Peer)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *LobbyCreatedResponseData) HasPeers() bool`

HasPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


