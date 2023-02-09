// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_config_configuration_aggregator", configurationAggregatorDataSource)
}

// configurationAggregatorDataSource returns the Terraform awscc_config_configuration_aggregator data source.
// This Terraform data source corresponds to the CloudFormation AWS::Config::ConfigurationAggregator resource.
func configurationAggregatorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountAggregationSources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AccountIds": {
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "AllAwsRegions": {
		//	        "type": "boolean"
		//	      },
		//	      "AwsRegions": {
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "required": [
		//	      "AccountIds"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"account_aggregation_sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AccountIds
					"account_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: AllAwsRegions
					"all_aws_regions": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: AwsRegions
					"aws_regions": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationAggregatorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the aggregator.",
		//	  "type": "string"
		//	}
		"configuration_aggregator_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the aggregator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationAggregatorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the aggregator.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[\\w\\-]+",
		//	  "type": "string"
		//	}
		"configuration_aggregator_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the aggregator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OrganizationAggregationSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AllAwsRegions": {
		//	      "type": "boolean"
		//	    },
		//	    "AwsRegions": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "RoleArn": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"organization_aggregation_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllAwsRegions
				"all_aws_regions": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AwsRegions
				"aws_regions": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the configuration aggregator.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags for the configuration aggregator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Config::ConfigurationAggregator",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConfigurationAggregator").WithTerraformTypeName("awscc_config_configuration_aggregator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_aggregation_sources":     "AccountAggregationSources",
		"account_ids":                     "AccountIds",
		"all_aws_regions":                 "AllAwsRegions",
		"aws_regions":                     "AwsRegions",
		"configuration_aggregator_arn":    "ConfigurationAggregatorArn",
		"configuration_aggregator_name":   "ConfigurationAggregatorName",
		"key":                             "Key",
		"organization_aggregation_source": "OrganizationAggregationSource",
		"role_arn":                        "RoleArn",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
