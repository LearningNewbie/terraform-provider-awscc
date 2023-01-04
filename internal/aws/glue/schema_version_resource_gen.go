// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_glue_schema_version", schemaVersionResource)
}

// schemaVersionResource returns the Terraform awscc_glue_schema_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::Glue::SchemaVersion resource.
func schemaVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Identifier for the schema where the schema version will be created.",
		//	  "properties": {
		//	    "RegistryName": {
		//	      "description": "Name of the registry to identify where the Schema is located.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "SchemaArn": {
		//	      "description": "Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.",
		//	      "pattern": "arn:(aws|aws-us-gov|aws-cn):glue:.*",
		//	      "type": "string"
		//	    },
		//	    "SchemaName": {
		//	      "description": "Name of the schema. This parameter requires RegistryName to be provided.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RegistryName
				"registry_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Name of the registry to identify where the Schema is located.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SchemaArn
				"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):glue:.*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SchemaName
				"schema_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Name of the schema. This parameter requires RegistryName to be provided.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Identifier for the schema where the schema version will be created.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Complete definition of the schema in plain-text.",
		//	  "maxLength": 170000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"schema_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Complete definition of the schema in plain-text.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 170000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents the version ID associated with the schema version.",
		//	  "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents the version ID associated with the schema version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "This resource represents an individual schema version of a schema defined in Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::SchemaVersion").WithTerraformTypeName("awscc_glue_schema_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"registry_name":     "RegistryName",
		"schema":            "Schema",
		"schema_arn":        "SchemaArn",
		"schema_definition": "SchemaDefinition",
		"schema_name":       "SchemaName",
		"version_id":        "VersionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
