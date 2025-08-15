package akp

import (
	"fmt"
	"os"
	"testing"

	"github.com/akuity/terraform-provider-akp/akp/testhelpers"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	orgName        = ""
	providerConfig = ""
)

func TestMain(m *testing.M) {
	if v := os.Getenv("AKUITY_PROVIDER_ORG_NAME"); v == "" {
		orgName = "terraform-provider-acceptance-test"
	} else {
		orgName = v
	}

	providerConfig = fmt.Sprintf(`
provider "akp" {
	org_name = "%s"
}
`, orgName)

	testhelpers.TestMain(m)
}

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"akp": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("AKUITY_API_KEY_ID"); v == "" {
		t.Fatal("AKUITY_API_KEY_ID must be set for acceptance tests")
	}
	if v := os.Getenv("AKUITY_API_KEY_SECRET"); v == "" {
		t.Fatal("AKUITY_API_KEY_SECRET must be set for acceptance tests")
	}
}
