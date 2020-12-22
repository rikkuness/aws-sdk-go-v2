// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewaymanagementapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
)

type awsRestjson1_deserializeOpDeleteConnection struct {
}

func (*awsRestjson1_deserializeOpDeleteConnection) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpDeleteConnection) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorDeleteConnection(response, &metadata)
	}
	output := &DeleteConnectionOutput{}
	out.Result = output

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorDeleteConnection(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ForbiddenException", errorCode):
		return awsRestjson1_deserializeErrorForbiddenException(response, errorBody)

	case strings.EqualFold("GoneException", errorCode):
		return awsRestjson1_deserializeErrorGoneException(response, errorBody)

	case strings.EqualFold("LimitExceededException", errorCode):
		return awsRestjson1_deserializeErrorLimitExceededException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsRestjson1_deserializeOpGetConnection struct {
}

func (*awsRestjson1_deserializeOpGetConnection) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetConnection) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetConnection(response, &metadata)
	}
	output := &GetConnectionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetConnectionOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetConnection(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ForbiddenException", errorCode):
		return awsRestjson1_deserializeErrorForbiddenException(response, errorBody)

	case strings.EqualFold("GoneException", errorCode):
		return awsRestjson1_deserializeErrorGoneException(response, errorBody)

	case strings.EqualFold("LimitExceededException", errorCode):
		return awsRestjson1_deserializeErrorLimitExceededException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetConnectionOutput(v **GetConnectionOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GetConnectionOutput
	if *v == nil {
		sv = &GetConnectionOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "connectedAt":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __timestampIso8601 to be of type string, got %T instead", value)
				}
				t, err := smithytime.ParseDateTime(jtv)
				if err != nil {
					return err
				}
				sv.ConnectedAt = ptr.Time(t)
			}

		case "identity":
			if err := awsRestjson1_deserializeDocumentIdentity(&sv.Identity, value); err != nil {
				return err
			}

		case "lastActiveAt":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __timestampIso8601 to be of type string, got %T instead", value)
				}
				t, err := smithytime.ParseDateTime(jtv)
				if err != nil {
					return err
				}
				sv.LastActiveAt = ptr.Time(t)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpPostToConnection struct {
}

func (*awsRestjson1_deserializeOpPostToConnection) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpPostToConnection) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorPostToConnection(response, &metadata)
	}
	output := &PostToConnectionOutput{}
	out.Result = output

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorPostToConnection(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ForbiddenException", errorCode):
		return awsRestjson1_deserializeErrorForbiddenException(response, errorBody)

	case strings.EqualFold("GoneException", errorCode):
		return awsRestjson1_deserializeErrorGoneException(response, errorBody)

	case strings.EqualFold("LimitExceededException", errorCode):
		return awsRestjson1_deserializeErrorLimitExceededException(response, errorBody)

	case strings.EqualFold("PayloadTooLargeException", errorCode):
		return awsRestjson1_deserializeErrorPayloadTooLargeException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeErrorForbiddenException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ForbiddenException{}
	return output
}

func awsRestjson1_deserializeErrorGoneException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.GoneException{}
	return output
}

func awsRestjson1_deserializeErrorLimitExceededException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.LimitExceededException{}
	return output
}

func awsRestjson1_deserializeErrorPayloadTooLargeException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.PayloadTooLargeException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentPayloadTooLargeException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentForbiddenException(v **types.ForbiddenException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ForbiddenException
	if *v == nil {
		sv = &types.ForbiddenException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentGoneException(v **types.GoneException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.GoneException
	if *v == nil {
		sv = &types.GoneException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentIdentity(v **types.Identity, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Identity
	if *v == nil {
		sv = &types.Identity{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "sourceIp":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.SourceIp = ptr.String(jtv)
			}

		case "userAgent":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.UserAgent = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentLimitExceededException(v **types.LimitExceededException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.LimitExceededException
	if *v == nil {
		sv = &types.LimitExceededException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentPayloadTooLargeException(v **types.PayloadTooLargeException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.PayloadTooLargeException
	if *v == nil {
		sv = &types.PayloadTooLargeException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
