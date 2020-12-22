// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration for a web distribution. When you update a
// distribution, there are more required fields than when you create a
// distribution. When you update your distribution by using this API action, follow
// the steps here to get the current configuration and then make your updates, to
// make sure that you include all of the required fields. To view a summary, see
// Required Fields for Create Distribution and Update Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide. The update process includes getting
// the current distribution configuration, updating the XML document that is
// returned to make your changes, and then submitting an UpdateDistribution request
// to make the updates. For information about updating a distribution using the
// CloudFront console instead, see Creating a Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-creating-console.html)
// in the Amazon CloudFront Developer Guide. To update a web distribution using the
// CloudFront API
//
// * Submit a GetDistributionConfig
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html)
// request to get the current configuration and an Etag header
//
// for the
// distribution. If you update the distribution again, you must get a new Etag
// header.
//
// * Update the XML document that was returned in the response to your
// GetDistributionConfig request to include your changes. When you edit the XML
// file, be aware of the following:
//
// * You must strip out the ETag parameter that
// is returned.
//
// * Additional fields are required when you update a distribution.
// There may be fields included in the XML file for features that you haven't
// configured for your distribution. This is expected and required to successfully
// update the distribution.
//
// * You can't change the value of CallerReference. If
// you try to change this value, CloudFront returns an
//
// IllegalUpdate error.
//
// * The
// new configuration replaces the existing configuration; the values that you
// specify in an UpdateDistribution request are not merged into your existing
// configuration. When you add, delete, or replace values in an element that allows
// multiple values (for example, CNAME), you must specify all of the values that
// you want to appear in the updated distribution. In addition,
//
// you must update
// the corresponding Quantity element.
//
// * Submit an UpdateDistribution request to
// update the configuration for your distribution:
//
// * In the request body, include
// the XML document that you updated in Step 2. The request body must include
// an
//
// XML document with a DistributionConfig element.
//
// * Set the value of the HTTP
// If-Match header to the value of the ETag header that CloudFront returned
//
// when
// you submitted the GetDistributionConfig request in Step 1.
//
// * Review the
// response to the UpdateDistribution request to confirm that the configuration
// was
//
// successfully updated.
//
// * Optional: Submit a GetDistribution
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html)
// request to confirm that your changes have propagated.
//
// When propagation is
// complete, the value of Status is Deployed.
func (c *Client) UpdateDistribution(ctx context.Context, params *UpdateDistributionInput, optFns ...func(*Options)) (*UpdateDistributionOutput, error) {
	if params == nil {
		params = &UpdateDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDistribution", params, optFns, addOperationUpdateDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to update a distribution.
type UpdateDistributionInput struct {

	// The distribution's configuration information.
	//
	// This member is required.
	DistributionConfig *types.DistributionConfig

	// The distribution's id.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the
	// distribution's configuration. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

// The returned result of the corresponding request.
type UpdateDistributionOutput struct {

	// The distribution's information.
	Distribution *types.Distribution

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateDistribution{}, middleware.After)
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
	if err = addOpUpdateDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateDistribution",
	}
}
