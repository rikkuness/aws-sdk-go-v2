// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the MSK configurations in this Region.
func (c *Client) ListConfigurationRevisions(ctx context.Context, params *ListConfigurationRevisionsInput, optFns ...func(*Options)) (*ListConfigurationRevisionsOutput, error) {
	if params == nil {
		params = &ListConfigurationRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationRevisions", params, optFns, addOperationListConfigurationRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationRevisionsInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration and
	// all of its revisions.
	//
	// This member is required.
	Arn *string

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults int32

	// The paginated results marker. When the result of the operation is truncated, the
	// call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string
}

type ListConfigurationRevisionsOutput struct {

	// Paginated results marker.
	NextToken *string

	// List of ConfigurationRevision objects.
	Revisions []types.ConfigurationRevision

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConfigurationRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationRevisions{}, middleware.After)
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
	if err = addOpListConfigurationRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationRevisions(options.Region), middleware.Before); err != nil {
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

// ListConfigurationRevisionsAPIClient is a client that implements the
// ListConfigurationRevisions operation.
type ListConfigurationRevisionsAPIClient interface {
	ListConfigurationRevisions(context.Context, *ListConfigurationRevisionsInput, ...func(*Options)) (*ListConfigurationRevisionsOutput, error)
}

var _ ListConfigurationRevisionsAPIClient = (*Client)(nil)

// ListConfigurationRevisionsPaginatorOptions is the paginator options for
// ListConfigurationRevisions
type ListConfigurationRevisionsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfigurationRevisionsPaginator is a paginator for
// ListConfigurationRevisions
type ListConfigurationRevisionsPaginator struct {
	options   ListConfigurationRevisionsPaginatorOptions
	client    ListConfigurationRevisionsAPIClient
	params    *ListConfigurationRevisionsInput
	nextToken *string
	firstPage bool
}

// NewListConfigurationRevisionsPaginator returns a new
// ListConfigurationRevisionsPaginator
func NewListConfigurationRevisionsPaginator(client ListConfigurationRevisionsAPIClient, params *ListConfigurationRevisionsInput, optFns ...func(*ListConfigurationRevisionsPaginatorOptions)) *ListConfigurationRevisionsPaginator {
	options := ListConfigurationRevisionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListConfigurationRevisionsInput{}
	}

	return &ListConfigurationRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfigurationRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListConfigurationRevisions page.
func (p *ListConfigurationRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfigurationRevisionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListConfigurationRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConfigurationRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListConfigurationRevisions",
	}
}
