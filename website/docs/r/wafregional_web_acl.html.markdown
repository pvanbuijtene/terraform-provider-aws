---
layout: "aws"
page_title: "AWS: aws_wafregional_web_acl"
sidebar_current: "docs-aws-resource-wafregional-web-acl"
description: |-
  Provides a AWS WAF Regional web access control group (ACL) resource for use with ALB.
---

# aws\_wafregional\_web\_acl

Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.

## Example Usage

```hcl
resource "aws_wafregional_ipset" "ipset" {
  name = "tfIPSet"
  
  ip_set_descriptors {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_wafregional_rule" "wafrule" {
  depends_on  = ["aws_wafregional_ipset.ipset"]
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"
  
  predicates {
    data_id = "${aws_wafregional_ipset.ipset.id}"
    negated = false
    type    = "IPMatch"
  }
}

resource "aws_wafregional_web_acl" "wafacl" {
  depends_on  = ["aws_wafregional_ipset.ipset", "aws_wafregional_rule.wafrule"]
  name        = "tfWebACL"
  metric_name = "tfWebACL"
  
  default_action {
    type = "ALLOW"
  }
  
  rules {
    action {
       type = "BLOCK"
    }
    
    priority = 1
    rule_id  = "${aws_wafregional_rule.wafrule.id}"
  }
}
```

## Argument Reference

The following arguments are supported:

* `default_action` - (Required) The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
* `metric_name` - (Required) The name or description for the Amazon CloudWatch metric of this web ACL.
* `name` - (Required) The name or description of the web ACL.
* `rules` - (Required) The rules to associate with the web ACL and the settings for each rule.

## Nested Blocks

### `default_action`

#### Arguments

* `type` - (Required) Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule.
  e.g. `ALLOW`, `BLOCK` or `COUNT`

### `rules`

See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_ActivatedRule.html) for all details and supported values.

#### Arguments

* `action` - (Required) The action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
  e.g. `ALLOW`, `BLOCK` or `COUNT`
* `priority` - (Required) Specifies the order in which the rules in a WebACL are evaluated.
  Rules with a lower value are evaluated before rules with a higher value.
* `rule_id` - (Required) ID of the associated [rule](/docs/providers/aws/r/wafregional_rule.html)
* `type` - (Optional) The rule type, either `REGULAR`, as defined by [Rule](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_Rule.html), or `RATE_BASED`, as defined by [RateBasedRule](http://docs.aws.amazon.com/waf/latest/APIReference/API_RateBasedRule.html). The default is REGULAR. If you add a RATE_BASED rule, you need to set `type` as `RATE_BASED`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the WAF Regional WebACL.
