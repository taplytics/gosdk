# \BucketingApi

All URIs are relative to *https://universal-api.taplytics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketingGet**](BucketingApi.md#BucketingGet) | **Get** /bucketing | Get Experiments and Variations for the user.
[**BucketingPost**](BucketingApi.md#BucketingPost) | **Post** /bucketing | Get Experiments and Variations for the user.



## BucketingGet

> BucketingGet(ctx).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Verbose(verbose).Execute()

Get Experiments and Variations for the user.



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
    attributes := *openapiclient.NewUserAttributes() // UserAttributes | serialized attributes object
    customData := TODO // map[string]interface{} | serialized custom data object
    verbose := true // bool | Flag to return object of experiments and variations with variables (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketingApi.BucketingGet(context.Background()).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketingApi.BucketingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **attributes** | [**UserAttributes**](UserAttributes.md) | serialized attributes object | 
 **customData** | [**map[string]interface{}**](map[string]interface{}.md) | serialized custom data object | 
 **verbose** | **bool** | Flag to return object of experiments and variations with variables | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketingPost

> BucketingPost(ctx).Token(token).UserId(userId).Verbose(verbose).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get Experiments and Variations for the user.



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
    verbose := true // bool | Flag to return object of experiments and variations with variables (optional)
    userAttributesWithCustomData := *openapiclient.NewUserAttributesWithCustomData() // UserAttributesWithCustomData | All user attributes and optional custom data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketingApi.BucketingPost(context.Background()).Token(token).UserId(userId).Verbose(verbose).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketingApi.BucketingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **verbose** | **bool** | Flag to return object of experiments and variations with variables | 
 **userAttributesWithCustomData** | [**UserAttributesWithCustomData**](UserAttributesWithCustomData.md) | All user attributes and optional custom data | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/jason
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

