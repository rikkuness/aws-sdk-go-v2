// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestClient_JsonLists_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonListsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes JSON lists
		"RestJsonLists": {
			Params: &JsonListsInput{
				StringList: []string{
					"foo",
					"bar",
				},
				StringSet: []string{
					"foo",
					"bar",
				},
				IntegerList: []int32{
					1,
					2,
				},
				BooleanList: []bool{
					true,
					false,
				},
				TimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1398796238),
					smithytime.ParseEpochSeconds(1398796238),
				},
				EnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				NestedStringList: [][]string{
					{
						"foo",
						"bar",
					},
					{
						"baz",
						"qux",
					},
				},
				StructureList: []types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringList": [
			        "foo",
			        "bar"
			    ],
			    "stringSet": [
			        "foo",
			        "bar"
			    ],
			    "integerList": [
			        1,
			        2
			    ],
			    "booleanList": [
			        true,
			        false
			    ],
			    "timestampList": [
			        1398796238,
			        1398796238
			    ],
			    "enumList": [
			        "Foo",
			        "0"
			    ],
			    "nestedStringList": [
			        [
			            "foo",
			            "bar"
			        ],
			        [
			            "baz",
			            "qux"
			        ]
			    ],
			    "myStructureList": [
			        {
			            "value": "1",
			            "other": "2"
			        },
			        {
			            "value": "3",
			            "other": "4"
			        }
			    ]
			}`))
			},
		},
		// Serializes empty JSON lists
		"RestJsonListsEmpty": {
			Params: &JsonListsInput{
				StringList: []string{},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringList": []
			}`))
			},
		},
		// Serializes null values in lists
		"RestJsonListsSerializeNull": {
			Params: &JsonListsInput{
				SparseStringList: []*string{
					nil,
					ptr.String("hi"),
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonLists",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "sparseStringList": [
			        null,
			        "hi"
			    ]
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "RestJsonListsSerializeNull" {
				t.Skip("disabled test aws.protocoltests.restjson#RestJson aws.protocoltests.restjson#JsonLists")
			}

			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.JsonLists(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonLists_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonListsOutput
	}{
		// Serializes JSON lists
		"RestJsonLists": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringList": [
			        "foo",
			        "bar"
			    ],
			    "stringSet": [
			        "foo",
			        "bar"
			    ],
			    "integerList": [
			        1,
			        2
			    ],
			    "booleanList": [
			        true,
			        false
			    ],
			    "timestampList": [
			        1398796238,
			        1398796238
			    ],
			    "enumList": [
			        "Foo",
			        "0"
			    ],
			    "nestedStringList": [
			        [
			            "foo",
			            "bar"
			        ],
			        [
			            "baz",
			            "qux"
			        ]
			    ],
			    "myStructureList": [
			        {
			            "value": "1",
			            "other": "2"
			        },
			        {
			            "value": "3",
			            "other": "4"
			        }
			    ]
			}`),
			ExpectResult: &JsonListsOutput{
				StringList: []string{
					"foo",
					"bar",
				},
				StringSet: []string{
					"foo",
					"bar",
				},
				IntegerList: []int32{
					1,
					2,
				},
				BooleanList: []bool{
					true,
					false,
				},
				TimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1398796238),
					smithytime.ParseEpochSeconds(1398796238),
				},
				EnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				NestedStringList: [][]string{
					{
						"foo",
						"bar",
					},
					{
						"baz",
						"qux",
					},
				},
				StructureList: []types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
		},
		// Serializes empty JSON lists
		"RestJsonListsEmpty": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringList": []
			}`),
			ExpectResult: &JsonListsOutput{
				StringList: []string{},
			},
		},
		// Serializes null values in sparse lists
		"RestJsonListsSerializeNull": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "sparseStringList": [
			        null,
			        "hi"
			    ]
			}`),
			ExpectResult: &JsonListsOutput{
				SparseStringList: []*string{
					nil,
					ptr.String("hi"),
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params JsonListsInput
			result, err := client.JsonLists(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
