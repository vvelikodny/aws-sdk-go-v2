// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutAppReplicationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ID of the application tassociated with the replication configuration.
	AppId *string `locationName:"appId" type:"string"`

	// Replication configurations for server groups in the application.
	ServerGroupReplicationConfigurations []ServerGroupReplicationConfiguration `locationName:"serverGroupReplicationConfigurations" type:"list"`
}

// String returns the string representation
func (s PutAppReplicationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

type PutAppReplicationConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAppReplicationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutAppReplicationConfiguration = "PutAppReplicationConfiguration"

// PutAppReplicationConfigurationRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Creates or updates a replication configuration for an application.
//
//    // Example sending a request using PutAppReplicationConfigurationRequest.
//    req := client.PutAppReplicationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/PutAppReplicationConfiguration
func (c *Client) PutAppReplicationConfigurationRequest(input *PutAppReplicationConfigurationInput) PutAppReplicationConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutAppReplicationConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAppReplicationConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutAppReplicationConfigurationOutput{})
	return PutAppReplicationConfigurationRequest{Request: req, Input: input, Copy: c.PutAppReplicationConfigurationRequest}
}

// PutAppReplicationConfigurationRequest is the request type for the
// PutAppReplicationConfiguration API operation.
type PutAppReplicationConfigurationRequest struct {
	*aws.Request
	Input *PutAppReplicationConfigurationInput
	Copy  func(*PutAppReplicationConfigurationInput) PutAppReplicationConfigurationRequest
}

// Send marshals and sends the PutAppReplicationConfiguration API request.
func (r PutAppReplicationConfigurationRequest) Send(ctx context.Context) (*PutAppReplicationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAppReplicationConfigurationResponse{
		PutAppReplicationConfigurationOutput: r.Request.Data.(*PutAppReplicationConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAppReplicationConfigurationResponse is the response type for the
// PutAppReplicationConfiguration API operation.
type PutAppReplicationConfigurationResponse struct {
	*PutAppReplicationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAppReplicationConfiguration request.
func (r *PutAppReplicationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
