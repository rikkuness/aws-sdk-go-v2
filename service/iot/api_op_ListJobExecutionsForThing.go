// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the job executions for the specified thing.
func (c *Client) ListJobExecutionsForThing(ctx context.Context, params *ListJobExecutionsForThingInput, optFns ...func(*Options)) (*ListJobExecutionsForThingOutput, error) {
	if params == nil {
		params = &ListJobExecutionsForThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobExecutionsForThing", params, optFns, addOperationListJobExecutionsForThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobExecutionsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobExecutionsForThingInput struct {

	// The thing name.
	//
	// This member is required.
	ThingName *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The namespace used to indicate that a job is a customer-managed job. When you
	// specify a value for this parameter, AWS IoT Core sends jobs notifications to
	// MQTT topics that contain the value in the following format.
	// $aws/things/THING_NAME/jobs/JOB_ID/notify-namespace-NAMESPACE_ID/ The
	// namespaceId feature is in public preview.
	NamespaceId *string

	// The token to retrieve the next set of results.
	NextToken *string

	// An optional filter that lets you search for jobs that have the specified status.
	Status types.JobExecutionStatus
}

type ListJobExecutionsForThingOutput struct {

	// A list of job execution summaries.
	ExecutionSummaries []types.JobExecutionSummaryForThing

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobExecutionsForThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobExecutionsForThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobExecutionsForThing{}, middleware.After)
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
	if err = addOpListJobExecutionsForThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobExecutionsForThing(options.Region), middleware.Before); err != nil {
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

// ListJobExecutionsForThingAPIClient is a client that implements the
// ListJobExecutionsForThing operation.
type ListJobExecutionsForThingAPIClient interface {
	ListJobExecutionsForThing(context.Context, *ListJobExecutionsForThingInput, ...func(*Options)) (*ListJobExecutionsForThingOutput, error)
}

var _ ListJobExecutionsForThingAPIClient = (*Client)(nil)

// ListJobExecutionsForThingPaginatorOptions is the paginator options for
// ListJobExecutionsForThing
type ListJobExecutionsForThingPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobExecutionsForThingPaginator is a paginator for ListJobExecutionsForThing
type ListJobExecutionsForThingPaginator struct {
	options   ListJobExecutionsForThingPaginatorOptions
	client    ListJobExecutionsForThingAPIClient
	params    *ListJobExecutionsForThingInput
	nextToken *string
	firstPage bool
}

// NewListJobExecutionsForThingPaginator returns a new
// ListJobExecutionsForThingPaginator
func NewListJobExecutionsForThingPaginator(client ListJobExecutionsForThingAPIClient, params *ListJobExecutionsForThingInput, optFns ...func(*ListJobExecutionsForThingPaginatorOptions)) *ListJobExecutionsForThingPaginator {
	options := ListJobExecutionsForThingPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListJobExecutionsForThingInput{}
	}

	return &ListJobExecutionsForThingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobExecutionsForThingPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJobExecutionsForThing page.
func (p *ListJobExecutionsForThingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobExecutionsForThingOutput, error) {
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

	result, err := p.client.ListJobExecutionsForThing(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobExecutionsForThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListJobExecutionsForThing",
	}
}
