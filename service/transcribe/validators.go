// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateLanguageModel struct {
}

func (*validateOpCreateLanguageModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLanguageModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLanguageModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLanguageModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMedicalVocabulary struct {
}

func (*validateOpCreateMedicalVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMedicalVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMedicalVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMedicalVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateVocabularyFilter struct {
}

func (*validateOpCreateVocabularyFilter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateVocabularyFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateVocabularyFilterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateVocabularyFilterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateVocabulary struct {
}

func (*validateOpCreateVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLanguageModel struct {
}

func (*validateOpDeleteLanguageModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLanguageModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLanguageModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLanguageModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMedicalTranscriptionJob struct {
}

func (*validateOpDeleteMedicalTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMedicalTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMedicalTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMedicalTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMedicalVocabulary struct {
}

func (*validateOpDeleteMedicalVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMedicalVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMedicalVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMedicalVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTranscriptionJob struct {
}

func (*validateOpDeleteTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteVocabularyFilter struct {
}

func (*validateOpDeleteVocabularyFilter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteVocabularyFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteVocabularyFilterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteVocabularyFilterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteVocabulary struct {
}

func (*validateOpDeleteVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLanguageModel struct {
}

func (*validateOpDescribeLanguageModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLanguageModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLanguageModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLanguageModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMedicalTranscriptionJob struct {
}

func (*validateOpGetMedicalTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMedicalTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMedicalTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMedicalTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMedicalVocabulary struct {
}

func (*validateOpGetMedicalVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMedicalVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMedicalVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMedicalVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTranscriptionJob struct {
}

func (*validateOpGetTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetVocabularyFilter struct {
}

func (*validateOpGetVocabularyFilter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetVocabularyFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetVocabularyFilterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetVocabularyFilterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetVocabulary struct {
}

func (*validateOpGetVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMedicalTranscriptionJob struct {
}

func (*validateOpStartMedicalTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMedicalTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMedicalTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMedicalTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTranscriptionJob struct {
}

func (*validateOpStartTranscriptionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTranscriptionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTranscriptionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTranscriptionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMedicalVocabulary struct {
}

func (*validateOpUpdateMedicalVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMedicalVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMedicalVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMedicalVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateVocabularyFilter struct {
}

func (*validateOpUpdateVocabularyFilter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateVocabularyFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateVocabularyFilterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateVocabularyFilterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateVocabulary struct {
}

func (*validateOpUpdateVocabulary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateVocabulary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateVocabularyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateVocabularyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateLanguageModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLanguageModel{}, middleware.After)
}

func addOpCreateMedicalVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMedicalVocabulary{}, middleware.After)
}

func addOpCreateVocabularyFilterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateVocabularyFilter{}, middleware.After)
}

func addOpCreateVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateVocabulary{}, middleware.After)
}

func addOpDeleteLanguageModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLanguageModel{}, middleware.After)
}

func addOpDeleteMedicalTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMedicalTranscriptionJob{}, middleware.After)
}

func addOpDeleteMedicalVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMedicalVocabulary{}, middleware.After)
}

func addOpDeleteTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTranscriptionJob{}, middleware.After)
}

func addOpDeleteVocabularyFilterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteVocabularyFilter{}, middleware.After)
}

func addOpDeleteVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteVocabulary{}, middleware.After)
}

func addOpDescribeLanguageModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLanguageModel{}, middleware.After)
}

func addOpGetMedicalTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMedicalTranscriptionJob{}, middleware.After)
}

func addOpGetMedicalVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMedicalVocabulary{}, middleware.After)
}

func addOpGetTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTranscriptionJob{}, middleware.After)
}

func addOpGetVocabularyFilterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetVocabularyFilter{}, middleware.After)
}

func addOpGetVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetVocabulary{}, middleware.After)
}

func addOpStartMedicalTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMedicalTranscriptionJob{}, middleware.After)
}

func addOpStartTranscriptionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTranscriptionJob{}, middleware.After)
}

func addOpUpdateMedicalVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMedicalVocabulary{}, middleware.After)
}

func addOpUpdateVocabularyFilterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateVocabularyFilter{}, middleware.After)
}

func addOpUpdateVocabularyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateVocabulary{}, middleware.After)
}

func validateContentRedaction(v *types.ContentRedaction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContentRedaction"}
	if len(v.RedactionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RedactionType"))
	}
	if len(v.RedactionOutput) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RedactionOutput"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInputDataConfig(v *types.InputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputDataConfig"}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLanguageModelInput(v *CreateLanguageModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLanguageModelInput"}
	if v.ModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelName"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if len(v.BaseModelName) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("BaseModelName"))
	}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMedicalVocabularyInput(v *CreateMedicalVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMedicalVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.VocabularyFileUri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyFileUri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateVocabularyFilterInput(v *CreateVocabularyFilterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateVocabularyFilterInput"}
	if v.VocabularyFilterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyFilterName"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateVocabularyInput(v *CreateVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateVocabularyInput"}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLanguageModelInput(v *DeleteLanguageModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLanguageModelInput"}
	if v.ModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMedicalTranscriptionJobInput(v *DeleteMedicalTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMedicalTranscriptionJobInput"}
	if v.MedicalTranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MedicalTranscriptionJobName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMedicalVocabularyInput(v *DeleteMedicalVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMedicalVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTranscriptionJobInput(v *DeleteTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTranscriptionJobInput"}
	if v.TranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TranscriptionJobName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteVocabularyFilterInput(v *DeleteVocabularyFilterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteVocabularyFilterInput"}
	if v.VocabularyFilterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyFilterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteVocabularyInput(v *DeleteVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeLanguageModelInput(v *DescribeLanguageModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLanguageModelInput"}
	if v.ModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMedicalTranscriptionJobInput(v *GetMedicalTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMedicalTranscriptionJobInput"}
	if v.MedicalTranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MedicalTranscriptionJobName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMedicalVocabularyInput(v *GetMedicalVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMedicalVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTranscriptionJobInput(v *GetTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTranscriptionJobInput"}
	if v.TranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TranscriptionJobName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetVocabularyFilterInput(v *GetVocabularyFilterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetVocabularyFilterInput"}
	if v.VocabularyFilterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyFilterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetVocabularyInput(v *GetVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMedicalTranscriptionJobInput(v *StartMedicalTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMedicalTranscriptionJobInput"}
	if v.OutputBucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputBucketName"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.MedicalTranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MedicalTranscriptionJobName"))
	}
	if len(v.Specialty) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Specialty"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Media == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Media"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTranscriptionJobInput(v *StartTranscriptionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTranscriptionJobInput"}
	if v.ContentRedaction != nil {
		if err := validateContentRedaction(v.ContentRedaction); err != nil {
			invalidParams.AddNested("ContentRedaction", err.(smithy.InvalidParamsError))
		}
	}
	if v.Media == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Media"))
	}
	if v.TranscriptionJobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TranscriptionJobName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMedicalVocabularyInput(v *UpdateMedicalVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMedicalVocabularyInput"}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateVocabularyFilterInput(v *UpdateVocabularyFilterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateVocabularyFilterInput"}
	if v.VocabularyFilterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyFilterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateVocabularyInput(v *UpdateVocabularyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateVocabularyInput"}
	if v.VocabularyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VocabularyName"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
