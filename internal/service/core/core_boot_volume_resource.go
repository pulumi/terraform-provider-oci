// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreBootVolumeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreBootVolume,
		Read:     readCoreBootVolume,
		Update:   updateCoreBootVolume,
		Delete:   deleteCoreBootVolume,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"bootVolume",
								"bootVolumeBackup",
								"bootVolumeReplica",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"backup_policy_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherResource("backup_policy_id", "oci_core_volume_backup_policy_assignment"),
			},
			"boot_volume_replicas": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"availability_domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"boot_volume_replica_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"boot_volume_replicas_deletion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_auto_tune_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"size_in_gbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"vpus_per_gb": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"auto_tuned_vpus_per_gb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hydrated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"size_in_mbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func updateCoreBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreBootVolumeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.BootVolume
	DisableNotFoundRetries bool
}

func (s *CoreBootVolumeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreBootVolumeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateProvisioning),
		string(oci_core.BootVolumeLifecycleStateRestoring),
	}
}

func (s *CoreBootVolumeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
	}
}

func (s *CoreBootVolumeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateTerminating),
	}
}

func (s *CoreBootVolumeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateTerminated),
	}
}

func (s *CoreBootVolumeResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateProvisioning),
	}
}

func (s *CoreBootVolumeResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
	}
}

func (s *CoreBootVolumeResourceCrud) Create() error {
	request := oci_core.CreateBootVolumeRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists("backup_policy_id"); ok {
		tmp := backupPolicyId.(string)
		request.BackupPolicyId = &tmp
	}

	if bootVolumeReplicas, ok := s.D.GetOkExists("boot_volume_replicas"); ok {
		interfaces := bootVolumeReplicas.([]interface{})
		tmp := make([]oci_core.BootVolumeReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "boot_volume_replicas", stateDataIndex)
			converted, err := s.mapToBootVolumeReplicaDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("boot_volume_replicas") {
			request.BootVolumeReplicas = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoTuneEnabled, ok := s.D.GetOkExists("is_auto_tune_enabled"); ok {
		tmp := isAutoTuneEnabled.(bool)
		request.IsAutoTuneEnabled = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if sizeInGBs, ok := s.D.GetOkExists("size_in_gbs"); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInGBs = &tmpInt64
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToBootVolumeSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = tmp
		}
	}

	if vpusPerGB, ok := s.D.GetOkExists("vpus_per_gb"); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VpusPerGB = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *CoreBootVolumeResourceCrud) Get() error {
	request := oci_core.GetBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *CoreBootVolumeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

	if bootVolumeReplicas, ok := s.D.GetOkExists("boot_volume_replicas"); ok {
		interfaces := bootVolumeReplicas.([]interface{})
		tmp := make([]oci_core.BootVolumeReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "boot_volume_replicas", stateDataIndex)
			converted, err := s.mapToBootVolumeReplicaDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("boot_volume_replicas") {
			request.BootVolumeReplicas = tmp
		}
	}

	if bootVolumeReplicasDeletion, ok := s.D.GetOkExists("boot_volume_replicas_deletion"); ok {
		tmp := bootVolumeReplicasDeletion.(bool)
		if tmp == true {
			request.BootVolumeReplicas = []oci_core.BootVolumeReplicaDetails{}
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoTuneEnabled, ok := s.D.GetOkExists("is_auto_tune_enabled"); ok {
		tmp := isAutoTuneEnabled.(bool)
		request.IsAutoTuneEnabled = &tmp
	}

	if s.D.HasChange("kms_key_id") {
		keyUpdateRequest := oci_core.UpdateBootVolumeKmsKeyRequest{}

		bootVolumeId := s.D.Id()
		keyUpdateRequest.BootVolumeId = &bootVolumeId

		tmp := s.D.Get("kms_key_id").(string)
		keyUpdateRequest.KmsKeyId = &tmp

		keyUpdateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		_, err := s.Client.UpdateBootVolumeKmsKey(context.Background(), keyUpdateRequest)
		if err != nil {
			return err
		}
	}

	if sizeInGBs, ok := s.D.GetOkExists("size_in_gbs"); ok && s.D.HasChange("size_in_gbs") {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInGBs = &tmpInt64
	}

	if vpusPerGB, ok := s.D.GetOkExists("vpus_per_gb"); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VpusPerGB = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *CoreBootVolumeResourceCrud) Delete() error {
	request := oci_core.DeleteBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteBootVolume(context.Background(), request)
	return err
}

func (s *CoreBootVolumeResourceCrud) SetData() error {
	if s.Res.AutoTunedVpusPerGB != nil {
		s.D.Set("auto_tuned_vpus_per_gb", strconv.FormatInt(*s.Res.AutoTunedVpusPerGB, 10))
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	bootVolumeReplicas := []interface{}{}
	for _, item := range s.Res.BootVolumeReplicas {
		bootVolumeReplicas = append(bootVolumeReplicas, BootVolumeReplicaInfoToMap(item))
	}
	s.D.Set("boot_volume_replicas", bootVolumeReplicas)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.IsAutoTuneEnabled != nil {
		s.D.Set("is_auto_tune_enabled", *s.Res.IsAutoTuneEnabled)
	}

	if s.Res.IsHydrated != nil {
		s.D.Set("is_hydrated", *s.Res.IsHydrated)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	if s.Res.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := BootVolumeSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	if s.Res.VpusPerGB != nil {
		s.D.Set("vpus_per_gb", strconv.FormatInt(*s.Res.VpusPerGB, 10))
	}

	// Add backup policy id from the other API
	backupPolicyId, err := getBackupPolicyId(s.Res.Id, s.Client)
	if err != nil {
		log.Printf("[ERROR] Received an error when fetching backup policy id %v", err)
	} else if backupPolicyId != nil {
		s.D.Set("backup_policy_id", backupPolicyId)
	}
	return nil
}

func (s *CoreBootVolumeResourceCrud) mapToBootVolumeReplicaDetails(fieldKeyFormat string) (oci_core.BootVolumeReplicaDetails, error) {
	result := oci_core.BootVolumeReplicaDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func BootVolumeReplicaInfoToMap(obj oci_core.BootVolumeReplicaInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BootVolumeReplicaId != nil {
		result["boot_volume_replica_id"] = string(*obj.BootVolumeReplicaId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *CoreBootVolumeResourceCrud) mapToBootVolumeSourceDetails(fieldKeyFormat string) (oci_core.BootVolumeSourceDetails, error) {
	var baseObject oci_core.BootVolumeSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("bootVolume"):
		details := oci_core.BootVolumeSourceFromBootVolumeDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("bootVolumeBackup"):
		details := oci_core.BootVolumeSourceFromBootVolumeBackupDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("bootVolumeReplica"):
		details := oci_core.BootVolumeSourceFromBootVolumeReplicaDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func BootVolumeSourceDetailsToMap(obj *oci_core.BootVolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.BootVolumeSourceFromBootVolumeDetails:
		result["type"] = "bootVolume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.BootVolumeSourceFromBootVolumeBackupDetails:
		result["type"] = "bootVolumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.BootVolumeSourceFromBootVolumeReplicaDetails:
		result["type"] = "bootVolumeReplica"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreBootVolumeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeBootVolumeCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BootVolumeId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeBootVolumeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
