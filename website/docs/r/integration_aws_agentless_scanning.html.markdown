---
subcategory: "Cloud Account Integrations"
layout: "lacework"
page_title: "Lacework: lacework_integration_aws_agentless_scanning"
description: |-
  Create and manage AWS Agentless Scanning integration
---

# lacework\_integration\_aws\_agentless\_scanning

Use this resource to configure an AWS Agentless Scanning integration.

## Example Usage

```hcl
resource "lacework_integration_aws_agentless_scanning" "account_abc" {
  name                      = "account ABC"
  scan_frequency            = 24
  query_text                = var.query_text
  scan_containers           = true
  scan_host_vulnerabilities = true
  account_id = "0123456789"
  bucket_arn = "arn:aws:s3:::bucket-arn"
	credentials { 
	  role_arn = "arn:aws:iam::0123456789:role/iam-123"
	  external_id = "0123456789"
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The AWS Agentless Scanning integration name.
* `scan_frequency` - (Required) How often, in hours, the scan will run.
* `query_text` - (Optional) The lql query.
* `scan_containers` - (Optional) Whether to includes scanning for containers.
* `scan_host_vulnerabilities` - (Optional) Whether to includes scanning for host vulnerabilities.
* `account_id` - (Optional) The aws account id.
* `bucket_arn` - (Optional) The bucket arn.
* `credentials` - (Optional) The credentials needed by the integration. See [Credentials](#credentials) below for details.
* `enabled` - (Optional) The state of the external integration. Defaults to `true`.
* `retries` - (Optional) The number of attempts to create the external integration. Defaults to `5`.

### Credentials

  `credentials` supports the following arguments:

* `role_arn` - (Optional) The role arn.
* `external_id` - (Optional) The external id.

## Import

A Lacework AWS Agentless Scanning integration can be imported using a `INT_GUID`, e.g.

```
$ terraform import lacework_integration_aws_agentless_scanning.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5
```
-> **Note:** To retrieve the `INT_GUID` from existing integrations in your account, use the
	Lacework CLI command `lacework cloud-accounts list`. To install this tool follow
	[this documentation](https://docs.lacework.com/cli/).
