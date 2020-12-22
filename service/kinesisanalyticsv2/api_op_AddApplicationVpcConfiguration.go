// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a Virtual Private Cloud (VPC) configuration to the application.
// Applications can use VPCs to store and access resources securely. Note the
// following about VPC configurations for Kinesis Data Analytics applications:
//
// *
// VPC configurations are not supported for SQL applications.
//
// * When a VPC is
// added to a Kinesis Data Analytics application, the application can no longer be
// accessed from the Internet directly. To enable Internet access to the
// application, add an Internet gateway to your VPC.
func (c *Client) AddApplicationVpcConfiguration(ctx context.Context, params *AddApplicationVpcConfigurationInput, optFns ...func(*Options)) (*AddApplicationVpcConfigurationOutput, error) {
	if params == nil {
		params = &AddApplicationVpcConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationVpcConfiguration", params, optFns, addOperationAddApplicationVpcConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationVpcConfigurationInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The version of the application to which you want to add the VPC configuration.
	// You can use the DescribeApplication operation to get the current application
	// version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// Description of the VPC to add to the application.
	//
	// This member is required.
	VpcConfiguration *types.VpcConfiguration
}

type AddApplicationVpcConfigurationOutput struct {

	// The ARN of the application.
	ApplicationARN *string

	// Provides the current application version. Kinesis Data Analytics updates the
	// ApplicationVersionId each time you update the application.
	ApplicationVersionId *int64

	// The parameters of the new VPC configuration.
	VpcConfigurationDescription *types.VpcConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationVpcConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationVpcConfiguration{}, middleware.After)
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
	if err = addOpAddApplicationVpcConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationVpcConfiguration",
	}
}
