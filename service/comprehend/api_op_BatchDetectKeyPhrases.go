// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectKeyPhrasesRequest
type BatchDetectKeyPhrasesInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify English ("en") or Spanish
	// ("es"). All documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document must contain fewer that 5,000 bytes
	// of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectKeyPhrasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectKeyPhrasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectKeyPhrasesInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectKeyPhrasesResponse
type BatchDetectKeyPhrasesOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectKeyPhrasesItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectKeyPhrasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDetectKeyPhrases = "BatchDetectKeyPhrases"

// BatchDetectKeyPhrasesRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Detects the key noun phrases found in a batch of documents.
//
//    // Example sending a request using BatchDetectKeyPhrasesRequest.
//    req := client.BatchDetectKeyPhrasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectKeyPhrases
func (c *Client) BatchDetectKeyPhrasesRequest(input *BatchDetectKeyPhrasesInput) BatchDetectKeyPhrasesRequest {
	op := &aws.Operation{
		Name:       opBatchDetectKeyPhrases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDetectKeyPhrasesInput{}
	}

	req := c.newRequest(op, input, &BatchDetectKeyPhrasesOutput{})
	return BatchDetectKeyPhrasesRequest{Request: req, Input: input, Copy: c.BatchDetectKeyPhrasesRequest}
}

// BatchDetectKeyPhrasesRequest is the request type for the
// BatchDetectKeyPhrases API operation.
type BatchDetectKeyPhrasesRequest struct {
	*aws.Request
	Input *BatchDetectKeyPhrasesInput
	Copy  func(*BatchDetectKeyPhrasesInput) BatchDetectKeyPhrasesRequest
}

// Send marshals and sends the BatchDetectKeyPhrases API request.
func (r BatchDetectKeyPhrasesRequest) Send(ctx context.Context) (*BatchDetectKeyPhrasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectKeyPhrasesResponse{
		BatchDetectKeyPhrasesOutput: r.Request.Data.(*BatchDetectKeyPhrasesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectKeyPhrasesResponse is the response type for the
// BatchDetectKeyPhrases API operation.
type BatchDetectKeyPhrasesResponse struct {
	*BatchDetectKeyPhrasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectKeyPhrases request.
func (r *BatchDetectKeyPhrasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}