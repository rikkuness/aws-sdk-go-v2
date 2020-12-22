// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique endpoint specific to the AWS account making the call.
func (c *Client) DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) {
	if params == nil {
		params = &DescribeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoint", params, optFns, addOperationDescribeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeEndpoint operation.
type DescribeEndpointInput struct {

	// The endpoint type. Valid endpoint types include:
	//
	// * iot:Data - Returns a
	// VeriSign signed data endpoint.
	//
	// * iot:Data-ATS - Returns an ATS signed data
	// endpoint.
	//
	// * iot:CredentialProvider - Returns an AWS IoT credentials provider
	// API endpoint.
	//
	// * iot:Jobs - Returns an AWS IoT device management Jobs API
	// endpoint.
	//
	// We strongly recommend that customers use the newer iot:Data-ATS
	// endpoint type to avoid issues related to the widespread distrust of Symantec
	// certificate authorities.
	EndpointType *string
}

// The output from the DescribeEndpoint operation.
type DescribeEndpointOutput struct {

	// The endpoint. The format of the endpoint is as follows:
	// identifier.iot.region.amazonaws.com.
	EndpointAddress *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEndpoint{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeEndpoint",
	}
}
