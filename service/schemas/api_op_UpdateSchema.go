// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the schema definition Inactive schemas will be deleted after two years.
func (c *Client) UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) {
	if params == nil {
		params = &UpdateSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSchema", params, optFns, addOperationUpdateSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSchemaInput struct {

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	// The ID of the client token.
	ClientTokenId *string

	// The source of the schema definition.
	Content *string

	// The description of the schema.
	Description *string

	// The schema type for the events schema.
	Type types.Type
}

type UpdateSchemaOutput struct {

	// The description of the schema.
	Description *string

	// The date and time that schema was modified.
	LastModified *time.Time

	// The ARN of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// The version number of the schema
	SchemaVersion *string

	// Key-value pairs associated with a resource.
	Tags map[string]string

	// The type of the schema.
	Type *string

	// The date the schema version was created.
	VersionCreatedDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSchema{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateSchemaMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSchema(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSchema struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSchema) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSchemaInput ")
	}

	if input.ClientTokenId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientTokenId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSchemaMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSchema{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "UpdateSchema",
	}
}
