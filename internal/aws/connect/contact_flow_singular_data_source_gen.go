// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_contact_flow", contactFlowDataSource)
}

// contactFlowDataSource returns the Terraform awscc_connect_contact_flow data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::ContactFlow resource.
func contactFlowDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContactFlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the contact flow (ARN).",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the contact flow (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The content of the contact flow in JSON format.",
		//	  "maxLength": 256000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"content": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The content of the contact flow in JSON format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the contact flow.",
		//	  "maxLength": 500,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the contact flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance (ARN).",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance (ARN).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the contact flow.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the contact flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the contact flow.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "ARCHIVED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the contact flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the contact flow.",
		//	  "enum": [
		//	    "CONTACT_FLOW",
		//	    "CUSTOMER_QUEUE",
		//	    "CUSTOMER_HOLD",
		//	    "CUSTOMER_WHISPER",
		//	    "AGENT_HOLD",
		//	    "AGENT_WHISPER",
		//	    "OUTBOUND_WHISPER",
		//	    "AGENT_TRANSFER",
		//	    "QUEUE_TRANSFER"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the contact flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::ContactFlow",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::ContactFlow").WithTerraformTypeName("awscc_connect_contact_flow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_flow_arn": "ContactFlowArn",
		"content":          "Content",
		"description":      "Description",
		"instance_arn":     "InstanceArn",
		"key":              "Key",
		"name":             "Name",
		"state":            "State",
		"tags":             "Tags",
		"type":             "Type",
		"value":            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
