// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideosignaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API allows you to connect WebRTC-enabled devices with Alexa display
// devices. When invoked, it sends the Alexa Session Description Protocol (SDP)
// offer to the master peer. The offer is delivered as soon as the master is
// connected to the specified signaling channel. This API returns the SDP answer
// from the connected master. If the master is not connected to the signaling
// channel, redelivery requests are made until the message expires.
func (c *Client) SendAlexaOfferToMaster(ctx context.Context, params *SendAlexaOfferToMasterInput, optFns ...func(*Options)) (*SendAlexaOfferToMasterOutput, error) {
	if params == nil {
		params = &SendAlexaOfferToMasterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendAlexaOfferToMaster", params, optFns, addOperationSendAlexaOfferToMasterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendAlexaOfferToMasterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendAlexaOfferToMasterInput struct {

	// The ARN of the signaling channel by which Alexa and the master peer communicate.
	//
	// This member is required.
	ChannelARN *string

	// The base64-encoded SDP offer content.
	//
	// This member is required.
	MessagePayload *string

	// The unique identifier for the sender client.
	//
	// This member is required.
	SenderClientId *string
}

type SendAlexaOfferToMasterOutput struct {

	// The base64-encoded SDP answer content.
	Answer *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendAlexaOfferToMasterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendAlexaOfferToMaster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendAlexaOfferToMaster{}, middleware.After)
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
	if err = addOpSendAlexaOfferToMasterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendAlexaOfferToMaster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendAlexaOfferToMaster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "SendAlexaOfferToMaster",
	}
}
