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

package watsonxdatav2_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`WatsonxDataV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(watsonxDataService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
				URL: "https://watsonxdatav2/api",
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
				"WATSONX_DATA_URL": "https://watsonxdatav2/api",
				"WATSONX_DATA_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(&watsonxdatav2.WatsonxDataV2Options{
				})
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(&watsonxdatav2.WatsonxDataV2Options{
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(&watsonxdatav2.WatsonxDataV2Options{
				})
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
				"WATSONX_DATA_URL": "https://watsonxdatav2/api",
				"WATSONX_DATA_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(&watsonxdatav2.WatsonxDataV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(watsonxDataService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WATSONX_DATA_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(&watsonxdatav2.WatsonxDataV2Options{
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
			url, err = watsonxdatav2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions) - Operation response error`, func() {
		listBucketRegistrationsPath := "/bucket_registrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBucketRegistrations with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := new(watsonxdatav2.ListBucketRegistrationsOptions)
				listBucketRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions)`, func() {
		listBucketRegistrationsPath := "/bucket_registrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_registrations": [{"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListBucketRegistrations successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := new(watsonxdatav2.ListBucketRegistrationsOptions)
				listBucketRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListBucketRegistrationsWithContext(ctx, listBucketRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListBucketRegistrationsWithContext(ctx, listBucketRegistrationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBucketRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_registrations": [{"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListBucketRegistrations successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListBucketRegistrations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := new(watsonxdatav2.ListBucketRegistrationsOptions)
				listBucketRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBucketRegistrations with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := new(watsonxdatav2.ListBucketRegistrationsOptions)
				listBucketRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
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
			It(`Invoke ListBucketRegistrations successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := new(watsonxdatav2.ListBucketRegistrationsOptions)
				listBucketRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptionsModel)
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
	Describe(`CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions) - Operation response error`, func() {
		createBucketRegistrationPath := "/bucket_registrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketRegistrationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBucketRegistration with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")

				// Construct an instance of the CreateBucketRegistrationOptions model
				createBucketRegistrationOptionsModel := new(watsonxdatav2.CreateBucketRegistrationOptions)
				createBucketRegistrationOptionsModel.BucketDetails = bucketDetailsModel
				createBucketRegistrationOptionsModel.BucketType = core.StringPtr("ibm_cos")
				createBucketRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createBucketRegistrationOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				createBucketRegistrationOptionsModel.ManagedBy = core.StringPtr("ibm")
				createBucketRegistrationOptionsModel.TableType = core.StringPtr("iceberg")
				createBucketRegistrationOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				createBucketRegistrationOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				createBucketRegistrationOptionsModel.Region = core.StringPtr("us-south")
				createBucketRegistrationOptionsModel.State = core.StringPtr("active")
				createBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions)`, func() {
		createBucketRegistrationPath := "/bucket_registrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateBucketRegistration successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")

				// Construct an instance of the CreateBucketRegistrationOptions model
				createBucketRegistrationOptionsModel := new(watsonxdatav2.CreateBucketRegistrationOptions)
				createBucketRegistrationOptionsModel.BucketDetails = bucketDetailsModel
				createBucketRegistrationOptionsModel.BucketType = core.StringPtr("ibm_cos")
				createBucketRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createBucketRegistrationOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				createBucketRegistrationOptionsModel.ManagedBy = core.StringPtr("ibm")
				createBucketRegistrationOptionsModel.TableType = core.StringPtr("iceberg")
				createBucketRegistrationOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				createBucketRegistrationOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				createBucketRegistrationOptionsModel.Region = core.StringPtr("us-south")
				createBucketRegistrationOptionsModel.State = core.StringPtr("active")
				createBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateBucketRegistrationWithContext(ctx, createBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateBucketRegistrationWithContext(ctx, createBucketRegistrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createBucketRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateBucketRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")

				// Construct an instance of the CreateBucketRegistrationOptions model
				createBucketRegistrationOptionsModel := new(watsonxdatav2.CreateBucketRegistrationOptions)
				createBucketRegistrationOptionsModel.BucketDetails = bucketDetailsModel
				createBucketRegistrationOptionsModel.BucketType = core.StringPtr("ibm_cos")
				createBucketRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createBucketRegistrationOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				createBucketRegistrationOptionsModel.ManagedBy = core.StringPtr("ibm")
				createBucketRegistrationOptionsModel.TableType = core.StringPtr("iceberg")
				createBucketRegistrationOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				createBucketRegistrationOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				createBucketRegistrationOptionsModel.Region = core.StringPtr("us-south")
				createBucketRegistrationOptionsModel.State = core.StringPtr("active")
				createBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBucketRegistration with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")

				// Construct an instance of the CreateBucketRegistrationOptions model
				createBucketRegistrationOptionsModel := new(watsonxdatav2.CreateBucketRegistrationOptions)
				createBucketRegistrationOptionsModel.BucketDetails = bucketDetailsModel
				createBucketRegistrationOptionsModel.BucketType = core.StringPtr("ibm_cos")
				createBucketRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createBucketRegistrationOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				createBucketRegistrationOptionsModel.ManagedBy = core.StringPtr("ibm")
				createBucketRegistrationOptionsModel.TableType = core.StringPtr("iceberg")
				createBucketRegistrationOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				createBucketRegistrationOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				createBucketRegistrationOptionsModel.Region = core.StringPtr("us-south")
				createBucketRegistrationOptionsModel.State = core.StringPtr("active")
				createBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBucketRegistrationOptions model with no property values
				createBucketRegistrationOptionsModelNew := new(watsonxdatav2.CreateBucketRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModelNew)
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
			It(`Invoke CreateBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")

				// Construct an instance of the CreateBucketRegistrationOptions model
				createBucketRegistrationOptionsModel := new(watsonxdatav2.CreateBucketRegistrationOptions)
				createBucketRegistrationOptionsModel.BucketDetails = bucketDetailsModel
				createBucketRegistrationOptionsModel.BucketType = core.StringPtr("ibm_cos")
				createBucketRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createBucketRegistrationOptionsModel.Description = core.StringPtr("COS bucket for customer data")
				createBucketRegistrationOptionsModel.ManagedBy = core.StringPtr("ibm")
				createBucketRegistrationOptionsModel.TableType = core.StringPtr("iceberg")
				createBucketRegistrationOptionsModel.BucketDisplayName = core.StringPtr("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.BucketTags = []string{"read customer data", "write customer data'"}
				createBucketRegistrationOptionsModel.CatalogTags = []string{"catalog_tag_1", "catalog_tag_2"}
				createBucketRegistrationOptionsModel.Region = core.StringPtr("us-south")
				createBucketRegistrationOptionsModel.State = core.StringPtr("active")
				createBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptionsModel)
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
	Describe(`GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions) - Operation response error`, func() {
		getBucketRegistrationPath := "/bucket_registrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketRegistrationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketRegistration with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketRegistrationOptions model
				getBucketRegistrationOptionsModel := new(watsonxdatav2.GetBucketRegistrationOptions)
				getBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions)`, func() {
		getBucketRegistrationPath := "/bucket_registrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetBucketRegistration successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketRegistrationOptions model
				getBucketRegistrationOptionsModel := new(watsonxdatav2.GetBucketRegistrationOptions)
				getBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetBucketRegistrationWithContext(ctx, getBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetBucketRegistrationWithContext(ctx, getBucketRegistrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBucketRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetBucketRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketRegistrationOptions model
				getBucketRegistrationOptionsModel := new(watsonxdatav2.GetBucketRegistrationOptions)
				getBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketRegistration with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketRegistrationOptions model
				getBucketRegistrationOptionsModel := new(watsonxdatav2.GetBucketRegistrationOptions)
				getBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketRegistrationOptions model with no property values
				getBucketRegistrationOptionsModelNew := new(watsonxdatav2.GetBucketRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModelNew)
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
			It(`Invoke GetBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetBucketRegistrationOptions model
				getBucketRegistrationOptionsModel := new(watsonxdatav2.GetBucketRegistrationOptions)
				getBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptionsModel)
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
	Describe(`DeleteBucketRegistration(deleteBucketRegistrationOptions *DeleteBucketRegistrationOptions)`, func() {
		deleteBucketRegistrationPath := "/bucket_registrations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketRegistrationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteBucketRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketRegistrationOptions model
				deleteBucketRegistrationOptionsModel := new(watsonxdatav2.DeleteBucketRegistrationOptions)
				deleteBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				deleteBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteBucketRegistration(deleteBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucketRegistration with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketRegistrationOptions model
				deleteBucketRegistrationOptionsModel := new(watsonxdatav2.DeleteBucketRegistrationOptions)
				deleteBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				deleteBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteBucketRegistration(deleteBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketRegistrationOptions model with no property values
				deleteBucketRegistrationOptionsModelNew := new(watsonxdatav2.DeleteBucketRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteBucketRegistration(deleteBucketRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions) - Operation response error`, func() {
		updateBucketRegistrationPath := "/bucket_registrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBucketRegistration with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateBucketRegistrationOptions model
				updateBucketRegistrationOptionsModel := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				updateBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions)`, func() {
		updateBucketRegistrationPath := "/bucket_registrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateBucketRegistration successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateBucketRegistrationOptions model
				updateBucketRegistrationOptionsModel := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				updateBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateBucketRegistrationWithContext(ctx, updateBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateBucketRegistrationWithContext(ctx, updateBucketRegistrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"bucket_registration": {"access_key": "<access_key>", "actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "bucket_display_name": "sample-bucket-displayname", "bucket_id": "samplebucket123", "bucket_name": "sample-bucket", "bucket_type": "ibm_cos", "created_by": "<username>@<domain>.com", "created_on": "1686120645", "description": "COS bucket for customer data", "endpoint": "https://s3.<region>.cloud-object-storage.appdomain.cloud/", "managed_by": "ibm", "region": "us-south", "secret_key": "secret_key", "state": "active", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateBucketRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateBucketRegistrationOptions model
				updateBucketRegistrationOptionsModel := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				updateBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBucketRegistration with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateBucketRegistrationOptions model
				updateBucketRegistrationOptionsModel := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				updateBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBucketRegistrationOptions model with no property values
				updateBucketRegistrationOptionsModelNew := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModelNew)
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
			It(`Invoke UpdateBucketRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateBucketRegistrationOptions model
				updateBucketRegistrationOptionsModel := new(watsonxdatav2.UpdateBucketRegistrationOptions)
				updateBucketRegistrationOptionsModel.BucketID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateBucketRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateBucketRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptionsModel)
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
	Describe(`CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions) - Operation response error`, func() {
		createActivateBucketPath := "/bucket_registrations/testString/activate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createActivateBucketPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateActivateBucket with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateActivateBucketOptions model
				createActivateBucketOptionsModel := new(watsonxdatav2.CreateActivateBucketOptions)
				createActivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				createActivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createActivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions)`, func() {
		createActivateBucketPath := "/bucket_registrations/testString/activate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createActivateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateActivateBucket successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateActivateBucketOptions model
				createActivateBucketOptionsModel := new(watsonxdatav2.CreateActivateBucketOptions)
				createActivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				createActivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createActivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateActivateBucketWithContext(ctx, createActivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateActivateBucketWithContext(ctx, createActivateBucketOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createActivateBucketPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateActivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateActivateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateActivateBucketOptions model
				createActivateBucketOptionsModel := new(watsonxdatav2.CreateActivateBucketOptions)
				createActivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				createActivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createActivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateActivateBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateActivateBucketOptions model
				createActivateBucketOptionsModel := new(watsonxdatav2.CreateActivateBucketOptions)
				createActivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				createActivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createActivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateActivateBucketOptions model with no property values
				createActivateBucketOptionsModelNew := new(watsonxdatav2.CreateActivateBucketOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModelNew)
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
			It(`Invoke CreateActivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateActivateBucketOptions model
				createActivateBucketOptionsModel := new(watsonxdatav2.CreateActivateBucketOptions)
				createActivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				createActivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createActivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateActivateBucket(createActivateBucketOptionsModel)
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
	Describe(`DeleteDeactivateBucket(deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions)`, func() {
		deleteDeactivateBucketPath := "/bucket_registrations/testString/deactivate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDeactivateBucketPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDeactivateBucket successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDeactivateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDeactivateBucketOptions model
				deleteDeactivateBucketOptionsModel := new(watsonxdatav2.DeleteDeactivateBucketOptions)
				deleteDeactivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				deleteDeactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDeactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDeactivateBucket(deleteDeactivateBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDeactivateBucket with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDeactivateBucketOptions model
				deleteDeactivateBucketOptionsModel := new(watsonxdatav2.DeleteDeactivateBucketOptions)
				deleteDeactivateBucketOptionsModel.BucketID = core.StringPtr("testString")
				deleteDeactivateBucketOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDeactivateBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDeactivateBucket(deleteDeactivateBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDeactivateBucketOptions model with no property values
				deleteDeactivateBucketOptionsModelNew := new(watsonxdatav2.DeleteDeactivateBucketOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteDeactivateBucket(deleteDeactivateBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions) - Operation response error`, func() {
		listBucketObjectsPath := "/bucket_registrations/testString/objects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBucketObjects with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketObjectsOptions model
				listBucketObjectsOptionsModel := new(watsonxdatav2.ListBucketObjectsOptions)
				listBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions)`, func() {
		listBucketObjectsPath := "/bucket_registrations/testString/objects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"objects": ["object_1"], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListBucketObjects successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListBucketObjectsOptions model
				listBucketObjectsOptionsModel := new(watsonxdatav2.ListBucketObjectsOptions)
				listBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListBucketObjectsWithContext(ctx, listBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListBucketObjectsWithContext(ctx, listBucketObjectsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBucketObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"objects": ["object_1"], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListBucketObjects successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListBucketObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBucketObjectsOptions model
				listBucketObjectsOptionsModel := new(watsonxdatav2.ListBucketObjectsOptions)
				listBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBucketObjects with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketObjectsOptions model
				listBucketObjectsOptionsModel := new(watsonxdatav2.ListBucketObjectsOptions)
				listBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBucketObjectsOptions model with no property values
				listBucketObjectsOptionsModelNew := new(watsonxdatav2.ListBucketObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModelNew)
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
			It(`Invoke ListBucketObjects successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListBucketObjectsOptions model
				listBucketObjectsOptionsModel := new(watsonxdatav2.ListBucketObjectsOptions)
				listBucketObjectsOptionsModel.BucketID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listBucketObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListBucketObjects(listBucketObjectsOptionsModel)
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
	Describe(`TestBucketConnection(testBucketConnectionOptions *TestBucketConnectionOptions) - Operation response error`, func() {
		testBucketConnectionPath := "/test_bucket_connection"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testBucketConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TestBucketConnection with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsModel := new(watsonxdatav2.TestBucketConnectionOptions)
				testBucketConnectionOptionsModel.AccessKey = core.StringPtr("<access_key>")
				testBucketConnectionOptionsModel.BucketName = core.StringPtr("sample-bucket")
				testBucketConnectionOptionsModel.BucketType = core.StringPtr("ibm_cos")
				testBucketConnectionOptionsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.Region = core.StringPtr("us-south")
				testBucketConnectionOptionsModel.SecretKey = core.StringPtr("secret_key")
				testBucketConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				testBucketConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestBucketConnection(testBucketConnectionOptions *TestBucketConnectionOptions)`, func() {
		testBucketConnectionPath := "/test_bucket_connection"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testBucketConnectionPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_status": {"state": true, "state_message": "bucket does not exist or the credentials provided are not valid."}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke TestBucketConnection successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsModel := new(watsonxdatav2.TestBucketConnectionOptions)
				testBucketConnectionOptionsModel.AccessKey = core.StringPtr("<access_key>")
				testBucketConnectionOptionsModel.BucketName = core.StringPtr("sample-bucket")
				testBucketConnectionOptionsModel.BucketType = core.StringPtr("ibm_cos")
				testBucketConnectionOptionsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.Region = core.StringPtr("us-south")
				testBucketConnectionOptionsModel.SecretKey = core.StringPtr("secret_key")
				testBucketConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				testBucketConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.TestBucketConnectionWithContext(ctx, testBucketConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.TestBucketConnectionWithContext(ctx, testBucketConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(testBucketConnectionPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bucket_status": {"state": true, "state_message": "bucket does not exist or the credentials provided are not valid."}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke TestBucketConnection successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.TestBucketConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsModel := new(watsonxdatav2.TestBucketConnectionOptions)
				testBucketConnectionOptionsModel.AccessKey = core.StringPtr("<access_key>")
				testBucketConnectionOptionsModel.BucketName = core.StringPtr("sample-bucket")
				testBucketConnectionOptionsModel.BucketType = core.StringPtr("ibm_cos")
				testBucketConnectionOptionsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.Region = core.StringPtr("us-south")
				testBucketConnectionOptionsModel.SecretKey = core.StringPtr("secret_key")
				testBucketConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				testBucketConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke TestBucketConnection with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsModel := new(watsonxdatav2.TestBucketConnectionOptions)
				testBucketConnectionOptionsModel.AccessKey = core.StringPtr("<access_key>")
				testBucketConnectionOptionsModel.BucketName = core.StringPtr("sample-bucket")
				testBucketConnectionOptionsModel.BucketType = core.StringPtr("ibm_cos")
				testBucketConnectionOptionsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.Region = core.StringPtr("us-south")
				testBucketConnectionOptionsModel.SecretKey = core.StringPtr("secret_key")
				testBucketConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				testBucketConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TestBucketConnectionOptions model with no property values
				testBucketConnectionOptionsModelNew := new(watsonxdatav2.TestBucketConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModelNew)
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
			It(`Invoke TestBucketConnection successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsModel := new(watsonxdatav2.TestBucketConnectionOptions)
				testBucketConnectionOptionsModel.AccessKey = core.StringPtr("<access_key>")
				testBucketConnectionOptionsModel.BucketName = core.StringPtr("sample-bucket")
				testBucketConnectionOptionsModel.BucketType = core.StringPtr("ibm_cos")
				testBucketConnectionOptionsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.Region = core.StringPtr("us-south")
				testBucketConnectionOptionsModel.SecretKey = core.StringPtr("secret_key")
				testBucketConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				testBucketConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.TestBucketConnection(testBucketConnectionOptionsModel)
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
	Describe(`CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptions *CreateDriverDatabaseCatalogOptions) - Operation response error`, func() {
		createDriverDatabaseCatalogPath := "/database_driver_registrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDriverDatabaseCatalogPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDriverDatabaseCatalog with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				createDriverDatabaseCatalogOptionsModel := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Hostname = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Port = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Driver = CreateMockReader("This is a mock file.")
				createDriverDatabaseCatalogOptionsModel.DriverContentType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DriverFileName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Certificate = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CertificateExtension = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Ssl = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Username = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Password = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Description = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CreatedOn = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptions *CreateDriverDatabaseCatalogOptions)`, func() {
		createDriverDatabaseCatalogPath := "/database_driver_registrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDriverDatabaseCatalogPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"database": {"database_display_name": "DatabaseDisplayName", "database_id": "DatabaseID"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDriverDatabaseCatalog successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				createDriverDatabaseCatalogOptionsModel := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Hostname = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Port = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Driver = CreateMockReader("This is a mock file.")
				createDriverDatabaseCatalogOptionsModel.DriverContentType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DriverFileName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Certificate = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CertificateExtension = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Ssl = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Username = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Password = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Description = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CreatedOn = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDriverDatabaseCatalogWithContext(ctx, createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDriverDatabaseCatalogWithContext(ctx, createDriverDatabaseCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDriverDatabaseCatalogPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"database": {"database_display_name": "DatabaseDisplayName", "database_id": "DatabaseID"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDriverDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDriverDatabaseCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				createDriverDatabaseCatalogOptionsModel := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Hostname = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Port = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Driver = CreateMockReader("This is a mock file.")
				createDriverDatabaseCatalogOptionsModel.DriverContentType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DriverFileName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Certificate = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CertificateExtension = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Ssl = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Username = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Password = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Description = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CreatedOn = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDriverDatabaseCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				createDriverDatabaseCatalogOptionsModel := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Hostname = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Port = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Driver = CreateMockReader("This is a mock file.")
				createDriverDatabaseCatalogOptionsModel.DriverContentType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DriverFileName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Certificate = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CertificateExtension = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Ssl = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Username = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Password = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Description = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CreatedOn = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDriverDatabaseCatalogOptions model with no property values
				createDriverDatabaseCatalogOptionsModelNew := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModelNew)
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
			It(`Invoke CreateDriverDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				createDriverDatabaseCatalogOptionsModel := new(watsonxdatav2.CreateDriverDatabaseCatalogOptions)
				createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CatalogName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Hostname = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Port = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Driver = CreateMockReader("This is a mock file.")
				createDriverDatabaseCatalogOptionsModel.DriverContentType = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DriverFileName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Certificate = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CertificateExtension = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Ssl = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Username = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Password = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.DatabaseName = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Description = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.CreatedOn = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDriverDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptionsModel)
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
	Describe(`ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions) - Operation response error`, func() {
		listDatabaseRegistrationsPath := "/database_registrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabaseRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDatabaseRegistrations with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := new(watsonxdatav2.ListDatabaseRegistrationsOptions)
				listDatabaseRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDatabaseRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions)`, func() {
		listDatabaseRegistrationsPath := "/database_registrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabaseRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database_registrations": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListDatabaseRegistrations successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := new(watsonxdatav2.ListDatabaseRegistrationsOptions)
				listDatabaseRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDatabaseRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListDatabaseRegistrationsWithContext(ctx, listDatabaseRegistrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListDatabaseRegistrationsWithContext(ctx, listDatabaseRegistrationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDatabaseRegistrationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database_registrations": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListDatabaseRegistrations successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListDatabaseRegistrations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := new(watsonxdatav2.ListDatabaseRegistrationsOptions)
				listDatabaseRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDatabaseRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDatabaseRegistrations with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := new(watsonxdatav2.ListDatabaseRegistrationsOptions)
				listDatabaseRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDatabaseRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
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
			It(`Invoke ListDatabaseRegistrations successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := new(watsonxdatav2.ListDatabaseRegistrationsOptions)
				listDatabaseRegistrationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDatabaseRegistrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptionsModel)
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
	Describe(`CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions) - Operation response error`, func() {
		createDatabaseRegistrationPath := "/database_registrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseRegistrationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDatabaseRegistration with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsModel := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				createDatabaseRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseRegistrationOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseRegistrationOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseRegistrationOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseRegistrationOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseRegistrationOptionsModel.DatabaseProperties = []watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}
				createDatabaseRegistrationOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions)`, func() {
		createDatabaseRegistrationPath := "/database_registrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"database_registration": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDatabaseRegistration successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsModel := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				createDatabaseRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseRegistrationOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseRegistrationOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseRegistrationOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseRegistrationOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseRegistrationOptionsModel.DatabaseProperties = []watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}
				createDatabaseRegistrationOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDatabaseRegistrationWithContext(ctx, createDatabaseRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDatabaseRegistrationWithContext(ctx, createDatabaseRegistrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseRegistrationPath))
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
					fmt.Fprintf(res, "%s", `{"database_registration": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDatabaseRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDatabaseRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsModel := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				createDatabaseRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseRegistrationOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseRegistrationOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseRegistrationOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseRegistrationOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseRegistrationOptionsModel.DatabaseProperties = []watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}
				createDatabaseRegistrationOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDatabaseRegistration with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsModel := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				createDatabaseRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseRegistrationOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseRegistrationOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseRegistrationOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseRegistrationOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseRegistrationOptionsModel.DatabaseProperties = []watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}
				createDatabaseRegistrationOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDatabaseRegistrationOptions model with no property values
				createDatabaseRegistrationOptionsModelNew := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModelNew)
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
			It(`Invoke CreateDatabaseRegistration successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsModel := new(watsonxdatav2.CreateDatabaseRegistrationOptions)
				createDatabaseRegistrationOptionsModel.CatalogName = core.StringPtr("sampleCatalog")
				createDatabaseRegistrationOptionsModel.DatabaseDisplayName = core.StringPtr("new_database")
				createDatabaseRegistrationOptionsModel.DatabaseType = core.StringPtr("db2")
				createDatabaseRegistrationOptionsModel.CreatedOn = core.Int64Ptr(int64(38))
				createDatabaseRegistrationOptionsModel.DatabaseDetails = registerDatabaseCatalogBodyDatabaseDetailsModel
				createDatabaseRegistrationOptionsModel.DatabaseProperties = []watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}
				createDatabaseRegistrationOptionsModel.Description = core.StringPtr("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.Tags = []string{"tag_1", "tag_2"}
				createDatabaseRegistrationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDatabaseRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptionsModel)
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
	Describe(`GetDatabase(getDatabaseOptions *GetDatabaseOptions) - Operation response error`, func() {
		getDatabasePath := "/database_registrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabasePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDatabase with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseOptions model
				getDatabaseOptionsModel := new(watsonxdatav2.GetDatabaseOptions)
				getDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				getDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetDatabase(getDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetDatabase(getDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDatabase(getDatabaseOptions *GetDatabaseOptions)`, func() {
		getDatabasePath := "/database_registrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabasePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetDatabase successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDatabaseOptions model
				getDatabaseOptionsModel := new(watsonxdatav2.GetDatabaseOptions)
				getDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				getDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetDatabaseWithContext(ctx, getDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetDatabase(getDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetDatabaseWithContext(ctx, getDatabaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDatabasePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetDatabase successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetDatabase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDatabaseOptions model
				getDatabaseOptionsModel := new(watsonxdatav2.GetDatabaseOptions)
				getDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				getDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDatabase(getDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDatabase with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseOptions model
				getDatabaseOptionsModel := new(watsonxdatav2.GetDatabaseOptions)
				getDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				getDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetDatabase(getDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDatabaseOptions model with no property values
				getDatabaseOptionsModelNew := new(watsonxdatav2.GetDatabaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetDatabase(getDatabaseOptionsModelNew)
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
			It(`Invoke GetDatabase successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseOptions model
				getDatabaseOptionsModel := new(watsonxdatav2.GetDatabaseOptions)
				getDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				getDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDatabase(getDatabaseOptionsModel)
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
	Describe(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
		deleteDatabaseCatalogPath := "/database_registrations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseCatalogPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDatabaseCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				deleteDatabaseCatalogOptionsModel := new(watsonxdatav2.DeleteDatabaseCatalogOptions)
				deleteDatabaseCatalogOptionsModel.DatabaseID = core.StringPtr("testString")
				deleteDatabaseCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDatabaseCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDatabaseCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseCatalogOptions model
				deleteDatabaseCatalogOptionsModel := new(watsonxdatav2.DeleteDatabaseCatalogOptions)
				deleteDatabaseCatalogOptionsModel.DatabaseID = core.StringPtr("testString")
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
				deleteDatabaseCatalogOptionsModelNew := new(watsonxdatav2.DeleteDatabaseCatalogOptions)
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
	Describe(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions) - Operation response error`, func() {
		updateDatabasePath := "/database_registrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabasePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDatabase with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav2.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions)`, func() {
		updateDatabasePath := "/database_registrations/testString"
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateDatabase successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav2.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "created_by": "user1@bim.com", "created_on": "1686792721", "database_details": {"database_name": "new_database", "hostname": "netezza://ps.fyre.com", "password": "samplepassword", "port": 4543, "sasl": true, "ssl": true, "tables": "kafka_table_name", "username": "sampleuser"}, "database_display_name": "new_database", "database_id": "new_database_id", "database_properties": ["DatabaseProperties"], "database_type": "netezza", "description": "Description of the external Database", "tags": ["Tags"]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateDatabase successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav2.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDatabase with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav2.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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
				updateDatabaseOptionsModelNew := new(watsonxdatav2.UpdateDatabaseOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDatabaseOptions model
				updateDatabaseOptionsModel := new(watsonxdatav2.UpdateDatabaseOptions)
				updateDatabaseOptionsModel.DatabaseID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDatabaseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateDatabase(updateDatabaseOptionsModel)
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
	Describe(`ValidateDatabaseConnection(validateDatabaseConnectionOptions *ValidateDatabaseConnectionOptions) - Operation response error`, func() {
		validateDatabaseConnectionPath := "/test_database_connection"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateDatabaseConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateDatabaseConnection with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				validateDatabaseConnectionOptionsModel := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				validateDatabaseConnectionOptionsModel.DatabaseDetails = validateDatabaseBodyDatabaseDetailsModel
				validateDatabaseConnectionOptionsModel.DatabaseType = core.StringPtr("netezza")
				validateDatabaseConnectionOptionsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				validateDatabaseConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateDatabaseConnection(validateDatabaseConnectionOptions *ValidateDatabaseConnectionOptions)`, func() {
		validateDatabaseConnectionPath := "/test_database_connection"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateDatabaseConnectionPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection_response": {"state": false, "state_message": "StateMessage"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ValidateDatabaseConnection successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				validateDatabaseConnectionOptionsModel := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				validateDatabaseConnectionOptionsModel.DatabaseDetails = validateDatabaseBodyDatabaseDetailsModel
				validateDatabaseConnectionOptionsModel.DatabaseType = core.StringPtr("netezza")
				validateDatabaseConnectionOptionsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				validateDatabaseConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ValidateDatabaseConnectionWithContext(ctx, validateDatabaseConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ValidateDatabaseConnectionWithContext(ctx, validateDatabaseConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(validateDatabaseConnectionPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection_response": {"state": false, "state_message": "StateMessage"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ValidateDatabaseConnection successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ValidateDatabaseConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				validateDatabaseConnectionOptionsModel := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				validateDatabaseConnectionOptionsModel.DatabaseDetails = validateDatabaseBodyDatabaseDetailsModel
				validateDatabaseConnectionOptionsModel.DatabaseType = core.StringPtr("netezza")
				validateDatabaseConnectionOptionsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				validateDatabaseConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateDatabaseConnection with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				validateDatabaseConnectionOptionsModel := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				validateDatabaseConnectionOptionsModel.DatabaseDetails = validateDatabaseBodyDatabaseDetailsModel
				validateDatabaseConnectionOptionsModel.DatabaseType = core.StringPtr("netezza")
				validateDatabaseConnectionOptionsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				validateDatabaseConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ValidateDatabaseConnectionOptions model with no property values
				validateDatabaseConnectionOptionsModelNew := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModelNew)
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
			It(`Invoke ValidateDatabaseConnection successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				validateDatabaseConnectionOptionsModel := new(watsonxdatav2.ValidateDatabaseConnectionOptions)
				validateDatabaseConnectionOptionsModel.DatabaseDetails = validateDatabaseBodyDatabaseDetailsModel
				validateDatabaseConnectionOptionsModel.DatabaseType = core.StringPtr("netezza")
				validateDatabaseConnectionOptionsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.AuthInstanceID = core.StringPtr("testString")
				validateDatabaseConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptionsModel)
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
	Describe(`ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions) - Operation response error`, func() {
		listDb2EnginesPath := "/db2_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDb2EnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDb2Engines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := new(watsonxdatav2.ListDb2EnginesOptions)
				listDb2EnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDb2EnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions)`, func() {
		listDb2EnginesPath := "/db2_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDb2EnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"db2_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListDb2Engines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := new(watsonxdatav2.ListDb2EnginesOptions)
				listDb2EnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDb2EnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListDb2EnginesWithContext(ctx, listDb2EnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListDb2EnginesWithContext(ctx, listDb2EnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDb2EnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"db2_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListDb2Engines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListDb2Engines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := new(watsonxdatav2.ListDb2EnginesOptions)
				listDb2EnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDb2EnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDb2Engines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := new(watsonxdatav2.ListDb2EnginesOptions)
				listDb2EnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDb2EnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
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
			It(`Invoke ListDb2Engines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := new(watsonxdatav2.ListDb2EnginesOptions)
				listDb2EnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listDb2EnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListDb2Engines(listDb2EnginesOptionsModel)
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
	Describe(`CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions) - Operation response error`, func() {
		createDb2EnginePath := "/db2_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDb2EnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDb2Engine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsModel := new(watsonxdatav2.CreateDb2EngineOptions)
				createDb2EngineOptionsModel.Origin = core.StringPtr("external")
				createDb2EngineOptionsModel.Type = core.StringPtr("db2")
				createDb2EngineOptionsModel.Description = core.StringPtr("db2 engine description")
				createDb2EngineOptionsModel.EngineDetails = createDb2EngineDetailsModel
				createDb2EngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createDb2EngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions)`, func() {
		createDb2EnginePath := "/db2_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDb2EnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDb2Engine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsModel := new(watsonxdatav2.CreateDb2EngineOptions)
				createDb2EngineOptionsModel.Origin = core.StringPtr("external")
				createDb2EngineOptionsModel.Type = core.StringPtr("db2")
				createDb2EngineOptionsModel.Description = core.StringPtr("db2 engine description")
				createDb2EngineOptionsModel.EngineDetails = createDb2EngineDetailsModel
				createDb2EngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createDb2EngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateDb2EngineWithContext(ctx, createDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateDb2EngineWithContext(ctx, createDb2EngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDb2EnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateDb2Engine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateDb2Engine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsModel := new(watsonxdatav2.CreateDb2EngineOptions)
				createDb2EngineOptionsModel.Origin = core.StringPtr("external")
				createDb2EngineOptionsModel.Type = core.StringPtr("db2")
				createDb2EngineOptionsModel.Description = core.StringPtr("db2 engine description")
				createDb2EngineOptionsModel.EngineDetails = createDb2EngineDetailsModel
				createDb2EngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createDb2EngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDb2Engine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsModel := new(watsonxdatav2.CreateDb2EngineOptions)
				createDb2EngineOptionsModel.Origin = core.StringPtr("external")
				createDb2EngineOptionsModel.Type = core.StringPtr("db2")
				createDb2EngineOptionsModel.Description = core.StringPtr("db2 engine description")
				createDb2EngineOptionsModel.EngineDetails = createDb2EngineDetailsModel
				createDb2EngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createDb2EngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDb2EngineOptions model with no property values
				createDb2EngineOptionsModelNew := new(watsonxdatav2.CreateDb2EngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModelNew)
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
			It(`Invoke CreateDb2Engine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsModel := new(watsonxdatav2.CreateDb2EngineOptions)
				createDb2EngineOptionsModel.Origin = core.StringPtr("external")
				createDb2EngineOptionsModel.Type = core.StringPtr("db2")
				createDb2EngineOptionsModel.Description = core.StringPtr("db2 engine description")
				createDb2EngineOptionsModel.EngineDetails = createDb2EngineDetailsModel
				createDb2EngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createDb2EngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateDb2Engine(createDb2EngineOptionsModel)
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
	Describe(`DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions)`, func() {
		deleteDb2EnginePath := "/db2_engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDb2EnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDb2Engine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteDb2Engine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDb2EngineOptions model
				deleteDb2EngineOptionsModel := new(watsonxdatav2.DeleteDb2EngineOptions)
				deleteDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDb2Engine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteDb2EngineOptions model
				deleteDb2EngineOptionsModel := new(watsonxdatav2.DeleteDb2EngineOptions)
				deleteDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDb2EngineOptions model with no property values
				deleteDb2EngineOptionsModelNew := new(watsonxdatav2.DeleteDb2EngineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions) - Operation response error`, func() {
		updateDb2EnginePath := "/db2_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDb2EnginePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDb2Engine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDb2EngineOptions model
				updateDb2EngineOptionsModel := new(watsonxdatav2.UpdateDb2EngineOptions)
				updateDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions)`, func() {
		updateDb2EnginePath := "/db2_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDb2EnginePath))
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
					fmt.Fprintf(res, "%s", `{"db2_engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateDb2Engine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDb2EngineOptions model
				updateDb2EngineOptionsModel := new(watsonxdatav2.UpdateDb2EngineOptions)
				updateDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateDb2EngineWithContext(ctx, updateDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateDb2EngineWithContext(ctx, updateDb2EngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateDb2EnginePath))
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
					fmt.Fprintf(res, "%s", `{"db2_engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateDb2Engine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateDb2Engine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDb2EngineOptions model
				updateDb2EngineOptionsModel := new(watsonxdatav2.UpdateDb2EngineOptions)
				updateDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDb2Engine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDb2EngineOptions model
				updateDb2EngineOptionsModel := new(watsonxdatav2.UpdateDb2EngineOptions)
				updateDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDb2EngineOptions model with no property values
				updateDb2EngineOptionsModelNew := new(watsonxdatav2.UpdateDb2EngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModelNew)
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
			It(`Invoke UpdateDb2Engine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDb2EngineOptions model
				updateDb2EngineOptionsModel := new(watsonxdatav2.UpdateDb2EngineOptions)
				updateDb2EngineOptionsModel.EngineID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateDb2EngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateDb2EngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptionsModel)
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
	Describe(`ListEngines(listEnginesOptions *ListEnginesOptions) - Operation response error`, func() {
		listEnginesPath := "/engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := new(watsonxdatav2.ListEnginesOptions)
				listEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListEngines(listEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListEngines(listEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEngines(listEnginesOptions *ListEnginesOptions)`, func() {
		listEnginesPath := "/engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engines": {"db2_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}], "milvus_services": [{"actions": ["Actions"], "created_by": "<username>@<domain>.com", "created_on": 9, "description": "milvus service for running sql queries", "grpc_port": 8, "host_name": "sampleMilvus", "https_port": 9, "origin": "native", "service_display_name": "sampleService", "service_id": "sampleService123", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "milvus"}], "netezza_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}], "prestissimo_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "prestissimo engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "ibm-lh-lakehouse-prestissimo-01-prestissimo-svc-cpd-instance.apps.wkclhconnectortest.cp.fyre.ibm.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-prestissimo-01-prestissimo-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "prestissimo", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "presto_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "spark_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := new(watsonxdatav2.ListEnginesOptions)
				listEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListEnginesWithContext(ctx, listEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListEngines(listEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListEnginesWithContext(ctx, listEnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engines": {"db2_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "db2 engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-db2-01-db2-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "db2"}], "milvus_services": [{"actions": ["Actions"], "created_by": "<username>@<domain>.com", "created_on": 9, "description": "milvus service for running sql queries", "grpc_port": 8, "host_name": "sampleMilvus", "https_port": 9, "origin": "native", "service_display_name": "sampleService", "service_id": "sampleService123", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "milvus"}], "netezza_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}], "prestissimo_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "prestissimo engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "ibm-lh-lakehouse-prestissimo-01-prestissimo-svc-cpd-instance.apps.wkclhconnectortest.cp.fyre.ibm.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-prestissimo-01-prestissimo-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "prestissimo", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "presto_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "spark_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}]}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := new(watsonxdatav2.ListEnginesOptions)
				listEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListEngines(listEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := new(watsonxdatav2.ListEnginesOptions)
				listEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListEngines(listEnginesOptionsModel)
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
			It(`Invoke ListEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := new(watsonxdatav2.ListEnginesOptions)
				listEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListEngines(listEnginesOptionsModel)
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
	Describe(`GetDeployments(getDeploymentsOptions *GetDeploymentsOptions) - Operation response error`, func() {
		getDeploymentsPath := "/instance"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeployments with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav2.GetDeploymentsOptions)
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deploymentresponse": {"deployment": {"cloud_type": "awq", "enable_private_endpoints": true, "enable_public_endpoints": true, "first_time_use": false, "formation_id": "new_form_id", "id": "dep_id", "plan_id": "new_plan_id", "platform_options": {"backup_encryption_key_crn": "<backup_encryption_key_crn>", "disk_encryption_key_crn": "<disk_encryption_key_crn>", "key_protect_key_id": "<key_protect_key_id>"}, "region": "us-south", "resource_group_crn": "crn:v1:staging:public:resource-controller::a/hddrtnjjj27dh38xbw::resource-group:c02a6a94f16e4ca", "type": "deployment_type", "version": "1.0.2"}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetDeployments successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav2.GetDeploymentsOptions)
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deploymentresponse": {"deployment": {"cloud_type": "awq", "enable_private_endpoints": true, "enable_public_endpoints": true, "first_time_use": false, "formation_id": "new_form_id", "id": "dep_id", "plan_id": "new_plan_id", "platform_options": {"backup_encryption_key_crn": "<backup_encryption_key_crn>", "disk_encryption_key_crn": "<disk_encryption_key_crn>", "key_protect_key_id": "<key_protect_key_id>"}, "region": "us-south", "resource_group_crn": "crn:v1:staging:public:resource-controller::a/hddrtnjjj27dh38xbw::resource-group:c02a6a94f16e4ca", "type": "deployment_type", "version": "1.0.2"}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetDeployments successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				getDeploymentsOptionsModel := new(watsonxdatav2.GetDeploymentsOptions)
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeployments with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav2.GetDeploymentsOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := new(watsonxdatav2.GetDeploymentsOptions)
				getDeploymentsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetDeployments(getDeploymentsOptionsModel)
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
	Describe(`ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions) - Operation response error`, func() {
		listNetezzaEnginesPath := "/netezza_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNetezzaEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNetezzaEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := new(watsonxdatav2.ListNetezzaEnginesOptions)
				listNetezzaEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listNetezzaEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions)`, func() {
		listNetezzaEnginesPath := "/netezza_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNetezzaEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"netezza_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListNetezzaEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := new(watsonxdatav2.ListNetezzaEnginesOptions)
				listNetezzaEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listNetezzaEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListNetezzaEnginesWithContext(ctx, listNetezzaEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListNetezzaEnginesWithContext(ctx, listNetezzaEnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listNetezzaEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"netezza_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListNetezzaEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListNetezzaEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := new(watsonxdatav2.ListNetezzaEnginesOptions)
				listNetezzaEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listNetezzaEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNetezzaEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := new(watsonxdatav2.ListNetezzaEnginesOptions)
				listNetezzaEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listNetezzaEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
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
			It(`Invoke ListNetezzaEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := new(watsonxdatav2.ListNetezzaEnginesOptions)
				listNetezzaEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listNetezzaEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptionsModel)
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
	Describe(`CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions) - Operation response error`, func() {
		createNetezzaEnginePath := "/netezza_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNetezzaEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateNetezzaEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsModel := new(watsonxdatav2.CreateNetezzaEngineOptions)
				createNetezzaEngineOptionsModel.Origin = core.StringPtr("external")
				createNetezzaEngineOptionsModel.Type = core.StringPtr("netezza")
				createNetezzaEngineOptionsModel.Description = core.StringPtr("netezza engine description")
				createNetezzaEngineOptionsModel.EngineDetails = createNetezzaEngineDetailsModel
				createNetezzaEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createNetezzaEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions)`, func() {
		createNetezzaEnginePath := "/netezza_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNetezzaEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateNetezzaEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsModel := new(watsonxdatav2.CreateNetezzaEngineOptions)
				createNetezzaEngineOptionsModel.Origin = core.StringPtr("external")
				createNetezzaEngineOptionsModel.Type = core.StringPtr("netezza")
				createNetezzaEngineOptionsModel.Description = core.StringPtr("netezza engine description")
				createNetezzaEngineOptionsModel.EngineDetails = createNetezzaEngineDetailsModel
				createNetezzaEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createNetezzaEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateNetezzaEngineWithContext(ctx, createNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateNetezzaEngineWithContext(ctx, createNetezzaEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createNetezzaEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateNetezzaEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateNetezzaEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsModel := new(watsonxdatav2.CreateNetezzaEngineOptions)
				createNetezzaEngineOptionsModel.Origin = core.StringPtr("external")
				createNetezzaEngineOptionsModel.Type = core.StringPtr("netezza")
				createNetezzaEngineOptionsModel.Description = core.StringPtr("netezza engine description")
				createNetezzaEngineOptionsModel.EngineDetails = createNetezzaEngineDetailsModel
				createNetezzaEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createNetezzaEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateNetezzaEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsModel := new(watsonxdatav2.CreateNetezzaEngineOptions)
				createNetezzaEngineOptionsModel.Origin = core.StringPtr("external")
				createNetezzaEngineOptionsModel.Type = core.StringPtr("netezza")
				createNetezzaEngineOptionsModel.Description = core.StringPtr("netezza engine description")
				createNetezzaEngineOptionsModel.EngineDetails = createNetezzaEngineDetailsModel
				createNetezzaEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createNetezzaEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateNetezzaEngineOptions model with no property values
				createNetezzaEngineOptionsModelNew := new(watsonxdatav2.CreateNetezzaEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModelNew)
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
			It(`Invoke CreateNetezzaEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsModel := new(watsonxdatav2.CreateNetezzaEngineOptions)
				createNetezzaEngineOptionsModel.Origin = core.StringPtr("external")
				createNetezzaEngineOptionsModel.Type = core.StringPtr("netezza")
				createNetezzaEngineOptionsModel.Description = core.StringPtr("netezza engine description")
				createNetezzaEngineOptionsModel.EngineDetails = createNetezzaEngineDetailsModel
				createNetezzaEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createNetezzaEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptionsModel)
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
	Describe(`DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions)`, func() {
		deleteNetezzaEnginePath := "/netezza_engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNetezzaEnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteNetezzaEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteNetezzaEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNetezzaEngineOptions model
				deleteNetezzaEngineOptionsModel := new(watsonxdatav2.DeleteNetezzaEngineOptions)
				deleteNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNetezzaEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteNetezzaEngineOptions model
				deleteNetezzaEngineOptionsModel := new(watsonxdatav2.DeleteNetezzaEngineOptions)
				deleteNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteNetezzaEngineOptions model with no property values
				deleteNetezzaEngineOptionsModelNew := new(watsonxdatav2.DeleteNetezzaEngineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions) - Operation response error`, func() {
		updateNetezzaEnginePath := "/netezza_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateNetezzaEnginePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateNetezzaEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateNetezzaEngineOptions model
				updateNetezzaEngineOptionsModel := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				updateNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions)`, func() {
		updateNetezzaEnginePath := "/netezza_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateNetezzaEnginePath))
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
					fmt.Fprintf(res, "%s", `{"netezza_engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateNetezzaEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateNetezzaEngineOptions model
				updateNetezzaEngineOptionsModel := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				updateNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateNetezzaEngineWithContext(ctx, updateNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateNetezzaEngineWithContext(ctx, updateNetezzaEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateNetezzaEnginePath))
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
					fmt.Fprintf(res, "%s", `{"netezza_engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "netezza engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "host_name": "xyz-netezza-01-netezza-svc", "origin": "ibm", "port": 4, "status": "REGISTERED", "tags": ["Tags"], "type": "netezza"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateNetezzaEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateNetezzaEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateNetezzaEngineOptions model
				updateNetezzaEngineOptionsModel := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				updateNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateNetezzaEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateNetezzaEngineOptions model
				updateNetezzaEngineOptionsModel := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				updateNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateNetezzaEngineOptions model with no property values
				updateNetezzaEngineOptionsModelNew := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModelNew)
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
			It(`Invoke UpdateNetezzaEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateNetezzaEngineOptions model
				updateNetezzaEngineOptionsModel := new(watsonxdatav2.UpdateNetezzaEngineOptions)
				updateNetezzaEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateNetezzaEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateNetezzaEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptionsModel)
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
	Describe(`ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions) - Operation response error`, func() {
		listOtherEnginesPath := "/other_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOtherEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOtherEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := new(watsonxdatav2.ListOtherEnginesOptions)
				listOtherEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listOtherEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions)`, func() {
		listOtherEnginesPath := "/other_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOtherEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"other_engines": [{"created_by": "<username>@<domain>.com", "created_on": 9, "description": "engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "engine_type": "netezza", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "registered", "status_code": 10, "tags": ["Tags"], "type": "external"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListOtherEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := new(watsonxdatav2.ListOtherEnginesOptions)
				listOtherEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listOtherEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListOtherEnginesWithContext(ctx, listOtherEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListOtherEnginesWithContext(ctx, listOtherEnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listOtherEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"other_engines": [{"created_by": "<username>@<domain>.com", "created_on": 9, "description": "engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "engine_type": "netezza", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "registered", "status_code": 10, "tags": ["Tags"], "type": "external"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListOtherEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListOtherEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := new(watsonxdatav2.ListOtherEnginesOptions)
				listOtherEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listOtherEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOtherEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := new(watsonxdatav2.ListOtherEnginesOptions)
				listOtherEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listOtherEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
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
			It(`Invoke ListOtherEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := new(watsonxdatav2.ListOtherEnginesOptions)
				listOtherEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listOtherEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListOtherEngines(listOtherEnginesOptionsModel)
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
	Describe(`CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions) - Operation response error`, func() {
		createOtherEnginePath := "/other_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOtherEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOtherEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := new(watsonxdatav2.CreateOtherEngineOptions)
				createOtherEngineOptionsModel.Description = core.StringPtr("external engine description")
				createOtherEngineOptionsModel.EngineDetails = otherEngineDetailsModel
				createOtherEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine01")
				createOtherEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions)`, func() {
		createOtherEnginePath := "/other_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOtherEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"created_by": "<username>@<domain>.com", "created_on": 9, "description": "engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "engine_type": "netezza", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "registered", "status_code": 10, "tags": ["Tags"], "type": "external"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateOtherEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := new(watsonxdatav2.CreateOtherEngineOptions)
				createOtherEngineOptionsModel.Description = core.StringPtr("external engine description")
				createOtherEngineOptionsModel.EngineDetails = otherEngineDetailsModel
				createOtherEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine01")
				createOtherEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateOtherEngineWithContext(ctx, createOtherEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateOtherEngineWithContext(ctx, createOtherEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createOtherEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"created_by": "<username>@<domain>.com", "created_on": 9, "description": "engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "engine_type": "netezza", "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "registered", "status_code": 10, "tags": ["Tags"], "type": "external"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateOtherEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateOtherEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := new(watsonxdatav2.CreateOtherEngineOptions)
				createOtherEngineOptionsModel.Description = core.StringPtr("external engine description")
				createOtherEngineOptionsModel.EngineDetails = otherEngineDetailsModel
				createOtherEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine01")
				createOtherEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateOtherEngine with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := new(watsonxdatav2.CreateOtherEngineOptions)
				createOtherEngineOptionsModel.Description = core.StringPtr("external engine description")
				createOtherEngineOptionsModel.EngineDetails = otherEngineDetailsModel
				createOtherEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine01")
				createOtherEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateOtherEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := new(watsonxdatav2.CreateOtherEngineOptions)
				createOtherEngineOptionsModel.Description = core.StringPtr("external engine description")
				createOtherEngineOptionsModel.EngineDetails = otherEngineDetailsModel
				createOtherEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine01")
				createOtherEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateOtherEngine(createOtherEngineOptionsModel)
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
	Describe(`DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions)`, func() {
		deleteOtherEnginePath := "/other_engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteOtherEnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteOtherEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteOtherEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteOtherEngineOptions model
				deleteOtherEngineOptionsModel := new(watsonxdatav2.DeleteOtherEngineOptions)
				deleteOtherEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteOtherEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteOtherEngineOptions model
				deleteOtherEngineOptionsModel := new(watsonxdatav2.DeleteOtherEngineOptions)
				deleteOtherEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteOtherEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteOtherEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteOtherEngineOptions model with no property values
				deleteOtherEngineOptionsModelNew := new(watsonxdatav2.DeleteOtherEngineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions) - Operation response error`, func() {
		listPrestoEnginesPath := "/presto_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPrestoEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := new(watsonxdatav2.ListPrestoEnginesOptions)
				listPrestoEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions)`, func() {
		listPrestoEnginesPath := "/presto_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"presto_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListPrestoEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := new(watsonxdatav2.ListPrestoEnginesOptions)
				listPrestoEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListPrestoEnginesWithContext(ctx, listPrestoEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListPrestoEnginesWithContext(ctx, listPrestoEnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"presto_engines": [{"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListPrestoEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListPrestoEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := new(watsonxdatav2.ListPrestoEnginesOptions)
				listPrestoEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPrestoEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := new(watsonxdatav2.ListPrestoEnginesOptions)
				listPrestoEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
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
			It(`Invoke ListPrestoEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := new(watsonxdatav2.ListPrestoEnginesOptions)
				listPrestoEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptionsModel)
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
	Describe(`CreateEngine(createEngineOptions *CreateEngineOptions) - Operation response error`, func() {
		createEnginePath := "/presto_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav2.CreateEngineOptions)
				createEngineOptionsModel.Origin = core.StringPtr("native")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.AssociatedCatalogs = []string{"iceberg_data", "hive_data"}
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngine(createEngineOptions *CreateEngineOptions)`, func() {
		createEnginePath := "/presto_engines"
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav2.CreateEngineOptions)
				createEngineOptionsModel.Origin = core.StringPtr("native")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.AssociatedCatalogs = []string{"iceberg_data", "hive_data"}
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav2.CreateEngineOptions)
				createEngineOptionsModel.Origin = core.StringPtr("native")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.AssociatedCatalogs = []string{"iceberg_data", "hive_data"}
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngine(createEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav2.CreateEngineOptions)
				createEngineOptionsModel.Origin = core.StringPtr("native")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.AssociatedCatalogs = []string{"iceberg_data", "hive_data"}
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
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
				createEngineOptionsModelNew := new(watsonxdatav2.CreateEngineOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsModel := new(watsonxdatav2.CreateEngineOptions)
				createEngineOptionsModel.Origin = core.StringPtr("native")
				createEngineOptionsModel.Type = core.StringPtr("presto")
				createEngineOptionsModel.AssociatedCatalogs = []string{"iceberg_data", "hive_data"}
				createEngineOptionsModel.Description = core.StringPtr("presto engine description")
				createEngineOptionsModel.EngineDetails = engineDetailsBodyModel
				createEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createEngineOptionsModel.FirstTimeUse = core.BoolPtr(true)
				createEngineOptionsModel.Region = core.StringPtr("us-south")
				createEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createEngineOptionsModel.Version = core.StringPtr("1.2.3")
				createEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngine(createEngineOptionsModel)
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
	Describe(`GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions) - Operation response error`, func() {
		getPrestoEnginePath := "/presto_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEnginePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPrestoEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineOptions model
				getPrestoEngineOptionsModel := new(watsonxdatav2.GetPrestoEngineOptions)
				getPrestoEngineOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions)`, func() {
		getPrestoEnginePath := "/presto_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEnginePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetPrestoEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetPrestoEngineOptions model
				getPrestoEngineOptionsModel := new(watsonxdatav2.GetPrestoEngineOptions)
				getPrestoEngineOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetPrestoEngineWithContext(ctx, getPrestoEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetPrestoEngineWithContext(ctx, getPrestoEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEnginePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetPrestoEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetPrestoEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPrestoEngineOptions model
				getPrestoEngineOptionsModel := new(watsonxdatav2.GetPrestoEngineOptions)
				getPrestoEngineOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPrestoEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineOptions model
				getPrestoEngineOptionsModel := new(watsonxdatav2.GetPrestoEngineOptions)
				getPrestoEngineOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPrestoEngineOptions model with no property values
				getPrestoEngineOptionsModelNew := new(watsonxdatav2.GetPrestoEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModelNew)
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
			It(`Invoke GetPrestoEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineOptions model
				getPrestoEngineOptionsModel := new(watsonxdatav2.GetPrestoEngineOptions)
				getPrestoEngineOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetPrestoEngine(getPrestoEngineOptionsModel)
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
	Describe(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
		deleteEnginePath := "/presto_engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				deleteEngineOptionsModel := new(watsonxdatav2.DeleteEngineOptions)
				deleteEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteEngine(deleteEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteEngineOptions model
				deleteEngineOptionsModel := new(watsonxdatav2.DeleteEngineOptions)
				deleteEngineOptionsModel.EngineID = core.StringPtr("testString")
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
				deleteEngineOptionsModelNew := new(watsonxdatav2.DeleteEngineOptions)
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
	Describe(`UpdateEngine(updateEngineOptions *UpdateEngineOptions) - Operation response error`, func() {
		updateEnginePath := "/presto_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnginePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav2.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEngine(updateEngineOptions *UpdateEngineOptions)`, func() {
		updateEnginePath := "/presto_engines/testString"
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav2.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "associated_catalogs": ["AssociatedCatalogs"], "build_version": "1.0.3.0.0", "coordinator": {"node_type": "worker", "quantity": 8}, "created_by": "<username>@<domain>.com", "created_on": 9, "description": "presto engine for running sql queries", "engine_details": {"connection_string": "1.2.3.4", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}, "metastore_host": "1.2.3.4"}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "external_host_name": "your-hostname.apps.your-domain.com", "group_id": "new_group_id", "host_name": "ibm-lh-lakehouse-presto-01-presto-svc", "origin": "ibm", "port": 4, "region": "us-south", "size_config": "starter", "status": "running", "status_code": 10, "tags": ["Tags"], "type": "presto", "version": "1.2.0", "worker": {"node_type": "worker", "quantity": 8}}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav2.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateEngine(updateEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav2.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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
				updateEngineOptionsModelNew := new(watsonxdatav2.UpdateEngineOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateEngineOptions model
				updateEngineOptionsModel := new(watsonxdatav2.UpdateEngineOptions)
				updateEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateEngine(updateEngineOptionsModel)
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
	Describe(`ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions) - Operation response error`, func() {
		listPrestoEngineCatalogsPath := "/presto_engines/testString/catalogs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPrestoEngineCatalogs with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				listPrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				listPrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions)`, func() {
		listPrestoEngineCatalogsPath := "/presto_engines/testString/catalogs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "creation_date": "16073847388"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListPrestoEngineCatalogs successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				listPrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				listPrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListPrestoEngineCatalogsWithContext(ctx, listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListPrestoEngineCatalogsWithContext(ctx, listPrestoEngineCatalogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "creation_date": "16073847388"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListPrestoEngineCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListPrestoEngineCatalogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				listPrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				listPrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPrestoEngineCatalogs with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				listPrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				listPrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPrestoEngineCatalogsOptions model with no property values
				listPrestoEngineCatalogsOptionsModelNew := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModelNew)
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
			It(`Invoke ListPrestoEngineCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				listPrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ListPrestoEngineCatalogsOptions)
				listPrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listPrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptionsModel)
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
	Describe(`ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptions *ReplacePrestoEngineCatalogsOptions) - Operation response error`, func() {
		replacePrestoEngineCatalogsPath := "/presto_engines/testString/catalogs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_names"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplacePrestoEngineCatalogs with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				replacePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				replacePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptions *ReplacePrestoEngineCatalogsOptions)`, func() {
		replacePrestoEngineCatalogsPath := "/presto_engines/testString/catalogs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_names"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "creation_date": "16073847388"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ReplacePrestoEngineCatalogs successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				replacePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				replacePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ReplacePrestoEngineCatalogsWithContext(ctx, replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ReplacePrestoEngineCatalogsWithContext(ctx, replacePrestoEngineCatalogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replacePrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_names"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"catalog_name": "sampleCatalog", "creation_date": "16073847388"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ReplacePrestoEngineCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ReplacePrestoEngineCatalogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				replacePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				replacePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplacePrestoEngineCatalogs with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				replacePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				replacePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplacePrestoEngineCatalogsOptions model with no property values
				replacePrestoEngineCatalogsOptionsModelNew := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModelNew)
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
			It(`Invoke ReplacePrestoEngineCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				replacePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.ReplacePrestoEngineCatalogsOptions)
				replacePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replacePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptionsModel)
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
	Describe(`DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions)`, func() {
		deletePrestoEngineCatalogsPath := "/presto_engines/testString/catalogs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePrestoEngineCatalogsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["catalog_names"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePrestoEngineCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeletePrestoEngineCatalogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePrestoEngineCatalogsOptions model
				deletePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.DeletePrestoEngineCatalogsOptions)
				deletePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePrestoEngineCatalogs with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeletePrestoEngineCatalogsOptions model
				deletePrestoEngineCatalogsOptionsModel := new(watsonxdatav2.DeletePrestoEngineCatalogsOptions)
				deletePrestoEngineCatalogsOptionsModel.EngineID = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.CatalogNames = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deletePrestoEngineCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePrestoEngineCatalogsOptions model with no property values
				deletePrestoEngineCatalogsOptionsModelNew := new(watsonxdatav2.DeletePrestoEngineCatalogsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions) - Operation response error`, func() {
		getPrestoEngineCatalogPath := "/presto_engines/testString/catalogs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEngineCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPrestoEngineCatalog with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineCatalogOptions model
				getPrestoEngineCatalogOptionsModel := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				getPrestoEngineCatalogOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions)`, func() {
		getPrestoEngineCatalogPath := "/presto_engines/testString/catalogs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEngineCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog": {"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetPrestoEngineCatalog successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetPrestoEngineCatalogOptions model
				getPrestoEngineCatalogOptionsModel := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				getPrestoEngineCatalogOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetPrestoEngineCatalogWithContext(ctx, getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetPrestoEngineCatalogWithContext(ctx, getPrestoEngineCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPrestoEngineCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog": {"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetPrestoEngineCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetPrestoEngineCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPrestoEngineCatalogOptions model
				getPrestoEngineCatalogOptionsModel := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				getPrestoEngineCatalogOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPrestoEngineCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineCatalogOptions model
				getPrestoEngineCatalogOptionsModel := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				getPrestoEngineCatalogOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPrestoEngineCatalogOptions model with no property values
				getPrestoEngineCatalogOptionsModelNew := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModelNew)
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
			It(`Invoke GetPrestoEngineCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetPrestoEngineCatalogOptions model
				getPrestoEngineCatalogOptionsModel := new(watsonxdatav2.GetPrestoEngineCatalogOptions)
				getPrestoEngineCatalogOptionsModel.EngineID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getPrestoEngineCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptionsModel)
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
	Describe(`CreateEnginePause(createEnginePauseOptions *CreateEnginePauseOptions) - Operation response error`, func() {
		createEnginePausePath := "/presto_engines/testString/pause"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePausePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnginePause with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEnginePauseOptions model
				createEnginePauseOptionsModel := new(watsonxdatav2.CreateEnginePauseOptions)
				createEnginePauseOptionsModel.EngineID = core.StringPtr("testString")
				createEnginePauseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEnginePauseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnginePause(createEnginePauseOptions *CreateEnginePauseOptions)`, func() {
		createEnginePausePath := "/presto_engines/testString/pause"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePausePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEnginePause successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateEnginePauseOptions model
				createEnginePauseOptionsModel := new(watsonxdatav2.CreateEnginePauseOptions)
				createEnginePauseOptionsModel.EngineID = core.StringPtr("testString")
				createEnginePauseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEnginePauseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEnginePauseWithContext(ctx, createEnginePauseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEnginePauseWithContext(ctx, createEnginePauseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEnginePausePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEnginePause successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEnginePause(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnginePauseOptions model
				createEnginePauseOptionsModel := new(watsonxdatav2.CreateEnginePauseOptions)
				createEnginePauseOptionsModel.EngineID = core.StringPtr("testString")
				createEnginePauseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEnginePauseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEnginePause with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEnginePauseOptions model
				createEnginePauseOptionsModel := new(watsonxdatav2.CreateEnginePauseOptions)
				createEnginePauseOptionsModel.EngineID = core.StringPtr("testString")
				createEnginePauseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEnginePauseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnginePauseOptions model with no property values
				createEnginePauseOptionsModelNew := new(watsonxdatav2.CreateEnginePauseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEnginePause(createEnginePauseOptionsModelNew)
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
			It(`Invoke CreateEnginePause successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEnginePauseOptions model
				createEnginePauseOptionsModel := new(watsonxdatav2.CreateEnginePauseOptions)
				createEnginePauseOptionsModel.EngineID = core.StringPtr("testString")
				createEnginePauseOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEnginePauseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEnginePause(createEnginePauseOptionsModel)
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
	Describe(`RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions) - Operation response error`, func() {
		runExplainStatementPath := "/presto_engines/testString/query_explain"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runExplainStatementPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunExplainStatement with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainStatementOptions model
				runExplainStatementOptionsModel := new(watsonxdatav2.RunExplainStatementOptions)
				runExplainStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainStatementOptionsModel.Format = core.StringPtr("json")
				runExplainStatementOptionsModel.Type = core.StringPtr("io")
				runExplainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions)`, func() {
		runExplainStatementPath := "/presto_engines/testString/query_explain"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runExplainStatementPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "result": "Result"}`)
				}))
			})
			It(`Invoke RunExplainStatement successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RunExplainStatementOptions model
				runExplainStatementOptionsModel := new(watsonxdatav2.RunExplainStatementOptions)
				runExplainStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainStatementOptionsModel.Format = core.StringPtr("json")
				runExplainStatementOptionsModel.Type = core.StringPtr("io")
				runExplainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.RunExplainStatementWithContext(ctx, runExplainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.RunExplainStatementWithContext(ctx, runExplainStatementOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(runExplainStatementPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "result": "Result"}`)
				}))
			})
			It(`Invoke RunExplainStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.RunExplainStatement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RunExplainStatementOptions model
				runExplainStatementOptionsModel := new(watsonxdatav2.RunExplainStatementOptions)
				runExplainStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainStatementOptionsModel.Format = core.StringPtr("json")
				runExplainStatementOptionsModel.Type = core.StringPtr("io")
				runExplainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RunExplainStatement with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainStatementOptions model
				runExplainStatementOptionsModel := new(watsonxdatav2.RunExplainStatementOptions)
				runExplainStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainStatementOptionsModel.Format = core.StringPtr("json")
				runExplainStatementOptionsModel.Type = core.StringPtr("io")
				runExplainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunExplainStatementOptions model with no property values
				runExplainStatementOptionsModelNew := new(watsonxdatav2.RunExplainStatementOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.RunExplainStatement(runExplainStatementOptionsModelNew)
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
			It(`Invoke RunExplainStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainStatementOptions model
				runExplainStatementOptionsModel := new(watsonxdatav2.RunExplainStatementOptions)
				runExplainStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainStatementOptionsModel.Format = core.StringPtr("json")
				runExplainStatementOptionsModel.Type = core.StringPtr("io")
				runExplainStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.RunExplainStatement(runExplainStatementOptionsModel)
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
	Describe(`RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions) - Operation response error`, func() {
		runExplainAnalyzeStatementPath := "/presto_engines/testString/query_explain_analyze"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runExplainAnalyzeStatementPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunExplainAnalyzeStatement with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				runExplainAnalyzeStatementOptionsModel := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				runExplainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				runExplainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions)`, func() {
		runExplainAnalyzeStatementPath := "/presto_engines/testString/query_explain_analyze"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runExplainAnalyzeStatementPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "result": "Result"}`)
				}))
			})
			It(`Invoke RunExplainAnalyzeStatement successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				runExplainAnalyzeStatementOptionsModel := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				runExplainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				runExplainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.RunExplainAnalyzeStatementWithContext(ctx, runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.RunExplainAnalyzeStatementWithContext(ctx, runExplainAnalyzeStatementOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(runExplainAnalyzeStatementPath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "result": "Result"}`)
				}))
			})
			It(`Invoke RunExplainAnalyzeStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.RunExplainAnalyzeStatement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				runExplainAnalyzeStatementOptionsModel := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				runExplainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				runExplainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RunExplainAnalyzeStatement with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				runExplainAnalyzeStatementOptionsModel := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				runExplainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				runExplainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunExplainAnalyzeStatementOptions model with no property values
				runExplainAnalyzeStatementOptionsModelNew := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModelNew)
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
			It(`Invoke RunExplainAnalyzeStatement successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				runExplainAnalyzeStatementOptionsModel := new(watsonxdatav2.RunExplainAnalyzeStatementOptions)
				runExplainAnalyzeStatementOptionsModel.EngineID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Statement = core.StringPtr("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.Verbose = core.BoolPtr(true)
				runExplainAnalyzeStatementOptionsModel.AuthInstanceID = core.StringPtr("testString")
				runExplainAnalyzeStatementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptionsModel)
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
	Describe(`CreateEngineRestart(createEngineRestartOptions *CreateEngineRestartOptions) - Operation response error`, func() {
		createEngineRestartPath := "/presto_engines/testString/restart"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineRestartPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEngineRestart with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineRestartOptions model
				createEngineRestartOptionsModel := new(watsonxdatav2.CreateEngineRestartOptions)
				createEngineRestartOptionsModel.EngineID = core.StringPtr("testString")
				createEngineRestartOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineRestartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngineRestart(createEngineRestartOptions *CreateEngineRestartOptions)`, func() {
		createEngineRestartPath := "/presto_engines/testString/restart"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineRestartPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineRestart successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateEngineRestartOptions model
				createEngineRestartOptionsModel := new(watsonxdatav2.CreateEngineRestartOptions)
				createEngineRestartOptionsModel.EngineID = core.StringPtr("testString")
				createEngineRestartOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineRestartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEngineRestartWithContext(ctx, createEngineRestartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEngineRestartWithContext(ctx, createEngineRestartOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEngineRestartPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineRestart successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEngineRestart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEngineRestartOptions model
				createEngineRestartOptionsModel := new(watsonxdatav2.CreateEngineRestartOptions)
				createEngineRestartOptionsModel.EngineID = core.StringPtr("testString")
				createEngineRestartOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineRestartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngineRestart with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineRestartOptions model
				createEngineRestartOptionsModel := new(watsonxdatav2.CreateEngineRestartOptions)
				createEngineRestartOptionsModel.EngineID = core.StringPtr("testString")
				createEngineRestartOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineRestartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEngineRestartOptions model with no property values
				createEngineRestartOptionsModelNew := new(watsonxdatav2.CreateEngineRestartOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModelNew)
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
			It(`Invoke CreateEngineRestart successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineRestartOptions model
				createEngineRestartOptionsModel := new(watsonxdatav2.CreateEngineRestartOptions)
				createEngineRestartOptionsModel.EngineID = core.StringPtr("testString")
				createEngineRestartOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineRestartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngineRestart(createEngineRestartOptionsModel)
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
	Describe(`CreateEngineResume(createEngineResumeOptions *CreateEngineResumeOptions) - Operation response error`, func() {
		createEngineResumePath := "/presto_engines/testString/resume"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineResumePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEngineResume with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineResumeOptions model
				createEngineResumeOptionsModel := new(watsonxdatav2.CreateEngineResumeOptions)
				createEngineResumeOptionsModel.EngineID = core.StringPtr("testString")
				createEngineResumeOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineResumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngineResume(createEngineResumeOptions *CreateEngineResumeOptions)`, func() {
		createEngineResumePath := "/presto_engines/testString/resume"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineResumePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineResume successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateEngineResumeOptions model
				createEngineResumeOptionsModel := new(watsonxdatav2.CreateEngineResumeOptions)
				createEngineResumeOptionsModel.EngineID = core.StringPtr("testString")
				createEngineResumeOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineResumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEngineResumeWithContext(ctx, createEngineResumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEngineResumeWithContext(ctx, createEngineResumeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEngineResumePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineResume successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEngineResume(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEngineResumeOptions model
				createEngineResumeOptionsModel := new(watsonxdatav2.CreateEngineResumeOptions)
				createEngineResumeOptionsModel.EngineID = core.StringPtr("testString")
				createEngineResumeOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineResumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngineResume with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineResumeOptions model
				createEngineResumeOptionsModel := new(watsonxdatav2.CreateEngineResumeOptions)
				createEngineResumeOptionsModel.EngineID = core.StringPtr("testString")
				createEngineResumeOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineResumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEngineResumeOptions model with no property values
				createEngineResumeOptionsModelNew := new(watsonxdatav2.CreateEngineResumeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEngineResume(createEngineResumeOptionsModelNew)
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
			It(`Invoke CreateEngineResume successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateEngineResumeOptions model
				createEngineResumeOptionsModel := new(watsonxdatav2.CreateEngineResumeOptions)
				createEngineResumeOptionsModel.EngineID = core.StringPtr("testString")
				createEngineResumeOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineResumeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngineResume(createEngineResumeOptionsModel)
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
	Describe(`CreateEngineScale(createEngineScaleOptions *CreateEngineScaleOptions) - Operation response error`, func() {
		createEngineScalePath := "/presto_engines/testString/scale"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineScalePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEngineScale with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateEngineScaleOptions model
				createEngineScaleOptionsModel := new(watsonxdatav2.CreateEngineScaleOptions)
				createEngineScaleOptionsModel.EngineID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Coordinator = nodeDescriptionModel
				createEngineScaleOptionsModel.Worker = nodeDescriptionModel
				createEngineScaleOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEngineScale(createEngineScaleOptions *CreateEngineScaleOptions)`, func() {
		createEngineScalePath := "/presto_engines/testString/scale"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEngineScalePath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineScale successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateEngineScaleOptions model
				createEngineScaleOptionsModel := new(watsonxdatav2.CreateEngineScaleOptions)
				createEngineScaleOptionsModel.EngineID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Coordinator = nodeDescriptionModel
				createEngineScaleOptionsModel.Worker = nodeDescriptionModel
				createEngineScaleOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateEngineScaleWithContext(ctx, createEngineScaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateEngineScaleWithContext(ctx, createEngineScaleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEngineScalePath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateEngineScale successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateEngineScale(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateEngineScaleOptions model
				createEngineScaleOptionsModel := new(watsonxdatav2.CreateEngineScaleOptions)
				createEngineScaleOptionsModel.EngineID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Coordinator = nodeDescriptionModel
				createEngineScaleOptionsModel.Worker = nodeDescriptionModel
				createEngineScaleOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEngineScale with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateEngineScaleOptions model
				createEngineScaleOptionsModel := new(watsonxdatav2.CreateEngineScaleOptions)
				createEngineScaleOptionsModel.EngineID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Coordinator = nodeDescriptionModel
				createEngineScaleOptionsModel.Worker = nodeDescriptionModel
				createEngineScaleOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEngineScaleOptions model with no property values
				createEngineScaleOptionsModelNew := new(watsonxdatav2.CreateEngineScaleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateEngineScale(createEngineScaleOptionsModelNew)
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
			It(`Invoke CreateEngineScale successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateEngineScaleOptions model
				createEngineScaleOptionsModel := new(watsonxdatav2.CreateEngineScaleOptions)
				createEngineScaleOptionsModel.EngineID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Coordinator = nodeDescriptionModel
				createEngineScaleOptionsModel.Worker = nodeDescriptionModel
				createEngineScaleOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createEngineScaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateEngineScale(createEngineScaleOptionsModel)
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
	Describe(`ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions) - Operation response error`, func() {
		listSparkEnginesPath := "/spark_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEnginesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSparkEngines with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := new(watsonxdatav2.ListSparkEnginesOptions)
				listSparkEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions)`, func() {
		listSparkEnginesPath := "/spark_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "spark_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}]}`)
				}))
			})
			It(`Invoke ListSparkEngines successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := new(watsonxdatav2.ListSparkEnginesOptions)
				listSparkEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListSparkEnginesWithContext(ctx, listSparkEnginesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListSparkEnginesWithContext(ctx, listSparkEnginesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEnginesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "spark_engines": [{"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}]}`)
				}))
			})
			It(`Invoke ListSparkEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListSparkEngines(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := new(watsonxdatav2.ListSparkEnginesOptions)
				listSparkEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSparkEngines with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := new(watsonxdatav2.ListSparkEnginesOptions)
				listSparkEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
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
			It(`Invoke ListSparkEngines successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := new(watsonxdatav2.ListSparkEnginesOptions)
				listSparkEnginesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEnginesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListSparkEngines(listSparkEnginesOptionsModel)
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
	Describe(`CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions) - Operation response error`, func() {
		createSparkEnginePath := "/spark_engines"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEnginePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSparkEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsModel := new(watsonxdatav2.CreateSparkEngineOptions)
				createSparkEngineOptionsModel.Origin = core.StringPtr("native")
				createSparkEngineOptionsModel.Type = core.StringPtr("spark")
				createSparkEngineOptionsModel.Description = core.StringPtr("spark engine description")
				createSparkEngineOptionsModel.EngineDetails = sparkEngineDetailsPrototypeModel
				createSparkEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createSparkEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions)`, func() {
		createSparkEnginePath := "/spark_engines"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateSparkEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsModel := new(watsonxdatav2.CreateSparkEngineOptions)
				createSparkEngineOptionsModel.Origin = core.StringPtr("native")
				createSparkEngineOptionsModel.Type = core.StringPtr("spark")
				createSparkEngineOptionsModel.Description = core.StringPtr("spark engine description")
				createSparkEngineOptionsModel.EngineDetails = sparkEngineDetailsPrototypeModel
				createSparkEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createSparkEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateSparkEngineWithContext(ctx, createSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateSparkEngineWithContext(ctx, createSparkEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateSparkEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateSparkEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsModel := new(watsonxdatav2.CreateSparkEngineOptions)
				createSparkEngineOptionsModel.Origin = core.StringPtr("native")
				createSparkEngineOptionsModel.Type = core.StringPtr("spark")
				createSparkEngineOptionsModel.Description = core.StringPtr("spark engine description")
				createSparkEngineOptionsModel.EngineDetails = sparkEngineDetailsPrototypeModel
				createSparkEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createSparkEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSparkEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsModel := new(watsonxdatav2.CreateSparkEngineOptions)
				createSparkEngineOptionsModel.Origin = core.StringPtr("native")
				createSparkEngineOptionsModel.Type = core.StringPtr("spark")
				createSparkEngineOptionsModel.Description = core.StringPtr("spark engine description")
				createSparkEngineOptionsModel.EngineDetails = sparkEngineDetailsPrototypeModel
				createSparkEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createSparkEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSparkEngineOptions model with no property values
				createSparkEngineOptionsModelNew := new(watsonxdatav2.CreateSparkEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModelNew)
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
			It(`Invoke CreateSparkEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsModel := new(watsonxdatav2.CreateSparkEngineOptions)
				createSparkEngineOptionsModel.Origin = core.StringPtr("native")
				createSparkEngineOptionsModel.Type = core.StringPtr("spark")
				createSparkEngineOptionsModel.Description = core.StringPtr("spark engine description")
				createSparkEngineOptionsModel.EngineDetails = sparkEngineDetailsPrototypeModel
				createSparkEngineOptionsModel.EngineDisplayName = core.StringPtr("sampleEngine")
				createSparkEngineOptionsModel.Tags = []string{"tag1", "tag2"}
				createSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateSparkEngine(createSparkEngineOptionsModel)
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
	Describe(`DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions)`, func() {
		deleteSparkEnginePath := "/spark_engines/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSparkEnginePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSparkEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteSparkEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSparkEngineOptions model
				deleteSparkEngineOptionsModel := new(watsonxdatav2.DeleteSparkEngineOptions)
				deleteSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSparkEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteSparkEngineOptions model
				deleteSparkEngineOptionsModel := new(watsonxdatav2.DeleteSparkEngineOptions)
				deleteSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				deleteSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSparkEngineOptions model with no property values
				deleteSparkEngineOptionsModelNew := new(watsonxdatav2.DeleteSparkEngineOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions) - Operation response error`, func() {
		updateSparkEnginePath := "/spark_engines/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSparkEnginePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSparkEngine with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSparkEngineOptions model
				updateSparkEngineOptionsModel := new(watsonxdatav2.UpdateSparkEngineOptions)
				updateSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions)`, func() {
		updateSparkEnginePath := "/spark_engines/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSparkEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateSparkEngine successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSparkEngineOptions model
				updateSparkEngineOptionsModel := new(watsonxdatav2.UpdateSparkEngineOptions)
				updateSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateSparkEngineWithContext(ctx, updateSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateSparkEngineWithContext(ctx, updateSparkEngineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSparkEnginePath))
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
					fmt.Fprintf(res, "%s", `{"engine": {"actions": ["Actions"], "build_version": "1.0.3.0.0", "created_by": "<username>@<domain>.com", "created_on": 9, "description": "spark engine for running sql queries", "engine_details": {"connection_string": "https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>", "endpoints": {"applications_api": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>", "history_server_endpoint": "$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server", "spark_access_endpoint": "$HOST/analytics-engine/details/spark-<instance_id>", "spark_jobs_v4_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications", "spark_kernel_endpoint": "$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels", "view_history_server": "ViewHistoryServer", "wxd_application_endpoint": "$HOST/v1/1698311655308796/engines/spark817/applications"}}, "engine_display_name": "sampleEngine", "engine_id": "sampleEngine123", "origin": "ibm", "status": "Registered", "tags": ["Tags"], "type": "spark"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateSparkEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateSparkEngine(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSparkEngineOptions model
				updateSparkEngineOptionsModel := new(watsonxdatav2.UpdateSparkEngineOptions)
				updateSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSparkEngine with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSparkEngineOptions model
				updateSparkEngineOptionsModel := new(watsonxdatav2.UpdateSparkEngineOptions)
				updateSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSparkEngineOptions model with no property values
				updateSparkEngineOptionsModelNew := new(watsonxdatav2.UpdateSparkEngineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModelNew)
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
			It(`Invoke UpdateSparkEngine successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSparkEngineOptions model
				updateSparkEngineOptionsModel := new(watsonxdatav2.UpdateSparkEngineOptions)
				updateSparkEngineOptionsModel.EngineID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSparkEngineOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSparkEngineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptionsModel)
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
	Describe(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) - Operation response error`, func() {
		listSparkEngineApplicationsPath := "/spark_engines/testString/applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEngineApplicationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSparkEngineApplications with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEngineApplicationsOptions model
				listSparkEngineApplicationsOptionsModel := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				listSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions)`, func() {
		listSparkEngineApplicationsPath := "/spark_engines/testString/applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEngineApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"application_id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "auto_termination_time": "2020-12-08T10:00:00.000Z", "creation_time": "2020-12-08T10:00:00.000Z", "end_time": "2020-12-08T10:00:00.000Z", "failed_time": "2020-12-08T10:00:00.000Z", "finish_time": "2020-12-08T10:00:00.000Z", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "runtime": {"spark_version": "3.3"}, "spark_application_id": "application_16073847388_0001", "spark_application_name": "SparkApplicationName", "start_time": "2020-12-08T10:00:00.000Z", "state": "RUNNING", "submission_time": "2023-11-01T11:18:49.758Z", "template_id": "spark-3.3-jaas-v2-cp4d-template"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListSparkEngineApplications successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListSparkEngineApplicationsOptions model
				listSparkEngineApplicationsOptionsModel := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				listSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListSparkEngineApplicationsWithContext(ctx, listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListSparkEngineApplicationsWithContext(ctx, listSparkEngineApplicationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSparkEngineApplicationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applications": [{"application_id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "auto_termination_time": "2020-12-08T10:00:00.000Z", "creation_time": "2020-12-08T10:00:00.000Z", "end_time": "2020-12-08T10:00:00.000Z", "failed_time": "2020-12-08T10:00:00.000Z", "finish_time": "2020-12-08T10:00:00.000Z", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "runtime": {"spark_version": "3.3"}, "spark_application_id": "application_16073847388_0001", "spark_application_name": "SparkApplicationName", "start_time": "2020-12-08T10:00:00.000Z", "state": "RUNNING", "submission_time": "2023-11-01T11:18:49.758Z", "template_id": "spark-3.3-jaas-v2-cp4d-template"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListSparkEngineApplications successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListSparkEngineApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSparkEngineApplicationsOptions model
				listSparkEngineApplicationsOptionsModel := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				listSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSparkEngineApplications with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEngineApplicationsOptions model
				listSparkEngineApplicationsOptionsModel := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				listSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSparkEngineApplicationsOptions model with no property values
				listSparkEngineApplicationsOptionsModelNew := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModelNew)
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
			It(`Invoke ListSparkEngineApplications successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSparkEngineApplicationsOptions model
				listSparkEngineApplicationsOptionsModel := new(watsonxdatav2.ListSparkEngineApplicationsOptions)
				listSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptionsModel)
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
	Describe(`CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions) - Operation response error`, func() {
		createSparkEngineApplicationPath := "/spark_engines/testString/applications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEngineApplicationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSparkEngineApplication with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				createSparkEngineApplicationOptionsModel := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				createSparkEngineApplicationOptionsModel.EngineID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.ApplicationDetails = sparkApplicationDetailsModel
				createSparkEngineApplicationOptionsModel.JobEndpoint = core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Type = core.StringPtr("iae")
				createSparkEngineApplicationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions)`, func() {
		createSparkEngineApplicationPath := "/spark_engines/testString/applications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEngineApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "spark_engine_application": {"application_id": "23c99c14-3af8-467d-9703-cc8163c60d35", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "state": "ACCEPTED"}}`)
				}))
			})
			It(`Invoke CreateSparkEngineApplication successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				createSparkEngineApplicationOptionsModel := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				createSparkEngineApplicationOptionsModel.EngineID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.ApplicationDetails = sparkApplicationDetailsModel
				createSparkEngineApplicationOptionsModel.JobEndpoint = core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Type = core.StringPtr("iae")
				createSparkEngineApplicationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.CreateSparkEngineApplicationWithContext(ctx, createSparkEngineApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.CreateSparkEngineApplicationWithContext(ctx, createSparkEngineApplicationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSparkEngineApplicationPath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "spark_engine_application": {"application_id": "23c99c14-3af8-467d-9703-cc8163c60d35", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "state": "ACCEPTED"}}`)
				}))
			})
			It(`Invoke CreateSparkEngineApplication successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.CreateSparkEngineApplication(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				createSparkEngineApplicationOptionsModel := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				createSparkEngineApplicationOptionsModel.EngineID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.ApplicationDetails = sparkApplicationDetailsModel
				createSparkEngineApplicationOptionsModel.JobEndpoint = core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Type = core.StringPtr("iae")
				createSparkEngineApplicationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSparkEngineApplication with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				createSparkEngineApplicationOptionsModel := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				createSparkEngineApplicationOptionsModel.EngineID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.ApplicationDetails = sparkApplicationDetailsModel
				createSparkEngineApplicationOptionsModel.JobEndpoint = core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Type = core.StringPtr("iae")
				createSparkEngineApplicationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSparkEngineApplicationOptions model with no property values
				createSparkEngineApplicationOptionsModelNew := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModelNew)
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
			It(`Invoke CreateSparkEngineApplication successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				createSparkEngineApplicationOptionsModel := new(watsonxdatav2.CreateSparkEngineApplicationOptions)
				createSparkEngineApplicationOptionsModel.EngineID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.ApplicationDetails = sparkApplicationDetailsModel
				createSparkEngineApplicationOptionsModel.JobEndpoint = core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Type = core.StringPtr("iae")
				createSparkEngineApplicationOptionsModel.AuthInstanceID = core.StringPtr("testString")
				createSparkEngineApplicationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptionsModel)
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
	Describe(`DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions)`, func() {
		deleteSparkEngineApplicationsPath := "/spark_engines/testString/applications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSparkEngineApplicationsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["application_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSparkEngineApplications successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteSparkEngineApplications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSparkEngineApplicationsOptions model
				deleteSparkEngineApplicationsOptionsModel := new(watsonxdatav2.DeleteSparkEngineApplicationsOptions)
				deleteSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSparkEngineApplications with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteSparkEngineApplicationsOptions model
				deleteSparkEngineApplicationsOptionsModel := new(watsonxdatav2.DeleteSparkEngineApplicationsOptions)
				deleteSparkEngineApplicationsOptionsModel.EngineID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSparkEngineApplicationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := watsonxDataService.DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSparkEngineApplicationsOptions model with no property values
				deleteSparkEngineApplicationsOptionsModelNew := new(watsonxdatav2.DeleteSparkEngineApplicationsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = watsonxDataService.DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions) - Operation response error`, func() {
		getSparkEngineApplicationStatusPath := "/spark_engines/testString/applications/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSparkEngineApplicationStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSparkEngineApplicationStatus with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				getSparkEngineApplicationStatusOptionsModel := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				getSparkEngineApplicationStatusOptionsModel.EngineID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.ApplicationID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions)`, func() {
		getSparkEngineApplicationStatusPath := "/spark_engines/testString/applications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSparkEngineApplicationStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application": {"application_details": {"application": "/opt/ibm/spark/examples/src/main/python/wordcount.py", "arguments": ["Arguments"], "conf": {"spark_app_name": "MyJob", "spark_hive_metastore_client_auth_mode": "PLAIN", "spark_hive_metastore_client_plain_password": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...", "spark_hive_metastore_client_plain_username": "ibm_lh_token_admin", "spark_hive_metastore_truststore_password": "changeit", "spark_hive_metastore_truststore_path": "file:///opt/ibm/jdk/lib/security/cacerts", "spark_hive_metastore_truststore_type": "JKS", "spark_hive_metastore_use_ssl": "true", "spark_sql_catalog_implementation": "hive", "spark_sql_catalog_lakehouse": "org.apache.iceberg.spark.SparkCatalog", "spark_sql_catalog_lakehouse_type": "hive", "spark_sql_catalog_lakehouse_uri": "SparkSqlCatalogLakehouseURI", "spark_sql_extensions": "org.apache.iceberg.spark.extensions.IcebergSparkSessionExtensions", "spark_sql_iceberg_vectorization_enabled": "false"}, "env": {"anyKey": "anyValue"}, "name": "SparkApplication1"}, "application_id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "auto_termination_time": "2020-12-08T10:00:00.000Z", "creation_time": "Saturday 28 October 2023 07:17:06.856+0000", "deploy_mode": "stand-alone", "end_time": "2020-12-08T10:00:00.000Z", "failed_time": "FailedTime", "finish_time": "Saturday 28 October 2023 07:17:38.966+0000", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "return_code": "0", "spark_application_id": "app-20231028071726-0000", "spark_application_name": "PythonWordCount", "start_time": "Saturday 28 October 2023 07:17:26.649+0000", "state": "FINISHED", "state_details": [{"code": "Code", "message": "Message", "type": "Type"}], "submission_time": "2023-11-01T11:18:49.758Z", "template_id": "spark-3.3-jaas-v2-cp4d-template"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetSparkEngineApplicationStatus successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				getSparkEngineApplicationStatusOptionsModel := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				getSparkEngineApplicationStatusOptionsModel.EngineID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.ApplicationID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetSparkEngineApplicationStatusWithContext(ctx, getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetSparkEngineApplicationStatusWithContext(ctx, getSparkEngineApplicationStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSparkEngineApplicationStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"application": {"application_details": {"application": "/opt/ibm/spark/examples/src/main/python/wordcount.py", "arguments": ["Arguments"], "conf": {"spark_app_name": "MyJob", "spark_hive_metastore_client_auth_mode": "PLAIN", "spark_hive_metastore_client_plain_password": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...", "spark_hive_metastore_client_plain_username": "ibm_lh_token_admin", "spark_hive_metastore_truststore_password": "changeit", "spark_hive_metastore_truststore_path": "file:///opt/ibm/jdk/lib/security/cacerts", "spark_hive_metastore_truststore_type": "JKS", "spark_hive_metastore_use_ssl": "true", "spark_sql_catalog_implementation": "hive", "spark_sql_catalog_lakehouse": "org.apache.iceberg.spark.SparkCatalog", "spark_sql_catalog_lakehouse_type": "hive", "spark_sql_catalog_lakehouse_uri": "SparkSqlCatalogLakehouseURI", "spark_sql_extensions": "org.apache.iceberg.spark.extensions.IcebergSparkSessionExtensions", "spark_sql_iceberg_vectorization_enabled": "false"}, "env": {"anyKey": "anyValue"}, "name": "SparkApplication1"}, "application_id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "auto_termination_time": "2020-12-08T10:00:00.000Z", "creation_time": "Saturday 28 October 2023 07:17:06.856+0000", "deploy_mode": "stand-alone", "end_time": "2020-12-08T10:00:00.000Z", "failed_time": "FailedTime", "finish_time": "Saturday 28 October 2023 07:17:38.966+0000", "id": "cd7cbf1f-8893-4c51-aa3d-d92729f05e99", "return_code": "0", "spark_application_id": "app-20231028071726-0000", "spark_application_name": "PythonWordCount", "start_time": "Saturday 28 October 2023 07:17:26.649+0000", "state": "FINISHED", "state_details": [{"code": "Code", "message": "Message", "type": "Type"}], "submission_time": "2023-11-01T11:18:49.758Z", "template_id": "spark-3.3-jaas-v2-cp4d-template"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetSparkEngineApplicationStatus successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetSparkEngineApplicationStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				getSparkEngineApplicationStatusOptionsModel := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				getSparkEngineApplicationStatusOptionsModel.EngineID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.ApplicationID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSparkEngineApplicationStatus with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				getSparkEngineApplicationStatusOptionsModel := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				getSparkEngineApplicationStatusOptionsModel.EngineID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.ApplicationID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSparkEngineApplicationStatusOptions model with no property values
				getSparkEngineApplicationStatusOptionsModelNew := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModelNew)
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
			It(`Invoke GetSparkEngineApplicationStatus successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				getSparkEngineApplicationStatusOptionsModel := new(watsonxdatav2.GetSparkEngineApplicationStatusOptions)
				getSparkEngineApplicationStatusOptionsModel.EngineID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.ApplicationID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getSparkEngineApplicationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptionsModel)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav2.TestLHConsoleOptions)
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
					fmt.Fprintf(res, "%s", `{"message": "Successful message", "message_code": "successfulCode"}`)
				}))
			})
			It(`Invoke TestLHConsole successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav2.TestLHConsoleOptions)
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
					fmt.Fprintf(res, "%s", `{"message": "Successful message", "message_code": "successfulCode"}`)
				}))
			})
			It(`Invoke TestLHConsole successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				testLhConsoleOptionsModel := new(watsonxdatav2.TestLHConsoleOptions)
				testLhConsoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.TestLHConsole(testLhConsoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke TestLHConsole with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav2.TestLHConsoleOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := new(watsonxdatav2.TestLHConsoleOptions)
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
	Describe(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions) - Operation response error`, func() {
		listCatalogsPath := "/catalogs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalogs with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(watsonxdatav2.ListCatalogsOptions)
				listCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
		listCatalogsPath := "/catalogs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListCatalogs successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(watsonxdatav2.ListCatalogsOptions)
				listCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListCatalogsWithContext(ctx, listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListCatalogsWithContext(ctx, listCatalogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalogs": [{"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ListCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListCatalogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(watsonxdatav2.ListCatalogsOptions)
				listCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCatalogs with error: Operation request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(watsonxdatav2.ListCatalogsOptions)
				listCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
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
			It(`Invoke ListCatalogs successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(watsonxdatav2.ListCatalogsOptions)
				listCatalogsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListCatalogs(listCatalogsOptionsModel)
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
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions) - Operation response error`, func() {
		getCatalogPath := "/catalogs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalog with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(watsonxdatav2.GetCatalogOptions)
				getCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
		getCatalogPath := "/catalogs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog": {"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetCatalog successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(watsonxdatav2.GetCatalogOptions)
				getCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"catalog": {"actions": ["Actions"], "associated_buckets": ["AssociatedBuckets"], "associated_databases": ["AssociatedDatabases"], "associated_engines": ["AssociatedEngines"], "catalog_name": "sampleCatalog", "catalog_type": "iceberg", "created_by": "<username>@<domain>.com", "created_on": "1602839833", "description": "Iceberg catalog description", "hostname": "s3a://samplehost.com", "last_sync_at": "1602839833", "managed_by": "ibm", "metastore": "glue", "port": "3232", "status": "running", "sync_description": "Table registration was successful", "sync_exception": ["SyncException"], "sync_status": "SUCCESS", "tags": ["Tags"], "thrift_uri": "thrift://samplehost-catalog:4354"}, "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(watsonxdatav2.GetCatalogOptions)
				getCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(watsonxdatav2.GetCatalogOptions)
				getCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogOptions model with no property values
				getCatalogOptionsModelNew := new(watsonxdatav2.GetCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetCatalog(getCatalogOptionsModelNew)
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
			It(`Invoke GetCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(watsonxdatav2.GetCatalogOptions)
				getCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				getCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetCatalog(getCatalogOptionsModel)
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
	Describe(`ListSchemas(listSchemasOptions *ListSchemasOptions) - Operation response error`, func() {
		listSchemasPath := "/catalogs/testString/schemas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSchemasPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSchemas with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(watsonxdatav2.ListSchemasOptions)
				listSchemasOptionsModel.EngineID = core.StringPtr("testString")
				listSchemasOptionsModel.CatalogID = core.StringPtr("testString")
				listSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSchemas(listSchemasOptions *ListSchemasOptions)`, func() {
		listSchemasPath := "/catalogs/testString/schemas"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "schemas": ["Schemas"]}`)
				}))
			})
			It(`Invoke ListSchemas successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(watsonxdatav2.ListSchemasOptions)
				listSchemasOptionsModel.EngineID = core.StringPtr("testString")
				listSchemasOptionsModel.CatalogID = core.StringPtr("testString")
				listSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListSchemasWithContext(ctx, listSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListSchemasWithContext(ctx, listSchemasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "schemas": ["Schemas"]}`)
				}))
			})
			It(`Invoke ListSchemas successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListSchemas(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(watsonxdatav2.ListSchemasOptions)
				listSchemasOptionsModel.EngineID = core.StringPtr("testString")
				listSchemasOptionsModel.CatalogID = core.StringPtr("testString")
				listSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSchemas with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(watsonxdatav2.ListSchemasOptions)
				listSchemasOptionsModel.EngineID = core.StringPtr("testString")
				listSchemasOptionsModel.CatalogID = core.StringPtr("testString")
				listSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSchemasOptions model with no property values
				listSchemasOptionsModelNew := new(watsonxdatav2.ListSchemasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListSchemas(listSchemasOptionsModelNew)
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
			It(`Invoke ListSchemas successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(watsonxdatav2.ListSchemasOptions)
				listSchemasOptionsModel.EngineID = core.StringPtr("testString")
				listSchemasOptionsModel.CatalogID = core.StringPtr("testString")
				listSchemasOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListSchemas(listSchemasOptionsModel)
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
		createSchemaPath := "/catalogs/testString/schemas"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSchema with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav2.CreateSchemaOptions)
				createSchemaOptionsModel.EngineID = core.StringPtr("testString")
				createSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				createSchemaOptionsModel.CustomPath = core.StringPtr("sample-path")
				createSchemaOptionsModel.SchemaName = core.StringPtr("SampleSchema1")
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
		createSchemaPath := "/catalogs/testString/schemas"
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
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateSchema successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav2.CreateSchemaOptions)
				createSchemaOptionsModel.EngineID = core.StringPtr("testString")
				createSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				createSchemaOptionsModel.CustomPath = core.StringPtr("sample-path")
				createSchemaOptionsModel.SchemaName = core.StringPtr("SampleSchema1")
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
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke CreateSchema successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				createSchemaOptionsModel := new(watsonxdatav2.CreateSchemaOptions)
				createSchemaOptionsModel.EngineID = core.StringPtr("testString")
				createSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				createSchemaOptionsModel.CustomPath = core.StringPtr("sample-path")
				createSchemaOptionsModel.SchemaName = core.StringPtr("SampleSchema1")
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav2.CreateSchemaOptions)
				createSchemaOptionsModel.EngineID = core.StringPtr("testString")
				createSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				createSchemaOptionsModel.CustomPath = core.StringPtr("sample-path")
				createSchemaOptionsModel.SchemaName = core.StringPtr("SampleSchema1")
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
				createSchemaOptionsModelNew := new(watsonxdatav2.CreateSchemaOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(watsonxdatav2.CreateSchemaOptions)
				createSchemaOptionsModel.EngineID = core.StringPtr("testString")
				createSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				createSchemaOptionsModel.CustomPath = core.StringPtr("sample-path")
				createSchemaOptionsModel.SchemaName = core.StringPtr("SampleSchema1")
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
		deleteSchemaPath := "/catalogs/testString/schemas/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSchemaPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSchema successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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
				deleteSchemaOptionsModel := new(watsonxdatav2.DeleteSchemaOptions)
				deleteSchemaOptionsModel.EngineID = core.StringPtr("testString")
				deleteSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				deleteSchemaOptionsModel.SchemaID = core.StringPtr("testString")
				deleteSchemaOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteSchema(deleteSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSchema with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsModel := new(watsonxdatav2.DeleteSchemaOptions)
				deleteSchemaOptionsModel.EngineID = core.StringPtr("testString")
				deleteSchemaOptionsModel.CatalogID = core.StringPtr("testString")
				deleteSchemaOptionsModel.SchemaID = core.StringPtr("testString")
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
				deleteSchemaOptionsModelNew := new(watsonxdatav2.DeleteSchemaOptions)
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
	Describe(`ListTables(listTablesOptions *ListTablesOptions) - Operation response error`, func() {
		listTablesPath := "/catalogs/testString/schemas/testString/tables"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTables with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(watsonxdatav2.ListTablesOptions)
				listTablesOptionsModel.CatalogID = core.StringPtr("testString")
				listTablesOptionsModel.SchemaID = core.StringPtr("testString")
				listTablesOptionsModel.EngineID = core.StringPtr("testString")
				listTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTables(listTablesOptions *ListTablesOptions)`, func() {
		listTablesPath := "/catalogs/testString/schemas/testString/tables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "tables": ["Tables"]}`)
				}))
			})
			It(`Invoke ListTables successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(watsonxdatav2.ListTablesOptions)
				listTablesOptionsModel.CatalogID = core.StringPtr("testString")
				listTablesOptionsModel.SchemaID = core.StringPtr("testString")
				listTablesOptionsModel.EngineID = core.StringPtr("testString")
				listTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListTablesWithContext(ctx, listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListTables(listTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListTablesWithContext(ctx, listTablesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTablesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "tables": ["Tables"]}`)
				}))
			})
			It(`Invoke ListTables successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListTables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(watsonxdatav2.ListTablesOptions)
				listTablesOptionsModel.CatalogID = core.StringPtr("testString")
				listTablesOptionsModel.SchemaID = core.StringPtr("testString")
				listTablesOptionsModel.EngineID = core.StringPtr("testString")
				listTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListTables(listTablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTables with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(watsonxdatav2.ListTablesOptions)
				listTablesOptionsModel.CatalogID = core.StringPtr("testString")
				listTablesOptionsModel.SchemaID = core.StringPtr("testString")
				listTablesOptionsModel.EngineID = core.StringPtr("testString")
				listTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListTables(listTablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTablesOptions model with no property values
				listTablesOptionsModelNew := new(watsonxdatav2.ListTablesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListTables(listTablesOptionsModelNew)
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
			It(`Invoke ListTables successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTablesOptions model
				listTablesOptionsModel := new(watsonxdatav2.ListTablesOptions)
				listTablesOptionsModel.CatalogID = core.StringPtr("testString")
				listTablesOptionsModel.SchemaID = core.StringPtr("testString")
				listTablesOptionsModel.EngineID = core.StringPtr("testString")
				listTablesOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListTables(listTablesOptionsModel)
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
	Describe(`GetTable(getTableOptions *GetTableOptions) - Operation response error`, func() {
		getTablePath := "/catalogs/testString/schemas/testString/tables/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTable with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(watsonxdatav2.GetTableOptions)
				getTableOptionsModel.CatalogID = core.StringPtr("testString")
				getTableOptionsModel.SchemaID = core.StringPtr("testString")
				getTableOptionsModel.TableID = core.StringPtr("testString")
				getTableOptionsModel.EngineID = core.StringPtr("testString")
				getTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTable(getTableOptions *GetTableOptions)`, func() {
		getTablePath := "/catalogs/testString/schemas/testString/tables/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"columns": [{"column_name": "expenses", "comment": "expenses column", "extra": "varchar", "type": "varchar"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetTable successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(watsonxdatav2.GetTableOptions)
				getTableOptionsModel.CatalogID = core.StringPtr("testString")
				getTableOptionsModel.SchemaID = core.StringPtr("testString")
				getTableOptionsModel.TableID = core.StringPtr("testString")
				getTableOptionsModel.EngineID = core.StringPtr("testString")
				getTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.GetTableWithContext(ctx, getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.GetTable(getTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.GetTableWithContext(ctx, getTableOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTablePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"columns": [{"column_name": "expenses", "comment": "expenses column", "extra": "varchar", "type": "varchar"}], "response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke GetTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.GetTable(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(watsonxdatav2.GetTableOptions)
				getTableOptionsModel.CatalogID = core.StringPtr("testString")
				getTableOptionsModel.SchemaID = core.StringPtr("testString")
				getTableOptionsModel.TableID = core.StringPtr("testString")
				getTableOptionsModel.EngineID = core.StringPtr("testString")
				getTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.GetTable(getTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTable with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(watsonxdatav2.GetTableOptions)
				getTableOptionsModel.CatalogID = core.StringPtr("testString")
				getTableOptionsModel.SchemaID = core.StringPtr("testString")
				getTableOptionsModel.TableID = core.StringPtr("testString")
				getTableOptionsModel.EngineID = core.StringPtr("testString")
				getTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.GetTable(getTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTableOptions model with no property values
				getTableOptionsModelNew := new(watsonxdatav2.GetTableOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.GetTable(getTableOptionsModelNew)
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
			It(`Invoke GetTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the GetTableOptions model
				getTableOptionsModel := new(watsonxdatav2.GetTableOptions)
				getTableOptionsModel.CatalogID = core.StringPtr("testString")
				getTableOptionsModel.SchemaID = core.StringPtr("testString")
				getTableOptionsModel.TableID = core.StringPtr("testString")
				getTableOptionsModel.EngineID = core.StringPtr("testString")
				getTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				getTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.GetTable(getTableOptionsModel)
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
	Describe(`DeleteTable(deleteTableOptions *DeleteTableOptions)`, func() {
		deleteTablePath := "/catalogs/testString/schemas/testString/tables/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTablePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := watsonxDataService.DeleteTable(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTableOptions model
				deleteTableOptionsModel := new(watsonxdatav2.DeleteTableOptions)
				deleteTableOptionsModel.CatalogID = core.StringPtr("testString")
				deleteTableOptionsModel.SchemaID = core.StringPtr("testString")
				deleteTableOptionsModel.TableID = core.StringPtr("testString")
				deleteTableOptionsModel.EngineID = core.StringPtr("testString")
				deleteTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				deleteTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = watsonxDataService.DeleteTable(deleteTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTable with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the DeleteTableOptions model
				deleteTableOptionsModel := new(watsonxdatav2.DeleteTableOptions)
				deleteTableOptionsModel.CatalogID = core.StringPtr("testString")
				deleteTableOptionsModel.SchemaID = core.StringPtr("testString")
				deleteTableOptionsModel.TableID = core.StringPtr("testString")
				deleteTableOptionsModel.EngineID = core.StringPtr("testString")
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
				deleteTableOptionsModelNew := new(watsonxdatav2.DeleteTableOptions)
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
	Describe(`UpdateTable(updateTableOptions *UpdateTableOptions) - Operation response error`, func() {
		updateTablePath := "/catalogs/testString/schemas/testString/tables/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTablePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTable with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav2.UpdateTableOptions)
				updateTableOptionsModel.CatalogID = core.StringPtr("testString")
				updateTableOptionsModel.SchemaID = core.StringPtr("testString")
				updateTableOptionsModel.TableID = core.StringPtr("testString")
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTable(updateTableOptions *UpdateTableOptions)`, func() {
		updateTablePath := "/catalogs/testString/schemas/testString/tables/testString"
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateTable successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav2.UpdateTableOptions)
				updateTableOptionsModel.CatalogID = core.StringPtr("testString")
				updateTableOptionsModel.SchemaID = core.StringPtr("testString")
				updateTableOptionsModel.TableID = core.StringPtr("testString")
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateTable successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
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

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav2.UpdateTableOptions)
				updateTableOptionsModel.CatalogID = core.StringPtr("testString")
				updateTableOptionsModel.SchemaID = core.StringPtr("testString")
				updateTableOptionsModel.TableID = core.StringPtr("testString")
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateTable(updateTableOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTable with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav2.UpdateTableOptions)
				updateTableOptionsModel.CatalogID = core.StringPtr("testString")
				updateTableOptionsModel.SchemaID = core.StringPtr("testString")
				updateTableOptionsModel.TableID = core.StringPtr("testString")
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
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
				updateTableOptionsModelNew := new(watsonxdatav2.UpdateTableOptions)
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
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateTableOptions model
				updateTableOptionsModel := new(watsonxdatav2.UpdateTableOptions)
				updateTableOptionsModel.CatalogID = core.StringPtr("testString")
				updateTableOptionsModel.SchemaID = core.StringPtr("testString")
				updateTableOptionsModel.TableID = core.StringPtr("testString")
				updateTableOptionsModel.EngineID = core.StringPtr("testString")
				updateTableOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateTableOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateTableOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateTable(updateTableOptionsModel)
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
	Describe(`ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions) - Operation response error`, func() {
		listTableSnapshotsPath := "/catalogs/testString/schemas/testString/tables/testString/snapshots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTableSnapshots with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTableSnapshotsOptions model
				listTableSnapshotsOptionsModel := new(watsonxdatav2.ListTableSnapshotsOptions)
				listTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.CatalogID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.SchemaID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.TableID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions)`, func() {
		listTableSnapshotsPath := "/catalogs/testString/schemas/testString/tables/testString/snapshots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "snapshots": [{"committed_at": "1609379392", "operation": "alter", "snapshot_id": "2332342122211222", "summary": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListTableSnapshots successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ListTableSnapshotsOptions model
				listTableSnapshotsOptionsModel := new(watsonxdatav2.ListTableSnapshotsOptions)
				listTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.CatalogID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.SchemaID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.TableID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ListTableSnapshotsWithContext(ctx, listTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ListTableSnapshotsWithContext(ctx, listTableSnapshotsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTableSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}, "snapshots": [{"committed_at": "1609379392", "operation": "alter", "snapshot_id": "2332342122211222", "summary": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListTableSnapshots successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ListTableSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTableSnapshotsOptions model
				listTableSnapshotsOptionsModel := new(watsonxdatav2.ListTableSnapshotsOptions)
				listTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.CatalogID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.SchemaID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.TableID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTableSnapshots with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTableSnapshotsOptions model
				listTableSnapshotsOptionsModel := new(watsonxdatav2.ListTableSnapshotsOptions)
				listTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.CatalogID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.SchemaID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.TableID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTableSnapshotsOptions model with no property values
				listTableSnapshotsOptionsModelNew := new(watsonxdatav2.ListTableSnapshotsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModelNew)
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
			It(`Invoke ListTableSnapshots successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ListTableSnapshotsOptions model
				listTableSnapshotsOptionsModel := new(watsonxdatav2.ListTableSnapshotsOptions)
				listTableSnapshotsOptionsModel.EngineID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.CatalogID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.SchemaID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.TableID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.AuthInstanceID = core.StringPtr("testString")
				listTableSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptionsModel)
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
	Describe(`ReplaceSnapshot(replaceSnapshotOptions *ReplaceSnapshotOptions) - Operation response error`, func() {
		replaceSnapshotPath := "/catalogs/testString/schemas/testString/tables/testString/snapshots/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSnapshotPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceSnapshot with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplaceSnapshotOptions model
				replaceSnapshotOptionsModel := new(watsonxdatav2.ReplaceSnapshotOptions)
				replaceSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.CatalogID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SchemaID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.TableID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SnapshotID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSnapshot(replaceSnapshotOptions *ReplaceSnapshotOptions)`, func() {
		replaceSnapshotPath := "/catalogs/testString/schemas/testString/tables/testString/snapshots/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSnapshotPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ReplaceSnapshot successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceSnapshotOptions model
				replaceSnapshotOptionsModel := new(watsonxdatav2.ReplaceSnapshotOptions)
				replaceSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.CatalogID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SchemaID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.TableID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SnapshotID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.ReplaceSnapshotWithContext(ctx, replaceSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.ReplaceSnapshotWithContext(ctx, replaceSnapshotOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceSnapshotPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["engine_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke ReplaceSnapshot successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.ReplaceSnapshot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceSnapshotOptions model
				replaceSnapshotOptionsModel := new(watsonxdatav2.ReplaceSnapshotOptions)
				replaceSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.CatalogID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SchemaID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.TableID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SnapshotID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceSnapshot with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplaceSnapshotOptions model
				replaceSnapshotOptionsModel := new(watsonxdatav2.ReplaceSnapshotOptions)
				replaceSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.CatalogID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SchemaID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.TableID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SnapshotID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceSnapshotOptions model with no property values
				replaceSnapshotOptionsModelNew := new(watsonxdatav2.ReplaceSnapshotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModelNew)
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
			It(`Invoke ReplaceSnapshot successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the ReplaceSnapshotOptions model
				replaceSnapshotOptionsModel := new(watsonxdatav2.ReplaceSnapshotOptions)
				replaceSnapshotOptionsModel.EngineID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.CatalogID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SchemaID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.TableID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.SnapshotID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.AuthInstanceID = core.StringPtr("testString")
				replaceSnapshotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.ReplaceSnapshot(replaceSnapshotOptionsModel)
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
	Describe(`UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions) - Operation response error`, func() {
		updateSyncCatalogPath := "/catalogs/testString/sync"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSyncCatalogPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Authinstanceid"]).ToNot(BeNil())
					Expect(req.Header["Authinstanceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSyncCatalog with error: Operation response processing error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSyncCatalogOptions model
				updateSyncCatalogOptionsModel := new(watsonxdatav2.UpdateSyncCatalogOptions)
				updateSyncCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSyncCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				watsonxDataService.EnableRetries(0, 0)
				result, response, operationErr = watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions)`, func() {
		updateSyncCatalogPath := "/catalogs/testString/sync"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSyncCatalogPath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateSyncCatalog successfully with retries`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())
				watsonxDataService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSyncCatalogOptions model
				updateSyncCatalogOptionsModel := new(watsonxdatav2.UpdateSyncCatalogOptions)
				updateSyncCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSyncCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := watsonxDataService.UpdateSyncCatalogWithContext(ctx, updateSyncCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				watsonxDataService.DisableRetries()
				result, response, operationErr := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = watsonxDataService.UpdateSyncCatalogWithContext(ctx, updateSyncCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSyncCatalogPath))
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
					fmt.Fprintf(res, "%s", `{"response": {"message": "Successful message", "message_code": "successfulCode"}}`)
				}))
			})
			It(`Invoke UpdateSyncCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := watsonxDataService.UpdateSyncCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSyncCatalogOptions model
				updateSyncCatalogOptionsModel := new(watsonxdatav2.UpdateSyncCatalogOptions)
				updateSyncCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSyncCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSyncCatalog with error: Operation validation and request error`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSyncCatalogOptions model
				updateSyncCatalogOptionsModel := new(watsonxdatav2.UpdateSyncCatalogOptions)
				updateSyncCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSyncCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := watsonxDataService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSyncCatalogOptions model with no property values
				updateSyncCatalogOptionsModelNew := new(watsonxdatav2.UpdateSyncCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModelNew)
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
			It(`Invoke UpdateSyncCatalog successfully`, func() {
				watsonxDataService, serviceErr := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(watsonxDataService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateSyncCatalogOptions model
				updateSyncCatalogOptionsModel := new(watsonxdatav2.UpdateSyncCatalogOptions)
				updateSyncCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Body = []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}
				updateSyncCatalogOptionsModel.AuthInstanceID = core.StringPtr("testString")
				updateSyncCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			watsonxDataService, _ := watsonxdatav2.NewWatsonxDataV2(&watsonxdatav2.WatsonxDataV2Options{
				URL:           "http://watsonxdatav2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewBucketDetails successfully`, func() {
				bucketName := "sample-bucket"
				_model, err := watsonxDataService.NewBucketDetails(bucketName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateActivateBucketOptions successfully`, func() {
				// Construct an instance of the CreateActivateBucketOptions model
				bucketID := "testString"
				createActivateBucketOptionsModel := watsonxDataService.NewCreateActivateBucketOptions(bucketID)
				createActivateBucketOptionsModel.SetBucketID("testString")
				createActivateBucketOptionsModel.SetAuthInstanceID("testString")
				createActivateBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createActivateBucketOptionsModel).ToNot(BeNil())
				Expect(createActivateBucketOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(createActivateBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createActivateBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBucketRegistrationOptions successfully`, func() {
				// Construct an instance of the BucketDetails model
				bucketDetailsModel := new(watsonxdatav2.BucketDetails)
				Expect(bucketDetailsModel).ToNot(BeNil())
				bucketDetailsModel.AccessKey = core.StringPtr("<access_key>")
				bucketDetailsModel.BucketName = core.StringPtr("sample-bucket")
				bucketDetailsModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketDetailsModel.SecretKey = core.StringPtr("secret_key")
				Expect(bucketDetailsModel.AccessKey).To(Equal(core.StringPtr("<access_key>")))
				Expect(bucketDetailsModel.BucketName).To(Equal(core.StringPtr("sample-bucket")))
				Expect(bucketDetailsModel.Endpoint).To(Equal(core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")))
				Expect(bucketDetailsModel.SecretKey).To(Equal(core.StringPtr("secret_key")))

				// Construct an instance of the CreateBucketRegistrationOptions model
				var createBucketRegistrationOptionsBucketDetails *watsonxdatav2.BucketDetails = nil
				createBucketRegistrationOptionsBucketType := "ibm_cos"
				createBucketRegistrationOptionsCatalogName := "sampleCatalog"
				createBucketRegistrationOptionsDescription := "COS bucket for customer data"
				createBucketRegistrationOptionsManagedBy := "ibm"
				createBucketRegistrationOptionsTableType := "iceberg"
				createBucketRegistrationOptionsModel := watsonxDataService.NewCreateBucketRegistrationOptions(createBucketRegistrationOptionsBucketDetails, createBucketRegistrationOptionsBucketType, createBucketRegistrationOptionsCatalogName, createBucketRegistrationOptionsDescription, createBucketRegistrationOptionsManagedBy, createBucketRegistrationOptionsTableType)
				createBucketRegistrationOptionsModel.SetBucketDetails(bucketDetailsModel)
				createBucketRegistrationOptionsModel.SetBucketType("ibm_cos")
				createBucketRegistrationOptionsModel.SetCatalogName("sampleCatalog")
				createBucketRegistrationOptionsModel.SetDescription("COS bucket for customer data")
				createBucketRegistrationOptionsModel.SetManagedBy("ibm")
				createBucketRegistrationOptionsModel.SetTableType("iceberg")
				createBucketRegistrationOptionsModel.SetBucketDisplayName("sample-bucket-displayname")
				createBucketRegistrationOptionsModel.SetBucketTags([]string{"read customer data", "write customer data'"})
				createBucketRegistrationOptionsModel.SetCatalogTags([]string{"catalog_tag_1", "catalog_tag_2"})
				createBucketRegistrationOptionsModel.SetRegion("us-south")
				createBucketRegistrationOptionsModel.SetState("active")
				createBucketRegistrationOptionsModel.SetAuthInstanceID("testString")
				createBucketRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBucketRegistrationOptionsModel).ToNot(BeNil())
				Expect(createBucketRegistrationOptionsModel.BucketDetails).To(Equal(bucketDetailsModel))
				Expect(createBucketRegistrationOptionsModel.BucketType).To(Equal(core.StringPtr("ibm_cos")))
				Expect(createBucketRegistrationOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(createBucketRegistrationOptionsModel.Description).To(Equal(core.StringPtr("COS bucket for customer data")))
				Expect(createBucketRegistrationOptionsModel.ManagedBy).To(Equal(core.StringPtr("ibm")))
				Expect(createBucketRegistrationOptionsModel.TableType).To(Equal(core.StringPtr("iceberg")))
				Expect(createBucketRegistrationOptionsModel.BucketDisplayName).To(Equal(core.StringPtr("sample-bucket-displayname")))
				Expect(createBucketRegistrationOptionsModel.BucketTags).To(Equal([]string{"read customer data", "write customer data'"}))
				Expect(createBucketRegistrationOptionsModel.CatalogTags).To(Equal([]string{"catalog_tag_1", "catalog_tag_2"}))
				Expect(createBucketRegistrationOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createBucketRegistrationOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(createBucketRegistrationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createBucketRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDatabaseRegistrationOptions successfully`, func() {
				// Construct an instance of the RegisterDatabaseCatalogBodyDatabaseDetails model
				registerDatabaseCatalogBodyDatabaseDetailsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabaseDetails)
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel).ToNot(BeNil())
				registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate = core.StringPtr("contents of a pem/crt file")
				registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension = core.StringPtr("pem/crt")
				registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@<hostname>.com")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts = core.StringPtr("abc.com:1234,xyz.com:4321")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				registerDatabaseCatalogBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Certificate).To(Equal(core.StringPtr("contents of a pem/crt file")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.CertificateExtension).To(Equal(core.StringPtr("pem/crt")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.DatabaseName).To(Equal(core.StringPtr("new_database")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Hostname).To(Equal(core.StringPtr("db2@<hostname>.com")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Hosts).To(Equal(core.StringPtr("abc.com:1234,xyz.com:4321")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Password).To(Equal(core.StringPtr("samplepassword")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Port).To(Equal(core.Int64Ptr(int64(4553))))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Sasl).To(Equal(core.BoolPtr(true)))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Ssl).To(Equal(core.BoolPtr(true)))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Tables).To(Equal(core.StringPtr("kafka_table_name")))
				Expect(registerDatabaseCatalogBodyDatabaseDetailsModel.Username).To(Equal(core.StringPtr("sampleuser")))

				// Construct an instance of the RegisterDatabaseCatalogBodyDatabasePropertiesItems model
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel := new(watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems)
				Expect(registerDatabaseCatalogBodyDatabasePropertiesItemsModel).ToNot(BeNil())
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt = core.BoolPtr(true)
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key = core.StringPtr("abc")
				registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value = core.StringPtr("xyz")
				Expect(registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Encrypt).To(Equal(core.BoolPtr(true)))
				Expect(registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Key).To(Equal(core.StringPtr("abc")))
				Expect(registerDatabaseCatalogBodyDatabasePropertiesItemsModel.Value).To(Equal(core.StringPtr("xyz")))

				// Construct an instance of the CreateDatabaseRegistrationOptions model
				createDatabaseRegistrationOptionsCatalogName := "sampleCatalog"
				createDatabaseRegistrationOptionsDatabaseDisplayName := "new_database"
				createDatabaseRegistrationOptionsDatabaseType := "db2"
				createDatabaseRegistrationOptionsModel := watsonxDataService.NewCreateDatabaseRegistrationOptions(createDatabaseRegistrationOptionsCatalogName, createDatabaseRegistrationOptionsDatabaseDisplayName, createDatabaseRegistrationOptionsDatabaseType)
				createDatabaseRegistrationOptionsModel.SetCatalogName("sampleCatalog")
				createDatabaseRegistrationOptionsModel.SetDatabaseDisplayName("new_database")
				createDatabaseRegistrationOptionsModel.SetDatabaseType("db2")
				createDatabaseRegistrationOptionsModel.SetCreatedOn(int64(38))
				createDatabaseRegistrationOptionsModel.SetDatabaseDetails(registerDatabaseCatalogBodyDatabaseDetailsModel)
				createDatabaseRegistrationOptionsModel.SetDatabaseProperties([]watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel})
				createDatabaseRegistrationOptionsModel.SetDescription("db2 extenal database description")
				createDatabaseRegistrationOptionsModel.SetTags([]string{"tag_1", "tag_2"})
				createDatabaseRegistrationOptionsModel.SetAuthInstanceID("testString")
				createDatabaseRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDatabaseRegistrationOptionsModel).ToNot(BeNil())
				Expect(createDatabaseRegistrationOptionsModel.CatalogName).To(Equal(core.StringPtr("sampleCatalog")))
				Expect(createDatabaseRegistrationOptionsModel.DatabaseDisplayName).To(Equal(core.StringPtr("new_database")))
				Expect(createDatabaseRegistrationOptionsModel.DatabaseType).To(Equal(core.StringPtr("db2")))
				Expect(createDatabaseRegistrationOptionsModel.CreatedOn).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createDatabaseRegistrationOptionsModel.DatabaseDetails).To(Equal(registerDatabaseCatalogBodyDatabaseDetailsModel))
				Expect(createDatabaseRegistrationOptionsModel.DatabaseProperties).To(Equal([]watsonxdatav2.RegisterDatabaseCatalogBodyDatabasePropertiesItems{*registerDatabaseCatalogBodyDatabasePropertiesItemsModel}))
				Expect(createDatabaseRegistrationOptionsModel.Description).To(Equal(core.StringPtr("db2 extenal database description")))
				Expect(createDatabaseRegistrationOptionsModel.Tags).To(Equal([]string{"tag_1", "tag_2"}))
				Expect(createDatabaseRegistrationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDb2EngineOptions successfully`, func() {
				// Construct an instance of the CreateDb2EngineDetails model
				createDb2EngineDetailsModel := new(watsonxdatav2.CreateDb2EngineDetails)
				Expect(createDb2EngineDetailsModel).ToNot(BeNil())
				createDb2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				Expect(createDb2EngineDetailsModel.ConnectionString).To(Equal(core.StringPtr("1.2.3.4")))

				// Construct an instance of the CreateDb2EngineOptions model
				createDb2EngineOptionsOrigin := "external"
				createDb2EngineOptionsType := "db2"
				createDb2EngineOptionsModel := watsonxDataService.NewCreateDb2EngineOptions(createDb2EngineOptionsOrigin, createDb2EngineOptionsType)
				createDb2EngineOptionsModel.SetOrigin("external")
				createDb2EngineOptionsModel.SetType("db2")
				createDb2EngineOptionsModel.SetDescription("db2 engine description")
				createDb2EngineOptionsModel.SetEngineDetails(createDb2EngineDetailsModel)
				createDb2EngineOptionsModel.SetEngineDisplayName("sampleEngine")
				createDb2EngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				createDb2EngineOptionsModel.SetAuthInstanceID("testString")
				createDb2EngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDb2EngineOptionsModel).ToNot(BeNil())
				Expect(createDb2EngineOptionsModel.Origin).To(Equal(core.StringPtr("external")))
				Expect(createDb2EngineOptionsModel.Type).To(Equal(core.StringPtr("db2")))
				Expect(createDb2EngineOptionsModel.Description).To(Equal(core.StringPtr("db2 engine description")))
				Expect(createDb2EngineOptionsModel.EngineDetails).To(Equal(createDb2EngineDetailsModel))
				Expect(createDb2EngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(createDb2EngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(createDb2EngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDb2EngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDriverDatabaseCatalogOptions successfully`, func() {
				// Construct an instance of the CreateDriverDatabaseCatalogOptions model
				databaseDisplayName := "testString"
				databaseType := "testString"
				catalogName := "testString"
				hostname := "testString"
				port := "testString"
				createDriverDatabaseCatalogOptionsModel := watsonxDataService.NewCreateDriverDatabaseCatalogOptions(databaseDisplayName, databaseType, catalogName, hostname, port)
				createDriverDatabaseCatalogOptionsModel.SetDatabaseDisplayName("testString")
				createDriverDatabaseCatalogOptionsModel.SetDatabaseType("testString")
				createDriverDatabaseCatalogOptionsModel.SetCatalogName("testString")
				createDriverDatabaseCatalogOptionsModel.SetHostname("testString")
				createDriverDatabaseCatalogOptionsModel.SetPort("testString")
				createDriverDatabaseCatalogOptionsModel.SetDriver(CreateMockReader("This is a mock file."))
				createDriverDatabaseCatalogOptionsModel.SetDriverContentType("testString")
				createDriverDatabaseCatalogOptionsModel.SetDriverFileName("testString")
				createDriverDatabaseCatalogOptionsModel.SetCertificate("testString")
				createDriverDatabaseCatalogOptionsModel.SetCertificateExtension("testString")
				createDriverDatabaseCatalogOptionsModel.SetSsl("testString")
				createDriverDatabaseCatalogOptionsModel.SetUsername("testString")
				createDriverDatabaseCatalogOptionsModel.SetPassword("testString")
				createDriverDatabaseCatalogOptionsModel.SetDatabaseName("testString")
				createDriverDatabaseCatalogOptionsModel.SetDescription("testString")
				createDriverDatabaseCatalogOptionsModel.SetCreatedOn("testString")
				createDriverDatabaseCatalogOptionsModel.SetAuthInstanceID("testString")
				createDriverDatabaseCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDriverDatabaseCatalogOptionsModel).ToNot(BeNil())
				Expect(createDriverDatabaseCatalogOptionsModel.DatabaseDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.DatabaseType).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Hostname).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Port).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Driver).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDriverDatabaseCatalogOptionsModel.DriverContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.DriverFileName).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Certificate).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.CertificateExtension).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Ssl).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.DatabaseName).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.CreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDriverDatabaseCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineOptions successfully`, func() {
				// Construct an instance of the NodeDescriptionBody model
				nodeDescriptionBodyModel := new(watsonxdatav2.NodeDescriptionBody)
				Expect(nodeDescriptionBodyModel).ToNot(BeNil())
				nodeDescriptionBodyModel.NodeType = core.StringPtr("worker")
				nodeDescriptionBodyModel.Quantity = core.Int64Ptr(int64(38))
				Expect(nodeDescriptionBodyModel.NodeType).To(Equal(core.StringPtr("worker")))
				Expect(nodeDescriptionBodyModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the EngineDetailsBody model
				engineDetailsBodyModel := new(watsonxdatav2.EngineDetailsBody)
				Expect(engineDetailsBodyModel).ToNot(BeNil())
				engineDetailsBodyModel.ApiKey = core.StringPtr("<api_key>")
				engineDetailsBodyModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsBodyModel.Coordinator = nodeDescriptionBodyModel
				engineDetailsBodyModel.InstanceID = core.StringPtr("instance_id")
				engineDetailsBodyModel.ManagedBy = core.StringPtr("fully/self")
				engineDetailsBodyModel.SizeConfig = core.StringPtr("starter")
				engineDetailsBodyModel.Worker = nodeDescriptionBodyModel
				Expect(engineDetailsBodyModel.ApiKey).To(Equal(core.StringPtr("<api_key>")))
				Expect(engineDetailsBodyModel.ConnectionString).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(engineDetailsBodyModel.Coordinator).To(Equal(nodeDescriptionBodyModel))
				Expect(engineDetailsBodyModel.InstanceID).To(Equal(core.StringPtr("instance_id")))
				Expect(engineDetailsBodyModel.ManagedBy).To(Equal(core.StringPtr("fully/self")))
				Expect(engineDetailsBodyModel.SizeConfig).To(Equal(core.StringPtr("starter")))
				Expect(engineDetailsBodyModel.Worker).To(Equal(nodeDescriptionBodyModel))

				// Construct an instance of the CreateEngineOptions model
				createEngineOptionsOrigin := "native"
				createEngineOptionsType := "presto"
				createEngineOptionsModel := watsonxDataService.NewCreateEngineOptions(createEngineOptionsOrigin, createEngineOptionsType)
				createEngineOptionsModel.SetOrigin("native")
				createEngineOptionsModel.SetType("presto")
				createEngineOptionsModel.SetAssociatedCatalogs([]string{"iceberg_data", "hive_data"})
				createEngineOptionsModel.SetDescription("presto engine description")
				createEngineOptionsModel.SetEngineDetails(engineDetailsBodyModel)
				createEngineOptionsModel.SetEngineDisplayName("sampleEngine")
				createEngineOptionsModel.SetFirstTimeUse(true)
				createEngineOptionsModel.SetRegion("us-south")
				createEngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				createEngineOptionsModel.SetVersion("1.2.3")
				createEngineOptionsModel.SetAuthInstanceID("testString")
				createEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineOptionsModel).ToNot(BeNil())
				Expect(createEngineOptionsModel.Origin).To(Equal(core.StringPtr("native")))
				Expect(createEngineOptionsModel.Type).To(Equal(core.StringPtr("presto")))
				Expect(createEngineOptionsModel.AssociatedCatalogs).To(Equal([]string{"iceberg_data", "hive_data"}))
				Expect(createEngineOptionsModel.Description).To(Equal(core.StringPtr("presto engine description")))
				Expect(createEngineOptionsModel.EngineDetails).To(Equal(engineDetailsBodyModel))
				Expect(createEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(createEngineOptionsModel.FirstTimeUse).To(Equal(core.BoolPtr(true)))
				Expect(createEngineOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createEngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(createEngineOptionsModel.Version).To(Equal(core.StringPtr("1.2.3")))
				Expect(createEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnginePauseOptions successfully`, func() {
				// Construct an instance of the CreateEnginePauseOptions model
				engineID := "testString"
				createEnginePauseOptionsModel := watsonxDataService.NewCreateEnginePauseOptions(engineID)
				createEnginePauseOptionsModel.SetEngineID("testString")
				createEnginePauseOptionsModel.SetAuthInstanceID("testString")
				createEnginePauseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnginePauseOptionsModel).ToNot(BeNil())
				Expect(createEnginePauseOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createEnginePauseOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEnginePauseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineRestartOptions successfully`, func() {
				// Construct an instance of the CreateEngineRestartOptions model
				engineID := "testString"
				createEngineRestartOptionsModel := watsonxDataService.NewCreateEngineRestartOptions(engineID)
				createEngineRestartOptionsModel.SetEngineID("testString")
				createEngineRestartOptionsModel.SetAuthInstanceID("testString")
				createEngineRestartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineRestartOptionsModel).ToNot(BeNil())
				Expect(createEngineRestartOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineRestartOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineRestartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineResumeOptions successfully`, func() {
				// Construct an instance of the CreateEngineResumeOptions model
				engineID := "testString"
				createEngineResumeOptionsModel := watsonxDataService.NewCreateEngineResumeOptions(engineID)
				createEngineResumeOptionsModel.SetEngineID("testString")
				createEngineResumeOptionsModel.SetAuthInstanceID("testString")
				createEngineResumeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineResumeOptionsModel).ToNot(BeNil())
				Expect(createEngineResumeOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineResumeOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineResumeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEngineScaleOptions successfully`, func() {
				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				Expect(nodeDescriptionModel).ToNot(BeNil())
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(38))
				Expect(nodeDescriptionModel.NodeType).To(Equal(core.StringPtr("worker")))
				Expect(nodeDescriptionModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the CreateEngineScaleOptions model
				engineID := "testString"
				createEngineScaleOptionsModel := watsonxDataService.NewCreateEngineScaleOptions(engineID)
				createEngineScaleOptionsModel.SetEngineID("testString")
				createEngineScaleOptionsModel.SetCoordinator(nodeDescriptionModel)
				createEngineScaleOptionsModel.SetWorker(nodeDescriptionModel)
				createEngineScaleOptionsModel.SetAuthInstanceID("testString")
				createEngineScaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEngineScaleOptionsModel).ToNot(BeNil())
				Expect(createEngineScaleOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineScaleOptionsModel.Coordinator).To(Equal(nodeDescriptionModel))
				Expect(createEngineScaleOptionsModel.Worker).To(Equal(nodeDescriptionModel))
				Expect(createEngineScaleOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createEngineScaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateNetezzaEngineOptions successfully`, func() {
				// Construct an instance of the CreateNetezzaEngineDetails model
				createNetezzaEngineDetailsModel := new(watsonxdatav2.CreateNetezzaEngineDetails)
				Expect(createNetezzaEngineDetailsModel).ToNot(BeNil())
				createNetezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				Expect(createNetezzaEngineDetailsModel.ConnectionString).To(Equal(core.StringPtr("1.2.3.4")))

				// Construct an instance of the CreateNetezzaEngineOptions model
				createNetezzaEngineOptionsOrigin := "external"
				createNetezzaEngineOptionsType := "netezza"
				createNetezzaEngineOptionsModel := watsonxDataService.NewCreateNetezzaEngineOptions(createNetezzaEngineOptionsOrigin, createNetezzaEngineOptionsType)
				createNetezzaEngineOptionsModel.SetOrigin("external")
				createNetezzaEngineOptionsModel.SetType("netezza")
				createNetezzaEngineOptionsModel.SetDescription("netezza engine description")
				createNetezzaEngineOptionsModel.SetEngineDetails(createNetezzaEngineDetailsModel)
				createNetezzaEngineOptionsModel.SetEngineDisplayName("sampleEngine")
				createNetezzaEngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				createNetezzaEngineOptionsModel.SetAuthInstanceID("testString")
				createNetezzaEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createNetezzaEngineOptionsModel).ToNot(BeNil())
				Expect(createNetezzaEngineOptionsModel.Origin).To(Equal(core.StringPtr("external")))
				Expect(createNetezzaEngineOptionsModel.Type).To(Equal(core.StringPtr("netezza")))
				Expect(createNetezzaEngineOptionsModel.Description).To(Equal(core.StringPtr("netezza engine description")))
				Expect(createNetezzaEngineOptionsModel.EngineDetails).To(Equal(createNetezzaEngineDetailsModel))
				Expect(createNetezzaEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(createNetezzaEngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(createNetezzaEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createNetezzaEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOtherEngineOptions successfully`, func() {
				// Construct an instance of the OtherEngineDetails model
				otherEngineDetailsModel := new(watsonxdatav2.OtherEngineDetails)
				Expect(otherEngineDetailsModel).ToNot(BeNil())
				otherEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				otherEngineDetailsModel.EngineType = core.StringPtr("netezza")
				otherEngineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")
				Expect(otherEngineDetailsModel.ConnectionString).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(otherEngineDetailsModel.EngineType).To(Equal(core.StringPtr("netezza")))
				Expect(otherEngineDetailsModel.MetastoreHost).To(Equal(core.StringPtr("1.2.3.4")))

				// Construct an instance of the CreateOtherEngineOptions model
				createOtherEngineOptionsModel := watsonxDataService.NewCreateOtherEngineOptions()
				createOtherEngineOptionsModel.SetDescription("external engine description")
				createOtherEngineOptionsModel.SetEngineDetails(otherEngineDetailsModel)
				createOtherEngineOptionsModel.SetEngineDisplayName("sampleEngine01")
				createOtherEngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				createOtherEngineOptionsModel.SetAuthInstanceID("testString")
				createOtherEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOtherEngineOptionsModel).ToNot(BeNil())
				Expect(createOtherEngineOptionsModel.Description).To(Equal(core.StringPtr("external engine description")))
				Expect(createOtherEngineOptionsModel.EngineDetails).To(Equal(otherEngineDetailsModel))
				Expect(createOtherEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine01")))
				Expect(createOtherEngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(createOtherEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createOtherEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSchemaOptions successfully`, func() {
				// Construct an instance of the CreateSchemaOptions model
				engineID := "testString"
				catalogID := "testString"
				createSchemaOptionsCustomPath := "sample-path"
				createSchemaOptionsSchemaName := "SampleSchema1"
				createSchemaOptionsModel := watsonxDataService.NewCreateSchemaOptions(engineID, catalogID, createSchemaOptionsCustomPath, createSchemaOptionsSchemaName)
				createSchemaOptionsModel.SetEngineID("testString")
				createSchemaOptionsModel.SetCatalogID("testString")
				createSchemaOptionsModel.SetCustomPath("sample-path")
				createSchemaOptionsModel.SetSchemaName("SampleSchema1")
				createSchemaOptionsModel.SetBucketName("sample-bucket")
				createSchemaOptionsModel.SetAuthInstanceID("testString")
				createSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSchemaOptionsModel).ToNot(BeNil())
				Expect(createSchemaOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaOptionsModel.CustomPath).To(Equal(core.StringPtr("sample-path")))
				Expect(createSchemaOptionsModel.SchemaName).To(Equal(core.StringPtr("SampleSchema1")))
				Expect(createSchemaOptionsModel.BucketName).To(Equal(core.StringPtr("sample-bucket")))
				Expect(createSchemaOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSparkEngineApplicationOptions successfully`, func() {
				// Construct an instance of the SparkApplicationDetails model
				sparkApplicationDetailsModel := new(watsonxdatav2.SparkApplicationDetails)
				Expect(sparkApplicationDetailsModel).ToNot(BeNil())
				sparkApplicationDetailsModel.Application = core.StringPtr("s3://mybucket/wordcount.py")
				sparkApplicationDetailsModel.Arguments = []string{"people.txt"}
				sparkApplicationDetailsModel.Conf = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Env = map[string]string{"key1": "key:value"}
				sparkApplicationDetailsModel.Name = core.StringPtr("SparkApplicaton1")
				Expect(sparkApplicationDetailsModel.Application).To(Equal(core.StringPtr("s3://mybucket/wordcount.py")))
				Expect(sparkApplicationDetailsModel.Arguments).To(Equal([]string{"people.txt"}))
				Expect(sparkApplicationDetailsModel.Conf).To(Equal(map[string]string{"key1": "key:value"}))
				Expect(sparkApplicationDetailsModel.Env).To(Equal(map[string]string{"key1": "key:value"}))
				Expect(sparkApplicationDetailsModel.Name).To(Equal(core.StringPtr("SparkApplicaton1")))

				// Construct an instance of the CreateSparkEngineApplicationOptions model
				engineID := "testString"
				var createSparkEngineApplicationOptionsApplicationDetails *watsonxdatav2.SparkApplicationDetails = nil
				createSparkEngineApplicationOptionsModel := watsonxDataService.NewCreateSparkEngineApplicationOptions(engineID, createSparkEngineApplicationOptionsApplicationDetails)
				createSparkEngineApplicationOptionsModel.SetEngineID("testString")
				createSparkEngineApplicationOptionsModel.SetApplicationDetails(sparkApplicationDetailsModel)
				createSparkEngineApplicationOptionsModel.SetJobEndpoint("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")
				createSparkEngineApplicationOptionsModel.SetServiceInstanceID("testString")
				createSparkEngineApplicationOptionsModel.SetType("iae")
				createSparkEngineApplicationOptionsModel.SetAuthInstanceID("testString")
				createSparkEngineApplicationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSparkEngineApplicationOptionsModel).ToNot(BeNil())
				Expect(createSparkEngineApplicationOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(createSparkEngineApplicationOptionsModel.ApplicationDetails).To(Equal(sparkApplicationDetailsModel))
				Expect(createSparkEngineApplicationOptionsModel.JobEndpoint).To(Equal(core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications")))
				Expect(createSparkEngineApplicationOptionsModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSparkEngineApplicationOptionsModel.Type).To(Equal(core.StringPtr("iae")))
				Expect(createSparkEngineApplicationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSparkEngineApplicationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSparkEngineOptions successfully`, func() {
				// Construct an instance of the SparkEngineDetailsPrototype model
				sparkEngineDetailsPrototypeModel := new(watsonxdatav2.SparkEngineDetailsPrototype)
				Expect(sparkEngineDetailsPrototypeModel).ToNot(BeNil())
				sparkEngineDetailsPrototypeModel.ApiKey = core.StringPtr("apikey")
				sparkEngineDetailsPrototypeModel.ConnectionString = core.StringPtr("1.2.3.4")
				sparkEngineDetailsPrototypeModel.InstanceID = core.StringPtr("spark-id")
				sparkEngineDetailsPrototypeModel.ManagedBy = core.StringPtr("fully/self")
				Expect(sparkEngineDetailsPrototypeModel.ApiKey).To(Equal(core.StringPtr("apikey")))
				Expect(sparkEngineDetailsPrototypeModel.ConnectionString).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(sparkEngineDetailsPrototypeModel.InstanceID).To(Equal(core.StringPtr("spark-id")))
				Expect(sparkEngineDetailsPrototypeModel.ManagedBy).To(Equal(core.StringPtr("fully/self")))

				// Construct an instance of the CreateSparkEngineOptions model
				createSparkEngineOptionsOrigin := "native"
				createSparkEngineOptionsType := "spark"
				createSparkEngineOptionsModel := watsonxDataService.NewCreateSparkEngineOptions(createSparkEngineOptionsOrigin, createSparkEngineOptionsType)
				createSparkEngineOptionsModel.SetOrigin("native")
				createSparkEngineOptionsModel.SetType("spark")
				createSparkEngineOptionsModel.SetDescription("spark engine description")
				createSparkEngineOptionsModel.SetEngineDetails(sparkEngineDetailsPrototypeModel)
				createSparkEngineOptionsModel.SetEngineDisplayName("sampleEngine")
				createSparkEngineOptionsModel.SetTags([]string{"tag1", "tag2"})
				createSparkEngineOptionsModel.SetAuthInstanceID("testString")
				createSparkEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSparkEngineOptionsModel).ToNot(BeNil())
				Expect(createSparkEngineOptionsModel.Origin).To(Equal(core.StringPtr("native")))
				Expect(createSparkEngineOptionsModel.Type).To(Equal(core.StringPtr("spark")))
				Expect(createSparkEngineOptionsModel.Description).To(Equal(core.StringPtr("spark engine description")))
				Expect(createSparkEngineOptionsModel.EngineDetails).To(Equal(sparkEngineDetailsPrototypeModel))
				Expect(createSparkEngineOptionsModel.EngineDisplayName).To(Equal(core.StringPtr("sampleEngine")))
				Expect(createSparkEngineOptionsModel.Tags).To(Equal([]string{"tag1", "tag2"}))
				Expect(createSparkEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSparkEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketRegistrationOptions successfully`, func() {
				// Construct an instance of the DeleteBucketRegistrationOptions model
				bucketID := "testString"
				deleteBucketRegistrationOptionsModel := watsonxDataService.NewDeleteBucketRegistrationOptions(bucketID)
				deleteBucketRegistrationOptionsModel.SetBucketID("testString")
				deleteBucketRegistrationOptionsModel.SetAuthInstanceID("testString")
				deleteBucketRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketRegistrationOptionsModel).ToNot(BeNil())
				Expect(deleteBucketRegistrationOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketRegistrationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDatabaseCatalogOptions successfully`, func() {
				// Construct an instance of the DeleteDatabaseCatalogOptions model
				databaseID := "testString"
				deleteDatabaseCatalogOptionsModel := watsonxDataService.NewDeleteDatabaseCatalogOptions(databaseID)
				deleteDatabaseCatalogOptionsModel.SetDatabaseID("testString")
				deleteDatabaseCatalogOptionsModel.SetAuthInstanceID("testString")
				deleteDatabaseCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseCatalogOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseCatalogOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDb2EngineOptions successfully`, func() {
				// Construct an instance of the DeleteDb2EngineOptions model
				engineID := "testString"
				deleteDb2EngineOptionsModel := watsonxDataService.NewDeleteDb2EngineOptions(engineID)
				deleteDb2EngineOptionsModel.SetEngineID("testString")
				deleteDb2EngineOptionsModel.SetAuthInstanceID("testString")
				deleteDb2EngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDb2EngineOptionsModel).ToNot(BeNil())
				Expect(deleteDb2EngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDb2EngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDb2EngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDeactivateBucketOptions successfully`, func() {
				// Construct an instance of the DeleteDeactivateBucketOptions model
				bucketID := "testString"
				deleteDeactivateBucketOptionsModel := watsonxDataService.NewDeleteDeactivateBucketOptions(bucketID)
				deleteDeactivateBucketOptionsModel.SetBucketID("testString")
				deleteDeactivateBucketOptionsModel.SetAuthInstanceID("testString")
				deleteDeactivateBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDeactivateBucketOptionsModel).ToNot(BeNil())
				Expect(deleteDeactivateBucketOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDeactivateBucketOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDeactivateBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEngineOptions successfully`, func() {
				// Construct an instance of the DeleteEngineOptions model
				engineID := "testString"
				deleteEngineOptionsModel := watsonxDataService.NewDeleteEngineOptions(engineID)
				deleteEngineOptionsModel.SetEngineID("testString")
				deleteEngineOptionsModel.SetAuthInstanceID("testString")
				deleteEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEngineOptionsModel).ToNot(BeNil())
				Expect(deleteEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNetezzaEngineOptions successfully`, func() {
				// Construct an instance of the DeleteNetezzaEngineOptions model
				engineID := "testString"
				deleteNetezzaEngineOptionsModel := watsonxDataService.NewDeleteNetezzaEngineOptions(engineID)
				deleteNetezzaEngineOptionsModel.SetEngineID("testString")
				deleteNetezzaEngineOptionsModel.SetAuthInstanceID("testString")
				deleteNetezzaEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNetezzaEngineOptionsModel).ToNot(BeNil())
				Expect(deleteNetezzaEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNetezzaEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNetezzaEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteOtherEngineOptions successfully`, func() {
				// Construct an instance of the DeleteOtherEngineOptions model
				engineID := "testString"
				deleteOtherEngineOptionsModel := watsonxDataService.NewDeleteOtherEngineOptions(engineID)
				deleteOtherEngineOptionsModel.SetEngineID("testString")
				deleteOtherEngineOptionsModel.SetAuthInstanceID("testString")
				deleteOtherEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteOtherEngineOptionsModel).ToNot(BeNil())
				Expect(deleteOtherEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOtherEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOtherEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePrestoEngineCatalogsOptions successfully`, func() {
				// Construct an instance of the DeletePrestoEngineCatalogsOptions model
				engineID := "testString"
				catalogNames := "testString"
				deletePrestoEngineCatalogsOptionsModel := watsonxDataService.NewDeletePrestoEngineCatalogsOptions(engineID, catalogNames)
				deletePrestoEngineCatalogsOptionsModel.SetEngineID("testString")
				deletePrestoEngineCatalogsOptionsModel.SetCatalogNames("testString")
				deletePrestoEngineCatalogsOptionsModel.SetAuthInstanceID("testString")
				deletePrestoEngineCatalogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePrestoEngineCatalogsOptionsModel).ToNot(BeNil())
				Expect(deletePrestoEngineCatalogsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deletePrestoEngineCatalogsOptionsModel.CatalogNames).To(Equal(core.StringPtr("testString")))
				Expect(deletePrestoEngineCatalogsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deletePrestoEngineCatalogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSchemaOptions successfully`, func() {
				// Construct an instance of the DeleteSchemaOptions model
				engineID := "testString"
				catalogID := "testString"
				schemaID := "testString"
				deleteSchemaOptionsModel := watsonxDataService.NewDeleteSchemaOptions(engineID, catalogID, schemaID)
				deleteSchemaOptionsModel.SetEngineID("testString")
				deleteSchemaOptionsModel.SetCatalogID("testString")
				deleteSchemaOptionsModel.SetSchemaID("testString")
				deleteSchemaOptionsModel.SetAuthInstanceID("testString")
				deleteSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSchemaOptionsModel).ToNot(BeNil())
				Expect(deleteSchemaOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSparkEngineApplicationsOptions successfully`, func() {
				// Construct an instance of the DeleteSparkEngineApplicationsOptions model
				engineID := "testString"
				applicationID := "testString"
				deleteSparkEngineApplicationsOptionsModel := watsonxDataService.NewDeleteSparkEngineApplicationsOptions(engineID, applicationID)
				deleteSparkEngineApplicationsOptionsModel.SetEngineID("testString")
				deleteSparkEngineApplicationsOptionsModel.SetApplicationID("testString")
				deleteSparkEngineApplicationsOptionsModel.SetAuthInstanceID("testString")
				deleteSparkEngineApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSparkEngineApplicationsOptionsModel).ToNot(BeNil())
				Expect(deleteSparkEngineApplicationsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSparkEngineApplicationsOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSparkEngineApplicationsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSparkEngineApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSparkEngineOptions successfully`, func() {
				// Construct an instance of the DeleteSparkEngineOptions model
				engineID := "testString"
				deleteSparkEngineOptionsModel := watsonxDataService.NewDeleteSparkEngineOptions(engineID)
				deleteSparkEngineOptionsModel.SetEngineID("testString")
				deleteSparkEngineOptionsModel.SetAuthInstanceID("testString")
				deleteSparkEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSparkEngineOptionsModel).ToNot(BeNil())
				Expect(deleteSparkEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSparkEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSparkEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTableOptions successfully`, func() {
				// Construct an instance of the DeleteTableOptions model
				catalogID := "testString"
				schemaID := "testString"
				tableID := "testString"
				engineID := "testString"
				deleteTableOptionsModel := watsonxDataService.NewDeleteTableOptions(catalogID, schemaID, tableID, engineID)
				deleteTableOptionsModel.SetCatalogID("testString")
				deleteTableOptionsModel.SetSchemaID("testString")
				deleteTableOptionsModel.SetTableID("testString")
				deleteTableOptionsModel.SetEngineID("testString")
				deleteTableOptionsModel.SetAuthInstanceID("testString")
				deleteTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTableOptionsModel).ToNot(BeNil())
				Expect(deleteTableOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.TableID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketRegistrationOptions successfully`, func() {
				// Construct an instance of the GetBucketRegistrationOptions model
				bucketID := "testString"
				getBucketRegistrationOptionsModel := watsonxDataService.NewGetBucketRegistrationOptions(bucketID)
				getBucketRegistrationOptionsModel.SetBucketID("testString")
				getBucketRegistrationOptionsModel.SetAuthInstanceID("testString")
				getBucketRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketRegistrationOptionsModel).ToNot(BeNil())
				Expect(getBucketRegistrationOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketRegistrationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getBucketRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				catalogID := "testString"
				getCatalogOptionsModel := watsonxDataService.NewGetCatalogOptions(catalogID)
				getCatalogOptionsModel.SetCatalogID("testString")
				getCatalogOptionsModel.SetAuthInstanceID("testString")
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDatabaseOptions successfully`, func() {
				// Construct an instance of the GetDatabaseOptions model
				databaseID := "testString"
				getDatabaseOptionsModel := watsonxDataService.NewGetDatabaseOptions(databaseID)
				getDatabaseOptionsModel.SetDatabaseID("testString")
				getDatabaseOptionsModel.SetAuthInstanceID("testString")
				getDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDatabaseOptionsModel).ToNot(BeNil())
				Expect(getDatabaseOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(getDatabaseOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentsOptions successfully`, func() {
				// Construct an instance of the GetDeploymentsOptions model
				getDeploymentsOptionsModel := watsonxDataService.NewGetDeploymentsOptions()
				getDeploymentsOptionsModel.SetAuthInstanceID("testString")
				getDeploymentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentsOptionsModel).ToNot(BeNil())
				Expect(getDeploymentsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPrestoEngineCatalogOptions successfully`, func() {
				// Construct an instance of the GetPrestoEngineCatalogOptions model
				engineID := "testString"
				catalogID := "testString"
				getPrestoEngineCatalogOptionsModel := watsonxDataService.NewGetPrestoEngineCatalogOptions(engineID, catalogID)
				getPrestoEngineCatalogOptionsModel.SetEngineID("testString")
				getPrestoEngineCatalogOptionsModel.SetCatalogID("testString")
				getPrestoEngineCatalogOptionsModel.SetAuthInstanceID("testString")
				getPrestoEngineCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPrestoEngineCatalogOptionsModel).ToNot(BeNil())
				Expect(getPrestoEngineCatalogOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getPrestoEngineCatalogOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(getPrestoEngineCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPrestoEngineCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPrestoEngineOptions successfully`, func() {
				// Construct an instance of the GetPrestoEngineOptions model
				engineID := "testString"
				getPrestoEngineOptionsModel := watsonxDataService.NewGetPrestoEngineOptions(engineID)
				getPrestoEngineOptionsModel.SetEngineID("testString")
				getPrestoEngineOptionsModel.SetAuthInstanceID("testString")
				getPrestoEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPrestoEngineOptionsModel).ToNot(BeNil())
				Expect(getPrestoEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getPrestoEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getPrestoEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSparkEngineApplicationStatusOptions successfully`, func() {
				// Construct an instance of the GetSparkEngineApplicationStatusOptions model
				engineID := "testString"
				applicationID := "testString"
				getSparkEngineApplicationStatusOptionsModel := watsonxDataService.NewGetSparkEngineApplicationStatusOptions(engineID, applicationID)
				getSparkEngineApplicationStatusOptionsModel.SetEngineID("testString")
				getSparkEngineApplicationStatusOptionsModel.SetApplicationID("testString")
				getSparkEngineApplicationStatusOptionsModel.SetAuthInstanceID("testString")
				getSparkEngineApplicationStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSparkEngineApplicationStatusOptionsModel).ToNot(BeNil())
				Expect(getSparkEngineApplicationStatusOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getSparkEngineApplicationStatusOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getSparkEngineApplicationStatusOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getSparkEngineApplicationStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTableOptions successfully`, func() {
				// Construct an instance of the GetTableOptions model
				catalogID := "testString"
				schemaID := "testString"
				tableID := "testString"
				engineID := "testString"
				getTableOptionsModel := watsonxDataService.NewGetTableOptions(catalogID, schemaID, tableID, engineID)
				getTableOptionsModel.SetCatalogID("testString")
				getTableOptionsModel.SetSchemaID("testString")
				getTableOptionsModel.SetTableID("testString")
				getTableOptionsModel.SetEngineID("testString")
				getTableOptionsModel.SetAuthInstanceID("testString")
				getTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTableOptionsModel).ToNot(BeNil())
				Expect(getTableOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.TableID).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := watsonxDataService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListBucketObjectsOptions successfully`, func() {
				// Construct an instance of the ListBucketObjectsOptions model
				bucketID := "testString"
				listBucketObjectsOptionsModel := watsonxDataService.NewListBucketObjectsOptions(bucketID)
				listBucketObjectsOptionsModel.SetBucketID("testString")
				listBucketObjectsOptionsModel.SetAuthInstanceID("testString")
				listBucketObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBucketObjectsOptionsModel).ToNot(BeNil())
				Expect(listBucketObjectsOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(listBucketObjectsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listBucketObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBucketRegistrationsOptions successfully`, func() {
				// Construct an instance of the ListBucketRegistrationsOptions model
				listBucketRegistrationsOptionsModel := watsonxDataService.NewListBucketRegistrationsOptions()
				listBucketRegistrationsOptionsModel.SetAuthInstanceID("testString")
				listBucketRegistrationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBucketRegistrationsOptionsModel).ToNot(BeNil())
				Expect(listBucketRegistrationsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listBucketRegistrationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogsOptions successfully`, func() {
				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := watsonxDataService.NewListCatalogsOptions()
				listCatalogsOptionsModel.SetAuthInstanceID("testString")
				listCatalogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogsOptionsModel).ToNot(BeNil())
				Expect(listCatalogsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDatabaseRegistrationsOptions successfully`, func() {
				// Construct an instance of the ListDatabaseRegistrationsOptions model
				listDatabaseRegistrationsOptionsModel := watsonxDataService.NewListDatabaseRegistrationsOptions()
				listDatabaseRegistrationsOptionsModel.SetAuthInstanceID("testString")
				listDatabaseRegistrationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDatabaseRegistrationsOptionsModel).ToNot(BeNil())
				Expect(listDatabaseRegistrationsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDatabaseRegistrationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDb2EnginesOptions successfully`, func() {
				// Construct an instance of the ListDb2EnginesOptions model
				listDb2EnginesOptionsModel := watsonxDataService.NewListDb2EnginesOptions()
				listDb2EnginesOptionsModel.SetAuthInstanceID("testString")
				listDb2EnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDb2EnginesOptionsModel).ToNot(BeNil())
				Expect(listDb2EnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDb2EnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnginesOptions successfully`, func() {
				// Construct an instance of the ListEnginesOptions model
				listEnginesOptionsModel := watsonxDataService.NewListEnginesOptions()
				listEnginesOptionsModel.SetAuthInstanceID("testString")
				listEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnginesOptionsModel).ToNot(BeNil())
				Expect(listEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNetezzaEnginesOptions successfully`, func() {
				// Construct an instance of the ListNetezzaEnginesOptions model
				listNetezzaEnginesOptionsModel := watsonxDataService.NewListNetezzaEnginesOptions()
				listNetezzaEnginesOptionsModel.SetAuthInstanceID("testString")
				listNetezzaEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNetezzaEnginesOptionsModel).ToNot(BeNil())
				Expect(listNetezzaEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listNetezzaEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOtherEnginesOptions successfully`, func() {
				// Construct an instance of the ListOtherEnginesOptions model
				listOtherEnginesOptionsModel := watsonxDataService.NewListOtherEnginesOptions()
				listOtherEnginesOptionsModel.SetAuthInstanceID("testString")
				listOtherEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOtherEnginesOptionsModel).ToNot(BeNil())
				Expect(listOtherEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listOtherEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPrestoEngineCatalogsOptions successfully`, func() {
				// Construct an instance of the ListPrestoEngineCatalogsOptions model
				engineID := "testString"
				listPrestoEngineCatalogsOptionsModel := watsonxDataService.NewListPrestoEngineCatalogsOptions(engineID)
				listPrestoEngineCatalogsOptionsModel.SetEngineID("testString")
				listPrestoEngineCatalogsOptionsModel.SetAuthInstanceID("testString")
				listPrestoEngineCatalogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPrestoEngineCatalogsOptionsModel).ToNot(BeNil())
				Expect(listPrestoEngineCatalogsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(listPrestoEngineCatalogsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPrestoEngineCatalogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPrestoEnginesOptions successfully`, func() {
				// Construct an instance of the ListPrestoEnginesOptions model
				listPrestoEnginesOptionsModel := watsonxDataService.NewListPrestoEnginesOptions()
				listPrestoEnginesOptionsModel.SetAuthInstanceID("testString")
				listPrestoEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPrestoEnginesOptionsModel).ToNot(BeNil())
				Expect(listPrestoEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listPrestoEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSchemasOptions successfully`, func() {
				// Construct an instance of the ListSchemasOptions model
				engineID := "testString"
				catalogID := "testString"
				listSchemasOptionsModel := watsonxDataService.NewListSchemasOptions(engineID, catalogID)
				listSchemasOptionsModel.SetEngineID("testString")
				listSchemasOptionsModel.SetCatalogID("testString")
				listSchemasOptionsModel.SetAuthInstanceID("testString")
				listSchemasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSchemasOptionsModel).ToNot(BeNil())
				Expect(listSchemasOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(listSchemasOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(listSchemasOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSchemasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSparkEngineApplicationsOptions successfully`, func() {
				// Construct an instance of the ListSparkEngineApplicationsOptions model
				engineID := "testString"
				listSparkEngineApplicationsOptionsModel := watsonxDataService.NewListSparkEngineApplicationsOptions(engineID)
				listSparkEngineApplicationsOptionsModel.SetEngineID("testString")
				listSparkEngineApplicationsOptionsModel.SetAuthInstanceID("testString")
				listSparkEngineApplicationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSparkEngineApplicationsOptionsModel).ToNot(BeNil())
				Expect(listSparkEngineApplicationsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(listSparkEngineApplicationsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSparkEngineApplicationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSparkEnginesOptions successfully`, func() {
				// Construct an instance of the ListSparkEnginesOptions model
				listSparkEnginesOptionsModel := watsonxDataService.NewListSparkEnginesOptions()
				listSparkEnginesOptionsModel.SetAuthInstanceID("testString")
				listSparkEnginesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSparkEnginesOptionsModel).ToNot(BeNil())
				Expect(listSparkEnginesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSparkEnginesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTableSnapshotsOptions successfully`, func() {
				// Construct an instance of the ListTableSnapshotsOptions model
				engineID := "testString"
				catalogID := "testString"
				schemaID := "testString"
				tableID := "testString"
				listTableSnapshotsOptionsModel := watsonxDataService.NewListTableSnapshotsOptions(engineID, catalogID, schemaID, tableID)
				listTableSnapshotsOptionsModel.SetEngineID("testString")
				listTableSnapshotsOptionsModel.SetCatalogID("testString")
				listTableSnapshotsOptionsModel.SetSchemaID("testString")
				listTableSnapshotsOptionsModel.SetTableID("testString")
				listTableSnapshotsOptionsModel.SetAuthInstanceID("testString")
				listTableSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTableSnapshotsOptionsModel).ToNot(BeNil())
				Expect(listTableSnapshotsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(listTableSnapshotsOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(listTableSnapshotsOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(listTableSnapshotsOptionsModel.TableID).To(Equal(core.StringPtr("testString")))
				Expect(listTableSnapshotsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTableSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTablesOptions successfully`, func() {
				// Construct an instance of the ListTablesOptions model
				catalogID := "testString"
				schemaID := "testString"
				engineID := "testString"
				listTablesOptionsModel := watsonxDataService.NewListTablesOptions(catalogID, schemaID, engineID)
				listTablesOptionsModel.SetCatalogID("testString")
				listTablesOptionsModel.SetSchemaID("testString")
				listTablesOptionsModel.SetEngineID("testString")
				listTablesOptionsModel.SetAuthInstanceID("testString")
				listTablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTablesOptionsModel).ToNot(BeNil())
				Expect(listTablesOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRegisterDatabaseCatalogBodyDatabaseDetails successfully`, func() {
				hostname := "db2@<hostname>.com"
				port := int64(4553)
				_model, err := watsonxDataService.NewRegisterDatabaseCatalogBodyDatabaseDetails(hostname, port)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRegisterDatabaseCatalogBodyDatabasePropertiesItems successfully`, func() {
				encrypt := true
				key := "hive.metastore"
				value := "glue"
				_model, err := watsonxDataService.NewRegisterDatabaseCatalogBodyDatabasePropertiesItems(encrypt, key, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplacePrestoEngineCatalogsOptions successfully`, func() {
				// Construct an instance of the ReplacePrestoEngineCatalogsOptions model
				engineID := "testString"
				catalogNames := "testString"
				replacePrestoEngineCatalogsOptionsModel := watsonxDataService.NewReplacePrestoEngineCatalogsOptions(engineID, catalogNames)
				replacePrestoEngineCatalogsOptionsModel.SetEngineID("testString")
				replacePrestoEngineCatalogsOptionsModel.SetCatalogNames("testString")
				replacePrestoEngineCatalogsOptionsModel.SetAuthInstanceID("testString")
				replacePrestoEngineCatalogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replacePrestoEngineCatalogsOptionsModel).ToNot(BeNil())
				Expect(replacePrestoEngineCatalogsOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(replacePrestoEngineCatalogsOptionsModel.CatalogNames).To(Equal(core.StringPtr("testString")))
				Expect(replacePrestoEngineCatalogsOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replacePrestoEngineCatalogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceSnapshotOptions successfully`, func() {
				// Construct an instance of the ReplaceSnapshotOptions model
				engineID := "testString"
				catalogID := "testString"
				schemaID := "testString"
				tableID := "testString"
				snapshotID := "testString"
				replaceSnapshotOptionsModel := watsonxDataService.NewReplaceSnapshotOptions(engineID, catalogID, schemaID, tableID, snapshotID)
				replaceSnapshotOptionsModel.SetEngineID("testString")
				replaceSnapshotOptionsModel.SetCatalogID("testString")
				replaceSnapshotOptionsModel.SetSchemaID("testString")
				replaceSnapshotOptionsModel.SetTableID("testString")
				replaceSnapshotOptionsModel.SetSnapshotID("testString")
				replaceSnapshotOptionsModel.SetAuthInstanceID("testString")
				replaceSnapshotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceSnapshotOptionsModel).ToNot(BeNil())
				Expect(replaceSnapshotOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.TableID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.SnapshotID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSnapshotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunExplainAnalyzeStatementOptions successfully`, func() {
				// Construct an instance of the RunExplainAnalyzeStatementOptions model
				engineID := "testString"
				runExplainAnalyzeStatementOptionsStatement := "show schemas in catalog_name"
				runExplainAnalyzeStatementOptionsModel := watsonxDataService.NewRunExplainAnalyzeStatementOptions(engineID, runExplainAnalyzeStatementOptionsStatement)
				runExplainAnalyzeStatementOptionsModel.SetEngineID("testString")
				runExplainAnalyzeStatementOptionsModel.SetStatement("show schemas in catalog_name")
				runExplainAnalyzeStatementOptionsModel.SetVerbose(true)
				runExplainAnalyzeStatementOptionsModel.SetAuthInstanceID("testString")
				runExplainAnalyzeStatementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runExplainAnalyzeStatementOptionsModel).ToNot(BeNil())
				Expect(runExplainAnalyzeStatementOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(runExplainAnalyzeStatementOptionsModel.Statement).To(Equal(core.StringPtr("show schemas in catalog_name")))
				Expect(runExplainAnalyzeStatementOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(runExplainAnalyzeStatementOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(runExplainAnalyzeStatementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunExplainStatementOptions successfully`, func() {
				// Construct an instance of the RunExplainStatementOptions model
				engineID := "testString"
				runExplainStatementOptionsStatement := "show schemas in catalog_name"
				runExplainStatementOptionsModel := watsonxDataService.NewRunExplainStatementOptions(engineID, runExplainStatementOptionsStatement)
				runExplainStatementOptionsModel.SetEngineID("testString")
				runExplainStatementOptionsModel.SetStatement("show schemas in catalog_name")
				runExplainStatementOptionsModel.SetFormat("json")
				runExplainStatementOptionsModel.SetType("io")
				runExplainStatementOptionsModel.SetAuthInstanceID("testString")
				runExplainStatementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runExplainStatementOptionsModel).ToNot(BeNil())
				Expect(runExplainStatementOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(runExplainStatementOptionsModel.Statement).To(Equal(core.StringPtr("show schemas in catalog_name")))
				Expect(runExplainStatementOptionsModel.Format).To(Equal(core.StringPtr("json")))
				Expect(runExplainStatementOptionsModel.Type).To(Equal(core.StringPtr("io")))
				Expect(runExplainStatementOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(runExplainStatementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSparkApplicationDetails successfully`, func() {
				application := "s3://mybucket/wordcount.py"
				arguments := []string{"people.txt"}
				conf := map[string]string{"key1": "key:value"}
				env := map[string]string{"key1": "key:value"}
				_model, err := watsonxDataService.NewSparkApplicationDetails(application, arguments, conf, env)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTestBucketConnectionOptions successfully`, func() {
				// Construct an instance of the TestBucketConnectionOptions model
				testBucketConnectionOptionsAccessKey := "<access_key>"
				testBucketConnectionOptionsBucketName := "sample-bucket"
				testBucketConnectionOptionsBucketType := "ibm_cos"
				testBucketConnectionOptionsEndpoint := "https://s3.<region>.cloud-object-storage.appdomain.cloud/"
				testBucketConnectionOptionsRegion := "us-south"
				testBucketConnectionOptionsSecretKey := "secret_key"
				testBucketConnectionOptionsModel := watsonxDataService.NewTestBucketConnectionOptions(testBucketConnectionOptionsAccessKey, testBucketConnectionOptionsBucketName, testBucketConnectionOptionsBucketType, testBucketConnectionOptionsEndpoint, testBucketConnectionOptionsRegion, testBucketConnectionOptionsSecretKey)
				testBucketConnectionOptionsModel.SetAccessKey("<access_key>")
				testBucketConnectionOptionsModel.SetBucketName("sample-bucket")
				testBucketConnectionOptionsModel.SetBucketType("ibm_cos")
				testBucketConnectionOptionsModel.SetEndpoint("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				testBucketConnectionOptionsModel.SetRegion("us-south")
				testBucketConnectionOptionsModel.SetSecretKey("secret_key")
				testBucketConnectionOptionsModel.SetAuthInstanceID("testString")
				testBucketConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testBucketConnectionOptionsModel).ToNot(BeNil())
				Expect(testBucketConnectionOptionsModel.AccessKey).To(Equal(core.StringPtr("<access_key>")))
				Expect(testBucketConnectionOptionsModel.BucketName).To(Equal(core.StringPtr("sample-bucket")))
				Expect(testBucketConnectionOptionsModel.BucketType).To(Equal(core.StringPtr("ibm_cos")))
				Expect(testBucketConnectionOptionsModel.Endpoint).To(Equal(core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")))
				Expect(testBucketConnectionOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(testBucketConnectionOptionsModel.SecretKey).To(Equal(core.StringPtr("secret_key")))
				Expect(testBucketConnectionOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(testBucketConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTestLHConsoleOptions successfully`, func() {
				// Construct an instance of the TestLHConsoleOptions model
				testLhConsoleOptionsModel := watsonxDataService.NewTestLHConsoleOptions()
				testLhConsoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testLhConsoleOptionsModel).ToNot(BeNil())
				Expect(testLhConsoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBucketRegistrationOKBodyPatch successfully`, func() {
				// Construct an instance of the BucketRegistration model
				bucketRegistrationModel := new(watsonxdatav2.BucketRegistration)
				bucketRegistrationModel.AccessKey = core.StringPtr("<access_key>")
				bucketRegistrationModel.Actions = []string{"create", "update"}
				bucketRegistrationModel.AssociatedCatalogs = []string{"iceberg_catalog", "hive_catalog"}
				bucketRegistrationModel.BucketDisplayName = core.StringPtr("samplebucketdisplayname")
				bucketRegistrationModel.BucketID = core.StringPtr("samplebucketid")
				bucketRegistrationModel.BucketName = core.StringPtr("samplebucket")
				bucketRegistrationModel.BucketType = core.StringPtr("minio")
				bucketRegistrationModel.CreatedBy = core.StringPtr("username@domain.com")
				bucketRegistrationModel.CreatedOn = core.StringPtr("1699457595")
				bucketRegistrationModel.Description = core.StringPtr("default bucket")
				bucketRegistrationModel.Endpoint = core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/")
				bucketRegistrationModel.ManagedBy = core.StringPtr("ibm")
				bucketRegistrationModel.Region = core.StringPtr("us-south")
				bucketRegistrationModel.SecretKey = core.StringPtr("secret_key")
				bucketRegistrationModel.State = core.StringPtr("active")
				bucketRegistrationModel.Tags = []string{"tag1", "tag2"}

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("Update bucket details")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateBucketRegistrationOKBody model
				updateBucketRegistrationOKBody := new(watsonxdatav2.UpdateBucketRegistrationOKBody)
				updateBucketRegistrationOKBody.BucketRegistration = bucketRegistrationModel
				updateBucketRegistrationOKBody.Response = successResponseModel

				updateBucketRegistrationOKBodyPatch := watsonxDataService.NewUpdateBucketRegistrationOKBodyPatch(updateBucketRegistrationOKBody)
				Expect(updateBucketRegistrationOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateBucketRegistrationOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/bucket_registration": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/bucket_registration")),
					"From": BeNil(),
					"Value": Equal(updateBucketRegistrationOKBody.BucketRegistration),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateBucketRegistrationOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateBucketRegistrationOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateBucketRegistrationOptions model
				bucketID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateBucketRegistrationOptionsModel := watsonxDataService.NewUpdateBucketRegistrationOptions(bucketID, body)
				updateBucketRegistrationOptionsModel.SetBucketID("testString")
				updateBucketRegistrationOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateBucketRegistrationOptionsModel.SetAuthInstanceID("testString")
				updateBucketRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBucketRegistrationOptionsModel).ToNot(BeNil())
				Expect(updateBucketRegistrationOptionsModel.BucketID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketRegistrationOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateBucketRegistrationOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDatabaseOKBodyPatch successfully`, func() {
				// Construct an instance of the DatabaseRegistrationDatabaseDetails model
				databaseRegistrationDatabaseDetailsModel := new(watsonxdatav2.DatabaseRegistrationDatabaseDetails)
				databaseRegistrationDatabaseDetailsModel.DatabaseName = core.StringPtr("new_database")
				databaseRegistrationDatabaseDetailsModel.Hostname = core.StringPtr("netezza://abc.efg.com")
				databaseRegistrationDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				databaseRegistrationDatabaseDetailsModel.Port = core.Int64Ptr(int64(4353))
				databaseRegistrationDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				databaseRegistrationDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				databaseRegistrationDatabaseDetailsModel.Tables = core.StringPtr("netezza_table_name")
				databaseRegistrationDatabaseDetailsModel.Username = core.StringPtr("sampleuser")

				// Construct an instance of the DatabaseRegistration model
				databaseRegistrationModel := new(watsonxdatav2.DatabaseRegistration)
				databaseRegistrationModel.Actions = []string{"update", "delete"}
				databaseRegistrationModel.AssociatedCatalogs = []string{"iceberg_catalog", "hive_catalog"}
				databaseRegistrationModel.CreatedBy = core.StringPtr("user1@bim.com")
				databaseRegistrationModel.CreatedOn = core.StringPtr("1686792721")
				databaseRegistrationModel.DatabaseDetails = databaseRegistrationDatabaseDetailsModel
				databaseRegistrationModel.DatabaseDisplayName = core.StringPtr("new_database,")
				databaseRegistrationModel.DatabaseID = core.StringPtr("new_database_id,")
				databaseRegistrationModel.DatabaseProperties = []string{"key1", "key2"}
				databaseRegistrationModel.DatabaseType = core.StringPtr("netezza")
				databaseRegistrationModel.Description = core.StringPtr("Description of the database")
				databaseRegistrationModel.Tags = []string{"testdatabase", "userdatabase"}

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("Update database")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateDatabaseOKBody model
				updateDatabaseOKBody := new(watsonxdatav2.UpdateDatabaseOKBody)
				updateDatabaseOKBody.Database = databaseRegistrationModel
				updateDatabaseOKBody.Response = successResponseModel

				updateDatabaseOKBodyPatch := watsonxDataService.NewUpdateDatabaseOKBodyPatch(updateDatabaseOKBody)
				Expect(updateDatabaseOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateDatabaseOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/database": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/database")),
					"From": BeNil(),
					"Value": Equal(updateDatabaseOKBody.Database),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateDatabaseOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateDatabaseOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateDatabaseOptions model
				databaseID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateDatabaseOptionsModel := watsonxDataService.NewUpdateDatabaseOptions(databaseID, body)
				updateDatabaseOptionsModel.SetDatabaseID("testString")
				updateDatabaseOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateDatabaseOptionsModel.SetAuthInstanceID("testString")
				updateDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDatabaseOptionsModel).ToNot(BeNil())
				Expect(updateDatabaseOptionsModel.DatabaseID).To(Equal(core.StringPtr("testString")))
				Expect(updateDatabaseOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDatabaseOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDb2EngineOKBodyPatch successfully`, func() {
				// Construct an instance of the Db2EngineDetails model
				db2EngineDetailsModel := new(watsonxdatav2.Db2EngineDetails)
				db2EngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				db2EngineDetailsModel.MetastoreHost = core.StringPtr("thrift://mh-connection-string-sample.com")

				// Construct an instance of the Db2Engine model
				db2EngineModel := new(watsonxdatav2.Db2Engine)
				db2EngineModel.Actions = []string{"update", "delete"}
				db2EngineModel.BuildVersion = core.StringPtr("1.0.3.0.0")
				db2EngineModel.CreatedBy = core.StringPtr("user@test.com")
				db2EngineModel.CreatedOn = core.Int64Ptr(int64(1700322469))
				db2EngineModel.Description = core.StringPtr("updated description for db2 engine.")
				db2EngineModel.EngineDetails = db2EngineDetailsModel
				db2EngineModel.EngineDisplayName = core.StringPtr("sample db2 Engine Display Name")
				db2EngineModel.EngineID = core.StringPtr("sample db2 Engine Name")
				db2EngineModel.HostName = core.StringPtr("xyz-db2-01-db2-svc")
				db2EngineModel.Origin = core.StringPtr("external")
				db2EngineModel.Port = core.Int64Ptr(int64(38))
				db2EngineModel.Status = core.StringPtr("REGISTERED")
				db2EngineModel.Tags = []string{"tag1", "tag2"}
				db2EngineModel.Type = core.StringPtr("db2")

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("update db2 engine")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateDb2EngineOKBody model
				updateDb2EngineOKBody := new(watsonxdatav2.UpdateDb2EngineOKBody)
				updateDb2EngineOKBody.Db2Engine = db2EngineModel
				updateDb2EngineOKBody.Response = successResponseModel

				updateDb2EngineOKBodyPatch := watsonxDataService.NewUpdateDb2EngineOKBodyPatch(updateDb2EngineOKBody)
				Expect(updateDb2EngineOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateDb2EngineOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/db2_engine": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/db2_engine")),
					"From": BeNil(),
					"Value": Equal(updateDb2EngineOKBody.Db2Engine),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateDb2EngineOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateDb2EngineOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateDb2EngineOptions model
				engineID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateDb2EngineOptionsModel := watsonxDataService.NewUpdateDb2EngineOptions(engineID, body)
				updateDb2EngineOptionsModel.SetEngineID("testString")
				updateDb2EngineOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateDb2EngineOptionsModel.SetAuthInstanceID("testString")
				updateDb2EngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDb2EngineOptionsModel).ToNot(BeNil())
				Expect(updateDb2EngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateDb2EngineOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDb2EngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDb2EngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEngineOKBodyPatch successfully`, func() {
				// Construct an instance of the NodeDescription model
				nodeDescriptionModel := new(watsonxdatav2.NodeDescription)
				nodeDescriptionModel.NodeType = core.StringPtr("worker")
				nodeDescriptionModel.Quantity = core.Int64Ptr(int64(1))

				// Construct an instance of the Endpoints model
				endpointsModel := new(watsonxdatav2.Endpoints)
				endpointsModel.ApplicationsApi = core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>")
				endpointsModel.HistoryServerEndpoint = core.StringPtr("$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server")
				endpointsModel.SparkAccessEndpoint = core.StringPtr("$HOST/analytics-engine/details/spark-<instance_id>")
				endpointsModel.SparkJobsV4Endpoint = core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications")
				endpointsModel.SparkKernelEndpoint = core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels")
				endpointsModel.ViewHistoryServer = core.StringPtr("testString")
				endpointsModel.WxdApplicationEndpoint = core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/applications")

				// Construct an instance of the EngineDetails model
				engineDetailsModel := new(watsonxdatav2.EngineDetails)
				engineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				engineDetailsModel.Endpoints = endpointsModel
				engineDetailsModel.MetastoreHost = core.StringPtr("1.2.3.4")

				// Construct an instance of the PrestoEngine model
				prestoEngineModel := new(watsonxdatav2.PrestoEngine)
				prestoEngineModel.Actions = []string{"update", "delete"}
				prestoEngineModel.AssociatedCatalogs = []string{"new_catalog_1", "new_catalog_2"}
				prestoEngineModel.BuildVersion = core.StringPtr("1.0.3.0.0")
				prestoEngineModel.Coordinator = nodeDescriptionModel
				prestoEngineModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				prestoEngineModel.CreatedOn = core.Int64Ptr(int64(163788384993))
				prestoEngineModel.Description = core.StringPtr("updated description for presto engine")
				prestoEngineModel.EngineDetails = engineDetailsModel
				prestoEngineModel.EngineDisplayName = core.StringPtr("sampleEngine")
				prestoEngineModel.EngineID = core.StringPtr("sampleEngine123")
				prestoEngineModel.ExternalHostName = core.StringPtr("your-hostname.apps.your-domain.com")
				prestoEngineModel.GroupID = core.StringPtr("new_group_id")
				prestoEngineModel.HostName = core.StringPtr("your-hostname.apps.your-domain.com")
				prestoEngineModel.Origin = core.StringPtr("ibm")
				prestoEngineModel.Port = core.Int64Ptr(int64(38))
				prestoEngineModel.Region = core.StringPtr("us-south")
				prestoEngineModel.SizeConfig = core.StringPtr("starter")
				prestoEngineModel.Status = core.StringPtr("running")
				prestoEngineModel.StatusCode = core.Int64Ptr(int64(0))
				prestoEngineModel.Tags = []string{"tag1", "tag2"}
				prestoEngineModel.Type = core.StringPtr("presto")
				prestoEngineModel.Version = core.StringPtr("1.2.0")
				prestoEngineModel.Worker = nodeDescriptionModel

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("update engine")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateEngineOKBody model
				updateEngineOKBody := new(watsonxdatav2.UpdateEngineOKBody)
				updateEngineOKBody.Engine = prestoEngineModel
				updateEngineOKBody.Response = successResponseModel

				updateEngineOKBodyPatch := watsonxDataService.NewUpdateEngineOKBodyPatch(updateEngineOKBody)
				Expect(updateEngineOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateEngineOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/engine": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/engine")),
					"From": BeNil(),
					"Value": Equal(updateEngineOKBody.Engine),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateEngineOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateEngineOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateEngineOptions model
				engineID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateEngineOptionsModel := watsonxDataService.NewUpdateEngineOptions(engineID, body)
				updateEngineOptionsModel.SetEngineID("testString")
				updateEngineOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateEngineOptionsModel.SetAuthInstanceID("testString")
				updateEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEngineOptionsModel).ToNot(BeNil())
				Expect(updateEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateNetezzaEngineOKBodyPatch successfully`, func() {
				// Construct an instance of the NetezzaEngineDetails model
				netezzaEngineDetailsModel := new(watsonxdatav2.NetezzaEngineDetails)
				netezzaEngineDetailsModel.ConnectionString = core.StringPtr("1.2.3.4")
				netezzaEngineDetailsModel.MetastoreHost = core.StringPtr("thrift://mh-connection-string-sample.com")

				// Construct an instance of the NetezzaEngine model
				netezzaEngineModel := new(watsonxdatav2.NetezzaEngine)
				netezzaEngineModel.Actions = []string{"update", "delete"}
				netezzaEngineModel.BuildVersion = core.StringPtr("1.0.3.0.0")
				netezzaEngineModel.CreatedBy = core.StringPtr("user@test.com")
				netezzaEngineModel.CreatedOn = core.Int64Ptr(int64(1700322469))
				netezzaEngineModel.Description = core.StringPtr("updated description for netezza engine.")
				netezzaEngineModel.EngineDetails = netezzaEngineDetailsModel
				netezzaEngineModel.EngineDisplayName = core.StringPtr("sample Netezza Engine Display Name")
				netezzaEngineModel.EngineID = core.StringPtr("sample Netezza Engine Name")
				netezzaEngineModel.HostName = core.StringPtr("xyz-netezza-01-netezza-svc")
				netezzaEngineModel.Origin = core.StringPtr("external")
				netezzaEngineModel.Port = core.Int64Ptr(int64(38))
				netezzaEngineModel.Status = core.StringPtr("REGISTERED")
				netezzaEngineModel.Tags = []string{"tag1", "tag2"}
				netezzaEngineModel.Type = core.StringPtr("netezza")

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("update netezza engine")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateNetezzaEngineOKBody model
				updateNetezzaEngineOKBody := new(watsonxdatav2.UpdateNetezzaEngineOKBody)
				updateNetezzaEngineOKBody.NetezzaEngine = netezzaEngineModel
				updateNetezzaEngineOKBody.Response = successResponseModel

				updateNetezzaEngineOKBodyPatch := watsonxDataService.NewUpdateNetezzaEngineOKBodyPatch(updateNetezzaEngineOKBody)
				Expect(updateNetezzaEngineOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateNetezzaEngineOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/netezza_engine": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/netezza_engine")),
					"From": BeNil(),
					"Value": Equal(updateNetezzaEngineOKBody.NetezzaEngine),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateNetezzaEngineOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateNetezzaEngineOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateNetezzaEngineOptions model
				engineID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateNetezzaEngineOptionsModel := watsonxDataService.NewUpdateNetezzaEngineOptions(engineID, body)
				updateNetezzaEngineOptionsModel.SetEngineID("testString")
				updateNetezzaEngineOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateNetezzaEngineOptionsModel.SetAuthInstanceID("testString")
				updateNetezzaEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateNetezzaEngineOptionsModel).ToNot(BeNil())
				Expect(updateNetezzaEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateNetezzaEngineOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateNetezzaEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateNetezzaEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSparkEngineOKBodyPatch successfully`, func() {
				// Construct an instance of the SparkEngineDetailsEndpoints model
				sparkEngineDetailsEndpointsModel := new(watsonxdatav2.SparkEngineDetailsEndpoints)
				sparkEngineDetailsEndpointsModel.ApplicationsApi = core.StringPtr("$HOST/v4/analytics_engines/<spark_id>/spark_applications/<application_id>")
				sparkEngineDetailsEndpointsModel.HistoryServerEndpoint = core.StringPtr("$HOST/v2/spark/v3/instances/<spark_id>/spark_history_server")
				sparkEngineDetailsEndpointsModel.SparkAccessEndpoint = core.StringPtr("$HOST/analytics-engine/details/spark-<instance_id>")
				sparkEngineDetailsEndpointsModel.SparkJobsV4Endpoint = core.StringPtr("$HOST/v4/analytics_engines/<spark_id>/spark_applications")
				sparkEngineDetailsEndpointsModel.SparkKernelEndpoint = core.StringPtr("$HOST/v4/analytics_engines/<spark_id>/jkg/api/kernels")
				sparkEngineDetailsEndpointsModel.ViewHistoryServer = core.StringPtr("View history server")
				sparkEngineDetailsEndpointsModel.WxdApplicationEndpoint = core.StringPtr("$HOST/v1/<wxd_instance_id>/engines/<engine_id>/applications")

				// Construct an instance of the SparkEngineDetails model
				sparkEngineDetailsModel := new(watsonxdatav2.SparkEngineDetails)
				sparkEngineDetailsModel.ConnectionString = core.StringPtr("https://xyz.<region>.ae.cloud.123.com/v3/analytics_engines/<spark_iae_id>")
				sparkEngineDetailsModel.Endpoints = sparkEngineDetailsEndpointsModel

				// Construct an instance of the SparkEngine model
				sparkEngineModel := new(watsonxdatav2.SparkEngine)
				sparkEngineModel.Actions = []string{"update", "delete"}
				sparkEngineModel.BuildVersion = core.StringPtr("1.0.3.0.0")
				sparkEngineModel.CreatedBy = core.StringPtr("<username>@<domain>.com")
				sparkEngineModel.CreatedOn = core.Int64Ptr(int64(163788384993))
				sparkEngineModel.Description = core.StringPtr("Spark engines for running spark applications")
				sparkEngineModel.EngineDetails = sparkEngineDetailsModel
				sparkEngineModel.EngineDisplayName = core.StringPtr("sampleEngine")
				sparkEngineModel.EngineID = core.StringPtr("sampleEngine123")
				sparkEngineModel.Origin = core.StringPtr("discover")
				sparkEngineModel.Status = core.StringPtr("REGISTERED")
				sparkEngineModel.Tags = []string{"tag1", "tag2"}
				sparkEngineModel.Type = core.StringPtr("spark")

				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("Spark engines list")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateSparkEngineOKBody model
				updateSparkEngineOKBody := new(watsonxdatav2.UpdateSparkEngineOKBody)
				updateSparkEngineOKBody.Engine = sparkEngineModel
				updateSparkEngineOKBody.Response = successResponseModel

				updateSparkEngineOKBodyPatch := watsonxDataService.NewUpdateSparkEngineOKBodyPatch(updateSparkEngineOKBody)
				Expect(updateSparkEngineOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateSparkEngineOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/engine": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/engine")),
					"From": BeNil(),
					"Value": Equal(updateSparkEngineOKBody.Engine),
					}),
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateSparkEngineOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateSparkEngineOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateSparkEngineOptions model
				engineID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateSparkEngineOptionsModel := watsonxDataService.NewUpdateSparkEngineOptions(engineID, body)
				updateSparkEngineOptionsModel.SetEngineID("testString")
				updateSparkEngineOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateSparkEngineOptionsModel.SetAuthInstanceID("testString")
				updateSparkEngineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSparkEngineOptionsModel).ToNot(BeNil())
				Expect(updateSparkEngineOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateSparkEngineOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateSparkEngineOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSparkEngineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSyncCatalogOKBodyPatch successfully`, func() {
				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("sync catalog")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateSyncCatalogOKBody model
				updateSyncCatalogOKBody := new(watsonxdatav2.UpdateSyncCatalogOKBody)
				updateSyncCatalogOKBody.Response = successResponseModel

				updateSyncCatalogOKBodyPatch := watsonxDataService.NewUpdateSyncCatalogOKBodyPatch(updateSyncCatalogOKBody)
				Expect(updateSyncCatalogOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateSyncCatalogOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateSyncCatalogOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateSyncCatalogOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateSyncCatalogOptions model
				catalogID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateSyncCatalogOptionsModel := watsonxDataService.NewUpdateSyncCatalogOptions(catalogID, body)
				updateSyncCatalogOptionsModel.SetCatalogID("testString")
				updateSyncCatalogOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateSyncCatalogOptionsModel.SetAuthInstanceID("testString")
				updateSyncCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSyncCatalogOptionsModel).ToNot(BeNil())
				Expect(updateSyncCatalogOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(updateSyncCatalogOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateSyncCatalogOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSyncCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTableOKBodyPatch successfully`, func() {
				// Construct an instance of the SuccessResponse model
				successResponseModel := new(watsonxdatav2.SuccessResponse)
				successResponseModel.Message = core.StringPtr("update table")
				successResponseModel.MessageCode = core.StringPtr("success")

				// Construct an instance of the UpdateTableOKBody model
				updateTableOKBody := new(watsonxdatav2.UpdateTableOKBody)
				updateTableOKBody.Response = successResponseModel

				updateTableOKBodyPatch := watsonxDataService.NewUpdateTableOKBodyPatch(updateTableOKBody)
				Expect(updateTableOKBodyPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(watsonxdatav2.JSONPatchOperation).Path
				}
				Expect(updateTableOKBodyPatch).To(MatchAllElements(_path, Elements{
				"/response": MatchAllFields(Fields{
					"Op": PointTo(Equal(watsonxdatav2.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/response")),
					"From": BeNil(),
					"Value": Equal(updateTableOKBody.Response),
					}),
				}))
			})
			It(`Invoke NewUpdateTableOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(watsonxdatav2.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateTableOptions model
				catalogID := "testString"
				schemaID := "testString"
				tableID := "testString"
				engineID := "testString"
				body := []watsonxdatav2.JSONPatchOperation{}
				updateTableOptionsModel := watsonxDataService.NewUpdateTableOptions(catalogID, schemaID, tableID, engineID, body)
				updateTableOptionsModel.SetCatalogID("testString")
				updateTableOptionsModel.SetSchemaID("testString")
				updateTableOptionsModel.SetTableID("testString")
				updateTableOptionsModel.SetEngineID("testString")
				updateTableOptionsModel.SetBody([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel})
				updateTableOptionsModel.SetAuthInstanceID("testString")
				updateTableOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTableOptionsModel).ToNot(BeNil())
				Expect(updateTableOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.SchemaID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.TableID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.EngineID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.Body).To(Equal([]watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateTableOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateTableOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidateDatabaseBodyDatabaseDetails successfully`, func() {
				hostname := "db2@hostname.com"
				port := int64(4553)
				_model, err := watsonxDataService.NewValidateDatabaseBodyDatabaseDetails(hostname, port)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewValidateDatabaseConnectionOptions successfully`, func() {
				// Construct an instance of the ValidateDatabaseBodyDatabaseDetails model
				validateDatabaseBodyDatabaseDetailsModel := new(watsonxdatav2.ValidateDatabaseBodyDatabaseDetails)
				Expect(validateDatabaseBodyDatabaseDetailsModel).ToNot(BeNil())
				validateDatabaseBodyDatabaseDetailsModel.DatabaseName = core.StringPtr("sampledatabase")
				validateDatabaseBodyDatabaseDetailsModel.Hostname = core.StringPtr("db2@hostname.com")
				validateDatabaseBodyDatabaseDetailsModel.Password = core.StringPtr("samplepassword")
				validateDatabaseBodyDatabaseDetailsModel.Port = core.Int64Ptr(int64(4553))
				validateDatabaseBodyDatabaseDetailsModel.Sasl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Ssl = core.BoolPtr(true)
				validateDatabaseBodyDatabaseDetailsModel.Tables = core.StringPtr("kafka_table_name")
				validateDatabaseBodyDatabaseDetailsModel.Username = core.StringPtr("sampleuser")
				Expect(validateDatabaseBodyDatabaseDetailsModel.DatabaseName).To(Equal(core.StringPtr("sampledatabase")))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Hostname).To(Equal(core.StringPtr("db2@hostname.com")))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Password).To(Equal(core.StringPtr("samplepassword")))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Port).To(Equal(core.Int64Ptr(int64(4553))))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Sasl).To(Equal(core.BoolPtr(true)))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Ssl).To(Equal(core.BoolPtr(true)))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Tables).To(Equal(core.StringPtr("kafka_table_name")))
				Expect(validateDatabaseBodyDatabaseDetailsModel.Username).To(Equal(core.StringPtr("sampleuser")))

				// Construct an instance of the ValidateDatabaseConnectionOptions model
				var validateDatabaseConnectionOptionsDatabaseDetails *watsonxdatav2.ValidateDatabaseBodyDatabaseDetails = nil
				validateDatabaseConnectionOptionsDatabaseType := "netezza"
				validateDatabaseConnectionOptionsModel := watsonxDataService.NewValidateDatabaseConnectionOptions(validateDatabaseConnectionOptionsDatabaseDetails, validateDatabaseConnectionOptionsDatabaseType)
				validateDatabaseConnectionOptionsModel.SetDatabaseDetails(validateDatabaseBodyDatabaseDetailsModel)
				validateDatabaseConnectionOptionsModel.SetDatabaseType("netezza")
				validateDatabaseConnectionOptionsModel.SetCertificate("contents of a pem/crt file")
				validateDatabaseConnectionOptionsModel.SetAuthInstanceID("testString")
				validateDatabaseConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateDatabaseConnectionOptionsModel).ToNot(BeNil())
				Expect(validateDatabaseConnectionOptionsModel.DatabaseDetails).To(Equal(validateDatabaseBodyDatabaseDetailsModel))
				Expect(validateDatabaseConnectionOptionsModel.DatabaseType).To(Equal(core.StringPtr("netezza")))
				Expect(validateDatabaseConnectionOptionsModel.Certificate).To(Equal(core.StringPtr("contents of a pem/crt file")))
				Expect(validateDatabaseConnectionOptionsModel.AuthInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(validateDatabaseConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
