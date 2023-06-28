---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_refresh_schedule Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::QuickSight::RefreshSchedule
---

# awscc_quicksight_refresh_schedule (Data Source)

Data Source schema for AWS::QuickSight::RefreshSchedule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) of the data source.</p>
- `aws_account_id` (String)
- `data_set_id` (String)
- `schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule))

<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Read-Only:

- `refresh_type` (String)
- `schedule_frequency` (Attributes) <p>Information about the schedule frequency.</p> (see [below for nested schema](#nestedatt--schedule--schedule_frequency))
- `schedule_id` (String) <p>An unique identifier for the refresh schedule.</p>
- `start_after_date_time` (String) <p>The date time after which refresh is to be scheduled</p>

<a id="nestedatt--schedule--schedule_frequency"></a>
### Nested Schema for `schedule.schedule_frequency`

Read-Only:

- `interval` (String)
- `refresh_on_day` (Attributes) <p>The day scheduled for refresh.</p> (see [below for nested schema](#nestedatt--schedule--schedule_frequency--refresh_on_day))
- `time_of_the_day` (String) <p>The time of the day for scheduled refresh.</p>
- `time_zone` (String) <p>The timezone for scheduled refresh.</p>

<a id="nestedatt--schedule--schedule_frequency--refresh_on_day"></a>
### Nested Schema for `schedule.schedule_frequency.refresh_on_day`

Read-Only:

- `day_of_month` (String) <p>The Day Of Month for scheduled refresh.</p>
- `day_of_week` (String)