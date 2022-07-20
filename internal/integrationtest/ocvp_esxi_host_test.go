// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	tf_client "terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/resourcediscovery"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"
)

var (
	OcvpEsxiHostRequiredOnlyResource = EsxiHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostRepresentation)

	OcvpEsxiHostResourceConfig = EsxiHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Update, OcvpEsxiHostRepresentation)

	ReplacementEsxiHostResourceConfig = EsxiHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation)

	OcvpOcvpEsxiHostSingularDataSourceRepresentation = map[string]interface{}{
		"esxi_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_esxi_host.test_esxi_host.id}`},
	}
	OcvpOcvpEsxiHostDataSourceRepresentation = map[string]interface{}{
		"compute_instance_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_esxi_host.test_esxi_host.compute_instance_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sddc_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpEsxiHostDataSourceFilterRepresentation}}
	OcvpEsxiHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_esxi_host.test_esxi_host.id}`}},
	}
	existingEsxiHostDataSourceRepresentation = map[string]interface{}{
		"sddc_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.display_name}-1`},
	}

	OcvpEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"capacity_reservation_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"compute_availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"current_sku":                 acctest.Representation{RepType: acctest.Optional, Create: `MONTH`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"next_sku":                    acctest.Representation{RepType: acctest.Optional, Create: `MONTH`, Update: `HOUR`},
		"host_ocpu_count":             acctest.Representation{RepType: acctest.Optional, Create: `52`},
		"host_shape_name":             acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO2.52`},
	}
	replacementEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"failed_esxi_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_esxi_hosts.existing_esxi_hosts.esxi_host_collection[0].id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `replacement`},
	}

	EsxiHostResourceDependencies = OcvpSddcRequiredOnlyResource + acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "existing_esxi_hosts", acctest.Optional, acctest.Create, existingEsxiHostDataSourceRepresentation)
)

// issue-routing-tag: ocvp/default
func TestOcvpEsxiHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpEsxiHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ocvp_esxi_host.test_esxi_host"
	datasourceName := "data.oci_ocvp_esxi_hosts.test_esxi_hosts"
	singularDatasourceName := "data.oci_ocvp_esxi_host.test_esxi_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EsxiHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostRepresentation), "ocvp", "esxiHost", t)

	acctest.ResourceTest(t, testAccCheckOcvpEsxiHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "current_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "host_ocpu_count"),
				resource.TestCheckResourceAttrSet(resourceName, "host_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies,
		},
		// verify replace node
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "display_name", `replacement`),
				resource.TestCheckResourceAttrSet(resourceName, "failed_esxi_host_id"),
				resource.TestCheckResourceAttrSet(resourceName, "host_ocpu_count"),
				resource.TestCheckResourceAttrSet(resourceName, "host_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// verify singular datasource for replace node
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpOcvpEsxiHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ReplacementEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", `replacement`),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grace_period_end_date"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_shape_name", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_sku", "MONTH"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// delete replace node before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "MONTH"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Update, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "HOUR"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "test_esxi_hosts", acctest.Optional, acctest.Update, OcvpOcvpEsxiHostDataSourceRepresentation) +
				compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Update, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_shape_name", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpOcvpEsxiHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_shape_name", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_sku", "HOUR"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpEsxiHostRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpEsxiHostDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EsxiHostClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_esxi_host" {
			noResourceFound = false
			request := oci_ocvp.GetEsxiHostRequest{}

			tmp := rs.Primary.ID
			request.EsxiHostId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetEsxiHost(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpEsxiHost") {
		resource.AddTestSweepers("OcvpEsxiHost", &resource.Sweeper{
			Name:         "OcvpEsxiHost",
			Dependencies: acctest.DependencyGraph["esxiHost"],
			F:            sweepOcvpEsxiHostResource,
		})
	}
}

func sweepOcvpEsxiHostResource(compartment string) error {
	esxiHostClient := acctest.GetTestClients(&schema.ResourceData{}).EsxiHostClient()
	esxiHostIds, err := getOcvpEsxiHostIds(compartment)
	if err != nil {
		return err
	}
	for _, esxiHostId := range esxiHostIds {
		if ok := acctest.SweeperDefaultResourceId[esxiHostId]; !ok {
			deleteEsxiHostRequest := oci_ocvp.DeleteEsxiHostRequest{}

			deleteEsxiHostRequest.EsxiHostId = &esxiHostId

			deleteEsxiHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := esxiHostClient.DeleteEsxiHost(context.Background(), deleteEsxiHostRequest)
			if error != nil {
				fmt.Printf("Error deleting EsxiHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", esxiHostId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &esxiHostId, OcvpEsxiHostSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpEsxiHostSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpEsxiHostIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EsxiHostId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	esxiHostClient := acctest.GetTestClients(&schema.ResourceData{}).EsxiHostClient()

	listEsxiHostsRequest := oci_ocvp.ListEsxiHostsRequest{}
	listEsxiHostsRequest.LifecycleState = oci_ocvp.ListEsxiHostsLifecycleStateActive
	listEsxiHostsResponse, err := esxiHostClient.ListEsxiHosts(context.Background(), listEsxiHostsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EsxiHost list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, esxiHost := range listEsxiHostsResponse.Items {
		id := *esxiHost.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EsxiHostId", id)
	}
	return resourceIds, nil
}

func OcvpEsxiHostSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if esxiHostResponse, ok := response.Response.(oci_ocvp.GetEsxiHostResponse); ok {
		return esxiHostResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpEsxiHostSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EsxiHostClient().GetEsxiHost(context.Background(), oci_ocvp.GetEsxiHostRequest{
		EsxiHostId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
