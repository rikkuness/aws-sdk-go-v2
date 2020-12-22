// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a table that shows counts of resources that are noncompliant with their
// tag policies. For more information on tag policies, see Tag Policies
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
// in the AWS Organizations User Guide. You can call this operation only from the
// organization's master account and from the us-east-1 Region.
func (c *Client) GetComplianceSummary(ctx context.Context, params *GetComplianceSummaryInput, optFns ...func(*Options)) (*GetComplianceSummaryOutput, error) {
	if params == nil {
		params = &GetComplianceSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceSummary", params, optFns, addOperationGetComplianceSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceSummaryInput struct {

	// A list of attributes to group the counts of noncompliant resources by. If
	// supplied, the counts are sorted by those attributes.
	GroupBy []types.GroupByAttribute

	// A limit that restricts the number of results that are returned per page.
	MaxResults *int32

	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken, use
	// that string for this value to request an additional page of data.
	PaginationToken *string

	// A list of Regions to limit the output by. If you use this parameter, the count
	// of returned noncompliant resources includes only resources in the specified
	// Regions.
	RegionFilters []string

	// The constraints on the resources that you want returned. The format of each
	// resource type is service[:resourceType]. For example, specifying a resource type
	// of ec2 returns all Amazon EC2 resources (which includes EC2 instances).
	// Specifying a resource type of ec2:instance returns only EC2 instances. The
	// string for each service name and resource type is the same as that embedded in a
	// resource's Amazon Resource Name (ARN). Consult the AWS General Reference for the
	// following:
	//
	// * For a list of service name strings, see AWS Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces).
	//
	// *
	// For resource type strings, see Example ARNs
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax).
	//
	// *
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// You
	// can specify multiple resource types by using an array. The array can include up
	// to 100 items. Note that the length constraint requirement applies to each
	// resource type filter.
	ResourceTypeFilters []string

	// A list of tag keys to limit the output by. If you use this parameter, the count
	// of returned noncompliant resources includes only resources that have the
	// specified tag keys.
	TagKeyFilters []string

	// The target identifiers (usually, specific account IDs) to limit the output by.
	// If you use this parameter, the count of returned noncompliant resources includes
	// only resources with the specified target IDs.
	TargetIdFilters []string
}

type GetComplianceSummaryOutput struct {

	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string

	// A table that shows counts of noncompliant resources.
	SummaryList []types.Summary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetComplianceSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceSummary{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceSummary(options.Region), middleware.Before); err != nil {
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

// GetComplianceSummaryAPIClient is a client that implements the
// GetComplianceSummary operation.
type GetComplianceSummaryAPIClient interface {
	GetComplianceSummary(context.Context, *GetComplianceSummaryInput, ...func(*Options)) (*GetComplianceSummaryOutput, error)
}

var _ GetComplianceSummaryAPIClient = (*Client)(nil)

// GetComplianceSummaryPaginatorOptions is the paginator options for
// GetComplianceSummary
type GetComplianceSummaryPaginatorOptions struct {
	// A limit that restricts the number of results that are returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetComplianceSummaryPaginator is a paginator for GetComplianceSummary
type GetComplianceSummaryPaginator struct {
	options   GetComplianceSummaryPaginatorOptions
	client    GetComplianceSummaryAPIClient
	params    *GetComplianceSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetComplianceSummaryPaginator returns a new GetComplianceSummaryPaginator
func NewGetComplianceSummaryPaginator(client GetComplianceSummaryAPIClient, params *GetComplianceSummaryInput, optFns ...func(*GetComplianceSummaryPaginatorOptions)) *GetComplianceSummaryPaginator {
	options := GetComplianceSummaryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetComplianceSummaryInput{}
	}

	return &GetComplianceSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetComplianceSummaryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetComplianceSummary page.
func (p *GetComplianceSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetComplianceSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetComplianceSummary(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetComplianceSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "GetComplianceSummary",
	}
}
