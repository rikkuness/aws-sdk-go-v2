// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the attachment that has the specified ID. Attachments can include
// screenshots, error logs, or other files that describe your issue. Attachment IDs
// are generated by the case management system when you add an attachment to a case
// or case communication. Attachment IDs are returned in the AttachmentDetails
// objects that are returned by the DescribeCommunications operation.
//
// * You must
// have a Business or Enterprise support plan to use the AWS Support API.
//
// * If you
// call the AWS Support API from an account that does not have a Business or
// Enterprise support plan, the SubscriptionRequiredException error message
// appears. For information about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeAttachment(ctx context.Context, params *DescribeAttachmentInput, optFns ...func(*Options)) (*DescribeAttachmentOutput, error) {
	if params == nil {
		params = &DescribeAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAttachment", params, optFns, addOperationDescribeAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAttachmentInput struct {

	// The ID of the attachment to return. Attachment IDs are returned by the
	// DescribeCommunications operation.
	//
	// This member is required.
	AttachmentId *string
}

// The content and file name of the attachment returned by the DescribeAttachment
// operation.
type DescribeAttachmentOutput struct {

	// This object includes the attachment content and file name. In the previous
	// response syntax, the value for the data parameter appears as blob, which is
	// represented as a base64-encoded string. The value for fileName is the name of
	// the attachment, such as troubleshoot-screenshot.png.
	Attachment *types.Attachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAttachment{}, middleware.After)
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
	if err = addOpDescribeAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeAttachment",
	}
}
