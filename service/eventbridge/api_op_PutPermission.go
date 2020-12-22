// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Running PutPermission permits the specified AWS account or AWS organization to
// put events to the specified event bus. Amazon EventBridge (CloudWatch Events)
// rules in your account are triggered by these events arriving to an event bus in
// your account. For another account to send events to your account, that external
// account must have an EventBridge rule with your account's event bus as a target.
// To enable multiple AWS accounts to put events to your event bus, run
// PutPermission once for each of these accounts. Or, if all the accounts are
// members of the same AWS organization, you can run PutPermission once specifying
// Principal as "*" and specifying the AWS organization ID in Condition, to grant
// permissions to all accounts in that organization. If you grant permissions using
// an organization, then accounts in that organization must specify a RoleArn with
// proper permissions when they use PutTarget to add your account's event bus as a
// target. For more information, see Sending and Receiving Events Between AWS
// Accounts
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide. The permission policy on the default event
// bus cannot exceed 10 KB in size.
func (c *Client) PutPermission(ctx context.Context, params *PutPermissionInput, optFns ...func(*Options)) (*PutPermissionOutput, error) {
	if params == nil {
		params = &PutPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPermission", params, optFns, addOperationPutPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionInput struct {

	// The action that you are enabling the other account to perform. Currently, this
	// must be events:PutEvents.
	Action *string

	// This parameter enables you to limit the permission to accounts that fulfill a
	// certain condition, such as being a member of a certain AWS organization. For
	// more information about AWS Organizations, see What Is AWS Organizations
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html)
	// in the AWS Organizations User Guide. If you specify Condition with an AWS
	// organization ID, and specify "*" as the value for Principal, you grant
	// permission to all the accounts in the named organization. The Condition is a
	// JSON string which must contain Type, Key, and Value fields.
	Condition *types.Condition

	// The name of the event bus associated with the rule. If you omit this, the
	// default event bus is used.
	EventBusName *string

	// A JSON string that describes the permission policy statement. You can include a
	// Policy parameter in the request instead of using the StatementId, Action,
	// Principal, or Condition parameters.
	Policy *string

	// The 12-digit AWS account ID that you are permitting to put events to your
	// default event bus. Specify "*" to permit any account to put events to your
	// default event bus. If you specify "*" without specifying Condition, avoid
	// creating rules that may match undesirable events. To create more secure rules,
	// make sure that the event pattern for each rule contains an account field with a
	// specific account ID from which to receive events. Rules with an account field do
	// not match any events sent from other accounts.
	Principal *string

	// An identifier string for the external account that you are granting permissions
	// to. If you later want to revoke the permission for this external account,
	// specify this StatementId when you run RemovePermission.
	StatementId *string
}

type PutPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermission{}, middleware.After)
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
	if err = addOpPutPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutPermission",
	}
}
