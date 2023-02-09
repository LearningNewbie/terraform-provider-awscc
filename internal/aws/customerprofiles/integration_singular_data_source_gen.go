// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package customerprofiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_customerprofiles_integration", integrationDataSource)
}

// integrationDataSource returns the Terraform awscc_customerprofiles_integration data source.
// This Terraform data source corresponds to the CloudFormation AWS::CustomerProfiles::Integration resource.
func integrationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this integration got created",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this integration got created",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the domain.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FlowDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Description": {
		//	      "maxLength": 2048,
		//	      "pattern": "[\\w!@#\\-.?,\\s]*",
		//	      "type": "string"
		//	    },
		//	    "FlowName": {
		//	      "maxLength": 256,
		//	      "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
		//	      "type": "string"
		//	    },
		//	    "KmsArn": {
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "pattern": "arn:aws:kms:.*:[0-9]+:.*",
		//	      "type": "string"
		//	    },
		//	    "SourceFlowConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ConnectorProfileName": {
		//	          "maxLength": 256,
		//	          "pattern": "[\\w/!@#+=.-]+",
		//	          "type": "string"
		//	        },
		//	        "ConnectorType": {
		//	          "enum": [
		//	            "Salesforce",
		//	            "Marketo",
		//	            "ServiceNow",
		//	            "Zendesk",
		//	            "S3"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "IncrementalPullConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "DatetimeTypeFieldName": {
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "SourceConnectorProperties": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Marketo": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Object": {
		//	                  "additionalProperties": false,
		//	                  "maxLength": 512,
		//	                  "pattern": "\\S+",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Object"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "S3": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "BucketName": {
		//	                  "maxLength": 63,
		//	                  "minLength": 3,
		//	                  "pattern": "\\S+",
		//	                  "type": "string"
		//	                },
		//	                "BucketPrefix": {
		//	                  "maxLength": 512,
		//	                  "pattern": ".*",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "BucketName"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "Salesforce": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "EnableDynamicFieldUpdate": {
		//	                  "type": "boolean"
		//	                },
		//	                "IncludeDeletedRecords": {
		//	                  "type": "boolean"
		//	                },
		//	                "Object": {
		//	                  "additionalProperties": false,
		//	                  "maxLength": 512,
		//	                  "pattern": "\\S+",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Object"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "ServiceNow": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Object": {
		//	                  "additionalProperties": false,
		//	                  "maxLength": 512,
		//	                  "pattern": "\\S+",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Object"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "Zendesk": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Object": {
		//	                  "additionalProperties": false,
		//	                  "maxLength": 512,
		//	                  "pattern": "\\S+",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Object"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ConnectorType",
		//	        "SourceConnectorProperties"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Tasks": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ConnectorOperator": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Marketo": {
		//	                "enum": [
		//	                  "PROJECTION",
		//	                  "LESS_THAN",
		//	                  "GREATER_THAN",
		//	                  "BETWEEN",
		//	                  "ADDITION",
		//	                  "MULTIPLICATION",
		//	                  "DIVISION",
		//	                  "SUBTRACTION",
		//	                  "MASK_ALL",
		//	                  "MASK_FIRST_N",
		//	                  "MASK_LAST_N",
		//	                  "VALIDATE_NON_NULL",
		//	                  "VALIDATE_NON_ZERO",
		//	                  "VALIDATE_NON_NEGATIVE",
		//	                  "VALIDATE_NUMERIC",
		//	                  "NO_OP"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "S3": {
		//	                "enum": [
		//	                  "PROJECTION",
		//	                  "LESS_THAN",
		//	                  "GREATER_THAN",
		//	                  "BETWEEN",
		//	                  "LESS_THAN_OR_EQUAL_TO",
		//	                  "GREATER_THAN_OR_EQUAL_TO",
		//	                  "EQUAL_TO",
		//	                  "NOT_EQUAL_TO",
		//	                  "ADDITION",
		//	                  "MULTIPLICATION",
		//	                  "DIVISION",
		//	                  "SUBTRACTION",
		//	                  "MASK_ALL",
		//	                  "MASK_FIRST_N",
		//	                  "MASK_LAST_N",
		//	                  "VALIDATE_NON_NULL",
		//	                  "VALIDATE_NON_ZERO",
		//	                  "VALIDATE_NON_NEGATIVE",
		//	                  "VALIDATE_NUMERIC",
		//	                  "NO_OP"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Salesforce": {
		//	                "enum": [
		//	                  "PROJECTION",
		//	                  "LESS_THAN",
		//	                  "GREATER_THAN",
		//	                  "CONTAINS",
		//	                  "BETWEEN",
		//	                  "LESS_THAN_OR_EQUAL_TO",
		//	                  "GREATER_THAN_OR_EQUAL_TO",
		//	                  "EQUAL_TO",
		//	                  "NOT_EQUAL_TO",
		//	                  "ADDITION",
		//	                  "MULTIPLICATION",
		//	                  "DIVISION",
		//	                  "SUBTRACTION",
		//	                  "MASK_ALL",
		//	                  "MASK_FIRST_N",
		//	                  "MASK_LAST_N",
		//	                  "VALIDATE_NON_NULL",
		//	                  "VALIDATE_NON_ZERO",
		//	                  "VALIDATE_NON_NEGATIVE",
		//	                  "VALIDATE_NUMERIC",
		//	                  "NO_OP"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "ServiceNow": {
		//	                "enum": [
		//	                  "PROJECTION",
		//	                  "LESS_THAN",
		//	                  "GREATER_THAN",
		//	                  "CONTAINS",
		//	                  "BETWEEN",
		//	                  "LESS_THAN_OR_EQUAL_TO",
		//	                  "GREATER_THAN_OR_EQUAL_TO",
		//	                  "EQUAL_TO",
		//	                  "NOT_EQUAL_TO",
		//	                  "ADDITION",
		//	                  "MULTIPLICATION",
		//	                  "DIVISION",
		//	                  "SUBTRACTION",
		//	                  "MASK_ALL",
		//	                  "MASK_FIRST_N",
		//	                  "MASK_LAST_N",
		//	                  "VALIDATE_NON_NULL",
		//	                  "VALIDATE_NON_ZERO",
		//	                  "VALIDATE_NON_NEGATIVE",
		//	                  "VALIDATE_NUMERIC",
		//	                  "NO_OP"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Zendesk": {
		//	                "enum": [
		//	                  "PROJECTION",
		//	                  "GREATER_THAN",
		//	                  "ADDITION",
		//	                  "MULTIPLICATION",
		//	                  "DIVISION",
		//	                  "SUBTRACTION",
		//	                  "MASK_ALL",
		//	                  "MASK_FIRST_N",
		//	                  "MASK_LAST_N",
		//	                  "VALIDATE_NON_NULL",
		//	                  "VALIDATE_NON_ZERO",
		//	                  "VALIDATE_NON_NEGATIVE",
		//	                  "VALIDATE_NUMERIC",
		//	                  "NO_OP"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "DestinationField": {
		//	            "maxLength": 256,
		//	            "pattern": ".*",
		//	            "type": "string"
		//	          },
		//	          "SourceFields": {
		//	            "items": {
		//	              "maxLength": 2048,
		//	              "pattern": ".*",
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "TaskProperties": {
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "OperatorPropertyKey": {
		//	                  "enum": [
		//	                    "VALUE",
		//	                    "VALUES",
		//	                    "DATA_TYPE",
		//	                    "UPPER_BOUND",
		//	                    "LOWER_BOUND",
		//	                    "SOURCE_DATA_TYPE",
		//	                    "DESTINATION_DATA_TYPE",
		//	                    "VALIDATION_ACTION",
		//	                    "MASK_VALUE",
		//	                    "MASK_LENGTH",
		//	                    "TRUNCATE_LENGTH",
		//	                    "MATH_OPERATION_FIELDS_ORDER",
		//	                    "CONCAT_FORMAT",
		//	                    "SUBFIELD_CATEGORY_MAP"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Property": {
		//	                  "maxLength": 2048,
		//	                  "pattern": ".+",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "OperatorPropertyKey",
		//	                "Property"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "TaskType": {
		//	            "enum": [
		//	              "Arithmetic",
		//	              "Filter",
		//	              "Map",
		//	              "Mask",
		//	              "Merge",
		//	              "Truncate",
		//	              "Validate"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "SourceFields",
		//	          "TaskType"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "TriggerConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "TriggerProperties": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Scheduled": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "DataPullMode": {
		//	                  "enum": [
		//	                    "Incremental",
		//	                    "Complete"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "FirstExecutionFrom": {
		//	                  "type": "number"
		//	                },
		//	                "ScheduleEndTime": {
		//	                  "type": "number"
		//	                },
		//	                "ScheduleExpression": {
		//	                  "maxLength": 256,
		//	                  "pattern": ".*",
		//	                  "type": "string"
		//	                },
		//	                "ScheduleOffset": {
		//	                  "maximum": 36000,
		//	                  "minimum": 0,
		//	                  "type": "integer"
		//	                },
		//	                "ScheduleStartTime": {
		//	                  "type": "number"
		//	                },
		//	                "Timezone": {
		//	                  "maxLength": 256,
		//	                  "pattern": ".*",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "ScheduleExpression"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "TriggerType": {
		//	          "enum": [
		//	            "Scheduled",
		//	            "Event",
		//	            "OnDemand"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TriggerType"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "FlowName",
		//	    "KmsArn",
		//	    "Tasks",
		//	    "TriggerConfig",
		//	    "SourceFlowConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"flow_definition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FlowName
				"flow_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: KmsArn
				"kms_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourceFlowConfig
				"source_flow_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ConnectorProfileName
						"connector_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ConnectorType
						"connector_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: IncrementalPullConfig
						"incremental_pull_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DatetimeTypeFieldName
								"datetime_type_field_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SourceConnectorProperties
						"source_connector_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Marketo
								"marketo": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Object
										"object": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: S3
								"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: BucketName
										"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: BucketPrefix
										"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Salesforce
								"salesforce": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: EnableDynamicFieldUpdate
										"enable_dynamic_field_update": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: IncludeDeletedRecords
										"include_deleted_records": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Object
										"object": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: ServiceNow
								"service_now": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Object
										"object": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Zendesk
								"zendesk": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Object
										"object": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Tasks
				"tasks": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectorOperator
							"connector_operator": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Marketo
									"marketo": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: S3
									"s3": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Salesforce
									"salesforce": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: ServiceNow
									"service_now": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Zendesk
									"zendesk": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: DestinationField
							"destination_field": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: SourceFields
							"source_fields": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TaskProperties
							"task_properties": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: OperatorPropertyKey
										"operator_property_key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Property
										"property": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: TaskType
							"task_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TriggerConfig
				"trigger_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TriggerProperties
						"trigger_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Scheduled
								"scheduled": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DataPullMode
										"data_pull_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: FirstExecutionFrom
										"first_execution_from": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ScheduleEndTime
										"schedule_end_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ScheduleExpression
										"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ScheduleOffset
										"schedule_offset": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ScheduleStartTime
										"schedule_start_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Timezone
										"timezone": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TriggerType
						"trigger_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this integration got last updated at",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this integration got last updated at",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObjectTypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the ObjectType defined for the 3rd party data in Profile Service",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"object_type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the ObjectType defined for the 3rd party data in Profile Service",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObjectTypeNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The mapping between 3rd party event types and ObjectType names",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"object_type_names": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The mapping between 3rd party event types and ObjectType names",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags (keys and values) associated with the integration",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags (keys and values) associated with the integration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Uri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URI of the S3 bucket or any other type of data source.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URI of the S3 bucket or any other type of data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CustomerProfiles::Integration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Integration").WithTerraformTypeName("awscc_customerprofiles_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":                 "BucketName",
		"bucket_prefix":               "BucketPrefix",
		"connector_operator":          "ConnectorOperator",
		"connector_profile_name":      "ConnectorProfileName",
		"connector_type":              "ConnectorType",
		"created_at":                  "CreatedAt",
		"data_pull_mode":              "DataPullMode",
		"datetime_type_field_name":    "DatetimeTypeFieldName",
		"description":                 "Description",
		"destination_field":           "DestinationField",
		"domain_name":                 "DomainName",
		"enable_dynamic_field_update": "EnableDynamicFieldUpdate",
		"first_execution_from":        "FirstExecutionFrom",
		"flow_definition":             "FlowDefinition",
		"flow_name":                   "FlowName",
		"include_deleted_records":     "IncludeDeletedRecords",
		"incremental_pull_config":     "IncrementalPullConfig",
		"key":                         "Key",
		"kms_arn":                     "KmsArn",
		"last_updated_at":             "LastUpdatedAt",
		"marketo":                     "Marketo",
		"object":                      "Object",
		"object_type_name":            "ObjectTypeName",
		"object_type_names":           "ObjectTypeNames",
		"operator_property_key":       "OperatorPropertyKey",
		"property":                    "Property",
		"s3":                          "S3",
		"salesforce":                  "Salesforce",
		"schedule_end_time":           "ScheduleEndTime",
		"schedule_expression":         "ScheduleExpression",
		"schedule_offset":             "ScheduleOffset",
		"schedule_start_time":         "ScheduleStartTime",
		"scheduled":                   "Scheduled",
		"service_now":                 "ServiceNow",
		"source_connector_properties": "SourceConnectorProperties",
		"source_fields":               "SourceFields",
		"source_flow_config":          "SourceFlowConfig",
		"tags":                        "Tags",
		"task_properties":             "TaskProperties",
		"task_type":                   "TaskType",
		"tasks":                       "Tasks",
		"timezone":                    "Timezone",
		"trigger_config":              "TriggerConfig",
		"trigger_properties":          "TriggerProperties",
		"trigger_type":                "TriggerType",
		"uri":                         "Uri",
		"value":                       "Value",
		"zendesk":                     "Zendesk",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
