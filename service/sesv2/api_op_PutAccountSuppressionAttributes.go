// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Change the settings for the account-level suppression list.
func (c *Client) PutAccountSuppressionAttributes(ctx context.Context, params *PutAccountSuppressionAttributesInput, optFns ...func(*Options)) (*PutAccountSuppressionAttributesOutput, error) {
	if params == nil {
		params = &PutAccountSuppressionAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountSuppressionAttributes", params, optFns, addOperationPutAccountSuppressionAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountSuppressionAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change your account's suppression preferences.
type PutAccountSuppressionAttributesInput struct {

	// A list that contains the reasons that email addresses will be automatically
	// added to the suppression list for your account. This list can contain any or all
	// of the following:
	//
	// * COMPLAINT – Amazon SES adds an email address to the
	// suppression list for your account when a message sent to that address results in
	// a complaint.
	//
	// * BOUNCE – Amazon SES adds an email address to the suppression
	// list for your account when a message sent to that address results in a hard
	// bounce.
	SuppressedReasons []types.SuppressionListReason
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutAccountSuppressionAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAccountSuppressionAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountSuppressionAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountSuppressionAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSuppressionAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountSuppressionAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutAccountSuppressionAttributes",
	}
}
