resource "volcenginecc_cbr_backup_plan" "CBRBackupPlanDemo" {
  name      = "CBRBackupPlanDemo"
  policy_id = "policy-krn3z1wf9ubujxxxxx"
  resource_list = [
    {
      resource_type = "ECS"
      resource_id   = "res-krn4bq3sctbwvxxxxx"
      meta_information = {
        vepfs_meta = ""
        ecs_meta   = "{\"AutoBackupAllAttachedVolumes\":true,\"VolumeList\":[{\"VolumeId\":\"vol-3wt6m9uygg4kgxxxxx\",\"VolumeName\":\"emr-d6af13acb6966xxxxx-volume-1-2\",\"VolumeType\":\"ESSD_FlexPL\"},{\"VolumeId\":\"vol-3wt6m9vh3a3qxxxxx\",\"VolumeName\":\"emr-d6af13acb6966xxxxx-volume-2-2\",\"VolumeType\":\"ESSD_FlexPL\"}]}"
      }
    }
  ]
}