// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a task definition. You can specify a family and revision to find
// information about a specific task definition, or you can simply specify the
// family to find the latest ACTIVE revision in that family. You can only describe
// INACTIVE task definitions while an active task or service references them.
func (c *Client) DescribeTaskDefinition(ctx context.Context, params *DescribeTaskDefinitionInput, optFns ...func(*Options)) (*DescribeTaskDefinitionOutput, error) {
	if params == nil {
		params = &DescribeTaskDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTaskDefinition", params, optFns, addOperationDescribeTaskDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTaskDefinitionInput struct {

	// The family for the latest ACTIVE revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN) of the
	// task definition to describe.
	//
	// This member is required.
	TaskDefinition *string

	// Specifies whether to see the resource tags for the task definition. If TAGS is
	// specified, the tags are included in the response. If this field is omitted, tags
	// are not included in the response.
	Include []types.TaskDefinitionField
}

type DescribeTaskDefinitionOutput struct {

	// The metadata that is applied to the task definition to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	// * Maximum number of
	// tags per resource - 50
	//
	// * For each resource, each tag key must be unique, and
	// each tag key can have only one value.
	//
	// * Maximum key length - 128 Unicode
	// characters in UTF-8
	//
	// * Maximum value length - 256 Unicode characters in UTF-8
	//
	// *
	// If your tagging schema is used across multiple services and resources, remember
	// that other services may have restrictions on allowed characters. Generally
	// allowed characters are: letters, numbers, and spaces representable in UTF-8, and
	// the following characters: + - = . _ : / @.
	//
	// * Tag keys and values are
	// case-sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or lowercase combination
	// of such as a prefix for either keys or values as it is reserved for AWS use. You
	// cannot edit or delete tag keys or values with this prefix. Tags with this prefix
	// do not count against your tags per resource limit.
	Tags []types.Tag

	// The full task definition description.
	TaskDefinition *types.TaskDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTaskDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskDefinition{}, middleware.After)
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
	if err = addOpDescribeTaskDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTaskDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeTaskDefinition",
	}
}
