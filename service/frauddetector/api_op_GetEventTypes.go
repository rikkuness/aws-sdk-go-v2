// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all event types or a specific event type if name is provided. This is a
// paginated API. If you provide a null maxResults, this action retrieves a maximum
// of 10 records per page. If you provide a maxResults, the value must be between 5
// and 10. To get the next page results, provide the pagination token from the
// GetEventTypesResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func (c *Client) GetEventTypes(ctx context.Context, params *GetEventTypesInput, optFns ...func(*Options)) (*GetEventTypesOutput, error) {
	if params == nil {
		params = &GetEventTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventTypes", params, optFns, addOperationGetEventTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventTypesInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name.
	Name *string

	// The next token for the subsequent request.
	NextToken *string
}

type GetEventTypesOutput struct {

	// An array of event types.
	EventTypes []*types.EventType

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEventTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEventTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEventTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetEventTypes",
	}
}