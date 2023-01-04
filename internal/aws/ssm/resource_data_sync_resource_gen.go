// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ssm_resource_data_sync", resourceDataSyncResource)
}

// resourceDataSyncResource returns the Terraform awscc_ssm_resource_data_sync resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSM::ResourceDataSync resource.
func resourceDataSyncResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BucketPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BucketRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"bucket_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KMSKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 512),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3Destination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BucketName": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "BucketPrefix": {
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "BucketRegion": {
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "KMSKeyArn": {
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "SyncFormat": {
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "BucketName",
		//	    "BucketRegion",
		//	    "SyncFormat"
		//	  ],
		//	  "type": "object"
		//	}
		"s3_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BucketName
				"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: BucketPrefix
				"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: BucketRegion
				"bucket_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 64),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: KMSKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 512),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SyncFormat
				"sync_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 1024),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SyncFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"sync_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SyncName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"sync_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SyncSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AwsOrganizationsSource": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "OrganizationSourceType": {
		//	          "maxLength": 64,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "OrganizationalUnits": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "OrganizationSourceType"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IncludeFutureRegions": {
		//	      "type": "boolean"
		//	    },
		//	    "SourceRegions": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "SourceType": {
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SourceType",
		//	    "SourceRegions"
		//	  ],
		//	  "type": "object"
		//	}
		"sync_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AwsOrganizationsSource
				"aws_organizations_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OrganizationSourceType
						"organization_source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 64),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: OrganizationalUnits
						"organizational_units": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IncludeFutureRegions
				"include_future_regions": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SourceRegions
				"source_regions": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceType
				"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 64),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SyncType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"sync_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::SSM::ResourceDataSync",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourceDataSync").WithTerraformTypeName("awscc_ssm_resource_data_sync")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_organizations_source": "AwsOrganizationsSource",
		"bucket_name":              "BucketName",
		"bucket_prefix":            "BucketPrefix",
		"bucket_region":            "BucketRegion",
		"include_future_regions":   "IncludeFutureRegions",
		"kms_key_arn":              "KMSKeyArn",
		"organization_source_type": "OrganizationSourceType",
		"organizational_units":     "OrganizationalUnits",
		"s3_destination":           "S3Destination",
		"source_regions":           "SourceRegions",
		"source_type":              "SourceType",
		"sync_format":              "SyncFormat",
		"sync_name":                "SyncName",
		"sync_source":              "SyncSource",
		"sync_type":                "SyncType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
