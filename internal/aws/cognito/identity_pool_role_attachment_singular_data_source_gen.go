// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cognito

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cognito_identity_pool_role_attachment", identityPoolRoleAttachmentDataSource)
}

// identityPoolRoleAttachmentDataSource returns the Terraform awscc_cognito_identity_pool_role_attachment data source.
// This Terraform data source corresponds to the CloudFormation AWS::Cognito::IdentityPoolRoleAttachment resource.
func identityPoolRoleAttachmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"identity_pool_role_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"identity_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RoleMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AmbiguousRoleResolution": {
		//	          "type": "string"
		//	        },
		//	        "IdentityProvider": {
		//	          "type": "string"
		//	        },
		//	        "RulesConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Rules": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Claim": {
		//	                    "type": "string"
		//	                  },
		//	                  "MatchType": {
		//	                    "type": "string"
		//	                  },
		//	                  "RoleARN": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Claim",
		//	                  "MatchType",
		//	                  "RoleARN",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "Rules"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Type": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Type"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  }
		//	}
		"role_mappings":           // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AmbiguousRoleResolution
					"ambiguous_role_resolution": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: IdentityProvider
					"identity_provider": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RulesConfiguration
					"rules_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Rules
							"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Claim
										"claim": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: MatchType
										"match_type": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: RoleARN
										"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
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
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Roles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  }
		//	}
		"roles":             // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Cognito::IdentityPoolRoleAttachment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Cognito::IdentityPoolRoleAttachment").WithTerraformTypeName("awscc_cognito_identity_pool_role_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ambiguous_role_resolution":        "AmbiguousRoleResolution",
		"claim":                            "Claim",
		"identity_pool_id":                 "IdentityPoolId",
		"identity_pool_role_attachment_id": "Id",
		"identity_provider":                "IdentityProvider",
		"match_type":                       "MatchType",
		"role_arn":                         "RoleARN",
		"role_mappings":                    "RoleMappings",
		"roles":                            "Roles",
		"rules":                            "Rules",
		"rules_configuration":              "RulesConfiguration",
		"type":                             "Type",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
