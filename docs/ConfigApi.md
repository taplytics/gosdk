# \ConfigApi

All URIs are relative to *https://universal-api.taplytics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigGet**](ConfigApi.md#ConfigGet) | **Get** /config | Get Verbose Project Config Document for User
[**ConfigPost**](ConfigApi.md#ConfigPost) | **Post** /config | Get Verbose Project Config Document for User



## ConfigGet

> map[string]interface{} ConfigGet(ctx).Token(token).UserId(userId).Execute()

Get Verbose Project Config Document for User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | SDK token for the project
    userId := "userId_example" // string | ID for current user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ConfigGet(context.Background()).Token(token).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ConfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ConfigGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPost

> map[string]interface{} ConfigPost(ctx).Token(token).UserId(userId).Execute()

Get Verbose Project Config Document for User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | SDK token for the project
    userId := "userId_example" // string | ID for current user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ConfigPost(context.Background()).Token(token).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ConfigPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.ConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

