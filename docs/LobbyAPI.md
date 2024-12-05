# \LobbyAPI

All URIs are relative to *https://lobby.blazium.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LobbyCreate**](LobbyAPI.md#LobbyCreate) | **Post** /lobby/create | Create a new game lobby.



## LobbyCreate

> LobbyCreatedResponse LobbyCreate(ctx).LobbyCreateRequest(lobbyCreateRequest).Execute()

Create a new game lobby.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	lobbyCreateRequest := *openapiclient.NewLobbyCreateRequest("Hangman", int32(4)) // LobbyCreateRequest | JSON payload containing lobby creation details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LobbyAPI.LobbyCreate(context.Background()).LobbyCreateRequest(lobbyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LobbyAPI.LobbyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LobbyCreate`: LobbyCreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `LobbyAPI.LobbyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLobbyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lobbyCreateRequest** | [**LobbyCreateRequest**](LobbyCreateRequest.md) | JSON payload containing lobby creation details. | 

### Return type

[**LobbyCreatedResponse**](LobbyCreatedResponse.md)

### Authorization

[GameSession](../README.md#GameSession), [GameId](../README.md#GameId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

