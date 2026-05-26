//go:build integration

/**
 * (C) Copyright IBM Corp. 2026.
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

package watsonxdatav3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the watsonxdatav3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`WatsonxDataV3 Integration Tests`, func() {
	const externalConfigFile = "../watsonx_data_v3.env"

	var (
		err          error
		watsonxDataService *watsonxdatav3.WatsonxDataV3
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
			config, err = core.GetServiceProperties(watsonxdatav3.DefaultServiceName)
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
			watsonxDataServiceOptions := &watsonxdatav3.WatsonxDataV3Options{}

			watsonxDataService, err = watsonxdatav3.NewWatsonxDataV3UsingExternalConfig(watsonxDataServiceOptions)
			Expect(err).To(BeNil())
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(watsonxDataService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			watsonxDataService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateHdfsStorage - Add/Create HDFS storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateHdfsStorage(createHdfsStorageOptions *CreateHdfsStorageOptions)`, func() {
			createHdfsStorageOptions := &watsonxdatav3.CreateHdfsStorageOptions{
				DisplayName: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				HmsThriftURI: core.StringPtr("testString"),
				HmsThriftPort: core.Int64Ptr(int64(1)),
				CoreSite: core.StringPtr("testString"),
				HdfsSite: core.StringPtr("testString"),
				Kerberos: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				CatalogType: core.StringPtr("testString"),
				Krb5Config: core.StringPtr("testString"),
				HiveKeytab: CreateMockReader("This is a mock file."),
				HiveKeytabContentType: core.StringPtr("testString"),
				HdfsKeytab: CreateMockReader("This is a mock file."),
				HdfsKeytabContentType: core.StringPtr("testString"),
				HiveServerPrincipal: core.StringPtr("testString"),
				HiveClientPrincipal: core.StringPtr("testString"),
				HdfsPrincipal: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				CreatedAt: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			hdfsStorageRegistration, response, err := watsonxDataService.CreateHdfsStorage(createHdfsStorageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(hdfsStorageRegistration).ToNot(BeNil())
		})
	})

	Describe(`ListStorageRegistrations - Get list of storage registrations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListStorageRegistrations(listStorageRegistrationsOptions *ListStorageRegistrationsOptions)`, func() {
			listStorageRegistrationsOptions := &watsonxdatav3.ListStorageRegistrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			storageRegistrationCollection, response, err := watsonxDataService.ListStorageRegistrations(listStorageRegistrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageRegistrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateStorageRegistration - Register storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateStorageRegistration(createStorageRegistrationOptions *CreateStorageRegistrationOptions)`, func() {
			storageCatalogPrototypeModel := &watsonxdatav3.StorageCatalogPrototype{
				BasePath: core.StringPtr("/abc/def"),
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				CatalogType: core.StringPtr("iceberg"),
			}

			storageDetailsAccesskeyVaultModel := &watsonxdatav3.StorageDetailsAccesskeyVault{
				Key: core.StringPtr("secret key"),
				SecretName: core.StringPtr("secret name"),
				SecretUrn: core.StringPtr("secret urn"),
			}

			storageDetailsModel := &watsonxdatav3.StorageDetails{
				AccessKey: core.StringPtr("<access_key>"),
				AccessKeyVault: storageDetailsAccesskeyVaultModel,
				AccountName: core.StringPtr("sample-storage"),
				ApplicationID: core.StringPtr("application-id"),
				AuthMode: core.StringPtr("iam"),
				ContainerName: core.StringPtr("sample-container"),
				DirectoryID: core.StringPtr("directory-id"),
				Endpoint: core.StringPtr("https://s3.us-south.cloud-object-storage.appdomain.cloud/"),
				KeyFile: core.StringPtr("key_file"),
				Name: core.StringPtr("sample-storage"),
				Provider: core.StringPtr("ibm-cos"),
				Region: core.StringPtr("us-south"),
				RoleArn: core.StringPtr("arn:aws:iam::5ssdd5467-002c-a4f8cac3f3f9"),
				SasToken: core.StringPtr("<sas-token>"),
				SecretKey: core.StringPtr("secret_key"),
				SecretKeyVault: storageDetailsAccesskeyVaultModel,
				VaultEnabled: core.BoolPtr(true),
			}

			createStorageRegistrationOptions := &watsonxdatav3.CreateStorageRegistrationOptions{
				Description: core.StringPtr("COS storage for customer data"),
				DisplayName: core.StringPtr("sample-storage-displayname"),
				ManagedBy: core.StringPtr("ibm"),
				Type: core.StringPtr("ibm_cos"),
				AssociatedCatalog: storageCatalogPrototypeModel,
				Connection: storageDetailsModel,
				Region: core.StringPtr("us-south"),
				StorageUse: core.StringPtr("acl"),
				Tags: []string{"storage-tag1", "storage-tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			storageRegistration, response, err := watsonxDataService.CreateStorageRegistration(createStorageRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(storageRegistration).ToNot(BeNil())
		})
	})

	Describe(`GetStorageRegistration - Get storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetStorageRegistration(getStorageRegistrationOptions *GetStorageRegistrationOptions)`, func() {
			getStorageRegistrationOptions := &watsonxdatav3.GetStorageRegistrationOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				SkipMdsCall: core.BoolPtr(false),
			}

			storageRegistration, response, err := watsonxDataService.GetStorageRegistration(getStorageRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageRegistration).ToNot(BeNil())
		})
	})

	Describe(`UpdateStorageRegistration - Update storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateStorageRegistration(updateStorageRegistrationOptions *UpdateStorageRegistrationOptions)`, func() {
			storageDetailsAccesskeyVaultModel := &watsonxdatav3.StorageDetailsAccesskeyVault{
				Key: core.StringPtr("secret key"),
				SecretName: core.StringPtr("secret name"),
				SecretUrn: core.StringPtr("secret urn"),
			}

			storageDetailsModel := &watsonxdatav3.StorageDetails{
				AccessKey: core.StringPtr("<access_key>"),
				AccessKeyVault: storageDetailsAccesskeyVaultModel,
				AccountName: core.StringPtr("sample-storage"),
				ApplicationID: core.StringPtr("application-id"),
				AuthMode: core.StringPtr("iam"),
				ContainerName: core.StringPtr("sample-container"),
				DirectoryID: core.StringPtr("directory-id"),
				Endpoint: core.StringPtr("https://s3.us-south.cloud-object-storage.appdomain.cloud/"),
				KeyFile: core.StringPtr("key_file"),
				Name: core.StringPtr("sample-storage"),
				Provider: core.StringPtr("ibm-cos"),
				Region: core.StringPtr("us-south"),
				RoleArn: core.StringPtr("arn:aws:iam::5ssdd5467-002c-a4f8cac3f3f9"),
				SasToken: core.StringPtr("<sas-token>"),
				SecretKey: core.StringPtr("secret_key"),
				SecretKeyVault: storageDetailsAccesskeyVaultModel,
				VaultEnabled: core.BoolPtr(true),
			}

			storageRegistrationPatchModel := &watsonxdatav3.StorageRegistrationPatch{
				Connection: storageDetailsModel,
				Description: core.StringPtr("COS storage for customer data"),
				DisplayName: core.StringPtr("sample-storage-displayname"),
				SystemStorageUpdateCredentials: core.BoolPtr(true),
				Tags: []string{"teststorage", "userstorage"},
			}
			storageRegistrationPatchModelAsPatch, asPatchErr := storageRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateStorageRegistrationOptions := &watsonxdatav3.UpdateStorageRegistrationOptions{
				ID: core.StringPtr("testString"),
				Body: storageRegistrationPatchModelAsPatch,
				SkipMdsCall: core.BoolPtr(false),
				AuthInstanceID: core.StringPtr("testString"),
			}

			storageRegistration, response, err := watsonxDataService.UpdateStorageRegistration(updateStorageRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageRegistration).ToNot(BeNil())
		})
	})

	Describe(`AddStorageCatalog - Add storage catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddStorageCatalog(addStorageCatalogOptions *AddStorageCatalogOptions)`, func() {
			addStorageCatalogOptions := &watsonxdatav3.AddStorageCatalogOptions{
				StorageID: core.StringPtr("testString"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				BasePath: core.StringPtr("/abc/def"),
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogType: core.StringPtr("iceberg"),
				SkipMdsCall: core.BoolPtr(false),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.AddStorageCatalog(addStorageCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetStorageObjectProperties - Get storage object properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetStorageObjectProperties(getStorageObjectPropertiesOptions *GetStorageObjectPropertiesOptions)`, func() {
			pathModel := &watsonxdatav3.Path{
				Path: core.StringPtr("testString"),
			}

			getStorageObjectPropertiesOptions := &watsonxdatav3.GetStorageObjectPropertiesOptions{
				StorageID: core.StringPtr("testString"),
				Paths: []watsonxdatav3.Path{*pathModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			storageObjectProperties, response, err := watsonxDataService.GetStorageObjectProperties(getStorageObjectPropertiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(storageObjectProperties).ToNot(BeNil())
		})
	})

	Describe(`ListStorageRegistrationsObjects - List storage objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListStorageRegistrationsObjects(listStorageRegistrationsObjectsOptions *ListStorageRegistrationsObjectsOptions)`, func() {
			listStorageRegistrationsObjectsOptions := &watsonxdatav3.ListStorageRegistrationsObjectsOptions{
				StorageID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				Path: core.StringPtr("testString"),
				Paginated: core.BoolPtr(true),
				PageSize: core.Int64Ptr(int64(1)),
				Prefix: core.StringPtr("testString"),
				StartAfter: core.StringPtr("testString"),
			}

			storageRegistrationObjectCollection, response, err := watsonxDataService.ListStorageRegistrationsObjects(listStorageRegistrationsObjectsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageRegistrationObjectCollection).ToNot(BeNil())
		})
	})

	Describe(`ListDatabaseRegistrations - Get list of databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions)`, func() {
			listDatabaseRegistrationsOptions := &watsonxdatav3.ListDatabaseRegistrationsOptions{
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
			databaseCatalogPrototypeModel := &watsonxdatav3.DatabaseCatalogPrototype{
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogType: core.StringPtr("iceberg"),
			}

			secretDetailsModel := &watsonxdatav3.SecretDetails{
				Key: core.StringPtr("secret key"),
				SecretName: core.StringPtr("secret name"),
				SecretUrn: core.StringPtr("secret urn"),
			}

			databaseDetailsPrototypeModel := &watsonxdatav3.DatabaseDetailsPrototype{
				AuthenticationType: core.StringPtr("LDAP, NTLM, ActiveDirectoryServicePrincipal, ActiveDirectoryPassword"),
				AuthenticationValue: core.StringPtr("LDAP"),
				AuthenticationValueKeyVault: secretDetailsModel,
				BrokerAuthenticationPassword: core.StringPtr("samplepassword"),
				BrokerAuthenticationType: core.StringPtr("PASSWORD"),
				BrokerAuthenticationUser: core.StringPtr("sampleuser"),
				BrokerHost: core.StringPtr("samplehost"),
				BrokerPort: core.Int64Ptr(int64(4553)),
				Certificate: core.StringPtr("exampleCertificate"),
				CertificateExtension: core.StringPtr("pem"),
				ConnectionMethod: core.StringPtr("basic, apikey"),
				ConnectionMode: core.StringPtr("service_name"),
				ConnectionModeValue: core.StringPtr("orclpdb"),
				ConnectionType: core.StringPtr("JDBC, Arrow flight"),
				ControllerAuthenticationPassword: core.StringPtr("samplepassword"),
				ControllerAuthenticationType: core.StringPtr("PASSWORD"),
				ControllerAuthenticationUser: core.StringPtr("sampleuser"),
				CoordinatorHost: core.StringPtr("samplehost"),
				CoordinatorPort: core.Int64Ptr(int64(4553)),
				CpdHostname: core.StringPtr("samplecpdhostname"),
				CredentialsKey: core.StringPtr("eyJ0eXBlIjoic2VydmljZV9hY2NvdW50IiwicHJvamVjdF9pZCI6ImNvbm9wcy1iaWdxdWVyeSIsInByaXZhdGVfa2V5X2lkIjoiMGY3......"),
				DomainName: core.StringPtr("conops-mssql.conops.local"),
				Hostname: core.StringPtr("http://db2@localhost:9900.com"),
				HostnameInCertificate: core.StringPtr("samplehostname"),
				Hosts: core.StringPtr("abc.com:1234,xyz.com:4321"),
				InformixServer: core.StringPtr("ol_informix1410"),
				Name: core.StringPtr("new_database"),
				Password: core.StringPtr("samplepassword"),
				PasswordKeyVault: secretDetailsModel,
				Port: core.Int64Ptr(int64(4553)),
				ProjectID: core.StringPtr("conops-bigquery"),
				Sasl: core.BoolPtr(true),
				SaslMechanism: core.StringPtr("plain"),
				SchemaName: core.StringPtr("sampleSchema"),
				Schemas: core.StringPtr("redis__name"),
				ServiceApiKey: core.StringPtr("sampleapikey"),
				ServiceHostname: core.StringPtr("api.dataplatform.dev.cloud.ibm.com"),
				ServicePassword: core.StringPtr("samplepassword"),
				ServicePort: core.Int64Ptr(int64(443)),
				ServiceSsl: core.BoolPtr(true),
				ServiceTokenURL: core.StringPtr("sampletoakenurl"),
				ServiceUsername: core.StringPtr("sampleusername"),
				Ssl: core.BoolPtr(true),
				SslcertificateKeyVault: secretDetailsModel,
				Tables: core.StringPtr("kafka_table_name, redis_table_name"),
				Username: core.StringPtr("sampleuser"),
				UsernameKeyVault: secretDetailsModel,
				ValidateServerCertificate: core.BoolPtr(true),
				VaultEnabled: core.BoolPtr(true),
				VerifyHostName: core.BoolPtr(true),
				WarehouseName: core.StringPtr("samplewrehouse"),
			}

			databaseRegistrationPrototypeDatabasePropertiesItemsModel := &watsonxdatav3.DatabaseRegistrationPrototypeDatabasePropertiesItems{
				Encrypt: core.BoolPtr(true),
				Key: core.StringPtr("abc"),
				Value: core.StringPtr("xyz"),
			}

			createDatabaseRegistrationOptions := &watsonxdatav3.CreateDatabaseRegistrationOptions{
				DisplayName: core.StringPtr("new_database"),
				Type: core.StringPtr("db2"),
				AssociatedCatalog: databaseCatalogPrototypeModel,
				Connection: databaseDetailsPrototypeModel,
				CreatedAt: core.StringPtr("1686792721"),
				Description: core.StringPtr("db2 extenal database description"),
				Properties: []watsonxdatav3.DatabaseRegistrationPrototypeDatabasePropertiesItems{*databaseRegistrationPrototypeDatabasePropertiesItemsModel},
				SourceAssetID: core.StringPtr("cc85d899-9ec3-496a-be36-99cabc62f123"),
				SourceCatalogID: core.StringPtr("cc85d899-9ec3-496a-be36-99cff962f000"),
				SourceProjectID: core.StringPtr("cc85d899-9ec3-496a-be36-99cff9000116"),
				Tags: []string{"testdatabase", "userdatabase"},
				TargetCatalogID: core.StringPtr("cc85d899-9ec3-496a-be36-99cff9000116"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`AddDatabaseCatalog - Post database catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddDatabaseCatalog(addDatabaseCatalogOptions *AddDatabaseCatalogOptions)`, func() {
			addDatabaseCatalogOptions := &watsonxdatav3.AddDatabaseCatalogOptions{
				DatabaseID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("sampleCatalog"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.AddDatabaseCatalog(addDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDatabase - Get database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabase(getDatabaseOptions *GetDatabaseOptions)`, func() {
			getDatabaseOptions := &watsonxdatav3.GetDatabaseOptions{
				ID: core.StringPtr("testString"),
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
			databaseRegistrationPatchDatabaseDetailsDatabasePropertiesItemsModel := &watsonxdatav3.DatabaseRegistrationPatchDatabaseDetailsDatabasePropertiesItems{
				Encrypt: core.BoolPtr(true),
				Key: core.StringPtr("abc"),
				Value: core.StringPtr("xyz"),
			}

			databaseRegistrationPatchDatabaseDetailsModel := &watsonxdatav3.DatabaseRegistrationPatchDatabaseDetails{
				AuthenticationValue: core.StringPtr("LDAP"),
				BrokerAuthenticationPassword: core.StringPtr("samplepassword"),
				BrokerAuthenticationType: core.StringPtr("PASSWORD"),
				BrokerAuthenticationUser: core.StringPtr("sampleuser"),
				ControllerAuthenticationPassword: core.StringPtr("samplepassword"),
				ControllerAuthenticationType: core.StringPtr("PASSWORD"),
				ControllerAuthenticationUser: core.StringPtr("sampleuser"),
				CredentialsKey: core.StringPtr("eyJ0eXBlIjoic2VydmljZV9hY2NvdW50IiwicHJvamVjdF9pZCI6ImNvbm9wcy1iaWdxdWVyeSIsInByaXZhdGVfa2V5X2lkIjoiMGY3......"),
				Password: core.StringPtr("samplepassword"),
				Properties: []watsonxdatav3.DatabaseRegistrationPatchDatabaseDetailsDatabasePropertiesItems{*databaseRegistrationPatchDatabaseDetailsDatabasePropertiesItemsModel},
				Username: core.StringPtr("sampleuser"),
			}

			databaseRegistrationPatchTablesItemsModel := &watsonxdatav3.DatabaseRegistrationPatchTablesItems{
				CreatedAt: core.StringPtr("1686792721"),
				FileContents: core.StringPtr("sample file content"),
				FileName: core.StringPtr("test.json"),
				SchemaName: core.StringPtr("customer"),
				TableName: core.StringPtr("customer"),
			}

			databaseRegistrationPatchTopicsItemsModel := &watsonxdatav3.DatabaseRegistrationPatchTopicsItems{
				CreatedAt: core.StringPtr("1686792721"),
				FileContents: core.StringPtr("sample file contents"),
				FileName: core.StringPtr("test.json"),
				TopicName: core.StringPtr("customer"),
			}

			databaseRegistrationPatchModel := &watsonxdatav3.DatabaseRegistrationPatch{
				Connection: databaseRegistrationPatchDatabaseDetailsModel,
				Description: core.StringPtr("External database description"),
				DisplayName: core.StringPtr("new_database"),
				Tables: []watsonxdatav3.DatabaseRegistrationPatchTablesItems{*databaseRegistrationPatchTablesItemsModel},
				Tags: []string{"testdatabase", "userdatabase"},
				Topics: []watsonxdatav3.DatabaseRegistrationPatchTopicsItems{*databaseRegistrationPatchTopicsItemsModel},
			}
			databaseRegistrationPatchModelAsPatch, asPatchErr := databaseRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDatabaseOptions := &watsonxdatav3.UpdateDatabaseOptions{
				ID: core.StringPtr("testString"),
				Body: databaseRegistrationPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.UpdateDatabase(updateDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngines - Get list of Presto(Java) engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions)`, func() {
			listPrestoEnginesOptions := &watsonxdatav3.ListPrestoEnginesOptions{
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
			coordinatorNodeDescriptionBodyModel := &watsonxdatav3.CoordinatorNodeDescriptionBody{
				NodeType: core.StringPtr("starter"),
				Quantity: core.Int64Ptr(int64(1)),
			}

			workerNodeDescriptionBodyModel := &watsonxdatav3.WorkerNodeDescriptionBody{
				NodeType: core.StringPtr("starter"),
				Quantity: core.Int64Ptr(int64(1)),
			}

			autoscalingConfigModel := &watsonxdatav3.AutoscalingConfig{
				Type: core.StringPtr("cpu"),
				Target: core.Int64Ptr(int64(40)),
				MinWorkerQuantity: core.Int64Ptr(int64(1)),
				MaxWorkerQuantity: core.Int64Ptr(int64(18)),
				QueryTerminationGracePeriodMin: core.Int64Ptr(int64(1)),
				ScaleInStabilizationWindowMin: core.Int64Ptr(int64(5)),
				ScalingStepSize: core.Int64Ptr(int64(1)),
			}

			engineDetailsModel := &watsonxdatav3.EngineDetails{
				Coordinator: coordinatorNodeDescriptionBodyModel,
				SizeConfig: core.StringPtr("starter"),
				Worker: workerNodeDescriptionBodyModel,
				AutoscalingEnabled: core.BoolPtr(true),
				AutoscalingConfig: autoscalingConfigModel,
			}

			createPrestoEngineOptions := &watsonxdatav3.CreatePrestoEngineOptions{
				Configuration: engineDetailsModel,
				DisplayName: core.StringPtr("sampleEngine"),
				Origin: core.StringPtr("native"),
				AssociatedCatalogs: []string{"iceberg_data", "hive_data"},
				Description: core.StringPtr("presto engine for running sql queries"),
				ID: core.StringPtr("presto123"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.CreatePrestoEngine(createPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestoEngineAutoscaling - Update Autoscaling Configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestoEngineAutoscaling(updatePrestoEngineAutoscalingOptions *UpdatePrestoEngineAutoscalingOptions)`, func() {
			autoScalingConfigModel := &watsonxdatav3.AutoScalingConfig{
				Target: core.Int64Ptr(int64(40)),
				MinWorkerQuantity: core.Int64Ptr(int64(1)),
				MaxWorkerQuantity: core.Int64Ptr(int64(1)),
				QueryTerminationGracePeriodMin: core.Int64Ptr(int64(1)),
				ScaleInStabilizationWindowMin: core.Int64Ptr(int64(5)),
				ScalingStepSize: core.Int64Ptr(int64(1)),
			}

			updateAutoScalingRequestModel := &watsonxdatav3.UpdateAutoScalingRequest{
				AutoscalingConfig: autoScalingConfigModel,
			}
			updateAutoScalingRequestModelAsPatch, asPatchErr := updateAutoScalingRequestModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestoEngineAutoscalingOptions := &watsonxdatav3.UpdatePrestoEngineAutoscalingOptions{
				EngineID: core.StringPtr("testString"),
				UpdateAutoScalingRequestPatch: updateAutoScalingRequestModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			patchAutoscalingSuccess, response, err := watsonxDataService.UpdatePrestoEngineAutoscaling(updatePrestoEngineAutoscalingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(patchAutoscalingSuccess).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngineCatalogs - Get presto engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions)`, func() {
			listPrestoEngineCatalogsOptions := &watsonxdatav3.ListPrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestoEngineCatalogs - Associate catalogs to presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestoEngineCatalogs(createPrestoEngineCatalogsOptions *CreatePrestoEngineCatalogsOptions)`, func() {
			createPrestoEngineCatalogsOptions := &watsonxdatav3.CreatePrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: []string{"iceberg_catalog"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.CreatePrestoEngineCatalogs(createPrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngineCatalog - Get presto engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions)`, func() {
			getPrestoEngineCatalogOptions := &watsonxdatav3.GetPrestoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngineConfig - Get Presto Engine Configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngineConfig(getPrestoEngineConfigOptions *GetPrestoEngineConfigOptions)`, func() {
			getPrestoEngineConfigOptions := &watsonxdatav3.GetPrestoEngineConfigOptions{
				EngineID: core.StringPtr("presto860"),
				AuthInstanceID: core.StringPtr("crn:v1:staging:public:lakehouse:eu-de:a/810fe64a9e3446d1b919d0eff69d3c5f:87495456-c862-4b35-b8de-884b7f15eee4::"),
				Sections: core.StringPtr("catalog,configuration,jvm"),
			}

			prestoEnginePropertiesDetails, response, err := watsonxDataService.GetPrestoEngineConfig(getPrestoEngineConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEnginePropertiesDetails).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestoEngineConfig - Update Presto Engine Configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestoEngineConfig(updatePrestoEngineConfigOptions *UpdatePrestoEngineConfigOptions)`, func() {
			catalogPropertiesModel := &watsonxdatav3.CatalogProperties{
				Coordinator: map[string]string{"key1": "testString"},
				Worker: map[string]string{"key1": "testString"},
			}

			configurationPropertiesModel := &watsonxdatav3.ConfigurationProperties{
				Coordinator: map[string]string{"key1": "configuration_property_value"},
				Worker: map[string]string{"key1": "configuration_property_value"},
			}

			jvmPropertiesModel := &watsonxdatav3.JvmProperties{
				Coordinator: map[string]string{"key1": "JVM_property_value"},
				Worker: map[string]string{"key1": "JVM_property_value"},
			}

			logConfigPropertiesModel := &watsonxdatav3.LogConfigProperties{
				Coordinator: map[string]string{"key1": "log_config_property_value"},
				Worker: map[string]string{"key1": "log_config_property_value"},
			}

			prestoEnginePropertiesModel := &watsonxdatav3.PrestoEngineProperties{
				Catalog: map[string]watsonxdatav3.CatalogProperties{"key1": *catalogPropertiesModel},
				Configuration: configurationPropertiesModel,
				EventListener: map[string]string{"key1": "event_listener_property_value"},
				Global: map[string]string{"key1": "global_property_value"},
				JmxExporterConfig: map[string]string{"key1": "jmx_exporter_config_property_value"},
				Jvm: jvmPropertiesModel,
				LogConfig: logConfigPropertiesModel,
			}
			prestoEnginePropertiesModel.Catalog["foo"] = *catalogPropertiesModel

			prestoEnginePropertiesDetailsModel := &watsonxdatav3.PrestoEnginePropertiesDetails{
				EngineProperties: prestoEnginePropertiesModel,
			}
			prestoEnginePropertiesDetailsModelAsPatch, asPatchErr := prestoEnginePropertiesDetailsModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestoEngineConfigOptions := &watsonxdatav3.UpdatePrestoEngineConfigOptions{
				EngineID: core.StringPtr("presto93"),
				AuthInstanceID: core.StringPtr("crn:v1:staging:public:lakehouse:eu-de:a/810fe64a9e3446d1b919d0eff69d3c5f:87495456-c862-4b35-b8de-884b7f15eee4::"),
				PrestoEnginePropertiesDetailsPatch: prestoEnginePropertiesDetailsModelAsPatch,
			}

			response, err := watsonxDataService.UpdatePrestoEngineConfig(updatePrestoEngineConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetPrestoEngine - Get presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions)`, func() {
			getPrestoEngineOptions := &watsonxdatav3.GetPrestoEngineOptions{
				ID: core.StringPtr("testString"),
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
			enginePropertiesCatalogAdditionalPropertiesModel := &watsonxdatav3.EnginePropertiesCatalogAdditionalProperties{
				Coordinator: map[string]interface{}{"anyKey": "anyValue"},
				Worker: map[string]interface{}{"anyKey": "anyValue"},
			}

			enginePropertiesConfigurationModel := &watsonxdatav3.EnginePropertiesConfiguration{
				Coordinator: map[string]string{"key1": "configuration_property_value"},
				Worker: map[string]string{"key1": "configuration_property_value"},
			}

			enginePropertiesJvmModel := &watsonxdatav3.EnginePropertiesJvm{
				Coordinator: map[string]string{"key1": "JVM_property_value"},
				Worker: map[string]string{"key1": "JVM_property_value"},
			}

			enginePropertiesLogConfigModel := &watsonxdatav3.EnginePropertiesLogConfig{
				Coordinator: map[string]string{"key1": "testString"},
				Worker: map[string]string{"key1": "log_config_property_value"},
			}

			enginePropertiesModel := &watsonxdatav3.EngineProperties{
				Catalog: map[string]watsonxdatav3.EnginePropertiesCatalogAdditionalProperties{"key1": *enginePropertiesCatalogAdditionalPropertiesModel},
				Configuration: enginePropertiesConfigurationModel,
				EventListener: map[string]string{"key1": "event_listener_property_value"},
				Global: map[string]string{"key1": "global_property_value"},
				JmxExporterConfig: map[string]string{"key1": "jmx_exporter_config_property_value"},
				Jvm: enginePropertiesJvmModel,
				LogConfig: enginePropertiesLogConfigModel,
			}
			enginePropertiesModel.Catalog["foo"] = *enginePropertiesCatalogAdditionalPropertiesModel

			removeEnginePropertiesCatalogAdditionalPropertiesModel := &watsonxdatav3.RemoveEnginePropertiesCatalogAdditionalProperties{
				Coordinator: []string{"property_name"},
				Worker: []string{"property_name"},
			}

			removeEnginePropertiesConfigurationModel := &watsonxdatav3.RemoveEnginePropertiesConfiguration{
				Coordinator: []string{"property_name"},
				Worker: []string{"property_name"},
			}

			removeEnginePropertiesJvmModel := &watsonxdatav3.RemoveEnginePropertiesJvm{
				Coordinator: []string{"property_name"},
				Worker: []string{"property_name"},
			}

			removeEnginePropertiesLogConfigModel := &watsonxdatav3.RemoveEnginePropertiesLogConfig{
				Coordinator: []string{"property_name"},
				Worker: []string{"property_name"},
			}

			removeEnginePropertiesModel := &watsonxdatav3.RemoveEngineProperties{
				Catalog: map[string]watsonxdatav3.RemoveEnginePropertiesCatalogAdditionalProperties{"key1": *removeEnginePropertiesCatalogAdditionalPropertiesModel},
				Configuration: removeEnginePropertiesConfigurationModel,
				EventListener: []string{"property_name"},
				Global: []string{"property_name"},
				JmxExporterConfig: []string{"testString"},
				Jvm: removeEnginePropertiesJvmModel,
				LogConfig: removeEnginePropertiesLogConfigModel,
			}
			removeEnginePropertiesModel.Catalog["foo"] = *removeEnginePropertiesCatalogAdditionalPropertiesModel

			prestoEnginePatchModel := &watsonxdatav3.PrestoEnginePatch{
				Description: core.StringPtr("updated description for presto engine"),
				DisplayName: core.StringPtr("sampleEngine"),
				Properties: enginePropertiesModel,
				RemoveEngineProperties: removeEnginePropertiesModel,
				RestartType: core.StringPtr("force"),
				Tags: []string{"tag1", "tag2"},
			}
			prestoEnginePatchModelAsPatch, asPatchErr := prestoEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestoEngineOptions := &watsonxdatav3.UpdatePrestoEngineOptions{
				ID: core.StringPtr("testString"),
				Body: prestoEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.UpdatePrestoEngine(updatePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`PausePrestoEngine - Pause presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PausePrestoEngine(pausePrestoEngineOptions *PausePrestoEngineOptions)`, func() {
			pausePrestoEngineOptions := &watsonxdatav3.PausePrestoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineSuccessResponse, response, err := watsonxDataService.PausePrestoEngine(pausePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngineSuccessResponse).ToNot(BeNil())
		})
	})

	Describe(`RunExplainStatement - Explain presto query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions)`, func() {
			runExplainStatementOptions := &watsonxdatav3.RunExplainStatementOptions{
				ID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Catalog: core.StringPtr("catalog_name"),
				Format: core.StringPtr("json"),
				Schema: core.StringPtr("schema_name"),
				Type: core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoQueryExplain, response, err := watsonxDataService.RunExplainStatement(runExplainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoQueryExplain).ToNot(BeNil())
		})
	})

	Describe(`RunExplainAnalyzeStatement - Explain presto analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions)`, func() {
			runExplainAnalyzeStatementOptions := &watsonxdatav3.RunExplainAnalyzeStatementOptions{
				ID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Verbose: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoQueryExplain, response, err := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoQueryExplain).ToNot(BeNil())
		})
	})

	Describe(`RestartPrestoEngine - Restart a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestartPrestoEngine(restartPrestoEngineOptions *RestartPrestoEngineOptions)`, func() {
			restartPrestoEngineOptions := &watsonxdatav3.RestartPrestoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineSuccessResponse, response, err := watsonxDataService.RestartPrestoEngine(restartPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngineSuccessResponse).ToNot(BeNil())
		})
	})

	Describe(`ResumePrestoEngine - Resume Presto(Java) engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumePrestoEngine(resumePrestoEngineOptions *ResumePrestoEngineOptions)`, func() {
			resumePrestoEngineOptions := &watsonxdatav3.ResumePrestoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineSuccessResponse, response, err := watsonxDataService.ResumePrestoEngine(resumePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngineSuccessResponse).ToNot(BeNil())
		})
	})

	Describe(`ScalePrestoEngine - Scale a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScalePrestoEngine(scalePrestoEngineOptions *ScalePrestoEngineOptions)`, func() {
			nodeDescriptionModel := &watsonxdatav3.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(1)),
			}

			scalePrestoEngineOptions := &watsonxdatav3.ScalePrestoEngineOptions{
				ID: core.StringPtr("testString"),
				Coordinator: nodeDescriptionModel,
				Worker: nodeDescriptionModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineSuccessResponse, response, err := watsonxDataService.ScalePrestoEngine(scalePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(prestoEngineSuccessResponse).ToNot(BeNil())
		})
	})

	Describe(`ListPrestissimoEngines - Get list of prestissimo engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestissimoEngines(listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions)`, func() {
			listPrestissimoEnginesOptions := &watsonxdatav3.ListPrestissimoEnginesOptions{
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
			nodeDescriptionModel := &watsonxdatav3.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(1)),
			}

			prestissimoEngineDetailsModel := &watsonxdatav3.PrestissimoEngineDetails{
				Coordinator: nodeDescriptionModel,
				SizeConfig: core.StringPtr("starter"),
				Worker: nodeDescriptionModel,
			}

			createPrestissimoEngineOptions := &watsonxdatav3.CreatePrestissimoEngineOptions{
				Configuration: prestissimoEngineDetailsModel,
				DisplayName: core.StringPtr("sampleEngine"),
				Origin: core.StringPtr("native"),
				AssociatedCatalogs: []string{"hive_data"},
				Description: core.StringPtr("prestissimo engine description"),
				ID: core.StringPtr("prestissimo123"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.CreatePrestissimoEngine(createPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngineCatalog - Get prestissimo engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions)`, func() {
			getPrestissimoEngineCatalogOptions := &watsonxdatav3.GetPrestissimoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngine - Get prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngine(getPrestissimoEngineOptions *GetPrestissimoEngineOptions)`, func() {
			getPrestissimoEngineOptions := &watsonxdatav3.GetPrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
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
			prestissimoEnginePropertiesCatalogAdditionalPropertiesModel := &watsonxdatav3.PrestissimoEnginePropertiesCatalogAdditionalProperties{
				Coordinator: map[string]interface{}{"anyKey": "anyValue"},
				Worker: map[string]interface{}{"anyKey": "anyValue"},
			}

			prestissimoEnginePropertiesConfigurationModel := &watsonxdatav3.PrestissimoEnginePropertiesConfiguration{
				Coordinator: map[string]string{"key1": "testString"},
				Worker: map[string]string{"key1": "testString"},
			}

			prestissimoEnginePropertiesJvmModel := &watsonxdatav3.PrestissimoEnginePropertiesJvm{
				Coordinator: map[string]string{"key1": "testString"},
			}

			prestissimoEnginePropertiesLogConfigModel := &watsonxdatav3.PrestissimoEnginePropertiesLogConfig{
				Coordinator: map[string]string{"key1": "testString"},
				Worker: map[string]string{"key1": "testString"},
			}

			prestissimoEnginePropertiesModel := &watsonxdatav3.PrestissimoEngineProperties{
				Catalog: map[string]watsonxdatav3.PrestissimoEnginePropertiesCatalogAdditionalProperties{"key1": *prestissimoEnginePropertiesCatalogAdditionalPropertiesModel},
				Configuration: prestissimoEnginePropertiesConfigurationModel,
				Global: map[string]string{"key1": "testString"},
				Jvm: prestissimoEnginePropertiesJvmModel,
				LogConfig: prestissimoEnginePropertiesLogConfigModel,
				OptimizerProperties: map[string]string{"key1": "testString"},
				Velox: map[string]string{"key1": "testString"},
			}
			prestissimoEnginePropertiesModel.Catalog["foo"] = *prestissimoEnginePropertiesCatalogAdditionalPropertiesModel

			removePrestissimoEnginePropertiesCatalogAdditionalPropertiesModel := &watsonxdatav3.RemovePrestissimoEnginePropertiesCatalogAdditionalProperties{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			removePrestissimoEnginePropertiesConfigurationModel := &watsonxdatav3.RemovePrestissimoEnginePropertiesConfiguration{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			removePrestissimoEnginePropertiesJvmModel := &watsonxdatav3.RemovePrestissimoEnginePropertiesJvm{
				Coordinator: []string{"testString"},
			}

			removePrestissimoEnginePropertiesLogConfigModel := &watsonxdatav3.RemovePrestissimoEnginePropertiesLogConfig{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			removePrestissimoEnginePropertiesModel := &watsonxdatav3.RemovePrestissimoEngineProperties{
				Catalog: map[string]watsonxdatav3.RemovePrestissimoEnginePropertiesCatalogAdditionalProperties{"key1": *removePrestissimoEnginePropertiesCatalogAdditionalPropertiesModel},
				Configuration: removePrestissimoEnginePropertiesConfigurationModel,
				Global: []string{"testString"},
				Jvm: removePrestissimoEnginePropertiesJvmModel,
				LogConfig: removePrestissimoEnginePropertiesLogConfigModel,
				OptimizerProperties: []string{"testString"},
				Velox: []string{"testString"},
			}
			removePrestissimoEnginePropertiesModel.Catalog["foo"] = *removePrestissimoEnginePropertiesCatalogAdditionalPropertiesModel

			prestissimoEnginePatchModel := &watsonxdatav3.PrestissimoEnginePatch{
				Description: core.StringPtr("updated description for prestissimo engine"),
				DisplayName: core.StringPtr("sampleEngine"),
				Properties: prestissimoEnginePropertiesModel,
				RemoveEngineProperties: removePrestissimoEnginePropertiesModel,
				RestartType: core.StringPtr("force"),
				Tags: []string{"tag1", "tag2"},
			}
			prestissimoEnginePatchModelAsPatch, asPatchErr := prestissimoEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestissimoEngineOptions := &watsonxdatav3.UpdatePrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
				Body: prestissimoEnginePatchModelAsPatch,
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
			listPrestissimoEngineCatalogsOptions := &watsonxdatav3.ListPrestissimoEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngineCatalogs - Associate catalogs to a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngineCatalogs(createPrestissimoEngineCatalogsOptions *CreatePrestissimoEngineCatalogsOptions)`, func() {
			createPrestissimoEngineCatalogsOptions := &watsonxdatav3.CreatePrestissimoEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				CatalogNames: []string{"iceberg_catalog"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.CreatePrestissimoEngineCatalogs(createPrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`PausePrestissimoEngine - Pause prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PausePrestissimoEngine(pausePrestissimoEngineOptions *PausePrestissimoEngineOptions)`, func() {
			pausePrestissimoEngineOptions := &watsonxdatav3.PausePrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.PausePrestissimoEngine(pausePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`RunPrestissimoExplainStatement - Explain query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions)`, func() {
			runPrestissimoExplainStatementOptions := &watsonxdatav3.RunPrestissimoExplainStatementOptions{
				ID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Catalog: core.StringPtr("catalog_name"),
				Format: core.StringPtr("json"),
				Schema: core.StringPtr("schema_name"),
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
			runPrestissimoExplainAnalyzeStatementOptions := &watsonxdatav3.RunPrestissimoExplainAnalyzeStatementOptions{
				ID: core.StringPtr("testString"),
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

	Describe(`RestartPrestissimoEngine - Restart a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestartPrestissimoEngine(restartPrestissimoEngineOptions *RestartPrestissimoEngineOptions)`, func() {
			restartPrestissimoEngineOptions := &watsonxdatav3.RestartPrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.RestartPrestissimoEngine(restartPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ResumePrestissimoEngine - Resume prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumePrestissimoEngine(resumePrestissimoEngineOptions *ResumePrestissimoEngineOptions)`, func() {
			resumePrestissimoEngineOptions := &watsonxdatav3.ResumePrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ResumePrestissimoEngine(resumePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ScalePrestissimoEngine - Scale a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScalePrestissimoEngine(scalePrestissimoEngineOptions *ScalePrestissimoEngineOptions)`, func() {
			nodeDescriptionModel := &watsonxdatav3.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(1)),
			}

			scalePrestissimoEngineOptions := &watsonxdatav3.ScalePrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
				Coordinator: nodeDescriptionModel,
				Worker: nodeDescriptionModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ScalePrestissimoEngine(scalePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListDb2Engines - Get list of db2 engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions)`, func() {
			listDb2EnginesOptions := &watsonxdatav3.ListDb2EnginesOptions{
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
			db2EngineDetailsBodyModel := &watsonxdatav3.Db2EngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createDb2EngineOptions := &watsonxdatav3.CreateDb2EngineOptions{
				Configuration: db2EngineDetailsBodyModel,
				DisplayName: core.StringPtr("sampleEngine"),
				Origin: core.StringPtr("external"),
				Description: core.StringPtr("db2 engine description"),
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
			db2EnginePatchModel := &watsonxdatav3.Db2EnginePatch{
				Description: core.StringPtr("db2 engine updated description"),
				DisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
			}
			db2EnginePatchModelAsPatch, asPatchErr := db2EnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDb2EngineOptions := &watsonxdatav3.UpdateDb2EngineOptions{
				ID: core.StringPtr("testString"),
				Body: db2EnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2Engine, response, err := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(db2Engine).ToNot(BeNil())
		})
	})

	Describe(`ListOtherEngines - Get list of other engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions)`, func() {
			listOtherEnginesOptions := &watsonxdatav3.ListOtherEnginesOptions{
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
			otherEngineConfigurationBodyModel := &watsonxdatav3.OtherEngineConfigurationBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
				Type: core.StringPtr("netezza"),
			}

			createOtherEngineOptions := &watsonxdatav3.CreateOtherEngineOptions{
				Configuration: otherEngineConfigurationBodyModel,
				DisplayName: core.StringPtr("sampleEngine01"),
				Origin: core.StringPtr("external"),
				Description: core.StringPtr("external engine description"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			otherEngine, response, err := watsonxDataService.CreateOtherEngine(createOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(otherEngine).ToNot(BeNil())
		})
	})

	Describe(`ListNetezzaEngines - Get list of Netezza engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions)`, func() {
			listNetezzaEnginesOptions := &watsonxdatav3.ListNetezzaEnginesOptions{
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
			netezzaEngineConfigurationBodyModel := &watsonxdatav3.NetezzaEngineConfigurationBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createNetezzaEngineOptions := &watsonxdatav3.CreateNetezzaEngineOptions{
				Configuration: netezzaEngineConfigurationBodyModel,
				DisplayName: core.StringPtr("sampleEngine"),
				Origin: core.StringPtr("external"),
				Description: core.StringPtr("netezza engine description"),
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
			netezzaEnginePatchModel := &watsonxdatav3.NetezzaEnginePatch{
				Description: core.StringPtr("netezza engine updated description"),
				DisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
			}
			netezzaEnginePatchModelAsPatch, asPatchErr := netezzaEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetezzaEngineOptions := &watsonxdatav3.UpdateNetezzaEngineOptions{
				ID: core.StringPtr("testString"),
				Body: netezzaEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngine, response, err := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(netezzaEngine).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngines - Get list of spark engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions)`, func() {
			listSparkEnginesOptions := &watsonxdatav3.ListSparkEnginesOptions{
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
			sparkEndpointsModel := &watsonxdatav3.SparkEndpoints{
				ApplicationsApi: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/application_id"),
				HistoryServerEndpoint: core.StringPtr("$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server"),
				SparkAccessEndpoint: core.StringPtr("$HOST/analytics-engine/details/spark-instance_id"),
				SparkJobsV4Endpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications"),
				SparkKernelEndpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels"),
				ViewHistoryServer: core.StringPtr("testString"),
				WxdApplicationEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/applications"),
				WxdEngineEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817"),
				WxdHistoryServerEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/history_server"),
				WxdHistoryServerUiEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/history_server/ui"),
			}

			sparkEngineHomeModel := &watsonxdatav3.SparkEngineHome{
				Path: core.StringPtr("spark/spark1234"),
				StorageName: core.StringPtr("test-spark-storage"),
				Volume: core.StringPtr("test-spark-volume"),
				VolumeID: core.StringPtr("1704979825978585"),
				VolumeName: core.StringPtr("my-volume"),
			}

			sparkEngineResourceLimitModel := &watsonxdatav3.SparkEngineResourceLimit{
				Cores: core.StringPtr("1"),
				Memory: core.StringPtr("4G"),
			}

			sparkEngineResourceUtilisationModel := &watsonxdatav3.SparkEngineResourceUtilisation{
				Cores: core.StringPtr("1m"),
				Memory: core.StringPtr("4Gi"),
			}

			sparkScaleConfigModel := &watsonxdatav3.SparkScaleConfig{
				AutoScaleEnabled: core.BoolPtr(true),
				CurrentNumberOfNodes: core.Int64Ptr(int64(2)),
				MaximumNumberOfNodes: core.Int64Ptr(int64(5)),
				MinimumNumberOfNodes: core.Int64Ptr(int64(1)),
				NodeType: core.StringPtr("medium"),
				NumberOfNodes: core.Int64Ptr(int64(2)),
			}

			sparkVscodeConfigModel := &watsonxdatav3.SparkVscodeConfig{
				Crn: core.StringPtr("crn:v1:staging:public:lakehouse:us-east:a/9aa2b62f2a644ffb9e004451dc631307:00924abc-59a3-45bc-a2c0-58bbccf89ee2"),
				EnvironmentType: core.StringPtr("SaaS"),
				Host: core.StringPtr("us-south.lakehouse.dev.cloud.ibm.com"),
				UserName: core.StringPtr("user@example.com"),
			}

			sparkEngineDetailsModel := &watsonxdatav3.SparkEngineDetails{
				ApiKey: core.StringPtr("apikey"),
				ConnectionString: core.StringPtr("https://xyz.region.ae.cloud.123.com/v3/analytics_engines/spark_iae_id"),
				DefaultConfig: map[string]string{"key1": "configuration"},
				DefaultVersion: core.StringPtr("4.8.3"),
				Endpoints: sparkEndpointsModel,
				EngineHome: sparkEngineHomeModel,
				EngineSubType: core.StringPtr("java"),
				InstanceID: core.StringPtr("spark-id"),
				ManagedBy: core.StringPtr("fully"),
				ResourceLimitEnabled: core.BoolPtr(true),
				ResourceLimits: sparkEngineResourceLimitModel,
				ResourceUtilisation: sparkEngineResourceUtilisationModel,
				ScaleConfig: sparkScaleConfigModel,
				VscodeConfig: sparkVscodeConfigModel,
			}

			createSparkEngineOptions := &watsonxdatav3.CreateSparkEngineOptions{
				DisplayName: core.StringPtr("sampleEngine"),
				Origin: core.StringPtr("external"),
				AssociatedCatalogs: []string{"iceberg_data", "hive_data"},
				Configuration: sparkEngineDetailsModel,
				Description: core.StringPtr("spark engine description"),
				ID: core.StringPtr("spark123"),
				Status: core.StringPtr("provisioning"),
				Tags: []string{"tag1", "tag2"},
				Type: core.StringPtr("spark"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.CreateSparkEngine(createSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineCatalog - Get spark engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineCatalog(getSparkEngineCatalogOptions *GetSparkEngineCatalogOptions)`, func() {
			getSparkEngineCatalogOptions := &watsonxdatav3.GetSparkEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetSparkEngineCatalog(getSparkEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngine - Get spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngine(getSparkEngineOptions *GetSparkEngineOptions)`, func() {
			getSparkEngineOptions := &watsonxdatav3.GetSparkEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.GetSparkEngine(getSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdateSparkEngine - Update spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions)`, func() {
			sparkEngineHomePatchModel := &watsonxdatav3.SparkEngineHomePatch{
				StorageName: core.StringPtr("test-spark-storage"),
			}

			sparkEngineResourceLimitModel := &watsonxdatav3.SparkEngineResourceLimit{
				Cores: core.StringPtr("1"),
				Memory: core.StringPtr("4G"),
			}

			sparkEnginePatchEngineDetailsModel := &watsonxdatav3.SparkEnginePatchEngineDetails{
				DefaultConfig: map[string]string{"key1": "configuration"},
				DefaultVersion: core.StringPtr("4.8.3"),
				EngineHome: sparkEngineHomePatchModel,
				ResourceLimitEnabled: core.BoolPtr(true),
				ResourceLimits: sparkEngineResourceLimitModel,
			}

			sparkEnginePatchModel := &watsonxdatav3.SparkEnginePatch{
				Configuration: sparkEnginePatchEngineDetailsModel,
				Description: core.StringPtr("updated description for spark engine"),
				DisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
			}
			sparkEnginePatchModelAsPatch, asPatchErr := sparkEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSparkEngineOptions := &watsonxdatav3.UpdateSparkEngineOptions{
				ID: core.StringPtr("testString"),
				Body: sparkEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngineCatalogs - Get spark engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngineCatalogs(listSparkEngineCatalogsOptions *ListSparkEngineCatalogsOptions)`, func() {
			listSparkEngineCatalogsOptions := &watsonxdatav3.ListSparkEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListSparkEngineCatalogs(listSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngineCatalogs - Associate catalogs to spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineCatalogs(createSparkEngineCatalogsOptions *CreateSparkEngineCatalogsOptions)`, func() {
			createSparkEngineCatalogsOptions := &watsonxdatav3.CreateSparkEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				CatalogNames: []string{"iceberg_catalog"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.CreateSparkEngineCatalogs(createSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`PauseSparkEngine - Pause engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PauseSparkEngine(pauseSparkEngineOptions *PauseSparkEngineOptions)`, func() {
			pauseSparkEngineOptions := &watsonxdatav3.PauseSparkEngineOptions{
				ID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.PauseSparkEngine(pauseSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ResumeSparkEngine - Resume engine. This feature is only available in SAAS`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumeSparkEngine(resumeSparkEngineOptions *ResumeSparkEngineOptions)`, func() {
			resumeSparkEngineOptions := &watsonxdatav3.ResumeSparkEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ResumeSparkEngine(resumeSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ScaleSparkEngine - Scale engine. This feature is only available in SAAS`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScaleSparkEngine(scaleSparkEngineOptions *ScaleSparkEngineOptions)`, func() {
			scaleSparkEngineOptions := &watsonxdatav3.ScaleSparkEngineOptions{
				ID: core.StringPtr("testString"),
				NumberOfNodes: core.Int64Ptr(int64(2)),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ScaleSparkEngine(scaleSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineApplicationStatus - Get spark engine application details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions)`, func() {
			getSparkEngineApplicationStatusOptions := &watsonxdatav3.GetSparkEngineApplicationStatusOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationResponse, response, err := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineApplicationUi - Redirect to spark UI`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineApplicationUi(getSparkEngineApplicationUiOptions *GetSparkEngineApplicationUiOptions)`, func() {
			getSparkEngineApplicationUiOptions := &watsonxdatav3.GetSparkEngineApplicationUiOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.GetSparkEngineApplicationUi(getSparkEngineApplicationUiOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ListSparkEngineApplications - List all applications in a spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) with pagination`, func(){
			listSparkEngineApplicationsOptions := &watsonxdatav3.ListSparkEngineApplicationsOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				State: []string{"testString"},
				SubmissionTimeInterval: core.StringPtr("testString"),
				StartTimeInterval: core.StringPtr("testString"),
				EndTimeInterval: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
				Start: core.StringPtr("testString"),
			}

			listSparkEngineApplicationsOptions.Start = nil
			listSparkEngineApplicationsOptions.Limit = core.Int64Ptr(1)

			var allResults []watsonxdatav3.SparkEngineApplicationSummary
			for {
				sparkEngineApplicationCollection, response, err := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(sparkEngineApplicationCollection).ToNot(BeNil())
				allResults = append(allResults, sparkEngineApplicationCollection.Applications...)

				listSparkEngineApplicationsOptions.Start, err = sparkEngineApplicationCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listSparkEngineApplicationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) using SparkEngineApplicationsPager`, func(){
			listSparkEngineApplicationsOptions := &watsonxdatav3.ListSparkEngineApplicationsOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				State: []string{"testString"},
				SubmissionTimeInterval: core.StringPtr("testString"),
				StartTimeInterval: core.StringPtr("testString"),
				EndTimeInterval: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := watsonxDataService.NewSparkEngineApplicationsPager(listSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []watsonxdatav3.SparkEngineApplicationSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = watsonxDataService.NewSparkEngineApplicationsPager(listSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSparkEngineApplications() returned a total of %d item(s) using SparkEngineApplicationsPager.\n", len(allResults))
		})
	})

	Describe(`CreateSparkEngineApplication - Create spark engine application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions)`, func() {
			sparkApplicationRuntimeModel := &watsonxdatav3.SparkApplicationRuntime{
				SparkVersion: core.StringPtr("3.4"),
			}

			sparkApplicationDetailsModel := &watsonxdatav3.SparkApplicationDetails{
				Application: core.StringPtr("s3://mybucket/wordcount.py"),
				Archives: core.StringPtr("s3://mybucket/myarchive.zip"),
				Arguments: []string{"people.txt"},
				Class: core.StringPtr("org.apache.spark.examples.SparkPi"),
				Conf: map[string]string{"key1": "configuration"},
				Env: map[string]string{"key1": "configuration"},
				Files: core.StringPtr("s3://mybucket/myfile.txt"),
				Jars: core.StringPtr("testString"),
				Name: core.StringPtr("SparkApplicaton1"),
				Packages: core.StringPtr("org.apache.spark:example_1.2.3"),
				Repositories: core.StringPtr("https://repo1.maven.org/maven2/"),
				Runtime: sparkApplicationRuntimeModel,
				SparkVersion: core.StringPtr("3.3"),
			}

			sparkEngineApplicationCallbackModel := &watsonxdatav3.SparkEngineApplicationCallback{
				URL: core.StringPtr("testString"),
			}

			sparkVolumeDetailsModel := &watsonxdatav3.SparkVolumeDetails{
				MountPath: core.StringPtr("/mount/path"),
				Name: core.StringPtr("my-volume"),
				ReadOnly: core.BoolPtr(true),
				SourceSubPath: core.StringPtr("/source/path"),
			}

			createSparkEngineApplicationOptions := &watsonxdatav3.CreateSparkEngineApplicationOptions{
				ID: core.StringPtr("testString"),
				ApplicationDetails: sparkApplicationDetailsModel,
				Callback: sparkEngineApplicationCallbackModel,
				ContextID: core.StringPtr("testString"),
				ContextType: core.StringPtr("project"),
				DeployMode: core.StringPtr("local"),
				IdempotencyKey: core.StringPtr("testString"),
				InitScripts: []string{"file://init_scripts/test.sh"},
				JobEndpoint: core.StringPtr("<host>/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/engine_applications"),
				MaxRetries: core.StringPtr("3"),
				MinRetryIntervalInSeconds: core.StringPtr("30"),
				ServiceInstanceID: core.StringPtr("iae"),
				TimeoutInSeconds: core.StringPtr("60"),
				Type: core.StringPtr("spark"),
				Volumes: []watsonxdatav3.SparkVolumeDetails{*sparkVolumeDetailsModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationSummary, response, err := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationSummary).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineHistoryServer - Get Spark history server details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineHistoryServer(getSparkEngineHistoryServerOptions *GetSparkEngineHistoryServerOptions)`, func() {
			getSparkEngineHistoryServerOptions := &watsonxdatav3.GetSparkEngineHistoryServerOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkHistoryServer, response, err := watsonxDataService.GetSparkEngineHistoryServer(getSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkHistoryServer).ToNot(BeNil())
		})
	})

	Describe(`StartSparkEngineHistoryServer - Start spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`StartSparkEngineHistoryServer(startSparkEngineHistoryServerOptions *StartSparkEngineHistoryServerOptions)`, func() {
			startSparkEngineHistoryServerOptions := &watsonxdatav3.StartSparkEngineHistoryServerOptions{
				ID: core.StringPtr("testString"),
				Cores: core.StringPtr("1"),
				Memory: core.StringPtr("4G"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkHistoryServer, response, err := watsonxDataService.StartSparkEngineHistoryServer(startSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkHistoryServer).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineHistoryServerUi - Redirect to spark history server UI`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineHistoryServerUi(getSparkEngineHistoryServerUiOptions *GetSparkEngineHistoryServerUiOptions)`, func() {
			getSparkEngineHistoryServerUiOptions := &watsonxdatav3.GetSparkEngineHistoryServerUiOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.GetSparkEngineHistoryServerUi(getSparkEngineHistoryServerUiOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ValidateIntegration - To validate an integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateIntegration(validateIntegrationOptions *ValidateIntegrationOptions)`, func() {
			catalogsModel := &watsonxdatav3.Catalogs{
				CatalogNames: []string{"iceberg_data", "hive_data"},
			}

			validateIntegrationOptions := &watsonxdatav3.ValidateIntegrationOptions{
				Type: core.StringPtr("ranger"),
				AccessToken: core.StringPtr("Header.Payload.Signature"),
				Apikey: core.StringPtr("apikey"),
				Catalogs: catalogsModel,
				Certificate: core.StringPtr("certificate_content_base64_encoded"),
				Password: core.StringPtr("password"),
				Ssl: core.BoolPtr(true),
				URL: core.StringPtr("https://www.abcd.com"),
				Username: core.StringPtr("username"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			validateIntegration, response, err := watsonxDataService.ValidateIntegration(validateIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validateIntegration).ToNot(BeNil())
		})
	})

	Describe(`ListAllIntegrations - List integrations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAllIntegrations(listAllIntegrationsOptions *ListAllIntegrationsOptions)`, func() {
			listAllIntegrationsOptions := &watsonxdatav3.ListAllIntegrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Secret: core.StringPtr("testString"),
				Type: []string{"testString"},
				State: []string{"active"},
			}

			integrationCollection, response, err := watsonxDataService.ListAllIntegrations(listAllIntegrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateIntegration - Create an integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions)`, func() {
			catalogsModel := &watsonxdatav3.Catalogs{
				CatalogNames: []string{"iceberg_data", "hive_data"},
			}

			createIntegrationOptions := &watsonxdatav3.CreateIntegrationOptions{
				AccessToken: core.StringPtr("Header.Payload.Signature"),
				Apikey: core.StringPtr("apikey"),
				Catalogs: catalogsModel,
				Certificate: core.StringPtr("certificate_content_base64_encoded"),
				CertificateExtension: core.StringPtr("pem"),
				ConnectionMode: core.StringPtr("external"),
				CrossAccountIntegration: core.BoolPtr(false),
				EnableDataPolicyWithinWxd: core.BoolPtr(false),
				IkcUserAccountID: core.StringPtr("ikc_user_account_id"),
				Password: core.StringPtr("password"),
				PolicyCacheTimeConfiguration: core.StringPtr("123456789"),
				Resource: core.StringPtr("presto01"),
				Ssl: core.BoolPtr(true),
				Type: core.StringPtr("ranger"),
				URL: core.StringPtr("https://abcd.efgh.com"),
				Username: core.StringPtr("username"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			integration, response, err := watsonxDataService.CreateIntegration(createIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(integration).ToNot(BeNil())
		})
	})

	Describe(`GetIntegrations - Get an Integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIntegrations(getIntegrationsOptions *GetIntegrationsOptions)`, func() {
			getIntegrationsOptions := &watsonxdatav3.GetIntegrationsOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			integration, response, err := watsonxDataService.GetIntegrations(getIntegrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integration).ToNot(BeNil())
		})
	})

	Describe(`UpdateIntegration - Update an Integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateIntegration(updateIntegrationOptions *UpdateIntegrationOptions)`, func() {
			catalogsModel := &watsonxdatav3.Catalogs{
				CatalogNames: []string{"iceberg_data", "hive_data"},
			}

			integrationPatchModel := &watsonxdatav3.IntegrationPatch{
				AccessToken: core.StringPtr("Header.Payload.Signature"),
				Apikey: core.StringPtr("apikey"),
				Catalogs: catalogsModel,
				Certificate: core.StringPtr("certificate_content_base64_encoded"),
				CertificateExtension: core.StringPtr("pem"),
				ConnectionMode: core.StringPtr("external"),
				CrossAccountIntegration: core.BoolPtr(false),
				EnableDataPolicyWithinWxd: core.BoolPtr(false),
				IkcUserAccountID: core.StringPtr("ikc_user_account_id"),
				Password: core.StringPtr("password"),
				PolicyCacheTimeConfiguration: core.StringPtr("123456789"),
				Resource: core.StringPtr("presto01"),
				Ssl: core.BoolPtr(true),
				State: core.StringPtr("active"),
				URL: core.StringPtr("https://abcd.efgh.com"),
				Username: core.StringPtr("username"),
			}
			integrationPatchModelAsPatch, asPatchErr := integrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIntegrationOptions := &watsonxdatav3.UpdateIntegrationOptions{
				ID: core.StringPtr("testString"),
				Body: integrationPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
				Secret: core.StringPtr("testString"),
			}

			integration, response, err := watsonxDataService.UpdateIntegration(updateIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integration).ToNot(BeNil())
		})
	})

	Describe(`RegisterTable - Register delta and hudi tables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RegisterTable(registerTableOptions *RegisterTableOptions)`, func() {
			registerTableOptions := &watsonxdatav3.RegisterTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				MetadataLocation: core.StringPtr("s3a://storagename/path/to/table/metadata_location/_delta_log"),
				TableName: core.StringPtr("table1"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			registerTableCreatedBody, response, err := watsonxDataService.RegisterTable(registerTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(registerTableCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`LoadTable - Load delta and hudi tables metadata`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LoadTable(loadTableOptions *LoadTableOptions)`, func() {
			loadTableOptions := &watsonxdatav3.LoadTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			loadTableResponse, response, err := watsonxDataService.LoadTable(loadTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadTableResponse).ToNot(BeNil())
		})
	})

	Describe(`ListCatalogs - Get list of catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
			listCatalogsOptions := &watsonxdatav3.ListCatalogsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Secret: core.StringPtr("testString"),
				DefaultCatalogs: core.BoolPtr(false),
				View: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListCatalogs(listCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetCatalogEngineAssociation - Get catalog engine association`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogEngineAssociation(getCatalogEngineAssociationOptions *GetCatalogEngineAssociationOptions)`, func() {
			getCatalogEngineAssociationOptions := &watsonxdatav3.GetCatalogEngineAssociationOptions{
				CatalogName: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogEngineResponse, response, err := watsonxDataService.GetCatalogEngineAssociation(getCatalogEngineAssociationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEngineResponse).ToNot(BeNil())
		})
	})

	Describe(`ListTables - List all tables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTables(listTablesOptions *ListTablesOptions)`, func() {
			listTablesOptions := &watsonxdatav3.ListTablesOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
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
			getTableOptions := &watsonxdatav3.GetTableOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			table, response, err := watsonxDataService.GetTable(getTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(table).ToNot(BeNil())
		})
	})

	Describe(`UpdateTable - Rename table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTable(updateTableOptions *UpdateTableOptions)`, func() {
			tablePatchModel := &watsonxdatav3.TablePatch{
				Name: core.StringPtr("updated_table_name"),
			}
			tablePatchModelAsPatch, asPatchErr := tablePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTableOptions := &watsonxdatav3.UpdateTableOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				Body: tablePatchModelAsPatch,
				Type: core.StringPtr("testString"),
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
			listColumnsOptions := &watsonxdatav3.ListColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
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
			columnModel := &watsonxdatav3.Column{
				Comment: core.StringPtr("Expenses column for each department"),
				Extra: core.StringPtr("AUTO_INCREMENT"),
				Length: core.StringPtr("30"),
				Name: core.StringPtr("expenses"),
				Precision: core.StringPtr("10"),
				Scale: core.StringPtr("2"),
				Type: core.StringPtr("varchar"),
			}

			createColumnsOptions := &watsonxdatav3.CreateColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				Columns: []watsonxdatav3.Column{*columnModel},
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
			columnPatchModel := &watsonxdatav3.ColumnPatch{
				Name: core.StringPtr("expenses"),
			}
			columnPatchModelAsPatch, asPatchErr := columnPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateColumnOptions := &watsonxdatav3.UpdateColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				ColumnName: core.StringPtr("testString"),
				Body: columnPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			column, response, err := watsonxDataService.UpdateColumn(updateColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(column).ToNot(BeNil())
		})
	})

	Describe(`RollbackTable - Rollback snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RollbackTable(rollbackTableOptions *RollbackTableOptions)`, func() {
			rollbackTableOptions := &watsonxdatav3.RollbackTableOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("12357647"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponsePrototype, response, err := watsonxDataService.RollbackTable(rollbackTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponsePrototype).ToNot(BeNil())
		})
	})

	Describe(`ListTableSnapshots - Get table snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions)`, func() {
			listTableSnapshotsOptions := &watsonxdatav3.ListTableSnapshotsOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			tableSnapshotCollection, response, err := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableSnapshotCollection).ToNot(BeNil())
		})
	})

	Describe(`ListSchemas - List all schemas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchemas(listSchemasOptions *ListSchemasOptions)`, func() {
			listSchemasOptions := &watsonxdatav3.ListSchemasOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			schemasCollection, response, err := watsonxDataService.ListSchemas(listSchemasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schemasCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSchema - Create schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
			createSchemaOptions := &watsonxdatav3.CreateSchemaOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				CustomPath: core.StringPtr("sample-path"),
				Name: core.StringPtr("SampleSchema1"),
				Hostname: core.StringPtr("db2@hostname.com"),
				Port: core.Int64Ptr(int64(4553)),
				StorageName: core.StringPtr("sample-bucket"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			schemaPrototype, response, err := watsonxDataService.CreateSchema(createSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(schemaPrototype).ToNot(BeNil())
		})
	})

	Describe(`UpdateSyncCatalog - External Iceberg table registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions)`, func() {
			updateSyncCatalogOptions := &watsonxdatav3.UpdateSyncCatalogOptions{
				ID: core.StringPtr("testString"),
				AutoAddNewTables: core.BoolPtr(true),
				RegisterNewTables: core.BoolPtr(true),
				SyncExistingTables: core.BoolPtr(true),
				SyncIcebergMd: core.BoolPtr(true),
				SyncPath: core.StringPtr("sample-path"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponsePrototype, response, err := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponsePrototype).ToNot(BeNil())
		})
	})

	Describe(`GetCatalog - Get catalog properties by catalog_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
			getCatalogOptions := &watsonxdatav3.GetCatalogOptions{
				Name: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetCatalog(getCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusServices - Get list of milvus services`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusServices(listMilvusServicesOptions *ListMilvusServicesOptions)`, func() {
			listMilvusServicesOptions := &watsonxdatav3.ListMilvusServicesOptions{
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
			createMilvusServiceOptions := &watsonxdatav3.CreateMilvusServiceOptions{
				DisplayName: core.StringPtr("sampleService"),
				Origin: core.StringPtr("native"),
				RootPath: core.StringPtr("Sample/path"),
				TshirtSize: core.StringPtr("small"),
				DcCpu: core.Float64Ptr(float64(0.01)),
				DcMemory: core.Float64Ptr(float64(0.01)),
				DcReplicas: core.Int64Ptr(int64(1)),
				Description: core.StringPtr("milvus service for running sql queries"),
				DwCpu: core.Float64Ptr(float64(0.01)),
				DwMemory: core.Float64Ptr(float64(0.01)),
				DwReplicas: core.Int64Ptr(int64(1)),
				EtcdCpu: core.Float64Ptr(float64(0.01)),
				EtcdMemory: core.Float64Ptr(float64(0.01)),
				ID: core.StringPtr("milvus123"),
				IndexType: core.StringPtr("ivf_sq8"),
				IwCpu: core.Float64Ptr(float64(0.01)),
				IwMemory: core.Float64Ptr(float64(0.01)),
				IwReplicas: core.Int64Ptr(int64(1)),
				KafkaCpu: core.Float64Ptr(float64(0.01)),
				KafkaMemory: core.Float64Ptr(float64(0.01)),
				ProxyCpu: core.Float64Ptr(float64(0.01)),
				ProxyMemory: core.Float64Ptr(float64(0.01)),
				ProxyReplicas: core.Int64Ptr(int64(1)),
				QcCpu: core.Float64Ptr(float64(0.01)),
				QcMemory: core.Float64Ptr(float64(0.01)),
				QcReplicas: core.Int64Ptr(int64(1)),
				QwCpu: core.Float64Ptr(float64(0.01)),
				QwMemory: core.Float64Ptr(float64(0.01)),
				QwReplicas: core.Int64Ptr(int64(1)),
				RcCpu: core.Float64Ptr(float64(0.01)),
				RcMemory: core.Float64Ptr(float64(0.01)),
				RcReplicas: core.Int64Ptr(int64(1)),
				StorageName: core.StringPtr("Sample_storage_name"),
				Tags: []string{"tag1", "tag2"},
				Vector: core.Int64Ptr(int64(1)),
				VectorDimension: core.Int64Ptr(int64(384)),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.CreateMilvusService(createMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`GetMilvusService - Get milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMilvusService(getMilvusServiceOptions *GetMilvusServiceOptions)`, func() {
			getMilvusServiceOptions := &watsonxdatav3.GetMilvusServiceOptions{
				ID: core.StringPtr("testString"),
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
			milvusServicePatchModel := &watsonxdatav3.MilvusServicePatch{
				Description: core.StringPtr("updated description for milvus service"),
				DisplayName: core.StringPtr("sampleService"),
				Tags: []string{"tag1", "tag2"},
			}
			milvusServicePatchModelAsPatch, asPatchErr := milvusServicePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateMilvusServiceOptions := &watsonxdatav3.UpdateMilvusServiceOptions{
				ID: core.StringPtr("testString"),
				Body: milvusServicePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.UpdateMilvusService(updateMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`CreateMilvusServicePause - Pause milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMilvusServicePause(createMilvusServicePauseOptions *CreateMilvusServicePauseOptions)`, func() {
			createMilvusServicePauseOptions := &watsonxdatav3.CreateMilvusServicePauseOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateMilvusServicePause(createMilvusServicePauseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateMilvusServiceResume - Resume milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMilvusServiceResume(createMilvusServiceResumeOptions *CreateMilvusServiceResumeOptions)`, func() {
			createMilvusServiceResumeOptions := &watsonxdatav3.CreateMilvusServiceResumeOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateMilvusServiceResume(createMilvusServiceResumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateMilvusServiceScale - Scale a milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMilvusServiceScale(createMilvusServiceScaleOptions *CreateMilvusServiceScaleOptions)`, func() {
			createMilvusServiceScaleOptions := &watsonxdatav3.CreateMilvusServiceScaleOptions{
				ID: core.StringPtr("testString"),
				TshirtSize: core.StringPtr("small"),
				DcCpu: core.Float64Ptr(float64(0.01)),
				DcMemory: core.Float64Ptr(float64(0.01)),
				DcReplicas: core.Int64Ptr(int64(1)),
				DwCpu: core.Float64Ptr(float64(0.01)),
				DwMemory: core.Float64Ptr(float64(0.01)),
				DwReplicas: core.Int64Ptr(int64(1)),
				EtcdCpu: core.Float64Ptr(float64(0.01)),
				EtcdMemory: core.Float64Ptr(float64(0.01)),
				IndexType: core.StringPtr("flat"),
				IwCpu: core.Float64Ptr(float64(0.01)),
				IwMemory: core.Float64Ptr(float64(0.01)),
				IwReplicas: core.Int64Ptr(int64(1)),
				KafkaCpu: core.Float64Ptr(float64(0.01)),
				KafkaMemory: core.Float64Ptr(float64(0.01)),
				ProxyCpu: core.Float64Ptr(float64(0.01)),
				ProxyMemory: core.Float64Ptr(float64(0.01)),
				ProxyReplicas: core.Int64Ptr(int64(1)),
				QcCpu: core.Float64Ptr(float64(0.01)),
				QcMemory: core.Float64Ptr(float64(0.01)),
				QcReplicas: core.Int64Ptr(int64(1)),
				QwCpu: core.Float64Ptr(float64(0.01)),
				QwMemory: core.Float64Ptr(float64(0.01)),
				QwReplicas: core.Int64Ptr(int64(1)),
				RcCpu: core.Float64Ptr(float64(0.01)),
				RcMemory: core.Float64Ptr(float64(0.01)),
				RcReplicas: core.Int64Ptr(int64(1)),
				Vector: core.Int64Ptr(int64(1)),
				VectorDimension: core.Int64Ptr(int64(384)),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateMilvusServiceScale(createMilvusServiceScaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusServiceDatabases - Get milvus service databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusServiceDatabases(listMilvusServiceDatabasesOptions *ListMilvusServiceDatabasesOptions)`, func() {
			listMilvusServiceDatabasesOptions := &watsonxdatav3.ListMilvusServiceDatabasesOptions{
				ServiceID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusServiceDatabases, response, err := watsonxDataService.ListMilvusServiceDatabases(listMilvusServiceDatabasesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusServiceDatabases).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusDatabaseCollections - Get milvus database collections`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusDatabaseCollections(listMilvusDatabaseCollectionsOptions *ListMilvusDatabaseCollectionsOptions)`, func() {
			listMilvusDatabaseCollectionsOptions := &watsonxdatav3.ListMilvusDatabaseCollectionsOptions{
				ServiceID: core.StringPtr("testString"),
				DatabaseID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusDatabaseCollections, response, err := watsonxDataService.ListMilvusDatabaseCollections(listMilvusDatabaseCollectionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusDatabaseCollections).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusDatabasePartitions - Get milvus database collection partitions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusDatabasePartitions(listMilvusDatabasePartitionsOptions *ListMilvusDatabasePartitionsOptions)`, func() {
			listMilvusDatabasePartitionsOptions := &watsonxdatav3.ListMilvusDatabasePartitionsOptions{
				ServiceID: core.StringPtr("testString"),
				DatabaseID: core.StringPtr("testString"),
				CollectionName: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusDatabasePartitions, response, err := watsonxDataService.ListMilvusDatabasePartitions(listMilvusDatabasePartitionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusDatabasePartitions).ToNot(BeNil())
		})
	})

	Describe(`UpdateMilvusServiceBucket - Update milvus service storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateMilvusServiceBucket(updateMilvusServiceBucketOptions *UpdateMilvusServiceBucketOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav3.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: "testString",
			}

			updateMilvusServiceBucketOptions := &watsonxdatav3.UpdateMilvusServiceBucketOptions{
				ServiceID: core.StringPtr("testString"),
				Body: []watsonxdatav3.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.UpdateMilvusServiceBucket(updateMilvusServiceBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegration - Get semantic automation layer Integration of the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegration(getSalIntegrationOptions *GetSalIntegrationOptions)`, func() {
			getSalIntegrationOptions := &watsonxdatav3.GetSalIntegrationOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			salIntegration, response, err := watsonxDataService.GetSalIntegration(getSalIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salIntegration).ToNot(BeNil())
		})
	})

	Describe(`CreateSalIntegration - Create semantic automation layer integration with IBM Knowledge Catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSalIntegration(createSalIntegrationOptions *CreateSalIntegrationOptions)`, func() {
			createSalIntegrationOptions := &watsonxdatav3.CreateSalIntegrationOptions{
				Apikey: core.StringPtr("67GveYtUdovRFEfnMLYP8x0S1b2mY1BkEGqBYbJK"),
				EngineID: core.StringPtr("presto-01"),
				StorageResourceCrn: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/a7026b374f39f570d20984c1ac6ecf63:5778e94f-c8c7-46a8-9878-d5eeadb51161"),
				StorageType: core.StringPtr("bmcos_object_storage"),
				TrialPlan: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salIntegration, response, err := watsonxDataService.CreateSalIntegration(createSalIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(salIntegration).ToNot(BeNil())
		})
	})

	Describe(`UpdateSalIntegration - Patch Semantic automation layer integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSalIntegration(updateSalIntegrationOptions *UpdateSalIntegrationOptions)`, func() {
			salIntegrationPatchModel := &watsonxdatav3.SalIntegrationPatch{
				Apikey: core.StringPtr("67GveYtUdovRFEfnMLYP8x0S1b2mY1BkEGqBYbJK"),
				EngineID: core.StringPtr("presto-01"),
			}
			salIntegrationPatchModelAsPatch, asPatchErr := salIntegrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSalIntegrationOptions := &watsonxdatav3.UpdateSalIntegrationOptions{
				SalIntegrationPatch: salIntegrationPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			salIntegration, response, err := watsonxDataService.UpdateSalIntegration(updateSalIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salIntegration).ToNot(BeNil())
		})
	})

	Describe(`CreateSalIntegrationEnrichment - create enrichment jobs of schema(or tables in schema)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSalIntegrationEnrichment(createSalIntegrationEnrichmentOptions *CreateSalIntegrationEnrichmentOptions)`, func() {
			enrichmentObjModel := &watsonxdatav3.EnrichmentObj{
				Catalog: core.StringPtr("iceberg_data"),
				Operation: core.StringPtr("create"),
				Schema: core.StringPtr("schema1"),
				Tables: []string{"table1"},
			}

			createSalIntegrationEnrichmentOptions := &watsonxdatav3.CreateSalIntegrationEnrichmentOptions{
				Changes: []watsonxdatav3.EnrichmentObj{*enrichmentObjModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.CreateSalIntegrationEnrichment(createSalIntegrationEnrichmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListSalIntegrationEnrichmentAssets - List semantic enriched data_assets in IBM Knowledge catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSalIntegrationEnrichmentAssets(listSalIntegrationEnrichmentAssetsOptions *ListSalIntegrationEnrichmentAssetsOptions)`, func() {
			listSalIntegrationEnrichmentAssetsOptions := &watsonxdatav3.ListSalIntegrationEnrichmentAssetsOptions{
				ProjectID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentDataAssetCollection, response, err := watsonxDataService.ListSalIntegrationEnrichmentAssets(listSalIntegrationEnrichmentAssetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentDataAssetCollection).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegrationEnrichmentAssetsByID - Get semantic enrichment data_asset associated with the table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationEnrichmentAssetsByID(getSalIntegrationEnrichmentAssetsByIdOptions *GetSalIntegrationEnrichmentAssetsByIdOptions)`, func() {
			getSalIntegrationEnrichmentAssetsByIdOptions := &watsonxdatav3.GetSalIntegrationEnrichmentAssetsByIdOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentDataAsset, response, err := watsonxDataService.GetSalIntegrationEnrichmentAssetsByID(getSalIntegrationEnrichmentAssetsByIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentDataAsset).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegrationEnrichmentGlobalSettings - Get metadata enrichment global settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationEnrichmentGlobalSettings(getSalIntegrationEnrichmentGlobalSettingsOptions *GetSalIntegrationEnrichmentGlobalSettingsOptions)`, func() {
			getSalIntegrationEnrichmentGlobalSettingsOptions := &watsonxdatav3.GetSalIntegrationEnrichmentGlobalSettingsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentSettings, response, err := watsonxDataService.GetSalIntegrationEnrichmentGlobalSettings(getSalIntegrationEnrichmentGlobalSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentSettings).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSalIntegrationEnrichmentGlobalSettings - Create or update metadata enrichment global settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSalIntegrationEnrichmentGlobalSettings(replaceSalIntegrationEnrichmentGlobalSettingsOptions *ReplaceSalIntegrationEnrichmentGlobalSettingsOptions)`, func() {
			salEnrichmentSettingsExpansionDescriptionConfigurationModel := &watsonxdatav3.SalEnrichmentSettingsExpansionDescriptionConfiguration{
				AssignmentThreshold: core.Float64Ptr(float64(0.14)),
				SuggestionThreshold: core.Float64Ptr(float64(0.9)),
			}

			salEnrichmentSettingsExpansionNameConfigurationModel := &watsonxdatav3.SalEnrichmentSettingsExpansionNameConfiguration{
				AssignmentThreshold: core.Float64Ptr(float64(0.1)),
				SuggestionThreshold: core.Float64Ptr(float64(0.1)),
			}

			salEnrichmentSettingsExpansionModel := &watsonxdatav3.SalEnrichmentSettingsExpansion{
				Description: core.BoolPtr(true),
				DescriptionConfiguration: salEnrichmentSettingsExpansionDescriptionConfigurationModel,
				Name: core.BoolPtr(true),
				NameConfiguration: salEnrichmentSettingsExpansionNameConfigurationModel,
			}

			salEnrichmentSettingsTermAssignmentModel := &watsonxdatav3.SalEnrichmentSettingsTermAssignment{
				ClassBasedAssignments: core.BoolPtr(false),
				EvaluateNegativeAssignments: core.BoolPtr(false),
				LlmBasedAssignments: core.BoolPtr(false),
				MlBasedAssignmentsCustom: core.BoolPtr(false),
				MlBasedAssignmentsDefault: core.BoolPtr(false),
				NameMatching: core.BoolPtr(false),
				TermAssignmentThreshold: core.Float64Ptr(float64(0.3)),
				TermSuggestionThreshold: core.Float64Ptr(float64(0.4)),
			}

			replaceSalIntegrationEnrichmentGlobalSettingsOptions := &watsonxdatav3.ReplaceSalIntegrationEnrichmentGlobalSettingsOptions{
				Expansion: salEnrichmentSettingsExpansionModel,
				TermAssignment: salEnrichmentSettingsTermAssignmentModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentSettings, response, err := watsonxDataService.ReplaceSalIntegrationEnrichmentGlobalSettings(replaceSalIntegrationEnrichmentGlobalSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentSettings).ToNot(BeNil())
		})
	})

	Describe(`ListSalIntegrationEnrichmentJobs - List semantic enrichment jobs associated with the schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSalIntegrationEnrichmentJobs(listSalIntegrationEnrichmentJobsOptions *ListSalIntegrationEnrichmentJobsOptions)`, func() {
			listSalIntegrationEnrichmentJobsOptions := &watsonxdatav3.ListSalIntegrationEnrichmentJobsOptions{
				ProjectID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentJobs, response, err := watsonxDataService.ListSalIntegrationEnrichmentJobs(listSalIntegrationEnrichmentJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentJobs).ToNot(BeNil())
		})
	})

	Describe(`ListSalIntegrationEnrichmentJobRuns - List semantic enrichment job runs associated with the schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSalIntegrationEnrichmentJobRuns(listSalIntegrationEnrichmentJobRunsOptions *ListSalIntegrationEnrichmentJobRunsOptions)`, func() {
			listSalIntegrationEnrichmentJobRunsOptions := &watsonxdatav3.ListSalIntegrationEnrichmentJobRunsOptions{
				JobID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentJobRuns, response, err := watsonxDataService.ListSalIntegrationEnrichmentJobRuns(listSalIntegrationEnrichmentJobRunsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentJobRuns).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegrationEnrichmentJobRunLogs - Get semantic enrichment job run logs associated with the job run`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationEnrichmentJobRunLogs(getSalIntegrationEnrichmentJobRunLogsOptions *GetSalIntegrationEnrichmentJobRunLogsOptions)`, func() {
			getSalIntegrationEnrichmentJobRunLogsOptions := &watsonxdatav3.GetSalIntegrationEnrichmentJobRunLogsOptions{
				JobID: core.StringPtr("testString"),
				RunID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentJobRunLogs, response, err := watsonxDataService.GetSalIntegrationEnrichmentJobRunLogs(getSalIntegrationEnrichmentJobRunLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentJobRunLogs).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegrationEnrichmentProjectSettings - get metadata enrichment settings for a IBM Knowledge Catalog project(schema)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationEnrichmentProjectSettings(getSalIntegrationEnrichmentProjectSettingsOptions *GetSalIntegrationEnrichmentProjectSettingsOptions)`, func() {
			getSalIntegrationEnrichmentProjectSettingsOptions := &watsonxdatav3.GetSalIntegrationEnrichmentProjectSettingsOptions{
				ProjectID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salEnrichmentSettings, response, err := watsonxDataService.GetSalIntegrationEnrichmentProjectSettings(getSalIntegrationEnrichmentProjectSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salEnrichmentSettings).ToNot(BeNil())
		})
	})

	Describe(`ReplaceSalIntegrationEnrichmentProjectSettings - Create or update project level metadata enrichment settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSalIntegrationEnrichmentProjectSettings(replaceSalIntegrationEnrichmentProjectSettingsOptions *ReplaceSalIntegrationEnrichmentProjectSettingsOptions)`, func() {
			salEnrichmentSettingsExpansionDescriptionConfigurationModel := &watsonxdatav3.SalEnrichmentSettingsExpansionDescriptionConfiguration{
				AssignmentThreshold: core.Float64Ptr(float64(0.14)),
				SuggestionThreshold: core.Float64Ptr(float64(0.9)),
			}

			salEnrichmentSettingsExpansionNameConfigurationModel := &watsonxdatav3.SalEnrichmentSettingsExpansionNameConfiguration{
				AssignmentThreshold: core.Float64Ptr(float64(0.1)),
				SuggestionThreshold: core.Float64Ptr(float64(0.1)),
			}

			salEnrichmentSettingsExpansionModel := &watsonxdatav3.SalEnrichmentSettingsExpansion{
				Description: core.BoolPtr(true),
				DescriptionConfiguration: salEnrichmentSettingsExpansionDescriptionConfigurationModel,
				Name: core.BoolPtr(true),
				NameConfiguration: salEnrichmentSettingsExpansionNameConfigurationModel,
			}

			salEnrichmentSettingsTermAssignmentModel := &watsonxdatav3.SalEnrichmentSettingsTermAssignment{
				ClassBasedAssignments: core.BoolPtr(false),
				EvaluateNegativeAssignments: core.BoolPtr(false),
				LlmBasedAssignments: core.BoolPtr(false),
				MlBasedAssignmentsCustom: core.BoolPtr(false),
				MlBasedAssignmentsDefault: core.BoolPtr(false),
				NameMatching: core.BoolPtr(false),
				TermAssignmentThreshold: core.Float64Ptr(float64(0.3)),
				TermSuggestionThreshold: core.Float64Ptr(float64(0.4)),
			}

			replaceSalIntegrationEnrichmentProjectSettingsOptions := &watsonxdatav3.ReplaceSalIntegrationEnrichmentProjectSettingsOptions{
				ProjectID: core.StringPtr("testString"),
				Expansion: salEnrichmentSettingsExpansionModel,
				TermAssignment: salEnrichmentSettingsTermAssignmentModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.ReplaceSalIntegrationEnrichmentProjectSettings(replaceSalIntegrationEnrichmentProjectSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})

	Describe(`GetSalIntegrationGlossaryTerms - Get list of uploaded glossary terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationGlossaryTerms(getSalIntegrationGlossaryTermsOptions *GetSalIntegrationGlossaryTermsOptions)`, func() {
			getSalIntegrationGlossaryTermsOptions := &watsonxdatav3.GetSalIntegrationGlossaryTermsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			salGlossaryTerms, response, err := watsonxDataService.GetSalIntegrationGlossaryTerms(getSalIntegrationGlossaryTermsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salGlossaryTerms).ToNot(BeNil())
		})
	})

	Describe(`CreateSalIntegrationUploadGlossary - Upload semantic enrichment business terms glossary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSalIntegrationUploadGlossary(createSalIntegrationUploadGlossaryOptions *CreateSalIntegrationUploadGlossaryOptions)`, func() {
			createSalIntegrationUploadGlossaryOptions := &watsonxdatav3.CreateSalIntegrationUploadGlossaryOptions{
				ReplaceOption: core.StringPtr("all"),
				GlossaryCsv: CreateMockReader("This is a mock file."),
				GlossaryCsvContentType: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salGlossaryUploadProcess, response, err := watsonxDataService.CreateSalIntegrationUploadGlossary(createSalIntegrationUploadGlossaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(salGlossaryUploadProcess).ToNot(BeNil())
		})
	})

	Describe(`GetSalIntegrationUploadGlossaryStatus - Get status of glossary uploading process(job)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSalIntegrationUploadGlossaryStatus(getSalIntegrationUploadGlossaryStatusOptions *GetSalIntegrationUploadGlossaryStatusOptions)`, func() {
			getSalIntegrationUploadGlossaryStatusOptions := &watsonxdatav3.GetSalIntegrationUploadGlossaryStatusOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salGlossaryUploadStatus, response, err := watsonxDataService.GetSalIntegrationUploadGlossaryStatus(getSalIntegrationUploadGlossaryStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salGlossaryUploadStatus).ToNot(BeNil())
		})
	})

	Describe(`ListSalIntegrationEnrichmentMappings - Post watsonx.data schema to fetch mapped catalog and project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSalIntegrationEnrichmentMappings(listSalIntegrationEnrichmentMappingsOptions *ListSalIntegrationEnrichmentMappingsOptions)`, func() {
			listSalIntegrationEnrichmentMappingsOptions := &watsonxdatav3.ListSalIntegrationEnrichmentMappingsOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				Next: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			salIntegrationMappings, response, err := watsonxDataService.ListSalIntegrationEnrichmentMappings(listSalIntegrationEnrichmentMappingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(salIntegrationMappings).ToNot(BeNil())
		})
	})

	Describe(`ListSemanticSearchQueries - List Semantic Search queries snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSemanticSearchQueries(listSemanticSearchQueriesOptions *ListSemanticSearchQueriesOptions)`, func() {
			listSemanticSearchQueriesOptions := &watsonxdatav3.ListSemanticSearchQueriesOptions{
				EngineID: core.StringPtr("testString"),
				SchemaSearchEnabled: core.BoolPtr(true),
				ColumnSearchEnabled: core.BoolPtr(true),
				MaxResultNumber: core.Int64Ptr(int64(5)),
				RunSearch: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			semanticSearchList, response, err := watsonxDataService.ListSemanticSearchQueries(listSemanticSearchQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(semanticSearchList).ToNot(BeNil())
		})
	})

	Describe(`CreateSemanticSearchQueries - Trigger a global search among IBM Knowledge Catalog metadata by given query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSemanticSearchQueries(createSemanticSearchQueriesOptions *CreateSemanticSearchQueriesOptions)`, func() {
			semanticSearchBodySearchConfigModel := &watsonxdatav3.SemanticSearchBodySearchConfig{
				ColumnSearchEnabled: core.BoolPtr(true),
				Fields: []string{"metadata.name", "metadata.description", "metadata.tags"},
				MaxResultNumber: core.Int64Ptr(int64(1)),
				SchemaSearchEnabled: core.BoolPtr(true),
			}

			createSemanticSearchQueriesOptions := &watsonxdatav3.CreateSemanticSearchQueriesOptions{
				EngineID: core.StringPtr("presto01"),
				QueryInput: core.StringPtr("catalog_table"),
				SearchConfig: semanticSearchBodySearchConfigModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			semanticSearch, response, err := watsonxDataService.CreateSemanticSearchQueries(createSemanticSearchQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(semanticSearch).ToNot(BeNil())
		})
	})

	Describe(`ListIngestionJobs - List ingestion jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIngestionJobs(listIngestionJobsOptions *ListIngestionJobsOptions) with pagination`, func(){
			listIngestionJobsOptions := &watsonxdatav3.ListIngestionJobsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listIngestionJobsOptions.Start = nil
			listIngestionJobsOptions.Limit = core.Int64Ptr(1)

			var allResults []watsonxdatav3.IngestionJob
			for {
				ingestionJobCollection, response, err := watsonxDataService.ListIngestionJobs(listIngestionJobsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(ingestionJobCollection).ToNot(BeNil())
				allResults = append(allResults, ingestionJobCollection.Jobs...)

				listIngestionJobsOptions.Start, err = ingestionJobCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listIngestionJobsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListIngestionJobs(listIngestionJobsOptions *ListIngestionJobsOptions) using IngestionJobsPager`, func(){
			listIngestionJobsOptions := &watsonxdatav3.ListIngestionJobsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := watsonxDataService.NewIngestionJobsPager(listIngestionJobsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []watsonxdatav3.IngestionJob
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = watsonxDataService.NewIngestionJobsPager(listIngestionJobsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListIngestionJobs() returned a total of %d item(s) using IngestionJobsPager.\n", len(allResults))
		})
	})

	Describe(`CreateIngestionJob - Submit ingestion job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIngestionJob(createIngestionJobOptions *CreateIngestionJobOptions)`, func() {
			schemaTransformationModel := &watsonxdatav3.SchemaTransformation{
				OldColumn: core.StringPtr("old_column_name"),
				NewColumn: core.StringPtr("new_column_name"),
				NewType: core.StringPtr("string"),
			}

			bucketDetailsModel := &watsonxdatav3.BucketDetails{
				AccessKey: core.StringPtr("testString"),
				AccountName: core.StringPtr("testString"),
				ApplicationID: core.StringPtr("testString"),
				ContainerName: core.StringPtr("testString"),
				DirectoryID: core.StringPtr("testString"),
				Endpoint: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				SecretKey: core.StringPtr("testString"),
				AuthMode: core.StringPtr("aws_assume_role"),
				RoleArn: core.StringPtr("testString"),
				ManagedBy: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Type: core.StringPtr("adls_gen1"),
			}

			icebergSourceTableModel := &watsonxdatav3.IcebergSourceTable{
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				WarehouseName: core.StringPtr("testString"),
				SnapshotID: core.Int64Ptr(int64(0)),
			}

			dbConnectionModelModel := &watsonxdatav3.DbConnectionModel{
				DatabaseID: core.StringPtr("testString"),
				DbType: core.StringPtr("testString"),
				Host: core.StringPtr("testString"),
				Port: core.StringPtr("testString"),
				DbName: core.StringPtr("testString"),
				DbUsername: core.StringPtr("testString"),
				Password: core.StringPtr("testString"),
				AuthenticationValue: core.StringPtr("testString"),
				IsSsl: core.BoolPtr(false),
				CertExtension: core.StringPtr("testString"),
				CertContent: core.StringPtr("testString"),
				ConnectionMode: core.StringPtr("testString"),
				ConnectionModeValue: core.StringPtr("testString"),
				JdbcURL: core.StringPtr("testString"),
			}

			fileFormatPropertiesModel := &watsonxdatav3.FileFormatProperties{
				Encoding: core.StringPtr("testString"),
				EscapeCharacter: core.StringPtr("testString"),
				FieldDelimiter: core.StringPtr("testString"),
				Header: core.BoolPtr(true),
				LineDelimiter: core.StringPtr("testString"),
			}

			sourceDetailsModel := &watsonxdatav3.SourceDetails{
				FilePaths: core.StringPtr("testString"),
				FileType: core.StringPtr("csv"),
				SourceType: core.StringPtr("STORAGE"),
				SchemaTransformations: []watsonxdatav3.SchemaTransformation{*schemaTransformationModel},
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				BucketDetails: bucketDetailsModel,
				IcebergSourceTable: icebergSourceTableModel,
				SourceDatabase: dbConnectionModelModel,
				FileFormatProperties: fileFormatPropertiesModel,
				IsLocalIngestion: core.BoolPtr(false),
			}

			targetDetailsModel := &watsonxdatav3.TargetDetails{
				Catalog: core.StringPtr("testString"),
				Schema: core.StringPtr("testString"),
				Table: core.StringPtr("testString"),
				WriteMode: core.StringPtr("testString"),
				MergeOnRead: core.BoolPtr(false),
				SchemaMode: core.StringPtr("testString"),
				SchemaInfer: core.BoolPtr(true),
				CatalogURI: core.StringPtr("testString"),
				BucketDetails: bucketDetailsModel,
				Location: core.StringPtr("testString"),
				IsNewSchema: core.BoolPtr(false),
				IsNewTable: core.BoolPtr(false),
			}

			executeConfigModel := &watsonxdatav3.ExecuteConfig{
				DriverMemory: core.StringPtr("testString"),
				DriverCores: core.Int64Ptr(int64(1)),
				ExecutorMemory: core.StringPtr("testString"),
				ExecutorCores: core.Int64Ptr(int64(1)),
				NumExecutors: core.Int64Ptr(int64(1)),
			}

			ingestionEngineModel := &watsonxdatav3.IngestionEngine{
				EngineID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Endpoint: core.StringPtr("testString"),
				Origin: core.StringPtr("testString"),
				ExecuteConfig: executeConfigModel,
				BucketDetails: bucketDetailsModel,
				LogPath: core.StringPtr("testString"),
			}

			capacityDetailsModel := &watsonxdatav3.CapacityDetails{
				ID: CreateMockUUID("6f51248a-4cda-459a-b141-eb9b76bb6689"),
			}

			createIngestionJobOptions := &watsonxdatav3.CreateIngestionJobOptions{
				AuthInstanceID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Source: sourceDetailsModel,
				Target: targetDetailsModel,
				Engine: ingestionEngineModel,
				EngineID: core.StringPtr("spark123"),
				ExecuteConfig: executeConfigModel,
				SourceIcebergTable: icebergSourceTableModel,
				PartitionBy: core.StringPtr("column1,column2"),
				Capacity: capacityDetailsModel,
			}

			ingestionJob, response, err := watsonxDataService.CreateIngestionJob(createIngestionJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ingestionJob).ToNot(BeNil())
		})
	})

	Describe(`GetIngestionJob - Get ingestion job by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIngestionJob(getIngestionJobOptions *GetIngestionJobOptions)`, func() {
			getIngestionJobOptions := &watsonxdatav3.GetIngestionJobOptions{
				AuthInstanceID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			ingestionJob, response, err := watsonxDataService.GetIngestionJob(getIngestionJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ingestionJob).ToNot(BeNil())
		})
	})

	Describe(`ListResourceAccessPolicies - List Resource Access Policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceAccessPolicies(listResourceAccessPoliciesOptions *ListResourceAccessPoliciesOptions)`, func() {
			listResourceAccessPoliciesOptions := &watsonxdatav3.ListResourceAccessPoliciesOptions{
				AuthInstanceID: core.StringPtr("testString"),
				ResourceType: core.StringPtr("catalog"),
				ResourceID: []string{"testString"},
				ResourceName: []string{"testString"},
			}

			accessPolicies, response, err := watsonxDataService.ListResourceAccessPolicies(listResourceAccessPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessPolicies).ToNot(BeNil())
		})
	})

	Describe(`BulkUpdateResourceAccessPolicies - Bulk Update Resource Access Policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`BulkUpdateResourceAccessPolicies(bulkUpdateResourceAccessPoliciesOptions *BulkUpdateResourceAccessPoliciesOptions)`, func() {
			resourceDetailsBulkUpdateModel := &watsonxdatav3.ResourceDetailsBulkUpdate{
				ID: core.StringPtr("presto01"),
				Name: core.StringPtr("hive_data"),
				Type: core.StringPtr("presto"),
			}

			subjectModel := &watsonxdatav3.Subject{
				Type: core.StringPtr("user"),
				Value: core.StringPtr("user1"),
			}

			subjectBulkUpdateModel := &watsonxdatav3.SubjectBulkUpdate{
				Permissions: []string{"testString"},
				Subject: subjectModel,
			}

			accessPolicyBulkUpdateModel := &watsonxdatav3.AccessPolicyBulkUpdate{
				Resources: []watsonxdatav3.ResourceDetailsBulkUpdate{*resourceDetailsBulkUpdateModel},
				Subjects: []watsonxdatav3.SubjectBulkUpdate{*subjectBulkUpdateModel},
			}

			bulkUpdateResourceAccessPoliciesOptions := &watsonxdatav3.BulkUpdateResourceAccessPoliciesOptions{
				AccessPolicies: []watsonxdatav3.AccessPolicyBulkUpdate{*accessPolicyBulkUpdateModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			accessPolicyBulkUpdateResponse, response, err := watsonxDataService.BulkUpdateResourceAccessPolicies(bulkUpdateResourceAccessPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessPolicyBulkUpdateResponse).ToNot(BeNil())
		})
	})

	Describe(`RevokeResourceAccessPolicies - Revoke Resource Access Policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RevokeResourceAccessPolicies(revokeResourceAccessPoliciesOptions *RevokeResourceAccessPoliciesOptions)`, func() {
			resourceDetailsModel := &watsonxdatav3.ResourceDetails{
				ID: core.StringPtr("presto01"),
				Name: core.StringPtr("hive_data"),
				Type: core.StringPtr("presto"),
			}

			subjectModel := &watsonxdatav3.Subject{
				Type: core.StringPtr("user"),
				Value: core.StringPtr("user1"),
			}

			subjectRevokeModel := &watsonxdatav3.SubjectRevoke{
				Permissions: []string{"testString"},
				Subject: subjectModel,
			}

			revokeResourceAccessPoliciesOptions := &watsonxdatav3.RevokeResourceAccessPoliciesOptions{
				Resources: []watsonxdatav3.ResourceDetails{*resourceDetailsModel},
				Subjects: []watsonxdatav3.SubjectRevoke{*subjectRevokeModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.RevokeResourceAccessPolicies(revokeResourceAccessPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`FilterResourceAccessPoliciesOnUsersAndUsergroups - Filter Resource Access Policies On Users And User Groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`FilterResourceAccessPoliciesOnUsersAndUsergroups(filterResourceAccessPoliciesOnUsersAndUsergroupsOptions *FilterResourceAccessPoliciesOnUsersAndUsergroupsOptions)`, func() {
			resourceDetailsModel := &watsonxdatav3.ResourceDetails{
				ID: core.StringPtr("presto01"),
				Name: core.StringPtr("hive_data"),
				Type: core.StringPtr("presto"),
			}

			subjectModel := &watsonxdatav3.Subject{
				Type: core.StringPtr("user"),
				Value: core.StringPtr("user1"),
			}

			accessPoliciesSearchModel := &watsonxdatav3.AccessPoliciesSearch{
				Resources: []watsonxdatav3.ResourceDetails{*resourceDetailsModel},
				SubjectsToSearch: []watsonxdatav3.Subject{*subjectModel},
			}

			filterResourceAccessPoliciesOnUsersAndUsergroupsOptions := &watsonxdatav3.FilterResourceAccessPoliciesOnUsersAndUsergroupsOptions{
				AccessPoliciesSearch: []watsonxdatav3.AccessPoliciesSearch{*accessPoliciesSearchModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			accessPolicies, response, err := watsonxDataService.FilterResourceAccessPoliciesOnUsersAndUsergroups(filterResourceAccessPoliciesOnUsersAndUsergroupsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessPolicies).ToNot(BeNil())
		})
	})

	Describe(`ListDataPolicies - List data policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataPolicies(listDataPoliciesOptions *ListDataPoliciesOptions)`, func() {
			listDataPoliciesOptions := &watsonxdatav3.ListDataPoliciesOptions{
				AuthInstanceID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				Status: core.StringPtr("testString"),
				IncludeMetadata: core.BoolPtr(true),
				IncludeRules: core.BoolPtr(true),
				BucketName: core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				DataArtifact: core.StringPtr("testString"),
			}

			policyV2Collection, response, err := watsonxDataService.ListDataPolicies(listDataPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyV2Collection).ToNot(BeNil())
		})
	})

	Describe(`CreateDataPolicy - Create a new data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataPolicy(createDataPolicyOptions *CreateDataPolicyOptions)`, func() {
			ruleGranteeModel := &watsonxdatav3.RuleGrantee{
				Key: core.StringPtr("user_name"),
				Type: core.StringPtr("user_identity"),
				Value: core.StringPtr("user1"),
			}

			transformColumnPropertiesModel := &watsonxdatav3.TransformColumnProperties{
				MaskCondition: core.StringPtr("NIL"),
				MaskType: core.StringPtr("mask_show_last_4"),
				MaskValue: core.StringPtr("NIL"),
			}

			rowFilterPropertiesModel := &watsonxdatav3.RowFilterProperties{
				FilterCondition: core.StringPtr("NIL"),
				RowFilter: core.StringPtr("addr_country='US'"),
			}

			ruleV2Model := &watsonxdatav3.RuleV2{
				Actions: []string{"alter", "create"},
				Effect: core.StringPtr("allow"),
				Grantees: []watsonxdatav3.RuleGrantee{*ruleGranteeModel},
				TransformColumns: transformColumnPropertiesModel,
				TransformRows: rowFilterPropertiesModel,
			}

			createDataPolicyOptions := &watsonxdatav3.CreateDataPolicyOptions{
				DataArtifact: core.StringPtr("schema1/table1/(column1|column2)"),
				Rules: []watsonxdatav3.RuleV2{*ruleV2Model},
				CatalogName: core.StringPtr("catalog1"),
				CatalogType: core.StringPtr("catalog1"),
				Description: core.StringPtr("policy description"),
				PolicyName: core.StringPtr("policy1"),
				ResourceID: core.StringPtr("catalog1"),
				Status: core.StringPtr("active"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			policyV2, response, err := watsonxDataService.CreateDataPolicy(createDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyV2).ToNot(BeNil())
		})
	})

	Describe(`GetDataPolicy - Get data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataPolicy(getDataPolicyOptions *GetDataPolicyOptions)`, func() {
			getDataPolicyOptions := &watsonxdatav3.GetDataPolicyOptions{
				Name: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			policyV2, response, err := watsonxDataService.GetDataPolicy(getDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyV2).ToNot(BeNil())
		})
	})

	Describe(`ReplaceDataPolicy - Replace data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceDataPolicy(replaceDataPolicyOptions *ReplaceDataPolicyOptions)`, func() {
			ruleGranteeModel := &watsonxdatav3.RuleGrantee{
				Key: core.StringPtr("user_name"),
				Type: core.StringPtr("user_identity"),
				Value: core.StringPtr("user1"),
			}

			transformColumnPropertiesModel := &watsonxdatav3.TransformColumnProperties{
				MaskCondition: core.StringPtr("NIL"),
				MaskType: core.StringPtr("mask_show_last_4"),
				MaskValue: core.StringPtr("NIL"),
			}

			rowFilterPropertiesModel := &watsonxdatav3.RowFilterProperties{
				FilterCondition: core.StringPtr("NIL"),
				RowFilter: core.StringPtr("addr_country='US'"),
			}

			ruleV2Model := &watsonxdatav3.RuleV2{
				Actions: []string{"alter", "create"},
				Effect: core.StringPtr("allow"),
				Grantees: []watsonxdatav3.RuleGrantee{*ruleGranteeModel},
				TransformColumns: transformColumnPropertiesModel,
				TransformRows: rowFilterPropertiesModel,
			}

			replaceDataPolicyOptions := &watsonxdatav3.ReplaceDataPolicyOptions{
				Name: core.StringPtr("testString"),
				DataArtifact: core.StringPtr("schema1/table1/(column1|column2)"),
				Rules: []watsonxdatav3.RuleV2{*ruleV2Model},
				CatalogName: core.StringPtr("catalog1"),
				CatalogType: core.StringPtr("catalog1"),
				Description: core.StringPtr("policy description"),
				PolicyName: core.StringPtr("policy1"),
				ResourceID: core.StringPtr("catalog1"),
				Status: core.StringPtr("active"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			policyV2, response, err := watsonxDataService.ReplaceDataPolicy(replaceDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyV2).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataPolicy - Update data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataPolicy(updateDataPolicyOptions *UpdateDataPolicyOptions)`, func() {
			jsonPatchOperationModel := &watsonxdatav3.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: "testString",
			}

			updateDataPolicyOptions := &watsonxdatav3.UpdateDataPolicyOptions{
				Name: core.StringPtr("testString"),
				Body: []watsonxdatav3.JSONPatchOperation{*jsonPatchOperationModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			policyV2, response, err := watsonxDataService.UpdateDataPolicy(updateDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyV2).ToNot(BeNil())
		})
	})

	Describe(`DeleteStorageRegistration - Unregister Storage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteStorageRegistration(deleteStorageRegistrationOptions *DeleteStorageRegistrationOptions)`, func() {
			deleteStorageRegistrationOptions := &watsonxdatav3.DeleteStorageRegistrationOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				SkipMdsCall: core.BoolPtr(false),
			}

			response, err := watsonxDataService.DeleteStorageRegistration(deleteStorageRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDatabaseCatalog - Delete database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
			deleteDatabaseCatalogOptions := &watsonxdatav3.DeleteDatabaseCatalogOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestoEngineCatalogs - Disassociate catalogs from a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions)`, func() {
			deletePrestoEngineCatalogsOptions := &watsonxdatav3.DeletePrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteEngine - Delete presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
			deleteEngineOptions := &watsonxdatav3.DeleteEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteEngine(deleteEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestissimoEngine - Delete prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestissimoEngine(deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions)`, func() {
			deletePrestissimoEngineOptions := &watsonxdatav3.DeletePrestissimoEngineOptions{
				ID: core.StringPtr("testString"),
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
			deletePrestissimoEngineCatalogsOptions := &watsonxdatav3.DeletePrestissimoEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDb2Engine - Delete db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions)`, func() {
			deleteDb2EngineOptions := &watsonxdatav3.DeleteDb2EngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteOtherEngine - Delete engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions)`, func() {
			deleteOtherEngineOptions := &watsonxdatav3.DeleteOtherEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteNetezzaEngine - Delete netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions)`, func() {
			deleteNetezzaEngineOptions := &watsonxdatav3.DeleteNetezzaEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngine - Delete spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions)`, func() {
			deleteSparkEngineOptions := &watsonxdatav3.DeleteSparkEngineOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineCatalogs - Disassociate catalogs from a spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineCatalogs(deleteSparkEngineCatalogsOptions *DeleteSparkEngineCatalogsOptions)`, func() {
			deleteSparkEngineCatalogsOptions := &watsonxdatav3.DeleteSparkEngineCatalogsOptions{
				ID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineCatalogs(deleteSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineApplication - Stop Spark Applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineApplication(deleteSparkEngineApplicationOptions *DeleteSparkEngineApplicationOptions)`, func() {
			deleteSparkEngineApplicationOptions := &watsonxdatav3.DeleteSparkEngineApplicationOptions{
				EngineID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineApplication(deleteSparkEngineApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineHistoryServer - Stop spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineHistoryServer(deleteSparkEngineHistoryServerOptions *DeleteSparkEngineHistoryServerOptions)`, func() {
			deleteSparkEngineHistoryServerOptions := &watsonxdatav3.DeleteSparkEngineHistoryServerOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineHistoryServer(deleteSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteIntegration - Delete an Integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIntegration(deleteIntegrationOptions *DeleteIntegrationOptions)`, func() {
			deleteIntegrationOptions := &watsonxdatav3.DeleteIntegrationOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteIntegration(deleteIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSchema - Delete schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
			deleteSchemaOptions := &watsonxdatav3.DeleteSchemaOptions{
				EngineID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
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
			deleteTableOptions := &watsonxdatav3.DeleteTableOptions{
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
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
			deleteColumnOptions := &watsonxdatav3.DeleteColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogName: core.StringPtr("testString"),
				SchemaName: core.StringPtr("testString"),
				TableName: core.StringPtr("testString"),
				ColumnName: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteColumn(deleteColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteCatalog - Delete catalog catalog_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions)`, func() {
			deleteCatalogOptions := &watsonxdatav3.DeleteCatalogOptions{
				Name: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				SkipMdsCall: core.BoolPtr(false),
			}

			response, err := watsonxDataService.DeleteCatalog(deleteCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteMilvusService - Delete milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteMilvusService(deleteMilvusServiceOptions *DeleteMilvusServiceOptions)`, func() {
			deleteMilvusServiceOptions := &watsonxdatav3.DeleteMilvusServiceOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteMilvusService(deleteMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSalIntegration - Delete Semantic Automation Layer(SAL) integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSalIntegration(deleteSalIntegrationOptions *DeleteSalIntegrationOptions)`, func() {
			deleteSalIntegrationOptions := &watsonxdatav3.DeleteSalIntegrationOptions{
			}

			response, err := watsonxDataService.DeleteSalIntegration(deleteSalIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSemanticSearchQueries - Clear semantic search query history`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSemanticSearchQueries(deleteSemanticSearchQueriesOptions *DeleteSemanticSearchQueriesOptions)`, func() {
			deleteSemanticSearchQueriesOptions := &watsonxdatav3.DeleteSemanticSearchQueriesOptions{
				BatchSize: core.Int64Ptr(int64(1)),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSemanticSearchQueries(deleteSemanticSearchQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSemanticSearchQueriesByID - Delete semantic search query history by record id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSemanticSearchQueriesByID(deleteSemanticSearchQueriesByIdOptions *DeleteSemanticSearchQueriesByIdOptions)`, func() {
			deleteSemanticSearchQueriesByIdOptions := &watsonxdatav3.DeleteSemanticSearchQueriesByIdOptions{
				ID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSemanticSearchQueriesByID(deleteSemanticSearchQueriesByIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSalMetadata - Delete Semantic automation layer(SAL) integration metadata`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSalMetadata(deleteSalMetadataOptions *DeleteSalMetadataOptions)`, func() {
			deleteSalMetadataOptions := &watsonxdatav3.DeleteSalMetadataOptions{
			}

			response, err := watsonxDataService.DeleteSalMetadata(deleteSalMetadataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDataPolicies - Bulk delete data policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataPolicies(deleteDataPoliciesOptions *DeleteDataPoliciesOptions)`, func() {
			deleteDataPoliciesOptions := &watsonxdatav3.DeleteDataPoliciesOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Policies: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDataPolicies(deleteDataPoliciesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDataPolicy - Delete a data policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataPolicy(deleteDataPolicyOptions *DeleteDataPolicyOptions)`, func() {
			deleteDataPolicyOptions := &watsonxdatav3.DeleteDataPolicyOptions{
				Name: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDataPolicy(deleteDataPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
