// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This exception is thrown when trusted access has not been enabled between AWS
// CloudTrail and AWS Organizations. For more information, see Enabling Trusted
// Access with Other AWS Services
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// and Prepare For Creating a Trail For Your Organization
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html).
type CloudTrailAccessNotEnabledException struct {
	Message *string
}

func (e *CloudTrailAccessNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudTrailAccessNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudTrailAccessNotEnabledException) ErrorCode() string {
	return "CloudTrailAccessNotEnabledException"
}
func (e *CloudTrailAccessNotEnabledException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when an operation is called with an invalid trail ARN.
// The format of a trail ARN is:
// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
type CloudTrailARNInvalidException struct {
	Message *string
}

func (e *CloudTrailARNInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudTrailARNInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudTrailARNInvalidException) ErrorCode() string             { return "CloudTrailARNInvalidException" }
func (e *CloudTrailARNInvalidException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Cannot set a CloudWatch Logs delivery for this region.
type CloudWatchLogsDeliveryUnavailableException struct {
	Message *string
}

func (e *CloudWatchLogsDeliveryUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudWatchLogsDeliveryUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudWatchLogsDeliveryUnavailableException) ErrorCode() string {
	return "CloudWatchLogsDeliveryUnavailableException"
}
func (e *CloudWatchLogsDeliveryUnavailableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// If you run GetInsightSelectors on a trail that does not have Insights events
// enabled, the operation throws the exception InsightNotEnabledException.
type InsightNotEnabledException struct {
	Message *string
}

func (e *InsightNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsightNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsightNotEnabledException) ErrorCode() string             { return "InsightNotEnabledException" }
func (e *InsightNotEnabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the IAM user or role that is used to create the
// organization trail is lacking one or more required permissions for creating an
// organization trail in a required service. For more information, see Prepare For
// Creating a Trail For Your Organization
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html).
type InsufficientDependencyServiceAccessPermissionException struct {
	Message *string
}

func (e *InsufficientDependencyServiceAccessPermissionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDependencyServiceAccessPermissionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDependencyServiceAccessPermissionException) ErrorCode() string {
	return "InsufficientDependencyServiceAccessPermissionException"
}
func (e *InsufficientDependencyServiceAccessPermissionException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the policy on the S3 bucket or KMS key is not
// sufficient.
type InsufficientEncryptionPolicyException struct {
	Message *string
}

func (e *InsufficientEncryptionPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientEncryptionPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientEncryptionPolicyException) ErrorCode() string {
	return "InsufficientEncryptionPolicyException"
}
func (e *InsufficientEncryptionPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the policy on the S3 bucket is not sufficient.
type InsufficientS3BucketPolicyException struct {
	Message *string
}

func (e *InsufficientS3BucketPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientS3BucketPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientS3BucketPolicyException) ErrorCode() string {
	return "InsufficientS3BucketPolicyException"
}
func (e *InsufficientS3BucketPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the policy on the SNS topic is not sufficient.
type InsufficientSnsTopicPolicyException struct {
	Message *string
}

func (e *InsufficientSnsTopicPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientSnsTopicPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientSnsTopicPolicyException) ErrorCode() string {
	return "InsufficientSnsTopicPolicyException"
}
func (e *InsufficientSnsTopicPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the provided CloudWatch log group is not valid.
type InvalidCloudWatchLogsLogGroupArnException struct {
	Message *string
}

func (e *InvalidCloudWatchLogsLogGroupArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCloudWatchLogsLogGroupArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCloudWatchLogsLogGroupArnException) ErrorCode() string {
	return "InvalidCloudWatchLogsLogGroupArnException"
}
func (e *InvalidCloudWatchLogsLogGroupArnException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the provided role is not valid.
type InvalidCloudWatchLogsRoleArnException struct {
	Message *string
}

func (e *InvalidCloudWatchLogsRoleArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCloudWatchLogsRoleArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCloudWatchLogsRoleArnException) ErrorCode() string {
	return "InvalidCloudWatchLogsRoleArnException"
}
func (e *InvalidCloudWatchLogsRoleArnException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Occurs if an event category that is not valid is specified as a value of
// EventCategory.
type InvalidEventCategoryException struct {
	Message *string
}

func (e *InvalidEventCategoryException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEventCategoryException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEventCategoryException) ErrorCode() string             { return "InvalidEventCategoryException" }
func (e *InvalidEventCategoryException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the PutEventSelectors operation is called with a
// number of event selectors or data resources that is not valid. The combination
// of event selectors and data resources is not valid. A trail can have up to 5
// event selectors. A trail is limited to 250 data resources. These data resources
// can be distributed across event selectors, but the overall total cannot exceed
// 250. You can:
//
// * Specify a valid number of event selectors (1 to 5) for a
// trail.
//
// * Specify a valid number of data resources (1 to 250) for an event
// selector. The limit of number of resources on an individual event selector is
// configurable up to 250. However, this upper limit is allowed only if the total
// number of data resources does not exceed 250 across all event selectors for a
// trail.
//
// * Specify a valid value for a parameter. For example, specifying the
// ReadWriteType parameter with a value of read-only is invalid.
type InvalidEventSelectorsException struct {
	Message *string
}

func (e *InvalidEventSelectorsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEventSelectorsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEventSelectorsException) ErrorCode() string             { return "InvalidEventSelectorsException" }
func (e *InvalidEventSelectorsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when an operation is called on a trail from a region
// other than the region in which the trail was created.
type InvalidHomeRegionException struct {
	Message *string
}

func (e *InvalidHomeRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidHomeRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidHomeRegionException) ErrorCode() string             { return "InvalidHomeRegionException" }
func (e *InvalidHomeRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The formatting or syntax of the InsightSelectors JSON statement in your
// PutInsightSelectors or GetInsightSelectors request is not valid, or the
// specified insight type in the InsightSelectors statement is not a valid insight
// type.
type InvalidInsightSelectorsException struct {
	Message *string
}

func (e *InvalidInsightSelectorsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInsightSelectorsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInsightSelectorsException) ErrorCode() string {
	return "InvalidInsightSelectorsException"
}
func (e *InvalidInsightSelectorsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the KMS key ARN is invalid.
type InvalidKmsKeyIdException struct {
	Message *string
}

func (e *InvalidKmsKeyIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidKmsKeyIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidKmsKeyIdException) ErrorCode() string             { return "InvalidKmsKeyIdException" }
func (e *InvalidKmsKeyIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs when an invalid lookup attribute is specified.
type InvalidLookupAttributesException struct {
	Message *string
}

func (e *InvalidLookupAttributesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLookupAttributesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLookupAttributesException) ErrorCode() string {
	return "InvalidLookupAttributesException"
}
func (e *InvalidLookupAttributesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown if the limit specified is invalid.
type InvalidMaxResultsException struct {
	Message *string
}

func (e *InvalidMaxResultsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidMaxResultsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidMaxResultsException) ErrorCode() string             { return "InvalidMaxResultsException" }
func (e *InvalidMaxResultsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Invalid token or token that was previously used in a request with different
// parameters. This exception is thrown if the token is invalid.
type InvalidNextTokenException struct {
	Message *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the combination of parameters provided is not
// valid.
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

// This exception is thrown when the provided S3 bucket name is not valid.
type InvalidS3BucketNameException struct {
	Message *string
}

func (e *InvalidS3BucketNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3BucketNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3BucketNameException) ErrorCode() string             { return "InvalidS3BucketNameException" }
func (e *InvalidS3BucketNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the provided S3 prefix is not valid.
type InvalidS3PrefixException struct {
	Message *string
}

func (e *InvalidS3PrefixException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3PrefixException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3PrefixException) ErrorCode() string             { return "InvalidS3PrefixException" }
func (e *InvalidS3PrefixException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the provided SNS topic name is not valid.
type InvalidSnsTopicNameException struct {
	Message *string
}

func (e *InvalidSnsTopicNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnsTopicNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnsTopicNameException) ErrorCode() string             { return "InvalidSnsTopicNameException" }
func (e *InvalidSnsTopicNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified tag key or values are not valid. It
// can also occur if there are duplicate tags or too many tags on the resource.
type InvalidTagParameterException struct {
	Message *string
}

func (e *InvalidTagParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagParameterException) ErrorCode() string             { return "InvalidTagParameterException" }
func (e *InvalidTagParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Occurs if the timestamp values are invalid. Either the start time occurs after
// the end time or the time range is outside the range of possible values.
type InvalidTimeRangeException struct {
	Message *string
}

func (e *InvalidTimeRangeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTimeRangeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTimeRangeException) ErrorCode() string             { return "InvalidTimeRangeException" }
func (e *InvalidTimeRangeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Reserved for future use.
type InvalidTokenException struct {
	Message *string
}

func (e *InvalidTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTokenException) ErrorCode() string             { return "InvalidTokenException" }
func (e *InvalidTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the provided trail name is not valid. Trail names
// must meet the following requirements:
//
// * Contain only ASCII letters (a-z, A-Z),
// numbers (0-9), periods (.), underscores (_), or dashes (-)
//
// * Start with a
// letter or number, and end with a letter or number
//
// * Be between 3 and 128
// characters
//
// * Have no adjacent periods, underscores or dashes. Names like
// my-_namespace and my--namespace are invalid.
//
// * Not be in IP address format (for
// example, 192.168.5.4)
type InvalidTrailNameException struct {
	Message *string
}

func (e *InvalidTrailNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTrailNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTrailNameException) ErrorCode() string             { return "InvalidTrailNameException" }
func (e *InvalidTrailNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when there is an issue with the specified KMS key and
// the trail can’t be updated.
type KmsException struct {
	Message *string
}

func (e *KmsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsException) ErrorCode() string             { return "KmsException" }
func (e *KmsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is no longer in use.
type KmsKeyDisabledException struct {
	Message *string
}

func (e *KmsKeyDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsKeyDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsKeyDisabledException) ErrorCode() string             { return "KmsKeyDisabledException" }
func (e *KmsKeyDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the KMS key does not exist, or when the S3 bucket
// and the KMS key are not in the same region.
type KmsKeyNotFoundException struct {
	Message *string
}

func (e *KmsKeyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsKeyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsKeyNotFoundException) ErrorCode() string             { return "KmsKeyNotFoundException" }
func (e *KmsKeyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the maximum number of trails is reached.
type MaximumNumberOfTrailsExceededException struct {
	Message *string
}

func (e *MaximumNumberOfTrailsExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MaximumNumberOfTrailsExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MaximumNumberOfTrailsExceededException) ErrorCode() string {
	return "MaximumNumberOfTrailsExceededException"
}
func (e *MaximumNumberOfTrailsExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the AWS account making the request to create or
// update an organization trail is not the master account for an organization in
// AWS Organizations. For more information, see Prepare For Creating a Trail For
// Your Organization
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html).
type NotOrganizationMasterAccountException struct {
	Message *string
}

func (e *NotOrganizationMasterAccountException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotOrganizationMasterAccountException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotOrganizationMasterAccountException) ErrorCode() string {
	return "NotOrganizationMasterAccountException"
}
func (e *NotOrganizationMasterAccountException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the requested operation is not permitted.
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

// This exception is thrown when AWS Organizations is not configured to support all
// features. All features must be enabled in AWS Organization to support creating
// an organization trail. For more information, see Prepare For Creating a Trail
// For Your Organization
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html).
type OrganizationNotInAllFeaturesModeException struct {
	Message *string
}

func (e *OrganizationNotInAllFeaturesModeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OrganizationNotInAllFeaturesModeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OrganizationNotInAllFeaturesModeException) ErrorCode() string {
	return "OrganizationNotInAllFeaturesModeException"
}
func (e *OrganizationNotInAllFeaturesModeException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This exception is thrown when the request is made from an AWS account that is
// not a member of an organization. To make this request, sign in using the
// credentials of an account that belongs to an organization.
type OrganizationsNotInUseException struct {
	Message *string
}

func (e *OrganizationsNotInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OrganizationsNotInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OrganizationsNotInUseException) ErrorCode() string             { return "OrganizationsNotInUseException" }
func (e *OrganizationsNotInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified resource is not found.
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

// This exception is thrown when the specified resource type is not supported by
// CloudTrail.
type ResourceTypeNotSupportedException struct {
	Message *string
}

func (e *ResourceTypeNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceTypeNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceTypeNotSupportedException) ErrorCode() string {
	return "ResourceTypeNotSupportedException"
}
func (e *ResourceTypeNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified S3 bucket does not exist.
type S3BucketDoesNotExistException struct {
	Message *string
}

func (e *S3BucketDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3BucketDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3BucketDoesNotExistException) ErrorCode() string             { return "S3BucketDoesNotExistException" }
func (e *S3BucketDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of tags per trail has exceeded the permitted amount. Currently, the
// limit is 50.
type TagsLimitExceededException struct {
	Message *string
}

func (e *TagsLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagsLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagsLimitExceededException) ErrorCode() string             { return "TagsLimitExceededException" }
func (e *TagsLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the specified trail already exists.
type TrailAlreadyExistsException struct {
	Message *string
}

func (e *TrailAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrailAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrailAlreadyExistsException) ErrorCode() string             { return "TrailAlreadyExistsException" }
func (e *TrailAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the trail with the given name is not found.
type TrailNotFoundException struct {
	Message *string
}

func (e *TrailNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrailNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrailNotFoundException) ErrorCode() string             { return "TrailNotFoundException" }
func (e *TrailNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is no longer in use.
type TrailNotProvidedException struct {
	Message *string
}

func (e *TrailNotProvidedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrailNotProvidedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrailNotProvidedException) ErrorCode() string             { return "TrailNotProvidedException" }
func (e *TrailNotProvidedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception is thrown when the requested operation is not supported.
type UnsupportedOperationException struct {
	Message *string
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string             { return "UnsupportedOperationException" }
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
