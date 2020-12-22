// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a unique asymmetric data key pair. The
// GenerateDataKeyPairWithoutPlaintext operation returns a plaintext public key and
// a copy of the private key that is encrypted under the symmetric CMK you specify.
// Unlike GenerateDataKeyPair, this operation does not return a plaintext private
// key. To generate a data key pair, you must specify a symmetric customer master
// key (CMK) to encrypt the private key in the data key pair. You cannot use an
// asymmetric CMK or a CMK in a custom key store. To get the type and origin of
// your CMK, use the KeySpec field in the DescribeKey response. You can use the
// public key that GenerateDataKeyPairWithoutPlaintext returns to encrypt data or
// verify a signature outside of AWS KMS. Then, store the encrypted private key
// with the data. When you are ready to decrypt data or sign a message, you can use
// the Decrypt operation to decrypt the encrypted private key.
// GenerateDataKeyPairWithoutPlaintext returns a unique data key pair for each
// request. The bytes in the key are not related to the caller or CMK that is used
// to encrypt the private key. You can use the optional encryption context to add
// additional security to the encryption operation. If you specify an
// EncryptionContext, you must specify the same encryption context (a
// case-sensitive exact match) when decrypting the encrypted data key. Otherwise,
// the request to decrypt fails with an InvalidCiphertextException. For more
// information, see Encryption Context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the AWS Key Management Service Developer Guide. The CMK that you use for this
// operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) GenerateDataKeyPairWithoutPlaintext(ctx context.Context, params *GenerateDataKeyPairWithoutPlaintextInput, optFns ...func(*Options)) (*GenerateDataKeyPairWithoutPlaintextOutput, error) {
	if params == nil {
		params = &GenerateDataKeyPairWithoutPlaintextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataKeyPairWithoutPlaintext", params, optFns, addOperationGenerateDataKeyPairWithoutPlaintextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataKeyPairWithoutPlaintextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyPairWithoutPlaintextInput struct {

	// Specifies the CMK that encrypts the private key in the data key pair. You must
	// specify a symmetric CMK. You cannot use an asymmetric CMK or a CMK in a custom
	// key store. To get the type and origin of your CMK, use the DescribeKey
	// operation. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias
	// name, or alias ARN. When using an alias name, prefix it with "alias/". For
	// example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys or DescribeKey. To get the alias name and alias ARN,
	// use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Determines the type of data key pair that is generated. The AWS KMS rule that
	// restricts the use of asymmetric RSA CMKs to encrypt and decrypt or to sign and
	// verify (but not both), and the rule that permits you to use ECC CMKs only to
	// sign and verify, are not effective outside of AWS KMS.
	//
	// This member is required.
	KeyPairSpec types.DataKeyPairSpec

	// Specifies the encryption context that will be used when encrypting the private
	// key in the data key pair. An encryption context is a collection of non-secret
	// key-value pairs that represents additional authenticated data. When you use an
	// encryption context to encrypt data, you must specify the same (an exact
	// case-sensitive match) encryption context to decrypt the data. An encryption
	// context is optional when encrypting with a symmetric CMK, but it is highly
	// recommended. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string
}

type GenerateDataKeyPairWithoutPlaintextOutput struct {

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that encrypted the private key.
	KeyId *string

	// The type of data key pair that was generated.
	KeyPairSpec types.DataKeyPairSpec

	// The encrypted copy of the private key. When you use the HTTP API or the AWS CLI,
	// the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	PrivateKeyCiphertextBlob []byte

	// The public key (in plaintext).
	PublicKey []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGenerateDataKeyPairWithoutPlaintextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKeyPairWithoutPlaintext{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKeyPairWithoutPlaintext{}, middleware.After)
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
	if err = addOpGenerateDataKeyPairWithoutPlaintextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKeyPairWithoutPlaintext(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateDataKeyPairWithoutPlaintext(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GenerateDataKeyPairWithoutPlaintext",
	}
}
