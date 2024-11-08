// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bedrock_prompt_version", promptVersionDataSource)
}

// promptVersionDataSource returns the Terraform awscc_bedrock_prompt_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Bedrock::PromptVersion resource.
func promptVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of a prompt version resource",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10}:[0-9]{1,20})$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of a prompt version resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerEncryptionKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A KMS key ARN",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$",
		//	  "type": "string"
		//	}
		"customer_encryption_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A KMS key ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultVariant
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for a variant.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"default_variant": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for a variant.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description for a prompt version resource.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description for a prompt version resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for a prompt resource.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for a prompt resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PromptArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of a prompt resource possibly with a version",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10})$",
		//	  "type": "string"
		//	}
		"prompt_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of a prompt resource possibly with a version",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PromptId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier for a Prompt",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"prompt_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier for a Prompt",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map of tag keys and values",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Value of a tag",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of tag keys and values",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Variants
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of prompt variants",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Prompt variant",
		//	    "properties": {
		//	      "InferenceConfiguration": {
		//	        "description": "Model inference configuration",
		//	        "properties": {
		//	          "Text": {
		//	            "additionalProperties": false,
		//	            "description": "Prompt model inference configuration",
		//	            "properties": {
		//	              "MaxTokens": {
		//	                "description": "Maximum length of output",
		//	                "maximum": 4096,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "StopSequences": {
		//	                "description": "List of stop sequences",
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 4,
		//	                "minItems": 0,
		//	                "type": "array"
		//	              },
		//	              "Temperature": {
		//	                "description": "Controls randomness, higher values increase diversity",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "TopP": {
		//	                "description": "Cumulative probability cutoff for token selection",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ModelId": {
		//	        "description": "ARN or Id of a Bedrock Foundational Model or Inference Profile, or the ARN of a imported model, or a provisioned throughput ARN for custom models.",
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "^(arn:aws(-[^:]{1,12})?:(bedrock|sagemaker):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "Name for a variant.",
		//	        "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	        "type": "string"
		//	      },
		//	      "TemplateConfiguration": {
		//	        "description": "Prompt template configuration",
		//	        "properties": {
		//	          "Text": {
		//	            "additionalProperties": false,
		//	            "description": "Configuration for text prompt template",
		//	            "properties": {
		//	              "InputVariables": {
		//	                "description": "List of input variables",
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "additionalProperties": false,
		//	                  "description": "Input variable",
		//	                  "properties": {
		//	                    "Name": {
		//	                      "description": "Name for an input variable",
		//	                      "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "maxItems": 5,
		//	                "minItems": 0,
		//	                "type": "array"
		//	              },
		//	              "Text": {
		//	                "description": "Prompt content for String prompt template",
		//	                "maxLength": 200000,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Text"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "TemplateType": {
		//	        "description": "Prompt template type",
		//	        "enum": [
		//	          "TEXT"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "TemplateType",
		//	      "TemplateConfiguration"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"variants": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InferenceConfiguration
					"inference_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Text
							"text": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: MaxTokens
									"max_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum length of output",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: StopSequences
									"stop_sequences": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "List of stop sequences",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Temperature
									"temperature": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Controls randomness, higher values increase diversity",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: TopP
									"top_p": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Cumulative probability cutoff for token selection",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Prompt model inference configuration",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Model inference configuration",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ModelId
					"model_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "ARN or Id of a Bedrock Foundational Model or Inference Profile, or the ARN of a imported model, or a provisioned throughput ARN for custom models.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Name for a variant.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TemplateConfiguration
					"template_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Text
							"text": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: InputVariables
									"input_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
										NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Name
												"name": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "Name for an input variable",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
										}, /*END NESTED OBJECT*/
										Description: "List of input variables",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Text
									"text": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Prompt content for String prompt template",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Configuration for text prompt template",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Prompt template configuration",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TemplateType
					"template_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Prompt template type",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of prompt variants",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^(DRAFT|[0-9]{0,4}[1-9][0-9]{0,4})$",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Bedrock::PromptVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::PromptVersion").WithTerraformTypeName("awscc_bedrock_prompt_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"created_at":                  "CreatedAt",
		"customer_encryption_key_arn": "CustomerEncryptionKeyArn",
		"default_variant":             "DefaultVariant",
		"description":                 "Description",
		"inference_configuration":     "InferenceConfiguration",
		"input_variables":             "InputVariables",
		"max_tokens":                  "MaxTokens",
		"model_id":                    "ModelId",
		"name":                        "Name",
		"prompt_arn":                  "PromptArn",
		"prompt_id":                   "PromptId",
		"stop_sequences":              "StopSequences",
		"tags":                        "Tags",
		"temperature":                 "Temperature",
		"template_configuration":      "TemplateConfiguration",
		"template_type":               "TemplateType",
		"text":                        "Text",
		"top_p":                       "TopP",
		"updated_at":                  "UpdatedAt",
		"variants":                    "Variants",
		"version":                     "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
