// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_autoscaling_lifecycle_hook", lifecycleHookResource)
}

// lifecycleHookResource returns the Terraform awscc_autoscaling_lifecycle_hook resource.
// This Terraform resource corresponds to the CloudFormation AWS::AutoScaling::LifecycleHook resource.
func lifecycleHookResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoScalingGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Auto Scaling group for the lifecycle hook.",
		//	  "type": "string"
		//	}
		"auto_scaling_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Auto Scaling group for the lifecycle hook.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultResult
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
		//	  "type": "string"
		//	}
		"default_result": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HeartbeatTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
		//	  "type": "integer"
		//	}
		"heartbeat_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LifecycleHookName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the lifecycle hook.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"lifecycle_hook_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the lifecycle hook.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LifecycleTransition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The instance state to which you want to attach the lifecycle hook.",
		//	  "type": "string"
		//	}
		"lifecycle_transition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The instance state to which you want to attach the lifecycle hook.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotificationMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
		//	  "maxLength": 1023,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"notification_metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1023),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotificationTargetARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
		//	  "type": "string"
		//	}
		"notification_target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
			Optional:    true,
			Computed:    true,
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
		Description: "Resource Type definition for AWS::AutoScaling::LifecycleHook",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AutoScaling::LifecycleHook").WithTerraformTypeName("awscc_autoscaling_lifecycle_hook")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_name": "AutoScalingGroupName",
		"default_result":          "DefaultResult",
		"heartbeat_timeout":       "HeartbeatTimeout",
		"lifecycle_hook_name":     "LifecycleHookName",
		"lifecycle_transition":    "LifecycleTransition",
		"notification_metadata":   "NotificationMetadata",
		"notification_target_arn": "NotificationTargetARN",
		"role_arn":                "RoleARN",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
