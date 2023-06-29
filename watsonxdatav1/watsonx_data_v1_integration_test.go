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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the watsonxdatav1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`WatsonxDataV1 Integration Tests`, func() {
	const externalConfigFile = "../watsonx_data_v1.env"

	var (
		err                error
		watsonxDataService *watsonxdatav1.WatsonxDataV1
		serviceURL         string
		config             map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(watsonxdatav1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			bearerToken := "Bearer ****"
			authenticator, err := core.NewBearerTokenAuthenticator(bearerToken)
			if err != nil {
				panic(err)
			}

			watsonxDataServiceOptions := &watsonxdatav1.WatsonxDataV1Options{
				Authenticator: authenticator,
			}

			watsonxDataService, err = watsonxdatav1.NewWatsonxDataV1UsingExternalConfig(watsonxDataServiceOptions)
			Expect(err).To(BeNil())
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(watsonxDataService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			watsonxDataService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateDbConnUsers - Grant users and groups permission to the db connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDbConnUsers(createDbConnUsersOptions *CreateDbConnUsersOptions)`, func() {
			bucketDbConnGroupsMetadataModel := &watsonxdatav1.BucketDbConnGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			bucketDbConnUsersMetadataModel := &watsonxdatav1.BucketDbConnUsersMetadata{
				UserName:   core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			createDbConnUsersOptions := &watsonxdatav1.CreateDbConnUsersOptions{
				DatabaseID:     core.StringPtr("testString"),
				Groups:         []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel},
				Users:          []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateDbConnUsers(createDbConnUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListDataPolicies - Get policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataPolicies(listDataPoliciesOptions *ListDataPoliciesOptions)`, func() {
			listDataPoliciesOptions := &watsonxdatav1.ListDataPoliciesOptions{
				LhInstanceID:    core.StringPtr("sampleInstanceID"),
				AuthInstanceID:  core.StringPtr("sampleAuthInstanceID"),
				CatalogName:     core.StringPtr("sampleCatalogName"),
				Status:          core.StringPtr("testString"),
				IncludeMetadata: core.BoolPtr(true),
				IncludeRules:    core.BoolPtr(true),
			}

			policyListSchema, response, err := watsonxDataService.ListDataPolicies(listDataPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyListSchema).ToNot(BeNil())
		})
	})

	Describe(`CreateDataPolicy - Create new data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataPolicy(createDataPolicyOptions *CreateDataPolicyOptions)`, func() {
			ruleGranteeModel := &watsonxdatav1.RuleGrantee{
				Value: core.StringPtr("testString"),
				Key:   core.StringPtr("user_name"),
				Type:  core.StringPtr("user_identity"),
			}

			ruleModel := &watsonxdatav1.Rule{
				Actions: []string{"all"},
				Effect:  core.StringPtr("allow"),
				Grantee: ruleGranteeModel,
			}

			createDataPolicyOptions := &watsonxdatav1.CreateDataPolicyOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				DataArtifact:   core.StringPtr("schema1/table1/(column1|column2)"),
				PolicyName:     core.StringPtr("testString"),
				Rules:          []watsonxdatav1.Rule{*ruleModel},
				Description:    core.StringPtr("testString"),
				Status:         core.StringPtr("active"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			createDataPolicyCreatedBody, response, err := watsonxDataService.CreateDataPolicy(createDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createDataPolicyCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`GetEngineUsers - Get permission in the engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEngineUsers(getEngineUsersOptions *GetEngineUsersOptions)`, func() {
			getEngineUsersOptions := &watsonxdatav1.GetEngineUsersOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getEngineUsersSchema, response, err := watsonxDataService.GetEngineUsers(getEngineUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getEngineUsersSchema).ToNot(BeNil())
		})
	})

	Describe(`UpdateEngineUsers - Updates user and groups permission in the engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEngineUsers(updateEngineUsersOptions *UpdateEngineUsersOptions)`, func() {
			engineGroupsMetadataModel := &watsonxdatav1.EngineGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			engineUsersMetadataModel := &watsonxdatav1.EngineUsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			updateEngineUsersOptions := &watsonxdatav1.UpdateEngineUsersOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				Groups:         []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel},
				Users:          []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateEngineUsers(updateEngineUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateDbConnUsers - Updates user and groups permission in the db connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDbConnUsers(updateDbConnUsersOptions *UpdateDbConnUsersOptions)`, func() {
			bucketDbConnGroupsMetadataModel := &watsonxdatav1.BucketDbConnGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			bucketDbConnUsersMetadataModel := &watsonxdatav1.BucketDbConnUsersMetadata{
				UserName:   core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			updateDbConnUsersOptions := &watsonxdatav1.UpdateDbConnUsersOptions{
				DatabaseID:     core.StringPtr("testString"),
				Groups:         []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel},
				Users:          []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateDbConnUsers(updateDbConnUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDbConnUsers - Get permission in the db connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDbConnUsers(getDbConnUsersOptions *GetDbConnUsersOptions)`, func() {
			getDbConnUsersOptions := &watsonxdatav1.GetDbConnUsersOptions{
				DatabaseID:     core.StringPtr("testString"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getDbConnUsersSchema, response, err := watsonxDataService.GetDbConnUsers(getDbConnUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDbConnUsersSchema).ToNot(BeNil())
		})
	})

	Describe(`CreateCatalogUsers - Grant users and groups permission to the catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogUsers(createCatalogUsersOptions *CreateCatalogUsersOptions)`, func() {
			catalogGroupsMetadataModel := &watsonxdatav1.CatalogGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			catalogUsersMetadataModel := &watsonxdatav1.CatalogUsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			createCatalogUsersOptions := &watsonxdatav1.CreateCatalogUsersOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				Groups:         []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel},
				Users:          []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateCatalogUsers(createCatalogUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetCatalogUsers - Get users and groups permission in the catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogUsers(getCatalogUsersOptions *GetCatalogUsersOptions)`, func() {
			getCatalogUsersOptions := &watsonxdatav1.GetCatalogUsersOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getCatalogUsersSchema, response, err := watsonxDataService.GetCatalogUsers(getCatalogUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getCatalogUsersSchema).ToNot(BeNil())
		})
	})

	Describe(`UpdateCatalogUsers - Updates user and groups permission in the catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalogUsers(updateCatalogUsersOptions *UpdateCatalogUsersOptions)`, func() {
			catalogGroupsMetadataModel := &watsonxdatav1.CatalogGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			catalogUsersMetadataModel := &watsonxdatav1.CatalogUsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			updateCatalogUsersOptions := &watsonxdatav1.UpdateCatalogUsersOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				Groups:         []watsonxdatav1.CatalogGroupsMetadata{*catalogGroupsMetadataModel},
				Users:          []watsonxdatav1.CatalogUsersMetadata{*catalogUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateCatalogUsers(updateCatalogUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`Evaluate - Evaluate permission`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Evaluate(evaluateOptions *EvaluateOptions)`, func() {
			resourcesMetadataModel := &watsonxdatav1.ResourcesMetadata{
				Action:       core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				ResourceType: core.StringPtr("engine"),
			}

			evaluateOptions := &watsonxdatav1.EvaluateOptions{
				Resources:      []watsonxdatav1.ResourcesMetadata{*resourcesMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			evaluationResultSchema, response, err := watsonxDataService.Evaluate(evaluateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(evaluationResultSchema).ToNot(BeNil())
		})
	})

	Describe(`GetPoliciesList - Get policies for specific catalog in catalog_name list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPoliciesList(getPoliciesListOptions *GetPoliciesListOptions)`, func() {
			getPoliciesListOptions := &watsonxdatav1.GetPoliciesListOptions{
				LhInstanceID:        core.StringPtr("sampleInstanceID"),
				AuthInstanceID:      core.StringPtr("testString"),
				CatalogList:         []string{"testString"},
				EngineList:          []string{"testString"},
				DataPoliciesList:    []string{"testString"},
				IncludeDataPolicies: core.BoolPtr(true),
			}

			policySchemaList, response, err := watsonxDataService.GetPoliciesList(getPoliciesListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policySchemaList).ToNot(BeNil())
		})
	})

	Describe(`CreateMetastoreUsers - Grant users and groups permission to the metastore`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMetastoreUsers(createMetastoreUsersOptions *CreateMetastoreUsersOptions)`, func() {
			groupsMetadataModel := &watsonxdatav1.GroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			usersMetadataModel := &watsonxdatav1.UsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			createMetastoreUsersOptions := &watsonxdatav1.CreateMetastoreUsersOptions{
				MetastoreName:  core.StringPtr("testString"),
				Groups:         []watsonxdatav1.GroupsMetadata{*groupsMetadataModel},
				Users:          []watsonxdatav1.UsersMetadata{*usersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateMetastoreUsers(createMetastoreUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetMetastoreUsers - Get permission in the metastore`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMetastoreUsers(getMetastoreUsersOptions *GetMetastoreUsersOptions)`, func() {
			getMetastoreUsersOptions := &watsonxdatav1.GetMetastoreUsersOptions{
				MetastoreName:  core.StringPtr("testString"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getMetastoreUsersSchema, response, err := watsonxDataService.GetMetastoreUsers(getMetastoreUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getMetastoreUsersSchema).ToNot(BeNil())
		})
	})

	Describe(`UpdateMetastoreUsers - Updates user and groups permission in the metastore`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateMetastoreUsers(updateMetastoreUsersOptions *UpdateMetastoreUsersOptions)`, func() {
			groupsMetadataModel := &watsonxdatav1.GroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			usersMetadataModel := &watsonxdatav1.UsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			updateMetastoreUsersOptions := &watsonxdatav1.UpdateMetastoreUsersOptions{
				MetastoreName:  core.StringPtr("testString"),
				Groups:         []watsonxdatav1.GroupsMetadata{*groupsMetadataModel},
				Users:          []watsonxdatav1.UsersMetadata{*usersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateMetastoreUsers(updateMetastoreUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateBucketUsers - Grant users and groups permission to the bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBucketUsers(createBucketUsersOptions *CreateBucketUsersOptions)`, func() {
			bucketDbConnGroupsMetadataModel := &watsonxdatav1.BucketDbConnGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			bucketDbConnUsersMetadataModel := &watsonxdatav1.BucketDbConnUsersMetadata{
				UserName:   core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			createBucketUsersOptions := &watsonxdatav1.CreateBucketUsersOptions{
				BucketID:       core.StringPtr("testString"),
				Groups:         []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel},
				Users:          []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateBucketUsers(createBucketUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDefaultPolicies - Get AMS default policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDefaultPolicies(getDefaultPoliciesOptions *GetDefaultPoliciesOptions)`, func() {
			getDefaultPoliciesOptions := &watsonxdatav1.GetDefaultPoliciesOptions{
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			defaultPolicySchema, response, err := watsonxDataService.GetDefaultPolicies(getDefaultPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultPolicySchema).ToNot(BeNil())
		})
	})

	Describe(`GetPolicyVersion - Get AMS policies version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyVersion(getPolicyVersionOptions *GetPolicyVersionOptions)`, func() {
			getPolicyVersionOptions := &watsonxdatav1.GetPolicyVersionOptions{
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			policyVersionResultSchema, response, err := watsonxDataService.GetPolicyVersion(getPolicyVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyVersionResultSchema).ToNot(BeNil())
		})
	})

	Describe(`GetDataPolicy - Get policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataPolicy(getDataPolicyOptions *GetDataPolicyOptions)`, func() {
			getDataPolicyOptions := &watsonxdatav1.GetDataPolicyOptions{
				PolicyName:     core.StringPtr("testString"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			policySchema, response, err := watsonxDataService.GetDataPolicy(getDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policySchema).ToNot(BeNil())
		})
	})

	Describe(`ReplaceDataPolicy - Updates data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceDataPolicy(replaceDataPolicyOptions *ReplaceDataPolicyOptions)`, func() {
			ruleGranteeModel := &watsonxdatav1.RuleGrantee{
				Value: core.StringPtr("testString"),
				Key:   core.StringPtr("user_name"),
				Type:  core.StringPtr("user_identity"),
			}

			ruleModel := &watsonxdatav1.Rule{
				Actions: []string{"all"},
				Effect:  core.StringPtr("allow"),
				Grantee: ruleGranteeModel,
			}

			replaceDataPolicyOptions := &watsonxdatav1.ReplaceDataPolicyOptions{
				PolicyName:     core.StringPtr("testString"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				DataArtifact:   core.StringPtr("schema1/table1/(column1|column2)"),
				Rules:          []watsonxdatav1.Rule{*ruleModel},
				Description:    core.StringPtr("testString"),
				Status:         core.StringPtr("active"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			replaceDataPolicyCreatedBody, response, err := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(replaceDataPolicyCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`CreateEngineUsers - Grant permission to the engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEngineUsers(createEngineUsersOptions *CreateEngineUsersOptions)`, func() {
			engineGroupsMetadataModel := &watsonxdatav1.EngineGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			engineUsersMetadataModel := &watsonxdatav1.EngineUsersMetadata{
				Permission: core.StringPtr("can_administer"),
				UserName:   core.StringPtr("testString"),
			}

			createEngineUsersOptions := &watsonxdatav1.CreateEngineUsersOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				Groups:         []watsonxdatav1.EngineGroupsMetadata{*engineGroupsMetadataModel},
				Users:          []watsonxdatav1.EngineUsersMetadata{*engineUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateEngineUsers(createEngineUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetBucketUsers - Get permission in the bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBucketUsers(getBucketUsersOptions *GetBucketUsersOptions)`, func() {
			getBucketUsersOptions := &watsonxdatav1.GetBucketUsersOptions{
				BucketID:       core.StringPtr("testString"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getBucketUsersSchema, response, err := watsonxDataService.GetBucketUsers(getBucketUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBucketUsersSchema).ToNot(BeNil())
		})
	})

	Describe(`UpdateBucketUsers - Updates user and groups permission in the bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBucketUsers(updateBucketUsersOptions *UpdateBucketUsersOptions)`, func() {
			bucketDbConnGroupsMetadataModel := &watsonxdatav1.BucketDbConnGroupsMetadata{
				GroupID:    core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			bucketDbConnUsersMetadataModel := &watsonxdatav1.BucketDbConnUsersMetadata{
				UserName:   core.StringPtr("testString"),
				Permission: core.StringPtr("can_administer"),
			}

			updateBucketUsersOptions := &watsonxdatav1.UpdateBucketUsersOptions{
				BucketID:       core.StringPtr("testString"),
				Groups:         []watsonxdatav1.BucketDbConnGroupsMetadata{*bucketDbConnGroupsMetadataModel},
				Users:          []watsonxdatav1.BucketDbConnUsersMetadata{*bucketDbConnUsersMetadataModel},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateBucketUsers(updateBucketUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`RegisterBucket - Register bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RegisterBucket(registerBucketOptions *RegisterBucketOptions)`, func() {
			bucketDetailsModel := &watsonxdatav1.BucketDetails{
				AccessKey:  core.StringPtr("<access_key>"),
				BucketName: core.StringPtr("sample-bucket1"),
				Endpoint:   core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/"),
				SecretKey:  core.StringPtr("<secret_key>"),
			}

			registerBucketOptions := &watsonxdatav1.RegisterBucketOptions{
				BucketDetails:     bucketDetailsModel,
				Description:       core.StringPtr("COS bucket for customer data"),
				TableType:         core.StringPtr("iceberg"),
				BucketType:        core.StringPtr("ibm_cos"),
				CatalogName:       core.StringPtr("samplecatalog1"),
				ManagedBy:         core.StringPtr("customer"),
				BucketDisplayName: core.StringPtr("sample-bucket-displayname"),
				BucketTags:        []string{"read customer data", "write customer data'"},
				CatalogTags:       []string{"catalog_tag_1", "catalog_tag_2"},
				ThriftURI:         core.StringPtr("thrift://samplehost-metastore:4354"),
				AuthInstanceID:    core.StringPtr("sampleAuthInstanceID"),
			}

			registerBucketCreatedBody, response, err := watsonxDataService.RegisterBucket(registerBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(registerBucketCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`GetBuckets - Get buckets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBuckets(getBucketsOptions *GetBucketsOptions)`, func() {
			getBucketsOptions := &watsonxdatav1.GetBucketsOptions{
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getBucketsOkBody, response, err := watsonxDataService.GetBuckets(getBucketsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBucketsOkBody).ToNot(BeNil())
		})
	})

	Describe(`GetBucketObjects - Get bucket objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBucketObjects(getBucketObjectsOptions *GetBucketObjectsOptions)`, func() {
			getBucketObjectsOptions := &watsonxdatav1.GetBucketObjectsOptions{
				BucketID:       core.StringPtr("sample-bucket1"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getBucketObjectsOkBody, response, err := watsonxDataService.GetBucketObjects(getBucketObjectsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBucketObjectsOkBody).ToNot(BeNil())
		})
	})

	Describe(`UpdateBucket - Update bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBucket(updateBucketOptions *UpdateBucketOptions)`, func() {
			updateBucketOptions := &watsonxdatav1.UpdateBucketOptions{
				BucketID:          core.StringPtr("sample-bucket1"),
				AccessKey:         core.StringPtr("<access_key>"),
				BucketDisplayName: core.StringPtr("sample-bucket1-displayname"),
				Description:       core.StringPtr("COS bucket for customer data"),
				SecretKey:         core.StringPtr("<secret_key>"),
				Tags:              []string{"testbucket", "userbucket"},
				AuthInstanceID:    core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateBucket(updateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`DeactivateBucket - Deactivate bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeactivateBucket(deactivateBucketOptions *DeactivateBucketOptions)`, func() {
			deactivateBucketOptions := &watsonxdatav1.DeactivateBucketOptions{
				BucketID:       core.StringPtr("sample-bucket1"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.DeactivateBucket(deactivateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ActivateBucket - Active bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ActivateBucket(activateBucketOptions *ActivateBucketOptions)`, func() {
			activateBucketOptions := &watsonxdatav1.ActivateBucketOptions{
				BucketID:       core.StringPtr("sample-bucket1"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.ActivateBucket(activateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDatabases - Get databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabases(getDatabasesOptions *GetDatabasesOptions)`, func() {
			getDatabasesOptions := &watsonxdatav1.GetDatabasesOptions{
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.GetDatabases(getDatabasesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CreateDatabaseCatalog - Add/Create database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDatabaseCatalog(createDatabaseCatalogOptions *CreateDatabaseCatalogOptions)`, func() {
			registerDatabaseCatalogBodyDatabaseDetailsModel := &watsonxdatav1.RegisterDatabaseCatalogBodyDatabaseDetails{
				Password:     core.StringPtr("samplepassword"),
				Port:         core.StringPtr("4553"),
				Ssl:          core.BoolPtr(true),
				Tables:       core.StringPtr("kafka_table_name"),
				Username:     core.StringPtr("sampleuser"),
				DatabaseName: core.StringPtr("new_database"),
				Hostname:     core.StringPtr("db2@<hostname>.com"),
			}

			createDatabaseCatalogOptions := &watsonxdatav1.CreateDatabaseCatalogOptions{
				DatabaseDisplayName: core.StringPtr("new_database"),
				DatabaseType:        core.StringPtr("db2"),
				CatalogName:         core.StringPtr("sampledbcatalog"),
				DatabaseDetails:     registerDatabaseCatalogBodyDatabaseDetailsModel,
				Description:         core.StringPtr("db2 extenal database description"),
				Tags:                []string{"tag_1", "tag_2"},
				CreatedBy:           core.StringPtr("<username>@<domain>.com"),
				Accept:              core.StringPtr("application/json"),
				AuthInstanceID:      core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.CreateDatabaseCatalog(createDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`UpdateDatabase - Update database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions)`, func() {
			updateDatabaseBodyDatabaseDetailsModel := &watsonxdatav1.UpdateDatabaseBodyDatabaseDetails{
				Password: core.StringPtr("samplepassword"),
				Username: core.StringPtr("sampleuser"),
			}

			updateDatabaseOptions := &watsonxdatav1.UpdateDatabaseOptions{
				DatabaseID:          core.StringPtr("new_db_id"),
				DatabaseDetails:     updateDatabaseBodyDatabaseDetailsModel,
				DatabaseDisplayName: core.StringPtr("new_database"),
				Description:         core.StringPtr("External database description"),
				Tags:                []string{"testdatabase", "userdatabase"},
				Accept:              core.StringPtr("application/json"),
				AuthInstanceID:      core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.UpdateDatabase(updateDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDeployments - Get instance details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDeployments(getDeploymentsOptions *GetDeploymentsOptions)`, func() {
			getDeploymentsOptions := &watsonxdatav1.GetDeploymentsOptions{
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.GetDeployments(getDeploymentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CreateEngine - Create engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEngine(createEngineOptions *CreateEngineOptions)`, func() {
			nodeDescriptionBodyModel := &watsonxdatav1.NodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			engineDetailsBodyModel := &watsonxdatav1.EngineDetailsBody{
				Worker:      nodeDescriptionBodyModel,
				Coordinator: nodeDescriptionBodyModel,
				SizeConfig:  core.StringPtr("starter"),
			}

			createEngineOptions := &watsonxdatav1.CreateEngineOptions{
				Version:            core.StringPtr("1.2.3"),
				EngineDetails:      engineDetailsBodyModel,
				Origin:             core.StringPtr("ibm"),
				Type:               core.StringPtr("presto"),
				Description:        core.StringPtr("presto engine description"),
				EngineDisplayName:  core.StringPtr("sampleEngine"),
				FirstTimeUse:       core.BoolPtr(false),
				Region:             core.StringPtr("us-south"),
				AssociatedCatalogs: []string{"new_catalog_1", "new_catalog_2"},
				Accept:             core.StringPtr("application/json"),
				AuthInstanceID:     core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.CreateEngine(createEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetEngines - Get engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEngines(getEnginesOptions *GetEnginesOptions)`, func() {
			getEnginesOptions := &watsonxdatav1.GetEnginesOptions{
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getEnginesOkBody, response, err := watsonxDataService.GetEngines(getEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getEnginesOkBody).ToNot(BeNil())
		})
	})

	Describe(`UpdateEngine - Update engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEngine(updateEngineOptions *UpdateEngineOptions)`, func() {
			nodeDescriptionModel := &watsonxdatav1.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			updateEngineOptions := &watsonxdatav1.UpdateEngineOptions{
				EngineID:          core.StringPtr("sampleEngineID"),
				Coordinator:       nodeDescriptionModel,
				Description:       core.StringPtr("presto engine updated description"),
				EngineDisplayName: core.StringPtr("starter"),
				Tags:              []string{"tag1", "tag2"},
				Worker:            nodeDescriptionModel,
				Accept:            core.StringPtr("application/json"),
				AuthInstanceID:    core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.UpdateEngine(updateEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`TestLHConsole - Readiness API`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions)`, func() {
			testLhConsoleOptions := &watsonxdatav1.TestLHConsoleOptions{}

			successResponse, response, err := watsonxDataService.TestLHConsole(testLhConsoleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetMetastores - Get Catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMetastores(getMetastoresOptions *GetMetastoresOptions)`, func() {
			getMetastoresOptions := &watsonxdatav1.GetMetastoresOptions{
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getMetastoresOkBody, response, err := watsonxDataService.GetMetastores(getMetastoresOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getMetastoresOkBody).ToNot(BeNil())
		})
	})

	Describe(`GetHMS - Get Metastore`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetHMS(getHMSOptions *GetHMSOptions)`, func() {
			getHmsOptions := &watsonxdatav1.GetHMSOptions{
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.GetHMS(getHmsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CreateSchema - Create schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
			createSchemaOptions := &watsonxdatav1.CreateSchemaOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				EngineID:       core.StringPtr("sampleEngineID"),
				SchemaName:     core.StringPtr("new_schema1"),
				BucketName:     core.StringPtr("sampleBucket"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.CreateSchema(createSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSchemas - Get schemas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchemas(getSchemasOptions *GetSchemasOptions)`, func() {
			getSchemasOptions := &watsonxdatav1.GetSchemasOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getSchemasOkBody, response, err := watsonxDataService.GetSchemas(getSchemasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getSchemasOkBody).ToNot(BeNil())
		})
	})

	Describe(`SaveQuery - Save query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SaveQuery(saveQueryOptions *SaveQueryOptions)`, func() {
			saveQueryOptions := &watsonxdatav1.SaveQueryOptions{
				QueryName:      core.StringPtr("query1"),
				CreatedBy:      core.StringPtr("<username>@<domain>.com"),
				Description:    core.StringPtr("query to get expense data"),
				QueryString:    core.StringPtr("select expenses from expenditure"),
				CreatedOn:      core.StringPtr("1608437933"),
				EngineID:       core.StringPtr("sampleEngingID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.SaveQuery(saveQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateQuery - Update query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateQuery(updateQueryOptions *UpdateQueryOptions)`, func() {
			updateQueryOptions := &watsonxdatav1.UpdateQueryOptions{
				QueryName:      core.StringPtr("query1"),
				QueryString:    core.StringPtr("testString"),
				Description:    core.StringPtr("testString"),
				NewQueryName:   core.StringPtr("query2"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.UpdateQuery(updateQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetQueries - Get queries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetQueries(getQueriesOptions *GetQueriesOptions)`, func() {
			getQueriesOptions := &watsonxdatav1.GetQueriesOptions{
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getQueriesOkBody, response, err := watsonxDataService.GetQueries(getQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getQueriesOkBody).ToNot(BeNil())
		})
	})

	Describe(`DeleteQuery - Delete query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteQuery(deleteQueryOptions *DeleteQueryOptions)`, func() {
			deleteQueryOptions := &watsonxdatav1.DeleteQueryOptions{
				QueryName:      core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteQuery(deleteQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`PostQuery - Run SQL statement`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostQuery(postQueryOptions *PostQueryOptions)`, func() {
			postQueryOptions := &watsonxdatav1.PostQueryOptions{
				Engine:         core.StringPtr("SampleEngine"),
				Catalog:        core.StringPtr("sampleCatalogName"),
				Schema:         core.StringPtr("testString"),
				SqlQuery:       core.StringPtr("testString"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.PostQuery(postQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ExplainAnalyzeStatement - Explain analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ExplainAnalyzeStatement(explainAnalyzeStatementOptions *ExplainAnalyzeStatementOptions)`, func() {
			explainAnalyzeStatementOptions := &watsonxdatav1.ExplainAnalyzeStatementOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				EngineID:       core.StringPtr("sampleEngineID"),
				SchemaName:     core.StringPtr("sampleSchemaName"),
				Statement:      core.StringPtr("show schemas in catalog"),
				Verbose:        core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			explainAnalyzeStatementCreatedBody, response, err := watsonxDataService.ExplainAnalyzeStatement(explainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(explainAnalyzeStatementCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ExplainStatement - Explain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ExplainStatement(explainStatementOptions *ExplainStatementOptions)`, func() {
			explainStatementOptions := &watsonxdatav1.ExplainStatementOptions{
				EngineID:       core.StringPtr("sampleEngingID"),
				Statement:      core.StringPtr("show schemas in catalog"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				Format:         core.StringPtr("json"),
				SchemaName:     core.StringPtr("sampleSchemaName"),
				Type:           core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			explainStatementCreatedBody, response, err := watsonxDataService.ExplainStatement(explainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(explainStatementCreatedBody).ToNot(BeNil())
		})
	})

	// ----
	Describe(`AddMetastoreToEngine - Add catalog to engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddMetastoreToEngine(addMetastoreToEngineOptions *AddMetastoreToEngineOptions)`, func() {
			addMetastoreToEngineOptions := &watsonxdatav1.AddMetastoreToEngineOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				EngineID:       core.StringPtr("sampleEngingID"),
				CreatedBy:      core.StringPtr("<username>@<domain>.com"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.AddMetastoreToEngine(addMetastoreToEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`RemoveCatalogFromEngine - Remove catalog from engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveCatalogFromEngine(removeCatalogFromEngineOptions *RemoveCatalogFromEngineOptions)`, func() {
			removeCatalogFromEngineOptions := &watsonxdatav1.RemoveCatalogFromEngineOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				EngineID:       core.StringPtr("sampleEngingID"),
				CreatedBy:      core.StringPtr("testString"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.RemoveCatalogFromEngine(removeCatalogFromEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetTables - Get tables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTables(getTablesOptions *GetTablesOptions)`, func() {
			getTablesOptions := &watsonxdatav1.GetTablesOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				SchemaName:     core.StringPtr("new_schema2"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getTablesOkBody, response, err := watsonxDataService.GetTables(getTablesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTablesOkBody).ToNot(BeNil())
		})
	})

	Describe(`UpdateTable - Update table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTable(updateTableOptions *UpdateTableOptions)`, func() {
			updateTableBodyAddColumnsItemsModel := &watsonxdatav1.UpdateTableBodyAddColumnsItems{
				ColumnComment: core.StringPtr("income column"),
				ColumnName:    core.StringPtr("income"),
				DataType:      core.StringPtr("varchar"),
			}

			updateTableBodyDropColumnsItemsModel := &watsonxdatav1.UpdateTableBodyDropColumnsItems{
				ColumnName: core.StringPtr("expenditure"),
			}

			updateTableBodyRenameColumnsItemsModel := &watsonxdatav1.UpdateTableBodyRenameColumnsItems{
				ColumnName:    core.StringPtr("expenditure"),
				NewColumnName: core.StringPtr("expenses"),
			}

			updateTableOptions := &watsonxdatav1.UpdateTableOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				SchemaName:     core.StringPtr("new_schema2"),
				TableName:      core.StringPtr("c1"),
				AddColumns:     []watsonxdatav1.UpdateTableBodyAddColumnsItems{*updateTableBodyAddColumnsItemsModel},
				DropColumns:    []watsonxdatav1.UpdateTableBodyDropColumnsItems{*updateTableBodyDropColumnsItemsModel},
				NewTableName:   core.StringPtr("updated_table_name"),
				RenameColumns:  []watsonxdatav1.UpdateTableBodyRenameColumnsItems{*updateTableBodyRenameColumnsItemsModel},
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.UpdateTable(updateTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetTableSnapshots - Get table snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTableSnapshots(getTableSnapshotsOptions *GetTableSnapshotsOptions)`, func() {
			getTableSnapshotsOptions := &watsonxdatav1.GetTableSnapshotsOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				SchemaName:     core.StringPtr("new_schema2"),
				TableName:      core.StringPtr("updated_table_name"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			getTableSnapshotsOkBody, response, err := watsonxDataService.GetTableSnapshots(getTableSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTableSnapshotsOkBody).ToNot(BeNil())
		})
	})

	Describe(`RollbackSnapshot - Rollback snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RollbackSnapshot(rollbackSnapshotOptions *RollbackSnapshotOptions)`, func() {
			rollbackSnapshotOptions := &watsonxdatav1.RollbackSnapshotOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CatalogName:    core.StringPtr("sampleCatalogName"),
				SchemaName:     core.StringPtr("new_schema2"),
				SnapshotID:     core.StringPtr("2332342122211222"),
				TableName:      core.StringPtr("updated_table_name"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			successResponse, response, err := watsonxDataService.RollbackSnapshot(rollbackSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ParseCsv - Parse CSV for table creation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ParseCsv(parseCsvOptions *ParseCsvOptions)`, func() {
			parseCsvOptions := &watsonxdatav1.ParseCsvOptions{
				Engine:         core.StringPtr("sampleEngine"),
				ParseFile:      core.StringPtr("testString"),
				FileType:       core.StringPtr("testString"),
				Accept:         core.StringPtr("application/json"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			result, response, err := watsonxDataService.ParseCsv(parseCsvOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`UplaodCsv - Upload CSV for table creation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UplaodCsv(uplaodCsvOptions *UplaodCsvOptions)`, func() {
			uplaodCsvOptions := &watsonxdatav1.UplaodCsvOptions{
				Engine:           core.StringPtr("sampleEngine"),
				Catalog:          core.StringPtr("sampleCatalogName"),
				Schema:           core.StringPtr("sampleSchema"),
				TableName:        core.StringPtr("sampleTableName"),
				IngestionJobName: core.StringPtr("testString"),
				Scheduled:        core.StringPtr("testString"),
				CreatedBy:        core.StringPtr("testString"),
				TargetTable:      core.StringPtr("testString"),
				HeadersVar:       core.StringPtr("testString"),
				Csv:              core.StringPtr("testString"),
				Accept:           core.StringPtr("application/json"),
				AuthInstanceID:   core.StringPtr("testString"),
			}

			result, response, err := watsonxDataService.UplaodCsv(uplaodCsvOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeleteEngineUsers - Revoke permission to access engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEngineUsers(deleteEngineUsersOptions *DeleteEngineUsersOptions)`, func() {
			deleteEngineUsersOptions := &watsonxdatav1.DeleteEngineUsersOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				Groups:         []string{"testString"},
				Users:          []string{"sampleUser"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteEngineUsers(deleteEngineUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDbConnUsers - Revoke permission to access db connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDbConnUsers(deleteDbConnUsersOptions *DeleteDbConnUsersOptions)`, func() {
			deleteDbConnUsersOptions := &watsonxdatav1.DeleteDbConnUsersOptions{
				DatabaseID:     core.StringPtr("sampleDatabaseID"),
				Groups:         []string{"testString"},
				Users:          []string{"sampleUser"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteDbConnUsers(deleteDbConnUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteCatalogUsers - Revoke multiple users and groups permission to access catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalogUsers(deleteCatalogUsersOptions *DeleteCatalogUsersOptions)`, func() {
			deleteCatalogUsersOptions := &watsonxdatav1.DeleteCatalogUsersOptions{
				CatalogName:    core.StringPtr("sampledbcatalog"),
				Groups:         []string{"testString"},
				Users:          []string{"sampleUser"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteCatalogUsers(deleteCatalogUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteMetastoreUsers - Revoke permission to access metastore`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteMetastoreUsers(deleteMetastoreUsersOptions *DeleteMetastoreUsersOptions)`, func() {
			deleteMetastoreUsersOptions := &watsonxdatav1.DeleteMetastoreUsersOptions{
				MetastoreName:  core.StringPtr("sampledbcatalog"),
				Groups:         []string{"testString"},
				Users:          []string{"sampleUser"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteMetastoreUsers(deleteMetastoreUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
	Describe(`DeleteDataPolicies - Revoke data policy access management policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataPolicies(deleteDataPoliciesOptions *DeleteDataPoliciesOptions)`, func() {
			deleteDataPoliciesOptions := &watsonxdatav1.DeleteDataPoliciesOptions{
				DataPolicies:   []string{"sampleDataPolicy"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteDataPolicies(deleteDataPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDataPolicy - Revoke data policy access management policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataPolicy(deleteDataPolicyOptions *DeleteDataPolicyOptions)`, func() {
			deleteDataPolicyOptions := &watsonxdatav1.DeleteDataPolicyOptions{
				PolicyName:     core.StringPtr("samplePolicyName"),
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteDataPolicy(deleteDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteBucketUsers - Revoke permission to access bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBucketUsers(deleteBucketUsersOptions *DeleteBucketUsersOptions)`, func() {
			deleteBucketUsersOptions := &watsonxdatav1.DeleteBucketUsersOptions{
				BucketID:       core.StringPtr("sampleBucketID"),
				Groups:         []string{"sampleGroup1"},
				Users:          []string{"sampleUser"},
				LhInstanceID:   core.StringPtr("sampleInstanceID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteBucketUsers(deleteBucketUsersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDatabaseCatalog - Delete database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
			deleteDatabaseCatalogOptions := &watsonxdatav1.DeleteDatabaseCatalogOptions{
				DatabaseID:     core.StringPtr("sampleDatabaseID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSchema - Delete schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
			deleteSchemaOptions := &watsonxdatav1.DeleteSchemaOptions{
				CatalogName:    core.StringPtr("sampleCatalogName"),
				EngineID:       core.StringPtr("sampleEngineID"),
				SchemaName:     core.StringPtr("sampleSchemaName"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteSchema(deleteSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTable - Delete table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTable(deleteTableOptions *DeleteTableOptions)`, func() {
			deleteTableBodyDeleteTablesItemsModel := &watsonxdatav1.DeleteTableBodyDeleteTablesItems{
				CatalogName: core.StringPtr("sampleCatalogName"),
				SchemaName:  core.StringPtr("sampleSchemaName"),
				TableName:   core.StringPtr("sampleTableName"),
			}

			deleteTableOptions := &watsonxdatav1.DeleteTableOptions{
				DeleteTables:   []watsonxdatav1.DeleteTableBodyDeleteTablesItems{*deleteTableBodyDeleteTablesItemsModel},
				EngineID:       core.StringPtr("sampleEngineID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteTable(deleteTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`UnregisterBucket - Unregister Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnregisterBucket(unregisterBucketOptions *UnregisterBucketOptions)`, func() {
			unregisterBucketOptions := &watsonxdatav1.UnregisterBucketOptions{
				BucketID:       core.StringPtr("sampleBucketID"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.UnregisterBucket(unregisterBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`PauseEngine - Pause engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PauseEngine(pauseEngineOptions *PauseEngineOptions)`, func() {
			pauseEngineOptions := &watsonxdatav1.PauseEngineOptions{
				EngineID:       core.StringPtr("sampleEngingID"),
				CreatedBy:      core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			pauseEngineCreatedBody, response, err := watsonxDataService.PauseEngine(pauseEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pauseEngineCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ResumeEngine - Resume engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumeEngine(resumeEngineOptions *ResumeEngineOptions)`, func() {
			resumeEngineOptions := &watsonxdatav1.ResumeEngineOptions{
				EngineID:       core.StringPtr("sampleEngineID"),
				CreatedBy:      core.StringPtr("<username>@<domain>.com"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			resumeEngineCreatedBody, response, err := watsonxDataService.ResumeEngine(resumeEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resumeEngineCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`DeleteEngine - Delete engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
			deleteEngineOptions := &watsonxdatav1.DeleteEngineOptions{
				EngineID:       core.StringPtr("sampleEngingID"),
				CreatedBy:      core.StringPtr("<username>@<domain>.com"),
				AuthInstanceID: core.StringPtr("sampleAuthInstanceID"),
			}

			response, err := watsonxDataService.DeleteEngine(deleteEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
