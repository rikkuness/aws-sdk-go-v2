// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets usage information for an identity, including number of datasets and data
// usage. This API can be called with temporary user credentials provided by
// Cognito Identity or with developer credentials.
func (c *Client) DescribeIdentityUsage(ctx context.Context, params *DescribeIdentityUsageInput, optFns ...func(*Options)) (*DescribeIdentityUsageOutput, error) {
	if params == nil {
		params = &DescribeIdentityUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdentityUsage", params, optFns, addOperationDescribeIdentityUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdentityUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for information about the usage of an identity pool.
type DescribeIdentityUsageInput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string
}

// The response to a successful DescribeIdentityUsage request.
type DescribeIdentityUsageOutput struct {

	// Usage information for the identity.
	IdentityUsage *types.IdentityUsage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIdentityUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIdentityUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIdentityUsage{}, middleware.After)
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
	if err = addOpDescribeIdentityUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityUsage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeIdentityUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "DescribeIdentityUsage",
	}
}
