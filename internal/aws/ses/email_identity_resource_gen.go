// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ses_email_identity", emailIdentityResource)
}

// emailIdentityResource returns the Terraform awscc_ses_email_identity resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::EmailIdentity resource.
func emailIdentityResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConfigurationSetAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Used to associate a configuration set with an email identity.",
		//	  "properties": {
		//	    "ConfigurationSetName": {
		//	      "description": "The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration_set_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConfigurationSetName
				"configuration_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Used to associate a configuration set with an email identity.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Used to enable or disable DKIM authentication for an email identity.",
		//	  "properties": {
		//	    "SigningEnabled": {
		//	      "description": "Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dkim_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SigningEnabled
				"signing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Used to enable or disable DKIM authentication for an email identity.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenName1
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_name_1": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenName2
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_name_2": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenName3
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_name_3": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenValue1
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_value_1": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenValue2
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_value_2": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimDNSTokenValue3
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dkim_dns_token_value_3": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DkimSigningAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.",
		//	  "properties": {
		//	    "DomainSigningPrivateKey": {
		//	      "description": "[Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.",
		//	      "type": "string"
		//	    },
		//	    "DomainSigningSelector": {
		//	      "description": "[Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.",
		//	      "type": "string"
		//	    },
		//	    "NextSigningKeyLength": {
		//	      "description": "[Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.",
		//	      "pattern": "RSA_1024_BIT|RSA_2048_BIT",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dkim_signing_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DomainSigningPrivateKey
				"domain_signing_private_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "[Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// DomainSigningPrivateKey is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: DomainSigningSelector
				"domain_signing_selector": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "[Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// DomainSigningSelector is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: NextSigningKeyLength
				"next_signing_key_length": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "[Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("RSA_1024_BIT|RSA_2048_BIT"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EmailIdentity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The email address or domain to verify.",
		//	  "type": "string"
		//	}
		"email_identity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The email address or domain to verify.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FeedbackAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Used to enable or disable feedback forwarding for an identity.",
		//	  "properties": {
		//	    "EmailForwardingEnabled": {
		//	      "description": "If the value is true, you receive email notifications when bounce or complaint events occur",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"feedback_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EmailForwardingEnabled
				"email_forwarding_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If the value is true, you receive email notifications when bounce or complaint events occur",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Used to enable or disable feedback forwarding for an identity.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MailFromAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
		//	  "properties": {
		//	    "BehaviorOnMxFailure": {
		//	      "description": "The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.",
		//	      "pattern": "USE_DEFAULT_VALUE|REJECT_MESSAGE",
		//	      "type": "string"
		//	    },
		//	    "MailFromDomain": {
		//	      "description": "The custom MAIL FROM domain that you want the verified identity to use",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"mail_from_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BehaviorOnMxFailure
				"behavior_on_mx_failure": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("USE_DEFAULT_VALUE|REJECT_MESSAGE"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MailFromDomain
				"mail_from_domain": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The custom MAIL FROM domain that you want the verified identity to use",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::SES::EmailIdentity",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::EmailIdentity").WithTerraformTypeName("awscc_ses_email_identity")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"behavior_on_mx_failure":       "BehaviorOnMxFailure",
		"configuration_set_attributes": "ConfigurationSetAttributes",
		"configuration_set_name":       "ConfigurationSetName",
		"dkim_attributes":              "DkimAttributes",
		"dkim_dns_token_name_1":        "DkimDNSTokenName1",
		"dkim_dns_token_name_2":        "DkimDNSTokenName2",
		"dkim_dns_token_name_3":        "DkimDNSTokenName3",
		"dkim_dns_token_value_1":       "DkimDNSTokenValue1",
		"dkim_dns_token_value_2":       "DkimDNSTokenValue2",
		"dkim_dns_token_value_3":       "DkimDNSTokenValue3",
		"dkim_signing_attributes":      "DkimSigningAttributes",
		"domain_signing_private_key":   "DomainSigningPrivateKey",
		"domain_signing_selector":      "DomainSigningSelector",
		"email_forwarding_enabled":     "EmailForwardingEnabled",
		"email_identity":               "EmailIdentity",
		"feedback_attributes":          "FeedbackAttributes",
		"mail_from_attributes":         "MailFromAttributes",
		"mail_from_domain":             "MailFromDomain",
		"next_signing_key_length":      "NextSigningKeyLength",
		"signing_enabled":              "SigningEnabled",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DkimSigningAttributes/DomainSigningSelector",
		"/properties/DkimSigningAttributes/DomainSigningPrivateKey",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
