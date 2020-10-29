// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or modifies OwnershipControls for an Amazon S3 bucket. To use this
// operation, you must have the s3:GetBucketOwnershipControls permission. For more
// information about Amazon S3 permissions, see Specifying Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// For information about Amazon S3 Object Ownership, see Using Object Ownership
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html).
// The following operations are related to GetBucketOwnershipControls:
//
// *
// GetBucketOwnershipControls
//
// * DeleteBucketOwnershipControls
func (c *Client) PutBucketOwnershipControls(ctx context.Context, params *PutBucketOwnershipControlsInput, optFns ...func(*Options)) (*PutBucketOwnershipControlsOutput, error) {
	if params == nil {
		params = &PutBucketOwnershipControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketOwnershipControls", params, optFns, addOperationPutBucketOwnershipControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketOwnershipControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketOwnershipControlsInput struct {

	// The name of the Amazon S3 bucket whose OwnershipControls you want to set.
	//
	// This member is required.
	Bucket *string

	// The OwnershipControls (BucketOwnerPreferred or ObjectWriter) that you want to
	// apply to this Amazon S3 bucket.
	//
	// This member is required.
	OwnershipControls *types.OwnershipControls

	// The MD5 hash of the OwnershipControls request body.
	ContentMD5 *string

	ExpectedBucketOwner *string
}

type PutBucketOwnershipControlsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketOwnershipControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketOwnershipControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketOwnershipControls{}, middleware.After)
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
	addOpPutBucketOwnershipControlsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketOwnershipControls(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBucketOwnershipControls(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketOwnershipControls",
	}
}