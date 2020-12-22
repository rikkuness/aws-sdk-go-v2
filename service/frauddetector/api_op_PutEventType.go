// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an event type. An event is a business activity that is
// evaluated for fraud risk. With Amazon Fraud Detector, you generate fraud
// predictions for events. An event type defines the structure for an event sent to
// Amazon Fraud Detector. This includes the variables sent as part of the event,
// the entity performing the event (such as a customer), and the labels that
// classify the event. Example event types include online payment transactions,
// account registrations, and authentications.
func (c *Client) PutEventType(ctx context.Context, params *PutEventTypeInput, optFns ...func(*Options)) (*PutEventTypeOutput, error) {
	if params == nil {
		params = &PutEventTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEventType", params, optFns, addOperationPutEventTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventTypeInput struct {

	// The entity type for the event type. Example entity types: customer, merchant,
	// account.
	//
	// This member is required.
	EntityTypes []string

	// The event type variables.
	//
	// This member is required.
	EventVariables []string

	// The name.
	//
	// This member is required.
	Name *string

	// The description of the event type.
	Description *string

	// The event type labels.
	Labels []string

	// A collection of key and value pairs.
	Tags []types.Tag
}

type PutEventTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutEventTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEventType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEventType{}, middleware.After)
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
	if err = addOpPutEventTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEventType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "PutEventType",
	}
}
