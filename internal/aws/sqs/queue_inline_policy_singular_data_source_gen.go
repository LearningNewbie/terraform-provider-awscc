// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sqs_queue_inline_policy", queueInlinePolicyDataSource)
}

// queueInlinePolicyDataSource returns the Terraform awscc_sqs_queue_inline_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::SQS::QueueInlinePolicy resource.
func queueInlinePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document that contains permissions to add to the specified SQS queue",
		//	  "type": "object"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "A policy document that contains permissions to add to the specified SQS queue",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Queue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the SQS queue.",
		//	  "type": "string"
		//	}
		"queue": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the SQS queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SQS::QueueInlinePolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SQS::QueueInlinePolicy").WithTerraformTypeName("awscc_sqs_queue_inline_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"policy_document": "PolicyDocument",
		"queue":           "Queue",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
