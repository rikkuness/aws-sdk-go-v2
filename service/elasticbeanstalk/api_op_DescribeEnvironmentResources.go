// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns AWS resources for this environment.
func (c *Client) DescribeEnvironmentResources(ctx context.Context, params *DescribeEnvironmentResourcesInput, optFns ...func(*Options)) (*DescribeEnvironmentResourcesOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentResources", params, optFns, addOperationDescribeEnvironmentResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe the resources in an environment.
type DescribeEnvironmentResourcesInput struct {

	// The ID of the environment to retrieve AWS resource usage data. Condition: You
	// must specify either this or an EnvironmentName, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentId *string

	// The name of the environment to retrieve AWS resource usage data. Condition: You
	// must specify either this or an EnvironmentId, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string
}

// Result message containing a list of environment resource descriptions.
type DescribeEnvironmentResourcesOutput struct {

	// A list of EnvironmentResourceDescription.
	EnvironmentResources *types.EnvironmentResourceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentResources{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentResources",
	}
}
