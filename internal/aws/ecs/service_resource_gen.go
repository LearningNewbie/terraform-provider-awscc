// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ecs_service", serviceResource)
}

// serviceResource returns the Terraform awscc_ecs_service resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECS::Service resource.
func serviceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CapacityProviderStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Base": {
		//	        "type": "integer"
		//	      },
		//	      "CapacityProvider": {
		//	        "type": "string"
		//	      },
		//	      "Weight": {
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"capacity_provider_strategy": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Base
					"base": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CapacityProvider
					"capacity_provider": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Weight
					"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
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
		// Property: Cluster
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"cluster": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Alarms": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AlarmNames": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Enable": {
		//	          "type": "boolean"
		//	        },
		//	        "Rollback": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "AlarmNames",
		//	        "Rollback",
		//	        "Enable"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DeploymentCircuitBreaker": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Enable": {
		//	          "type": "boolean"
		//	        },
		//	        "Rollback": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enable",
		//	        "Rollback"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "MaximumPercent": {
		//	      "type": "integer"
		//	    },
		//	    "MinimumHealthyPercent": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"deployment_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Alarms
				"alarms": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AlarmNames
						"alarm_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: Enable
						"enable": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: Rollback
						"rollback": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DeploymentCircuitBreaker
				"deployment_circuit_breaker": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enable
						"enable": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: Rollback
						"rollback": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaximumPercent
				"maximum_percent": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MinimumHealthyPercent
				"minimum_healthy_percent": schema.Int64Attribute{ /*START ATTRIBUTE*/
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
		// Property: DeploymentController
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Type": {
		//	      "enum": [
		//	        "CODE_DEPLOY",
		//	        "ECS",
		//	        "EXTERNAL"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"deployment_controller": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"CODE_DEPLOY",
							"ECS",
							"EXTERNAL",
						),
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
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesiredCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"desired_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableECSManagedTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_ecs_managed_tags": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableExecuteCommand
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_execute_command": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckGracePeriodSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"health_check_grace_period_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LaunchType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "EC2",
		//	    "FARGATE",
		//	    "EXTERNAL"
		//	  ],
		//	  "type": "string"
		//	}
		"launch_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"EC2",
					"FARGATE",
					"EXTERNAL",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ContainerName": {
		//	        "type": "string"
		//	      },
		//	      "ContainerPort": {
		//	        "type": "integer"
		//	      },
		//	      "LoadBalancerName": {
		//	        "type": "string"
		//	      },
		//	      "TargetGroupArn": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"load_balancers": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ContainerName
					"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ContainerPort
					"container_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LoadBalancerName
					"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TargetGroupArn
					"target_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
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
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AwsvpcConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AssignPublicIp": {
		//	          "enum": [
		//	            "DISABLED",
		//	            "ENABLED"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "SecurityGroups": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Subnets": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AwsvpcConfiguration
				"awsvpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AssignPublicIp
						"assign_public_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DISABLED",
									"ENABLED",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SecurityGroups
						"security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Subnets
						"subnets": schema.ListAttribute{ /*START ATTRIBUTE*/
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlacementConstraints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Expression": {
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "distinctInstance",
		//	          "memberOf"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"placement_constraints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Expression
					"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"distinctInstance",
								"memberOf",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlacementStrategies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Field": {
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "binpack",
		//	          "random",
		//	          "spread"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"placement_strategies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Field
					"field": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"binpack",
								"random",
								"spread",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlatformVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "LATEST",
		//	  "type": "string"
		//	}
		"platform_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("LATEST"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PropagateTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "SERVICE",
		//	    "TASK_DEFINITION"
		//	  ],
		//	  "type": "string"
		//	}
		"propagate_tags": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SERVICE",
					"TASK_DEFINITION",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Role
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchedulingStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DAEMON",
		//	    "REPLICA"
		//	  ],
		//	  "type": "string"
		//	}
		"scheduling_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DAEMON",
					"REPLICA",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceConnectConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Enabled": {
		//	      "type": "boolean"
		//	    },
		//	    "LogConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "LogDriver": {
		//	          "type": "string"
		//	        },
		//	        "Options": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "SecretOptions": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Name": {
		//	                "type": "string"
		//	              },
		//	              "ValueFrom": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Name",
		//	              "ValueFrom"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Namespace": {
		//	      "type": "string"
		//	    },
		//	    "Services": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ClientAliases": {
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "DnsName": {
		//	                  "type": "string"
		//	                },
		//	                "Port": {
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "Port"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "DiscoveryName": {
		//	            "type": "string"
		//	          },
		//	          "IngressPortOverride": {
		//	            "type": "integer"
		//	          },
		//	          "PortName": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "PortName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "Enabled"
		//	  ],
		//	  "type": "object"
		//	}
		"service_connect_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: LogConfiguration
				"log_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogDriver
						"log_driver": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Options
						"options":           // Pattern: ""
						schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SecretOptions
						"secret_options": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
									// Property: ValueFrom
									"value_from": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
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
				// Property: Namespace
				"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Services
				"services": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ClientAliases
							"client_aliases": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DnsName
										"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Port
										"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Required: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: DiscoveryName
							"discovery_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: IngressPortOverride
							"ingress_port_override": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: PortName
							"port_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
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
			// ServiceConnectConfiguration is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceRegistries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ContainerName": {
		//	        "type": "string"
		//	      },
		//	      "ContainerPort": {
		//	        "type": "integer"
		//	      },
		//	      "Port": {
		//	        "type": "integer"
		//	      },
		//	      "RegistryArn": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"service_registries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ContainerName
					"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ContainerPort
					"container_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
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
					// Property: RegistryArn
					"registry_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
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
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
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
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
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
		// Property: TaskDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"task_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::ECS::Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::Service").WithTerraformTypeName("awscc_ecs_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_names":                       "AlarmNames",
		"alarms":                            "Alarms",
		"assign_public_ip":                  "AssignPublicIp",
		"awsvpc_configuration":              "AwsvpcConfiguration",
		"base":                              "Base",
		"capacity_provider":                 "CapacityProvider",
		"capacity_provider_strategy":        "CapacityProviderStrategy",
		"client_aliases":                    "ClientAliases",
		"cluster":                           "Cluster",
		"container_name":                    "ContainerName",
		"container_port":                    "ContainerPort",
		"deployment_circuit_breaker":        "DeploymentCircuitBreaker",
		"deployment_configuration":          "DeploymentConfiguration",
		"deployment_controller":             "DeploymentController",
		"desired_count":                     "DesiredCount",
		"discovery_name":                    "DiscoveryName",
		"dns_name":                          "DnsName",
		"enable":                            "Enable",
		"enable_ecs_managed_tags":           "EnableECSManagedTags",
		"enable_execute_command":            "EnableExecuteCommand",
		"enabled":                           "Enabled",
		"expression":                        "Expression",
		"field":                             "Field",
		"health_check_grace_period_seconds": "HealthCheckGracePeriodSeconds",
		"ingress_port_override":             "IngressPortOverride",
		"key":                               "Key",
		"launch_type":                       "LaunchType",
		"load_balancer_name":                "LoadBalancerName",
		"load_balancers":                    "LoadBalancers",
		"log_configuration":                 "LogConfiguration",
		"log_driver":                        "LogDriver",
		"maximum_percent":                   "MaximumPercent",
		"minimum_healthy_percent":           "MinimumHealthyPercent",
		"name":                              "Name",
		"namespace":                         "Namespace",
		"network_configuration":             "NetworkConfiguration",
		"options":                           "Options",
		"placement_constraints":             "PlacementConstraints",
		"placement_strategies":              "PlacementStrategies",
		"platform_version":                  "PlatformVersion",
		"port":                              "Port",
		"port_name":                         "PortName",
		"propagate_tags":                    "PropagateTags",
		"registry_arn":                      "RegistryArn",
		"role":                              "Role",
		"rollback":                          "Rollback",
		"scheduling_strategy":               "SchedulingStrategy",
		"secret_options":                    "SecretOptions",
		"security_groups":                   "SecurityGroups",
		"service_arn":                       "ServiceArn",
		"service_connect_configuration":     "ServiceConnectConfiguration",
		"service_name":                      "ServiceName",
		"service_registries":                "ServiceRegistries",
		"services":                          "Services",
		"subnets":                           "Subnets",
		"tags":                              "Tags",
		"target_group_arn":                  "TargetGroupArn",
		"task_definition":                   "TaskDefinition",
		"type":                              "Type",
		"value":                             "Value",
		"value_from":                        "ValueFrom",
		"weight":                            "Weight",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServiceConnectConfiguration",
	})
	opts = opts.WithCreateTimeoutInMinutes(180).WithDeleteTimeoutInMinutes(30)

	opts = opts.WithUpdateTimeoutInMinutes(180)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
