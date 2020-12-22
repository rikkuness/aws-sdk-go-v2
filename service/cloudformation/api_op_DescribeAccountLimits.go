// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves your account's AWS CloudFormation limits, such as the maximum number
// of stacks that you can create in your account. For more information about
// account limits, see AWS CloudFormation Limits
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html)
// in the AWS CloudFormation User Guide.
func (c *Client) DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) {
	if params == nil {
		params = &DescribeAccountLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountLimits", params, optFns, addOperationDescribeAccountLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeAccountLimits action.
type DescribeAccountLimitsInput struct {

	// A string that identifies the next page of limits that you want to retrieve.
	NextToken *string
}

// The output for the DescribeAccountLimits action.
type DescribeAccountLimitsOutput struct {

	// An account limit structure that contain a list of AWS CloudFormation account
	// limits and their values.
	AccountLimits []types.AccountLimit

	// If the output exceeds 1 MB in size, a string that identifies the next page of
	// limits. If no additional page exists, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAccountLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAccountLimits{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountLimits(options.Region), middleware.Before); err != nil {
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

// DescribeAccountLimitsAPIClient is a client that implements the
// DescribeAccountLimits operation.
type DescribeAccountLimitsAPIClient interface {
	DescribeAccountLimits(context.Context, *DescribeAccountLimitsInput, ...func(*Options)) (*DescribeAccountLimitsOutput, error)
}

var _ DescribeAccountLimitsAPIClient = (*Client)(nil)

// DescribeAccountLimitsPaginatorOptions is the paginator options for
// DescribeAccountLimits
type DescribeAccountLimitsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAccountLimitsPaginator is a paginator for DescribeAccountLimits
type DescribeAccountLimitsPaginator struct {
	options   DescribeAccountLimitsPaginatorOptions
	client    DescribeAccountLimitsAPIClient
	params    *DescribeAccountLimitsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAccountLimitsPaginator returns a new DescribeAccountLimitsPaginator
func NewDescribeAccountLimitsPaginator(client DescribeAccountLimitsAPIClient, params *DescribeAccountLimitsInput, optFns ...func(*DescribeAccountLimitsPaginatorOptions)) *DescribeAccountLimitsPaginator {
	options := DescribeAccountLimitsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeAccountLimitsInput{}
	}

	return &DescribeAccountLimitsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAccountLimitsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAccountLimits page.
func (p *DescribeAccountLimitsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeAccountLimits(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeAccountLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeAccountLimits",
	}
}
