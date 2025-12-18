// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package davescloudserver_test

import (
	"context"
	"os"
	"testing"

	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/internal/testutil"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := davescloudserver.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	order, err := client.Store.Orders.New(context.TODO(), davescloudserver.StoreOrderNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", order.ID)
}
