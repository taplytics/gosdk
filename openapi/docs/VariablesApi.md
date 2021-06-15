# \VariablesApi

All URIs are relative to *https://universal-api.taplytics.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VariablesGet**](VariablesApi.md#VariablesGet) | **Get** /variables | Get all active variables for user
[**VariablesPost**](VariablesApi.md#VariablesPost) | **Post** /variables | Get all active variables for user
[**VariablevalueGet**](VariablesApi.md#VariablevalueGet) | **Get** /variablevalue | Get value for a Taplytics Variable
[**VariablevaluePost**](VariablesApi.md#VariablevaluePost) | **Post** /variablevalue | Get value for a Taplytics Variable



## VariablesGet

> []Variable VariablesGet(ctx).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Execute()

Get all active variables for user



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
    resp, r, err := api_client.VariablesApi.VariablesGet(context.Background()).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.VariablesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VariablesGet`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.VariablesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **attributes** | [**UserAttributes**](UserAttributes.md) | serialized attributes object | 
 **customData** | [**map[string]interface{}**](map[string]interface{}.md) | serialized custom data object | 

### Return type

[**[]Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariablesPost

> []Variable VariablesPost(ctx).Token(token).UserId(userId).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get all active variables for user



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
    resp, r, err := api_client.VariablesApi.VariablesPost(context.Background()).Token(token).UserId(userId).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.VariablesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VariablesPost`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.VariablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **userAttributesWithCustomData** | [**UserAttributesWithCustomData**](UserAttributesWithCustomData.md) | All user attributes and optional custom data | 

### Return type

[**[]Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/jason
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VariablevalueGet

> VariablevalueGet(ctx).Token(token).UserId(userId).VarName(varName).DefaultValue(defaultValue).Attributes(attributes).CustomData(customData).Execute()

Get value for a Taplytics Variable



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
    varName := "varName_example" // string | name of variable
    defaultValue := "defaultValue_example" // string | default value to be used if user does not have variable available.
    attributes := *openapiclient.NewUserAttributes() // UserAttributes | serialized attributes object
    customData := TODO // map[string]interface{} | serialized custom data object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VariablesApi.VariablevalueGet(context.Background()).Token(token).UserId(userId).VarName(varName).DefaultValue(defaultValue).Attributes(attributes).CustomData(customData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.VariablevalueGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablevalueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **varName** | **string** | name of variable | 
 **defaultValue** | **string** | default value to be used if user does not have variable available. | 
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


## VariablevaluePost

> VariablevaluePost(ctx).Token(token).UserId(userId).VarName(varName).DefaultValue(defaultValue).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()

Get value for a Taplytics Variable



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
    varName := "varName_example" // string | name of variable
    defaultValue := "defaultValue_example" // string | default value to be used if user does not have variable available.
    userAttributesWithCustomData := *openapiclient.NewUserAttributesWithCustomData() // UserAttributesWithCustomData | All user attributes and optional custom data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VariablesApi.VariablevaluePost(context.Background()).Token(token).UserId(userId).VarName(varName).DefaultValue(defaultValue).UserAttributesWithCustomData(userAttributesWithCustomData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.VariablevaluePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVariablevaluePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | SDK token for the project | 
 **userId** | **string** | ID for current user | 
 **varName** | **string** | name of variable | 
 **defaultValue** | **string** | default value to be used if user does not have variable available. | 
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

