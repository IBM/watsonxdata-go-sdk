/**
 * (C) Copyright IBM Corp. 2024.
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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.84.1-55f6d880-20240110-194020
 */

// Package watsonxdatav2 : Operations and models for the WatsonxDataV2 service
package watsonxdatav2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/watsonxdata-go-sdk/common"
)

// WatsonxDataV2 : This is the Public API for IBM watsonx.data
//
// API Version: 2.0.0
type WatsonxDataV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://region.lakehouse.cloud.ibm.com/lakehouse/api/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watsonx_data"

// WatsonxDataV2Options : Service options
type WatsonxDataV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewWatsonxDataV2UsingExternalConfig : constructs an instance of WatsonxDataV2 with passed in options and external configuration.
func NewWatsonxDataV2UsingExternalConfig(options *WatsonxDataV2Options) (watsonxData *WatsonxDataV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	watsonxData, err = NewWatsonxDataV2(options)
	if err != nil {
		return
	}

	err = watsonxData.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = watsonxData.Service.SetServiceURL(options.URL)
	}
	return
}

// NewWatsonxDataV2 : constructs an instance of WatsonxDataV2 with passed in options.
func NewWatsonxDataV2(options *WatsonxDataV2Options) (service *WatsonxDataV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &WatsonxDataV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "watsonxData" suitable for processing requests.
func (watsonxData *WatsonxDataV2) Clone() *WatsonxDataV2 {
	if core.IsNil(watsonxData) {
		return nil
	}
	clone := *watsonxData
	clone.Service = watsonxData.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (watsonxData *WatsonxDataV2) SetServiceURL(url string) error {
	return watsonxData.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (watsonxData *WatsonxDataV2) GetServiceURL() string {
	return watsonxData.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (watsonxData *WatsonxDataV2) SetDefaultHeaders(headers http.Header) {
	watsonxData.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV2) SetEnableGzipCompression(enableGzip bool) {
	watsonxData.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV2) GetEnableGzipCompression() bool {
	return watsonxData.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (watsonxData *WatsonxDataV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	watsonxData.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (watsonxData *WatsonxDataV2) DisableRetries() {
	watsonxData.Service.DisableRetries()
}

// ListBucketRegistrations : Get bucket registrations
// Get list of registered buckets.
func (watsonxData *WatsonxDataV2) ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions) (result *BucketRegistrationCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListBucketRegistrationsWithContext(context.Background(), listBucketRegistrationsOptions)
}

// ListBucketRegistrationsWithContext is an alternate form of the ListBucketRegistrations method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListBucketRegistrationsWithContext(ctx context.Context, listBucketRegistrationsOptions *ListBucketRegistrationsOptions) (result *BucketRegistrationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listBucketRegistrationsOptions, "listBucketRegistrationsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBucketRegistrationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListBucketRegistrations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listBucketRegistrationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listBucketRegistrationsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistrationCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateBucketRegistration : Register bucket
// Register a new bucket.
func (watsonxData *WatsonxDataV2) CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.CreateBucketRegistrationWithContext(context.Background(), createBucketRegistrationOptions)
}

// CreateBucketRegistrationWithContext is an alternate form of the CreateBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateBucketRegistrationWithContext(ctx context.Context, createBucketRegistrationOptions *CreateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBucketRegistrationOptions, "createBucketRegistrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBucketRegistrationOptions, "createBucketRegistrationOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createBucketRegistrationOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createBucketRegistrationOptions.BucketDetails != nil {
		body["bucket_details"] = createBucketRegistrationOptions.BucketDetails
	}
	if createBucketRegistrationOptions.BucketType != nil {
		body["bucket_type"] = createBucketRegistrationOptions.BucketType
	}
	if createBucketRegistrationOptions.Description != nil {
		body["description"] = createBucketRegistrationOptions.Description
	}
	if createBucketRegistrationOptions.ManagedBy != nil {
		body["managed_by"] = createBucketRegistrationOptions.ManagedBy
	}
	if createBucketRegistrationOptions.AssociatedCatalog != nil {
		body["associated_catalog"] = createBucketRegistrationOptions.AssociatedCatalog
	}
	if createBucketRegistrationOptions.BucketDisplayName != nil {
		body["bucket_display_name"] = createBucketRegistrationOptions.BucketDisplayName
	}
	if createBucketRegistrationOptions.Region != nil {
		body["region"] = createBucketRegistrationOptions.Region
	}
	if createBucketRegistrationOptions.Tags != nil {
		body["tags"] = createBucketRegistrationOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBucketRegistration : Get bucket
// Get a registered bucket.
func (watsonxData *WatsonxDataV2) GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.GetBucketRegistrationWithContext(context.Background(), getBucketRegistrationOptions)
}

// GetBucketRegistrationWithContext is an alternate form of the GetBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetBucketRegistrationWithContext(ctx context.Context, getBucketRegistrationOptions *GetBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketRegistrationOptions, "getBucketRegistrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketRegistrationOptions, "getBucketRegistrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *getBucketRegistrationOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getBucketRegistrationOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteBucketRegistration : Unregister Bucket
// Unregister a bucket.
func (watsonxData *WatsonxDataV2) DeleteBucketRegistration(deleteBucketRegistrationOptions *DeleteBucketRegistrationOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteBucketRegistrationWithContext(context.Background(), deleteBucketRegistrationOptions)
}

// DeleteBucketRegistrationWithContext is an alternate form of the DeleteBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteBucketRegistrationWithContext(ctx context.Context, deleteBucketRegistrationOptions *DeleteBucketRegistrationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketRegistrationOptions, "deleteBucketRegistrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketRegistrationOptions, "deleteBucketRegistrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *deleteBucketRegistrationOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteBucketRegistrationOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateBucketRegistration : Update bucket
// Update bucket details & credentials.
func (watsonxData *WatsonxDataV2) UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateBucketRegistrationWithContext(context.Background(), updateBucketRegistrationOptions)
}

// UpdateBucketRegistrationWithContext is an alternate form of the UpdateBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateBucketRegistrationWithContext(ctx context.Context, updateBucketRegistrationOptions *UpdateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBucketRegistrationOptions, "updateBucketRegistrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBucketRegistrationOptions, "updateBucketRegistrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *updateBucketRegistrationOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateBucketRegistrationOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateBucketRegistrationOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateActivateBucket : Activate Bucket
// Activate a registered bucket.
func (watsonxData *WatsonxDataV2) CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions) (result *CreateActivateBucketCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateActivateBucketWithContext(context.Background(), createActivateBucketOptions)
}

// CreateActivateBucketWithContext is an alternate form of the CreateActivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateActivateBucketWithContext(ctx context.Context, createActivateBucketOptions *CreateActivateBucketOptions) (result *CreateActivateBucketCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createActivateBucketOptions, "createActivateBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createActivateBucketOptions, "createActivateBucketOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *createActivateBucketOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/activate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createActivateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateActivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createActivateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createActivateBucketOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateActivateBucketCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDeactivateBucket : Deactivate Bucket
// Deactivate a bucket.
func (watsonxData *WatsonxDataV2) DeleteDeactivateBucket(deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDeactivateBucketWithContext(context.Background(), deleteDeactivateBucketOptions)
}

// DeleteDeactivateBucketWithContext is an alternate form of the DeleteDeactivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDeactivateBucketWithContext(ctx context.Context, deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDeactivateBucketOptions, "deleteDeactivateBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDeactivateBucketOptions, "deleteDeactivateBucketOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *deleteDeactivateBucketOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/deactivate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDeactivateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDeactivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDeactivateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDeactivateBucketOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// ListBucketObjects : List bucket objects
// Fetch all objects from a given bucket.
func (watsonxData *WatsonxDataV2) ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions) (result *BucketRegistrationObjectCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListBucketObjectsWithContext(context.Background(), listBucketObjectsOptions)
}

// ListBucketObjectsWithContext is an alternate form of the ListBucketObjects method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListBucketObjectsWithContext(ctx context.Context, listBucketObjectsOptions *ListBucketObjectsOptions) (result *BucketRegistrationObjectCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBucketObjectsOptions, "listBucketObjectsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBucketObjectsOptions, "listBucketObjectsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *listBucketObjectsOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/objects`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBucketObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListBucketObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listBucketObjectsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listBucketObjectsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistrationObjectCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TestBucketConnection : Check bucket credentials to be valid
// Check whether provided bucket credentials are valid or not.
func (watsonxData *WatsonxDataV2) TestBucketConnection(testBucketConnectionOptions *TestBucketConnectionOptions) (result *TestBucketConnectionOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.TestBucketConnectionWithContext(context.Background(), testBucketConnectionOptions)
}

// TestBucketConnectionWithContext is an alternate form of the TestBucketConnection method which supports a Context parameter
func (watsonxData *WatsonxDataV2) TestBucketConnectionWithContext(ctx context.Context, testBucketConnectionOptions *TestBucketConnectionOptions) (result *TestBucketConnectionOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(testBucketConnectionOptions, "testBucketConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(testBucketConnectionOptions, "testBucketConnectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/test_bucket_connection`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range testBucketConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "TestBucketConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if testBucketConnectionOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*testBucketConnectionOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if testBucketConnectionOptions.AccessKey != nil {
		body["access_key"] = testBucketConnectionOptions.AccessKey
	}
	if testBucketConnectionOptions.BucketName != nil {
		body["bucket_name"] = testBucketConnectionOptions.BucketName
	}
	if testBucketConnectionOptions.BucketType != nil {
		body["bucket_type"] = testBucketConnectionOptions.BucketType
	}
	if testBucketConnectionOptions.Endpoint != nil {
		body["endpoint"] = testBucketConnectionOptions.Endpoint
	}
	if testBucketConnectionOptions.Region != nil {
		body["region"] = testBucketConnectionOptions.Region
	}
	if testBucketConnectionOptions.SecretKey != nil {
		body["secret_key"] = testBucketConnectionOptions.SecretKey
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestBucketConnectionOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDriverDatabaseCatalog : Add/Create database with driver
// Add or create a new database with driver.
func (watsonxData *WatsonxDataV2) CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptions *CreateDriverDatabaseCatalogOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDriverDatabaseCatalogWithContext(context.Background(), createDriverDatabaseCatalogOptions)
}

// CreateDriverDatabaseCatalogWithContext is an alternate form of the CreateDriverDatabaseCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateDriverDatabaseCatalogWithContext(ctx context.Context, createDriverDatabaseCatalogOptions *CreateDriverDatabaseCatalogOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDriverDatabaseCatalogOptions, "createDriverDatabaseCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDriverDatabaseCatalogOptions, "createDriverDatabaseCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_driver_registrations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDriverDatabaseCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateDriverDatabaseCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createDriverDatabaseCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDriverDatabaseCatalogOptions.AuthInstanceID))
	}

	builder.AddFormData("driver", "filename",
		core.StringNilMapper(createDriverDatabaseCatalogOptions.DriverContentType), createDriverDatabaseCatalogOptions.Driver)
	builder.AddFormData("driver_file_name", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.DriverFileName))
	builder.AddFormData("database_display_name", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.DatabaseDisplayName))
	builder.AddFormData("database_type", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.DatabaseType))
	builder.AddFormData("catalog_name", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.CatalogName))
	builder.AddFormData("hostname", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Hostname))
	builder.AddFormData("port", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Port))
	builder.AddFormData("username", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Username))
	builder.AddFormData("password", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Password))
	builder.AddFormData("database_name", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.DatabaseName))
	if createDriverDatabaseCatalogOptions.Certificate != nil {
		builder.AddFormData("certificate", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Certificate))
	}
	if createDriverDatabaseCatalogOptions.CertificateExtension != nil {
		builder.AddFormData("certificate_extension", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.CertificateExtension))
	}
	if createDriverDatabaseCatalogOptions.Ssl != nil {
		builder.AddFormData("ssl", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Ssl))
	}
	if createDriverDatabaseCatalogOptions.Description != nil {
		builder.AddFormData("description", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.Description))
	}
	if createDriverDatabaseCatalogOptions.CreatedOn != nil {
		builder.AddFormData("created_on", "", "", fmt.Sprint(*createDriverDatabaseCatalogOptions.CreatedOn))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDatabaseRegistrations : Get databases
// Get list of databases.
func (watsonxData *WatsonxDataV2) ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions) (result *DatabaseRegistrationCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListDatabaseRegistrationsWithContext(context.Background(), listDatabaseRegistrationsOptions)
}

// ListDatabaseRegistrationsWithContext is an alternate form of the ListDatabaseRegistrations method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListDatabaseRegistrationsWithContext(ctx context.Context, listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions) (result *DatabaseRegistrationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDatabaseRegistrationsOptions, "listDatabaseRegistrationsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDatabaseRegistrationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListDatabaseRegistrations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDatabaseRegistrationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listDatabaseRegistrationsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistrationCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDatabaseRegistration : Add/Create database
// Add or create a new database.
func (watsonxData *WatsonxDataV2) CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDatabaseRegistrationWithContext(context.Background(), createDatabaseRegistrationOptions)
}

// CreateDatabaseRegistrationWithContext is an alternate form of the CreateDatabaseRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateDatabaseRegistrationWithContext(ctx context.Context, createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDatabaseRegistrationOptions, "createDatabaseRegistrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDatabaseRegistrationOptions, "createDatabaseRegistrationOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDatabaseRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateDatabaseRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDatabaseRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDatabaseRegistrationOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDatabaseRegistrationOptions.DatabaseDisplayName != nil {
		body["database_display_name"] = createDatabaseRegistrationOptions.DatabaseDisplayName
	}
	if createDatabaseRegistrationOptions.DatabaseType != nil {
		body["database_type"] = createDatabaseRegistrationOptions.DatabaseType
	}
	if createDatabaseRegistrationOptions.AssociatedCatalog != nil {
		body["associated_catalog"] = createDatabaseRegistrationOptions.AssociatedCatalog
	}
	if createDatabaseRegistrationOptions.CreatedOn != nil {
		body["created_on"] = createDatabaseRegistrationOptions.CreatedOn
	}
	if createDatabaseRegistrationOptions.DatabaseDetails != nil {
		body["database_details"] = createDatabaseRegistrationOptions.DatabaseDetails
	}
	if createDatabaseRegistrationOptions.DatabaseProperties != nil {
		body["database_properties"] = createDatabaseRegistrationOptions.DatabaseProperties
	}
	if createDatabaseRegistrationOptions.Description != nil {
		body["description"] = createDatabaseRegistrationOptions.Description
	}
	if createDatabaseRegistrationOptions.Tags != nil {
		body["tags"] = createDatabaseRegistrationOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDatabase : Get database
// Get a registered databases.
func (watsonxData *WatsonxDataV2) GetDatabase(getDatabaseOptions *GetDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.GetDatabaseWithContext(context.Background(), getDatabaseOptions)
}

// GetDatabaseWithContext is an alternate form of the GetDatabase method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetDatabaseWithContext(ctx context.Context, getDatabaseOptions *GetDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDatabaseOptions, "getDatabaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDatabaseOptions, "getDatabaseOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *getDatabaseOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDatabaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetDatabase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDatabaseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDatabaseOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDatabaseCatalog : Delete database
// Delete a database.
func (watsonxData *WatsonxDataV2) DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDatabaseCatalogWithContext(context.Background(), deleteDatabaseCatalogOptions)
}

// DeleteDatabaseCatalogWithContext is an alternate form of the DeleteDatabaseCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDatabaseCatalogWithContext(ctx context.Context, deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *deleteDatabaseCatalogOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDatabaseCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDatabaseCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDatabaseCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDatabaseCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateDatabase : Update database
// Update database details.
func (watsonxData *WatsonxDataV2) UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateDatabaseWithContext(context.Background(), updateDatabaseOptions)
}

// UpdateDatabaseWithContext is an alternate form of the UpdateDatabase method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateDatabaseWithContext(ctx context.Context, updateDatabaseOptions *UpdateDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDatabaseOptions, "updateDatabaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDatabaseOptions, "updateDatabaseOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *updateDatabaseOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDatabaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateDatabase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateDatabaseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDatabaseOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateDatabaseOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ValidateDatabaseConnection : Validate database connection
// API to validate the database connection.
func (watsonxData *WatsonxDataV2) ValidateDatabaseConnection(validateDatabaseConnectionOptions *ValidateDatabaseConnectionOptions) (result *TestDatabaseConnectionResponse, response *core.DetailedResponse, err error) {
	return watsonxData.ValidateDatabaseConnectionWithContext(context.Background(), validateDatabaseConnectionOptions)
}

// ValidateDatabaseConnectionWithContext is an alternate form of the ValidateDatabaseConnection method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ValidateDatabaseConnectionWithContext(ctx context.Context, validateDatabaseConnectionOptions *ValidateDatabaseConnectionOptions) (result *TestDatabaseConnectionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(validateDatabaseConnectionOptions, "validateDatabaseConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(validateDatabaseConnectionOptions, "validateDatabaseConnectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/test_database_connection`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range validateDatabaseConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ValidateDatabaseConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if validateDatabaseConnectionOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*validateDatabaseConnectionOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if validateDatabaseConnectionOptions.DatabaseDetails != nil {
		body["database_details"] = validateDatabaseConnectionOptions.DatabaseDetails
	}
	if validateDatabaseConnectionOptions.DatabaseType != nil {
		body["database_type"] = validateDatabaseConnectionOptions.DatabaseType
	}
	if validateDatabaseConnectionOptions.Certificate != nil {
		body["certificate"] = validateDatabaseConnectionOptions.Certificate
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestDatabaseConnectionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDb2Engines : Get list of db2 engines
// Get list of all db2 engines.
func (watsonxData *WatsonxDataV2) ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions) (result *Db2EngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListDb2EnginesWithContext(context.Background(), listDb2EnginesOptions)
}

// ListDb2EnginesWithContext is an alternate form of the ListDb2Engines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListDb2EnginesWithContext(ctx context.Context, listDb2EnginesOptions *ListDb2EnginesOptions) (result *Db2EngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDb2EnginesOptions, "listDb2EnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDb2EnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListDb2Engines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDb2EnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listDb2EnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2EngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDb2Engine : Create db2 engine
// Create a new db2 engine.
func (watsonxData *WatsonxDataV2) CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDb2EngineWithContext(context.Background(), createDb2EngineOptions)
}

// CreateDb2EngineWithContext is an alternate form of the CreateDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateDb2EngineWithContext(ctx context.Context, createDb2EngineOptions *CreateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDb2EngineOptions, "createDb2EngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDb2EngineOptions, "createDb2EngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDb2EngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDb2EngineOptions.Origin != nil {
		body["origin"] = createDb2EngineOptions.Origin
	}
	if createDb2EngineOptions.Type != nil {
		body["type"] = createDb2EngineOptions.Type
	}
	if createDb2EngineOptions.Description != nil {
		body["description"] = createDb2EngineOptions.Description
	}
	if createDb2EngineOptions.EngineDetails != nil {
		body["engine_details"] = createDb2EngineOptions.EngineDetails
	}
	if createDb2EngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createDb2EngineOptions.EngineDisplayName
	}
	if createDb2EngineOptions.Tags != nil {
		body["tags"] = createDb2EngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2Engine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDb2Engine : Delete db2 engine
// Delete a db2 engine.
func (watsonxData *WatsonxDataV2) DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDb2EngineWithContext(context.Background(), deleteDb2EngineOptions)
}

// DeleteDb2EngineWithContext is an alternate form of the DeleteDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDb2EngineWithContext(ctx context.Context, deleteDb2EngineOptions *DeleteDb2EngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDb2EngineOptions, "deleteDb2EngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDb2EngineOptions, "deleteDb2EngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteDb2EngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDb2EngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateDb2Engine : Update db2 engine
// Update details of db2 engine.
func (watsonxData *WatsonxDataV2) UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateDb2EngineWithContext(context.Background(), updateDb2EngineOptions)
}

// UpdateDb2EngineWithContext is an alternate form of the UpdateDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateDb2EngineWithContext(ctx context.Context, updateDb2EngineOptions *UpdateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDb2EngineOptions, "updateDb2EngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDb2EngineOptions, "updateDb2EngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateDb2EngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDb2EngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateDb2EngineOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2Engine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEngines : Get all engines
// Get all engine details.
func (watsonxData *WatsonxDataV2) GetEngines(getEnginesOptions *GetEnginesOptions) (result *Engines, response *core.DetailedResponse, err error) {
	return watsonxData.GetEnginesWithContext(context.Background(), getEnginesOptions)
}

// GetEnginesWithContext is an alternate form of the GetEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetEnginesWithContext(ctx context.Context, getEnginesOptions *GetEnginesOptions) (result *Engines, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getEnginesOptions, "getEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEngines)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDeployments : Get deployments
// Get list of all deployments.
func (watsonxData *WatsonxDataV2) GetDeployments(getDeploymentsOptions *GetDeploymentsOptions) (result *GetDeploymentsOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetDeploymentsWithContext(context.Background(), getDeploymentsOptions)
}

// GetDeploymentsWithContext is an alternate form of the GetDeployments method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetDeploymentsWithContext(ctx context.Context, getDeploymentsOptions *GetDeploymentsOptions) (result *GetDeploymentsOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getDeploymentsOptions, "getDeploymentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/instance`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetDeployments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDeploymentsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDeploymentsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDeploymentsOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListNetezzaEngines : Get list of netezza engines
// Get list of all netezza engines.
func (watsonxData *WatsonxDataV2) ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions) (result *NetezzaEngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListNetezzaEnginesWithContext(context.Background(), listNetezzaEnginesOptions)
}

// ListNetezzaEnginesWithContext is an alternate form of the ListNetezzaEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListNetezzaEnginesWithContext(ctx context.Context, listNetezzaEnginesOptions *ListNetezzaEnginesOptions) (result *NetezzaEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNetezzaEnginesOptions, "listNetezzaEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listNetezzaEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListNetezzaEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listNetezzaEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listNetezzaEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateNetezzaEngine : Create netezza engine
// Create a new netezza engine.
func (watsonxData *WatsonxDataV2) CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	return watsonxData.CreateNetezzaEngineWithContext(context.Background(), createNetezzaEngineOptions)
}

// CreateNetezzaEngineWithContext is an alternate form of the CreateNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateNetezzaEngineWithContext(ctx context.Context, createNetezzaEngineOptions *CreateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNetezzaEngineOptions, "createNetezzaEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createNetezzaEngineOptions, "createNetezzaEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createNetezzaEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createNetezzaEngineOptions.Origin != nil {
		body["origin"] = createNetezzaEngineOptions.Origin
	}
	if createNetezzaEngineOptions.Type != nil {
		body["type"] = createNetezzaEngineOptions.Type
	}
	if createNetezzaEngineOptions.Description != nil {
		body["description"] = createNetezzaEngineOptions.Description
	}
	if createNetezzaEngineOptions.EngineDetails != nil {
		body["engine_details"] = createNetezzaEngineOptions.EngineDetails
	}
	if createNetezzaEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createNetezzaEngineOptions.EngineDisplayName
	}
	if createNetezzaEngineOptions.Tags != nil {
		body["tags"] = createNetezzaEngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteNetezzaEngine : Delete netezza engine
// Delete a netezza engine.
func (watsonxData *WatsonxDataV2) DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteNetezzaEngineWithContext(context.Background(), deleteNetezzaEngineOptions)
}

// DeleteNetezzaEngineWithContext is an alternate form of the DeleteNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteNetezzaEngineWithContext(ctx context.Context, deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNetezzaEngineOptions, "deleteNetezzaEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNetezzaEngineOptions, "deleteNetezzaEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteNetezzaEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteNetezzaEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateNetezzaEngine : Update netezza engine
// Update details of netezza engine.
func (watsonxData *WatsonxDataV2) UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateNetezzaEngineWithContext(context.Background(), updateNetezzaEngineOptions)
}

// UpdateNetezzaEngineWithContext is an alternate form of the UpdateNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateNetezzaEngineWithContext(ctx context.Context, updateNetezzaEngineOptions *UpdateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNetezzaEngineOptions, "updateNetezzaEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateNetezzaEngineOptions, "updateNetezzaEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateNetezzaEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateNetezzaEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateNetezzaEngineOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListOtherEngines : List other engines
// list all other engine details.
func (watsonxData *WatsonxDataV2) ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions) (result *OtherEngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListOtherEnginesWithContext(context.Background(), listOtherEnginesOptions)
}

// ListOtherEnginesWithContext is an alternate form of the ListOtherEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListOtherEnginesWithContext(ctx context.Context, listOtherEnginesOptions *ListOtherEnginesOptions) (result *OtherEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listOtherEnginesOptions, "listOtherEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listOtherEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListOtherEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listOtherEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listOtherEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOtherEngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateOtherEngine : Create other engine
// Create a new engine.
func (watsonxData *WatsonxDataV2) CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions) (result *OtherEngine, response *core.DetailedResponse, err error) {
	return watsonxData.CreateOtherEngineWithContext(context.Background(), createOtherEngineOptions)
}

// CreateOtherEngineWithContext is an alternate form of the CreateOtherEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateOtherEngineWithContext(ctx context.Context, createOtherEngineOptions *CreateOtherEngineOptions) (result *OtherEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOtherEngineOptions, "createOtherEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createOtherEngineOptions, "createOtherEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createOtherEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateOtherEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createOtherEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createOtherEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createOtherEngineOptions.EngineDetails != nil {
		body["engine_details"] = createOtherEngineOptions.EngineDetails
	}
	if createOtherEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createOtherEngineOptions.EngineDisplayName
	}
	if createOtherEngineOptions.Description != nil {
		body["description"] = createOtherEngineOptions.Description
	}
	if createOtherEngineOptions.Origin != nil {
		body["origin"] = createOtherEngineOptions.Origin
	}
	if createOtherEngineOptions.Tags != nil {
		body["tags"] = createOtherEngineOptions.Tags
	}
	if createOtherEngineOptions.Type != nil {
		body["type"] = createOtherEngineOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOtherEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteOtherEngine : Delete engine
// Delete an engine from lakehouse.
func (watsonxData *WatsonxDataV2) DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteOtherEngineWithContext(context.Background(), deleteOtherEngineOptions)
}

// DeleteOtherEngineWithContext is an alternate form of the DeleteOtherEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteOtherEngineWithContext(ctx context.Context, deleteOtherEngineOptions *DeleteOtherEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteOtherEngineOptions, "deleteOtherEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteOtherEngineOptions, "deleteOtherEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteOtherEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteOtherEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteOtherEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteOtherEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteOtherEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// ListPrestissimoEngines : Get list of prestissimo engines
// Get list of all prestissimo engines.
func (watsonxData *WatsonxDataV2) ListPrestissimoEngines(listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions) (result *PrestissimoEngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListPrestissimoEnginesWithContext(context.Background(), listPrestissimoEnginesOptions)
}

// ListPrestissimoEnginesWithContext is an alternate form of the ListPrestissimoEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestissimoEnginesWithContext(ctx context.Context, listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions) (result *PrestissimoEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPrestissimoEnginesOptions, "listPrestissimoEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPrestissimoEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestissimoEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestissimoEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestissimoEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEngine : Create prestissimo engine
// Create a new prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngine(createPrestissimoEngineOptions *CreatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestissimoEngineWithContext(context.Background(), createPrestissimoEngineOptions)
}

// CreatePrestissimoEngineWithContext is an alternate form of the CreatePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineWithContext(ctx context.Context, createPrestissimoEngineOptions *CreatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEngineOptions, "createPrestissimoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestissimoEngineOptions, "createPrestissimoEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createPrestissimoEngineOptions.Origin != nil {
		body["origin"] = createPrestissimoEngineOptions.Origin
	}
	if createPrestissimoEngineOptions.Type != nil {
		body["type"] = createPrestissimoEngineOptions.Type
	}
	if createPrestissimoEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createPrestissimoEngineOptions.AssociatedCatalogs
	}
	if createPrestissimoEngineOptions.Description != nil {
		body["description"] = createPrestissimoEngineOptions.Description
	}
	if createPrestissimoEngineOptions.EngineDetails != nil {
		body["engine_details"] = createPrestissimoEngineOptions.EngineDetails
	}
	if createPrestissimoEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createPrestissimoEngineOptions.EngineDisplayName
	}
	if createPrestissimoEngineOptions.Region != nil {
		body["region"] = createPrestissimoEngineOptions.Region
	}
	if createPrestissimoEngineOptions.Tags != nil {
		body["tags"] = createPrestissimoEngineOptions.Tags
	}
	if createPrestissimoEngineOptions.Version != nil {
		body["version"] = createPrestissimoEngineOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPrestissimoEngine : Get prestissimo engine
// Get details of one prestissimo engine.
func (watsonxData *WatsonxDataV2) GetPrestissimoEngine(getPrestissimoEngineOptions *GetPrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.GetPrestissimoEngineWithContext(context.Background(), getPrestissimoEngineOptions)
}

// GetPrestissimoEngineWithContext is an alternate form of the GetPrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineWithContext(ctx context.Context, getPrestissimoEngineOptions *GetPrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestissimoEngineOptions, "getPrestissimoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPrestissimoEngineOptions, "getPrestissimoEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestissimoEngine : Delete prestissimo engine
// Delete a prestissimo engine.
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngine(deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeletePrestissimoEngineWithContext(context.Background(), deletePrestissimoEngineOptions)
}

// DeletePrestissimoEngineWithContext is an alternate form of the DeletePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineWithContext(ctx context.Context, deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestissimoEngineOptions, "deletePrestissimoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePrestissimoEngineOptions, "deletePrestissimoEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdatePrestissimoEngine : Update prestissimo engine
// Update details of prestissimo engine.
func (watsonxData *WatsonxDataV2) UpdatePrestissimoEngine(updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.UpdatePrestissimoEngineWithContext(context.Background(), updatePrestissimoEngineOptions)
}

// UpdatePrestissimoEngineWithContext is an alternate form of the UpdatePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdatePrestissimoEngineWithContext(ctx context.Context, updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePrestissimoEngineOptions, "updatePrestissimoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePrestissimoEngineOptions, "updatePrestissimoEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updatePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdatePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updatePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updatePrestissimoEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updatePrestissimoEngineOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPrestissimoEngineCatalogs : Get prestissimo engine catalogs
// Get list of all catalogs attached a prestissimo engine.
func (watsonxData *WatsonxDataV2) ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListPrestissimoEngineCatalogsWithContext(context.Background(), listPrestissimoEngineCatalogsOptions)
}

// ListPrestissimoEngineCatalogsWithContext is an alternate form of the ListPrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestissimoEngineCatalogsWithContext(ctx context.Context, listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPrestissimoEngineCatalogsOptions, "listPrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPrestissimoEngineCatalogsOptions, "listPrestissimoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listPrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplacePrestissimoEngineCatalogs : Associate catalogs to a prestissimo engine
// Associate one or more catalogs to a prestissimo engine.
func (watsonxData *WatsonxDataV2) ReplacePrestissimoEngineCatalogs(replacePrestissimoEngineCatalogsOptions *ReplacePrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ReplacePrestissimoEngineCatalogsWithContext(context.Background(), replacePrestissimoEngineCatalogsOptions)
}

// ReplacePrestissimoEngineCatalogsWithContext is an alternate form of the ReplacePrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ReplacePrestissimoEngineCatalogsWithContext(ctx context.Context, replacePrestissimoEngineCatalogsOptions *ReplacePrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replacePrestissimoEngineCatalogsOptions, "replacePrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replacePrestissimoEngineCatalogsOptions, "replacePrestissimoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *replacePrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replacePrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ReplacePrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if replacePrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*replacePrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*replacePrestissimoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestissimoEngineCatalogs : Disassociate catalogs from a prestissimo engine
// Disassociate one or more catalogs from a prestissimo engine.
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeletePrestissimoEngineCatalogsWithContext(context.Background(), deletePrestissimoEngineCatalogsOptions)
}

// DeletePrestissimoEngineCatalogsWithContext is an alternate form of the DeletePrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineCatalogsWithContext(ctx context.Context, deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestissimoEngineCatalogsOptions, "deletePrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePrestissimoEngineCatalogsOptions, "deletePrestissimoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*deletePrestissimoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// GetPrestissimoEngineCatalog : Get prestissimo engine catalog
// Get catalog attached to a prestissimo engine.
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	return watsonxData.GetPrestissimoEngineCatalogWithContext(context.Background(), getPrestissimoEngineCatalogOptions)
}

// GetPrestissimoEngineCatalogWithContext is an alternate form of the GetPrestissimoEngineCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineCatalogWithContext(ctx context.Context, getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestissimoEngineCatalogOptions, "getPrestissimoEngineCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPrestissimoEngineCatalogOptions, "getPrestissimoEngineCatalogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestissimoEngineCatalogOptions.EngineID,
		"catalog_id": *getPrestissimoEngineCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPrestissimoEngineCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestissimoEngineCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestissimoEngineCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestissimoEngineCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEnginePause : Pause prestissimo engine
// Pause a running prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEnginePause(createPrestissimoEnginePauseOptions *CreatePrestissimoEnginePauseOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestissimoEnginePauseWithContext(context.Background(), createPrestissimoEnginePauseOptions)
}

// CreatePrestissimoEnginePauseWithContext is an alternate form of the CreatePrestissimoEnginePause method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEnginePauseWithContext(ctx context.Context, createPrestissimoEnginePauseOptions *CreatePrestissimoEnginePauseOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEnginePauseOptions, "createPrestissimoEnginePauseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestissimoEnginePauseOptions, "createPrestissimoEnginePauseOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createPrestissimoEnginePauseOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/pause`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestissimoEnginePauseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEnginePause")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createPrestissimoEnginePauseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEnginePauseOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunPrestissimoExplainStatement : Explain query
// Explain a query statement.
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions) (result *ResultPrestissimoExplainStatement, response *core.DetailedResponse, err error) {
	return watsonxData.RunPrestissimoExplainStatementWithContext(context.Background(), runPrestissimoExplainStatementOptions)
}

// RunPrestissimoExplainStatementWithContext is an alternate form of the RunPrestissimoExplainStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainStatementWithContext(ctx context.Context, runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions) (result *ResultPrestissimoExplainStatement, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runPrestissimoExplainStatementOptions, "runPrestissimoExplainStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runPrestissimoExplainStatementOptions, "runPrestissimoExplainStatementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runPrestissimoExplainStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/query_explain`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runPrestissimoExplainStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunPrestissimoExplainStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runPrestissimoExplainStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runPrestissimoExplainStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runPrestissimoExplainStatementOptions.Statement != nil {
		body["statement"] = runPrestissimoExplainStatementOptions.Statement
	}
	if runPrestissimoExplainStatementOptions.Format != nil {
		body["format"] = runPrestissimoExplainStatementOptions.Format
	}
	if runPrestissimoExplainStatementOptions.Type != nil {
		body["type"] = runPrestissimoExplainStatementOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResultPrestissimoExplainStatement)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunPrestissimoExplainAnalyzeStatement : Explain analyze
// Return query metrics after query is complete.
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions) (result *ResultRunPrestissimoExplainAnalyzeStatement, response *core.DetailedResponse, err error) {
	return watsonxData.RunPrestissimoExplainAnalyzeStatementWithContext(context.Background(), runPrestissimoExplainAnalyzeStatementOptions)
}

// RunPrestissimoExplainAnalyzeStatementWithContext is an alternate form of the RunPrestissimoExplainAnalyzeStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainAnalyzeStatementWithContext(ctx context.Context, runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions) (result *ResultRunPrestissimoExplainAnalyzeStatement, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runPrestissimoExplainAnalyzeStatementOptions, "runPrestissimoExplainAnalyzeStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runPrestissimoExplainAnalyzeStatementOptions, "runPrestissimoExplainAnalyzeStatementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runPrestissimoExplainAnalyzeStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/query_explain_analyze`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runPrestissimoExplainAnalyzeStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunPrestissimoExplainAnalyzeStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runPrestissimoExplainAnalyzeStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runPrestissimoExplainAnalyzeStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runPrestissimoExplainAnalyzeStatementOptions.Statement != nil {
		body["statement"] = runPrestissimoExplainAnalyzeStatementOptions.Statement
	}
	if runPrestissimoExplainAnalyzeStatementOptions.Verbose != nil {
		body["verbose"] = runPrestissimoExplainAnalyzeStatementOptions.Verbose
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResultRunPrestissimoExplainAnalyzeStatement)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEngineRestart : Restart a prestissimo engine
// Restart an existing prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineRestart(createPrestissimoEngineRestartOptions *CreatePrestissimoEngineRestartOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestissimoEngineRestartWithContext(context.Background(), createPrestissimoEngineRestartOptions)
}

// CreatePrestissimoEngineRestartWithContext is an alternate form of the CreatePrestissimoEngineRestart method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineRestartWithContext(ctx context.Context, createPrestissimoEngineRestartOptions *CreatePrestissimoEngineRestartOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEngineRestartOptions, "createPrestissimoEngineRestartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestissimoEngineRestartOptions, "createPrestissimoEngineRestartOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createPrestissimoEngineRestartOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/restart`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestissimoEngineRestartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEngineRestart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createPrestissimoEngineRestartOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEngineRestartOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEngineResume : Resume prestissimo engine
// Resume a paused prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineResume(createPrestissimoEngineResumeOptions *CreatePrestissimoEngineResumeOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestissimoEngineResumeWithContext(context.Background(), createPrestissimoEngineResumeOptions)
}

// CreatePrestissimoEngineResumeWithContext is an alternate form of the CreatePrestissimoEngineResume method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineResumeWithContext(ctx context.Context, createPrestissimoEngineResumeOptions *CreatePrestissimoEngineResumeOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEngineResumeOptions, "createPrestissimoEngineResumeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestissimoEngineResumeOptions, "createPrestissimoEngineResumeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createPrestissimoEngineResumeOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/resume`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestissimoEngineResumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEngineResume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createPrestissimoEngineResumeOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEngineResumeOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEngineScale : Scale a prestissimo engine
// Scale an existing prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineScale(createPrestissimoEngineScaleOptions *CreatePrestissimoEngineScaleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestissimoEngineScaleWithContext(context.Background(), createPrestissimoEngineScaleOptions)
}

// CreatePrestissimoEngineScaleWithContext is an alternate form of the CreatePrestissimoEngineScale method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineScaleWithContext(ctx context.Context, createPrestissimoEngineScaleOptions *CreatePrestissimoEngineScaleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEngineScaleOptions, "createPrestissimoEngineScaleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestissimoEngineScaleOptions, "createPrestissimoEngineScaleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createPrestissimoEngineScaleOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/scale`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestissimoEngineScaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEngineScale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPrestissimoEngineScaleOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEngineScaleOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createPrestissimoEngineScaleOptions.Coordinator != nil {
		body["coordinator"] = createPrestissimoEngineScaleOptions.Coordinator
	}
	if createPrestissimoEngineScaleOptions.Worker != nil {
		body["worker"] = createPrestissimoEngineScaleOptions.Worker
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPrestoEngines : Get list of presto engines
// Get list of all presto engines.
func (watsonxData *WatsonxDataV2) ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions) (result *PrestoEngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListPrestoEnginesWithContext(context.Background(), listPrestoEnginesOptions)
}

// ListPrestoEnginesWithContext is an alternate form of the ListPrestoEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestoEnginesWithContext(ctx context.Context, listPrestoEnginesOptions *ListPrestoEnginesOptions) (result *PrestoEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPrestoEnginesOptions, "listPrestoEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPrestoEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestoEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestoEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestoEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestoEngine : Create presto engine
// Create a new presto engine.
func (watsonxData *WatsonxDataV2) CreatePrestoEngine(createPrestoEngineOptions *CreatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.CreatePrestoEngineWithContext(context.Background(), createPrestoEngineOptions)
}

// CreatePrestoEngineWithContext is an alternate form of the CreatePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestoEngineWithContext(ctx context.Context, createPrestoEngineOptions *CreatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestoEngineOptions, "createPrestoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPrestoEngineOptions, "createPrestoEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createPrestoEngineOptions.Origin != nil {
		body["origin"] = createPrestoEngineOptions.Origin
	}
	if createPrestoEngineOptions.Type != nil {
		body["type"] = createPrestoEngineOptions.Type
	}
	if createPrestoEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createPrestoEngineOptions.AssociatedCatalogs
	}
	if createPrestoEngineOptions.Description != nil {
		body["description"] = createPrestoEngineOptions.Description
	}
	if createPrestoEngineOptions.EngineDetails != nil {
		body["engine_details"] = createPrestoEngineOptions.EngineDetails
	}
	if createPrestoEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createPrestoEngineOptions.EngineDisplayName
	}
	if createPrestoEngineOptions.Region != nil {
		body["region"] = createPrestoEngineOptions.Region
	}
	if createPrestoEngineOptions.Tags != nil {
		body["tags"] = createPrestoEngineOptions.Tags
	}
	if createPrestoEngineOptions.Version != nil {
		body["version"] = createPrestoEngineOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPrestoEngine : Get presto engine
// Get details of one presto engine.
func (watsonxData *WatsonxDataV2) GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.GetPrestoEngineWithContext(context.Background(), getPrestoEngineOptions)
}

// GetPrestoEngineWithContext is an alternate form of the GetPrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestoEngineWithContext(ctx context.Context, getPrestoEngineOptions *GetPrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestoEngineOptions, "getPrestoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPrestoEngineOptions, "getPrestoEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteEngine : Delete presto engine
// Delete a presto engine.
func (watsonxData *WatsonxDataV2) DeleteEngine(deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteEngineWithContext(context.Background(), deleteEngineOptions)
}

// DeleteEngineWithContext is an alternate form of the DeleteEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteEngineWithContext(ctx context.Context, deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEngineOptions, "deleteEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEngineOptions, "deleteEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdatePrestoEngine : Update presto engine
// Update details of presto engine.
func (watsonxData *WatsonxDataV2) UpdatePrestoEngine(updatePrestoEngineOptions *UpdatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	return watsonxData.UpdatePrestoEngineWithContext(context.Background(), updatePrestoEngineOptions)
}

// UpdatePrestoEngineWithContext is an alternate form of the UpdatePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdatePrestoEngineWithContext(ctx context.Context, updatePrestoEngineOptions *UpdatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePrestoEngineOptions, "updatePrestoEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePrestoEngineOptions, "updatePrestoEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updatePrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdatePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updatePrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updatePrestoEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updatePrestoEngineOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPrestoEngineCatalogs : Get presto engine catalogs
// Get list of all catalogs attached to a presto engine.
func (watsonxData *WatsonxDataV2) ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListPrestoEngineCatalogsWithContext(context.Background(), listPrestoEngineCatalogsOptions)
}

// ListPrestoEngineCatalogsWithContext is an alternate form of the ListPrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestoEngineCatalogsWithContext(ctx context.Context, listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPrestoEngineCatalogsOptions, "listPrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPrestoEngineCatalogsOptions, "listPrestoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listPrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestoEngineCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplacePrestoEngineCatalogs : Associate catalogs to presto engine
// Associate one or more catalogs to a presto engine.
func (watsonxData *WatsonxDataV2) ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptions *ReplacePrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ReplacePrestoEngineCatalogsWithContext(context.Background(), replacePrestoEngineCatalogsOptions)
}

// ReplacePrestoEngineCatalogsWithContext is an alternate form of the ReplacePrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ReplacePrestoEngineCatalogsWithContext(ctx context.Context, replacePrestoEngineCatalogsOptions *ReplacePrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replacePrestoEngineCatalogsOptions, "replacePrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replacePrestoEngineCatalogsOptions, "replacePrestoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *replacePrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replacePrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ReplacePrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if replacePrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*replacePrestoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*replacePrestoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestoEngineCatalogs : Disassociate catalogs from a presto engine
// Disassociate one or more catalogs from a presto engine.
func (watsonxData *WatsonxDataV2) DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeletePrestoEngineCatalogsWithContext(context.Background(), deletePrestoEngineCatalogsOptions)
}

// DeletePrestoEngineCatalogsWithContext is an alternate form of the DeletePrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestoEngineCatalogsWithContext(ctx context.Context, deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestoEngineCatalogsOptions, "deletePrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePrestoEngineCatalogsOptions, "deletePrestoEngineCatalogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*deletePrestoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// GetPrestoEngineCatalog : Get presto engine catalog
// Get catalog attached to presto engine.
func (watsonxData *WatsonxDataV2) GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	return watsonxData.GetPrestoEngineCatalogWithContext(context.Background(), getPrestoEngineCatalogOptions)
}

// GetPrestoEngineCatalogWithContext is an alternate form of the GetPrestoEngineCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestoEngineCatalogWithContext(ctx context.Context, getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestoEngineCatalogOptions, "getPrestoEngineCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPrestoEngineCatalogOptions, "getPrestoEngineCatalogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestoEngineCatalogOptions.EngineID,
		"catalog_id": *getPrestoEngineCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPrestoEngineCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestoEngineCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestoEngineCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestoEngineCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateEnginePause : Pause presto engine
// Pause a running presto engine.
func (watsonxData *WatsonxDataV2) CreateEnginePause(createEnginePauseOptions *CreateEnginePauseOptions) (result *CreateEnginePauseCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEnginePauseWithContext(context.Background(), createEnginePauseOptions)
}

// CreateEnginePauseWithContext is an alternate form of the CreateEnginePause method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateEnginePauseWithContext(ctx context.Context, createEnginePauseOptions *CreateEnginePauseOptions) (result *CreateEnginePauseCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEnginePauseOptions, "createEnginePauseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEnginePauseOptions, "createEnginePauseOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createEnginePauseOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/pause`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEnginePauseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateEnginePause")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createEnginePauseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEnginePauseOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEnginePauseCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunExplainStatement : Explain presto query
// Explain a query statement.
func (watsonxData *WatsonxDataV2) RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions) (result *RunExplainStatementOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.RunExplainStatementWithContext(context.Background(), runExplainStatementOptions)
}

// RunExplainStatementWithContext is an alternate form of the RunExplainStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunExplainStatementWithContext(ctx context.Context, runExplainStatementOptions *RunExplainStatementOptions) (result *RunExplainStatementOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runExplainStatementOptions, "runExplainStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runExplainStatementOptions, "runExplainStatementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runExplainStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/query_explain`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runExplainStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunExplainStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runExplainStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runExplainStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runExplainStatementOptions.Statement != nil {
		body["statement"] = runExplainStatementOptions.Statement
	}
	if runExplainStatementOptions.Format != nil {
		body["format"] = runExplainStatementOptions.Format
	}
	if runExplainStatementOptions.Type != nil {
		body["type"] = runExplainStatementOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRunExplainStatementOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunExplainAnalyzeStatement : Explain presto analyze
// Return query metrics after query is complete.
func (watsonxData *WatsonxDataV2) RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions) (result *RunExplainAnalyzeStatementOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.RunExplainAnalyzeStatementWithContext(context.Background(), runExplainAnalyzeStatementOptions)
}

// RunExplainAnalyzeStatementWithContext is an alternate form of the RunExplainAnalyzeStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunExplainAnalyzeStatementWithContext(ctx context.Context, runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions) (result *RunExplainAnalyzeStatementOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runExplainAnalyzeStatementOptions, "runExplainAnalyzeStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runExplainAnalyzeStatementOptions, "runExplainAnalyzeStatementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runExplainAnalyzeStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/query_explain_analyze`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runExplainAnalyzeStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunExplainAnalyzeStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runExplainAnalyzeStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runExplainAnalyzeStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runExplainAnalyzeStatementOptions.Statement != nil {
		body["statement"] = runExplainAnalyzeStatementOptions.Statement
	}
	if runExplainAnalyzeStatementOptions.Verbose != nil {
		body["verbose"] = runExplainAnalyzeStatementOptions.Verbose
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRunExplainAnalyzeStatementOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateEngineRestart : Restart a presto engine
// Restart an existing presto engine.
func (watsonxData *WatsonxDataV2) CreateEngineRestart(createEngineRestartOptions *CreateEngineRestartOptions) (result *CreateEngineRestartCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEngineRestartWithContext(context.Background(), createEngineRestartOptions)
}

// CreateEngineRestartWithContext is an alternate form of the CreateEngineRestart method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateEngineRestartWithContext(ctx context.Context, createEngineRestartOptions *CreateEngineRestartOptions) (result *CreateEngineRestartCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEngineRestartOptions, "createEngineRestartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEngineRestartOptions, "createEngineRestartOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createEngineRestartOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/restart`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEngineRestartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateEngineRestart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createEngineRestartOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEngineRestartOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineRestartCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateEngineResume : Resume presto engine
// Resume a paused presto engine.
func (watsonxData *WatsonxDataV2) CreateEngineResume(createEngineResumeOptions *CreateEngineResumeOptions) (result *CreateEngineResumeCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEngineResumeWithContext(context.Background(), createEngineResumeOptions)
}

// CreateEngineResumeWithContext is an alternate form of the CreateEngineResume method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateEngineResumeWithContext(ctx context.Context, createEngineResumeOptions *CreateEngineResumeOptions) (result *CreateEngineResumeCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEngineResumeOptions, "createEngineResumeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEngineResumeOptions, "createEngineResumeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createEngineResumeOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/resume`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEngineResumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateEngineResume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createEngineResumeOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEngineResumeOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineResumeCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateEngineScale : Scale a presto engine
// Scale an existing presto engine.
func (watsonxData *WatsonxDataV2) CreateEngineScale(createEngineScaleOptions *CreateEngineScaleOptions) (result *CreateEngineScaleCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEngineScaleWithContext(context.Background(), createEngineScaleOptions)
}

// CreateEngineScaleWithContext is an alternate form of the CreateEngineScale method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateEngineScaleWithContext(ctx context.Context, createEngineScaleOptions *CreateEngineScaleOptions) (result *CreateEngineScaleCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEngineScaleOptions, "createEngineScaleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEngineScaleOptions, "createEngineScaleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createEngineScaleOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/scale`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEngineScaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateEngineScale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createEngineScaleOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEngineScaleOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createEngineScaleOptions.Coordinator != nil {
		body["coordinator"] = createEngineScaleOptions.Coordinator
	}
	if createEngineScaleOptions.Worker != nil {
		body["worker"] = createEngineScaleOptions.Worker
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineScaleCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListSparkEngines : List all spark engines
// List all spark engines.
func (watsonxData *WatsonxDataV2) ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions) (result *SparkEngineCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListSparkEnginesWithContext(context.Background(), listSparkEnginesOptions)
}

// ListSparkEnginesWithContext is an alternate form of the ListSparkEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkEnginesWithContext(ctx context.Context, listSparkEnginesOptions *ListSparkEnginesOptions) (result *SparkEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSparkEnginesOptions, "listSparkEnginesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSparkEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngine : Create spark engine
// Create a new spark  engine.
func (watsonxData *WatsonxDataV2) CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	return watsonxData.CreateSparkEngineWithContext(context.Background(), createSparkEngineOptions)
}

// CreateSparkEngineWithContext is an alternate form of the CreateSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineWithContext(ctx context.Context, createSparkEngineOptions *CreateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineOptions, "createSparkEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSparkEngineOptions, "createSparkEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createSparkEngineOptions.Origin != nil {
		body["origin"] = createSparkEngineOptions.Origin
	}
	if createSparkEngineOptions.Type != nil {
		body["type"] = createSparkEngineOptions.Type
	}
	if createSparkEngineOptions.Description != nil {
		body["description"] = createSparkEngineOptions.Description
	}
	if createSparkEngineOptions.EngineDetails != nil {
		body["engine_details"] = createSparkEngineOptions.EngineDetails
	}
	if createSparkEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createSparkEngineOptions.EngineDisplayName
	}
	if createSparkEngineOptions.Tags != nil {
		body["tags"] = createSparkEngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngine : Delete spark engine
// Delete a spark engine.
func (watsonxData *WatsonxDataV2) DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteSparkEngineWithContext(context.Background(), deleteSparkEngineOptions)
}

// DeleteSparkEngineWithContext is an alternate form of the DeleteSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineWithContext(ctx context.Context, deleteSparkEngineOptions *DeleteSparkEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineOptions, "deleteSparkEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSparkEngineOptions, "deleteSparkEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateSparkEngine : Update spark engine
// Update details of spark engine.
func (watsonxData *WatsonxDataV2) UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateSparkEngineWithContext(context.Background(), updateSparkEngineOptions)
}

// UpdateSparkEngineWithContext is an alternate form of the UpdateSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateSparkEngineWithContext(ctx context.Context, updateSparkEngineOptions *UpdateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSparkEngineOptions, "updateSparkEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSparkEngineOptions, "updateSparkEngineOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateSparkEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateSparkEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateSparkEngineOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngine)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListSparkEngineApplications : List all applications in a spark engine
// List all applications in a spark engine.
func (watsonxData *WatsonxDataV2) ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) (result *SparkEngineApplicationStatusCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListSparkEngineApplicationsWithContext(context.Background(), listSparkEngineApplicationsOptions)
}

// ListSparkEngineApplicationsWithContext is an alternate form of the ListSparkEngineApplications method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkEngineApplicationsWithContext(ctx context.Context, listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) (result *SparkEngineApplicationStatusCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSparkEngineApplicationsOptions, "listSparkEngineApplicationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listSparkEngineApplicationsOptions, "listSparkEngineApplicationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listSparkEngineApplicationsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSparkEngineApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkEngineApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkEngineApplicationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkEngineApplicationsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatusCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngineApplication : Submit engine applications
// Submit engine applications.
func (watsonxData *WatsonxDataV2) CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	return watsonxData.CreateSparkEngineApplicationWithContext(context.Background(), createSparkEngineApplicationOptions)
}

// CreateSparkEngineApplicationWithContext is an alternate form of the CreateSparkEngineApplication method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineApplicationWithContext(ctx context.Context, createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineApplicationOptions, "createSparkEngineApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSparkEngineApplicationOptions, "createSparkEngineApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createSparkEngineApplicationOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSparkEngineApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngineApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSparkEngineApplicationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineApplicationOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createSparkEngineApplicationOptions.ApplicationDetails != nil {
		body["application_details"] = createSparkEngineApplicationOptions.ApplicationDetails
	}
	if createSparkEngineApplicationOptions.JobEndpoint != nil {
		body["job_endpoint"] = createSparkEngineApplicationOptions.JobEndpoint
	}
	if createSparkEngineApplicationOptions.ServiceInstanceID != nil {
		body["service_instance_id"] = createSparkEngineApplicationOptions.ServiceInstanceID
	}
	if createSparkEngineApplicationOptions.Type != nil {
		body["type"] = createSparkEngineApplicationOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngineApplications : Stop Spark Applications
// Stop a running spark application.
func (watsonxData *WatsonxDataV2) DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteSparkEngineApplicationsWithContext(context.Background(), deleteSparkEngineApplicationsOptions)
}

// DeleteSparkEngineApplicationsWithContext is an alternate form of the DeleteSparkEngineApplications method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineApplicationsWithContext(ctx context.Context, deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineApplicationsOptions, "deleteSparkEngineApplicationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSparkEngineApplicationsOptions, "deleteSparkEngineApplicationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineApplicationsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSparkEngineApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngineApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineApplicationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineApplicationsOptions.AuthInstanceID))
	}

	builder.AddQuery("application_id", fmt.Sprint(*deleteSparkEngineApplicationsOptions.ApplicationID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// GetSparkEngineApplicationStatus : Get spark application
// Get status of spark application.
func (watsonxData *WatsonxDataV2) GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	return watsonxData.GetSparkEngineApplicationStatusWithContext(context.Background(), getSparkEngineApplicationStatusOptions)
}

// GetSparkEngineApplicationStatusWithContext is an alternate form of the GetSparkEngineApplicationStatus method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetSparkEngineApplicationStatusWithContext(ctx context.Context, getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkEngineApplicationStatusOptions, "getSparkEngineApplicationStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSparkEngineApplicationStatusOptions, "getSparkEngineApplicationStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getSparkEngineApplicationStatusOptions.EngineID,
		"application_id": *getSparkEngineApplicationStatusOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications/{application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSparkEngineApplicationStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetSparkEngineApplicationStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSparkEngineApplicationStatusOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSparkEngineApplicationStatusOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TestLHConsole : Readiness API
// Verify lhconsole server is up and running.
func (watsonxData *WatsonxDataV2) TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.TestLHConsoleWithContext(context.Background(), testLHConsoleOptions)
}

// TestLHConsoleWithContext is an alternate form of the TestLHConsole method which supports a Context parameter
func (watsonxData *WatsonxDataV2) TestLHConsoleWithContext(ctx context.Context, testLHConsoleOptions *TestLHConsoleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(testLHConsoleOptions, "testLHConsoleOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/ready`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range testLHConsoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "TestLHConsole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCatalogs : List all registered catalogs
// List all registered catalogs.
func (watsonxData *WatsonxDataV2) ListCatalogs(listCatalogsOptions *ListCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListCatalogsWithContext(context.Background(), listCatalogsOptions)
}

// ListCatalogsWithContext is an alternate form of the ListCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListCatalogsWithContext(ctx context.Context, listCatalogsOptions *ListCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogsOptions, "listCatalogsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCatalog : Get catalog properties by catalog_id
// Get catalog properties of a catalog identified by catalog_id.
func (watsonxData *WatsonxDataV2) GetCatalog(getCatalogOptions *GetCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	return watsonxData.GetCatalogWithContext(context.Background(), getCatalogOptions)
}

// GetCatalogWithContext is an alternate form of the GetCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetCatalogWithContext(ctx context.Context, getCatalogOptions *GetCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCatalogOptions, "getCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCatalogOptions, "getCatalogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *getCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListSchemas : List all schemas
// List all schemas in catalog.
func (watsonxData *WatsonxDataV2) ListSchemas(listSchemasOptions *ListSchemasOptions) (result *ListSchemasOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.ListSchemasWithContext(context.Background(), listSchemasOptions)
}

// ListSchemasWithContext is an alternate form of the ListSchemas method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSchemasWithContext(ctx context.Context, listSchemasOptions *ListSchemasOptions) (result *ListSchemasOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSchemasOptions, "listSchemasOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listSchemasOptions, "listSchemasOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listSchemasOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSchemasOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSchemasOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listSchemasOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListSchemasOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSchema : Create schema
// Create a new schema.
func (watsonxData *WatsonxDataV2) CreateSchema(createSchemaOptions *CreateSchemaOptions) (result *CreateSchemaCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateSchemaWithContext(context.Background(), createSchemaOptions)
}

// CreateSchemaWithContext is an alternate form of the CreateSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSchemaWithContext(ctx context.Context, createSchemaOptions *CreateSchemaOptions) (result *CreateSchemaCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaOptions, "createSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSchemaOptions, "createSchemaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *createSchemaOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSchemaOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*createSchemaOptions.EngineID))

	body := make(map[string]interface{})
	if createSchemaOptions.CustomPath != nil {
		body["custom_path"] = createSchemaOptions.CustomPath
	}
	if createSchemaOptions.SchemaName != nil {
		body["schema_name"] = createSchemaOptions.SchemaName
	}
	if createSchemaOptions.BucketName != nil {
		body["bucket_name"] = createSchemaOptions.BucketName
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateSchemaCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSchema : Delete schema
// Delete a schema.
func (watsonxData *WatsonxDataV2) DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteSchemaWithContext(context.Background(), deleteSchemaOptions)
}

// DeleteSchemaWithContext is an alternate form of the DeleteSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSchemaWithContext(ctx context.Context, deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSchemaOptions, "deleteSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSchemaOptions, "deleteSchemaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteSchemaOptions.CatalogID,
		"schema_id": *deleteSchemaOptions.SchemaID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSchemaOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteSchemaOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// ListTables : List all tables
// List all tables in a schema in a catalog for a given engine.
func (watsonxData *WatsonxDataV2) ListTables(listTablesOptions *ListTablesOptions) (result *TableCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListTablesWithContext(context.Background(), listTablesOptions)
}

// ListTablesWithContext is an alternate form of the ListTables method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListTablesWithContext(ctx context.Context, listTablesOptions *ListTablesOptions) (result *TableCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTablesOptions, "listTablesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTablesOptions, "listTablesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listTablesOptions.CatalogID,
		"schema_id": *listTablesOptions.SchemaID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTablesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listTablesOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listTablesOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTable : Get table details
// Get details of a given table in a catalog and schema.
func (watsonxData *WatsonxDataV2) GetTable(getTableOptions *GetTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	return watsonxData.GetTableWithContext(context.Background(), getTableOptions)
}

// GetTableWithContext is an alternate form of the GetTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetTableWithContext(ctx context.Context, getTableOptions *GetTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTableOptions, "getTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTableOptions, "getTableOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *getTableOptions.CatalogID,
		"schema_id": *getTableOptions.SchemaID,
		"table_id": *getTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*getTableOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTable)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTable : Delete table
// Delete table for a given schema and catalog.
func (watsonxData *WatsonxDataV2) DeleteTable(deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteTableWithContext(context.Background(), deleteTableOptions)
}

// DeleteTableWithContext is an alternate form of the DeleteTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteTableWithContext(ctx context.Context, deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTableOptions, "deleteTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTableOptions, "deleteTableOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteTableOptions.CatalogID,
		"schema_id": *deleteTableOptions.SchemaID,
		"table_id": *deleteTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteTableOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateTable : Alter table
// Rename table.
func (watsonxData *WatsonxDataV2) UpdateTable(updateTableOptions *UpdateTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateTableWithContext(context.Background(), updateTableOptions)
}

// UpdateTableWithContext is an alternate form of the UpdateTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateTableWithContext(ctx context.Context, updateTableOptions *UpdateTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTableOptions, "updateTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTableOptions, "updateTableOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *updateTableOptions.CatalogID,
		"schema_id": *updateTableOptions.SchemaID,
		"table_id": *updateTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*updateTableOptions.EngineID))

	_, err = builder.SetBodyContentJSON(updateTableOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTable)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListColumns : List all columns of a table
// List all columns of a table in a given a schema for a given catalog.
func (watsonxData *WatsonxDataV2) ListColumns(listColumnsOptions *ListColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListColumnsWithContext(context.Background(), listColumnsOptions)
}

// ListColumnsWithContext is an alternate form of the ListColumns method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListColumnsWithContext(ctx context.Context, listColumnsOptions *ListColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listColumnsOptions, "listColumnsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listColumnsOptions, "listColumnsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listColumnsOptions.CatalogID,
		"schema_id": *listColumnsOptions.SchemaID,
		"table_id": *listColumnsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listColumnsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListColumns")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listColumnsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listColumnsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listColumnsOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumnCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateColumns : Add column(s)
// Add one or multiple columns to a table in a schema for a given catalog.
func (watsonxData *WatsonxDataV2) CreateColumns(createColumnsOptions *CreateColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	return watsonxData.CreateColumnsWithContext(context.Background(), createColumnsOptions)
}

// CreateColumnsWithContext is an alternate form of the CreateColumns method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateColumnsWithContext(ctx context.Context, createColumnsOptions *CreateColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createColumnsOptions, "createColumnsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createColumnsOptions, "createColumnsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *createColumnsOptions.CatalogID,
		"schema_id": *createColumnsOptions.SchemaID,
		"table_id": *createColumnsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createColumnsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateColumns")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createColumnsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createColumnsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*createColumnsOptions.EngineID))

	body := make(map[string]interface{})
	if createColumnsOptions.Columns != nil {
		body["columns"] = createColumnsOptions.Columns
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumnCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteColumn : Delete column
// Delete column in a table for a given schema and catalog.
func (watsonxData *WatsonxDataV2) DeleteColumn(deleteColumnOptions *DeleteColumnOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteColumnWithContext(context.Background(), deleteColumnOptions)
}

// DeleteColumnWithContext is an alternate form of the DeleteColumn method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteColumnWithContext(ctx context.Context, deleteColumnOptions *DeleteColumnOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteColumnOptions, "deleteColumnOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteColumnOptions, "deleteColumnOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteColumnOptions.CatalogID,
		"schema_id": *deleteColumnOptions.SchemaID,
		"table_id": *deleteColumnOptions.TableID,
		"column_id": *deleteColumnOptions.ColumnID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns/{column_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteColumnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteColumn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteColumnOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteColumnOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteColumnOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateColumn : Alter column
// Update the given column - rename column.
func (watsonxData *WatsonxDataV2) UpdateColumn(updateColumnOptions *UpdateColumnOptions) (result *Column, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateColumnWithContext(context.Background(), updateColumnOptions)
}

// UpdateColumnWithContext is an alternate form of the UpdateColumn method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateColumnWithContext(ctx context.Context, updateColumnOptions *UpdateColumnOptions) (result *Column, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateColumnOptions, "updateColumnOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateColumnOptions, "updateColumnOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *updateColumnOptions.CatalogID,
		"schema_id": *updateColumnOptions.SchemaID,
		"table_id": *updateColumnOptions.TableID,
		"column_id": *updateColumnOptions.ColumnID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns/{column_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateColumnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateColumn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateColumnOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateColumnOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*updateColumnOptions.EngineID))

	_, err = builder.SetBodyContentJSON(updateColumnOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumn)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTableSnapshots : Get table snapshots
// List all table snapshots.
func (watsonxData *WatsonxDataV2) ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions) (result *TableSnapshotCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListTableSnapshotsWithContext(context.Background(), listTableSnapshotsOptions)
}

// ListTableSnapshotsWithContext is an alternate form of the ListTableSnapshots method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListTableSnapshotsWithContext(ctx context.Context, listTableSnapshotsOptions *ListTableSnapshotsOptions) (result *TableSnapshotCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTableSnapshotsOptions, "listTableSnapshotsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTableSnapshotsOptions, "listTableSnapshotsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listTableSnapshotsOptions.CatalogID,
		"schema_id": *listTableSnapshotsOptions.SchemaID,
		"table_id": *listTableSnapshotsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/snapshots`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTableSnapshotsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListTableSnapshots")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTableSnapshotsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listTableSnapshotsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listTableSnapshotsOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableSnapshotCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceSnapshot : Rollback snapshot
// Rollback to a table snapshot.
func (watsonxData *WatsonxDataV2) ReplaceSnapshot(replaceSnapshotOptions *ReplaceSnapshotOptions) (result *ReplaceSnapshotCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.ReplaceSnapshotWithContext(context.Background(), replaceSnapshotOptions)
}

// ReplaceSnapshotWithContext is an alternate form of the ReplaceSnapshot method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ReplaceSnapshotWithContext(ctx context.Context, replaceSnapshotOptions *ReplaceSnapshotOptions) (result *ReplaceSnapshotCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceSnapshotOptions, "replaceSnapshotOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceSnapshotOptions, "replaceSnapshotOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *replaceSnapshotOptions.CatalogID,
		"schema_id": *replaceSnapshotOptions.SchemaID,
		"table_id": *replaceSnapshotOptions.TableID,
		"snapshot_id": *replaceSnapshotOptions.SnapshotID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/snapshots/{snapshot_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ReplaceSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if replaceSnapshotOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*replaceSnapshotOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*replaceSnapshotOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReplaceSnapshotCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateSyncCatalog : External Iceberg table registration
// Synchronize the external Iceberg table registration for a catalog identified by catalog_id.
func (watsonxData *WatsonxDataV2) UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions) (result *UpdateSyncCatalogOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateSyncCatalogWithContext(context.Background(), updateSyncCatalogOptions)
}

// UpdateSyncCatalogWithContext is an alternate form of the UpdateSyncCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateSyncCatalogWithContext(ctx context.Context, updateSyncCatalogOptions *UpdateSyncCatalogOptions) (result *UpdateSyncCatalogOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSyncCatalogOptions, "updateSyncCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSyncCatalogOptions, "updateSyncCatalogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *updateSyncCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/sync`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSyncCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateSyncCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateSyncCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateSyncCatalogOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateSyncCatalogOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateSyncCatalogOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListMilvusServices : Get list of milvus services
// Get list milvus services.
func (watsonxData *WatsonxDataV2) ListMilvusServices(listMilvusServicesOptions *ListMilvusServicesOptions) (result *MilvusServiceCollection, response *core.DetailedResponse, err error) {
	return watsonxData.ListMilvusServicesWithContext(context.Background(), listMilvusServicesOptions)
}

// ListMilvusServicesWithContext is an alternate form of the ListMilvusServices method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListMilvusServicesWithContext(ctx context.Context, listMilvusServicesOptions *ListMilvusServicesOptions) (result *MilvusServiceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listMilvusServicesOptions, "listMilvusServicesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMilvusServicesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListMilvusServices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listMilvusServicesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listMilvusServicesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusServiceCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateMilvusService : Create milvus service
// Create milvus service.
func (watsonxData *WatsonxDataV2) CreateMilvusService(createMilvusServiceOptions *CreateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	return watsonxData.CreateMilvusServiceWithContext(context.Background(), createMilvusServiceOptions)
}

// CreateMilvusServiceWithContext is an alternate form of the CreateMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateMilvusServiceWithContext(ctx context.Context, createMilvusServiceOptions *CreateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMilvusServiceOptions, "createMilvusServiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createMilvusServiceOptions, "createMilvusServiceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createMilvusServiceOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createMilvusServiceOptions.Origin != nil {
		body["origin"] = createMilvusServiceOptions.Origin
	}
	if createMilvusServiceOptions.Type != nil {
		body["type"] = createMilvusServiceOptions.Type
	}
	if createMilvusServiceOptions.Description != nil {
		body["description"] = createMilvusServiceOptions.Description
	}
	if createMilvusServiceOptions.ServiceDisplayName != nil {
		body["service_display_name"] = createMilvusServiceOptions.ServiceDisplayName
	}
	if createMilvusServiceOptions.Tags != nil {
		body["tags"] = createMilvusServiceOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMilvusService : Get milvus service
// Get milvus service.
func (watsonxData *WatsonxDataV2) GetMilvusService(getMilvusServiceOptions *GetMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	return watsonxData.GetMilvusServiceWithContext(context.Background(), getMilvusServiceOptions)
}

// GetMilvusServiceWithContext is an alternate form of the GetMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetMilvusServiceWithContext(ctx context.Context, getMilvusServiceOptions *GetMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMilvusServiceOptions, "getMilvusServiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMilvusServiceOptions, "getMilvusServiceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *getMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getMilvusServiceOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteMilvusService : Delete milvus service
// Delete milvus service.
func (watsonxData *WatsonxDataV2) DeleteMilvusService(deleteMilvusServiceOptions *DeleteMilvusServiceOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteMilvusServiceWithContext(context.Background(), deleteMilvusServiceOptions)
}

// DeleteMilvusServiceWithContext is an alternate form of the DeleteMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteMilvusServiceWithContext(ctx context.Context, deleteMilvusServiceOptions *DeleteMilvusServiceOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMilvusServiceOptions, "deleteMilvusServiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteMilvusServiceOptions, "deleteMilvusServiceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *deleteMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteMilvusServiceOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateMilvusService : Update milvus service
// Update details of milvus service.
func (watsonxData *WatsonxDataV2) UpdateMilvusService(updateMilvusServiceOptions *UpdateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateMilvusServiceWithContext(context.Background(), updateMilvusServiceOptions)
}

// UpdateMilvusServiceWithContext is an alternate form of the UpdateMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateMilvusServiceWithContext(ctx context.Context, updateMilvusServiceOptions *UpdateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMilvusServiceOptions, "updateMilvusServiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMilvusServiceOptions, "updateMilvusServiceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *updateMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateMilvusServiceOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateMilvusServiceOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BucketCatalog : bucket catalog.
type BucketCatalog struct {
	// catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// catalog tags.
	CatalogTags []string `json:"catalog_tags,omitempty"`

	// catalog type.
	CatalogType *string `json:"catalog_type,omitempty"`
}

// UnmarshalBucketCatalog unmarshals an instance of BucketCatalog from the specified map of raw messages.
func UnmarshalBucketCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_tags", &obj.CatalogTags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketDetails : bucket details.
type BucketDetails struct {
	// Access key ID, encrypted during bucket registration.
	AccessKey *string `json:"access_key,omitempty"`

	// actual bucket name.
	BucketName *string `json:"bucket_name" validate:"required"`

	// Cos endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// Secret access key, encrypted during bucket registration.
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewBucketDetails : Instantiate BucketDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewBucketDetails(bucketName string) (_model *BucketDetails, err error) {
	_model = &BucketDetails{
		BucketName: core.StringPtr(bucketName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalBucketDetails unmarshals an instance of BucketDetails from the specified map of raw messages.
func UnmarshalBucketDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketDetails)
	err = core.UnmarshalPrimitive(m, "access_key", &obj.AccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_key", &obj.SecretKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketRegistration : Bucket.
type BucketRegistration struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// bucket catalog.
	AssociatedCatalog *BucketCatalog `json:"associated_catalog" validate:"required"`

	// bucket details.
	BucketDetails *BucketDetails `json:"bucket_details,omitempty"`

	// bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// bucket ID auto generated during bucket registration.
	BucketID *string `json:"bucket_id,omitempty"`

	// bucket type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// Username who created the bucket.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Creation date.
	CreatedOn *string `json:"created_on" validate:"required"`

	// bucket description.
	Description *string `json:"description" validate:"required"`

	// managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// Region where the bucket is located.
	Region *string `json:"region,omitempty"`

	// mark bucket active or inactive.
	State *string `json:"state" validate:"required"`

	// tags.
	Tags []string `json:"tags,omitempty"`
}

// Constants associated with the BucketRegistration.BucketType property.
// bucket type.
const (
	BucketRegistration_BucketType_AmazonS3 = "amazon_s3"
	BucketRegistration_BucketType_AwsS3 = "aws_s3"
	BucketRegistration_BucketType_IbmCeph = "ibm_ceph"
	BucketRegistration_BucketType_IbmCos = "ibm_cos"
	BucketRegistration_BucketType_Minio = "minio"
)

// Constants associated with the BucketRegistration.ManagedBy property.
// managed by.
const (
	BucketRegistration_ManagedBy_Customer = "customer"
	BucketRegistration_ManagedBy_Ibm = "ibm"
)

// Constants associated with the BucketRegistration.State property.
// mark bucket active or inactive.
const (
	BucketRegistration_State_Active = "active"
	BucketRegistration_State_Inactive = "inactive"
)

// UnmarshalBucketRegistration unmarshals an instance of BucketRegistration from the specified map of raw messages.
func UnmarshalBucketRegistration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistration)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "associated_catalog", &obj.AssociatedCatalog, UnmarshalBucketCatalog)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "bucket_details", &obj.BucketDetails, UnmarshalBucketDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_display_name", &obj.BucketDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_id", &obj.BucketID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_type", &obj.BucketType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewBucketRegistrationPatch(bucketRegistration *BucketRegistration) (_patch []JSONPatchOperation) {
	if (bucketRegistration.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: bucketRegistration.Actions,
		})
	}
	if (bucketRegistration.AssociatedCatalog != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/associated_catalog"),
			Value: bucketRegistration.AssociatedCatalog,
		})
	}
	if (bucketRegistration.BucketDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/bucket_details"),
			Value: bucketRegistration.BucketDetails,
		})
	}
	if (bucketRegistration.BucketDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/bucket_display_name"),
			Value: bucketRegistration.BucketDisplayName,
		})
	}
	if (bucketRegistration.BucketID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/bucket_id"),
			Value: bucketRegistration.BucketID,
		})
	}
	if (bucketRegistration.BucketType != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/bucket_type"),
			Value: bucketRegistration.BucketType,
		})
	}
	if (bucketRegistration.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: bucketRegistration.CreatedBy,
		})
	}
	if (bucketRegistration.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: bucketRegistration.CreatedOn,
		})
	}
	if (bucketRegistration.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: bucketRegistration.Description,
		})
	}
	if (bucketRegistration.ManagedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/managed_by"),
			Value: bucketRegistration.ManagedBy,
		})
	}
	if (bucketRegistration.Region != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/region"),
			Value: bucketRegistration.Region,
		})
	}
	if (bucketRegistration.State != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/state"),
			Value: bucketRegistration.State,
		})
	}
	if (bucketRegistration.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: bucketRegistration.Tags,
		})
	}
	return
}

// BucketRegistrationCollection : List bucket registrations.
type BucketRegistrationCollection struct {
	// Buckets.
	BucketRegistrations []BucketRegistration `json:"bucket_registrations,omitempty"`
}

// UnmarshalBucketRegistrationCollection unmarshals an instance of BucketRegistrationCollection from the specified map of raw messages.
func UnmarshalBucketRegistrationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistrationCollection)
	err = core.UnmarshalModel(m, "bucket_registrations", &obj.BucketRegistrations, UnmarshalBucketRegistration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketRegistrationObjectCollection : List bucket objects.
type BucketRegistrationObjectCollection struct {
	// bucket object.
	Objects []string `json:"objects,omitempty"`
}

// UnmarshalBucketRegistrationObjectCollection unmarshals an instance of BucketRegistrationObjectCollection from the specified map of raw messages.
func UnmarshalBucketRegistrationObjectCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistrationObjectCollection)
	err = core.UnmarshalPrimitive(m, "objects", &obj.Objects)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketStatusResponse : object defining the response of checking if the credentials of a bucket are valid.
type BucketStatusResponse struct {
	// bucket credentials are valid or not.
	State *bool `json:"state" validate:"required"`

	// message response as per the credentials validated.
	StateMessage *string `json:"state_message" validate:"required"`
}

// UnmarshalBucketStatusResponse unmarshals an instance of BucketStatusResponse from the specified map of raw messages.
func UnmarshalBucketStatusResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketStatusResponse)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_message", &obj.StateMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Catalog : Define the catalog details.
type Catalog struct {
	// list of allowed actions.
	Actions []string `json:"actions,omitempty"`

	// Associated buckets items.
	AssociatedBuckets []string `json:"associated_buckets,omitempty"`

	// Associated databases items.
	AssociatedDatabases []string `json:"associated_databases,omitempty"`

	// Associated engines items.
	AssociatedEngines []string `json:"associated_engines,omitempty"`

	// Name for the catalog.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Table type.
	CatalogType *string `json:"catalog_type,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`

	// IBM thrift uri hostname.
	Hostname *string `json:"hostname,omitempty"`

	// Last sync time.
	LastSyncAt *string `json:"last_sync_at,omitempty"`

	// Managed by.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Catalog name.
	Metastore *string `json:"metastore,omitempty"`

	// IBM thrift uri port.
	Port *string `json:"port,omitempty"`

	// Catalog status.
	Status *string `json:"status,omitempty"`

	// Sync description.
	SyncDescription *string `json:"sync_description,omitempty"`

	// Tables not sync because data is corrupted.
	SyncException []string `json:"sync_exception,omitempty"`

	// Sync status.
	SyncStatus *string `json:"sync_status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Customer thrift uri.
	ThriftURI *string `json:"thrift_uri,omitempty"`
}

// Constants associated with the Catalog.ManagedBy property.
// Managed by.
const (
	Catalog_ManagedBy_Customer = "customer"
	Catalog_ManagedBy_Ibm = "ibm"
)

// UnmarshalCatalog unmarshals an instance of Catalog from the specified map of raw messages.
func UnmarshalCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Catalog)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_buckets", &obj.AssociatedBuckets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_databases", &obj.AssociatedDatabases)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_engines", &obj.AssociatedEngines)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_sync_at", &obj.LastSyncAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore", &obj.Metastore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_description", &obj.SyncDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_exception", &obj.SyncException)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_status", &obj.SyncStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "thrift_uri", &obj.ThriftURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogCollection : GetCatalogs OK.
type CatalogCollection struct {
	// Catalogs.
	Catalogs []Catalog `json:"catalogs,omitempty"`
}

// UnmarshalCatalogCollection unmarshals an instance of CatalogCollection from the specified map of raw messages.
func UnmarshalCatalogCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogCollection)
	err = core.UnmarshalModel(m, "catalogs", &obj.Catalogs, UnmarshalCatalog)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Column : Column.
type Column struct {
	// Column name.
	ColumnName *string `json:"column_name,omitempty"`

	// Comment.
	Comment *string `json:"comment,omitempty"`

	// Extra.
	Extra *string `json:"extra,omitempty"`

	// length.
	Length *string `json:"length,omitempty"`

	// scale.
	Scale *string `json:"scale,omitempty"`

	// Data type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalColumn unmarshals an instance of Column from the specified map of raw messages.
func UnmarshalColumn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Column)
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "extra", &obj.Extra)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "length", &obj.Length)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scale", &obj.Scale)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewColumnPatch(column *Column) (_patch []JSONPatchOperation) {
	if (column.ColumnName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/column_name"),
			Value: column.ColumnName,
		})
	}
	if (column.Comment != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/comment"),
			Value: column.Comment,
		})
	}
	if (column.Extra != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/extra"),
			Value: column.Extra,
		})
	}
	if (column.Length != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/length"),
			Value: column.Length,
		})
	}
	if (column.Scale != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/scale"),
			Value: column.Scale,
		})
	}
	if (column.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: column.Type,
		})
	}
	return
}

// ColumnCollection : list of columns in a table.
type ColumnCollection struct {
	// List of the columns present in the table.
	Columns []Column `json:"columns,omitempty"`
}

// UnmarshalColumnCollection unmarshals an instance of ColumnCollection from the specified map of raw messages.
func UnmarshalColumnCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ColumnCollection)
	err = core.UnmarshalModel(m, "columns", &obj.Columns, UnmarshalColumn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionResponse : check connection response details are valid or not.
type ConnectionResponse struct {
	// Whether the connection details are valid or not.
	State *bool `json:"state,omitempty"`

	// Connection message received by connector libraries for failed connection.
	StateMessage *string `json:"state_message,omitempty"`
}

// UnmarshalConnectionResponse unmarshals an instance of ConnectionResponse from the specified map of raw messages.
func UnmarshalConnectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionResponse)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_message", &obj.StateMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateActivateBucketCreatedBody : Activate bucket.
type CreateActivateBucketCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateActivateBucketCreatedBody unmarshals an instance of CreateActivateBucketCreatedBody from the specified map of raw messages.
func UnmarshalCreateActivateBucketCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateActivateBucketCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateActivateBucketOptions : The CreateActivateBucket options.
type CreateActivateBucketOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateActivateBucketOptions : Instantiate CreateActivateBucketOptions
func (*WatsonxDataV2) NewCreateActivateBucketOptions(bucketID string) *CreateActivateBucketOptions {
	return &CreateActivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *CreateActivateBucketOptions) SetBucketID(bucketID string) *CreateActivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateActivateBucketOptions) SetAuthInstanceID(authInstanceID string) *CreateActivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateActivateBucketOptions) SetHeaders(param map[string]string) *CreateActivateBucketOptions {
	options.Headers = param
	return options
}

// CreateBucketRegistrationOptions : The CreateBucketRegistration options.
type CreateBucketRegistrationOptions struct {
	// bucket details.
	BucketDetails *BucketDetails `json:"bucket_details" validate:"required"`

	// bucket type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// bucket description.
	Description *string `json:"description" validate:"required"`

	// managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// bucket catalog.
	AssociatedCatalog *BucketCatalog `json:"associated_catalog,omitempty"`

	// bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// region where the bucket is located.
	Region *string `json:"region,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateBucketRegistrationOptions.BucketType property.
// bucket type.
const (
	CreateBucketRegistrationOptions_BucketType_AwsS3 = "aws_s3"
	CreateBucketRegistrationOptions_BucketType_IbmCeph = "ibm_ceph"
	CreateBucketRegistrationOptions_BucketType_IbmCos = "ibm_cos"
	CreateBucketRegistrationOptions_BucketType_Minio = "minio"
)

// Constants associated with the CreateBucketRegistrationOptions.ManagedBy property.
// managed by.
const (
	CreateBucketRegistrationOptions_ManagedBy_Customer = "customer"
	CreateBucketRegistrationOptions_ManagedBy_Ibm = "ibm"
)

// NewCreateBucketRegistrationOptions : Instantiate CreateBucketRegistrationOptions
func (*WatsonxDataV2) NewCreateBucketRegistrationOptions(bucketDetails *BucketDetails, bucketType string, description string, managedBy string) *CreateBucketRegistrationOptions {
	return &CreateBucketRegistrationOptions{
		BucketDetails: bucketDetails,
		BucketType: core.StringPtr(bucketType),
		Description: core.StringPtr(description),
		ManagedBy: core.StringPtr(managedBy),
	}
}

// SetBucketDetails : Allow user to set BucketDetails
func (_options *CreateBucketRegistrationOptions) SetBucketDetails(bucketDetails *BucketDetails) *CreateBucketRegistrationOptions {
	_options.BucketDetails = bucketDetails
	return _options
}

// SetBucketType : Allow user to set BucketType
func (_options *CreateBucketRegistrationOptions) SetBucketType(bucketType string) *CreateBucketRegistrationOptions {
	_options.BucketType = core.StringPtr(bucketType)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateBucketRegistrationOptions) SetDescription(description string) *CreateBucketRegistrationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *CreateBucketRegistrationOptions) SetManagedBy(managedBy string) *CreateBucketRegistrationOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetAssociatedCatalog : Allow user to set AssociatedCatalog
func (_options *CreateBucketRegistrationOptions) SetAssociatedCatalog(associatedCatalog *BucketCatalog) *CreateBucketRegistrationOptions {
	_options.AssociatedCatalog = associatedCatalog
	return _options
}

// SetBucketDisplayName : Allow user to set BucketDisplayName
func (_options *CreateBucketRegistrationOptions) SetBucketDisplayName(bucketDisplayName string) *CreateBucketRegistrationOptions {
	_options.BucketDisplayName = core.StringPtr(bucketDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateBucketRegistrationOptions) SetRegion(region string) *CreateBucketRegistrationOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateBucketRegistrationOptions) SetTags(tags []string) *CreateBucketRegistrationOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *CreateBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBucketRegistrationOptions) SetHeaders(param map[string]string) *CreateBucketRegistrationOptions {
	options.Headers = param
	return options
}

// CreateColumnsOptions : The CreateColumns options.
type CreateColumnsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// List of the tables present in the schema.
	Columns []Column `json:"columns,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateColumnsOptions : Instantiate CreateColumnsOptions
func (*WatsonxDataV2) NewCreateColumnsOptions(engineID string, catalogID string, schemaID string, tableID string) *CreateColumnsOptions {
	return &CreateColumnsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateColumnsOptions) SetEngineID(engineID string) *CreateColumnsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *CreateColumnsOptions) SetCatalogID(catalogID string) *CreateColumnsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *CreateColumnsOptions) SetSchemaID(schemaID string) *CreateColumnsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *CreateColumnsOptions) SetTableID(tableID string) *CreateColumnsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumns : Allow user to set Columns
func (_options *CreateColumnsOptions) SetColumns(columns []Column) *CreateColumnsOptions {
	_options.Columns = columns
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateColumnsOptions) SetAuthInstanceID(authInstanceID string) *CreateColumnsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateColumnsOptions) SetHeaders(param map[string]string) *CreateColumnsOptions {
	options.Headers = param
	return options
}

// CreateDatabaseRegistrationOptions : The CreateDatabaseRegistration options.
type CreateDatabaseRegistrationOptions struct {
	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// database catalog.
	AssociatedCatalog *DatabaseCatalog `json:"associated_catalog,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// database details.
	DatabaseDetails *DatabaseDetails `json:"database_details,omitempty"`

	// This will hold all the properties for a custom database.
	DatabaseProperties []DatabaseRegistrationPrototypeDatabasePropertiesItems `json:"database_properties,omitempty"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDatabaseRegistrationOptions : Instantiate CreateDatabaseRegistrationOptions
func (*WatsonxDataV2) NewCreateDatabaseRegistrationOptions(databaseDisplayName string, databaseType string) *CreateDatabaseRegistrationOptions {
	return &CreateDatabaseRegistrationOptions{
		DatabaseDisplayName: core.StringPtr(databaseDisplayName),
		DatabaseType: core.StringPtr(databaseType),
	}
}

// SetDatabaseDisplayName : Allow user to set DatabaseDisplayName
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseDisplayName(databaseDisplayName string) *CreateDatabaseRegistrationOptions {
	_options.DatabaseDisplayName = core.StringPtr(databaseDisplayName)
	return _options
}

// SetDatabaseType : Allow user to set DatabaseType
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseType(databaseType string) *CreateDatabaseRegistrationOptions {
	_options.DatabaseType = core.StringPtr(databaseType)
	return _options
}

// SetAssociatedCatalog : Allow user to set AssociatedCatalog
func (_options *CreateDatabaseRegistrationOptions) SetAssociatedCatalog(associatedCatalog *DatabaseCatalog) *CreateDatabaseRegistrationOptions {
	_options.AssociatedCatalog = associatedCatalog
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *CreateDatabaseRegistrationOptions) SetCreatedOn(createdOn string) *CreateDatabaseRegistrationOptions {
	_options.CreatedOn = core.StringPtr(createdOn)
	return _options
}

// SetDatabaseDetails : Allow user to set DatabaseDetails
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseDetails(databaseDetails *DatabaseDetails) *CreateDatabaseRegistrationOptions {
	_options.DatabaseDetails = databaseDetails
	return _options
}

// SetDatabaseProperties : Allow user to set DatabaseProperties
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseProperties(databaseProperties []DatabaseRegistrationPrototypeDatabasePropertiesItems) *CreateDatabaseRegistrationOptions {
	_options.DatabaseProperties = databaseProperties
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDatabaseRegistrationOptions) SetDescription(description string) *CreateDatabaseRegistrationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDatabaseRegistrationOptions) SetTags(tags []string) *CreateDatabaseRegistrationOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDatabaseRegistrationOptions) SetAuthInstanceID(authInstanceID string) *CreateDatabaseRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDatabaseRegistrationOptions) SetHeaders(param map[string]string) *CreateDatabaseRegistrationOptions {
	options.Headers = param
	return options
}

// CreateDb2EngineOptions : The CreateDb2Engine options.
type CreateDb2EngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type.
	Type *string `json:"type" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *Db2EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDb2EngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateDb2EngineOptions_Origin_Discover = "discover"
	CreateDb2EngineOptions_Origin_External = "external"
	CreateDb2EngineOptions_Origin_Native = "native"
)

// NewCreateDb2EngineOptions : Instantiate CreateDb2EngineOptions
func (*WatsonxDataV2) NewCreateDb2EngineOptions(origin string, typeVar string) *CreateDb2EngineOptions {
	return &CreateDb2EngineOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateDb2EngineOptions) SetOrigin(origin string) *CreateDb2EngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateDb2EngineOptions) SetType(typeVar string) *CreateDb2EngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDb2EngineOptions) SetDescription(description string) *CreateDb2EngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateDb2EngineOptions) SetEngineDetails(engineDetails *Db2EngineDetailsBody) *CreateDb2EngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateDb2EngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateDb2EngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDb2EngineOptions) SetTags(tags []string) *CreateDb2EngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *CreateDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDb2EngineOptions) SetHeaders(param map[string]string) *CreateDb2EngineOptions {
	options.Headers = param
	return options
}

// CreateDriverDatabaseCatalogOptions : The CreateDriverDatabaseCatalog options.
type CreateDriverDatabaseCatalogOptions struct {
	// Driver file to upload.
	Driver io.ReadCloser `json:"driver" validate:"required"`

	// Name of the driver file.
	DriverFileName *string `json:"driver_file_name" validate:"required"`

	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Host name.
	Hostname *string `json:"hostname" validate:"required"`

	// Port.
	Port *string `json:"port" validate:"required"`

	// Username.
	Username *string `json:"username" validate:"required"`

	// Psssword.
	Password *string `json:"password" validate:"required"`

	// Database name.
	DatabaseName *string `json:"database_name" validate:"required"`

	// The content type of driver.
	DriverContentType *string `json:"driver_content_type,omitempty"`

	// contents of a pem/crt file.
	Certificate *string `json:"certificate,omitempty"`

	// extension of the certificate file.
	CertificateExtension *string `json:"certificate_extension,omitempty"`

	// SSL Mode.
	Ssl *string `json:"ssl,omitempty"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDriverDatabaseCatalogOptions : Instantiate CreateDriverDatabaseCatalogOptions
func (*WatsonxDataV2) NewCreateDriverDatabaseCatalogOptions(driver io.ReadCloser, driverFileName string, databaseDisplayName string, databaseType string, catalogName string, hostname string, port string, username string, password string, databaseName string) *CreateDriverDatabaseCatalogOptions {
	return &CreateDriverDatabaseCatalogOptions{
		Driver: driver,
		DriverFileName: core.StringPtr(driverFileName),
		DatabaseDisplayName: core.StringPtr(databaseDisplayName),
		DatabaseType: core.StringPtr(databaseType),
		CatalogName: core.StringPtr(catalogName),
		Hostname: core.StringPtr(hostname),
		Port: core.StringPtr(port),
		Username: core.StringPtr(username),
		Password: core.StringPtr(password),
		DatabaseName: core.StringPtr(databaseName),
	}
}

// SetDriver : Allow user to set Driver
func (_options *CreateDriverDatabaseCatalogOptions) SetDriver(driver io.ReadCloser) *CreateDriverDatabaseCatalogOptions {
	_options.Driver = driver
	return _options
}

// SetDriverFileName : Allow user to set DriverFileName
func (_options *CreateDriverDatabaseCatalogOptions) SetDriverFileName(driverFileName string) *CreateDriverDatabaseCatalogOptions {
	_options.DriverFileName = core.StringPtr(driverFileName)
	return _options
}

// SetDatabaseDisplayName : Allow user to set DatabaseDisplayName
func (_options *CreateDriverDatabaseCatalogOptions) SetDatabaseDisplayName(databaseDisplayName string) *CreateDriverDatabaseCatalogOptions {
	_options.DatabaseDisplayName = core.StringPtr(databaseDisplayName)
	return _options
}

// SetDatabaseType : Allow user to set DatabaseType
func (_options *CreateDriverDatabaseCatalogOptions) SetDatabaseType(databaseType string) *CreateDriverDatabaseCatalogOptions {
	_options.DatabaseType = core.StringPtr(databaseType)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *CreateDriverDatabaseCatalogOptions) SetCatalogName(catalogName string) *CreateDriverDatabaseCatalogOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetHostname : Allow user to set Hostname
func (_options *CreateDriverDatabaseCatalogOptions) SetHostname(hostname string) *CreateDriverDatabaseCatalogOptions {
	_options.Hostname = core.StringPtr(hostname)
	return _options
}

// SetPort : Allow user to set Port
func (_options *CreateDriverDatabaseCatalogOptions) SetPort(port string) *CreateDriverDatabaseCatalogOptions {
	_options.Port = core.StringPtr(port)
	return _options
}

// SetUsername : Allow user to set Username
func (_options *CreateDriverDatabaseCatalogOptions) SetUsername(username string) *CreateDriverDatabaseCatalogOptions {
	_options.Username = core.StringPtr(username)
	return _options
}

// SetPassword : Allow user to set Password
func (_options *CreateDriverDatabaseCatalogOptions) SetPassword(password string) *CreateDriverDatabaseCatalogOptions {
	_options.Password = core.StringPtr(password)
	return _options
}

// SetDatabaseName : Allow user to set DatabaseName
func (_options *CreateDriverDatabaseCatalogOptions) SetDatabaseName(databaseName string) *CreateDriverDatabaseCatalogOptions {
	_options.DatabaseName = core.StringPtr(databaseName)
	return _options
}

// SetDriverContentType : Allow user to set DriverContentType
func (_options *CreateDriverDatabaseCatalogOptions) SetDriverContentType(driverContentType string) *CreateDriverDatabaseCatalogOptions {
	_options.DriverContentType = core.StringPtr(driverContentType)
	return _options
}

// SetCertificate : Allow user to set Certificate
func (_options *CreateDriverDatabaseCatalogOptions) SetCertificate(certificate string) *CreateDriverDatabaseCatalogOptions {
	_options.Certificate = core.StringPtr(certificate)
	return _options
}

// SetCertificateExtension : Allow user to set CertificateExtension
func (_options *CreateDriverDatabaseCatalogOptions) SetCertificateExtension(certificateExtension string) *CreateDriverDatabaseCatalogOptions {
	_options.CertificateExtension = core.StringPtr(certificateExtension)
	return _options
}

// SetSsl : Allow user to set Ssl
func (_options *CreateDriverDatabaseCatalogOptions) SetSsl(ssl string) *CreateDriverDatabaseCatalogOptions {
	_options.Ssl = core.StringPtr(ssl)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDriverDatabaseCatalogOptions) SetDescription(description string) *CreateDriverDatabaseCatalogOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *CreateDriverDatabaseCatalogOptions) SetCreatedOn(createdOn string) *CreateDriverDatabaseCatalogOptions {
	_options.CreatedOn = core.StringPtr(createdOn)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDriverDatabaseCatalogOptions) SetAuthInstanceID(authInstanceID string) *CreateDriverDatabaseCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDriverDatabaseCatalogOptions) SetHeaders(param map[string]string) *CreateDriverDatabaseCatalogOptions {
	options.Headers = param
	return options
}

// CreateEnginePauseCreatedBody : Pause.
type CreateEnginePauseCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEnginePauseCreatedBody unmarshals an instance of CreateEnginePauseCreatedBody from the specified map of raw messages.
func UnmarshalCreateEnginePauseCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEnginePauseCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEnginePauseOptions : The CreateEnginePause options.
type CreateEnginePauseOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEnginePauseOptions : Instantiate CreateEnginePauseOptions
func (*WatsonxDataV2) NewCreateEnginePauseOptions(engineID string) *CreateEnginePauseOptions {
	return &CreateEnginePauseOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateEnginePauseOptions) SetEngineID(engineID string) *CreateEnginePauseOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEnginePauseOptions) SetAuthInstanceID(authInstanceID string) *CreateEnginePauseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnginePauseOptions) SetHeaders(param map[string]string) *CreateEnginePauseOptions {
	options.Headers = param
	return options
}

// CreateEngineRestartCreatedBody : restart engine.
type CreateEngineRestartCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineRestartCreatedBody unmarshals an instance of CreateEngineRestartCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineRestartCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineRestartCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineRestartOptions : The CreateEngineRestart options.
type CreateEngineRestartOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEngineRestartOptions : Instantiate CreateEngineRestartOptions
func (*WatsonxDataV2) NewCreateEngineRestartOptions(engineID string) *CreateEngineRestartOptions {
	return &CreateEngineRestartOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateEngineRestartOptions) SetEngineID(engineID string) *CreateEngineRestartOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEngineRestartOptions) SetAuthInstanceID(authInstanceID string) *CreateEngineRestartOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEngineRestartOptions) SetHeaders(param map[string]string) *CreateEngineRestartOptions {
	options.Headers = param
	return options
}

// CreateEngineResumeCreatedBody : resume.
type CreateEngineResumeCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineResumeCreatedBody unmarshals an instance of CreateEngineResumeCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineResumeCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineResumeCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineResumeOptions : The CreateEngineResume options.
type CreateEngineResumeOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEngineResumeOptions : Instantiate CreateEngineResumeOptions
func (*WatsonxDataV2) NewCreateEngineResumeOptions(engineID string) *CreateEngineResumeOptions {
	return &CreateEngineResumeOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateEngineResumeOptions) SetEngineID(engineID string) *CreateEngineResumeOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEngineResumeOptions) SetAuthInstanceID(authInstanceID string) *CreateEngineResumeOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEngineResumeOptions) SetHeaders(param map[string]string) *CreateEngineResumeOptions {
	options.Headers = param
	return options
}

// CreateEngineScaleCreatedBody : scale engine.
type CreateEngineScaleCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineScaleCreatedBody unmarshals an instance of CreateEngineScaleCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineScaleCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineScaleCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineScaleOptions : The CreateEngineScale options.
type CreateEngineScaleOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEngineScaleOptions : Instantiate CreateEngineScaleOptions
func (*WatsonxDataV2) NewCreateEngineScaleOptions(engineID string) *CreateEngineScaleOptions {
	return &CreateEngineScaleOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateEngineScaleOptions) SetEngineID(engineID string) *CreateEngineScaleOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCoordinator : Allow user to set Coordinator
func (_options *CreateEngineScaleOptions) SetCoordinator(coordinator *NodeDescription) *CreateEngineScaleOptions {
	_options.Coordinator = coordinator
	return _options
}

// SetWorker : Allow user to set Worker
func (_options *CreateEngineScaleOptions) SetWorker(worker *NodeDescription) *CreateEngineScaleOptions {
	_options.Worker = worker
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEngineScaleOptions) SetAuthInstanceID(authInstanceID string) *CreateEngineScaleOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEngineScaleOptions) SetHeaders(param map[string]string) *CreateEngineScaleOptions {
	options.Headers = param
	return options
}

// CreateMilvusServiceOptions : The CreateMilvusService options.
type CreateMilvusServiceOptions struct {
	// Origin - place holder.
	Origin *string `json:"origin" validate:"required"`

	// service type.
	Type *string `json:"type" validate:"required"`

	// Service description.
	Description *string `json:"description,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateMilvusServiceOptions : Instantiate CreateMilvusServiceOptions
func (*WatsonxDataV2) NewCreateMilvusServiceOptions(origin string, typeVar string) *CreateMilvusServiceOptions {
	return &CreateMilvusServiceOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateMilvusServiceOptions) SetOrigin(origin string) *CreateMilvusServiceOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateMilvusServiceOptions) SetType(typeVar string) *CreateMilvusServiceOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateMilvusServiceOptions) SetDescription(description string) *CreateMilvusServiceOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetServiceDisplayName : Allow user to set ServiceDisplayName
func (_options *CreateMilvusServiceOptions) SetServiceDisplayName(serviceDisplayName string) *CreateMilvusServiceOptions {
	_options.ServiceDisplayName = core.StringPtr(serviceDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateMilvusServiceOptions) SetTags(tags []string) *CreateMilvusServiceOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *CreateMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMilvusServiceOptions) SetHeaders(param map[string]string) *CreateMilvusServiceOptions {
	options.Headers = param
	return options
}

// CreateNetezzaEngineOptions : The CreateNetezzaEngine options.
type CreateNetezzaEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type.
	Type *string `json:"type" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *NetezzaEngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateNetezzaEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateNetezzaEngineOptions_Origin_Discover = "discover"
	CreateNetezzaEngineOptions_Origin_External = "external"
	CreateNetezzaEngineOptions_Origin_Native = "native"
)

// NewCreateNetezzaEngineOptions : Instantiate CreateNetezzaEngineOptions
func (*WatsonxDataV2) NewCreateNetezzaEngineOptions(origin string, typeVar string) *CreateNetezzaEngineOptions {
	return &CreateNetezzaEngineOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateNetezzaEngineOptions) SetOrigin(origin string) *CreateNetezzaEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateNetezzaEngineOptions) SetType(typeVar string) *CreateNetezzaEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateNetezzaEngineOptions) SetDescription(description string) *CreateNetezzaEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateNetezzaEngineOptions) SetEngineDetails(engineDetails *NetezzaEngineDetailsBody) *CreateNetezzaEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateNetezzaEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateNetezzaEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateNetezzaEngineOptions) SetTags(tags []string) *CreateNetezzaEngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNetezzaEngineOptions) SetHeaders(param map[string]string) *CreateNetezzaEngineOptions {
	options.Headers = param
	return options
}

// CreateOtherEngineOptions : The CreateOtherEngine options.
type CreateOtherEngineOptions struct {
	// External engine details.
	EngineDetails *OtherEngineDetailsBody `json:"engine_details" validate:"required"`

	// engine display name.
	EngineDisplayName *string `json:"engine_display_name" validate:"required"`

	// engine description.
	Description *string `json:"description,omitempty"`

	// Origin - created or registered.
	Origin *string `json:"origin,omitempty"`

	// other engine tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateOtherEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateOtherEngineOptions_Origin_Discover = "discover"
	CreateOtherEngineOptions_Origin_External = "external"
	CreateOtherEngineOptions_Origin_Native = "native"
)

// NewCreateOtherEngineOptions : Instantiate CreateOtherEngineOptions
func (*WatsonxDataV2) NewCreateOtherEngineOptions(engineDetails *OtherEngineDetailsBody, engineDisplayName string) *CreateOtherEngineOptions {
	return &CreateOtherEngineOptions{
		EngineDetails: engineDetails,
		EngineDisplayName: core.StringPtr(engineDisplayName),
	}
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateOtherEngineOptions) SetEngineDetails(engineDetails *OtherEngineDetailsBody) *CreateOtherEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateOtherEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateOtherEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateOtherEngineOptions) SetDescription(description string) *CreateOtherEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetOrigin : Allow user to set Origin
func (_options *CreateOtherEngineOptions) SetOrigin(origin string) *CreateOtherEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateOtherEngineOptions) SetTags(tags []string) *CreateOtherEngineOptions {
	_options.Tags = tags
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateOtherEngineOptions) SetType(typeVar string) *CreateOtherEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateOtherEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateOtherEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateOtherEngineOptions) SetHeaders(param map[string]string) *CreateOtherEngineOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEngineOptions : The CreatePrestissimoEngine options.
type CreatePrestissimoEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type prestissimo, others like netezza.
	Type *string `json:"type" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *PrestissimoEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Version like 0.278 for prestissimo or else.
	Version *string `json:"version,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePrestissimoEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreatePrestissimoEngineOptions_Origin_Discover = "discover"
	CreatePrestissimoEngineOptions_Origin_External = "external"
	CreatePrestissimoEngineOptions_Origin_Native = "native"
)

// NewCreatePrestissimoEngineOptions : Instantiate CreatePrestissimoEngineOptions
func (*WatsonxDataV2) NewCreatePrestissimoEngineOptions(origin string, typeVar string) *CreatePrestissimoEngineOptions {
	return &CreatePrestissimoEngineOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreatePrestissimoEngineOptions) SetOrigin(origin string) *CreatePrestissimoEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreatePrestissimoEngineOptions) SetType(typeVar string) *CreatePrestissimoEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreatePrestissimoEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreatePrestissimoEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePrestissimoEngineOptions) SetDescription(description string) *CreatePrestissimoEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreatePrestissimoEngineOptions) SetEngineDetails(engineDetails *PrestissimoEngineDetails) *CreatePrestissimoEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreatePrestissimoEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreatePrestissimoEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreatePrestissimoEngineOptions) SetRegion(region string) *CreatePrestissimoEngineOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreatePrestissimoEngineOptions) SetTags(tags []string) *CreatePrestissimoEngineOptions {
	_options.Tags = tags
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreatePrestissimoEngineOptions) SetVersion(version string) *CreatePrestissimoEngineOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEngineOptions) SetHeaders(param map[string]string) *CreatePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEnginePauseOptions : The CreatePrestissimoEnginePause options.
type CreatePrestissimoEnginePauseOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePrestissimoEnginePauseOptions : Instantiate CreatePrestissimoEnginePauseOptions
func (*WatsonxDataV2) NewCreatePrestissimoEnginePauseOptions(engineID string) *CreatePrestissimoEnginePauseOptions {
	return &CreatePrestissimoEnginePauseOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreatePrestissimoEnginePauseOptions) SetEngineID(engineID string) *CreatePrestissimoEnginePauseOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEnginePauseOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEnginePauseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEnginePauseOptions) SetHeaders(param map[string]string) *CreatePrestissimoEnginePauseOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEngineRestartOptions : The CreatePrestissimoEngineRestart options.
type CreatePrestissimoEngineRestartOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePrestissimoEngineRestartOptions : Instantiate CreatePrestissimoEngineRestartOptions
func (*WatsonxDataV2) NewCreatePrestissimoEngineRestartOptions(engineID string) *CreatePrestissimoEngineRestartOptions {
	return &CreatePrestissimoEngineRestartOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreatePrestissimoEngineRestartOptions) SetEngineID(engineID string) *CreatePrestissimoEngineRestartOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEngineRestartOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEngineRestartOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEngineRestartOptions) SetHeaders(param map[string]string) *CreatePrestissimoEngineRestartOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEngineResumeOptions : The CreatePrestissimoEngineResume options.
type CreatePrestissimoEngineResumeOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePrestissimoEngineResumeOptions : Instantiate CreatePrestissimoEngineResumeOptions
func (*WatsonxDataV2) NewCreatePrestissimoEngineResumeOptions(engineID string) *CreatePrestissimoEngineResumeOptions {
	return &CreatePrestissimoEngineResumeOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreatePrestissimoEngineResumeOptions) SetEngineID(engineID string) *CreatePrestissimoEngineResumeOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEngineResumeOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEngineResumeOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEngineResumeOptions) SetHeaders(param map[string]string) *CreatePrestissimoEngineResumeOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEngineScaleOptions : The CreatePrestissimoEngineScale options.
type CreatePrestissimoEngineScaleOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePrestissimoEngineScaleOptions : Instantiate CreatePrestissimoEngineScaleOptions
func (*WatsonxDataV2) NewCreatePrestissimoEngineScaleOptions(engineID string) *CreatePrestissimoEngineScaleOptions {
	return &CreatePrestissimoEngineScaleOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreatePrestissimoEngineScaleOptions) SetEngineID(engineID string) *CreatePrestissimoEngineScaleOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCoordinator : Allow user to set Coordinator
func (_options *CreatePrestissimoEngineScaleOptions) SetCoordinator(coordinator *PrestissimoNodeDescriptionBody) *CreatePrestissimoEngineScaleOptions {
	_options.Coordinator = coordinator
	return _options
}

// SetWorker : Allow user to set Worker
func (_options *CreatePrestissimoEngineScaleOptions) SetWorker(worker *PrestissimoNodeDescriptionBody) *CreatePrestissimoEngineScaleOptions {
	_options.Worker = worker
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEngineScaleOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEngineScaleOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEngineScaleOptions) SetHeaders(param map[string]string) *CreatePrestissimoEngineScaleOptions {
	options.Headers = param
	return options
}

// CreatePrestoEngineOptions : The CreatePrestoEngine options.
type CreatePrestoEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type presto.
	Type *string `json:"type" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Node details.
	EngineDetails *EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Version like 0.278 for presto or else.
	Version *string `json:"version,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePrestoEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreatePrestoEngineOptions_Origin_Discover = "discover"
	CreatePrestoEngineOptions_Origin_External = "external"
	CreatePrestoEngineOptions_Origin_Native = "native"
)

// NewCreatePrestoEngineOptions : Instantiate CreatePrestoEngineOptions
func (*WatsonxDataV2) NewCreatePrestoEngineOptions(origin string, typeVar string) *CreatePrestoEngineOptions {
	return &CreatePrestoEngineOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreatePrestoEngineOptions) SetOrigin(origin string) *CreatePrestoEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreatePrestoEngineOptions) SetType(typeVar string) *CreatePrestoEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreatePrestoEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreatePrestoEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePrestoEngineOptions) SetDescription(description string) *CreatePrestoEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreatePrestoEngineOptions) SetEngineDetails(engineDetails *EngineDetailsBody) *CreatePrestoEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreatePrestoEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreatePrestoEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreatePrestoEngineOptions) SetRegion(region string) *CreatePrestoEngineOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreatePrestoEngineOptions) SetTags(tags []string) *CreatePrestoEngineOptions {
	_options.Tags = tags
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreatePrestoEngineOptions) SetVersion(version string) *CreatePrestoEngineOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestoEngineOptions) SetHeaders(param map[string]string) *CreatePrestoEngineOptions {
	options.Headers = param
	return options
}

// CreateSchemaCreatedBody : success response.
type CreateSchemaCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateSchemaCreatedBody unmarshals an instance of CreateSchemaCreatedBody from the specified map of raw messages.
func UnmarshalCreateSchemaCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateSchemaCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateSchemaOptions : The CreateSchema options.
type CreateSchemaOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Path associated with bucket.
	CustomPath *string `json:"custom_path" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Bucket associated to metastore where schema will be added.
	BucketName *string `json:"bucket_name,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSchemaOptions : Instantiate CreateSchemaOptions
func (*WatsonxDataV2) NewCreateSchemaOptions(engineID string, catalogID string, customPath string, schemaName string) *CreateSchemaOptions {
	return &CreateSchemaOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		CustomPath: core.StringPtr(customPath),
		SchemaName: core.StringPtr(schemaName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSchemaOptions) SetEngineID(engineID string) *CreateSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *CreateSchemaOptions) SetCatalogID(catalogID string) *CreateSchemaOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetCustomPath : Allow user to set CustomPath
func (_options *CreateSchemaOptions) SetCustomPath(customPath string) *CreateSchemaOptions {
	_options.CustomPath = core.StringPtr(customPath)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *CreateSchemaOptions) SetSchemaName(schemaName string) *CreateSchemaOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetBucketName : Allow user to set BucketName
func (_options *CreateSchemaOptions) SetBucketName(bucketName string) *CreateSchemaOptions {
	_options.BucketName = core.StringPtr(bucketName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSchemaOptions) SetAuthInstanceID(authInstanceID string) *CreateSchemaOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSchemaOptions) SetHeaders(param map[string]string) *CreateSchemaOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineApplicationOptions : The CreateSparkEngineApplication options.
type CreateSparkEngineApplicationOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application details.
	ApplicationDetails *SparkApplicationDetails `json:"application_details" validate:"required"`

	// Job endpoint.
	JobEndpoint *string `json:"job_endpoint,omitempty"`

	// Service Instance ID for POST.
	ServiceInstanceID *string `json:"service_instance_id,omitempty"`

	// Engine Type.
	Type *string `json:"type,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateSparkEngineApplicationOptions.Type property.
// Engine Type.
const (
	CreateSparkEngineApplicationOptions_Type_Emr = "emr"
	CreateSparkEngineApplicationOptions_Type_Iae = "iae"
)

// NewCreateSparkEngineApplicationOptions : Instantiate CreateSparkEngineApplicationOptions
func (*WatsonxDataV2) NewCreateSparkEngineApplicationOptions(engineID string, applicationDetails *SparkApplicationDetails) *CreateSparkEngineApplicationOptions {
	return &CreateSparkEngineApplicationOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationDetails: applicationDetails,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSparkEngineApplicationOptions) SetEngineID(engineID string) *CreateSparkEngineApplicationOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationDetails : Allow user to set ApplicationDetails
func (_options *CreateSparkEngineApplicationOptions) SetApplicationDetails(applicationDetails *SparkApplicationDetails) *CreateSparkEngineApplicationOptions {
	_options.ApplicationDetails = applicationDetails
	return _options
}

// SetJobEndpoint : Allow user to set JobEndpoint
func (_options *CreateSparkEngineApplicationOptions) SetJobEndpoint(jobEndpoint string) *CreateSparkEngineApplicationOptions {
	_options.JobEndpoint = core.StringPtr(jobEndpoint)
	return _options
}

// SetServiceInstanceID : Allow user to set ServiceInstanceID
func (_options *CreateSparkEngineApplicationOptions) SetServiceInstanceID(serviceInstanceID string) *CreateSparkEngineApplicationOptions {
	_options.ServiceInstanceID = core.StringPtr(serviceInstanceID)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateSparkEngineApplicationOptions) SetType(typeVar string) *CreateSparkEngineApplicationOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineApplicationOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineApplicationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineApplicationOptions) SetHeaders(param map[string]string) *CreateSparkEngineApplicationOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineOptions : The CreateSparkEngine options.
type CreateSparkEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type spark, others like netezza.
	Type *string `json:"type" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Node details.
	EngineDetails *SparkEngineDetailsPrototype `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateSparkEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateSparkEngineOptions_Origin_Discover = "discover"
	CreateSparkEngineOptions_Origin_External = "external"
)

// NewCreateSparkEngineOptions : Instantiate CreateSparkEngineOptions
func (*WatsonxDataV2) NewCreateSparkEngineOptions(origin string, typeVar string) *CreateSparkEngineOptions {
	return &CreateSparkEngineOptions{
		Origin: core.StringPtr(origin),
		Type: core.StringPtr(typeVar),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateSparkEngineOptions) SetOrigin(origin string) *CreateSparkEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateSparkEngineOptions) SetType(typeVar string) *CreateSparkEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateSparkEngineOptions) SetDescription(description string) *CreateSparkEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateSparkEngineOptions) SetEngineDetails(engineDetails *SparkEngineDetailsPrototype) *CreateSparkEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateSparkEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateSparkEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateSparkEngineOptions) SetTags(tags []string) *CreateSparkEngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineOptions) SetHeaders(param map[string]string) *CreateSparkEngineOptions {
	options.Headers = param
	return options
}

// DatabaseCatalog : database catalog.
type DatabaseCatalog struct {
	// catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// catalog tags.
	CatalogTags []string `json:"catalog_tags,omitempty"`

	// catalog type.
	CatalogType *string `json:"catalog_type,omitempty"`
}

// UnmarshalDatabaseCatalog unmarshals an instance of DatabaseCatalog from the specified map of raw messages.
func UnmarshalDatabaseCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_tags", &obj.CatalogTags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseDetails : database details.
type DatabaseDetails struct {
	// contents of a pem/crt file.
	Certificate *string `json:"certificate,omitempty"`

	// extension of the certificate file.
	CertificateExtension *string `json:"certificate_extension,omitempty"`

	// Database name.
	DatabaseName *string `json:"database_name,omitempty"`

	// Host name.
	Hostname *string `json:"hostname" validate:"required"`

	// Hostname in certificate.
	HostnameInCertificate *string `json:"hostname_in_certificate,omitempty"`

	// String of hostname:port.
	Hosts *string `json:"hosts,omitempty"`

	// Psssword.
	Password *string `json:"password,omitempty"`

	// Port.
	Port *int64 `json:"port" validate:"required"`

	// SASL Mode.
	Sasl *bool `json:"sasl,omitempty"`

	// SSL Mode.
	Ssl *bool `json:"ssl,omitempty"`

	// Only for Kafka - Add kafka tables.
	Tables *string `json:"tables,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`

	// Verify certificate.
	ValidateServerCertificate *bool `json:"validate_server_certificate,omitempty"`
}

// NewDatabaseDetails : Instantiate DatabaseDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewDatabaseDetails(hostname string, port int64) (_model *DatabaseDetails, err error) {
	_model = &DatabaseDetails{
		Hostname: core.StringPtr(hostname),
		Port: core.Int64Ptr(port),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDatabaseDetails unmarshals an instance of DatabaseDetails from the specified map of raw messages.
func UnmarshalDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseDetails)
	err = core.UnmarshalPrimitive(m, "certificate", &obj.Certificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_extension", &obj.CertificateExtension)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname_in_certificate", &obj.HostnameInCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hosts", &obj.Hosts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sasl", &obj.Sasl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "validate_server_certificate", &obj.ValidateServerCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistration : database registration object.
type DatabaseRegistration struct {
	// actions.
	Actions []string `json:"actions,omitempty"`

	// database catalog.
	AssociatedCatalog *DatabaseCatalog `json:"associated_catalog,omitempty"`

	// Catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// database details.
	DatabaseDetails *DatabaseDetails `json:"database_details" validate:"required"`

	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Database ID.
	DatabaseID *string `json:"database_id,omitempty"`

	// This will hold all the properties for a custom database.
	DatabaseProperties []DatabaseRegistrationDatabasePropertiesItems `json:"database_properties,omitempty"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalDatabaseRegistration unmarshals an instance of DatabaseRegistration from the specified map of raw messages.
func UnmarshalDatabaseRegistration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistration)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "associated_catalog", &obj.AssociatedCatalog, UnmarshalDatabaseCatalog)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "database_details", &obj.DatabaseDetails, UnmarshalDatabaseDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_display_name", &obj.DatabaseDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_id", &obj.DatabaseID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "database_properties", &obj.DatabaseProperties, UnmarshalDatabaseRegistrationDatabasePropertiesItems)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_type", &obj.DatabaseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewDatabaseRegistrationPatch(databaseRegistration *DatabaseRegistration) (_patch []JSONPatchOperation) {
	if (databaseRegistration.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: databaseRegistration.Actions,
		})
	}
	if (databaseRegistration.AssociatedCatalog != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/associated_catalog"),
			Value: databaseRegistration.AssociatedCatalog,
		})
	}
	if (databaseRegistration.CatalogName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/catalog_name"),
			Value: databaseRegistration.CatalogName,
		})
	}
	if (databaseRegistration.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: databaseRegistration.CreatedBy,
		})
	}
	if (databaseRegistration.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: databaseRegistration.CreatedOn,
		})
	}
	if (databaseRegistration.DatabaseDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/database_details"),
			Value: databaseRegistration.DatabaseDetails,
		})
	}
	if (databaseRegistration.DatabaseDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/database_display_name"),
			Value: databaseRegistration.DatabaseDisplayName,
		})
	}
	if (databaseRegistration.DatabaseID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/database_id"),
			Value: databaseRegistration.DatabaseID,
		})
	}
	if (databaseRegistration.DatabaseProperties != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/database_properties"),
			Value: databaseRegistration.DatabaseProperties,
		})
	}
	if (databaseRegistration.DatabaseType != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/database_type"),
			Value: databaseRegistration.DatabaseType,
		})
	}
	if (databaseRegistration.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: databaseRegistration.Description,
		})
	}
	if (databaseRegistration.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: databaseRegistration.Tags,
		})
	}
	return
}

// DatabaseRegistrationCollection : list database registrations.
type DatabaseRegistrationCollection struct {
	// Database body.
	DatabaseRegistrations []DatabaseRegistration `json:"database_registrations,omitempty"`
}

// UnmarshalDatabaseRegistrationCollection unmarshals an instance of DatabaseRegistrationCollection from the specified map of raw messages.
func UnmarshalDatabaseRegistrationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationCollection)
	err = core.UnmarshalModel(m, "database_registrations", &obj.DatabaseRegistrations, UnmarshalDatabaseRegistration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistrationDatabasePropertiesItems : Key value object.
type DatabaseRegistrationDatabasePropertiesItems struct {
	// Wether the value is to be encrypted before storing.
	Encrypt *bool `json:"encrypt" validate:"required"`

	// Key of the database property.
	Key *string `json:"key" validate:"required"`

	// Value of the database property.
	Value *string `json:"value" validate:"required"`
}

// UnmarshalDatabaseRegistrationDatabasePropertiesItems unmarshals an instance of DatabaseRegistrationDatabasePropertiesItems from the specified map of raw messages.
func UnmarshalDatabaseRegistrationDatabasePropertiesItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationDatabasePropertiesItems)
	err = core.UnmarshalPrimitive(m, "encrypt", &obj.Encrypt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistrationPrototypeDatabasePropertiesItems : Key value object.
type DatabaseRegistrationPrototypeDatabasePropertiesItems struct {
	// Wether the value is to be encrypted before storing.
	Encrypt *bool `json:"encrypt" validate:"required"`

	// Key of the database property.
	Key *string `json:"key" validate:"required"`

	// Value of the database property.
	Value *string `json:"value" validate:"required"`
}

// NewDatabaseRegistrationPrototypeDatabasePropertiesItems : Instantiate DatabaseRegistrationPrototypeDatabasePropertiesItems (Generic Model Constructor)
func (*WatsonxDataV2) NewDatabaseRegistrationPrototypeDatabasePropertiesItems(encrypt bool, key string, value string) (_model *DatabaseRegistrationPrototypeDatabasePropertiesItems, err error) {
	_model = &DatabaseRegistrationPrototypeDatabasePropertiesItems{
		Encrypt: core.BoolPtr(encrypt),
		Key: core.StringPtr(key),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDatabaseRegistrationPrototypeDatabasePropertiesItems unmarshals an instance of DatabaseRegistrationPrototypeDatabasePropertiesItems from the specified map of raw messages.
func UnmarshalDatabaseRegistrationPrototypeDatabasePropertiesItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationPrototypeDatabasePropertiesItems)
	err = core.UnmarshalPrimitive(m, "encrypt", &obj.Encrypt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2Engine : Db2 engine details.
type Db2Engine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *Db2EngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalDb2Engine unmarshals an instance of Db2Engine from the specified map of raw messages.
func UnmarshalDb2Engine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2Engine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalDb2EngineDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewDb2EnginePatch(db2Engine *Db2Engine) (_patch []JSONPatchOperation) {
	if (db2Engine.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: db2Engine.Actions,
		})
	}
	if (db2Engine.BuildVersion != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/build_version"),
			Value: db2Engine.BuildVersion,
		})
	}
	if (db2Engine.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: db2Engine.CreatedBy,
		})
	}
	if (db2Engine.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: db2Engine.CreatedOn,
		})
	}
	if (db2Engine.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: db2Engine.Description,
		})
	}
	if (db2Engine.EngineDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_details"),
			Value: db2Engine.EngineDetails,
		})
	}
	if (db2Engine.EngineDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_display_name"),
			Value: db2Engine.EngineDisplayName,
		})
	}
	if (db2Engine.EngineID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_id"),
			Value: db2Engine.EngineID,
		})
	}
	if (db2Engine.HostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/host_name"),
			Value: db2Engine.HostName,
		})
	}
	if (db2Engine.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: db2Engine.Origin,
		})
	}
	if (db2Engine.Port != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/port"),
			Value: db2Engine.Port,
		})
	}
	if (db2Engine.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: db2Engine.Status,
		})
	}
	if (db2Engine.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: db2Engine.Tags,
		})
	}
	if (db2Engine.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: db2Engine.Type,
		})
	}
	return
}

// Db2EngineCollection : list db2 engines.
type Db2EngineCollection struct {
	// list db2 engines.
	Db2Engines []Db2Engine `json:"db2_engines,omitempty"`
}

// UnmarshalDb2EngineCollection unmarshals an instance of Db2EngineCollection from the specified map of raw messages.
func UnmarshalDb2EngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineCollection)
	err = core.UnmarshalModel(m, "db2_engines", &obj.Db2Engines, UnmarshalDb2Engine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EngineDetails : External engine details.
type Db2EngineDetails struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalDb2EngineDetails unmarshals an instance of Db2EngineDetails from the specified map of raw messages.
func UnmarshalDb2EngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EngineDetailsBody : External engine details.
type Db2EngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`
}

// UnmarshalDb2EngineDetailsBody unmarshals an instance of Db2EngineDetailsBody from the specified map of raw messages.
func UnmarshalDb2EngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteBucketRegistrationOptions : The DeleteBucketRegistration options.
type DeleteBucketRegistrationOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketRegistrationOptions : Instantiate DeleteBucketRegistrationOptions
func (*WatsonxDataV2) NewDeleteBucketRegistrationOptions(bucketID string) *DeleteBucketRegistrationOptions {
	return &DeleteBucketRegistrationOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeleteBucketRegistrationOptions) SetBucketID(bucketID string) *DeleteBucketRegistrationOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *DeleteBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketRegistrationOptions) SetHeaders(param map[string]string) *DeleteBucketRegistrationOptions {
	options.Headers = param
	return options
}

// DeleteColumnOptions : The DeleteColumn options.
type DeleteColumnOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// URL encoded schema name.
	ColumnID *string `json:"column_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteColumnOptions : Instantiate DeleteColumnOptions
func (*WatsonxDataV2) NewDeleteColumnOptions(engineID string, catalogID string, schemaID string, tableID string, columnID string) *DeleteColumnOptions {
	return &DeleteColumnOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		ColumnID: core.StringPtr(columnID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteColumnOptions) SetEngineID(engineID string) *DeleteColumnOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteColumnOptions) SetCatalogID(catalogID string) *DeleteColumnOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteColumnOptions) SetSchemaID(schemaID string) *DeleteColumnOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *DeleteColumnOptions) SetTableID(tableID string) *DeleteColumnOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumnID : Allow user to set ColumnID
func (_options *DeleteColumnOptions) SetColumnID(columnID string) *DeleteColumnOptions {
	_options.ColumnID = core.StringPtr(columnID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteColumnOptions) SetAuthInstanceID(authInstanceID string) *DeleteColumnOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteColumnOptions) SetHeaders(param map[string]string) *DeleteColumnOptions {
	options.Headers = param
	return options
}

// DeleteDatabaseCatalogOptions : The DeleteDatabaseCatalog options.
type DeleteDatabaseCatalogOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDatabaseCatalogOptions : Instantiate DeleteDatabaseCatalogOptions
func (*WatsonxDataV2) NewDeleteDatabaseCatalogOptions(databaseID string) *DeleteDatabaseCatalogOptions {
	return &DeleteDatabaseCatalogOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *DeleteDatabaseCatalogOptions) SetDatabaseID(databaseID string) *DeleteDatabaseCatalogOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDatabaseCatalogOptions) SetAuthInstanceID(authInstanceID string) *DeleteDatabaseCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDatabaseCatalogOptions) SetHeaders(param map[string]string) *DeleteDatabaseCatalogOptions {
	options.Headers = param
	return options
}

// DeleteDb2EngineOptions : The DeleteDb2Engine options.
type DeleteDb2EngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDb2EngineOptions : Instantiate DeleteDb2EngineOptions
func (*WatsonxDataV2) NewDeleteDb2EngineOptions(engineID string) *DeleteDb2EngineOptions {
	return &DeleteDb2EngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteDb2EngineOptions) SetEngineID(engineID string) *DeleteDb2EngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDb2EngineOptions) SetHeaders(param map[string]string) *DeleteDb2EngineOptions {
	options.Headers = param
	return options
}

// DeleteDeactivateBucketOptions : The DeleteDeactivateBucket options.
type DeleteDeactivateBucketOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDeactivateBucketOptions : Instantiate DeleteDeactivateBucketOptions
func (*WatsonxDataV2) NewDeleteDeactivateBucketOptions(bucketID string) *DeleteDeactivateBucketOptions {
	return &DeleteDeactivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeleteDeactivateBucketOptions) SetBucketID(bucketID string) *DeleteDeactivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDeactivateBucketOptions) SetAuthInstanceID(authInstanceID string) *DeleteDeactivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDeactivateBucketOptions) SetHeaders(param map[string]string) *DeleteDeactivateBucketOptions {
	options.Headers = param
	return options
}

// DeleteEngineOptions : The DeleteEngine options.
type DeleteEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEngineOptions : Instantiate DeleteEngineOptions
func (*WatsonxDataV2) NewDeleteEngineOptions(engineID string) *DeleteEngineOptions {
	return &DeleteEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteEngineOptions) SetEngineID(engineID string) *DeleteEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEngineOptions) SetHeaders(param map[string]string) *DeleteEngineOptions {
	options.Headers = param
	return options
}

// DeleteMilvusServiceOptions : The DeleteMilvusService options.
type DeleteMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteMilvusServiceOptions : Instantiate DeleteMilvusServiceOptions
func (*WatsonxDataV2) NewDeleteMilvusServiceOptions(serviceID string) *DeleteMilvusServiceOptions {
	return &DeleteMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *DeleteMilvusServiceOptions) SetServiceID(serviceID string) *DeleteMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *DeleteMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMilvusServiceOptions) SetHeaders(param map[string]string) *DeleteMilvusServiceOptions {
	options.Headers = param
	return options
}

// DeleteNetezzaEngineOptions : The DeleteNetezzaEngine options.
type DeleteNetezzaEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNetezzaEngineOptions : Instantiate DeleteNetezzaEngineOptions
func (*WatsonxDataV2) NewDeleteNetezzaEngineOptions(engineID string) *DeleteNetezzaEngineOptions {
	return &DeleteNetezzaEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteNetezzaEngineOptions) SetEngineID(engineID string) *DeleteNetezzaEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNetezzaEngineOptions) SetHeaders(param map[string]string) *DeleteNetezzaEngineOptions {
	options.Headers = param
	return options
}

// DeleteOtherEngineOptions : The DeleteOtherEngine options.
type DeleteOtherEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteOtherEngineOptions : Instantiate DeleteOtherEngineOptions
func (*WatsonxDataV2) NewDeleteOtherEngineOptions(engineID string) *DeleteOtherEngineOptions {
	return &DeleteOtherEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteOtherEngineOptions) SetEngineID(engineID string) *DeleteOtherEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteOtherEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteOtherEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteOtherEngineOptions) SetHeaders(param map[string]string) *DeleteOtherEngineOptions {
	options.Headers = param
	return options
}

// DeletePrestissimoEngineCatalogsOptions : The DeletePrestissimoEngineCatalogs options.
type DeletePrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Catalog id(s) to be stopped, comma separated.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePrestissimoEngineCatalogsOptions : Instantiate DeletePrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewDeletePrestissimoEngineCatalogsOptions(engineID string, catalogNames string) *DeletePrestissimoEngineCatalogsOptions {
	return &DeletePrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *DeletePrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *DeletePrestissimoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *DeletePrestissimoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *DeletePrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// DeletePrestissimoEngineOptions : The DeletePrestissimoEngine options.
type DeletePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePrestissimoEngineOptions : Instantiate DeletePrestissimoEngineOptions
func (*WatsonxDataV2) NewDeletePrestissimoEngineOptions(engineID string) *DeletePrestissimoEngineOptions {
	return &DeletePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestissimoEngineOptions) SetEngineID(engineID string) *DeletePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestissimoEngineOptions) SetHeaders(param map[string]string) *DeletePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// DeletePrestoEngineCatalogsOptions : The DeletePrestoEngineCatalogs options.
type DeletePrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Catalog id(s) to be stopped, comma separated.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePrestoEngineCatalogsOptions : Instantiate DeletePrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewDeletePrestoEngineCatalogsOptions(engineID string, catalogNames string) *DeletePrestoEngineCatalogsOptions {
	return &DeletePrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestoEngineCatalogsOptions) SetEngineID(engineID string) *DeletePrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *DeletePrestoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *DeletePrestoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *DeletePrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// DeleteSchemaOptions : The DeleteSchema options.
type DeleteSchemaOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded Schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSchemaOptions : Instantiate DeleteSchemaOptions
func (*WatsonxDataV2) NewDeleteSchemaOptions(engineID string, catalogID string, schemaID string) *DeleteSchemaOptions {
	return &DeleteSchemaOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSchemaOptions) SetEngineID(engineID string) *DeleteSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteSchemaOptions) SetCatalogID(catalogID string) *DeleteSchemaOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteSchemaOptions) SetSchemaID(schemaID string) *DeleteSchemaOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSchemaOptions) SetAuthInstanceID(authInstanceID string) *DeleteSchemaOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSchemaOptions) SetHeaders(param map[string]string) *DeleteSchemaOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineApplicationsOptions : The DeleteSparkEngineApplications options.
type DeleteSparkEngineApplicationsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application id(s) to be stopped, comma separated.
	ApplicationID *string `json:"application_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSparkEngineApplicationsOptions : Instantiate DeleteSparkEngineApplicationsOptions
func (*WatsonxDataV2) NewDeleteSparkEngineApplicationsOptions(engineID string, applicationID string) *DeleteSparkEngineApplicationsOptions {
	return &DeleteSparkEngineApplicationsOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineApplicationsOptions) SetEngineID(engineID string) *DeleteSparkEngineApplicationsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *DeleteSparkEngineApplicationsOptions) SetApplicationID(applicationID string) *DeleteSparkEngineApplicationsOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineApplicationsOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineApplicationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineApplicationsOptions) SetHeaders(param map[string]string) *DeleteSparkEngineApplicationsOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineOptions : The DeleteSparkEngine options.
type DeleteSparkEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSparkEngineOptions : Instantiate DeleteSparkEngineOptions
func (*WatsonxDataV2) NewDeleteSparkEngineOptions(engineID string) *DeleteSparkEngineOptions {
	return &DeleteSparkEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineOptions) SetEngineID(engineID string) *DeleteSparkEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineOptions) SetHeaders(param map[string]string) *DeleteSparkEngineOptions {
	options.Headers = param
	return options
}

// DeleteTableOptions : The DeleteTable options.
type DeleteTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTableOptions : Instantiate DeleteTableOptions
func (*WatsonxDataV2) NewDeleteTableOptions(catalogID string, schemaID string, tableID string, engineID string) *DeleteTableOptions {
	return &DeleteTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteTableOptions) SetCatalogID(catalogID string) *DeleteTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteTableOptions) SetSchemaID(schemaID string) *DeleteTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *DeleteTableOptions) SetTableID(tableID string) *DeleteTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteTableOptions) SetEngineID(engineID string) *DeleteTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteTableOptions) SetAuthInstanceID(authInstanceID string) *DeleteTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTableOptions) SetHeaders(param map[string]string) *DeleteTableOptions {
	options.Headers = param
	return options
}

// Deployment : Deployment.
type Deployment struct {
	// Cloud type.
	CloudType *string `json:"cloud_type,omitempty"`

	// Enable private endpoints.
	EnablePrivateEndpoints *bool `json:"enable_private_endpoints,omitempty"`

	// Enable public endpoints.
	EnablePublicEndpoints *bool `json:"enable_public_endpoints,omitempty"`

	// Parameter for UI to validate if console is used for the first time.
	FirstTimeUse *bool `json:"first_time_use" validate:"required"`

	// Formation id.
	FormationID *string `json:"formation_id,omitempty"`

	// Id.
	ID *string `json:"id,omitempty"`

	// Plan id.
	PlanID *string `json:"plan_id,omitempty"`

	// Platform options.
	PlatformOptions *DeploymentPlatformOptions `json:"platform_options,omitempty"`

	// Region.
	Region *string `json:"region,omitempty"`

	// Resource group crn for the formation.
	ResourceGroupCrn *string `json:"resource_group_crn,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// Version.
	Version *string `json:"version,omitempty"`
}

// UnmarshalDeployment unmarshals an instance of Deployment from the specified map of raw messages.
func UnmarshalDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployment)
	err = core.UnmarshalPrimitive(m, "cloud_type", &obj.CloudType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_private_endpoints", &obj.EnablePrivateEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_public_endpoints", &obj.EnablePublicEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first_time_use", &obj.FirstTimeUse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "formation_id", &obj.FormationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "platform_options", &obj.PlatformOptions, UnmarshalDeploymentPlatformOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_crn", &obj.ResourceGroupCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentPlatformOptions : Platform options.
type DeploymentPlatformOptions struct {
	// Backup encryption key crn.
	BackupEncryptionKeyCrn *string `json:"backup_encryption_key_crn,omitempty"`

	// Disk encryption key crn.
	DiskEncryptionKeyCrn *string `json:"disk_encryption_key_crn,omitempty"`

	// Key protect key id.
	KeyProtectKeyID *string `json:"key_protect_key_id,omitempty"`
}

// UnmarshalDeploymentPlatformOptions unmarshals an instance of DeploymentPlatformOptions from the specified map of raw messages.
func UnmarshalDeploymentPlatformOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentPlatformOptions)
	err = core.UnmarshalPrimitive(m, "backup_encryption_key_crn", &obj.BackupEncryptionKeyCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disk_encryption_key_crn", &obj.DiskEncryptionKeyCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_protect_key_id", &obj.KeyProtectKeyID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentsResponse : DeploymentsResponse.
type DeploymentsResponse struct {
	// Deployment.
	Deployment *Deployment `json:"deployment,omitempty"`
}

// UnmarshalDeploymentsResponse unmarshals an instance of DeploymentsResponse from the specified map of raw messages.
func UnmarshalDeploymentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentsResponse)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalDeployment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Engine : All engine details.
type Engine struct {
	// list of db2 engines.
	Db2Engines []Db2Engine `json:"db2_engines,omitempty"`

	// list of milvus engines.
	MilvusServices []MilvusService `json:"milvus_services,omitempty"`

	// list of netezza engines.
	NetezzaEngines []NetezzaEngine `json:"netezza_engines,omitempty"`

	// list of prestissimo engines.
	PrestissimoEngines []PrestissimoEngine `json:"prestissimo_engines,omitempty"`

	// list of presto engines.
	PrestoEngines []PrestoEngine `json:"presto_engines,omitempty"`

	// list of spark engines.
	SparkEngines []SparkEngine `json:"spark_engines,omitempty"`
}

// UnmarshalEngine unmarshals an instance of Engine from the specified map of raw messages.
func UnmarshalEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Engine)
	err = core.UnmarshalModel(m, "db2_engines", &obj.Db2Engines, UnmarshalDb2Engine)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "milvus_services", &obj.MilvusServices, UnmarshalMilvusService)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "netezza_engines", &obj.NetezzaEngines, UnmarshalNetezzaEngine)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prestissimo_engines", &obj.PrestissimoEngines, UnmarshalPrestissimoEngine)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "presto_engines", &obj.PrestoEngines, UnmarshalPrestoEngine)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spark_engines", &obj.SparkEngines, UnmarshalSparkEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineDetailsBody : Node details.
type EngineDetailsBody struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Node details.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Node details.
	Worker *NodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the EngineDetailsBody.SizeConfig property.
// Size config.
const (
	EngineDetailsBody_SizeConfig_CacheOptimized = "cache_optimized"
	EngineDetailsBody_SizeConfig_ComputeOptimized = "compute_optimized"
	EngineDetailsBody_SizeConfig_Custom = "custom"
	EngineDetailsBody_SizeConfig_Large = "large"
	EngineDetailsBody_SizeConfig_Medium = "medium"
	EngineDetailsBody_SizeConfig_Small = "small"
	EngineDetailsBody_SizeConfig_Starter = "starter"
)

// UnmarshalEngineDetailsBody unmarshals an instance of EngineDetailsBody from the specified map of raw messages.
func UnmarshalEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescriptionBody)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Engines : List all engines.
type Engines struct {
	// All engine details.
	Engines *Engine `json:"engines,omitempty"`
}

// UnmarshalEngines unmarshals an instance of Engines from the specified map of raw messages.
func UnmarshalEngines(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Engines)
	err = core.UnmarshalModel(m, "engines", &obj.Engines, UnmarshalEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketRegistrationOptions : The GetBucketRegistration options.
type GetBucketRegistrationOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketRegistrationOptions : Instantiate GetBucketRegistrationOptions
func (*WatsonxDataV2) NewGetBucketRegistrationOptions(bucketID string) *GetBucketRegistrationOptions {
	return &GetBucketRegistrationOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *GetBucketRegistrationOptions) SetBucketID(bucketID string) *GetBucketRegistrationOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *GetBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketRegistrationOptions) SetHeaders(param map[string]string) *GetBucketRegistrationOptions {
	options.Headers = param
	return options
}

// GetCatalogOptions : The GetCatalog options.
type GetCatalogOptions struct {
	// catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogOptions : Instantiate GetCatalogOptions
func (*WatsonxDataV2) NewGetCatalogOptions(catalogID string) *GetCatalogOptions {
	return &GetCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetCatalogOptions) SetCatalogID(catalogID string) *GetCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogOptions) SetHeaders(param map[string]string) *GetCatalogOptions {
	options.Headers = param
	return options
}

// GetDatabaseOptions : The GetDatabase options.
type GetDatabaseOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDatabaseOptions : Instantiate GetDatabaseOptions
func (*WatsonxDataV2) NewGetDatabaseOptions(databaseID string) *GetDatabaseOptions {
	return &GetDatabaseOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *GetDatabaseOptions) SetDatabaseID(databaseID string) *GetDatabaseOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDatabaseOptions) SetAuthInstanceID(authInstanceID string) *GetDatabaseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDatabaseOptions) SetHeaders(param map[string]string) *GetDatabaseOptions {
	options.Headers = param
	return options
}

// GetDeploymentsOKBody : Response body structure for get deployments.
type GetDeploymentsOKBody struct {
	// DeploymentsResponse.
	Deploymentresponse *DeploymentsResponse `json:"deploymentresponse" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetDeploymentsOKBody unmarshals an instance of GetDeploymentsOKBody from the specified map of raw messages.
func UnmarshalGetDeploymentsOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDeploymentsOKBody)
	err = core.UnmarshalModel(m, "deploymentresponse", &obj.Deploymentresponse, UnmarshalDeploymentsResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDeploymentsOptions : The GetDeployments options.
type GetDeploymentsOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentsOptions : Instantiate GetDeploymentsOptions
func (*WatsonxDataV2) NewGetDeploymentsOptions() *GetDeploymentsOptions {
	return &GetDeploymentsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDeploymentsOptions) SetAuthInstanceID(authInstanceID string) *GetDeploymentsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentsOptions) SetHeaders(param map[string]string) *GetDeploymentsOptions {
	options.Headers = param
	return options
}

// GetEnginesOptions : The GetEngines options.
type GetEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnginesOptions : Instantiate GetEnginesOptions
func (*WatsonxDataV2) NewGetEnginesOptions() *GetEnginesOptions {
	return &GetEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetEnginesOptions) SetAuthInstanceID(authInstanceID string) *GetEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnginesOptions) SetHeaders(param map[string]string) *GetEnginesOptions {
	options.Headers = param
	return options
}

// GetMilvusServiceOptions : The GetMilvusService options.
type GetMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMilvusServiceOptions : Instantiate GetMilvusServiceOptions
func (*WatsonxDataV2) NewGetMilvusServiceOptions(serviceID string) *GetMilvusServiceOptions {
	return &GetMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *GetMilvusServiceOptions) SetServiceID(serviceID string) *GetMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *GetMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMilvusServiceOptions) SetHeaders(param map[string]string) *GetMilvusServiceOptions {
	options.Headers = param
	return options
}

// GetPrestissimoEngineCatalogOptions : The GetPrestissimoEngineCatalog options.
type GetPrestissimoEngineCatalogOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPrestissimoEngineCatalogOptions : Instantiate GetPrestissimoEngineCatalogOptions
func (*WatsonxDataV2) NewGetPrestissimoEngineCatalogOptions(engineID string, catalogID string) *GetPrestissimoEngineCatalogOptions {
	return &GetPrestissimoEngineCatalogOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestissimoEngineCatalogOptions) SetEngineID(engineID string) *GetPrestissimoEngineCatalogOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetPrestissimoEngineCatalogOptions) SetCatalogID(catalogID string) *GetPrestissimoEngineCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestissimoEngineCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetPrestissimoEngineCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestissimoEngineCatalogOptions) SetHeaders(param map[string]string) *GetPrestissimoEngineCatalogOptions {
	options.Headers = param
	return options
}

// GetPrestissimoEngineOptions : The GetPrestissimoEngine options.
type GetPrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPrestissimoEngineOptions : Instantiate GetPrestissimoEngineOptions
func (*WatsonxDataV2) NewGetPrestissimoEngineOptions(engineID string) *GetPrestissimoEngineOptions {
	return &GetPrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestissimoEngineOptions) SetEngineID(engineID string) *GetPrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *GetPrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestissimoEngineOptions) SetHeaders(param map[string]string) *GetPrestissimoEngineOptions {
	options.Headers = param
	return options
}

// GetPrestoEngineCatalogOptions : The GetPrestoEngineCatalog options.
type GetPrestoEngineCatalogOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPrestoEngineCatalogOptions : Instantiate GetPrestoEngineCatalogOptions
func (*WatsonxDataV2) NewGetPrestoEngineCatalogOptions(engineID string, catalogID string) *GetPrestoEngineCatalogOptions {
	return &GetPrestoEngineCatalogOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestoEngineCatalogOptions) SetEngineID(engineID string) *GetPrestoEngineCatalogOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetPrestoEngineCatalogOptions) SetCatalogID(catalogID string) *GetPrestoEngineCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestoEngineCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetPrestoEngineCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestoEngineCatalogOptions) SetHeaders(param map[string]string) *GetPrestoEngineCatalogOptions {
	options.Headers = param
	return options
}

// GetPrestoEngineOptions : The GetPrestoEngine options.
type GetPrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPrestoEngineOptions : Instantiate GetPrestoEngineOptions
func (*WatsonxDataV2) NewGetPrestoEngineOptions(engineID string) *GetPrestoEngineOptions {
	return &GetPrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestoEngineOptions) SetEngineID(engineID string) *GetPrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *GetPrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestoEngineOptions) SetHeaders(param map[string]string) *GetPrestoEngineOptions {
	options.Headers = param
	return options
}

// GetSparkEngineApplicationStatusOptions : The GetSparkEngineApplicationStatus options.
type GetSparkEngineApplicationStatusOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application id.
	ApplicationID *string `json:"application_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSparkEngineApplicationStatusOptions : Instantiate GetSparkEngineApplicationStatusOptions
func (*WatsonxDataV2) NewGetSparkEngineApplicationStatusOptions(engineID string, applicationID string) *GetSparkEngineApplicationStatusOptions {
	return &GetSparkEngineApplicationStatusOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSparkEngineApplicationStatusOptions) SetEngineID(engineID string) *GetSparkEngineApplicationStatusOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *GetSparkEngineApplicationStatusOptions) SetApplicationID(applicationID string) *GetSparkEngineApplicationStatusOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSparkEngineApplicationStatusOptions) SetAuthInstanceID(authInstanceID string) *GetSparkEngineApplicationStatusOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkEngineApplicationStatusOptions) SetHeaders(param map[string]string) *GetSparkEngineApplicationStatusOptions {
	options.Headers = param
	return options
}

// GetTableOptions : The GetTable options.
type GetTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTableOptions : Instantiate GetTableOptions
func (*WatsonxDataV2) NewGetTableOptions(catalogID string, schemaID string, tableID string, engineID string) *GetTableOptions {
	return &GetTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetTableOptions) SetCatalogID(catalogID string) *GetTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *GetTableOptions) SetSchemaID(schemaID string) *GetTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *GetTableOptions) SetTableID(tableID string) *GetTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *GetTableOptions) SetEngineID(engineID string) *GetTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetTableOptions) SetAuthInstanceID(authInstanceID string) *GetTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTableOptions) SetHeaders(param map[string]string) *GetTableOptions {
	options.Headers = param
	return options
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add = "add"
	JSONPatchOperation_Op_Copy = "copy"
	JSONPatchOperation_Op_Move = "move"
	JSONPatchOperation_Op_Remove = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*WatsonxDataV2) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBucketObjectsOptions : The ListBucketObjects options.
type ListBucketObjectsOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBucketObjectsOptions : Instantiate ListBucketObjectsOptions
func (*WatsonxDataV2) NewListBucketObjectsOptions(bucketID string) *ListBucketObjectsOptions {
	return &ListBucketObjectsOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *ListBucketObjectsOptions) SetBucketID(bucketID string) *ListBucketObjectsOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListBucketObjectsOptions) SetAuthInstanceID(authInstanceID string) *ListBucketObjectsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBucketObjectsOptions) SetHeaders(param map[string]string) *ListBucketObjectsOptions {
	options.Headers = param
	return options
}

// ListBucketRegistrationsOptions : The ListBucketRegistrations options.
type ListBucketRegistrationsOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBucketRegistrationsOptions : Instantiate ListBucketRegistrationsOptions
func (*WatsonxDataV2) NewListBucketRegistrationsOptions() *ListBucketRegistrationsOptions {
	return &ListBucketRegistrationsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListBucketRegistrationsOptions) SetAuthInstanceID(authInstanceID string) *ListBucketRegistrationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBucketRegistrationsOptions) SetHeaders(param map[string]string) *ListBucketRegistrationsOptions {
	options.Headers = param
	return options
}

// ListCatalogsOptions : The ListCatalogs options.
type ListCatalogsOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCatalogsOptions : Instantiate ListCatalogsOptions
func (*WatsonxDataV2) NewListCatalogsOptions() *ListCatalogsOptions {
	return &ListCatalogsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogsOptions) SetHeaders(param map[string]string) *ListCatalogsOptions {
	options.Headers = param
	return options
}

// ListColumnsOptions : The ListColumns options.
type ListColumnsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListColumnsOptions : Instantiate ListColumnsOptions
func (*WatsonxDataV2) NewListColumnsOptions(engineID string, catalogID string, schemaID string, tableID string) *ListColumnsOptions {
	return &ListColumnsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListColumnsOptions) SetEngineID(engineID string) *ListColumnsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListColumnsOptions) SetCatalogID(catalogID string) *ListColumnsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListColumnsOptions) SetSchemaID(schemaID string) *ListColumnsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *ListColumnsOptions) SetTableID(tableID string) *ListColumnsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListColumnsOptions) SetAuthInstanceID(authInstanceID string) *ListColumnsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListColumnsOptions) SetHeaders(param map[string]string) *ListColumnsOptions {
	options.Headers = param
	return options
}

// ListDatabaseRegistrationsOptions : The ListDatabaseRegistrations options.
type ListDatabaseRegistrationsOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDatabaseRegistrationsOptions : Instantiate ListDatabaseRegistrationsOptions
func (*WatsonxDataV2) NewListDatabaseRegistrationsOptions() *ListDatabaseRegistrationsOptions {
	return &ListDatabaseRegistrationsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListDatabaseRegistrationsOptions) SetAuthInstanceID(authInstanceID string) *ListDatabaseRegistrationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDatabaseRegistrationsOptions) SetHeaders(param map[string]string) *ListDatabaseRegistrationsOptions {
	options.Headers = param
	return options
}

// ListDb2EnginesOptions : The ListDb2Engines options.
type ListDb2EnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDb2EnginesOptions : Instantiate ListDb2EnginesOptions
func (*WatsonxDataV2) NewListDb2EnginesOptions() *ListDb2EnginesOptions {
	return &ListDb2EnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListDb2EnginesOptions) SetAuthInstanceID(authInstanceID string) *ListDb2EnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDb2EnginesOptions) SetHeaders(param map[string]string) *ListDb2EnginesOptions {
	options.Headers = param
	return options
}

// ListMilvusServicesOptions : The ListMilvusServices options.
type ListMilvusServicesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListMilvusServicesOptions : Instantiate ListMilvusServicesOptions
func (*WatsonxDataV2) NewListMilvusServicesOptions() *ListMilvusServicesOptions {
	return &ListMilvusServicesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListMilvusServicesOptions) SetAuthInstanceID(authInstanceID string) *ListMilvusServicesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListMilvusServicesOptions) SetHeaders(param map[string]string) *ListMilvusServicesOptions {
	options.Headers = param
	return options
}

// ListNetezzaEnginesOptions : The ListNetezzaEngines options.
type ListNetezzaEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListNetezzaEnginesOptions : Instantiate ListNetezzaEnginesOptions
func (*WatsonxDataV2) NewListNetezzaEnginesOptions() *ListNetezzaEnginesOptions {
	return &ListNetezzaEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListNetezzaEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListNetezzaEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListNetezzaEnginesOptions) SetHeaders(param map[string]string) *ListNetezzaEnginesOptions {
	options.Headers = param
	return options
}

// ListOtherEnginesOptions : The ListOtherEngines options.
type ListOtherEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListOtherEnginesOptions : Instantiate ListOtherEnginesOptions
func (*WatsonxDataV2) NewListOtherEnginesOptions() *ListOtherEnginesOptions {
	return &ListOtherEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListOtherEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListOtherEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListOtherEnginesOptions) SetHeaders(param map[string]string) *ListOtherEnginesOptions {
	options.Headers = param
	return options
}

// ListPrestissimoEngineCatalogsOptions : The ListPrestissimoEngineCatalogs options.
type ListPrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPrestissimoEngineCatalogsOptions : Instantiate ListPrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewListPrestissimoEngineCatalogsOptions(engineID string) *ListPrestissimoEngineCatalogsOptions {
	return &ListPrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListPrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *ListPrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListPrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *ListPrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ListPrestissimoEnginesOptions : The ListPrestissimoEngines options.
type ListPrestissimoEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPrestissimoEnginesOptions : Instantiate ListPrestissimoEnginesOptions
func (*WatsonxDataV2) NewListPrestissimoEnginesOptions() *ListPrestissimoEnginesOptions {
	return &ListPrestissimoEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestissimoEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListPrestissimoEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestissimoEnginesOptions) SetHeaders(param map[string]string) *ListPrestissimoEnginesOptions {
	options.Headers = param
	return options
}

// ListPrestoEngineCatalogsOptions : The ListPrestoEngineCatalogs options.
type ListPrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPrestoEngineCatalogsOptions : Instantiate ListPrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewListPrestoEngineCatalogsOptions(engineID string) *ListPrestoEngineCatalogsOptions {
	return &ListPrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListPrestoEngineCatalogsOptions) SetEngineID(engineID string) *ListPrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListPrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *ListPrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ListPrestoEnginesOptions : The ListPrestoEngines options.
type ListPrestoEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPrestoEnginesOptions : Instantiate ListPrestoEnginesOptions
func (*WatsonxDataV2) NewListPrestoEnginesOptions() *ListPrestoEnginesOptions {
	return &ListPrestoEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestoEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListPrestoEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestoEnginesOptions) SetHeaders(param map[string]string) *ListPrestoEnginesOptions {
	options.Headers = param
	return options
}

// ListSchemasOKBody : GetSchemas OK.
type ListSchemasOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Schemas.
	Schemas []string `json:"schemas" validate:"required"`
}

// UnmarshalListSchemasOKBody unmarshals an instance of ListSchemasOKBody from the specified map of raw messages.
func UnmarshalListSchemasOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListSchemasOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schemas", &obj.Schemas)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListSchemasOptions : The ListSchemas options.
type ListSchemasOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSchemasOptions : Instantiate ListSchemasOptions
func (*WatsonxDataV2) NewListSchemasOptions(engineID string, catalogID string) *ListSchemasOptions {
	return &ListSchemasOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListSchemasOptions) SetEngineID(engineID string) *ListSchemasOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListSchemasOptions) SetCatalogID(catalogID string) *ListSchemasOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSchemasOptions) SetAuthInstanceID(authInstanceID string) *ListSchemasOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSchemasOptions) SetHeaders(param map[string]string) *ListSchemasOptions {
	options.Headers = param
	return options
}

// ListSparkEngineApplicationsOptions : The ListSparkEngineApplications options.
type ListSparkEngineApplicationsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSparkEngineApplicationsOptions : Instantiate ListSparkEngineApplicationsOptions
func (*WatsonxDataV2) NewListSparkEngineApplicationsOptions(engineID string) *ListSparkEngineApplicationsOptions {
	return &ListSparkEngineApplicationsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListSparkEngineApplicationsOptions) SetEngineID(engineID string) *ListSparkEngineApplicationsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkEngineApplicationsOptions) SetAuthInstanceID(authInstanceID string) *ListSparkEngineApplicationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkEngineApplicationsOptions) SetHeaders(param map[string]string) *ListSparkEngineApplicationsOptions {
	options.Headers = param
	return options
}

// ListSparkEnginesOptions : The ListSparkEngines options.
type ListSparkEnginesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSparkEnginesOptions : Instantiate ListSparkEnginesOptions
func (*WatsonxDataV2) NewListSparkEnginesOptions() *ListSparkEnginesOptions {
	return &ListSparkEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListSparkEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkEnginesOptions) SetHeaders(param map[string]string) *ListSparkEnginesOptions {
	options.Headers = param
	return options
}

// ListTableSnapshotsOptions : The ListTableSnapshots options.
type ListTableSnapshotsOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Schema ID.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// Table ID.
	TableID *string `json:"table_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTableSnapshotsOptions : Instantiate ListTableSnapshotsOptions
func (*WatsonxDataV2) NewListTableSnapshotsOptions(engineID string, catalogID string, schemaID string, tableID string) *ListTableSnapshotsOptions {
	return &ListTableSnapshotsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListTableSnapshotsOptions) SetEngineID(engineID string) *ListTableSnapshotsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListTableSnapshotsOptions) SetCatalogID(catalogID string) *ListTableSnapshotsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListTableSnapshotsOptions) SetSchemaID(schemaID string) *ListTableSnapshotsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *ListTableSnapshotsOptions) SetTableID(tableID string) *ListTableSnapshotsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListTableSnapshotsOptions) SetAuthInstanceID(authInstanceID string) *ListTableSnapshotsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTableSnapshotsOptions) SetHeaders(param map[string]string) *ListTableSnapshotsOptions {
	options.Headers = param
	return options
}

// ListTablesOptions : The ListTables options.
type ListTablesOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTablesOptions : Instantiate ListTablesOptions
func (*WatsonxDataV2) NewListTablesOptions(catalogID string, schemaID string, engineID string) *ListTablesOptions {
	return &ListTablesOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListTablesOptions) SetCatalogID(catalogID string) *ListTablesOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListTablesOptions) SetSchemaID(schemaID string) *ListTablesOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *ListTablesOptions) SetEngineID(engineID string) *ListTablesOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListTablesOptions) SetAuthInstanceID(authInstanceID string) *ListTablesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTablesOptions) SetHeaders(param map[string]string) *ListTablesOptions {
	options.Headers = param
	return options
}

// MilvusService : milvus service details.
type MilvusService struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Service description.
	Description *string `json:"description,omitempty"`

	// milvus grpc_host.
	GrpcHost *string `json:"grpc_host,omitempty"`

	// milvus port.
	GrpcPort *int64 `json:"grpc_port,omitempty"`

	// milvus display name.
	HostName *string `json:"host_name,omitempty"`

	// milvus https_host.
	HttpsHost *string `json:"https_host,omitempty"`

	// milvus port.
	HttpsPort *int64 `json:"https_port,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Service programmatic name.
	ServiceID *string `json:"service_id,omitempty"`

	// milvus status.
	Status *string `json:"status,omitempty"`

	// milvus status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// service type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the MilvusService.Status property.
// milvus status.
const (
	MilvusService_Status_Pending = "pending"
	MilvusService_Status_Running = "running"
	MilvusService_Status_Stopped = "stopped"
)

// UnmarshalMilvusService unmarshals an instance of MilvusService from the specified map of raw messages.
func UnmarshalMilvusService(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MilvusService)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "grpc_host", &obj.GrpcHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "grpc_port", &obj.GrpcPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "https_host", &obj.HttpsHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "https_port", &obj.HttpsPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_display_name", &obj.ServiceDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_id", &obj.ServiceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewMilvusServicePatch(milvusService *MilvusService) (_patch []JSONPatchOperation) {
	if (milvusService.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: milvusService.Actions,
		})
	}
	if (milvusService.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: milvusService.CreatedBy,
		})
	}
	if (milvusService.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: milvusService.CreatedOn,
		})
	}
	if (milvusService.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: milvusService.Description,
		})
	}
	if (milvusService.GrpcHost != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/grpc_host"),
			Value: milvusService.GrpcHost,
		})
	}
	if (milvusService.GrpcPort != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/grpc_port"),
			Value: milvusService.GrpcPort,
		})
	}
	if (milvusService.HostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/host_name"),
			Value: milvusService.HostName,
		})
	}
	if (milvusService.HttpsHost != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/https_host"),
			Value: milvusService.HttpsHost,
		})
	}
	if (milvusService.HttpsPort != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/https_port"),
			Value: milvusService.HttpsPort,
		})
	}
	if (milvusService.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: milvusService.Origin,
		})
	}
	if (milvusService.ServiceDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/service_display_name"),
			Value: milvusService.ServiceDisplayName,
		})
	}
	if (milvusService.ServiceID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/service_id"),
			Value: milvusService.ServiceID,
		})
	}
	if (milvusService.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: milvusService.Status,
		})
	}
	if (milvusService.StatusCode != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status_code"),
			Value: milvusService.StatusCode,
		})
	}
	if (milvusService.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: milvusService.Tags,
		})
	}
	if (milvusService.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: milvusService.Type,
		})
	}
	return
}

// MilvusServiceCollection : List milvus services.
type MilvusServiceCollection struct {
	// milvus service body.
	MilvusServices []MilvusService `json:"milvus_services,omitempty"`
}

// UnmarshalMilvusServiceCollection unmarshals an instance of MilvusServiceCollection from the specified map of raw messages.
func UnmarshalMilvusServiceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MilvusServiceCollection)
	err = core.UnmarshalModel(m, "milvus_services", &obj.MilvusServices, UnmarshalMilvusService)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngine : Netezza engine details.
type NetezzaEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *NetezzaEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalNetezzaEngine unmarshals an instance of NetezzaEngine from the specified map of raw messages.
func UnmarshalNetezzaEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalNetezzaEngineDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewNetezzaEnginePatch(netezzaEngine *NetezzaEngine) (_patch []JSONPatchOperation) {
	if (netezzaEngine.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: netezzaEngine.Actions,
		})
	}
	if (netezzaEngine.BuildVersion != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/build_version"),
			Value: netezzaEngine.BuildVersion,
		})
	}
	if (netezzaEngine.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: netezzaEngine.CreatedBy,
		})
	}
	if (netezzaEngine.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: netezzaEngine.CreatedOn,
		})
	}
	if (netezzaEngine.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: netezzaEngine.Description,
		})
	}
	if (netezzaEngine.EngineDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_details"),
			Value: netezzaEngine.EngineDetails,
		})
	}
	if (netezzaEngine.EngineDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_display_name"),
			Value: netezzaEngine.EngineDisplayName,
		})
	}
	if (netezzaEngine.EngineID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_id"),
			Value: netezzaEngine.EngineID,
		})
	}
	if (netezzaEngine.HostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/host_name"),
			Value: netezzaEngine.HostName,
		})
	}
	if (netezzaEngine.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: netezzaEngine.Origin,
		})
	}
	if (netezzaEngine.Port != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/port"),
			Value: netezzaEngine.Port,
		})
	}
	if (netezzaEngine.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: netezzaEngine.Status,
		})
	}
	if (netezzaEngine.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: netezzaEngine.Tags,
		})
	}
	if (netezzaEngine.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: netezzaEngine.Type,
		})
	}
	return
}

// NetezzaEngineCollection : list Netezza engines.
type NetezzaEngineCollection struct {
	// list Netezza engines.
	NetezzaEngines []NetezzaEngine `json:"netezza_engines,omitempty"`
}

// UnmarshalNetezzaEngineCollection unmarshals an instance of NetezzaEngineCollection from the specified map of raw messages.
func UnmarshalNetezzaEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineCollection)
	err = core.UnmarshalModel(m, "netezza_engines", &obj.NetezzaEngines, UnmarshalNetezzaEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngineDetails : External engine details.
type NetezzaEngineDetails struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalNetezzaEngineDetails unmarshals an instance of NetezzaEngineDetails from the specified map of raw messages.
func UnmarshalNetezzaEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngineDetailsBody : External engine details.
type NetezzaEngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`
}

// UnmarshalNetezzaEngineDetailsBody unmarshals an instance of NetezzaEngineDetailsBody from the specified map of raw messages.
func UnmarshalNetezzaEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NodeDescription : NodeDescription.
type NodeDescription struct {
	// Node type.
	NodeType *string `json:"node_type,omitempty"`

	// Quantity.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalNodeDescription unmarshals an instance of NodeDescription from the specified map of raw messages.
func UnmarshalNodeDescription(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeDescription)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NodeDescriptionBody : Node details.
type NodeDescriptionBody struct {
	// Node Type, r5, m, i..
	NodeType *string `json:"node_type,omitempty"`

	// Number of nodes.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalNodeDescriptionBody unmarshals an instance of NodeDescriptionBody from the specified map of raw messages.
func UnmarshalNodeDescriptionBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeDescriptionBody)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngine : external engine details.
type OtherEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *OtherEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// origin.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Type like presto, netezza, external,..
	Type *string `json:"type,omitempty"`
}

// UnmarshalOtherEngine unmarshals an instance of OtherEngine from the specified map of raw messages.
func UnmarshalOtherEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalOtherEngineDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineCollection : list other engines.
type OtherEngineCollection struct {
	// list other engines.
	OtherEngines []OtherEngine `json:"other_engines,omitempty"`
}

// UnmarshalOtherEngineCollection unmarshals an instance of OtherEngineCollection from the specified map of raw messages.
func UnmarshalOtherEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineCollection)
	err = core.UnmarshalModel(m, "other_engines", &obj.OtherEngines, UnmarshalOtherEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineDetails : External engine details.
type OtherEngineDetails struct {
	// external engine connection string.
	ConnectionString *string `json:"connection_string" validate:"required"`

	// Actual engine type.
	EngineType *string `json:"engine_type" validate:"required"`

	// metastore host - not required while registering an engine.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalOtherEngineDetails unmarshals an instance of OtherEngineDetails from the specified map of raw messages.
func UnmarshalOtherEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_type", &obj.EngineType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineDetailsBody : External engine details.
type OtherEngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string" validate:"required"`

	// Actual engine type.
	EngineType *string `json:"engine_type" validate:"required"`
}

// NewOtherEngineDetailsBody : Instantiate OtherEngineDetailsBody (Generic Model Constructor)
func (*WatsonxDataV2) NewOtherEngineDetailsBody(connectionString string, engineType string) (_model *OtherEngineDetailsBody, err error) {
	_model = &OtherEngineDetailsBody{
		ConnectionString: core.StringPtr(connectionString),
		EngineType: core.StringPtr(engineType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOtherEngineDetailsBody unmarshals an instance of OtherEngineDetailsBody from the specified map of raw messages.
func UnmarshalOtherEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_type", &obj.EngineType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEndpoints : Endpoints.
type PrestissimoEndpoints struct {
	// Application API.
	ApplicationsApi *string `json:"applications_api,omitempty"`

	// History server endpoint.
	HistoryServerEndpoint *string `json:"history_server_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkAccessEndpoint *string `json:"spark_access_endpoint,omitempty"`

	// Spark jobs V4 endpoint.
	SparkJobsV4Endpoint *string `json:"spark_jobs_v4_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkKernelEndpoint *string `json:"spark_kernel_endpoint,omitempty"`

	// View history server.
	ViewHistoryServer *string `json:"view_history_server,omitempty"`

	// Wxd application endpoint.
	WxdApplicationEndpoint *string `json:"wxd_application_endpoint,omitempty"`
}

// UnmarshalPrestissimoEndpoints unmarshals an instance of PrestissimoEndpoints from the specified map of raw messages.
func UnmarshalPrestissimoEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEndpoints)
	err = core.UnmarshalPrimitive(m, "applications_api", &obj.ApplicationsApi)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "history_server_endpoint", &obj.HistoryServerEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_access_endpoint", &obj.SparkAccessEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_jobs_v4_endpoint", &obj.SparkJobsV4Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_kernel_endpoint", &obj.SparkKernelEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "view_history_server", &obj.ViewHistoryServer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_application_endpoint", &obj.WxdApplicationEndpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngine : EngineDetail.
type PrestissimoEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalog.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *PrestissimoEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Applicable only for OCP based clusters.  This is typically  servicename+route.
	ExternalHostName *string `json:"external_host_name" validate:"required"`

	// Group ID.
	GroupID *string `json:"group_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Region - place holder.
	Region *string `json:"region,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Engine status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`

	// Version of the engine.
	Version *string `json:"version,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the PrestissimoEngine.Origin property.
// Origin - place holder.
const (
	PrestissimoEngine_Origin_Discover = "discover"
	PrestissimoEngine_Origin_External = "external"
	PrestissimoEngine_Origin_Native = "native"
)

// Constants associated with the PrestissimoEngine.Status property.
// Engine status.
const (
	PrestissimoEngine_Status_Pending = "pending"
	PrestissimoEngine_Status_Running = "running"
	PrestissimoEngine_Status_Stopped = "stopped"
)

// UnmarshalPrestissimoEngine unmarshals an instance of PrestissimoEngine from the specified map of raw messages.
func UnmarshalPrestissimoEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalPrestissimoEngineDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "external_host_name", &obj.ExternalHostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewPrestissimoEnginePatch(prestissimoEngine *PrestissimoEngine) (_patch []JSONPatchOperation) {
	if (prestissimoEngine.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: prestissimoEngine.Actions,
		})
	}
	if (prestissimoEngine.AssociatedCatalogs != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/associated_catalogs"),
			Value: prestissimoEngine.AssociatedCatalogs,
		})
	}
	if (prestissimoEngine.BuildVersion != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/build_version"),
			Value: prestissimoEngine.BuildVersion,
		})
	}
	if (prestissimoEngine.Coordinator != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/coordinator"),
			Value: prestissimoEngine.Coordinator,
		})
	}
	if (prestissimoEngine.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: prestissimoEngine.CreatedBy,
		})
	}
	if (prestissimoEngine.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: prestissimoEngine.CreatedOn,
		})
	}
	if (prestissimoEngine.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: prestissimoEngine.Description,
		})
	}
	if (prestissimoEngine.EngineDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_details"),
			Value: prestissimoEngine.EngineDetails,
		})
	}
	if (prestissimoEngine.EngineDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_display_name"),
			Value: prestissimoEngine.EngineDisplayName,
		})
	}
	if (prestissimoEngine.EngineID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_id"),
			Value: prestissimoEngine.EngineID,
		})
	}
	if (prestissimoEngine.ExternalHostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/external_host_name"),
			Value: prestissimoEngine.ExternalHostName,
		})
	}
	if (prestissimoEngine.GroupID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/group_id"),
			Value: prestissimoEngine.GroupID,
		})
	}
	if (prestissimoEngine.HostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/host_name"),
			Value: prestissimoEngine.HostName,
		})
	}
	if (prestissimoEngine.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: prestissimoEngine.Origin,
		})
	}
	if (prestissimoEngine.Port != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/port"),
			Value: prestissimoEngine.Port,
		})
	}
	if (prestissimoEngine.Region != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/region"),
			Value: prestissimoEngine.Region,
		})
	}
	if (prestissimoEngine.SizeConfig != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/size_config"),
			Value: prestissimoEngine.SizeConfig,
		})
	}
	if (prestissimoEngine.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: prestissimoEngine.Status,
		})
	}
	if (prestissimoEngine.StatusCode != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status_code"),
			Value: prestissimoEngine.StatusCode,
		})
	}
	if (prestissimoEngine.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: prestissimoEngine.Tags,
		})
	}
	if (prestissimoEngine.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: prestissimoEngine.Type,
		})
	}
	if (prestissimoEngine.Version != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/version"),
			Value: prestissimoEngine.Version,
		})
	}
	if (prestissimoEngine.Worker != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/worker"),
			Value: prestissimoEngine.Worker,
		})
	}
	return
}

// PrestissimoEngineCollection : list Prestissimo Engines.
type PrestissimoEngineCollection struct {
	// list prestissimo engines.
	PrestissimoEngines []PrestissimoEngine `json:"prestissimo_engines,omitempty"`
}

// UnmarshalPrestissimoEngineCollection unmarshals an instance of PrestissimoEngineCollection from the specified map of raw messages.
func UnmarshalPrestissimoEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngineCollection)
	err = core.UnmarshalModel(m, "prestissimo_engines", &obj.PrestissimoEngines, UnmarshalPrestissimoEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngineDetails : External engine details.
type PrestissimoEngineDetails struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Endpoints.
	Endpoints *PrestissimoEndpoints `json:"endpoints,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the PrestissimoEngineDetails.SizeConfig property.
// Size config.
const (
	PrestissimoEngineDetails_SizeConfig_CacheOptimized = "cache_optimized"
	PrestissimoEngineDetails_SizeConfig_ComputeOptimized = "compute_optimized"
	PrestissimoEngineDetails_SizeConfig_Custom = "custom"
	PrestissimoEngineDetails_SizeConfig_Large = "large"
	PrestissimoEngineDetails_SizeConfig_Medium = "medium"
	PrestissimoEngineDetails_SizeConfig_Small = "small"
	PrestissimoEngineDetails_SizeConfig_Starter = "starter"
)

// UnmarshalPrestissimoEngineDetails unmarshals an instance of PrestissimoEngineDetails from the specified map of raw messages.
func UnmarshalPrestissimoEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngineDetails)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "endpoints", &obj.Endpoints, UnmarshalPrestissimoEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoNodeDescriptionBody : Node details.
type PrestissimoNodeDescriptionBody struct {
	// Node Type, r5, m, i..
	NodeType *string `json:"node_type,omitempty"`

	// Number of nodes.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalPrestissimoNodeDescriptionBody unmarshals an instance of PrestissimoNodeDescriptionBody from the specified map of raw messages.
func UnmarshalPrestissimoNodeDescriptionBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoNodeDescriptionBody)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestoEngine : EngineDetail.
type PrestoEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Node details.
	EngineDetails *EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Applicable only for OCP based clusters.  This is typically  servicename+route.
	ExternalHostName *string `json:"external_host_name" validate:"required"`

	// Group ID.
	GroupID *string `json:"group_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - created or registered.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Engine status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type presto.
	Type *string `json:"type,omitempty"`

	// Version of the engine.
	Version *string `json:"version,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`
}

// Constants associated with the PrestoEngine.Origin property.
// Origin - created or registered.
const (
	PrestoEngine_Origin_Discover = "discover"
	PrestoEngine_Origin_External = "external"
	PrestoEngine_Origin_Native = "native"
)

// Constants associated with the PrestoEngine.Status property.
// Engine status.
const (
	PrestoEngine_Status_Pending = "pending"
	PrestoEngine_Status_Running = "running"
	PrestoEngine_Status_Stopped = "stopped"
)

// UnmarshalPrestoEngine unmarshals an instance of PrestoEngine from the specified map of raw messages.
func UnmarshalPrestoEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalEngineDetailsBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "external_host_name", &obj.ExternalHostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescription)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewPrestoEnginePatch(prestoEngine *PrestoEngine) (_patch []JSONPatchOperation) {
	if (prestoEngine.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: prestoEngine.Actions,
		})
	}
	if (prestoEngine.AssociatedCatalogs != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/associated_catalogs"),
			Value: prestoEngine.AssociatedCatalogs,
		})
	}
	if (prestoEngine.BuildVersion != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/build_version"),
			Value: prestoEngine.BuildVersion,
		})
	}
	if (prestoEngine.Coordinator != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/coordinator"),
			Value: prestoEngine.Coordinator,
		})
	}
	if (prestoEngine.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: prestoEngine.CreatedBy,
		})
	}
	if (prestoEngine.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: prestoEngine.CreatedOn,
		})
	}
	if (prestoEngine.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: prestoEngine.Description,
		})
	}
	if (prestoEngine.EngineDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_details"),
			Value: prestoEngine.EngineDetails,
		})
	}
	if (prestoEngine.EngineDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_display_name"),
			Value: prestoEngine.EngineDisplayName,
		})
	}
	if (prestoEngine.EngineID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_id"),
			Value: prestoEngine.EngineID,
		})
	}
	if (prestoEngine.ExternalHostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/external_host_name"),
			Value: prestoEngine.ExternalHostName,
		})
	}
	if (prestoEngine.GroupID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/group_id"),
			Value: prestoEngine.GroupID,
		})
	}
	if (prestoEngine.HostName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/host_name"),
			Value: prestoEngine.HostName,
		})
	}
	if (prestoEngine.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: prestoEngine.Origin,
		})
	}
	if (prestoEngine.Port != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/port"),
			Value: prestoEngine.Port,
		})
	}
	if (prestoEngine.Region != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/region"),
			Value: prestoEngine.Region,
		})
	}
	if (prestoEngine.SizeConfig != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/size_config"),
			Value: prestoEngine.SizeConfig,
		})
	}
	if (prestoEngine.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: prestoEngine.Status,
		})
	}
	if (prestoEngine.StatusCode != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status_code"),
			Value: prestoEngine.StatusCode,
		})
	}
	if (prestoEngine.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: prestoEngine.Tags,
		})
	}
	if (prestoEngine.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: prestoEngine.Type,
		})
	}
	if (prestoEngine.Version != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/version"),
			Value: prestoEngine.Version,
		})
	}
	if (prestoEngine.Worker != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/worker"),
			Value: prestoEngine.Worker,
		})
	}
	return
}

// PrestoEngineCollection : List Presto engines.
type PrestoEngineCollection struct {
	// Presto engine.
	PrestoEngines []PrestoEngine `json:"presto_engines,omitempty"`
}

// UnmarshalPrestoEngineCollection unmarshals an instance of PrestoEngineCollection from the specified map of raw messages.
func UnmarshalPrestoEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEngineCollection)
	err = core.UnmarshalModel(m, "presto_engines", &obj.PrestoEngines, UnmarshalPrestoEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplacePrestissimoEngineCatalogsOptions : The ReplacePrestissimoEngineCatalogs options.
type ReplacePrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// comma separated catalog names.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplacePrestissimoEngineCatalogsOptions : Instantiate ReplacePrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewReplacePrestissimoEngineCatalogsOptions(engineID string, catalogNames string) *ReplacePrestissimoEngineCatalogsOptions {
	return &ReplacePrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ReplacePrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *ReplacePrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *ReplacePrestissimoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *ReplacePrestissimoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ReplacePrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ReplacePrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplacePrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *ReplacePrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ReplacePrestoEngineCatalogsOptions : The ReplacePrestoEngineCatalogs options.
type ReplacePrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// comma separated catalog names.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplacePrestoEngineCatalogsOptions : Instantiate ReplacePrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewReplacePrestoEngineCatalogsOptions(engineID string, catalogNames string) *ReplacePrestoEngineCatalogsOptions {
	return &ReplacePrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ReplacePrestoEngineCatalogsOptions) SetEngineID(engineID string) *ReplacePrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *ReplacePrestoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *ReplacePrestoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ReplacePrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ReplacePrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplacePrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *ReplacePrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ReplaceSnapshotCreatedBody : success response.
type ReplaceSnapshotCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalReplaceSnapshotCreatedBody unmarshals an instance of ReplaceSnapshotCreatedBody from the specified map of raw messages.
func UnmarshalReplaceSnapshotCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplaceSnapshotCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceSnapshotOptions : The ReplaceSnapshot options.
type ReplaceSnapshotOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Schema ID.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// Table ID.
	TableID *string `json:"table_id" validate:"required,ne="`

	// Snapshot ID.
	SnapshotID *string `json:"snapshot_id" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceSnapshotOptions : Instantiate ReplaceSnapshotOptions
func (*WatsonxDataV2) NewReplaceSnapshotOptions(engineID string, catalogID string, schemaID string, tableID string, snapshotID string) *ReplaceSnapshotOptions {
	return &ReplaceSnapshotOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		SnapshotID: core.StringPtr(snapshotID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ReplaceSnapshotOptions) SetEngineID(engineID string) *ReplaceSnapshotOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ReplaceSnapshotOptions) SetCatalogID(catalogID string) *ReplaceSnapshotOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ReplaceSnapshotOptions) SetSchemaID(schemaID string) *ReplaceSnapshotOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *ReplaceSnapshotOptions) SetTableID(tableID string) *ReplaceSnapshotOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetSnapshotID : Allow user to set SnapshotID
func (_options *ReplaceSnapshotOptions) SetSnapshotID(snapshotID string) *ReplaceSnapshotOptions {
	_options.SnapshotID = core.StringPtr(snapshotID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ReplaceSnapshotOptions) SetAuthInstanceID(authInstanceID string) *ReplaceSnapshotOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceSnapshotOptions) SetHeaders(param map[string]string) *ReplaceSnapshotOptions {
	options.Headers = param
	return options
}

// ResultPrestissimoExplainStatement : ExplainStatement OK.
type ResultPrestissimoExplainStatement struct {
	// Result.
	Result *string `json:"result,omitempty"`
}

// UnmarshalResultPrestissimoExplainStatement unmarshals an instance of ResultPrestissimoExplainStatement from the specified map of raw messages.
func UnmarshalResultPrestissimoExplainStatement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultPrestissimoExplainStatement)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResultRunPrestissimoExplainAnalyzeStatement : explainAnalyzeStatement OK.
type ResultRunPrestissimoExplainAnalyzeStatement struct {
	// explainAnalyzeStatement result.
	Result *string `json:"result,omitempty"`
}

// UnmarshalResultRunPrestissimoExplainAnalyzeStatement unmarshals an instance of ResultRunPrestissimoExplainAnalyzeStatement from the specified map of raw messages.
func UnmarshalResultRunPrestissimoExplainAnalyzeStatement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultRunPrestissimoExplainAnalyzeStatement)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunExplainAnalyzeStatementOKBody : explainAnalyzeStatement OK.
type RunExplainAnalyzeStatementOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// explainAnalyzeStatement result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalRunExplainAnalyzeStatementOKBody unmarshals an instance of RunExplainAnalyzeStatementOKBody from the specified map of raw messages.
func UnmarshalRunExplainAnalyzeStatementOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RunExplainAnalyzeStatementOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunExplainAnalyzeStatementOptions : The RunExplainAnalyzeStatement options.
type RunExplainAnalyzeStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to show explain analyze.
	Statement *string `json:"statement" validate:"required"`

	// Verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunExplainAnalyzeStatementOptions : Instantiate RunExplainAnalyzeStatementOptions
func (*WatsonxDataV2) NewRunExplainAnalyzeStatementOptions(engineID string, statement string) *RunExplainAnalyzeStatementOptions {
	return &RunExplainAnalyzeStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunExplainAnalyzeStatementOptions) SetEngineID(engineID string) *RunExplainAnalyzeStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunExplainAnalyzeStatementOptions) SetStatement(statement string) *RunExplainAnalyzeStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *RunExplainAnalyzeStatementOptions) SetVerbose(verbose bool) *RunExplainAnalyzeStatementOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunExplainAnalyzeStatementOptions) SetAuthInstanceID(authInstanceID string) *RunExplainAnalyzeStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunExplainAnalyzeStatementOptions) SetHeaders(param map[string]string) *RunExplainAnalyzeStatementOptions {
	options.Headers = param
	return options
}

// RunExplainStatementOKBody : ExplainStatement OK.
type RunExplainStatementOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalRunExplainStatementOKBody unmarshals an instance of RunExplainStatementOKBody from the specified map of raw messages.
func UnmarshalRunExplainStatementOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RunExplainStatementOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunExplainStatementOptions : The RunExplainStatement options.
type RunExplainStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to determine explain plan.
	Statement *string `json:"statement" validate:"required"`

	// Format.
	Format *string `json:"format,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RunExplainStatementOptions.Format property.
// Format.
const (
	RunExplainStatementOptions_Format_Graphviz = "graphviz"
	RunExplainStatementOptions_Format_JSON = "json"
	RunExplainStatementOptions_Format_Text = "text"
)

// Constants associated with the RunExplainStatementOptions.Type property.
// Type.
const (
	RunExplainStatementOptions_Type_Distributed = "distributed"
	RunExplainStatementOptions_Type_Io = "io"
	RunExplainStatementOptions_Type_Logical = "logical"
	RunExplainStatementOptions_Type_Validate = "validate"
)

// NewRunExplainStatementOptions : Instantiate RunExplainStatementOptions
func (*WatsonxDataV2) NewRunExplainStatementOptions(engineID string, statement string) *RunExplainStatementOptions {
	return &RunExplainStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunExplainStatementOptions) SetEngineID(engineID string) *RunExplainStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunExplainStatementOptions) SetStatement(statement string) *RunExplainStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *RunExplainStatementOptions) SetFormat(format string) *RunExplainStatementOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetType : Allow user to set Type
func (_options *RunExplainStatementOptions) SetType(typeVar string) *RunExplainStatementOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunExplainStatementOptions) SetAuthInstanceID(authInstanceID string) *RunExplainStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunExplainStatementOptions) SetHeaders(param map[string]string) *RunExplainStatementOptions {
	options.Headers = param
	return options
}

// RunPrestissimoExplainAnalyzeStatementOptions : The RunPrestissimoExplainAnalyzeStatement options.
type RunPrestissimoExplainAnalyzeStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to show explain analyze.
	Statement *string `json:"statement" validate:"required"`

	// Verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunPrestissimoExplainAnalyzeStatementOptions : Instantiate RunPrestissimoExplainAnalyzeStatementOptions
func (*WatsonxDataV2) NewRunPrestissimoExplainAnalyzeStatementOptions(engineID string, statement string) *RunPrestissimoExplainAnalyzeStatementOptions {
	return &RunPrestissimoExplainAnalyzeStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetEngineID(engineID string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetStatement(statement string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetVerbose(verbose bool) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetAuthInstanceID(authInstanceID string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunPrestissimoExplainAnalyzeStatementOptions) SetHeaders(param map[string]string) *RunPrestissimoExplainAnalyzeStatementOptions {
	options.Headers = param
	return options
}

// RunPrestissimoExplainStatementOptions : The RunPrestissimoExplainStatement options.
type RunPrestissimoExplainStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to determine explain plan.
	Statement *string `json:"statement" validate:"required"`

	// Format.
	Format *string `json:"format,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RunPrestissimoExplainStatementOptions.Format property.
// Format.
const (
	RunPrestissimoExplainStatementOptions_Format_Graphviz = "graphviz"
	RunPrestissimoExplainStatementOptions_Format_JSON = "json"
	RunPrestissimoExplainStatementOptions_Format_Text = "text"
)

// Constants associated with the RunPrestissimoExplainStatementOptions.Type property.
// Type.
const (
	RunPrestissimoExplainStatementOptions_Type_Distributed = "distributed"
	RunPrestissimoExplainStatementOptions_Type_Io = "io"
	RunPrestissimoExplainStatementOptions_Type_Logical = "logical"
	RunPrestissimoExplainStatementOptions_Type_Validate = "validate"
)

// NewRunPrestissimoExplainStatementOptions : Instantiate RunPrestissimoExplainStatementOptions
func (*WatsonxDataV2) NewRunPrestissimoExplainStatementOptions(engineID string, statement string) *RunPrestissimoExplainStatementOptions {
	return &RunPrestissimoExplainStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunPrestissimoExplainStatementOptions) SetEngineID(engineID string) *RunPrestissimoExplainStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunPrestissimoExplainStatementOptions) SetStatement(statement string) *RunPrestissimoExplainStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *RunPrestissimoExplainStatementOptions) SetFormat(format string) *RunPrestissimoExplainStatementOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetType : Allow user to set Type
func (_options *RunPrestissimoExplainStatementOptions) SetType(typeVar string) *RunPrestissimoExplainStatementOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunPrestissimoExplainStatementOptions) SetAuthInstanceID(authInstanceID string) *RunPrestissimoExplainStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunPrestissimoExplainStatementOptions) SetHeaders(param map[string]string) *RunPrestissimoExplainStatementOptions {
	options.Headers = param
	return options
}

// SparkApplicationDetails : Application details.
type SparkApplicationDetails struct {
	// Application.
	Application *string `json:"application" validate:"required"`

	// List of arguments.
	Arguments []string `json:"arguments" validate:"required"`

	// Application.
	Conf *SparkApplicationDetailsConf `json:"conf" validate:"required"`

	// Application.
	Env map[string]interface{} `json:"env" validate:"required"`

	// Display name of the spark application.
	Name *string `json:"name,omitempty"`
}

// NewSparkApplicationDetails : Instantiate SparkApplicationDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewSparkApplicationDetails(application string, arguments []string, conf *SparkApplicationDetailsConf, env map[string]interface{}) (_model *SparkApplicationDetails, err error) {
	_model = &SparkApplicationDetails{
		Application: core.StringPtr(application),
		Arguments: arguments,
		Conf: conf,
		Env: env,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSparkApplicationDetails unmarshals an instance of SparkApplicationDetails from the specified map of raw messages.
func UnmarshalSparkApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conf", &obj.Conf, UnmarshalSparkApplicationDetailsConf)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "env", &obj.Env)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkApplicationDetailsConf : Application.
type SparkApplicationDetailsConf struct {
	// Spark application name.
	SparkAppName *string `json:"spark_app_name,omitempty"`

	// Hive Metastore authentication mode.
	SparkHiveMetastoreClientAuthMode *string `json:"spark_hive_metastore_client_auth_mode,omitempty"`

	// Hive Metastore plain password.
	SparkHiveMetastoreClientPlainPassword *string `json:"spark_hive_metastore_client_plain_password,omitempty"`

	// Hive Metastore plain username.
	SparkHiveMetastoreClientPlainUsername *string `json:"spark_hive_metastore_client_plain_username,omitempty"`

	// Truststore password.
	SparkHiveMetastoreTruststorePassword *string `json:"spark_hive_metastore_truststore_password,omitempty"`

	// Truststore path.
	SparkHiveMetastoreTruststorePath *string `json:"spark_hive_metastore_truststore_path,omitempty"`

	// Truststore type.
	SparkHiveMetastoreTruststoreType *string `json:"spark_hive_metastore_truststore_type,omitempty"`

	// Enable or disable SSL for Hive Metastore.
	SparkHiveMetastoreUseSsl *string `json:"spark_hive_metastore_use_ssl,omitempty"`

	// SQL catalog implementation.
	SparkSqlCatalogImplementation *string `json:"spark_sql_catalog_implementation,omitempty"`

	// Lakehouse catalog name.
	SparkSqlCatalogLakehouse *string `json:"spark_sql_catalog_lakehouse,omitempty"`

	// Lakehouse catalog type.
	SparkSqlCatalogLakehouseType *string `json:"spark_sql_catalog_lakehouse_type,omitempty"`

	// Lakehouse catalog URI.
	SparkSqlCatalogLakehouseURI *string `json:"spark_sql_catalog_lakehouse_uri,omitempty"`

	// SQL extensions.
	SparkSqlExtensions *string `json:"spark_sql_extensions,omitempty"`

	// Enable or disable Iceberg vectorization.
	SparkSqlIcebergVectorizationEnabled *string `json:"spark_sql_iceberg_vectorization_enabled,omitempty"`
}

// UnmarshalSparkApplicationDetailsConf unmarshals an instance of SparkApplicationDetailsConf from the specified map of raw messages.
func UnmarshalSparkApplicationDetailsConf(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkApplicationDetailsConf)
	err = core.UnmarshalPrimitive(m, "spark_app_name", &obj.SparkAppName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_auth_mode", &obj.SparkHiveMetastoreClientAuthMode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_plain_password", &obj.SparkHiveMetastoreClientPlainPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_plain_username", &obj.SparkHiveMetastoreClientPlainUsername)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_password", &obj.SparkHiveMetastoreTruststorePassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_path", &obj.SparkHiveMetastoreTruststorePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_type", &obj.SparkHiveMetastoreTruststoreType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_use_ssl", &obj.SparkHiveMetastoreUseSsl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_implementation", &obj.SparkSqlCatalogImplementation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse", &obj.SparkSqlCatalogLakehouse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse_type", &obj.SparkSqlCatalogLakehouseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse_uri", &obj.SparkSqlCatalogLakehouseURI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_extensions", &obj.SparkSqlExtensions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_iceberg_vectorization_enabled", &obj.SparkSqlIcebergVectorizationEnabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEndpoints : Application Endpoints.
type SparkEndpoints struct {
	// Application API.
	ApplicationsApi *string `json:"applications_api,omitempty"`

	// History server endpoint.
	HistoryServerEndpoint *string `json:"history_server_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkAccessEndpoint *string `json:"spark_access_endpoint,omitempty"`

	// Spark jobs V4 endpoint.
	SparkJobsV4Endpoint *string `json:"spark_jobs_v4_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkKernelEndpoint *string `json:"spark_kernel_endpoint,omitempty"`

	// View history server.
	ViewHistoryServer *string `json:"view_history_server,omitempty"`

	// Wxd application endpoint.
	WxdApplicationEndpoint *string `json:"wxd_application_endpoint,omitempty"`
}

// UnmarshalSparkEndpoints unmarshals an instance of SparkEndpoints from the specified map of raw messages.
func UnmarshalSparkEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEndpoints)
	err = core.UnmarshalPrimitive(m, "applications_api", &obj.ApplicationsApi)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "history_server_endpoint", &obj.HistoryServerEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_access_endpoint", &obj.SparkAccessEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_jobs_v4_endpoint", &obj.SparkJobsV4Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_kernel_endpoint", &obj.SparkKernelEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "view_history_server", &obj.ViewHistoryServer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_application_endpoint", &obj.WxdApplicationEndpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngine : EngineDetail.
type SparkEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *SparkEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Type like spark, netezza,..
	Type *string `json:"type,omitempty"`
}

// UnmarshalSparkEngine unmarshals an instance of SparkEngine from the specified map of raw messages.
func UnmarshalSparkEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalSparkEngineDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewSparkEnginePatch(sparkEngine *SparkEngine) (_patch []JSONPatchOperation) {
	if (sparkEngine.Actions != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/actions"),
			Value: sparkEngine.Actions,
		})
	}
	if (sparkEngine.BuildVersion != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/build_version"),
			Value: sparkEngine.BuildVersion,
		})
	}
	if (sparkEngine.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_by"),
			Value: sparkEngine.CreatedBy,
		})
	}
	if (sparkEngine.CreatedOn != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/created_on"),
			Value: sparkEngine.CreatedOn,
		})
	}
	if (sparkEngine.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: sparkEngine.Description,
		})
	}
	if (sparkEngine.EngineDetails != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_details"),
			Value: sparkEngine.EngineDetails,
		})
	}
	if (sparkEngine.EngineDisplayName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_display_name"),
			Value: sparkEngine.EngineDisplayName,
		})
	}
	if (sparkEngine.EngineID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/engine_id"),
			Value: sparkEngine.EngineID,
		})
	}
	if (sparkEngine.Origin != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/origin"),
			Value: sparkEngine.Origin,
		})
	}
	if (sparkEngine.Status != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/status"),
			Value: sparkEngine.Status,
		})
	}
	if (sparkEngine.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/tags"),
			Value: sparkEngine.Tags,
		})
	}
	if (sparkEngine.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: sparkEngine.Type,
		})
	}
	return
}

// SparkEngineApplicationStatus : Engine Application Status.
type SparkEngineApplicationStatus struct {
	// Application Details.
	ApplicationDetails *SparkEngineApplicationStatusApplicationDetails `json:"application_details,omitempty"`

	// Application ID.
	ApplicationID *string `json:"application_id,omitempty"`

	// Auto Termination Time.
	AutoTerminationTime *string `json:"auto_termination_time,omitempty"`

	// Creation time.
	CreationTime *string `json:"creation_time,omitempty"`

	// Deployment mode.
	DeployMode *string `json:"deploy_mode,omitempty"`

	// End Time.
	EndTime *string `json:"end_time,omitempty"`

	// Failed time.
	FailedTime *string `json:"failed_time,omitempty"`

	// Finish time.
	FinishTime *string `json:"finish_time,omitempty"`

	// Application ID.
	ID *string `json:"id,omitempty"`

	// Job endpoint.
	JobEndpoint *string `json:"job_endpoint,omitempty"`

	// Return code.
	ReturnCode *string `json:"return_code,omitempty"`

	// application run time.
	Runtime *SparkEngineApplicationStatusRuntime `json:"runtime,omitempty"`

	// Service Instance ID for POST.
	ServiceInstanceID *string `json:"service_instance_id,omitempty"`

	// Spark application ID.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Spark application name.
	SparkApplicationName *string `json:"spark_application_name,omitempty"`

	// Start time.
	StartTime *string `json:"start_time,omitempty"`

	// Application state.
	State *string `json:"state,omitempty"`

	// Application state details.
	StateDetails []SparkEngineApplicationStatusStateDetailsItems `json:"state_details,omitempty"`

	// Application submission time.
	SubmissionTime *string `json:"submission_time,omitempty"`

	// Template ID.
	TemplateID *string `json:"template_id,omitempty"`

	// Engine Type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the SparkEngineApplicationStatus.Type property.
// Engine Type.
const (
	SparkEngineApplicationStatus_Type_Emr = "emr"
	SparkEngineApplicationStatus_Type_Iae = "iae"
)

// UnmarshalSparkEngineApplicationStatus unmarshals an instance of SparkEngineApplicationStatus from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatus)
	err = core.UnmarshalModel(m, "application_details", &obj.ApplicationDetails, UnmarshalSparkEngineApplicationStatusApplicationDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creation_time", &obj.CreationTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deploy_mode", &obj.DeployMode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed_time", &obj.FailedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "job_endpoint", &obj.JobEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "return_code", &obj.ReturnCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "runtime", &obj.Runtime, UnmarshalSparkEngineApplicationStatusRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_instance_id", &obj.ServiceInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_id", &obj.SparkApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_name", &obj.SparkApplicationName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "state_details", &obj.StateDetails, UnmarshalSparkEngineApplicationStatusStateDetailsItems)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "submission_time", &obj.SubmissionTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_id", &obj.TemplateID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusApplicationDetails : Application Details.
type SparkEngineApplicationStatusApplicationDetails struct {
	// Application.
	Application *string `json:"application,omitempty"`

	// List of arguments.
	Arguments []string `json:"arguments,omitempty"`

	// Application.
	Conf *SparkEngineApplicationStatusApplicationDetailsConf `json:"conf,omitempty"`

	// Application.
	Env map[string]interface{} `json:"env,omitempty"`

	// Display name of the spark application.
	Name *string `json:"name,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusApplicationDetails unmarshals an instance of SparkEngineApplicationStatusApplicationDetails from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conf", &obj.Conf, UnmarshalSparkEngineApplicationStatusApplicationDetailsConf)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "env", &obj.Env)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusApplicationDetailsConf : Application.
type SparkEngineApplicationStatusApplicationDetailsConf struct {
	// Spark application name.
	SparkAppName *string `json:"spark_app_name,omitempty"`

	// Hive Metastore authentication mode.
	SparkHiveMetastoreClientAuthMode *string `json:"spark_hive_metastore_client_auth_mode,omitempty"`

	// Hive Metastore plain password.
	SparkHiveMetastoreClientPlainPassword *string `json:"spark_hive_metastore_client_plain_password,omitempty"`

	// Hive Metastore plain username.
	SparkHiveMetastoreClientPlainUsername *string `json:"spark_hive_metastore_client_plain_username,omitempty"`

	// Truststore password.
	SparkHiveMetastoreTruststorePassword *string `json:"spark_hive_metastore_truststore_password,omitempty"`

	// Truststore path.
	SparkHiveMetastoreTruststorePath *string `json:"spark_hive_metastore_truststore_path,omitempty"`

	// Truststore type.
	SparkHiveMetastoreTruststoreType *string `json:"spark_hive_metastore_truststore_type,omitempty"`

	// Enable or disable SSL for Hive Metastore.
	SparkHiveMetastoreUseSsl *string `json:"spark_hive_metastore_use_ssl,omitempty"`

	// SQL catalog implementation.
	SparkSqlCatalogImplementation *string `json:"spark_sql_catalog_implementation,omitempty"`

	// Lakehouse catalog name.
	SparkSqlCatalogLakehouse *string `json:"spark_sql_catalog_lakehouse,omitempty"`

	// Lakehouse catalog type.
	SparkSqlCatalogLakehouseType *string `json:"spark_sql_catalog_lakehouse_type,omitempty"`

	// Lakehouse catalog URI.
	SparkSqlCatalogLakehouseURI *string `json:"spark_sql_catalog_lakehouse_uri,omitempty"`

	// SQL extensions.
	SparkSqlExtensions *string `json:"spark_sql_extensions,omitempty"`

	// Enable or disable Iceberg vectorization.
	SparkSqlIcebergVectorizationEnabled *string `json:"spark_sql_iceberg_vectorization_enabled,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusApplicationDetailsConf unmarshals an instance of SparkEngineApplicationStatusApplicationDetailsConf from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusApplicationDetailsConf(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusApplicationDetailsConf)
	err = core.UnmarshalPrimitive(m, "spark_app_name", &obj.SparkAppName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_auth_mode", &obj.SparkHiveMetastoreClientAuthMode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_plain_password", &obj.SparkHiveMetastoreClientPlainPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_client_plain_username", &obj.SparkHiveMetastoreClientPlainUsername)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_password", &obj.SparkHiveMetastoreTruststorePassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_path", &obj.SparkHiveMetastoreTruststorePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_truststore_type", &obj.SparkHiveMetastoreTruststoreType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_hive_metastore_use_ssl", &obj.SparkHiveMetastoreUseSsl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_implementation", &obj.SparkSqlCatalogImplementation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse", &obj.SparkSqlCatalogLakehouse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse_type", &obj.SparkSqlCatalogLakehouseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_catalog_lakehouse_uri", &obj.SparkSqlCatalogLakehouseURI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_extensions", &obj.SparkSqlExtensions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_sql_iceberg_vectorization_enabled", &obj.SparkSqlIcebergVectorizationEnabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusCollection : Engine Application Detail.
type SparkEngineApplicationStatusCollection struct {
	// Application body.
	Applications []SparkEngineApplicationStatus `json:"applications,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusCollection unmarshals an instance of SparkEngineApplicationStatusCollection from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusCollection)
	err = core.UnmarshalModel(m, "applications", &obj.Applications, UnmarshalSparkEngineApplicationStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusRuntime : application run time.
type SparkEngineApplicationStatusRuntime struct {
	// Spark Version.
	SparkVersion *string `json:"spark_version,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusRuntime unmarshals an instance of SparkEngineApplicationStatusRuntime from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusRuntime)
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusStateDetailsItems : State details.
type SparkEngineApplicationStatusStateDetailsItems struct {
	// State details code.
	Code *string `json:"code,omitempty"`

	// State details message.
	Message *string `json:"message,omitempty"`

	// State details type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusStateDetailsItems unmarshals an instance of SparkEngineApplicationStatusStateDetailsItems from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusStateDetailsItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusStateDetailsItems)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineCollection : List spark engines.
type SparkEngineCollection struct {
	// List spark engines.
	SparkEngines []SparkEngine `json:"spark_engines,omitempty"`
}

// UnmarshalSparkEngineCollection unmarshals an instance of SparkEngineCollection from the specified map of raw messages.
func UnmarshalSparkEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineCollection)
	err = core.UnmarshalModel(m, "spark_engines", &obj.SparkEngines, UnmarshalSparkEngine)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineDetails : External engine details.
type SparkEngineDetails struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Application Endpoints.
	Endpoints *SparkEndpoints `json:"endpoints,omitempty"`
}

// UnmarshalSparkEngineDetails unmarshals an instance of SparkEngineDetails from the specified map of raw messages.
func UnmarshalSparkEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "endpoints", &obj.Endpoints, UnmarshalSparkEndpoints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineDetailsPrototype : Node details.
type SparkEngineDetailsPrototype struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`
}

// UnmarshalSparkEngineDetailsPrototype unmarshals an instance of SparkEngineDetailsPrototype from the specified map of raw messages.
func UnmarshalSparkEngineDetailsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineDetailsPrototype)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessResponse : Response of success.
type SuccessResponse struct {
	// Message.
	Message *string `json:"message,omitempty"`

	// Message code.
	MessageCode *string `json:"message_code,omitempty"`
}

// UnmarshalSuccessResponse unmarshals an instance of SuccessResponse from the specified map of raw messages.
func UnmarshalSuccessResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_code", &obj.MessageCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Table : GetColumns OK.
type Table struct {
	// Columns.
	Columns []Column `json:"columns,omitempty"`

	// Table name.
	TableName *string `json:"table_name,omitempty"`
}

// UnmarshalTable unmarshals an instance of Table from the specified map of raw messages.
func UnmarshalTable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Table)
	err = core.UnmarshalModel(m, "columns", &obj.Columns, UnmarshalColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewTablePatch(table *Table) (_patch []JSONPatchOperation) {
	if (table.Columns != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/columns"),
			Value: table.Columns,
		})
	}
	if (table.TableName != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/table_name"),
			Value: table.TableName,
		})
	}
	return
}

// TableCollection : tables list.
type TableCollection struct {
	// List of the tables present in the schema.
	Tables []string `json:"tables,omitempty"`
}

// UnmarshalTableCollection unmarshals an instance of TableCollection from the specified map of raw messages.
func UnmarshalTableCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableCollection)
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableSnapshot : TableSnapshot.
type TableSnapshot struct {
	// Committed at.
	CommittedAt *string `json:"committed_at,omitempty"`

	// Operation.
	Operation *string `json:"operation,omitempty"`

	// Snapshot id.
	SnapshotID *string `json:"snapshot_id,omitempty"`

	// Summary.
	Summary map[string]interface{} `json:"summary,omitempty"`
}

// UnmarshalTableSnapshot unmarshals an instance of TableSnapshot from the specified map of raw messages.
func UnmarshalTableSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableSnapshot)
	err = core.UnmarshalPrimitive(m, "committed_at", &obj.CommittedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snapshot_id", &obj.SnapshotID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "summary", &obj.Summary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableSnapshotCollection : TableSnapshot OK.
type TableSnapshotCollection struct {
	// Snapshots.
	Snapshots []TableSnapshot `json:"snapshots,omitempty"`
}

// UnmarshalTableSnapshotCollection unmarshals an instance of TableSnapshotCollection from the specified map of raw messages.
func UnmarshalTableSnapshotCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableSnapshotCollection)
	err = core.UnmarshalModel(m, "snapshots", &obj.Snapshots, UnmarshalTableSnapshot)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestBucketConnectionOKBody : ValidateBucketRegistrationCredentials OK.
type TestBucketConnectionOKBody struct {
	// object defining the response of checking if the credentials of a bucket are valid.
	BucketStatus *BucketStatusResponse `json:"bucket_status" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalTestBucketConnectionOKBody unmarshals an instance of TestBucketConnectionOKBody from the specified map of raw messages.
func UnmarshalTestBucketConnectionOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestBucketConnectionOKBody)
	err = core.UnmarshalModel(m, "bucket_status", &obj.BucketStatus, UnmarshalBucketStatusResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestBucketConnectionOptions : The TestBucketConnection options.
type TestBucketConnectionOptions struct {
	// access key to access the bucket.
	AccessKey *string `json:"access_key" validate:"required"`

	// name of the bucket to be checked.
	BucketName *string `json:"bucket_name" validate:"required"`

	// type of bucket that is selected.
	BucketType *string `json:"bucket_type" validate:"required"`

	// endpoint to reach the bucket.
	Endpoint *string `json:"endpoint" validate:"required"`

	// bucket region.
	Region *string `json:"region" validate:"required"`

	// secret key to access the bucket.
	SecretKey *string `json:"secret_key" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the TestBucketConnectionOptions.BucketType property.
// type of bucket that is selected.
const (
	TestBucketConnectionOptions_BucketType_AmazonS3 = "amazon_s3"
	TestBucketConnectionOptions_BucketType_AwsS3 = "aws_s3"
	TestBucketConnectionOptions_BucketType_IbmCeph = "ibm_ceph"
	TestBucketConnectionOptions_BucketType_IbmCos = "ibm_cos"
	TestBucketConnectionOptions_BucketType_Minio = "minio"
)

// NewTestBucketConnectionOptions : Instantiate TestBucketConnectionOptions
func (*WatsonxDataV2) NewTestBucketConnectionOptions(accessKey string, bucketName string, bucketType string, endpoint string, region string, secretKey string) *TestBucketConnectionOptions {
	return &TestBucketConnectionOptions{
		AccessKey: core.StringPtr(accessKey),
		BucketName: core.StringPtr(bucketName),
		BucketType: core.StringPtr(bucketType),
		Endpoint: core.StringPtr(endpoint),
		Region: core.StringPtr(region),
		SecretKey: core.StringPtr(secretKey),
	}
}

// SetAccessKey : Allow user to set AccessKey
func (_options *TestBucketConnectionOptions) SetAccessKey(accessKey string) *TestBucketConnectionOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetBucketName : Allow user to set BucketName
func (_options *TestBucketConnectionOptions) SetBucketName(bucketName string) *TestBucketConnectionOptions {
	_options.BucketName = core.StringPtr(bucketName)
	return _options
}

// SetBucketType : Allow user to set BucketType
func (_options *TestBucketConnectionOptions) SetBucketType(bucketType string) *TestBucketConnectionOptions {
	_options.BucketType = core.StringPtr(bucketType)
	return _options
}

// SetEndpoint : Allow user to set Endpoint
func (_options *TestBucketConnectionOptions) SetEndpoint(endpoint string) *TestBucketConnectionOptions {
	_options.Endpoint = core.StringPtr(endpoint)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *TestBucketConnectionOptions) SetRegion(region string) *TestBucketConnectionOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetSecretKey : Allow user to set SecretKey
func (_options *TestBucketConnectionOptions) SetSecretKey(secretKey string) *TestBucketConnectionOptions {
	_options.SecretKey = core.StringPtr(secretKey)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *TestBucketConnectionOptions) SetAuthInstanceID(authInstanceID string) *TestBucketConnectionOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TestBucketConnectionOptions) SetHeaders(param map[string]string) *TestBucketConnectionOptions {
	options.Headers = param
	return options
}

// TestDatabaseConnectionResponse : Success response.
type TestDatabaseConnectionResponse struct {
	// check connection response details are valid or not.
	ConnectionResponse *ConnectionResponse `json:"connection_response,omitempty"`
}

// UnmarshalTestDatabaseConnectionResponse unmarshals an instance of TestDatabaseConnectionResponse from the specified map of raw messages.
func UnmarshalTestDatabaseConnectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDatabaseConnectionResponse)
	err = core.UnmarshalModel(m, "connection_response", &obj.ConnectionResponse, UnmarshalConnectionResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestLHConsoleOptions : The TestLHConsole options.
type TestLHConsoleOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTestLHConsoleOptions : Instantiate TestLHConsoleOptions
func (*WatsonxDataV2) NewTestLHConsoleOptions() *TestLHConsoleOptions {
	return &TestLHConsoleOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *TestLHConsoleOptions) SetHeaders(param map[string]string) *TestLHConsoleOptions {
	options.Headers = param
	return options
}

// UpdateBucketRegistrationOptions : The UpdateBucketRegistration options.
type UpdateBucketRegistrationOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Request body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBucketRegistrationOptions : Instantiate UpdateBucketRegistrationOptions
func (*WatsonxDataV2) NewUpdateBucketRegistrationOptions(bucketID string, body []JSONPatchOperation) *UpdateBucketRegistrationOptions {
	return &UpdateBucketRegistrationOptions{
		BucketID: core.StringPtr(bucketID),
		Body: body,
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *UpdateBucketRegistrationOptions) SetBucketID(bucketID string) *UpdateBucketRegistrationOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateBucketRegistrationOptions) SetBody(body []JSONPatchOperation) *UpdateBucketRegistrationOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *UpdateBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBucketRegistrationOptions) SetHeaders(param map[string]string) *UpdateBucketRegistrationOptions {
	options.Headers = param
	return options
}

// UpdateColumnOptions : The UpdateColumn options.
type UpdateColumnOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// URL encoded schema name.
	ColumnID *string `json:"column_id" validate:"required,ne="`

	// Request body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateColumnOptions : Instantiate UpdateColumnOptions
func (*WatsonxDataV2) NewUpdateColumnOptions(engineID string, catalogID string, schemaID string, tableID string, columnID string, body []JSONPatchOperation) *UpdateColumnOptions {
	return &UpdateColumnOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		ColumnID: core.StringPtr(columnID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateColumnOptions) SetEngineID(engineID string) *UpdateColumnOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *UpdateColumnOptions) SetCatalogID(catalogID string) *UpdateColumnOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *UpdateColumnOptions) SetSchemaID(schemaID string) *UpdateColumnOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *UpdateColumnOptions) SetTableID(tableID string) *UpdateColumnOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumnID : Allow user to set ColumnID
func (_options *UpdateColumnOptions) SetColumnID(columnID string) *UpdateColumnOptions {
	_options.ColumnID = core.StringPtr(columnID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateColumnOptions) SetBody(body []JSONPatchOperation) *UpdateColumnOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateColumnOptions) SetAuthInstanceID(authInstanceID string) *UpdateColumnOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateColumnOptions) SetHeaders(param map[string]string) *UpdateColumnOptions {
	options.Headers = param
	return options
}

// UpdateDatabaseOptions : The UpdateDatabase options.
type UpdateDatabaseOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// Request body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDatabaseOptions : Instantiate UpdateDatabaseOptions
func (*WatsonxDataV2) NewUpdateDatabaseOptions(databaseID string, body []JSONPatchOperation) *UpdateDatabaseOptions {
	return &UpdateDatabaseOptions{
		DatabaseID: core.StringPtr(databaseID),
		Body: body,
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *UpdateDatabaseOptions) SetDatabaseID(databaseID string) *UpdateDatabaseOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateDatabaseOptions) SetBody(body []JSONPatchOperation) *UpdateDatabaseOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateDatabaseOptions) SetAuthInstanceID(authInstanceID string) *UpdateDatabaseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDatabaseOptions) SetHeaders(param map[string]string) *UpdateDatabaseOptions {
	options.Headers = param
	return options
}

// UpdateDb2EngineOptions : The UpdateDb2Engine options.
type UpdateDb2EngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDb2EngineOptions : Instantiate UpdateDb2EngineOptions
func (*WatsonxDataV2) NewUpdateDb2EngineOptions(engineID string, body []JSONPatchOperation) *UpdateDb2EngineOptions {
	return &UpdateDb2EngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateDb2EngineOptions) SetEngineID(engineID string) *UpdateDb2EngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateDb2EngineOptions) SetBody(body []JSONPatchOperation) *UpdateDb2EngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDb2EngineOptions) SetHeaders(param map[string]string) *UpdateDb2EngineOptions {
	options.Headers = param
	return options
}

// UpdateMilvusServiceOptions : The UpdateMilvusService options.
type UpdateMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// Update milvus service.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateMilvusServiceOptions : Instantiate UpdateMilvusServiceOptions
func (*WatsonxDataV2) NewUpdateMilvusServiceOptions(serviceID string, body []JSONPatchOperation) *UpdateMilvusServiceOptions {
	return &UpdateMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
		Body: body,
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *UpdateMilvusServiceOptions) SetServiceID(serviceID string) *UpdateMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateMilvusServiceOptions) SetBody(body []JSONPatchOperation) *UpdateMilvusServiceOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *UpdateMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMilvusServiceOptions) SetHeaders(param map[string]string) *UpdateMilvusServiceOptions {
	options.Headers = param
	return options
}

// UpdateNetezzaEngineOptions : The UpdateNetezzaEngine options.
type UpdateNetezzaEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateNetezzaEngineOptions : Instantiate UpdateNetezzaEngineOptions
func (*WatsonxDataV2) NewUpdateNetezzaEngineOptions(engineID string, body []JSONPatchOperation) *UpdateNetezzaEngineOptions {
	return &UpdateNetezzaEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateNetezzaEngineOptions) SetEngineID(engineID string) *UpdateNetezzaEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateNetezzaEngineOptions) SetBody(body []JSONPatchOperation) *UpdateNetezzaEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNetezzaEngineOptions) SetHeaders(param map[string]string) *UpdateNetezzaEngineOptions {
	options.Headers = param
	return options
}

// UpdatePrestissimoEngineOptions : The UpdatePrestissimoEngine options.
type UpdatePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update prestissimo engine body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdatePrestissimoEngineOptions : Instantiate UpdatePrestissimoEngineOptions
func (*WatsonxDataV2) NewUpdatePrestissimoEngineOptions(engineID string, body []JSONPatchOperation) *UpdatePrestissimoEngineOptions {
	return &UpdatePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdatePrestissimoEngineOptions) SetEngineID(engineID string) *UpdatePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdatePrestissimoEngineOptions) SetBody(body []JSONPatchOperation) *UpdatePrestissimoEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdatePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdatePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrestissimoEngineOptions) SetHeaders(param map[string]string) *UpdatePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// UpdatePrestoEngineOptions : The UpdatePrestoEngine options.
type UpdatePrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdatePrestoEngineOptions : Instantiate UpdatePrestoEngineOptions
func (*WatsonxDataV2) NewUpdatePrestoEngineOptions(engineID string, body []JSONPatchOperation) *UpdatePrestoEngineOptions {
	return &UpdatePrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdatePrestoEngineOptions) SetEngineID(engineID string) *UpdatePrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdatePrestoEngineOptions) SetBody(body []JSONPatchOperation) *UpdatePrestoEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdatePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdatePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrestoEngineOptions) SetHeaders(param map[string]string) *UpdatePrestoEngineOptions {
	options.Headers = param
	return options
}

// UpdateSparkEngineOptions : The UpdateSparkEngine options.
type UpdateSparkEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSparkEngineOptions : Instantiate UpdateSparkEngineOptions
func (*WatsonxDataV2) NewUpdateSparkEngineOptions(engineID string, body []JSONPatchOperation) *UpdateSparkEngineOptions {
	return &UpdateSparkEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateSparkEngineOptions) SetEngineID(engineID string) *UpdateSparkEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateSparkEngineOptions) SetBody(body []JSONPatchOperation) *UpdateSparkEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSparkEngineOptions) SetHeaders(param map[string]string) *UpdateSparkEngineOptions {
	options.Headers = param
	return options
}

// UpdateSyncCatalogOKBody : success response.
type UpdateSyncCatalogOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalUpdateSyncCatalogOKBody unmarshals an instance of UpdateSyncCatalogOKBody from the specified map of raw messages.
func UnmarshalUpdateSyncCatalogOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateSyncCatalogOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonxDataV2) NewUpdateSyncCatalogOKBodyPatch(updateSyncCatalogOKBody *UpdateSyncCatalogOKBody) (_patch []JSONPatchOperation) {
	if (updateSyncCatalogOKBody.Response != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/response"),
			Value: updateSyncCatalogOKBody.Response,
		})
	}
	return
}

// UpdateSyncCatalogOptions : The UpdateSyncCatalog options.
type UpdateSyncCatalogOptions struct {
	// catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Request body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSyncCatalogOptions : Instantiate UpdateSyncCatalogOptions
func (*WatsonxDataV2) NewUpdateSyncCatalogOptions(catalogID string, body []JSONPatchOperation) *UpdateSyncCatalogOptions {
	return &UpdateSyncCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
		Body: body,
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *UpdateSyncCatalogOptions) SetCatalogID(catalogID string) *UpdateSyncCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateSyncCatalogOptions) SetBody(body []JSONPatchOperation) *UpdateSyncCatalogOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateSyncCatalogOptions) SetAuthInstanceID(authInstanceID string) *UpdateSyncCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSyncCatalogOptions) SetHeaders(param map[string]string) *UpdateSyncCatalogOptions {
	options.Headers = param
	return options
}

// UpdateTableOptions : The UpdateTable options.
type UpdateTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// Request body.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTableOptions : Instantiate UpdateTableOptions
func (*WatsonxDataV2) NewUpdateTableOptions(catalogID string, schemaID string, tableID string, engineID string, body []JSONPatchOperation) *UpdateTableOptions {
	return &UpdateTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *UpdateTableOptions) SetCatalogID(catalogID string) *UpdateTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *UpdateTableOptions) SetSchemaID(schemaID string) *UpdateTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *UpdateTableOptions) SetTableID(tableID string) *UpdateTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateTableOptions) SetEngineID(engineID string) *UpdateTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateTableOptions) SetBody(body []JSONPatchOperation) *UpdateTableOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateTableOptions) SetAuthInstanceID(authInstanceID string) *UpdateTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTableOptions) SetHeaders(param map[string]string) *UpdateTableOptions {
	options.Headers = param
	return options
}

// ValidateDatabaseBodyDatabaseDetails : database details.
type ValidateDatabaseBodyDatabaseDetails struct {
	// db name.
	DatabaseName *string `json:"database_name,omitempty"`

	// Host name.
	Hostname *string `json:"hostname" validate:"required"`

	// Psssword.
	Password *string `json:"password,omitempty"`

	// Port.
	Port *int64 `json:"port" validate:"required"`

	// SASL Mode.
	Sasl *bool `json:"sasl,omitempty"`

	// SSL Mode.
	Ssl *bool `json:"ssl,omitempty"`

	// Only for Kafka - Add kafka tables.
	Tables *string `json:"tables,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`

	// Verify certificate.
	ValidateServerCertificate *bool `json:"validate_server_certificate,omitempty"`
}

// NewValidateDatabaseBodyDatabaseDetails : Instantiate ValidateDatabaseBodyDatabaseDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewValidateDatabaseBodyDatabaseDetails(hostname string, port int64) (_model *ValidateDatabaseBodyDatabaseDetails, err error) {
	_model = &ValidateDatabaseBodyDatabaseDetails{
		Hostname: core.StringPtr(hostname),
		Port: core.Int64Ptr(port),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalValidateDatabaseBodyDatabaseDetails unmarshals an instance of ValidateDatabaseBodyDatabaseDetails from the specified map of raw messages.
func UnmarshalValidateDatabaseBodyDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ValidateDatabaseBodyDatabaseDetails)
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sasl", &obj.Sasl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "validate_server_certificate", &obj.ValidateServerCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ValidateDatabaseConnectionOptions : The ValidateDatabaseConnection options.
type ValidateDatabaseConnectionOptions struct {
	// database details.
	DatabaseDetails *ValidateDatabaseBodyDatabaseDetails `json:"database_details" validate:"required"`

	// Type of db connection.
	DatabaseType *string `json:"database_type" validate:"required"`

	// contents of a pem/crt file.
	Certificate *string `json:"certificate,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ValidateDatabaseConnectionOptions.DatabaseType property.
// Type of db connection.
const (
	ValidateDatabaseConnectionOptions_DatabaseType_Db2 = "db2"
	ValidateDatabaseConnectionOptions_DatabaseType_Kafka = "kafka"
	ValidateDatabaseConnectionOptions_DatabaseType_Mongodb = "mongodb"
	ValidateDatabaseConnectionOptions_DatabaseType_Mycustomdb = "mycustomdb"
	ValidateDatabaseConnectionOptions_DatabaseType_Mysql = "mysql"
	ValidateDatabaseConnectionOptions_DatabaseType_Netezza = "netezza"
	ValidateDatabaseConnectionOptions_DatabaseType_Postgresql = "postgresql"
	ValidateDatabaseConnectionOptions_DatabaseType_Sqlserver = "sqlserver"
)

// NewValidateDatabaseConnectionOptions : Instantiate ValidateDatabaseConnectionOptions
func (*WatsonxDataV2) NewValidateDatabaseConnectionOptions(databaseDetails *ValidateDatabaseBodyDatabaseDetails, databaseType string) *ValidateDatabaseConnectionOptions {
	return &ValidateDatabaseConnectionOptions{
		DatabaseDetails: databaseDetails,
		DatabaseType: core.StringPtr(databaseType),
	}
}

// SetDatabaseDetails : Allow user to set DatabaseDetails
func (_options *ValidateDatabaseConnectionOptions) SetDatabaseDetails(databaseDetails *ValidateDatabaseBodyDatabaseDetails) *ValidateDatabaseConnectionOptions {
	_options.DatabaseDetails = databaseDetails
	return _options
}

// SetDatabaseType : Allow user to set DatabaseType
func (_options *ValidateDatabaseConnectionOptions) SetDatabaseType(databaseType string) *ValidateDatabaseConnectionOptions {
	_options.DatabaseType = core.StringPtr(databaseType)
	return _options
}

// SetCertificate : Allow user to set Certificate
func (_options *ValidateDatabaseConnectionOptions) SetCertificate(certificate string) *ValidateDatabaseConnectionOptions {
	_options.Certificate = core.StringPtr(certificate)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ValidateDatabaseConnectionOptions) SetAuthInstanceID(authInstanceID string) *ValidateDatabaseConnectionOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ValidateDatabaseConnectionOptions) SetHeaders(param map[string]string) *ValidateDatabaseConnectionOptions {
	options.Headers = param
	return options
}
