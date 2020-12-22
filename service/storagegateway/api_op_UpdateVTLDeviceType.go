// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the type of medium changer in a tape gateway. When you activate a tape
// gateway, you select a medium changer type for the tape gateway. This operation
// enables you to select a different type of medium changer after a tape gateway is
// activated. This operation is only supported in the tape gateway type.
func (c *Client) UpdateVTLDeviceType(ctx context.Context, params *UpdateVTLDeviceTypeInput, optFns ...func(*Options)) (*UpdateVTLDeviceTypeOutput, error) {
	if params == nil {
		params = &UpdateVTLDeviceTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVTLDeviceType", params, optFns, addOperationUpdateVTLDeviceTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVTLDeviceTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVTLDeviceTypeInput struct {

	// The type of medium changer you want to select. Valid Values: STK-L700 |
	// AWS-Gateway-VTL | IBM-03584L32-0402
	//
	// This member is required.
	DeviceType *string

	// The Amazon Resource Name (ARN) of the medium changer you want to select.
	//
	// This member is required.
	VTLDeviceARN *string
}

// UpdateVTLDeviceTypeOutput
type UpdateVTLDeviceTypeOutput struct {

	// The Amazon Resource Name (ARN) of the medium changer you have selected.
	VTLDeviceARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateVTLDeviceTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVTLDeviceType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVTLDeviceType{}, middleware.After)
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
	if err = addOpUpdateVTLDeviceTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVTLDeviceType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVTLDeviceType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateVTLDeviceType",
	}
}
