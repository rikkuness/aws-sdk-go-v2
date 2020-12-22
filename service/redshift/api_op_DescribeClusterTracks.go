// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the available maintenance tracks.
func (c *Client) DescribeClusterTracks(ctx context.Context, params *DescribeClusterTracksInput, optFns ...func(*Options)) (*DescribeClusterTracksOutput, error) {
	if params == nil {
		params = &DescribeClusterTracksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterTracks", params, optFns, addOperationDescribeClusterTracksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterTracksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterTracksInput struct {

	// The name of the maintenance track.
	MaintenanceTrackName *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterTracks request exceed the
	// value specified in MaxRecords, Amazon Redshift returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// An integer value for the maximum number of maintenance tracks to return.
	MaxRecords *int32
}

type DescribeClusterTracksOutput struct {

	// A list of maintenance tracks output by the DescribeClusterTracks operation.
	MaintenanceTracks []types.MaintenanceTrack

	// The starting point to return a set of response tracklist records. You can
	// retrieve the next set of response records by providing the returned marker value
	// in the Marker parameter and retrying the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClusterTracksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterTracks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterTracks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterTracks(options.Region), middleware.Before); err != nil {
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

// DescribeClusterTracksAPIClient is a client that implements the
// DescribeClusterTracks operation.
type DescribeClusterTracksAPIClient interface {
	DescribeClusterTracks(context.Context, *DescribeClusterTracksInput, ...func(*Options)) (*DescribeClusterTracksOutput, error)
}

var _ DescribeClusterTracksAPIClient = (*Client)(nil)

// DescribeClusterTracksPaginatorOptions is the paginator options for
// DescribeClusterTracks
type DescribeClusterTracksPaginatorOptions struct {
	// An integer value for the maximum number of maintenance tracks to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterTracksPaginator is a paginator for DescribeClusterTracks
type DescribeClusterTracksPaginator struct {
	options   DescribeClusterTracksPaginatorOptions
	client    DescribeClusterTracksAPIClient
	params    *DescribeClusterTracksInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterTracksPaginator returns a new DescribeClusterTracksPaginator
func NewDescribeClusterTracksPaginator(client DescribeClusterTracksAPIClient, params *DescribeClusterTracksInput, optFns ...func(*DescribeClusterTracksPaginatorOptions)) *DescribeClusterTracksPaginator {
	options := DescribeClusterTracksPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeClusterTracksInput{}
	}

	return &DescribeClusterTracksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterTracksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeClusterTracks page.
func (p *DescribeClusterTracksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterTracksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeClusterTracks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeClusterTracks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterTracks",
	}
}
