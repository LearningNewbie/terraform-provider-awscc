// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iottwinmaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iottwinmaker_sync_job", syncJobDataSource)
}

// syncJobDataSource returns the Terraform awscc_iottwinmaker_sync_job data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTTwinMaker::SyncJob resource.
func syncJobDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the SyncJob.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iottwinmaker:[a-z0-9-]+:[0-9]{12}:[\\/a-zA-Z0-9_\\-\\.:]+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the SyncJob.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationDateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the sync job was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_date_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the sync job was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of SyncJob.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z_\\-0-9]+",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of SyncJob.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SyncRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM Role that execute SyncJob.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iam::[0-9]{12}:role/.*",
		//	  "type": "string"
		//	}
		"sync_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM Role that execute SyncJob.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SyncSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source of the SyncJob.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"sync_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source of the SyncJob.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdateDateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the sync job was updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"update_date_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the sync job was updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkspaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the workspace.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
		//	  "type": "string"
		//	}
		"workspace_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the workspace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTTwinMaker::SyncJob",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTTwinMaker::SyncJob").WithTerraformTypeName("awscc_iottwinmaker_sync_job")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"creation_date_time": "CreationDateTime",
		"state":              "State",
		"sync_role":          "SyncRole",
		"sync_source":        "SyncSource",
		"tags":               "Tags",
		"update_date_time":   "UpdateDateTime",
		"workspace_id":       "WorkspaceId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
