// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codedeploy_deployment_config", deploymentConfigDataSource)
}

// deploymentConfigDataSource returns the Terraform awscc_codedeploy_deployment_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeDeploy::DeploymentConfig resource.
func deploymentConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ComputePlatform
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The destination platform type for the deployment (Lambda, Server, or ECS).",
		//	  "type": "string"
		//	}
		"compute_platform": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The destination platform type for the deployment (Lambda, Server, or ECS).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentConfigName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
		//	  "type": "string"
		//	}
		"deployment_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinimumHealthyHosts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
		//	  "properties": {
		//	    "Type": {
		//	      "type": "string"
		//	    },
		//	    "Value": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "Value"
		//	  ],
		//	  "type": "object"
		//	}
		"minimum_healthy_hosts": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Value
				"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficRoutingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration that specifies how the deployment traffic is routed.",
		//	  "properties": {
		//	    "TimeBasedCanary": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CanaryInterval": {
		//	          "type": "integer"
		//	        },
		//	        "CanaryPercentage": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "CanaryPercentage",
		//	        "CanaryInterval"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeBasedLinear": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "LinearInterval": {
		//	          "type": "integer"
		//	        },
		//	        "LinearPercentage": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "LinearInterval",
		//	        "LinearPercentage"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"traffic_routing_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TimeBasedCanary
				"time_based_canary": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CanaryInterval
						"canary_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: CanaryPercentage
						"canary_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TimeBasedLinear
				"time_based_linear": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LinearInterval
						"linear_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: LinearPercentage
						"linear_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration that specifies how the deployment traffic is routed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeDeploy::DeploymentConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeDeploy::DeploymentConfig").WithTerraformTypeName("awscc_codedeploy_deployment_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"canary_interval":        "CanaryInterval",
		"canary_percentage":      "CanaryPercentage",
		"compute_platform":       "ComputePlatform",
		"deployment_config_name": "DeploymentConfigName",
		"linear_interval":        "LinearInterval",
		"linear_percentage":      "LinearPercentage",
		"minimum_healthy_hosts":  "MinimumHealthyHosts",
		"time_based_canary":      "TimeBasedCanary",
		"time_based_linear":      "TimeBasedLinear",
		"traffic_routing_config": "TrafficRoutingConfig",
		"type":                   "Type",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
