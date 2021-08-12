// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

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
	registry.AddResourceTypeFactory("aws_apigateway_documentation_version", documentationVersionResourceType)
}

// documentationVersionResourceType returns the Terraform aws_apigateway_documentation_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ApiGateway::DocumentationVersion resource type.
func documentationVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the API documentation snapshot.",
			//   "type": "string"
			// }
			Description: "The description of the API documentation snapshot.",
			Type:        types.StringType,
			Optional:    true,
		},
		"documentation_version": {
			// Property: DocumentationVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version identifier of the API documentation snapshot.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The version identifier of the API documentation snapshot.",
			Type:        types.StringType,
			Required:    true,
			// DocumentationVersion is a force-new attribute.
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the API.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The identifier of the API.",
			Type:        types.StringType,
			Required:    true,
			// RestApiId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A snapshot of the documentation of an API.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationVersion").WithTerraformTypeName("aws_apigateway_documentation_version").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_apigateway_documentation_version", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
