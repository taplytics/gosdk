# \VariationsApi

All URIs are relative to *https://universal-api.taplytics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VariationGet**](VariationsApi.md#VariationGet) | **Get** /variation | Get Variation for a Taplytics Experiment
[**VariationPost**](VariationsApi.md#VariationPost) | **Post** /variation | Get Variation for a Taplytics Experiment



## VariationGet

> VariationGet(ctx).Token(token).UserId(userId).ExperimentName(experimentName).Attributes(attributes).CustomData(customData).Execute()

Get Variation for a Taplytics Experiment



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
    experimentName := "experimentName_example" // string | Name of an Experiment
    attributes := *openapiclient.NewUserAttributes() // UserAttributes | serialized attributes object
    customData := TODO // map[string]interface{} | serialized custom data object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VariationsApi.VariationGet(context.Background()).Token(token).UserId(userId).ExperimentName(experimentName).Attributes(attributes).CustomData(customData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariationsApi.VariationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **experimentName** | **string** | Name of an Experiment | 
 **attributes** | [**UserAttributes**](UserAttributes.md) | serialized attributes object | 
 **customData** | [**map[string]interface{}**](map[string]interface{}.md) | serialized custom data object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariationPost

> VariationPost(ctx).Token(token).UserId(userId).ExperimentName(experimentName).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get Variation for a Taplytics Experiment



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
    experimentName := "experimentName_example" // string | Name of an Experiment
    userAttributesWithCustomData := *openapiclient.NewUserAttributesWithCustomData() // UserAttributesWithCustomData | All user attributes and optional custom data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VariationsApi.VariationPost(context.Background()).Token(token).UserId(userId).ExperimentName(experimentName).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariationsApi.VariationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **experimentName** | **string** | Name of an Experiment | 
 **userAttributesWithCustomData** | [**UserAttributesWithCustomData**](UserAttributesWithCustomData.md) | All user attributes and optional custom data | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/jason
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

