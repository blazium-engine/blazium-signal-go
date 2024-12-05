# Peer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the peer. | [optional] 
**Ready** | Pointer to **bool** | Indicates if the peer is ready. | [optional] 
**Name** | Pointer to **string** | Name of the peer. | [optional] 

## Methods

### NewPeer

`func NewPeer() *Peer`

NewPeer instantiates a new Peer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerWithDefaults

`func NewPeerWithDefaults() *Peer`

NewPeerWithDefaults instantiates a new Peer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Peer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Peer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Peer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Peer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReady

`func (o *Peer) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Peer) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Peer) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *Peer) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetName

`func (o *Peer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Peer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Peer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Peer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


