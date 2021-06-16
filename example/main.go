package main

import (
	"context"
	"fmt"
	"os"

	openapi "github.com/taplytics/go_sdk"
)

func main() {

	token := "SDK TOKEN"                       // string | SDK token for the project
	userId := "USER ID"                        // string | ID for current user
	attributes := *openapi.NewUserAttributes() // UserAttributes | serialized attributes object
	customData := make(map[string]interface{}) // map[string]interface{} | serialized custom data object

	configuration := openapi.NewConfiguration()
	api_client := openapi.NewAPIClient(configuration)
	resp, err := api_client.FeatureFlagsApi.FeatureflagsGet(context.Background()).Token(token).UserId(userId).Attributes(attributes).CustomData(customData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.FeatureflagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}

	fmt.Print(resp.Body)
}
