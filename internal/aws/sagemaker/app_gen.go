// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_sagemaker_app", appResourceType)
}

// appResourceType returns the Terraform aws_sagemaker_app resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::App resource type.
func appResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"app_arn": {
			// Property: AppArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the app.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the app.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_name": {
			// Property: AppName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the app.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the app.",
			Type:        types.StringType,
			Required:    true,
			// AppName is a force-new attribute.
		},
		"app_type": {
			// Property: AppType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of app.",
			//   "enum": [
			//     "JupyterServer",
			//     "KernelGateway"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of app.",
			Type:        types.StringType,
			Required:    true,
			// AppType is a force-new attribute.
		},
		"domain_id": {
			// Property: DomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain ID.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The domain ID.",
			Type:        types.StringType,
			Required:    true,
			// DomainId is a force-new attribute.
		},
		"resource_spec": {
			// Property: ResourceSpec
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "InstanceType": {
			//       "description": "The instance type that the image version runs on.",
			//       "enum": [
			//         "system",
			//         "ml.t3.micro",
			//         "ml.t3.small",
			//         "ml.t3.medium",
			//         "ml.t3.large",
			//         "ml.t3.xlarge",
			//         "ml.t3.2xlarge",
			//         "ml.m5.large",
			//         "ml.m5.xlarge",
			//         "ml.m5.2xlarge",
			//         "ml.m5.4xlarge",
			//         "ml.m5.8xlarge",
			//         "ml.m5.12xlarge",
			//         "ml.m5.16xlarge",
			//         "ml.m5.24xlarge",
			//         "ml.c5.large",
			//         "ml.c5.xlarge",
			//         "ml.c5.2xlarge",
			//         "ml.c5.4xlarge",
			//         "ml.c5.9xlarge",
			//         "ml.c5.12xlarge",
			//         "ml.c5.18xlarge",
			//         "ml.c5.24xlarge",
			//         "ml.p3.2xlarge",
			//         "ml.p3.8xlarge",
			//         "ml.p3.16xlarge",
			//         "ml.g4dn.xlarge",
			//         "ml.g4dn.2xlarge",
			//         "ml.g4dn.4xlarge",
			//         "ml.g4dn.8xlarge",
			//         "ml.g4dn.12xlarge",
			//         "ml.g4dn.16xlarge"
			//       ],
			//       "type": "string"
			//     },
			//     "SageMakerImageArn": {
			//       "description": "The ARN of the SageMaker image that the image version belongs to.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SageMakerImageVersionArn": {
			//       "description": "The ARN of the image version created on the instance.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"instance_type": {
						// Property: InstanceType
						Description: "The instance type that the image version runs on.",
						Type:        types.StringType,
						Optional:    true,
					},
					"sage_maker_image_arn": {
						// Property: SageMakerImageArn
						Description: "The ARN of the SageMaker image that the image version belongs to.",
						Type:        types.StringType,
						Optional:    true,
					},
					"sage_maker_image_version_arn": {
						// Property: SageMakerImageVersionArn
						Description: "The ARN of the image version created on the instance.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the app.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of tags to apply to the app.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
			// Tags is a write-only attribute.
		},
		"user_profile_name": {
			// Property: UserProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "The user profile name.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The user profile name.",
			Type:        types.StringType,
			Required:    true,
			// UserProfileName is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SageMaker::App",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::App").WithTerraformTypeName("aws_sagemaker_app").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sagemaker_app", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
