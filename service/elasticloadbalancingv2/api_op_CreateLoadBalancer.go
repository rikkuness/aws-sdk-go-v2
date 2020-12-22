// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Application Load Balancer, Network Load Balancer, or Gateway Load
// Balancer. For more information, see the following:
//
// * Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html)
//
// *
// Network Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html)
//
// *
// Gateway Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-load-balancers.html)
//
// This
// operation is idempotent, which means that it completes at most one time. If you
// attempt to create multiple load balancers with the same settings, each call
// succeeds.
func (c *Client) CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancer", params, optFns, addOperationCreateLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoadBalancerInput struct {

	// The name of the load balancer. This name must be unique per region per account,
	// can have a maximum of 32 characters, must contain only alphanumeric characters
	// or hyphens, must not begin or end with a hyphen, and must not begin with
	// "internal-".
	//
	// This member is required.
	Name *string

	// [Application Load Balancers on Outposts] The ID of the customer-owned address
	// pool (CoIP pool).
	CustomerOwnedIpv4Pool *string

	// The type of IP addresses used by the subnets for your load balancer. The
	// possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6
	// addresses). Internal load balancers must use ipv4.
	IpAddressType types.IpAddressType

	// The nodes of an Internet-facing load balancer have public IP addresses. The DNS
	// name of an Internet-facing load balancer is publicly resolvable to the public IP
	// addresses of the nodes. Therefore, Internet-facing load balancers can route
	// requests from clients over the internet. The nodes of an internal load balancer
	// have only private IP addresses. The DNS name of an internal load balancer is
	// publicly resolvable to the private IP addresses of the nodes. Therefore,
	// internal load balancers can route requests only from clients with access to the
	// VPC for the load balancer. The default is an Internet-facing load balancer. You
	// cannot specify a scheme for a Gateway Load Balancer.
	Scheme types.LoadBalancerSchemeEnum

	// [Application Load Balancers] The IDs of the security groups for the load
	// balancer.
	SecurityGroups []string

	// The IDs of the public subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings. [Application Load
	// Balancers] You must specify subnets from at least two Availability Zones. You
	// cannot specify Elastic IP addresses for your subnets. [Application Load
	// Balancers on Outposts] You must specify one Outpost subnet. [Application Load
	// Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	// [Network Load Balancers] You can specify subnets from one or more Availability
	// Zones. You can specify one Elastic IP address per subnet if you need static IP
	// addresses for your internet-facing load balancer. For internal load balancers,
	// you can specify one private IP address per subnet from the IPv4 range of the
	// subnet. For internet-facing load balancer, you can specify one IPv6 address per
	// subnet. [Gateway Load Balancers] You can specify subnets from one or more
	// Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	SubnetMappings []types.SubnetMapping

	// The IDs of the public subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings. [Application Load
	// Balancers] You must specify subnets from at least two Availability Zones.
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	// [Application Load Balancers on Local Zones] You can specify subnets from one or
	// more Local Zones. [Network Load Balancers] You can specify subnets from one or
	// more Availability Zones. [Gateway Load Balancers] You can specify subnets from
	// one or more Availability Zones.
	Subnets []string

	// The tags to assign to the load balancer.
	Tags []types.Tag

	// The type of load balancer. The default is application.
	Type types.LoadBalancerTypeEnum
}

type CreateLoadBalancerOutput struct {

	// Information about the load balancer.
	LoadBalancers []types.LoadBalancer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancer{}, middleware.After)
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
	if err = addOpCreateLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancer",
	}
}
