// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an environment member to an AWS Cloud9 development environment.
func (c *Client) CreateEnvironmentMembership(ctx context.Context, params *CreateEnvironmentMembershipInput, optFns ...func(*Options)) (*CreateEnvironmentMembershipOutput, error) {
	if params == nil {
		params = &CreateEnvironmentMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentMembership", params, optFns, addOperationCreateEnvironmentMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentMembershipInput struct {

	// The ID of the environment that contains the environment member you want to add.
	//
	// This member is required.
	EnvironmentId *string

	// The type of environment member permissions you want to associate with this
	// environment member. Available values include:
	//
	// * read-only: Has read-only access
	// to the environment.
	//
	// * read-write: Has read-write access to the environment.
	//
	// This member is required.
	Permissions types.MemberPermissions

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	//
	// This member is required.
	UserArn *string
}

type CreateEnvironmentMembershipOutput struct {

	// Information about the environment member that was added.
	Membership *types.EnvironmentMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEnvironmentMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEnvironmentMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEnvironmentMembership{}, middleware.After)
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
	if err = addOpCreateEnvironmentMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "CreateEnvironmentMembership",
	}
}
