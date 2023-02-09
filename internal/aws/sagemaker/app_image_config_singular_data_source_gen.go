// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_app_image_config", appImageConfigDataSource)
}

// appImageConfigDataSource returns the Terraform awscc_sagemaker_app_image_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::AppImageConfig resource.
func appImageConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppImageConfigArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AppImageConfig.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:app-image-config/.*",
		//	  "type": "string"
		//	}
		"app_image_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AppImageConfig.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AppImageConfigName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Name of the AppImageConfig.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
		//	  "type": "string"
		//	}
		"app_image_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Name of the AppImageConfig.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KernelGatewayImageConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The KernelGatewayImageConfig.",
		//	  "properties": {
		//	    "FileSystemConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
		//	      "properties": {
		//	        "DefaultGid": {
		//	          "description": "The default POSIX group ID (GID). If not specified, defaults to 100.",
		//	          "maximum": 65535,
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "DefaultUid": {
		//	          "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
		//	          "maximum": 65535,
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "MountPath": {
		//	          "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
		//	          "maxLength": 1024,
		//	          "minLength": 1,
		//	          "pattern": "^/.*",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "KernelSpecs": {
		//	      "description": "The specification of the Jupyter kernels in the image.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "DisplayName": {
		//	            "description": "The display name of the kernel.",
		//	            "maxLength": 1024,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the kernel.",
		//	            "maxLength": 1024,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "KernelSpecs"
		//	  ],
		//	  "type": "object"
		//	}
		"kernel_gateway_image_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FileSystemConfig
				"file_system_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DefaultGid
						"default_gid": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The default POSIX group ID (GID). If not specified, defaults to 100.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DefaultUid
						"default_uid": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The default POSIX user ID (UID). If not specified, defaults to 1000.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MountPath
						"mount_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KernelSpecs
				"kernel_specs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DisplayName
							"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The display name of the kernel.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the kernel.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The specification of the Jupyter kernels in the image.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The KernelGatewayImageConfig.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to the AppImageConfig.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
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
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "A list of tags to apply to the AppImageConfig.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::AppImageConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::AppImageConfig").WithTerraformTypeName("awscc_sagemaker_app_image_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_arn":        "AppImageConfigArn",
		"app_image_config_name":       "AppImageConfigName",
		"default_gid":                 "DefaultGid",
		"default_uid":                 "DefaultUid",
		"display_name":                "DisplayName",
		"file_system_config":          "FileSystemConfig",
		"kernel_gateway_image_config": "KernelGatewayImageConfig",
		"kernel_specs":                "KernelSpecs",
		"key":                         "Key",
		"mount_path":                  "MountPath",
		"name":                        "Name",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
