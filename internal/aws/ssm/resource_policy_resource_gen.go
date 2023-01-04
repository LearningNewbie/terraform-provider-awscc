// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ssm_resource_policy", resourcePolicyResource)
}

// resourcePolicyResource returns the Terraform awscc_ssm_resource_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSM::ResourcePolicy resource.
func resourcePolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Actual policy statement.",
		//	  "type": "string"
		//	}
		"policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Actual policy statement.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyHash
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A snapshot identifier for the policy over time.",
		//	  "type": "string"
		//	}
		"policy_hash": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A snapshot identifier for the policy over time.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An unique identifier within the policies of a resource. ",
		//	  "type": "string"
		//	}
		"policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An unique identifier within the policies of a resource. ",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of OpsItemGroup etc.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of OpsItemGroup etc.",
			Required:    true,
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
		Description: "Resource Type definition for AWS::SSM::ResourcePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourcePolicy").WithTerraformTypeName("awscc_ssm_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"policy":       "Policy",
		"policy_hash":  "PolicyHash",
		"policy_id":    "PolicyId",
		"resource_arn": "ResourceArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
