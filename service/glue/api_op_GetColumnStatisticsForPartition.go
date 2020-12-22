// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves partition statistics of columns. The Identity and Access Management
// (IAM) permission required for this operation is GetPartition.
func (c *Client) GetColumnStatisticsForPartition(ctx context.Context, params *GetColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*GetColumnStatisticsForPartitionOutput, error) {
	if params == nil {
		params = &GetColumnStatisticsForPartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetColumnStatisticsForPartition", params, optFns, addOperationGetColumnStatisticsForPartitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetColumnStatisticsForPartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetColumnStatisticsForPartitionInput struct {

	// A list of the column names.
	//
	// This member is required.
	ColumnNames []string

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// A list of partition values identifying the partition.
	//
	// This member is required.
	PartitionValues []string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
}

type GetColumnStatisticsForPartitionOutput struct {

	// List of ColumnStatistics that failed to be retrieved.
	ColumnStatisticsList []types.ColumnStatistics

	// Error occurred during retrieving column statistics data.
	Errors []types.ColumnError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetColumnStatisticsForPartitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetColumnStatisticsForPartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetColumnStatisticsForPartition{}, middleware.After)
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
	if err = addOpGetColumnStatisticsForPartitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetColumnStatisticsForPartition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetColumnStatisticsForPartition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetColumnStatisticsForPartition",
	}
}
