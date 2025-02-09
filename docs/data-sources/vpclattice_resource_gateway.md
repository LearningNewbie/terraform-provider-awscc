---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_resource_gateway Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::ResourceGateway
---

# awscc_vpclattice_resource_gateway (Data Source)

Data Source schema for AWS::VpcLattice::ResourceGateway



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `ip_address_type` (String)
- `name` (String)
- `resource_gateway_id` (String)
- `security_group_ids` (Set of String) The ID of one or more security groups to associate with the endpoint network interface.
- `subnet_ids` (Set of String) The ID of one or more subnets in which to create an endpoint network interface.
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `vpc_identifier` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
