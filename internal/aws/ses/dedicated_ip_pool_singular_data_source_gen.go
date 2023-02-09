// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ses_dedicated_ip_pool", dedicatedIpPoolDataSource)
}

// dedicatedIpPoolDataSource returns the Terraform awscc_ses_dedicated_ip_pool data source.
// This Terraform data source corresponds to the CloudFormation AWS::SES::DedicatedIpPool resource.
func dedicatedIpPoolDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: PoolName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the dedicated IP pool.",
		//	  "pattern": "^[a-z0-9_-]{0,64}$",
		//	  "type": "string"
		//	}
		"pool_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the dedicated IP pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScalingMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.",
		//	  "pattern": "^(STANDARD|MANAGED)$",
		//	  "type": "string"
		//	}
		"scaling_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SES::DedicatedIpPool",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::DedicatedIpPool").WithTerraformTypeName("awscc_ses_dedicated_ip_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"pool_name":    "PoolName",
		"scaling_mode": "ScalingMode",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
