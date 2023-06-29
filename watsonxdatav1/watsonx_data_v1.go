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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.72.2-2bede9d2-20230601-202845
 */

// Package watsonxdatav1 : Operations and models for the WatsonxDataV1 service
package watsonxdatav1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/watsonxdata-go-sdk/common"
)

// WatsonxDataV1 : This is the Public API for IBM watsonx.data
//
// API Version: SaaS-GA-1.0.0
type WatsonxDataV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://lakehouse/api/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watsonx_data"

// WatsonxDataV1Options : Service options
type WatsonxDataV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewWatsonxDataV1UsingExternalConfig : constructs an instance of WatsonxDataV1 with passed in options and external configuration.
func NewWatsonxDataV1UsingExternalConfig(options *WatsonxDataV1Options) (watsonxData *WatsonxDataV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	watsonxData, err = NewWatsonxDataV1(options)
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

// NewWatsonxDataV1 : constructs an instance of WatsonxDataV1 with passed in options.
func NewWatsonxDataV1(options *WatsonxDataV1Options) (service *WatsonxDataV1, err error) {
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

	service = &WatsonxDataV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "watsonxData" suitable for processing requests.
func (watsonxData *WatsonxDataV1) Clone() *WatsonxDataV1 {
	if core.IsNil(watsonxData) {
		return nil
	}
	clone := *watsonxData
	clone.Service = watsonxData.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (watsonxData *WatsonxDataV1) SetServiceURL(url string) error {
	return watsonxData.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (watsonxData *WatsonxDataV1) GetServiceURL() string {
	return watsonxData.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (watsonxData *WatsonxDataV1) SetDefaultHeaders(headers http.Header) {
	watsonxData.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV1) SetEnableGzipCompression(enableGzip bool) {
	watsonxData.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV1) GetEnableGzipCompression() bool {
	return watsonxData.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (watsonxData *WatsonxDataV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	watsonxData.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (watsonxData *WatsonxDataV1) DisableRetries() {
	watsonxData.Service.DisableRetries()
}

// CreateDbConnUsers : Grant users and groups permission to the db connection
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) CreateDbConnUsers(createDbConnUsersOptions *CreateDbConnUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDbConnUsersWithContext(context.Background(), createDbConnUsersOptions)
}

// CreateDbConnUsersWithContext is an alternate form of the CreateDbConnUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateDbConnUsersWithContext(ctx context.Context, createDbConnUsersOptions *CreateDbConnUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDbConnUsersOptions, "createDbConnUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDbConnUsersOptions, "createDbConnUsersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/databases`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDbConnUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateDbConnUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDbConnUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createDbConnUsersOptions.LhInstanceID))
	}
	if createDbConnUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDbConnUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDbConnUsersOptions.DatabaseID != nil {
		body["database_id"] = createDbConnUsersOptions.DatabaseID
	}
	if createDbConnUsersOptions.Groups != nil {
		body["groups"] = createDbConnUsersOptions.Groups
	}
	if createDbConnUsersOptions.Users != nil {
		body["users"] = createDbConnUsersOptions.Users
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

// ListDataPolicies : Get policies
// Get list of all data policies.
func (watsonxData *WatsonxDataV1) ListDataPolicies(listDataPoliciesOptions *ListDataPoliciesOptions) (result *PolicyListSchema, response *core.DetailedResponse, err error) {
	return watsonxData.ListDataPoliciesWithContext(context.Background(), listDataPoliciesOptions)
}

// ListDataPoliciesWithContext is an alternate form of the ListDataPolicies method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ListDataPoliciesWithContext(ctx context.Context, listDataPoliciesOptions *ListDataPoliciesOptions) (result *PolicyListSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataPoliciesOptions, "listDataPoliciesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataPoliciesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ListDataPolicies")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDataPoliciesOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*listDataPoliciesOptions.LhInstanceID))
	}
	if listDataPoliciesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listDataPoliciesOptions.AuthInstanceID))
	}

	if listDataPoliciesOptions.CatalogName != nil {
		builder.AddQuery("catalog_name", fmt.Sprint(*listDataPoliciesOptions.CatalogName))
	}
	if listDataPoliciesOptions.Status != nil {
		builder.AddQuery("status", fmt.Sprint(*listDataPoliciesOptions.Status))
	}
	if listDataPoliciesOptions.IncludeMetadata != nil {
		builder.AddQuery("include_metadata", fmt.Sprint(*listDataPoliciesOptions.IncludeMetadata))
	}
	if listDataPoliciesOptions.IncludeRules != nil {
		builder.AddQuery("include_rules", fmt.Sprint(*listDataPoliciesOptions.IncludeRules))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicyListSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDataPolicy : Create new data policy
// Create new data policy.
func (watsonxData *WatsonxDataV1) CreateDataPolicy(createDataPolicyOptions *CreateDataPolicyOptions) (result *CreateDataPolicyCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDataPolicyWithContext(context.Background(), createDataPolicyOptions)
}

// CreateDataPolicyWithContext is an alternate form of the CreateDataPolicy method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateDataPolicyWithContext(ctx context.Context, createDataPolicyOptions *CreateDataPolicyOptions) (result *CreateDataPolicyCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataPolicyOptions, "createDataPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDataPolicyOptions, "createDataPolicyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDataPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateDataPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDataPolicyOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createDataPolicyOptions.LhInstanceID))
	}
	if createDataPolicyOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDataPolicyOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDataPolicyOptions.CatalogName != nil {
		body["catalog_name"] = createDataPolicyOptions.CatalogName
	}
	if createDataPolicyOptions.DataArtifact != nil {
		body["data_artifact"] = createDataPolicyOptions.DataArtifact
	}
	if createDataPolicyOptions.PolicyName != nil {
		body["policy_name"] = createDataPolicyOptions.PolicyName
	}
	if createDataPolicyOptions.Rules != nil {
		body["rules"] = createDataPolicyOptions.Rules
	}
	if createDataPolicyOptions.Description != nil {
		body["description"] = createDataPolicyOptions.Description
	}
	if createDataPolicyOptions.Status != nil {
		body["status"] = createDataPolicyOptions.Status
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateDataPolicyCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataPolicies : Revoke data policy access management policy
// You require catalog can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteDataPolicies(deleteDataPoliciesOptions *DeleteDataPoliciesOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDataPoliciesWithContext(context.Background(), deleteDataPoliciesOptions)
}

// DeleteDataPoliciesWithContext is an alternate form of the DeleteDataPolicies method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteDataPoliciesWithContext(ctx context.Context, deleteDataPoliciesOptions *DeleteDataPoliciesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataPoliciesOptions, "deleteDataPoliciesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDataPoliciesOptions, "deleteDataPoliciesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDataPoliciesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteDataPolicies")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteDataPoliciesOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteDataPoliciesOptions.LhInstanceID))
	}
	if deleteDataPoliciesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDataPoliciesOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteDataPoliciesOptions.DataPolicies != nil {
		body["data_policies"] = deleteDataPoliciesOptions.DataPolicies
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// GetEngineUsers : Get permission in the engine
// Get users and groups permission in the engine.
func (watsonxData *WatsonxDataV1) GetEngineUsers(getEngineUsersOptions *GetEngineUsersOptions) (result *GetEngineUsersSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetEngineUsersWithContext(context.Background(), getEngineUsersOptions)
}

// GetEngineUsersWithContext is an alternate form of the GetEngineUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetEngineUsersWithContext(ctx context.Context, getEngineUsersOptions *GetEngineUsersOptions) (result *GetEngineUsersSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEngineUsersOptions, "getEngineUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEngineUsersOptions, "getEngineUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getEngineUsersOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEngineUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetEngineUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getEngineUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getEngineUsersOptions.LhInstanceID))
	}
	if getEngineUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getEngineUsersOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetEngineUsersSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteEngineUsers : Revoke permission to access engine
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteEngineUsers(deleteEngineUsersOptions *DeleteEngineUsersOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteEngineUsersWithContext(context.Background(), deleteEngineUsersOptions)
}

// DeleteEngineUsersWithContext is an alternate form of the DeleteEngineUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteEngineUsersWithContext(ctx context.Context, deleteEngineUsersOptions *DeleteEngineUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEngineUsersOptions, "deleteEngineUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEngineUsersOptions, "deleteEngineUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteEngineUsersOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEngineUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteEngineUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteEngineUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteEngineUsersOptions.LhInstanceID))
	}
	if deleteEngineUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteEngineUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteEngineUsersOptions.Groups != nil {
		body["groups"] = deleteEngineUsersOptions.Groups
	}
	if deleteEngineUsersOptions.Users != nil {
		body["users"] = deleteEngineUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateEngineUsers : Updates user and groups permission in the engine
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) UpdateEngineUsers(updateEngineUsersOptions *UpdateEngineUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateEngineUsersWithContext(context.Background(), updateEngineUsersOptions)
}

// UpdateEngineUsersWithContext is an alternate form of the UpdateEngineUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateEngineUsersWithContext(ctx context.Context, updateEngineUsersOptions *UpdateEngineUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEngineUsersOptions, "updateEngineUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEngineUsersOptions, "updateEngineUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateEngineUsersOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/engines/{engine_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEngineUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateEngineUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateEngineUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*updateEngineUsersOptions.LhInstanceID))
	}
	if updateEngineUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateEngineUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateEngineUsersOptions.Groups != nil {
		body["groups"] = updateEngineUsersOptions.Groups
	}
	if updateEngineUsersOptions.Users != nil {
		body["users"] = updateEngineUsersOptions.Users
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

// DeleteDbConnUsers : Revoke permission to access db connection
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteDbConnUsers(deleteDbConnUsersOptions *DeleteDbConnUsersOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDbConnUsersWithContext(context.Background(), deleteDbConnUsersOptions)
}

// DeleteDbConnUsersWithContext is an alternate form of the DeleteDbConnUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteDbConnUsersWithContext(ctx context.Context, deleteDbConnUsersOptions *DeleteDbConnUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDbConnUsersOptions, "deleteDbConnUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDbConnUsersOptions, "deleteDbConnUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *deleteDbConnUsersOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/databases/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDbConnUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteDbConnUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteDbConnUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteDbConnUsersOptions.LhInstanceID))
	}
	if deleteDbConnUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDbConnUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteDbConnUsersOptions.Groups != nil {
		body["groups"] = deleteDbConnUsersOptions.Groups
	}
	if deleteDbConnUsersOptions.Users != nil {
		body["users"] = deleteDbConnUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateDbConnUsers : Updates user and groups permission in the db connection
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) UpdateDbConnUsers(updateDbConnUsersOptions *UpdateDbConnUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateDbConnUsersWithContext(context.Background(), updateDbConnUsersOptions)
}

// UpdateDbConnUsersWithContext is an alternate form of the UpdateDbConnUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateDbConnUsersWithContext(ctx context.Context, updateDbConnUsersOptions *UpdateDbConnUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDbConnUsersOptions, "updateDbConnUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDbConnUsersOptions, "updateDbConnUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *updateDbConnUsersOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/databases/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDbConnUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateDbConnUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateDbConnUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*updateDbConnUsersOptions.LhInstanceID))
	}
	if updateDbConnUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDbConnUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateDbConnUsersOptions.Groups != nil {
		body["groups"] = updateDbConnUsersOptions.Groups
	}
	if updateDbConnUsersOptions.Users != nil {
		body["users"] = updateDbConnUsersOptions.Users
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

// GetDbConnUsers : Get permission in the db connection
// Get users and groups permission in the db connection.
func (watsonxData *WatsonxDataV1) GetDbConnUsers(getDbConnUsersOptions *GetDbConnUsersOptions) (result *GetDbConnUsersSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetDbConnUsersWithContext(context.Background(), getDbConnUsersOptions)
}

// GetDbConnUsersWithContext is an alternate form of the GetDbConnUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetDbConnUsersWithContext(ctx context.Context, getDbConnUsersOptions *GetDbConnUsersOptions) (result *GetDbConnUsersSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDbConnUsersOptions, "getDbConnUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDbConnUsersOptions, "getDbConnUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *getDbConnUsersOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/databases/{database_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDbConnUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetDbConnUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDbConnUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getDbConnUsersOptions.LhInstanceID))
	}
	if getDbConnUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDbConnUsersOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDbConnUsersSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCatalogUsers : Grant users and groups permission to the catalog
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) CreateCatalogUsers(createCatalogUsersOptions *CreateCatalogUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateCatalogUsersWithContext(context.Background(), createCatalogUsersOptions)
}

// CreateCatalogUsersWithContext is an alternate form of the CreateCatalogUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateCatalogUsersWithContext(ctx context.Context, createCatalogUsersOptions *CreateCatalogUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCatalogUsersOptions, "createCatalogUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCatalogUsersOptions, "createCatalogUsersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/catalogs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCatalogUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateCatalogUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createCatalogUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createCatalogUsersOptions.LhInstanceID))
	}
	if createCatalogUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createCatalogUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createCatalogUsersOptions.CatalogName != nil {
		body["catalog_name"] = createCatalogUsersOptions.CatalogName
	}
	if createCatalogUsersOptions.Groups != nil {
		body["groups"] = createCatalogUsersOptions.Groups
	}
	if createCatalogUsersOptions.Users != nil {
		body["users"] = createCatalogUsersOptions.Users
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

// GetCatalogUsers : Get users and groups permission in the catalog
// Get users and groups permission in the catalog.
func (watsonxData *WatsonxDataV1) GetCatalogUsers(getCatalogUsersOptions *GetCatalogUsersOptions) (result *GetCatalogUsersSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetCatalogUsersWithContext(context.Background(), getCatalogUsersOptions)
}

// GetCatalogUsersWithContext is an alternate form of the GetCatalogUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetCatalogUsersWithContext(ctx context.Context, getCatalogUsersOptions *GetCatalogUsersOptions) (result *GetCatalogUsersSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCatalogUsersOptions, "getCatalogUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCatalogUsersOptions, "getCatalogUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_name": *getCatalogUsersOptions.CatalogName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/catalogs/{catalog_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetCatalogUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getCatalogUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getCatalogUsersOptions.LhInstanceID))
	}
	if getCatalogUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getCatalogUsersOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetCatalogUsersSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCatalogUsers : Revoke multiple users and groups permission to access catalog
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteCatalogUsers(deleteCatalogUsersOptions *DeleteCatalogUsersOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteCatalogUsersWithContext(context.Background(), deleteCatalogUsersOptions)
}

// DeleteCatalogUsersWithContext is an alternate form of the DeleteCatalogUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteCatalogUsersWithContext(ctx context.Context, deleteCatalogUsersOptions *DeleteCatalogUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCatalogUsersOptions, "deleteCatalogUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCatalogUsersOptions, "deleteCatalogUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_name": *deleteCatalogUsersOptions.CatalogName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/catalogs/{catalog_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCatalogUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteCatalogUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteCatalogUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteCatalogUsersOptions.LhInstanceID))
	}
	if deleteCatalogUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteCatalogUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteCatalogUsersOptions.Groups != nil {
		body["groups"] = deleteCatalogUsersOptions.Groups
	}
	if deleteCatalogUsersOptions.Users != nil {
		body["users"] = deleteCatalogUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateCatalogUsers : Updates user and groups permission in the catalog
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) UpdateCatalogUsers(updateCatalogUsersOptions *UpdateCatalogUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateCatalogUsersWithContext(context.Background(), updateCatalogUsersOptions)
}

// UpdateCatalogUsersWithContext is an alternate form of the UpdateCatalogUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateCatalogUsersWithContext(ctx context.Context, updateCatalogUsersOptions *UpdateCatalogUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCatalogUsersOptions, "updateCatalogUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCatalogUsersOptions, "updateCatalogUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"catalog_name": *updateCatalogUsersOptions.CatalogName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/catalogs/{catalog_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCatalogUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateCatalogUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateCatalogUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*updateCatalogUsersOptions.LhInstanceID))
	}
	if updateCatalogUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateCatalogUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateCatalogUsersOptions.Groups != nil {
		body["groups"] = updateCatalogUsersOptions.Groups
	}
	if updateCatalogUsersOptions.Users != nil {
		body["users"] = updateCatalogUsersOptions.Users
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

// Evaluate : Evaluate permission
// Evaluate user has permission to access resource or not.
func (watsonxData *WatsonxDataV1) Evaluate(evaluateOptions *EvaluateOptions) (result *EvaluationResultSchema, response *core.DetailedResponse, err error) {
	return watsonxData.EvaluateWithContext(context.Background(), evaluateOptions)
}

// EvaluateWithContext is an alternate form of the Evaluate method which supports a Context parameter
func (watsonxData *WatsonxDataV1) EvaluateWithContext(ctx context.Context, evaluateOptions *EvaluateOptions) (result *EvaluationResultSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(evaluateOptions, "evaluateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(evaluateOptions, "evaluateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/evaluation`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range evaluateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "Evaluate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if evaluateOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*evaluateOptions.LhInstanceID))
	}
	if evaluateOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*evaluateOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if evaluateOptions.Resources != nil {
		body["resources"] = evaluateOptions.Resources
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEvaluationResultSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPoliciesList : Get policies for specific catalog in catalog_name list
// Get policies list.
func (watsonxData *WatsonxDataV1) GetPoliciesList(getPoliciesListOptions *GetPoliciesListOptions) (result *PolicySchemaList, response *core.DetailedResponse, err error) {
	return watsonxData.GetPoliciesListWithContext(context.Background(), getPoliciesListOptions)
}

// GetPoliciesListWithContext is an alternate form of the GetPoliciesList method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetPoliciesListWithContext(ctx context.Context, getPoliciesListOptions *GetPoliciesListOptions) (result *PolicySchemaList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getPoliciesListOptions, "getPoliciesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPoliciesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetPoliciesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPoliciesListOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getPoliciesListOptions.LhInstanceID))
	}
	if getPoliciesListOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPoliciesListOptions.AuthInstanceID))
	}

	if getPoliciesListOptions.CatalogList != nil {
		builder.AddQuery("catalog_list", strings.Join(getPoliciesListOptions.CatalogList, ","))
	}
	if getPoliciesListOptions.EngineList != nil {
		builder.AddQuery("engine_list", strings.Join(getPoliciesListOptions.EngineList, ","))
	}
	if getPoliciesListOptions.DataPoliciesList != nil {
		builder.AddQuery("data_policies_list", strings.Join(getPoliciesListOptions.DataPoliciesList, ","))
	}
	if getPoliciesListOptions.IncludeDataPolicies != nil {
		builder.AddQuery("include_data_policies", fmt.Sprint(*getPoliciesListOptions.IncludeDataPolicies))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicySchemaList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateMetastoreUsers : Grant users and groups permission to the metastore
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) CreateMetastoreUsers(createMetastoreUsersOptions *CreateMetastoreUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateMetastoreUsersWithContext(context.Background(), createMetastoreUsersOptions)
}

// CreateMetastoreUsersWithContext is an alternate form of the CreateMetastoreUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateMetastoreUsersWithContext(ctx context.Context, createMetastoreUsersOptions *CreateMetastoreUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMetastoreUsersOptions, "createMetastoreUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createMetastoreUsersOptions, "createMetastoreUsersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/metastores`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createMetastoreUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateMetastoreUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMetastoreUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createMetastoreUsersOptions.LhInstanceID))
	}
	if createMetastoreUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createMetastoreUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createMetastoreUsersOptions.MetastoreName != nil {
		body["metastore_name"] = createMetastoreUsersOptions.MetastoreName
	}
	if createMetastoreUsersOptions.Groups != nil {
		body["groups"] = createMetastoreUsersOptions.Groups
	}
	if createMetastoreUsersOptions.Users != nil {
		body["users"] = createMetastoreUsersOptions.Users
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

// GetMetastoreUsers : Get permission in the metastore
// Get users and groups permission in the metastore.
func (watsonxData *WatsonxDataV1) GetMetastoreUsers(getMetastoreUsersOptions *GetMetastoreUsersOptions) (result *GetMetastoreUsersSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetMetastoreUsersWithContext(context.Background(), getMetastoreUsersOptions)
}

// GetMetastoreUsersWithContext is an alternate form of the GetMetastoreUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetMetastoreUsersWithContext(ctx context.Context, getMetastoreUsersOptions *GetMetastoreUsersOptions) (result *GetMetastoreUsersSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMetastoreUsersOptions, "getMetastoreUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMetastoreUsersOptions, "getMetastoreUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"metastore_name": *getMetastoreUsersOptions.MetastoreName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/metastores/{metastore_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMetastoreUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetMetastoreUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMetastoreUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getMetastoreUsersOptions.LhInstanceID))
	}
	if getMetastoreUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getMetastoreUsersOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMetastoreUsersSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteMetastoreUsers : Revoke permission to access metastore
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteMetastoreUsers(deleteMetastoreUsersOptions *DeleteMetastoreUsersOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteMetastoreUsersWithContext(context.Background(), deleteMetastoreUsersOptions)
}

// DeleteMetastoreUsersWithContext is an alternate form of the DeleteMetastoreUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteMetastoreUsersWithContext(ctx context.Context, deleteMetastoreUsersOptions *DeleteMetastoreUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMetastoreUsersOptions, "deleteMetastoreUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteMetastoreUsersOptions, "deleteMetastoreUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"metastore_name": *deleteMetastoreUsersOptions.MetastoreName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/metastores/{metastore_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteMetastoreUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteMetastoreUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteMetastoreUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteMetastoreUsersOptions.LhInstanceID))
	}
	if deleteMetastoreUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteMetastoreUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteMetastoreUsersOptions.Groups != nil {
		body["groups"] = deleteMetastoreUsersOptions.Groups
	}
	if deleteMetastoreUsersOptions.Users != nil {
		body["users"] = deleteMetastoreUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateMetastoreUsers : Updates user and groups permission in the metastore
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) UpdateMetastoreUsers(updateMetastoreUsersOptions *UpdateMetastoreUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateMetastoreUsersWithContext(context.Background(), updateMetastoreUsersOptions)
}

// UpdateMetastoreUsersWithContext is an alternate form of the UpdateMetastoreUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateMetastoreUsersWithContext(ctx context.Context, updateMetastoreUsersOptions *UpdateMetastoreUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMetastoreUsersOptions, "updateMetastoreUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMetastoreUsersOptions, "updateMetastoreUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"metastore_name": *updateMetastoreUsersOptions.MetastoreName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/metastores/{metastore_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMetastoreUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateMetastoreUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateMetastoreUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*updateMetastoreUsersOptions.LhInstanceID))
	}
	if updateMetastoreUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateMetastoreUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateMetastoreUsersOptions.Groups != nil {
		body["groups"] = updateMetastoreUsersOptions.Groups
	}
	if updateMetastoreUsersOptions.Users != nil {
		body["users"] = updateMetastoreUsersOptions.Users
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

// CreateBucketUsers : Grant users and groups permission to the bucket
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) CreateBucketUsers(createBucketUsersOptions *CreateBucketUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateBucketUsersWithContext(context.Background(), createBucketUsersOptions)
}

// CreateBucketUsersWithContext is an alternate form of the CreateBucketUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateBucketUsersWithContext(ctx context.Context, createBucketUsersOptions *CreateBucketUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBucketUsersOptions, "createBucketUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBucketUsersOptions, "createBucketUsersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/buckets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBucketUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateBucketUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createBucketUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createBucketUsersOptions.LhInstanceID))
	}
	if createBucketUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createBucketUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createBucketUsersOptions.BucketID != nil {
		body["bucket_id"] = createBucketUsersOptions.BucketID
	}
	if createBucketUsersOptions.Groups != nil {
		body["groups"] = createBucketUsersOptions.Groups
	}
	if createBucketUsersOptions.Users != nil {
		body["users"] = createBucketUsersOptions.Users
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

// GetDefaultPolicies : Get AMS default policies
// Get AMS default policies.
func (watsonxData *WatsonxDataV1) GetDefaultPolicies(getDefaultPoliciesOptions *GetDefaultPoliciesOptions) (result *DefaultPolicySchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetDefaultPoliciesWithContext(context.Background(), getDefaultPoliciesOptions)
}

// GetDefaultPoliciesWithContext is an alternate form of the GetDefaultPolicies method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetDefaultPoliciesWithContext(ctx context.Context, getDefaultPoliciesOptions *GetDefaultPoliciesOptions) (result *DefaultPolicySchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getDefaultPoliciesOptions, "getDefaultPoliciesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/default_policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDefaultPoliciesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetDefaultPolicies")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDefaultPoliciesOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getDefaultPoliciesOptions.LhInstanceID))
	}
	if getDefaultPoliciesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDefaultPoliciesOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDefaultPolicySchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPolicyVersion : Get AMS policies version
// Get AMS policies version.
func (watsonxData *WatsonxDataV1) GetPolicyVersion(getPolicyVersionOptions *GetPolicyVersionOptions) (result *PolicyVersionResultSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetPolicyVersionWithContext(context.Background(), getPolicyVersionOptions)
}

// GetPolicyVersionWithContext is an alternate form of the GetPolicyVersion method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetPolicyVersionWithContext(ctx context.Context, getPolicyVersionOptions *GetPolicyVersionOptions) (result *PolicyVersionResultSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getPolicyVersionOptions, "getPolicyVersionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/policy_versions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPolicyVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetPolicyVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPolicyVersionOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getPolicyVersionOptions.LhInstanceID))
	}
	if getPolicyVersionOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPolicyVersionOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicyVersionResultSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDataPolicy : Get policy
// Get policy detail.
func (watsonxData *WatsonxDataV1) GetDataPolicy(getDataPolicyOptions *GetDataPolicyOptions) (result *PolicySchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetDataPolicyWithContext(context.Background(), getDataPolicyOptions)
}

// GetDataPolicyWithContext is an alternate form of the GetDataPolicy method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetDataPolicyWithContext(ctx context.Context, getDataPolicyOptions *GetDataPolicyOptions) (result *PolicySchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataPolicyOptions, "getDataPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDataPolicyOptions, "getDataPolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_name": *getDataPolicyOptions.PolicyName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies/{policy_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetDataPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDataPolicyOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getDataPolicyOptions.LhInstanceID))
	}
	if getDataPolicyOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDataPolicyOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicySchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceDataPolicy : Updates data policy
// You require catalog can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) ReplaceDataPolicy(replaceDataPolicyOptions *ReplaceDataPolicyOptions) (result *ReplaceDataPolicyCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.ReplaceDataPolicyWithContext(context.Background(), replaceDataPolicyOptions)
}

// ReplaceDataPolicyWithContext is an alternate form of the ReplaceDataPolicy method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ReplaceDataPolicyWithContext(ctx context.Context, replaceDataPolicyOptions *ReplaceDataPolicyOptions) (result *ReplaceDataPolicyCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceDataPolicyOptions, "replaceDataPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceDataPolicyOptions, "replaceDataPolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_name": *replaceDataPolicyOptions.PolicyName,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies/{policy_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceDataPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ReplaceDataPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceDataPolicyOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*replaceDataPolicyOptions.LhInstanceID))
	}
	if replaceDataPolicyOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*replaceDataPolicyOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if replaceDataPolicyOptions.CatalogName != nil {
		body["catalog_name"] = replaceDataPolicyOptions.CatalogName
	}
	if replaceDataPolicyOptions.DataArtifact != nil {
		body["data_artifact"] = replaceDataPolicyOptions.DataArtifact
	}
	if replaceDataPolicyOptions.Rules != nil {
		body["rules"] = replaceDataPolicyOptions.Rules
	}
	if replaceDataPolicyOptions.Description != nil {
		body["description"] = replaceDataPolicyOptions.Description
	}
	if replaceDataPolicyOptions.Status != nil {
		body["status"] = replaceDataPolicyOptions.Status
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReplaceDataPolicyCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataPolicy : Revoke data policy access management policy
// You require catalog can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteDataPolicy(deleteDataPolicyOptions *DeleteDataPolicyOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDataPolicyWithContext(context.Background(), deleteDataPolicyOptions)
}

// DeleteDataPolicyWithContext is an alternate form of the DeleteDataPolicy method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteDataPolicyWithContext(ctx context.Context, deleteDataPolicyOptions *DeleteDataPolicyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataPolicyOptions, "deleteDataPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDataPolicyOptions, "deleteDataPolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_name": *deleteDataPolicyOptions.PolicyName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/data_policies/{policy_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDataPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteDataPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDataPolicyOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteDataPolicyOptions.LhInstanceID))
	}
	if deleteDataPolicyOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDataPolicyOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// CreateEngineUsers : Grant permission to the engine
// You require administrator role or can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) CreateEngineUsers(createEngineUsersOptions *CreateEngineUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEngineUsersWithContext(context.Background(), createEngineUsersOptions)
}

// CreateEngineUsersWithContext is an alternate form of the CreateEngineUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateEngineUsersWithContext(ctx context.Context, createEngineUsersOptions *CreateEngineUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEngineUsersOptions, "createEngineUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEngineUsersOptions, "createEngineUsersOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/engines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEngineUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateEngineUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createEngineUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*createEngineUsersOptions.LhInstanceID))
	}
	if createEngineUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEngineUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createEngineUsersOptions.EngineID != nil {
		body["engine_id"] = createEngineUsersOptions.EngineID
	}
	if createEngineUsersOptions.Groups != nil {
		body["groups"] = createEngineUsersOptions.Groups
	}
	if createEngineUsersOptions.Users != nil {
		body["users"] = createEngineUsersOptions.Users
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

// GetBucketUsers : Get permission in the bucket
// Get users and groups permission in the bucket.
func (watsonxData *WatsonxDataV1) GetBucketUsers(getBucketUsersOptions *GetBucketUsersOptions) (result *GetBucketUsersSchema, response *core.DetailedResponse, err error) {
	return watsonxData.GetBucketUsersWithContext(context.Background(), getBucketUsersOptions)
}

// GetBucketUsersWithContext is an alternate form of the GetBucketUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetBucketUsersWithContext(ctx context.Context, getBucketUsersOptions *GetBucketUsersOptions) (result *GetBucketUsersSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketUsersOptions, "getBucketUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketUsersOptions, "getBucketUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *getBucketUsersOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/buckets/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetBucketUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getBucketUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*getBucketUsersOptions.LhInstanceID))
	}
	if getBucketUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getBucketUsersOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketUsersSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteBucketUsers : Revoke permission to access bucket
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) DeleteBucketUsers(deleteBucketUsersOptions *DeleteBucketUsersOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteBucketUsersWithContext(context.Background(), deleteBucketUsersOptions)
}

// DeleteBucketUsersWithContext is an alternate form of the DeleteBucketUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteBucketUsersWithContext(ctx context.Context, deleteBucketUsersOptions *DeleteBucketUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketUsersOptions, "deleteBucketUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketUsersOptions, "deleteBucketUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *deleteBucketUsersOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/buckets/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteBucketUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteBucketUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*deleteBucketUsersOptions.LhInstanceID))
	}
	if deleteBucketUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteBucketUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteBucketUsersOptions.Groups != nil {
		body["groups"] = deleteBucketUsersOptions.Groups
	}
	if deleteBucketUsersOptions.Users != nil {
		body["users"] = deleteBucketUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateBucketUsers : Updates user and groups permission in the bucket
// You require can_administer permission to perform this action.
func (watsonxData *WatsonxDataV1) UpdateBucketUsers(updateBucketUsersOptions *UpdateBucketUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateBucketUsersWithContext(context.Background(), updateBucketUsersOptions)
}

// UpdateBucketUsersWithContext is an alternate form of the UpdateBucketUsers method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateBucketUsersWithContext(ctx context.Context, updateBucketUsersOptions *UpdateBucketUsersOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBucketUsersOptions, "updateBucketUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBucketUsersOptions, "updateBucketUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *updateBucketUsersOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/access/buckets/{bucket_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBucketUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateBucketUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateBucketUsersOptions.LhInstanceID != nil {
		builder.AddHeader("LhInstanceId", fmt.Sprint(*updateBucketUsersOptions.LhInstanceID))
	}
	if updateBucketUsersOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateBucketUsersOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateBucketUsersOptions.Groups != nil {
		body["groups"] = updateBucketUsersOptions.Groups
	}
	if updateBucketUsersOptions.Users != nil {
		body["users"] = updateBucketUsersOptions.Users
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

// GetBuckets : Get buckets
// Get list of all buckets registered to Lakehouse.
func (watsonxData *WatsonxDataV1) GetBuckets(getBucketsOptions *GetBucketsOptions) (result *GetBucketsOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetBucketsWithContext(context.Background(), getBucketsOptions)
}

// GetBucketsWithContext is an alternate form of the GetBuckets method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetBucketsWithContext(ctx context.Context, getBucketsOptions *GetBucketsOptions) (result *GetBucketsOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getBucketsOptions, "getBucketsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetBuckets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getBucketsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getBucketsOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketsOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBucketObjects : Get bucket objects
// Fetch all objects from a given bucket.
func (watsonxData *WatsonxDataV1) GetBucketObjects(getBucketObjectsOptions *GetBucketObjectsOptions) (result *GetBucketObjectsOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetBucketObjectsWithContext(context.Background(), getBucketObjectsOptions)
}

// GetBucketObjectsWithContext is an alternate form of the GetBucketObjects method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetBucketObjectsWithContext(ctx context.Context, getBucketObjectsOptions *GetBucketObjectsOptions) (result *GetBucketObjectsOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketObjectsOptions, "getBucketObjectsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketObjectsOptions, "getBucketObjectsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket/objects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetBucketObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getBucketObjectsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getBucketObjectsOptions.AuthInstanceID))
	}

	builder.AddQuery("bucket_id", fmt.Sprint(*getBucketObjectsOptions.BucketID))

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketObjectsOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeactivateBucket : Deactivate bucket
// Deactivate an active bucket in Lakehouse.
func (watsonxData *WatsonxDataV1) DeactivateBucket(deactivateBucketOptions *DeactivateBucketOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.DeactivateBucketWithContext(context.Background(), deactivateBucketOptions)
}

// DeactivateBucketWithContext is an alternate form of the DeactivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeactivateBucketWithContext(ctx context.Context, deactivateBucketOptions *DeactivateBucketOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deactivateBucketOptions, "deactivateBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deactivateBucketOptions, "deactivateBucketOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket/deactivate`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deactivateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeactivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if deactivateBucketOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*deactivateBucketOptions.Accept))
	}
	if deactivateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deactivateBucketOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deactivateBucketOptions.BucketID != nil {
		body["bucket_id"] = deactivateBucketOptions.BucketID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// RegisterBucket : Register bucket
// Register a new bucket in Lakehouse.
func (watsonxData *WatsonxDataV1) RegisterBucket(registerBucketOptions *RegisterBucketOptions) (result *RegisterBucketCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.RegisterBucketWithContext(context.Background(), registerBucketOptions)
}

// RegisterBucketWithContext is an alternate form of the RegisterBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV1) RegisterBucketWithContext(ctx context.Context, registerBucketOptions *RegisterBucketOptions) (result *RegisterBucketCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(registerBucketOptions, "registerBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(registerBucketOptions, "registerBucketOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range registerBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "RegisterBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if registerBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*registerBucketOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if registerBucketOptions.BucketDetails != nil {
		body["bucket_details"] = registerBucketOptions.BucketDetails
	}
	if registerBucketOptions.Description != nil {
		body["description"] = registerBucketOptions.Description
	}
	if registerBucketOptions.TableType != nil {
		body["table_type"] = registerBucketOptions.TableType
	}
	if registerBucketOptions.BucketType != nil {
		body["bucket_type"] = registerBucketOptions.BucketType
	}
	if registerBucketOptions.CatalogName != nil {
		body["catalog_name"] = registerBucketOptions.CatalogName
	}
	if registerBucketOptions.ManagedBy != nil {
		body["managed_by"] = registerBucketOptions.ManagedBy
	}
	if registerBucketOptions.BucketDisplayName != nil {
		body["bucket_display_name"] = registerBucketOptions.BucketDisplayName
	}
	if registerBucketOptions.BucketTags != nil {
		body["bucket_tags"] = registerBucketOptions.BucketTags
	}
	if registerBucketOptions.CatalogTags != nil {
		body["catalog_tags"] = registerBucketOptions.CatalogTags
	}
	if registerBucketOptions.ThriftURI != nil {
		body["thrift_uri"] = registerBucketOptions.ThriftURI
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRegisterBucketCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UnregisterBucket : Unregister Bucket
// Unregister a bucket from Lakehouse.
func (watsonxData *WatsonxDataV1) UnregisterBucket(unregisterBucketOptions *UnregisterBucketOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.UnregisterBucketWithContext(context.Background(), unregisterBucketOptions)
}

// UnregisterBucketWithContext is an alternate form of the UnregisterBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UnregisterBucketWithContext(ctx context.Context, unregisterBucketOptions *UnregisterBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unregisterBucketOptions, "unregisterBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unregisterBucketOptions, "unregisterBucketOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range unregisterBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UnregisterBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if unregisterBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*unregisterBucketOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if unregisterBucketOptions.BucketID != nil {
		body["bucket_id"] = unregisterBucketOptions.BucketID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateBucket : Update bucket
// Update bucket details/credentials.
func (watsonxData *WatsonxDataV1) UpdateBucket(updateBucketOptions *UpdateBucketOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateBucketWithContext(context.Background(), updateBucketOptions)
}

// UpdateBucketWithContext is an alternate form of the UpdateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateBucketWithContext(ctx context.Context, updateBucketOptions *UpdateBucketOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBucketOptions, "updateBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBucketOptions, "updateBucketOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateBucketOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateBucketOptions.BucketID != nil {
		body["bucket_id"] = updateBucketOptions.BucketID
	}
	if updateBucketOptions.AccessKey != nil {
		body["access_key"] = updateBucketOptions.AccessKey
	}
	if updateBucketOptions.BucketDisplayName != nil {
		body["bucket_display_name"] = updateBucketOptions.BucketDisplayName
	}
	if updateBucketOptions.Description != nil {
		body["description"] = updateBucketOptions.Description
	}
	if updateBucketOptions.SecretKey != nil {
		body["secret_key"] = updateBucketOptions.SecretKey
	}
	if updateBucketOptions.Tags != nil {
		body["tags"] = updateBucketOptions.Tags
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

// ActivateBucket : Active bucket
// Activate an invalid bucket in Lakehouse.
func (watsonxData *WatsonxDataV1) ActivateBucket(activateBucketOptions *ActivateBucketOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.ActivateBucketWithContext(context.Background(), activateBucketOptions)
}

// ActivateBucketWithContext is an alternate form of the ActivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ActivateBucketWithContext(ctx context.Context, activateBucketOptions *ActivateBucketOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(activateBucketOptions, "activateBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(activateBucketOptions, "activateBucketOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/buckets/bucket/activate`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range activateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ActivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if activateBucketOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*activateBucketOptions.Accept))
	}
	if activateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*activateBucketOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if activateBucketOptions.BucketID != nil {
		body["bucket_id"] = activateBucketOptions.BucketID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// GetDatabases : Get databases
// Get list of all databases in Lakehouse.
func (watsonxData *WatsonxDataV1) GetDatabases(getDatabasesOptions *GetDatabasesOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.GetDatabasesWithContext(context.Background(), getDatabasesOptions)
}

// GetDatabasesWithContext is an alternate form of the GetDatabases method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetDatabasesWithContext(ctx context.Context, getDatabasesOptions *GetDatabasesOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getDatabasesOptions, "getDatabasesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/databases`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDatabasesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetDatabases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if getDatabasesOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getDatabasesOptions.Accept))
	}
	if getDatabasesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDatabasesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// CreateDatabaseCatalog : Add/Create database
// Add or create a new database in Lakehouse.
func (watsonxData *WatsonxDataV1) CreateDatabaseCatalog(createDatabaseCatalogOptions *CreateDatabaseCatalogOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.CreateDatabaseCatalogWithContext(context.Background(), createDatabaseCatalogOptions)
}

// CreateDatabaseCatalogWithContext is an alternate form of the CreateDatabaseCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateDatabaseCatalogWithContext(ctx context.Context, createDatabaseCatalogOptions *CreateDatabaseCatalogOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDatabaseCatalogOptions, "createDatabaseCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDatabaseCatalogOptions, "createDatabaseCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/databases/database`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDatabaseCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateDatabaseCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if createDatabaseCatalogOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*createDatabaseCatalogOptions.Accept))
	}
	if createDatabaseCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDatabaseCatalogOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDatabaseCatalogOptions.DatabaseDisplayName != nil {
		body["database_display_name"] = createDatabaseCatalogOptions.DatabaseDisplayName
	}
	if createDatabaseCatalogOptions.DatabaseType != nil {
		body["database_type"] = createDatabaseCatalogOptions.DatabaseType
	}
	if createDatabaseCatalogOptions.CatalogName != nil {
		body["catalog_name"] = createDatabaseCatalogOptions.CatalogName
	}
	if createDatabaseCatalogOptions.DatabaseDetails != nil {
		body["database_details"] = createDatabaseCatalogOptions.DatabaseDetails
	}
	if createDatabaseCatalogOptions.Description != nil {
		body["description"] = createDatabaseCatalogOptions.Description
	}
	if createDatabaseCatalogOptions.Tags != nil {
		body["tags"] = createDatabaseCatalogOptions.Tags
	}
	if createDatabaseCatalogOptions.CreatedBy != nil {
		body["created_by"] = createDatabaseCatalogOptions.CreatedBy
	}
	if createDatabaseCatalogOptions.CreatedOn != nil {
		body["created_on"] = createDatabaseCatalogOptions.CreatedOn
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// DeleteDatabaseCatalog : Delete database
// Delete a database from Lakehouse.
func (watsonxData *WatsonxDataV1) DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteDatabaseCatalogWithContext(context.Background(), deleteDatabaseCatalogOptions)
}

// DeleteDatabaseCatalogWithContext is an alternate form of the DeleteDatabaseCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteDatabaseCatalogWithContext(ctx context.Context, deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/databases/database`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDatabaseCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteDatabaseCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteDatabaseCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDatabaseCatalogOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteDatabaseCatalogOptions.DatabaseID != nil {
		body["database_id"] = deleteDatabaseCatalogOptions.DatabaseID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
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
func (watsonxData *WatsonxDataV1) UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateDatabaseWithContext(context.Background(), updateDatabaseOptions)
}

// UpdateDatabaseWithContext is an alternate form of the UpdateDatabase method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateDatabaseWithContext(ctx context.Context, updateDatabaseOptions *UpdateDatabaseOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDatabaseOptions, "updateDatabaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDatabaseOptions, "updateDatabaseOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/databases/database`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDatabaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateDatabase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if updateDatabaseOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*updateDatabaseOptions.Accept))
	}
	if updateDatabaseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDatabaseOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateDatabaseOptions.DatabaseID != nil {
		body["database_id"] = updateDatabaseOptions.DatabaseID
	}
	if updateDatabaseOptions.DatabaseDetails != nil {
		body["database_details"] = updateDatabaseOptions.DatabaseDetails
	}
	if updateDatabaseOptions.DatabaseDisplayName != nil {
		body["database_display_name"] = updateDatabaseOptions.DatabaseDisplayName
	}
	if updateDatabaseOptions.Description != nil {
		body["description"] = updateDatabaseOptions.Description
	}
	if updateDatabaseOptions.Tags != nil {
		body["tags"] = updateDatabaseOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// PauseEngine : Pause engine
// Pause a running engine.
func (watsonxData *WatsonxDataV1) PauseEngine(pauseEngineOptions *PauseEngineOptions) (result *PauseEngineCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.PauseEngineWithContext(context.Background(), pauseEngineOptions)
}

// PauseEngineWithContext is an alternate form of the PauseEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) PauseEngineWithContext(ctx context.Context, pauseEngineOptions *PauseEngineOptions) (result *PauseEngineCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pauseEngineOptions, "pauseEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pauseEngineOptions, "pauseEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines/engine/pause`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range pauseEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "PauseEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if pauseEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*pauseEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if pauseEngineOptions.EngineID != nil {
		body["engine_id"] = pauseEngineOptions.EngineID
	}
	if pauseEngineOptions.CreatedBy != nil {
		body["created_by"] = pauseEngineOptions.CreatedBy
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPauseEngineCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEngines : Get engines
// Get all engine details.
func (watsonxData *WatsonxDataV1) GetEngines(getEnginesOptions *GetEnginesOptions) (result *GetEnginesOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetEnginesWithContext(context.Background(), getEnginesOptions)
}

// GetEnginesWithContext is an alternate form of the GetEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetEnginesWithContext(ctx context.Context, getEnginesOptions *GetEnginesOptions) (result *GetEnginesOKBody, response *core.DetailedResponse, err error) {
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

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetEngines")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetEnginesOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDeployments : Get instance details
// Get instance details.
func (watsonxData *WatsonxDataV1) GetDeployments(getDeploymentsOptions *GetDeploymentsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.GetDeploymentsWithContext(context.Background(), getDeploymentsOptions)
}

// GetDeploymentsWithContext is an alternate form of the GetDeployments method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetDeploymentsWithContext(ctx context.Context, getDeploymentsOptions *GetDeploymentsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
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

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetDeployments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if getDeploymentsOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getDeploymentsOptions.Accept))
	}
	if getDeploymentsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDeploymentsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// UpdateEngine : Update engine
// Update engine.
func (watsonxData *WatsonxDataV1) UpdateEngine(updateEngineOptions *UpdateEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateEngineWithContext(context.Background(), updateEngineOptions)
}

// UpdateEngineWithContext is an alternate form of the UpdateEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateEngineWithContext(ctx context.Context, updateEngineOptions *UpdateEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEngineOptions, "updateEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEngineOptions, "updateEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines/engine`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if updateEngineOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*updateEngineOptions.Accept))
	}
	if updateEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateEngineOptions.EngineID != nil {
		body["engine_id"] = updateEngineOptions.EngineID
	}
	if updateEngineOptions.Coordinator != nil {
		body["coordinator"] = updateEngineOptions.Coordinator
	}
	if updateEngineOptions.Description != nil {
		body["description"] = updateEngineOptions.Description
	}
	if updateEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = updateEngineOptions.EngineDisplayName
	}
	if updateEngineOptions.Tags != nil {
		body["tags"] = updateEngineOptions.Tags
	}
	if updateEngineOptions.Worker != nil {
		body["worker"] = updateEngineOptions.Worker
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// CreateEngine : Create engine
// Create a new engine.
func (watsonxData *WatsonxDataV1) CreateEngine(createEngineOptions *CreateEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.CreateEngineWithContext(context.Background(), createEngineOptions)
}

// CreateEngineWithContext is an alternate form of the CreateEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateEngineWithContext(ctx context.Context, createEngineOptions *CreateEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEngineOptions, "createEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEngineOptions, "createEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines/engine`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if createEngineOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*createEngineOptions.Accept))
	}
	if createEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createEngineOptions.Version != nil {
		body["version"] = createEngineOptions.Version
	}
	if createEngineOptions.EngineDetails != nil {
		body["engine_details"] = createEngineOptions.EngineDetails
	}
	if createEngineOptions.Origin != nil {
		body["origin"] = createEngineOptions.Origin
	}
	if createEngineOptions.Type != nil {
		body["type"] = createEngineOptions.Type
	}
	if createEngineOptions.Description != nil {
		body["description"] = createEngineOptions.Description
	}
	if createEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createEngineOptions.EngineDisplayName
	}
	if createEngineOptions.FirstTimeUse != nil {
		body["first_time_use"] = createEngineOptions.FirstTimeUse
	}
	if createEngineOptions.Region != nil {
		body["region"] = createEngineOptions.Region
	}
	if createEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createEngineOptions.AssociatedCatalogs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// DeleteEngine : Delete engine
// Delete an engine from lakehouse.
func (watsonxData *WatsonxDataV1) DeleteEngine(deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteEngineWithContext(context.Background(), deleteEngineOptions)
}

// DeleteEngineWithContext is an alternate form of the DeleteEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteEngineWithContext(ctx context.Context, deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEngineOptions, "deleteEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEngineOptions, "deleteEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines/engine`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteEngineOptions.EngineID != nil {
		body["engine_id"] = deleteEngineOptions.EngineID
	}
	if deleteEngineOptions.CreatedBy != nil {
		body["created_by"] = deleteEngineOptions.CreatedBy
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// ResumeEngine : Resume engine
// Resume a paused engine.
func (watsonxData *WatsonxDataV1) ResumeEngine(resumeEngineOptions *ResumeEngineOptions) (result *ResumeEngineCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.ResumeEngineWithContext(context.Background(), resumeEngineOptions)
}

// ResumeEngineWithContext is an alternate form of the ResumeEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ResumeEngineWithContext(ctx context.Context, resumeEngineOptions *ResumeEngineOptions) (result *ResumeEngineCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resumeEngineOptions, "resumeEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resumeEngineOptions, "resumeEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/engines/engine/resume`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range resumeEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ResumeEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if resumeEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*resumeEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if resumeEngineOptions.EngineID != nil {
		body["engine_id"] = resumeEngineOptions.EngineID
	}
	if resumeEngineOptions.CreatedBy != nil {
		body["created_by"] = resumeEngineOptions.CreatedBy
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResumeEngineCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExplainAnalyzeStatement : Explain analyze
// Return query metrics after query is complete.
func (watsonxData *WatsonxDataV1) ExplainAnalyzeStatement(explainAnalyzeStatementOptions *ExplainAnalyzeStatementOptions) (result *ExplainAnalyzeStatementCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.ExplainAnalyzeStatementWithContext(context.Background(), explainAnalyzeStatementOptions)
}

// ExplainAnalyzeStatementWithContext is an alternate form of the ExplainAnalyzeStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ExplainAnalyzeStatementWithContext(ctx context.Context, explainAnalyzeStatementOptions *ExplainAnalyzeStatementOptions) (result *ExplainAnalyzeStatementCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(explainAnalyzeStatementOptions, "explainAnalyzeStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(explainAnalyzeStatementOptions, "explainAnalyzeStatementOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/explainanalyze`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range explainAnalyzeStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ExplainAnalyzeStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if explainAnalyzeStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*explainAnalyzeStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if explainAnalyzeStatementOptions.CatalogName != nil {
		body["catalog_name"] = explainAnalyzeStatementOptions.CatalogName
	}
	if explainAnalyzeStatementOptions.EngineID != nil {
		body["engine_id"] = explainAnalyzeStatementOptions.EngineID
	}
	if explainAnalyzeStatementOptions.SchemaName != nil {
		body["schema_name"] = explainAnalyzeStatementOptions.SchemaName
	}
	if explainAnalyzeStatementOptions.Statement != nil {
		body["statement"] = explainAnalyzeStatementOptions.Statement
	}
	if explainAnalyzeStatementOptions.Verbose != nil {
		body["verbose"] = explainAnalyzeStatementOptions.Verbose
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExplainAnalyzeStatementCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExplainStatement : Explain
// Explain a query statement.
func (watsonxData *WatsonxDataV1) ExplainStatement(explainStatementOptions *ExplainStatementOptions) (result *ExplainStatementCreatedBody, response *core.DetailedResponse, err error) {
	return watsonxData.ExplainStatementWithContext(context.Background(), explainStatementOptions)
}

// ExplainStatementWithContext is an alternate form of the ExplainStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ExplainStatementWithContext(ctx context.Context, explainStatementOptions *ExplainStatementOptions) (result *ExplainStatementCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(explainStatementOptions, "explainStatementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(explainStatementOptions, "explainStatementOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/explain`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range explainStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ExplainStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if explainStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*explainStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if explainStatementOptions.EngineID != nil {
		body["engine_id"] = explainStatementOptions.EngineID
	}
	if explainStatementOptions.Statement != nil {
		body["statement"] = explainStatementOptions.Statement
	}
	if explainStatementOptions.CatalogName != nil {
		body["catalog_name"] = explainStatementOptions.CatalogName
	}
	if explainStatementOptions.Format != nil {
		body["format"] = explainStatementOptions.Format
	}
	if explainStatementOptions.SchemaName != nil {
		body["schema_name"] = explainStatementOptions.SchemaName
	}
	if explainStatementOptions.Type != nil {
		body["type"] = explainStatementOptions.Type
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExplainStatementCreatedBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TestLHConsole : Readiness API
// Verify lhconsole server is up and running.
func (watsonxData *WatsonxDataV1) TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.TestLHConsoleWithContext(context.Background(), testLHConsoleOptions)
}

// TestLHConsoleWithContext is an alternate form of the TestLHConsole method which supports a Context parameter
func (watsonxData *WatsonxDataV1) TestLHConsoleWithContext(ctx context.Context, testLHConsoleOptions *TestLHConsoleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
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

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "TestLHConsole")
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

// GetMetastores : Get Catalogs
// Get list of all registered metastores.
func (watsonxData *WatsonxDataV1) GetMetastores(getMetastoresOptions *GetMetastoresOptions) (result *GetMetastoresOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetMetastoresWithContext(context.Background(), getMetastoresOptions)
}

// GetMetastoresWithContext is an alternate form of the GetMetastores method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetMetastoresWithContext(ctx context.Context, getMetastoresOptions *GetMetastoresOptions) (result *GetMetastoresOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMetastoresOptions, "getMetastoresOptions")
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

	for headerName, headerValue := range getMetastoresOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetMetastores")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMetastoresOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getMetastoresOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMetastoresOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetHMS : Get Metastore
// Get list of all registered HMS metastores.
func (watsonxData *WatsonxDataV1) GetHMS(getHMSOptions *GetHMSOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.GetHMSWithContext(context.Background(), getHMSOptions)
}

// GetHMSWithContext is an alternate form of the GetHMS method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetHMSWithContext(ctx context.Context, getHMSOptions *GetHMSOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHMSOptions, "getHMSOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/metastores`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHMSOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetHMS")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if getHMSOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getHMSOptions.Accept))
	}
	if getHMSOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getHMSOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// AddMetastoreToEngine : Add catalog to engine
// Associate a catalog to an engine.
func (watsonxData *WatsonxDataV1) AddMetastoreToEngine(addMetastoreToEngineOptions *AddMetastoreToEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.AddMetastoreToEngineWithContext(context.Background(), addMetastoreToEngineOptions)
}

// AddMetastoreToEngineWithContext is an alternate form of the AddMetastoreToEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) AddMetastoreToEngineWithContext(ctx context.Context, addMetastoreToEngineOptions *AddMetastoreToEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addMetastoreToEngineOptions, "addMetastoreToEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addMetastoreToEngineOptions, "addMetastoreToEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/add_catalog_to_engine`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range addMetastoreToEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "AddMetastoreToEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if addMetastoreToEngineOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*addMetastoreToEngineOptions.Accept))
	}
	if addMetastoreToEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*addMetastoreToEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if addMetastoreToEngineOptions.CatalogName != nil {
		body["catalog_name"] = addMetastoreToEngineOptions.CatalogName
	}
	if addMetastoreToEngineOptions.EngineID != nil {
		body["engine_id"] = addMetastoreToEngineOptions.EngineID
	}
	if addMetastoreToEngineOptions.CreatedBy != nil {
		body["created_by"] = addMetastoreToEngineOptions.CreatedBy
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// RemoveCatalogFromEngine : Remove catalog from engine
// Remove a catalog from an engine.
func (watsonxData *WatsonxDataV1) RemoveCatalogFromEngine(removeCatalogFromEngineOptions *RemoveCatalogFromEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.RemoveCatalogFromEngineWithContext(context.Background(), removeCatalogFromEngineOptions)
}

// RemoveCatalogFromEngineWithContext is an alternate form of the RemoveCatalogFromEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV1) RemoveCatalogFromEngineWithContext(ctx context.Context, removeCatalogFromEngineOptions *RemoveCatalogFromEngineOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeCatalogFromEngineOptions, "removeCatalogFromEngineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeCatalogFromEngineOptions, "removeCatalogFromEngineOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/remove_catalog_from_engine`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeCatalogFromEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "RemoveCatalogFromEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if removeCatalogFromEngineOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*removeCatalogFromEngineOptions.Accept))
	}
	if removeCatalogFromEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*removeCatalogFromEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if removeCatalogFromEngineOptions.CatalogName != nil {
		body["catalog_name"] = removeCatalogFromEngineOptions.CatalogName
	}
	if removeCatalogFromEngineOptions.EngineID != nil {
		body["engine_id"] = removeCatalogFromEngineOptions.EngineID
	}
	if removeCatalogFromEngineOptions.CreatedBy != nil {
		body["created_by"] = removeCatalogFromEngineOptions.CreatedBy
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// SaveQuery : Save query
// Save a new query.
func (watsonxData *WatsonxDataV1) SaveQuery(saveQueryOptions *SaveQueryOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.SaveQueryWithContext(context.Background(), saveQueryOptions)
}

// SaveQueryWithContext is an alternate form of the SaveQuery method which supports a Context parameter
func (watsonxData *WatsonxDataV1) SaveQueryWithContext(ctx context.Context, saveQueryOptions *SaveQueryOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(saveQueryOptions, "saveQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(saveQueryOptions, "saveQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"query_name": *saveQueryOptions.QueryName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/queries/{query_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range saveQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "SaveQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if saveQueryOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*saveQueryOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if saveQueryOptions.CreatedBy != nil {
		body["created_by"] = saveQueryOptions.CreatedBy
	}
	if saveQueryOptions.Description != nil {
		body["description"] = saveQueryOptions.Description
	}
	if saveQueryOptions.QueryString != nil {
		body["query_string"] = saveQueryOptions.QueryString
	}
	if saveQueryOptions.CreatedOn != nil {
		body["created_on"] = saveQueryOptions.CreatedOn
	}
	if saveQueryOptions.EngineID != nil {
		body["engine_id"] = saveQueryOptions.EngineID
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

// DeleteQuery : Delete query
// Delete a saved query.
func (watsonxData *WatsonxDataV1) DeleteQuery(deleteQueryOptions *DeleteQueryOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteQueryWithContext(context.Background(), deleteQueryOptions)
}

// DeleteQueryWithContext is an alternate form of the DeleteQuery method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteQueryWithContext(ctx context.Context, deleteQueryOptions *DeleteQueryOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteQueryOptions, "deleteQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteQueryOptions, "deleteQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"query_name": *deleteQueryOptions.QueryName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/queries/{query_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteQueryOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteQueryOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateQuery : Update query
// Update a saved query.
func (watsonxData *WatsonxDataV1) UpdateQuery(updateQueryOptions *UpdateQueryOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateQueryWithContext(context.Background(), updateQueryOptions)
}

// UpdateQueryWithContext is an alternate form of the UpdateQuery method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateQueryWithContext(ctx context.Context, updateQueryOptions *UpdateQueryOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateQueryOptions, "updateQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateQueryOptions, "updateQueryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"query_name": *updateQueryOptions.QueryName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/queries/{query_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateQueryOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateQueryOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if updateQueryOptions.QueryString != nil {
		body["query_string"] = updateQueryOptions.QueryString
	}
	if updateQueryOptions.Description != nil {
		body["description"] = updateQueryOptions.Description
	}
	if updateQueryOptions.NewQueryName != nil {
		body["new_query_name"] = updateQueryOptions.NewQueryName
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

// GetQueries : Get queries
// List all saved queries.
func (watsonxData *WatsonxDataV1) GetQueries(getQueriesOptions *GetQueriesOptions) (result *GetQueriesOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetQueriesWithContext(context.Background(), getQueriesOptions)
}

// GetQueriesWithContext is an alternate form of the GetQueries method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetQueriesWithContext(ctx context.Context, getQueriesOptions *GetQueriesOptions) (result *GetQueriesOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getQueriesOptions, "getQueriesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/queries`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getQueriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetQueries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getQueriesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getQueriesOptions.AuthInstanceID))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetQueriesOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSchema : Create schema
// Create a new schema.
func (watsonxData *WatsonxDataV1) CreateSchema(createSchemaOptions *CreateSchemaOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.CreateSchemaWithContext(context.Background(), createSchemaOptions)
}

// CreateSchemaWithContext is an alternate form of the CreateSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV1) CreateSchemaWithContext(ctx context.Context, createSchemaOptions *CreateSchemaOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaOptions, "createSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSchemaOptions, "createSchemaOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/schemas/schema`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "CreateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSchemaOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createSchemaOptions.CatalogName != nil {
		body["catalog_name"] = createSchemaOptions.CatalogName
	}
	if createSchemaOptions.EngineID != nil {
		body["engine_id"] = createSchemaOptions.EngineID
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSchema : Delete schema
// Delete a schema.
func (watsonxData *WatsonxDataV1) DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteSchemaWithContext(context.Background(), deleteSchemaOptions)
}

// DeleteSchemaWithContext is an alternate form of the DeleteSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteSchemaWithContext(ctx context.Context, deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSchemaOptions, "deleteSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSchemaOptions, "deleteSchemaOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/schemas/schema`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSchemaOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteSchemaOptions.CatalogName != nil {
		body["catalog_name"] = deleteSchemaOptions.CatalogName
	}
	if deleteSchemaOptions.EngineID != nil {
		body["engine_id"] = deleteSchemaOptions.EngineID
	}
	if deleteSchemaOptions.SchemaName != nil {
		body["schema_name"] = deleteSchemaOptions.SchemaName
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// GetSchemas : Get schemas
// List schemas in catalog.
func (watsonxData *WatsonxDataV1) GetSchemas(getSchemasOptions *GetSchemasOptions) (result *GetSchemasOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetSchemasWithContext(context.Background(), getSchemasOptions)
}

// GetSchemasWithContext is an alternate form of the GetSchemas method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetSchemasWithContext(ctx context.Context, getSchemasOptions *GetSchemasOptions) (result *GetSchemasOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSchemasOptions, "getSchemasOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSchemasOptions, "getSchemasOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/schemas`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetSchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSchemasOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSchemasOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*getSchemasOptions.EngineID))
	builder.AddQuery("catalog_name", fmt.Sprint(*getSchemasOptions.CatalogName))

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetSchemasOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostQuery : Run SQL statement
// Using this API to run a SQL statement.
func (watsonxData *WatsonxDataV1) PostQuery(postQueryOptions *PostQueryOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.PostQueryWithContext(context.Background(), postQueryOptions)
}

// PostQueryWithContext is an alternate form of the PostQuery method which supports a Context parameter
func (watsonxData *WatsonxDataV1) PostQueryWithContext(ctx context.Context, postQueryOptions *PostQueryOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postQueryOptions, "postQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postQueryOptions, "postQueryOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/v1/statement`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range postQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "PostQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if postQueryOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*postQueryOptions.Accept))
	}
	if postQueryOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*postQueryOptions.AuthInstanceID))
	}

	builder.AddQuery("engine", fmt.Sprint(*postQueryOptions.Engine))

	builder.AddFormData("catalog", "", "", fmt.Sprint(*postQueryOptions.Catalog))
	builder.AddFormData("schema", "", "", fmt.Sprint(*postQueryOptions.Schema))
	builder.AddFormData("sqlQuery", "", "", fmt.Sprint(*postQueryOptions.SqlQuery))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// DeleteTable : Delete table
// Delete one or multiple tables for a given schema and catalog.
func (watsonxData *WatsonxDataV1) DeleteTable(deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	return watsonxData.DeleteTableWithContext(context.Background(), deleteTableOptions)
}

// DeleteTableWithContext is an alternate form of the DeleteTable method which supports a Context parameter
func (watsonxData *WatsonxDataV1) DeleteTableWithContext(ctx context.Context, deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTableOptions, "deleteTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTableOptions, "deleteTableOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/tables/table`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "DeleteTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if deleteTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteTableOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if deleteTableOptions.DeleteTables != nil {
		body["delete_tables"] = deleteTableOptions.DeleteTables
	}
	if deleteTableOptions.EngineID != nil {
		body["engine_id"] = deleteTableOptions.EngineID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, nil)

	return
}

// UpdateTable : Update table
// Update the given table - rename table, add/drop/rename columns.
func (watsonxData *WatsonxDataV1) UpdateTable(updateTableOptions *UpdateTableOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.UpdateTableWithContext(context.Background(), updateTableOptions)
}

// UpdateTableWithContext is an alternate form of the UpdateTable method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UpdateTableWithContext(ctx context.Context, updateTableOptions *UpdateTableOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTableOptions, "updateTableOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTableOptions, "updateTableOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/tables/table`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UpdateTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if updateTableOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*updateTableOptions.Accept))
	}
	if updateTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*updateTableOptions.EngineID))
	builder.AddQuery("catalog_name", fmt.Sprint(*updateTableOptions.CatalogName))
	builder.AddQuery("schema_name", fmt.Sprint(*updateTableOptions.SchemaName))
	builder.AddQuery("table_name", fmt.Sprint(*updateTableOptions.TableName))

	body := make(map[string]interface{})
	if updateTableOptions.AddColumns != nil {
		body["add_columns"] = updateTableOptions.AddColumns
	}
	if updateTableOptions.DropColumns != nil {
		body["drop_columns"] = updateTableOptions.DropColumns
	}
	if updateTableOptions.NewTableName != nil {
		body["new_table_name"] = updateTableOptions.NewTableName
	}
	if updateTableOptions.RenameColumns != nil {
		body["rename_columns"] = updateTableOptions.RenameColumns
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// GetTableSnapshots : Get table snapshots
// List all table snapshots.
func (watsonxData *WatsonxDataV1) GetTableSnapshots(getTableSnapshotsOptions *GetTableSnapshotsOptions) (result *GetTableSnapshotsOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetTableSnapshotsWithContext(context.Background(), getTableSnapshotsOptions)
}

// GetTableSnapshotsWithContext is an alternate form of the GetTableSnapshots method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetTableSnapshotsWithContext(ctx context.Context, getTableSnapshotsOptions *GetTableSnapshotsOptions) (result *GetTableSnapshotsOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTableSnapshotsOptions, "getTableSnapshotsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTableSnapshotsOptions, "getTableSnapshotsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/tables/table/snapshots`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTableSnapshotsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetTableSnapshots")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTableSnapshotsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getTableSnapshotsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*getTableSnapshotsOptions.EngineID))
	builder.AddQuery("catalog_name", fmt.Sprint(*getTableSnapshotsOptions.CatalogName))
	builder.AddQuery("schema_name", fmt.Sprint(*getTableSnapshotsOptions.SchemaName))
	builder.AddQuery("table_name", fmt.Sprint(*getTableSnapshotsOptions.TableName))

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTableSnapshotsOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RollbackSnapshot : Rollback snapshot
// Rollback to a table snapshot.
func (watsonxData *WatsonxDataV1) RollbackSnapshot(rollbackSnapshotOptions *RollbackSnapshotOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	return watsonxData.RollbackSnapshotWithContext(context.Background(), rollbackSnapshotOptions)
}

// RollbackSnapshotWithContext is an alternate form of the RollbackSnapshot method which supports a Context parameter
func (watsonxData *WatsonxDataV1) RollbackSnapshotWithContext(ctx context.Context, rollbackSnapshotOptions *RollbackSnapshotOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(rollbackSnapshotOptions, "rollbackSnapshotOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(rollbackSnapshotOptions, "rollbackSnapshotOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/tables/table/rollback`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range rollbackSnapshotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "RollbackSnapshot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if rollbackSnapshotOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*rollbackSnapshotOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*rollbackSnapshotOptions.EngineID))
	builder.AddQuery("catalog_name", fmt.Sprint(*rollbackSnapshotOptions.CatalogName))
	builder.AddQuery("schema_name", fmt.Sprint(*rollbackSnapshotOptions.SchemaName))

	body := make(map[string]interface{})
	if rollbackSnapshotOptions.SnapshotID != nil {
		body["snapshot_id"] = rollbackSnapshotOptions.SnapshotID
	}
	if rollbackSnapshotOptions.TableName != nil {
		body["table_name"] = rollbackSnapshotOptions.TableName
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

// GetTables : Get tables
// List all tables in a schema in a catalog for a given engine.
func (watsonxData *WatsonxDataV1) GetTables(getTablesOptions *GetTablesOptions) (result *GetTablesOKBody, response *core.DetailedResponse, err error) {
	return watsonxData.GetTablesWithContext(context.Background(), getTablesOptions)
}

// GetTablesWithContext is an alternate form of the GetTables method which supports a Context parameter
func (watsonxData *WatsonxDataV1) GetTablesWithContext(ctx context.Context, getTablesOptions *GetTablesOptions) (result *GetTablesOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTablesOptions, "getTablesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTablesOptions, "getTablesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/tables`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "GetTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTablesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getTablesOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*getTablesOptions.EngineID))
	builder.AddQuery("catalog_name", fmt.Sprint(*getTablesOptions.CatalogName))
	builder.AddQuery("schema_name", fmt.Sprint(*getTablesOptions.SchemaName))

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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTablesOKBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ParseCsv : Parse CSV for table creation
// When creating a tabble, parse the CSV file.
func (watsonxData *WatsonxDataV1) ParseCsv(parseCsvOptions *ParseCsvOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.ParseCsvWithContext(context.Background(), parseCsvOptions)
}

// ParseCsvWithContext is an alternate form of the ParseCsv method which supports a Context parameter
func (watsonxData *WatsonxDataV1) ParseCsvWithContext(ctx context.Context, parseCsvOptions *ParseCsvOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(parseCsvOptions, "parseCsvOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(parseCsvOptions, "parseCsvOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/parse/csv`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range parseCsvOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "ParseCsv")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if parseCsvOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*parseCsvOptions.Accept))
	}
	if parseCsvOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*parseCsvOptions.AuthInstanceID))
	}

	builder.AddQuery("engine", fmt.Sprint(*parseCsvOptions.Engine))

	builder.AddFormData("parse_file", "", "", fmt.Sprint(*parseCsvOptions.ParseFile))
	builder.AddFormData("file_type", "", "", fmt.Sprint(*parseCsvOptions.FileType))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// UplaodCsv : Upload CSV for table creation
// When creating a table, upload a CSV file.
func (watsonxData *WatsonxDataV1) UplaodCsv(uplaodCsvOptions *UplaodCsvOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonxData.UplaodCsvWithContext(context.Background(), uplaodCsvOptions)
}

// UplaodCsvWithContext is an alternate form of the UplaodCsv method which supports a Context parameter
func (watsonxData *WatsonxDataV1) UplaodCsvWithContext(ctx context.Context, uplaodCsvOptions *UplaodCsvOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uplaodCsvOptions, "uplaodCsvOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uplaodCsvOptions, "uplaodCsvOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/v2/upload/csv`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range uplaodCsvOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V1", "UplaodCsv")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if uplaodCsvOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*uplaodCsvOptions.Accept))
	}
	if uplaodCsvOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*uplaodCsvOptions.AuthInstanceID))
	}

	builder.AddQuery("engine", fmt.Sprint(*uplaodCsvOptions.Engine))

	builder.AddFormData("catalog", "", "", fmt.Sprint(*uplaodCsvOptions.Catalog))
	builder.AddFormData("schema", "", "", fmt.Sprint(*uplaodCsvOptions.Schema))
	builder.AddFormData("tableName", "", "", fmt.Sprint(*uplaodCsvOptions.TableName))
	builder.AddFormData("ingestionJobName", "", "", fmt.Sprint(*uplaodCsvOptions.IngestionJobName))
	builder.AddFormData("scheduled", "", "", fmt.Sprint(*uplaodCsvOptions.Scheduled))
	builder.AddFormData("created_by", "", "", fmt.Sprint(*uplaodCsvOptions.CreatedBy))
	builder.AddFormData("targetTable", "", "", fmt.Sprint(*uplaodCsvOptions.TargetTable))
	builder.AddFormData("headers", "", "", fmt.Sprint(*uplaodCsvOptions.HeadersVar))
	builder.AddFormData("csv", "", "", fmt.Sprint(*uplaodCsvOptions.Csv))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonxData.Service.Request(request, &result)

	return
}

// ActivateBucketOptions : The ActivateBucket options.
type ActivateBucketOptions struct {
	// Bucket name.
	BucketID *string `json:"bucket_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewActivateBucketOptions : Instantiate ActivateBucketOptions
func (*WatsonxDataV1) NewActivateBucketOptions(bucketID string) *ActivateBucketOptions {
	return &ActivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *ActivateBucketOptions) SetBucketID(bucketID string) *ActivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *ActivateBucketOptions) SetAccept(accept string) *ActivateBucketOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ActivateBucketOptions) SetAuthInstanceID(authInstanceID string) *ActivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ActivateBucketOptions) SetHeaders(param map[string]string) *ActivateBucketOptions {
	options.Headers = param
	return options
}

// AddMetastoreToEngineOptions : The AddMetastoreToEngine options.
type AddMetastoreToEngineOptions struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddMetastoreToEngineOptions : Instantiate AddMetastoreToEngineOptions
func (*WatsonxDataV1) NewAddMetastoreToEngineOptions(catalogName string, engineID string) *AddMetastoreToEngineOptions {
	return &AddMetastoreToEngineOptions{
		CatalogName: core.StringPtr(catalogName),
		EngineID:    core.StringPtr(engineID),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *AddMetastoreToEngineOptions) SetCatalogName(catalogName string) *AddMetastoreToEngineOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *AddMetastoreToEngineOptions) SetEngineID(engineID string) *AddMetastoreToEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *AddMetastoreToEngineOptions) SetAccept(accept string) *AddMetastoreToEngineOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *AddMetastoreToEngineOptions) SetCreatedBy(createdBy string) *AddMetastoreToEngineOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *AddMetastoreToEngineOptions) SetAuthInstanceID(authInstanceID string) *AddMetastoreToEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddMetastoreToEngineOptions) SetHeaders(param map[string]string) *AddMetastoreToEngineOptions {
	options.Headers = param
	return options
}

// CreateBucketUsersOptions : The CreateBucketUsers options.
type CreateBucketUsersOptions struct {
	// The bucket id.
	BucketID *string `json:"bucket_id" validate:"required"`

	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateBucketUsersOptions : Instantiate CreateBucketUsersOptions
func (*WatsonxDataV1) NewCreateBucketUsersOptions(bucketID string) *CreateBucketUsersOptions {
	return &CreateBucketUsersOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *CreateBucketUsersOptions) SetBucketID(bucketID string) *CreateBucketUsersOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *CreateBucketUsersOptions) SetGroups(groups []BucketDbConnGroupsMetadata) *CreateBucketUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *CreateBucketUsersOptions) SetUsers(users []BucketDbConnUsersMetadata) *CreateBucketUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateBucketUsersOptions) SetLhInstanceID(lhInstanceID string) *CreateBucketUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateBucketUsersOptions) SetAuthInstanceID(authInstanceID string) *CreateBucketUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBucketUsersOptions) SetHeaders(param map[string]string) *CreateBucketUsersOptions {
	options.Headers = param
	return options
}

// CreateCatalogUsersOptions : The CreateCatalogUsers options.
type CreateCatalogUsersOptions struct {
	// The catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// The group list.
	Groups []CatalogGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []CatalogUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCatalogUsersOptions : Instantiate CreateCatalogUsersOptions
func (*WatsonxDataV1) NewCreateCatalogUsersOptions(catalogName string) *CreateCatalogUsersOptions {
	return &CreateCatalogUsersOptions{
		CatalogName: core.StringPtr(catalogName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *CreateCatalogUsersOptions) SetCatalogName(catalogName string) *CreateCatalogUsersOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *CreateCatalogUsersOptions) SetGroups(groups []CatalogGroupsMetadata) *CreateCatalogUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *CreateCatalogUsersOptions) SetUsers(users []CatalogUsersMetadata) *CreateCatalogUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateCatalogUsersOptions) SetLhInstanceID(lhInstanceID string) *CreateCatalogUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateCatalogUsersOptions) SetAuthInstanceID(authInstanceID string) *CreateCatalogUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCatalogUsersOptions) SetHeaders(param map[string]string) *CreateCatalogUsersOptions {
	options.Headers = param
	return options
}

// CreateDataPolicyOptions : The CreateDataPolicy options.
type CreateDataPolicyOptions struct {
	// catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// data artifact.
	DataArtifact *string `json:"data_artifact" validate:"required"`

	// the displayed name for data policy.
	PolicyName *string `json:"policy_name" validate:"required"`

	// rules.
	Rules []Rule `json:"rules" validate:"required"`

	// a more detailed description of the policy.
	Description *string `json:"description,omitempty"`

	// data policy status.
	Status *string `json:"status,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDataPolicyOptions.Status property.
// data policy status.
const (
	CreateDataPolicyOptions_Status_Active   = "active"
	CreateDataPolicyOptions_Status_Inactive = "inactive"
)

// NewCreateDataPolicyOptions : Instantiate CreateDataPolicyOptions
func (*WatsonxDataV1) NewCreateDataPolicyOptions(catalogName string, dataArtifact string, policyName string, rules []Rule) *CreateDataPolicyOptions {
	return &CreateDataPolicyOptions{
		CatalogName:  core.StringPtr(catalogName),
		DataArtifact: core.StringPtr(dataArtifact),
		PolicyName:   core.StringPtr(policyName),
		Rules:        rules,
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *CreateDataPolicyOptions) SetCatalogName(catalogName string) *CreateDataPolicyOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetDataArtifact : Allow user to set DataArtifact
func (_options *CreateDataPolicyOptions) SetDataArtifact(dataArtifact string) *CreateDataPolicyOptions {
	_options.DataArtifact = core.StringPtr(dataArtifact)
	return _options
}

// SetPolicyName : Allow user to set PolicyName
func (_options *CreateDataPolicyOptions) SetPolicyName(policyName string) *CreateDataPolicyOptions {
	_options.PolicyName = core.StringPtr(policyName)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *CreateDataPolicyOptions) SetRules(rules []Rule) *CreateDataPolicyOptions {
	_options.Rules = rules
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDataPolicyOptions) SetDescription(description string) *CreateDataPolicyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *CreateDataPolicyOptions) SetStatus(status string) *CreateDataPolicyOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateDataPolicyOptions) SetLhInstanceID(lhInstanceID string) *CreateDataPolicyOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDataPolicyOptions) SetAuthInstanceID(authInstanceID string) *CreateDataPolicyOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataPolicyOptions) SetHeaders(param map[string]string) *CreateDataPolicyOptions {
	options.Headers = param
	return options
}

// CreateDatabaseCatalogOptions : The CreateDatabaseCatalog options.
type CreateDatabaseCatalogOptions struct {
	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// Catalog name of the new catalog to be created with database.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// database details.
	DatabaseDetails *RegisterDatabaseCatalogBodyDatabaseDetails `json:"database_details,omitempty"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDatabaseCatalogOptions.DatabaseType property.
// Connector type.
const (
	CreateDatabaseCatalogOptions_DatabaseType_Db2        = "db2"
	CreateDatabaseCatalogOptions_DatabaseType_Jmx        = "jmx"
	CreateDatabaseCatalogOptions_DatabaseType_Kafka      = "kafka"
	CreateDatabaseCatalogOptions_DatabaseType_Memory     = "memory"
	CreateDatabaseCatalogOptions_DatabaseType_Mongodb    = "mongodb"
	CreateDatabaseCatalogOptions_DatabaseType_Mysql      = "mysql"
	CreateDatabaseCatalogOptions_DatabaseType_Netezza    = "netezza"
	CreateDatabaseCatalogOptions_DatabaseType_Postgresql = "postgresql"
	CreateDatabaseCatalogOptions_DatabaseType_System     = "system"
	CreateDatabaseCatalogOptions_DatabaseType_Tpcds      = "tpcds"
	CreateDatabaseCatalogOptions_DatabaseType_Tpch       = "tpch"
)

// NewCreateDatabaseCatalogOptions : Instantiate CreateDatabaseCatalogOptions
func (*WatsonxDataV1) NewCreateDatabaseCatalogOptions(databaseDisplayName string, databaseType string, catalogName string) *CreateDatabaseCatalogOptions {
	return &CreateDatabaseCatalogOptions{
		DatabaseDisplayName: core.StringPtr(databaseDisplayName),
		DatabaseType:        core.StringPtr(databaseType),
		CatalogName:         core.StringPtr(catalogName),
	}
}

// SetDatabaseDisplayName : Allow user to set DatabaseDisplayName
func (_options *CreateDatabaseCatalogOptions) SetDatabaseDisplayName(databaseDisplayName string) *CreateDatabaseCatalogOptions {
	_options.DatabaseDisplayName = core.StringPtr(databaseDisplayName)
	return _options
}

// SetDatabaseType : Allow user to set DatabaseType
func (_options *CreateDatabaseCatalogOptions) SetDatabaseType(databaseType string) *CreateDatabaseCatalogOptions {
	_options.DatabaseType = core.StringPtr(databaseType)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *CreateDatabaseCatalogOptions) SetCatalogName(catalogName string) *CreateDatabaseCatalogOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *CreateDatabaseCatalogOptions) SetAccept(accept string) *CreateDatabaseCatalogOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetDatabaseDetails : Allow user to set DatabaseDetails
func (_options *CreateDatabaseCatalogOptions) SetDatabaseDetails(databaseDetails *RegisterDatabaseCatalogBodyDatabaseDetails) *CreateDatabaseCatalogOptions {
	_options.DatabaseDetails = databaseDetails
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDatabaseCatalogOptions) SetDescription(description string) *CreateDatabaseCatalogOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDatabaseCatalogOptions) SetTags(tags []string) *CreateDatabaseCatalogOptions {
	_options.Tags = tags
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *CreateDatabaseCatalogOptions) SetCreatedBy(createdBy string) *CreateDatabaseCatalogOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *CreateDatabaseCatalogOptions) SetCreatedOn(createdOn int64) *CreateDatabaseCatalogOptions {
	_options.CreatedOn = core.Int64Ptr(createdOn)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDatabaseCatalogOptions) SetAuthInstanceID(authInstanceID string) *CreateDatabaseCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDatabaseCatalogOptions) SetHeaders(param map[string]string) *CreateDatabaseCatalogOptions {
	options.Headers = param
	return options
}

// CreateDbConnUsersOptions : The CreateDbConnUsers options.
type CreateDbConnUsersOptions struct {
	// The db connection id.
	DatabaseID *string `json:"database_id" validate:"required"`

	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDbConnUsersOptions : Instantiate CreateDbConnUsersOptions
func (*WatsonxDataV1) NewCreateDbConnUsersOptions(databaseID string) *CreateDbConnUsersOptions {
	return &CreateDbConnUsersOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *CreateDbConnUsersOptions) SetDatabaseID(databaseID string) *CreateDbConnUsersOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *CreateDbConnUsersOptions) SetGroups(groups []BucketDbConnGroupsMetadata) *CreateDbConnUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *CreateDbConnUsersOptions) SetUsers(users []BucketDbConnUsersMetadata) *CreateDbConnUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateDbConnUsersOptions) SetLhInstanceID(lhInstanceID string) *CreateDbConnUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDbConnUsersOptions) SetAuthInstanceID(authInstanceID string) *CreateDbConnUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDbConnUsersOptions) SetHeaders(param map[string]string) *CreateDbConnUsersOptions {
	options.Headers = param
	return options
}

// CreateEngineOptions : The CreateEngine options.
type CreateEngineOptions struct {
	// Version like 0.278 for presto or else.
	Version *string `json:"version" validate:"required"`

	// Node details.
	EngineDetails *EngineDetailsBody `json:"engine_details" validate:"required"`

	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine type presto, others like netezza.
	Type *string `json:"type" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Optional parameter for UI - set as true when first time use.
	FirstTimeUse *bool `json:"first_time_use,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateEngineOptions_Origin_Customer = "customer"
	CreateEngineOptions_Origin_Ibm      = "ibm"
)

// NewCreateEngineOptions : Instantiate CreateEngineOptions
func (*WatsonxDataV1) NewCreateEngineOptions(version string, engineDetails *EngineDetailsBody, origin string, typeVar string) *CreateEngineOptions {
	return &CreateEngineOptions{
		Version:       core.StringPtr(version),
		EngineDetails: engineDetails,
		Origin:        core.StringPtr(origin),
		Type:          core.StringPtr(typeVar),
	}
}

// SetVersion : Allow user to set Version
func (_options *CreateEngineOptions) SetVersion(version string) *CreateEngineOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateEngineOptions) SetEngineDetails(engineDetails *EngineDetailsBody) *CreateEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetOrigin : Allow user to set Origin
func (_options *CreateEngineOptions) SetOrigin(origin string) *CreateEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateEngineOptions) SetType(typeVar string) *CreateEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *CreateEngineOptions) SetAccept(accept string) *CreateEngineOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateEngineOptions) SetDescription(description string) *CreateEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetFirstTimeUse : Allow user to set FirstTimeUse
func (_options *CreateEngineOptions) SetFirstTimeUse(firstTimeUse bool) *CreateEngineOptions {
	_options.FirstTimeUse = core.BoolPtr(firstTimeUse)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateEngineOptions) SetRegion(region string) *CreateEngineOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreateEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreateEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEngineOptions) SetHeaders(param map[string]string) *CreateEngineOptions {
	options.Headers = param
	return options
}

// CreateEngineUsersOptions : The CreateEngineUsers options.
type CreateEngineUsersOptions struct {
	// The engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// The group list.
	Groups []EngineGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []EngineUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEngineUsersOptions : Instantiate CreateEngineUsersOptions
func (*WatsonxDataV1) NewCreateEngineUsersOptions(engineID string) *CreateEngineUsersOptions {
	return &CreateEngineUsersOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateEngineUsersOptions) SetEngineID(engineID string) *CreateEngineUsersOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *CreateEngineUsersOptions) SetGroups(groups []EngineGroupsMetadata) *CreateEngineUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *CreateEngineUsersOptions) SetUsers(users []EngineUsersMetadata) *CreateEngineUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateEngineUsersOptions) SetLhInstanceID(lhInstanceID string) *CreateEngineUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateEngineUsersOptions) SetAuthInstanceID(authInstanceID string) *CreateEngineUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEngineUsersOptions) SetHeaders(param map[string]string) *CreateEngineUsersOptions {
	options.Headers = param
	return options
}

// CreateMetastoreUsersOptions : The CreateMetastoreUsers options.
type CreateMetastoreUsersOptions struct {
	// The metastore name.
	MetastoreName *string `json:"metastore_name" validate:"required"`

	// The group list.
	Groups []GroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []UsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateMetastoreUsersOptions : Instantiate CreateMetastoreUsersOptions
func (*WatsonxDataV1) NewCreateMetastoreUsersOptions(metastoreName string) *CreateMetastoreUsersOptions {
	return &CreateMetastoreUsersOptions{
		MetastoreName: core.StringPtr(metastoreName),
	}
}

// SetMetastoreName : Allow user to set MetastoreName
func (_options *CreateMetastoreUsersOptions) SetMetastoreName(metastoreName string) *CreateMetastoreUsersOptions {
	_options.MetastoreName = core.StringPtr(metastoreName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *CreateMetastoreUsersOptions) SetGroups(groups []GroupsMetadata) *CreateMetastoreUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *CreateMetastoreUsersOptions) SetUsers(users []UsersMetadata) *CreateMetastoreUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *CreateMetastoreUsersOptions) SetLhInstanceID(lhInstanceID string) *CreateMetastoreUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateMetastoreUsersOptions) SetAuthInstanceID(authInstanceID string) *CreateMetastoreUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMetastoreUsersOptions) SetHeaders(param map[string]string) *CreateMetastoreUsersOptions {
	options.Headers = param
	return options
}

// CreateSchemaOptions : The CreateSchema options.
type CreateSchemaOptions struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

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
func (*WatsonxDataV1) NewCreateSchemaOptions(catalogName string, engineID string, schemaName string) *CreateSchemaOptions {
	return &CreateSchemaOptions{
		CatalogName: core.StringPtr(catalogName),
		EngineID:    core.StringPtr(engineID),
		SchemaName:  core.StringPtr(schemaName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *CreateSchemaOptions) SetCatalogName(catalogName string) *CreateSchemaOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSchemaOptions) SetEngineID(engineID string) *CreateSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
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

// DeactivateBucketOptions : The DeactivateBucket options.
type DeactivateBucketOptions struct {
	// Bucket name.
	BucketID *string `json:"bucket_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeactivateBucketOptions : Instantiate DeactivateBucketOptions
func (*WatsonxDataV1) NewDeactivateBucketOptions(bucketID string) *DeactivateBucketOptions {
	return &DeactivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeactivateBucketOptions) SetBucketID(bucketID string) *DeactivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *DeactivateBucketOptions) SetAccept(accept string) *DeactivateBucketOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeactivateBucketOptions) SetAuthInstanceID(authInstanceID string) *DeactivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeactivateBucketOptions) SetHeaders(param map[string]string) *DeactivateBucketOptions {
	options.Headers = param
	return options
}

// DeleteBucketUsersOptions : The DeleteBucketUsers options.
type DeleteBucketUsersOptions struct {
	// Bucket ID for DELETE.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// The group ids array to be deleted.
	Groups []string `json:"groups,omitempty"`

	// The user names array to be deleted.
	Users []string `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketUsersOptions : Instantiate DeleteBucketUsersOptions
func (*WatsonxDataV1) NewDeleteBucketUsersOptions(bucketID string) *DeleteBucketUsersOptions {
	return &DeleteBucketUsersOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeleteBucketUsersOptions) SetBucketID(bucketID string) *DeleteBucketUsersOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *DeleteBucketUsersOptions) SetGroups(groups []string) *DeleteBucketUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *DeleteBucketUsersOptions) SetUsers(users []string) *DeleteBucketUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteBucketUsersOptions) SetLhInstanceID(lhInstanceID string) *DeleteBucketUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteBucketUsersOptions) SetAuthInstanceID(authInstanceID string) *DeleteBucketUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketUsersOptions) SetHeaders(param map[string]string) *DeleteBucketUsersOptions {
	options.Headers = param
	return options
}

// DeleteCatalogUsersOptions : The DeleteCatalogUsers options.
type DeleteCatalogUsersOptions struct {
	// Catalog name for DELETE.
	CatalogName *string `json:"catalog_name" validate:"required,ne="`

	// The group ids array to be deleted.
	Groups []string `json:"groups,omitempty"`

	// The user names array to be deleted.
	Users []string `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCatalogUsersOptions : Instantiate DeleteCatalogUsersOptions
func (*WatsonxDataV1) NewDeleteCatalogUsersOptions(catalogName string) *DeleteCatalogUsersOptions {
	return &DeleteCatalogUsersOptions{
		CatalogName: core.StringPtr(catalogName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *DeleteCatalogUsersOptions) SetCatalogName(catalogName string) *DeleteCatalogUsersOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *DeleteCatalogUsersOptions) SetGroups(groups []string) *DeleteCatalogUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *DeleteCatalogUsersOptions) SetUsers(users []string) *DeleteCatalogUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteCatalogUsersOptions) SetLhInstanceID(lhInstanceID string) *DeleteCatalogUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteCatalogUsersOptions) SetAuthInstanceID(authInstanceID string) *DeleteCatalogUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCatalogUsersOptions) SetHeaders(param map[string]string) *DeleteCatalogUsersOptions {
	options.Headers = param
	return options
}

// DeleteDataPoliciesOptions : The DeleteDataPolicies options.
type DeleteDataPoliciesOptions struct {
	// data policy names array to be deleted.
	DataPolicies []string `json:"data_policies,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDataPoliciesOptions : Instantiate DeleteDataPoliciesOptions
func (*WatsonxDataV1) NewDeleteDataPoliciesOptions() *DeleteDataPoliciesOptions {
	return &DeleteDataPoliciesOptions{}
}

// SetDataPolicies : Allow user to set DataPolicies
func (_options *DeleteDataPoliciesOptions) SetDataPolicies(dataPolicies []string) *DeleteDataPoliciesOptions {
	_options.DataPolicies = dataPolicies
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteDataPoliciesOptions) SetLhInstanceID(lhInstanceID string) *DeleteDataPoliciesOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDataPoliciesOptions) SetAuthInstanceID(authInstanceID string) *DeleteDataPoliciesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDataPoliciesOptions) SetHeaders(param map[string]string) *DeleteDataPoliciesOptions {
	options.Headers = param
	return options
}

// DeleteDataPolicyOptions : The DeleteDataPolicy options.
type DeleteDataPolicyOptions struct {
	// Policy name for DELETE.
	PolicyName *string `json:"policy_name" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDataPolicyOptions : Instantiate DeleteDataPolicyOptions
func (*WatsonxDataV1) NewDeleteDataPolicyOptions(policyName string) *DeleteDataPolicyOptions {
	return &DeleteDataPolicyOptions{
		PolicyName: core.StringPtr(policyName),
	}
}

// SetPolicyName : Allow user to set PolicyName
func (_options *DeleteDataPolicyOptions) SetPolicyName(policyName string) *DeleteDataPolicyOptions {
	_options.PolicyName = core.StringPtr(policyName)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteDataPolicyOptions) SetLhInstanceID(lhInstanceID string) *DeleteDataPolicyOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDataPolicyOptions) SetAuthInstanceID(authInstanceID string) *DeleteDataPolicyOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDataPolicyOptions) SetHeaders(param map[string]string) *DeleteDataPolicyOptions {
	options.Headers = param
	return options
}

// DeleteDatabaseCatalogOptions : The DeleteDatabaseCatalog options.
type DeleteDatabaseCatalogOptions struct {
	// Database ID.
	DatabaseID *string `json:"database_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDatabaseCatalogOptions : Instantiate DeleteDatabaseCatalogOptions
func (*WatsonxDataV1) NewDeleteDatabaseCatalogOptions(databaseID string) *DeleteDatabaseCatalogOptions {
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

// DeleteDbConnUsersOptions : The DeleteDbConnUsers options.
type DeleteDbConnUsersOptions struct {
	// Db connection id for DELETE.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// The group ids array to be deleted.
	Groups []string `json:"groups,omitempty"`

	// The user names array to be deleted.
	Users []string `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDbConnUsersOptions : Instantiate DeleteDbConnUsersOptions
func (*WatsonxDataV1) NewDeleteDbConnUsersOptions(databaseID string) *DeleteDbConnUsersOptions {
	return &DeleteDbConnUsersOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *DeleteDbConnUsersOptions) SetDatabaseID(databaseID string) *DeleteDbConnUsersOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *DeleteDbConnUsersOptions) SetGroups(groups []string) *DeleteDbConnUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *DeleteDbConnUsersOptions) SetUsers(users []string) *DeleteDbConnUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteDbConnUsersOptions) SetLhInstanceID(lhInstanceID string) *DeleteDbConnUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDbConnUsersOptions) SetAuthInstanceID(authInstanceID string) *DeleteDbConnUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDbConnUsersOptions) SetHeaders(param map[string]string) *DeleteDbConnUsersOptions {
	options.Headers = param
	return options
}

// DeleteEngineOptions : The DeleteEngine options.
type DeleteEngineOptions struct {
	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEngineOptions : Instantiate DeleteEngineOptions
func (*WatsonxDataV1) NewDeleteEngineOptions(engineID string) *DeleteEngineOptions {
	return &DeleteEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteEngineOptions) SetEngineID(engineID string) *DeleteEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *DeleteEngineOptions) SetCreatedBy(createdBy string) *DeleteEngineOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
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

// DeleteEngineUsersOptions : The DeleteEngineUsers options.
type DeleteEngineUsersOptions struct {
	// Engine ID for DELETE.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// The group ids array to be deleted.
	Groups []string `json:"groups,omitempty"`

	// The user names array to be deleted.
	Users []string `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEngineUsersOptions : Instantiate DeleteEngineUsersOptions
func (*WatsonxDataV1) NewDeleteEngineUsersOptions(engineID string) *DeleteEngineUsersOptions {
	return &DeleteEngineUsersOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteEngineUsersOptions) SetEngineID(engineID string) *DeleteEngineUsersOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *DeleteEngineUsersOptions) SetGroups(groups []string) *DeleteEngineUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *DeleteEngineUsersOptions) SetUsers(users []string) *DeleteEngineUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteEngineUsersOptions) SetLhInstanceID(lhInstanceID string) *DeleteEngineUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteEngineUsersOptions) SetAuthInstanceID(authInstanceID string) *DeleteEngineUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEngineUsersOptions) SetHeaders(param map[string]string) *DeleteEngineUsersOptions {
	options.Headers = param
	return options
}

// DeleteMetastoreUsersOptions : The DeleteMetastoreUsers options.
type DeleteMetastoreUsersOptions struct {
	// Metastore name for DELETE.
	MetastoreName *string `json:"metastore_name" validate:"required,ne="`

	// The group ids array to be deleted.
	Groups []string `json:"groups,omitempty"`

	// The user names array to be deleted.
	Users []string `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteMetastoreUsersOptions : Instantiate DeleteMetastoreUsersOptions
func (*WatsonxDataV1) NewDeleteMetastoreUsersOptions(metastoreName string) *DeleteMetastoreUsersOptions {
	return &DeleteMetastoreUsersOptions{
		MetastoreName: core.StringPtr(metastoreName),
	}
}

// SetMetastoreName : Allow user to set MetastoreName
func (_options *DeleteMetastoreUsersOptions) SetMetastoreName(metastoreName string) *DeleteMetastoreUsersOptions {
	_options.MetastoreName = core.StringPtr(metastoreName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *DeleteMetastoreUsersOptions) SetGroups(groups []string) *DeleteMetastoreUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *DeleteMetastoreUsersOptions) SetUsers(users []string) *DeleteMetastoreUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *DeleteMetastoreUsersOptions) SetLhInstanceID(lhInstanceID string) *DeleteMetastoreUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteMetastoreUsersOptions) SetAuthInstanceID(authInstanceID string) *DeleteMetastoreUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMetastoreUsersOptions) SetHeaders(param map[string]string) *DeleteMetastoreUsersOptions {
	options.Headers = param
	return options
}

// DeleteQueryOptions : The DeleteQuery options.
type DeleteQueryOptions struct {
	// Query name.
	QueryName *string `json:"query_name" validate:"required,ne="`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteQueryOptions : Instantiate DeleteQueryOptions
func (*WatsonxDataV1) NewDeleteQueryOptions(queryName string) *DeleteQueryOptions {
	return &DeleteQueryOptions{
		QueryName: core.StringPtr(queryName),
	}
}

// SetQueryName : Allow user to set QueryName
func (_options *DeleteQueryOptions) SetQueryName(queryName string) *DeleteQueryOptions {
	_options.QueryName = core.StringPtr(queryName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteQueryOptions) SetAuthInstanceID(authInstanceID string) *DeleteQueryOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteQueryOptions) SetHeaders(param map[string]string) *DeleteQueryOptions {
	options.Headers = param
	return options
}

// DeleteSchemaOptions : The DeleteSchema options.
type DeleteSchemaOptions struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSchemaOptions : Instantiate DeleteSchemaOptions
func (*WatsonxDataV1) NewDeleteSchemaOptions(catalogName string, engineID string, schemaName string) *DeleteSchemaOptions {
	return &DeleteSchemaOptions{
		CatalogName: core.StringPtr(catalogName),
		EngineID:    core.StringPtr(engineID),
		SchemaName:  core.StringPtr(schemaName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *DeleteSchemaOptions) SetCatalogName(catalogName string) *DeleteSchemaOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSchemaOptions) SetEngineID(engineID string) *DeleteSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *DeleteSchemaOptions) SetSchemaName(schemaName string) *DeleteSchemaOptions {
	_options.SchemaName = core.StringPtr(schemaName)
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

// DeleteTableOptions : The DeleteTable options.
type DeleteTableOptions struct {
	// Delete table list.
	DeleteTables []DeleteTableBodyDeleteTablesItems `json:"delete_tables" validate:"required"`

	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTableOptions : Instantiate DeleteTableOptions
func (*WatsonxDataV1) NewDeleteTableOptions(deleteTables []DeleteTableBodyDeleteTablesItems, engineID string) *DeleteTableOptions {
	return &DeleteTableOptions{
		DeleteTables: deleteTables,
		EngineID:     core.StringPtr(engineID),
	}
}

// SetDeleteTables : Allow user to set DeleteTables
func (_options *DeleteTableOptions) SetDeleteTables(deleteTables []DeleteTableBodyDeleteTablesItems) *DeleteTableOptions {
	_options.DeleteTables = deleteTables
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

// EvaluateOptions : The Evaluate options.
type EvaluateOptions struct {
	// resource list.
	Resources []ResourcesMetadata `json:"resources,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEvaluateOptions : Instantiate EvaluateOptions
func (*WatsonxDataV1) NewEvaluateOptions() *EvaluateOptions {
	return &EvaluateOptions{}
}

// SetResources : Allow user to set Resources
func (_options *EvaluateOptions) SetResources(resources []ResourcesMetadata) *EvaluateOptions {
	_options.Resources = resources
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *EvaluateOptions) SetLhInstanceID(lhInstanceID string) *EvaluateOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *EvaluateOptions) SetAuthInstanceID(authInstanceID string) *EvaluateOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *EvaluateOptions) SetHeaders(param map[string]string) *EvaluateOptions {
	options.Headers = param
	return options
}

// ExplainAnalyzeStatementOptions : The ExplainAnalyzeStatement options.
type ExplainAnalyzeStatementOptions struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Statement.
	Statement *string `json:"statement" validate:"required"`

	// Verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExplainAnalyzeStatementOptions : Instantiate ExplainAnalyzeStatementOptions
func (*WatsonxDataV1) NewExplainAnalyzeStatementOptions(catalogName string, engineID string, schemaName string, statement string) *ExplainAnalyzeStatementOptions {
	return &ExplainAnalyzeStatementOptions{
		CatalogName: core.StringPtr(catalogName),
		EngineID:    core.StringPtr(engineID),
		SchemaName:  core.StringPtr(schemaName),
		Statement:   core.StringPtr(statement),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *ExplainAnalyzeStatementOptions) SetCatalogName(catalogName string) *ExplainAnalyzeStatementOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *ExplainAnalyzeStatementOptions) SetEngineID(engineID string) *ExplainAnalyzeStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *ExplainAnalyzeStatementOptions) SetSchemaName(schemaName string) *ExplainAnalyzeStatementOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *ExplainAnalyzeStatementOptions) SetStatement(statement string) *ExplainAnalyzeStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *ExplainAnalyzeStatementOptions) SetVerbose(verbose bool) *ExplainAnalyzeStatementOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ExplainAnalyzeStatementOptions) SetAuthInstanceID(authInstanceID string) *ExplainAnalyzeStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExplainAnalyzeStatementOptions) SetHeaders(param map[string]string) *ExplainAnalyzeStatementOptions {
	options.Headers = param
	return options
}

// ExplainStatementOptions : The ExplainStatement options.
type ExplainStatementOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Statement.
	Statement *string `json:"statement" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Format.
	Format *string `json:"format,omitempty"`

	// Schema name.
	SchemaName *string `json:"schema_name,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ExplainStatementOptions.Format property.
// Format.
const (
	ExplainStatementOptions_Format_Graphviz = "graphviz"
	ExplainStatementOptions_Format_JSON     = "json"
	ExplainStatementOptions_Format_Text     = "text"
)

// Constants associated with the ExplainStatementOptions.Type property.
// Type.
const (
	ExplainStatementOptions_Type_Distributed = "distributed"
	ExplainStatementOptions_Type_Io          = "io"
	ExplainStatementOptions_Type_Logical     = "logical"
	ExplainStatementOptions_Type_Validate    = "validate"
)

// NewExplainStatementOptions : Instantiate ExplainStatementOptions
func (*WatsonxDataV1) NewExplainStatementOptions(engineID string, statement string) *ExplainStatementOptions {
	return &ExplainStatementOptions{
		EngineID:  core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ExplainStatementOptions) SetEngineID(engineID string) *ExplainStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *ExplainStatementOptions) SetStatement(statement string) *ExplainStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *ExplainStatementOptions) SetCatalogName(catalogName string) *ExplainStatementOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *ExplainStatementOptions) SetFormat(format string) *ExplainStatementOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *ExplainStatementOptions) SetSchemaName(schemaName string) *ExplainStatementOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetType : Allow user to set Type
func (_options *ExplainStatementOptions) SetType(typeVar string) *ExplainStatementOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ExplainStatementOptions) SetAuthInstanceID(authInstanceID string) *ExplainStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExplainStatementOptions) SetHeaders(param map[string]string) *ExplainStatementOptions {
	options.Headers = param
	return options
}

// GetBucketObjectsOptions : The GetBucketObjects options.
type GetBucketObjectsOptions struct {
	// Bucket ID.
	BucketID *string `json:"bucket_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketObjectsOptions : Instantiate GetBucketObjectsOptions
func (*WatsonxDataV1) NewGetBucketObjectsOptions(bucketID string) *GetBucketObjectsOptions {
	return &GetBucketObjectsOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *GetBucketObjectsOptions) SetBucketID(bucketID string) *GetBucketObjectsOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetBucketObjectsOptions) SetAuthInstanceID(authInstanceID string) *GetBucketObjectsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketObjectsOptions) SetHeaders(param map[string]string) *GetBucketObjectsOptions {
	options.Headers = param
	return options
}

// GetBucketUsersOptions : The GetBucketUsers options.
type GetBucketUsersOptions struct {
	// Bucket name for GET.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketUsersOptions : Instantiate GetBucketUsersOptions
func (*WatsonxDataV1) NewGetBucketUsersOptions(bucketID string) *GetBucketUsersOptions {
	return &GetBucketUsersOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *GetBucketUsersOptions) SetBucketID(bucketID string) *GetBucketUsersOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetBucketUsersOptions) SetLhInstanceID(lhInstanceID string) *GetBucketUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetBucketUsersOptions) SetAuthInstanceID(authInstanceID string) *GetBucketUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketUsersOptions) SetHeaders(param map[string]string) *GetBucketUsersOptions {
	options.Headers = param
	return options
}

// GetBucketsOptions : The GetBuckets options.
type GetBucketsOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketsOptions : Instantiate GetBucketsOptions
func (*WatsonxDataV1) NewGetBucketsOptions() *GetBucketsOptions {
	return &GetBucketsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetBucketsOptions) SetAuthInstanceID(authInstanceID string) *GetBucketsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketsOptions) SetHeaders(param map[string]string) *GetBucketsOptions {
	options.Headers = param
	return options
}

// GetCatalogUsersOptions : The GetCatalogUsers options.
type GetCatalogUsersOptions struct {
	// catalog name for GET.
	CatalogName *string `json:"catalog_name" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogUsersOptions : Instantiate GetCatalogUsersOptions
func (*WatsonxDataV1) NewGetCatalogUsersOptions(catalogName string) *GetCatalogUsersOptions {
	return &GetCatalogUsersOptions{
		CatalogName: core.StringPtr(catalogName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *GetCatalogUsersOptions) SetCatalogName(catalogName string) *GetCatalogUsersOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetCatalogUsersOptions) SetLhInstanceID(lhInstanceID string) *GetCatalogUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetCatalogUsersOptions) SetAuthInstanceID(authInstanceID string) *GetCatalogUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogUsersOptions) SetHeaders(param map[string]string) *GetCatalogUsersOptions {
	options.Headers = param
	return options
}

// GetDataPolicyOptions : The GetDataPolicy options.
type GetDataPolicyOptions struct {
	// policy name to get.
	PolicyName *string `json:"policy_name" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataPolicyOptions : Instantiate GetDataPolicyOptions
func (*WatsonxDataV1) NewGetDataPolicyOptions(policyName string) *GetDataPolicyOptions {
	return &GetDataPolicyOptions{
		PolicyName: core.StringPtr(policyName),
	}
}

// SetPolicyName : Allow user to set PolicyName
func (_options *GetDataPolicyOptions) SetPolicyName(policyName string) *GetDataPolicyOptions {
	_options.PolicyName = core.StringPtr(policyName)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetDataPolicyOptions) SetLhInstanceID(lhInstanceID string) *GetDataPolicyOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDataPolicyOptions) SetAuthInstanceID(authInstanceID string) *GetDataPolicyOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataPolicyOptions) SetHeaders(param map[string]string) *GetDataPolicyOptions {
	options.Headers = param
	return options
}

// GetDatabasesOptions : The GetDatabases options.
type GetDatabasesOptions struct {
	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDatabasesOptions : Instantiate GetDatabasesOptions
func (*WatsonxDataV1) NewGetDatabasesOptions() *GetDatabasesOptions {
	return &GetDatabasesOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *GetDatabasesOptions) SetAccept(accept string) *GetDatabasesOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDatabasesOptions) SetAuthInstanceID(authInstanceID string) *GetDatabasesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDatabasesOptions) SetHeaders(param map[string]string) *GetDatabasesOptions {
	options.Headers = param
	return options
}

// GetDbConnUsersOptions : The GetDbConnUsers options.
type GetDbConnUsersOptions struct {
	// Db connection id for GET.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDbConnUsersOptions : Instantiate GetDbConnUsersOptions
func (*WatsonxDataV1) NewGetDbConnUsersOptions(databaseID string) *GetDbConnUsersOptions {
	return &GetDbConnUsersOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *GetDbConnUsersOptions) SetDatabaseID(databaseID string) *GetDbConnUsersOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetDbConnUsersOptions) SetLhInstanceID(lhInstanceID string) *GetDbConnUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDbConnUsersOptions) SetAuthInstanceID(authInstanceID string) *GetDbConnUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDbConnUsersOptions) SetHeaders(param map[string]string) *GetDbConnUsersOptions {
	options.Headers = param
	return options
}

// GetDefaultPoliciesOptions : The GetDefaultPolicies options.
type GetDefaultPoliciesOptions struct {
	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDefaultPoliciesOptions : Instantiate GetDefaultPoliciesOptions
func (*WatsonxDataV1) NewGetDefaultPoliciesOptions() *GetDefaultPoliciesOptions {
	return &GetDefaultPoliciesOptions{}
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetDefaultPoliciesOptions) SetLhInstanceID(lhInstanceID string) *GetDefaultPoliciesOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDefaultPoliciesOptions) SetAuthInstanceID(authInstanceID string) *GetDefaultPoliciesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDefaultPoliciesOptions) SetHeaders(param map[string]string) *GetDefaultPoliciesOptions {
	options.Headers = param
	return options
}

// GetDeploymentsOptions : The GetDeployments options.
type GetDeploymentsOptions struct {
	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentsOptions : Instantiate GetDeploymentsOptions
func (*WatsonxDataV1) NewGetDeploymentsOptions() *GetDeploymentsOptions {
	return &GetDeploymentsOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *GetDeploymentsOptions) SetAccept(accept string) *GetDeploymentsOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
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

// GetEngineUsersOptions : The GetEngineUsers options.
type GetEngineUsersOptions struct {
	// Engine ID for GET.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEngineUsersOptions : Instantiate GetEngineUsersOptions
func (*WatsonxDataV1) NewGetEngineUsersOptions(engineID string) *GetEngineUsersOptions {
	return &GetEngineUsersOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetEngineUsersOptions) SetEngineID(engineID string) *GetEngineUsersOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetEngineUsersOptions) SetLhInstanceID(lhInstanceID string) *GetEngineUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetEngineUsersOptions) SetAuthInstanceID(authInstanceID string) *GetEngineUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEngineUsersOptions) SetHeaders(param map[string]string) *GetEngineUsersOptions {
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
func (*WatsonxDataV1) NewGetEnginesOptions() *GetEnginesOptions {
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

// GetHMSOptions : The GetHMS options.
type GetHMSOptions struct {
	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetHMSOptions : Instantiate GetHMSOptions
func (*WatsonxDataV1) NewGetHMSOptions() *GetHMSOptions {
	return &GetHMSOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *GetHMSOptions) SetAccept(accept string) *GetHMSOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetHMSOptions) SetAuthInstanceID(authInstanceID string) *GetHMSOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetHMSOptions) SetHeaders(param map[string]string) *GetHMSOptions {
	options.Headers = param
	return options
}

// GetMetastoreUsersOptions : The GetMetastoreUsers options.
type GetMetastoreUsersOptions struct {
	// Metastore name for GET.
	MetastoreName *string `json:"metastore_name" validate:"required,ne="`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMetastoreUsersOptions : Instantiate GetMetastoreUsersOptions
func (*WatsonxDataV1) NewGetMetastoreUsersOptions(metastoreName string) *GetMetastoreUsersOptions {
	return &GetMetastoreUsersOptions{
		MetastoreName: core.StringPtr(metastoreName),
	}
}

// SetMetastoreName : Allow user to set MetastoreName
func (_options *GetMetastoreUsersOptions) SetMetastoreName(metastoreName string) *GetMetastoreUsersOptions {
	_options.MetastoreName = core.StringPtr(metastoreName)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetMetastoreUsersOptions) SetLhInstanceID(lhInstanceID string) *GetMetastoreUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetMetastoreUsersOptions) SetAuthInstanceID(authInstanceID string) *GetMetastoreUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetastoreUsersOptions) SetHeaders(param map[string]string) *GetMetastoreUsersOptions {
	options.Headers = param
	return options
}

// GetMetastoresOptions : The GetMetastores options.
type GetMetastoresOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMetastoresOptions : Instantiate GetMetastoresOptions
func (*WatsonxDataV1) NewGetMetastoresOptions() *GetMetastoresOptions {
	return &GetMetastoresOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetMetastoresOptions) SetAuthInstanceID(authInstanceID string) *GetMetastoresOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMetastoresOptions) SetHeaders(param map[string]string) *GetMetastoresOptions {
	options.Headers = param
	return options
}

// GetPoliciesListOptions : The GetPoliciesList options.
type GetPoliciesListOptions struct {
	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// policies for specific catalogs list.
	CatalogList []string `json:"catalog_list,omitempty"`

	// policies for specific engines list.
	EngineList []string `json:"engine_list,omitempty"`

	// policies for specific Data Polices list.
	DataPoliciesList []string `json:"data_policies_list,omitempty"`

	// include policies for specific catalogs or not.
	IncludeDataPolicies *bool `json:"include_data_policies,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPoliciesListOptions : Instantiate GetPoliciesListOptions
func (*WatsonxDataV1) NewGetPoliciesListOptions() *GetPoliciesListOptions {
	return &GetPoliciesListOptions{}
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetPoliciesListOptions) SetLhInstanceID(lhInstanceID string) *GetPoliciesListOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPoliciesListOptions) SetAuthInstanceID(authInstanceID string) *GetPoliciesListOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetCatalogList : Allow user to set CatalogList
func (_options *GetPoliciesListOptions) SetCatalogList(catalogList []string) *GetPoliciesListOptions {
	_options.CatalogList = catalogList
	return _options
}

// SetEngineList : Allow user to set EngineList
func (_options *GetPoliciesListOptions) SetEngineList(engineList []string) *GetPoliciesListOptions {
	_options.EngineList = engineList
	return _options
}

// SetDataPoliciesList : Allow user to set DataPoliciesList
func (_options *GetPoliciesListOptions) SetDataPoliciesList(dataPoliciesList []string) *GetPoliciesListOptions {
	_options.DataPoliciesList = dataPoliciesList
	return _options
}

// SetIncludeDataPolicies : Allow user to set IncludeDataPolicies
func (_options *GetPoliciesListOptions) SetIncludeDataPolicies(includeDataPolicies bool) *GetPoliciesListOptions {
	_options.IncludeDataPolicies = core.BoolPtr(includeDataPolicies)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPoliciesListOptions) SetHeaders(param map[string]string) *GetPoliciesListOptions {
	options.Headers = param
	return options
}

// GetPolicyVersionOptions : The GetPolicyVersion options.
type GetPolicyVersionOptions struct {
	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPolicyVersionOptions : Instantiate GetPolicyVersionOptions
func (*WatsonxDataV1) NewGetPolicyVersionOptions() *GetPolicyVersionOptions {
	return &GetPolicyVersionOptions{}
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *GetPolicyVersionOptions) SetLhInstanceID(lhInstanceID string) *GetPolicyVersionOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPolicyVersionOptions) SetAuthInstanceID(authInstanceID string) *GetPolicyVersionOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPolicyVersionOptions) SetHeaders(param map[string]string) *GetPolicyVersionOptions {
	options.Headers = param
	return options
}

// GetQueriesOptions : The GetQueries options.
type GetQueriesOptions struct {
	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetQueriesOptions : Instantiate GetQueriesOptions
func (*WatsonxDataV1) NewGetQueriesOptions() *GetQueriesOptions {
	return &GetQueriesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetQueriesOptions) SetAuthInstanceID(authInstanceID string) *GetQueriesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetQueriesOptions) SetHeaders(param map[string]string) *GetQueriesOptions {
	options.Headers = param
	return options
}

// GetSchemasOptions : The GetSchemas options.
type GetSchemasOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSchemasOptions : Instantiate GetSchemasOptions
func (*WatsonxDataV1) NewGetSchemasOptions(engineID string, catalogName string) *GetSchemasOptions {
	return &GetSchemasOptions{
		EngineID:    core.StringPtr(engineID),
		CatalogName: core.StringPtr(catalogName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSchemasOptions) SetEngineID(engineID string) *GetSchemasOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *GetSchemasOptions) SetCatalogName(catalogName string) *GetSchemasOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSchemasOptions) SetAuthInstanceID(authInstanceID string) *GetSchemasOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSchemasOptions) SetHeaders(param map[string]string) *GetSchemasOptions {
	options.Headers = param
	return options
}

// GetTableSnapshotsOptions : The GetTableSnapshots options.
type GetTableSnapshotsOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Table name.
	TableName *string `json:"table_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTableSnapshotsOptions : Instantiate GetTableSnapshotsOptions
func (*WatsonxDataV1) NewGetTableSnapshotsOptions(engineID string, catalogName string, schemaName string, tableName string) *GetTableSnapshotsOptions {
	return &GetTableSnapshotsOptions{
		EngineID:    core.StringPtr(engineID),
		CatalogName: core.StringPtr(catalogName),
		SchemaName:  core.StringPtr(schemaName),
		TableName:   core.StringPtr(tableName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetTableSnapshotsOptions) SetEngineID(engineID string) *GetTableSnapshotsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *GetTableSnapshotsOptions) SetCatalogName(catalogName string) *GetTableSnapshotsOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *GetTableSnapshotsOptions) SetSchemaName(schemaName string) *GetTableSnapshotsOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetTableName : Allow user to set TableName
func (_options *GetTableSnapshotsOptions) SetTableName(tableName string) *GetTableSnapshotsOptions {
	_options.TableName = core.StringPtr(tableName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetTableSnapshotsOptions) SetAuthInstanceID(authInstanceID string) *GetTableSnapshotsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTableSnapshotsOptions) SetHeaders(param map[string]string) *GetTableSnapshotsOptions {
	options.Headers = param
	return options
}

// GetTablesOptions : The GetTables options.
type GetTablesOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTablesOptions : Instantiate GetTablesOptions
func (*WatsonxDataV1) NewGetTablesOptions(engineID string, catalogName string, schemaName string) *GetTablesOptions {
	return &GetTablesOptions{
		EngineID:    core.StringPtr(engineID),
		CatalogName: core.StringPtr(catalogName),
		SchemaName:  core.StringPtr(schemaName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetTablesOptions) SetEngineID(engineID string) *GetTablesOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *GetTablesOptions) SetCatalogName(catalogName string) *GetTablesOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *GetTablesOptions) SetSchemaName(schemaName string) *GetTablesOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetTablesOptions) SetAuthInstanceID(authInstanceID string) *GetTablesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTablesOptions) SetHeaders(param map[string]string) *GetTablesOptions {
	options.Headers = param
	return options
}

// ListDataPoliciesOptions : The ListDataPolicies options.
type ListDataPoliciesOptions struct {
	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// catalog name to filter.
	CatalogName *string `json:"catalog_name,omitempty"`

	// policy status to filter.
	Status *string `json:"status,omitempty"`

	// response will include data policy meta data or not.
	IncludeMetadata *bool `json:"include_metadata,omitempty"`

	// response will include data policy rules or not.
	IncludeRules *bool `json:"include_rules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDataPoliciesOptions : Instantiate ListDataPoliciesOptions
func (*WatsonxDataV1) NewListDataPoliciesOptions() *ListDataPoliciesOptions {
	return &ListDataPoliciesOptions{}
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *ListDataPoliciesOptions) SetLhInstanceID(lhInstanceID string) *ListDataPoliciesOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListDataPoliciesOptions) SetAuthInstanceID(authInstanceID string) *ListDataPoliciesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *ListDataPoliciesOptions) SetCatalogName(catalogName string) *ListDataPoliciesOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ListDataPoliciesOptions) SetStatus(status string) *ListDataPoliciesOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetIncludeMetadata : Allow user to set IncludeMetadata
func (_options *ListDataPoliciesOptions) SetIncludeMetadata(includeMetadata bool) *ListDataPoliciesOptions {
	_options.IncludeMetadata = core.BoolPtr(includeMetadata)
	return _options
}

// SetIncludeRules : Allow user to set IncludeRules
func (_options *ListDataPoliciesOptions) SetIncludeRules(includeRules bool) *ListDataPoliciesOptions {
	_options.IncludeRules = core.BoolPtr(includeRules)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataPoliciesOptions) SetHeaders(param map[string]string) *ListDataPoliciesOptions {
	options.Headers = param
	return options
}

// ParseCsvOptions : The ParseCsv options.
type ParseCsvOptions struct {
	// Presto engine name.
	Engine *string `json:"engine" validate:"required"`

	// parse file to data type.
	ParseFile *string `json:"parse_file" validate:"required"`

	// File type.
	FileType *string `json:"file_type" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewParseCsvOptions : Instantiate ParseCsvOptions
func (*WatsonxDataV1) NewParseCsvOptions(engine string, parseFile string, fileType string) *ParseCsvOptions {
	return &ParseCsvOptions{
		Engine:    core.StringPtr(engine),
		ParseFile: core.StringPtr(parseFile),
		FileType:  core.StringPtr(fileType),
	}
}

// SetEngine : Allow user to set Engine
func (_options *ParseCsvOptions) SetEngine(engine string) *ParseCsvOptions {
	_options.Engine = core.StringPtr(engine)
	return _options
}

// SetParseFile : Allow user to set ParseFile
func (_options *ParseCsvOptions) SetParseFile(parseFile string) *ParseCsvOptions {
	_options.ParseFile = core.StringPtr(parseFile)
	return _options
}

// SetFileType : Allow user to set FileType
func (_options *ParseCsvOptions) SetFileType(fileType string) *ParseCsvOptions {
	_options.FileType = core.StringPtr(fileType)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *ParseCsvOptions) SetAccept(accept string) *ParseCsvOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ParseCsvOptions) SetAuthInstanceID(authInstanceID string) *ParseCsvOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ParseCsvOptions) SetHeaders(param map[string]string) *ParseCsvOptions {
	options.Headers = param
	return options
}

// PauseEngineOptions : The PauseEngine options.
type PauseEngineOptions struct {
	// Engine ID to be paused.
	EngineID *string `json:"engine_id" validate:"required"`

	// Created by - Logged in username.
	CreatedBy *string `json:"created_by,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPauseEngineOptions : Instantiate PauseEngineOptions
func (*WatsonxDataV1) NewPauseEngineOptions(engineID string) *PauseEngineOptions {
	return &PauseEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *PauseEngineOptions) SetEngineID(engineID string) *PauseEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *PauseEngineOptions) SetCreatedBy(createdBy string) *PauseEngineOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *PauseEngineOptions) SetAuthInstanceID(authInstanceID string) *PauseEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PauseEngineOptions) SetHeaders(param map[string]string) *PauseEngineOptions {
	options.Headers = param
	return options
}

// PostQueryOptions : The PostQuery options.
type PostQueryOptions struct {
	// Presto engine name.
	Engine *string `json:"engine" validate:"required"`

	// Catalog name.
	Catalog *string `json:"catalog" validate:"required"`

	// Schema name.
	Schema *string `json:"schema" validate:"required"`

	// SQL Query.
	SqlQuery *string `json:"sqlQuery" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostQueryOptions : Instantiate PostQueryOptions
func (*WatsonxDataV1) NewPostQueryOptions(engine string, catalog string, schema string, sqlQuery string) *PostQueryOptions {
	return &PostQueryOptions{
		Engine:   core.StringPtr(engine),
		Catalog:  core.StringPtr(catalog),
		Schema:   core.StringPtr(schema),
		SqlQuery: core.StringPtr(sqlQuery),
	}
}

// SetEngine : Allow user to set Engine
func (_options *PostQueryOptions) SetEngine(engine string) *PostQueryOptions {
	_options.Engine = core.StringPtr(engine)
	return _options
}

// SetCatalog : Allow user to set Catalog
func (_options *PostQueryOptions) SetCatalog(catalog string) *PostQueryOptions {
	_options.Catalog = core.StringPtr(catalog)
	return _options
}

// SetSchema : Allow user to set Schema
func (_options *PostQueryOptions) SetSchema(schema string) *PostQueryOptions {
	_options.Schema = core.StringPtr(schema)
	return _options
}

// SetSqlQuery : Allow user to set SqlQuery
func (_options *PostQueryOptions) SetSqlQuery(sqlQuery string) *PostQueryOptions {
	_options.SqlQuery = core.StringPtr(sqlQuery)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *PostQueryOptions) SetAccept(accept string) *PostQueryOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *PostQueryOptions) SetAuthInstanceID(authInstanceID string) *PostQueryOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostQueryOptions) SetHeaders(param map[string]string) *PostQueryOptions {
	options.Headers = param
	return options
}

// RegisterBucketOptions : The RegisterBucket options.
type RegisterBucketOptions struct {
	// Bucket Details.
	BucketDetails *BucketDetails `json:"bucket_details" validate:"required"`

	// Bucket description.
	Description *string `json:"description" validate:"required"`

	// Table type.
	TableType *string `json:"table_type" validate:"required"`

	// Bucket Type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// Catalog name for the new catalog to be created with bucket.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// Bucket Display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// tags.
	BucketTags []string `json:"bucket_tags,omitempty"`

	// Catalog tags.
	CatalogTags []string `json:"catalog_tags,omitempty"`

	// Thrift URI.
	ThriftURI *string `json:"thrift_uri,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RegisterBucketOptions.TableType property.
// Table type.
const (
	RegisterBucketOptions_TableType_HiveHadoop2 = "hive-hadoop2"
	RegisterBucketOptions_TableType_Iceberg     = "iceberg"
)

// Constants associated with the RegisterBucketOptions.BucketType property.
// Bucket Type.
const (
	RegisterBucketOptions_BucketType_AwsS3  = "aws_s3"
	RegisterBucketOptions_BucketType_IbmCos = "ibm_cos"
	RegisterBucketOptions_BucketType_Minio  = "minio"
)

// Constants associated with the RegisterBucketOptions.ManagedBy property.
// Managed by.
const (
	RegisterBucketOptions_ManagedBy_Customer = "customer"
	RegisterBucketOptions_ManagedBy_Ibm      = "ibm"
)

// NewRegisterBucketOptions : Instantiate RegisterBucketOptions
func (*WatsonxDataV1) NewRegisterBucketOptions(bucketDetails *BucketDetails, description string, tableType string, bucketType string, catalogName string, managedBy string) *RegisterBucketOptions {
	return &RegisterBucketOptions{
		BucketDetails: bucketDetails,
		Description:   core.StringPtr(description),
		TableType:     core.StringPtr(tableType),
		BucketType:    core.StringPtr(bucketType),
		CatalogName:   core.StringPtr(catalogName),
		ManagedBy:     core.StringPtr(managedBy),
	}
}

// SetBucketDetails : Allow user to set BucketDetails
func (_options *RegisterBucketOptions) SetBucketDetails(bucketDetails *BucketDetails) *RegisterBucketOptions {
	_options.BucketDetails = bucketDetails
	return _options
}

// SetDescription : Allow user to set Description
func (_options *RegisterBucketOptions) SetDescription(description string) *RegisterBucketOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTableType : Allow user to set TableType
func (_options *RegisterBucketOptions) SetTableType(tableType string) *RegisterBucketOptions {
	_options.TableType = core.StringPtr(tableType)
	return _options
}

// SetBucketType : Allow user to set BucketType
func (_options *RegisterBucketOptions) SetBucketType(bucketType string) *RegisterBucketOptions {
	_options.BucketType = core.StringPtr(bucketType)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *RegisterBucketOptions) SetCatalogName(catalogName string) *RegisterBucketOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *RegisterBucketOptions) SetManagedBy(managedBy string) *RegisterBucketOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetBucketDisplayName : Allow user to set BucketDisplayName
func (_options *RegisterBucketOptions) SetBucketDisplayName(bucketDisplayName string) *RegisterBucketOptions {
	_options.BucketDisplayName = core.StringPtr(bucketDisplayName)
	return _options
}

// SetBucketTags : Allow user to set BucketTags
func (_options *RegisterBucketOptions) SetBucketTags(bucketTags []string) *RegisterBucketOptions {
	_options.BucketTags = bucketTags
	return _options
}

// SetCatalogTags : Allow user to set CatalogTags
func (_options *RegisterBucketOptions) SetCatalogTags(catalogTags []string) *RegisterBucketOptions {
	_options.CatalogTags = catalogTags
	return _options
}

// SetThriftURI : Allow user to set ThriftURI
func (_options *RegisterBucketOptions) SetThriftURI(thriftURI string) *RegisterBucketOptions {
	_options.ThriftURI = core.StringPtr(thriftURI)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RegisterBucketOptions) SetAuthInstanceID(authInstanceID string) *RegisterBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterBucketOptions) SetHeaders(param map[string]string) *RegisterBucketOptions {
	options.Headers = param
	return options
}

// RemoveCatalogFromEngineOptions : The RemoveCatalogFromEngine options.
type RemoveCatalogFromEngineOptions struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveCatalogFromEngineOptions : Instantiate RemoveCatalogFromEngineOptions
func (*WatsonxDataV1) NewRemoveCatalogFromEngineOptions(catalogName string, engineID string) *RemoveCatalogFromEngineOptions {
	return &RemoveCatalogFromEngineOptions{
		CatalogName: core.StringPtr(catalogName),
		EngineID:    core.StringPtr(engineID),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *RemoveCatalogFromEngineOptions) SetCatalogName(catalogName string) *RemoveCatalogFromEngineOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *RemoveCatalogFromEngineOptions) SetEngineID(engineID string) *RemoveCatalogFromEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *RemoveCatalogFromEngineOptions) SetAccept(accept string) *RemoveCatalogFromEngineOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *RemoveCatalogFromEngineOptions) SetCreatedBy(createdBy string) *RemoveCatalogFromEngineOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RemoveCatalogFromEngineOptions) SetAuthInstanceID(authInstanceID string) *RemoveCatalogFromEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveCatalogFromEngineOptions) SetHeaders(param map[string]string) *RemoveCatalogFromEngineOptions {
	options.Headers = param
	return options
}

// ReplaceDataPolicyOptions : The ReplaceDataPolicy options.
type ReplaceDataPolicyOptions struct {
	// Policy name for PATCH.
	PolicyName *string `json:"policy_name" validate:"required,ne="`

	// catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// data artifact.
	DataArtifact *string `json:"data_artifact" validate:"required"`

	// rules.
	Rules []Rule `json:"rules" validate:"required"`

	// a more detailed description of the policy.
	Description *string `json:"description,omitempty"`

	// data policy status.
	Status *string `json:"status,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceDataPolicyOptions.Status property.
// data policy status.
const (
	ReplaceDataPolicyOptions_Status_Active   = "active"
	ReplaceDataPolicyOptions_Status_Inactive = "inactive"
)

// NewReplaceDataPolicyOptions : Instantiate ReplaceDataPolicyOptions
func (*WatsonxDataV1) NewReplaceDataPolicyOptions(policyName string, catalogName string, dataArtifact string, rules []Rule) *ReplaceDataPolicyOptions {
	return &ReplaceDataPolicyOptions{
		PolicyName:   core.StringPtr(policyName),
		CatalogName:  core.StringPtr(catalogName),
		DataArtifact: core.StringPtr(dataArtifact),
		Rules:        rules,
	}
}

// SetPolicyName : Allow user to set PolicyName
func (_options *ReplaceDataPolicyOptions) SetPolicyName(policyName string) *ReplaceDataPolicyOptions {
	_options.PolicyName = core.StringPtr(policyName)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *ReplaceDataPolicyOptions) SetCatalogName(catalogName string) *ReplaceDataPolicyOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetDataArtifact : Allow user to set DataArtifact
func (_options *ReplaceDataPolicyOptions) SetDataArtifact(dataArtifact string) *ReplaceDataPolicyOptions {
	_options.DataArtifact = core.StringPtr(dataArtifact)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *ReplaceDataPolicyOptions) SetRules(rules []Rule) *ReplaceDataPolicyOptions {
	_options.Rules = rules
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ReplaceDataPolicyOptions) SetDescription(description string) *ReplaceDataPolicyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ReplaceDataPolicyOptions) SetStatus(status string) *ReplaceDataPolicyOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *ReplaceDataPolicyOptions) SetLhInstanceID(lhInstanceID string) *ReplaceDataPolicyOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ReplaceDataPolicyOptions) SetAuthInstanceID(authInstanceID string) *ReplaceDataPolicyOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceDataPolicyOptions) SetHeaders(param map[string]string) *ReplaceDataPolicyOptions {
	options.Headers = param
	return options
}

// ResumeEngineOptions : The ResumeEngine options.
type ResumeEngineOptions struct {
	// Engine ID to be resumed.
	EngineID *string `json:"engine_id" validate:"required"`

	// Created by - logged in username.
	CreatedBy *string `json:"created_by,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewResumeEngineOptions : Instantiate ResumeEngineOptions
func (*WatsonxDataV1) NewResumeEngineOptions(engineID string) *ResumeEngineOptions {
	return &ResumeEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ResumeEngineOptions) SetEngineID(engineID string) *ResumeEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *ResumeEngineOptions) SetCreatedBy(createdBy string) *ResumeEngineOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ResumeEngineOptions) SetAuthInstanceID(authInstanceID string) *ResumeEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResumeEngineOptions) SetHeaders(param map[string]string) *ResumeEngineOptions {
	options.Headers = param
	return options
}

// RollbackSnapshotOptions : The RollbackSnapshot options.
type RollbackSnapshotOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Snapshot id.
	SnapshotID *string `json:"snapshot_id" validate:"required"`

	// Table name.
	TableName *string `json:"table_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRollbackSnapshotOptions : Instantiate RollbackSnapshotOptions
func (*WatsonxDataV1) NewRollbackSnapshotOptions(engineID string, catalogName string, schemaName string, snapshotID string, tableName string) *RollbackSnapshotOptions {
	return &RollbackSnapshotOptions{
		EngineID:    core.StringPtr(engineID),
		CatalogName: core.StringPtr(catalogName),
		SchemaName:  core.StringPtr(schemaName),
		SnapshotID:  core.StringPtr(snapshotID),
		TableName:   core.StringPtr(tableName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RollbackSnapshotOptions) SetEngineID(engineID string) *RollbackSnapshotOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *RollbackSnapshotOptions) SetCatalogName(catalogName string) *RollbackSnapshotOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *RollbackSnapshotOptions) SetSchemaName(schemaName string) *RollbackSnapshotOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetSnapshotID : Allow user to set SnapshotID
func (_options *RollbackSnapshotOptions) SetSnapshotID(snapshotID string) *RollbackSnapshotOptions {
	_options.SnapshotID = core.StringPtr(snapshotID)
	return _options
}

// SetTableName : Allow user to set TableName
func (_options *RollbackSnapshotOptions) SetTableName(tableName string) *RollbackSnapshotOptions {
	_options.TableName = core.StringPtr(tableName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RollbackSnapshotOptions) SetAuthInstanceID(authInstanceID string) *RollbackSnapshotOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RollbackSnapshotOptions) SetHeaders(param map[string]string) *RollbackSnapshotOptions {
	options.Headers = param
	return options
}

// SaveQueryOptions : The SaveQuery options.
type SaveQueryOptions struct {
	// Query name.
	QueryName *string `json:"query_name" validate:"required,ne="`

	// Created by.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Description.
	Description *string `json:"description" validate:"required"`

	// Query string.
	QueryString *string `json:"query_string" validate:"required"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// Engine ID.
	EngineID *string `json:"engine_id,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSaveQueryOptions : Instantiate SaveQueryOptions
func (*WatsonxDataV1) NewSaveQueryOptions(queryName string, createdBy string, description string, queryString string) *SaveQueryOptions {
	return &SaveQueryOptions{
		QueryName:   core.StringPtr(queryName),
		CreatedBy:   core.StringPtr(createdBy),
		Description: core.StringPtr(description),
		QueryString: core.StringPtr(queryString),
	}
}

// SetQueryName : Allow user to set QueryName
func (_options *SaveQueryOptions) SetQueryName(queryName string) *SaveQueryOptions {
	_options.QueryName = core.StringPtr(queryName)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *SaveQueryOptions) SetCreatedBy(createdBy string) *SaveQueryOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *SaveQueryOptions) SetDescription(description string) *SaveQueryOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetQueryString : Allow user to set QueryString
func (_options *SaveQueryOptions) SetQueryString(queryString string) *SaveQueryOptions {
	_options.QueryString = core.StringPtr(queryString)
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *SaveQueryOptions) SetCreatedOn(createdOn string) *SaveQueryOptions {
	_options.CreatedOn = core.StringPtr(createdOn)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *SaveQueryOptions) SetEngineID(engineID string) *SaveQueryOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *SaveQueryOptions) SetAuthInstanceID(authInstanceID string) *SaveQueryOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SaveQueryOptions) SetHeaders(param map[string]string) *SaveQueryOptions {
	options.Headers = param
	return options
}

// TestLHConsoleOptions : The TestLHConsole options.
type TestLHConsoleOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTestLHConsoleOptions : Instantiate TestLHConsoleOptions
func (*WatsonxDataV1) NewTestLHConsoleOptions() *TestLHConsoleOptions {
	return &TestLHConsoleOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *TestLHConsoleOptions) SetHeaders(param map[string]string) *TestLHConsoleOptions {
	options.Headers = param
	return options
}

// UnregisterBucketOptions : The UnregisterBucket options.
type UnregisterBucketOptions struct {
	// Bucket name.
	BucketID *string `json:"bucket_id" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnregisterBucketOptions : Instantiate UnregisterBucketOptions
func (*WatsonxDataV1) NewUnregisterBucketOptions(bucketID string) *UnregisterBucketOptions {
	return &UnregisterBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *UnregisterBucketOptions) SetBucketID(bucketID string) *UnregisterBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UnregisterBucketOptions) SetAuthInstanceID(authInstanceID string) *UnregisterBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UnregisterBucketOptions) SetHeaders(param map[string]string) *UnregisterBucketOptions {
	options.Headers = param
	return options
}

// UpdateBucketOptions : The UpdateBucket options.
type UpdateBucketOptions struct {
	// Bucket ID auto generated during bucket registration.
	BucketID *string `json:"bucket_id" validate:"required"`

	// Access key ID, encrypted during bucket registration.
	AccessKey *string `json:"access_key,omitempty"`

	// Bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// Modified description.
	Description *string `json:"description,omitempty"`

	// Secret access key, encrypted during bucket registration.
	SecretKey *string `json:"secret_key,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBucketOptions : Instantiate UpdateBucketOptions
func (*WatsonxDataV1) NewUpdateBucketOptions(bucketID string) *UpdateBucketOptions {
	return &UpdateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *UpdateBucketOptions) SetBucketID(bucketID string) *UpdateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAccessKey : Allow user to set AccessKey
func (_options *UpdateBucketOptions) SetAccessKey(accessKey string) *UpdateBucketOptions {
	_options.AccessKey = core.StringPtr(accessKey)
	return _options
}

// SetBucketDisplayName : Allow user to set BucketDisplayName
func (_options *UpdateBucketOptions) SetBucketDisplayName(bucketDisplayName string) *UpdateBucketOptions {
	_options.BucketDisplayName = core.StringPtr(bucketDisplayName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateBucketOptions) SetDescription(description string) *UpdateBucketOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetSecretKey : Allow user to set SecretKey
func (_options *UpdateBucketOptions) SetSecretKey(secretKey string) *UpdateBucketOptions {
	_options.SecretKey = core.StringPtr(secretKey)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *UpdateBucketOptions) SetTags(tags []string) *UpdateBucketOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateBucketOptions) SetAuthInstanceID(authInstanceID string) *UpdateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBucketOptions) SetHeaders(param map[string]string) *UpdateBucketOptions {
	options.Headers = param
	return options
}

// UpdateBucketUsersOptions : The UpdateBucketUsers options.
type UpdateBucketUsersOptions struct {
	// Bucket ID for PATCH.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBucketUsersOptions : Instantiate UpdateBucketUsersOptions
func (*WatsonxDataV1) NewUpdateBucketUsersOptions(bucketID string) *UpdateBucketUsersOptions {
	return &UpdateBucketUsersOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *UpdateBucketUsersOptions) SetBucketID(bucketID string) *UpdateBucketUsersOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *UpdateBucketUsersOptions) SetGroups(groups []BucketDbConnGroupsMetadata) *UpdateBucketUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *UpdateBucketUsersOptions) SetUsers(users []BucketDbConnUsersMetadata) *UpdateBucketUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *UpdateBucketUsersOptions) SetLhInstanceID(lhInstanceID string) *UpdateBucketUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateBucketUsersOptions) SetAuthInstanceID(authInstanceID string) *UpdateBucketUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBucketUsersOptions) SetHeaders(param map[string]string) *UpdateBucketUsersOptions {
	options.Headers = param
	return options
}

// UpdateCatalogUsersOptions : The UpdateCatalogUsers options.
type UpdateCatalogUsersOptions struct {
	// Catalog name for PATCH.
	CatalogName *string `json:"catalog_name" validate:"required,ne="`

	// The group list.
	Groups []CatalogGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []CatalogUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCatalogUsersOptions : Instantiate UpdateCatalogUsersOptions
func (*WatsonxDataV1) NewUpdateCatalogUsersOptions(catalogName string) *UpdateCatalogUsersOptions {
	return &UpdateCatalogUsersOptions{
		CatalogName: core.StringPtr(catalogName),
	}
}

// SetCatalogName : Allow user to set CatalogName
func (_options *UpdateCatalogUsersOptions) SetCatalogName(catalogName string) *UpdateCatalogUsersOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *UpdateCatalogUsersOptions) SetGroups(groups []CatalogGroupsMetadata) *UpdateCatalogUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *UpdateCatalogUsersOptions) SetUsers(users []CatalogUsersMetadata) *UpdateCatalogUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *UpdateCatalogUsersOptions) SetLhInstanceID(lhInstanceID string) *UpdateCatalogUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateCatalogUsersOptions) SetAuthInstanceID(authInstanceID string) *UpdateCatalogUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCatalogUsersOptions) SetHeaders(param map[string]string) *UpdateCatalogUsersOptions {
	options.Headers = param
	return options
}

// UpdateDatabaseOptions : The UpdateDatabase options.
type UpdateDatabaseOptions struct {
	// Database ID.
	DatabaseID *string `json:"database_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// database details.
	DatabaseDetails *UpdateDatabaseBodyDatabaseDetails `json:"database_details,omitempty"`

	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name,omitempty"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDatabaseOptions : Instantiate UpdateDatabaseOptions
func (*WatsonxDataV1) NewUpdateDatabaseOptions(databaseID string) *UpdateDatabaseOptions {
	return &UpdateDatabaseOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *UpdateDatabaseOptions) SetDatabaseID(databaseID string) *UpdateDatabaseOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *UpdateDatabaseOptions) SetAccept(accept string) *UpdateDatabaseOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetDatabaseDetails : Allow user to set DatabaseDetails
func (_options *UpdateDatabaseOptions) SetDatabaseDetails(databaseDetails *UpdateDatabaseBodyDatabaseDetails) *UpdateDatabaseOptions {
	_options.DatabaseDetails = databaseDetails
	return _options
}

// SetDatabaseDisplayName : Allow user to set DatabaseDisplayName
func (_options *UpdateDatabaseOptions) SetDatabaseDisplayName(databaseDisplayName string) *UpdateDatabaseOptions {
	_options.DatabaseDisplayName = core.StringPtr(databaseDisplayName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateDatabaseOptions) SetDescription(description string) *UpdateDatabaseOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *UpdateDatabaseOptions) SetTags(tags []string) *UpdateDatabaseOptions {
	_options.Tags = tags
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

// UpdateDbConnUsersOptions : The UpdateDbConnUsers options.
type UpdateDbConnUsersOptions struct {
	// Db connection id for PATCH.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDbConnUsersOptions : Instantiate UpdateDbConnUsersOptions
func (*WatsonxDataV1) NewUpdateDbConnUsersOptions(databaseID string) *UpdateDbConnUsersOptions {
	return &UpdateDbConnUsersOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *UpdateDbConnUsersOptions) SetDatabaseID(databaseID string) *UpdateDbConnUsersOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *UpdateDbConnUsersOptions) SetGroups(groups []BucketDbConnGroupsMetadata) *UpdateDbConnUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *UpdateDbConnUsersOptions) SetUsers(users []BucketDbConnUsersMetadata) *UpdateDbConnUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *UpdateDbConnUsersOptions) SetLhInstanceID(lhInstanceID string) *UpdateDbConnUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateDbConnUsersOptions) SetAuthInstanceID(authInstanceID string) *UpdateDbConnUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDbConnUsersOptions) SetHeaders(param map[string]string) *UpdateDbConnUsersOptions {
	options.Headers = param
	return options
}

// UpdateEngineOptions : The UpdateEngine options.
type UpdateEngineOptions struct {
	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateEngineOptions : Instantiate UpdateEngineOptions
func (*WatsonxDataV1) NewUpdateEngineOptions(engineID string) *UpdateEngineOptions {
	return &UpdateEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateEngineOptions) SetEngineID(engineID string) *UpdateEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *UpdateEngineOptions) SetAccept(accept string) *UpdateEngineOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetCoordinator : Allow user to set Coordinator
func (_options *UpdateEngineOptions) SetCoordinator(coordinator *NodeDescription) *UpdateEngineOptions {
	_options.Coordinator = coordinator
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateEngineOptions) SetDescription(description string) *UpdateEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *UpdateEngineOptions) SetEngineDisplayName(engineDisplayName string) *UpdateEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *UpdateEngineOptions) SetTags(tags []string) *UpdateEngineOptions {
	_options.Tags = tags
	return _options
}

// SetWorker : Allow user to set Worker
func (_options *UpdateEngineOptions) SetWorker(worker *NodeDescription) *UpdateEngineOptions {
	_options.Worker = worker
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEngineOptions) SetHeaders(param map[string]string) *UpdateEngineOptions {
	options.Headers = param
	return options
}

// UpdateEngineUsersOptions : The UpdateEngineUsers options.
type UpdateEngineUsersOptions struct {
	// Engine ID for PATCH.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// The group list.
	Groups []EngineGroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []EngineUsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateEngineUsersOptions : Instantiate UpdateEngineUsersOptions
func (*WatsonxDataV1) NewUpdateEngineUsersOptions(engineID string) *UpdateEngineUsersOptions {
	return &UpdateEngineUsersOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateEngineUsersOptions) SetEngineID(engineID string) *UpdateEngineUsersOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *UpdateEngineUsersOptions) SetGroups(groups []EngineGroupsMetadata) *UpdateEngineUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *UpdateEngineUsersOptions) SetUsers(users []EngineUsersMetadata) *UpdateEngineUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *UpdateEngineUsersOptions) SetLhInstanceID(lhInstanceID string) *UpdateEngineUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateEngineUsersOptions) SetAuthInstanceID(authInstanceID string) *UpdateEngineUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEngineUsersOptions) SetHeaders(param map[string]string) *UpdateEngineUsersOptions {
	options.Headers = param
	return options
}

// UpdateMetastoreUsersOptions : The UpdateMetastoreUsers options.
type UpdateMetastoreUsersOptions struct {
	// Metastore name for PATCH.
	MetastoreName *string `json:"metastore_name" validate:"required,ne="`

	// The group list.
	Groups []GroupsMetadata `json:"groups,omitempty"`

	// The user list.
	Users []UsersMetadata `json:"users,omitempty"`

	// Lake House Instance ID.
	LhInstanceID *string `json:"LhInstanceId,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateMetastoreUsersOptions : Instantiate UpdateMetastoreUsersOptions
func (*WatsonxDataV1) NewUpdateMetastoreUsersOptions(metastoreName string) *UpdateMetastoreUsersOptions {
	return &UpdateMetastoreUsersOptions{
		MetastoreName: core.StringPtr(metastoreName),
	}
}

// SetMetastoreName : Allow user to set MetastoreName
func (_options *UpdateMetastoreUsersOptions) SetMetastoreName(metastoreName string) *UpdateMetastoreUsersOptions {
	_options.MetastoreName = core.StringPtr(metastoreName)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *UpdateMetastoreUsersOptions) SetGroups(groups []GroupsMetadata) *UpdateMetastoreUsersOptions {
	_options.Groups = groups
	return _options
}

// SetUsers : Allow user to set Users
func (_options *UpdateMetastoreUsersOptions) SetUsers(users []UsersMetadata) *UpdateMetastoreUsersOptions {
	_options.Users = users
	return _options
}

// SetLhInstanceID : Allow user to set LhInstanceID
func (_options *UpdateMetastoreUsersOptions) SetLhInstanceID(lhInstanceID string) *UpdateMetastoreUsersOptions {
	_options.LhInstanceID = core.StringPtr(lhInstanceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateMetastoreUsersOptions) SetAuthInstanceID(authInstanceID string) *UpdateMetastoreUsersOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMetastoreUsersOptions) SetHeaders(param map[string]string) *UpdateMetastoreUsersOptions {
	options.Headers = param
	return options
}

// UpdateQueryOptions : The UpdateQuery options.
type UpdateQueryOptions struct {
	// Query name.
	QueryName *string `json:"query_name" validate:"required,ne="`

	// Query string.
	QueryString *string `json:"query_string" validate:"required"`

	// Description.
	Description *string `json:"description" validate:"required"`

	// New query name.
	NewQueryName *string `json:"new_query_name" validate:"required"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateQueryOptions : Instantiate UpdateQueryOptions
func (*WatsonxDataV1) NewUpdateQueryOptions(queryName string, queryString string, description string, newQueryName string) *UpdateQueryOptions {
	return &UpdateQueryOptions{
		QueryName:    core.StringPtr(queryName),
		QueryString:  core.StringPtr(queryString),
		Description:  core.StringPtr(description),
		NewQueryName: core.StringPtr(newQueryName),
	}
}

// SetQueryName : Allow user to set QueryName
func (_options *UpdateQueryOptions) SetQueryName(queryName string) *UpdateQueryOptions {
	_options.QueryName = core.StringPtr(queryName)
	return _options
}

// SetQueryString : Allow user to set QueryString
func (_options *UpdateQueryOptions) SetQueryString(queryString string) *UpdateQueryOptions {
	_options.QueryString = core.StringPtr(queryString)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateQueryOptions) SetDescription(description string) *UpdateQueryOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetNewQueryName : Allow user to set NewQueryName
func (_options *UpdateQueryOptions) SetNewQueryName(newQueryName string) *UpdateQueryOptions {
	_options.NewQueryName = core.StringPtr(newQueryName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateQueryOptions) SetAuthInstanceID(authInstanceID string) *UpdateQueryOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateQueryOptions) SetHeaders(param map[string]string) *UpdateQueryOptions {
	options.Headers = param
	return options
}

// UpdateTableOptions : The UpdateTable options.
type UpdateTableOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Table name.
	TableName *string `json:"table_name" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Add columns.
	AddColumns []UpdateTableBodyAddColumnsItems `json:"add_columns,omitempty"`

	// Drop columns.
	DropColumns []UpdateTableBodyDropColumnsItems `json:"drop_columns,omitempty"`

	// New table name.
	NewTableName *string `json:"new_table_name,omitempty"`

	// Rename columns.
	RenameColumns []UpdateTableBodyRenameColumnsItems `json:"rename_columns,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTableOptions : Instantiate UpdateTableOptions
func (*WatsonxDataV1) NewUpdateTableOptions(engineID string, catalogName string, schemaName string, tableName string) *UpdateTableOptions {
	return &UpdateTableOptions{
		EngineID:    core.StringPtr(engineID),
		CatalogName: core.StringPtr(catalogName),
		SchemaName:  core.StringPtr(schemaName),
		TableName:   core.StringPtr(tableName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateTableOptions) SetEngineID(engineID string) *UpdateTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogName : Allow user to set CatalogName
func (_options *UpdateTableOptions) SetCatalogName(catalogName string) *UpdateTableOptions {
	_options.CatalogName = core.StringPtr(catalogName)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *UpdateTableOptions) SetSchemaName(schemaName string) *UpdateTableOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetTableName : Allow user to set TableName
func (_options *UpdateTableOptions) SetTableName(tableName string) *UpdateTableOptions {
	_options.TableName = core.StringPtr(tableName)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *UpdateTableOptions) SetAccept(accept string) *UpdateTableOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAddColumns : Allow user to set AddColumns
func (_options *UpdateTableOptions) SetAddColumns(addColumns []UpdateTableBodyAddColumnsItems) *UpdateTableOptions {
	_options.AddColumns = addColumns
	return _options
}

// SetDropColumns : Allow user to set DropColumns
func (_options *UpdateTableOptions) SetDropColumns(dropColumns []UpdateTableBodyDropColumnsItems) *UpdateTableOptions {
	_options.DropColumns = dropColumns
	return _options
}

// SetNewTableName : Allow user to set NewTableName
func (_options *UpdateTableOptions) SetNewTableName(newTableName string) *UpdateTableOptions {
	_options.NewTableName = core.StringPtr(newTableName)
	return _options
}

// SetRenameColumns : Allow user to set RenameColumns
func (_options *UpdateTableOptions) SetRenameColumns(renameColumns []UpdateTableBodyRenameColumnsItems) *UpdateTableOptions {
	_options.RenameColumns = renameColumns
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

// UplaodCsvOptions : The UplaodCsv options.
type UplaodCsvOptions struct {
	// Presto engine name.
	Engine *string `json:"engine" validate:"required"`

	// Catalog name.
	Catalog *string `json:"catalog" validate:"required"`

	// Schema name.
	Schema *string `json:"schema" validate:"required"`

	// table name.
	TableName *string `json:"tableName" validate:"required"`

	// ingestion job name.
	IngestionJobName *string `json:"ingestionJobName" validate:"required"`

	// Scheduled.
	Scheduled *string `json:"scheduled" validate:"required"`

	// Created by.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Target table.
	TargetTable *string `json:"targetTable" validate:"required"`

	// Headers.
	HeadersVar *string `json:"headers" validate:"required"`

	// csv.
	Csv *string `json:"csv" validate:"required"`

	// The type of the response:  or *_/_*.
	Accept *string `json:"Accept,omitempty"`

	// Instance ID.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUplaodCsvOptions : Instantiate UplaodCsvOptions
func (*WatsonxDataV1) NewUplaodCsvOptions(engine string, catalog string, schema string, tableName string, ingestionJobName string, scheduled string, createdBy string, targetTable string, headersVar string, csv string) *UplaodCsvOptions {
	return &UplaodCsvOptions{
		Engine:           core.StringPtr(engine),
		Catalog:          core.StringPtr(catalog),
		Schema:           core.StringPtr(schema),
		TableName:        core.StringPtr(tableName),
		IngestionJobName: core.StringPtr(ingestionJobName),
		Scheduled:        core.StringPtr(scheduled),
		CreatedBy:        core.StringPtr(createdBy),
		TargetTable:      core.StringPtr(targetTable),
		HeadersVar:       core.StringPtr(headersVar),
		Csv:              core.StringPtr(csv),
	}
}

// SetEngine : Allow user to set Engine
func (_options *UplaodCsvOptions) SetEngine(engine string) *UplaodCsvOptions {
	_options.Engine = core.StringPtr(engine)
	return _options
}

// SetCatalog : Allow user to set Catalog
func (_options *UplaodCsvOptions) SetCatalog(catalog string) *UplaodCsvOptions {
	_options.Catalog = core.StringPtr(catalog)
	return _options
}

// SetSchema : Allow user to set Schema
func (_options *UplaodCsvOptions) SetSchema(schema string) *UplaodCsvOptions {
	_options.Schema = core.StringPtr(schema)
	return _options
}

// SetTableName : Allow user to set TableName
func (_options *UplaodCsvOptions) SetTableName(tableName string) *UplaodCsvOptions {
	_options.TableName = core.StringPtr(tableName)
	return _options
}

// SetIngestionJobName : Allow user to set IngestionJobName
func (_options *UplaodCsvOptions) SetIngestionJobName(ingestionJobName string) *UplaodCsvOptions {
	_options.IngestionJobName = core.StringPtr(ingestionJobName)
	return _options
}

// SetScheduled : Allow user to set Scheduled
func (_options *UplaodCsvOptions) SetScheduled(scheduled string) *UplaodCsvOptions {
	_options.Scheduled = core.StringPtr(scheduled)
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *UplaodCsvOptions) SetCreatedBy(createdBy string) *UplaodCsvOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetTargetTable : Allow user to set TargetTable
func (_options *UplaodCsvOptions) SetTargetTable(targetTable string) *UplaodCsvOptions {
	_options.TargetTable = core.StringPtr(targetTable)
	return _options
}

// SetHeadersVar : Allow user to set HeadersVar
func (_options *UplaodCsvOptions) SetHeadersVar(headersVar string) *UplaodCsvOptions {
	_options.HeadersVar = core.StringPtr(headersVar)
	return _options
}

// SetCsv : Allow user to set Csv
func (_options *UplaodCsvOptions) SetCsv(csv string) *UplaodCsvOptions {
	_options.Csv = core.StringPtr(csv)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *UplaodCsvOptions) SetAccept(accept string) *UplaodCsvOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UplaodCsvOptions) SetAuthInstanceID(authInstanceID string) *UplaodCsvOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UplaodCsvOptions) SetHeaders(param map[string]string) *UplaodCsvOptions {
	options.Headers = param
	return options
}

// Bucket : Bucket.
type Bucket struct {
	// Username who created the bucket.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Creation date.
	CreatedOn *string `json:"created_on" validate:"required"`

	// Bucket Description.
	Description *string `json:"description" validate:"required"`

	// Bucket endpoint.
	Endpoint *string `json:"endpoint" validate:"required"`

	// Managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// Mark bucket active or inactive.
	State *string `json:"state" validate:"required"`

	// Tags.
	Tags []string `json:"tags" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs" validate:"required"`

	// Bucket Display Name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// Bucket ID auto generated during bucket registration.
	BucketID *string `json:"bucket_id,omitempty"`

	// Actual bucket name.
	BucketName *string `json:"bucket_name" validate:"required"`

	// Bucket Type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// Actions.
	Actions []string `json:"actions,omitempty"`
}

// Constants associated with the Bucket.ManagedBy property.
// Managed by.
const (
	Bucket_ManagedBy_Customer = "Customer"
	Bucket_ManagedBy_Ibm      = "IBM"
)

// Constants associated with the Bucket.State property.
// Mark bucket active or inactive.
const (
	Bucket_State_Active   = "active"
	Bucket_State_Inactive = "inactive"
)

// Constants associated with the Bucket.BucketType property.
// Bucket Type.
const (
	Bucket_BucketType_AmazonS3 = "amazon_s3"
	Bucket_BucketType_AwsS3    = "aws_s3"
	Bucket_BucketType_IbmCos   = "ibm_cos"
	Bucket_BucketType_Minio    = "minio"
)

// UnmarshalBucket unmarshals an instance of Bucket from the specified map of raw messages.
func UnmarshalBucket(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Bucket)
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
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
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
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
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
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_type", &obj.BucketType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketDbConnGroupsMetadata : BucketDbConnGroupsMetadata struct
type BucketDbConnGroupsMetadata struct {
	// The group id.
	GroupID *string `json:"group_id" validate:"required"`

	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`
}

// Constants associated with the BucketDbConnGroupsMetadata.Permission property.
// Eligible permission to the resource.
const (
	BucketDbConnGroupsMetadata_Permission_CanAdminister = "can_administer"
	BucketDbConnGroupsMetadata_Permission_CanRead       = "can_read"
	BucketDbConnGroupsMetadata_Permission_CanWrite      = "can_write"
)

// NewBucketDbConnGroupsMetadata : Instantiate BucketDbConnGroupsMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewBucketDbConnGroupsMetadata(groupID string, permission string) (_model *BucketDbConnGroupsMetadata, err error) {
	_model = &BucketDbConnGroupsMetadata{
		GroupID:    core.StringPtr(groupID),
		Permission: core.StringPtr(permission),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalBucketDbConnGroupsMetadata unmarshals an instance of BucketDbConnGroupsMetadata from the specified map of raw messages.
func UnmarshalBucketDbConnGroupsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketDbConnGroupsMetadata)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketDbConnUsersMetadata : BucketDbConnUsersMetadata struct
type BucketDbConnUsersMetadata struct {
	// The user name.
	UserName *string `json:"user_name" validate:"required"`

	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`
}

// Constants associated with the BucketDbConnUsersMetadata.Permission property.
// Eligible permission to the resource.
const (
	BucketDbConnUsersMetadata_Permission_CanAdminister = "can_administer"
	BucketDbConnUsersMetadata_Permission_CanRead       = "can_read"
	BucketDbConnUsersMetadata_Permission_CanWrite      = "can_write"
)

// NewBucketDbConnUsersMetadata : Instantiate BucketDbConnUsersMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewBucketDbConnUsersMetadata(userName string, permission string) (_model *BucketDbConnUsersMetadata, err error) {
	_model = &BucketDbConnUsersMetadata{
		UserName:   core.StringPtr(userName),
		Permission: core.StringPtr(permission),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalBucketDbConnUsersMetadata unmarshals an instance of BucketDbConnUsersMetadata from the specified map of raw messages.
func UnmarshalBucketDbConnUsersMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketDbConnUsersMetadata)
	err = core.UnmarshalPrimitive(m, "user_name", &obj.UserName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketDetails : Bucket Details.
type BucketDetails struct {
	// Access key ID, encrypted during bucket registration.
	AccessKey *string `json:"access_key,omitempty"`

	// Actual bucket name.
	BucketName *string `json:"bucket_name" validate:"required"`

	// Cos endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// Secret access key, encrypted during bucket registration.
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewBucketDetails : Instantiate BucketDetails (Generic Model Constructor)
func (*WatsonxDataV1) NewBucketDetails(bucketName string) (_model *BucketDetails, err error) {
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

// BucketPolicies : BucketPolicies struct
type BucketPolicies struct {
	// Policy version.
	PolicyVersion *string `json:"policy_version,omitempty"`

	// The policy name.
	PolicyName *string `json:"policy_name,omitempty"`
}

// UnmarshalBucketPolicies unmarshals an instance of BucketPolicies from the specified map of raw messages.
func UnmarshalBucketPolicies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketPolicies)
	err = core.UnmarshalPrimitive(m, "policy_version", &obj.PolicyVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogGroupsMetadata : CatalogGroupsMetadata struct
type CatalogGroupsMetadata struct {
	// The group id.
	GroupID *string `json:"group_id" validate:"required"`

	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`
}

// Constants associated with the CatalogGroupsMetadata.Permission property.
// Eligible permission to the resource.
const (
	CatalogGroupsMetadata_Permission_CanAdminister = "can_administer"
	CatalogGroupsMetadata_Permission_CanUse        = "can_use"
)

// NewCatalogGroupsMetadata : Instantiate CatalogGroupsMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewCatalogGroupsMetadata(groupID string, permission string) (_model *CatalogGroupsMetadata, err error) {
	_model = &CatalogGroupsMetadata{
		GroupID:    core.StringPtr(groupID),
		Permission: core.StringPtr(permission),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCatalogGroupsMetadata unmarshals an instance of CatalogGroupsMetadata from the specified map of raw messages.
func UnmarshalCatalogGroupsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogGroupsMetadata)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogPolicies : CatalogPolicies struct
type CatalogPolicies struct {
	// The policy name.
	PolicyName *string `json:"policy_name,omitempty"`

	// Policy version.
	PolicyVersion *string `json:"policy_version,omitempty"`
}

// UnmarshalCatalogPolicies unmarshals an instance of CatalogPolicies from the specified map of raw messages.
func UnmarshalCatalogPolicies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogPolicies)
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_version", &obj.PolicyVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogUsersMetadata : CatalogUsersMetadata struct
type CatalogUsersMetadata struct {
	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`

	// The user name.
	UserName *string `json:"user_name" validate:"required"`
}

// Constants associated with the CatalogUsersMetadata.Permission property.
// Eligible permission to the resource.
const (
	CatalogUsersMetadata_Permission_CanAdminister = "can_administer"
	CatalogUsersMetadata_Permission_CanUse        = "can_use"
)

// NewCatalogUsersMetadata : Instantiate CatalogUsersMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewCatalogUsersMetadata(permission string, userName string) (_model *CatalogUsersMetadata, err error) {
	_model = &CatalogUsersMetadata{
		Permission: core.StringPtr(permission),
		UserName:   core.StringPtr(userName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCatalogUsersMetadata unmarshals an instance of CatalogUsersMetadata from the specified map of raw messages.
func UnmarshalCatalogUsersMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogUsersMetadata)
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_name", &obj.UserName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDataPolicyCreatedBody : Create data policy success.
type CreateDataPolicyCreatedBody struct {
	// create data policy.
	DataPolicy *CreateDataPolicySchema `json:"data_policy" validate:"required"`

	Metadata *DataPolicyMetadata `json:"metadata" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalCreateDataPolicyCreatedBody unmarshals an instance of CreateDataPolicyCreatedBody from the specified map of raw messages.
func UnmarshalCreateDataPolicyCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDataPolicyCreatedBody)
	err = core.UnmarshalModel(m, "data_policy", &obj.DataPolicy, UnmarshalCreateDataPolicySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalDataPolicyMetadata)
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

// CreateDataPolicySchema : create data policy.
type CreateDataPolicySchema struct {
	// catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// data artifact.
	DataArtifact *string `json:"data_artifact" validate:"required"`

	// a more detailed description of the policy.
	Description *string `json:"description,omitempty"`

	// the displayed name for data policy.
	PolicyName *string `json:"policy_name" validate:"required"`

	// rules.
	Rules []Rule `json:"rules" validate:"required"`

	// data policy status.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the CreateDataPolicySchema.Status property.
// data policy status.
const (
	CreateDataPolicySchema_Status_Active   = "active"
	CreateDataPolicySchema_Status_Inactive = "inactive"
)

// NewCreateDataPolicySchema : Instantiate CreateDataPolicySchema (Generic Model Constructor)
func (*WatsonxDataV1) NewCreateDataPolicySchema(catalogName string, dataArtifact string, policyName string, rules []Rule) (_model *CreateDataPolicySchema, err error) {
	_model = &CreateDataPolicySchema{
		CatalogName:  core.StringPtr(catalogName),
		DataArtifact: core.StringPtr(dataArtifact),
		PolicyName:   core.StringPtr(policyName),
		Rules:        rules,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCreateDataPolicySchema unmarshals an instance of CreateDataPolicySchema from the specified map of raw messages.
func UnmarshalCreateDataPolicySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDataPolicySchema)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_artifact", &obj.DataArtifact)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataPolicies : DataPolicies struct
type DataPolicies struct {
	// Associate catalog.
	AssociateCatalog *string `json:"associate_catalog,omitempty"`

	// For resource policy, it's resource name like engin1. And for data policy it's policy name.
	PolicyName *string `json:"policy_name,omitempty"`

	// Policy version.
	PolicyVersion *string `json:"policy_version,omitempty"`
}

// UnmarshalDataPolicies unmarshals an instance of DataPolicies from the specified map of raw messages.
func UnmarshalDataPolicies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataPolicies)
	err = core.UnmarshalPrimitive(m, "associate_catalog", &obj.AssociateCatalog)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_version", &obj.PolicyVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataPolicyMetadata : DataPolicyMetadata struct
type DataPolicyMetadata struct {
	// an identifier for the creator of the policy.
	Creator *string `json:"creator,omitempty"`

	// a more detailed description of the rule.
	Description *string `json:"description,omitempty"`

	// an identifier for the last modifier of the policy.
	Modifier *string `json:"modifier,omitempty"`

	// an unique identifier for the policy.
	Pid *string `json:"pid,omitempty"`

	// policy name.
	PolicyName *string `json:"policy_name,omitempty"`

	// time when the policy was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// data policy version.
	Version *string `json:"version,omitempty"`

	// time when the policy was created.
	CreatedAt *string `json:"created_at,omitempty"`
}

// UnmarshalDataPolicyMetadata unmarshals an instance of DataPolicyMetadata from the specified map of raw messages.
func UnmarshalDataPolicyMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataPolicyMetadata)
	err = core.UnmarshalPrimitive(m, "creator", &obj.Creator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modifier", &obj.Modifier)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pid", &obj.Pid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DbConnPolicies : DbConnPolicies struct
type DbConnPolicies struct {
	// The policy name.
	PolicyName *string `json:"policy_name,omitempty"`

	// Policy version.
	PolicyVersion *string `json:"policy_version,omitempty"`
}

// UnmarshalDbConnPolicies unmarshals an instance of DbConnPolicies from the specified map of raw messages.
func UnmarshalDbConnPolicies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DbConnPolicies)
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_version", &obj.PolicyVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DefaultPolicySchema : AMS default schema.
type DefaultPolicySchema struct {
	// default grouping policies.
	GroupingPolicies []GroupingPolicyMetadata `json:"grouping_policies,omitempty"`

	// casbin model.
	Model *string `json:"model,omitempty"`

	// default policies.
	Policies []PolicyMetadata `json:"policies,omitempty"`
}

// UnmarshalDefaultPolicySchema unmarshals an instance of DefaultPolicySchema from the specified map of raw messages.
func UnmarshalDefaultPolicySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultPolicySchema)
	err = core.UnmarshalModel(m, "grouping_policies", &obj.GroupingPolicies, UnmarshalGroupingPolicyMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "policies", &obj.Policies, UnmarshalPolicyMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteTableBodyDeleteTablesItems : Delete tables items.
type DeleteTableBodyDeleteTablesItems struct {
	// Catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Schema name.
	SchemaName *string `json:"schema_name,omitempty"`

	// Table name.
	TableName *string `json:"table_name,omitempty"`
}

// UnmarshalDeleteTableBodyDeleteTablesItems unmarshals an instance of DeleteTableBodyDeleteTablesItems from the specified map of raw messages.
func UnmarshalDeleteTableBodyDeleteTablesItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteTableBodyDeleteTablesItems)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schema_name", &obj.SchemaName)
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

// EngineDetail : EngineDetail.
type EngineDetail struct {
	// Group ID.
	GroupID *string `json:"group_id,omitempty"`

	// Region - place holder.
	Region *string `json:"region,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Type like presto, netezza,..
	Type *string `json:"type,omitempty"`

	// Version of the engine.
	Version *string `json:"version,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`

	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Engine host name.
	HostName *string `json:"host_name,omitempty"`

	// Engine status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`
}

// UnmarshalEngineDetail unmarshals an instance of EngineDetail from the specified map of raw messages.
func UnmarshalEngineDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineDetail)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
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
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
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
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
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
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineDetailsBody : Node details.
type EngineDetailsBody struct {
	// Node details.
	Worker *NodeDescriptionBody `json:"worker,omitempty"`

	// Node details.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`
}

// Constants associated with the EngineDetailsBody.SizeConfig property.
// Size config.
const (
	EngineDetailsBody_SizeConfig_ComputeOptimized = "compute_optimized"
	EngineDetailsBody_SizeConfig_Custom           = "custom"
	EngineDetailsBody_SizeConfig_Large            = "large"
	EngineDetailsBody_SizeConfig_Medium           = "medium"
	EngineDetailsBody_SizeConfig_Small            = "small"
	EngineDetailsBody_SizeConfig_Starter          = "starter"
	EngineDetailsBody_SizeConfig_StorageOptimized = "storage_optimized"
)

// UnmarshalEngineDetailsBody unmarshals an instance of EngineDetailsBody from the specified map of raw messages.
func UnmarshalEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineDetailsBody)
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescriptionBody)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineGroupsMetadata : EngineGroupsMetadata struct
type EngineGroupsMetadata struct {
	// The group id.
	GroupID *string `json:"group_id" validate:"required"`

	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`
}

// Constants associated with the EngineGroupsMetadata.Permission property.
// Eligible permission to the resource.
const (
	EngineGroupsMetadata_Permission_CanAdminister = "can_administer"
	EngineGroupsMetadata_Permission_CanManage     = "can_manage"
	EngineGroupsMetadata_Permission_CanUse        = "can_use"
)

// NewEngineGroupsMetadata : Instantiate EngineGroupsMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewEngineGroupsMetadata(groupID string, permission string) (_model *EngineGroupsMetadata, err error) {
	_model = &EngineGroupsMetadata{
		GroupID:    core.StringPtr(groupID),
		Permission: core.StringPtr(permission),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEngineGroupsMetadata unmarshals an instance of EngineGroupsMetadata from the specified map of raw messages.
func UnmarshalEngineGroupsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineGroupsMetadata)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnginePolicies : EnginePolicies struct
type EnginePolicies struct {
	// The policy name.
	PolicyName *string `json:"policy_name,omitempty"`

	// Policy version.
	PolicyVersion *string `json:"policy_version,omitempty"`
}

// UnmarshalEnginePolicies unmarshals an instance of EnginePolicies from the specified map of raw messages.
func UnmarshalEnginePolicies(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnginePolicies)
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_version", &obj.PolicyVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineUsersMetadata : EngineUsersMetadata struct
type EngineUsersMetadata struct {
	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`

	// The user name.
	UserName *string `json:"user_name" validate:"required"`
}

// Constants associated with the EngineUsersMetadata.Permission property.
// Eligible permission to the resource.
const (
	EngineUsersMetadata_Permission_CanAdminister = "can_administer"
	EngineUsersMetadata_Permission_CanManage     = "can_manage"
	EngineUsersMetadata_Permission_CanUse        = "can_use"
)

// NewEngineUsersMetadata : Instantiate EngineUsersMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewEngineUsersMetadata(permission string, userName string) (_model *EngineUsersMetadata, err error) {
	_model = &EngineUsersMetadata{
		Permission: core.StringPtr(permission),
		UserName:   core.StringPtr(userName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEngineUsersMetadata unmarshals an instance of EngineUsersMetadata from the specified map of raw messages.
func UnmarshalEngineUsersMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineUsersMetadata)
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_name", &obj.UserName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvaluationResultSchema : Evaluation result schema.
type EvaluationResultSchema struct {
	// resource list.
	Resources []ResourceWithResult `json:"resources,omitempty"`
}

// UnmarshalEvaluationResultSchema unmarshals an instance of EvaluationResultSchema from the specified map of raw messages.
func UnmarshalEvaluationResultSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvaluationResultSchema)
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceWithResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplainAnalyzeStatementCreatedBody : explainAnalyzeStatement OK.
type ExplainAnalyzeStatementCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// explainAnalyzeStatement result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalExplainAnalyzeStatementCreatedBody unmarshals an instance of ExplainAnalyzeStatementCreatedBody from the specified map of raw messages.
func UnmarshalExplainAnalyzeStatementCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplainAnalyzeStatementCreatedBody)
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

// ExplainStatementCreatedBody : ExplainStatement OK.
type ExplainStatementCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalExplainStatementCreatedBody unmarshals an instance of ExplainStatementCreatedBody from the specified map of raw messages.
func UnmarshalExplainStatementCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplainStatementCreatedBody)
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

// GetBucketObjectsOKBody : GetBucketObjects OK.
type GetBucketObjectsOKBody struct {
	// Bucket objects.
	Objects []string `json:"objects" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetBucketObjectsOKBody unmarshals an instance of GetBucketObjectsOKBody from the specified map of raw messages.
func UnmarshalGetBucketObjectsOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketObjectsOKBody)
	err = core.UnmarshalPrimitive(m, "objects", &obj.Objects)
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

// GetBucketUsersSchema : Get bucket users schema.
type GetBucketUsersSchema struct {
	// The bucket id.
	BucketID *string `json:"bucket_id" validate:"required"`

	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// Total number of users and groups.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`
}

// UnmarshalGetBucketUsersSchema unmarshals an instance of GetBucketUsersSchema from the specified map of raw messages.
func UnmarshalGetBucketUsersSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketUsersSchema)
	err = core.UnmarshalPrimitive(m, "bucket_id", &obj.BucketID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalBucketDbConnGroupsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalBucketDbConnUsersMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketsOKBody : GetBuckets OK.
type GetBucketsOKBody struct {
	// Buckets.
	Buckets []Bucket `json:"buckets" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetBucketsOKBody unmarshals an instance of GetBucketsOKBody from the specified map of raw messages.
func UnmarshalGetBucketsOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketsOKBody)
	err = core.UnmarshalModel(m, "buckets", &obj.Buckets, UnmarshalBucket)
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

// GetCatalogUsersSchema : Get catalog users schema.
type GetCatalogUsersSchema struct {
	// Total number of users and groups.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The user list.
	Users []CatalogUsersMetadata `json:"users,omitempty"`

	// The catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// The group list.
	Groups []CatalogGroupsMetadata `json:"groups,omitempty"`
}

// UnmarshalGetCatalogUsersSchema unmarshals an instance of GetCatalogUsersSchema from the specified map of raw messages.
func UnmarshalGetCatalogUsersSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetCatalogUsersSchema)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalCatalogUsersMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalCatalogGroupsMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDbConnUsersSchema : Get Db connection users schema.
type GetDbConnUsersSchema struct {
	// The group list.
	Groups []BucketDbConnGroupsMetadata `json:"groups,omitempty"`

	// Total number of users and groups.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The user list.
	Users []BucketDbConnUsersMetadata `json:"users,omitempty"`

	// The db connection id.
	DatabaseID *string `json:"database_id" validate:"required"`
}

// UnmarshalGetDbConnUsersSchema unmarshals an instance of GetDbConnUsersSchema from the specified map of raw messages.
func UnmarshalGetDbConnUsersSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDbConnUsersSchema)
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalBucketDbConnGroupsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalBucketDbConnUsersMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_id", &obj.DatabaseID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEngineUsersSchema : Get engine users schema.
type GetEngineUsersSchema struct {
	// The engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// The group list.
	Groups []EngineGroupsMetadata `json:"groups,omitempty"`

	// Total number of users and groups.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The user list.
	Users []EngineUsersMetadata `json:"users,omitempty"`
}

// UnmarshalGetEngineUsersSchema unmarshals an instance of GetEngineUsersSchema from the specified map of raw messages.
func UnmarshalGetEngineUsersSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEngineUsersSchema)
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalEngineGroupsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalEngineUsersMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEnginesOKBody : getEngines.
type GetEnginesOKBody struct {
	Engines []EngineDetail `json:"engines" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetEnginesOKBody unmarshals an instance of GetEnginesOKBody from the specified map of raw messages.
func UnmarshalGetEnginesOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEnginesOKBody)
	err = core.UnmarshalModel(m, "engines", &obj.Engines, UnmarshalEngineDetail)
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

// GetMetastoreUsersSchema : Get metastore users schema.
type GetMetastoreUsersSchema struct {
	// The group list.
	Groups []GroupsMetadata `json:"groups,omitempty"`

	// The metastore name.
	MetastoreName *string `json:"metastore_name" validate:"required"`

	// Total number of users and groups.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The user list.
	Users []UsersMetadata `json:"users,omitempty"`
}

// UnmarshalGetMetastoreUsersSchema unmarshals an instance of GetMetastoreUsersSchema from the specified map of raw messages.
func UnmarshalGetMetastoreUsersSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMetastoreUsersSchema)
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalGroupsMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_name", &obj.MetastoreName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalUsersMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetMetastoresOKBody : GetMetastores OK.
type GetMetastoresOKBody struct {
	// Metastores.
	Catalogs []Metastore `json:"catalogs" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetMetastoresOKBody unmarshals an instance of GetMetastoresOKBody from the specified map of raw messages.
func UnmarshalGetMetastoresOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMetastoresOKBody)
	err = core.UnmarshalModel(m, "catalogs", &obj.Catalogs, UnmarshalMetastore)
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

// GetQueriesOKBody : GetQueries OK.
type GetQueriesOKBody struct {
	// Queries.
	Queries []Query `json:"queries" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalGetQueriesOKBody unmarshals an instance of GetQueriesOKBody from the specified map of raw messages.
func UnmarshalGetQueriesOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetQueriesOKBody)
	err = core.UnmarshalModel(m, "queries", &obj.Queries, UnmarshalQuery)
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

// GetSchemasOKBody : GetSchemas OK.
type GetSchemasOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Schemas.
	Schemas []string `json:"schemas" validate:"required"`
}

// UnmarshalGetSchemasOKBody unmarshals an instance of GetSchemasOKBody from the specified map of raw messages.
func UnmarshalGetSchemasOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSchemasOKBody)
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

// GetTableSnapshotsOKBody : TableSnapshot OK.
type GetTableSnapshotsOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Snapshots.
	Snapshots []TableSnapshot `json:"snapshots" validate:"required"`
}

// UnmarshalGetTableSnapshotsOKBody unmarshals an instance of GetTableSnapshotsOKBody from the specified map of raw messages.
func UnmarshalGetTableSnapshotsOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTableSnapshotsOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "snapshots", &obj.Snapshots, UnmarshalTableSnapshot)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTablesOKBody : GetTables OK.
type GetTablesOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Tables.
	Tables []string `json:"tables" validate:"required"`
}

// UnmarshalGetTablesOKBody unmarshals an instance of GetTablesOKBody from the specified map of raw messages.
func UnmarshalGetTablesOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTablesOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupingPolicyMetadata : GroupingPolicyMetadata struct
type GroupingPolicyMetadata struct {
	// domain.
	Domain *string `json:"domain,omitempty"`

	// inheritor.
	Inheritor *string `json:"inheritor,omitempty"`

	// role.
	Role *string `json:"role,omitempty"`
}

// UnmarshalGroupingPolicyMetadata unmarshals an instance of GroupingPolicyMetadata from the specified map of raw messages.
func UnmarshalGroupingPolicyMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupingPolicyMetadata)
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "inheritor", &obj.Inheritor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupsMetadata : Groups metadata.
type GroupsMetadata struct {
	// The group id.
	GroupID *string `json:"group_id" validate:"required"`

	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`
}

// Constants associated with the GroupsMetadata.Permission property.
// Eligible permission to the resource.
const (
	GroupsMetadata_Permission_CanAdminister = "can_administer"
	GroupsMetadata_Permission_CanManage     = "can_manage"
	GroupsMetadata_Permission_CanUse        = "can_use"
)

// NewGroupsMetadata : Instantiate GroupsMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewGroupsMetadata(groupID string, permission string) (_model *GroupsMetadata, err error) {
	_model = &GroupsMetadata{
		GroupID:    core.StringPtr(groupID),
		Permission: core.StringPtr(permission),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalGroupsMetadata unmarshals an instance of GroupsMetadata from the specified map of raw messages.
func UnmarshalGroupsMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupsMetadata)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Metastore : Metastore.
type Metastore struct {
	// Name for the metastore.
	CatalogName *string `json:"catalog_name,omitempty"`

	// IBM thrift uri hostname.
	Hostname *string `json:"hostname,omitempty"`

	// Managed by.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Metastore status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	Actions []string `json:"actions,omitempty"`

	// Associated buckets items.
	AssociatedBuckets []string `json:"associated_buckets,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Customer thrift uri.
	ThriftURI *string `json:"thrift_uri,omitempty"`

	// Table type.
	CatalogType *string `json:"catalog_type,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`

	// Associated databases items.
	AssociatedDatabases []string `json:"associated_databases,omitempty"`

	// Associated engines items.
	AssociatedEngines []string `json:"associated_engines,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// IBM thrift uri port.
	Port *string `json:"port,omitempty"`
}

// Constants associated with the Metastore.ManagedBy property.
// Managed by.
const (
	Metastore_ManagedBy_Customer = "customer"
	Metastore_ManagedBy_Ibm      = "ibm"
)

// UnmarshalMetastore unmarshals an instance of Metastore from the specified map of raw messages.
func UnmarshalMetastore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metastore)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
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
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_buckets", &obj.AssociatedBuckets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "thrift_uri", &obj.ThriftURI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
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
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
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

// PauseEngineCreatedBody : PauseEngineBody OK.
type PauseEngineCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalPauseEngineCreatedBody unmarshals an instance of PauseEngineCreatedBody from the specified map of raw messages.
func UnmarshalPauseEngineCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PauseEngineCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyListSchema : PolicyListSchema struct
type PolicyListSchema struct {
	// policy collection.
	Policies []PolicySchema `json:"policies" validate:"required"`

	// Total number of policies.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalPolicyListSchema unmarshals an instance of PolicyListSchema from the specified map of raw messages.
func UnmarshalPolicyListSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyListSchema)
	err = core.UnmarshalModel(m, "policies", &obj.Policies, UnmarshalPolicySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyMetadata : PolicyMetadata struct
type PolicyMetadata struct {
	// subject.
	Subject *string `json:"subject,omitempty"`

	// action array.
	Actions []string `json:"actions,omitempty"`

	// domain.
	Domain *string `json:"domain,omitempty"`

	// object.
	Object *string `json:"object,omitempty"`
}

// UnmarshalPolicyMetadata unmarshals an instance of PolicyMetadata from the specified map of raw messages.
func UnmarshalPolicyMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyMetadata)
	err = core.UnmarshalPrimitive(m, "subject", &obj.Subject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "object", &obj.Object)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicySchema : data policy.
type PolicySchema struct {
	// Total number of rules.
	RuleCount *int64 `json:"rule_count,omitempty"`

	// rules.
	Rules []Rule `json:"rules,omitempty"`

	// data policy status.
	Status *string `json:"status,omitempty"`

	// catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// data artifact.
	DataArtifact *string `json:"data_artifact,omitempty"`

	Metadata *DataPolicyMetadata `json:"metadata,omitempty"`

	// the displayed name for the policy.
	PolicyName *string `json:"policy_name,omitempty"`
}

// Constants associated with the PolicySchema.Status property.
// data policy status.
const (
	PolicySchema_Status_Active   = "active"
	PolicySchema_Status_Inactive = "inactive"
)

// UnmarshalPolicySchema unmarshals an instance of PolicySchema from the specified map of raw messages.
func UnmarshalPolicySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicySchema)
	err = core.UnmarshalPrimitive(m, "rule_count", &obj.RuleCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_artifact", &obj.DataArtifact)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalDataPolicyMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_name", &obj.PolicyName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicySchemaList : AMS schema List.
type PolicySchemaList struct {
	// catalog policies definition.
	CatalogPolicies []GetCatalogUsersSchema `json:"catalog_policies,omitempty"`

	// data policies definition.
	DataPolicies []PolicySchema `json:"data_policies,omitempty"`

	// engine policies definition.
	EnginePolicies []GetEngineUsersSchema `json:"engine_policies,omitempty"`
}

// UnmarshalPolicySchemaList unmarshals an instance of PolicySchemaList from the specified map of raw messages.
func UnmarshalPolicySchemaList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicySchemaList)
	err = core.UnmarshalModel(m, "catalog_policies", &obj.CatalogPolicies, UnmarshalGetCatalogUsersSchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_policies", &obj.DataPolicies, UnmarshalPolicySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_policies", &obj.EnginePolicies, UnmarshalGetEngineUsersSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyVersionResultSchema : AMS policy version result.
type PolicyVersionResultSchema struct {
	// The catalog policy version list.
	CatalogPolicies []CatalogPolicies `json:"catalog_policies,omitempty"`

	// The data policy version list.
	DataPolicies []DataPolicies `json:"data_policies,omitempty"`

	// The Db connection policy version list.
	DatabasePolicies []DbConnPolicies `json:"database_policies,omitempty"`

	// The engine policy version list.
	EnginePolicies []EnginePolicies `json:"engine_policies,omitempty"`

	// The bucket policy version list.
	BucketPolicies []BucketPolicies `json:"bucket_policies,omitempty"`
}

// UnmarshalPolicyVersionResultSchema unmarshals an instance of PolicyVersionResultSchema from the specified map of raw messages.
func UnmarshalPolicyVersionResultSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyVersionResultSchema)
	err = core.UnmarshalModel(m, "catalog_policies", &obj.CatalogPolicies, UnmarshalCatalogPolicies)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_policies", &obj.DataPolicies, UnmarshalDataPolicies)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "database_policies", &obj.DatabasePolicies, UnmarshalDbConnPolicies)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "engine_policies", &obj.EnginePolicies, UnmarshalEnginePolicies)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "bucket_policies", &obj.BucketPolicies, UnmarshalBucketPolicies)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Query : Query.
type Query struct {
	// Created by.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Created on.
	CreatedOn *string `json:"created_on" validate:"required"`

	// Description.
	Description *string `json:"description" validate:"required"`

	// Engine ID.
	EngineID *string `json:"engine_id" validate:"required"`

	// Query name.
	QueryName *string `json:"query_name" validate:"required"`

	// Query string.
	QueryString *string `json:"query_string" validate:"required"`
}

// UnmarshalQuery unmarshals an instance of Query from the specified map of raw messages.
func UnmarshalQuery(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Query)
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
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_name", &obj.QueryName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_string", &obj.QueryString)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RegisterBucketCreatedBody : RegisterBucketCreatedBody struct
type RegisterBucketCreatedBody struct {
	Bucket *RegisterBucketCreatedBodyBucket `json:"bucket" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalRegisterBucketCreatedBody unmarshals an instance of RegisterBucketCreatedBody from the specified map of raw messages.
func UnmarshalRegisterBucketCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegisterBucketCreatedBody)
	err = core.UnmarshalModel(m, "bucket", &obj.Bucket, UnmarshalRegisterBucketCreatedBodyBucket)
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

// RegisterBucketCreatedBodyBucket : RegisterBucketCreatedBodyBucket struct
type RegisterBucketCreatedBodyBucket struct {
	// Bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// Bucket ID.
	BucketID *string `json:"bucket_id,omitempty"`
}

// UnmarshalRegisterBucketCreatedBodyBucket unmarshals an instance of RegisterBucketCreatedBodyBucket from the specified map of raw messages.
func UnmarshalRegisterBucketCreatedBodyBucket(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegisterBucketCreatedBodyBucket)
	err = core.UnmarshalPrimitive(m, "bucket_display_name", &obj.BucketDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_id", &obj.BucketID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RegisterDatabaseCatalogBodyDatabaseDetails : database details.
type RegisterDatabaseCatalogBodyDatabaseDetails struct {
	// Psssword.
	Password *string `json:"password,omitempty"`

	// Port.
	Port *string `json:"port,omitempty"`

	// SSL Mode.
	Ssl *bool `json:"ssl,omitempty"`

	// Only for Kafka - Add kafka tables.
	Tables *string `json:"tables,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`

	// Database name.
	DatabaseName *string `json:"database_name,omitempty"`

	// Host name.
	Hostname *string `json:"hostname,omitempty"`
}

// UnmarshalRegisterDatabaseCatalogBodyDatabaseDetails unmarshals an instance of RegisterDatabaseCatalogBodyDatabaseDetails from the specified map of raw messages.
func UnmarshalRegisterDatabaseCatalogBodyDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegisterDatabaseCatalogBodyDatabaseDetails)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
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
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceDataPolicyCreatedBody : Replace data policy success.
type ReplaceDataPolicyCreatedBody struct {
	// Replace data policy.
	DataPolicy *ReplaceDataPolicySchema `json:"data_policy" validate:"required"`

	Metadata *DataPolicyMetadata `json:"metadata" validate:"required"`

	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`
}

// UnmarshalReplaceDataPolicyCreatedBody unmarshals an instance of ReplaceDataPolicyCreatedBody from the specified map of raw messages.
func UnmarshalReplaceDataPolicyCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplaceDataPolicyCreatedBody)
	err = core.UnmarshalModel(m, "data_policy", &obj.DataPolicy, UnmarshalReplaceDataPolicySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalDataPolicyMetadata)
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

// ReplaceDataPolicySchema : Replace data policy.
type ReplaceDataPolicySchema struct {
	// catalog name.
	CatalogName *string `json:"catalog_name" validate:"required"`

	// data artifact.
	DataArtifact *string `json:"data_artifact" validate:"required"`

	// a more detailed description of the policy.
	Description *string `json:"description,omitempty"`

	// rules.
	Rules []Rule `json:"rules" validate:"required"`

	// data policy status.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the ReplaceDataPolicySchema.Status property.
// data policy status.
const (
	ReplaceDataPolicySchema_Status_Active   = "active"
	ReplaceDataPolicySchema_Status_Inactive = "inactive"
)

// NewReplaceDataPolicySchema : Instantiate ReplaceDataPolicySchema (Generic Model Constructor)
func (*WatsonxDataV1) NewReplaceDataPolicySchema(catalogName string, dataArtifact string, rules []Rule) (_model *ReplaceDataPolicySchema, err error) {
	_model = &ReplaceDataPolicySchema{
		CatalogName:  core.StringPtr(catalogName),
		DataArtifact: core.StringPtr(dataArtifact),
		Rules:        rules,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalReplaceDataPolicySchema unmarshals an instance of ReplaceDataPolicySchema from the specified map of raw messages.
func UnmarshalReplaceDataPolicySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplaceDataPolicySchema)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_artifact", &obj.DataArtifact)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceWithResult : Resource with result.
type ResourceWithResult struct {
	// action.
	Action *string `json:"action" validate:"required"`

	// Resource name.
	ResourceName *string `json:"resource_name" validate:"required"`

	// Resource type.
	ResourceType *string `json:"resource_type" validate:"required"`

	// resource evaluation result.
	Result *bool `json:"result" validate:"required"`
}

// UnmarshalResourceWithResult unmarshals an instance of ResourceWithResult from the specified map of raw messages.
func UnmarshalResourceWithResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceWithResult)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
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

// ResourcesMetadata : Resource.
type ResourcesMetadata struct {
	// resource action to be evaluated.
	Action *string `json:"action" validate:"required"`

	// Resource name.
	ResourceName *string `json:"resource_name" validate:"required"`

	// Resource type.
	ResourceType *string `json:"resource_type" validate:"required"`
}

// Constants associated with the ResourcesMetadata.ResourceType property.
// Resource type.
const (
	ResourcesMetadata_ResourceType_Bucket   = "bucket"
	ResourcesMetadata_ResourceType_Catalog  = "catalog"
	ResourcesMetadata_ResourceType_Database = "database"
	ResourcesMetadata_ResourceType_Engine   = "engine"
)

// NewResourcesMetadata : Instantiate ResourcesMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewResourcesMetadata(action string, resourceName string, resourceType string) (_model *ResourcesMetadata, err error) {
	_model = &ResourcesMetadata{
		Action:       core.StringPtr(action),
		ResourceName: core.StringPtr(resourceName),
		ResourceType: core.StringPtr(resourceType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResourcesMetadata unmarshals an instance of ResourcesMetadata from the specified map of raw messages.
func UnmarshalResourcesMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourcesMetadata)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResumeEngineCreatedBody : resumeEngine OK.
type ResumeEngineCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalResumeEngineCreatedBody unmarshals an instance of ResumeEngineCreatedBody from the specified map of raw messages.
func UnmarshalResumeEngineCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResumeEngineCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : Rule struct
type Rule struct {
	// the actions to enforce when the data policy triggers.
	Actions []string `json:"actions" validate:"required"`

	// data policy effect.
	Effect *string `json:"effect,omitempty"`

	// user name, group id or tag value.
	Grantee *RuleGrantee `json:"grantee" validate:"required"`
}

// Constants associated with the Rule.Actions property.
const (
	Rule_Actions_All      = "all"
	Rule_Actions_Alter    = "alter"
	Rule_Actions_Create   = "create"
	Rule_Actions_Delete   = "delete"
	Rule_Actions_Drop     = "drop"
	Rule_Actions_Grant    = "grant"
	Rule_Actions_Insert   = "insert"
	Rule_Actions_Revoke   = "revoke"
	Rule_Actions_Select   = "select"
	Rule_Actions_Show     = "show"
	Rule_Actions_Truncate = "truncate"
	Rule_Actions_Use      = "use"
	Rule_Actions_View     = "view"
)

// Constants associated with the Rule.Effect property.
// data policy effect.
const (
	Rule_Effect_Allow = "allow"
	Rule_Effect_Deny  = "deny"
)

// NewRule : Instantiate Rule (Generic Model Constructor)
func (*WatsonxDataV1) NewRule(actions []string, grantee *RuleGrantee) (_model *Rule, err error) {
	_model = &Rule{
		Actions: actions,
		Grantee: grantee,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effect", &obj.Effect)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "grantee", &obj.Grantee, UnmarshalRuleGrantee)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleGrantee : user name, group id or tag value.
type RuleGrantee struct {
	// grantee value.
	Value *string `json:"value" validate:"required"`

	// grantee key.
	Key *string `json:"key" validate:"required"`

	// grantee type.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the RuleGrantee.Key property.
// grantee key.
const (
	RuleGrantee_Key_AttributeName = "attribute_name"
	RuleGrantee_Key_GroupID       = "group_id"
	RuleGrantee_Key_UserName      = "user_name"
)

// Constants associated with the RuleGrantee.Type property.
// grantee type.
const (
	RuleGrantee_Type_Tag          = "tag"
	RuleGrantee_Type_UserIdentity = "user_identity"
)

// NewRuleGrantee : Instantiate RuleGrantee (Generic Model Constructor)
func (*WatsonxDataV1) NewRuleGrantee(value string, key string, typeVar string) (_model *RuleGrantee, err error) {
	_model = &RuleGrantee{
		Value: core.StringPtr(value),
		Key:   core.StringPtr(key),
		Type:  core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuleGrantee unmarshals an instance of RuleGrantee from the specified map of raw messages.
func UnmarshalRuleGrantee(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleGrantee)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
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

// SuccessResponse : Response of success.
type SuccessResponse struct {
	// Message code.
	MessageCode *string `json:"_messageCode_,omitempty"`

	// Message.
	Message *string `json:"message,omitempty"`
}

// UnmarshalSuccessResponse unmarshals an instance of SuccessResponse from the specified map of raw messages.
func UnmarshalSuccessResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessResponse)
	err = core.UnmarshalPrimitive(m, "_messageCode_", &obj.MessageCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableSnapshot : TableSnapshot.
type TableSnapshot struct {
	// Operation.
	Operation *string `json:"operation,omitempty"`

	// Snapshot id.
	SnapshotID *string `json:"snapshot_id,omitempty"`

	// Summary.
	Summary map[string]interface{} `json:"summary,omitempty"`

	// Committed at.
	CommittedAt *string `json:"committed_at,omitempty"`
}

// UnmarshalTableSnapshot unmarshals an instance of TableSnapshot from the specified map of raw messages.
func UnmarshalTableSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableSnapshot)
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
	err = core.UnmarshalPrimitive(m, "committed_at", &obj.CommittedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateDatabaseBodyDatabaseDetails : database details.
type UpdateDatabaseBodyDatabaseDetails struct {
	// Password.
	Password *string `json:"password,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`
}

// UnmarshalUpdateDatabaseBodyDatabaseDetails unmarshals an instance of UpdateDatabaseBodyDatabaseDetails from the specified map of raw messages.
func UnmarshalUpdateDatabaseBodyDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateDatabaseBodyDatabaseDetails)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateTableBodyAddColumnsItems : Add_columns items.
type UpdateTableBodyAddColumnsItems struct {
	// Comment.
	ColumnComment *string `json:"column_comment,omitempty"`

	// Column name.
	ColumnName *string `json:"column_name,omitempty"`

	// Data type.
	DataType *string `json:"data_type,omitempty"`
}

// UnmarshalUpdateTableBodyAddColumnsItems unmarshals an instance of UpdateTableBodyAddColumnsItems from the specified map of raw messages.
func UnmarshalUpdateTableBodyAddColumnsItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateTableBodyAddColumnsItems)
	err = core.UnmarshalPrimitive(m, "column_comment", &obj.ColumnComment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_type", &obj.DataType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateTableBodyDropColumnsItems : Drop_columns items.
type UpdateTableBodyDropColumnsItems struct {
	// Column name.
	ColumnName *string `json:"column_name,omitempty"`
}

// UnmarshalUpdateTableBodyDropColumnsItems unmarshals an instance of UpdateTableBodyDropColumnsItems from the specified map of raw messages.
func UnmarshalUpdateTableBodyDropColumnsItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateTableBodyDropColumnsItems)
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateTableBodyRenameColumnsItems : Rename_columns items.
type UpdateTableBodyRenameColumnsItems struct {
	// Column name.
	ColumnName *string `json:"column_name,omitempty"`

	// New column name.
	NewColumnName *string `json:"new_column_name,omitempty"`
}

// UnmarshalUpdateTableBodyRenameColumnsItems unmarshals an instance of UpdateTableBodyRenameColumnsItems from the specified map of raw messages.
func UnmarshalUpdateTableBodyRenameColumnsItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateTableBodyRenameColumnsItems)
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "new_column_name", &obj.NewColumnName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UsersMetadata : Users metadata.
type UsersMetadata struct {
	// Eligible permission to the resource.
	Permission *string `json:"permission" validate:"required"`

	// The user name.
	UserName *string `json:"user_name" validate:"required"`
}

// Constants associated with the UsersMetadata.Permission property.
// Eligible permission to the resource.
const (
	UsersMetadata_Permission_CanAdminister = "can_administer"
	UsersMetadata_Permission_CanManage     = "can_manage"
	UsersMetadata_Permission_CanUse        = "can_use"
)

// NewUsersMetadata : Instantiate UsersMetadata (Generic Model Constructor)
func (*WatsonxDataV1) NewUsersMetadata(permission string, userName string) (_model *UsersMetadata, err error) {
	_model = &UsersMetadata{
		Permission: core.StringPtr(permission),
		UserName:   core.StringPtr(userName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUsersMetadata unmarshals an instance of UsersMetadata from the specified map of raw messages.
func UnmarshalUsersMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UsersMetadata)
	err = core.UnmarshalPrimitive(m, "permission", &obj.Permission)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_name", &obj.UserName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
