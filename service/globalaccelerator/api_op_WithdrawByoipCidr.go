// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops advertising an address range that is provisioned as an address pool. You
// can perform this operation at most once every 10 seconds, even if you specify
// different address ranges each time. To see an AWS CLI example of withdrawing an
// address range for BYOIP so it will no longer be advertised by AWS, scroll down
// to Example. It can take a few minutes before traffic to the specified addresses
// stops routing to AWS because of propagation delays. For more information, see
// Bring Your Own IP Addresses (BYOIP)
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in
// the AWS Global Accelerator Developer Guide.
func (c *Client) WithdrawByoipCidr(ctx context.Context, params *WithdrawByoipCidrInput, optFns ...func(*Options)) (*WithdrawByoipCidrOutput, error) {
	if params == nil {
		params = &WithdrawByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "WithdrawByoipCidr", params, optFns, addOperationWithdrawByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*WithdrawByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type WithdrawByoipCidrInput struct {

	// The address range, in CIDR notation.
	//
	// This member is required.
	Cidr *string
}

type WithdrawByoipCidrOutput struct {

	// Information about the address pool.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationWithdrawByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpWithdrawByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpWithdrawByoipCidr{}, middleware.After)
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
	if err = addOpWithdrawByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opWithdrawByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opWithdrawByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "WithdrawByoipCidr",
	}
}
