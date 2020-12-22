// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelJob struct {
}

func (*validateOpCancelJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateComputeEnvironment struct {
}

func (*validateOpCreateComputeEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateComputeEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateComputeEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateComputeEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateJobQueue struct {
}

func (*validateOpCreateJobQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateJobQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateJobQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateJobQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteComputeEnvironment struct {
}

func (*validateOpDeleteComputeEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteComputeEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteComputeEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteComputeEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteJobQueue struct {
}

func (*validateOpDeleteJobQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteJobQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteJobQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteJobQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeregisterJobDefinition struct {
}

func (*validateOpDeregisterJobDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeregisterJobDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeregisterJobDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeregisterJobDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJobs struct {
}

func (*validateOpDescribeJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterJobDefinition struct {
}

func (*validateOpRegisterJobDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterJobDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterJobDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterJobDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSubmitJob struct {
}

func (*validateOpSubmitJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSubmitJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SubmitJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSubmitJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTerminateJob struct {
}

func (*validateOpTerminateJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTerminateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TerminateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTerminateJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateComputeEnvironment struct {
}

func (*validateOpUpdateComputeEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateComputeEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateComputeEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateComputeEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateJobQueue struct {
}

func (*validateOpUpdateJobQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateJobQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateJobQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateJobQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelJob{}, middleware.After)
}

func addOpCreateComputeEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateComputeEnvironment{}, middleware.After)
}

func addOpCreateJobQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateJobQueue{}, middleware.After)
}

func addOpDeleteComputeEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteComputeEnvironment{}, middleware.After)
}

func addOpDeleteJobQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteJobQueue{}, middleware.After)
}

func addOpDeregisterJobDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeregisterJobDefinition{}, middleware.After)
}

func addOpDescribeJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJobs{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRegisterJobDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterJobDefinition{}, middleware.After)
}

func addOpSubmitJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSubmitJob{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpTerminateJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTerminateJob{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateComputeEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateComputeEnvironment{}, middleware.After)
}

func addOpUpdateJobQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateJobQueue{}, middleware.After)
}

func validateComputeEnvironmentOrder(v *types.ComputeEnvironmentOrder) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComputeEnvironmentOrder"}
	if v.ComputeEnvironment == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComputeEnvironment"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComputeEnvironmentOrders(v []types.ComputeEnvironmentOrder) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComputeEnvironmentOrders"}
	for i := range v {
		if err := validateComputeEnvironmentOrder(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComputeResource(v *types.ComputeResource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComputeResource"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.InstanceRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceRole"))
	}
	if v.InstanceTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceTypes"))
	}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContainerOverrides(v *types.ContainerOverrides) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContainerOverrides"}
	if v.ResourceRequirements != nil {
		if err := validateResourceRequirements(v.ResourceRequirements); err != nil {
			invalidParams.AddNested("ResourceRequirements", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContainerProperties(v *types.ContainerProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContainerProperties"}
	if v.Ulimits != nil {
		if err := validateUlimits(v.Ulimits); err != nil {
			invalidParams.AddNested("Ulimits", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceRequirements != nil {
		if err := validateResourceRequirements(v.ResourceRequirements); err != nil {
			invalidParams.AddNested("ResourceRequirements", err.(smithy.InvalidParamsError))
		}
	}
	if v.LinuxParameters != nil {
		if err := validateLinuxParameters(v.LinuxParameters); err != nil {
			invalidParams.AddNested("LinuxParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.LogConfiguration != nil {
		if err := validateLogConfiguration(v.LogConfiguration); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.Secrets != nil {
		if err := validateSecretList(v.Secrets); err != nil {
			invalidParams.AddNested("Secrets", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDevice(v *types.Device) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Device"}
	if v.HostPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HostPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDevicesList(v []types.Device) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DevicesList"}
	for i := range v {
		if err := validateDevice(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEvaluateOnExit(v *types.EvaluateOnExit) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EvaluateOnExit"}
	if len(v.Action) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEvaluateOnExitList(v []types.EvaluateOnExit) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EvaluateOnExitList"}
	for i := range v {
		if err := validateEvaluateOnExit(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLinuxParameters(v *types.LinuxParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LinuxParameters"}
	if v.Tmpfs != nil {
		if err := validateTmpfsList(v.Tmpfs); err != nil {
			invalidParams.AddNested("Tmpfs", err.(smithy.InvalidParamsError))
		}
	}
	if v.Devices != nil {
		if err := validateDevicesList(v.Devices); err != nil {
			invalidParams.AddNested("Devices", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLogConfiguration(v *types.LogConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LogConfiguration"}
	if v.SecretOptions != nil {
		if err := validateSecretList(v.SecretOptions); err != nil {
			invalidParams.AddNested("SecretOptions", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.LogDriver) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LogDriver"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodeOverrides(v *types.NodeOverrides) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodeOverrides"}
	if v.NodePropertyOverrides != nil {
		if err := validateNodePropertyOverrides(v.NodePropertyOverrides); err != nil {
			invalidParams.AddNested("NodePropertyOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodeProperties(v *types.NodeProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodeProperties"}
	if v.NodeRangeProperties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeRangeProperties"))
	} else if v.NodeRangeProperties != nil {
		if err := validateNodeRangeProperties(v.NodeRangeProperties); err != nil {
			invalidParams.AddNested("NodeRangeProperties", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodePropertyOverride(v *types.NodePropertyOverride) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodePropertyOverride"}
	if v.ContainerOverrides != nil {
		if err := validateContainerOverrides(v.ContainerOverrides); err != nil {
			invalidParams.AddNested("ContainerOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetNodes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetNodes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodePropertyOverrides(v []types.NodePropertyOverride) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodePropertyOverrides"}
	for i := range v {
		if err := validateNodePropertyOverride(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodeRangeProperties(v []types.NodeRangeProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodeRangeProperties"}
	for i := range v {
		if err := validateNodeRangeProperty(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodeRangeProperty(v *types.NodeRangeProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodeRangeProperty"}
	if v.Container != nil {
		if err := validateContainerProperties(v.Container); err != nil {
			invalidParams.AddNested("Container", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetNodes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetNodes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceRequirement(v *types.ResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceRequirement"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceRequirements(v []types.ResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceRequirements"}
	for i := range v {
		if err := validateResourceRequirement(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetryStrategy(v *types.RetryStrategy) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetryStrategy"}
	if v.EvaluateOnExit != nil {
		if err := validateEvaluateOnExitList(v.EvaluateOnExit); err != nil {
			invalidParams.AddNested("EvaluateOnExit", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSecret(v *types.Secret) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Secret"}
	if v.ValueFrom == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ValueFrom"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSecretList(v []types.Secret) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SecretList"}
	for i := range v {
		if err := validateSecret(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTmpfs(v *types.Tmpfs) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tmpfs"}
	if v.ContainerPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTmpfsList(v []types.Tmpfs) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TmpfsList"}
	for i := range v {
		if err := validateTmpfs(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUlimit(v *types.Ulimit) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Ulimit"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUlimits(v []types.Ulimit) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Ulimits"}
	for i := range v {
		if err := validateUlimit(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelJobInput(v *CancelJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelJobInput"}
	if v.Reason == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Reason"))
	}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateComputeEnvironmentInput(v *CreateComputeEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateComputeEnvironmentInput"}
	if v.ComputeEnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComputeEnvironmentName"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.ComputeResources != nil {
		if err := validateComputeResource(v.ComputeResources); err != nil {
			invalidParams.AddNested("ComputeResources", err.(smithy.InvalidParamsError))
		}
	}
	if v.ServiceRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceRole"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateJobQueueInput(v *CreateJobQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateJobQueueInput"}
	if v.JobQueueName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobQueueName"))
	}
	if v.ComputeEnvironmentOrder == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComputeEnvironmentOrder"))
	} else if v.ComputeEnvironmentOrder != nil {
		if err := validateComputeEnvironmentOrders(v.ComputeEnvironmentOrder); err != nil {
			invalidParams.AddNested("ComputeEnvironmentOrder", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteComputeEnvironmentInput(v *DeleteComputeEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteComputeEnvironmentInput"}
	if v.ComputeEnvironment == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComputeEnvironment"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteJobQueueInput(v *DeleteJobQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteJobQueueInput"}
	if v.JobQueue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobQueue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeregisterJobDefinitionInput(v *DeregisterJobDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeregisterJobDefinitionInput"}
	if v.JobDefinition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobDefinition"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeJobsInput(v *DescribeJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJobsInput"}
	if v.Jobs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Jobs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterJobDefinitionInput(v *RegisterJobDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterJobDefinitionInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.NodeProperties != nil {
		if err := validateNodeProperties(v.NodeProperties); err != nil {
			invalidParams.AddNested("NodeProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.ContainerProperties != nil {
		if err := validateContainerProperties(v.ContainerProperties); err != nil {
			invalidParams.AddNested("ContainerProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.RetryStrategy != nil {
		if err := validateRetryStrategy(v.RetryStrategy); err != nil {
			invalidParams.AddNested("RetryStrategy", err.(smithy.InvalidParamsError))
		}
	}
	if v.JobDefinitionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobDefinitionName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSubmitJobInput(v *SubmitJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SubmitJobInput"}
	if v.JobQueue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobQueue"))
	}
	if v.JobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobName"))
	}
	if v.JobDefinition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobDefinition"))
	}
	if v.NodeOverrides != nil {
		if err := validateNodeOverrides(v.NodeOverrides); err != nil {
			invalidParams.AddNested("NodeOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if v.RetryStrategy != nil {
		if err := validateRetryStrategy(v.RetryStrategy); err != nil {
			invalidParams.AddNested("RetryStrategy", err.(smithy.InvalidParamsError))
		}
	}
	if v.ContainerOverrides != nil {
		if err := validateContainerOverrides(v.ContainerOverrides); err != nil {
			invalidParams.AddNested("ContainerOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTerminateJobInput(v *TerminateJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TerminateJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.Reason == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Reason"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateComputeEnvironmentInput(v *UpdateComputeEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateComputeEnvironmentInput"}
	if v.ComputeEnvironment == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComputeEnvironment"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateJobQueueInput(v *UpdateJobQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateJobQueueInput"}
	if v.ComputeEnvironmentOrder != nil {
		if err := validateComputeEnvironmentOrders(v.ComputeEnvironmentOrder); err != nil {
			invalidParams.AddNested("ComputeEnvironmentOrder", err.(smithy.InvalidParamsError))
		}
	}
	if v.JobQueue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobQueue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
