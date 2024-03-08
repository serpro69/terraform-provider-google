// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package apphub_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccApphubApplication_applicationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApphubApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApphubApplication_applicationBasicExample(context),
			},
			{
				ResourceName:            "google_apphub_application.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "application_id"},
			},
		},
	})
}

func testAccApphubApplication_applicationBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_apphub_application" "example" {
  location = "us-east1"
  application_id = "tf-test-example-application%{random_suffix}"
  scope {
    type = "REGIONAL"
  }
}
`, context)
}

func TestAccApphubApplication_applicationFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApphubApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApphubApplication_applicationFullExample(context),
			},
			{
				ResourceName:            "google_apphub_application.example2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "application_id"},
			},
		},
	})
}

func testAccApphubApplication_applicationFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_apphub_application" "example2" {
  location = "us-east1"
  application_id = "tf-test-example-application%{random_suffix}"
  display_name = "Application Full%{random_suffix}"
  scope {
    type = "REGIONAL"
  }
  description = "Application for testing%{random_suffix}"
  attributes {
    environment {
      type = "STAGING"
		}
		criticality {  
      type = "MISSION_CRITICAL"
		}
		business_owners {
		  display_name =  "Alice%{random_suffix}"
		  email        =  "alice@google.com%{random_suffix}"
		}
		developer_owners {
		  display_name =  "Bob%{random_suffix}"
		  email        =  "bob@google.com%{random_suffix}"
		}
		operator_owners {
		  display_name =  "Charlie%{random_suffix}"
		  email        =  "charlie@google.com%{random_suffix}"
		}
  }
}
`, context)
}

func testAccCheckApphubApplicationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apphub_application" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ApphubApplication still exists at %s", url)
			}
		}

		return nil
	}
}
