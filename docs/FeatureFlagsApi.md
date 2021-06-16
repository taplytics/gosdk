# \FeatureFlagsApi

All URIs are relative to *https://universal-api.taplytics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeatureflagsGet**](FeatureFlagsApi.md#FeatureflagsGet) | **Get** /featureflags | Get enabled Feature Flags for the user
[**FeatureflagsPost**](FeatureFlagsApi.md#FeatureflagsPost) | **Post** /featureflags | Get enabled Feature Flags for the user
[**IsfeatureflagenabledGet**](FeatureFlagsApi.md#IsfeatureflagenabledGet) | **Get** /isfeatureflagenabled | Get if feature flag is enabled
[**IsfeatureflagenabledPost**](FeatureFlagsApi.md#IsfeatureflagenabledPost) | **Post** /isfeatureflagenabled | Get if feature flag is enabled



## FeatureflagsGet

> FeatureflagsGet(ctx).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Execute()

Get enabled Feature Flags for the user



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.FeatureflagsGet(context.Background()).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.FeatureflagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeatureflagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **attributes** | [**UserAttributes**](UserAttributes.md) | serialized attributes object | 
 **customData** | [**map[string]interface{}**](map[string]interface{}.md) | serialized custom data object | 

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


## FeatureflagsPost

> FeatureflagsPost(ctx).Token(token).UserId(userId).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get enabled Feature Flags for the user



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
    userAttributesWithCustomData := *openapiclient.NewUserAttributesWithCustomData() // UserAttributesWithCustomData | All user attributes and optional custom data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.FeatureflagsPost(context.Background()).Token(token).UserId(userId).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.FeatureflagsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeatureflagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
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


## IsfeatureflagenabledGet

> IsfeatureflagenabledGet(ctx).Token(token).UserId(userId).FeatureFlagKey(featureFlagKey).Attributes(attributes).CustomData(customData).Execute()

Get if feature flag is enabled



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
    featureFlagKey := "featureFlagKey_example" // string | key for the feature flag you want to check
    attributes := *openapiclient.NewUserAttributes() // UserAttributes | serialized attributes object
    customData := TODO // map[string]interface{} | serialized custom data object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.IsfeatureflagenabledGet(context.Background()).Token(token).UserId(userId).FeatureFlagKey(featureFlagKey).Attributes(attributes).CustomData(customData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.IsfeatureflagenabledGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsfeatureflagenabledGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **featureFlagKey** | **string** | key for the feature flag you want to check | 
 **attributes** | [**UserAttributes**](UserAttributes.md) | serialized attributes object | 
 **customData** | [**map[string]interface{}**](map[string]interface{}.md) | serialized custom data object | 

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


## IsfeatureflagenabledPost

> IsfeatureflagenabledPost(ctx).Token(token).UserId(userId).FeatureFlagKey(featureFlagKey).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get if feature flag is enabled



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
    featureFlagKey := "featureFlagKey_example" // string | key for the feature flag you want to check
    userAttributesWithCustomData := *openapiclient.NewUserAttributesWithCustomData() // UserAttributesWithCustomData | All user attributes and optional custom data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.IsfeatureflagenabledPost(context.Background()).Token(token).UserId(userId).FeatureFlagKey(featureFlagKey).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.IsfeatureflagenabledPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsfeatureflagenabledPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **featureFlagKey** | **string** | key for the feature flag you want to check | 
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

