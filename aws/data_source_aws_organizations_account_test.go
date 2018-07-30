package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAWSOrganizationsAccount_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsOrganizationsAccountConfigBasic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.aws_organizations_account.current", "arn"),
					resource.TestCheckResourceAttrSet("data.aws_organizations_account.current", "email"),
					resource.TestCheckResourceAttrSet("data.aws_organizations_account.current", "id"),
					resource.TestCheckResourceAttrSet("data.aws_organizations_account.current", "joined_method"),
					resource.TestCheckResourceAttrSet("data.aws_organizations_account.current", "status"),
				),
			},
		},
	})
}

const testAccCheckAwsOrganizationsAccountConfigBasic = `
data "aws_caller_identity" "current" { }

data "aws_organizations_account" "current" {
	account_id = "${data.aws_caller_identity.current.account_id}"
}
`
