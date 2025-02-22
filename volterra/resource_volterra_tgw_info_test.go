//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestResourceTGWInfo(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testTGWInfo(name),
			},
		},
	})
}

func testTGWInfo(name string) string {

	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  tgw_id = "tgw-0d9bf32a5171e5a49"
		  vpc_id = "vpc-02d0d70860b330393"
		  subnet_ids {
			outside_subnet_id = "subnet-12345678"
			inside_subnet_id = "subnet-12345678"
			workload_subnet_id = "subnet-12345678"
			az = "us-east-2a"
		  }
		}
		`, setTGWInfo, name, name)
}
