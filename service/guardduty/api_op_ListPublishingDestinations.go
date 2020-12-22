// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of publishing destinations associated with the specified
// dectectorId.
func (c *Client) ListPublishingDestinations(ctx context.Context, params *ListPublishingDestinationsInput, optFns ...func(*Options)) (*ListPublishingDestinationsOutput, error) {
	if params == nil {
		params = &ListPublishingDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublishingDestinations", params, optFns, addOperationListPublishingDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublishingDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublishingDestinationsInput struct {

	// The ID of the detector to retrieve publishing destinations for.
	//
	// This member is required.
	DetectorId *string

	// The maximum number of results to return in the response.
	MaxResults int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string
}

type ListPublishingDestinationsOutput struct {

	// A Destinations object that includes information about each publishing
	// destination returned.
	//
	// This member is required.
	Destinations []types.Destination

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPublishingDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPublishingDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublishingDestinations{}, middleware.After)
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
	if err = addOpListPublishingDestinationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublishingDestinations(options.Region), middleware.Before); err != nil {
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

// ListPublishingDestinationsAPIClient is a client that implements the
// ListPublishingDestinations operation.
type ListPublishingDestinationsAPIClient interface {
	ListPublishingDestinations(context.Context, *ListPublishingDestinationsInput, ...func(*Options)) (*ListPublishingDestinationsOutput, error)
}

var _ ListPublishingDestinationsAPIClient = (*Client)(nil)

// ListPublishingDestinationsPaginatorOptions is the paginator options for
// ListPublishingDestinations
type ListPublishingDestinationsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPublishingDestinationsPaginator is a paginator for
// ListPublishingDestinations
type ListPublishingDestinationsPaginator struct {
	options   ListPublishingDestinationsPaginatorOptions
	client    ListPublishingDestinationsAPIClient
	params    *ListPublishingDestinationsInput
	nextToken *string
	firstPage bool
}

// NewListPublishingDestinationsPaginator returns a new
// ListPublishingDestinationsPaginator
func NewListPublishingDestinationsPaginator(client ListPublishingDestinationsAPIClient, params *ListPublishingDestinationsInput, optFns ...func(*ListPublishingDestinationsPaginatorOptions)) *ListPublishingDestinationsPaginator {
	options := ListPublishingDestinationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListPublishingDestinationsInput{}
	}

	return &ListPublishingDestinationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPublishingDestinationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPublishingDestinations page.
func (p *ListPublishingDestinationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPublishingDestinationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPublishingDestinations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPublishingDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListPublishingDestinations",
	}
}
