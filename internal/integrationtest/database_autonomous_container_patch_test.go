// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousContainerPatchDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatabaseAutonomousContainerPatchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation) +
		DatabaseAutonomousExadataInfrastructureResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousContainerPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_patches.test_autonomous_container_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_patches", "test_autonomous_container_patches", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerPatchDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.patch_model"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.quarter"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.year"),
			),
		},
	})
}
