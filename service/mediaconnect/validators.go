// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddFlowOutputs struct {
}

func (*validateOpAddFlowOutputs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddFlowOutputs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddFlowOutputsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddFlowOutputsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAddFlowSources struct {
}

func (*validateOpAddFlowSources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddFlowSources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddFlowSourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddFlowSourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAddFlowVpcInterfaces struct {
}

func (*validateOpAddFlowVpcInterfaces) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddFlowVpcInterfaces) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddFlowVpcInterfacesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddFlowVpcInterfacesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateFlow struct {
}

func (*validateOpCreateFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateFlowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFlow struct {
}

func (*validateOpDeleteFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFlowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFlow struct {
}

func (*validateOpDescribeFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFlowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeOffering struct {
}

func (*validateOpDescribeOffering) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeOffering) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeOfferingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeOfferingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeReservation struct {
}

func (*validateOpDescribeReservation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeReservation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeReservationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeReservationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGrantFlowEntitlements struct {
}

func (*validateOpGrantFlowEntitlements) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGrantFlowEntitlements) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GrantFlowEntitlementsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGrantFlowEntitlementsInput(input); err != nil {
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

type validateOpPurchaseOffering struct {
}

func (*validateOpPurchaseOffering) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPurchaseOffering) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PurchaseOfferingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPurchaseOfferingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveFlowOutput struct {
}

func (*validateOpRemoveFlowOutput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveFlowOutput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveFlowOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveFlowOutputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveFlowSource struct {
}

func (*validateOpRemoveFlowSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveFlowSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveFlowSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveFlowSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveFlowVpcInterface struct {
}

func (*validateOpRemoveFlowVpcInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveFlowVpcInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveFlowVpcInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveFlowVpcInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRevokeFlowEntitlement struct {
}

func (*validateOpRevokeFlowEntitlement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRevokeFlowEntitlement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RevokeFlowEntitlementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRevokeFlowEntitlementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartFlow struct {
}

func (*validateOpStartFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartFlowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopFlow struct {
}

func (*validateOpStopFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopFlowInput(input); err != nil {
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

type validateOpUpdateFlowEntitlement struct {
}

func (*validateOpUpdateFlowEntitlement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFlowEntitlement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFlowEntitlementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFlowEntitlementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateFlow struct {
}

func (*validateOpUpdateFlow) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFlowInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateFlowOutput struct {
}

func (*validateOpUpdateFlowOutput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFlowOutput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFlowOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFlowOutputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateFlowSource struct {
}

func (*validateOpUpdateFlowSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFlowSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFlowSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFlowSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddFlowOutputsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddFlowOutputs{}, middleware.After)
}

func addOpAddFlowSourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddFlowSources{}, middleware.After)
}

func addOpAddFlowVpcInterfacesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddFlowVpcInterfaces{}, middleware.After)
}

func addOpCreateFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateFlow{}, middleware.After)
}

func addOpDeleteFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFlow{}, middleware.After)
}

func addOpDescribeFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFlow{}, middleware.After)
}

func addOpDescribeOfferingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeOffering{}, middleware.After)
}

func addOpDescribeReservationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeReservation{}, middleware.After)
}

func addOpGrantFlowEntitlementsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGrantFlowEntitlements{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPurchaseOfferingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPurchaseOffering{}, middleware.After)
}

func addOpRemoveFlowOutputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveFlowOutput{}, middleware.After)
}

func addOpRemoveFlowSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveFlowSource{}, middleware.After)
}

func addOpRemoveFlowVpcInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveFlowVpcInterface{}, middleware.After)
}

func addOpRevokeFlowEntitlementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRevokeFlowEntitlement{}, middleware.After)
}

func addOpStartFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartFlow{}, middleware.After)
}

func addOpStopFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopFlow{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateFlowEntitlementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFlowEntitlement{}, middleware.After)
}

func addOpUpdateFlowValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFlow{}, middleware.After)
}

func addOpUpdateFlowOutputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFlowOutput{}, middleware.After)
}

func addOpUpdateFlowSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFlowSource{}, middleware.After)
}

func validate__listOfAddOutputRequest(v []types.AddOutputRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfAddOutputRequest"}
	for i := range v {
		if err := validateAddOutputRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validate__listOfGrantEntitlementRequest(v []types.GrantEntitlementRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfGrantEntitlementRequest"}
	for i := range v {
		if err := validateGrantEntitlementRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validate__listOfSetSourceRequest(v []types.SetSourceRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfSetSourceRequest"}
	for i := range v {
		if err := validateSetSourceRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validate__listOfVpcInterfaceRequest(v []types.VpcInterfaceRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfVpcInterfaceRequest"}
	for i := range v {
		if err := validateVpcInterfaceRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAddOutputRequest(v *types.AddOutputRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddOutputRequest"}
	if len(v.Protocol) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Protocol"))
	}
	if v.Encryption != nil {
		if err := validateEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEncryption(v *types.Encryption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Encryption"}
	if len(v.Algorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Algorithm"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGrantEntitlementRequest(v *types.GrantEntitlementRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GrantEntitlementRequest"}
	if v.Encryption != nil {
		if err := validateEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if v.Subscribers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subscribers"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSetSourceRequest(v *types.SetSourceRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetSourceRequest"}
	if v.Decryption != nil {
		if err := validateEncryption(v.Decryption); err != nil {
			invalidParams.AddNested("Decryption", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVpcInterfaceRequest(v *types.VpcInterfaceRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VpcInterfaceRequest"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.SubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddFlowOutputsInput(v *AddFlowOutputsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddFlowOutputsInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.Outputs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Outputs"))
	} else if v.Outputs != nil {
		if err := validate__listOfAddOutputRequest(v.Outputs); err != nil {
			invalidParams.AddNested("Outputs", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddFlowSourcesInput(v *AddFlowSourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddFlowSourcesInput"}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	} else if v.Sources != nil {
		if err := validate__listOfSetSourceRequest(v.Sources); err != nil {
			invalidParams.AddNested("Sources", err.(smithy.InvalidParamsError))
		}
	}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddFlowVpcInterfacesInput(v *AddFlowVpcInterfacesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddFlowVpcInterfacesInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.VpcInterfaces == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcInterfaces"))
	} else if v.VpcInterfaces != nil {
		if err := validate__listOfVpcInterfaceRequest(v.VpcInterfaces); err != nil {
			invalidParams.AddNested("VpcInterfaces", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateFlowInput(v *CreateFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFlowInput"}
	if v.Sources != nil {
		if err := validate__listOfSetSourceRequest(v.Sources); err != nil {
			invalidParams.AddNested("Sources", err.(smithy.InvalidParamsError))
		}
	}
	if v.Source != nil {
		if err := validateSetSourceRequest(v.Source); err != nil {
			invalidParams.AddNested("Source", err.(smithy.InvalidParamsError))
		}
	}
	if v.Outputs != nil {
		if err := validate__listOfAddOutputRequest(v.Outputs); err != nil {
			invalidParams.AddNested("Outputs", err.(smithy.InvalidParamsError))
		}
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Entitlements != nil {
		if err := validate__listOfGrantEntitlementRequest(v.Entitlements); err != nil {
			invalidParams.AddNested("Entitlements", err.(smithy.InvalidParamsError))
		}
	}
	if v.VpcInterfaces != nil {
		if err := validate__listOfVpcInterfaceRequest(v.VpcInterfaces); err != nil {
			invalidParams.AddNested("VpcInterfaces", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFlowInput(v *DeleteFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFlowInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFlowInput(v *DescribeFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFlowInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeOfferingInput(v *DescribeOfferingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeOfferingInput"}
	if v.OfferingArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OfferingArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeReservationInput(v *DescribeReservationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeReservationInput"}
	if v.ReservationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReservationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGrantFlowEntitlementsInput(v *GrantFlowEntitlementsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GrantFlowEntitlementsInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.Entitlements == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Entitlements"))
	} else if v.Entitlements != nil {
		if err := validate__listOfGrantEntitlementRequest(v.Entitlements); err != nil {
			invalidParams.AddNested("Entitlements", err.(smithy.InvalidParamsError))
		}
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

func validateOpPurchaseOfferingInput(v *PurchaseOfferingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PurchaseOfferingInput"}
	if v.OfferingArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OfferingArn"))
	}
	if v.Start == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Start"))
	}
	if v.ReservationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReservationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveFlowOutputInput(v *RemoveFlowOutputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveFlowOutputInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.OutputArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveFlowSourceInput(v *RemoveFlowSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveFlowSourceInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveFlowVpcInterfaceInput(v *RemoveFlowVpcInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveFlowVpcInterfaceInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.VpcInterfaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcInterfaceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRevokeFlowEntitlementInput(v *RevokeFlowEntitlementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RevokeFlowEntitlementInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if v.EntitlementArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EntitlementArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartFlowInput(v *StartFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartFlowInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopFlowInput(v *StopFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopFlowInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
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
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
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

func validateOpUpdateFlowEntitlementInput(v *UpdateFlowEntitlementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFlowEntitlementInput"}
	if v.EntitlementArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EntitlementArn"))
	}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateFlowInput(v *UpdateFlowInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFlowInput"}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateFlowOutputInput(v *UpdateFlowOutputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFlowOutputInput"}
	if v.OutputArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputArn"))
	}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateFlowSourceInput(v *UpdateFlowSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFlowSourceInput"}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if v.FlowArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
