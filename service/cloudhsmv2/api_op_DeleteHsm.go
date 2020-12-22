// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified HSM. To specify an HSM, you can use its identifier (ID),
// the IP address of the HSM's elastic network interface (ENI), or the ID of the
// HSM's ENI. You need to specify only one of these values. To find these values,
// use DescribeClusters.
func (c *Client) DeleteHsm(ctx context.Context, params *DeleteHsmInput, optFns ...func(*Options)) (*DeleteHsmOutput, error) {
	if params == nil {
		params = &DeleteHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHsm", params, optFns, addOperationDeleteHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteHsmInput struct {

	// The identifier (ID) of the cluster that contains the HSM that you are deleting.
	//
	// This member is required.
	ClusterId *string

	// The identifier (ID) of the elastic network interface (ENI) of the HSM that you
	// are deleting.
	EniId *string

	// The IP address of the elastic network interface (ENI) of the HSM that you are
	// deleting.
	EniIp *string

	// The identifier (ID) of the HSM that you are deleting.
	HsmId *string
}

type DeleteHsmOutput struct {

	// The identifier (ID) of the HSM that was deleted.
	HsmId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHsm{}, middleware.After)
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
	if err = addOpDeleteHsmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHsm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DeleteHsm",
	}
}
