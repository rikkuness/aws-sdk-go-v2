// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified policies. If you specify a load balancer name, the
// action returns the descriptions of all policies created for the load balancer.
// If you specify a policy name associated with your load balancer, the action
// returns the description of that policy. If you don't specify a load balancer
// name, the action returns descriptions of the specified sample policies, or
// descriptions of all sample policies. The names of the sample policies have the
// ELBSample- prefix.
func (c *Client) DescribeLoadBalancerPolicies(ctx context.Context, params *DescribeLoadBalancerPoliciesInput, optFns ...func(*Options)) (*DescribeLoadBalancerPoliciesOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancerPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancerPolicies", params, optFns, addOperationDescribeLoadBalancerPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancerPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancerPolicies.
type DescribeLoadBalancerPoliciesInput struct {

	// The name of the load balancer.
	LoadBalancerName *string

	// The names of the policies.
	PolicyNames []string
}

// Contains the output of DescribeLoadBalancerPolicies.
type DescribeLoadBalancerPoliciesOutput struct {

	// Information about the policies.
	PolicyDescriptions []types.PolicyDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLoadBalancerPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLoadBalancerPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancerPolicies",
	}
}
