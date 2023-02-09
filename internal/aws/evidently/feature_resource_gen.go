// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package evidently

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_evidently_feature", featureResource)
}

// featureResource returns the Terraform awscc_evidently_feature resource.
// This Terraform resource corresponds to the CloudFormation AWS::Evidently::Feature resource.
func featureResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/feature/[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultVariation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"default_variation": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 127),
				stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 160,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 160),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EntityOverrides
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "EntityId": {
		//	        "type": "string"
		//	      },
		//	      "Variation": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "[-a-zA-Z0-9._]*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"entity_overrides": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EntityId
					"entity_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Variation
					"variation": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
							stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 20),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EvaluationStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ALL_RULES",
		//	    "DEFAULT_VARIATION"
		//	  ],
		//	  "type": "string"
		//	}
		"evaluation_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ALL_RULES",
					"DEFAULT_VARIATION",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 127),
				stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Project
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
		//	  "type": "string"
		//	}
		"project": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Variations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "oneOf": [
		//	      {
		//	        "required": [
		//	          "VariationName",
		//	          "StringValue"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "VariationName",
		//	          "BooleanValue"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "VariationName",
		//	          "LongValue"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "VariationName",
		//	          "DoubleValue"
		//	        ]
		//	      }
		//	    ],
		//	    "properties": {
		//	      "BooleanValue": {
		//	        "type": "boolean"
		//	      },
		//	      "DoubleValue": {
		//	        "type": "number"
		//	      },
		//	      "LongValue": {
		//	        "type": "number"
		//	      },
		//	      "StringValue": {
		//	        "maxLength": 512,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "VariationName": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "[-a-zA-Z0-9._]*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"variations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BooleanValue
					"boolean_value": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DoubleValue
					"double_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
							float64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LongValue
					"long_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
							float64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: StringValue
					"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 512),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: VariationName
					"variation_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
							stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 5),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
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
		Description: "Resource Type definition for AWS::Evidently::Feature.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Feature").WithTerraformTypeName("awscc_evidently_feature")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"boolean_value":       "BooleanValue",
		"default_variation":   "DefaultVariation",
		"description":         "Description",
		"double_value":        "DoubleValue",
		"entity_id":           "EntityId",
		"entity_overrides":    "EntityOverrides",
		"evaluation_strategy": "EvaluationStrategy",
		"key":                 "Key",
		"long_value":          "LongValue",
		"name":                "Name",
		"project":             "Project",
		"string_value":        "StringValue",
		"tags":                "Tags",
		"value":               "Value",
		"variation":           "Variation",
		"variation_name":      "VariationName",
		"variations":          "Variations",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
