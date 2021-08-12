// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_iotwireless_task_definition", taskDefinitionResourceType)
}

// taskDefinitionResourceType returns the Terraform aws_iotwireless_task_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::TaskDefinition resource type.
func taskDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "TaskDefinition arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "TaskDefinition arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_create_tasks": {
			// Property: AutoCreateTasks
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			//   "type": "boolean"
			// }
			Description: "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			Type:        types.BoolType,
			Required:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the new wireless gateway task definition",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the new wireless gateway task definition",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan_update_gateway_task_entry": {
			// Property: LoRaWANUpdateGatewayTaskEntry
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CurrentVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"current_version": {
						// Property: CurrentVersion
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Optional: true,
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Optional: true,
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"update_version": {
						// Property: UpdateVersion
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Optional: true,
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Optional: true,
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the new resource.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the new resource.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"task_definition_type": {
			// Property: TaskDefinitionType
			// CloudFormation resource type schema:
			// {
			//   "description": "A filter to list only the wireless gateway task definitions that use this task definition type",
			//   "enum": [
			//     "UPDATE"
			//   ],
			//   "type": "string"
			// }
			Description: "A filter to list only the wireless gateway task definitions that use this task definition type",
			Type:        types.StringType,
			Optional:    true,
		},
		"update": {
			// Property: Update
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "LoRaWAN": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CurrentVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SigKeyCrc": {
			//           "format": "int64",
			//           "type": "integer"
			//         },
			//         "UpdateSignature": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "UpdateVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateDataRole": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "UpdateDataSource": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"lo_ra_wan": {
						// Property: LoRaWAN
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"current_version": {
									// Property: CurrentVersion
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Optional: true,
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Optional: true,
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"sig_key_crc": {
									// Property: SigKeyCrc
									Type:     types.NumberType,
									Optional: true,
								},
								"update_signature": {
									// Property: UpdateSignature
									Type:     types.StringType,
									Optional: true,
								},
								"update_version": {
									// Property: UpdateVersion
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Optional: true,
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Optional: true,
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"update_data_role": {
						// Property: UpdateDataRole
						Type:     types.StringType,
						Optional: true,
					},
					"update_data_source": {
						// Property: UpdateDataSource
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Creates a gateway task definition.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::TaskDefinition").WithTerraformTypeName("aws_iotwireless_task_definition").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_iotwireless_task_definition", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
