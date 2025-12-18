// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package davescloudserver_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/internal/testutil"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/option"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/shared"
)

func TestStoreOrderNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Store.Orders.New(context.TODO(), davescloudserver.StoreOrderNewParams{
		Order: shared.OrderParam{
			ID:       davescloudserver.Int(10),
			Complete: davescloudserver.Bool(true),
			PetID:    davescloudserver.Int(198772),
			Quantity: davescloudserver.Int(7),
			ShipDate: davescloudserver.Time(time.Now()),
			Status:   shared.OrderStatusApproved,
		},
	})
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Store.Orders.Get(context.TODO(), 0)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.Store.Orders.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
