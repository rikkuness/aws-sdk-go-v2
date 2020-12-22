// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the path tracking results of a Amazon Rekognition Video analysis started by
// StartPersonTracking. The person path tracking operation is started by a call to
// StartPersonTracking which returns a job identifier (JobId). When the operation
// finishes, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to
// StartPersonTracking. To get the results of the person path tracking operation,
// first check that the status value published to the Amazon SNS topic is
// SUCCEEDED. If so, call GetPersonTracking and pass the job identifier (JobId)
// from the initial call to StartPersonTracking. GetPersonTracking returns an
// array, Persons, of tracked persons and the time(s) their paths were tracked in
// the video. GetPersonTracking only returns the default
//
// facial attributes
// (BoundingBox, Confidence, Landmarks, Pose, and Quality). The other facial
// attributes listed in the Face object of the following response syntax are not
// returned. For more information, see FaceDetail in the Amazon Rekognition
// Developer Guide. By default, the array is sorted by the time(s) a person's path
// is tracked in the video. You can sort by tracked persons by specifying INDEX for
// the SortBy input parameter. Use the MaxResults parameter to limit the number of
// items returned. If there are more results than specified in MaxResults, the
// value of NextToken in the operation response contains a pagination token for
// getting the next set of results. To get the next page of results, call
// GetPersonTracking and populate the NextToken request parameter with the token
// value returned from the previous call to GetPersonTracking.
func (c *Client) GetPersonTracking(ctx context.Context, params *GetPersonTrackingInput, optFns ...func(*Options)) (*GetPersonTrackingOutput, error) {
	if params == nil {
		params = &GetPersonTrackingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPersonTracking", params, optFns, addOperationGetPersonTrackingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPersonTrackingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPersonTrackingInput struct {

	// The identifier for a job that tracks persons in a video. You get the JobId from
	// a call to StartPersonTracking.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more persons to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of persons.
	NextToken *string

	// Sort to use for elements in the Persons array. Use TIMESTAMP to sort array
	// elements by the time persons are detected. Use INDEX to sort by the tracked
	// persons. If you sort by INDEX, the array elements for each person are sorted by
	// detection confidence. The default sort is by TIMESTAMP.
	SortBy types.PersonTrackingSortBy
}

type GetPersonTrackingOutput struct {

	// The current status of the person tracking job.
	JobStatus types.VideoJobStatus

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of persons.
	NextToken *string

	// An array of the persons detected in the video and the time(s) their path was
	// tracked throughout the video. An array element will exist for each time a
	// person's path is tracked.
	Persons []types.PersonDetection

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition Video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPersonTrackingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPersonTracking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPersonTracking{}, middleware.After)
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
	if err = addOpGetPersonTrackingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPersonTracking(options.Region), middleware.Before); err != nil {
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

// GetPersonTrackingAPIClient is a client that implements the GetPersonTracking
// operation.
type GetPersonTrackingAPIClient interface {
	GetPersonTracking(context.Context, *GetPersonTrackingInput, ...func(*Options)) (*GetPersonTrackingOutput, error)
}

var _ GetPersonTrackingAPIClient = (*Client)(nil)

// GetPersonTrackingPaginatorOptions is the paginator options for GetPersonTracking
type GetPersonTrackingPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPersonTrackingPaginator is a paginator for GetPersonTracking
type GetPersonTrackingPaginator struct {
	options   GetPersonTrackingPaginatorOptions
	client    GetPersonTrackingAPIClient
	params    *GetPersonTrackingInput
	nextToken *string
	firstPage bool
}

// NewGetPersonTrackingPaginator returns a new GetPersonTrackingPaginator
func NewGetPersonTrackingPaginator(client GetPersonTrackingAPIClient, params *GetPersonTrackingInput, optFns ...func(*GetPersonTrackingPaginatorOptions)) *GetPersonTrackingPaginator {
	options := GetPersonTrackingPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetPersonTrackingInput{}
	}

	return &GetPersonTrackingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPersonTrackingPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetPersonTracking page.
func (p *GetPersonTrackingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPersonTrackingOutput, error) {
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

	result, err := p.client.GetPersonTracking(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetPersonTracking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetPersonTracking",
	}
}
