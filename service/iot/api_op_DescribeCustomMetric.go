// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Defender detect custom metric. Requires
// permission to access the DescribeCustomMetric
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeCustomMetric(ctx context.Context, params *DescribeCustomMetricInput, optFns ...func(*Options)) (*DescribeCustomMetricOutput, error) {
	if params == nil {
		params = &DescribeCustomMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomMetric", params, optFns, c.addOperationDescribeCustomMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomMetricInput struct {

	// The name of the custom metric.
	//
	// This member is required.
	MetricName *string

	noSmithyDocumentSerde
}

type DescribeCustomMetricOutput struct {

	// The creation date of the custom metric in milliseconds since epoch.
	CreationDate *time.Time

	// Field represents a friendly name in the console for the custom metric; doesn't
	// have to be unique. Don't use this name as the metric identifier in the device
	// metric report. Can be updated.
	DisplayName *string

	// The time the custom metric was last modified in milliseconds since epoch.
	LastModifiedDate *time.Time

	// The Amazon Resource Number (ARN) of the custom metric.
	MetricArn *string

	// The name of the custom metric.
	MetricName *string

	// The type of the custom metric. Types include string-list, ip-address-list,
	// number-list, and number.
	MetricType types.CustomMetricType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCustomMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCustomMetric{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpDescribeCustomMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomMetric(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCustomMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeCustomMetric",
	}
}