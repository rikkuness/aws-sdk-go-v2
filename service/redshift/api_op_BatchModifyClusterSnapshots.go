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

// Modifies the settings for a set of cluster snapshots.
func (c *Client) BatchModifyClusterSnapshots(ctx context.Context, params *BatchModifyClusterSnapshotsInput, optFns ...func(*Options)) (*BatchModifyClusterSnapshotsOutput, error) {
	if params == nil {
		params = &BatchModifyClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchModifyClusterSnapshots", params, optFns, addOperationBatchModifyClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchModifyClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchModifyClusterSnapshotsInput struct {

	// A list of snapshot identifiers you want to modify.
	//
	// This member is required.
	SnapshotIdentifierList []string

	// A boolean value indicating whether to override an exception if the retention
	// period has passed.
	Force bool

	// The number of days that a manual snapshot is retained. If you specify the value
	// -1, the manual snapshot is retained indefinitely. The number must be either -1
	// or an integer between 1 and 3,653. If you decrease the manual snapshot retention
	// period from its current value, existing manual snapshots that fall outside of
	// the new retention period will return an error. If you want to suppress the
	// errors and delete the snapshots, use the force option.
	ManualSnapshotRetentionPeriod *int32
}

type BatchModifyClusterSnapshotsOutput struct {

	// A list of any errors returned.
	Errors []types.SnapshotErrorMessage

	// A list of the snapshots that were modified.
	Resources []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchModifyClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchModifyClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchModifyClusterSnapshots{}, middleware.After)
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
	if err = addOpBatchModifyClusterSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchModifyClusterSnapshots(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchModifyClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "BatchModifyClusterSnapshots",
	}
}
