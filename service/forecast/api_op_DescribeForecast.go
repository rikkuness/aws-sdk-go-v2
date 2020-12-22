// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a forecast created using the CreateForecast operation. In addition to
// listing the properties provided in the CreateForecast request, this operation
// lists the following properties:
//
// * DatasetGroupArn - The dataset group that
// provided the training data.
//
// * CreationTime
//
// * LastModificationTime
//
// * Status
//
// *
// Message - If an error occurred, information about the error.
func (c *Client) DescribeForecast(ctx context.Context, params *DescribeForecastInput, optFns ...func(*Options)) (*DescribeForecastOutput, error) {
	if params == nil {
		params = &DescribeForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeForecast", params, optFns, addOperationDescribeForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeForecastInput struct {

	// The Amazon Resource Name (ARN) of the forecast.
	//
	// This member is required.
	ForecastArn *string
}

type DescribeForecastOutput struct {

	// When the forecast creation task was created.
	CreationTime *time.Time

	// The ARN of the dataset group that provided the data used to train the predictor.
	DatasetGroupArn *string

	// The forecast ARN as specified in the request.
	ForecastArn *string

	// The name of the forecast.
	ForecastName *string

	// The quantiles at which probabilistic forecasts were generated.
	ForecastTypes []string

	// Initially, the same as CreationTime (status is CREATE_PENDING). Updated when
	// inference (creating the forecast) starts (status changed to CREATE_IN_PROGRESS),
	// and when inference is complete (status changed to ACTIVE) or fails (status
	// changed to CREATE_FAILED).
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The ARN of the predictor used to generate the forecast.
	PredictorArn *string

	// The status of the forecast. States include:
	//
	// * ACTIVE
	//
	// * CREATE_PENDING,
	// CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * DELETE_PENDING, DELETE_IN_PROGRESS,
	// DELETE_FAILED
	//
	// The Status of the forecast must be ACTIVE before you can query or
	// export the forecast.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeForecast{}, middleware.After)
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
	if err = addOpDescribeForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeForecast",
	}
}
