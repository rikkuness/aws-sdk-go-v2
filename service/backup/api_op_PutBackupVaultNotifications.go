// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Turns on notifications on a backup vault for the specified topic and events.
func (c *Client) PutBackupVaultNotifications(ctx context.Context, params *PutBackupVaultNotificationsInput, optFns ...func(*Options)) (*PutBackupVaultNotificationsOutput, error) {
	if params == nil {
		params = &PutBackupVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBackupVaultNotifications", params, optFns, addOperationPutBackupVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBackupVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupVaultNotificationsInput struct {

	// An array of events that indicate the status of jobs to back up resources to the
	// backup vault.
	//
	// This member is required.
	BackupVaultEvents []types.BackupVaultEvent

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s
	// events; for example, arn:aws:sns:us-west-2:111122223333:MyVaultTopic.
	//
	// This member is required.
	SNSTopicArn *string
}

type PutBackupVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBackupVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupVaultNotifications{}, middleware.After)
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
	if err = addOpPutBackupVaultNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupVaultNotifications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBackupVaultNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "PutBackupVaultNotifications",
	}
}
