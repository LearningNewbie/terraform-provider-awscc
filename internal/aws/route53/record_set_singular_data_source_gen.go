// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53_record_set", recordSetDataSource)
}

// recordSetDataSource returns the Terraform awscc_route53_record_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53::RecordSet resource.
func recordSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AliasTarget
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DNSName": {
		//	      "type": "string"
		//	    },
		//	    "EvaluateTargetHealth": {
		//	      "type": "boolean"
		//	    },
		//	    "HostedZoneId": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "HostedZoneId",
		//	    "DNSName"
		//	  ],
		//	  "type": "object"
		//	}
		"alias_target": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DNSName
				"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EvaluateTargetHealth
				"evaluate_target_health": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HostedZoneId
				"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CidrRoutingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CollectionId": {
		//	      "type": "string"
		//	    },
		//	    "LocationName": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "CollectionId",
		//	    "LocationName"
		//	  ],
		//	  "type": "object"
		//	}
		"cidr_routing_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CollectionId
				"collection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LocationName
				"location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Comment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Failover
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"failover": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: GeoLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ContinentCode": {
		//	      "type": "string"
		//	    },
		//	    "CountryCode": {
		//	      "type": "string"
		//	    },
		//	    "SubdivisionCode": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"geo_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ContinentCode
				"continent_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CountryCode
				"country_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SubdivisionCode
				"subdivision_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: GeoProximityLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AWSRegion": {
		//	      "type": "string"
		//	    },
		//	    "Bias": {
		//	      "type": "integer"
		//	    },
		//	    "Coordinates": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Latitude": {
		//	          "type": "string"
		//	        },
		//	        "Longitude": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Latitude",
		//	        "Longitude"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LocalZoneGroup": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"geo_proximity_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AWSRegion
				"aws_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Bias
				"bias": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Coordinates
				"coordinates": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Latitude
						"latitude": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Longitude
						"longitude": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LocalZoneGroup
				"local_zone_group": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"health_check_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: HostedZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: HostedZoneName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"hosted_zone_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"record_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MultiValueAnswer
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"multi_value_answer": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Region
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceRecords
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"resource_records": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SetIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"set_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TTL
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"ttl": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Weight
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53::RecordSet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::RecordSet").WithTerraformTypeName("awscc_route53_record_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_target":           "AliasTarget",
		"aws_region":             "AWSRegion",
		"bias":                   "Bias",
		"cidr_routing_config":    "CidrRoutingConfig",
		"collection_id":          "CollectionId",
		"comment":                "Comment",
		"continent_code":         "ContinentCode",
		"coordinates":            "Coordinates",
		"country_code":           "CountryCode",
		"dns_name":               "DNSName",
		"evaluate_target_health": "EvaluateTargetHealth",
		"failover":               "Failover",
		"geo_location":           "GeoLocation",
		"geo_proximity_location": "GeoProximityLocation",
		"health_check_id":        "HealthCheckId",
		"hosted_zone_id":         "HostedZoneId",
		"hosted_zone_name":       "HostedZoneName",
		"latitude":               "Latitude",
		"local_zone_group":       "LocalZoneGroup",
		"location_name":          "LocationName",
		"longitude":              "Longitude",
		"multi_value_answer":     "MultiValueAnswer",
		"name":                   "Name",
		"record_set_id":          "Id",
		"region":                 "Region",
		"resource_records":       "ResourceRecords",
		"set_identifier":         "SetIdentifier",
		"subdivision_code":       "SubdivisionCode",
		"ttl":                    "TTL",
		"type":                   "Type",
		"weight":                 "Weight",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
