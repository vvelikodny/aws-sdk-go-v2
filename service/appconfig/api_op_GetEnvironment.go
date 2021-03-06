// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the application that includes the environment you want to get.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The ID of the environment you wnat to get.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `location:"uri" locationName:"EnvironmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEnvironmentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEnvironmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentId != nil {
		v := *s.EnvironmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EnvironmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetEnvironmentOutput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	ApplicationId *string `type:"string"`

	// The description of the environment.
	Description *string `type:"string"`

	// The environment ID.
	Id *string `type:"string"`

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []Monitor `type:"list"`

	// The name of the environment.
	Name *string `min:"1" type:"string"`

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State EnvironmentState `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEnvironmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Monitors != nil {
		v := s.Monitors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Monitors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetEnvironment = "GetEnvironment"

// GetEnvironmentRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Retrieve information about an environment. An environment is a logical deployment
// group of AppConfig applications, such as applications in a Production environment
// or in an EU_Region environment. Each configuration deployment targets an
// environment. You can enable one or more Amazon CloudWatch alarms for an environment.
// If an alarm is triggered during a deployment, AppConfig roles back the configuration.
//
//    // Example sending a request using GetEnvironmentRequest.
//    req := client.GetEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/GetEnvironment
func (c *Client) GetEnvironmentRequest(input *GetEnvironmentInput) GetEnvironmentRequest {
	op := &aws.Operation{
		Name:       opGetEnvironment,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{ApplicationId}/environments/{EnvironmentId}",
	}

	if input == nil {
		input = &GetEnvironmentInput{}
	}

	req := c.newRequest(op, input, &GetEnvironmentOutput{})
	return GetEnvironmentRequest{Request: req, Input: input, Copy: c.GetEnvironmentRequest}
}

// GetEnvironmentRequest is the request type for the
// GetEnvironment API operation.
type GetEnvironmentRequest struct {
	*aws.Request
	Input *GetEnvironmentInput
	Copy  func(*GetEnvironmentInput) GetEnvironmentRequest
}

// Send marshals and sends the GetEnvironment API request.
func (r GetEnvironmentRequest) Send(ctx context.Context) (*GetEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEnvironmentResponse{
		GetEnvironmentOutput: r.Request.Data.(*GetEnvironmentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEnvironmentResponse is the response type for the
// GetEnvironment API operation.
type GetEnvironmentResponse struct {
	*GetEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEnvironment request.
func (r *GetEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
