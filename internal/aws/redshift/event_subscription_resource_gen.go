// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_redshift_event_subscription", eventSubscriptionResource)
}

// eventSubscriptionResource returns the Terraform awscc_redshift_event_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::Redshift::EventSubscription resource.
func eventSubscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CustSubscriptionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon Redshift event notification subscription.",
		//	  "type": "string"
		//	}
		"cust_subscription_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon Redshift event notification subscription.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomerAwsId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account associated with the Amazon Redshift event notification subscription.",
		//	  "type": "string"
		//	}
		"customer_aws_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account associated with the Amazon Redshift event notification subscription.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A boolean value; set to true to activate the subscription, and set to false to create the subscription but not activate it.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventCategories
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "configuration",
		//	      "management",
		//	      "monitoring",
		//	      "security",
		//	      "pending"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"event_categories": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Specifies the Amazon Redshift event categories to be published by the event notification subscription.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"configuration",
						"management",
						"monitoring",
						"security",
						"pending",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventCategoriesList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of Amazon Redshift event categories specified in the event notification subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"event_categories_list": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of Amazon Redshift event categories specified in the event notification subscription.",
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Severity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
		//	  "enum": [
		//	    "ERROR",
		//	    "INFO"
		//	  ],
		//	  "type": "string"
		//	}
		"severity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the Amazon Redshift event severity to be published by the event notification subscription.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ERROR",
					"INFO",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
		//	  "type": "string"
		//	}
		"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of one or more identifiers of Amazon Redshift source objects.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"source_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of one or more identifiers of Amazon Redshift source objects.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceIdsList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"source_ids_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of source that will be generating the events.",
		//	  "enum": [
		//	    "cluster",
		//	    "cluster-parameter-group",
		//	    "cluster-security-group",
		//	    "cluster-snapshot",
		//	    "scheduled-action"
		//	  ],
		//	  "type": "string"
		//	}
		"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of source that will be generating the events.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"cluster",
					"cluster-parameter-group",
					"cluster-security-group",
					"cluster-snapshot",
					"scheduled-action",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Amazon Redshift event notification subscription.",
		//	  "enum": [
		//	    "active",
		//	    "no-permission",
		//	    "topic-not-exist"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Amazon Redshift event notification subscription.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionCreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time the Amazon Redshift event notification subscription was created.",
		//	  "type": "string"
		//	}
		"subscription_creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time the Amazon Redshift event notification subscription was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon Redshift event notification subscription",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"subscription_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon Redshift event notification subscription",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The `AWS::Redshift::EventSubscription` resource creates an Amazon Redshift Event Subscription.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EventSubscription").WithTerraformTypeName("awscc_redshift_event_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cust_subscription_id":       "CustSubscriptionId",
		"customer_aws_id":            "CustomerAwsId",
		"enabled":                    "Enabled",
		"event_categories":           "EventCategories",
		"event_categories_list":      "EventCategoriesList",
		"key":                        "Key",
		"severity":                   "Severity",
		"sns_topic_arn":              "SnsTopicArn",
		"source_ids":                 "SourceIds",
		"source_ids_list":            "SourceIdsList",
		"source_type":                "SourceType",
		"status":                     "Status",
		"subscription_creation_time": "SubscriptionCreationTime",
		"subscription_name":          "SubscriptionName",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
