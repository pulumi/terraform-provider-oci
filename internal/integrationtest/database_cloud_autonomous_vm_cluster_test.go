// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CloudAutonomousVmClusterRequiredOnlyResource = CloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, cloudAutonomousVmClusterRepresentation)

	CloudAutonomousVmClusterResourceConfig = CloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Update, cloudAutonomousVmClusterRepresentation)

	cloudAutonomousVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
	}

	cloudAutonomousVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName1`, Update: `displayName2`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudAutonomousVmClusterDataSourceFilterRepresentation}}
	cloudAutonomousVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`}},
	}

	cloudAutonomousVmClusterRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "UpdatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `CloudAutonomousVmCluster`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}},
	}

	cloudAutonomousNetworkSecurityGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNsgRepresentation},
	}

	CloudAutonomousVmClusterResourceDependencies = DefinedTagsDependencies + AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, cloudExadataInfrastructureRepresentation) +
		`
	#dataguard requires the port to be open on the subnet
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}
	data "oci_identity_availability_domain" "ad" {
		compartment_id 		= "${var.compartment_id}"
		ad_number      		= 1
	}
	resource "oci_core_subnet" "exadata_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.22.0/24"
		display_name        = "ExadataSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}", "${oci_core_security_list.exadata_shapes_security_list.id}"]
		dns_label           = "subnetexadata"
	}

	resource "oci_core_subnet" "exadata_backup_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.23.0/24"
		display_name        = "ExadataBackupSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.ExampleRT.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "subnetexadata2"
	}

	resource "oci_core_internet_gateway" "ExampleIG" {
		compartment_id = "${var.compartment_id}"
		display_name   = "TFExampleIG"
		vcn_id         = "${oci_core_virtual_network.t.id}"
	}

	resource "oci_core_route_table" "ExampleRT" {
		compartment_id = "${var.compartment_id}"
		vcn_id         = "${oci_core_virtual_network.t.id}"
		display_name   = "TFExampleRouteTable"

		route_rules {
			destination       = "0.0.0.0/0"
			destination_type  = "CIDR_BLOCK"
			network_entity_id = "${oci_core_internet_gateway.ExampleIG.id}"
		}
	}

	resource "oci_core_security_list" "exadata_shapes_security_list" {
		compartment_id = "${var.compartment_id}"
		vcn_id         = "${oci_core_virtual_network.t.id}"
		display_name   = "ExadataSecurityList"

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "6"
		}

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "1"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "6"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "1"
		}
	}
`
)

func TestDatabaseCloudAutonomousVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudAutonomousVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	var resId string
	resourceName := "oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster"

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	datasourceName := "data.oci_database_cloud_autonomous_vm_clusters.test_cloud_autonomous_vm_clusters"
	singularDatasourceName := "data.oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster"

	var resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseCloudAutonomousVmClusterDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config +
					compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, cloudAutonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "CloudAutonomousVmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						print(resId)
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, cloudAutonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "CloudAutonomousVmCluster"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(cloudAutonomousVmClusterRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "CloudAutonomousVmCluster"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource reCreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Update, cloudAutonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource reCreated when it was supposed to be Updated.")
						}
						return err
					},
				),
			},
			{
				Config: config + compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(cloudAutonomousVmClusterRepresentation, map[string]interface{}{})),
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource reCreated when it was supposed to be Updated.")
						}
						return err
					},
				),
			},
			{
				Config: config + compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(cloudAutonomousVmClusterRepresentation, map[string]interface{}{})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource reCreated when it was supposed to be Updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_autonomous_vm_clusters", "test_cloud_autonomous_vm_clusters", acctest.Optional, acctest.Update, cloudAutonomousVmClusterDataSourceRepresentation) +
					compartmentIdVariableStr + CloudAutonomousVmClusterResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Update, cloudAutonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.cpu_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.data_storage_size_in_gb"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.domain"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.hostname"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "cloud_autonomous_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.ocpu_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.shape"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_clusters.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, cloudAutonomousVmClusterSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CloudAutonomousVmClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_autonomous_vm_cluster_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_gb"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ocpu_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},

			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CloudAutonomousVmClusterResourceConfig,
			},

			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseCloudAutonomousVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_cloud_autonomous_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetCloudAutonomousVmClusterRequest{}

			tmp := rs.Primary.ID
			request.CloudAutonomousVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetCloudAutonomousVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.CloudAutonomousVmClusterLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseCloudAutonomousVmCluster") {
		resource.AddTestSweepers("DatabaseCloudAutonomousVmCluster", &resource.Sweeper{
			Name:         "DatabaseCloudAutonomousVmCluster",
			Dependencies: acctest.DependencyGraph["cloudAutonomousVmCluster"],
			F:            sweepDatabaseCloudAutonomousVmClusterResource,
		})
	}
}

func sweepDatabaseCloudAutonomousVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	cloudAutonomousVmClusterIds, err := getCloudAutonomousVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudAutonomousVmClusterId := range cloudAutonomousVmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[cloudAutonomousVmClusterId]; !ok {
			deleteCloudAutonomousVmClusterRequest := oci_database.DeleteCloudAutonomousVmClusterRequest{}

			deleteCloudAutonomousVmClusterRequest.CloudAutonomousVmClusterId = &cloudAutonomousVmClusterId

			deleteCloudAutonomousVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteCloudAutonomousVmCluster(context.Background(), deleteCloudAutonomousVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudAutonomousVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudAutonomousVmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudAutonomousVmClusterId, cloudAutonomousVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				cloudAutonomousVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getCloudAutonomousVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudAutonomousVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listCloudAutonomousVmClustersRequest := oci_database.ListCloudAutonomousVmClustersRequest{}
	listCloudAutonomousVmClustersRequest.CompartmentId = &compartmentId
	listCloudAutonomousVmClustersRequest.LifecycleState = oci_database.CloudAutonomousVmClusterSummaryLifecycleStateAvailable
	listCloudAutonomousVmClustersResponse, err := databaseClient.ListCloudAutonomousVmClusters(context.Background(), listCloudAutonomousVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudAutonomousVmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudAutonomousVmCluster := range listCloudAutonomousVmClustersResponse.Items {
		id := *cloudAutonomousVmCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudAutonomousVmClusterId", id)
	}
	return resourceIds, nil
}

func cloudAutonomousVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudAutonomousVmClusterResponse, ok := response.Response.(oci_database.GetCloudAutonomousVmClusterResponse); ok {
		return cloudAutonomousVmClusterResponse.LifecycleState != oci_database.CloudAutonomousVmClusterLifecycleStateTerminated
	}
	return false
}

func cloudAutonomousVmClusterSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetCloudAutonomousVmCluster(context.Background(), oci_database.GetCloudAutonomousVmClusterRequest{
		CloudAutonomousVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
