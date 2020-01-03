// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to get a SqlInjectionMatchSet.
type GetSqlInjectionMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to get.
	// SqlInjectionMatchSetId is returned by CreateSqlInjectionMatchSet and by ListSqlInjectionMatchSets.
	//
	// SqlInjectionMatchSetId is a required field
	SqlInjectionMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSqlInjectionMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSqlInjectionMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSqlInjectionMatchSetInput"}

	if s.SqlInjectionMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SqlInjectionMatchSetId"))
	}
	if s.SqlInjectionMatchSetId != nil && len(*s.SqlInjectionMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SqlInjectionMatchSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a GetSqlInjectionMatchSet request.
type GetSqlInjectionMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the SqlInjectionMatchSet that you specified in the GetSqlInjectionMatchSet
	// request. For more information, see the following topics:
	//
	//    * SqlInjectionMatchSet: Contains Name, SqlInjectionMatchSetId, and an
	//    array of SqlInjectionMatchTuple objects
	//
	//    * SqlInjectionMatchTuple: Each SqlInjectionMatchTuple object contains
	//    FieldToMatch and TextTransformation
	//
	//    * FieldToMatch: Contains Data and Type
	SqlInjectionMatchSet *SqlInjectionMatchSet `type:"structure"`
}

// String returns the string representation
func (s GetSqlInjectionMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSqlInjectionMatchSet = "GetSqlInjectionMatchSet"

// GetSqlInjectionMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the SqlInjectionMatchSet that is specified by SqlInjectionMatchSetId.
//
//    // Example sending a request using GetSqlInjectionMatchSetRequest.
//    req := client.GetSqlInjectionMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetSqlInjectionMatchSet
func (c *Client) GetSqlInjectionMatchSetRequest(input *GetSqlInjectionMatchSetInput) GetSqlInjectionMatchSetRequest {
	op := &aws.Operation{
		Name:       opGetSqlInjectionMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSqlInjectionMatchSetInput{}
	}

	req := c.newRequest(op, input, &GetSqlInjectionMatchSetOutput{})
	return GetSqlInjectionMatchSetRequest{Request: req, Input: input, Copy: c.GetSqlInjectionMatchSetRequest}
}

// GetSqlInjectionMatchSetRequest is the request type for the
// GetSqlInjectionMatchSet API operation.
type GetSqlInjectionMatchSetRequest struct {
	*aws.Request
	Input *GetSqlInjectionMatchSetInput
	Copy  func(*GetSqlInjectionMatchSetInput) GetSqlInjectionMatchSetRequest
}

// Send marshals and sends the GetSqlInjectionMatchSet API request.
func (r GetSqlInjectionMatchSetRequest) Send(ctx context.Context) (*GetSqlInjectionMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSqlInjectionMatchSetResponse{
		GetSqlInjectionMatchSetOutput: r.Request.Data.(*GetSqlInjectionMatchSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSqlInjectionMatchSetResponse is the response type for the
// GetSqlInjectionMatchSet API operation.
type GetSqlInjectionMatchSetResponse struct {
	*GetSqlInjectionMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSqlInjectionMatchSet request.
func (r *GetSqlInjectionMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
