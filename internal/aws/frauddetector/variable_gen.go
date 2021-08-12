// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_frauddetector_variable", variableResourceType)
}

// variableResourceType returns the Terraform aws_frauddetector_variable resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FraudDetector::Variable resource type.
func variableResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the variable.",
			//   "type": "string"
			// }
			Description: "The ARN of the variable.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the variable was created.",
			//   "type": "string"
			// }
			Description: "The time when the variable was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_source": {
			// Property: DataSource
			// CloudFormation resource type schema:
			// {
			//   "description": "The source of the data.",
			//   "enum": [
			//     "EVENT",
			//     "EXTERNAL_MODEL_SCORE"
			//   ],
			//   "type": "string"
			// }
			Description: "The source of the data.",
			Type:        types.StringType,
			Required:    true,
		},
		"data_type": {
			// Property: DataType
			// CloudFormation resource type schema:
			// {
			//   "description": "The data type.",
			//   "enum": [
			//     "STRING",
			//     "INTEGER",
			//     "FLOAT",
			//     "BOOLEAN"
			//   ],
			//   "type": "string"
			// }
			Description: "The data type.",
			Type:        types.StringType,
			Required:    true,
		},
		"default_value": {
			// Property: DefaultValue
			// CloudFormation resource type schema:
			// {
			//   "description": "The default value for the variable when no value is received.",
			//   "type": "string"
			// }
			Description: "The default value for the variable when no value is received.",
			Type:        types.StringType,
			Required:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description.",
			Type:        types.StringType,
			Optional:    true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the variable was last updated.",
			//   "type": "string"
			// }
			Description: "The time when the variable was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the variable.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the variable.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags associated with this variable.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Tags associated with this variable.",
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
		"variable_type": {
			// Property: VariableType
			// CloudFormation resource type schema:
			// {
			//   "description": "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
			//   "enum": [
			//     "AUTH_CODE",
			//     "AVS",
			//     "BILLING_ADDRESS_L1",
			//     "BILLING_ADDRESS_L2",
			//     "BILLING_CITY",
			//     "BILLING_COUNTRY",
			//     "BILLING_NAME",
			//     "BILLING_PHONE",
			//     "BILLING_STATE",
			//     "BILLING_ZIP",
			//     "CARD_BIN",
			//     "CATEGORICAL",
			//     "CURRENCY_CODE",
			//     "EMAIL_ADDRESS",
			//     "FINGERPRINT",
			//     "FRAUD_LABEL",
			//     "FREE_FORM_TEXT",
			//     "IP_ADDRESS",
			//     "NUMERIC",
			//     "ORDER_ID",
			//     "PAYMENT_TYPE",
			//     "PHONE_NUMBER",
			//     "PRICE",
			//     "PRODUCT_CATEGORY",
			//     "SHIPPING_ADDRESS_L1",
			//     "SHIPPING_ADDRESS_L2",
			//     "SHIPPING_CITY",
			//     "SHIPPING_COUNTRY",
			//     "SHIPPING_NAME",
			//     "SHIPPING_PHONE",
			//     "SHIPPING_STATE",
			//     "SHIPPING_ZIP",
			//     "USERAGENT"
			//   ],
			//   "type": "string"
			// }
			Description: "The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A resource schema for a Variable in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Variable").WithTerraformTypeName("aws_frauddetector_variable").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_frauddetector_variable", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
