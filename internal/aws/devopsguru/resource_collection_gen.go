// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_devopsguru_resource_collection", resourceCollectionResourceType)
}

// resourceCollectionResourceType returns the Terraform awscc_devopsguru_resource_collection resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DevOpsGuru::ResourceCollection resource type.
func resourceCollectionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"resource_collection_filter": {
			// Property: ResourceCollectionFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			//   "properties": {
			//     "CloudFormation": {
			//       "additionalProperties": false,
			//       "description": "CloudFormation resource for DevOps Guru to monitor",
			//       "properties": {
			//         "StackNames": {
			//           "description": "An array of CloudFormation stack names.",
			//           "items": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "maxItems": 200,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cloudformation": {
						// Property: CloudFormation
						Description: "CloudFormation resource for DevOps Guru to monitor",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"stack_names": {
									// Property: StackNames
									Description: "An array of CloudFormation stack names.",
									Type:        types.ListType{ElemType: types.StringType},
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLength(1, 200),
									},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"resource_collection_type": {
			// Property: ResourceCollectionType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of ResourceCollection",
			//   "enum": [
			//     "AWS_CLOUD_FORMATION"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of ResourceCollection",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::ResourceCollection").WithTerraformTypeName("awscc_devopsguru_resource_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudformation":             "CloudFormation",
		"resource_collection_filter": "ResourceCollectionFilter",
		"resource_collection_type":   "ResourceCollectionType",
		"stack_names":                "StackNames",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_devopsguru_resource_collection", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
