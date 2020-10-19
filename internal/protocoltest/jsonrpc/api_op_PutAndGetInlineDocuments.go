// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as part of the payload.
func (c *Client) PutAndGetInlineDocuments(ctx context.Context, params *PutAndGetInlineDocumentsInput, optFns ...func(*Options)) (*PutAndGetInlineDocumentsOutput, error) {
	if params == nil {
		params = &PutAndGetInlineDocumentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAndGetInlineDocuments", params, optFns, addOperationPutAndGetInlineDocumentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAndGetInlineDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAndGetInlineDocumentsInput struct {
	InlineDocument smithy.Document
}

type PutAndGetInlineDocumentsOutput struct {
	InlineDocument smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAndGetInlineDocumentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAndGetInlineDocuments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAndGetInlineDocuments{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAndGetInlineDocuments(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutAndGetInlineDocuments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "foo",
		OperationName: "PutAndGetInlineDocuments",
	}
}