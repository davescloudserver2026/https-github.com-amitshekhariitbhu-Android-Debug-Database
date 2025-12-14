// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package davescloudserver_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/davescloudserver-go"
	"github.com/stainless-sdks/davescloudserver-go/internal/testutil"
	"github.com/stainless-sdks/davescloudserver-go/option"
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
	order, err := client.Store.Orders.New(context.TODO(), davescloudserver.StoreOrderNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", order.ID)
}
