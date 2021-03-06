// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type PutResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// The configuration object of the resource in valid JSON format. It must match
	// the schema registered with AWS CloudFormation.
	//
	// The configuration JSON must not exceed 64 KB.
	//
	// Configuration is a required field
	Configuration *string `type:"string" required:"true"`

	// Unique identifier of the resource.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// Name of the resource.
	ResourceName *string `type:"string"`

	// The type of the resource. The custom resource type must be registered with
	// AWS CloudFormation.
	//
	// You cannot use the organization names “aws”, “amzn”, “amazon”,
	// “alexa”, “custom” with custom resource types. It is the first part
	// of the ResourceType up to the first ::.
	//
	// ResourceType is a required field
	ResourceType *string `min:"1" type:"string" required:"true"`

	// Version of the schema registered for the ResourceType in AWS CloudFormation.
	//
	// SchemaVersionId is a required field
	SchemaVersionId *string `min:"1" type:"string" required:"true"`

	// Tags associated with the resource.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s PutResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourceConfigInput"}

	if s.Configuration == nil {
		invalidParams.Add(aws.NewErrParamRequired("Configuration"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if s.ResourceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if s.ResourceType != nil && len(*s.ResourceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceType", 1))
	}

	if s.SchemaVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaVersionId"))
	}
	if s.SchemaVersionId != nil && len(*s.SchemaVersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SchemaVersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutResourceConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutResourceConfig = "PutResourceConfig"

// PutResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Records the configuration state for the resource provided in the request.
// The configuration state of a resource is represented in AWS Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the
// list of configuration items for the custom resource type using existing AWS
// Config APIs.
//
// The custom resource type must be registered with AWS CloudFormation. This
// API accepts the configuration item registered with AWS CloudFormation.
//
// When you call this API, AWS Config only stores configuration state of the
// resource provided in the request. This API does not change or remediate the
// configuration of the resource.
//
//    // Example sending a request using PutResourceConfigRequest.
//    req := client.PutResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutResourceConfig
func (c *Client) PutResourceConfigRequest(input *PutResourceConfigInput) PutResourceConfigRequest {
	op := &aws.Operation{
		Name:       opPutResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutResourceConfigInput{}
	}

	req := c.newRequest(op, input, &PutResourceConfigOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutResourceConfigRequest{Request: req, Input: input, Copy: c.PutResourceConfigRequest}
}

// PutResourceConfigRequest is the request type for the
// PutResourceConfig API operation.
type PutResourceConfigRequest struct {
	*aws.Request
	Input *PutResourceConfigInput
	Copy  func(*PutResourceConfigInput) PutResourceConfigRequest
}

// Send marshals and sends the PutResourceConfig API request.
func (r PutResourceConfigRequest) Send(ctx context.Context) (*PutResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResourceConfigResponse{
		PutResourceConfigOutput: r.Request.Data.(*PutResourceConfigOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResourceConfigResponse is the response type for the
// PutResourceConfig API operation.
type PutResourceConfigResponse struct {
	*PutResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResourceConfig request.
func (r *PutResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
