// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_networkmanager_core_network", coreNetworkResource)
}

// coreNetworkResource returns the Terraform awscc_networkmanager_core_network resource.
// This Terraform resource corresponds to the CloudFormation AWS::NetworkManager::CoreNetwork resource.
func coreNetworkResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CoreNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon resource name) of core network",
		//	  "type": "string"
		//	}
		"core_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon resource name) of core network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of core network",
		//	  "type": "string"
		//	}
		"core_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of core network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The creation time of core network",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The creation time of core network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of core network",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of core network",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Edges
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The edges within a core network.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Asn": {
		//	        "description": "The ASN of a core network edge.",
		//	        "type": "number"
		//	      },
		//	      "EdgeLocation": {
		//	        "description": "The Region where a core network edge is located.",
		//	        "type": "string"
		//	      },
		//	      "InsideCidrBlocks": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "description": "The inside IP addresses used for core network edges.",
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"edges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Asn
					"asn": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "The ASN of a core network edge.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: EdgeLocation
					"edge_location": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Region where a core network edge is located.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: InsideCidrBlocks
					"inside_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The edges within a core network.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GlobalNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the global network that your core network is a part of.",
		//	  "type": "string"
		//	}
		"global_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the global network that your core network is a part of.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OwnerAccount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Owner of the core network",
		//	  "type": "string"
		//	}
		"owner_account": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Owner of the core network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Live policy document for the core network, you must provide PolicyDocument in Json Format",
		//	  "type": "object"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  generic.JSONStringType,
			Description: "Live policy document for the core network, you must provide PolicyDocument in Json Format",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.JSONStringType.AttributePlanModifier(),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Segments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The segments within a core network.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "EdgeLocations": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "description": "The Regions where the edges are located.",
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "Name": {
		//	        "description": "Name of segment",
		//	        "type": "string"
		//	      },
		//	      "SharedSegments": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "description": "The shared segments of a core network.",
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"segments": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EdgeLocations
					"edge_locations": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Name of segment",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SharedSegments
					"shared_segments": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The segments within a core network.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of core network",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of core network",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the global network.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags for the global network.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "AWS::NetworkManager::CoreNetwork Resource Type Definition.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::CoreNetwork").WithTerraformTypeName("awscc_networkmanager_core_network")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asn":                "Asn",
		"core_network_arn":   "CoreNetworkArn",
		"core_network_id":    "CoreNetworkId",
		"created_at":         "CreatedAt",
		"description":        "Description",
		"edge_location":      "EdgeLocation",
		"edge_locations":     "EdgeLocations",
		"edges":              "Edges",
		"global_network_id":  "GlobalNetworkId",
		"inside_cidr_blocks": "InsideCidrBlocks",
		"key":                "Key",
		"name":               "Name",
		"owner_account":      "OwnerAccount",
		"policy_document":    "PolicyDocument",
		"segments":           "Segments",
		"shared_segments":    "SharedSegments",
		"state":              "State",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(720).WithDeleteTimeoutInMinutes(720)

	opts = opts.WithUpdateTimeoutInMinutes(720)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
