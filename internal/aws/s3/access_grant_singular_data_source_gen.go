// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_access_grant", accessGrantDataSource)
}

// accessGrantDataSource returns the Terraform awscc_s3_access_grant data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::AccessGrant resource.
func accessGrantDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessGrantArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified access grant.",
		//	  "type": "string"
		//	}
		"access_grant_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified access grant.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID assigned to this access grant.",
		//	  "examples": [
		//	    "7c89cbd1-0f4e-40e3-861d-afb906952b77"
		//	  ],
		//	  "type": "string"
		//	}
		"access_grant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID assigned to this access grant.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantsLocationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options of the grant location, which is the S3 path to the data to which you are granting access.",
		//	  "properties": {
		//	    "S3SubPrefix": {
		//	      "description": "The S3 sub prefix of a registered location in your S3 Access Grants instance",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3SubPrefix"
		//	  ],
		//	  "type": "object"
		//	}
		"access_grants_location_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3SubPrefix
				"s3_sub_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The S3 sub prefix of a registered location in your S3 Access Grants instance",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options of the grant location, which is the S3 path to the data to which you are granting access.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessGrantsLocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The custom S3 location to be accessed by the grantee",
		//	  "examples": [
		//	    "125f332b-a499-4eb6-806f-8a6a1aa4cb96"
		//	  ],
		//	  "type": "string"
		//	}
		"access_grants_location_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The custom S3 location to be accessed by the grantee",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ApplicationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the application grantees will use to access the location",
		//	  "type": "string"
		//	}
		"application_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the application grantees will use to access the location",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GrantScope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.",
		//	  "type": "string"
		//	}
		"grant_scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Grantee
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The principal who will be granted permission to access S3.",
		//	  "properties": {
		//	    "GranteeIdentifier": {
		//	      "description": "The unique identifier of the Grantee",
		//	      "type": "string"
		//	    },
		//	    "GranteeType": {
		//	      "description": "Configures the transfer acceleration state for an Amazon S3 bucket.",
		//	      "enum": [
		//	        "IAM",
		//	        "DIRECTORY_USER",
		//	        "DIRECTORY_GROUP"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "GranteeType",
		//	    "GranteeIdentifier"
		//	  ],
		//	  "type": "object"
		//	}
		"grantee": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GranteeIdentifier
				"grantee_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The unique identifier of the Grantee",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GranteeType
				"grantee_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Configures the transfer acceleration state for an Amazon S3 bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The principal who will be granted permission to access S3.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Permission
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The level of access to be afforded to the grantee",
		//	  "enum": [
		//	    "READ",
		//	    "WRITE",
		//	    "READWRITE"
		//	  ],
		//	  "type": "string"
		//	}
		"permission": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The level of access to be afforded to the grantee",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3PrefixType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of S3SubPrefix.",
		//	  "enum": [
		//	    "Object"
		//	  ],
		//	  "type": "string"
		//	}
		"s3_prefix_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of S3SubPrefix.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3::AccessGrant",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::AccessGrant").WithTerraformTypeName("awscc_s3_access_grant")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_grant_arn":                     "AccessGrantArn",
		"access_grant_id":                      "AccessGrantId",
		"access_grants_location_configuration": "AccessGrantsLocationConfiguration",
		"access_grants_location_id":            "AccessGrantsLocationId",
		"application_arn":                      "ApplicationArn",
		"grant_scope":                          "GrantScope",
		"grantee":                              "Grantee",
		"grantee_identifier":                   "GranteeIdentifier",
		"grantee_type":                         "GranteeType",
		"key":                                  "Key",
		"permission":                           "Permission",
		"s3_prefix_type":                       "S3PrefixType",
		"s3_sub_prefix":                        "S3SubPrefix",
		"tags":                                 "Tags",
		"value":                                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
