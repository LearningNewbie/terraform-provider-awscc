// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ssmcontacts_contact", contactResource)
}

// contactResource returns the Terraform awscc_ssmcontacts_contact resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSMContacts::Contact resource.
func contactResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_\\-\\.]*$",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9_\\-\\.]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the contact.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\s]*$",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\s]*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Plan
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.",
		//	    "properties": {
		//	      "DurationInMinutes": {
		//	        "description": "The time to wait until beginning the next stage.",
		//	        "type": "integer"
		//	      },
		//	      "Targets": {
		//	        "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	          "oneOf": [
		//	            {
		//	              "required": [
		//	                "ChannelTargetInfo"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "ContactTargetInfo"
		//	              ]
		//	            }
		//	          ],
		//	          "properties": {
		//	            "ChannelTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
		//	              "properties": {
		//	                "ChannelId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact channel.",
		//	                  "type": "string"
		//	                },
		//	                "RetryIntervalInMinutes": {
		//	                  "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "ChannelId",
		//	                "RetryIntervalInMinutes"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "ContactTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "The contact that SSM Incident Manager is engaging during an incident.",
		//	              "properties": {
		//	                "ContactId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	                  "type": "string"
		//	                },
		//	                "IsEssential": {
		//	                  "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "required": [
		//	                "ContactId",
		//	                "IsEssential"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "required": [
		//	      "DurationInMinutes"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"plan": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DurationInMinutes
					"duration_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The time to wait until beginning the next stage.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Targets
					"targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ChannelTargetInfo
								"channel_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ChannelId
										"channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact channel.",
											Required:    true,
										}, /*END ATTRIBUTE*/
										// Property: RetryIntervalInMinutes
										"retry_interval_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Description: "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
											Required:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ContactTargetInfo
								"contact_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ContactId
										"contact_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact.",
											Required:    true,
										}, /*END ATTRIBUTE*/
										// Property: IsEssential
										"is_essential": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
											Required:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The contact that SSM Incident Manager is engaging during an incident.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			Required:    true,
			// Plan is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
		//	  "enum": [
		//	    "PERSONAL",
		//	    "CUSTOM",
		//	    "SERVICE",
		//	    "ESCALATION"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"PERSONAL",
					"CUSTOM",
					"SERVICE",
					"ESCALATION",
				),
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
		Description: "Resource Type definition for AWS::SSMContacts::Contact",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::Contact").WithTerraformTypeName("awscc_ssmcontacts_contact")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                     "Alias",
		"arn":                       "Arn",
		"channel_id":                "ChannelId",
		"channel_target_info":       "ChannelTargetInfo",
		"contact_id":                "ContactId",
		"contact_target_info":       "ContactTargetInfo",
		"display_name":              "DisplayName",
		"duration_in_minutes":       "DurationInMinutes",
		"is_essential":              "IsEssential",
		"plan":                      "Plan",
		"retry_interval_in_minutes": "RetryIntervalInMinutes",
		"targets":                   "Targets",
		"type":                      "Type",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Plan",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
