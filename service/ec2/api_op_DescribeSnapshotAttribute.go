// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified attribute of the specified snapshot. You can specify
// only one attribute at a time. For more information about EBS snapshots, see
// Amazon EBS Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeSnapshotAttribute(ctx context.Context, params *DescribeSnapshotAttributeInput, optFns ...func(*Options)) (*DescribeSnapshotAttributeOutput, error) {
	if params == nil {
		params = &DescribeSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshotAttribute", params, optFns, addOperationDescribeSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotAttributeInput struct {

	// The snapshot attribute you would like to view.
	//
	// This member is required.
	Attribute types.SnapshotAttributeName

	// The ID of the EBS snapshot.
	//
	// This member is required.
	SnapshotId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type DescribeSnapshotAttributeOutput struct {

	// The users and groups that have the permissions for creating volumes from the
	// snapshot.
	CreateVolumePermissions []types.CreateVolumePermission

	// The product codes.
	ProductCodes []types.ProductCode

	// The ID of the EBS snapshot.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSnapshotAttribute{}, middleware.After)
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
	if err = addOpDescribeSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSnapshotAttribute",
	}
}
