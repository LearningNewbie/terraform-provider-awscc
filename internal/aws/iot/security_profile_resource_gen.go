// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_iot_security_profile", securityProfileResource)
}

// securityProfileResource returns the Terraform awscc_iot_security_profile resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoT::SecurityProfile resource.
func securityProfileResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalMetricsToRetainV2
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The metric you want to retain. Dimensions are optional.",
		//	    "properties": {
		//	      "Metric": {
		//	        "description": "What is measured by the behavior.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[a-zA-Z0-9:_-]+",
		//	        "type": "string"
		//	      },
		//	      "MetricDimension": {
		//	        "additionalProperties": false,
		//	        "description": "The dimension of a metric.",
		//	        "properties": {
		//	          "DimensionName": {
		//	            "description": "A unique identifier for the dimension.",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "[a-zA-Z0-9:_-]+",
		//	            "type": "string"
		//	          },
		//	          "Operator": {
		//	            "description": "Defines how the dimensionValues of a dimension are interpreted.",
		//	            "enum": [
		//	              "IN",
		//	              "NOT_IN"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DimensionName"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Metric"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"additional_metrics_to_retain_v2": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Metric
					"metric": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "What is measured by the behavior.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: MetricDimension
					"metric_dimension": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DimensionName
							"dimension_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A unique identifier for the dimension.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 128),
									stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Operator
							"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Defines how the dimensionValues of a dimension are interpreted.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"IN",
										"NOT_IN",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The dimension of a metric.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AlertTargets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the destinations to which alerts are sent.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "A structure containing the alert target ARN and the role ARN.",
		//	      "properties": {
		//	        "AlertTargetArn": {
		//	          "description": "The ARN of the notification target to which alerts are sent.",
		//	          "maxLength": 2048,
		//	          "type": "string"
		//	        },
		//	        "RoleArn": {
		//	          "description": "The ARN of the role that grants permission to send alerts to the notification target.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "AlertTargetArn",
		//	        "RoleArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"alert_targets":           // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AlertTargetArn
					"alert_target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ARN of the notification target to which alerts are sent.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(2048),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: RoleArn
					"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ARN of the role that grants permission to send alerts to the notification target.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(20, 2048),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies the destinations to which alerts are sent.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Behaviors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the behaviors that, when violated by a device (thing), cause an alert.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A security profile behavior.",
		//	    "properties": {
		//	      "Criteria": {
		//	        "additionalProperties": false,
		//	        "description": "The criteria by which the behavior is determined to be normal.",
		//	        "properties": {
		//	          "ComparisonOperator": {
		//	            "description": "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).",
		//	            "enum": [
		//	              "less-than",
		//	              "less-than-equals",
		//	              "greater-than",
		//	              "greater-than-equals",
		//	              "in-cidr-set",
		//	              "not-in-cidr-set",
		//	              "in-port-set",
		//	              "not-in-port-set",
		//	              "in-set",
		//	              "not-in-set"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "ConsecutiveDatapointsToAlarm": {
		//	            "description": "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.",
		//	            "maximum": 10,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "ConsecutiveDatapointsToClear": {
		//	            "description": "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.",
		//	            "maximum": 10,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "DurationSeconds": {
		//	            "description": "Use this to specify the time duration over which the behavior is evaluated.",
		//	            "type": "integer"
		//	          },
		//	          "MlDetectionConfig": {
		//	            "additionalProperties": false,
		//	            "description": "The configuration of an ML Detect Security Profile.",
		//	            "properties": {
		//	              "ConfidenceLevel": {
		//	                "description": "The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.",
		//	                "enum": [
		//	                  "LOW",
		//	                  "MEDIUM",
		//	                  "HIGH"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "StatisticalThreshold": {
		//	            "additionalProperties": false,
		//	            "description": "A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.",
		//	            "properties": {
		//	              "Statistic": {
		//	                "description": "The percentile which resolves to a threshold value by which compliance with a behavior is determined",
		//	                "enum": [
		//	                  "Average",
		//	                  "p0",
		//	                  "p0.1",
		//	                  "p0.01",
		//	                  "p1",
		//	                  "p10",
		//	                  "p50",
		//	                  "p90",
		//	                  "p99",
		//	                  "p99.9",
		//	                  "p99.99",
		//	                  "p100"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Value": {
		//	            "additionalProperties": false,
		//	            "description": "The value to be compared with the metric.",
		//	            "properties": {
		//	              "Cidrs": {
		//	                "description": "If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array",
		//	                "uniqueItems": true
		//	              },
		//	              "Count": {
		//	                "description": "If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.",
		//	                "minimum": 0,
		//	                "type": "string"
		//	              },
		//	              "Number": {
		//	                "description": "The numeral value of a metric.",
		//	                "type": "number"
		//	              },
		//	              "Numbers": {
		//	                "description": "The numeral values of a metric.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "number"
		//	                },
		//	                "type": "array",
		//	                "uniqueItems": true
		//	              },
		//	              "Ports": {
		//	                "description": "If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "maximum": 65535,
		//	                  "minimum": 0,
		//	                  "type": "integer"
		//	                },
		//	                "type": "array",
		//	                "uniqueItems": true
		//	              },
		//	              "Strings": {
		//	                "description": "The string values of a metric.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array",
		//	                "uniqueItems": true
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Metric": {
		//	        "description": "What is measured by the behavior.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[a-zA-Z0-9:_-]+",
		//	        "type": "string"
		//	      },
		//	      "MetricDimension": {
		//	        "additionalProperties": false,
		//	        "description": "The dimension of a metric.",
		//	        "properties": {
		//	          "DimensionName": {
		//	            "description": "A unique identifier for the dimension.",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "[a-zA-Z0-9:_-]+",
		//	            "type": "string"
		//	          },
		//	          "Operator": {
		//	            "description": "Defines how the dimensionValues of a dimension are interpreted.",
		//	            "enum": [
		//	              "IN",
		//	              "NOT_IN"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DimensionName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Name": {
		//	        "description": "The name for the behavior.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[a-zA-Z0-9:_-]+",
		//	        "type": "string"
		//	      },
		//	      "SuppressAlerts": {
		//	        "description": "Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.",
		//	        "type": "boolean"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxLength": 100,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"behaviors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Criteria
					"criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ComparisonOperator
							"comparison_operator": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"less-than",
										"less-than-equals",
										"greater-than",
										"greater-than-equals",
										"in-cidr-set",
										"not-in-cidr-set",
										"in-port-set",
										"not-in-port-set",
										"in-set",
										"not-in-set",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ConsecutiveDatapointsToAlarm
							"consecutive_datapoints_to_alarm": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 10),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ConsecutiveDatapointsToClear
							"consecutive_datapoints_to_clear": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 10),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DurationSeconds
							"duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Use this to specify the time duration over which the behavior is evaluated.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MlDetectionConfig
							"ml_detection_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ConfidenceLevel
									"confidence_level": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.",
										Optional:    true,
										Computed:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"LOW",
												"MEDIUM",
												"HIGH",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The configuration of an ML Detect Security Profile.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: StatisticalThreshold
							"statistical_threshold": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Statistic
									"statistic": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The percentile which resolves to a threshold value by which compliance with a behavior is determined",
										Optional:    true,
										Computed:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"Average",
												"p0",
												"p0.1",
												"p0.01",
												"p1",
												"p10",
												"p50",
												"p90",
												"p99",
												"p99.9",
												"p99.99",
												"p100",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Cidrs
									"cidrs": schema.SetAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
											setplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Count
									"count": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Number
									"number": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "The numeral value of a metric.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
											float64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Numbers
									"numbers": schema.SetAttribute{ /*START ATTRIBUTE*/
										ElementType: types.Float64Type,
										Description: "The numeral values of a metric.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
											setplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Ports
									"ports": schema.SetAttribute{ /*START ATTRIBUTE*/
										ElementType: types.Int64Type,
										Description: "If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.",
										Optional:    true,
										Computed:    true,
										Validators: []validator.Set{ /*START VALIDATORS*/
											setvalidator.ValueInt64sAre(
												int64validator.Between(0, 65535),
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
											setplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Strings
									"strings": schema.SetAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The string values of a metric.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
											setplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The value to be compared with the metric.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The criteria by which the behavior is determined to be normal.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Metric
					"metric": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "What is measured by the behavior.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: MetricDimension
					"metric_dimension": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DimensionName
							"dimension_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A unique identifier for the dimension.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 128),
									stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Operator
							"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Defines how the dimensionValues of a dimension are interpreted.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"IN",
										"NOT_IN",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The dimension of a metric.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name for the behavior.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: SuppressAlerts
					"suppress_alerts": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies the behaviors that, when violated by a device (thing), cause an alert.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon resource name) of the created security profile.",
		//	  "type": "string"
		//	}
		"security_profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon resource name) of the created security profile.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityProfileDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the security profile.",
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"security_profile_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the security profile.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the security profile.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9:_-]+",
		//	  "type": "string"
		//	}
		"security_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the security profile.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata that can be used to manage the security profile.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Metadata that can be used to manage the security profile.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of target ARNs that the security profile is attached to.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The ARN of the target to which the security profile is attached.",
		//	    "maxLength": 2048,
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"target_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A set of target ARNs that the security profile is attached to.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(2048),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "A security profile defines a set of expected behaviors for devices in your account.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::SecurityProfile").WithTerraformTypeName("awscc_iot_security_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_metrics_to_retain_v2": "AdditionalMetricsToRetainV2",
		"alert_target_arn":                "AlertTargetArn",
		"alert_targets":                   "AlertTargets",
		"behaviors":                       "Behaviors",
		"cidrs":                           "Cidrs",
		"comparison_operator":             "ComparisonOperator",
		"confidence_level":                "ConfidenceLevel",
		"consecutive_datapoints_to_alarm": "ConsecutiveDatapointsToAlarm",
		"consecutive_datapoints_to_clear": "ConsecutiveDatapointsToClear",
		"count":                           "Count",
		"criteria":                        "Criteria",
		"dimension_name":                  "DimensionName",
		"duration_seconds":                "DurationSeconds",
		"key":                             "Key",
		"metric":                          "Metric",
		"metric_dimension":                "MetricDimension",
		"ml_detection_config":             "MlDetectionConfig",
		"name":                            "Name",
		"number":                          "Number",
		"numbers":                         "Numbers",
		"operator":                        "Operator",
		"ports":                           "Ports",
		"role_arn":                        "RoleArn",
		"security_profile_arn":            "SecurityProfileArn",
		"security_profile_description":    "SecurityProfileDescription",
		"security_profile_name":           "SecurityProfileName",
		"statistic":                       "Statistic",
		"statistical_threshold":           "StatisticalThreshold",
		"strings":                         "Strings",
		"suppress_alerts":                 "SuppressAlerts",
		"tags":                            "Tags",
		"target_arns":                     "TargetArns",
		"value":                           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
