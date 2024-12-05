# LobbyCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | The command returned by the server. | [optional] 
**Message** | Pointer to **string** | Informational message about the operation. | [optional] 
**Data** | Pointer to [**LobbyCreatedResponseData**](LobbyCreatedResponseData.md) |  | [optional] 

## Methods

### NewLobbyCreatedResponse

`func NewLobbyCreatedResponse() *LobbyCreatedResponse`

NewLobbyCreatedResponse instantiates a new LobbyCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyCreatedResponseWithDefaults

`func NewLobbyCreatedResponseWithDefaults() *LobbyCreatedResponse`

NewLobbyCreatedResponseWithDefaults instantiates a new LobbyCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *LobbyCreatedResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *LobbyCreatedResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *LobbyCreatedResponse) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *LobbyCreatedResponse) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetMessage

`func (o *LobbyCreatedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LobbyCreatedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LobbyCreatedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LobbyCreatedResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *LobbyCreatedResponse) GetData() LobbyCreatedResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LobbyCreatedResponse) GetDataOk() (*LobbyCreatedResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LobbyCreatedResponse) SetData(v LobbyCreatedResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LobbyCreatedResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


