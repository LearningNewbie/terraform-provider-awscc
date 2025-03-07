// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_subscription_filter", subscriptionFilterDataSource)
}

// subscriptionFilterDataSource returns the Terraform awscc_logs_subscription_filter data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::SubscriptionFilter resource.
func subscriptionFilterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplyOnTransformedLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This parameter is valid only for log groups that have an active log transformer. For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).\n If this value is ``true``, the subscription filter is applied on the transformed version of the log events instead of the original ingested log events.",
		//	  "type": "boolean"
		//	}
		"apply_on_transformed_logs": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "This parameter is valid only for log groups that have an active log transformer. For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).\n If this value is ``true``, the subscription filter is applied on the transformed version of the log events instead of the original ingested log events.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the destination.",
		//	  "type": "string"
		//	}
		"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Distribution
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The method used to distribute log data to the destination, which can be either random or grouped by log stream.",
		//	  "enum": [
		//	    "Random",
		//	    "ByLogStream"
		//	  ],
		//	  "type": "string"
		//	}
		"distribution": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The method used to distribute log data to the destination, which can be either random or grouped by log stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FilterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subscription filter.",
		//	  "type": "string"
		//	}
		"filter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subscription filter.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FilterPattern
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The filtering expressions that restrict what gets delivered to the destination AWS resource. For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).",
		//	  "type": "string"
		//	}
		"filter_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The filtering expressions that restrict what gets delivered to the destination AWS resource. For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The log group to associate with the subscription filter. All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.",
		//	  "type": "string"
		//	}
		"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The log group to associate with the subscription filter. All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of an IAM role that grants CWL permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::SubscriptionFilter",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::SubscriptionFilter").WithTerraformTypeName("awscc_logs_subscription_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"apply_on_transformed_logs": "ApplyOnTransformedLogs",
		"destination_arn":           "DestinationArn",
		"distribution":              "Distribution",
		"filter_name":               "FilterName",
		"filter_pattern":            "FilterPattern",
		"log_group_name":            "LogGroupName",
		"role_arn":                  "RoleArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
