// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
)

type awsAwsjson11_deserializeOpDescribeDimensionKeys struct {
}

func (*awsAwsjson11_deserializeOpDescribeDimensionKeys) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDescribeDimensionKeys) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeDimensionKeys(response, &metadata)
	}
	output := &DescribeDimensionKeysOutput{}
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

	err = awsAwsjson11_deserializeOpDocumentDescribeDimensionKeysOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorDescribeDimensionKeys(response *smithyhttp.Response, metadata *middleware.Metadata) error {
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
	case strings.EqualFold("InternalServiceError", errorCode):
		return awsAwsjson11_deserializeErrorInternalServiceError(response, errorBody)

	case strings.EqualFold("InvalidArgumentException", errorCode):
		return awsAwsjson11_deserializeErrorInvalidArgumentException(response, errorBody)

	case strings.EqualFold("NotAuthorizedException", errorCode):
		return awsAwsjson11_deserializeErrorNotAuthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpGetResourceMetrics struct {
}

func (*awsAwsjson11_deserializeOpGetResourceMetrics) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpGetResourceMetrics) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorGetResourceMetrics(response, &metadata)
	}
	output := &GetResourceMetricsOutput{}
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

	err = awsAwsjson11_deserializeOpDocumentGetResourceMetricsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson11_deserializeOpErrorGetResourceMetrics(response *smithyhttp.Response, metadata *middleware.Metadata) error {
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
	case strings.EqualFold("InternalServiceError", errorCode):
		return awsAwsjson11_deserializeErrorInternalServiceError(response, errorBody)

	case strings.EqualFold("InvalidArgumentException", errorCode):
		return awsAwsjson11_deserializeErrorInvalidArgumentException(response, errorBody)

	case strings.EqualFold("NotAuthorizedException", errorCode):
		return awsAwsjson11_deserializeErrorNotAuthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson11_deserializeErrorInternalServiceError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.InternalServiceError{}
	err := awsAwsjson11_deserializeDocumentInternalServiceError(&output, shape)

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

func awsAwsjson11_deserializeErrorInvalidArgumentException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.InvalidArgumentException{}
	err := awsAwsjson11_deserializeDocumentInvalidArgumentException(&output, shape)

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

func awsAwsjson11_deserializeErrorNotAuthorizedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
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

	output := &types.NotAuthorizedException{}
	err := awsAwsjson11_deserializeDocumentNotAuthorizedException(&output, shape)

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

func awsAwsjson11_deserializeDocumentDataPoint(v **types.DataPoint, value interface{}) error {
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

	var sv *types.DataPoint
	if *v == nil {
		sv = &types.DataPoint{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Timestamp":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Timestamp = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Value":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Double to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Value = ptr.Float64(f64)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentDataPointsList(v *[]types.DataPoint, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.DataPoint
	if *v == nil {
		cv = []types.DataPoint{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DataPoint
		destAddr := &col
		if err := awsAwsjson11_deserializeDocumentDataPoint(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionKeyDescription(v **types.DimensionKeyDescription, value interface{}) error {
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

	var sv *types.DimensionKeyDescription
	if *v == nil {
		sv = &types.DimensionKeyDescription{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, value); err != nil {
				return err
			}

		case "Partitions":
			if err := awsAwsjson11_deserializeDocumentMetricValuesList(&sv.Partitions, value); err != nil {
				return err
			}

		case "Total":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Double to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.Total = ptr.Float64(f64)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionKeyDescriptionList(v *[]types.DimensionKeyDescription, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.DimensionKeyDescription
	if *v == nil {
		cv = []types.DimensionKeyDescription{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DimensionKeyDescription
		destAddr := &col
		if err := awsAwsjson11_deserializeDocumentDimensionKeyDescription(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentDimensionMap(v *map[string]string, value interface{}) error {
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

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsAwsjson11_deserializeDocumentInternalServiceError(v **types.InternalServiceError, value interface{}) error {
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

	var sv *types.InternalServiceError
	if *v == nil {
		sv = &types.InternalServiceError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
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

func awsAwsjson11_deserializeDocumentInvalidArgumentException(v **types.InvalidArgumentException, value interface{}) error {
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

	var sv *types.InvalidArgumentException
	if *v == nil {
		sv = &types.InvalidArgumentException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
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

func awsAwsjson11_deserializeDocumentMetricKeyDataPoints(v **types.MetricKeyDataPoints, value interface{}) error {
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

	var sv *types.MetricKeyDataPoints
	if *v == nil {
		sv = &types.MetricKeyDataPoints{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "DataPoints":
			if err := awsAwsjson11_deserializeDocumentDataPointsList(&sv.DataPoints, value); err != nil {
				return err
			}

		case "Key":
			if err := awsAwsjson11_deserializeDocumentResponseResourceMetricKey(&sv.Key, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentMetricKeyDataPointsList(v *[]types.MetricKeyDataPoints, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.MetricKeyDataPoints
	if *v == nil {
		cv = []types.MetricKeyDataPoints{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.MetricKeyDataPoints
		destAddr := &col
		if err := awsAwsjson11_deserializeDocumentMetricKeyDataPoints(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentMetricValuesList(v *[]float64, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []float64
	if *v == nil {
		cv = []float64{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col float64
		if value != nil {
			jtv, ok := value.(json.Number)
			if !ok {
				return fmt.Errorf("expected Double to be json.Number, got %T instead", value)
			}
			f64, err := jtv.Float64()
			if err != nil {
				return err
			}
			col = f64
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentNotAuthorizedException(v **types.NotAuthorizedException, value interface{}) error {
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

	var sv *types.NotAuthorizedException
	if *v == nil {
		sv = &types.NotAuthorizedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
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

func awsAwsjson11_deserializeDocumentResponsePartitionKey(v **types.ResponsePartitionKey, value interface{}) error {
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

	var sv *types.ResponsePartitionKey
	if *v == nil {
		sv = &types.ResponsePartitionKey{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeDocumentResponsePartitionKeyList(v *[]types.ResponsePartitionKey, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.ResponsePartitionKey
	if *v == nil {
		cv = []types.ResponsePartitionKey{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.ResponsePartitionKey
		destAddr := &col
		if err := awsAwsjson11_deserializeDocumentResponsePartitionKey(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson11_deserializeDocumentResponseResourceMetricKey(v **types.ResponseResourceMetricKey, value interface{}) error {
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

	var sv *types.ResponseResourceMetricKey
	if *v == nil {
		sv = &types.ResponseResourceMetricKey{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Dimensions":
			if err := awsAwsjson11_deserializeDocumentDimensionMap(&sv.Dimensions, value); err != nil {
				return err
			}

		case "Metric":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Metric = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentDescribeDimensionKeysOutput(v **DescribeDimensionKeysOutput, value interface{}) error {
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

	var sv *DescribeDimensionKeysOutput
	if *v == nil {
		sv = &DescribeDimensionKeysOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AlignedEndTime":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedEndTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "AlignedStartTime":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedStartTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Keys":
			if err := awsAwsjson11_deserializeDocumentDimensionKeyDescriptionList(&sv.Keys, value); err != nil {
				return err
			}

		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		case "PartitionKeys":
			if err := awsAwsjson11_deserializeDocumentResponsePartitionKeyList(&sv.PartitionKeys, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson11_deserializeOpDocumentGetResourceMetricsOutput(v **GetResourceMetricsOutput, value interface{}) error {
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

	var sv *GetResourceMetricsOutput
	if *v == nil {
		sv = &GetResourceMetricsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AlignedEndTime":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedEndTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "AlignedStartTime":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected ISOTimestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				sv.AlignedStartTime = ptr.Time(smithytime.ParseEpochSeconds(f64))
			}

		case "Identifier":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Identifier = ptr.String(jtv)
			}

		case "MetricList":
			if err := awsAwsjson11_deserializeDocumentMetricKeyDataPointsList(&sv.MetricList, value); err != nil {
				return err
			}

		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
