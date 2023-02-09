// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_frauddetector_variable", variableResource)
}

// variableResource returns the Terraform awscc_frauddetector_variable resource.
// This Terraform resource corresponds to the CloudFormation AWS::FraudDetector::Variable resource.
func variableResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the variable.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the variable.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the variable was created.",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the variable was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source of the data.",
		//	  "enum": [
		//	    "EVENT",
		//	    "EXTERNAL_MODEL_SCORE"
		//	  ],
		//	  "type": "string"
		//	}
		"data_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source of the data.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"EVENT",
					"EXTERNAL_MODEL_SCORE",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DataType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The data type.",
		//	  "enum": [
		//	    "STRING",
		//	    "INTEGER",
		//	    "FLOAT",
		//	    "BOOLEAN"
		//	  ],
		//	  "type": "string"
		//	}
		"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The data type.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"STRING",
					"INTEGER",
					"FLOAT",
					"BOOLEAN",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultValue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default value for the variable when no value is received.",
		//	  "type": "string"
		//	}
		"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default value for the variable when no value is received.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the variable was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the variable was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the variable.",
		//	  "pattern": "^[a-z_][a-z0-9_]{0,99}?$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the variable.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z_][a-z0-9_]{0,99}?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with this variable.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags associated with this variable.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VariableType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
		//	  "enum": [
		//	    "AUTH_CODE",
		//	    "AVS",
		//	    "BILLING_ADDRESS_L1",
		//	    "BILLING_ADDRESS_L2",
		//	    "BILLING_CITY",
		//	    "BILLING_COUNTRY",
		//	    "BILLING_NAME",
		//	    "BILLING_PHONE",
		//	    "BILLING_STATE",
		//	    "BILLING_ZIP",
		//	    "CARD_BIN",
		//	    "CATEGORICAL",
		//	    "CURRENCY_CODE",
		//	    "EMAIL_ADDRESS",
		//	    "FINGERPRINT",
		//	    "FRAUD_LABEL",
		//	    "FREE_FORM_TEXT",
		//	    "IP_ADDRESS",
		//	    "NUMERIC",
		//	    "ORDER_ID",
		//	    "PAYMENT_TYPE",
		//	    "PHONE_NUMBER",
		//	    "PRICE",
		//	    "PRODUCT_CATEGORY",
		//	    "SHIPPING_ADDRESS_L1",
		//	    "SHIPPING_ADDRESS_L2",
		//	    "SHIPPING_CITY",
		//	    "SHIPPING_COUNTRY",
		//	    "SHIPPING_NAME",
		//	    "SHIPPING_PHONE",
		//	    "SHIPPING_STATE",
		//	    "SHIPPING_ZIP",
		//	    "USERAGENT"
		//	  ],
		//	  "type": "string"
		//	}
		"variable_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"AUTH_CODE",
					"AVS",
					"BILLING_ADDRESS_L1",
					"BILLING_ADDRESS_L2",
					"BILLING_CITY",
					"BILLING_COUNTRY",
					"BILLING_NAME",
					"BILLING_PHONE",
					"BILLING_STATE",
					"BILLING_ZIP",
					"CARD_BIN",
					"CATEGORICAL",
					"CURRENCY_CODE",
					"EMAIL_ADDRESS",
					"FINGERPRINT",
					"FRAUD_LABEL",
					"FREE_FORM_TEXT",
					"IP_ADDRESS",
					"NUMERIC",
					"ORDER_ID",
					"PAYMENT_TYPE",
					"PHONE_NUMBER",
					"PRICE",
					"PRODUCT_CATEGORY",
					"SHIPPING_ADDRESS_L1",
					"SHIPPING_ADDRESS_L2",
					"SHIPPING_CITY",
					"SHIPPING_COUNTRY",
					"SHIPPING_NAME",
					"SHIPPING_PHONE",
					"SHIPPING_STATE",
					"SHIPPING_ZIP",
					"USERAGENT",
				),
			}, /*END VALIDATORS*/
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
		Description: "A resource schema for a Variable in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Variable").WithTerraformTypeName("awscc_frauddetector_variable")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"data_source":       "DataSource",
		"data_type":         "DataType",
		"default_value":     "DefaultValue",
		"description":       "Description",
		"key":               "Key",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
		"variable_type":     "VariableType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
