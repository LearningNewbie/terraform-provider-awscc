// Code generated by generators/resource/main.go; DO NOT EDIT.

package codestarconnections

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
	registry.AddResourceTypeFactory("aws_codestarconnections_connection", connectionResourceType)
}

// connectionResourceType returns the Terraform aws_codestarconnections_connection resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeStarConnections::Connection resource type.
func connectionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connection_name": {
			// Property: ConnectionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the connection. Connection names must be unique in an AWS user account.",
			//   "maxLength": 32,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the connection. Connection names must be unique in an AWS user account.",
			Type:        types.StringType,
			Required:    true,
			// ConnectionName is a force-new attribute.
		},
		"connection_status": {
			// Property: ConnectionStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The current status of the connection.",
			//   "type": "string"
			// }
			Description: "The current status of the connection.",
			Type:        types.StringType,
			Computed:    true,
		},
		"host_arn": {
			// Property: HostArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// HostArn is a force-new attribute.
		},
		"owner_account_id": {
			// Property: OwnerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.",
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.",
			Type:        types.StringType,
			Computed:    true,
		},
		"provider_type": {
			// Property: ProviderType
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.",
			//   "type": "string"
			// }
			Description: "The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ProviderType is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the tags applied to a connection.",
			//   "items": {
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Specifies the tags applied to a connection.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Schema for AWS::CodeStarConnections::Connection resource which can be used to connect external source providers with AWS CodePipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarConnections::Connection").WithTerraformTypeName("aws_codestarconnections_connection").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_codestarconnections_connection", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
