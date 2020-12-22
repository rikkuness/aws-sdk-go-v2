// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastoredata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of metadata entries about folders and objects in the specified
// folder.
func (c *Client) ListItems(ctx context.Context, params *ListItemsInput, optFns ...func(*Options)) (*ListItemsOutput, error) {
	if params == nil {
		params = &ListItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListItems", params, optFns, addOperationListItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListItemsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListItems request with MaxResults set at 500. Although 2,000 items match your
	// request, the service returns no more than the first 500 items. (The service also
	// returns a NextToken value that you can use to fetch the next batch of results.)
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 1,000 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListItems request with MaxResults set at 500. The service
	// returns the first batch of results (up to 500) and a NextToken value. To see the
	// next batch of results, you can submit the ListItems request a second time and
	// specify the NextToken value. Tokens expire after 15 minutes.
	NextToken *string

	// The path in the container from which to retrieve items. Format: //
	Path *string
}

type ListItemsOutput struct {

	// The metadata entries for the folders and objects at the requested path.
	Items []types.Item

	// The token that can be used in a request to view the next set of results. For
	// example, you submit a ListItems request that matches 2,000 items with MaxResults
	// set at 500. The service returns the first batch of results (up to 500) and a
	// NextToken value that can be used to fetch the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListItems{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListItems(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "ListItems",
	}
}
