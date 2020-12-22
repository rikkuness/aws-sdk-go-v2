// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the AWS Glue DataBrew recipes in the current AWS account.
func (c *Client) ListRecipes(ctx context.Context, params *ListRecipesInput, optFns ...func(*Options)) (*ListRecipesOutput, error) {
	if params == nil {
		params = &ListRecipesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecipes", params, optFns, addOperationListRecipesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecipesInput struct {

	// The maximum number of results to return in this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// A version identifier. Using this parameter indicates to return only those
	// recipes that have this version identifier.
	RecipeVersion *string
}

type ListRecipesOutput struct {

	// A list of recipes that are defined in the current AWS account.
	//
	// This member is required.
	Recipes []types.Recipe

	// A token generated by DataBrew that specifies where to continue pagination if a
	// previous request was truncated. To get the next set of pages, pass in the
	// NextToken value from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRecipesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecipes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecipes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecipes(options.Region), middleware.Before); err != nil {
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

// ListRecipesAPIClient is a client that implements the ListRecipes operation.
type ListRecipesAPIClient interface {
	ListRecipes(context.Context, *ListRecipesInput, ...func(*Options)) (*ListRecipesOutput, error)
}

var _ ListRecipesAPIClient = (*Client)(nil)

// ListRecipesPaginatorOptions is the paginator options for ListRecipes
type ListRecipesPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecipesPaginator is a paginator for ListRecipes
type ListRecipesPaginator struct {
	options   ListRecipesPaginatorOptions
	client    ListRecipesAPIClient
	params    *ListRecipesInput
	nextToken *string
	firstPage bool
}

// NewListRecipesPaginator returns a new ListRecipesPaginator
func NewListRecipesPaginator(client ListRecipesAPIClient, params *ListRecipesInput, optFns ...func(*ListRecipesPaginatorOptions)) *ListRecipesPaginator {
	options := ListRecipesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListRecipesInput{}
	}

	return &ListRecipesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecipesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRecipes page.
func (p *ListRecipesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecipesOutput, error) {
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

	result, err := p.client.ListRecipes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecipes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "ListRecipes",
	}
}
