// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stops logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func (c *Client) DisableLogging(ctx context.Context, params *DisableLoggingInput, optFns ...func(*Options)) (*DisableLoggingOutput, error) {
	if params == nil {
		params = &DisableLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableLogging", params, optFns, addOperationDisableLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DisableLoggingInput struct {

	// The identifier of the cluster on which logging is to be stopped. Example:
	// examplecluster
	//
	// This member is required.
	ClusterIdentifier *string
}

// Describes the status of logging for a cluster.
type DisableLoggingOutput struct {

	// The name of the S3 bucket where the log files are stored.
	BucketName *string

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// true if logging is on, false if logging is off.
	LoggingEnabled bool

	// The prefix applied to the log file names.
	S3KeyPrefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableLogging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableLogging{}, middleware.After)
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
	if err = addOpDisableLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableLogging(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DisableLogging",
	}
}
