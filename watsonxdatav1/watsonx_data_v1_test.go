/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package watsonxdatav1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`WatsonxDataV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(watsonxDataService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
				URL: "https://watsonxdatav1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(watsonxDataService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSONX_DATA_URL":       "https://watsonxdatav1/api",
				"WATSONX_DATA_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(&watsonxdatav1.WatsonxDataV1Options{})
				Expect(watsonxDataService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := watsonxDataService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != watsonxDataService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(watsonxDataService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(watsonxDataService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(&watsonxdatav1.WatsonxDataV1Options{
					URL: "https://testService/api",
				})
				Expect(watsonxDataService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := watsonxDataService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != watsonxDataService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(watsonxDataService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(watsonxDataService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(&watsonxdatav1.WatsonxDataV1Options{})
				err := watsonxDataService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := watsonxDataService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != watsonxDataService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(watsonxDataService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(watsonxDataService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSONX_DATA_URL":       "https://watsonxdatav1/api",
				"WATSONX_DATA_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(&watsonxdatav1.WatsonxDataV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(watsonxDataService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSONX_DATA_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(&watsonxdatav1.WatsonxDataV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(watsonxDataService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = watsonxdatav1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateDbConnUsers(createDbConnUsersOptions *CreateDbConnUsersOptions) - Operation response error`, func() {
		createDbConnUsersPath := "/access/databases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDbConnUsersPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDbConnUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsModel := new(watsonxdatav1.CreateDbConnUsersOptions)
				createDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDbConnUsers(createDbConnUsersOptions *CreateDbConnUsersOptions)`, func() {
		createDbConnUsersPath := "/access/databases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDbConnUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateDbConnUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsModel := new(watsonxdatav1.CreateDbConnUsersOptions)
				createDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDbConnUsersWithContext(ctx, createDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDbConnUsersWithContext(ctx, createDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDbConnUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDbConnUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsModel := new(watsonxdatav1.CreateDbConnUsersOptions)
				createDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDbConnUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsModel := new(watsonxdatav1.CreateDbConnUsersOptions)
				createDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDbConnUsersOptions model with no property values
				createDbConnUsersOptionsModelNew := new(watsonxdatav1.CreateDbConnUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsModel := new(watsonxdatav1.CreateDbConnUsersOptions)
				createDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDbConnUsers(createDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataPolicies(listDataPoliciesOptions *ListDataPoliciesOptions) - Operation response error`, func() {
		listDataPoliciesPath := "/access/data_policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_metadata query parameter
					// TODO: Add check for include_rules query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataPolicies with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := new(watsonxdatav1.ListDataPoliciesOptions)
				listDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.CatalogName = core.StringPtr("testString")
				listDataPoliciesOptionsModel.Status = core.StringPtr("testString")
				listDataPoliciesOptionsModel.IncludeMetadata = core.BoolPtr(true)
				listDataPoliciesOptionsModel.IncludeRules = core.BoolPtr(true)
				listDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataPolicies(listDataPoliciesOptions *ListDataPoliciesOptions)`, func() {
		listDataPoliciesPath := "/access/data_policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_metadata query parameter
					// TODO: Add check for include_rules query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}], "total_count": 10}`)
				}))
			})
			It(`Invoke ListDataPolicies successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := new(watsonxdatav1.ListDataPoliciesOptions)
				listDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.CatalogName = core.StringPtr("testString")
				listDataPoliciesOptionsModel.Status = core.StringPtr("testString")
				listDataPoliciesOptionsModel.IncludeMetadata = core.BoolPtr(true)
				listDataPoliciesOptionsModel.IncludeRules = core.BoolPtr(true)
				listDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListDataPoliciesWithContext(ctx, listDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListDataPoliciesWithContext(ctx, listDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"testString"}))
					// TODO: Add check for include_metadata query parameter
					// TODO: Add check for include_rules query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}], "total_count": 10}`)
				}))
			})
			It(`Invoke ListDataPolicies successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListDataPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := new(watsonxdatav1.ListDataPoliciesOptions)
				listDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.CatalogName = core.StringPtr("testString")
				listDataPoliciesOptionsModel.Status = core.StringPtr("testString")
				listDataPoliciesOptionsModel.IncludeMetadata = core.BoolPtr(true)
				listDataPoliciesOptionsModel.IncludeRules = core.BoolPtr(true)
				listDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataPolicies with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := new(watsonxdatav1.ListDataPoliciesOptions)
				listDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.CatalogName = core.StringPtr("testString")
				listDataPoliciesOptionsModel.Status = core.StringPtr("testString")
				listDataPoliciesOptionsModel.IncludeMetadata = core.BoolPtr(true)
				listDataPoliciesOptionsModel.IncludeRules = core.BoolPtr(true)
				listDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataPolicies successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := new(watsonxdatav1.ListDataPoliciesOptions)
				listDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDataPoliciesOptionsModel.CatalogName = core.StringPtr("testString")
				listDataPoliciesOptionsModel.Status = core.StringPtr("testString")
				listDataPoliciesOptionsModel.IncludeMetadata = core.BoolPtr(true)
				listDataPoliciesOptionsModel.IncludeRules = core.BoolPtr(true)
				listDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListDataPolicies(listDataPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataPolicy(createDataPolicyOptions *CreateDataPolicyOptions) - Operation response error`, func() {
		createDataPolicyPath := "/access/data_policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataPolicy with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsModel := new(watsonxdatav1.CreateDataPolicyOptions)
				createDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				createDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				createDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				createDataPolicyOptionsModel.Description = core.StringPtr("testString")
				createDataPolicyOptionsModel.Status = core.StringPtr("active")
				createDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataPolicy(createDataPolicyOptions *CreateDataPolicyOptions)`, func() {
		createDataPolicyPath := "/access/data_policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataPolicyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"data_policy": {"catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "description": "Description", "policy_name": "PolicyName", "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active"}, "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke CreateDataPolicy successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsModel := new(watsonxdatav1.CreateDataPolicyOptions)
				createDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				createDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				createDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				createDataPolicyOptionsModel.Description = core.StringPtr("testString")
				createDataPolicyOptionsModel.Status = core.StringPtr("active")
				createDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDataPolicyWithContext(ctx, createDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDataPolicyWithContext(ctx, createDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataPolicyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"data_policy": {"catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "description": "Description", "policy_name": "PolicyName", "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active"}, "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke CreateDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDataPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsModel := new(watsonxdatav1.CreateDataPolicyOptions)
				createDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				createDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				createDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				createDataPolicyOptionsModel.Description = core.StringPtr("testString")
				createDataPolicyOptionsModel.Status = core.StringPtr("active")
				createDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataPolicy with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsModel := new(watsonxdatav1.CreateDataPolicyOptions)
				createDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				createDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				createDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				createDataPolicyOptionsModel.Description = core.StringPtr("testString")
				createDataPolicyOptionsModel.Status = core.StringPtr("active")
				createDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataPolicyOptions model with no property values
				createDataPolicyOptionsModelNew := new(watsonxdatav1.CreateDataPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsModel := new(watsonxdatav1.CreateDataPolicyOptions)
				createDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				createDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				createDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				createDataPolicyOptionsModel.Description = core.StringPtr("testString")
				createDataPolicyOptionsModel.Status = core.StringPtr("active")
				createDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDataPolicy(createDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDataPolicies(deleteDataPoliciesOptions *DeleteDataPoliciesOptions)`, func() {
		deleteDataPoliciesPath := "/access/data_policies"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDataPoliciesPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDataPolicies successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDataPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataPoliciesOptions model
				deleteDataPoliciesOptionsModel := new(watsonxdatav1.DeleteDataPoliciesOptions)
				deleteDataPoliciesOptionsModel.DataPolicies = []string{"testString"}
				deleteDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDataPolicies(deleteDataPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataPolicies with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDataPoliciesOptions model
				deleteDataPoliciesOptionsModel := new(watsonxdatav1.DeleteDataPoliciesOptions)
				deleteDataPoliciesOptionsModel.DataPolicies = []string{"testString"}
				deleteDataPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDataPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDataPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDataPolicies(deleteDataPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEngineUsers(getEngineUsersOptions *GetEngineUsersOptions) - Operation response error`, func() {
		getEngineUsersPath := "/access/engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEngineUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEngineUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEngineUsersOptions model
				getEngineUsersOptionsModel := new(watsonxdatav1.GetEngineUsersOptions)
				getEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				getEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEngineUsers(getEngineUsersOptions *GetEngineUsersOptions)`, func() {
		getEngineUsersPath := "/access/engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEngineUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine_id": "EngineID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}`)
				}))
			})
			It(`Invoke GetEngineUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetEngineUsersOptions model
				getEngineUsersOptionsModel := new(watsonxdatav1.GetEngineUsersOptions)
				getEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				getEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetEngineUsersWithContext(ctx, getEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetEngineUsersWithContext(ctx, getEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEngineUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine_id": "EngineID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}`)
				}))
			})
			It(`Invoke GetEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetEngineUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEngineUsersOptions model
				getEngineUsersOptionsModel := new(watsonxdatav1.GetEngineUsersOptions)
				getEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				getEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEngineUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEngineUsersOptions model
				getEngineUsersOptionsModel := new(watsonxdatav1.GetEngineUsersOptions)
				getEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				getEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEngineUsersOptions model with no property values
				getEngineUsersOptionsModelNew := new(watsonxdatav1.GetEngineUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetEngineUsers(getEngineUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEngineUsersOptions model
				getEngineUsersOptionsModel := new(watsonxdatav1.GetEngineUsersOptions)
				getEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				getEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetEngineUsers(getEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEngineUsers(deleteEngineUsersOptions *DeleteEngineUsersOptions)`, func() {
		deleteEngineUsersPath := "/access/engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEngineUsersPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteEngineUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEngineUsersOptions model
				deleteEngineUsersOptionsModel := new(watsonxdatav1.DeleteEngineUsersOptions)
				deleteEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.Groups = []string{"testString"}
				deleteEngineUsersOptionsModel.Users = []string{"testString"}
				deleteEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteEngineUsers(deleteEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEngineUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteEngineUsersOptions model
				deleteEngineUsersOptionsModel := new(watsonxdatav1.DeleteEngineUsersOptions)
				deleteEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.Groups = []string{"testString"}
				deleteEngineUsersOptionsModel.Users = []string{"testString"}
				deleteEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteEngineUsers(deleteEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEngineUsersOptions model with no property values
				deleteEngineUsersOptionsModelNew := new(watsonxdatav1.DeleteEngineUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteEngineUsers(deleteEngineUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEngineUsers(updateEngineUsersOptions *UpdateEngineUsersOptions) - Operation response error`, func() {
		updateEngineUsersPath := "/access/engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEngineUsersPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEngineUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineUsersOptions model
				updateEngineUsersOptionsModel := new(watsonxdatav1.UpdateEngineUsersOptions)
				updateEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				updateEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				updateEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEngineUsers(updateEngineUsersOptions *UpdateEngineUsersOptions)`, func() {
		updateEngineUsersPath := "/access/engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEngineUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateEngineUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineUsersOptions model
				updateEngineUsersOptionsModel := new(watsonxdatav1.UpdateEngineUsersOptions)
				updateEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				updateEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				updateEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateEngineUsersWithContext(ctx, updateEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateEngineUsersWithContext(ctx, updateEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEngineUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateEngineUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineUsersOptions model
				updateEngineUsersOptionsModel := new(watsonxdatav1.UpdateEngineUsersOptions)
				updateEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				updateEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				updateEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEngineUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineUsersOptions model
				updateEngineUsersOptionsModel := new(watsonxdatav1.UpdateEngineUsersOptions)
				updateEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				updateEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				updateEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEngineUsersOptions model with no property values
				updateEngineUsersOptionsModelNew := new(watsonxdatav1.UpdateEngineUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineUsersOptions model
				updateEngineUsersOptionsModel := new(watsonxdatav1.UpdateEngineUsersOptions)
				updateEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				updateEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				updateEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateEngineUsers(updateEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDbConnUsers(deleteDbConnUsersOptions *DeleteDbConnUsersOptions)`, func() {
		deleteDbConnUsersPath := "/access/databases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDbConnUsersPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDbConnUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDbConnUsersOptions model
				deleteDbConnUsersOptionsModel := new(watsonxdatav1.DeleteDbConnUsersOptions)
				deleteDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.Groups = []string{"testString"}
				deleteDbConnUsersOptionsModel.Users = []string{"testString"}
				deleteDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDbConnUsers(deleteDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDbConnUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDbConnUsersOptions model
				deleteDbConnUsersOptionsModel := new(watsonxdatav1.DeleteDbConnUsersOptions)
				deleteDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.Groups = []string{"testString"}
				deleteDbConnUsersOptionsModel.Users = []string{"testString"}
				deleteDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDbConnUsers(deleteDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDbConnUsersOptions model with no property values
				deleteDbConnUsersOptionsModelNew := new(watsonxdatav1.DeleteDbConnUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteDbConnUsers(deleteDbConnUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDbConnUsers(updateDbConnUsersOptions *UpdateDbConnUsersOptions) - Operation response error`, func() {
		updateDbConnUsersPath := "/access/databases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDbConnUsersPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDbConnUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateDbConnUsersOptions model
				updateDbConnUsersOptionsModel := new(watsonxdatav1.UpdateDbConnUsersOptions)
				updateDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDbConnUsers(updateDbConnUsersOptions *UpdateDbConnUsersOptions)`, func() {
		updateDbConnUsersPath := "/access/databases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDbConnUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateDbConnUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateDbConnUsersOptions model
				updateDbConnUsersOptionsModel := new(watsonxdatav1.UpdateDbConnUsersOptions)
				updateDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateDbConnUsersWithContext(ctx, updateDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateDbConnUsersWithContext(ctx, updateDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDbConnUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateDbConnUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateDbConnUsersOptions model
				updateDbConnUsersOptionsModel := new(watsonxdatav1.UpdateDbConnUsersOptions)
				updateDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDbConnUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateDbConnUsersOptions model
				updateDbConnUsersOptionsModel := new(watsonxdatav1.UpdateDbConnUsersOptions)
				updateDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDbConnUsersOptions model with no property values
				updateDbConnUsersOptionsModelNew := new(watsonxdatav1.UpdateDbConnUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateDbConnUsersOptions model
				updateDbConnUsersOptionsModel := new(watsonxdatav1.UpdateDbConnUsersOptions)
				updateDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateDbConnUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDbConnUsers(getDbConnUsersOptions *GetDbConnUsersOptions) - Operation response error`, func() {
		getDbConnUsersPath := "/access/databases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDbConnUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDbConnUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDbConnUsersOptions model
				getDbConnUsersOptionsModel := new(watsonxdatav1.GetDbConnUsersOptions)
				getDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDbConnUsers(getDbConnUsersOptions *GetDbConnUsersOptions)`, func() {
		getDbConnUsersPath := "/access/databases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDbConnUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"user_name": "UserName", "permission": "can_administer"}], "database_id": "DatabaseID"}`)
				}))
			})
			It(`Invoke GetDbConnUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDbConnUsersOptions model
				getDbConnUsersOptionsModel := new(watsonxdatav1.GetDbConnUsersOptions)
				getDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDbConnUsersWithContext(ctx, getDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDbConnUsersWithContext(ctx, getDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDbConnUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"user_name": "UserName", "permission": "can_administer"}], "database_id": "DatabaseID"}`)
				}))
			})
			It(`Invoke GetDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDbConnUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDbConnUsersOptions model
				getDbConnUsersOptionsModel := new(watsonxdatav1.GetDbConnUsersOptions)
				getDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDbConnUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDbConnUsersOptions model
				getDbConnUsersOptionsModel := new(watsonxdatav1.GetDbConnUsersOptions)
				getDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDbConnUsersOptions model with no property values
				getDbConnUsersOptionsModelNew := new(watsonxdatav1.GetDbConnUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDbConnUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDbConnUsersOptions model
				getDbConnUsersOptionsModel := new(watsonxdatav1.GetDbConnUsersOptions)
				getDbConnUsersOptionsModel.DatabaseID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDbConnUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDbConnUsers(getDbConnUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogUsers(createCatalogUsersOptions *CreateCatalogUsersOptions) - Operation response error`, func() {
		createCatalogUsersPath := "/access/catalogs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogUsersPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsModel := new(watsonxdatav1.CreateCatalogUsersOptions)
				createCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				createCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				createCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogUsers(createCatalogUsersOptions *CreateCatalogUsersOptions)`, func() {
		createCatalogUsersPath := "/access/catalogs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateCatalogUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsModel := new(watsonxdatav1.CreateCatalogUsersOptions)
				createCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				createCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				createCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateCatalogUsersWithContext(ctx, createCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateCatalogUsersWithContext(ctx, createCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateCatalogUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsModel := new(watsonxdatav1.CreateCatalogUsersOptions)
				createCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				createCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				createCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCatalogUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsModel := new(watsonxdatav1.CreateCatalogUsersOptions)
				createCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				createCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				createCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogUsersOptions model with no property values
				createCatalogUsersOptionsModelNew := new(watsonxdatav1.CreateCatalogUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsModel := new(watsonxdatav1.CreateCatalogUsersOptions)
				createCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				createCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				createCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateCatalogUsers(createCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogUsers(getCatalogUsersOptions *GetCatalogUsersOptions) - Operation response error`, func() {
		getCatalogUsersPath := "/access/catalogs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogUsersOptions model
				getCatalogUsersOptionsModel := new(watsonxdatav1.GetCatalogUsersOptions)
				getCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				getCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogUsers(getCatalogUsersOptions *GetCatalogUsersOptions)`, func() {
		getCatalogUsersPath := "/access/catalogs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}], "catalog_name": "CatalogName", "groups": [{"group_id": "GroupID", "permission": "can_administer"}]}`)
				}))
			})
			It(`Invoke GetCatalogUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogUsersOptions model
				getCatalogUsersOptionsModel := new(watsonxdatav1.GetCatalogUsersOptions)
				getCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				getCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetCatalogUsersWithContext(ctx, getCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetCatalogUsersWithContext(ctx, getCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}], "catalog_name": "CatalogName", "groups": [{"group_id": "GroupID", "permission": "can_administer"}]}`)
				}))
			})
			It(`Invoke GetCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetCatalogUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogUsersOptions model
				getCatalogUsersOptionsModel := new(watsonxdatav1.GetCatalogUsersOptions)
				getCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				getCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalogUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogUsersOptions model
				getCatalogUsersOptionsModel := new(watsonxdatav1.GetCatalogUsersOptions)
				getCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				getCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogUsersOptions model with no property values
				getCatalogUsersOptionsModelNew := new(watsonxdatav1.GetCatalogUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogUsersOptions model
				getCatalogUsersOptionsModel := new(watsonxdatav1.GetCatalogUsersOptions)
				getCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				getCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetCatalogUsers(getCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCatalogUsers(deleteCatalogUsersOptions *DeleteCatalogUsersOptions)`, func() {
		deleteCatalogUsersPath := "/access/catalogs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogUsersPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteCatalogUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogUsersOptions model
				deleteCatalogUsersOptionsModel := new(watsonxdatav1.DeleteCatalogUsersOptions)
				deleteCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.Groups = []string{"testString"}
				deleteCatalogUsersOptionsModel.Users = []string{"testString"}
				deleteCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteCatalogUsers(deleteCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogUsersOptions model
				deleteCatalogUsersOptionsModel := new(watsonxdatav1.DeleteCatalogUsersOptions)
				deleteCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.Groups = []string{"testString"}
				deleteCatalogUsersOptionsModel.Users = []string{"testString"}
				deleteCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteCatalogUsers(deleteCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogUsersOptions model with no property values
				deleteCatalogUsersOptionsModelNew := new(watsonxdatav1.DeleteCatalogUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteCatalogUsers(deleteCatalogUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogUsers(updateCatalogUsersOptions *UpdateCatalogUsersOptions) - Operation response error`, func() {
		updateCatalogUsersPath := "/access/catalogs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogUsersPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogUsersOptions model
				updateCatalogUsersOptionsModel := new(watsonxdatav1.UpdateCatalogUsersOptions)
				updateCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				updateCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				updateCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogUsers(updateCatalogUsersOptions *UpdateCatalogUsersOptions)`, func() {
		updateCatalogUsersPath := "/access/catalogs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateCatalogUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogUsersOptions model
				updateCatalogUsersOptionsModel := new(watsonxdatav1.UpdateCatalogUsersOptions)
				updateCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				updateCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				updateCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateCatalogUsersWithContext(ctx, updateCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateCatalogUsersWithContext(ctx, updateCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateCatalogUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogUsersOptions model
				updateCatalogUsersOptionsModel := new(watsonxdatav1.UpdateCatalogUsersOptions)
				updateCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				updateCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				updateCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalogUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogUsersOptions model
				updateCatalogUsersOptionsModel := new(watsonxdatav1.UpdateCatalogUsersOptions)
				updateCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				updateCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				updateCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogUsersOptions model with no property values
				updateCatalogUsersOptionsModelNew := new(watsonxdatav1.UpdateCatalogUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCatalogUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogUsersOptions model
				updateCatalogUsersOptionsModel := new(watsonxdatav1.UpdateCatalogUsersOptions)
				updateCatalogUsersOptionsModel.CatalogName = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Groups = []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}
				updateCatalogUsersOptionsModel.Users = []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}
				updateCatalogUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateCatalogUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Evaluate(evaluateOptions *EvaluateOptions) - Operation response error`, func() {
		evaluatePath := "/access/evaluation"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(evaluatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Evaluate with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := new(watsonxdatav1.EvaluateOptions)
				evaluateOptionsModel.Resources = []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}
				evaluateOptionsModel.LhInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.AuthInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Evaluate(evaluateOptions *EvaluateOptions)`, func() {
		evaluatePath := "/access/evaluation"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(evaluatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"action": "Action", "resource_name": "ResourceName", "resource_type": "ResourceType", "result": true}]}`)
				}))
			})
			It(`Invoke Evaluate successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := new(watsonxdatav1.EvaluateOptions)
				evaluateOptionsModel.Resources = []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}
				evaluateOptionsModel.LhInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.AuthInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.EvaluateWithContext(ctx, evaluateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.EvaluateWithContext(ctx, evaluateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(evaluatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"action": "Action", "resource_name": "ResourceName", "resource_type": "ResourceType", "result": true}]}`)
				}))
			})
			It(`Invoke Evaluate successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.Evaluate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := new(watsonxdatav1.EvaluateOptions)
				evaluateOptionsModel.Resources = []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}
				evaluateOptionsModel.LhInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.AuthInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Evaluate with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := new(watsonxdatav1.EvaluateOptions)
				evaluateOptionsModel.Resources = []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}
				evaluateOptionsModel.LhInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.AuthInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Evaluate successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := new(watsonxdatav1.EvaluateOptions)
				evaluateOptionsModel.Resources = []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}
				evaluateOptionsModel.LhInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.AuthInstanceID = core.StringPtr("testString")
				evaluateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.Evaluate(evaluateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPoliciesList(getPoliciesListOptions *GetPoliciesListOptions) - Operation response error`, func() {
		getPoliciesListPath := "/access/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPoliciesListPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_data_policies query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPoliciesList with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := new(watsonxdatav1.GetPoliciesListOptions)
				getPoliciesListOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.CatalogList = []string{"testString"}
				getPoliciesListOptionsModel.EngineList = []string{"testString"}
				getPoliciesListOptionsModel.DataPoliciesList = []string{"testString"}
				getPoliciesListOptionsModel.IncludeDataPolicies = core.BoolPtr(true)
				getPoliciesListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPoliciesList(getPoliciesListOptions *GetPoliciesListOptions)`, func() {
		getPoliciesListPath := "/access/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPoliciesListPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_data_policies query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog_policies": [{"total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}], "catalog_name": "CatalogName", "groups": [{"group_id": "GroupID", "permission": "can_administer"}]}], "data_policies": [{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}], "engine_policies": [{"engine_id": "EngineID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}]}`)
				}))
			})
			It(`Invoke GetPoliciesList successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := new(watsonxdatav1.GetPoliciesListOptions)
				getPoliciesListOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.CatalogList = []string{"testString"}
				getPoliciesListOptionsModel.EngineList = []string{"testString"}
				getPoliciesListOptionsModel.DataPoliciesList = []string{"testString"}
				getPoliciesListOptionsModel.IncludeDataPolicies = core.BoolPtr(true)
				getPoliciesListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetPoliciesListWithContext(ctx, getPoliciesListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetPoliciesListWithContext(ctx, getPoliciesListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPoliciesListPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for include_data_policies query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog_policies": [{"total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}], "catalog_name": "CatalogName", "groups": [{"group_id": "GroupID", "permission": "can_administer"}]}], "data_policies": [{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}], "engine_policies": [{"engine_id": "EngineID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}]}`)
				}))
			})
			It(`Invoke GetPoliciesList successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetPoliciesList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := new(watsonxdatav1.GetPoliciesListOptions)
				getPoliciesListOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.CatalogList = []string{"testString"}
				getPoliciesListOptionsModel.EngineList = []string{"testString"}
				getPoliciesListOptionsModel.DataPoliciesList = []string{"testString"}
				getPoliciesListOptionsModel.IncludeDataPolicies = core.BoolPtr(true)
				getPoliciesListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPoliciesList with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := new(watsonxdatav1.GetPoliciesListOptions)
				getPoliciesListOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.CatalogList = []string{"testString"}
				getPoliciesListOptionsModel.EngineList = []string{"testString"}
				getPoliciesListOptionsModel.DataPoliciesList = []string{"testString"}
				getPoliciesListOptionsModel.IncludeDataPolicies = core.BoolPtr(true)
				getPoliciesListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPoliciesList successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := new(watsonxdatav1.GetPoliciesListOptions)
				getPoliciesListOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPoliciesListOptionsModel.CatalogList = []string{"testString"}
				getPoliciesListOptionsModel.EngineList = []string{"testString"}
				getPoliciesListOptionsModel.DataPoliciesList = []string{"testString"}
				getPoliciesListOptionsModel.IncludeDataPolicies = core.BoolPtr(true)
				getPoliciesListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetPoliciesList(getPoliciesListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateMetastoreUsers(createMetastoreUsersOptions *CreateMetastoreUsersOptions) - Operation response error`, func() {
		createMetastoreUsersPath := "/access/metastores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMetastoreUsersPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateMetastoreUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsModel := new(watsonxdatav1.CreateMetastoreUsersOptions)
				createMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				createMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				createMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateMetastoreUsers(createMetastoreUsersOptions *CreateMetastoreUsersOptions)`, func() {
		createMetastoreUsersPath := "/access/metastores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMetastoreUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateMetastoreUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsModel := new(watsonxdatav1.CreateMetastoreUsersOptions)
				createMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				createMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				createMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateMetastoreUsersWithContext(ctx, createMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateMetastoreUsersWithContext(ctx, createMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createMetastoreUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateMetastoreUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsModel := new(watsonxdatav1.CreateMetastoreUsersOptions)
				createMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				createMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				createMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateMetastoreUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsModel := new(watsonxdatav1.CreateMetastoreUsersOptions)
				createMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				createMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				createMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateMetastoreUsersOptions model with no property values
				createMetastoreUsersOptionsModelNew := new(watsonxdatav1.CreateMetastoreUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsModel := new(watsonxdatav1.CreateMetastoreUsersOptions)
				createMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				createMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				createMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetastoreUsers(getMetastoreUsersOptions *GetMetastoreUsersOptions) - Operation response error`, func() {
		getMetastoreUsersPath := "/access/metastores/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoreUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetastoreUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoreUsersOptions model
				getMetastoreUsersOptionsModel := new(watsonxdatav1.GetMetastoreUsersOptions)
				getMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetastoreUsers(getMetastoreUsersOptions *GetMetastoreUsersOptions)`, func() {
		getMetastoreUsersPath := "/access/metastores/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoreUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"group_id": "GroupID", "permission": "can_administer"}], "metastore_name": "MetastoreName", "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}`)
				}))
			})
			It(`Invoke GetMetastoreUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetMetastoreUsersOptions model
				getMetastoreUsersOptionsModel := new(watsonxdatav1.GetMetastoreUsersOptions)
				getMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetMetastoreUsersWithContext(ctx, getMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetMetastoreUsersWithContext(ctx, getMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoreUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"group_id": "GroupID", "permission": "can_administer"}], "metastore_name": "MetastoreName", "total_count": 10, "users": [{"permission": "can_administer", "user_name": "UserName"}]}`)
				}))
			})
			It(`Invoke GetMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetMetastoreUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetastoreUsersOptions model
				getMetastoreUsersOptionsModel := new(watsonxdatav1.GetMetastoreUsersOptions)
				getMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetastoreUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoreUsersOptions model
				getMetastoreUsersOptionsModel := new(watsonxdatav1.GetMetastoreUsersOptions)
				getMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMetastoreUsersOptions model with no property values
				getMetastoreUsersOptionsModelNew := new(watsonxdatav1.GetMetastoreUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoreUsersOptions model
				getMetastoreUsersOptionsModel := new(watsonxdatav1.GetMetastoreUsersOptions)
				getMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteMetastoreUsers(deleteMetastoreUsersOptions *DeleteMetastoreUsersOptions)`, func() {
		deleteMetastoreUsersPath := "/access/metastores/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteMetastoreUsersPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteMetastoreUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteMetastoreUsersOptions model
				deleteMetastoreUsersOptionsModel := new(watsonxdatav1.DeleteMetastoreUsersOptions)
				deleteMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.Groups = []string{"testString"}
				deleteMetastoreUsersOptionsModel.Users = []string{"testString"}
				deleteMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteMetastoreUsers(deleteMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteMetastoreUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteMetastoreUsersOptions model
				deleteMetastoreUsersOptionsModel := new(watsonxdatav1.DeleteMetastoreUsersOptions)
				deleteMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.Groups = []string{"testString"}
				deleteMetastoreUsersOptionsModel.Users = []string{"testString"}
				deleteMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteMetastoreUsers(deleteMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteMetastoreUsersOptions model with no property values
				deleteMetastoreUsersOptionsModelNew := new(watsonxdatav1.DeleteMetastoreUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteMetastoreUsers(deleteMetastoreUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMetastoreUsers(updateMetastoreUsersOptions *UpdateMetastoreUsersOptions) - Operation response error`, func() {
		updateMetastoreUsersPath := "/access/metastores/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMetastoreUsersPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMetastoreUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateMetastoreUsersOptions model
				updateMetastoreUsersOptionsModel := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				updateMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				updateMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				updateMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMetastoreUsers(updateMetastoreUsersOptions *UpdateMetastoreUsersOptions)`, func() {
		updateMetastoreUsersPath := "/access/metastores/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMetastoreUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateMetastoreUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateMetastoreUsersOptions model
				updateMetastoreUsersOptionsModel := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				updateMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				updateMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				updateMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateMetastoreUsersWithContext(ctx, updateMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateMetastoreUsersWithContext(ctx, updateMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMetastoreUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateMetastoreUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateMetastoreUsersOptions model
				updateMetastoreUsersOptionsModel := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				updateMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				updateMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				updateMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateMetastoreUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateMetastoreUsersOptions model
				updateMetastoreUsersOptionsModel := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				updateMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				updateMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				updateMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateMetastoreUsersOptions model with no property values
				updateMetastoreUsersOptionsModelNew := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateMetastoreUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the UpdateMetastoreUsersOptions model
				updateMetastoreUsersOptionsModel := new(watsonxdatav1.UpdateMetastoreUsersOptions)
				updateMetastoreUsersOptionsModel.MetastoreName = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Groups = []watsonxdatav1.GroupsMetadata{*groupsMetadataModel}
				updateMetastoreUsersOptionsModel.Users = []watsonxdatav1.UsersMetadata{*usersMetadataModel}
				updateMetastoreUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateMetastoreUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBucketUsers(createBucketUsersOptions *CreateBucketUsersOptions) - Operation response error`, func() {
		createBucketUsersPath := "/access/buckets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketUsersPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBucketUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsModel := new(watsonxdatav1.CreateBucketUsersOptions)
				createBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBucketUsers(createBucketUsersOptions *CreateBucketUsersOptions)`, func() {
		createBucketUsersPath := "/access/buckets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateBucketUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsModel := new(watsonxdatav1.CreateBucketUsersOptions)
				createBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateBucketUsersWithContext(ctx, createBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateBucketUsersWithContext(ctx, createBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateBucketUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsModel := new(watsonxdatav1.CreateBucketUsersOptions)
				createBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBucketUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsModel := new(watsonxdatav1.CreateBucketUsersOptions)
				createBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBucketUsersOptions model with no property values
				createBucketUsersOptionsModelNew := new(watsonxdatav1.CreateBucketUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsModel := new(watsonxdatav1.CreateBucketUsersOptions)
				createBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				createBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				createBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateBucketUsers(createBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDefaultPolicies(getDefaultPoliciesOptions *GetDefaultPoliciesOptions) - Operation response error`, func() {
		getDefaultPoliciesPath := "/access/default_policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDefaultPolicies with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := new(watsonxdatav1.GetDefaultPoliciesOptions)
				getDefaultPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDefaultPolicies(getDefaultPoliciesOptions *GetDefaultPoliciesOptions)`, func() {
		getDefaultPoliciesPath := "/access/default_policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"grouping_policies": [{"domain": "Domain", "inheritor": "Inheritor", "role": "Role"}], "model": "Model", "policies": [{"subject": "Subject", "actions": ["Actions"], "domain": "Domain", "object": "Object"}]}`)
				}))
			})
			It(`Invoke GetDefaultPolicies successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := new(watsonxdatav1.GetDefaultPoliciesOptions)
				getDefaultPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDefaultPoliciesWithContext(ctx, getDefaultPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDefaultPoliciesWithContext(ctx, getDefaultPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"grouping_policies": [{"domain": "Domain", "inheritor": "Inheritor", "role": "Role"}], "model": "Model", "policies": [{"subject": "Subject", "actions": ["Actions"], "domain": "Domain", "object": "Object"}]}`)
				}))
			})
			It(`Invoke GetDefaultPolicies successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDefaultPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := new(watsonxdatav1.GetDefaultPoliciesOptions)
				getDefaultPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDefaultPolicies with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := new(watsonxdatav1.GetDefaultPoliciesOptions)
				getDefaultPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDefaultPolicies successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := new(watsonxdatav1.GetDefaultPoliciesOptions)
				getDefaultPoliciesOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDefaultPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyVersion(getPolicyVersionOptions *GetPolicyVersionOptions) - Operation response error`, func() {
		getPolicyVersionPath := "/access/policy_versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyVersionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicyVersion with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := new(watsonxdatav1.GetPolicyVersionOptions)
				getPolicyVersionOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyVersion(getPolicyVersionOptions *GetPolicyVersionOptions)`, func() {
		getPolicyVersionPath := "/access/policy_versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "data_policies": [{"associate_catalog": "AssociateCatalog", "policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "database_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "engine_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "bucket_policies": [{"policy_version": "PolicyVersion", "policy_name": "PolicyName"}]}`)
				}))
			})
			It(`Invoke GetPolicyVersion successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := new(watsonxdatav1.GetPolicyVersionOptions)
				getPolicyVersionOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetPolicyVersionWithContext(ctx, getPolicyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetPolicyVersionWithContext(ctx, getPolicyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "data_policies": [{"associate_catalog": "AssociateCatalog", "policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "database_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "engine_policies": [{"policy_name": "PolicyName", "policy_version": "PolicyVersion"}], "bucket_policies": [{"policy_version": "PolicyVersion", "policy_name": "PolicyName"}]}`)
				}))
			})
			It(`Invoke GetPolicyVersion successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetPolicyVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := new(watsonxdatav1.GetPolicyVersionOptions)
				getPolicyVersionOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPolicyVersion with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := new(watsonxdatav1.GetPolicyVersionOptions)
				getPolicyVersionOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPolicyVersion successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := new(watsonxdatav1.GetPolicyVersionOptions)
				getPolicyVersionOptionsModel.LhInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPolicyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetPolicyVersion(getPolicyVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataPolicy(getDataPolicyOptions *GetDataPolicyOptions) - Operation response error`, func() {
		getDataPolicyPath := "/access/data_policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataPolicy with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDataPolicyOptions model
				getDataPolicyOptionsModel := new(watsonxdatav1.GetDataPolicyOptions)
				getDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				getDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataPolicy(getDataPolicyOptions *GetDataPolicyOptions)`, func() {
		getDataPolicyPath := "/access/data_policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}`)
				}))
			})
			It(`Invoke GetDataPolicy successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDataPolicyOptions model
				getDataPolicyOptionsModel := new(watsonxdatav1.GetDataPolicyOptions)
				getDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				getDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDataPolicyWithContext(ctx, getDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDataPolicyWithContext(ctx, getDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rule_count": 9, "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active", "catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "policy_name": "PolicyName"}`)
				}))
			})
			It(`Invoke GetDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDataPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataPolicyOptions model
				getDataPolicyOptionsModel := new(watsonxdatav1.GetDataPolicyOptions)
				getDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				getDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataPolicy with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDataPolicyOptions model
				getDataPolicyOptionsModel := new(watsonxdatav1.GetDataPolicyOptions)
				getDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				getDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataPolicyOptions model with no property values
				getDataPolicyOptionsModelNew := new(watsonxdatav1.GetDataPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetDataPolicy(getDataPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDataPolicyOptions model
				getDataPolicyOptionsModel := new(watsonxdatav1.GetDataPolicyOptions)
				getDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				getDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDataPolicy(getDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceDataPolicy(replaceDataPolicyOptions *ReplaceDataPolicyOptions) - Operation response error`, func() {
		replaceDataPolicyPath := "/access/data_policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceDataPolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceDataPolicy with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the ReplaceDataPolicyOptions model
				replaceDataPolicyOptionsModel := new(watsonxdatav1.ReplaceDataPolicyOptions)
				replaceDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				replaceDataPolicyOptionsModel.Description = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Status = core.StringPtr("active")
				replaceDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceDataPolicy(replaceDataPolicyOptions *ReplaceDataPolicyOptions)`, func() {
		replaceDataPolicyPath := "/access/data_policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceDataPolicyPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"data_policy": {"catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "description": "Description", "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active"}, "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke ReplaceDataPolicy successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the ReplaceDataPolicyOptions model
				replaceDataPolicyOptionsModel := new(watsonxdatav1.ReplaceDataPolicyOptions)
				replaceDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				replaceDataPolicyOptionsModel.Description = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Status = core.StringPtr("active")
				replaceDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ReplaceDataPolicyWithContext(ctx, replaceDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ReplaceDataPolicyWithContext(ctx, replaceDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceDataPolicyPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"data_policy": {"catalog_name": "CatalogName", "data_artifact": "schema1/table1/(column1|column2)", "description": "Description", "rules": [{"actions": ["all"], "effect": "allow", "grantee": {"value": "Value", "key": "user_name", "type": "user_identity"}}], "status": "active"}, "metadata": {"creator": "Creator", "description": "Description", "modifier": "Modifier", "pid": "Pid", "policy_name": "PolicyName", "updated_at": "UpdatedAt", "version": "Version", "created_at": "CreatedAt"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke ReplaceDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ReplaceDataPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the ReplaceDataPolicyOptions model
				replaceDataPolicyOptionsModel := new(watsonxdatav1.ReplaceDataPolicyOptions)
				replaceDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				replaceDataPolicyOptionsModel.Description = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Status = core.StringPtr("active")
				replaceDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceDataPolicy with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the ReplaceDataPolicyOptions model
				replaceDataPolicyOptionsModel := new(watsonxdatav1.ReplaceDataPolicyOptions)
				replaceDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				replaceDataPolicyOptionsModel.Description = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Status = core.StringPtr("active")
				replaceDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceDataPolicyOptions model with no property values
				replaceDataPolicyOptionsModelNew := new(watsonxdatav1.ReplaceDataPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ReplaceDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel

				// Construct an instance of the ReplaceDataPolicyOptions model
				replaceDataPolicyOptionsModel := new(watsonxdatav1.ReplaceDataPolicyOptions)
				replaceDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.CatalogName = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.DataArtifact = core.StringPtr("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.Rules = []watsonxdatav1.Rule{*ruleModel}
				replaceDataPolicyOptionsModel.Description = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Status = core.StringPtr("active")
				replaceDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDataPolicy(deleteDataPolicyOptions *DeleteDataPolicyOptions)`, func() {
		deleteDataPolicyPath := "/access/data_policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDataPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDataPolicy successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDataPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataPolicyOptions model
				deleteDataPolicyOptionsModel := new(watsonxdatav1.DeleteDataPolicyOptions)
				deleteDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDataPolicy(deleteDataPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataPolicy with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDataPolicyOptions model
				deleteDataPolicyOptionsModel := new(watsonxdatav1.DeleteDataPolicyOptions)
				deleteDataPolicyOptionsModel.PolicyName = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDataPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDataPolicy(deleteDataPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDataPolicyOptions model with no property values
				deleteDataPolicyOptionsModelNew := new(watsonxdatav1.DeleteDataPolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteDataPolicy(deleteDataPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngineUsers(createEngineUsersOptions *CreateEngineUsersOptions) - Operation response error`, func() {
		createEngineUsersPath := "/access/engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineUsersPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEngineUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsModel := new(watsonxdatav1.CreateEngineUsersOptions)
				createEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				createEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				createEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngineUsers(createEngineUsersOptions *CreateEngineUsersOptions)`, func() {
		createEngineUsersPath := "/access/engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateEngineUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsModel := new(watsonxdatav1.CreateEngineUsersOptions)
				createEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				createEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				createEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEngineUsersWithContext(ctx, createEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEngineUsersWithContext(ctx, createEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEngineUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsModel := new(watsonxdatav1.CreateEngineUsersOptions)
				createEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				createEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				createEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngineUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsModel := new(watsonxdatav1.CreateEngineUsersOptions)
				createEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				createEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				createEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEngineUsersOptions model with no property values
				createEngineUsersOptionsModelNew := new(watsonxdatav1.CreateEngineUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateEngineUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsModel := new(watsonxdatav1.CreateEngineUsersOptions)
				createEngineUsersOptionsModel.EngineID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Groups = []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}
				createEngineUsersOptionsModel.Users = []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}
				createEngineUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngineUsers(createEngineUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketUsers(getBucketUsersOptions *GetBucketUsersOptions) - Operation response error`, func() {
		getBucketUsersPath := "/access/buckets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketUsersOptions model
				getBucketUsersOptionsModel := new(watsonxdatav1.GetBucketUsersOptions)
				getBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				getBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketUsers(getBucketUsersOptions *GetBucketUsersOptions)`, func() {
		getBucketUsersPath := "/access/buckets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_id": "BucketID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"user_name": "UserName", "permission": "can_administer"}]}`)
				}))
			})
			It(`Invoke GetBucketUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketUsersOptions model
				getBucketUsersOptionsModel := new(watsonxdatav1.GetBucketUsersOptions)
				getBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				getBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetBucketUsersWithContext(ctx, getBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetBucketUsersWithContext(ctx, getBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_id": "BucketID", "groups": [{"group_id": "GroupID", "permission": "can_administer"}], "total_count": 10, "users": [{"user_name": "UserName", "permission": "can_administer"}]}`)
				}))
			})
			It(`Invoke GetBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetBucketUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketUsersOptions model
				getBucketUsersOptionsModel := new(watsonxdatav1.GetBucketUsersOptions)
				getBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				getBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketUsersOptions model
				getBucketUsersOptionsModel := new(watsonxdatav1.GetBucketUsersOptions)
				getBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				getBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketUsersOptions model with no property values
				getBucketUsersOptionsModelNew := new(watsonxdatav1.GetBucketUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetBucketUsers(getBucketUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketUsersOptions model
				getBucketUsersOptionsModel := new(watsonxdatav1.GetBucketUsersOptions)
				getBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				getBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetBucketUsers(getBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteBucketUsers(deleteBucketUsersOptions *DeleteBucketUsersOptions)`, func() {
		deleteBucketUsersPath := "/access/buckets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketUsersPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteBucketUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketUsersOptions model
				deleteBucketUsersOptionsModel := new(watsonxdatav1.DeleteBucketUsersOptions)
				deleteBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.Groups = []string{"testString"}
				deleteBucketUsersOptionsModel.Users = []string{"testString"}
				deleteBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteBucketUsers(deleteBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucketUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketUsersOptions model
				deleteBucketUsersOptionsModel := new(watsonxdatav1.DeleteBucketUsersOptions)
				deleteBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.Groups = []string{"testString"}
				deleteBucketUsersOptionsModel.Users = []string{"testString"}
				deleteBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteBucketUsers(deleteBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketUsersOptions model with no property values
				deleteBucketUsersOptionsModelNew := new(watsonxdatav1.DeleteBucketUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteBucketUsers(deleteBucketUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucketUsers(updateBucketUsersOptions *UpdateBucketUsersOptions) - Operation response error`, func() {
		updateBucketUsersPath := "/access/buckets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketUsersPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBucketUsers with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateBucketUsersOptions model
				updateBucketUsersOptionsModel := new(watsonxdatav1.UpdateBucketUsersOptions)
				updateBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucketUsers(updateBucketUsersOptions *UpdateBucketUsersOptions)`, func() {
		updateBucketUsersPath := "/access/buckets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateBucketUsers successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateBucketUsersOptions model
				updateBucketUsersOptionsModel := new(watsonxdatav1.UpdateBucketUsersOptions)
				updateBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateBucketUsersWithContext(ctx, updateBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateBucketUsersWithContext(ctx, updateBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketUsersPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Lhinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Lhinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateBucketUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateBucketUsersOptions model
				updateBucketUsersOptionsModel := new(watsonxdatav1.UpdateBucketUsersOptions)
				updateBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBucketUsers with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateBucketUsersOptions model
				updateBucketUsersOptionsModel := new(watsonxdatav1.UpdateBucketUsersOptions)
				updateBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBucketUsersOptions model with no property values
				updateBucketUsersOptionsModelNew := new(watsonxdatav1.UpdateBucketUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBucketUsers successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")

				// Construct an instance of the UpdateBucketUsersOptions model
				updateBucketUsersOptionsModel := new(watsonxdatav1.UpdateBucketUsersOptions)
				updateBucketUsersOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Groups = []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}
				updateBucketUsersOptionsModel.Users = []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}
				updateBucketUsersOptionsModel.LhInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateBucketUsers(updateBucketUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuckets(getBucketsOptions *GetBucketsOptions) - Operation response error`, func() {
		getBucketsPath := "/buckets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBuckets with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := new(watsonxdatav1.GetBucketsOptions)
				getBucketsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBuckets(getBucketsOptions *GetBucketsOptions)`, func() {
		getBucketsPath := "/buckets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"buckets": [{"created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "IBM", "state": "active", "tags": ["Tags"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "actions": ["Actions"]}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetBuckets successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := new(watsonxdatav1.GetBucketsOptions)
				getBucketsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetBucketsWithContext(ctx, getBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetBucketsWithContext(ctx, getBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"buckets": [{"created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "IBM", "state": "active", "tags": ["Tags"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "actions": ["Actions"]}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetBuckets successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetBuckets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := new(watsonxdatav1.GetBucketsOptions)
				getBucketsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBuckets with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := new(watsonxdatav1.GetBucketsOptions)
				getBucketsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBuckets successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := new(watsonxdatav1.GetBucketsOptions)
				getBucketsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetBuckets(getBucketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketObjects(getBucketObjectsOptions *GetBucketObjectsOptions) - Operation response error`, func() {
		getBucketObjectsPath := "/buckets/bucket/objects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["bucket_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketObjects with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketObjectsOptions model
				getBucketObjectsOptionsModel := new(watsonxdatav1.GetBucketObjectsOptions)
				getBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketObjects(getBucketObjectsOptions *GetBucketObjectsOptions)`, func() {
		getBucketObjectsPath := "/buckets/bucket/objects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["bucket_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"objects": ["object_1"], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetBucketObjects successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketObjectsOptions model
				getBucketObjectsOptionsModel := new(watsonxdatav1.GetBucketObjectsOptions)
				getBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetBucketObjectsWithContext(ctx, getBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetBucketObjectsWithContext(ctx, getBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["bucket_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"objects": ["object_1"], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetBucketObjects successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetBucketObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketObjectsOptions model
				getBucketObjectsOptionsModel := new(watsonxdatav1.GetBucketObjectsOptions)
				getBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketObjects with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketObjectsOptions model
				getBucketObjectsOptionsModel := new(watsonxdatav1.GetBucketObjectsOptions)
				getBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketObjectsOptions model with no property values
				getBucketObjectsOptionsModelNew := new(watsonxdatav1.GetBucketObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBucketObjects successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketObjectsOptions model
				getBucketObjectsOptionsModel := new(watsonxdatav1.GetBucketObjectsOptions)
				getBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetBucketObjects(getBucketObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeactivateBucket(deactivateBucketOptions *DeactivateBucketOptions)`, func() {
		deactivateBucketPath := "/buckets/bucket/deactivate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deactivateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeactivateBucket successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the DeactivateBucketOptions model
				deactivateBucketOptionsModel := new(watsonxdatav1.DeactivateBucketOptions)
				deactivateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				deactivateBucketOptionsModel.Accept = core.StringPtr("testString")
				deactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.DeactivateBucketWithContext(ctx, deactivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.DeactivateBucket(deactivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.DeactivateBucketWithContext(ctx, deactivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deactivateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DeactivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.DeactivateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeactivateBucketOptions model
				deactivateBucketOptionsModel := new(watsonxdatav1.DeactivateBucketOptions)
				deactivateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				deactivateBucketOptionsModel.Accept = core.StringPtr("testString")
				deactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.DeactivateBucket(deactivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeactivateBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeactivateBucketOptions model
				deactivateBucketOptionsModel := new(watsonxdatav1.DeactivateBucketOptions)
				deactivateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				deactivateBucketOptionsModel.Accept = core.StringPtr("testString")
				deactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.DeactivateBucket(deactivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeactivateBucketOptions model with no property values
				deactivateBucketOptionsModelNew := new(watsonxdatav1.DeactivateBucketOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.DeactivateBucket(deactivateBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke DeactivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeactivateBucketOptions model
				deactivateBucketOptionsModel := new(watsonxdatav1.DeactivateBucketOptions)
				deactivateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				deactivateBucketOptionsModel.Accept = core.StringPtr("testString")
				deactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.DeactivateBucket(deactivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RegisterBucket(registerBucketOptions *RegisterBucketOptions) - Operation response error`, func() {
		registerBucketPath := "/buckets/bucket"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerBucketPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RegisterBucket with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")

				// Construct an instance of the RegisterBucketOptions model
				registerBucketOptionsModel := new(watsonxdatav1.RegisterBucketOptions)
				registerBucketOptionsModel.BucketDetails = bucketDetailsModel
				registerBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				registerBucketOptionsModel.TableType = core.StringPtr("iceberg")
				registerBucketOptionsModel.BucketType = core.StringPtr("ibm_cos")
				registerBucketOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				registerBucketOptionsModel.ManagedBy = core.StringPtr("ibm")
				registerBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				registerBucketOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				registerBucketOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				registerBucketOptionsModel.ThriftURI = core.StringPtr("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				registerBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RegisterBucket(registerBucketOptions *RegisterBucketOptions)`, func() {
		registerBucketPath := "/buckets/bucket"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bucket": {"bucket_display_name": "BucketDisplayName", "bucket_id": "BucketID"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke RegisterBucket successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")

				// Construct an instance of the RegisterBucketOptions model
				registerBucketOptionsModel := new(watsonxdatav1.RegisterBucketOptions)
				registerBucketOptionsModel.BucketDetails = bucketDetailsModel
				registerBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				registerBucketOptionsModel.TableType = core.StringPtr("iceberg")
				registerBucketOptionsModel.BucketType = core.StringPtr("ibm_cos")
				registerBucketOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				registerBucketOptionsModel.ManagedBy = core.StringPtr("ibm")
				registerBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				registerBucketOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				registerBucketOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				registerBucketOptionsModel.ThriftURI = core.StringPtr("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				registerBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.RegisterBucketWithContext(ctx, registerBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.RegisterBucketWithContext(ctx, registerBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(registerBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"bucket": {"bucket_display_name": "BucketDisplayName", "bucket_id": "BucketID"}, "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke RegisterBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.RegisterBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")

				// Construct an instance of the RegisterBucketOptions model
				registerBucketOptionsModel := new(watsonxdatav1.RegisterBucketOptions)
				registerBucketOptionsModel.BucketDetails = bucketDetailsModel
				registerBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				registerBucketOptionsModel.TableType = core.StringPtr("iceberg")
				registerBucketOptionsModel.BucketType = core.StringPtr("ibm_cos")
				registerBucketOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				registerBucketOptionsModel.ManagedBy = core.StringPtr("ibm")
				registerBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				registerBucketOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				registerBucketOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				registerBucketOptionsModel.ThriftURI = core.StringPtr("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				registerBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RegisterBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")

				// Construct an instance of the RegisterBucketOptions model
				registerBucketOptionsModel := new(watsonxdatav1.RegisterBucketOptions)
				registerBucketOptionsModel.BucketDetails = bucketDetailsModel
				registerBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				registerBucketOptionsModel.TableType = core.StringPtr("iceberg")
				registerBucketOptionsModel.BucketType = core.StringPtr("ibm_cos")
				registerBucketOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				registerBucketOptionsModel.ManagedBy = core.StringPtr("ibm")
				registerBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				registerBucketOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				registerBucketOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				registerBucketOptionsModel.ThriftURI = core.StringPtr("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				registerBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RegisterBucketOptions model with no property values
				registerBucketOptionsModelNew := new(watsonxdatav1.RegisterBucketOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.RegisterBucket(registerBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke RegisterBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")

				// Construct an instance of the RegisterBucketOptions model
				registerBucketOptionsModel := new(watsonxdatav1.RegisterBucketOptions)
				registerBucketOptionsModel.BucketDetails = bucketDetailsModel
				registerBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				registerBucketOptionsModel.TableType = core.StringPtr("iceberg")
				registerBucketOptionsModel.BucketType = core.StringPtr("ibm_cos")
				registerBucketOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				registerBucketOptionsModel.ManagedBy = core.StringPtr("ibm")
				registerBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				registerBucketOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				registerBucketOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				registerBucketOptionsModel.ThriftURI = core.StringPtr("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				registerBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.RegisterBucket(registerBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnregisterBucket(unregisterBucketOptions *UnregisterBucketOptions)`, func() {
		unregisterBucketPath := "/buckets/bucket"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unregisterBucketPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UnregisterBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.UnregisterBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UnregisterBucketOptions model
				unregisterBucketOptionsModel := new(watsonxdatav1.UnregisterBucketOptions)
				unregisterBucketOptionsModel.BucketID = core.StringPtr("bucket_id")
				unregisterBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				unregisterBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.UnregisterBucket(unregisterBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UnregisterBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UnregisterBucketOptions model
				unregisterBucketOptionsModel := new(watsonxdatav1.UnregisterBucketOptions)
				unregisterBucketOptionsModel.BucketID = core.StringPtr("bucket_id")
				unregisterBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				unregisterBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.UnregisterBucket(unregisterBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UnregisterBucketOptions model with no property values
				unregisterBucketOptionsModelNew := new(watsonxdatav1.UnregisterBucketOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.UnregisterBucket(unregisterBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucket(updateBucketOptions *UpdateBucketOptions) - Operation response error`, func() {
		updateBucketPath := "/buckets/bucket"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBucket with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsModel := new(watsonxdatav1.UpdateBucketOptions)
				updateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				updateBucketOptionsModel.AccessKey = core.StringPtr("<access_key>")
				updateBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				updateBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				updateBucketOptionsModel.SecretKey = core.StringPtr("<secret_key>")
				updateBucketOptionsModel.Tags = []string{"testbucket", "userbucket"}
				updateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucket(updateBucketOptions *UpdateBucketOptions)`, func() {
		updateBucketPath := "/buckets/bucket"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateBucket successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsModel := new(watsonxdatav1.UpdateBucketOptions)
				updateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				updateBucketOptionsModel.AccessKey = core.StringPtr("<access_key>")
				updateBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				updateBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				updateBucketOptionsModel.SecretKey = core.StringPtr("<secret_key>")
				updateBucketOptionsModel.Tags = []string{"testbucket", "userbucket"}
				updateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateBucketWithContext(ctx, updateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateBucketWithContext(ctx, updateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsModel := new(watsonxdatav1.UpdateBucketOptions)
				updateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				updateBucketOptionsModel.AccessKey = core.StringPtr("<access_key>")
				updateBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				updateBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				updateBucketOptionsModel.SecretKey = core.StringPtr("<secret_key>")
				updateBucketOptionsModel.Tags = []string{"testbucket", "userbucket"}
				updateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsModel := new(watsonxdatav1.UpdateBucketOptions)
				updateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				updateBucketOptionsModel.AccessKey = core.StringPtr("<access_key>")
				updateBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				updateBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				updateBucketOptionsModel.SecretKey = core.StringPtr("<secret_key>")
				updateBucketOptionsModel.Tags = []string{"testbucket", "userbucket"}
				updateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBucketOptions model with no property values
				updateBucketOptionsModelNew := new(watsonxdatav1.UpdateBucketOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateBucket(updateBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsModel := new(watsonxdatav1.UpdateBucketOptions)
				updateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				updateBucketOptionsModel.AccessKey = core.StringPtr("<access_key>")
				updateBucketOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				updateBucketOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				updateBucketOptionsModel.SecretKey = core.StringPtr("<secret_key>")
				updateBucketOptionsModel.Tags = []string{"testbucket", "userbucket"}
				updateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateBucket(updateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ActivateBucket(activateBucketOptions *ActivateBucketOptions)`, func() {
		activateBucketPath := "/buckets/bucket/activate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(activateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ActivateBucket successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ActivateBucketOptions model
				activateBucketOptionsModel := new(watsonxdatav1.ActivateBucketOptions)
				activateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				activateBucketOptionsModel.Accept = core.StringPtr("testString")
				activateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				activateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ActivateBucketWithContext(ctx, activateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ActivateBucket(activateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ActivateBucketWithContext(ctx, activateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(activateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ActivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ActivateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ActivateBucketOptions model
				activateBucketOptionsModel := new(watsonxdatav1.ActivateBucketOptions)
				activateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				activateBucketOptionsModel.Accept = core.StringPtr("testString")
				activateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				activateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ActivateBucket(activateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ActivateBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ActivateBucketOptions model
				activateBucketOptionsModel := new(watsonxdatav1.ActivateBucketOptions)
				activateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				activateBucketOptionsModel.Accept = core.StringPtr("testString")
				activateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				activateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ActivateBucket(activateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ActivateBucketOptions model with no property values
				activateBucketOptionsModelNew := new(watsonxdatav1.ActivateBucketOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ActivateBucket(activateBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ActivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ActivateBucketOptions model
				activateBucketOptionsModel := new(watsonxdatav1.ActivateBucketOptions)
				activateBucketOptionsModel.BucketID = core.StringPtr("samplebucket123")
				activateBucketOptionsModel.Accept = core.StringPtr("testString")
				activateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				activateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ActivateBucket(activateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDatabases(getDatabasesOptions *GetDatabasesOptions)`, func() {
		getDatabasesPath := "/databases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDatabases successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDatabasesOptions model
				getDatabasesOptionsModel := new(watsonxdatav1.GetDatabasesOptions)
				getDatabasesOptionsModel.Accept = core.StringPtr("testString")
				getDatabasesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDatabasesWithContext(ctx, getDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDatabases(getDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDatabasesWithContext(ctx, getDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDatabases successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDatabases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDatabasesOptions model
				getDatabasesOptionsModel := new(watsonxdatav1.GetDatabasesOptions)
				getDatabasesOptionsModel.Accept = core.StringPtr("testString")
				getDatabasesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDatabases(getDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDatabases with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDatabasesOptions model
				getDatabasesOptionsModel := new(watsonxdatav1.GetDatabasesOptions)
				getDatabasesOptionsModel.Accept = core.StringPtr("testString")
				getDatabasesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDatabases(getDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDatabases successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDatabasesOptions model
				getDatabasesOptionsModel := new(watsonxdatav1.GetDatabasesOptions)
				getDatabasesOptionsModel.Accept = core.StringPtr("testString")
				getDatabasesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDatabases(getDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDatabaseCatalog(createDatabaseCatalogOptions *CreateDatabaseCatalogOptions)`, func() {
		createDatabaseCatalogPath := "/databases/database"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseCatalogPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateDatabaseCatalog successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.StringPtr("4553")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")

				// Construct an instance of the CreateDatabaseCatalogOptions model
				createDatabaseCatalogOptionsModel := new(watsonxdatav1.CreateDatabaseCatalogOptions)
				createDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseCatalogOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseCatalogOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseCatalogOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseCatalogOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				createDatabaseCatalogOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseCatalogOptionsModel.Accept = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDatabaseCatalogWithContext(ctx, createDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDatabaseCatalogWithContext(ctx, createDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseCatalogPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDatabaseCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.StringPtr("4553")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")

				// Construct an instance of the CreateDatabaseCatalogOptions model
				createDatabaseCatalogOptionsModel := new(watsonxdatav1.CreateDatabaseCatalogOptions)
				createDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseCatalogOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseCatalogOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseCatalogOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseCatalogOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				createDatabaseCatalogOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseCatalogOptionsModel.Accept = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDatabaseCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.StringPtr("4553")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")

				// Construct an instance of the CreateDatabaseCatalogOptions model
				createDatabaseCatalogOptionsModel := new(watsonxdatav1.CreateDatabaseCatalogOptions)
				createDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseCatalogOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseCatalogOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseCatalogOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseCatalogOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				createDatabaseCatalogOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseCatalogOptionsModel.Accept = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDatabaseCatalogOptions model with no property values
				createDatabaseCatalogOptionsModelNew := new(watsonxdatav1.CreateDatabaseCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.StringPtr("4553")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")

				// Construct an instance of the CreateDatabaseCatalogOptions model
				createDatabaseCatalogOptionsModel := new(watsonxdatav1.CreateDatabaseCatalogOptions)
				createDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseCatalogOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseCatalogOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseCatalogOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseCatalogOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				createDatabaseCatalogOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseCatalogOptionsModel.Accept = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
		deleteDatabaseCatalogPath := "/databases/database"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseCatalogPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDatabaseCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDatabaseCatalogOptions model
				deleteDatabaseCatalogOptionsModel := new(watsonxdatav1.DeleteDatabaseCatalogOptions)
				deleteDatabaseCatalogOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				deleteDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDatabaseCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseCatalogOptions model
				deleteDatabaseCatalogOptionsModel := new(watsonxdatav1.DeleteDatabaseCatalogOptions)
				deleteDatabaseCatalogOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				deleteDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDatabaseCatalogOptions model with no property values
				deleteDatabaseCatalogOptionsModelNew := new(watsonxdatav1.DeleteDatabaseCatalogOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions)`, func() {
		updateDatabasePath := "/databases/database"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabasePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateDatabase successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the UpdateDatabaseBodyDatabaseDetails model
				updateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav1.UpdateDatabaseBodyDatabaseDetails)
				updateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				updateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav1.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				updateDatabaseOptionsModel.DatabaseDetails = updateDatabaseBodyDatabaseDetailsModel
				updateDatabaseOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				updateDatabaseOptionsModel.Description = core.StringPtr("External database description")
				updateDatabaseOptionsModel.Tags = []string{"testdatabase", "userdatabase"}
				updateDatabaseOptionsModel.Accept = core.StringPtr("testString")
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateDatabaseWithContext(ctx, updateDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateDatabaseWithContext(ctx, updateDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabasePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateDatabase successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateDatabase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDatabaseBodyDatabaseDetails model
				updateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav1.UpdateDatabaseBodyDatabaseDetails)
				updateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				updateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav1.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				updateDatabaseOptionsModel.DatabaseDetails = updateDatabaseBodyDatabaseDetailsModel
				updateDatabaseOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				updateDatabaseOptionsModel.Description = core.StringPtr("External database description")
				updateDatabaseOptionsModel.Tags = []string{"testdatabase", "userdatabase"}
				updateDatabaseOptionsModel.Accept = core.StringPtr("testString")
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDatabase with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateDatabaseBodyDatabaseDetails model
				updateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav1.UpdateDatabaseBodyDatabaseDetails)
				updateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				updateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav1.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				updateDatabaseOptionsModel.DatabaseDetails = updateDatabaseBodyDatabaseDetailsModel
				updateDatabaseOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				updateDatabaseOptionsModel.Description = core.StringPtr("External database description")
				updateDatabaseOptionsModel.Tags = []string{"testdatabase", "userdatabase"}
				updateDatabaseOptionsModel.Accept = core.StringPtr("testString")
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDatabaseOptions model with no property values
				updateDatabaseOptionsModelNew := new(watsonxdatav1.UpdateDatabaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateDatabase(updateDatabaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDatabase successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateDatabaseBodyDatabaseDetails model
				updateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav1.UpdateDatabaseBodyDatabaseDetails)
				updateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				updateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav1.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("new_db_id")
				updateDatabaseOptionsModel.DatabaseDetails = updateDatabaseBodyDatabaseDetailsModel
				updateDatabaseOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				updateDatabaseOptionsModel.Description = core.StringPtr("External database description")
				updateDatabaseOptionsModel.Tags = []string{"testdatabase", "userdatabase"}
				updateDatabaseOptionsModel.Accept = core.StringPtr("testString")
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PauseEngine(pauseEngineOptions *PauseEngineOptions) - Operation response error`, func() {
		pauseEnginePath := "/engines/engine/pause"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(pauseEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PauseEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsModel := new(watsonxdatav1.PauseEngineOptions)
				pauseEngineOptionsModel.EngineID = core.StringPtr("testString")
				pauseEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				pauseEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				pauseEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PauseEngine(pauseEngineOptions *PauseEngineOptions)`, func() {
		pauseEnginePath := "/engines/engine/pause"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(pauseEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke PauseEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsModel := new(watsonxdatav1.PauseEngineOptions)
				pauseEngineOptionsModel.EngineID = core.StringPtr("testString")
				pauseEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				pauseEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				pauseEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.PauseEngineWithContext(ctx, pauseEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.PauseEngineWithContext(ctx, pauseEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(pauseEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke PauseEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.PauseEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsModel := new(watsonxdatav1.PauseEngineOptions)
				pauseEngineOptionsModel.EngineID = core.StringPtr("testString")
				pauseEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				pauseEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				pauseEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PauseEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsModel := new(watsonxdatav1.PauseEngineOptions)
				pauseEngineOptionsModel.EngineID = core.StringPtr("testString")
				pauseEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				pauseEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				pauseEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PauseEngineOptions model with no property values
				pauseEngineOptionsModelNew := new(watsonxdatav1.PauseEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.PauseEngine(pauseEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke PauseEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsModel := new(watsonxdatav1.PauseEngineOptions)
				pauseEngineOptionsModel.EngineID = core.StringPtr("testString")
				pauseEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				pauseEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				pauseEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.PauseEngine(pauseEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEngines(getEnginesOptions *GetEnginesOptions) - Operation response error`, func() {
		getEnginesPath := "/engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := new(watsonxdatav1.GetEnginesOptions)
				getEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEngines(getEnginesOptions *GetEnginesOptions)`, func() {
		getEnginesPath := "/engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engines": [{"group_id": "new_group_id", "region": "us-south", "size_config": "starter", "created_on": 9, "engine_display_name": "sampleEngine", "origin": "ibm", "port": 4, "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}, "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "status": "running", "tags": ["Tags"], "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "host_name": "ibm-lh-presto-svc.com", "status_code": 10, "description": "presto engine for running sql queries", "engine_id": "sampleEngine123"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := new(watsonxdatav1.GetEnginesOptions)
				getEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetEnginesWithContext(ctx, getEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetEnginesWithContext(ctx, getEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engines": [{"group_id": "new_group_id", "region": "us-south", "size_config": "starter", "created_on": 9, "engine_display_name": "sampleEngine", "origin": "ibm", "port": 4, "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}, "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "status": "running", "tags": ["Tags"], "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "host_name": "ibm-lh-presto-svc.com", "status_code": 10, "description": "presto engine for running sql queries", "engine_id": "sampleEngine123"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := new(watsonxdatav1.GetEnginesOptions)
				getEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := new(watsonxdatav1.GetEnginesOptions)
				getEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := new(watsonxdatav1.GetEnginesOptions)
				getEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetEngines(getEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeployments(getDeploymentsOptions *GetDeploymentsOptions)`, func() {
		getDeploymentsPath := "/instance"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDeployments successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav1.GetDeploymentsOptions)
				getDeploymentsOptionsModel.Accept = core.StringPtr("testString")
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDeploymentsWithContext(ctx, getDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDeploymentsWithContext(ctx, getDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDeployments successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDeployments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav1.GetDeploymentsOptions)
				getDeploymentsOptionsModel.Accept = core.StringPtr("testString")
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeployments with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav1.GetDeploymentsOptions)
				getDeploymentsOptionsModel.Accept = core.StringPtr("testString")
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDeployments successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav1.GetDeploymentsOptions)
				getDeploymentsOptionsModel.Accept = core.StringPtr("testString")
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEngine(updateEngineOptions *UpdateEngineOptions)`, func() {
		updateEnginePath := "/engines/engine"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnginePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav1.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav1.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				updateEngineOptionsModel.Coordinator = nodeDescriptionModel
				updateEngineOptionsModel.Description = core.StringPtr("presto engine updated description")
				updateEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				updateEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				updateEngineOptionsModel.Worker = nodeDescriptionModel
				updateEngineOptionsModel.Accept = core.StringPtr("testString")
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateEngineWithContext(ctx, updateEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateEngineWithContext(ctx, updateEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnginePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav1.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav1.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				updateEngineOptionsModel.Coordinator = nodeDescriptionModel
				updateEngineOptionsModel.Description = core.StringPtr("presto engine updated description")
				updateEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				updateEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				updateEngineOptionsModel.Worker = nodeDescriptionModel
				updateEngineOptionsModel.Accept = core.StringPtr("testString")
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav1.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav1.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				updateEngineOptionsModel.Coordinator = nodeDescriptionModel
				updateEngineOptionsModel.Description = core.StringPtr("presto engine updated description")
				updateEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				updateEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				updateEngineOptionsModel.Worker = nodeDescriptionModel
				updateEngineOptionsModel.Accept = core.StringPtr("testString")
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEngineOptions model with no property values
				updateEngineOptionsModelNew := new(watsonxdatav1.UpdateEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateEngine(updateEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav1.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav1.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				updateEngineOptionsModel.Coordinator = nodeDescriptionModel
				updateEngineOptionsModel.Description = core.StringPtr("presto engine updated description")
				updateEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				updateEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				updateEngineOptionsModel.Worker = nodeDescriptionModel
				updateEngineOptionsModel.Accept = core.StringPtr("testString")
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngine(createEngineOptions *CreateEngineOptions)`, func() {
		createEnginePath := "/engines/engine"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav1.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav1.EngineDetailsBody)
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav1.CreateEngineOptions)
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.Origin = core.StringPtr("ibm")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.AssociatedCatalogs = []string{"new_catalog_1", "new_catalog_2"}
				createEngineOptionsModel.Accept = core.StringPtr("testString")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEngineWithContext(ctx, createEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEngineWithContext(ctx, createEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke CreateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav1.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav1.EngineDetailsBody)
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav1.CreateEngineOptions)
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.Origin = core.StringPtr("ibm")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.AssociatedCatalogs = []string{"new_catalog_1", "new_catalog_2"}
				createEngineOptionsModel.Accept = core.StringPtr("testString")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav1.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav1.EngineDetailsBody)
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav1.CreateEngineOptions)
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.Origin = core.StringPtr("ibm")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.AssociatedCatalogs = []string{"new_catalog_1", "new_catalog_2"}
				createEngineOptionsModel.Accept = core.StringPtr("testString")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEngineOptions model with no property values
				createEngineOptionsModelNew := new(watsonxdatav1.CreateEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEngine(createEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav1.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav1.EngineDetailsBody)
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav1.CreateEngineOptions)
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.Origin = core.StringPtr("ibm")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.AssociatedCatalogs = []string{"new_catalog_1", "new_catalog_2"}
				createEngineOptionsModel.Accept = core.StringPtr("testString")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
		deleteEnginePath := "/engines/engine"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEngineOptions model
				deleteEngineOptionsModel := new(watsonxdatav1.DeleteEngineOptions)
				deleteEngineOptionsModel.EngineID = core.StringPtr("eng_if")
				deleteEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				deleteEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteEngine(deleteEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteEngineOptions model
				deleteEngineOptionsModel := new(watsonxdatav1.DeleteEngineOptions)
				deleteEngineOptionsModel.EngineID = core.StringPtr("eng_if")
				deleteEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				deleteEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteEngine(deleteEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEngineOptions model with no property values
				deleteEngineOptionsModelNew := new(watsonxdatav1.DeleteEngineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteEngine(deleteEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResumeEngine(resumeEngineOptions *ResumeEngineOptions) - Operation response error`, func() {
		resumeEnginePath := "/engines/engine/resume"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resumeEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResumeEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsModel := new(watsonxdatav1.ResumeEngineOptions)
				resumeEngineOptionsModel.EngineID = core.StringPtr("eng_id")
				resumeEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				resumeEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				resumeEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResumeEngine(resumeEngineOptions *ResumeEngineOptions)`, func() {
		resumeEnginePath := "/engines/engine/resume"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resumeEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke ResumeEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsModel := new(watsonxdatav1.ResumeEngineOptions)
				resumeEngineOptionsModel.EngineID = core.StringPtr("eng_id")
				resumeEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				resumeEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				resumeEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ResumeEngineWithContext(ctx, resumeEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ResumeEngineWithContext(ctx, resumeEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resumeEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke ResumeEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ResumeEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsModel := new(watsonxdatav1.ResumeEngineOptions)
				resumeEngineOptionsModel.EngineID = core.StringPtr("eng_id")
				resumeEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				resumeEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				resumeEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResumeEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsModel := new(watsonxdatav1.ResumeEngineOptions)
				resumeEngineOptionsModel.EngineID = core.StringPtr("eng_id")
				resumeEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				resumeEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				resumeEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResumeEngineOptions model with no property values
				resumeEngineOptionsModelNew := new(watsonxdatav1.ResumeEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ResumeEngine(resumeEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ResumeEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsModel := new(watsonxdatav1.ResumeEngineOptions)
				resumeEngineOptionsModel.EngineID = core.StringPtr("eng_id")
				resumeEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				resumeEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				resumeEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ResumeEngine(resumeEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExplainAnalyzeStatement(explainAnalyzeStatementOptions *ExplainAnalyzeStatementOptions) - Operation response error`, func() {
		explainAnalyzeStatementPath := "/explainanalyze"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainAnalyzeStatementPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ExplainAnalyzeStatement with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsModel := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				explainAnalyzeStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				explainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExplainAnalyzeStatement(explainAnalyzeStatementOptions *ExplainAnalyzeStatementOptions)`, func() {
		explainAnalyzeStatementPath := "/explainanalyze"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainAnalyzeStatementPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "result": "Result"}`)
				}))
			})
			It(`Invoke ExplainAnalyzeStatement successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsModel := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				explainAnalyzeStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				explainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ExplainAnalyzeStatementWithContext(ctx, explainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ExplainAnalyzeStatementWithContext(ctx, explainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainAnalyzeStatementPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "result": "Result"}`)
				}))
			})
			It(`Invoke ExplainAnalyzeStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ExplainAnalyzeStatement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsModel := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				explainAnalyzeStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				explainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ExplainAnalyzeStatement with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsModel := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				explainAnalyzeStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				explainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExplainAnalyzeStatementOptions model with no property values
				explainAnalyzeStatementOptionsModelNew := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ExplainAnalyzeStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsModel := new(watsonxdatav1.ExplainAnalyzeStatementOptions)
				explainAnalyzeStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				explainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExplainStatement(explainStatementOptions *ExplainStatementOptions) - Operation response error`, func() {
		explainStatementPath := "/explain"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainStatementPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ExplainStatement with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsModel := new(watsonxdatav1.ExplainStatementOptions)
				explainStatementOptionsModel.EngineID = core.StringPtr("eng_id")
				explainStatementOptionsModel.Statement = core.StringPtr("show schemas")
				explainStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainStatementOptionsModel.Format = core.StringPtr("json")
				explainStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainStatementOptionsModel.Type = core.StringPtr("io")
				explainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExplainStatement(explainStatementOptions *ExplainStatementOptions)`, func() {
		explainStatementPath := "/explain"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainStatementPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "result": "Result"}`)
				}))
			})
			It(`Invoke ExplainStatement successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsModel := new(watsonxdatav1.ExplainStatementOptions)
				explainStatementOptionsModel.EngineID = core.StringPtr("eng_id")
				explainStatementOptionsModel.Statement = core.StringPtr("show schemas")
				explainStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainStatementOptionsModel.Format = core.StringPtr("json")
				explainStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainStatementOptionsModel.Type = core.StringPtr("io")
				explainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ExplainStatementWithContext(ctx, explainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ExplainStatementWithContext(ctx, explainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(explainStatementPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "result": "Result"}`)
				}))
			})
			It(`Invoke ExplainStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ExplainStatement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsModel := new(watsonxdatav1.ExplainStatementOptions)
				explainStatementOptionsModel.EngineID = core.StringPtr("eng_id")
				explainStatementOptionsModel.Statement = core.StringPtr("show schemas")
				explainStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainStatementOptionsModel.Format = core.StringPtr("json")
				explainStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainStatementOptionsModel.Type = core.StringPtr("io")
				explainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ExplainStatement with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsModel := new(watsonxdatav1.ExplainStatementOptions)
				explainStatementOptionsModel.EngineID = core.StringPtr("eng_id")
				explainStatementOptionsModel.Statement = core.StringPtr("show schemas")
				explainStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainStatementOptionsModel.Format = core.StringPtr("json")
				explainStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainStatementOptionsModel.Type = core.StringPtr("io")
				explainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExplainStatementOptions model with no property values
				explainStatementOptionsModelNew := new(watsonxdatav1.ExplainStatementOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ExplainStatement(explainStatementOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ExplainStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsModel := new(watsonxdatav1.ExplainStatementOptions)
				explainStatementOptionsModel.EngineID = core.StringPtr("eng_id")
				explainStatementOptionsModel.Statement = core.StringPtr("show schemas")
				explainStatementOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				explainStatementOptionsModel.Format = core.StringPtr("json")
				explainStatementOptionsModel.SchemaName = core.StringPtr("new_schema")
				explainStatementOptionsModel.Type = core.StringPtr("io")
				explainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				explainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ExplainStatement(explainStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions) - Operation response error`, func() {
		testLHConsolePath := "/ready"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testLHConsolePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TestLHConsole with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav1.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions)`, func() {
		testLHConsolePath := "/ready"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testLHConsolePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke TestLHConsole successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav1.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.TestLHConsoleWithContext(ctx, testLhConsoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.TestLHConsoleWithContext(ctx, testLhConsoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testLHConsolePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke TestLHConsole successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.TestLHConsole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav1.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke TestLHConsole with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav1.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke TestLHConsole successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav1.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetastores(getMetastoresOptions *GetMetastoresOptions) - Operation response error`, func() {
		getMetastoresPath := "/catalogs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoresPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetastores with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := new(watsonxdatav1.GetMetastoresOptions)
				getMetastoresOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetastores(getMetastoresOptions *GetMetastoresOptions)`, func() {
		getMetastoresPath := "/catalogs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "hostname": "s3a://samplehost.com", "managed_by": "ibm", "status": "running", "tags": ["Tags"], "actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "created_by": "<username>@<domain>.com", "thrift_uri": "thrift://samplehost-metastore:4354", "catalog_type": "iceberg", "description": "Iceberg catalog description", "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "created_on": "1602839833", "port": "3232"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetMetastores successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := new(watsonxdatav1.GetMetastoresOptions)
				getMetastoresOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetMetastoresWithContext(ctx, getMetastoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetMetastoresWithContext(ctx, getMetastoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetastoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "hostname": "s3a://samplehost.com", "managed_by": "ibm", "status": "running", "tags": ["Tags"], "actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "created_by": "<username>@<domain>.com", "thrift_uri": "thrift://samplehost-metastore:4354", "catalog_type": "iceberg", "description": "Iceberg catalog description", "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "created_on": "1602839833", "port": "3232"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetMetastores successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetMetastores(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := new(watsonxdatav1.GetMetastoresOptions)
				getMetastoresOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetastores with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := new(watsonxdatav1.GetMetastoresOptions)
				getMetastoresOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetMetastores successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := new(watsonxdatav1.GetMetastoresOptions)
				getMetastoresOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getMetastoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetMetastores(getMetastoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHMS(getHMSOptions *GetHMSOptions)`, func() {
		getHMSPath := "/metastores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHMSPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetHMS successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetHMSOptions model
				getHmsOptionsModel := new(watsonxdatav1.GetHMSOptions)
				getHmsOptionsModel.Accept = core.StringPtr("testString")
				getHmsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getHmsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetHMSWithContext(ctx, getHmsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetHMS(getHmsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetHMSWithContext(ctx, getHmsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHMSPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetHMS successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetHMS(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHMSOptions model
				getHmsOptionsModel := new(watsonxdatav1.GetHMSOptions)
				getHmsOptionsModel.Accept = core.StringPtr("testString")
				getHmsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getHmsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetHMS(getHmsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetHMS with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetHMSOptions model
				getHmsOptionsModel := new(watsonxdatav1.GetHMSOptions)
				getHmsOptionsModel.Accept = core.StringPtr("testString")
				getHmsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getHmsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetHMS(getHmsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetHMS successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetHMSOptions model
				getHmsOptionsModel := new(watsonxdatav1.GetHMSOptions)
				getHmsOptionsModel.Accept = core.StringPtr("testString")
				getHmsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getHmsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetHMS(getHmsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMetastoreToEngine(addMetastoreToEngineOptions *AddMetastoreToEngineOptions)`, func() {
		addMetastoreToEnginePath := "/catalogs/add_catalog_to_engine"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMetastoreToEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke AddMetastoreToEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the AddMetastoreToEngineOptions model
				addMetastoreToEngineOptionsModel := new(watsonxdatav1.AddMetastoreToEngineOptions)
				addMetastoreToEngineOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				addMetastoreToEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				addMetastoreToEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				addMetastoreToEngineOptionsModel.Accept = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.AddMetastoreToEngineWithContext(ctx, addMetastoreToEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.AddMetastoreToEngineWithContext(ctx, addMetastoreToEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMetastoreToEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke AddMetastoreToEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.AddMetastoreToEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddMetastoreToEngineOptions model
				addMetastoreToEngineOptionsModel := new(watsonxdatav1.AddMetastoreToEngineOptions)
				addMetastoreToEngineOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				addMetastoreToEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				addMetastoreToEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				addMetastoreToEngineOptionsModel.Accept = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddMetastoreToEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the AddMetastoreToEngineOptions model
				addMetastoreToEngineOptionsModel := new(watsonxdatav1.AddMetastoreToEngineOptions)
				addMetastoreToEngineOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				addMetastoreToEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				addMetastoreToEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				addMetastoreToEngineOptionsModel.Accept = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddMetastoreToEngineOptions model with no property values
				addMetastoreToEngineOptionsModelNew := new(watsonxdatav1.AddMetastoreToEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddMetastoreToEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the AddMetastoreToEngineOptions model
				addMetastoreToEngineOptionsModel := new(watsonxdatav1.AddMetastoreToEngineOptions)
				addMetastoreToEngineOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				addMetastoreToEngineOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				addMetastoreToEngineOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				addMetastoreToEngineOptionsModel.Accept = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				addMetastoreToEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveCatalogFromEngine(removeCatalogFromEngineOptions *RemoveCatalogFromEngineOptions)`, func() {
		removeCatalogFromEnginePath := "/catalogs/remove_catalog_from_engine"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeCatalogFromEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke RemoveCatalogFromEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RemoveCatalogFromEngineOptions model
				removeCatalogFromEngineOptionsModel := new(watsonxdatav1.RemoveCatalogFromEngineOptions)
				removeCatalogFromEngineOptionsModel.CatalogName = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.EngineID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Accept = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.RemoveCatalogFromEngineWithContext(ctx, removeCatalogFromEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.RemoveCatalogFromEngineWithContext(ctx, removeCatalogFromEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeCatalogFromEnginePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke RemoveCatalogFromEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.RemoveCatalogFromEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveCatalogFromEngineOptions model
				removeCatalogFromEngineOptionsModel := new(watsonxdatav1.RemoveCatalogFromEngineOptions)
				removeCatalogFromEngineOptionsModel.CatalogName = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.EngineID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Accept = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RemoveCatalogFromEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RemoveCatalogFromEngineOptions model
				removeCatalogFromEngineOptionsModel := new(watsonxdatav1.RemoveCatalogFromEngineOptions)
				removeCatalogFromEngineOptionsModel.CatalogName = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.EngineID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Accept = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveCatalogFromEngineOptions model with no property values
				removeCatalogFromEngineOptionsModelNew := new(watsonxdatav1.RemoveCatalogFromEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke RemoveCatalogFromEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RemoveCatalogFromEngineOptions model
				removeCatalogFromEngineOptionsModel := new(watsonxdatav1.RemoveCatalogFromEngineOptions)
				removeCatalogFromEngineOptionsModel.CatalogName = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.EngineID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.CreatedBy = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Accept = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				removeCatalogFromEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveQuery(saveQueryOptions *SaveQueryOptions) - Operation response error`, func() {
		saveQueryPath := "/queries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveQueryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveQuery with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SaveQueryOptions model
				saveQueryOptionsModel := new(watsonxdatav1.SaveQueryOptions)
				saveQueryOptionsModel.QueryName = core.StringPtr("testString")
				saveQueryOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				saveQueryOptionsModel.Description = core.StringPtr("query to get expense data")
				saveQueryOptionsModel.QueryString = core.StringPtr("select expenses from expenditure")
				saveQueryOptionsModel.CreatedOn = core.StringPtr("1608437933")
				saveQueryOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				saveQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				saveQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveQuery(saveQueryOptions *SaveQueryOptions)`, func() {
		saveQueryPath := "/queries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke SaveQuery successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the SaveQueryOptions model
				saveQueryOptionsModel := new(watsonxdatav1.SaveQueryOptions)
				saveQueryOptionsModel.QueryName = core.StringPtr("testString")
				saveQueryOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				saveQueryOptionsModel.Description = core.StringPtr("query to get expense data")
				saveQueryOptionsModel.QueryString = core.StringPtr("select expenses from expenditure")
				saveQueryOptionsModel.CreatedOn = core.StringPtr("1608437933")
				saveQueryOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				saveQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				saveQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.SaveQueryWithContext(ctx, saveQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.SaveQueryWithContext(ctx, saveQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke SaveQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.SaveQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveQueryOptions model
				saveQueryOptionsModel := new(watsonxdatav1.SaveQueryOptions)
				saveQueryOptionsModel.QueryName = core.StringPtr("testString")
				saveQueryOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				saveQueryOptionsModel.Description = core.StringPtr("query to get expense data")
				saveQueryOptionsModel.QueryString = core.StringPtr("select expenses from expenditure")
				saveQueryOptionsModel.CreatedOn = core.StringPtr("1608437933")
				saveQueryOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				saveQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				saveQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveQuery with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SaveQueryOptions model
				saveQueryOptionsModel := new(watsonxdatav1.SaveQueryOptions)
				saveQueryOptionsModel.QueryName = core.StringPtr("testString")
				saveQueryOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				saveQueryOptionsModel.Description = core.StringPtr("query to get expense data")
				saveQueryOptionsModel.QueryString = core.StringPtr("select expenses from expenditure")
				saveQueryOptionsModel.CreatedOn = core.StringPtr("1608437933")
				saveQueryOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				saveQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				saveQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveQueryOptions model with no property values
				saveQueryOptionsModelNew := new(watsonxdatav1.SaveQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.SaveQuery(saveQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke SaveQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SaveQueryOptions model
				saveQueryOptionsModel := new(watsonxdatav1.SaveQueryOptions)
				saveQueryOptionsModel.QueryName = core.StringPtr("testString")
				saveQueryOptionsModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				saveQueryOptionsModel.Description = core.StringPtr("query to get expense data")
				saveQueryOptionsModel.QueryString = core.StringPtr("select expenses from expenditure")
				saveQueryOptionsModel.CreatedOn = core.StringPtr("1608437933")
				saveQueryOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				saveQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				saveQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.SaveQuery(saveQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteQuery(deleteQueryOptions *DeleteQueryOptions)`, func() {
		deleteQueryPath := "/queries/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteQueryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteQueryOptions model
				deleteQueryOptionsModel := new(watsonxdatav1.DeleteQueryOptions)
				deleteQueryOptionsModel.QueryName = core.StringPtr("testString")
				deleteQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteQuery(deleteQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteQuery with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteQueryOptions model
				deleteQueryOptionsModel := new(watsonxdatav1.DeleteQueryOptions)
				deleteQueryOptionsModel.QueryName = core.StringPtr("testString")
				deleteQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteQuery(deleteQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteQueryOptions model with no property values
				deleteQueryOptionsModelNew := new(watsonxdatav1.DeleteQueryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteQuery(deleteQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateQuery(updateQueryOptions *UpdateQueryOptions) - Operation response error`, func() {
		updateQueryPath := "/queries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateQueryPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateQuery with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateQueryOptions model
				updateQueryOptionsModel := new(watsonxdatav1.UpdateQueryOptions)
				updateQueryOptionsModel.QueryName = core.StringPtr("testString")
				updateQueryOptionsModel.QueryString = core.StringPtr("testString")
				updateQueryOptionsModel.Description = core.StringPtr("testString")
				updateQueryOptionsModel.NewQueryName = core.StringPtr("testString")
				updateQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateQuery(updateQueryOptions *UpdateQueryOptions)`, func() {
		updateQueryPath := "/queries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateQueryPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateQuery successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the UpdateQueryOptions model
				updateQueryOptionsModel := new(watsonxdatav1.UpdateQueryOptions)
				updateQueryOptionsModel.QueryName = core.StringPtr("testString")
				updateQueryOptionsModel.QueryString = core.StringPtr("testString")
				updateQueryOptionsModel.Description = core.StringPtr("testString")
				updateQueryOptionsModel.NewQueryName = core.StringPtr("testString")
				updateQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateQueryWithContext(ctx, updateQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateQueryWithContext(ctx, updateQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateQueryPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke UpdateQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateQueryOptions model
				updateQueryOptionsModel := new(watsonxdatav1.UpdateQueryOptions)
				updateQueryOptionsModel.QueryName = core.StringPtr("testString")
				updateQueryOptionsModel.QueryString = core.StringPtr("testString")
				updateQueryOptionsModel.Description = core.StringPtr("testString")
				updateQueryOptionsModel.NewQueryName = core.StringPtr("testString")
				updateQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateQuery with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateQueryOptions model
				updateQueryOptionsModel := new(watsonxdatav1.UpdateQueryOptions)
				updateQueryOptionsModel.QueryName = core.StringPtr("testString")
				updateQueryOptionsModel.QueryString = core.StringPtr("testString")
				updateQueryOptionsModel.Description = core.StringPtr("testString")
				updateQueryOptionsModel.NewQueryName = core.StringPtr("testString")
				updateQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateQueryOptions model with no property values
				updateQueryOptionsModelNew := new(watsonxdatav1.UpdateQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateQuery(updateQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateQueryOptions model
				updateQueryOptionsModel := new(watsonxdatav1.UpdateQueryOptions)
				updateQueryOptionsModel.QueryName = core.StringPtr("testString")
				updateQueryOptionsModel.QueryString = core.StringPtr("testString")
				updateQueryOptionsModel.Description = core.StringPtr("testString")
				updateQueryOptionsModel.NewQueryName = core.StringPtr("testString")
				updateQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateQuery(updateQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetQueries(getQueriesOptions *GetQueriesOptions) - Operation response error`, func() {
		getQueriesPath := "/queries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQueriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetQueries with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := new(watsonxdatav1.GetQueriesOptions)
				getQueriesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetQueries(getQueriesOptions *GetQueriesOptions)`, func() {
		getQueriesPath := "/queries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQueriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"queries": [{"created_by": "<username>@<domain>.com", "created_on": "1608437933", "description": "query to get expense data", "engine_id": "sampleEngine123", "query_name": "new_query_name", "query_string": "select expenses from expenditure"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetQueries successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := new(watsonxdatav1.GetQueriesOptions)
				getQueriesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetQueriesWithContext(ctx, getQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetQueriesWithContext(ctx, getQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQueriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"queries": [{"created_by": "<username>@<domain>.com", "created_on": "1608437933", "description": "query to get expense data", "engine_id": "sampleEngine123", "query_name": "new_query_name", "query_string": "select expenses from expenditure"}], "response": {"_messageCode_": "<message code>", "message": "Success"}}`)
				}))
			})
			It(`Invoke GetQueries successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := new(watsonxdatav1.GetQueriesOptions)
				getQueriesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetQueries with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := new(watsonxdatav1.GetQueriesOptions)
				getQueriesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetQueries successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := new(watsonxdatav1.GetQueriesOptions)
				getQueriesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetQueries(getQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSchema(createSchemaOptions *CreateSchemaOptions) - Operation response error`, func() {
		createSchemaPath := "/schemas/schema"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSchema with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav1.CreateSchemaOptions)
				createSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				createSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				createSchemaOptionsModel.BucketName = core.StringPtr("sample-bucket")
				createSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
		createSchemaPath := "/schemas/schema"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateSchema successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav1.CreateSchemaOptions)
				createSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				createSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				createSchemaOptionsModel.BucketName = core.StringPtr("sample-bucket")
				createSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateSchemaWithContext(ctx, createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateSchemaWithContext(ctx, createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke CreateSchema successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav1.CreateSchemaOptions)
				createSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				createSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				createSchemaOptionsModel.BucketName = core.StringPtr("sample-bucket")
				createSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSchema with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav1.CreateSchemaOptions)
				createSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				createSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				createSchemaOptionsModel.BucketName = core.StringPtr("sample-bucket")
				createSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSchemaOptions model with no property values
				createSchemaOptionsModelNew := new(watsonxdatav1.CreateSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateSchema(createSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSchema successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav1.CreateSchemaOptions)
				createSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				createSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				createSchemaOptionsModel.BucketName = core.StringPtr("sample-bucket")
				createSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
		deleteSchemaPath := "/schemas/schema"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSchemaPath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSchema successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsModel := new(watsonxdatav1.DeleteSchemaOptions)
				deleteSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				deleteSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				deleteSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				deleteSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteSchema(deleteSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSchema with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsModel := new(watsonxdatav1.DeleteSchemaOptions)
				deleteSchemaOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				deleteSchemaOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				deleteSchemaOptionsModel.SchemaName = core.StringPtr("new_schema")
				deleteSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteSchema(deleteSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSchemaOptions model with no property values
				deleteSchemaOptionsModelNew := new(watsonxdatav1.DeleteSchemaOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteSchema(deleteSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchemas(getSchemasOptions *GetSchemasOptions) - Operation response error`, func() {
		getSchemasPath := "/schemas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchemasPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchemas with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSchemasOptions model
				getSchemasOptionsModel := new(watsonxdatav1.GetSchemasOptions)
				getSchemasOptionsModel.EngineID = core.StringPtr("testString")
				getSchemasOptionsModel.CatalogName = core.StringPtr("testString")
				getSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchemas(getSchemasOptions *GetSchemasOptions)`, func() {
		getSchemasPath := "/schemas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "schemas": ["Schemas"]}`)
				}))
			})
			It(`Invoke GetSchemas successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetSchemasOptions model
				getSchemasOptionsModel := new(watsonxdatav1.GetSchemasOptions)
				getSchemasOptionsModel.EngineID = core.StringPtr("testString")
				getSchemasOptionsModel.CatalogName = core.StringPtr("testString")
				getSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetSchemasWithContext(ctx, getSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetSchemasWithContext(ctx, getSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "schemas": ["Schemas"]}`)
				}))
			})
			It(`Invoke GetSchemas successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetSchemas(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchemasOptions model
				getSchemasOptionsModel := new(watsonxdatav1.GetSchemasOptions)
				getSchemasOptionsModel.EngineID = core.StringPtr("testString")
				getSchemasOptionsModel.CatalogName = core.StringPtr("testString")
				getSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSchemas with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSchemasOptions model
				getSchemasOptionsModel := new(watsonxdatav1.GetSchemasOptions)
				getSchemasOptionsModel.EngineID = core.StringPtr("testString")
				getSchemasOptionsModel.CatalogName = core.StringPtr("testString")
				getSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchemasOptions model with no property values
				getSchemasOptionsModelNew := new(watsonxdatav1.GetSchemasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetSchemas(getSchemasOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSchemas successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSchemasOptions model
				getSchemasOptionsModel := new(watsonxdatav1.GetSchemasOptions)
				getSchemasOptionsModel.EngineID = core.StringPtr("testString")
				getSchemasOptionsModel.CatalogName = core.StringPtr("testString")
				getSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetSchemas(getSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostQuery(postQueryOptions *PostQueryOptions)`, func() {
		postQueryPath := "/v1/statement"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postQueryPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke PostQuery successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the PostQueryOptions model
				postQueryOptionsModel := new(watsonxdatav1.PostQueryOptions)
				postQueryOptionsModel.Engine = core.StringPtr("testString")
				postQueryOptionsModel.Catalog = core.StringPtr("testString")
				postQueryOptionsModel.Schema = core.StringPtr("testString")
				postQueryOptionsModel.SqlQuery = core.StringPtr("testString")
				postQueryOptionsModel.Accept = core.StringPtr("testString")
				postQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				postQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.PostQueryWithContext(ctx, postQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.PostQuery(postQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.PostQueryWithContext(ctx, postQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postQueryPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke PostQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.PostQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostQueryOptions model
				postQueryOptionsModel := new(watsonxdatav1.PostQueryOptions)
				postQueryOptionsModel.Engine = core.StringPtr("testString")
				postQueryOptionsModel.Catalog = core.StringPtr("testString")
				postQueryOptionsModel.Schema = core.StringPtr("testString")
				postQueryOptionsModel.SqlQuery = core.StringPtr("testString")
				postQueryOptionsModel.Accept = core.StringPtr("testString")
				postQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				postQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.PostQuery(postQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostQuery with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the PostQueryOptions model
				postQueryOptionsModel := new(watsonxdatav1.PostQueryOptions)
				postQueryOptionsModel.Engine = core.StringPtr("testString")
				postQueryOptionsModel.Catalog = core.StringPtr("testString")
				postQueryOptionsModel.Schema = core.StringPtr("testString")
				postQueryOptionsModel.SqlQuery = core.StringPtr("testString")
				postQueryOptionsModel.Accept = core.StringPtr("testString")
				postQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				postQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.PostQuery(postQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostQueryOptions model with no property values
				postQueryOptionsModelNew := new(watsonxdatav1.PostQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.PostQuery(postQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostQuery successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the PostQueryOptions model
				postQueryOptionsModel := new(watsonxdatav1.PostQueryOptions)
				postQueryOptionsModel.Engine = core.StringPtr("testString")
				postQueryOptionsModel.Catalog = core.StringPtr("testString")
				postQueryOptionsModel.Schema = core.StringPtr("testString")
				postQueryOptionsModel.SqlQuery = core.StringPtr("testString")
				postQueryOptionsModel.Accept = core.StringPtr("testString")
				postQueryOptionsModel.AuthInstanceID = core.StringPtr("testString")
				postQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.PostQuery(postQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTable(deleteTableOptions *DeleteTableOptions)`, func() {
		deleteTablePath := "/tables/table"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTablePath))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteTable(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTableBodyDeleteTablesItems model
				deleteTableBodyDeleteTablesItemsModel := new(watsonxdatav1.DeleteTableBodyDeleteTablesItems)
				deleteTableBodyDeleteTablesItemsModel.CatalogName = core.StringPtr("sampleCatalog")
				deleteTableBodyDeleteTablesItemsModel.SchemaName = core.StringPtr("new_schema")
				deleteTableBodyDeleteTablesItemsModel.TableName = core.StringPtr("new_table")

				// Construct an instance of the DeleteTableOptions model
				deleteTableOptionsModel := new(watsonxdatav1.DeleteTableOptions)
				deleteTableOptionsModel.DeleteTables = []watsonxdatav1.DeleteTableBodyDeleteTablesItems{*deleteTableBodyDeleteTablesItemsModel}
				deleteTableOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				deleteTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteTable(deleteTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTable with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteTableBodyDeleteTablesItems model
				deleteTableBodyDeleteTablesItemsModel := new(watsonxdatav1.DeleteTableBodyDeleteTablesItems)
				deleteTableBodyDeleteTablesItemsModel.CatalogName = core.StringPtr("sampleCatalog")
				deleteTableBodyDeleteTablesItemsModel.SchemaName = core.StringPtr("new_schema")
				deleteTableBodyDeleteTablesItemsModel.TableName = core.StringPtr("new_table")

				// Construct an instance of the DeleteTableOptions model
				deleteTableOptionsModel := new(watsonxdatav1.DeleteTableOptions)
				deleteTableOptionsModel.DeleteTables = []watsonxdatav1.DeleteTableBodyDeleteTablesItems{*deleteTableBodyDeleteTablesItemsModel}
				deleteTableOptionsModel.EngineID = core.StringPtr("sampleEngine123")
				deleteTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteTable(deleteTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTableOptions model with no property values
				deleteTableOptionsModelNew := new(watsonxdatav1.DeleteTableOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteTable(deleteTableOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTable(updateTableOptions *UpdateTableOptions)`, func() {
		updateTablePath := "/tables/table"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTablePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["table_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateTable successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTableBodyAddColumnsItems model
				updateTableBodyAddColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyAddColumnsItems)
				updateTableBodyAddColumnsItemsModel.ColumnComment = core.StringPtr("income column")
				updateTableBodyAddColumnsItemsModel.ColumnName = core.StringPtr("income")
				updateTableBodyAddColumnsItemsModel.DataType = core.StringPtr("varchar")

				// Construct an instance of the UpdateTableBodyDropColumnsItems model
				updateTableBodyDropColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyDropColumnsItems)
				updateTableBodyDropColumnsItemsModel.ColumnName = core.StringPtr("expenditure")

				// Construct an instance of the UpdateTableBodyRenameColumnsItems model
				updateTableBodyRenameColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyRenameColumnsItems)
				updateTableBodyRenameColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				updateTableBodyRenameColumnsItemsModel.NewColumnName = core.StringPtr("expenses")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav1.UpdateTableOptions)
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.CatalogName = core.StringPtr("testString")
				updateTableOptionsModel.SchemaName = core.StringPtr("testString")
				updateTableOptionsModel.TableName = core.StringPtr("testString")
				updateTableOptionsModel.AddColumns = []watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel}
				updateTableOptionsModel.DropColumns = []watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel}
				updateTableOptionsModel.NewTableName = core.StringPtr("updated_table_name")
				updateTableOptionsModel.RenameColumns = []watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel}
				updateTableOptionsModel.Accept = core.StringPtr("testString")
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateTableWithContext(ctx, updateTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateTableWithContext(ctx, updateTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTablePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["table_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UpdateTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateTable(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTableBodyAddColumnsItems model
				updateTableBodyAddColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyAddColumnsItems)
				updateTableBodyAddColumnsItemsModel.ColumnComment = core.StringPtr("income column")
				updateTableBodyAddColumnsItemsModel.ColumnName = core.StringPtr("income")
				updateTableBodyAddColumnsItemsModel.DataType = core.StringPtr("varchar")

				// Construct an instance of the UpdateTableBodyDropColumnsItems model
				updateTableBodyDropColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyDropColumnsItems)
				updateTableBodyDropColumnsItemsModel.ColumnName = core.StringPtr("expenditure")

				// Construct an instance of the UpdateTableBodyRenameColumnsItems model
				updateTableBodyRenameColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyRenameColumnsItems)
				updateTableBodyRenameColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				updateTableBodyRenameColumnsItemsModel.NewColumnName = core.StringPtr("expenses")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav1.UpdateTableOptions)
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.CatalogName = core.StringPtr("testString")
				updateTableOptionsModel.SchemaName = core.StringPtr("testString")
				updateTableOptionsModel.TableName = core.StringPtr("testString")
				updateTableOptionsModel.AddColumns = []watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel}
				updateTableOptionsModel.DropColumns = []watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel}
				updateTableOptionsModel.NewTableName = core.StringPtr("updated_table_name")
				updateTableOptionsModel.RenameColumns = []watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel}
				updateTableOptionsModel.Accept = core.StringPtr("testString")
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTable with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateTableBodyAddColumnsItems model
				updateTableBodyAddColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyAddColumnsItems)
				updateTableBodyAddColumnsItemsModel.ColumnComment = core.StringPtr("income column")
				updateTableBodyAddColumnsItemsModel.ColumnName = core.StringPtr("income")
				updateTableBodyAddColumnsItemsModel.DataType = core.StringPtr("varchar")

				// Construct an instance of the UpdateTableBodyDropColumnsItems model
				updateTableBodyDropColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyDropColumnsItems)
				updateTableBodyDropColumnsItemsModel.ColumnName = core.StringPtr("expenditure")

				// Construct an instance of the UpdateTableBodyRenameColumnsItems model
				updateTableBodyRenameColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyRenameColumnsItems)
				updateTableBodyRenameColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				updateTableBodyRenameColumnsItemsModel.NewColumnName = core.StringPtr("expenses")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav1.UpdateTableOptions)
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.CatalogName = core.StringPtr("testString")
				updateTableOptionsModel.SchemaName = core.StringPtr("testString")
				updateTableOptionsModel.TableName = core.StringPtr("testString")
				updateTableOptionsModel.AddColumns = []watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel}
				updateTableOptionsModel.DropColumns = []watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel}
				updateTableOptionsModel.NewTableName = core.StringPtr("updated_table_name")
				updateTableOptionsModel.RenameColumns = []watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel}
				updateTableOptionsModel.Accept = core.StringPtr("testString")
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTableOptions model with no property values
				updateTableOptionsModelNew := new(watsonxdatav1.UpdateTableOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateTable(updateTableOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UpdateTableBodyAddColumnsItems model
				updateTableBodyAddColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyAddColumnsItems)
				updateTableBodyAddColumnsItemsModel.ColumnComment = core.StringPtr("income column")
				updateTableBodyAddColumnsItemsModel.ColumnName = core.StringPtr("income")
				updateTableBodyAddColumnsItemsModel.DataType = core.StringPtr("varchar")

				// Construct an instance of the UpdateTableBodyDropColumnsItems model
				updateTableBodyDropColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyDropColumnsItems)
				updateTableBodyDropColumnsItemsModel.ColumnName = core.StringPtr("expenditure")

				// Construct an instance of the UpdateTableBodyRenameColumnsItems model
				updateTableBodyRenameColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyRenameColumnsItems)
				updateTableBodyRenameColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				updateTableBodyRenameColumnsItemsModel.NewColumnName = core.StringPtr("expenses")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav1.UpdateTableOptions)
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.CatalogName = core.StringPtr("testString")
				updateTableOptionsModel.SchemaName = core.StringPtr("testString")
				updateTableOptionsModel.TableName = core.StringPtr("testString")
				updateTableOptionsModel.AddColumns = []watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel}
				updateTableOptionsModel.DropColumns = []watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel}
				updateTableOptionsModel.NewTableName = core.StringPtr("updated_table_name")
				updateTableOptionsModel.RenameColumns = []watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel}
				updateTableOptionsModel.Accept = core.StringPtr("testString")
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTableSnapshots(getTableSnapshotsOptions *GetTableSnapshotsOptions) - Operation response error`, func() {
		getTableSnapshotsPath := "/tables/table/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["table_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTableSnapshots with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableSnapshotsOptions model
				getTableSnapshotsOptionsModel := new(watsonxdatav1.GetTableSnapshotsOptions)
				getTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.CatalogName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.SchemaName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.TableName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTableSnapshots(getTableSnapshotsOptions *GetTableSnapshotsOptions)`, func() {
		getTableSnapshotsPath := "/tables/table/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["table_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "snapshots": [{"operation": "alter", "snapshot_id": "2332342122211222", "summary": {"anyKey": "anyValue"}, "committed_at": "1609379392"}]}`)
				}))
			})
			It(`Invoke GetTableSnapshots successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetTableSnapshotsOptions model
				getTableSnapshotsOptionsModel := new(watsonxdatav1.GetTableSnapshotsOptions)
				getTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.CatalogName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.SchemaName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.TableName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetTableSnapshotsWithContext(ctx, getTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetTableSnapshotsWithContext(ctx, getTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["table_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "snapshots": [{"operation": "alter", "snapshot_id": "2332342122211222", "summary": {"anyKey": "anyValue"}, "committed_at": "1609379392"}]}`)
				}))
			})
			It(`Invoke GetTableSnapshots successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetTableSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTableSnapshotsOptions model
				getTableSnapshotsOptionsModel := new(watsonxdatav1.GetTableSnapshotsOptions)
				getTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.CatalogName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.SchemaName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.TableName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTableSnapshots with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableSnapshotsOptions model
				getTableSnapshotsOptionsModel := new(watsonxdatav1.GetTableSnapshotsOptions)
				getTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.CatalogName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.SchemaName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.TableName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTableSnapshotsOptions model with no property values
				getTableSnapshotsOptionsModelNew := new(watsonxdatav1.GetTableSnapshotsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTableSnapshots successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableSnapshotsOptions model
				getTableSnapshotsOptionsModel := new(watsonxdatav1.GetTableSnapshotsOptions)
				getTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.CatalogName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.SchemaName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.TableName = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetTableSnapshots(getTableSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RollbackSnapshot(rollbackSnapshotOptions *RollbackSnapshotOptions) - Operation response error`, func() {
		rollbackSnapshotPath := "/tables/table/rollback"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rollbackSnapshotPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RollbackSnapshot with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RollbackSnapshotOptions model
				rollbackSnapshotOptionsModel := new(watsonxdatav1.RollbackSnapshotOptions)
				rollbackSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.CatalogName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SchemaName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SnapshotID = core.StringPtr("2332342122211222")
				rollbackSnapshotOptionsModel.TableName = core.StringPtr("new_table")
				rollbackSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RollbackSnapshot(rollbackSnapshotOptions *RollbackSnapshotOptions)`, func() {
		rollbackSnapshotPath := "/tables/table/rollback"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rollbackSnapshotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke RollbackSnapshot successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RollbackSnapshotOptions model
				rollbackSnapshotOptionsModel := new(watsonxdatav1.RollbackSnapshotOptions)
				rollbackSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.CatalogName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SchemaName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SnapshotID = core.StringPtr("2332342122211222")
				rollbackSnapshotOptionsModel.TableName = core.StringPtr("new_table")
				rollbackSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.RollbackSnapshotWithContext(ctx, rollbackSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.RollbackSnapshotWithContext(ctx, rollbackSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rollbackSnapshotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"_messageCode_": "<message code>", "message": "Success"}`)
				}))
			})
			It(`Invoke RollbackSnapshot successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.RollbackSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RollbackSnapshotOptions model
				rollbackSnapshotOptionsModel := new(watsonxdatav1.RollbackSnapshotOptions)
				rollbackSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.CatalogName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SchemaName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SnapshotID = core.StringPtr("2332342122211222")
				rollbackSnapshotOptionsModel.TableName = core.StringPtr("new_table")
				rollbackSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RollbackSnapshot with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RollbackSnapshotOptions model
				rollbackSnapshotOptionsModel := new(watsonxdatav1.RollbackSnapshotOptions)
				rollbackSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.CatalogName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SchemaName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SnapshotID = core.StringPtr("2332342122211222")
				rollbackSnapshotOptionsModel.TableName = core.StringPtr("new_table")
				rollbackSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RollbackSnapshotOptions model with no property values
				rollbackSnapshotOptionsModelNew := new(watsonxdatav1.RollbackSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke RollbackSnapshot successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RollbackSnapshotOptions model
				rollbackSnapshotOptionsModel := new(watsonxdatav1.RollbackSnapshotOptions)
				rollbackSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.CatalogName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SchemaName = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.SnapshotID = core.StringPtr("2332342122211222")
				rollbackSnapshotOptionsModel.TableName = core.StringPtr("new_table")
				rollbackSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				rollbackSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.RollbackSnapshot(rollbackSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTables(getTablesOptions *GetTablesOptions) - Operation response error`, func() {
		getTablesPath := "/tables"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTables with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTablesOptions model
				getTablesOptionsModel := new(watsonxdatav1.GetTablesOptions)
				getTablesOptionsModel.EngineID = core.StringPtr("testString")
				getTablesOptionsModel.CatalogName = core.StringPtr("testString")
				getTablesOptionsModel.SchemaName = core.StringPtr("testString")
				getTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTables(getTablesOptions *GetTablesOptions)`, func() {
		getTablesPath := "/tables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "tables": ["Tables"]}`)
				}))
			})
			It(`Invoke GetTables successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetTablesOptions model
				getTablesOptionsModel := new(watsonxdatav1.GetTablesOptions)
				getTablesOptionsModel.EngineID = core.StringPtr("testString")
				getTablesOptionsModel.CatalogName = core.StringPtr("testString")
				getTablesOptionsModel.SchemaName = core.StringPtr("testString")
				getTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetTablesWithContext(ctx, getTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetTablesWithContext(ctx, getTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["catalog_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["schema_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"_messageCode_": "<message code>", "message": "Success"}, "tables": ["Tables"]}`)
				}))
			})
			It(`Invoke GetTables successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetTables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTablesOptions model
				getTablesOptionsModel := new(watsonxdatav1.GetTablesOptions)
				getTablesOptionsModel.EngineID = core.StringPtr("testString")
				getTablesOptionsModel.CatalogName = core.StringPtr("testString")
				getTablesOptionsModel.SchemaName = core.StringPtr("testString")
				getTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTables with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTablesOptions model
				getTablesOptionsModel := new(watsonxdatav1.GetTablesOptions)
				getTablesOptionsModel.EngineID = core.StringPtr("testString")
				getTablesOptionsModel.CatalogName = core.StringPtr("testString")
				getTablesOptionsModel.SchemaName = core.StringPtr("testString")
				getTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTablesOptions model with no property values
				getTablesOptionsModelNew := new(watsonxdatav1.GetTablesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetTables(getTablesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTables successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTablesOptions model
				getTablesOptionsModel := new(watsonxdatav1.GetTablesOptions)
				getTablesOptionsModel.EngineID = core.StringPtr("testString")
				getTablesOptionsModel.CatalogName = core.StringPtr("testString")
				getTablesOptionsModel.SchemaName = core.StringPtr("testString")
				getTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetTables(getTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ParseCsv(parseCsvOptions *ParseCsvOptions)`, func() {
		parseCsvPath := "/parse/csv"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(parseCsvPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ParseCsv successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ParseCsvOptions model
				parseCsvOptionsModel := new(watsonxdatav1.ParseCsvOptions)
				parseCsvOptionsModel.Engine = core.StringPtr("testString")
				parseCsvOptionsModel.ParseFile = core.StringPtr("testString")
				parseCsvOptionsModel.FileType = core.StringPtr("testString")
				parseCsvOptionsModel.Accept = core.StringPtr("testString")
				parseCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				parseCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ParseCsvWithContext(ctx, parseCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ParseCsv(parseCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ParseCsvWithContext(ctx, parseCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(parseCsvPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke ParseCsv successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ParseCsv(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParseCsvOptions model
				parseCsvOptionsModel := new(watsonxdatav1.ParseCsvOptions)
				parseCsvOptionsModel.Engine = core.StringPtr("testString")
				parseCsvOptionsModel.ParseFile = core.StringPtr("testString")
				parseCsvOptionsModel.FileType = core.StringPtr("testString")
				parseCsvOptionsModel.Accept = core.StringPtr("testString")
				parseCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				parseCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ParseCsv(parseCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ParseCsv with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ParseCsvOptions model
				parseCsvOptionsModel := new(watsonxdatav1.ParseCsvOptions)
				parseCsvOptionsModel.Engine = core.StringPtr("testString")
				parseCsvOptionsModel.ParseFile = core.StringPtr("testString")
				parseCsvOptionsModel.FileType = core.StringPtr("testString")
				parseCsvOptionsModel.Accept = core.StringPtr("testString")
				parseCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				parseCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ParseCsv(parseCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ParseCsvOptions model with no property values
				parseCsvOptionsModelNew := new(watsonxdatav1.ParseCsvOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ParseCsv(parseCsvOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ParseCsv successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ParseCsvOptions model
				parseCsvOptionsModel := new(watsonxdatav1.ParseCsvOptions)
				parseCsvOptionsModel.Engine = core.StringPtr("testString")
				parseCsvOptionsModel.ParseFile = core.StringPtr("testString")
				parseCsvOptionsModel.FileType = core.StringPtr("testString")
				parseCsvOptionsModel.Accept = core.StringPtr("testString")
				parseCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				parseCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ParseCsv(parseCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UplaodCsv(uplaodCsvOptions *UplaodCsvOptions)`, func() {
		uplaodCsvPath := "/v2/upload/csv"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uplaodCsvPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UplaodCsv successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the UplaodCsvOptions model
				uplaodCsvOptionsModel := new(watsonxdatav1.UplaodCsvOptions)
				uplaodCsvOptionsModel.Engine = core.StringPtr("testString")
				uplaodCsvOptionsModel.Catalog = core.StringPtr("testString")
				uplaodCsvOptionsModel.Schema = core.StringPtr("testString")
				uplaodCsvOptionsModel.TableName = core.StringPtr("testString")
				uplaodCsvOptionsModel.IngestionJobName = core.StringPtr("testString")
				uplaodCsvOptionsModel.Scheduled = core.StringPtr("testString")
				uplaodCsvOptionsModel.CreatedBy = core.StringPtr("testString")
				uplaodCsvOptionsModel.TargetTable = core.StringPtr("testString")
				uplaodCsvOptionsModel.HeadersVar = core.StringPtr("testString")
				uplaodCsvOptionsModel.Csv = core.StringPtr("testString")
				uplaodCsvOptionsModel.Accept = core.StringPtr("testString")
				uplaodCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				uplaodCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UplaodCsvWithContext(ctx, uplaodCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UplaodCsv(uplaodCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UplaodCsvWithContext(ctx, uplaodCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uplaodCsvPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke UplaodCsv successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UplaodCsv(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UplaodCsvOptions model
				uplaodCsvOptionsModel := new(watsonxdatav1.UplaodCsvOptions)
				uplaodCsvOptionsModel.Engine = core.StringPtr("testString")
				uplaodCsvOptionsModel.Catalog = core.StringPtr("testString")
				uplaodCsvOptionsModel.Schema = core.StringPtr("testString")
				uplaodCsvOptionsModel.TableName = core.StringPtr("testString")
				uplaodCsvOptionsModel.IngestionJobName = core.StringPtr("testString")
				uplaodCsvOptionsModel.Scheduled = core.StringPtr("testString")
				uplaodCsvOptionsModel.CreatedBy = core.StringPtr("testString")
				uplaodCsvOptionsModel.TargetTable = core.StringPtr("testString")
				uplaodCsvOptionsModel.HeadersVar = core.StringPtr("testString")
				uplaodCsvOptionsModel.Csv = core.StringPtr("testString")
				uplaodCsvOptionsModel.Accept = core.StringPtr("testString")
				uplaodCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				uplaodCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UplaodCsv(uplaodCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UplaodCsv with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UplaodCsvOptions model
				uplaodCsvOptionsModel := new(watsonxdatav1.UplaodCsvOptions)
				uplaodCsvOptionsModel.Engine = core.StringPtr("testString")
				uplaodCsvOptionsModel.Catalog = core.StringPtr("testString")
				uplaodCsvOptionsModel.Schema = core.StringPtr("testString")
				uplaodCsvOptionsModel.TableName = core.StringPtr("testString")
				uplaodCsvOptionsModel.IngestionJobName = core.StringPtr("testString")
				uplaodCsvOptionsModel.Scheduled = core.StringPtr("testString")
				uplaodCsvOptionsModel.CreatedBy = core.StringPtr("testString")
				uplaodCsvOptionsModel.TargetTable = core.StringPtr("testString")
				uplaodCsvOptionsModel.HeadersVar = core.StringPtr("testString")
				uplaodCsvOptionsModel.Csv = core.StringPtr("testString")
				uplaodCsvOptionsModel.Accept = core.StringPtr("testString")
				uplaodCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				uplaodCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UplaodCsv(uplaodCsvOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UplaodCsvOptions model with no property values
				uplaodCsvOptionsModelNew := new(watsonxdatav1.UplaodCsvOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UplaodCsv(uplaodCsvOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UplaodCsv successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the UplaodCsvOptions model
				uplaodCsvOptionsModel := new(watsonxdatav1.UplaodCsvOptions)
				uplaodCsvOptionsModel.Engine = core.StringPtr("testString")
				uplaodCsvOptionsModel.Catalog = core.StringPtr("testString")
				uplaodCsvOptionsModel.Schema = core.StringPtr("testString")
				uplaodCsvOptionsModel.TableName = core.StringPtr("testString")
				uplaodCsvOptionsModel.IngestionJobName = core.StringPtr("testString")
				uplaodCsvOptionsModel.Scheduled = core.StringPtr("testString")
				uplaodCsvOptionsModel.CreatedBy = core.StringPtr("testString")
				uplaodCsvOptionsModel.TargetTable = core.StringPtr("testString")
				uplaodCsvOptionsModel.HeadersVar = core.StringPtr("testString")
				uplaodCsvOptionsModel.Csv = core.StringPtr("testString")
				uplaodCsvOptionsModel.Accept = core.StringPtr("testString")
				uplaodCsvOptionsModel.AuthInstanceID = core.StringPtr("testString")
				uplaodCsvOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UplaodCsv(uplaodCsvOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			watsonxDataService, _ := watsonxdatav1.NewWatsonxDataV1(&watsonxdatav1.WatsonxDataV1Options{
				URL:           "http://watsonxdatav1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewActivateBucketOptions successfully`, func() {
				// Construct an instance of the ActivateBucketOptions model
				activateBucketOptionsBucketID := "samplebucket123"
				activateBucketOptionsModel := watsonxDataService.NewActivateBucketOptions(activateBucketOptionsBucketID)
				activateBucketOptionsModel.SetBucketID("samplebucket123")
				activateBucketOptionsModel.SetAccept("testString")
				activateBucketOptionsModel.SetAuthInstanceID("testString")
				activateBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(activateBucketOptionsModel).ToNot(BeNil())
				Expect(activateBucketOptionsModel.BucketID).To(Equal(core.StringPtr("samplebucket123")))
				Expect(activateBucketOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(activateBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(activateBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddMetastoreToEngineOptions successfully`, func() {
				// Construct an instance of the AddMetastoreToEngineOptions model
				addMetastoreToEngineOptionsCatalogName := "sampleCatalog"
				addMetastoreToEngineOptionsEngineID := "sampleEngine123"
				addMetastoreToEngineOptionsModel := watsonxDataService.NewAddMetastoreToEngineOptions(addMetastoreToEngineOptionsCatalogName, addMetastoreToEngineOptionsEngineID)
				addMetastoreToEngineOptionsModel.SetCatalogName("sampleCatalog")
				addMetastoreToEngineOptionsModel.SetEngineID("sampleEngine123")
				addMetastoreToEngineOptionsModel.SetAccept("testString")
				addMetastoreToEngineOptionsModel.SetCreatedBy("<username>@<domain>.com")
				addMetastoreToEngineOptionsModel.SetAuthInstanceID("testString")
				addMetastoreToEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMetastoreToEngineOptionsModel).ToNot(BeNil())
				Expect(addMetastoreToEngineOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(addMetastoreToEngineOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(addMetastoreToEngineOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(addMetastoreToEngineOptionsModel.CreatedBy).To(Equal(core.StringPtr("<username>@<domain>.com")))
				Expect(addMetastoreToEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(addMetastoreToEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBucketUsersOptions successfully`, func() {
				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				Expect(bucketDbConnGroupsMetadataModel).ToNot(BeNil())
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				Expect(bucketDbConnUsersMetadataModel).ToNot(BeNil())
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the CreateBucketUsersOptions model
				createBucketUsersOptionsBucketID := "testString"
				createBucketUsersOptionsModel := watsonxDataService.NewCreateBucketUsersOptions(createBucketUsersOptionsBucketID)
				createBucketUsersOptionsModel.SetBucketID("testString")
				createBucketUsersOptionsModel.SetGroups([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel})
				createBucketUsersOptionsModel.SetUsers([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel})
				createBucketUsersOptionsModel.SetLhInstanceID("testString")
				createBucketUsersOptionsModel.SetAuthInstanceID("testString")
				createBucketUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBucketUsersOptionsModel).ToNot(BeNil())
				Expect(createBucketUsersOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(createBucketUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}))
				Expect(createBucketUsersOptionsModel.Users).To(Equal([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}))
				Expect(createBucketUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createBucketUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createBucketUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCatalogUsersOptions successfully`, func() {
				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				Expect(catalogGroupsMetadataModel).ToNot(BeNil())
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(catalogGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(catalogGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				Expect(catalogUsersMetadataModel).ToNot(BeNil())
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")
				Expect(catalogUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(catalogUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateCatalogUsersOptions model
				createCatalogUsersOptionsCatalogName := "testString"
				createCatalogUsersOptionsModel := watsonxDataService.NewCreateCatalogUsersOptions(createCatalogUsersOptionsCatalogName)
				createCatalogUsersOptionsModel.SetCatalogName("testString")
				createCatalogUsersOptionsModel.SetGroups([]watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel})
				createCatalogUsersOptionsModel.SetUsers([]watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel})
				createCatalogUsersOptionsModel.SetLhInstanceID("testString")
				createCatalogUsersOptionsModel.SetAuthInstanceID("testString")
				createCatalogUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogUsersOptionsModel).ToNot(BeNil())
				Expect(createCatalogUsersOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}))
				Expect(createCatalogUsersOptionsModel.Users).To(Equal([]watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}))
				Expect(createCatalogUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDataPolicyOptions successfully`, func() {
				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				Expect(ruleGranteeModel).ToNot(BeNil())
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")
				Expect(ruleGranteeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(ruleGranteeModel.Key).To(Equal(core.StringPtr("user_name")))
				Expect(ruleGranteeModel.Type).To(Equal(core.StringPtr("user_identity")))

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel
				Expect(ruleModel.Actions).To(Equal([]string{"all"}))
				Expect(ruleModel.Effect).To(Equal(core.StringPtr("allow")))
				Expect(ruleModel.Grantee).To(Equal(ruleGranteeModel))

				// Construct an instance of the CreateDataPolicyOptions model
				createDataPolicyOptionsCatalogName := "testString"
				createDataPolicyOptionsDataArtifact := "schema1/table1/(column1|column2)"
				createDataPolicyOptionsPolicyName := "testString"
				createDataPolicyOptionsRules := []watsonxdatav1.Rule{}
				createDataPolicyOptionsModel := watsonxDataService.NewCreateDataPolicyOptions(createDataPolicyOptionsCatalogName, createDataPolicyOptionsDataArtifact, createDataPolicyOptionsPolicyName, createDataPolicyOptionsRules)
				createDataPolicyOptionsModel.SetCatalogName("testString")
				createDataPolicyOptionsModel.SetDataArtifact("schema1/table1/(column1|column2)")
				createDataPolicyOptionsModel.SetPolicyName("testString")
				createDataPolicyOptionsModel.SetRules([]watsonxdatav1.Rule{*ruleModel})
				createDataPolicyOptionsModel.SetDescription("testString")
				createDataPolicyOptionsModel.SetStatus("active")
				createDataPolicyOptionsModel.SetLhInstanceID("testString")
				createDataPolicyOptionsModel.SetAuthInstanceID("testString")
				createDataPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataPolicyOptionsModel).ToNot(BeNil())
				Expect(createDataPolicyOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(createDataPolicyOptionsModel.DataArtifact).To(Equal(core.StringPtr("schema1/table1/(column1|column2)")))
				Expect(createDataPolicyOptionsModel.PolicyName).To(Equal(core.StringPtr("testString")))
				Expect(createDataPolicyOptionsModel.Rules).To(Equal([]watsonxdatav1.Rule{*ruleModel}))
				Expect(createDataPolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDataPolicyOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(createDataPolicyOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDataPolicyOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDataPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDatabaseCatalogOptions successfully`, func() {
				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails)
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel).ToNot(BeNil())
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.StringPtr("4553")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Password).To(Equal(core.StringPtr("samplepassword")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Port).To(Equal(core.StringPtr("4553")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl).To(Equal(core.BoolPtr(true)))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Tables).To(Equal(core.StringPtr("kafka_table_name")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Username).To(Equal(core.StringPtr("sampleuser")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName).To(Equal(core.StringPtr("new_database")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname).To(Equal(core.StringPtr("db2@<hostname>.com")))

				// Construct an instance of the CreateDatabaseCatalogOptions model
				createDatabaseCatalogOptionsDatabaseDisplayName := "new_database"
				createDatabaseCatalogOptionsDatabaseType := "db2"
				createDatabaseCatalogOptionsCatalogName := "sampleCatalog"
				createDatabaseCatalogOptionsModel := watsonxDataService.NewCreateDatabaseCatalogOptions(createDatabaseCatalogOptionsDatabaseDisplayName, createDatabaseCatalogOptionsDatabaseType, createDatabaseCatalogOptionsCatalogName)
				createDatabaseCatalogOptionsModel.SetDatabaseDisplayName("new_database")
				createDatabaseCatalogOptionsModel.SetDatabaseType("db2")
				createDatabaseCatalogOptionsModel.SetCatalogName("sampleCatalog")
				createDatabaseCatalogOptionsModel.SetAccept("testString")
				createDatabaseCatalogOptionsModel.SetDatabaseDetails(registerDatabaseCatalogBodyDatabaseDetailsModel)
				createDatabaseCatalogOptionsModel.SetDescription("db2 extenal database description")
				createDatabaseCatalogOptionsModel.SetTags([]string{"tag_1", "tag_2"})
				createDatabaseCatalogOptionsModel.SetCreatedBy("<username>@<domain>.com")
				createDatabaseCatalogOptionsModel.SetCreatedOn(int64(38))
				createDatabaseCatalogOptionsModel.SetAuthInstanceID("testString")
				createDatabaseCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDatabaseCatalogOptionsModel).ToNot(BeNil())
				Expect(createDatabaseCatalogOptionsModel.DatabaseDisplayName).To(Equal(core.StringPtr("new_database")))
				Expect(createDatabaseCatalogOptionsModel.DatabaseType).To(Equal(core.StringPtr("db2")))
				Expect(createDatabaseCatalogOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(createDatabaseCatalogOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseCatalogOptionsModel.DatabaseDetails).To(Equal(registerDatabaseCatalogBodyDatabaseDetailsModel))
				Expect(createDatabaseCatalogOptionsModel.Description).To(Equal(core.StringPtr("db2 extenal database description")))
				Expect(createDatabaseCatalogOptionsModel.Tags).To(Equal([]string{"tag_1", "tag_2"}))
				Expect(createDatabaseCatalogOptionsModel.CreatedBy).To(Equal(core.StringPtr("<username>@<domain>.com")))
				Expect(createDatabaseCatalogOptionsModel.CreatedOn).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createDatabaseCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDbConnUsersOptions successfully`, func() {
				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				Expect(bucketDbConnGroupsMetadataModel).ToNot(BeNil())
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				Expect(bucketDbConnUsersMetadataModel).ToNot(BeNil())
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the CreateDbConnUsersOptions model
				createDbConnUsersOptionsDatabaseID := "testString"
				createDbConnUsersOptionsModel := watsonxDataService.NewCreateDbConnUsersOptions(createDbConnUsersOptionsDatabaseID)
				createDbConnUsersOptionsModel.SetDatabaseID("testString")
				createDbConnUsersOptionsModel.SetGroups([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel})
				createDbConnUsersOptionsModel.SetUsers([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel})
				createDbConnUsersOptionsModel.SetLhInstanceID("testString")
				createDbConnUsersOptionsModel.SetAuthInstanceID("testString")
				createDbConnUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDbConnUsersOptionsModel).ToNot(BeNil())
				Expect(createDbConnUsersOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(createDbConnUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}))
				Expect(createDbConnUsersOptionsModel.Users).To(Equal([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}))
				Expect(createDbConnUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDbConnUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDbConnUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineOptions successfully`, func() {
				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav1.NodeDescriptionBody)
				Expect(nodeDescriptionBodyModel).ToNot(BeNil())
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))
				Expect(nodeDescriptionBodyModel.NodeType).To(Equal(core.StringPtr("worker")))
				Expect(nodeDescriptionBodyModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav1.EngineDetailsBody)
				Expect(engineDetailsBodyModel).ToNot(BeNil())
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				Expect(engineDetailsBodyModel.Worker).To(Equal(nodeDescriptionBodyModel))
				Expect(engineDetailsBodyModel.Coordinator).To(Equal(nodeDescriptionBodyModel))
				Expect(engineDetailsBodyModel.SizeConfig).To(Equal(core.StringPtr("starter")))

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsVersion := "1.2.3"
				var createEngineOptionsEngineDetails *watsonxdatav1.EngineDetailsBody = nil
				createEngineOptionsOrigin := "ibm"
				createEngineOptionsType := "presto"
				createEngineOptionsModel := watsonxDataService.NewCreateEngineOptions(createEngineOptionsVersion, createEngineOptionsEngineDetails, createEngineOptionsOrigin, createEngineOptionsType)
				createEngineOptionsModel.SetVersion("1.2.3")
				createEngineOptionsModel.SetEngineDetails(engineDetailsBodyModel)
				createEngineOptionsModel.SetOrigin("ibm")
				createEngineOptionsModel.SetType("presto")
				createEngineOptionsModel.SetAccept("testString")
				createEngineOptionsModel.SetDescription("presto engine description")
				createEngineOptionsModel.SetEngineDisplayName("sampleEngine")
				createEngineOptionsModel.SetFirstTimeUse(true)
				createEngineOptionsModel.SetRegion("us-south")
				createEngineOptionsModel.SetAssociatedCatalogs([]string{"new_catalog_1", "new_catalog_2"})
				createEngineOptionsModel.SetAuthInstanceID("testString")
				createEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineOptionsModel).ToNot(BeNil())
				Expect(createEngineOptionsModel.Version).To(Equal(core.StringPtr("1.2.3")))
				Expect(createEngineOptionsModel.EngineDetails).To(Equal(engineDetailsBodyModel))
				Expect(createEngineOptionsModel.Origin).To(Equal(core.StringPtr("ibm")))
				Expect(createEngineOptionsModel.Type).To(Equal(core.StringPtr("presto")))
				Expect(createEngineOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(createEngineOptionsModel.Description).To(Equal(core.StringPtr("presto engine description")))
				Expect(createEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(createEngineOptionsModel.FirstTimeUse).To(Equal(core.BoolPtr(true)))
				Expect(createEngineOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createEngineOptionsModel.AssociatedCatalogs).To(Equal([]string{"new_catalog_1", "new_catalog_2"}))
				Expect(createEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineUsersOptions successfully`, func() {
				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				Expect(engineGroupsMetadataModel).ToNot(BeNil())
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(engineGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(engineGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				Expect(engineUsersMetadataModel).ToNot(BeNil())
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")
				Expect(engineUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(engineUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateEngineUsersOptions model
				createEngineUsersOptionsEngineID := "testString"
				createEngineUsersOptionsModel := watsonxDataService.NewCreateEngineUsersOptions(createEngineUsersOptionsEngineID)
				createEngineUsersOptionsModel.SetEngineID("testString")
				createEngineUsersOptionsModel.SetGroups([]watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel})
				createEngineUsersOptionsModel.SetUsers([]watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel})
				createEngineUsersOptionsModel.SetLhInstanceID("testString")
				createEngineUsersOptionsModel.SetAuthInstanceID("testString")
				createEngineUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineUsersOptionsModel).ToNot(BeNil())
				Expect(createEngineUsersOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}))
				Expect(createEngineUsersOptionsModel.Users).To(Equal([]watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}))
				Expect(createEngineUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateMetastoreUsersOptions successfully`, func() {
				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				Expect(groupsMetadataModel).ToNot(BeNil())
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(groupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(groupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				Expect(usersMetadataModel).ToNot(BeNil())
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")
				Expect(usersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(usersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateMetastoreUsersOptions model
				createMetastoreUsersOptionsMetastoreName := "testString"
				createMetastoreUsersOptionsModel := watsonxDataService.NewCreateMetastoreUsersOptions(createMetastoreUsersOptionsMetastoreName)
				createMetastoreUsersOptionsModel.SetMetastoreName("testString")
				createMetastoreUsersOptionsModel.SetGroups([]watsonxdatav1.GroupsMetadata{*groupsMetadataModel})
				createMetastoreUsersOptionsModel.SetUsers([]watsonxdatav1.UsersMetadata{*usersMetadataModel})
				createMetastoreUsersOptionsModel.SetLhInstanceID("testString")
				createMetastoreUsersOptionsModel.SetAuthInstanceID("testString")
				createMetastoreUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createMetastoreUsersOptionsModel).ToNot(BeNil())
				Expect(createMetastoreUsersOptionsModel.MetastoreName).To(Equal(core.StringPtr("testString")))
				Expect(createMetastoreUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.GroupsMetadata{*groupsMetadataModel}))
				Expect(createMetastoreUsersOptionsModel.Users).To(Equal([]watsonxdatav1.UsersMetadata{*usersMetadataModel}))
				Expect(createMetastoreUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createMetastoreUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createMetastoreUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSchemaOptions successfully`, func() {
				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsCatalogName := "sampleCatalog"
				createSchemaOptionsEngineID := "sampleEngine123"
				createSchemaOptionsSchemaName := "new_schema"
				createSchemaOptionsModel := watsonxDataService.NewCreateSchemaOptions(createSchemaOptionsCatalogName, createSchemaOptionsEngineID, createSchemaOptionsSchemaName)
				createSchemaOptionsModel.SetCatalogName("sampleCatalog")
				createSchemaOptionsModel.SetEngineID("sampleEngine123")
				createSchemaOptionsModel.SetSchemaName("new_schema")
				createSchemaOptionsModel.SetBucketName("sample-bucket")
				createSchemaOptionsModel.SetAuthInstanceID("testString")
				createSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSchemaOptionsModel).ToNot(BeNil())
				Expect(createSchemaOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(createSchemaOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(createSchemaOptionsModel.SchemaName).To(Equal(core.StringPtr("new_schema")))
				Expect(createSchemaOptionsModel.BucketName).To(Equal(core.StringPtr("sample-bucket")))
				Expect(createSchemaOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeactivateBucketOptions successfully`, func() {
				// Construct an instance of the DeactivateBucketOptions model
				deactivateBucketOptionsBucketID := "samplebucket123"
				deactivateBucketOptionsModel := watsonxDataService.NewDeactivateBucketOptions(deactivateBucketOptionsBucketID)
				deactivateBucketOptionsModel.SetBucketID("samplebucket123")
				deactivateBucketOptionsModel.SetAccept("testString")
				deactivateBucketOptionsModel.SetAuthInstanceID("testString")
				deactivateBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deactivateBucketOptionsModel).ToNot(BeNil())
				Expect(deactivateBucketOptionsModel.BucketID).To(Equal(core.StringPtr("samplebucket123")))
				Expect(deactivateBucketOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(deactivateBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deactivateBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketUsersOptions successfully`, func() {
				// Construct an instance of the DeleteBucketUsersOptions model
				bucketID := "testString"
				deleteBucketUsersOptionsModel := watsonxDataService.NewDeleteBucketUsersOptions(bucketID)
				deleteBucketUsersOptionsModel.SetBucketID("testString")
				deleteBucketUsersOptionsModel.SetGroups([]string{"testString"})
				deleteBucketUsersOptionsModel.SetUsers([]string{"testString"})
				deleteBucketUsersOptionsModel.SetLhInstanceID("testString")
				deleteBucketUsersOptionsModel.SetAuthInstanceID("testString")
				deleteBucketUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketUsersOptionsModel).ToNot(BeNil())
				Expect(deleteBucketUsersOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketUsersOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(deleteBucketUsersOptionsModel.Users).To(Equal([]string{"testString"}))
				Expect(deleteBucketUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogUsersOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogUsersOptions model
				catalogName := "testString"
				deleteCatalogUsersOptionsModel := watsonxDataService.NewDeleteCatalogUsersOptions(catalogName)
				deleteCatalogUsersOptionsModel.SetCatalogName("testString")
				deleteCatalogUsersOptionsModel.SetGroups([]string{"testString"})
				deleteCatalogUsersOptionsModel.SetUsers([]string{"testString"})
				deleteCatalogUsersOptionsModel.SetLhInstanceID("testString")
				deleteCatalogUsersOptionsModel.SetAuthInstanceID("testString")
				deleteCatalogUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogUsersOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogUsersOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogUsersOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(deleteCatalogUsersOptionsModel.Users).To(Equal([]string{"testString"}))
				Expect(deleteCatalogUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDataPoliciesOptions successfully`, func() {
				// Construct an instance of the DeleteDataPoliciesOptions model
				deleteDataPoliciesOptionsModel := watsonxDataService.NewDeleteDataPoliciesOptions()
				deleteDataPoliciesOptionsModel.SetDataPolicies([]string{"testString"})
				deleteDataPoliciesOptionsModel.SetLhInstanceID("testString")
				deleteDataPoliciesOptionsModel.SetAuthInstanceID("testString")
				deleteDataPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataPoliciesOptionsModel).ToNot(BeNil())
				Expect(deleteDataPoliciesOptionsModel.DataPolicies).To(Equal([]string{"testString"}))
				Expect(deleteDataPoliciesOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataPoliciesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDataPolicyOptions successfully`, func() {
				// Construct an instance of the DeleteDataPolicyOptions model
				policyName := "testString"
				deleteDataPolicyOptionsModel := watsonxDataService.NewDeleteDataPolicyOptions(policyName)
				deleteDataPolicyOptionsModel.SetPolicyName("testString")
				deleteDataPolicyOptionsModel.SetLhInstanceID("testString")
				deleteDataPolicyOptionsModel.SetAuthInstanceID("testString")
				deleteDataPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataPolicyOptionsModel).ToNot(BeNil())
				Expect(deleteDataPolicyOptionsModel.PolicyName).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataPolicyOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataPolicyOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDatabaseCatalogOptions successfully`, func() {
				// Construct an instance of the DeleteDatabaseCatalogOptions model
				deleteDatabaseCatalogOptionsDatabaseID := "new_db_id"
				deleteDatabaseCatalogOptionsModel := watsonxDataService.NewDeleteDatabaseCatalogOptions(deleteDatabaseCatalogOptionsDatabaseID)
				deleteDatabaseCatalogOptionsModel.SetDatabaseID("new_db_id")
				deleteDatabaseCatalogOptionsModel.SetAuthInstanceID("testString")
				deleteDatabaseCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseCatalogOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseCatalogOptionsModel.DatabaseID).To(Equal(core.StringPtr("new_db_id")))
				Expect(deleteDatabaseCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDbConnUsersOptions successfully`, func() {
				// Construct an instance of the DeleteDbConnUsersOptions model
				databaseID := "testString"
				deleteDbConnUsersOptionsModel := watsonxDataService.NewDeleteDbConnUsersOptions(databaseID)
				deleteDbConnUsersOptionsModel.SetDatabaseID("testString")
				deleteDbConnUsersOptionsModel.SetGroups([]string{"testString"})
				deleteDbConnUsersOptionsModel.SetUsers([]string{"testString"})
				deleteDbConnUsersOptionsModel.SetLhInstanceID("testString")
				deleteDbConnUsersOptionsModel.SetAuthInstanceID("testString")
				deleteDbConnUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDbConnUsersOptionsModel).ToNot(BeNil())
				Expect(deleteDbConnUsersOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDbConnUsersOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(deleteDbConnUsersOptionsModel.Users).To(Equal([]string{"testString"}))
				Expect(deleteDbConnUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDbConnUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDbConnUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEngineOptions successfully`, func() {
				// Construct an instance of the DeleteEngineOptions model
				deleteEngineOptionsEngineID := "eng_if"
				deleteEngineOptionsModel := watsonxDataService.NewDeleteEngineOptions(deleteEngineOptionsEngineID)
				deleteEngineOptionsModel.SetEngineID("eng_if")
				deleteEngineOptionsModel.SetCreatedBy("<username>@<domain>.com")
				deleteEngineOptionsModel.SetAuthInstanceID("testString")
				deleteEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEngineOptionsModel).ToNot(BeNil())
				Expect(deleteEngineOptionsModel.EngineID).To(Equal(core.StringPtr("eng_if")))
				Expect(deleteEngineOptionsModel.CreatedBy).To(Equal(core.StringPtr("<username>@<domain>.com")))
				Expect(deleteEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEngineUsersOptions successfully`, func() {
				// Construct an instance of the DeleteEngineUsersOptions model
				engineID := "testString"
				deleteEngineUsersOptionsModel := watsonxDataService.NewDeleteEngineUsersOptions(engineID)
				deleteEngineUsersOptionsModel.SetEngineID("testString")
				deleteEngineUsersOptionsModel.SetGroups([]string{"testString"})
				deleteEngineUsersOptionsModel.SetUsers([]string{"testString"})
				deleteEngineUsersOptionsModel.SetLhInstanceID("testString")
				deleteEngineUsersOptionsModel.SetAuthInstanceID("testString")
				deleteEngineUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEngineUsersOptionsModel).ToNot(BeNil())
				Expect(deleteEngineUsersOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineUsersOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(deleteEngineUsersOptionsModel.Users).To(Equal([]string{"testString"}))
				Expect(deleteEngineUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteMetastoreUsersOptions successfully`, func() {
				// Construct an instance of the DeleteMetastoreUsersOptions model
				metastoreName := "testString"
				deleteMetastoreUsersOptionsModel := watsonxDataService.NewDeleteMetastoreUsersOptions(metastoreName)
				deleteMetastoreUsersOptionsModel.SetMetastoreName("testString")
				deleteMetastoreUsersOptionsModel.SetGroups([]string{"testString"})
				deleteMetastoreUsersOptionsModel.SetUsers([]string{"testString"})
				deleteMetastoreUsersOptionsModel.SetLhInstanceID("testString")
				deleteMetastoreUsersOptionsModel.SetAuthInstanceID("testString")
				deleteMetastoreUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteMetastoreUsersOptionsModel).ToNot(BeNil())
				Expect(deleteMetastoreUsersOptionsModel.MetastoreName).To(Equal(core.StringPtr("testString")))
				Expect(deleteMetastoreUsersOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(deleteMetastoreUsersOptionsModel.Users).To(Equal([]string{"testString"}))
				Expect(deleteMetastoreUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMetastoreUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteMetastoreUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteQueryOptions successfully`, func() {
				// Construct an instance of the DeleteQueryOptions model
				queryName := "testString"
				deleteQueryOptionsModel := watsonxDataService.NewDeleteQueryOptions(queryName)
				deleteQueryOptionsModel.SetQueryName("testString")
				deleteQueryOptionsModel.SetAuthInstanceID("testString")
				deleteQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteQueryOptionsModel).ToNot(BeNil())
				Expect(deleteQueryOptionsModel.QueryName).To(Equal(core.StringPtr("testString")))
				Expect(deleteQueryOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSchemaOptions successfully`, func() {
				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsCatalogName := "sampleCatalog"
				deleteSchemaOptionsEngineID := "sampleEngine123"
				deleteSchemaOptionsSchemaName := "new_schema"
				deleteSchemaOptionsModel := watsonxDataService.NewDeleteSchemaOptions(deleteSchemaOptionsCatalogName, deleteSchemaOptionsEngineID, deleteSchemaOptionsSchemaName)
				deleteSchemaOptionsModel.SetCatalogName("sampleCatalog")
				deleteSchemaOptionsModel.SetEngineID("sampleEngine123")
				deleteSchemaOptionsModel.SetSchemaName("new_schema")
				deleteSchemaOptionsModel.SetAuthInstanceID("testString")
				deleteSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSchemaOptionsModel).ToNot(BeNil())
				Expect(deleteSchemaOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(deleteSchemaOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(deleteSchemaOptionsModel.SchemaName).To(Equal(core.StringPtr("new_schema")))
				Expect(deleteSchemaOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTableOptions successfully`, func() {
				// Construct an instance of the DeleteTableBodyDeleteTablesItems model
				deleteTableBodyDeleteTablesItemsModel := new(watsonxdatav1.DeleteTableBodyDeleteTablesItems)
				Expect(deleteTableBodyDeleteTablesItemsModel).ToNot(BeNil())
				deleteTableBodyDeleteTablesItemsModel.CatalogName = core.StringPtr("sampleCatalog")
				deleteTableBodyDeleteTablesItemsModel.SchemaName = core.StringPtr("new_schema")
				deleteTableBodyDeleteTablesItemsModel.TableName = core.StringPtr("new_table")
				Expect(deleteTableBodyDeleteTablesItemsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(deleteTableBodyDeleteTablesItemsModel.SchemaName).To(Equal(core.StringPtr("new_schema")))
				Expect(deleteTableBodyDeleteTablesItemsModel.TableName).To(Equal(core.StringPtr("new_table")))

				// Construct an instance of the DeleteTableOptions model
				deleteTableOptionsDeleteTables := []watsonxdatav1.DeleteTableBodyDeleteTablesItems{}
				deleteTableOptionsEngineID := "sampleEngine123"
				deleteTableOptionsModel := watsonxDataService.NewDeleteTableOptions(deleteTableOptionsDeleteTables, deleteTableOptionsEngineID)
				deleteTableOptionsModel.SetDeleteTables([]watsonxdatav1.DeleteTableBodyDeleteTablesItems{*deleteTableBodyDeleteTablesItemsModel})
				deleteTableOptionsModel.SetEngineID("sampleEngine123")
				deleteTableOptionsModel.SetAuthInstanceID("testString")
				deleteTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTableOptionsModel).ToNot(BeNil())
				Expect(deleteTableOptionsModel.DeleteTables).To(Equal([]watsonxdatav1.DeleteTableBodyDeleteTablesItems{*deleteTableBodyDeleteTablesItemsModel}))
				Expect(deleteTableOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(deleteTableOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEvaluateOptions successfully`, func() {
				// Construct an instance of the ResourcesMetadata model
				resourcesMetadataModel := new(watsonxdatav1.ResourcesMetadata)
				Expect(resourcesMetadataModel).ToNot(BeNil())
				resourcesMetadataModel.Action = core.StringPtr("testString")
				resourcesMetadataModel.ResourceName = core.StringPtr("testString")
				resourcesMetadataModel.ResourceType = core.StringPtr("engine")
				Expect(resourcesMetadataModel.Action).To(Equal(core.StringPtr("testString")))
				Expect(resourcesMetadataModel.ResourceName).To(Equal(core.StringPtr("testString")))
				Expect(resourcesMetadataModel.ResourceType).To(Equal(core.StringPtr("engine")))

				// Construct an instance of the EvaluateOptions model
				evaluateOptionsModel := watsonxDataService.NewEvaluateOptions()
				evaluateOptionsModel.SetResources([]watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel})
				evaluateOptionsModel.SetLhInstanceID("testString")
				evaluateOptionsModel.SetAuthInstanceID("testString")
				evaluateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(evaluateOptionsModel).ToNot(BeNil())
				Expect(evaluateOptionsModel.Resources).To(Equal([]watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel}))
				Expect(evaluateOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(evaluateOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(evaluateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExplainAnalyzeStatementOptions successfully`, func() {
				// Construct an instance of the ExplainAnalyzeStatementOptions model
				explainAnalyzeStatementOptionsCatalogName := "sampleCatalog"
				explainAnalyzeStatementOptionsEngineID := "sampleEngine1"
				explainAnalyzeStatementOptionsSchemaName := "new_schema"
				explainAnalyzeStatementOptionsStatement := "show schemas in catalog"
				explainAnalyzeStatementOptionsModel := watsonxDataService.NewExplainAnalyzeStatementOptions(explainAnalyzeStatementOptionsCatalogName, explainAnalyzeStatementOptionsEngineID, explainAnalyzeStatementOptionsSchemaName, explainAnalyzeStatementOptionsStatement)
				explainAnalyzeStatementOptionsModel.SetCatalogName("sampleCatalog")
				explainAnalyzeStatementOptionsModel.SetEngineID("sampleEngine1")
				explainAnalyzeStatementOptionsModel.SetSchemaName("new_schema")
				explainAnalyzeStatementOptionsModel.SetStatement("show schemas in catalog")
				explainAnalyzeStatementOptionsModel.SetVerbose(true)
				explainAnalyzeStatementOptionsModel.SetAuthInstanceID("testString")
				explainAnalyzeStatementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(explainAnalyzeStatementOptionsModel).ToNot(BeNil())
				Expect(explainAnalyzeStatementOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(explainAnalyzeStatementOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine1")))
				Expect(explainAnalyzeStatementOptionsModel.SchemaName).To(Equal(core.StringPtr("new_schema")))
				Expect(explainAnalyzeStatementOptionsModel.Statement).To(Equal(core.StringPtr("show schemas in catalog")))
				Expect(explainAnalyzeStatementOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(explainAnalyzeStatementOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(explainAnalyzeStatementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExplainStatementOptions successfully`, func() {
				// Construct an instance of the ExplainStatementOptions model
				explainStatementOptionsEngineID := "eng_id"
				explainStatementOptionsStatement := "show schemas"
				explainStatementOptionsModel := watsonxDataService.NewExplainStatementOptions(explainStatementOptionsEngineID, explainStatementOptionsStatement)
				explainStatementOptionsModel.SetEngineID("eng_id")
				explainStatementOptionsModel.SetStatement("show schemas")
				explainStatementOptionsModel.SetCatalogName("sampleCatalog")
				explainStatementOptionsModel.SetFormat("json")
				explainStatementOptionsModel.SetSchemaName("new_schema")
				explainStatementOptionsModel.SetType("io")
				explainStatementOptionsModel.SetAuthInstanceID("testString")
				explainStatementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(explainStatementOptionsModel).ToNot(BeNil())
				Expect(explainStatementOptionsModel.EngineID).To(Equal(core.StringPtr("eng_id")))
				Expect(explainStatementOptionsModel.Statement).To(Equal(core.StringPtr("show schemas")))
				Expect(explainStatementOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(explainStatementOptionsModel.Format).To(Equal(core.StringPtr("json")))
				Expect(explainStatementOptionsModel.SchemaName).To(Equal(core.StringPtr("new_schema")))
				Expect(explainStatementOptionsModel.Type).To(Equal(core.StringPtr("io")))
				Expect(explainStatementOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(explainStatementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketObjectsOptions successfully`, func() {
				// Construct an instance of the GetBucketObjectsOptions model
				bucketID := "testString"
				getBucketObjectsOptionsModel := watsonxDataService.NewGetBucketObjectsOptions(bucketID)
				getBucketObjectsOptionsModel.SetBucketID("testString")
				getBucketObjectsOptionsModel.SetAuthInstanceID("testString")
				getBucketObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketObjectsOptionsModel).ToNot(BeNil())
				Expect(getBucketObjectsOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketObjectsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketUsersOptions successfully`, func() {
				// Construct an instance of the GetBucketUsersOptions model
				bucketID := "testString"
				getBucketUsersOptionsModel := watsonxDataService.NewGetBucketUsersOptions(bucketID)
				getBucketUsersOptionsModel.SetBucketID("testString")
				getBucketUsersOptionsModel.SetLhInstanceID("testString")
				getBucketUsersOptionsModel.SetAuthInstanceID("testString")
				getBucketUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketUsersOptionsModel).ToNot(BeNil())
				Expect(getBucketUsersOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketsOptions successfully`, func() {
				// Construct an instance of the GetBucketsOptions model
				getBucketsOptionsModel := watsonxDataService.NewGetBucketsOptions()
				getBucketsOptionsModel.SetAuthInstanceID("testString")
				getBucketsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketsOptionsModel).ToNot(BeNil())
				Expect(getBucketsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogUsersOptions successfully`, func() {
				// Construct an instance of the GetCatalogUsersOptions model
				catalogName := "testString"
				getCatalogUsersOptionsModel := watsonxDataService.NewGetCatalogUsersOptions(catalogName)
				getCatalogUsersOptionsModel.SetCatalogName("testString")
				getCatalogUsersOptionsModel.SetLhInstanceID("testString")
				getCatalogUsersOptionsModel.SetAuthInstanceID("testString")
				getCatalogUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogUsersOptionsModel).ToNot(BeNil())
				Expect(getCatalogUsersOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataPolicyOptions successfully`, func() {
				// Construct an instance of the GetDataPolicyOptions model
				policyName := "testString"
				getDataPolicyOptionsModel := watsonxDataService.NewGetDataPolicyOptions(policyName)
				getDataPolicyOptionsModel.SetPolicyName("testString")
				getDataPolicyOptionsModel.SetLhInstanceID("testString")
				getDataPolicyOptionsModel.SetAuthInstanceID("testString")
				getDataPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataPolicyOptionsModel).ToNot(BeNil())
				Expect(getDataPolicyOptionsModel.PolicyName).To(Equal(core.StringPtr("testString")))
				Expect(getDataPolicyOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDataPolicyOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDataPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDatabasesOptions successfully`, func() {
				// Construct an instance of the GetDatabasesOptions model
				getDatabasesOptionsModel := watsonxDataService.NewGetDatabasesOptions()
				getDatabasesOptionsModel.SetAccept("testString")
				getDatabasesOptionsModel.SetAuthInstanceID("testString")
				getDatabasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDatabasesOptionsModel).ToNot(BeNil())
				Expect(getDatabasesOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getDatabasesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDatabasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDbConnUsersOptions successfully`, func() {
				// Construct an instance of the GetDbConnUsersOptions model
				databaseID := "testString"
				getDbConnUsersOptionsModel := watsonxDataService.NewGetDbConnUsersOptions(databaseID)
				getDbConnUsersOptionsModel.SetDatabaseID("testString")
				getDbConnUsersOptionsModel.SetLhInstanceID("testString")
				getDbConnUsersOptionsModel.SetAuthInstanceID("testString")
				getDbConnUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDbConnUsersOptionsModel).ToNot(BeNil())
				Expect(getDbConnUsersOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(getDbConnUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDbConnUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDbConnUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDefaultPoliciesOptions successfully`, func() {
				// Construct an instance of the GetDefaultPoliciesOptions model
				getDefaultPoliciesOptionsModel := watsonxDataService.NewGetDefaultPoliciesOptions()
				getDefaultPoliciesOptionsModel.SetLhInstanceID("testString")
				getDefaultPoliciesOptionsModel.SetAuthInstanceID("testString")
				getDefaultPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDefaultPoliciesOptionsModel).ToNot(BeNil())
				Expect(getDefaultPoliciesOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDefaultPoliciesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDefaultPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentsOptions successfully`, func() {
				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := watsonxDataService.NewGetDeploymentsOptions()
				getDeploymentsOptionsModel.SetAccept("testString")
				getDeploymentsOptionsModel.SetAuthInstanceID("testString")
				getDeploymentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentsOptionsModel).ToNot(BeNil())
				Expect(getDeploymentsOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEngineUsersOptions successfully`, func() {
				// Construct an instance of the GetEngineUsersOptions model
				engineID := "testString"
				getEngineUsersOptionsModel := watsonxDataService.NewGetEngineUsersOptions(engineID)
				getEngineUsersOptionsModel.SetEngineID("testString")
				getEngineUsersOptionsModel.SetLhInstanceID("testString")
				getEngineUsersOptionsModel.SetAuthInstanceID("testString")
				getEngineUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEngineUsersOptionsModel).ToNot(BeNil())
				Expect(getEngineUsersOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getEngineUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getEngineUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getEngineUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnginesOptions successfully`, func() {
				// Construct an instance of the GetEnginesOptions model
				getEnginesOptionsModel := watsonxDataService.NewGetEnginesOptions()
				getEnginesOptionsModel.SetAuthInstanceID("testString")
				getEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnginesOptionsModel).ToNot(BeNil())
				Expect(getEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHMSOptions successfully`, func() {
				// Construct an instance of the GetHMSOptions model
				getHmsOptionsModel := watsonxDataService.NewGetHMSOptions()
				getHmsOptionsModel.SetAccept("testString")
				getHmsOptionsModel.SetAuthInstanceID("testString")
				getHmsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHmsOptionsModel).ToNot(BeNil())
				Expect(getHmsOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getHmsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getHmsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetastoreUsersOptions successfully`, func() {
				// Construct an instance of the GetMetastoreUsersOptions model
				metastoreName := "testString"
				getMetastoreUsersOptionsModel := watsonxDataService.NewGetMetastoreUsersOptions(metastoreName)
				getMetastoreUsersOptionsModel.SetMetastoreName("testString")
				getMetastoreUsersOptionsModel.SetLhInstanceID("testString")
				getMetastoreUsersOptionsModel.SetAuthInstanceID("testString")
				getMetastoreUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetastoreUsersOptionsModel).ToNot(BeNil())
				Expect(getMetastoreUsersOptionsModel.MetastoreName).To(Equal(core.StringPtr("testString")))
				Expect(getMetastoreUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getMetastoreUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getMetastoreUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetastoresOptions successfully`, func() {
				// Construct an instance of the GetMetastoresOptions model
				getMetastoresOptionsModel := watsonxDataService.NewGetMetastoresOptions()
				getMetastoresOptionsModel.SetAuthInstanceID("testString")
				getMetastoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetastoresOptionsModel).ToNot(BeNil())
				Expect(getMetastoresOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getMetastoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPoliciesListOptions successfully`, func() {
				// Construct an instance of the GetPoliciesListOptions model
				getPoliciesListOptionsModel := watsonxDataService.NewGetPoliciesListOptions()
				getPoliciesListOptionsModel.SetLhInstanceID("testString")
				getPoliciesListOptionsModel.SetAuthInstanceID("testString")
				getPoliciesListOptionsModel.SetCatalogList([]string{"testString"})
				getPoliciesListOptionsModel.SetEngineList([]string{"testString"})
				getPoliciesListOptionsModel.SetDataPoliciesList([]string{"testString"})
				getPoliciesListOptionsModel.SetIncludeDataPolicies(true)
				getPoliciesListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPoliciesListOptionsModel).ToNot(BeNil())
				Expect(getPoliciesListOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPoliciesListOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPoliciesListOptionsModel.CatalogList).To(Equal([]string{"testString"}))
				Expect(getPoliciesListOptionsModel.EngineList).To(Equal([]string{"testString"}))
				Expect(getPoliciesListOptionsModel.DataPoliciesList).To(Equal([]string{"testString"}))
				Expect(getPoliciesListOptionsModel.IncludeDataPolicies).To(Equal(core.BoolPtr(true)))
				Expect(getPoliciesListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyVersionOptions successfully`, func() {
				// Construct an instance of the GetPolicyVersionOptions model
				getPolicyVersionOptionsModel := watsonxDataService.NewGetPolicyVersionOptions()
				getPolicyVersionOptionsModel.SetLhInstanceID("testString")
				getPolicyVersionOptionsModel.SetAuthInstanceID("testString")
				getPolicyVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyVersionOptionsModel).ToNot(BeNil())
				Expect(getPolicyVersionOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyVersionOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetQueriesOptions successfully`, func() {
				// Construct an instance of the GetQueriesOptions model
				getQueriesOptionsModel := watsonxDataService.NewGetQueriesOptions()
				getQueriesOptionsModel.SetAuthInstanceID("testString")
				getQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getQueriesOptionsModel).ToNot(BeNil())
				Expect(getQueriesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchemasOptions successfully`, func() {
				// Construct an instance of the GetSchemasOptions model
				engineID := "testString"
				catalogName := "testString"
				getSchemasOptionsModel := watsonxDataService.NewGetSchemasOptions(engineID, catalogName)
				getSchemasOptionsModel.SetEngineID("testString")
				getSchemasOptionsModel.SetCatalogName("testString")
				getSchemasOptionsModel.SetAuthInstanceID("testString")
				getSchemasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchemasOptionsModel).ToNot(BeNil())
				Expect(getSchemasOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getSchemasOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(getSchemasOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getSchemasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTableSnapshotsOptions successfully`, func() {
				// Construct an instance of the GetTableSnapshotsOptions model
				engineID := "testString"
				catalogName := "testString"
				schemaName := "testString"
				tableName := "testString"
				getTableSnapshotsOptionsModel := watsonxDataService.NewGetTableSnapshotsOptions(engineID, catalogName, schemaName, tableName)
				getTableSnapshotsOptionsModel.SetEngineID("testString")
				getTableSnapshotsOptionsModel.SetCatalogName("testString")
				getTableSnapshotsOptionsModel.SetSchemaName("testString")
				getTableSnapshotsOptionsModel.SetTableName("testString")
				getTableSnapshotsOptionsModel.SetAuthInstanceID("testString")
				getTableSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTableSnapshotsOptionsModel).ToNot(BeNil())
				Expect(getTableSnapshotsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getTableSnapshotsOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(getTableSnapshotsOptionsModel.SchemaName).To(Equal(core.StringPtr("testString")))
				Expect(getTableSnapshotsOptionsModel.TableName).To(Equal(core.StringPtr("testString")))
				Expect(getTableSnapshotsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getTableSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTablesOptions successfully`, func() {
				// Construct an instance of the GetTablesOptions model
				engineID := "testString"
				catalogName := "testString"
				schemaName := "testString"
				getTablesOptionsModel := watsonxDataService.NewGetTablesOptions(engineID, catalogName, schemaName)
				getTablesOptionsModel.SetEngineID("testString")
				getTablesOptionsModel.SetCatalogName("testString")
				getTablesOptionsModel.SetSchemaName("testString")
				getTablesOptionsModel.SetAuthInstanceID("testString")
				getTablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTablesOptionsModel).ToNot(BeNil())
				Expect(getTablesOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getTablesOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(getTablesOptionsModel.SchemaName).To(Equal(core.StringPtr("testString")))
				Expect(getTablesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getTablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataPoliciesOptions successfully`, func() {
				// Construct an instance of the ListDataPoliciesOptions model
				listDataPoliciesOptionsModel := watsonxDataService.NewListDataPoliciesOptions()
				listDataPoliciesOptionsModel.SetLhInstanceID("testString")
				listDataPoliciesOptionsModel.SetAuthInstanceID("testString")
				listDataPoliciesOptionsModel.SetCatalogName("testString")
				listDataPoliciesOptionsModel.SetStatus("testString")
				listDataPoliciesOptionsModel.SetIncludeMetadata(true)
				listDataPoliciesOptionsModel.SetIncludeRules(true)
				listDataPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataPoliciesOptionsModel).ToNot(BeNil())
				Expect(listDataPoliciesOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDataPoliciesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDataPoliciesOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(listDataPoliciesOptionsModel.Status).To(Equal(core.StringPtr("testString")))
				Expect(listDataPoliciesOptionsModel.IncludeMetadata).To(Equal(core.BoolPtr(true)))
				Expect(listDataPoliciesOptionsModel.IncludeRules).To(Equal(core.BoolPtr(true)))
				Expect(listDataPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewParseCsvOptions successfully`, func() {
				// Construct an instance of the ParseCsvOptions model
				engine := "testString"
				parseFile := "testString"
				fileType := "testString"
				parseCsvOptionsModel := watsonxDataService.NewParseCsvOptions(engine, parseFile, fileType)
				parseCsvOptionsModel.SetEngine("testString")
				parseCsvOptionsModel.SetParseFile("testString")
				parseCsvOptionsModel.SetFileType("testString")
				parseCsvOptionsModel.SetAccept("testString")
				parseCsvOptionsModel.SetAuthInstanceID("testString")
				parseCsvOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(parseCsvOptionsModel).ToNot(BeNil())
				Expect(parseCsvOptionsModel.Engine).To(Equal(core.StringPtr("testString")))
				Expect(parseCsvOptionsModel.ParseFile).To(Equal(core.StringPtr("testString")))
				Expect(parseCsvOptionsModel.FileType).To(Equal(core.StringPtr("testString")))
				Expect(parseCsvOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(parseCsvOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(parseCsvOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPauseEngineOptions successfully`, func() {
				// Construct an instance of the PauseEngineOptions model
				pauseEngineOptionsEngineID := "testString"
				pauseEngineOptionsModel := watsonxDataService.NewPauseEngineOptions(pauseEngineOptionsEngineID)
				pauseEngineOptionsModel.SetEngineID("testString")
				pauseEngineOptionsModel.SetCreatedBy("testString")
				pauseEngineOptionsModel.SetAuthInstanceID("testString")
				pauseEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(pauseEngineOptionsModel).ToNot(BeNil())
				Expect(pauseEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(pauseEngineOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(pauseEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(pauseEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostQueryOptions successfully`, func() {
				// Construct an instance of the PostQueryOptions model
				engine := "testString"
				catalog := "testString"
				schema := "testString"
				sqlQuery := "testString"
				postQueryOptionsModel := watsonxDataService.NewPostQueryOptions(engine, catalog, schema, sqlQuery)
				postQueryOptionsModel.SetEngine("testString")
				postQueryOptionsModel.SetCatalog("testString")
				postQueryOptionsModel.SetSchema("testString")
				postQueryOptionsModel.SetSqlQuery("testString")
				postQueryOptionsModel.SetAccept("testString")
				postQueryOptionsModel.SetAuthInstanceID("testString")
				postQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postQueryOptionsModel).ToNot(BeNil())
				Expect(postQueryOptionsModel.Engine).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.Catalog).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.Schema).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.SqlQuery).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(postQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRegisterBucketOptions successfully`, func() {
				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav1.BucketDetails)
				Expect(bucketDetailsModel).ToNot(BeNil())
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("<secret_key>")
				Expect(bucketDetailsModel.AccessKey).To(Equal(core.StringPtr("<access_key>")))
				Expect(bucketDetailsModel.BucketName).To(Equal(core.StringPtr("sample-bucket")))
				Expect(bucketDetailsModel.Endpoint).To(Equal(core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")))
				Expect(bucketDetailsModel.SecretKey).To(Equal(core.StringPtr("<secret_key>")))

				// Construct an instance of the RegisterBucketOptions model
				var registerBucketOptionsBucketDetails *watsonxdatav1.BucketDetails = nil
				registerBucketOptionsDescription := "COS bucket for customer data"
				registerBucketOptionsTableType := "iceberg"
				registerBucketOptionsBucketType := "ibm_cos"
				registerBucketOptionsCatalogName := "sampleCatalog"
				registerBucketOptionsManagedBy := "ibm"
				registerBucketOptionsModel := watsonxDataService.NewRegisterBucketOptions(registerBucketOptionsBucketDetails, registerBucketOptionsDescription, registerBucketOptionsTableType, registerBucketOptionsBucketType, registerBucketOptionsCatalogName, registerBucketOptionsManagedBy)
				registerBucketOptionsModel.SetBucketDetails(bucketDetailsModel)
				registerBucketOptionsModel.SetDescription("COS bucket for customer data")
				registerBucketOptionsModel.SetTableType("iceberg")
				registerBucketOptionsModel.SetBucketType("ibm_cos")
				registerBucketOptionsModel.SetCatalogName("sampleCatalog")
				registerBucketOptionsModel.SetManagedBy("ibm")
				registerBucketOptionsModel.SetBucketDisplayName("sample-bucket-displayname")
				registerBucketOptionsModel.SetBucketTags([]string{"read customer data", "write customer data'"})
				registerBucketOptionsModel.SetCatalogTags([]string{"catalog_tag_1", "catalog_tag_2"})
				registerBucketOptionsModel.SetThriftURI("thrift://samplehost-metastore:4354")
				registerBucketOptionsModel.SetAuthInstanceID("testString")
				registerBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(registerBucketOptionsModel).ToNot(BeNil())
				Expect(registerBucketOptionsModel.BucketDetails).To(Equal(bucketDetailsModel))
				Expect(registerBucketOptionsModel.Description).To(Equal(core.StringPtr("COS bucket for customer data")))
				Expect(registerBucketOptionsModel.TableType).To(Equal(core.StringPtr("iceberg")))
				Expect(registerBucketOptionsModel.BucketType).To(Equal(core.StringPtr("ibm_cos")))
				Expect(registerBucketOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(registerBucketOptionsModel.ManagedBy).To(Equal(core.StringPtr("ibm")))
				Expect(registerBucketOptionsModel.BucketDisplayName).To(Equal(core.StringPtr("sample-bucket-displayname")))
				Expect(registerBucketOptionsModel.BucketTags).To(Equal([]string{"read customer data", "write customer data'"}))
				Expect(registerBucketOptionsModel.CatalogTags).To(Equal([]string{"catalog_tag_1", "catalog_tag_2"}))
				Expect(registerBucketOptionsModel.ThriftURI).To(Equal(core.StringPtr("thrift://samplehost-metastore:4354")))
				Expect(registerBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(registerBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveCatalogFromEngineOptions successfully`, func() {
				// Construct an instance of the RemoveCatalogFromEngineOptions model
				removeCatalogFromEngineOptionsCatalogName := "testString"
				removeCatalogFromEngineOptionsEngineID := "testString"
				removeCatalogFromEngineOptionsModel := watsonxDataService.NewRemoveCatalogFromEngineOptions(removeCatalogFromEngineOptionsCatalogName, removeCatalogFromEngineOptionsEngineID)
				removeCatalogFromEngineOptionsModel.SetCatalogName("testString")
				removeCatalogFromEngineOptionsModel.SetEngineID("testString")
				removeCatalogFromEngineOptionsModel.SetAccept("testString")
				removeCatalogFromEngineOptionsModel.SetCreatedBy("testString")
				removeCatalogFromEngineOptionsModel.SetAuthInstanceID("testString")
				removeCatalogFromEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeCatalogFromEngineOptionsModel).ToNot(BeNil())
				Expect(removeCatalogFromEngineOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(removeCatalogFromEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(removeCatalogFromEngineOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(removeCatalogFromEngineOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(removeCatalogFromEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(removeCatalogFromEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceDataPolicyOptions successfully`, func() {
				// Construct an instance of the RuleGrantee model
				ruleGranteeModel := new(watsonxdatav1.RuleGrantee)
				Expect(ruleGranteeModel).ToNot(BeNil())
				ruleGranteeModel.Value = core.StringPtr("testString")
				ruleGranteeModel.Key = core.StringPtr("user_name")
				ruleGranteeModel.Type = core.StringPtr("user_identity")
				Expect(ruleGranteeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(ruleGranteeModel.Key).To(Equal(core.StringPtr("user_name")))
				Expect(ruleGranteeModel.Type).To(Equal(core.StringPtr("user_identity")))

				// Construct an instance of the Rule model
				ruleModel := new(watsonxdatav1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.Actions = []string{"all"}
				ruleModel.Effect = core.StringPtr("allow")
				ruleModel.Grantee = ruleGranteeModel
				Expect(ruleModel.Actions).To(Equal([]string{"all"}))
				Expect(ruleModel.Effect).To(Equal(core.StringPtr("allow")))
				Expect(ruleModel.Grantee).To(Equal(ruleGranteeModel))

				// Construct an instance of the ReplaceDataPolicyOptions model
				policyName := "testString"
				replaceDataPolicyOptionsCatalogName := "testString"
				replaceDataPolicyOptionsDataArtifact := "schema1/table1/(column1|column2)"
				replaceDataPolicyOptionsRules := []watsonxdatav1.Rule{}
				replaceDataPolicyOptionsModel := watsonxDataService.NewReplaceDataPolicyOptions(policyName, replaceDataPolicyOptionsCatalogName, replaceDataPolicyOptionsDataArtifact, replaceDataPolicyOptionsRules)
				replaceDataPolicyOptionsModel.SetPolicyName("testString")
				replaceDataPolicyOptionsModel.SetCatalogName("testString")
				replaceDataPolicyOptionsModel.SetDataArtifact("schema1/table1/(column1|column2)")
				replaceDataPolicyOptionsModel.SetRules([]watsonxdatav1.Rule{*ruleModel})
				replaceDataPolicyOptionsModel.SetDescription("testString")
				replaceDataPolicyOptionsModel.SetStatus("active")
				replaceDataPolicyOptionsModel.SetLhInstanceID("testString")
				replaceDataPolicyOptionsModel.SetAuthInstanceID("testString")
				replaceDataPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceDataPolicyOptionsModel).ToNot(BeNil())
				Expect(replaceDataPolicyOptionsModel.PolicyName).To(Equal(core.StringPtr("testString")))
				Expect(replaceDataPolicyOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(replaceDataPolicyOptionsModel.DataArtifact).To(Equal(core.StringPtr("schema1/table1/(column1|column2)")))
				Expect(replaceDataPolicyOptionsModel.Rules).To(Equal([]watsonxdatav1.Rule{*ruleModel}))
				Expect(replaceDataPolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceDataPolicyOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(replaceDataPolicyOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceDataPolicyOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceDataPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResumeEngineOptions successfully`, func() {
				// Construct an instance of the ResumeEngineOptions model
				resumeEngineOptionsEngineID := "eng_id"
				resumeEngineOptionsModel := watsonxDataService.NewResumeEngineOptions(resumeEngineOptionsEngineID)
				resumeEngineOptionsModel.SetEngineID("eng_id")
				resumeEngineOptionsModel.SetCreatedBy("<username>@<domain>.com")
				resumeEngineOptionsModel.SetAuthInstanceID("testString")
				resumeEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resumeEngineOptionsModel).ToNot(BeNil())
				Expect(resumeEngineOptionsModel.EngineID).To(Equal(core.StringPtr("eng_id")))
				Expect(resumeEngineOptionsModel.CreatedBy).To(Equal(core.StringPtr("<username>@<domain>.com")))
				Expect(resumeEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(resumeEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRollbackSnapshotOptions successfully`, func() {
				// Construct an instance of the RollbackSnapshotOptions model
				engineID := "testString"
				catalogName := "testString"
				schemaName := "testString"
				rollbackSnapshotOptionsSnapshotID := "2332342122211222"
				rollbackSnapshotOptionsTableName := "new_table"
				rollbackSnapshotOptionsModel := watsonxDataService.NewRollbackSnapshotOptions(engineID, catalogName, schemaName, rollbackSnapshotOptionsSnapshotID, rollbackSnapshotOptionsTableName)
				rollbackSnapshotOptionsModel.SetEngineID("testString")
				rollbackSnapshotOptionsModel.SetCatalogName("testString")
				rollbackSnapshotOptionsModel.SetSchemaName("testString")
				rollbackSnapshotOptionsModel.SetSnapshotID("2332342122211222")
				rollbackSnapshotOptionsModel.SetTableName("new_table")
				rollbackSnapshotOptionsModel.SetAuthInstanceID("testString")
				rollbackSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(rollbackSnapshotOptionsModel).ToNot(BeNil())
				Expect(rollbackSnapshotOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(rollbackSnapshotOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(rollbackSnapshotOptionsModel.SchemaName).To(Equal(core.StringPtr("testString")))
				Expect(rollbackSnapshotOptionsModel.SnapshotID).To(Equal(core.StringPtr("2332342122211222")))
				Expect(rollbackSnapshotOptionsModel.TableName).To(Equal(core.StringPtr("new_table")))
				Expect(rollbackSnapshotOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(rollbackSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveQueryOptions successfully`, func() {
				// Construct an instance of the SaveQueryOptions model
				queryName := "testString"
				saveQueryOptionsCreatedBy := "<username>@<domain>.com"
				saveQueryOptionsDescription := "query to get expense data"
				saveQueryOptionsQueryString := "select expenses from expenditure"
				saveQueryOptionsModel := watsonxDataService.NewSaveQueryOptions(queryName, saveQueryOptionsCreatedBy, saveQueryOptionsDescription, saveQueryOptionsQueryString)
				saveQueryOptionsModel.SetQueryName("testString")
				saveQueryOptionsModel.SetCreatedBy("<username>@<domain>.com")
				saveQueryOptionsModel.SetDescription("query to get expense data")
				saveQueryOptionsModel.SetQueryString("select expenses from expenditure")
				saveQueryOptionsModel.SetCreatedOn("1608437933")
				saveQueryOptionsModel.SetEngineID("sampleEngine123")
				saveQueryOptionsModel.SetAuthInstanceID("testString")
				saveQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveQueryOptionsModel).ToNot(BeNil())
				Expect(saveQueryOptionsModel.QueryName).To(Equal(core.StringPtr("testString")))
				Expect(saveQueryOptionsModel.CreatedBy).To(Equal(core.StringPtr("<username>@<domain>.com")))
				Expect(saveQueryOptionsModel.Description).To(Equal(core.StringPtr("query to get expense data")))
				Expect(saveQueryOptionsModel.QueryString).To(Equal(core.StringPtr("select expenses from expenditure")))
				Expect(saveQueryOptionsModel.CreatedOn).To(Equal(core.StringPtr("1608437933")))
				Expect(saveQueryOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(saveQueryOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(saveQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTestLHConsoleOptions successfully`, func() {
				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := watsonxDataService.NewTestLHConsoleOptions()
				testLhConsoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testLhConsoleOptionsModel).ToNot(BeNil())
				Expect(testLhConsoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnregisterBucketOptions successfully`, func() {
				// Construct an instance of the UnregisterBucketOptions model
				unregisterBucketOptionsBucketID := "bucket_id"
				unregisterBucketOptionsModel := watsonxDataService.NewUnregisterBucketOptions(unregisterBucketOptionsBucketID)
				unregisterBucketOptionsModel.SetBucketID("bucket_id")
				unregisterBucketOptionsModel.SetAuthInstanceID("testString")
				unregisterBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unregisterBucketOptionsModel).ToNot(BeNil())
				Expect(unregisterBucketOptionsModel.BucketID).To(Equal(core.StringPtr("bucket_id")))
				Expect(unregisterBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(unregisterBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBucketOptions successfully`, func() {
				// Construct an instance of the UpdateBucketOptions model
				updateBucketOptionsBucketID := "samplebucket123"
				updateBucketOptionsModel := watsonxDataService.NewUpdateBucketOptions(updateBucketOptionsBucketID)
				updateBucketOptionsModel.SetBucketID("samplebucket123")
				updateBucketOptionsModel.SetAccessKey("<access_key>")
				updateBucketOptionsModel.SetBucketDisplayName("sample-bucket-displayname")
				updateBucketOptionsModel.SetDescription("COS bucket for customer data")
				updateBucketOptionsModel.SetSecretKey("<secret_key>")
				updateBucketOptionsModel.SetTags([]string{"testbucket", "userbucket"})
				updateBucketOptionsModel.SetAuthInstanceID("testString")
				updateBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBucketOptionsModel).ToNot(BeNil())
				Expect(updateBucketOptionsModel.BucketID).To(Equal(core.StringPtr("samplebucket123")))
				Expect(updateBucketOptionsModel.AccessKey).To(Equal(core.StringPtr("<access_key>")))
				Expect(updateBucketOptionsModel.BucketDisplayName).To(Equal(core.StringPtr("sample-bucket-displayname")))
				Expect(updateBucketOptionsModel.Description).To(Equal(core.StringPtr("COS bucket for customer data")))
				Expect(updateBucketOptionsModel.SecretKey).To(Equal(core.StringPtr("<secret_key>")))
				Expect(updateBucketOptionsModel.Tags).To(Equal([]string{"testbucket", "userbucket"}))
				Expect(updateBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBucketUsersOptions successfully`, func() {
				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				Expect(bucketDbConnGroupsMetadataModel).ToNot(BeNil())
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				Expect(bucketDbConnUsersMetadataModel).ToNot(BeNil())
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the UpdateBucketUsersOptions model
				bucketID := "testString"
				updateBucketUsersOptionsModel := watsonxDataService.NewUpdateBucketUsersOptions(bucketID)
				updateBucketUsersOptionsModel.SetBucketID("testString")
				updateBucketUsersOptionsModel.SetGroups([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel})
				updateBucketUsersOptionsModel.SetUsers([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel})
				updateBucketUsersOptionsModel.SetLhInstanceID("testString")
				updateBucketUsersOptionsModel.SetAuthInstanceID("testString")
				updateBucketUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBucketUsersOptionsModel).ToNot(BeNil())
				Expect(updateBucketUsersOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}))
				Expect(updateBucketUsersOptionsModel.Users).To(Equal([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}))
				Expect(updateBucketUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogUsersOptions successfully`, func() {
				// Construct an instance of the CatalogGroupsMetadata model
				catalogGroupsMetadataModel := new(watsonxdatav1.CatalogGroupsMetadata)
				Expect(catalogGroupsMetadataModel).ToNot(BeNil())
				catalogGroupsMetadataModel.GroupID = core.StringPtr("testString")
				catalogGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(catalogGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(catalogGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the CatalogUsersMetadata model
				catalogUsersMetadataModel := new(watsonxdatav1.CatalogUsersMetadata)
				Expect(catalogUsersMetadataModel).ToNot(BeNil())
				catalogUsersMetadataModel.Permission = core.StringPtr("can_administer")
				catalogUsersMetadataModel.UserName = core.StringPtr("testString")
				Expect(catalogUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(catalogUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateCatalogUsersOptions model
				catalogName := "testString"
				updateCatalogUsersOptionsModel := watsonxDataService.NewUpdateCatalogUsersOptions(catalogName)
				updateCatalogUsersOptionsModel.SetCatalogName("testString")
				updateCatalogUsersOptionsModel.SetGroups([]watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel})
				updateCatalogUsersOptionsModel.SetUsers([]watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel})
				updateCatalogUsersOptionsModel.SetLhInstanceID("testString")
				updateCatalogUsersOptionsModel.SetAuthInstanceID("testString")
				updateCatalogUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogUsersOptionsModel).ToNot(BeNil())
				Expect(updateCatalogUsersOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel}))
				Expect(updateCatalogUsersOptionsModel.Users).To(Equal([]watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel}))
				Expect(updateCatalogUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDatabaseOptions successfully`, func() {
				// Construct an instance of the UpdateDatabaseBodyDatabaseDetails model
				updateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav1.UpdateDatabaseBodyDatabaseDetails)
				Expect(updateDatabaseBodyDatabaseDetailsModel).ToNot(BeNil())
				updateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				updateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				Expect(updateDatabaseBodyDatabaseDetailsModel.Password).To(Equal(core.StringPtr("samplepassword")))
				Expect(updateDatabaseBodyDatabaseDetailsModel.Username).To(Equal(core.StringPtr("sampleuser")))

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsDatabaseID := "new_db_id"
				updateDatabaseOptionsModel := watsonxDataService.NewUpdateDatabaseOptions(updateDatabaseOptionsDatabaseID)
				updateDatabaseOptionsModel.SetDatabaseID("new_db_id")
				updateDatabaseOptionsModel.SetAccept("testString")
				updateDatabaseOptionsModel.SetDatabaseDetails(updateDatabaseBodyDatabaseDetailsModel)
				updateDatabaseOptionsModel.SetDatabaseDisplayName("new_database")
				updateDatabaseOptionsModel.SetDescription("External database description")
				updateDatabaseOptionsModel.SetTags([]string{"testdatabase", "userdatabase"})
				updateDatabaseOptionsModel.SetAuthInstanceID("testString")
				updateDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDatabaseOptionsModel).ToNot(BeNil())
				Expect(updateDatabaseOptionsModel.DatabaseID).To(Equal(core.StringPtr("new_db_id")))
				Expect(updateDatabaseOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateDatabaseOptionsModel.DatabaseDetails).To(Equal(updateDatabaseBodyDatabaseDetailsModel))
				Expect(updateDatabaseOptionsModel.DatabaseDisplayName).To(Equal(core.StringPtr("new_database")))
				Expect(updateDatabaseOptionsModel.Description).To(Equal(core.StringPtr("External database description")))
				Expect(updateDatabaseOptionsModel.Tags).To(Equal([]string{"testdatabase", "userdatabase"}))
				Expect(updateDatabaseOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDbConnUsersOptions successfully`, func() {
				// Construct an instance of the BucketDbConnGroupsMetadata model
				bucketDbConnGroupsMetadataModel := new(watsonxdatav1.BucketDbConnGroupsMetadata)
				Expect(bucketDbConnGroupsMetadataModel).ToNot(BeNil())
				bucketDbConnGroupsMetadataModel.GroupID = core.StringPtr("testString")
				bucketDbConnGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the BucketDbConnUsersMetadata model
				bucketDbConnUsersMetadataModel := new(watsonxdatav1.BucketDbConnUsersMetadata)
				Expect(bucketDbConnUsersMetadataModel).ToNot(BeNil())
				bucketDbConnUsersMetadataModel.UserName = core.StringPtr("testString")
				bucketDbConnUsersMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(bucketDbConnUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))
				Expect(bucketDbConnUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the UpdateDbConnUsersOptions model
				databaseID := "testString"
				updateDbConnUsersOptionsModel := watsonxDataService.NewUpdateDbConnUsersOptions(databaseID)
				updateDbConnUsersOptionsModel.SetDatabaseID("testString")
				updateDbConnUsersOptionsModel.SetGroups([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel})
				updateDbConnUsersOptionsModel.SetUsers([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel})
				updateDbConnUsersOptionsModel.SetLhInstanceID("testString")
				updateDbConnUsersOptionsModel.SetAuthInstanceID("testString")
				updateDbConnUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDbConnUsersOptionsModel).ToNot(BeNil())
				Expect(updateDbConnUsersOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(updateDbConnUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel}))
				Expect(updateDbConnUsersOptionsModel.Users).To(Equal([]watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel}))
				Expect(updateDbConnUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDbConnUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDbConnUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEngineOptions successfully`, func() {
				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav1.NodeDescription)
				Expect(nodeDescriptionModel).ToNot(BeNil())
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))
				Expect(nodeDescriptionModel.NodeType).To(Equal(core.StringPtr("worker")))
				Expect(nodeDescriptionModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsEngineID := "sampleEngine123"
				updateEngineOptionsModel := watsonxDataService.NewUpdateEngineOptions(updateEngineOptionsEngineID)
				updateEngineOptionsModel.SetEngineID("sampleEngine123")
				updateEngineOptionsModel.SetAccept("testString")
				updateEngineOptionsModel.SetCoordinator(nodeDescriptionModel)
				updateEngineOptionsModel.SetDescription("presto engine updated description")
				updateEngineOptionsModel.SetEngineDisplayName("sampleEngine")
				updateEngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				updateEngineOptionsModel.SetWorker(nodeDescriptionModel)
				updateEngineOptionsModel.SetAuthInstanceID("testString")
				updateEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEngineOptionsModel).ToNot(BeNil())
				Expect(updateEngineOptionsModel.EngineID).To(Equal(core.StringPtr("sampleEngine123")))
				Expect(updateEngineOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineOptionsModel.Coordinator).To(Equal(nodeDescriptionModel))
				Expect(updateEngineOptionsModel.Description).To(Equal(core.StringPtr("presto engine updated description")))
				Expect(updateEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(updateEngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(updateEngineOptionsModel.Worker).To(Equal(nodeDescriptionModel))
				Expect(updateEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEngineUsersOptions successfully`, func() {
				// Construct an instance of the EngineGroupsMetadata model
				engineGroupsMetadataModel := new(watsonxdatav1.EngineGroupsMetadata)
				Expect(engineGroupsMetadataModel).ToNot(BeNil())
				engineGroupsMetadataModel.GroupID = core.StringPtr("testString")
				engineGroupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(engineGroupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(engineGroupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the EngineUsersMetadata model
				engineUsersMetadataModel := new(watsonxdatav1.EngineUsersMetadata)
				Expect(engineUsersMetadataModel).ToNot(BeNil())
				engineUsersMetadataModel.Permission = core.StringPtr("can_administer")
				engineUsersMetadataModel.UserName = core.StringPtr("testString")
				Expect(engineUsersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(engineUsersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateEngineUsersOptions model
				engineID := "testString"
				updateEngineUsersOptionsModel := watsonxDataService.NewUpdateEngineUsersOptions(engineID)
				updateEngineUsersOptionsModel.SetEngineID("testString")
				updateEngineUsersOptionsModel.SetGroups([]watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel})
				updateEngineUsersOptionsModel.SetUsers([]watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel})
				updateEngineUsersOptionsModel.SetLhInstanceID("testString")
				updateEngineUsersOptionsModel.SetAuthInstanceID("testString")
				updateEngineUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEngineUsersOptionsModel).ToNot(BeNil())
				Expect(updateEngineUsersOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel}))
				Expect(updateEngineUsersOptionsModel.Users).To(Equal([]watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel}))
				Expect(updateEngineUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMetastoreUsersOptions successfully`, func() {
				// Construct an instance of the GroupsMetadata model
				groupsMetadataModel := new(watsonxdatav1.GroupsMetadata)
				Expect(groupsMetadataModel).ToNot(BeNil())
				groupsMetadataModel.GroupID = core.StringPtr("testString")
				groupsMetadataModel.Permission = core.StringPtr("can_administer")
				Expect(groupsMetadataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(groupsMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))

				// Construct an instance of the UsersMetadata model
				usersMetadataModel := new(watsonxdatav1.UsersMetadata)
				Expect(usersMetadataModel).ToNot(BeNil())
				usersMetadataModel.Permission = core.StringPtr("can_administer")
				usersMetadataModel.UserName = core.StringPtr("testString")
				Expect(usersMetadataModel.Permission).To(Equal(core.StringPtr("can_administer")))
				Expect(usersMetadataModel.UserName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateMetastoreUsersOptions model
				metastoreName := "testString"
				updateMetastoreUsersOptionsModel := watsonxDataService.NewUpdateMetastoreUsersOptions(metastoreName)
				updateMetastoreUsersOptionsModel.SetMetastoreName("testString")
				updateMetastoreUsersOptionsModel.SetGroups([]watsonxdatav1.GroupsMetadata{*groupsMetadataModel})
				updateMetastoreUsersOptionsModel.SetUsers([]watsonxdatav1.UsersMetadata{*usersMetadataModel})
				updateMetastoreUsersOptionsModel.SetLhInstanceID("testString")
				updateMetastoreUsersOptionsModel.SetAuthInstanceID("testString")
				updateMetastoreUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMetastoreUsersOptionsModel).ToNot(BeNil())
				Expect(updateMetastoreUsersOptionsModel.MetastoreName).To(Equal(core.StringPtr("testString")))
				Expect(updateMetastoreUsersOptionsModel.Groups).To(Equal([]watsonxdatav1.GroupsMetadata{*groupsMetadataModel}))
				Expect(updateMetastoreUsersOptionsModel.Users).To(Equal([]watsonxdatav1.UsersMetadata{*usersMetadataModel}))
				Expect(updateMetastoreUsersOptionsModel.LhInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateMetastoreUsersOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateMetastoreUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateQueryOptions successfully`, func() {
				// Construct an instance of the UpdateQueryOptions model
				queryName := "testString"
				updateQueryOptionsQueryString := "testString"
				updateQueryOptionsDescription := "testString"
				updateQueryOptionsNewQueryName := "testString"
				updateQueryOptionsModel := watsonxDataService.NewUpdateQueryOptions(queryName, updateQueryOptionsQueryString, updateQueryOptionsDescription, updateQueryOptionsNewQueryName)
				updateQueryOptionsModel.SetQueryName("testString")
				updateQueryOptionsModel.SetQueryString("testString")
				updateQueryOptionsModel.SetDescription("testString")
				updateQueryOptionsModel.SetNewQueryName("testString")
				updateQueryOptionsModel.SetAuthInstanceID("testString")
				updateQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateQueryOptionsModel).ToNot(BeNil())
				Expect(updateQueryOptionsModel.QueryName).To(Equal(core.StringPtr("testString")))
				Expect(updateQueryOptionsModel.QueryString).To(Equal(core.StringPtr("testString")))
				Expect(updateQueryOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateQueryOptionsModel.NewQueryName).To(Equal(core.StringPtr("testString")))
				Expect(updateQueryOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTableOptions successfully`, func() {
				// Construct an instance of the UpdateTableBodyAddColumnsItems model
				updateTableBodyAddColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyAddColumnsItems)
				Expect(updateTableBodyAddColumnsItemsModel).ToNot(BeNil())
				updateTableBodyAddColumnsItemsModel.ColumnComment = core.StringPtr("income column")
				updateTableBodyAddColumnsItemsModel.ColumnName = core.StringPtr("income")
				updateTableBodyAddColumnsItemsModel.DataType = core.StringPtr("varchar")
				Expect(updateTableBodyAddColumnsItemsModel.ColumnComment).To(Equal(core.StringPtr("income column")))
				Expect(updateTableBodyAddColumnsItemsModel.ColumnName).To(Equal(core.StringPtr("income")))
				Expect(updateTableBodyAddColumnsItemsModel.DataType).To(Equal(core.StringPtr("varchar")))

				// Construct an instance of the UpdateTableBodyDropColumnsItems model
				updateTableBodyDropColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyDropColumnsItems)
				Expect(updateTableBodyDropColumnsItemsModel).ToNot(BeNil())
				updateTableBodyDropColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				Expect(updateTableBodyDropColumnsItemsModel.ColumnName).To(Equal(core.StringPtr("expenditure")))

				// Construct an instance of the UpdateTableBodyRenameColumnsItems model
				updateTableBodyRenameColumnsItemsModel := new(watsonxdatav1.UpdateTableBodyRenameColumnsItems)
				Expect(updateTableBodyRenameColumnsItemsModel).ToNot(BeNil())
				updateTableBodyRenameColumnsItemsModel.ColumnName = core.StringPtr("expenditure")
				updateTableBodyRenameColumnsItemsModel.NewColumnName = core.StringPtr("expenses")
				Expect(updateTableBodyRenameColumnsItemsModel.ColumnName).To(Equal(core.StringPtr("expenditure")))
				Expect(updateTableBodyRenameColumnsItemsModel.NewColumnName).To(Equal(core.StringPtr("expenses")))

				// Construct an instance of the UpdateTableOptions model
				engineID := "testString"
				catalogName := "testString"
				schemaName := "testString"
				tableName := "testString"
				updateTableOptionsModel := watsonxDataService.NewUpdateTableOptions(engineID, catalogName, schemaName, tableName)
				updateTableOptionsModel.SetEngineID("testString")
				updateTableOptionsModel.SetCatalogName("testString")
				updateTableOptionsModel.SetSchemaName("testString")
				updateTableOptionsModel.SetTableName("testString")
				updateTableOptionsModel.SetAccept("testString")
				updateTableOptionsModel.SetAddColumns([]watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel})
				updateTableOptionsModel.SetDropColumns([]watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel})
				updateTableOptionsModel.SetNewTableName("updated_table_name")
				updateTableOptionsModel.SetRenameColumns([]watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel})
				updateTableOptionsModel.SetAuthInstanceID("testString")
				updateTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTableOptionsModel).ToNot(BeNil())
				Expect(updateTableOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.SchemaName).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.TableName).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.AddColumns).To(Equal([]watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel}))
				Expect(updateTableOptionsModel.DropColumns).To(Equal([]watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel}))
				Expect(updateTableOptionsModel.NewTableName).To(Equal(core.StringPtr("updated_table_name")))
				Expect(updateTableOptionsModel.RenameColumns).To(Equal([]watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel}))
				Expect(updateTableOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUplaodCsvOptions successfully`, func() {
				// Construct an instance of the UplaodCsvOptions model
				engine := "testString"
				catalog := "testString"
				schema := "testString"
				tableName := "testString"
				ingestionJobName := "testString"
				scheduled := "testString"
				createdBy := "testString"
				targetTable := "testString"
				headers := "testString"
				csv := "testString"
				uplaodCsvOptionsModel := watsonxDataService.NewUplaodCsvOptions(engine, catalog, schema, tableName, ingestionJobName, scheduled, createdBy, targetTable, headers, csv)
				uplaodCsvOptionsModel.SetEngine("testString")
				uplaodCsvOptionsModel.SetCatalog("testString")
				uplaodCsvOptionsModel.SetSchema("testString")
				uplaodCsvOptionsModel.SetTableName("testString")
				uplaodCsvOptionsModel.SetIngestionJobName("testString")
				uplaodCsvOptionsModel.SetScheduled("testString")
				uplaodCsvOptionsModel.SetCreatedBy("testString")
				uplaodCsvOptionsModel.SetTargetTable("testString")
				uplaodCsvOptionsModel.SetHeadersVar("testString")
				uplaodCsvOptionsModel.SetCsv("testString")
				uplaodCsvOptionsModel.SetAccept("testString")
				uplaodCsvOptionsModel.SetAuthInstanceID("testString")
				uplaodCsvOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uplaodCsvOptionsModel).ToNot(BeNil())
				Expect(uplaodCsvOptionsModel.Engine).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Catalog).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Schema).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.TableName).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.IngestionJobName).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Scheduled).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.TargetTable).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.HeadersVar).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Csv).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(uplaodCsvOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBucketDbConnGroupsMetadata successfully`, func() {
				groupID := "testString"
				permission := "can_administer"
				_model, err := watsonxDataService.NewBucketDbConnGroupsMetadata(groupID, permission)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBucketDbConnUsersMetadata successfully`, func() {
				userName := "testString"
				permission := "can_administer"
				_model, err := watsonxDataService.NewBucketDbConnUsersMetadata(userName, permission)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBucketDetails successfully`, func() {
				bucketName := "sample-bucket"
				_model, err := watsonxDataService.NewBucketDetails(bucketName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCatalogGroupsMetadata successfully`, func() {
				groupID := "testString"
				permission := "can_administer"
				_model, err := watsonxDataService.NewCatalogGroupsMetadata(groupID, permission)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCatalogUsersMetadata successfully`, func() {
				permission := "can_administer"
				userName := "testString"
				_model, err := watsonxDataService.NewCatalogUsersMetadata(permission, userName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateDataPolicySchema successfully`, func() {
				catalogName := "testString"
				dataArtifact := "schema1/table1/(column1|column2)"
				policyName := "testString"
				rules := []watsonxdatav1.Rule{}
				_model, err := watsonxDataService.NewCreateDataPolicySchema(catalogName, dataArtifact, policyName, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEngineGroupsMetadata successfully`, func() {
				groupID := "testString"
				permission := "can_administer"
				_model, err := watsonxDataService.NewEngineGroupsMetadata(groupID, permission)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEngineUsersMetadata successfully`, func() {
				permission := "can_administer"
				userName := "testString"
				_model, err := watsonxDataService.NewEngineUsersMetadata(permission, userName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGroupsMetadata successfully`, func() {
				groupID := "testString"
				permission := "can_administer"
				_model, err := watsonxDataService.NewGroupsMetadata(groupID, permission)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceDataPolicySchema successfully`, func() {
				catalogName := "testString"
				dataArtifact := "schema1/table1/(column1|column2)"
				rules := []watsonxdatav1.Rule{}
				_model, err := watsonxDataService.NewReplaceDataPolicySchema(catalogName, dataArtifact, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourcesMetadata successfully`, func() {
				action := "testString"
				resourceName := "testString"
				resourceType := "engine"
				_model, err := watsonxDataService.NewResourcesMetadata(action, resourceName, resourceType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRule successfully`, func() {
				actions := []string{"all"}
				var grantee *watsonxdatav1.RuleGrantee = nil
				_, err := watsonxDataService.NewRule(actions, grantee)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRuleGrantee successfully`, func() {
				value := "testString"
				key := "user_name"
				typeVar := "user_identity"
				_model, err := watsonxDataService.NewRuleGrantee(value, key, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUsersMetadata successfully`, func() {
				permission := "can_administer"
				userName := "testString"
				_model, err := watsonxDataService.NewUsersMetadata(permission, userName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
