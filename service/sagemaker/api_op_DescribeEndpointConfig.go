// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the description of an endpoint configuration created using the
// CreateEndpointConfig API.
func (c *Client) DescribeEndpointConfig(ctx context.Context, params *DescribeEndpointConfigInput, optFns ...func(*Options)) (*DescribeEndpointConfigOutput, error) {
	if params == nil {
		params = &DescribeEndpointConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointConfig", params, optFns, addOperationDescribeEndpointConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointConfigInput struct {

	// The name of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string
}

type DescribeEndpointConfigOutput struct {

	// A timestamp that shows when the endpoint configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigArn *string

	// Name of the Amazon SageMaker endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint.
	//
	// This member is required.
	ProductionVariants []types.ProductionVariant

	//
	DataCaptureConfig *types.DataCaptureConfig

	// AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML
	// storage volume attached to the instance.
	KmsKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEndpointConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpointConfig{}, middleware.After)
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
	if err = addOpDescribeEndpointConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEndpointConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeEndpointConfig",
	}
}
