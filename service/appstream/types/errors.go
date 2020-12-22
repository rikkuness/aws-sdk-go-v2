// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An API error occurred. Wait a few minutes and try again.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The image does not support storage connectors.
type IncompatibleImageException struct {
	Message *string
}

func (e *IncompatibleImageException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleImageException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleImageException) ErrorCode() string             { return "IncompatibleImageException" }
func (e *IncompatibleImageException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource cannot be created because your AWS account is suspended. For
// assistance, contact AWS Support.
type InvalidAccountStatusException struct {
	Message *string
}

func (e *InvalidAccountStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAccountStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAccountStatusException) ErrorCode() string             { return "InvalidAccountStatusException" }
func (e *InvalidAccountStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates an incorrect combination of parameters, or a missing parameter.
type InvalidParameterCombinationException struct {
	Message *string
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	return "InvalidParameterCombinationException"
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified role is invalid.
type InvalidRoleException struct {
	Message *string
}

func (e *InvalidRoleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRoleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRoleException) ErrorCode() string             { return "InvalidRoleException" }
func (e *InvalidRoleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested limit exceeds the permitted limit for an account.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The attempted operation is not permitted.
type OperationNotPermittedException struct {
	Message *string
}

func (e *OperationNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationNotPermittedException) ErrorCode() string             { return "OperationNotPermittedException" }
func (e *OperationNotPermittedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// AppStream 2.0 can’t process the request right now because the Describe calls
// from your AWS account are being throttled by Amazon EC2. Try again later.
type RequestLimitExceededException struct {
	Message *string
}

func (e *RequestLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestLimitExceededException) ErrorCode() string             { return "RequestLimitExceededException" }
func (e *RequestLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource is in use.
type ResourceInUseException struct {
	Message *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource exists and is not in use, but isn't available.
type ResourceNotAvailableException struct {
	Message *string
}

func (e *ResourceNotAvailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotAvailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotAvailableException) ErrorCode() string             { return "ResourceNotAvailableException" }
func (e *ResourceNotAvailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource was not found.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
