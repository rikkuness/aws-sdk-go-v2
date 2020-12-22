// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a GatewayResponse of a specified response type on the given RestApi.
func (c *Client) UpdateGatewayResponse(ctx context.Context, params *UpdateGatewayResponseInput, optFns ...func(*Options)) (*UpdateGatewayResponseOutput, error) {
	if params == nil {
		params = &UpdateGatewayResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGatewayResponse", params, optFns, addOperationUpdateGatewayResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGatewayResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a GatewayResponse of a specified response type on the given RestApi.
type UpdateGatewayResponseInput struct {

	// [Required] The response type of the associated GatewayResponse. Valid values
	// are
	//
	// * ACCESS_DENIED
	//
	// * API_CONFIGURATION_ERROR
	//
	// * AUTHORIZER_FAILURE
	//
	// *
	// AUTHORIZER_CONFIGURATION_ERROR
	//
	// * BAD_REQUEST_PARAMETERS
	//
	// * BAD_REQUEST_BODY
	//
	// *
	// DEFAULT_4XX
	//
	// * DEFAULT_5XX
	//
	// * EXPIRED_TOKEN
	//
	// * INVALID_SIGNATURE
	//
	// *
	// INTEGRATION_FAILURE
	//
	// * INTEGRATION_TIMEOUT
	//
	// * INVALID_API_KEY
	//
	// *
	// MISSING_AUTHENTICATION_TOKEN
	//
	// * QUOTA_EXCEEDED
	//
	// * REQUEST_TOO_LARGE
	//
	// *
	// RESOURCE_NOT_FOUND
	//
	// * THROTTLED
	//
	// * UNAUTHORIZED
	//
	// * UNSUPPORTED_MEDIA_TYPE
	//
	// This member is required.
	ResponseType types.GatewayResponseType

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []types.PatchOperation
}

// A gateway response of a given response type and status code, with optional
// response parameters and mapping templates. For more information about valid
// gateway response types, see Gateway Response Types Supported by API Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html)
// Example:
// Get a Gateway Response of a given response type
//
// Request
//
// This example shows how
// to get a gateway response of the MISSING_AUTHENTICATION_TOKEN type. GET
// /restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN HTTP/1.1
// Host: beta-apigateway.us-east-1.amazonaws.com Content-Type: application/json
// X-Amz-Date: 20170503T202516Z Authorization: AWS4-HMAC-SHA256
// Credential={access-key-id}/20170503/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date,
// Signature=1b52460e3159c1a26cff29093855d50ea141c1c5b937528fecaf60f51129697a
// Cache-Control: no-cache Postman-Token: 3b2a1ce9-c848-2e26-2e2f-9c2caefbed45  The
// response type is specified as a URL path.
// Response
//
// The successful operation
// returns the 200 OK status code and a payload similar to the following: {
// "_links": { "curies": { "href":
// "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-gatewayresponse-{rel}.html",
// "name": "gatewayresponse", "templated": true }, "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" },
// "gatewayresponse:delete": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" } },
// "defaultResponse": false, "responseParameters": {
// "gatewayresponse.header.x-request-path": "method.request.path.petId",
// "gatewayresponse.header.Access-Control-Allow-Origin": "'a.b.c'",
// "gatewayresponse.header.x-request-query": "method.request.querystring.q",
// "gatewayresponse.header.x-request-header": "method.request.header.Accept" },
// "responseTemplates": { "application/json": "{\n \"message\":
// $context.error.messageString,\n \"type\": \"$context.error.responseType\",\n
// \"stage\": \"$context.stage\",\n \"resourcePath\": \"$context.resourcePath\",\n
// \"stageVariables.a\": \"$stageVariables.a\",\n \"statusCode\": \"'404'\"\n}" },
// "responseType": "MISSING_AUTHENTICATION_TOKEN", "statusCode": "404" }Customize
// Gateway Responses
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html)
type UpdateGatewayResponseOutput struct {

	// A Boolean flag to indicate whether this GatewayResponse is the default gateway
	// response (true) or not (false). A default gateway response is one generated by
	// API Gateway without any customization by an API developer.
	DefaultResponse bool

	// Response parameters (paths, query strings and headers) of the GatewayResponse as
	// a string-to-string map of key-value pairs.
	ResponseParameters map[string]string

	// Response templates of the GatewayResponse as a string-to-string map of key-value
	// pairs.
	ResponseTemplates map[string]string

	// The response type of the associated GatewayResponse. Valid values are
	//
	// *
	// ACCESS_DENIED
	//
	// * API_CONFIGURATION_ERROR
	//
	// * AUTHORIZER_FAILURE
	//
	// *
	// AUTHORIZER_CONFIGURATION_ERROR
	//
	// * BAD_REQUEST_PARAMETERS
	//
	// * BAD_REQUEST_BODY
	//
	// *
	// DEFAULT_4XX
	//
	// * DEFAULT_5XX
	//
	// * EXPIRED_TOKEN
	//
	// * INVALID_SIGNATURE
	//
	// *
	// INTEGRATION_FAILURE
	//
	// * INTEGRATION_TIMEOUT
	//
	// * INVALID_API_KEY
	//
	// *
	// MISSING_AUTHENTICATION_TOKEN
	//
	// * QUOTA_EXCEEDED
	//
	// * REQUEST_TOO_LARGE
	//
	// *
	// RESOURCE_NOT_FOUND
	//
	// * THROTTLED
	//
	// * UNAUTHORIZED
	//
	// * UNSUPPORTED_MEDIA_TYPE
	ResponseType types.GatewayResponseType

	// The HTTP status code for this GatewayResponse.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGatewayResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGatewayResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGatewayResponse{}, middleware.After)
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
	if err = addOpUpdateGatewayResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewayResponse(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateGatewayResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateGatewayResponse",
	}
}
