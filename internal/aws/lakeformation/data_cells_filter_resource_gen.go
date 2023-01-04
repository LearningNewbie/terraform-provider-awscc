// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lakeformation_data_cells_filter", dataCellsFilterResource)
}

// dataCellsFilterResource returns the Terraform awscc_lakeformation_data_cells_filter resource.
// This Terraform resource corresponds to the CloudFormation AWS::LakeFormation::DataCellsFilter resource.
func dataCellsFilterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ColumnNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of columns to be included in this Data Cells Filter.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "A string representing a resource's name.",
		//	    "maxLength": 255,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of columns to be included in this Data Cells Filter.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 255),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ColumnWildcard
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required",
		//	  "properties": {
		//	    "ExcludedColumnNames": {
		//	      "description": "A list of column names to be excluded from the Data Cells Filter.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "A string representing a resource's name.",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"column_wildcard": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExcludedColumnNames
				"excluded_column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of column names to be excluded from the Data Cells Filter.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 255),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatabaseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Database that the Table resides in.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Database that the Table resides in.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired name of the Data Cells Filter.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The desired name of the Data Cells Filter.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RowFilter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required",
		//	  "properties": {
		//	    "AllRowsWildcard": {
		//	      "additionalProperties": false,
		//	      "description": "An empty object representing a row wildcard.",
		//	      "type": "object"
		//	    },
		//	    "FilterExpression": {
		//	      "description": "A PartiQL predicate.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"row_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllRowsWildcard
				"all_rows_wildcard": schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "An empty object representing a row wildcard.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FilterExpression
				"filter_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A PartiQL predicate.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableCatalogId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Catalog Id of the Table on which to create a Data Cells Filter.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"table_catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Catalog Id of the Table on which to create a Data Cells Filter.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(12, 12),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Table to create a Data Cells Filter for.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Table to create a Data Cells Filter for.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "A resource schema representing a Lake Formation Data Cells Filter.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::DataCellsFilter").WithTerraformTypeName("awscc_lakeformation_data_cells_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"all_rows_wildcard":     "AllRowsWildcard",
		"column_names":          "ColumnNames",
		"column_wildcard":       "ColumnWildcard",
		"database_name":         "DatabaseName",
		"excluded_column_names": "ExcludedColumnNames",
		"filter_expression":     "FilterExpression",
		"name":                  "Name",
		"row_filter":            "RowFilter",
		"table_catalog_id":      "TableCatalogId",
		"table_name":            "TableName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
