package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceAwsWafv2WebAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsWafv2WebAclCreate,
		Read:   resourceAwsWafv2WebAclRead,
		// Update: resourceAwsWafWebAclUpdate,
		Delete: resourceAwsWafv2WebAclDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					wafv2.ScopeCloudfront,
					wafv2.ScopeRegional,
				}, false),
			},

			"default_action": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				ForceNew: true, // REMOVE THIS
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							ForceNew: true, // REMOVE THIS
							Elem: &schema.Resource{
								//Schema: map[string]*schema.Schema{
								//	"cloudwatch_metrics_enabled": {
								//		Type:     schema.TypeBool,
								//		Required: true,
								//	},
								//	"metric_name": {
								//		Type:     schema.TypeString,
								//		Required: true,
								//	},
								//	"sampled_requests_enabled": {
								//		Type:     schema.TypeBool,
								//		Required: true,
								//	},
								//},
							},
						},
					},
				},
			},

			"visibility_config": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				ForceNew: true, // REMOVE THIS
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_metrics_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"metric_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"sampled_requests_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			// "metric_name": {
			// 	Type:         schema.TypeString,
			// 	Required:     true,
			// 	ForceNew:     true,
			// 	ValidateFunc: validateWafMetricName,
			// },
			// "logging_configuration": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	MaxItems: 1,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"log_destination": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"redacted_fields": {
			// 				Type:     schema.TypeList,
			// 				Optional: true,
			// 				MaxItems: 1,
			// 				Elem: &schema.Resource{
			// 					Schema: map[string]*schema.Schema{
			// 						"field_to_match": {
			// 							Type:     schema.TypeSet,
			// 							Required: true,
			// 							Elem: &schema.Resource{
			// 								Schema: map[string]*schema.Schema{
			// 									"data": {
			// 										Type:     schema.TypeString,
			// 										Optional: true,
			// 									},
			// 									"type": {
			// 										Type:     schema.TypeString,
			// 										Required: true,
			// 									},
			// 								},
			// 							},
			// 						},
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
			// "rules": {
			// 	Type:     schema.TypeSet,
			// 	Optional: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"action": {
			// 				Type:     schema.TypeList,
			// 				Optional: true,
			// 				MaxItems: 1,
			// 				Elem: &schema.Resource{
			// 					Schema: map[string]*schema.Schema{
			// 						"type": {
			// 							Type:     schema.TypeString,
			// 							Required: true,
			// 						},
			// 					},
			// 				},
			// 			},
			// 			"override_action": {
			// 				Type:     schema.TypeList,
			// 				Optional: true,
			// 				MaxItems: 1,
			// 				Elem: &schema.Resource{
			// 					Schema: map[string]*schema.Schema{
			// 						"type": {
			// 							Type:     schema.TypeString,
			// 							Required: true,
			// 						},
			// 					},
			// 				},
			// 			},
			// 			"priority": {
			// 				Type:     schema.TypeInt,
			// 				Required: true,
			// 			},
			// 			"type": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 				Default:  waf.WafRuleTypeRegular,
			// 				ValidateFunc: validation.StringInSlice([]string{
			// 					waf.WafRuleTypeRegular,
			// 					waf.WafRuleTypeRateBased,
			// 					waf.WafRuleTypeGroup,
			// 				}, false),
			// 			},
			// 			"rule_id": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 		},
			// 	},
			// },
			// "tags": tagsSchema(),
		},
	}
}

func resourceAwsWafv2WebAclCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafv2conn
	//tags := keyvaluetags.New(d.Get("tags").(map[string]interface{})).IgnoreAws().WafTags()
	//
	//wr := newWafRetryer(conn)
	//out, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
	params := &wafv2.CreateWebACLInput{
		Name:  aws.String(d.Get("name").(string)),
		Scope: aws.String(d.Get("scope").(string)),
		//DefaultAction: expandWafAction(d.Get("default_action").(*schema.Set).List()),
		VisibilityConfig: expandWafv2VisibilityConfig(d.Get("visibility_config").(*schema.Set).List()),
	}

	//if len(tags) > 0 {
	//	params.Tags = tags
	//}

	resp, err := conn.CreateWebACL(params)
	//})
	if err != nil {
		return err
	}

	//resp := out.(*wafv2.CreateWebACLOutput)
	d.SetId(*resp.Summary.Id)
	//
	//arn := arn.ARN{
	//	Partition: meta.(*AWSClient).partition,
	//	Service:   "waf",
	//	AccountID: meta.(*AWSClient).accountid,
	//	Resource:  fmt.Sprintf("webacl/%s", d.Id()),
	//}.String()
	//
	//loggingConfiguration := d.Get("logging_configuration").([]interface{})
	//if len(loggingConfiguration) == 1 {
	//	input := &waf.PutLoggingConfigurationInput{
	//		LoggingConfiguration: expandWAFLoggingConfiguration(loggingConfiguration, arn),
	//	}
	//
	//	log.Printf("[DEBUG] Updating WAF Web ACL (%s) Logging Configuration: %s", d.Id(), input)
	//	if _, err := conn.PutLoggingConfiguration(input); err != nil {
	//		return fmt.Errorf("error updating WAF Web ACL (%s) Logging Configuration: %s", d.Id(), err)
	//	}
	//}
	//
	//rules := d.Get("rules").(*schema.Set).List()
	//if len(rules) > 0 {
	//	wr := newWafRetryer(conn)
	//	_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
	//		req := &waf.UpdateWebACLInput{
	//			ChangeToken:   token,
	//			DefaultAction: expandWafAction(d.Get("default_action").(*schema.Set).List()),
	//			Updates:       diffWafWebAclRules([]interface{}{}, rules),
	//			WebACLId:      aws.String(d.Id()),
	//		}
	//		return conn.UpdateWebACL(req)
	//	})
	//	if err != nil {
	//		return fmt.Errorf("Error Updating WAF ACL: %s", err)
	//	}
	//}

	return resourceAwsWafv2WebAclRead(d, meta)
}

func resourceAwsWafv2WebAclRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafv2conn
	params := &wafv2.GetWebACLInput{
		Id:    aws.String(d.Id()),
		Name:  aws.String(d.Get("name").(string)),
		Scope: aws.String(d.Get("scope").(string)),
	}

	resp, err := conn.GetWebACL(params)
	if err != nil {
		if isAWSErr(err, waf.ErrCodeNonexistentItemException, "") {
			log.Printf("[WARN] WAFV2 ACL (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	if resp == nil || resp.WebACL == nil {
		log.Printf("[WARN] WAFV2 ACL (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("arn", resp.WebACL.ARN)

	//arn := *resp.WebACL.ARN

	// 	if err := d.Set("default_action", flattenWafAction(resp.WebACL.DefaultAction)); err != nil {
	// 		return fmt.Errorf("error setting default_action: %s", err)
	// 	}
	d.Set("name", resp.WebACL.Name)
	// 	d.Set("metric_name", resp.WebACL.MetricName)

	// 	tags, err := keyvaluetags.WafListTags(conn, arn)
	// 	if err != nil {
	// 		return fmt.Errorf("error listing tags for WAF ACL (%s): %s", arn, err)
	// 	}
	// 	if err := d.Set("tags", tags.IgnoreAws().Map()); err != nil {
	// 		return fmt.Errorf("error setting tags: %s", err)
	// 	}

	// 	if err := d.Set("rules", flattenWafWebAclRules(resp.WebACL.Rules)); err != nil {
	// 		return fmt.Errorf("error setting rules: %s", err)
	// 	}

	// 	getLoggingConfigurationInput := &waf.GetLoggingConfigurationInput{
	// 		ResourceArn: aws.String(d.Get("arn").(string)),
	// 	}
	// 	loggingConfiguration := []interface{}{}

	// 	log.Printf("[DEBUG] Getting WAF Web ACL (%s) Logging Configuration: %s", d.Id(), getLoggingConfigurationInput)
	// 	getLoggingConfigurationOutput, err := conn.GetLoggingConfiguration(getLoggingConfigurationInput)

	// 	if err != nil && !isAWSErr(err, waf.ErrCodeNonexistentItemException, "") {
	// 		return fmt.Errorf("error getting WAF Web ACL (%s) Logging Configuration: %s", d.Id(), err)
	// 	}

	// 	if getLoggingConfigurationOutput != nil {
	// 		loggingConfiguration = flattenWAFLoggingConfiguration(getLoggingConfigurationOutput.LoggingConfiguration)
	// 	}

	// 	if err := d.Set("logging_configuration", loggingConfiguration); err != nil {
	// 		return fmt.Errorf("error setting logging_configuration: %s", err)
	// 	}

	return nil
}

// func resourceAwsWafWebAclUpdate(d *schema.ResourceData, meta interface{}) error {
// 	conn := meta.(*AWSClient).wafconn

// 	if d.HasChange("default_action") || d.HasChange("rules") {
// 		o, n := d.GetChange("rules")
// 		oldR, newR := o.(*schema.Set).List(), n.(*schema.Set).List()

// 		wr := newWafRetryer(conn)
// 		_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
// 			req := &waf.UpdateWebACLInput{
// 				ChangeToken:   token,
// 				DefaultAction: expandWafAction(d.Get("default_action").(*schema.Set).List()),
// 				Updates:       diffWafWebAclRules(oldR, newR),
// 				WebACLId:      aws.String(d.Id()),
// 			}
// 			return conn.UpdateWebACL(req)
// 		})
// 		if err != nil {
// 			return fmt.Errorf("Error Updating WAF ACL: %s", err)
// 		}
// 	}

// 	if d.HasChange("logging_configuration") {
// 		loggingConfiguration := d.Get("logging_configuration").([]interface{})

// 		if len(loggingConfiguration) == 1 {
// 			input := &waf.PutLoggingConfigurationInput{
// 				LoggingConfiguration: expandWAFLoggingConfiguration(loggingConfiguration, d.Get("arn").(string)),
// 			}

// 			log.Printf("[DEBUG] Updating WAF Web ACL (%s) Logging Configuration: %s", d.Id(), input)
// 			if _, err := conn.PutLoggingConfiguration(input); err != nil {
// 				return fmt.Errorf("error updating WAF Web ACL (%s) Logging Configuration: %s", d.Id(), err)
// 			}
// 		} else {
// 			input := &waf.DeleteLoggingConfigurationInput{
// 				ResourceArn: aws.String(d.Get("arn").(string)),
// 			}

// 			log.Printf("[DEBUG] Deleting WAF Web ACL (%s) Logging Configuration: %s", d.Id(), input)
// 			if _, err := conn.DeleteLoggingConfiguration(input); err != nil {
// 				return fmt.Errorf("error deleting WAF Web ACL (%s) Logging Configuration: %s", d.Id(), err)
// 			}
// 		}

// 	}

// 	if d.HasChange("tags") {
// 		o, n := d.GetChange("tags")

// 		if err := keyvaluetags.WafUpdateTags(conn, d.Get("arn").(string), o, n); err != nil {
// 			return fmt.Errorf("error updating tags: %s", err)
// 		}
// 	}

// 	return resourceAwsWafWebAclRead(d, meta)
// }

func resourceAwsWafv2WebAclDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).wafv2conn

	//// First, need to delete all rules
	//rules := d.Get("rules").(*schema.Set).List()
	//if len(rules) > 0 {
	//	wr := newWafRetryer(conn)
	//	_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
	//		req := &waf.UpdateWebACLInput{
	//			ChangeToken:   token,
	//			DefaultAction: expandWafAction(d.Get("default_action").(*schema.Set).List()),
	//			Updates:       diffWafWebAclRules(rules, []interface{}{}),
	//			WebACLId:      aws.String(d.Id()),
	//		}
	//		return conn.UpdateWebACL(req)
	//	})
	//	if err != nil {
	//		return fmt.Errorf("Error Removing WAF Regional ACL Rules: %s", err)
	//	}
	//}

	//wr := newWafRetryer(conn)
	//_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {

	params := &wafv2.GetWebACLInput{
		Id:    aws.String(d.Id()),
		Name:  aws.String(d.Get("name").(string)),
		Scope: aws.String(d.Get("scope").(string)),
	}

	resp, err := conn.GetWebACL(params)
	if err != nil {
		if isAWSErr(err, wafv2.ErrCodeWAFNonexistentItemException, "") {
			log.Printf("[WARN] WAFV2 ACL (%s) not found, removing from state", d.Id())
			return nil
		}

		return err
	}

	if resp == nil || resp.WebACL == nil {
		log.Printf("[WARN] WAFV2 ACL (%s) not found, removing from state", d.Id())
		return nil
	}

	req := &wafv2.DeleteWebACLInput{
		LockToken: resp.LockToken,
		Id:        aws.String(d.Id()),
	}

	log.Printf("[INFO] Deleting WAFV2 ACL")
	conn.DeleteWebACL(req)
	//})
	if err != nil {
		return fmt.Errorf("Error Deleting WAFV2 ACL: %s", err)
	}
	return nil
}

// func expandWAFLoggingConfiguration(l []interface{}, resourceARN string) *waf.LoggingConfiguration {
// 	if len(l) == 0 || l[0] == nil {
// 		return nil
// 	}

// 	m := l[0].(map[string]interface{})

// 	loggingConfiguration := &waf.LoggingConfiguration{
// 		LogDestinationConfigs: []*string{
// 			aws.String(m["log_destination"].(string)),
// 		},
// 		RedactedFields: expandWAFRedactedFields(m["redacted_fields"].([]interface{})),
// 		ResourceArn:    aws.String(resourceARN),
// 	}

// 	return loggingConfiguration
// }

func expandWafv2VisibilityConfig(l []interface{}) *wafv2.VisibilityConfig {
	m := l[0].(map[string]interface{})

	loggingConfiguration := &wafv2.VisibilityConfig{
		CloudWatchMetricsEnabled: aws.Bool(m["cloudwatch_metrics_enabled"].(bool)),
		MetricName:               aws.String(m["metric_name"].(string)),
		SampledRequestsEnabled:   aws.Bool(m["sampled_requests_enabled"].(bool)),
	}

	return loggingConfiguration
}

// func expandWAFRedactedFields(l []interface{}) []*waf.FieldToMatch {
// 	if len(l) == 0 || l[0] == nil {
// 		return nil
// 	}

// 	m := l[0].(map[string]interface{})

// 	if m["field_to_match"] == nil {
// 		return nil
// 	}

// 	redactedFields := make([]*waf.FieldToMatch, 0)

// 	for _, fieldToMatch := range m["field_to_match"].(*schema.Set).List() {
// 		if fieldToMatch == nil {
// 			continue
// 		}

// 		redactedFields = append(redactedFields, expandFieldToMatch(fieldToMatch.(map[string]interface{})))
// 	}

// 	return redactedFields
// }

// func flattenWAFLoggingConfiguration(loggingConfiguration *waf.LoggingConfiguration) []interface{} {
// 	if loggingConfiguration == nil {
// 		return []interface{}{}
// 	}

// 	m := map[string]interface{}{
// 		"log_destination": "",
// 		"redacted_fields": flattenWAFRedactedFields(loggingConfiguration.RedactedFields),
// 	}

// 	if len(loggingConfiguration.LogDestinationConfigs) > 0 {
// 		m["log_destination"] = aws.StringValue(loggingConfiguration.LogDestinationConfigs[0])
// 	}

// 	return []interface{}{m}
// }

// func flattenWAFRedactedFields(fieldToMatches []*waf.FieldToMatch) []interface{} {
// 	if len(fieldToMatches) == 0 {
// 		return []interface{}{}
// 	}

// 	fieldToMatchResource := &schema.Resource{
// 		Schema: map[string]*schema.Schema{
// 			"data": {
// 				Type:     schema.TypeString,
// 				Optional: true,
// 			},
// 			"type": {
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 		},
// 	}
// 	l := make([]interface{}, len(fieldToMatches))

// 	for i, fieldToMatch := range fieldToMatches {
// 		l[i] = flattenFieldToMatch(fieldToMatch)[0]
// 	}

// 	m := map[string]interface{}{
// 		"field_to_match": schema.NewSet(schema.HashResource(fieldToMatchResource), l),
// 	}

// 	return []interface{}{m}
// }
