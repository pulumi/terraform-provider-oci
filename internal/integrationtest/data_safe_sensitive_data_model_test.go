// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"terraform-provider-oci/internal/resourcediscovery"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/client"
	tf_client "terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"terraform-provider-oci/httpreplay"
)

var (
	DataSafeSensitiveDataModelRequiredOnlyResource = DataSafeSensitiveDataModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)

	DataSafeSensitiveDataModelResourceConfig = DataSafeSensitiveDataModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Update, sensitiveDataModelRepresentation)

	DataSafesensitiveDataModelSingularDataSourceRepresentation = map[string]interface{}{
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
	}

	DataSafesensitiveDataModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sensitive_data_model_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"time_created_less_than":    acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: sensitiveDataModelDataSourceFilterRepresentation}}
	sensitiveDataModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`}},
	}

	sensitiveDataModelRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"app_suite_name": acctest.Representation{RepType: acctest.Optional, Create: `appSuiteName`, Update: `appSuiteName2`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_app_defined_relation_discovery_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_sample_data_collection_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"schemas_for_discovery":                     acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{}},
		"sensitive_type_ids_for_discovery":          acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{}},
	}

	DataSafeSensitiveDataModelResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveDataModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_sensitive_data_model.test_sensitive_data_model"
	datasourceName := "data.oci_data_safe_sensitive_data_models.test_sensitive_data_models"
	singularDatasourceName := "data.oci_data_safe_sensitive_data_model.test_sensitive_data_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+DataSafeSensitiveDataModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Create, sensitiveDataModelRepresentation), "datasafe", "sensitiveDataModel", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveDataModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveDataModelResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveDataModelResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveDataModelResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Create, sensitiveDataModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app_suite_name", "appSuiteName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_app_defined_relation_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sample_data_collection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas_for_discovery.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "sensitive_type_ids_for_discovery.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSensitiveDataModelResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(sensitiveDataModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app_suite_name", "appSuiteName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_app_defined_relation_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sample_data_collection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schemas_for_discovery.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "sensitive_type_ids_for_discovery.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveDataModelResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Update, sensitiveDataModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "app_suite_name", "appSuiteName2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "schemas_for_discovery.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "sensitive_type_ids_for_discovery.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_models", "test_sensitive_data_models", acctest.Optional, acctest.Update, DataSafesensitiveDataModelDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Optional, acctest.Update, sensitiveDataModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "sensitive_data_model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_data_model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, DataSafesensitiveDataModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_data_model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "app_suite_name", "appSuiteName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_app_defined_relation_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sample_data_collection_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas_for_discovery.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sensitive_type_ids_for_discovery.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + targetIdVariableStr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveDataModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_data_model" {
			noResourceFound = false
			request := oci_data_safe.GetSensitiveDataModelRequest{}

			tmp := rs.Primary.ID
			request.SensitiveDataModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSensitiveDataModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.DiscoveryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveDataModel") {
		resource.AddTestSweepers("DataSafeSensitiveDataModel", &resource.Sweeper{
			Name:         "DataSafeSensitiveDataModel",
			Dependencies: acctest.DependencyGraph["sensitiveDataModel"],
			F:            sweepDataSafeSensitiveDataModelResource,
		})
	}
}

func sweepDataSafeSensitiveDataModelResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveDataModelIds, err := getDataSafeSensitiveDataModelIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveDataModelId := range sensitiveDataModelIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveDataModelId]; !ok {
			deleteSensitiveDataModelRequest := oci_data_safe.DeleteSensitiveDataModelRequest{}

			deleteSensitiveDataModelRequest.SensitiveDataModelId = &sensitiveDataModelId

			deleteSensitiveDataModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSensitiveDataModel(context.Background(), deleteSensitiveDataModelRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveDataModel %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveDataModelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sensitiveDataModelId, DataSafesensitiveDataModelsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafesensitiveDataModelsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSensitiveDataModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveDataModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSensitiveDataModelsRequest := oci_data_safe.ListSensitiveDataModelsRequest{}
	listSensitiveDataModelsRequest.CompartmentId = &compartmentId
	listSensitiveDataModelsRequest.LifecycleState = oci_data_safe.ListSensitiveDataModelsLifecycleStateActive
	listSensitiveDataModelsResponse, err := dataSafeClient.ListSensitiveDataModels(context.Background(), listSensitiveDataModelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SensitiveDataModel list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sensitiveDataModel := range listSensitiveDataModelsResponse.Items {
		id := *sensitiveDataModel.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveDataModelId", id)
	}
	return resourceIds, nil
}

func DataSafesensitiveDataModelsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sensitiveDataModelResponse, ok := response.Response.(oci_data_safe.GetSensitiveDataModelResponse); ok {
		return sensitiveDataModelResponse.LifecycleState != oci_data_safe.DiscoveryLifecycleStateDeleted
	}
	return false
}

func DataSafesensitiveDataModelsSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSensitiveDataModel(context.Background(), oci_data_safe.GetSensitiveDataModelRequest{
		SensitiveDataModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
