// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotevents_input", inputDataSource)
}

// inputDataSource returns the Terraform awscc_iotevents_input data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTEvents::Input resource.
func inputDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: InputDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The definition of the input.",
		//	  "properties": {
		//	    "Attributes": {
		//	      "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
		//	        "properties": {
		//	          "JsonPath": {
		//	            "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "^((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+))(\\.((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+)))*$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "JsonPath"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 200,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "Attributes"
		//	  ],
		//	  "type": "object"
		//	}
		"input_definition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: JsonPath
							"json_path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `<field-name>.<field-name>...`",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The definition of the input.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InputDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A brief description of the input.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"input_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A brief description of the input.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InputName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the input.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z][a-zA-Z0-9_]*$",
		//	  "type": "string"
		//	}
		"input_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the input.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Tags to be applied to Input.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Key of the Tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Value of the Tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Key of the Tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Value of the Tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTEvents::Input",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTEvents::Input").WithTerraformTypeName("awscc_iotevents_input")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":        "Attributes",
		"input_definition":  "InputDefinition",
		"input_description": "InputDescription",
		"input_name":        "InputName",
		"json_path":         "JsonPath",
		"key":               "Key",
		"tags":              "Tags",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
