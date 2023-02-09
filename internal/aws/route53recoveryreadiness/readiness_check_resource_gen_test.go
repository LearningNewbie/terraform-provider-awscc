// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoveryreadiness_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53RecoveryReadinessReadinessCheck_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryReadiness::ReadinessCheck", "awscc_route53recoveryreadiness_readiness_check", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSRoute53RecoveryReadinessReadinessCheck_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryReadiness::ReadinessCheck", "awscc_route53recoveryreadiness_readiness_check", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
