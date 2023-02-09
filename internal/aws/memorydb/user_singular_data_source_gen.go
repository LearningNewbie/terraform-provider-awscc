// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_memorydb_user", userDataSource)
}

// userDataSource returns the Terraform awscc_memorydb_user data source.
// This Terraform data source corresponds to the CloudFormation AWS::MemoryDB::User resource.
func userDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Access permissions string used for this user account.",
		//	  "type": "string"
		//	}
		"access_string": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Access permissions string used for this user account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the user account.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the user account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthenticationMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Passwords": {
		//	      "$comment": "List of passwords.",
		//	      "description": "Passwords used for this user account. You can create up to two passwords for each user.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Type": {
		//	      "description": "Type of authentication strategy for this user.",
		//	      "enum": [
		//	        "password"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"authentication_mode": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Passwords
				"passwords": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Passwords used for this user account. You can create up to two passwords for each user.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Type of authentication strategy for this user.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this user.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the user.",
		//	  "pattern": "[a-z][a-z0-9\\\\-]*",
		//	  "type": "string"
		//	}
		"user_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MemoryDB::User",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::User").WithTerraformTypeName("awscc_memorydb_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_string":       "AccessString",
		"arn":                 "Arn",
		"authentication_mode": "AuthenticationMode",
		"key":                 "Key",
		"passwords":           "Passwords",
		"status":              "Status",
		"tags":                "Tags",
		"type":                "Type",
		"user_name":           "UserName",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
