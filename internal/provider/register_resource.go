// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	tf_adm "terraform-provider-oci/internal/service/adm"
	tf_ai_anomaly_detection "terraform-provider-oci/internal/service/ai_anomaly_detection"
	tf_ai_vision "terraform-provider-oci/internal/service/ai_vision"
	tf_analytics "terraform-provider-oci/internal/service/analytics"
	tf_apigateway "terraform-provider-oci/internal/service/apigateway"
	tf_apm "terraform-provider-oci/internal/service/apm"
	tf_apm_config "terraform-provider-oci/internal/service/apm_config"
	tf_apm_synthetics "terraform-provider-oci/internal/service/apm_synthetics"
	tf_appmgmt_control "terraform-provider-oci/internal/service/appmgmt_control"
	tf_artifacts "terraform-provider-oci/internal/service/artifacts"
	tf_audit "terraform-provider-oci/internal/service/audit"
	tf_autoscaling "terraform-provider-oci/internal/service/autoscaling"
	tf_bastion "terraform-provider-oci/internal/service/bastion"
	tf_bds "terraform-provider-oci/internal/service/bds"
	tf_blockchain "terraform-provider-oci/internal/service/blockchain"
	tf_budget "terraform-provider-oci/internal/service/budget"
	tf_certificates_management "terraform-provider-oci/internal/service/certificates_management"
	tf_cloud_guard "terraform-provider-oci/internal/service/cloud_guard"
	tf_containerengine "terraform-provider-oci/internal/service/containerengine"
	tf_core "terraform-provider-oci/internal/service/core"
	tf_data_connectivity "terraform-provider-oci/internal/service/data_connectivity"
	tf_data_labeling_service "terraform-provider-oci/internal/service/data_labeling_service"
	tf_data_safe "terraform-provider-oci/internal/service/data_safe"
	tf_database "terraform-provider-oci/internal/service/database"
	tf_database_management "terraform-provider-oci/internal/service/database_management"
	tf_database_migration "terraform-provider-oci/internal/service/database_migration"
	tf_database_tools "terraform-provider-oci/internal/service/database_tools"
	tf_datacatalog "terraform-provider-oci/internal/service/datacatalog"
	tf_dataflow "terraform-provider-oci/internal/service/dataflow"
	tf_dataintegration "terraform-provider-oci/internal/service/dataintegration"
	tf_datascience "terraform-provider-oci/internal/service/datascience"
	tf_devops "terraform-provider-oci/internal/service/devops"
	tf_dns "terraform-provider-oci/internal/service/dns"
	tf_em_warehouse "terraform-provider-oci/internal/service/em_warehouse"
	tf_email "terraform-provider-oci/internal/service/email"
	tf_events "terraform-provider-oci/internal/service/events"
	tf_file_storage "terraform-provider-oci/internal/service/file_storage"
	tf_functions "terraform-provider-oci/internal/service/functions"
	tf_generic_artifacts_content "terraform-provider-oci/internal/service/generic_artifacts_content"
	tf_golden_gate "terraform-provider-oci/internal/service/golden_gate"
	tf_health_checks "terraform-provider-oci/internal/service/health_checks"
	tf_identity "terraform-provider-oci/internal/service/identity"
	tf_identity_data_plane "terraform-provider-oci/internal/service/identity_data_plane"
	tf_integration "terraform-provider-oci/internal/service/integration"
	tf_jms "terraform-provider-oci/internal/service/jms"
	tf_kms "terraform-provider-oci/internal/service/kms"
	tf_license_manager "terraform-provider-oci/internal/service/license_manager"
	tf_limits "terraform-provider-oci/internal/service/limits"
	tf_load_balancer "terraform-provider-oci/internal/service/load_balancer"
	tf_log_analytics "terraform-provider-oci/internal/service/log_analytics"
	tf_logging "terraform-provider-oci/internal/service/logging"
	tf_management_agent "terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "terraform-provider-oci/internal/service/marketplace"
	tf_metering_computation "terraform-provider-oci/internal/service/metering_computation"
	tf_monitoring "terraform-provider-oci/internal/service/monitoring"
	tf_mysql "terraform-provider-oci/internal/service/mysql"
	tf_network_load_balancer "terraform-provider-oci/internal/service/network_load_balancer"
	tf_nosql "terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "terraform-provider-oci/internal/service/objectstorage"
	tf_oce "terraform-provider-oci/internal/service/oce"
	tf_ocvp "terraform-provider-oci/internal/service/ocvp"
	tf_oda "terraform-provider-oci/internal/service/oda"
	tf_ons "terraform-provider-oci/internal/service/ons"
	tf_operator_access_control "terraform-provider-oci/internal/service/operator_access_control"
	tf_opsi "terraform-provider-oci/internal/service/opsi"
	tf_optimizer "terraform-provider-oci/internal/service/optimizer"
	tf_osmanagement "terraform-provider-oci/internal/service/osmanagement"
	tf_osp_gateway "terraform-provider-oci/internal/service/osp_gateway"
	tf_resourcemanager "terraform-provider-oci/internal/service/resourcemanager"
	tf_sch "terraform-provider-oci/internal/service/sch"
	tf_service_catalog "terraform-provider-oci/internal/service/service_catalog"
	tf_service_mesh "terraform-provider-oci/internal/service/service_mesh"
	tf_stack_monitoring "terraform-provider-oci/internal/service/stack_monitoring"
	tf_streaming "terraform-provider-oci/internal/service/streaming"
	tf_usage_proxy "terraform-provider-oci/internal/service/usage_proxy"
	tf_vault "terraform-provider-oci/internal/service/vault"
	tf_visual_builder "terraform-provider-oci/internal/service/visual_builder"
	tf_vn_monitoring "terraform-provider-oci/internal/service/vn_monitoring"
	tf_vulnerability_scanning "terraform-provider-oci/internal/service/vulnerability_scanning"
	tf_waa "terraform-provider-oci/internal/service/waa"
	tf_waas "terraform-provider-oci/internal/service/waas"
	tf_waf "terraform-provider-oci/internal/service/waf"
)

func init() {
	// adm service
	RegisterResource("oci_adm_knowledge_base", tf_adm.AdmKnowledgeBaseResource())
	RegisterResource("oci_adm_vulnerability_audit", tf_adm.AdmVulnerabilityAuditResource())
	// ai_anomaly_detection service
	RegisterResource("oci_ai_anomaly_detection_ai_private_endpoint", tf_ai_anomaly_detection.AiAnomalyDetectionAiPrivateEndpointResource())
	RegisterResource("oci_ai_anomaly_detection_data_asset", tf_ai_anomaly_detection.AiAnomalyDetectionDataAssetResource())
	RegisterResource("oci_ai_anomaly_detection_model", tf_ai_anomaly_detection.AiAnomalyDetectionModelResource())
	RegisterResource("oci_ai_anomaly_detection_project", tf_ai_anomaly_detection.AiAnomalyDetectionProjectResource())
	// analytics service
	RegisterResource("oci_analytics_analytics_instance", tf_analytics.AnalyticsAnalyticsInstanceResource())
	RegisterResource("oci_analytics_analytics_instance_private_access_channel", tf_analytics.AnalyticsAnalyticsInstancePrivateAccessChannelResource())
	RegisterResource("oci_analytics_analytics_instance_vanity_url", tf_analytics.AnalyticsAnalyticsInstanceVanityUrlResource())
	// apigateway service
	RegisterResource("oci_apigateway_api", tf_apigateway.ApigatewayApiResource())
	RegisterResource("oci_apigateway_certificate", tf_apigateway.ApigatewayCertificateResource())
	RegisterResource("oci_apigateway_deployment", tf_apigateway.ApigatewayDeploymentResource())
	RegisterResource("oci_apigateway_gateway", tf_apigateway.ApigatewayGatewayResource())
	RegisterResource("oci_apigateway_subscriber", tf_apigateway.ApigatewaySubscriberResource())
	RegisterResource("oci_apigateway_usage_plan", tf_apigateway.ApigatewayUsagePlanResource())
	// apm_config service
	RegisterResource("oci_apm_apm_domain", tf_apm.ApmApmDomainResource())
	RegisterResource("oci_apm_config_config", tf_apm_config.ApmConfigConfigResource())

	//apm synthetics
	RegisterResource("oci_apm_synthetics_dedicated_vantage_point", tf_apm_synthetics.ApmSyntheticsDedicatedVantagePointResource())
	RegisterResource("oci_apm_synthetics_monitor", tf_apm_synthetics.ApmSyntheticsMonitorResource())
	RegisterResource("oci_apm_synthetics_script", tf_apm_synthetics.ApmSyntheticsScriptResource())
	// appmgmt_control service
	RegisterResource("oci_appmgmt_control_monitor_plugin_management", tf_appmgmt_control.AppmgmtControlMonitorPluginManagementResource())
	// artifacts service
	RegisterResource("oci_artifacts_container_configuration", tf_artifacts.ArtifactsContainerConfigurationResource())
	RegisterResource("oci_artifacts_container_image_signature", tf_artifacts.ArtifactsContainerImageSignatureResource())
	RegisterResource("oci_artifacts_container_repository", tf_artifacts.ArtifactsContainerRepositoryResource())
	RegisterResource("oci_artifacts_generic_artifact", tf_artifacts.ArtifactsGenericArtifactResource())
	RegisterResource("oci_artifacts_repository", tf_artifacts.ArtifactsRepositoryResource())
	// audit service
	RegisterResource("oci_audit_configuration", tf_audit.AuditConfigurationResource())
	// autoscaling service
	RegisterResource("oci_autoscaling_auto_scaling_configuration", tf_autoscaling.AutoScalingAutoScalingConfigurationResource())
	// bastion service
	RegisterResource("oci_bastion_bastion", tf_bastion.BastionBastionResource())
	RegisterResource("oci_bastion_session", tf_bastion.BastionSessionResource())
	// bds service
	RegisterResource("oci_bds_auto_scaling_configuration", tf_bds.BdsAutoScalingConfigurationResource())
	RegisterResource("oci_bds_bds_instance", tf_bds.BdsBdsInstanceResource())
	RegisterResource("oci_bds_bds_instance_api_key", tf_bds.BdsBdsInstanceApiKeyResource())
	RegisterResource("oci_bds_bds_instance_metastore_config", tf_bds.BdsBdsInstanceMetastoreConfigResource())
	// blockchain service
	RegisterResource("oci_blockchain_blockchain_platform", tf_blockchain.BlockchainBlockchainPlatformResource())
	RegisterResource("oci_blockchain_osn", tf_blockchain.BlockchainOsnResource())
	RegisterResource("oci_blockchain_peer", tf_blockchain.BlockchainPeerResource())
	// budget service
	RegisterResource("oci_budget_alert_rule", tf_budget.BudgetAlertRuleResource())
	RegisterResource("oci_budget_budget", tf_budget.BudgetBudgetResource())
	// certificates_management service
	RegisterResource("oci_certificates_management_ca_bundle", tf_certificates_management.CertificatesManagementCaBundleResource())
	RegisterResource("oci_certificates_management_certificate", tf_certificates_management.CertificatesManagementCertificateResource())
	RegisterResource("oci_certificates_management_certificate_authority", tf_certificates_management.CertificatesManagementCertificateAuthorityResource())

	//data_safe service
	RegisterResource("oci_data_safe_alert", tf_data_safe.DataSafeAlertResource())
	RegisterResource("oci_data_safe_audit_archive_retrieval", tf_data_safe.DataSafeAuditArchiveRetrievalResource())
	RegisterResource("oci_data_safe_audit_policy", tf_data_safe.DataSafeAuditPolicyResource())
	RegisterResource("oci_data_safe_audit_profile", tf_data_safe.DataSafeAuditProfileResource())
	RegisterResource("oci_data_safe_audit_trail", tf_data_safe.DataSafeAuditTrailResource())
	RegisterResource("oci_data_safe_compare_security_assessment", tf_data_safe.DataSafeCompareSecurityAssessmentResource())
	RegisterResource("oci_data_safe_compare_user_assessment", tf_data_safe.DataSafeCompareUserAssessmentResource())
	RegisterResource("oci_data_safe_data_safe_configuration", tf_data_safe.DataSafeDataSafeConfigurationResource())
	RegisterResource("oci_data_safe_data_safe_private_endpoint", tf_data_safe.DataSafeDataSafePrivateEndpointResource())
	RegisterResource("oci_data_safe_on_prem_connector", tf_data_safe.DataSafeOnPremConnectorResource())
	RegisterResource("oci_data_safe_report_definition", tf_data_safe.DataSafeReportDefinitionResource())
	RegisterResource("oci_data_safe_security_assessment", tf_data_safe.DataSafeSecurityAssessmentResource())
	RegisterResource("oci_data_safe_target_alert_policy_association", tf_data_safe.DataSafeTargetAlertPolicyAssociationResource())
	RegisterResource("oci_data_safe_set_security_assessment_baseline", tf_data_safe.DataSafeSetSecurityAssessmentBaselineResource())
	RegisterResource("oci_data_safe_set_user_assessment_baseline", tf_data_safe.DataSafeSetUserAssessmentBaselineResource())
	RegisterResource("oci_data_safe_target_database", tf_data_safe.DataSafeTargetDatabaseResource())
	RegisterResource("oci_data_safe_unset_security_assessment_baseline", tf_data_safe.DataSafeUnsetSecurityAssessmentBaselineResource())
	RegisterResource("oci_data_safe_unset_user_assessment_baseline", tf_data_safe.DataSafeUnsetUserAssessmentBaselineResource())
	RegisterResource("oci_data_safe_user_assessment", tf_data_safe.DataSafeUserAssessmentResource())
	RegisterResource("oci_data_safe_sensitive_type", tf_data_safe.DataSafeSensitiveTypeResource())
	RegisterResource("oci_data_safe_discovery_job", tf_data_safe.DataSafeDiscoveryJobResource())
	RegisterResource("oci_data_safe_discovery_jobs_result", tf_data_safe.DataSafeDiscoveryJobsResultResource())
	RegisterResource("oci_data_safe_library_masking_format", tf_data_safe.DataSafeLibraryMaskingFormatResource())
	RegisterResource("oci_data_safe_masking_policy", tf_data_safe.DataSafeMaskingPolicyResource())
	RegisterResource("oci_data_safe_mask_data", tf_data_safe.DataSafeMaskDataResource())
	RegisterResource("oci_data_safe_sensitive_data_model", tf_data_safe.DataSafeSensitiveDataModelResource())
	RegisterResource("oci_data_safe_sensitive_data_models_sensitive_column", tf_data_safe.DataSafeSensitiveDataModelsSensitiveColumnResource())
	RegisterResource("oci_data_safe_masking_policies_masking_column", tf_data_safe.DataSafeMaskingPoliciesMaskingColumnResource())
	RegisterResource("oci_data_safe_sensitive_data_models_sensitive_column", tf_data_safe.DataSafeSensitiveDataModelsSensitiveColumnResource())
	RegisterResource("oci_data_safe_add_sdm_columns", tf_data_safe.DataSafeAddColumnsFromSdmResource())
	RegisterResource("oci_data_safe_sensitive_data_models_apply_discovery_job_results", tf_data_safe.DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource())

	// cloud_guard service
	RegisterResource("oci_cloud_guard_cloud_guard_configuration", tf_cloud_guard.CloudGuardCloudGuardConfigurationResource())
	RegisterResource("oci_cloud_guard_data_mask_rule", tf_cloud_guard.CloudGuardDataMaskRuleResource())
	RegisterResource("oci_cloud_guard_detector_recipe", tf_cloud_guard.CloudGuardDetectorRecipeResource())
	RegisterResource("oci_cloud_guard_managed_list", tf_cloud_guard.CloudGuardManagedListResource())
	RegisterResource("oci_cloud_guard_responder_recipe", tf_cloud_guard.CloudGuardResponderRecipeResource())
	RegisterResource("oci_cloud_guard_security_recipe", tf_cloud_guard.CloudGuardSecurityRecipeResource())
	RegisterResource("oci_cloud_guard_security_zone", tf_cloud_guard.CloudGuardSecurityZoneResource())
	RegisterResource("oci_cloud_guard_target", tf_cloud_guard.CloudGuardTargetResource())
	// computeinstanceagent service
	// containerengine service
	RegisterResource("oci_containerengine_cluster", tf_containerengine.ContainerengineClusterResource())
	RegisterResource("oci_containerengine_node_pool", tf_containerengine.ContainerengineNodePoolResource())
	// core service
	RegisterResource("oci_core_app_catalog_listing_resource_version_agreement", tf_core.AppCatalogListingResourceVersionAgreementResource())
	RegisterResource("oci_core_app_catalog_subscription", tf_core.CoreAppCatalogSubscriptionResource())
	RegisterResource("oci_core_boot_volume", tf_core.CoreBootVolumeResource())
	RegisterResource("oci_core_boot_volume_backup", tf_core.CoreBootVolumeBackupResource())
	RegisterResource("oci_core_capture_filter", tf_core.CoreCaptureFilterResource())
	RegisterResource("oci_core_cluster_network", tf_core.CoreClusterNetworkResource())
	RegisterResource("oci_core_compute_capacity_reservation", tf_core.CoreComputeCapacityReservationResource())
	RegisterResource("oci_core_compute_image_capability_schema", tf_core.CoreComputeImageCapabilitySchemaResource())
	RegisterResource("oci_core_console_history", tf_core.CoreConsoleHistoryResource())
	RegisterResource("oci_core_cpe", tf_core.CoreCpeResource())
	RegisterResource("oci_core_cross_connect", tf_core.CoreCrossConnectResource())
	RegisterResource("oci_core_cross_connect_group", tf_core.CoreCrossConnectGroupResource())
	RegisterResource("oci_core_dedicated_vm_host", tf_core.CoreDedicatedVmHostResource())
	RegisterResource("oci_core_default_dhcp_options", tf_core.DefaultCoreDhcpOptionsResource())
	RegisterResource("oci_core_default_route_table", tf_core.DefaultCoreRouteTableResource())
	RegisterResource("oci_core_default_security_list", tf_core.CoreDefaultSecurityListResource())
	RegisterResource("oci_core_dhcp_options", tf_core.CoreDhcpOptionsResource())
	RegisterResource("oci_core_drg", tf_core.CoreDrgResource())
	RegisterResource("oci_core_drg_attachment", tf_core.CoreDrgAttachmentResource())
	RegisterResource("oci_core_drg_attachment_management", tf_core.CoreDrgAttachmentManagementResource())
	RegisterResource("oci_core_drg_attachments_list", tf_core.CoreDrgAttachmentsListResource())
	RegisterResource("oci_core_drg_route_distribution", tf_core.CoreDrgRouteDistributionResource())
	RegisterResource("oci_core_drg_route_distribution_statement", tf_core.CoreDrgRouteDistributionStatementResource())
	RegisterResource("oci_core_drg_route_table", tf_core.CoreDrgRouteTableResource())
	RegisterResource("oci_core_drg_route_table_route_rule", tf_core.CoreDrgRouteTableRouteRuleResource())
	RegisterResource("oci_core_image", tf_core.CoreImageResource())
	RegisterResource("oci_core_instance", tf_core.CoreInstanceResource())
	RegisterResource("oci_core_instance_configuration", tf_core.CoreInstanceConfigurationResource())
	RegisterResource("oci_core_instance_console_connection", tf_core.CoreInstanceConsoleConnectionResource())
	RegisterResource("oci_core_instance_pool", tf_core.CoreInstancePoolResource())
	RegisterResource("oci_core_instance_pool_instance", tf_core.CoreInstancePoolInstanceResource())
	RegisterResource("oci_core_internet_gateway", tf_core.CoreInternetGatewayResource())
	RegisterResource("oci_core_ipsec", tf_core.CoreIpSecConnectionResource())
	RegisterResource("oci_core_ipsec_connection_tunnel_management", tf_core.CoreIpSecConnectionTunnelManagementResource())
	RegisterResource("oci_core_ipv6", tf_core.CoreIpv6Resource())
	RegisterResource("oci_core_listing_resource_version_agreement", tf_core.AppCatalogListingResourceVersionAgreementResource())
	RegisterResource("oci_core_local_peering_gateway", tf_core.CoreLocalPeeringGatewayResource())
	RegisterResource("oci_core_nat_gateway", tf_core.CoreNatGatewayResource())
	RegisterResource("oci_core_network_security_group", tf_core.CoreNetworkSecurityGroupResource())
	RegisterResource("oci_core_network_security_group_security_rule", tf_core.CoreNetworkSecurityGroupSecurityRuleResource())
	RegisterResource("oci_core_private_ip", tf_core.CorePrivateIpResource())
	RegisterResource("oci_core_public_ip", tf_core.CorePublicIpResource())
	RegisterResource("oci_core_public_ip_pool", tf_core.CorePublicIpPoolResource())
	RegisterResource("oci_core_public_ip_pool_capacity", tf_core.PublicIpPoolCapacityResource())
	RegisterResource("oci_core_remote_peering_connection", tf_core.CoreRemotePeeringConnectionResource())
	RegisterResource("oci_core_route_table", tf_core.CoreRouteTableResource())
	RegisterResource("oci_core_route_table_attachment", tf_core.CoreRouteTableAttachmentResource())
	RegisterResource("oci_core_security_list", tf_core.CoreSecurityListResource())
	RegisterResource("oci_core_service_gateway", tf_core.CoreServiceGatewayResource())
	RegisterResource("oci_core_shape_management", tf_core.CoreShapeResource())
	RegisterResource("oci_core_subnet", tf_core.CoreSubnetResource())
	RegisterResource("oci_core_vcn", tf_core.CoreVcnResource())
	RegisterResource("oci_core_virtual_circuit", tf_core.CoreVirtualCircuitResource())
	RegisterResource("oci_core_vlan", tf_core.CoreVlanResource())
	RegisterResource("oci_core_vnic_attachment", tf_core.CoreVnicAttachmentResource())
	RegisterResource("oci_core_volume", tf_core.CoreVolumeResource())
	RegisterResource("oci_core_volume_attachment", tf_core.CoreVolumeAttachmentResource())
	RegisterResource("oci_core_volume_backup", tf_core.CoreVolumeBackupResource())
	RegisterResource("oci_core_volume_backup_policy", tf_core.CoreVolumeBackupPolicyResource())
	RegisterResource("oci_core_volume_backup_policy_assignment", tf_core.CoreVolumeBackupPolicyAssignmentResource())
	RegisterResource("oci_core_volume_group", tf_core.CoreVolumeGroupResource())
	RegisterResource("oci_core_volume_group_backup", tf_core.CoreVolumeGroupBackupResource())
	RegisterResource("oci_core_vtap", tf_core.CoreVtapResource())
	// data_connectivity service
	RegisterResource("oci_data_connectivity_registry", tf_data_connectivity.DataConnectivityRegistryResource())
	RegisterResource("oci_data_connectivity_registry_connection", tf_data_connectivity.DataConnectivityRegistryConnectionResource())
	RegisterResource("oci_data_connectivity_registry_data_asset", tf_data_connectivity.DataConnectivityRegistryDataAssetResource())
	RegisterResource("oci_data_connectivity_registry_folder", tf_data_connectivity.DataConnectivityRegistryFolderResource())
	// data_labeling_service service
	RegisterResource("oci_data_labeling_service_dataset", tf_data_labeling_service.DataLabelingServiceDatasetResource())

	// database service
	RegisterResource("oci_database_autonomous_container_database", tf_database.DatabaseAutonomousContainerDatabaseResource())
	RegisterResource("oci_database_autonomous_container_database_dataguard_association", tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationResource())
	RegisterResource("oci_database_autonomous_container_database_dataguard_association_operation", tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource())
	RegisterResource("oci_database_autonomous_database", tf_database.DatabaseAutonomousDatabaseResource())
	RegisterResource("oci_database_autonomous_database_backup", tf_database.DatabaseAutonomousDatabaseBackupResource())
	RegisterResource("oci_database_autonomous_database_instance_wallet_management", tf_database.DatabaseAutonomousDatabaseInstanceWalletManagementResource())
	RegisterResource("oci_database_autonomous_database_regional_wallet_management", tf_database.DatabaseAutonomousDatabaseRegionalWalletManagementResource())
	RegisterResource("oci_database_autonomous_database_wallet", tf_database.DatabaseAutonomousDatabaseWalletResource())
	RegisterResource("oci_database_autonomous_exadata_infrastructure", tf_database.DatabaseAutonomousExadataInfrastructureResource())
	RegisterResource("oci_database_autonomous_vm_cluster", tf_database.DatabaseAutonomousVmClusterResource())
	RegisterResource("oci_database_backup", tf_database.DatabaseBackupResource())
	RegisterResource("oci_database_backup_destination", tf_database.DatabaseBackupDestinationResource())
	RegisterResource("oci_database_cloud_autonomous_vm_cluster", tf_database.DatabaseCloudAutonomousVmClusterResource())
	RegisterResource("oci_database_cloud_database_management", tf_database.DatabaseCloudDatabaseManagementResource())
	RegisterResource("oci_database_cloud_exadata_infrastructure", tf_database.DatabaseCloudExadataInfrastructureResource())
	RegisterResource("oci_database_cloud_vm_cluster", tf_database.DatabaseCloudVmClusterResource())
	RegisterResource("oci_database_cloud_vm_cluster_iorm_config", tf_database.DatabaseCloudVmClusterIormConfigResource())
	RegisterResource("oci_database_data_guard_association", tf_database.DatabaseDataGuardAssociationResource())
	RegisterResource("oci_database_database", tf_database.DatabaseDatabaseResource())
	RegisterResource("oci_database_database_software_image", tf_database.DatabaseDatabaseSoftwareImageResource())
	RegisterResource("oci_database_database_upgrade", tf_database.DatabaseDatabaseUpgradeResource())
	RegisterResource("oci_database_db_home", tf_database.DatabaseDbHomeResource())
	RegisterResource("oci_database_db_node_console_connection", tf_database.DatabaseDbNodeConsoleConnectionResource())
	RegisterResource("oci_database_db_system", tf_database.DatabaseDbSystemResource())
	RegisterResource("oci_database_db_systems_upgrade", tf_database.DatabaseDbSystemsUpgradeResource())
	RegisterResource("oci_database_exadata_infrastructure", tf_database.DatabaseExadataInfrastructureResource())
	RegisterResource("oci_database_exadata_infrastructure", tf_database.DatabaseExadataInfrastructureResource())
	RegisterResource("oci_database_exadata_infrastructure_storage", tf_database.DatabaseExadataInfrastructureStorageResource())
	RegisterResource("oci_database_exadata_iorm_config", tf_database.DatabaseExadataIormConfigResource())
	RegisterResource("oci_database_external_container_database", tf_database.DatabaseExternalContainerDatabaseResource())
	RegisterResource("oci_database_external_container_database_management", tf_database.DatabaseExternalContainerDatabaseManagementResource())
	RegisterResource("oci_database_external_database_connector", tf_database.DatabaseExternalDatabaseConnectorResource())
	RegisterResource("oci_database_external_non_container_database", tf_database.DatabaseExternalNonContainerDatabaseResource())
	RegisterResource("oci_database_external_non_container_database_management", tf_database.DatabaseExternalNonContainerDatabaseManagementResource())
	RegisterResource("oci_database_external_non_container_database_operations_insights_management", tf_database.DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource())
	RegisterResource("oci_database_external_pluggable_database", tf_database.DatabaseExternalPluggableDatabaseResource())
	RegisterResource("oci_database_external_pluggable_database_management", tf_database.DatabaseExternalPluggableDatabaseManagementResource())
	RegisterResource("oci_database_external_pluggable_database_operations_insights_management", tf_database.DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource())
	RegisterResource("oci_database_externalcontainerdatabases_stack_monitoring", tf_database.DatabaseExternalcontainerdatabasesStackMonitoringResource())
	RegisterResource("oci_database_externalnoncontainerdatabases_stack_monitoring", tf_database.DatabaseExternalnoncontainerdatabasesStackMonitoringResource())
	RegisterResource("oci_database_externalpluggabledatabases_stack_monitoring", tf_database.DatabaseExternalpluggabledatabasesStackMonitoringResource())
	RegisterResource("oci_database_key_store", tf_database.DatabaseKeyStoreResource())
	RegisterResource("oci_database_maintenance_run", tf_database.DatabaseMaintenanceRunResource())
	RegisterResource("oci_database_pluggable_database", tf_database.DatabasePluggableDatabaseResource())
	RegisterResource("oci_database_pluggable_databases_local_clone", tf_database.DatabasePluggableDatabasesLocalCloneResource())
	RegisterResource("oci_database_pluggable_databases_remote_clone", tf_database.DatabasePluggableDatabasesRemoteCloneResource())
	RegisterResource("oci_database_vm_cluster", tf_database.DatabaseVmClusterResource())
	RegisterResource("oci_database_vm_cluster_add_virtual_machine", tf_database.DatabaseVmClusterAddVirtualMachineResource())
	RegisterResource("oci_database_vm_cluster_network", tf_database.DatabaseVmClusterNetworkResource())
	RegisterResource("oci_database_vm_cluster_remove_virtual_machine", tf_database.DatabaseVmClusterRemoveVirtualMachineResource())
	// database_management service
	RegisterResource("oci_database_management_db_management_private_endpoint", tf_database_management.DatabaseManagementDbManagementPrivateEndpointResource())
	RegisterResource("oci_database_management_managed_database_group", tf_database_management.DatabaseManagementManagedDatabaseGroupResource())
	RegisterResource("oci_database_management_managed_databases_change_database_parameter", tf_database_management.DatabaseManagementManagedDatabasesChangeDatabaseParameterResource())
	RegisterResource("oci_database_management_managed_databases_reset_database_parameter", tf_database_management.DatabaseManagementManagedDatabasesResetDatabaseParameterResource())
	// database_migration service
	RegisterResource("oci_database_migration", tf_database_migration.DatabaseMigrationResource())
	RegisterResource("oci_database_migration_agent", tf_database_migration.DatabaseMigrationAgentResource())
	RegisterResource("oci_database_migration_connection", tf_database_migration.DatabaseMigrationConnectionResource())
	RegisterResource("oci_database_migration_job", tf_database_migration.DatabaseMigrationJobResource())
	RegisterResource("oci_database_migration_migration", tf_database_migration.DatabaseMigrationMigrationResource())
	// database_tools service
	RegisterResource("oci_database_tools_database_tools_connection", tf_database_tools.DatabaseToolsDatabaseToolsConnectionResource())
	RegisterResource("oci_database_tools_database_tools_private_endpoint", tf_database_tools.DatabaseToolsDatabaseToolsPrivateEndpointResource())
	// datacatalog service
	RegisterResource("oci_datacatalog_catalog", tf_datacatalog.DatacatalogCatalogResource())
	RegisterResource("oci_datacatalog_catalog_private_endpoint", tf_datacatalog.DatacatalogCatalogPrivateEndpointResource())
	RegisterResource("oci_datacatalog_connection", tf_datacatalog.DatacatalogConnectionResource())
	RegisterResource("oci_datacatalog_data_asset", tf_datacatalog.DatacatalogDataAssetResource())
	RegisterResource("oci_datacatalog_metastore", tf_datacatalog.DatacatalogMetastoreResource())
	// dataflow service
	RegisterResource("oci_dataflow_application", tf_dataflow.DataflowApplicationResource())
	RegisterResource("oci_dataflow_invoke_run", tf_dataflow.DataflowInvokeRunResource())
	RegisterResource("oci_dataflow_private_endpoint", tf_dataflow.DataflowPrivateEndpointResource())
	// dataintegration service
	RegisterResource("oci_dataintegration_workspace", tf_dataintegration.DataintegrationWorkspaceResource())
	// datascience service
	RegisterResource("oci_datascience_job", tf_datascience.DatascienceJobResource())
	RegisterResource("oci_datascience_job_run", tf_datascience.DatascienceJobRunResource())
	RegisterResource("oci_datascience_model", tf_datascience.DatascienceModelResource())
	RegisterResource("oci_datascience_model_deployment", tf_datascience.DatascienceModelDeploymentResource())
	RegisterResource("oci_datascience_model_provenance", tf_datascience.DatascienceModelProvenanceResource())
	RegisterResource("oci_datascience_notebook_session", tf_datascience.DatascienceNotebookSessionResource())
	RegisterResource("oci_datascience_project", tf_datascience.DatascienceProjectResource())
	// devops service
	RegisterResource("oci_devops_build_pipeline", tf_devops.DevopsBuildPipelineResource())
	RegisterResource("oci_devops_build_pipeline_stage", tf_devops.DevopsBuildPipelineStageResource())
	RegisterResource("oci_devops_build_run", tf_devops.DevopsBuildRunResource())
	RegisterResource("oci_devops_connection", tf_devops.DevopsConnectionResource())
	RegisterResource("oci_devops_deploy_artifact", tf_devops.DevopsDeployArtifactResource())
	RegisterResource("oci_devops_deploy_environment", tf_devops.DevopsDeployEnvironmentResource())
	RegisterResource("oci_devops_deploy_pipeline", tf_devops.DevopsDeployPipelineResource())
	RegisterResource("oci_devops_deploy_stage", tf_devops.DevopsDeployStageResource())
	RegisterResource("oci_devops_deployment", tf_devops.DevopsDeploymentResource())
	RegisterResource("oci_devops_project", tf_devops.DevopsProjectResource())
	RegisterResource("oci_devops_repository", tf_devops.DevopsRepositoryResource())
	RegisterResource("oci_devops_repository_mirror", tf_devops.DevopsRepositoryMirrorResource())
	RegisterResource("oci_devops_repository_ref", tf_devops.DevopsRepositoryRefResource())
	RegisterResource("oci_devops_trigger", tf_devops.DevopsTriggerResource())
	// dns service
	RegisterResource("oci_dns_record", tf_dns.DnsRecordResource())
	RegisterResource("oci_dns_resolver", tf_dns.DnsResolverResource())
	RegisterResource("oci_dns_resolver_endpoint", tf_dns.DnsResolverEndpointResource())
	RegisterResource("oci_dns_rrset", tf_dns.DnsRrsetResource())
	RegisterResource("oci_dns_steering_policy", tf_dns.DnsSteeringPolicyResource())
	RegisterResource("oci_dns_steering_policy_attachment", tf_dns.DnsSteeringPolicyAttachmentResource())
	RegisterResource("oci_dns_tsig_key", tf_dns.DnsTsigKeyResource())
	RegisterResource("oci_dns_view", tf_dns.DnsViewResource())
	RegisterResource("oci_dns_zone", tf_dns.DnsZoneResource())
	// em_warehouse service
	RegisterResource("oci_em_warehouse_em_warehouse", tf_em_warehouse.EmWarehouseEmWarehouseResource())
	// email service
	RegisterResource("oci_email_dkim", tf_email.EmailDkimResource())
	RegisterResource("oci_email_email_domain", tf_email.EmailEmailDomainResource())
	RegisterResource("oci_email_sender", tf_email.EmailSenderResource())
	RegisterResource("oci_email_suppression", tf_email.EmailSuppressionResource())
	// events service
	RegisterResource("oci_events_rule", tf_events.EventsRuleResource())
	// file_storage service
	RegisterResource("oci_file_storage_export", tf_file_storage.FileStorageExportResource())
	RegisterResource("oci_file_storage_export_set", tf_file_storage.FileStorageExportSetResource())
	RegisterResource("oci_file_storage_file_system", tf_file_storage.FileStorageFileSystemResource())
	RegisterResource("oci_file_storage_mount_target", tf_file_storage.FileStorageMountTargetResource())
	RegisterResource("oci_file_storage_snapshot", tf_file_storage.FileStorageSnapshotResource())
	// functions service
	RegisterResource("oci_functions_application", tf_functions.FunctionsApplicationResource())
	RegisterResource("oci_functions_function", tf_functions.FunctionsFunctionResource())
	RegisterResource("oci_functions_invoke_function", tf_functions.FunctionsInvokeFunctionResource())
	// generic_artifacts_content service
	RegisterResource("oci_generic_artifacts_content_artifact_by_path", tf_generic_artifacts_content.GenericArtifactsContentArtifactByPathResource())
	// golden_gate service
	RegisterResource("oci_golden_gate_database_registration", tf_golden_gate.GoldenGateDatabaseRegistrationResource())
	RegisterResource("oci_golden_gate_deployment", tf_golden_gate.GoldenGateDeploymentResource())
	RegisterResource("oci_golden_gate_deployment_backup", tf_golden_gate.GoldenGateDeploymentBackupResource())
	// health_checks service
	RegisterResource("oci_health_checks_http_monitor", tf_health_checks.HealthChecksHttpMonitorResource())
	RegisterResource("oci_health_checks_http_probe", tf_health_checks.HealthChecksHttpProbeResource())
	RegisterResource("oci_health_checks_ping_monitor", tf_health_checks.HealthChecksPingMonitorResource())
	RegisterResource("oci_health_checks_ping_probe", tf_health_checks.HealthChecksPingProbeResource())
	// identity service
	RegisterResource("oci_identity_api_key", tf_identity.IdentityApiKeyResource())
	RegisterResource("oci_identity_api_key", tf_identity.IdentityApiKeyResource())
	RegisterResource("oci_identity_auth_token", tf_identity.IdentityAuthTokenResource())
	RegisterResource("oci_identity_auth_token", tf_identity.IdentityAuthTokenResource())
	RegisterResource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyResource())
	RegisterResource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyResource())
	RegisterResource("oci_identity_compartment", tf_identity.IdentityCompartmentResource())
	RegisterResource("oci_identity_compartment", tf_identity.IdentityCompartmentResource())
	RegisterResource("oci_identity_customer_secret_key", tf_identity.IdentityCustomerSecretKeyResource())
	RegisterResource("oci_identity_customer_secret_key", tf_identity.IdentityCustomerSecretKeyResource())
	RegisterResource("oci_identity_db_credential", tf_identity.IdentityDbCredentialResource())
	RegisterResource("oci_identity_domain", tf_identity.IdentityDomainResource())
	RegisterResource("oci_identity_domain_replication_to_region", tf_identity.IdentityDomainReplicationToRegionResource())
	RegisterResource("oci_identity_dynamic_group", tf_identity.IdentityDynamicGroupResource())
	RegisterResource("oci_identity_group", tf_identity.IdentityGroupResource())
	RegisterResource("oci_identity_identity_provider", tf_identity.IdentityIdentityProviderResource())
	RegisterResource("oci_identity_idp_group_mapping", tf_identity.IdentityIdpGroupMappingResource())
	RegisterResource("oci_identity_import_standard_tags_management", tf_identity.IdentityImportStandardTagsManagementResource())
	RegisterResource("oci_identity_network_source", tf_identity.IdentityNetworkSourceResource())
	RegisterResource("oci_identity_policy", tf_identity.IdentityPolicyResource())
	RegisterResource("oci_identity_smtp_credential", tf_identity.IdentitySmtpCredentialResource())
	RegisterResource("oci_identity_swift_password", tf_identity.IdentitySwiftPasswordResource())
	RegisterResource("oci_identity_tag", tf_identity.IdentityTagResource())
	RegisterResource("oci_identity_tag_default", tf_identity.IdentityTagDefaultResource())
	RegisterResource("oci_identity_tag_namespace", tf_identity.IdentityTagNamespaceResource())
	RegisterResource("oci_identity_ui_password", tf_identity.IdentityUiPasswordResource())
	RegisterResource("oci_identity_user", tf_identity.IdentityUserResource())
	RegisterResource("oci_identity_user_capabilities_management", tf_identity.IdentityUserCapabilitiesManagementResource())
	RegisterResource("oci_identity_user_group_membership", tf_identity.IdentityUserGroupMembershipResource())
	// identity_data_plane service
	RegisterResource("oci_identity_data_plane_generate_scoped_access_token", tf_identity.IdentityDataPlaneGenerateScopedAccessTokenResource())
	RegisterResource("oci_identity_data_plane_generate_scoped_access_token", tf_identity_data_plane.IdentityDataPlaneGenerateScopedAccessTokenResource())
	// integration service
	RegisterResource("oci_integration_integration_instance", tf_integration.IntegrationIntegrationInstanceResource())
	// jms service
	RegisterResource("oci_jms_fleet", tf_jms.JmsFleetResource())
	// kms service
	RegisterResource("oci_kms_encrypted_data", tf_kms.KmsEncryptedDataResource())
	RegisterResource("oci_kms_generated_key", tf_kms.KmsGeneratedKeyResource())
	RegisterResource("oci_kms_key", tf_kms.KmsKeyResource())
	RegisterResource("oci_kms_key_version", tf_kms.KmsKeyVersionResource())
	RegisterResource("oci_kms_sign", tf_kms.KmsSignResource())
	RegisterResource("oci_kms_vault", tf_kms.KmsVaultResource())
	RegisterResource("oci_kms_vault_replication", tf_kms.KmsVaultReplicationResource())
	RegisterResource("oci_kms_verify", tf_kms.KmsVerifyResource())
	// license_manager service
	RegisterResource("oci_license_manager_configuration", tf_license_manager.LicenseManagerConfigurationResource())
	RegisterResource("oci_license_manager_license_record", tf_license_manager.LicenseManagerLicenseRecordResource())
	RegisterResource("oci_license_manager_product_license", tf_license_manager.LicenseManagerProductLicenseResource())
	// limits service
	RegisterResource("oci_limits_quota", tf_limits.LimitsQuotaResource())
	// load_balancer service
	RegisterResource("oci_load_balancer_backend", tf_load_balancer.LoadBalancerBackendResource())
	RegisterResource("oci_load_balancer_backend_set", tf_load_balancer.LoadBalancerBackendSetResource())
	RegisterResource("oci_load_balancer_certificate", tf_load_balancer.LoadBalancerCertificateResource())
	RegisterResource("oci_load_balancer_hostname", tf_load_balancer.LoadBalancerHostnameResource())
	RegisterResource("oci_load_balancer_listener", tf_load_balancer.LoadBalancerListenerResource())
	RegisterResource("oci_load_balancer_load_balancer", tf_load_balancer.LoadBalancerLoadBalancerResource())
	RegisterResource("oci_load_balancer_load_balancer_routing_policy", tf_load_balancer.LoadBalancerLoadBalancerRoutingPolicyResource())
	RegisterResource("oci_load_balancer_path_route_set", tf_load_balancer.LoadBalancerPathRouteSetResource())
	RegisterResource("oci_load_balancer_rule_set", tf_load_balancer.LoadBalancerRuleSetResource())
	RegisterResource("oci_load_balancer_ssl_cipher_suite", tf_load_balancer.LoadBalancerSslCipherSuiteResource())
	// log_analytics service
	RegisterResource("oci_log_analytics_log_analytics_entity", tf_log_analytics.LogAnalyticsLogAnalyticsEntityResource())
	RegisterResource("oci_log_analytics_log_analytics_import_custom_content", tf_log_analytics.LogAnalyticsLogAnalyticsImportCustomContentResource())
	RegisterResource("oci_log_analytics_log_analytics_log_group", tf_log_analytics.LogAnalyticsLogAnalyticsLogGroupResource())
	RegisterResource("oci_log_analytics_log_analytics_object_collection_rule", tf_log_analytics.LogAnalyticsLogAnalyticsObjectCollectionRuleResource())
	RegisterResource("oci_log_analytics_log_analytics_preferences_management", tf_log_analytics.LogAnalyticsLogAnalyticsPreferencesManagementResource())
	RegisterResource("oci_log_analytics_log_analytics_resource_categories_management", tf_log_analytics.LogAnalyticsLogAnalyticsResourceCategoriesManagementResource())
	RegisterResource("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", tf_log_analytics.LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource())
	RegisterResource("oci_log_analytics_namespace", tf_log_analytics.LogAnalyticsNamespaceResource())
	RegisterResource("oci_log_analytics_namespace_scheduled_task", tf_log_analytics.LogAnalyticsNamespaceScheduledTaskResource())
	// logging service
	RegisterResource("oci_logging_log", tf_logging.LoggingLogResource())
	RegisterResource("oci_logging_log_group", tf_logging.LoggingLogGroupResource())
	RegisterResource("oci_logging_log_saved_search", tf_logging.LoggingLogSavedSearchResource())
	RegisterResource("oci_logging_unified_agent_configuration", tf_logging.LoggingUnifiedAgentConfigurationResource())
	// management_agent service
	RegisterResource("oci_management_agent_management_agent", tf_management_agent.ManagementAgentManagementAgentResource())
	RegisterResource("oci_management_agent_management_agent_install_key", tf_management_agent.ManagementAgentManagementAgentInstallKeyResource())
	// management_dashboard service
	RegisterResource("oci_management_dashboard_management_dashboards_import", tf_management_dashboard.ManagementDashboardManagementDashboardsImportResource())
	// marketplace service
	RegisterResource("oci_marketplace_accepted_agreement", tf_marketplace.MarketplaceAcceptedAgreementResource())
	RegisterResource("oci_marketplace_listing_package_agreement", tf_marketplace.MarketplaceListingPackageAgreementResource())
	RegisterResource("oci_marketplace_publication", tf_marketplace.MarketplacePublicationResource())
	// metering_computation service
	RegisterResource("oci_metering_computation_custom_table", tf_metering_computation.MeteringComputationCustomTableResource())
	RegisterResource("oci_metering_computation_query", tf_metering_computation.MeteringComputationQueryResource())
	RegisterResource("oci_metering_computation_schedule", tf_metering_computation.MeteringComputationScheduleResource())
	RegisterResource("oci_metering_computation_usage", tf_metering_computation.MeteringComputationUsageResource())
	// monitoring service
	RegisterResource("oci_monitoring_alarm", tf_monitoring.MonitoringAlarmResource())
	// mysql service
	RegisterResource("oci_mysql_analytics_cluster", tf_mysql.MysqlAnalyticsClusterResource())
	RegisterResource("oci_mysql_channel", tf_mysql.MysqlChannelResource())
	RegisterResource("oci_mysql_heat_wave_cluster", tf_mysql.MysqlHeatWaveClusterResource())
	RegisterResource("oci_mysql_mysql_backup", tf_mysql.MysqlMysqlBackupResource())
	RegisterResource("oci_mysql_mysql_db_system", tf_mysql.MysqlMysqlDbSystemResource())
	// network_load_balancer service
	RegisterResource("oci_network_load_balancer_backend", tf_network_load_balancer.NetworkLoadBalancerBackendResource())
	RegisterResource("oci_network_load_balancer_backend_set", tf_network_load_balancer.NetworkLoadBalancerBackendSetResource())
	RegisterResource("oci_network_load_balancer_listener", tf_network_load_balancer.NetworkLoadBalancerListenerResource())
	RegisterResource("oci_network_load_balancer_network_load_balancer", tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancerResource())
	RegisterResource("oci_network_load_balancer_network_load_balancers_backend_sets_unified", tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource())
	// nosql service
	RegisterResource("oci_nosql_index", tf_nosql.NosqlIndexResource())
	RegisterResource("oci_nosql_table", tf_nosql.NosqlTableResource())
	// objectstorage service
	RegisterResource("oci_objectstorage_bucket", tf_objectstorage.ObjectStorageBucketResource())
	RegisterResource("oci_objectstorage_namespace_metadata", tf_objectstorage.ObjectStorageNamespaceMetadataResource())
	RegisterResource("oci_objectstorage_object", tf_objectstorage.ObjectStorageObjectResource())
	RegisterResource("oci_objectstorage_object_lifecycle_policy", tf_objectstorage.ObjectStorageObjectLifecyclePolicyResource())
	RegisterResource("oci_objectstorage_preauthrequest", tf_objectstorage.ObjectStoragePreauthenticatedRequestResource())
	RegisterResource("oci_objectstorage_replication_policy", tf_objectstorage.ObjectStorageReplicationPolicyResource())
	// oce service
	RegisterResource("oci_oce_oce_instance", tf_oce.OceOceInstanceResource())
	// ocvp service
	RegisterResource("oci_ocvp_esxi_host", tf_ocvp.OcvpEsxiHostResource())
	RegisterResource("oci_ocvp_sddc", tf_ocvp.OcvpSddcResource())
	// oda service
	RegisterResource("oci_oda_oda_instance", tf_oda.OdaOdaInstanceResource())
	// onesubscription service
	// ons service
	RegisterResource("oci_ons_notification_topic", tf_ons.OnsNotificationTopicResource())
	RegisterResource("oci_ons_subscription", tf_ons.OnsSubscriptionResource())
	// operator_access_control service
	RegisterResource("oci_operator_access_control_operator_control", tf_operator_access_control.OperatorAccessControlOperatorControlResource())
	RegisterResource("oci_operator_access_control_operator_control_assignment", tf_operator_access_control.OperatorAccessControlOperatorControlAssignmentResource())
	// opsi service
	RegisterResource("oci_opsi_awr_hub", tf_opsi.OpsiAwrHubResource())
	RegisterResource("oci_opsi_database_insight", tf_opsi.OpsiDatabaseInsightResource())
	RegisterResource("oci_opsi_enterprise_manager_bridge", tf_opsi.OpsiEnterpriseManagerBridgeResource())
	RegisterResource("oci_opsi_exadata_insight", tf_opsi.OpsiExadataInsightResource())
	RegisterResource("oci_opsi_host_insight", tf_opsi.OpsiHostInsightResource())
	RegisterResource("oci_opsi_operations_insights_private_endpoint", tf_opsi.OpsiOperationsInsightsPrivateEndpointResource())
	RegisterResource("oci_opsi_operations_insights_warehouse", tf_opsi.OpsiOperationsInsightsWarehouseResource())
	RegisterResource("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", tf_opsi.OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource())
	RegisterResource("oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet", tf_opsi.OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource())
	RegisterResource("oci_opsi_operations_insights_warehouse_user", tf_opsi.OpsiOperationsInsightsWarehouseUserResource())
	// optimizer service
	RegisterResource("oci_optimizer_enrollment_status", tf_optimizer.OptimizerEnrollmentStatusResource())
	RegisterResource("oci_optimizer_profile", tf_optimizer.OptimizerProfileResource())
	RegisterResource("oci_optimizer_recommendation", tf_optimizer.OptimizerRecommendationResource())
	RegisterResource("oci_optimizer_resource_action", tf_optimizer.OptimizerResourceActionResource())
	// osmanagement service
	RegisterResource("oci_osmanagement_managed_instance", tf_osmanagement.OsmanagementManagedInstanceResource())
	RegisterResource("oci_osmanagement_managed_instance_group", tf_osmanagement.OsmanagementManagedInstanceGroupResource())
	RegisterResource("oci_osmanagement_managed_instance_management", tf_osmanagement.OsmanagementManagedInstanceManagementResource())
	RegisterResource("oci_osmanagement_software_source", tf_osmanagement.OsmanagementSoftwareSourceResource())
	// osp_gateway
	RegisterResource("oci_osp_gateway_subscription", tf_osp_gateway.OspGatewaySubscriptionResource())
	// osub_usage service
	// osub_subscription service
	// osub_organization_subscription service
	// osub_billing_schedule service
	// resourcemanager service
	RegisterResource("oci_resourcemanager_private_endpoint", tf_resourcemanager.ResourcemanagerPrivateEndpointResource())
	// sch service
	RegisterResource("oci_sch_service_connector", tf_sch.SchServiceConnectorResource())
	// secrets service
	// service_catalog service
	RegisterResource("oci_service_catalog_private_application", tf_service_catalog.ServiceCatalogPrivateApplicationResource())
	RegisterResource("oci_service_catalog_service_catalog", tf_service_catalog.ServiceCatalogServiceCatalogResource())
	RegisterResource("oci_service_catalog_service_catalog_association", tf_service_catalog.ServiceCatalogServiceCatalogAssociationResource())
	// service_manager_proxy service
	// service_mesh service
	RegisterResource("oci_service_mesh_access_policy", tf_service_mesh.ServiceMeshAccessPolicyResource())
	RegisterResource("oci_service_mesh_ingress_gateway", tf_service_mesh.ServiceMeshIngressGatewayResource())
	RegisterResource("oci_service_mesh_ingress_gateway_route_table", tf_service_mesh.ServiceMeshIngressGatewayRouteTableResource())
	RegisterResource("oci_service_mesh_mesh", tf_service_mesh.ServiceMeshMeshResource())
	RegisterResource("oci_service_mesh_virtual_deployment", tf_service_mesh.ServiceMeshVirtualDeploymentResource())
	RegisterResource("oci_service_mesh_virtual_service", tf_service_mesh.ServiceMeshVirtualServiceResource())
	RegisterResource("oci_service_mesh_virtual_service_route_table", tf_service_mesh.ServiceMeshVirtualServiceRouteTableResource())
	// stack_monitoring service
	RegisterResource("oci_stack_monitoring_discovery_job", tf_stack_monitoring.StackMonitoringDiscoveryJobResource())
	RegisterResource("oci_stack_monitoring_monitored_resource", tf_stack_monitoring.StackMonitoringMonitoredResourceResource())
	RegisterResource("oci_stack_monitoring_monitored_resources_associate_monitored_resource", tf_stack_monitoring.StackMonitoringMonitoredResourcesAssociateMonitoredResourceResource())
	RegisterResource("oci_stack_monitoring_monitored_resources_list_member", tf_stack_monitoring.StackMonitoringMonitoredResourcesListMemberResource())
	RegisterResource("oci_stack_monitoring_monitored_resources_search", tf_stack_monitoring.StackMonitoringMonitoredResourcesSearchResource())
	RegisterResource("oci_stack_monitoring_monitored_resources_search_association", tf_stack_monitoring.StackMonitoringMonitoredResourcesSearchAssociationResource())
	// streaming service
	RegisterResource("oci_streaming_connect_harness", tf_streaming.StreamingConnectHarnessResource())
	RegisterResource("oci_streaming_stream", tf_streaming.StreamingStreamResource())
	RegisterResource("oci_streaming_stream_pool", tf_streaming.StreamingStreamPoolResource())
	// usage_proxy service
	RegisterResource("oci_usage_proxy_subscription_redeemable_user", tf_usage_proxy.UsageProxySubscriptionRedeemableUserResource())
	// vault service
	RegisterResource("oci_vault_secret", tf_vault.VaultSecretResource())
	// vision service
	RegisterResource("oci_ai_vision_project", tf_ai_vision.AiVisionProjectResource())
	RegisterResource("oci_ai_vision_model", tf_ai_vision.AiVisionModelResource())
	// visual_builder service
	RegisterResource("oci_visual_builder_vb_instance", tf_visual_builder.VisualBuilderVbInstanceResource())
	// vn_monitoring service
	RegisterResource("oci_vn_monitoring_path_analysi", tf_vn_monitoring.VnMonitoringPathAnalysiResource())
	RegisterResource("oci_vn_monitoring_path_analyzer_test", tf_vn_monitoring.VnMonitoringPathAnalyzerTestResource())
	// vulnerability_scanning service
	RegisterResource("oci_vulnerability_scanning_container_scan_recipe", tf_vulnerability_scanning.VulnerabilityScanningContainerScanRecipeResource())
	RegisterResource("oci_vulnerability_scanning_container_scan_target", tf_vulnerability_scanning.VulnerabilityScanningContainerScanTargetResource())
	RegisterResource("oci_vulnerability_scanning_host_scan_recipe", tf_vulnerability_scanning.VulnerabilityScanningHostScanRecipeResource())
	RegisterResource("oci_vulnerability_scanning_host_scan_target", tf_vulnerability_scanning.VulnerabilityScanningHostScanTargetResource())
	// waa service
	RegisterResource("oci_waa_web_app_acceleration", tf_waa.WaaWebAppAccelerationResource())
	RegisterResource("oci_waa_web_app_acceleration_policy", tf_waa.WaaWebAppAccelerationPolicyResource())
	// waas service
	RegisterResource("oci_waas_address_list", tf_waas.WaasAddressListResource())
	RegisterResource("oci_waas_certificate", tf_waas.WaasCertificateResource())
	RegisterResource("oci_waas_custom_protection_rule", tf_waas.WaasCustomProtectionRuleResource())
	RegisterResource("oci_waas_http_redirect", tf_waas.WaasHttpRedirectResource())
	RegisterResource("oci_waas_protection_rule", tf_waas.WaasProtectionRuleResource())
	RegisterResource("oci_waas_purge_cache", tf_waas.WaasPurgeCacheResource())
	RegisterResource("oci_waas_waas_policy", tf_waas.WaasWaasPolicyResource())
	// waf service
	RegisterResource("oci_waf_network_address_list", tf_waf.WafNetworkAddressListResource())
	RegisterResource("oci_waf_web_app_firewall", tf_waf.WafWebAppFirewallResource())
	RegisterResource("oci_waf_web_app_firewall_policy", tf_waf.WafWebAppFirewallPolicyResource())
}
