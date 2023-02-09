// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_redshift_scheduled_action", scheduledActionResource)
}

// scheduledActionResource returns the Terraform awscc_redshift_scheduled_action resource.
// This Terraform resource corresponds to the CloudFormation AWS::Redshift::ScheduledAction resource.
func scheduledActionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Enable
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If true, the schedule is enabled. If false, the scheduled action does not trigger.",
		//	  "type": "boolean"
		//	}
		"enable": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If true, the schedule is enabled. If false, the scheduled action does not trigger.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.",
		//	  "type": "string"
		//	}
		"end_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IamRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM role to assume to run the target action.",
		//	  "type": "string"
		//	}
		"iam_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM role to assume to run the target action.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NextInvocations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of times when the scheduled action will run.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"next_invocations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of times when the scheduled action will run.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The schedule in `at( )` or `cron( )` format.",
		//	  "type": "string"
		//	}
		"schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The schedule in `at( )` or `cron( )` format.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScheduledActionDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the scheduled action.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"scheduled_action_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the scheduled action.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScheduledActionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the scheduled action. The name must be unique within an account.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"scheduled_action_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the scheduled action. The name must be unique within an account.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StartTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.",
		//	  "type": "string"
		//	}
		"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the scheduled action.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the scheduled action.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetAction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A JSON format string of the Amazon Redshift API operation with input parameters.",
		//	  "properties": {
		//	    "PauseCluster": {
		//	      "additionalProperties": false,
		//	      "description": "Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.",
		//	      "properties": {
		//	        "ClusterIdentifier": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ClusterIdentifier"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ResizeCluster": {
		//	      "additionalProperties": false,
		//	      "description": "Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.",
		//	      "properties": {
		//	        "Classic": {
		//	          "type": "boolean"
		//	        },
		//	        "ClusterIdentifier": {
		//	          "type": "string"
		//	        },
		//	        "ClusterType": {
		//	          "type": "string"
		//	        },
		//	        "NodeType": {
		//	          "type": "string"
		//	        },
		//	        "NumberOfNodes": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "ClusterIdentifier"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ResumeCluster": {
		//	      "additionalProperties": false,
		//	      "description": "Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.",
		//	      "properties": {
		//	        "ClusterIdentifier": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ClusterIdentifier"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"target_action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PauseCluster
				"pause_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ClusterIdentifier
						"cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResizeCluster
				"resize_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Classic
						"classic": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ClusterIdentifier
						"cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: ClusterType
						"cluster_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NodeType
						"node_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NumberOfNodes
						"number_of_nodes": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResumeCluster
				"resume_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ClusterIdentifier
						"cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A JSON format string of the Amazon Redshift API operation with input parameters.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::ScheduledAction").WithTerraformTypeName("awscc_redshift_scheduled_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"classic":                      "Classic",
		"cluster_identifier":           "ClusterIdentifier",
		"cluster_type":                 "ClusterType",
		"enable":                       "Enable",
		"end_time":                     "EndTime",
		"iam_role":                     "IamRole",
		"next_invocations":             "NextInvocations",
		"node_type":                    "NodeType",
		"number_of_nodes":              "NumberOfNodes",
		"pause_cluster":                "PauseCluster",
		"resize_cluster":               "ResizeCluster",
		"resume_cluster":               "ResumeCluster",
		"schedule":                     "Schedule",
		"scheduled_action_description": "ScheduledActionDescription",
		"scheduled_action_name":        "ScheduledActionName",
		"start_time":                   "StartTime",
		"state":                        "State",
		"target_action":                "TargetAction",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
