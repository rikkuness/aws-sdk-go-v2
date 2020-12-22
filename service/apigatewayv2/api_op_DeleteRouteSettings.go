// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the RouteSettings for a stage.
func (c *Client) DeleteRouteSettings(ctx context.Context, params *DeleteRouteSettingsInput, optFns ...func(*Options)) (*DeleteRouteSettingsOutput, error) {
	if params == nil {
		params = &DeleteRouteSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRouteSettings", params, optFns, addOperationDeleteRouteSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRouteSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRouteSettingsInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The route key.
	//
	// This member is required.
	RouteKey *string

	// The stage name. Stage names can only contain alphanumeric characters, hyphens,
	// and underscores. Maximum length is 128 characters.
	//
	// This member is required.
	StageName *string
}

type DeleteRouteSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteRouteSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRouteSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRouteSettings{}, middleware.After)
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
	if err = addOpDeleteRouteSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRouteSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRouteSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteRouteSettings",
	}
}
