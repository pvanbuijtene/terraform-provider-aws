package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsOrganizationsAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsOrganizationsAccountRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"joined_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsOrganizationsAccountRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).organizationsconn
	accountId := d.Get("account_id").(string)

	describeAccount := &organizations.DescribeAccountInput{
		AccountId: aws.String(accountId),
	}

	log.Printf("[DEBUG] Reading AWS Organizations account: %s", accountId)
	describeResp, err := conn.DescribeAccount(describeAccount)
	if err != nil {
		return errwrap.Wrapf("Error retrieving AWS Organizations account: {{err}}", err)
	}

	// if len(describeResp.LoadBalancers) != 1 {
	// 	return fmt.Errorf("Search returned %d results, please revise so only one is returned", len(describeResp.LoadBalancers))
	// }

	d.SetId(*describeResp.Account.Id)

	return flattenAwsOrganizationsAccount(d, meta, describeResp.Account)
}

func flattenAwsOrganizationsAccount(d *schema.ResourceData, meta interface{}, account *organizations.Account) error {

	d.Set("arn", account.Arn)
	d.Set("email", account.Email)
	d.Set("id", account.Id)
	d.Set("joined_method", account.JoinedMethod)
	d.Set("name", account.Name)
	d.Set("status", account.Status)

	return nil
}
