// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the URL of an existing Amazon SQS queue. To access a queue that belongs
// to another AWS account, use the QueueOwnerAWSAccountId parameter to specify the
// account ID of the queue's owner. The queue's owner must grant you permission to
// access the queue. For more information about shared queue access, see
// AddPermission or see Allow Developers to Write Messages to a Shared Queue
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) GetQueueUrl(ctx context.Context, params *GetQueueUrlInput, optFns ...func(*Options)) (*GetQueueUrlOutput, error) {
	if params == nil {
		params = &GetQueueUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueueUrl", params, optFns, addOperationGetQueueUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetQueueUrlInput struct {

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens (-), and underscores (_). Queue URLs
	// and names are case-sensitive.
	//
	// This member is required.
	QueueName *string

	// The AWS account ID of the account that created the queue.
	QueueOwnerAWSAccountId *string
}

// For more information, see Interpreting Responses
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
// in the Amazon Simple Queue Service Developer Guide.
type GetQueueUrlOutput struct {

	// The URL of the queue.
	QueueUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQueueUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetQueueUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetQueueUrl{}, middleware.After)
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
	if err = addOpGetQueueUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetQueueUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "GetQueueUrl",
	}
}
