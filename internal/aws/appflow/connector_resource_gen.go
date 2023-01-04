// Code generated by generators/resource/main.go; DO NOT EDIT.

package appflow

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
	registry.AddResourceFactory("awscc_appflow_connector", connectorResource)
}

// connectorResource returns the Terraform awscc_appflow_connector resource.
// This Terraform resource corresponds to the CloudFormation AWS::AppFlow::Connector resource.
func connectorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConnectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.",
		//	  "maxLength": 512,
		//	  "pattern": "arn:*:appflow:.*:[0-9]+:.*",
		//	  "type": "string"
		//	}
		"connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: " The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConnectorLabel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.",
		//	  "maxLength": 512,
		//	  "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
		//	  "type": "string"
		//	}
		"connector_label": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: " The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(512),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9][\\w!@#.-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConnectorProvisioningConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains information about the configuration of the connector being registered.",
		//	  "properties": {
		//	    "Lambda": {
		//	      "additionalProperties": false,
		//	      "description": "Contains information about the configuration of the lambda which is being registered as the connector.",
		//	      "properties": {
		//	        "LambdaArn": {
		//	          "description": "Lambda ARN of the connector being registered.",
		//	          "maxLength": 512,
		//	          "pattern": "arn:*:.*:.*:[0-9]+:.*",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LambdaArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"connector_provisioning_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Lambda
				"lambda": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LambdaArn
						"lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Lambda ARN of the connector being registered.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(512),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:*:.*:.*:[0-9]+:.*"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains information about the configuration of the lambda which is being registered as the connector.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains information about the configuration of the connector being registered.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectorProvisioningType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provisioning type of the connector. Currently the only supported value is LAMBDA. ",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
		//	  "type": "string"
		//	}
		"connector_provisioning_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provisioning type of the connector. Currently the only supported value is LAMBDA. ",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9][\\w!@#.-]+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description about the connector that's being registered.",
		//	  "maxLength": 2048,
		//	  "pattern": "[\\s\\w/!@#+=.-]*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description about the connector that's being registered.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\w/!@#+=.-]*"), ""),
			}, /*END VALIDATORS*/
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
		Description: "Resource schema for AWS::AppFlow::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppFlow::Connector").WithTerraformTypeName("awscc_appflow_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connector_arn":                 "ConnectorArn",
		"connector_label":               "ConnectorLabel",
		"connector_provisioning_config": "ConnectorProvisioningConfig",
		"connector_provisioning_type":   "ConnectorProvisioningType",
		"description":                   "Description",
		"lambda":                        "Lambda",
		"lambda_arn":                    "LambdaArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
