// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a training job.
func (c *Client) DescribeTrainingJob(ctx context.Context, params *DescribeTrainingJobInput, optFns ...func(*Options)) (*DescribeTrainingJobOutput, error) {
	if params == nil {
		params = &DescribeTrainingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrainingJob", params, optFns, addOperationDescribeTrainingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrainingJobInput struct {

	// The name of the training job.
	//
	// This member is required.
	TrainingJobName *string
}

type DescribeTrainingJobOutput struct {

	// Information about the algorithm used for training, and algorithm metadata.
	//
	// This member is required.
	AlgorithmSpecification *types.AlgorithmSpecification

	// A timestamp that indicates when the training job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Information about the Amazon S3 location that is configured for storing model
	// artifacts.
	//
	// This member is required.
	ModelArtifacts *types.ModelArtifacts

	// Resources, including ML compute instances and ML storage volumes, that are
	// configured for model training.
	//
	// This member is required.
	ResourceConfig *types.ResourceConfig

	// Provides detailed information about the state of the training job. For detailed
	// information on the secondary status of the training job, see StatusMessage under
	// SecondaryStatusTransition. Amazon SageMaker provides primary statuses and
	// secondary statuses that apply to each of them: InProgress
	//
	// * Starting - Starting
	// the training job.
	//
	// * Downloading - An optional stage for algorithms that support
	// File training input mode. It indicates that data is being downloaded to the ML
	// storage volumes.
	//
	// * Training - Training is in progress.
	//
	// * Interrupted - The job
	// stopped because the managed spot training instances were interrupted.
	//
	// *
	// Uploading - Training is complete and the model artifacts are being uploaded to
	// the S3 location.
	//
	// Completed
	//
	// * Completed - The training job has
	// completed.
	//
	// Failed
	//
	// * Failed - The training job has failed. The reason for the
	// failure is returned in the FailureReason field of
	// DescribeTrainingJobResponse.
	//
	// Stopped
	//
	// * MaxRuntimeExceeded - The job stopped
	// because it exceeded the maximum allowed runtime.
	//
	// * MaxWaitTimeExceeded - The
	// job stopped because it exceeded the maximum allowed wait time.
	//
	// * Stopped - The
	// training job has stopped.
	//
	// Stopping
	//
	// * Stopping - Stopping the training
	// job.
	//
	// Valid values for SecondaryStatus are subject to change. We no longer
	// support the following secondary statuses:
	//
	// * LaunchingMLInstances
	//
	// *
	// PreparingTrainingStack
	//
	// * DownloadingTrainingImage
	//
	// This member is required.
	SecondaryStatus types.SecondaryStatus

	// Specifies a limit to how long a model training job can run. It also specifies
	// the maximum time to wait for a spot instance. When the job reaches the time
	// limit, Amazon SageMaker ends the training job. Use this API to cap model
	// training costs. To stop a job, Amazon SageMaker sends the algorithm the SIGTERM
	// signal, which delays job termination for 120 seconds. Algorithms can use this
	// 120-second window to save the model artifacts, so the results of training are
	// not lost.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// The Amazon Resource Name (ARN) of the training job.
	//
	// This member is required.
	TrainingJobArn *string

	// Name of the model training job.
	//
	// This member is required.
	TrainingJobName *string

	// The status of the training job. Amazon SageMaker provides the following training
	// job statuses:
	//
	// * InProgress - The training is in progress.
	//
	// * Completed - The
	// training job has completed.
	//
	// * Failed - The training job has failed. To see the
	// reason for the failure, see the FailureReason field in the response to a
	// DescribeTrainingJobResponse call.
	//
	// * Stopping - The training job is stopping.
	//
	// *
	// Stopped - The training job has stopped.
	//
	// For more detailed information, see
	// SecondaryStatus.
	//
	// This member is required.
	TrainingJobStatus types.TrainingJobStatus

	// The Amazon Resource Name (ARN) of an AutoML job.
	AutoMLJobArn *string

	// The billable time in seconds. You can calculate the savings from using managed
	// spot training using the formula (1 - BillableTimeInSeconds /
	// TrainingTimeInSeconds) * 100. For example, if BillableTimeInSeconds is 100 and
	// TrainingTimeInSeconds is 500, the savings is 80%.
	BillableTimeInSeconds *int32

	// Contains information about the output location for managed spot training
	// checkpoint data.
	CheckpointConfig *types.CheckpointConfig

	// Configuration information for the debug hook parameters, collection
	// configuration, and storage paths.
	DebugHookConfig *types.DebugHookConfig

	// Configuration information for debugging rules.
	DebugRuleConfigurations []types.DebugRuleConfiguration

	// Status about the debug rule evaluation.
	DebugRuleEvaluationStatuses []types.DebugRuleEvaluationStatus

	// To encrypt all communications between ML compute instances in distributed
	// training, choose True. Encryption provides greater security for distributed
	// training, but training might take longer. How long it takes depends on the
	// amount of communication between compute instances, especially if you use a deep
	// learning algorithms in distributed training.
	EnableInterContainerTrafficEncryption bool

	// A Boolean indicating whether managed spot training is enabled (True) or not
	// (False).
	EnableManagedSpotTraining bool

	// If you want to allow inbound or outbound network calls, except for calls between
	// peers within a training cluster for distributed training, choose True. If you
	// enable network isolation for training jobs that are configured to use a VPC,
	// Amazon SageMaker downloads and uploads customer data and model artifacts through
	// the specified VPC, but the training container does not have network access.
	EnableNetworkIsolation bool

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	// * CreateProcessingJob
	//
	// *
	// CreateTrainingJob
	//
	// * CreateTransformJob
	ExperimentConfig *types.ExperimentConfig

	// If the training job failed, the reason it failed.
	FailureReason *string

	// A collection of MetricData objects that specify the names, values, and dates and
	// times that the training algorithm emitted to Amazon CloudWatch.
	FinalMetricDataList []types.MetricData

	// Algorithm-specific parameters.
	HyperParameters map[string]string

	// An array of Channel objects that describes each data input channel.
	InputDataConfig []types.Channel

	// The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job
	// that created the transform or training job.
	LabelingJobArn *string

	// A timestamp that indicates when the status of the training job was last
	// modified.
	LastModifiedTime *time.Time

	// The S3 path where model artifacts that you configured when creating the job are
	// stored. Amazon SageMaker creates subfolders for model artifacts.
	OutputDataConfig *types.OutputDataConfig

	// The AWS Identity and Access Management (IAM) role configured for the training
	// job.
	RoleArn *string

	// A history of all of the secondary statuses that the training job has
	// transitioned through.
	SecondaryStatusTransitions []types.SecondaryStatusTransition

	// Configuration of storage locations for TensorBoard output.
	TensorBoardOutputConfig *types.TensorBoardOutputConfig

	// Indicates the time when the training job ends on training instances. You are
	// billed for the time interval between the value of TrainingStartTime and this
	// time. For successful jobs and stopped jobs, this is the time after model
	// artifacts are uploaded. For failed jobs, this is the time when Amazon SageMaker
	// detects a job failure.
	TrainingEndTime *time.Time

	// Indicates the time when the training job starts on training instances. You are
	// billed for the time interval between this time and the value of TrainingEndTime.
	// The start time in CloudWatch Logs might be later than this time. The difference
	// is due to the time it takes to download the training data and to the size of the
	// training container.
	TrainingStartTime *time.Time

	// The training time in seconds.
	TrainingTimeInSeconds *int32

	// The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if
	// the training job was launched by a hyperparameter tuning job.
	TuningJobArn *string

	// A VpcConfig object that specifies the VPC that this training job has access to.
	// For more information, see Protect Training Jobs by Using an Amazon Virtual
	// Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).
	VpcConfig *types.VpcConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrainingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrainingJob{}, middleware.After)
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
	if err = addOpDescribeTrainingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrainingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrainingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeTrainingJob",
	}
}
