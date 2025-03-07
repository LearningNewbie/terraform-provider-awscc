// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package groundstation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_groundstation_dataflow_endpoint_group", dataflowEndpointGroupResource)
}

// dataflowEndpointGroupResource returns the Terraform awscc_groundstation_dataflow_endpoint_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::GroundStation::DataflowEndpointGroup resource.
func dataflowEndpointGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContactPostPassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.",
		//	  "type": "integer"
		//	}
		"contact_post_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContactPrePassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.",
		//	  "type": "integer"
		//	}
		"contact_pre_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "oneOf": [
		//	      {
		//	        "required": [
		//	          "Endpoint",
		//	          "SecurityDetails"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "AwsGroundStationAgentEndpoint"
		//	        ]
		//	      }
		//	    ],
		//	    "properties": {
		//	      "AwsGroundStationAgentEndpoint": {
		//	        "additionalProperties": false,
		//	        "description": "Information about AwsGroundStationAgentEndpoint.",
		//	        "properties": {
		//	          "AgentStatus": {
		//	            "description": "The status of AgentEndpoint.",
		//	            "enum": [
		//	              "SUCCESS",
		//	              "FAILED",
		//	              "ACTIVE",
		//	              "INACTIVE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "AuditResults": {
		//	            "description": "The results of the audit.",
		//	            "enum": [
		//	              "HEALTHY",
		//	              "UNHEALTHY"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "EgressAddress": {
		//	            "additionalProperties": false,
		//	            "description": "Egress address of AgentEndpoint with an optional mtu.",
		//	            "properties": {
		//	              "Mtu": {
		//	                "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
		//	                "type": "integer"
		//	              },
		//	              "SocketAddress": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Name": {
		//	                    "type": "string"
		//	                  },
		//	                  "Port": {
		//	                    "type": "integer"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "IngressAddress": {
		//	            "additionalProperties": false,
		//	            "description": "Ingress address of AgentEndpoint with a port range and an optional mtu.",
		//	            "properties": {
		//	              "Mtu": {
		//	                "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
		//	                "type": "integer"
		//	              },
		//	              "SocketAddress": {
		//	                "additionalProperties": false,
		//	                "description": "A socket address with a port range.",
		//	                "properties": {
		//	                  "Name": {
		//	                    "description": "IPv4 socket address.",
		//	                    "type": "string"
		//	                  },
		//	                  "PortRange": {
		//	                    "additionalProperties": false,
		//	                    "description": "Port range of a socket address.",
		//	                    "properties": {
		//	                      "Maximum": {
		//	                        "description": "A maximum value.",
		//	                        "type": "integer"
		//	                      },
		//	                      "Minimum": {
		//	                        "description": "A minimum value.",
		//	                        "type": "integer"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Endpoint": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Address": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Name": {
		//	                "type": "string"
		//	              },
		//	              "Port": {
		//	                "type": "integer"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Mtu": {
		//	            "type": "integer"
		//	          },
		//	          "Name": {
		//	            "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "SecurityDetails": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "RoleArn": {
		//	            "pattern": "^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$",
		//	            "type": "string"
		//	          },
		//	          "SecurityGroupIds": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "SubnetIds": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"endpoint_details": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AwsGroundStationAgentEndpoint
					"aws_ground_station_agent_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AgentStatus
							"agent_status": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The status of AgentEndpoint.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"SUCCESS",
										"FAILED",
										"ACTIVE",
										"INACTIVE",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: AuditResults
							"audit_results": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The results of the audit.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"HEALTHY",
										"UNHEALTHY",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: EgressAddress
							"egress_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Mtu
									"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: SocketAddress
									"socket_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Port
											"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
													int64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Egress address of AgentEndpoint with an optional mtu.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: IngressAddress
							"ingress_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Mtu
									"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: SocketAddress
									"socket_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "IPv4 socket address.",
												Optional:    true,
												Computed:    true,
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: PortRange
											"port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Maximum
													"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Description: "A maximum value.",
														Optional:    true,
														Computed:    true,
														PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
															int64planmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: Minimum
													"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Description: "A minimum value.",
														Optional:    true,
														Computed:    true,
														PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
															int64planmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
												Description: "Port range of a socket address.",
												Optional:    true,
												Computed:    true,
												PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
													objectplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "A socket address with a port range.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Ingress address of AgentEndpoint with a port range and an optional mtu.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z0-9_:-]{1,256}$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Information about AwsGroundStationAgentEndpoint.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Endpoint
					"endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Address
							"address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Port
									"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Mtu
							"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z0-9_:-]{1,256}$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SecurityDetails
					"security_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: RoleArn
							"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SecurityGroupIds
							"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SubnetIds
							"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dataflow_endpoint_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
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
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "AWS Ground Station DataflowEndpointGroup schema for CloudFormation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::DataflowEndpointGroup").WithTerraformTypeName("awscc_groundstation_dataflow_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                            "Address",
		"agent_status":                       "AgentStatus",
		"arn":                                "Arn",
		"audit_results":                      "AuditResults",
		"aws_ground_station_agent_endpoint":  "AwsGroundStationAgentEndpoint",
		"contact_post_pass_duration_seconds": "ContactPostPassDurationSeconds",
		"contact_pre_pass_duration_seconds":  "ContactPrePassDurationSeconds",
		"dataflow_endpoint_group_id":         "Id",
		"egress_address":                     "EgressAddress",
		"endpoint":                           "Endpoint",
		"endpoint_details":                   "EndpointDetails",
		"ingress_address":                    "IngressAddress",
		"key":                                "Key",
		"maximum":                            "Maximum",
		"minimum":                            "Minimum",
		"mtu":                                "Mtu",
		"name":                               "Name",
		"port":                               "Port",
		"port_range":                         "PortRange",
		"role_arn":                           "RoleArn",
		"security_details":                   "SecurityDetails",
		"security_group_ids":                 "SecurityGroupIds",
		"socket_address":                     "SocketAddress",
		"subnet_ids":                         "SubnetIds",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
