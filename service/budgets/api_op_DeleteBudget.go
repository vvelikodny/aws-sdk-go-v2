// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of DeleteBudget
type DeleteBudgetInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget that you want to delete.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget that you want to delete.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBudgetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBudgetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBudgetInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of DeleteBudget
type DeleteBudgetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBudgetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBudget = "DeleteBudget"

// DeleteBudgetRequest returns a request value for making API operation for
// AWS Budgets.
//
// Deletes a budget. You can delete your budget at any time.
//
// Deleting a budget also deletes the notifications and subscribers that are
// associated with that budget.
//
//    // Example sending a request using DeleteBudgetRequest.
//    req := client.DeleteBudgetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteBudgetRequest(input *DeleteBudgetInput) DeleteBudgetRequest {
	op := &aws.Operation{
		Name:       opDeleteBudget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBudgetInput{}
	}

	req := c.newRequest(op, input, &DeleteBudgetOutput{})
	return DeleteBudgetRequest{Request: req, Input: input, Copy: c.DeleteBudgetRequest}
}

// DeleteBudgetRequest is the request type for the
// DeleteBudget API operation.
type DeleteBudgetRequest struct {
	*aws.Request
	Input *DeleteBudgetInput
	Copy  func(*DeleteBudgetInput) DeleteBudgetRequest
}

// Send marshals and sends the DeleteBudget API request.
func (r DeleteBudgetRequest) Send(ctx context.Context) (*DeleteBudgetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBudgetResponse{
		DeleteBudgetOutput: r.Request.Data.(*DeleteBudgetOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBudgetResponse is the response type for the
// DeleteBudget API operation.
type DeleteBudgetResponse struct {
	*DeleteBudgetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBudget request.
func (r *DeleteBudgetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
