// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rum

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rum_app_monitor", appMonitorDataSource)
}

// appMonitorDataSource returns the Terraform awscc_rum_app_monitor data source.
// This Terraform data source corresponds to the CloudFormation AWS::RUM::AppMonitor resource.
func appMonitorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppMonitorConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "AppMonitor configuration",
		//	  "properties": {
		//	    "AllowCookies": {
		//	      "description": "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
		//	      "type": "boolean"
		//	    },
		//	    "EnableXRay": {
		//	      "description": "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludedPages": {
		//	      "description": "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Page Url",
		//	        "maxLength": 1260,
		//	        "minLength": 1,
		//	        "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "FavoritePages": {
		//	      "description": "A list of pages in the RUM console that are to be displayed with a favorite icon.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "GuestRoleArn": {
		//	      "description": "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
		//	      "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	      "type": "string"
		//	    },
		//	    "IdentityPoolId": {
		//	      "description": "The ID of the identity pool that is used to authorize the sending of data to RUM.",
		//	      "maxLength": 55,
		//	      "minLength": 1,
		//	      "pattern": "[\\w-]+:[0-9a-f-]+",
		//	      "type": "string"
		//	    },
		//	    "IncludedPages": {
		//	      "description": "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Page Url",
		//	        "maxLength": 1260,
		//	        "minLength": 1,
		//	        "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "MetricDestinations": {
		//	      "description": "An array of structures which define the destinations and the metrics that you want to send.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An structure which defines the destination and the metrics that you want to send.",
		//	        "properties": {
		//	          "Destination": {
		//	            "description": "Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.",
		//	            "enum": [
		//	              "CloudWatch",
		//	              "Evidently"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "DestinationArn": {
		//	            "description": "Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.",
		//	            "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	            "type": "string"
		//	          },
		//	          "IamRoleArn": {
		//	            "description": "This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.\n\nThis parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.",
		//	            "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	            "type": "string"
		//	          },
		//	          "MetricDefinitions": {
		//	            "description": "An array of structures which define the metrics that you want to send.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "A single metric definition",
		//	              "properties": {
		//	                "DimensionKeys": {
		//	                  "additionalProperties": false,
		//	                  "description": "Use this field only if you are sending the metric to CloudWatch.\n\nThis field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. Valid values for the entries in this field are the following:\n\n\"metadata.pageId\": \"PageId\"\n\n\"metadata.browserName\": \"BrowserName\"\n\n\"metadata.deviceType\": \"DeviceType\"\n\n\"metadata.osName\": \"OSName\"\n\n\"metadata.countryCode\": \"CountryCode\"\n\n\"event_details.fileType\": \"FileType\"\n\nAll dimensions listed in this field must also be included in EventPattern.",
		//	                  "patternProperties": {
		//	                    "": {
		//	                      "maxLength": 255,
		//	                      "minLength": 1,
		//	                      "pattern": ".*[^\\s].*",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "EventPattern": {
		//	                  "description": "The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.\n\nWhen you define extended metrics, the metric definition is not valid if EventPattern is omitted.\n\nExample event patterns:\n\n'{ \"event_type\": [\"com.amazon.rum.js_error_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Firefox\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003c\", 2000 ] }] } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], \"countryCode\": [ \"US\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003e=\", 2000, \"\u003c\", 8000 ] }] } }'\n\nIf the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.",
		//	                  "maxLength": 4000,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Name": {
		//	                  "description": "The name for the metric that is defined in this structure. Valid values are the following:\n\nPerformanceNavigationDuration\n\nPerformanceResourceDuration\n\nNavigationSatisfiedTransaction\n\nNavigationToleratedTransaction\n\nNavigationFrustratedTransaction\n\nWebVitalsCumulativeLayoutShift\n\nWebVitalsFirstInputDelay\n\nWebVitalsLargestContentfulPaint\n\nJsErrorCount\n\nHttpErrorCount\n\nSessionCount",
		//	                  "maxLength": 255,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "UnitLabel": {
		//	                  "description": "The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.",
		//	                  "maxLength": 256,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "ValueKey": {
		//	                  "description": "The field within the event object that the metric value is sourced from.\n\nIf you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.\n\nIf this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.",
		//	                  "maxLength": 256,
		//	                  "minLength": 1,
		//	                  "pattern": ".*",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Name"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 2000,
		//	            "minItems": 0,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SessionSampleRate": {
		//	      "description": "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "Telemetries": {
		//	      "description": "An array that lists the types of telemetry data that this app monitor is to collect.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "errors",
		//	          "performance",
		//	          "http"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"app_monitor_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCookies
				"allow_cookies": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EnableXRay
				"enable_x_ray": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludedPages
				"excluded_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FavoritePages
				"favorite_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of pages in the RUM console that are to be displayed with a favorite icon.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GuestRoleArn
				"guest_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IdentityPoolId
				"identity_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the identity pool that is used to authorize the sending of data to RUM.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IncludedPages
				"included_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MetricDestinations
				"metric_destinations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DestinationArn
							"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: IamRoleArn
							"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.\n\nThis parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: MetricDefinitions
							"metric_definitions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DimensionKeys
										"dimension_keys":    // Pattern: ""
										schema.MapAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Description: "Use this field only if you are sending the metric to CloudWatch.\n\nThis field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. Valid values for the entries in this field are the following:\n\n\"metadata.pageId\": \"PageId\"\n\n\"metadata.browserName\": \"BrowserName\"\n\n\"metadata.deviceType\": \"DeviceType\"\n\n\"metadata.osName\": \"OSName\"\n\n\"metadata.countryCode\": \"CountryCode\"\n\n\"event_details.fileType\": \"FileType\"\n\nAll dimensions listed in this field must also be included in EventPattern.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: EventPattern
										"event_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.\n\nWhen you define extended metrics, the metric definition is not valid if EventPattern is omitted.\n\nExample event patterns:\n\n'{ \"event_type\": [\"com.amazon.rum.js_error_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Firefox\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"<\", 2000 ] }] } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], \"countryCode\": [ \"US\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \">=\", 2000, \"<\", 8000 ] }] } }'\n\nIf the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Name
										"name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name for the metric that is defined in this structure. Valid values are the following:\n\nPerformanceNavigationDuration\n\nPerformanceResourceDuration\n\nNavigationSatisfiedTransaction\n\nNavigationToleratedTransaction\n\nNavigationFrustratedTransaction\n\nWebVitalsCumulativeLayoutShift\n\nWebVitalsFirstInputDelay\n\nWebVitalsLargestContentfulPaint\n\nJsErrorCount\n\nHttpErrorCount\n\nSessionCount",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: UnitLabel
										"unit_label": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: ValueKey
										"value_key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The field within the event object that the metric value is sourced from.\n\nIf you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.\n\nIf this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of structures which define the metrics that you want to send.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of structures which define the destinations and the metrics that you want to send.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SessionSampleRate
				"session_sample_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Telemetries
				"telemetries": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "An array that lists the types of telemetry data that this app monitor is to collect.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "AppMonitor configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CwLogEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
		//	  "type": "boolean"
		//	}
		"cw_log_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The top-level internet domain name for which your application has administrative authority.",
		//	  "maxLength": 253,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The top-level internet domain name for which your application has administrative authority.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the app monitor",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[\\.\\-_/#A-Za-z0-9]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the app monitor",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RUM::AppMonitor",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RUM::AppMonitor").WithTerraformTypeName("awscc_rum_app_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_cookies":             "AllowCookies",
		"app_monitor_configuration": "AppMonitorConfiguration",
		"cw_log_enabled":            "CwLogEnabled",
		"destination":               "Destination",
		"destination_arn":           "DestinationArn",
		"dimension_keys":            "DimensionKeys",
		"domain":                    "Domain",
		"enable_x_ray":              "EnableXRay",
		"event_pattern":             "EventPattern",
		"excluded_pages":            "ExcludedPages",
		"favorite_pages":            "FavoritePages",
		"guest_role_arn":            "GuestRoleArn",
		"iam_role_arn":              "IamRoleArn",
		"identity_pool_id":          "IdentityPoolId",
		"included_pages":            "IncludedPages",
		"key":                       "Key",
		"metric_definitions":        "MetricDefinitions",
		"metric_destinations":       "MetricDestinations",
		"name":                      "Name",
		"session_sample_rate":       "SessionSampleRate",
		"tags":                      "Tags",
		"telemetries":               "Telemetries",
		"unit_label":                "UnitLabel",
		"value":                     "Value",
		"value_key":                 "ValueKey",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
