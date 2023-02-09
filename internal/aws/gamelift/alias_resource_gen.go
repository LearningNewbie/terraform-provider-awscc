// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package gamelift

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_gamelift_alias", aliasResource)
}

// aliasResource returns the Terraform awscc_gamelift_alias resource.
// This Terraform resource corresponds to the CloudFormation AWS::GameLift::Alias resource.
func aliasResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AliasId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique alias ID",
		//	  "type": "string"
		//	}
		"alias_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique alias ID",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A human-readable description of the alias.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A human-readable description of the alias.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A descriptive label that is associated with an alias. Alias names do not need to be unique.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A descriptive label that is associated with an alias. Alias names do not need to be unique.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile(".*\\S.*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RoutingStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "anyOf": [
		//	    {
		//	      "required": [
		//	        "FleetId"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "Message"
		//	      ]
		//	    }
		//	  ],
		//	  "description": "A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.",
		//	  "properties": {
		//	    "FleetId": {
		//	      "description": "A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.",
		//	      "pattern": "^fleet-\\S+",
		//	      "type": "string"
		//	    },
		//	    "Message": {
		//	      "description": "The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.",
		//	      "enum": [
		//	        "SIMPLE",
		//	        "TERMINAL"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"routing_strategy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FleetId
				"fleet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^fleet-\\S+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Message
				"message": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SIMPLE",
							"TERMINAL",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.",
			Required:    true,
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
		Description: "The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::Alias").WithTerraformTypeName("awscc_gamelift_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_id":         "AliasId",
		"description":      "Description",
		"fleet_id":         "FleetId",
		"message":          "Message",
		"name":             "Name",
		"routing_strategy": "RoutingStrategy",
		"type":             "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
