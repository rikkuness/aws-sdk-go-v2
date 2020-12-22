// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Autopilot job. Find the best performing model after you run an
// Autopilot job by calling . Deploy that model by following the steps described in
// Step 6.1: Deploy the Model to Amazon SageMaker Hosting Services
// (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-deploy-model.html). For
// information about how to use Autopilot, see  Automate Model Development with
// Amazon SageMaker Autopilot
// (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html).
func (c *Client) CreateAutoMLJob(ctx context.Context, params *CreateAutoMLJobInput, optFns ...func(*Options)) (*CreateAutoMLJobOutput, error) {
	if params == nil {
		params = &CreateAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoMLJob", params, optFns, addOperationCreateAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoMLJobInput struct {

	// Identifies an Autopilot job. Must be unique to your account and is
	// case-insensitive.
	//
	// This member is required.
	AutoMLJobName *string

	// Similar to InputDataConfig supported by Tuning. Format(s) supported: CSV.
	// Minimum of 500 rows.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Similar to OutputDataConfig supported by Tuning. Format(s) supported: CSV.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the role that is used to access the data.
	//
	// This member is required.
	RoleArn *string

	// Contains CompletionCriteria and SecurityConfig.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Defines the objective of a an AutoML job. You provide a
	// AutoMLJobObjective$MetricName and Autopilot infers whether to minimize or
	// maximize it. If a metric is not specified, the most commonly used
	// ObjectiveMetric for problem type is automaically selected.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Generates possible candidates without training a model. A candidate is a
	// combination of data preprocessors, algorithms, and algorithm parameter settings.
	GenerateCandidateDefinitionsOnly bool

	// Defines the kind of preprocessing and algorithms intended for the candidates.
	// Options include: BinaryClassification, MulticlassClassification, and Regression.
	ProblemType types.ProblemType

	// Each tag consists of a key and an optional value. Tag keys must be unique per
	// resource.
	Tags []types.Tag
}

type CreateAutoMLJobOutput struct {

	// When a job is created, it is assigned a unique ARN.
	//
	// This member is required.
	AutoMLJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoMLJob{}, middleware.After)
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
	if err = addOpCreateAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoMLJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAutoMLJob",
	}
}
