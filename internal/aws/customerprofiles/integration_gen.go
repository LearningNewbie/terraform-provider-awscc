// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

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
	registry.AddResourceTypeFactory("aws_customerprofiles_integration", integrationResourceType)
}

// integrationResourceType returns the Terraform aws_customerprofiles_integration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CustomerProfiles::Integration resource type.
func integrationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got created",
			//   "type": "string"
			// }
			Description: "The time of this integration got created",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the domain.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Required:    true,
			// DomainName is a force-new attribute.
		},
		"flow_definition": {
			// Property: FlowDefinition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Description": {
			//       "maxLength": 2048,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "FlowName": {
			//       "maxLength": 256,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "KmsArn": {
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SourceFlowConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ConnectorProfileName": {
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "ConnectorType": {
			//           "enum": [
			//             "Salesforce",
			//             "Marketo",
			//             "ServiceNow",
			//             "Zendesk",
			//             "S3"
			//           ],
			//           "type": "string"
			//         },
			//         "IncrementalPullConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "DatetimeTypeFieldName": {
			//               "maxLength": 256,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SourceConnectorProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Marketo": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "S3": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "BucketName": {
			//                   "maxLength": 63,
			//                   "minLength": 3,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "BucketPrefix": {
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "BucketName"
			//               ],
			//               "type": "object"
			//             },
			//             "Salesforce": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "EnableDynamicFieldUpdate": {
			//                   "type": "boolean"
			//                 },
			//                 "IncludeDeletedRecords": {
			//                   "type": "boolean"
			//                 },
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "ServiceNow": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "Zendesk": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "ConnectorType",
			//         "SourceConnectorProperties"
			//       ],
			//       "type": "object"
			//     },
			//     "Tasks": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ConnectorOperator": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Marketo": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "S3": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Salesforce": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "ServiceNow": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Zendesk": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "GREATER_THAN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "DestinationField": {
			//             "maxLength": 256,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "SourceFields": {
			//             "items": {
			//               "maxLength": 2048,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "TaskProperties": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "OperatorPropertyKey": {
			//                   "enum": [
			//                     "VALUE",
			//                     "VALUES",
			//                     "DATA_TYPE",
			//                     "UPPER_BOUND",
			//                     "LOWER_BOUND",
			//                     "SOURCE_DATA_TYPE",
			//                     "DESTINATION_DATA_TYPE",
			//                     "VALIDATION_ACTION",
			//                     "MASK_VALUE",
			//                     "MASK_LENGTH",
			//                     "TRUNCATE_LENGTH",
			//                     "MATH_OPERATION_FIELDS_ORDER",
			//                     "CONCAT_FORMAT",
			//                     "SUBFIELD_CATEGORY_MAP"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Property": {
			//                   "maxLength": 2048,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "OperatorPropertyKey",
			//                 "Property"
			//               ],
			//               "type": "object"
			//             },
			//             "type": "array"
			//           },
			//           "TaskType": {
			//             "enum": [
			//               "Arithmetic",
			//               "Filter",
			//               "Map",
			//               "Mask",
			//               "Merge",
			//               "Truncate",
			//               "Validate"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "SourceFields",
			//           "TaskType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "TriggerConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "TriggerProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Scheduled": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "DataPullMode": {
			//                   "enum": [
			//                     "Incremental",
			//                     "Complete"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "FirstExecutionFrom": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleEndTime": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleExpression": {
			//                   "maxLength": 256,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "ScheduleOffset": {
			//                   "type": "integer"
			//                 },
			//                 "ScheduleStartTime": {
			//                   "type": "number"
			//                 },
			//                 "Timezone": {
			//                   "maxLength": 256,
			//                   "pattern": "",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "ScheduleExpression"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "TriggerType": {
			//           "enum": [
			//             "Scheduled",
			//             "Event",
			//             "OnDemand"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TriggerType"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "FlowName",
			//     "KmsArn",
			//     "Tasks",
			//     "TriggerConfig",
			//     "SourceFlowConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Optional: true,
					},
					"flow_name": {
						// Property: FlowName
						Type:     types.StringType,
						Required: true,
					},
					"kms_arn": {
						// Property: KmsArn
						Type:     types.StringType,
						Required: true,
					},
					"source_flow_config": {
						// Property: SourceFlowConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"connector_profile_name": {
									// Property: ConnectorProfileName
									Type:     types.StringType,
									Optional: true,
								},
								"connector_type": {
									// Property: ConnectorType
									Type:     types.StringType,
									Required: true,
								},
								"incremental_pull_config": {
									// Property: IncrementalPullConfig
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"datetime_type_field_name": {
												// Property: DatetimeTypeFieldName
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"source_connector_properties": {
									// Property: SourceConnectorProperties
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"marketo": {
												// Property: Marketo
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"s3": {
												// Property: S3
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"bucket_name": {
															// Property: BucketName
															Type:     types.StringType,
															Required: true,
														},
														"bucket_prefix": {
															// Property: BucketPrefix
															Type:     types.StringType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"salesforce": {
												// Property: Salesforce
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"enable_dynamic_field_update": {
															// Property: EnableDynamicFieldUpdate
															Type:     types.BoolType,
															Optional: true,
														},
														"include_deleted_records": {
															// Property: IncludeDeletedRecords
															Type:     types.BoolType,
															Optional: true,
														},
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"service_now": {
												// Property: ServiceNow
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"zendesk": {
												// Property: Zendesk
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
										},
									),
									Required: true,
								},
							},
						),
						Required: true,
					},
					"tasks": {
						// Property: Tasks
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"connector_operator": {
									// Property: ConnectorOperator
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"marketo": {
												// Property: Marketo
												Type:     types.StringType,
												Optional: true,
											},
											"s3": {
												// Property: S3
												Type:     types.StringType,
												Optional: true,
											},
											"salesforce": {
												// Property: Salesforce
												Type:     types.StringType,
												Optional: true,
											},
											"service_now": {
												// Property: ServiceNow
												Type:     types.StringType,
												Optional: true,
											},
											"zendesk": {
												// Property: Zendesk
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"destination_field": {
									// Property: DestinationField
									Type:     types.StringType,
									Optional: true,
								},
								"source_fields": {
									// Property: SourceFields
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
								},
								"task_properties": {
									// Property: TaskProperties
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"operator_property_key": {
												// Property: OperatorPropertyKey
												Type:     types.StringType,
												Required: true,
											},
											"property": {
												// Property: Property
												Type:     types.StringType,
												Required: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"task_type": {
									// Property: TaskType
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Required: true,
					},
					"trigger_config": {
						// Property: TriggerConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"trigger_properties": {
									// Property: TriggerProperties
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"scheduled": {
												// Property: Scheduled
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"data_pull_mode": {
															// Property: DataPullMode
															Type:     types.StringType,
															Optional: true,
														},
														"first_execution_from": {
															// Property: FirstExecutionFrom
															Type:     types.NumberType,
															Optional: true,
														},
														"schedule_end_time": {
															// Property: ScheduleEndTime
															Type:     types.NumberType,
															Optional: true,
														},
														"schedule_expression": {
															// Property: ScheduleExpression
															Type:     types.StringType,
															Required: true,
														},
														"schedule_offset": {
															// Property: ScheduleOffset
															Type:     types.NumberType,
															Optional: true,
														},
														"schedule_start_time": {
															// Property: ScheduleStartTime
															Type:     types.NumberType,
															Optional: true,
														},
														"timezone": {
															// Property: Timezone
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
								"trigger_type": {
									// Property: TriggerType
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			// FlowDefinition is a write-only attribute.
		},
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got last updated at",
			//   "type": "string"
			// }
			Description: "The time of this integration got last updated at",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_type_name": {
			// Property: ObjectTypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the ObjectType defined for the 3rd party data in Profile Service",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the ObjectType defined for the 3rd party data in Profile Service",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the integration",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the integration",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"uri": {
			// Property: Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URI of the S3 bucket or any other type of data source.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The URI of the S3 bucket or any other type of data source.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Uri is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The resource schema for creating an Amazon Connect Customer Profiles Integration.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Integration").WithTerraformTypeName("aws_customerprofiles_integration").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/FlowDefinition",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_customerprofiles_integration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
