// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInstanceInformationInput struct {
	_ struct{} `type:"structure"`

	// One or more filters. Use a filter to return a more specific list of instances.
	// You can filter on Amazon EC2 tag. Specify tags by using a key-value mapping.
	Filters []InstanceInformationStringFilter `type:"list"`

	// This is a legacy method. We recommend that you don't use this method. Instead,
	// use the InstanceInformationFilter action. The InstanceInformationFilter action
	// enables you to return instance information by using tags that are specified
	// as a key-value mapping.
	//
	// If you do use this method, then you can't use the InstanceInformationFilter
	// action. Using this method and the InstanceInformationFilter action causes
	// an exception error.
	InstanceInformationFilterList []InstanceInformationFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceInformationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceInformationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceInformationInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.InstanceInformationFilterList != nil {
		for i, v := range s.InstanceInformationFilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InstanceInformationFilterList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInstanceInformationOutput struct {
	_ struct{} `type:"structure"`

	// The instance information list.
	InstanceInformationList []InstanceInformation `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceInformationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstanceInformation = "DescribeInstanceInformation"

// DescribeInstanceInformationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Describes one or more of your instances. You can use this to get information
// about instances like the operating system platform, the SSM Agent version
// (Linux), status etc. If you specify one or more instance IDs, it returns
// information for those instances. If you do not specify instance IDs, it returns
// information for all your instances. If you specify an instance ID that is
// not valid or an instance that you do not own, you receive an error.
//
// The IamRole field for this API action is the Amazon Identity and Access Management
// (IAM) role assigned to on-premises instances. This call does not return the
// IAM role for Amazon EC2 instances.
//
//    // Example sending a request using DescribeInstanceInformationRequest.
//    req := client.DescribeInstanceInformationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstanceInformation
func (c *Client) DescribeInstanceInformationRequest(input *DescribeInstanceInformationInput) DescribeInstanceInformationRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceInformation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeInstanceInformationInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceInformationOutput{})
	return DescribeInstanceInformationRequest{Request: req, Input: input, Copy: c.DescribeInstanceInformationRequest}
}

// DescribeInstanceInformationRequest is the request type for the
// DescribeInstanceInformation API operation.
type DescribeInstanceInformationRequest struct {
	*aws.Request
	Input *DescribeInstanceInformationInput
	Copy  func(*DescribeInstanceInformationInput) DescribeInstanceInformationRequest
}

// Send marshals and sends the DescribeInstanceInformation API request.
func (r DescribeInstanceInformationRequest) Send(ctx context.Context) (*DescribeInstanceInformationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceInformationResponse{
		DescribeInstanceInformationOutput: r.Request.Data.(*DescribeInstanceInformationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInstanceInformationRequestPaginator returns a paginator for DescribeInstanceInformation.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInstanceInformationRequest(input)
//   p := ssm.NewDescribeInstanceInformationRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInstanceInformationPaginator(req DescribeInstanceInformationRequest) DescribeInstanceInformationPaginator {
	return DescribeInstanceInformationPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInstanceInformationInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeInstanceInformationPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInstanceInformationPaginator struct {
	aws.Pager
}

func (p *DescribeInstanceInformationPaginator) CurrentPage() *DescribeInstanceInformationOutput {
	return p.Pager.CurrentPage().(*DescribeInstanceInformationOutput)
}

// DescribeInstanceInformationResponse is the response type for the
// DescribeInstanceInformation API operation.
type DescribeInstanceInformationResponse struct {
	*DescribeInstanceInformationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceInformation request.
func (r *DescribeInstanceInformationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
