// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS IoT policy. The created policy is the default version for the
// policy. This operation creates a policy version with a version identifier of 1
// and sets 1 as the policy's default version.
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	if params == nil {
		params = &CreatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicy", params, optFns, addOperationCreatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreatePolicy operation.
type CreatePolicyInput struct {

	// The JSON document that describes the policy. policyDocument must have a minimum
	// length of 1, with a maximum length of 2048, excluding whitespace.
	//
	// This member is required.
	PolicyDocument *string

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// Metadata which can be used to manage the policy. For URI Request parameters use
	// format: ...key1=value1&key2=value2... For the CLI command-line parameter use
	// format: &&tags "key1=value1&key2=value2..." For the cli-input-json file use
	// format: "tags": "key1=value1&key2=value2..."
	Tags []types.Tag
}

// The output from the CreatePolicy operation.
type CreatePolicyOutput struct {

	// The policy ARN.
	PolicyArn *string

	// The JSON document that describes the policy.
	PolicyDocument *string

	// The policy name.
	PolicyName *string

	// The policy version ID.
	PolicyVersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePolicy{}, middleware.After)
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
	if err = addOpCreatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreatePolicy",
	}
}
