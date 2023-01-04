// Code generated by generators/resource/main.go; DO NOT EDIT.

package greengrassv2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_greengrassv2_deployment", deploymentResource)
}

// deploymentResource returns the Terraform awscc_greengrassv2_deployment resource.
// This Terraform resource corresponds to the CloudFormation AWS::GreengrassV2::Deployment resource.
func deploymentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Components
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ComponentVersion": {
		//	          "maxLength": 64,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "ConfigurationUpdate": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Merge": {
		//	              "maxLength": 10485760,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Reset": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 256,
		//	                "minLength": 0,
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "RunWith": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "PosixUser": {
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "SystemResourceLimits": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Cpus": {
		//	                  "minimum": 0,
		//	                  "type": "number"
		//	                },
		//	                "Memory": {
		//	                  "format": "int64",
		//	                  "maximum": 9223372036854771712,
		//	                  "minimum": 0,
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "WindowsUser": {
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"components":              // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComponentVersion
					"component_version": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 64),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ConfigurationUpdate
					"configuration_update": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Merge
							"merge": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 10485760),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Reset
							"reset": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.ValueStringsAre(
										stringvalidator.LengthBetween(0, 256),
									),
								}, /*END VALIDATORS*/
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
					// Property: RunWith
					"run_with": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: PosixUser
							"posix_user": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtLeast(1),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SystemResourceLimits
							"system_resource_limits": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Cpus
									"cpus": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.Float64{ /*START VALIDATORS*/
											float64validator.AtLeast(0.000000),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
											float64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Memory
									"memory": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.Int64{ /*START VALIDATORS*/
											int64validator.Between(0, 9223372036854771712),
										}, /*END VALIDATORS*/
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
							// Property: WindowsUser
							"windows_user": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtLeast(1),
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
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": ".+",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"deployment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentPolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ComponentUpdatePolicy": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Action": {
		//	          "enum": [
		//	            "NOTIFY_COMPONENTS",
		//	            "SKIP_NOTIFY_COMPONENTS"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "TimeoutInSeconds": {
		//	          "maximum": 2147483647,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ConfigurationValidationPolicy": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "TimeoutInSeconds": {
		//	          "maximum": 2147483647,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FailureHandlingPolicy": {
		//	      "enum": [
		//	        "ROLLBACK",
		//	        "DO_NOTHING"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"deployment_policies": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ComponentUpdatePolicy
				"component_update_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Action
						"action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"NOTIFY_COMPONENTS",
									"SKIP_NOTIFY_COMPONENTS",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: TimeoutInSeconds
						"timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(1, 2147483647),
							}, /*END VALIDATORS*/
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
				// Property: ConfigurationValidationPolicy
				"configuration_validation_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TimeoutInSeconds
						"timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(1, 2147483647),
							}, /*END VALIDATORS*/
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
				// Property: FailureHandlingPolicy
				"failure_handling_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ROLLBACK",
							"DO_NOTHING",
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
		// Property: IotJobConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AbortConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CriteriaList": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Action": {
		//	                "enum": [
		//	                  "CANCEL"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "FailureType": {
		//	                "enum": [
		//	                  "FAILED",
		//	                  "REJECTED",
		//	                  "TIMED_OUT",
		//	                  "ALL"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "MinNumberOfExecutedThings": {
		//	                "maximum": 2147483647,
		//	                "minimum": 1,
		//	                "type": "integer"
		//	              },
		//	              "ThresholdPercentage": {
		//	                "maximum": 100,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              }
		//	            },
		//	            "required": [
		//	              "FailureType",
		//	              "Action",
		//	              "ThresholdPercentage",
		//	              "MinNumberOfExecutedThings"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "CriteriaList"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "JobExecutionsRolloutConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ExponentialRate": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BaseRatePerMinute": {
		//	              "maximum": 1000,
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "IncrementFactor": {
		//	              "maximum": 5,
		//	              "minimum": 1,
		//	              "type": "number"
		//	            },
		//	            "RateIncreaseCriteria": {
		//	              "properties": {
		//	                "NumberOfNotifiedThings": {
		//	                  "maximum": 2147483647,
		//	                  "minimum": 1,
		//	                  "type": "integer"
		//	                },
		//	                "NumberOfSucceededThings": {
		//	                  "maximum": 2147483647,
		//	                  "minimum": 1,
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "BaseRatePerMinute",
		//	            "IncrementFactor",
		//	            "RateIncreaseCriteria"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "MaximumPerMinute": {
		//	          "maximum": 1000,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "TimeoutConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "InProgressTimeoutInMinutes": {
		//	          "maximum": 2147483647,
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"iot_job_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AbortConfig
				"abort_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CriteriaList
						"criteria_list": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Action
									"action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"CANCEL",
											),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: FailureType
									"failure_type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"FAILED",
												"REJECTED",
												"TIMED_OUT",
												"ALL",
											),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: MinNumberOfExecutedThings
									"min_number_of_executed_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Required: true,
										Validators: []validator.Int64{ /*START VALIDATORS*/
											int64validator.Between(1, 2147483647),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: ThresholdPercentage
									"threshold_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Required: true,
										Validators: []validator.Float64{ /*START VALIDATORS*/
											float64validator.Between(0.000000, 100.000000),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Required: true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeAtLeast(1),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: JobExecutionsRolloutConfig
				"job_executions_rollout_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ExponentialRate
						"exponential_rate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BaseRatePerMinute
								"base_rate_per_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(1, 1000),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: IncrementFactor
								"increment_factor": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.Between(1.000000, 5.000000),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: RateIncreaseCriteria
								"rate_increase_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: NumberOfNotifiedThings
										"number_of_notified_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.Int64{ /*START VALIDATORS*/
												int64validator.Between(1, 2147483647),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: NumberOfSucceededThings
										"number_of_succeeded_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.Int64{ /*START VALIDATORS*/
												int64validator.Between(1, 2147483647),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MaximumPerMinute
						"maximum_per_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(1, 1000),
							}, /*END VALIDATORS*/
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
				// Property: TimeoutConfig
				"timeout_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: InProgressTimeoutInMinutes
						"in_progress_timeout_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(0, 2147483647),
							}, /*END VALIDATORS*/
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
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[^:]*:iot:[^:]*:[0-9]+:(thing|thinggroup)/.+",
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:[^:]*:iot:[^:]*:[0-9]+:(thing|thinggroup)/.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource for Greengrass V2 deployment.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GreengrassV2::Deployment").WithTerraformTypeName("awscc_greengrassv2_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abort_config":                    "AbortConfig",
		"action":                          "Action",
		"base_rate_per_minute":            "BaseRatePerMinute",
		"component_update_policy":         "ComponentUpdatePolicy",
		"component_version":               "ComponentVersion",
		"components":                      "Components",
		"configuration_update":            "ConfigurationUpdate",
		"configuration_validation_policy": "ConfigurationValidationPolicy",
		"cpus":                            "Cpus",
		"criteria_list":                   "CriteriaList",
		"deployment_id":                   "DeploymentId",
		"deployment_name":                 "DeploymentName",
		"deployment_policies":             "DeploymentPolicies",
		"exponential_rate":                "ExponentialRate",
		"failure_handling_policy":         "FailureHandlingPolicy",
		"failure_type":                    "FailureType",
		"in_progress_timeout_in_minutes":  "InProgressTimeoutInMinutes",
		"increment_factor":                "IncrementFactor",
		"iot_job_configuration":           "IotJobConfiguration",
		"job_executions_rollout_config":   "JobExecutionsRolloutConfig",
		"maximum_per_minute":              "MaximumPerMinute",
		"memory":                          "Memory",
		"merge":                           "Merge",
		"min_number_of_executed_things":   "MinNumberOfExecutedThings",
		"number_of_notified_things":       "NumberOfNotifiedThings",
		"number_of_succeeded_things":      "NumberOfSucceededThings",
		"posix_user":                      "PosixUser",
		"rate_increase_criteria":          "RateIncreaseCriteria",
		"reset":                           "Reset",
		"run_with":                        "RunWith",
		"system_resource_limits":          "SystemResourceLimits",
		"tags":                            "Tags",
		"target_arn":                      "TargetArn",
		"threshold_percentage":            "ThresholdPercentage",
		"timeout_config":                  "TimeoutConfig",
		"timeout_in_seconds":              "TimeoutInSeconds",
		"windows_user":                    "WindowsUser",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
