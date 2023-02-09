// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediaconnect_flow_vpc_interface", flowVpcInterfaceDataSource)
}

// flowVpcInterfaceDataSource returns the Terraform awscc_mediaconnect_flow_vpc_interface data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaConnect::FlowVpcInterface resource.
func flowVpcInterfaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: FlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
		//	  "type": "string"
		//	}
		"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IDs of the network interfaces created in customer's account by MediaConnect.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"network_interface_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "IDs of the network interfaces created in customer's account by MediaConnect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Security Group IDs to be used on ENI.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Security Group IDs to be used on ENI.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Subnet must be in the AZ of the Flow",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Subnet must be in the AZ of the Flow",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaConnect::FlowVpcInterface",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowVpcInterface").WithTerraformTypeName("awscc_mediaconnect_flow_vpc_interface")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"flow_arn":              "FlowArn",
		"name":                  "Name",
		"network_interface_ids": "NetworkInterfaceIds",
		"role_arn":              "RoleArn",
		"security_group_ids":    "SecurityGroupIds",
		"subnet_id":             "SubnetId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
