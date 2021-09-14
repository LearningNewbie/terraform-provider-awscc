// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_network_insights_path", networkInsightsPathDataSourceType)
}

// networkInsightsPathDataSourceType returns the Terraform awscc_ec2_network_insights_path data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::NetworkInsightsPath resource type.
func networkInsightsPathDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_date": {
			// Property: CreatedDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"destination": {
			// Property: Destination
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"destination_ip": {
			// Property: DestinationIp
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"destination_port": {
			// Property: DestinationPort
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Computed: true,
		},
		"network_insights_path_arn": {
			// Property: NetworkInsightsPathArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"network_insights_path_id": {
			// Property: NetworkInsightsPathId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "tcp",
			//     "udp"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"source": {
			// Property: Source
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"source_ip": {
			// Property: SourceIp
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::NetworkInsightsPath",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsPath").WithTerraformTypeName("awscc_ec2_network_insights_path")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_date":              "CreatedDate",
		"destination":               "Destination",
		"destination_ip":            "DestinationIp",
		"destination_port":          "DestinationPort",
		"key":                       "Key",
		"network_insights_path_arn": "NetworkInsightsPathArn",
		"network_insights_path_id":  "NetworkInsightsPathId",
		"protocol":                  "Protocol",
		"source":                    "Source",
		"source_ip":                 "SourceIp",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}