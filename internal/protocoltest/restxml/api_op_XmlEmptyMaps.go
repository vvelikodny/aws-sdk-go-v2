// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) XmlEmptyMaps(ctx context.Context, params *XmlEmptyMapsInput, optFns ...func(*Options)) (*XmlEmptyMapsOutput, error) {
	if params == nil {
		params = &XmlEmptyMapsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlEmptyMaps", params, optFns, addOperationXmlEmptyMapsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlEmptyMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEmptyMapsInput struct {
	MyMap map[string]*types.GreetingStruct
}

type XmlEmptyMapsOutput struct {
	MyMap map[string]*types.GreetingStruct

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlEmptyMapsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpXmlEmptyMaps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpXmlEmptyMaps{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEmptyMaps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opXmlEmptyMaps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlEmptyMaps",
	}
}