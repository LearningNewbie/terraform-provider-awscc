// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package forecast

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
	registry.AddResourceFactory("awscc_forecast_dataset", datasetResource)
}

// datasetResource returns the Terraform awscc_forecast_dataset resource.
// This Terraform resource corresponds to the CloudFormation AWS::Forecast::Dataset resource.
func datasetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Frequency of data collection. This parameter is required for RELATED_TIME_SERIES",
		//	  "pattern": "^Y|M|W|D|H|30min|15min|10min|5min|1min$",
		//	  "type": "string"
		//	}
		"data_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Frequency of data collection. This parameter is required for RELATED_TIME_SERIES",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^Y|M|W|D|H|30min|15min|10min|5min|1min$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatasetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the dataset",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z][a-zA-Z0-9_]*",
		//	  "type": "string"
		//	}
		"dataset_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the dataset",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatasetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dataset type",
		//	  "enum": [
		//	    "TARGET_TIME_SERIES",
		//	    "RELATED_TIME_SERIES",
		//	    "ITEM_METADATA"
		//	  ],
		//	  "type": "string"
		//	}
		"dataset_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The dataset type",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"TARGET_TIME_SERIES",
					"RELATED_TIME_SERIES",
					"ITEM_METADATA",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain associated with the dataset",
		//	  "enum": [
		//	    "RETAIL",
		//	    "CUSTOM",
		//	    "INVENTORY_PLANNING",
		//	    "EC2_CAPACITY",
		//	    "WORK_FORCE",
		//	    "WEB_TRAFFIC",
		//	    "METRICS"
		//	  ],
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain associated with the dataset",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"RETAIL",
					"CUSTOM",
					"INVENTORY_PLANNING",
					"EC2_CAPACITY",
					"WORK_FORCE",
					"WEB_TRAFFIC",
					"METRICS",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "KmsKeyArn": {
		//	      "description": "KMS key used to encrypt the Dataset data",
		//	      "maxLength": 256,
		//	      "pattern": "arn:aws[-a-z]*:kms:.*:key/.*",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.",
		//	      "maxLength": 256,
		//	      "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"encryption_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "KMS key used to encrypt the Dataset data",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:aws[-a-z]*:kms:.*:key/.*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$"), ""),
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
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Attributes": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AttributeName": {
		//	            "description": "Name of the dataset field",
		//	            "pattern": "^[a-zA-Z][a-zA-Z0-9_]*",
		//	            "type": "string"
		//	          },
		//	          "AttributeType": {
		//	            "description": "Data type of the field",
		//	            "enum": [
		//	              "string",
		//	              "integer",
		//	              "float",
		//	              "timestamp",
		//	              "geolocation"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AttributeName
							"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Name of the dataset field",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: AttributeType
							"attribute_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Data type of the field",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"string",
										"integer",
										"float",
										"timestamp",
										"geolocation",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
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
		Description: "Resource Type Definition for AWS::Forecast::Dataset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Forecast::Dataset").WithTerraformTypeName("awscc_forecast_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"attribute_name":    "AttributeName",
		"attribute_type":    "AttributeType",
		"attributes":        "Attributes",
		"data_frequency":    "DataFrequency",
		"dataset_name":      "DatasetName",
		"dataset_type":      "DatasetType",
		"domain":            "Domain",
		"encryption_config": "EncryptionConfig",
		"key":               "Key",
		"kms_key_arn":       "KmsKeyArn",
		"role_arn":          "RoleArn",
		"schema":            "Schema",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
