// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the specified TagOptions or all TagOptions.
func (c *Client) ListTagOptions(ctx context.Context, params *ListTagOptionsInput, optFns ...func(*Options)) (*ListTagOptionsOutput, error) {
	if params == nil {
		params = &ListTagOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagOptions", params, optFns, addOperationListTagOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagOptionsInput struct {

	// The search filters. If no search filters are specified, the output includes all
	// TagOptions.
	Filters *types.ListTagOptionsFilters

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListTagOptionsOutput struct {

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// Information about the TagOptions.
	TagOptionDetails []types.TagOptionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagOptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagOptions(options.Region), middleware.Before); err != nil {
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

// ListTagOptionsAPIClient is a client that implements the ListTagOptions
// operation.
type ListTagOptionsAPIClient interface {
	ListTagOptions(context.Context, *ListTagOptionsInput, ...func(*Options)) (*ListTagOptionsOutput, error)
}

var _ ListTagOptionsAPIClient = (*Client)(nil)

// ListTagOptionsPaginatorOptions is the paginator options for ListTagOptions
type ListTagOptionsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTagOptionsPaginator is a paginator for ListTagOptions
type ListTagOptionsPaginator struct {
	options   ListTagOptionsPaginatorOptions
	client    ListTagOptionsAPIClient
	params    *ListTagOptionsInput
	nextToken *string
	firstPage bool
}

// NewListTagOptionsPaginator returns a new ListTagOptionsPaginator
func NewListTagOptionsPaginator(client ListTagOptionsAPIClient, params *ListTagOptionsInput, optFns ...func(*ListTagOptionsPaginatorOptions)) *ListTagOptionsPaginator {
	options := ListTagOptionsPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTagOptionsInput{}
	}

	return &ListTagOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTagOptionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTagOptions page.
func (p *ListTagOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTagOptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListTagOptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTagOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListTagOptions",
	}
}
