// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
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
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketOwnershipControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketOwnershipControls(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketOwnershipControlsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketOwnershipControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketOwnershipControls",
	}
}

// getPutBucketOwnershipControlsBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getPutBucketOwnershipControlsBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketOwnershipControlsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketOwnershipControlsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketOwnershipControlsBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
