// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/resourcediscovery"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"terraform-provider-oci/httpreplay"
)

var (
	BackupRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, DatabaseBackupRepresentation)

	DatabaseDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseBackupDataSourceFilterRepresentation}}
	DatabaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_backup.test_backup.id}`}},
	}

	DatabaseBackupRepresentation = map[string]interface{}{
		"database_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `Monthly Backup`},
	}

	DatabaseDatabaseBackupResourceDependencies = DbSystemResourceConfig + `
data "oci_database_databases" "db" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}`
)

// issue-routing-tag: database/default
func TestDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup.test_backup"
	datasourceName := "data.oci_database_backups.test_backups"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDatabaseBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, DatabaseBackupRepresentation), "database", "backup", t)

	acctest.ResourceTest(t, testAccCheckDatabaseBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdVariableStr + DatabaseDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, DatabaseBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", acctest.Optional, acctest.Update, DatabaseDatabaseBackupDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Optional, acctest.Update, DatabaseBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

				resource.TestCheckResourceAttr(datasourceName, "backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_edition"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.display_name", "Monthly Backup"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.vault_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.version"),
			),
		},
		// verify resource import
		{
			Config:            config + BackupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
				"database_size_in_gbs",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_backup" {
			noResourceFound = false
			request := oci_database.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.BackupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseBackup") {
		resource.AddTestSweepers("DatabaseBackup", &resource.Sweeper{
			Name:         "DatabaseBackup",
			Dependencies: acctest.DependencyGraph["backup"],
			F:            sweepDatabaseBackupResource,
		})
	}
}

func sweepDatabaseBackupResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	backupIds, err := getDatabaseBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, backupId := range backupIds {
		if ok := acctest.SweeperDefaultResourceId[backupId]; !ok {
			deleteBackupRequest := oci_database.DeleteBackupRequest{}

			deleteBackupRequest.BackupId = &backupId

			deleteBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteBackup(context.Background(), deleteBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting Backup %s %s, It is possible that the resource is already deleted. Please verify manually \n", backupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &backupId, DatabaseBackupSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseBackupSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listBackupsRequest := oci_database.ListBackupsRequest{}
	listBackupsRequest.CompartmentId = &compartmentId
	listBackupsResponse, err := databaseClient.ListBackups(context.Background(), listBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Backup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, backup := range listBackupsResponse.Items {
		id := *backup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackupId", id)
	}
	return resourceIds, nil
}

func DatabaseBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if backupResponse, ok := response.Response.(oci_database.GetBackupResponse); ok {
		return backupResponse.LifecycleState != oci_database.BackupLifecycleStateDeleted
	}
	return false
}

func DatabaseBackupSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetBackup(context.Background(), oci_database.GetBackupRequest{
		BackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
