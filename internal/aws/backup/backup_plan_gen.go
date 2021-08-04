// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_backup_backup_plan", backupPlan)
}

// backupPlan returns the Terraform aws_backup_backup_plan resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Backup::BackupPlan resource type.
func backupPlan(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"backup_plan": {
			// Property: BackupPlan
			// PrimaryIdentifier: true
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "AdvancedBackupSettings": {
			         "items": {
			           "additionalProperties": false,
			           "properties": {
			             "BackupOptions": {
			               "type": "object"
			             },
			             "ResourceType": {
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/AdvancedBackupSettingResourceType",
			           "required": [
			             "BackupOptions",
			             "ResourceType"
			           ],
			           "type": "object"
			         },
			         "type": "array",
			         "uniqueItems": false
			       },
			       "BackupPlanName": {
			         "type": "string"
			       },
			       "BackupPlanRule": {
			         "items": {
			           "additionalProperties": false,
			           "properties": {
			             "CompletionWindowMinutes": {
			               "type": "number"
			             },
			             "CopyActions": {
			               "items": {
			                 "additionalProperties": false,
			                 "properties": {
			                   "DestinationBackupVaultArn": {
			                     "type": "string"
			                   },
			                   "Lifecycle": {
			                     "additionalProperties": false,
			                     "properties": {
			                       "DeleteAfterDays": {
			                         "type": "number"
			                       },
			                       "MoveToColdStorageAfterDays": {
			                         "type": "number"
			                       }
			                     },
			                     "$ref": "#/definitions/LifecycleResourceType",
			                     "type": "object"
			                   }
			                 },
			                 "$ref": "#/definitions/CopyActionResourceType",
			                 "required": [
			                   "DestinationBackupVaultArn"
			                 ],
			                 "type": "object"
			               },
			               "type": "array",
			               "uniqueItems": false
			             },
			             "EnableContinuousBackup": {
			               "type": "boolean"
			             },
			             "Lifecycle": {
			               "additionalProperties": false,
			               "properties": {
			                 "DeleteAfterDays": {
			                   "type": "number"
			                 },
			                 "MoveToColdStorageAfterDays": {
			                   "type": "number"
			                 }
			               },
			               "$ref": "#/definitions/LifecycleResourceType",
			               "type": "object"
			             },
			             "RecoveryPointTags": {
			               "additionalProperties": false,
			               "patternProperties": {
			                 "^.{1,128}$": {
			                   "type": "string"
			                 }
			               },
			               "type": "object"
			             },
			             "RuleName": {
			               "type": "string"
			             },
			             "ScheduleExpression": {
			               "type": "string"
			             },
			             "StartWindowMinutes": {
			               "type": "number"
			             },
			             "TargetBackupVault": {
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/BackupRuleResourceType",
			           "required": [
			             "TargetBackupVault",
			             "RuleName"
			           ],
			           "type": "object"
			         },
			         "type": "array",
			         "uniqueItems": false
			       }
			     },
			     "$ref": "#/definitions/BackupPlanResourceType",
			     "required": [
			       "BackupPlanName",
			       "BackupPlanRule"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"advanced_backup_settings": {
						// Property: AdvancedBackupSettings
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "BackupOptions": {
						           "type": "object"
						         },
						         "ResourceType": {
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/AdvancedBackupSettingResourceType",
						       "required": [
						         "BackupOptions",
						         "ResourceType"
						       ],
						       "type": "object"
						     },
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"backup_options": {
									// Property: BackupOptions
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "object"
									   }
									*/
									Type:     types.MapType{ElemType: types.StringType},
									Required: true,
								},
								"resource_type": {
									// Property: ResourceType
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"backup_plan_name": {
						// Property: BackupPlanName
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"backup_plan_rule": {
						// Property: BackupPlanRule
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "CompletionWindowMinutes": {
						           "type": "number"
						         },
						         "CopyActions": {
						           "items": {
						             "additionalProperties": false,
						             "properties": {
						               "DestinationBackupVaultArn": {
						                 "type": "string"
						               },
						               "Lifecycle": {
						                 "additionalProperties": false,
						                 "properties": {
						                   "DeleteAfterDays": {
						                     "type": "number"
						                   },
						                   "MoveToColdStorageAfterDays": {
						                     "type": "number"
						                   }
						                 },
						                 "$ref": "#/definitions/LifecycleResourceType",
						                 "type": "object"
						               }
						             },
						             "$ref": "#/definitions/CopyActionResourceType",
						             "required": [
						               "DestinationBackupVaultArn"
						             ],
						             "type": "object"
						           },
						           "type": "array",
						           "uniqueItems": false
						         },
						         "EnableContinuousBackup": {
						           "type": "boolean"
						         },
						         "Lifecycle": {
						           "additionalProperties": false,
						           "properties": {
						             "DeleteAfterDays": {
						               "type": "number"
						             },
						             "MoveToColdStorageAfterDays": {
						               "type": "number"
						             }
						           },
						           "$ref": "#/definitions/LifecycleResourceType",
						           "type": "object"
						         },
						         "RecoveryPointTags": {
						           "additionalProperties": false,
						           "patternProperties": {
						             "^.{1,128}$": {
						               "type": "string"
						             }
						           },
						           "type": "object"
						         },
						         "RuleName": {
						           "type": "string"
						         },
						         "ScheduleExpression": {
						           "type": "string"
						         },
						         "StartWindowMinutes": {
						           "type": "number"
						         },
						         "TargetBackupVault": {
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/BackupRuleResourceType",
						       "required": [
						         "TargetBackupVault",
						         "RuleName"
						       ],
						       "type": "object"
						     },
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"completion_window_minutes": {
									// Property: CompletionWindowMinutes
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "number"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
								"copy_actions": {
									// Property: CopyActions
									// CloudFormation resource type schema:
									/*
									   {
									     "items": {
									       "additionalProperties": false,
									       "properties": {
									         "DestinationBackupVaultArn": {
									           "type": "string"
									         },
									         "Lifecycle": {
									           "additionalProperties": false,
									           "properties": {
									             "DeleteAfterDays": {
									               "type": "number"
									             },
									             "MoveToColdStorageAfterDays": {
									               "type": "number"
									             }
									           },
									           "$ref": "#/definitions/LifecycleResourceType",
									           "type": "object"
									         }
									       },
									       "$ref": "#/definitions/CopyActionResourceType",
									       "required": [
									         "DestinationBackupVaultArn"
									       ],
									       "type": "object"
									     },
									     "type": "array",
									     "uniqueItems": false
									   }
									*/
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"destination_backup_vault_arn": {
												// Property: DestinationBackupVaultArn
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Required: true,
											},
											"lifecycle": {
												// Property: Lifecycle
												// CloudFormation resource type schema:
												/*
												   {
												     "additionalProperties": false,
												     "properties": {
												       "DeleteAfterDays": {
												         "type": "number"
												       },
												       "MoveToColdStorageAfterDays": {
												         "type": "number"
												       }
												     },
												     "$ref": "#/definitions/LifecycleResourceType",
												     "type": "object"
												   }
												*/
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"delete_after_days": {
															// Property: DeleteAfterDays
															// CloudFormation resource type schema:
															/*
															   {
															     "type": "number"
															   }
															*/
															Type:     types.NumberType,
															Optional: true,
														},
														"move_to_cold_storage_after_days": {
															// Property: MoveToColdStorageAfterDays
															// CloudFormation resource type schema:
															/*
															   {
															     "type": "number"
															   }
															*/
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"enable_continuous_backup": {
									// Property: EnableContinuousBackup
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "boolean"
									   }
									*/
									Type:     types.BoolType,
									Optional: true,
								},
								"lifecycle": {
									// Property: Lifecycle
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "properties": {
									       "DeleteAfterDays": {
									         "type": "number"
									       },
									       "MoveToColdStorageAfterDays": {
									         "type": "number"
									       }
									     },
									     "$ref": "#/definitions/LifecycleResourceType",
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"delete_after_days": {
												// Property: DeleteAfterDays
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "number"
												   }
												*/
												Type:     types.NumberType,
												Optional: true,
											},
											"move_to_cold_storage_after_days": {
												// Property: MoveToColdStorageAfterDays
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "number"
												   }
												*/
												Type:     types.NumberType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"recovery_point_tags": {
									// Property: RecoveryPointTags
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "patternProperties": {
									       "^.{1,128}$": {
									         "type": "string"
									       }
									     },
									     "type": "object"
									   }
									*/
									// Pattern: "^.{1,128}$"
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"rule_name": {
									// Property: RuleName
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"schedule_expression": {
									// Property: ScheduleExpression
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"start_window_minutes": {
									// Property: StartWindowMinutes
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "number"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
								"target_backup_vault": {
									// Property: TargetBackupVault
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"backup_plan_arn": {
			// Property: BackupPlanArn
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"backup_plan_id": {
			// Property: BackupPlanId
			// PrimaryIdentifier: true
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"backup_plan_tags": {
			// Property: BackupPlanTags
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "patternProperties": {
			       "^.{1,128}$": {
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			// Pattern: "^.{1,128}$"
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: `Resource Type definition for AWS::Backup::BackupPlan`,
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupPlan").WithTerraformTypeName("aws_backup_backup_plan").WithTerraformSchema(schema).WithPrimaryIdentifierPath("/properties/BackupPlanId")

	opts = opts.WithCreateTimeoutInMinutes(0).WithUpdateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_backup_backup_plan", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
