// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of robots. You can optionally provide filters to retrieve
// specific robots.
func (c *Client) ListRobots(ctx context.Context, params *ListRobotsInput, optFns ...func(*Options)) (*ListRobotsOutput, error) {
	if params == nil {
		params = &ListRobotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRobots", params, optFns, addOperationListRobotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRobotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRobotsInput struct {

	// Optional filters to limit results. The filter names status and fleetName are
	// supported. When filtering, you must use the complete value of the filtered item.
	// You can use up to three filters, but they must be for the same named item. For
	// example, if you are looking for items with the status Registered or the status
	// Available.
	Filters []types.Filter

	// When this parameter is used, ListRobots only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListRobots request with the
	// returned nextToken value. This value can be between 1 and 200. If this parameter
	// is not used, then ListRobots returns up to 200 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListRobots again and assign that token to the
	// request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListRobotsOutput struct {

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListRobots again and assign that token to the
	// request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// A list of robots that meet the criteria of the request.
	Robots []types.Robot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRobotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRobots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRobots{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRobots(options.Region), middleware.Before); err != nil {
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

// ListRobotsAPIClient is a client that implements the ListRobots operation.
type ListRobotsAPIClient interface {
	ListRobots(context.Context, *ListRobotsInput, ...func(*Options)) (*ListRobotsOutput, error)
}

var _ ListRobotsAPIClient = (*Client)(nil)

// ListRobotsPaginatorOptions is the paginator options for ListRobots
type ListRobotsPaginatorOptions struct {
	// When this parameter is used, ListRobots only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListRobots request with the
	// returned nextToken value. This value can be between 1 and 200. If this parameter
	// is not used, then ListRobots returns up to 200 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRobotsPaginator is a paginator for ListRobots
type ListRobotsPaginator struct {
	options   ListRobotsPaginatorOptions
	client    ListRobotsAPIClient
	params    *ListRobotsInput
	nextToken *string
	firstPage bool
}

// NewListRobotsPaginator returns a new ListRobotsPaginator
func NewListRobotsPaginator(client ListRobotsAPIClient, params *ListRobotsInput, optFns ...func(*ListRobotsPaginatorOptions)) *ListRobotsPaginator {
	options := ListRobotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListRobotsInput{}
	}

	return &ListRobotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRobotsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRobots page.
func (p *ListRobotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRobotsOutput, error) {
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

	result, err := p.client.ListRobots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRobots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListRobots",
	}
}
