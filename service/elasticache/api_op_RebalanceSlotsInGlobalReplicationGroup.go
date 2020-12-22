// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Redistribute slots to ensure uniform distribution across existing shards in the
// cluster.
func (c *Client) RebalanceSlotsInGlobalReplicationGroup(ctx context.Context, params *RebalanceSlotsInGlobalReplicationGroupInput, optFns ...func(*Options)) (*RebalanceSlotsInGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &RebalanceSlotsInGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebalanceSlotsInGlobalReplicationGroup", params, optFns, addOperationRebalanceSlotsInGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebalanceSlotsInGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebalanceSlotsInGlobalReplicationGroupInput struct {

	// If True, redistribution is applied immediately.
	//
	// This member is required.
	ApplyImmediately bool

	// The name of the Global Datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string
}

type RebalanceSlotsInGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//
	// * The GlobalReplicationGroupIdSuffix represents the name of
	// the Global Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRebalanceSlotsInGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebalanceSlotsInGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebalanceSlotsInGlobalReplicationGroup{}, middleware.After)
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
	if err = addOpRebalanceSlotsInGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebalanceSlotsInGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebalanceSlotsInGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "RebalanceSlotsInGlobalReplicationGroup",
	}
}
