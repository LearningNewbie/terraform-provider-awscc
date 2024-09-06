// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package events

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_events_connection", connectionDataSource)
}

// connectionDataSource returns the Terraform awscc_events_connection data source.
// This Terraform data source corresponds to the CloudFormation AWS::Events::Connection resource.
func connectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The arn of the connection resource.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The arn of the connection resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "BasicAuthParameters"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "OAuthParameters"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "ApiKeyAuthParameters"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "ApiKeyAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ApiKeyName": {
		//	          "type": "string"
		//	        },
		//	        "ApiKeyValue": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ApiKeyName",
		//	        "ApiKeyValue"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "BasicAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Password": {
		//	          "type": "string"
		//	        },
		//	        "Username": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Username",
		//	        "Password"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "InvocationHttpParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BodyParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "HeaderParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "QueryStringParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AuthorizationEndpoint": {
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "ClientParameters": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ClientID": {
		//	              "type": "string"
		//	            },
		//	            "ClientSecret": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ClientID",
		//	            "ClientSecret"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "HttpMethod": {
		//	          "enum": [
		//	            "GET",
		//	            "POST",
		//	            "PUT"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "OAuthHttpParameters": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BodyParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "HeaderParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "QueryStringParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ClientParameters",
		//	        "AuthorizationEndpoint",
		//	        "HttpMethod"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApiKeyAuthParameters
				"api_key_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ApiKeyName
						"api_key_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ApiKeyValue
						"api_key_value": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: BasicAuthParameters
				"basic_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Password
						"password": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Username
						"username": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: InvocationHttpParameters
				"invocation_http_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BodyParameters
						"body_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
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
						// Property: HeaderParameters
						"header_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
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
						// Property: QueryStringParameters
						"query_string_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
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
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OAuthParameters
				"o_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AuthorizationEndpoint
						"authorization_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ClientParameters
						"client_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ClientID
								"client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: ClientSecret
								"client_secret": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: HttpMethod
						"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: OAuthHttpParameters
						"o_auth_http_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BodyParameters
								"body_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
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
								// Property: HeaderParameters
								"header_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
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
								// Property: QueryStringParameters
								"query_string_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
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
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "API_KEY",
		//	    "BASIC",
		//	    "OAUTH_CLIENT_CREDENTIALS"
		//	  ],
		//	  "type": "string"
		//	}
		"authorization_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the connection.",
		//	  "maxLength": 512,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the connection.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The arn of the secrets manager secret created in the customer account.",
		//	  "type": "string"
		//	}
		"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The arn of the secrets manager secret created in the customer account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Events::Connection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::Connection").WithTerraformTypeName("awscc_events_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_auth_parameters":    "ApiKeyAuthParameters",
		"api_key_name":               "ApiKeyName",
		"api_key_value":              "ApiKeyValue",
		"arn":                        "Arn",
		"auth_parameters":            "AuthParameters",
		"authorization_endpoint":     "AuthorizationEndpoint",
		"authorization_type":         "AuthorizationType",
		"basic_auth_parameters":      "BasicAuthParameters",
		"body_parameters":            "BodyParameters",
		"client_id":                  "ClientID",
		"client_parameters":          "ClientParameters",
		"client_secret":              "ClientSecret",
		"description":                "Description",
		"header_parameters":          "HeaderParameters",
		"http_method":                "HttpMethod",
		"invocation_http_parameters": "InvocationHttpParameters",
		"is_value_secret":            "IsValueSecret",
		"key":                        "Key",
		"name":                       "Name",
		"o_auth_http_parameters":     "OAuthHttpParameters",
		"o_auth_parameters":          "OAuthParameters",
		"password":                   "Password",
		"query_string_parameters":    "QueryStringParameters",
		"secret_arn":                 "SecretArn",
		"username":                   "Username",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
