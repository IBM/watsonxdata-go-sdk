//go:build integration

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

package watsonxdatav2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the watsonxdatav2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`WatsonxDataV2 Integration Tests`, func() {
	const externalConfigFile = "../watsonx_data_v2.env"

	var (
		err          error
		watsonxDataService *watsonxdatav2.WatsonxDataV2
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(watsonxdatav2.DefaultServiceName)
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

			bearerToken := "[YOUR_BEARER_TOKEN]"
			authenticator, err := core.NewBearerTokenAuthenticator(bearerToken)
			if err != nil {
				panic(err)
			}
			watsonxDataServiceOptions := &watsonxdatav2.WatsonxDataV2Options{
				Authenticator: authenticator,
			}			

			watsonxDataService, err = watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(watsonxDataServiceOptions)
			Expect(err).To(BeNil())
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(watsonxDataService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			watsonxDataService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListBucketRegistrations - Get bucket registrations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions)`, func() {
			listBucketRegistrationsOptions := &watsonxdatav2.ListBucketRegistrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistrationCollection, response, err := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateBucketRegistration - Register bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions)`, func() {
			bucketDetailsModel := &watsonxdatav2.BucketDetails{
				AccessKey: core.StringPtr("<access_key>"),
				BucketName: core.StringPtr("sample-bucket"),
				Endpoint: core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/"),
				SecretKey: core.StringPtr("secret_key"),
			}

			bucketCatalogModel := &watsonxdatav2.BucketCatalog{
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				CatalogType: core.StringPtr("iceberg"),
			}

			createBucketRegistrationOptions := &watsonxdatav2.CreateBucketRegistrationOptions{
				BucketDetails: bucketDetailsModel,
				BucketType: core.StringPtr("ibm_cos"),
				Description: core.StringPtr("COS bucket for customer data"),
				ManagedBy: core.StringPtr("ibm"),
				AssociatedCatalog: bucketCatalogModel,
				BucketDisplayName: core.StringPtr("sample-bucket-displayname"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"bucket-tag1", "bucket-tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`GetBucketRegistration - Get bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions)`, func() {
			getBucketRegistrationOptions := &watsonxdatav2.GetBucketRegistrationOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`UpdateBucketRegistration - Update bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateBucketRegistrationOptions := &watsonxdatav2.UpdateBucketRegistrationOptions{
				BucketID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`CreateActivateBucket - Activate Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions)`, func() {
			createActivateBucketOptions := &watsonxdatav2.CreateActivateBucketOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createActivateBucketCreatedBody, response, err := watsonxDataService.CreateActivateBucket(createActivateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createActivateBucketCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListBucketObjects - List bucket objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions)`, func() {
			listBucketObjectsOptions := &watsonxdatav2.ListBucketObjectsOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistrationObjectCollection, response, err := watsonxDataService.ListBucketObjects(listBucketObjectsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistrationObjectCollection).ToNot(BeNil())
		})
	})

	Describe(`TestBucketConnection - Check bucket credentials to be valid`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestBucketConnection(testBucketConnectionOptions *TestBucketConnectionOptions)`, func() {
			testBucketConnectionOptions := &watsonxdatav2.TestBucketConnectionOptions{
				AccessKey: core.StringPtr("<access_key>"),
				BucketName: core.StringPtr("sample-bucket"),
				BucketType: core.StringPtr("ibm_cos"),
				Endpoint: core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/"),
				Region: core.StringPtr("us-south"),
				SecretKey: core.StringPtr("secret_key"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			testBucketConnectionOkBody, response, err := watsonxDataService.TestBucketConnection(testBucketConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testBucketConnectionOkBody).ToNot(BeNil())
		})
	})

	Describe(`CreateDriverDatabaseCatalog - Add/Create database with driver`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptions *CreateDriverDatabaseCatalogOptions)`, func() {
			createDriverDatabaseCatalogOptions := &watsonxdatav2.CreateDriverDatabaseCatalogOptions{
				Driver: CreateMockReader("This is a mock file."),
				DriverFileName: core.StringPtr("testString"),
				DatabaseDisplayName: core.StringPtr("testString"),
				DatabaseType: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				Hostname: core.StringPtr("testString"),
				Port: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
				Password: core.StringPtr("testString"),
				DatabaseName: core.StringPtr("testString"),
				DriverContentType: core.StringPtr("testString"),
				Certificate: core.StringPtr("testString"),
				CertificateExtension: core.StringPtr("testString"),
				Ssl: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				CreatedOn: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.CreateDriverDatabaseCatalog(createDriverDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`ListDatabaseRegistrations - Get databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions)`, func() {
			listDatabaseRegistrationsOptions := &watsonxdatav2.ListDatabaseRegistrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistrationCollection, response, err := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDatabaseRegistration - Add/Create database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions)`, func() {
			databaseCatalogModel := &watsonxdatav2.DatabaseCatalog{
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				CatalogType: core.StringPtr("iceberg"),
			}

			databaseDetailsModel := &watsonxdatav2.DatabaseDetails{
				Certificate: core.StringPtr("contents of a pem/crt file"),
				CertificateExtension: core.StringPtr("pem/crt"),
				DatabaseName: core.StringPtr("new_database"),
				Hostname: core.StringPtr("db2@<hostname>.com"),
				HostnameInCertificate: core.StringPtr("samplehostname"),
				Hosts: core.StringPtr("abc.com:1234,xyz.com:4321"),
				Password: core.StringPtr("samplepassword"),
				Port: core.Int64Ptr(int64(4553)),
				Sasl: core.BoolPtr(true),
				Ssl: core.BoolPtr(true),
				Tables: core.StringPtr("kafka_table_name"),
				Username: core.StringPtr("sampleuser"),
				ValidateServerCertificate: core.BoolPtr(true),
			}

			databaseRegistrationPrototypeDatabasePropertiesItemsModel := &watsonxdatav2.DatabaseRegistrationPrototypeDatabasePropertiesItems{
				Encrypt: core.BoolPtr(true),
				Key: core.StringPtr("abc"),
				Value: core.StringPtr("xyz"),
			}

			createDatabaseRegistrationOptions := &watsonxdatav2.CreateDatabaseRegistrationOptions{
				DatabaseDisplayName: core.StringPtr("new_database"),
				DatabaseType: core.StringPtr("db2"),
				AssociatedCatalog: databaseCatalogModel,
				CreatedOn: core.StringPtr("1686792721"),
				DatabaseDetails: databaseDetailsModel,
				DatabaseProperties: []watsonxdatav2.DatabaseRegistrationPrototypeDatabasePropertiesItems{*databaseRegistrationPrototypeDatabasePropertiesItemsModel},
				Description: core.StringPtr("db2 extenal database description"),
				Tags: []string{"testdatabase", "userdatabase"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`GetDatabase - Get database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabase(getDatabaseOptions *GetDatabaseOptions)`, func() {
			getDatabaseOptions := &watsonxdatav2.GetDatabaseOptions{
				DatabaseID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.GetDatabase(getDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`UpdateDatabase - Update database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateDatabaseOptions := &watsonxdatav2.UpdateDatabaseOptions{
				DatabaseID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.UpdateDatabase(updateDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`ValidateDatabaseConnection - Validate database connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateDatabaseConnection(validateDatabaseConnectionOptions *ValidateDatabaseConnectionOptions)`, func() {
			validateDatabaseBodyDatabaseDetailsModel := &watsonxdatav2.ValidateDatabaseBodyDatabaseDetails{
				DatabaseName: core.StringPtr("sampledatabase"),
				Hostname: core.StringPtr("db2@hostname.com"),
				Password: core.StringPtr("samplepassword"),
				Port: core.Int64Ptr(int64(4553)),
				Sasl: core.BoolPtr(true),
				Ssl: core.BoolPtr(true),
				Tables: core.StringPtr("kafka_table_name"),
				Username: core.StringPtr("sampleuser"),
				ValidateServerCertificate: core.BoolPtr(true),
			}

			validateDatabaseConnectionOptions := &watsonxdatav2.ValidateDatabaseConnectionOptions{
				DatabaseDetails: validateDatabaseBodyDatabaseDetailsModel,
				DatabaseType: core.StringPtr("netezza"),
				Certificate: core.StringPtr("contents of a pem/crt file"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			testDatabaseConnectionResponse, response, err := watsonxDataService.ValidateDatabaseConnection(validateDatabaseConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testDatabaseConnectionResponse).ToNot(BeNil())
		})
	})

	Describe(`ListDb2Engines - Get list of db2 engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions)`, func() {
			listDb2EnginesOptions := &watsonxdatav2.ListDb2EnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2EngineCollection, response, err := watsonxDataService.ListDb2Engines(listDb2EnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(db2EngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDb2Engine - Create db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions)`, func() {
			db2EngineDetailsBodyModel := &watsonxdatav2.Db2EngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createDb2EngineOptions := &watsonxdatav2.CreateDb2EngineOptions{
				Origin: core.StringPtr("external"),
				Type: core.StringPtr("db2"),
				Description: core.StringPtr("db2 engine description"),
				EngineDetails: db2EngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2Engine, response, err := watsonxDataService.CreateDb2Engine(createDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(db2Engine).ToNot(BeNil())
		})
	})

	Describe(`UpdateDb2Engine - Update db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateDb2EngineOptions := &watsonxdatav2.UpdateDb2EngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2Engine, response, err := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(db2Engine).ToNot(BeNil())
		})
	})

	Describe(`GetEngines - Get all engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEngines(getEnginesOptions *GetEnginesOptions)`, func() {
			getEnginesOptions := &watsonxdatav2.GetEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			engines, response, err := watsonxDataService.GetEngines(getEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(engines).ToNot(BeNil())
		})
	})

	Describe(`GetDeployments - Get deployments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDeployments(getDeploymentsOptions *GetDeploymentsOptions)`, func() {
			getDeploymentsOptions := &watsonxdatav2.GetDeploymentsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			getDeploymentsOkBody, response, err := watsonxDataService.GetDeployments(getDeploymentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDeploymentsOkBody).ToNot(BeNil())
		})
	})

	Describe(`ListNetezzaEngines - Get list of netezza engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions)`, func() {
			listNetezzaEnginesOptions := &watsonxdatav2.ListNetezzaEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngineCollection, response, err := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(netezzaEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateNetezzaEngine - Create netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions)`, func() {
			netezzaEngineDetailsBodyModel := &watsonxdatav2.NetezzaEngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createNetezzaEngineOptions := &watsonxdatav2.CreateNetezzaEngineOptions{
				Origin: core.StringPtr("external"),
				Type: core.StringPtr("netezza"),
				Description: core.StringPtr("netezza engine description"),
				EngineDetails: netezzaEngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngine, response, err := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(netezzaEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdateNetezzaEngine - Update netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateNetezzaEngineOptions := &watsonxdatav2.UpdateNetezzaEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngine, response, err := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(netezzaEngine).ToNot(BeNil())
		})
	})

	Describe(`ListOtherEngines - List other engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions)`, func() {
			listOtherEnginesOptions := &watsonxdatav2.ListOtherEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			otherEngineCollection, response, err := watsonxDataService.ListOtherEngines(listOtherEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(otherEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateOtherEngine - Create other engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions)`, func() {
			otherEngineDetailsBodyModel := &watsonxdatav2.OtherEngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
				EngineType: core.StringPtr("netezza"),
			}

			createOtherEngineOptions := &watsonxdatav2.CreateOtherEngineOptions{
				EngineDetails: otherEngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine01"),
				Description: core.StringPtr("external engine description"),
				Origin: core.StringPtr("external"),
				Tags: []string{"tag1", "tag2"},
				Type: core.StringPtr("netezza"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			otherEngine, response, err := watsonxDataService.CreateOtherEngine(createOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(otherEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestissimoEngines - Get list of prestissimo engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestissimoEngines(listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions)`, func() {
			listPrestissimoEnginesOptions := &watsonxdatav2.ListPrestissimoEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngineCollection, response, err := watsonxDataService.ListPrestissimoEngines(listPrestissimoEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngine - Create prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngine(createPrestissimoEngineOptions *CreatePrestissimoEngineOptions)`, func() {
			prestissimoNodeDescriptionBodyModel := &watsonxdatav2.PrestissimoNodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			prestissimoEndpointsModel := &watsonxdatav2.PrestissimoEndpoints{
				ApplicationsApi: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>"),
				HistoryServerEndpoint: core.StringPtr("$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server"),
				SparkAccessEndpoint: core.StringPtr("$HOST/analytics-engine/details/spark-<instance_id>"),
				SparkJobsV4Endpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications"),
				SparkKernelEndpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels"),
				ViewHistoryServer: core.StringPtr("testString"),
				WxdApplicationEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/applications"),
			}

			prestissimoEngineDetailsModel := &watsonxdatav2.PrestissimoEngineDetails{
				ApiKey: core.StringPtr("<api_key>"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				Coordinator: prestissimoNodeDescriptionBodyModel,
				Endpoints: prestissimoEndpointsModel,
				InstanceID: core.StringPtr("instance_id"),
				ManagedBy: core.StringPtr("fully/self"),
				MetastoreHost: core.StringPtr("1.2.3.4"),
				SizeConfig: core.StringPtr("starter"),
				Worker: prestissimoNodeDescriptionBodyModel,
			}

			createPrestissimoEngineOptions := &watsonxdatav2.CreatePrestissimoEngineOptions{
				Origin: core.StringPtr("native"),
				Type: core.StringPtr("prestissimo"),
				AssociatedCatalogs: []string{"hive_data"},
				Description: core.StringPtr("prestissimo engine description"),
				EngineDetails: prestissimoEngineDetailsModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"tag1", "tag2"},
				Version: core.StringPtr("1.2.3"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.CreatePrestissimoEngine(createPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngine - Get prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngine(getPrestissimoEngineOptions *GetPrestissimoEngineOptions)`, func() {
			getPrestissimoEngineOptions := &watsonxdatav2.GetPrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.GetPrestissimoEngine(getPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestissimoEngine - Update prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestissimoEngine(updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updatePrestissimoEngineOptions := &watsonxdatav2.UpdatePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.UpdatePrestissimoEngine(updatePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestissimoEngineCatalogs - Get prestissimo engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions)`, func() {
			listPrestissimoEngineCatalogsOptions := &watsonxdatav2.ListPrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplacePrestissimoEngineCatalogs - Associate catalogs to a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplacePrestissimoEngineCatalogs(replacePrestissimoEngineCatalogsOptions *ReplacePrestissimoEngineCatalogsOptions)`, func() {
			replacePrestissimoEngineCatalogsOptions := &watsonxdatav2.ReplacePrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ReplacePrestissimoEngineCatalogs(replacePrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngineCatalog - Get prestissimo engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions)`, func() {
			getPrestissimoEngineCatalogOptions := &watsonxdatav2.GetPrestissimoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEnginePause - Pause prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEnginePause(createPrestissimoEnginePauseOptions *CreatePrestissimoEnginePauseOptions)`, func() {
			createPrestissimoEnginePauseOptions := &watsonxdatav2.CreatePrestissimoEnginePauseOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreatePrestissimoEnginePause(createPrestissimoEnginePauseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`RunPrestissimoExplainStatement - Explain query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions)`, func() {
			runPrestissimoExplainStatementOptions := &watsonxdatav2.RunPrestissimoExplainStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Format: core.StringPtr("json"),
				Type: core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			resultPrestissimoExplainStatement, response, err := watsonxDataService.RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resultPrestissimoExplainStatement).ToNot(BeNil())
		})
	})

	Describe(`RunPrestissimoExplainAnalyzeStatement - Explain analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions)`, func() {
			runPrestissimoExplainAnalyzeStatementOptions := &watsonxdatav2.RunPrestissimoExplainAnalyzeStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Verbose: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			resultRunPrestissimoExplainAnalyzeStatement, response, err := watsonxDataService.RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resultRunPrestissimoExplainAnalyzeStatement).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngineRestart - Restart a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngineRestart(createPrestissimoEngineRestartOptions *CreatePrestissimoEngineRestartOptions)`, func() {
			createPrestissimoEngineRestartOptions := &watsonxdatav2.CreatePrestissimoEngineRestartOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreatePrestissimoEngineRestart(createPrestissimoEngineRestartOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngineResume - Resume prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngineResume(createPrestissimoEngineResumeOptions *CreatePrestissimoEngineResumeOptions)`, func() {
			createPrestissimoEngineResumeOptions := &watsonxdatav2.CreatePrestissimoEngineResumeOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreatePrestissimoEngineResume(createPrestissimoEngineResumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngineScale - Scale a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngineScale(createPrestissimoEngineScaleOptions *CreatePrestissimoEngineScaleOptions)`, func() {
			prestissimoNodeDescriptionBodyModel := &watsonxdatav2.PrestissimoNodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			createPrestissimoEngineScaleOptions := &watsonxdatav2.CreatePrestissimoEngineScaleOptions{
				EngineID: core.StringPtr("testString"),
				Coordinator: prestissimoNodeDescriptionBodyModel,
				Worker: prestissimoNodeDescriptionBodyModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreatePrestissimoEngineScale(createPrestissimoEngineScaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngines - Get list of presto engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions)`, func() {
			listPrestoEnginesOptions := &watsonxdatav2.ListPrestoEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineCollection, response, err := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestoEngine - Create presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestoEngine(createPrestoEngineOptions *CreatePrestoEngineOptions)`, func() {
			nodeDescriptionBodyModel := &watsonxdatav2.NodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			engineDetailsBodyModel := &watsonxdatav2.EngineDetailsBody{
				ApiKey: core.StringPtr("<api_key>"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				Coordinator: nodeDescriptionBodyModel,
				InstanceID: core.StringPtr("instance_id"),
				ManagedBy: core.StringPtr("fully/self"),
				SizeConfig: core.StringPtr("starter"),
				Worker: nodeDescriptionBodyModel,
			}

			createPrestoEngineOptions := &watsonxdatav2.CreatePrestoEngineOptions{
				Origin: core.StringPtr("native"),
				Type: core.StringPtr("presto"),
				AssociatedCatalogs: []string{"iceberg_data", "hive_data"},
				Description: core.StringPtr("presto engine for running sql queries"),
				EngineDetails: engineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"tag1", "tag2"},
				Version: core.StringPtr("1.2.3"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.CreatePrestoEngine(createPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngine - Get presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions)`, func() {
			getPrestoEngineOptions := &watsonxdatav2.GetPrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.GetPrestoEngine(getPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestoEngine - Update presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestoEngine(updatePrestoEngineOptions *UpdatePrestoEngineOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updatePrestoEngineOptions := &watsonxdatav2.UpdatePrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.UpdatePrestoEngine(updatePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngineCatalogs - Get presto engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions)`, func() {
			listPrestoEngineCatalogsOptions := &watsonxdatav2.ListPrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplacePrestoEngineCatalogs - Associate catalogs to presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptions *ReplacePrestoEngineCatalogsOptions)`, func() {
			replacePrestoEngineCatalogsOptions := &watsonxdatav2.ReplacePrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ReplacePrestoEngineCatalogs(replacePrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngineCatalog - Get presto engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions)`, func() {
			getPrestoEngineCatalogOptions := &watsonxdatav2.GetPrestoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`CreateEnginePause - Pause presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEnginePause(createEnginePauseOptions *CreateEnginePauseOptions)`, func() {
			createEnginePauseOptions := &watsonxdatav2.CreateEnginePauseOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEnginePauseCreatedBody, response, err := watsonxDataService.CreateEnginePause(createEnginePauseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createEnginePauseCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`RunExplainStatement - Explain presto query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions)`, func() {
			runExplainStatementOptions := &watsonxdatav2.RunExplainStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Format: core.StringPtr("json"),
				Type: core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			runExplainStatementOkBody, response, err := watsonxDataService.RunExplainStatement(runExplainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runExplainStatementOkBody).ToNot(BeNil())
		})
	})

	Describe(`RunExplainAnalyzeStatement - Explain presto analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions)`, func() {
			runExplainAnalyzeStatementOptions := &watsonxdatav2.RunExplainAnalyzeStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Verbose: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			runExplainAnalyzeStatementOkBody, response, err := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runExplainAnalyzeStatementOkBody).ToNot(BeNil())
		})
	})

	Describe(`CreateEngineRestart - Restart a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEngineRestart(createEngineRestartOptions *CreateEngineRestartOptions)`, func() {
			createEngineRestartOptions := &watsonxdatav2.CreateEngineRestartOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineRestartCreatedBody, response, err := watsonxDataService.CreateEngineRestart(createEngineRestartOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createEngineRestartCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`CreateEngineResume - Resume presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEngineResume(createEngineResumeOptions *CreateEngineResumeOptions)`, func() {
			createEngineResumeOptions := &watsonxdatav2.CreateEngineResumeOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineResumeCreatedBody, response, err := watsonxDataService.CreateEngineResume(createEngineResumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createEngineResumeCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`CreateEngineScale - Scale a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEngineScale(createEngineScaleOptions *CreateEngineScaleOptions)`, func() {
			nodeDescriptionModel := &watsonxdatav2.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			createEngineScaleOptions := &watsonxdatav2.CreateEngineScaleOptions{
				EngineID: core.StringPtr("testString"),
				Coordinator: nodeDescriptionModel,
				Worker: nodeDescriptionModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineScaleCreatedBody, response, err := watsonxDataService.CreateEngineScale(createEngineScaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createEngineScaleCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngines - List all spark engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions)`, func() {
			listSparkEnginesOptions := &watsonxdatav2.ListSparkEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineCollection, response, err := watsonxDataService.ListSparkEngines(listSparkEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngine - Create spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions)`, func() {
			sparkEngineDetailsPrototypeModel := &watsonxdatav2.SparkEngineDetailsPrototype{
				ApiKey: core.StringPtr("apikey"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				InstanceID: core.StringPtr("spark-id"),
				ManagedBy: core.StringPtr("fully/self"),
			}

			createSparkEngineOptions := &watsonxdatav2.CreateSparkEngineOptions{
				Origin: core.StringPtr("external"),
				Type: core.StringPtr("spark"),
				Description: core.StringPtr("spark engine description"),
				EngineDetails: sparkEngineDetailsPrototypeModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.CreateSparkEngine(createSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdateSparkEngine - Update spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateSparkEngineOptions := &watsonxdatav2.UpdateSparkEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngineApplications - List all applications in a spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions)`, func() {
			listSparkEngineApplicationsOptions := &watsonxdatav2.ListSparkEngineApplicationsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationStatusCollection, response, err := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationStatusCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngineApplication - Submit engine applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions)`, func() {
			sparkApplicationDetailsConfModel := &watsonxdatav2.SparkApplicationDetailsConf{
				SparkAppName: core.StringPtr("MyJob"),
				SparkHiveMetastoreClientAuthMode: core.StringPtr("PLAIN"),
				SparkHiveMetastoreClientPlainPassword: core.StringPtr("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."),
				SparkHiveMetastoreClientPlainUsername: core.StringPtr("ibm_lh_token_admin"),
				SparkHiveMetastoreTruststorePassword: core.StringPtr("changeit"),
				SparkHiveMetastoreTruststorePath: core.StringPtr("file:///opt/ibm/jdk/lib/security/cacerts"),
				SparkHiveMetastoreTruststoreType: core.StringPtr("JKS"),
				SparkHiveMetastoreUseSsl: core.StringPtr("true"),
				SparkSqlCatalogImplementation: core.StringPtr("Spark Catalog Implementation"),
				SparkSqlCatalogLakehouse: core.StringPtr("org.apache.iceberg.spark.SparkCatalog"),
				SparkSqlCatalogLakehouseType: core.StringPtr("Spark Catalog Type"),
				SparkSqlCatalogLakehouseURI: core.StringPtr("Spark Catalog URI"),
				SparkSqlExtensions: core.StringPtr("org.apache.iceberg.spark.extensions.IcebergSparkSessionExtensions"),
				SparkSqlIcebergVectorizationEnabled: core.StringPtr("false"),
			}

			sparkApplicationDetailsModel := &watsonxdatav2.SparkApplicationDetails{
				Application: core.StringPtr("s3://mybucket/wordcount.py"),
				Arguments: []string{"people.txt"},
				Conf: sparkApplicationDetailsConfModel,
				Env: map[string]interface{}{"anyKey": "anyValue"},
				Name: core.StringPtr("SparkApplicaton1"),
			}

			createSparkEngineApplicationOptions := &watsonxdatav2.CreateSparkEngineApplicationOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationDetails: sparkApplicationDetailsModel,
				JobEndpoint: core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications"),
				ServiceInstanceID: core.StringPtr("testString"),
				Type: core.StringPtr("iae"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationStatus, response, err := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkEngineApplicationStatus).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineApplicationStatus - Get spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions)`, func() {
			getSparkEngineApplicationStatusOptions := &watsonxdatav2.GetSparkEngineApplicationStatusOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationStatus, response, err := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationStatus).ToNot(BeNil())
		})
	})

	Describe(`TestLHConsole - Readiness API`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestLHConsole(testLHConsoleOptions *TestLHConsoleOptions)`, func() {
			testLhConsoleOptions := &watsonxdatav2.TestLHConsoleOptions{
			}

			successResponse, response, err := watsonxDataService.TestLHConsole(testLhConsoleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListCatalogs - List all registered catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
			listCatalogsOptions := &watsonxdatav2.ListCatalogsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListCatalogs(listCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetCatalog - Get catalog properties by catalog_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
			getCatalogOptions := &watsonxdatav2.GetCatalogOptions{
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetCatalog(getCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`ListSchemas - List all schemas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchemas(listSchemasOptions *ListSchemasOptions)`, func() {
			listSchemasOptions := &watsonxdatav2.ListSchemasOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			listSchemasOkBody, response, err := watsonxDataService.ListSchemas(listSchemasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listSchemasOkBody).ToNot(BeNil())
		})
	})

	Describe(`CreateSchema - Create schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
			createSchemaOptions := &watsonxdatav2.CreateSchemaOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				CustomPath: core.StringPtr("sample-path"),
				SchemaName: core.StringPtr("SampleSchema1"),
				BucketName: core.StringPtr("sample-bucket"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createSchemaCreatedBody, response, err := watsonxDataService.CreateSchema(createSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createSchemaCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListTables - List all tables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTables(listTablesOptions *ListTablesOptions)`, func() {
			listTablesOptions := &watsonxdatav2.ListTablesOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			tableCollection, response, err := watsonxDataService.ListTables(listTablesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableCollection).ToNot(BeNil())
		})
	})

	Describe(`GetTable - Get table details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTable(getTableOptions *GetTableOptions)`, func() {
			getTableOptions := &watsonxdatav2.GetTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			table, response, err := watsonxDataService.GetTable(getTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(table).ToNot(BeNil())
		})
	})

	Describe(`UpdateTable - Alter table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTable(updateTableOptions *UpdateTableOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateTableOptions := &watsonxdatav2.UpdateTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			table, response, err := watsonxDataService.UpdateTable(updateTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(table).ToNot(BeNil())
		})
	})

	Describe(`ListColumns - List all columns of a table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListColumns(listColumnsOptions *ListColumnsOptions)`, func() {
			listColumnsOptions := &watsonxdatav2.ListColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			columnCollection, response, err := watsonxDataService.ListColumns(listColumnsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(columnCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateColumns - Add column(s)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateColumns(createColumnsOptions *CreateColumnsOptions)`, func() {
			columnModel := &watsonxdatav2.Column{
				ColumnName: core.StringPtr("expenses"),
				Comment: core.StringPtr("expenses column"),
				Extra: core.StringPtr("varchar"),
				Length: core.StringPtr("30"),
				Scale: core.StringPtr("2"),
				Type: core.StringPtr("varchar"),
			}

			createColumnsOptions := &watsonxdatav2.CreateColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				Columns: []watsonxdatav2.Column{*columnModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			columnCollection, response, err := watsonxDataService.CreateColumns(createColumnsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(columnCollection).ToNot(BeNil())
		})
	})

	Describe(`UpdateColumn - Alter column`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateColumn(updateColumnOptions *UpdateColumnOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateColumnOptions := &watsonxdatav2.UpdateColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				ColumnID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			column, response, err := watsonxDataService.UpdateColumn(updateColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(column).ToNot(BeNil())
		})
	})

	Describe(`ListTableSnapshots - Get table snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions)`, func() {
			listTableSnapshotsOptions := &watsonxdatav2.ListTableSnapshotsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			tableSnapshotCollection, response, err := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableSnapshotCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSnapshot - Rollback snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSnapshot(replaceSnapshotOptions *ReplaceSnapshotOptions)`, func() {
			replaceSnapshotOptions := &watsonxdatav2.ReplaceSnapshotOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			replaceSnapshotCreatedBody, response, err := watsonxDataService.ReplaceSnapshot(replaceSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(replaceSnapshotCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`UpdateSyncCatalog - External Iceberg table registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateSyncCatalogOptions := &watsonxdatav2.UpdateSyncCatalogOptions{
				CatalogID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			updateSyncCatalogOkBody, response, err := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateSyncCatalogOkBody).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusServices - Get list of milvus services`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusServices(listMilvusServicesOptions *ListMilvusServicesOptions)`, func() {
			listMilvusServicesOptions := &watsonxdatav2.ListMilvusServicesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusServiceCollection, response, err := watsonxDataService.ListMilvusServices(listMilvusServicesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusServiceCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateMilvusService - Create milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMilvusService(createMilvusServiceOptions *CreateMilvusServiceOptions)`, func() {
			createMilvusServiceOptions := &watsonxdatav2.CreateMilvusServiceOptions{
				Origin: core.StringPtr("native"),
				Type: core.StringPtr("milvus"),
				Description: core.StringPtr("milvus service for running sql queries"),
				ServiceDisplayName: core.StringPtr("sampleService"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.CreateMilvusService(createMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`GetMilvusService - Get milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMilvusService(getMilvusServiceOptions *GetMilvusServiceOptions)`, func() {
			getMilvusServiceOptions := &watsonxdatav2.GetMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.GetMilvusService(getMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`UpdateMilvusService - Update milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateMilvusService(updateMilvusServiceOptions *UpdateMilvusServiceOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav2.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateMilvusServiceOptions := &watsonxdatav2.UpdateMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				Body: []watsonxdatav2.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.UpdateMilvusService(updateMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`DeleteBucketRegistration - Unregister Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBucketRegistration(deleteBucketRegistrationOptions *DeleteBucketRegistrationOptions)`, func() {
			deleteBucketRegistrationOptions := &watsonxdatav2.DeleteBucketRegistrationOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteBucketRegistration(deleteBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDeactivateBucket - Deactivate Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDeactivateBucket(deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions)`, func() {
			deleteDeactivateBucketOptions := &watsonxdatav2.DeleteDeactivateBucketOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDeactivateBucket(deleteDeactivateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDatabaseCatalog - Delete database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
			deleteDatabaseCatalogOptions := &watsonxdatav2.DeleteDatabaseCatalogOptions{
				DatabaseID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDb2Engine - Delete db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions)`, func() {
			deleteDb2EngineOptions := &watsonxdatav2.DeleteDb2EngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteNetezzaEngine - Delete netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions)`, func() {
			deleteNetezzaEngineOptions := &watsonxdatav2.DeleteNetezzaEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteOtherEngine - Delete engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions)`, func() {
			deleteOtherEngineOptions := &watsonxdatav2.DeleteOtherEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestissimoEngine - Delete prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestissimoEngine(deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions)`, func() {
			deletePrestissimoEngineOptions := &watsonxdatav2.DeletePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestissimoEngine(deletePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestissimoEngineCatalogs - Disassociate catalogs from a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions)`, func() {
			deletePrestissimoEngineCatalogsOptions := &watsonxdatav2.DeletePrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteEngine - Delete presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
			deleteEngineOptions := &watsonxdatav2.DeleteEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteEngine(deleteEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestoEngineCatalogs - Disassociate catalogs from a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions)`, func() {
			deletePrestoEngineCatalogsOptions := &watsonxdatav2.DeletePrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngine - Delete spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions)`, func() {
			deleteSparkEngineOptions := &watsonxdatav2.DeleteSparkEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineApplications - Stop Spark Applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions)`, func() {
			deleteSparkEngineApplicationsOptions := &watsonxdatav2.DeleteSparkEngineApplicationsOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSchema - Delete schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
			deleteSchemaOptions := &watsonxdatav2.DeleteSchemaOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
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
			deleteTableOptions := &watsonxdatav2.DeleteTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteTable(deleteTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteColumn - Delete column`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteColumn(deleteColumnOptions *DeleteColumnOptions)`, func() {
			deleteColumnOptions := &watsonxdatav2.DeleteColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				ColumnID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteColumn(deleteColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteMilvusService - Delete milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteMilvusService(deleteMilvusServiceOptions *DeleteMilvusServiceOptions)`, func() {
			deleteMilvusServiceOptions := &watsonxdatav2.DeleteMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteMilvusService(deleteMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
