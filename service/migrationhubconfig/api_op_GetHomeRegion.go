// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the calling account’s home region, if configured. This API is used by
// other AWS services to determine the regional endpoint for calling AWS
// Application Discovery Service and Migration Hub. You must call GetHomeRegion at
// least once before you call any other AWS Application Discovery Service and AWS
// Migration Hub APIs, to obtain the account's Migration Hub home region.
func (c *Client) GetHomeRegion(ctx context.Context, params *GetHomeRegionInput, optFns ...func(*Options)) (*GetHomeRegionOutput, error) {
	if params == nil {
		params = &GetHomeRegionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHomeRegion", params, optFns, addOperationGetHomeRegionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHomeRegionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHomeRegionInput struct {
}

type GetHomeRegionOutput struct {

	// The name of the home region of the calling account.
	HomeRegion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHomeRegionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetHomeRegion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetHomeRegion{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHomeRegion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHomeRegion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "GetHomeRegion",
	}
}
