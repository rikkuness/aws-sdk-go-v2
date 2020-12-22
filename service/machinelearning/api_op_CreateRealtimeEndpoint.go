// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a real-time endpoint for the MLModel. The endpoint contains the URI of
// the MLModel; that is, the location to send real-time prediction requests for the
// specified MLModel.
func (c *Client) CreateRealtimeEndpoint(ctx context.Context, params *CreateRealtimeEndpointInput, optFns ...func(*Options)) (*CreateRealtimeEndpointOutput, error) {
	if params == nil {
		params = &CreateRealtimeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRealtimeEndpoint", params, optFns, addOperationCreateRealtimeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRealtimeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRealtimeEndpointInput struct {

	// The ID assigned to the MLModel during creation.
	//
	// This member is required.
	MLModelId *string
}

// Represents the output of an CreateRealtimeEndpoint operation. The result
// contains the MLModelId and the endpoint information for the MLModel. Note: The
// endpoint information includes the URI of the MLModel; that is, the location to
// send online prediction requests for the specified MLModel.
type CreateRealtimeEndpointOutput struct {

	// A user-supplied ID that uniquely identifies the MLModel. This value should be
	// identical to the value of the MLModelId in the request.
	MLModelId *string

	// The endpoint information of the MLModel
	RealtimeEndpointInfo *types.RealtimeEndpointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRealtimeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRealtimeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRealtimeEndpoint{}, middleware.After)
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
	if err = addOpCreateRealtimeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRealtimeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRealtimeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateRealtimeEndpoint",
	}
}
