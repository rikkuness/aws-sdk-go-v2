// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an analyzer for your account.
func (c *Client) CreateAnalyzer(ctx context.Context, params *CreateAnalyzerInput, optFns ...func(*Options)) (*CreateAnalyzerOutput, error) {
	if params == nil {
		params = &CreateAnalyzerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAnalyzer", params, optFns, addOperationCreateAnalyzerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAnalyzerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates an analyzer.
type CreateAnalyzerInput struct {

	// The name of the analyzer to create.
	//
	// This member is required.
	AnalyzerName *string

	// The type of analyzer to create. Only ACCOUNT analyzers are supported. You can
	// create only one analyzer per account per Region.
	//
	// This member is required.
	Type types.Type

	// Specifies the archive rules to add for the analyzer. Archive rules automatically
	// archive findings that meet the criteria you define for the rule.
	ArchiveRules []types.InlineArchiveRule

	// A client token.
	ClientToken *string

	// The tags to apply to the analyzer.
	Tags map[string]string
}

// The response to the request to create an analyzer.
type CreateAnalyzerOutput struct {

	// The ARN of the analyzer that was created by the request.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAnalyzerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAnalyzer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAnalyzer{}, middleware.After)
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
	if err = addOpCreateAnalyzerValidationMiddleware(stack); err != nil {
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
