// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecr_repository", repositoryDataSource)
}

// repositoryDataSource returns the Terraform awscc_ecr_repository data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECR::Repository resource.
func repositoryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EmptyOnDelete
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.",
		//	  "type": "boolean"
		//	}
		"empty_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
		//	  "properties": {
		//	    "EncryptionType": {
		//	      "description": "The encryption type to use.\n If you use the ``KMS`` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an key stored in (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.\n If you use the ``AES256`` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.",
		//	      "enum": [
		//	        "AES256",
		//	        "KMS",
		//	        "KMS_DSSE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KmsKey": {
		//	      "description": "If you use the ``KMS`` encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed KMS key for Amazon ECR will be used.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EncryptionType"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionType
				"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The encryption type to use.\n If you use the ``KMS`` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an key stored in (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.\n If you use the ``AES256`` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKey
				"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "If you use the ``KMS`` encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed KMS key for Amazon ECR will be used.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageScanningConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
		//	  "properties": {
		//	    "ScanOnPush": {
		//	      "description": "The setting that determines whether images are scanned after being pushed to a repository. If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_scanning_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ScanOnPush
				"scan_on_push": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "The setting that determines whether images are scanned after being pushed to a repository. If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageTagMutability
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
		//	  "enum": [
		//	    "MUTABLE",
		//	    "IMMUTABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"image_tag_mutability": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecyclePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).",
		//	  "properties": {
		//	    "LifecyclePolicyText": {
		//	      "description": "The JSON repository policy text to apply to the repository.",
		//	      "maxLength": 30720,
		//	      "minLength": 100,
		//	      "type": "string"
		//	    },
		//	    "RegistryId": {
		//	      "description": "The AWS account ID associated with the registry that contains the repository. If you do? not specify a registry, the default registry is assumed.",
		//	      "maxLength": 12,
		//	      "minLength": 12,
		//	      "pattern": "^[0-9]{12}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lifecycle_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LifecyclePolicyText
				"lifecycle_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON repository policy text to apply to the repository.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RegistryId
				"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS account ID associated with the registry that contains the repository. If you do? not specify a registry, the default registry is assumed.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
		//	  "maxLength": 256,
		//	  "minLength": 2,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryPolicyText
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.",
		//	  "type": "string"
		//	}
		"repository_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"repository_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The metadata to apply to a resource to help you categorize and organize them. Each tag consists of a key and a value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A ``value`` acts as a descriptor within a tag category (key).",
		//	        "maxLength": 255,
		//	        "minLength": 1,
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
						Description: "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A ``value`` acts as a descriptor within a tag category (key).",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECR::Repository",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::Repository").WithTerraformTypeName("awscc_ecr_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"empty_on_delete":              "EmptyOnDelete",
		"encryption_configuration":     "EncryptionConfiguration",
		"encryption_type":              "EncryptionType",
		"image_scanning_configuration": "ImageScanningConfiguration",
		"image_tag_mutability":         "ImageTagMutability",
		"key":                          "Key",
		"kms_key":                      "KmsKey",
		"lifecycle_policy":             "LifecyclePolicy",
		"lifecycle_policy_text":        "LifecyclePolicyText",
		"registry_id":                  "RegistryId",
		"repository_name":              "RepositoryName",
		"repository_policy_text":       "RepositoryPolicyText",
		"repository_uri":               "RepositoryUri",
		"scan_on_push":                 "ScanOnPush",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
