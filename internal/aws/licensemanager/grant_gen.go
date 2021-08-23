// Code generated by generators/resource/main.go; DO NOT EDIT.

package licensemanager

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
	registry.AddResourceTypeFactory("awscc_licensemanager_grant", grantResourceType)
}

// grantResourceType returns the Terraform awscc_licensemanager_grant resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::LicenseManager::Grant resource type.
func grantResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allowed_operations": {
			// Property: AllowedOperations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type: types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			Optional: true,
			// AllowedOperations is a write-only attribute.
		},
		"grant_arn": {
			// Property: GrantArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"grant_name": {
			// Property: GrantName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for the created Grant.",
			//   "type": "string"
			// }
			Description: "Name for the created Grant.",
			Type:        types.StringType,
			Optional:    true,
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "Home region for the created grant.",
			//   "type": "string"
			// }
			Description: "Home region for the created grant.",
			Type:        types.StringType,
			Optional:    true,
		},
		"license_arn": {
			// Property: LicenseArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"principals": {
			// Property: Principals
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 2048,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type: types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			Optional: true,
			// Principals is a write-only attribute.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the grant.",
			//   "type": "string"
			// }
			Description: "The version of the grant.",
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LicenseManager::Grant").WithTerraformTypeName("awscc_licensemanager_grant")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_operations": "AllowedOperations",
		"grant_arn":          "GrantArn",
		"grant_name":         "GrantName",
		"home_region":        "HomeRegion",
		"license_arn":        "LicenseArn",
		"principals":         "Principals",
		"status":             "Status",
		"version":            "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Principals",
		"/properties/AllowedOperations",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_licensemanager_grant", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
