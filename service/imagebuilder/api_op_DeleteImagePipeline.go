// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteImagePipelineInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image pipeline to delete.
	//
	// ImagePipelineArn is a required field
	ImagePipelineArn *string `location:"querystring" locationName:"imagePipelineArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImagePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImagePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImagePipelineInput"}

	if s.ImagePipelineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImagePipelineArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImagePipelineInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteImagePipelineOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image pipeline that was deleted.
	ImagePipelineArn *string `locationName:"imagePipelineArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteImagePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImagePipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteImagePipeline = "DeleteImagePipeline"

// DeleteImagePipelineRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Deletes an image pipeline.
//
//    // Example sending a request using DeleteImagePipelineRequest.
//    req := client.DeleteImagePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/DeleteImagePipeline
func (c *Client) DeleteImagePipelineRequest(input *DeleteImagePipelineInput) DeleteImagePipelineRequest {
	op := &aws.Operation{
		Name:       opDeleteImagePipeline,
		HTTPMethod: "DELETE",
		HTTPPath:   "/DeleteImagePipeline",
	}

	if input == nil {
		input = &DeleteImagePipelineInput{}
	}

	req := c.newRequest(op, input, &DeleteImagePipelineOutput{})
	return DeleteImagePipelineRequest{Request: req, Input: input, Copy: c.DeleteImagePipelineRequest}
}

// DeleteImagePipelineRequest is the request type for the
// DeleteImagePipeline API operation.
type DeleteImagePipelineRequest struct {
	*aws.Request
	Input *DeleteImagePipelineInput
	Copy  func(*DeleteImagePipelineInput) DeleteImagePipelineRequest
}

// Send marshals and sends the DeleteImagePipeline API request.
func (r DeleteImagePipelineRequest) Send(ctx context.Context) (*DeleteImagePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImagePipelineResponse{
		DeleteImagePipelineOutput: r.Request.Data.(*DeleteImagePipelineOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImagePipelineResponse is the response type for the
// DeleteImagePipeline API operation.
type DeleteImagePipelineResponse struct {
	*DeleteImagePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImagePipeline request.
func (r *DeleteImagePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
