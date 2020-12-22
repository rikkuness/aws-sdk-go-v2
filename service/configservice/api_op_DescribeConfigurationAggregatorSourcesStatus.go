// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns status information for sources within an aggregator. The status includes
// information about the last time AWS Config verified authorization between the
// source account and an aggregator account. In case of a failure, the status
// contains the related error code or message.
func (c *Client) DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigurationAggregatorSourcesStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationAggregatorSourcesStatus", params, optFns, addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationAggregatorSourcesStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationAggregatorSourcesStatusInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, AWS Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Filters the status type.
	//
	// * Valid value FAILED indicates errors while moving
	// data.
	//
	// * Valid value SUCCEEDED indicates the data was successfully moved.
	//
	// *
	// Valid value OUTDATED indicates the data is not the most recent.
	UpdateStatus []types.AggregatedSourceStatusType
}

type DescribeConfigurationAggregatorSourcesStatusOutput struct {

	// Returns an AggregatedSourceStatus object.
	AggregatedSourceStatusList []types.AggregatedSourceStatus

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
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
	if err = addOpDescribeConfigurationAggregatorSourcesStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationAggregatorSourcesStatus",
	}
}
