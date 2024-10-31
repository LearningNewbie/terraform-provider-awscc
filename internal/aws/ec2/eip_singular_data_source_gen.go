// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_eip", eIPDataSource)
}

// eIPDataSource returns the Terraform awscc_ec2_eip data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::EIP resource.
func eIPDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Address
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AllocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The network (``vpc``).\n If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The network (``vpc``).\n If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the instance.\n  Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the instance.\n  Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkBorderGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.\n Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.",
		//	  "type": "string"
		//	}
		"network_border_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.\n Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"public_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicIpv4Pool
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.\n  Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
		//	  "type": "string"
		//	}
		"public_ipv_4_pool": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.\n  Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the Elastic IP address.\n  Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Add tags to a resource](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#cloudformation-add-tag-specifications).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the Elastic IP address.\n  Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransferAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Elastic IP address you are accepting for transfer. You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide*.",
		//	  "type": "string"
		//	}
		"transfer_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Elastic IP address you are accepting for transfer. You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::EIP",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EIP").WithTerraformTypeName("awscc_ec2_eip")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":              "Address",
		"allocation_id":        "AllocationId",
		"domain":               "Domain",
		"instance_id":          "InstanceId",
		"ipam_pool_id":         "IpamPoolId",
		"key":                  "Key",
		"network_border_group": "NetworkBorderGroup",
		"public_ip":            "PublicIp",
		"public_ipv_4_pool":    "PublicIpv4Pool",
		"tags":                 "Tags",
		"transfer_address":     "TransferAddress",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
