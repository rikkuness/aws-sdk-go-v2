// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the updates associated with an Amazon EKS cluster or managed node group in
// your AWS account, in the specified Region.
func (c *Client) ListUpdates(ctx context.Context, params *ListUpdatesInput, optFns ...func(*Options)) (*ListUpdatesOutput, error) {
	if params == nil {
		params = &ListUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUpdates", params, optFns, addOperationListUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUpdatesInput struct {

	// The name of the Amazon EKS cluster to list updates for.
	//
	// This member is required.
	Name *string

	// The maximum number of update results returned by ListUpdates in paginated
	// output. When you use this parameter, ListUpdates returns only maxResults results
	// in a single page along with a nextToken response element. You can see the
	// remaining results of the initial request by sending another ListUpdates request
	// with the returned nextToken value. This value can be between 1 and 100. If you
	// don't use this parameter, ListUpdates returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListUpdates request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string

	// The name of the Amazon EKS managed node group to list updates for.
	NodegroupName *string
}

type ListUpdatesOutput struct {

	// The nextToken value to include in a future ListUpdates request. When the results
	// of a ListUpdates request exceed maxResults, you can use this value to retrieve
	// the next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// A list of all the updates for the specified cluster and Region.
	UpdateIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUpdates{}, middleware.After)
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
	if err = addOpListUpdatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUpdates(options.Region), middleware.Before); err != nil {
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

// ListUpdatesAPIClient is a client that implements the ListUpdates operation.
type ListUpdatesAPIClient interface {
	ListUpdates(context.Context, *ListUpdatesInput, ...func(*Options)) (*ListUpdatesOutput, error)
}

var _ ListUpdatesAPIClient = (*Client)(nil)

// ListUpdatesPaginatorOptions is the paginator options for ListUpdates
type ListUpdatesPaginatorOptions struct {
	// The maximum number of update results returned by ListUpdates in paginated
	// output. When you use this parameter, ListUpdates returns only maxResults results
	// in a single page along with a nextToken response element. You can see the
	// remaining results of the initial request by sending another ListUpdates request
	// with the returned nextToken value. This value can be between 1 and 100. If you
	// don't use this parameter, ListUpdates returns up to 100 results and a nextToken
	// value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUpdatesPaginator is a paginator for ListUpdates
type ListUpdatesPaginator struct {
	options   ListUpdatesPaginatorOptions
	client    ListUpdatesAPIClient
	params    *ListUpdatesInput
	nextToken *string
	firstPage bool
}

// NewListUpdatesPaginator returns a new ListUpdatesPaginator
func NewListUpdatesPaginator(client ListUpdatesAPIClient, params *ListUpdatesInput, optFns ...func(*ListUpdatesPaginatorOptions)) *ListUpdatesPaginator {
	options := ListUpdatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListUpdatesInput{}
	}

	return &ListUpdatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUpdatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUpdates page.
func (p *ListUpdatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUpdatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListUpdates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListUpdates",
	}
}
