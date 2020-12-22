// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a project, including project resources. Does not delete users associated
// with the project, but does delete the IAM roles that allowed access to the
// project.
func (c *Client) DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) {
	if params == nil {
		params = &DeleteProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProject", params, optFns, addOperationDeleteProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProjectInput struct {

	// The ID of the project to be deleted in AWS CodeStar.
	//
	// This member is required.
	Id *string

	// A user- or system-generated token that identifies the entity that requested
	// project deletion. This token can be used to repeat the request.
	ClientRequestToken *string

	// Whether to send a delete request for the primary stack in AWS CloudFormation
	// originally used to generate the project and its resources. This option will
	// delete all AWS resources for the project (except for any buckets in Amazon S3)
	// as well as deleting the project itself. Recommended for most use cases.
	DeleteStack bool
}

type DeleteProjectOutput struct {

	// The Amazon Resource Name (ARN) of the deleted project.
	ProjectArn *string

	// The ID of the primary stack in AWS CloudFormation that will be deleted as part
	// of deleting the project and its resources.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProject{}, middleware.After)
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
	if err = addOpDeleteProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "DeleteProject",
	}
}
