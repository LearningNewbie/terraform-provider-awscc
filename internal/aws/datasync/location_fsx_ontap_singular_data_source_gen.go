// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datasync_location_fsx_ontap", locationFSxONTAPDataSource)
}

// locationFSxONTAPDataSource returns the Terraform awscc_datasync_location_fsx_ontap data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataSync::LocationFSxONTAP resource.
func locationFSxONTAPDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: FsxFilesystemArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$",
		//	  "type": "string"
		//	}
		"fsx_filesystem_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the FSx ONTAP file system that was described.",
		//	  "maxLength": 4360,
		//	  "pattern": "^(efs|nfs|s3|smb|hdfs|fsx[a-z0-9-]+)://[a-zA-Z0-9.:/\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the FSx ONTAP file system that was described.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration settings for NFS or SMB protocol.",
		//	  "properties": {
		//	    "NFS": {
		//	      "additionalProperties": false,
		//	      "description": "NFS protocol configuration for FSx ONTAP file system.",
		//	      "properties": {
		//	        "MountOptions": {
		//	          "additionalProperties": false,
		//	          "description": "The NFS mount options that DataSync can use to mount your NFS share.",
		//	          "properties": {
		//	            "Version": {
		//	              "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
		//	              "enum": [
		//	                "AUTOMATIC",
		//	                "NFS3",
		//	                "NFS4_0",
		//	                "NFS4_1"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "MountOptions"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "SMB": {
		//	      "additionalProperties": false,
		//	      "description": "SMB protocol configuration for FSx ONTAP file system.",
		//	      "properties": {
		//	        "Domain": {
		//	          "description": "The name of the Windows domain that the SMB server belongs to.",
		//	          "maxLength": 253,
		//	          "pattern": "^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$",
		//	          "type": "string"
		//	        },
		//	        "MountOptions": {
		//	          "additionalProperties": false,
		//	          "description": "The mount options used by DataSync to access the SMB server.",
		//	          "properties": {
		//	            "Version": {
		//	              "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
		//	              "enum": [
		//	                "AUTOMATIC",
		//	                "SMB2",
		//	                "SMB3"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "Password": {
		//	          "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
		//	          "maxLength": 104,
		//	          "pattern": "^.{0,104}$",
		//	          "type": "string"
		//	        },
		//	        "User": {
		//	          "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
		//	          "maxLength": 104,
		//	          "pattern": "^[^\\x5B\\x5D\\\\/:;|=,+*?]{1,104}$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "User",
		//	        "Password",
		//	        "MountOptions"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"protocol": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NFS
				"nfs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MountOptions
						"mount_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Version
								"version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The specific NFS version that you want DataSync to use to mount your NFS share.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The NFS mount options that DataSync can use to mount your NFS share.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "NFS protocol configuration for FSx ONTAP file system.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SMB
				"smb": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Domain
						"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the Windows domain that the SMB server belongs to.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MountOptions
						"mount_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Version
								"version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The specific SMB version that you want DataSync to use to mount your SMB share.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The mount options used by DataSync to access the SMB server.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Password
						"password": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: User
						"user": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "SMB protocol configuration for FSx ONTAP file system.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration settings for NFS or SMB protocol.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/sg-[a-f0-9]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"security_group_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageVirtualMachineArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
		//	  "maxLength": 162,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:storage-virtual-machine/fs-[0-9a-f]+/svm-[0-9a-f]{17,}$",
		//	  "type": "string"
		//	}
		"storage_virtual_machine_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A subdirectory in the location's path.",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A subdirectory in the location's path.",
			Computed:    true,
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
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataSync::LocationFSxONTAP",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationFSxONTAP").WithTerraformTypeName("awscc_datasync_location_fsx_ontap")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain":                      "Domain",
		"fsx_filesystem_arn":          "FsxFilesystemArn",
		"key":                         "Key",
		"location_arn":                "LocationArn",
		"location_uri":                "LocationUri",
		"mount_options":               "MountOptions",
		"nfs":                         "NFS",
		"password":                    "Password",
		"protocol":                    "Protocol",
		"security_group_arns":         "SecurityGroupArns",
		"smb":                         "SMB",
		"storage_virtual_machine_arn": "StorageVirtualMachineArn",
		"subdirectory":                "Subdirectory",
		"tags":                        "Tags",
		"user":                        "User",
		"value":                       "Value",
		"version":                     "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
