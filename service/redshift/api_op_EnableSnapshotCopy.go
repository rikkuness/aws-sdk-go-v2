// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the automatic copy of snapshots from one region to another region for a
// specified cluster.
func (c *Client) EnableSnapshotCopy(ctx context.Context, params *EnableSnapshotCopyInput, optFns ...func(*Options)) (*EnableSnapshotCopyOutput, error) {
	if params == nil {
		params = &EnableSnapshotCopyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableSnapshotCopy", params, optFns, addOperationEnableSnapshotCopyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableSnapshotCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type EnableSnapshotCopyInput struct {

	// The unique identifier of the source cluster to copy snapshots from. Constraints:
	// Must be the valid name of an existing cluster that does not already have
	// cross-region snapshot copy enabled.
	//
	// This member is required.
	ClusterIdentifier *string

	// The destination AWS Region that you want to copy snapshots to. Constraints: Must
	// be the name of a valid AWS Region. For more information, see Regions and
	// Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	DestinationRegion *string

	// The number of days to retain newly copied snapshots in the destination AWS
	// Region after they are copied from the source AWS Region. If the value is -1, the
	// manual snapshot is retained indefinitely. The value must be either -1 or an
	// integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int32

	// The number of days to retain automated snapshots in the destination region after
	// they are copied from the source region. Default: 7. Constraints: Must be at
	// least 1 and no more than 35.
	RetentionPeriod *int32

	// The name of the snapshot copy grant to use when snapshots of an AWS
	// KMS-encrypted cluster are copied to the destination region.
	SnapshotCopyGrantName *string
}

type EnableSnapshotCopyOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableSnapshotCopyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableSnapshotCopy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableSnapshotCopy{}, middleware.After)
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
	if err = addOpEnableSnapshotCopyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSnapshotCopy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableSnapshotCopy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "EnableSnapshotCopy",
	}
}
