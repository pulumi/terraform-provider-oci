// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	loadBalancerShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	LoadBalancerShapeResourceConfig = ""
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_load_balancer_shapes.test_load_balancer_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_shapes", "test_load_balancer_shapes", Required, Create, loadBalancerShapeDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
			),
		},
	})
}
