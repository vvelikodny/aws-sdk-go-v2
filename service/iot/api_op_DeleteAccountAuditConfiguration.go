// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteAccountAuditConfigurationInput struct {
	_ struct{} `type:"structure"`

	// If true, all scheduled audits are deleted.
	DeleteScheduledAudits *bool `location:"querystring" locationName:"deleteScheduledAudits" type:"boolean"`
}

// String returns the string representation
func (s DeleteAccountAuditConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccountAuditConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeleteScheduledAudits != nil {
		v := *s.DeleteScheduledAudits

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "deleteScheduledAudits", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DeleteAccountAuditConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccountAuditConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccountAuditConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAccountAuditConfiguration = "DeleteAccountAuditConfiguration"

// DeleteAccountAuditConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Restores the default settings for Device Defender audits for this account.
// Any configuration data you entered is deleted and all audit checks are reset
// to disabled.
//
//    // Example sending a request using DeleteAccountAuditConfigurationRequest.
//    req := client.DeleteAccountAuditConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteAccountAuditConfigurationRequest(input *DeleteAccountAuditConfigurationInput) DeleteAccountAuditConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteAccountAuditConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/audit/configuration",
	}

	if input == nil {
		input = &DeleteAccountAuditConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteAccountAuditConfigurationOutput{})
	return DeleteAccountAuditConfigurationRequest{Request: req, Input: input, Copy: c.DeleteAccountAuditConfigurationRequest}
}

// DeleteAccountAuditConfigurationRequest is the request type for the
// DeleteAccountAuditConfiguration API operation.
type DeleteAccountAuditConfigurationRequest struct {
	*aws.Request
	Input *DeleteAccountAuditConfigurationInput
	Copy  func(*DeleteAccountAuditConfigurationInput) DeleteAccountAuditConfigurationRequest
}

// Send marshals and sends the DeleteAccountAuditConfiguration API request.
func (r DeleteAccountAuditConfigurationRequest) Send(ctx context.Context) (*DeleteAccountAuditConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccountAuditConfigurationResponse{
		DeleteAccountAuditConfigurationOutput: r.Request.Data.(*DeleteAccountAuditConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccountAuditConfigurationResponse is the response type for the
// DeleteAccountAuditConfiguration API operation.
type DeleteAccountAuditConfigurationResponse struct {
	*DeleteAccountAuditConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccountAuditConfiguration request.
func (r *DeleteAccountAuditConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}