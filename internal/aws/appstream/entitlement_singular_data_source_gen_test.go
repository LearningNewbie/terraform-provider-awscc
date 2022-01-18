// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAppStreamEntitlementDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AppStream::Entitlement", "awscc_appstream_entitlement", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSAppStreamEntitlementDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AppStream::Entitlement", "awscc_appstream_entitlement", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}