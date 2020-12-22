// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a target from a maintenance window.
func (c *Client) DeregisterTargetFromMaintenanceWindow(ctx context.Context, params *DeregisterTargetFromMaintenanceWindowInput, optFns ...func(*Options)) (*DeregisterTargetFromMaintenanceWindowOutput, error) {
	if params == nil {
		params = &DeregisterTargetFromMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterTargetFromMaintenanceWindow", params, optFns, addOperationDeregisterTargetFromMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterTargetFromMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTargetFromMaintenanceWindowInput struct {

	// The ID of the maintenance window the target should be removed from.
	//
	// This member is required.
	WindowId *string

	// The ID of the target definition to remove.
	//
	// This member is required.
	WindowTargetId *string

	// The system checks if the target is being referenced by a task. If the target is
	// being referenced, the system returns an error and does not deregister the target
	// from the maintenance window.
	Safe bool
}

type DeregisterTargetFromMaintenanceWindowOutput struct {

	// The ID of the maintenance window the target was removed from.
	WindowId *string

	// The ID of the removed target definition.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterTargetFromMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterTargetFromMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterTargetFromMaintenanceWindow{}, middleware.After)
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
	if err = addOpDeregisterTargetFromMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTargetFromMaintenanceWindow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterTargetFromMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeregisterTargetFromMaintenanceWindow",
	}
}
