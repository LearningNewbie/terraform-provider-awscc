// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_groundstation_mission_profile", missionProfileDataSource)
}

// missionProfileDataSource returns the Terraform awscc_groundstation_mission_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::GroundStation::MissionProfile resource.
func missionProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ContactPostPassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Post-pass time needed after the contact.",
		//	  "type": "integer"
		//	}
		"contact_post_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Post-pass time needed after the contact.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContactPrePassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pre-pass time needed before the contact.",
		//	  "type": "integer"
		//	}
		"contact_pre_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Pre-pass time needed before the contact.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataflowEdges
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Destination": {
		//	        "type": "string"
		//	      },
		//	      "Source": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"dataflow_edges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Destination
					"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Source
					"source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MinimumViableContactDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
		//	  "type": "integer"
		//	}
		"minimum_viable_contact_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name used to identify a mission profile.",
		//	  "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name used to identify a mission profile.",
			Computed:    true,
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
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrackingConfigArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"tracking_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GroundStation::MissionProfile",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::MissionProfile").WithTerraformTypeName("awscc_groundstation_mission_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"contact_post_pass_duration_seconds": "ContactPostPassDurationSeconds",
		"contact_pre_pass_duration_seconds":  "ContactPrePassDurationSeconds",
		"dataflow_edges":                     "DataflowEdges",
		"destination":                        "Destination",
		"id":                                 "Id",
		"key":                                "Key",
		"minimum_viable_contact_duration_seconds": "MinimumViableContactDurationSeconds",
		"name":                "Name",
		"region":              "Region",
		"source":              "Source",
		"tags":                "Tags",
		"tracking_config_arn": "TrackingConfigArn",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
