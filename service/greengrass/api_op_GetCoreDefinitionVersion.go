// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a core definition version.
func (c *Client) GetCoreDefinitionVersion(ctx context.Context, params *GetCoreDefinitionVersionInput, optFns ...func(*Options)) (*GetCoreDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetCoreDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoreDefinitionVersion", params, optFns, addOperationGetCoreDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoreDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoreDefinitionVersionInput struct {

	// The ID of the core definition.
	//
	// This member is required.
	CoreDefinitionId *string

	// The ID of the core definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListCoreDefinitionVersions'' requests. If the version is the last one that
	// was associated with a core definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	CoreDefinitionVersionId *string
}

type GetCoreDefinitionVersionOutput struct {

	// The ARN of the core definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the core definition version was
	// created.
	CreationTimestamp *string

	// Information about the core definition version.
	Definition *types.CoreDefinitionVersion

	// The ID of the core definition version.
	Id *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The version of the core definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCoreDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCoreDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCoreDefinitionVersion{}, middleware.After)
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
	if err = addOpGetCoreDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoreDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCoreDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetCoreDefinitionVersion",
	}
}
