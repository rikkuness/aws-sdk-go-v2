// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Search the schemas
func (c *Client) SearchSchemas(ctx context.Context, params *SearchSchemasInput, optFns ...func(*Options)) (*SearchSchemasOutput, error) {
	if params == nil {
		params = &SearchSchemasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSchemas", params, optFns, addOperationSearchSchemasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSchemasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSchemasInput struct {

	// Specifying this limits the results to only schemas that include the provided
	// keywords.
	//
	// This member is required.
	Keywords *string

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	Limit int32

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string
}

type SearchSchemasOutput struct {

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// An array of SearchSchemaSummary information.
	Schemas []types.SearchSchemaSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchSchemasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchSchemas{}, middleware.After)
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
	if err = addOpSearchSchemasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSchemas(options.Region), middleware.Before); err != nil {
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

// SearchSchemasAPIClient is a client that implements the SearchSchemas operation.
type SearchSchemasAPIClient interface {
	SearchSchemas(context.Context, *SearchSchemasInput, ...func(*Options)) (*SearchSchemasOutput, error)
}

var _ SearchSchemasAPIClient = (*Client)(nil)

// SearchSchemasPaginatorOptions is the paginator options for SearchSchemas
type SearchSchemasPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchSchemasPaginator is a paginator for SearchSchemas
type SearchSchemasPaginator struct {
	options   SearchSchemasPaginatorOptions
	client    SearchSchemasAPIClient
	params    *SearchSchemasInput
	nextToken *string
	firstPage bool
}

// NewSearchSchemasPaginator returns a new SearchSchemasPaginator
func NewSearchSchemasPaginator(client SearchSchemasAPIClient, params *SearchSchemasInput, optFns ...func(*SearchSchemasPaginatorOptions)) *SearchSchemasPaginator {
	options := SearchSchemasPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &SearchSchemasInput{}
	}

	return &SearchSchemasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchSchemasPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SearchSchemas page.
func (p *SearchSchemasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchSchemasOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.SearchSchemas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchSchemas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "SearchSchemas",
	}
}
