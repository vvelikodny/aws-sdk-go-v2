// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetEntitiesInput struct {
	_ struct{} `type:"structure"`

	// An array of entity IDs.
	//
	// The IDs should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	//
	// Ids is a required field
	Ids []string `locationName:"ids" type:"list" required:"true"`

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64 `locationName:"namespaceVersion" type:"long"`
}

// String returns the string representation
func (s GetEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEntitiesInput"}

	if s.Ids == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ids"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// An array of descriptions for the specified entities.
	Descriptions []EntityDescription `locationName:"descriptions" type:"list"`
}

// String returns the string representation
func (s GetEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEntities = "GetEntities"

// GetEntitiesRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Gets definitions of the specified entities. Uses the latest version of the
// user's namespace by default. This API returns the following TDM entities.
//
//    * Properties
//
//    * States
//
//    * Events
//
//    * Actions
//
//    * Capabilities
//
//    * Mappings
//
//    * Devices
//
//    * Device Models
//
//    * Services
//
// This action doesn't return definitions for systems, flows, and deployments.
//
//    // Example sending a request using GetEntitiesRequest.
//    req := client.GetEntitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetEntities
func (c *Client) GetEntitiesRequest(input *GetEntitiesInput) GetEntitiesRequest {
	op := &aws.Operation{
		Name:       opGetEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEntitiesInput{}
	}

	req := c.newRequest(op, input, &GetEntitiesOutput{})
	return GetEntitiesRequest{Request: req, Input: input, Copy: c.GetEntitiesRequest}
}

// GetEntitiesRequest is the request type for the
// GetEntities API operation.
type GetEntitiesRequest struct {
	*aws.Request
	Input *GetEntitiesInput
	Copy  func(*GetEntitiesInput) GetEntitiesRequest
}

// Send marshals and sends the GetEntities API request.
func (r GetEntitiesRequest) Send(ctx context.Context) (*GetEntitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEntitiesResponse{
		GetEntitiesOutput: r.Request.Data.(*GetEntitiesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEntitiesResponse is the response type for the
// GetEntities API operation.
type GetEntitiesResponse struct {
	*GetEntitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEntities request.
func (r *GetEntitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
