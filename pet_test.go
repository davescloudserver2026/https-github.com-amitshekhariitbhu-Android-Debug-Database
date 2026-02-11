// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package davescloudserver_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/internal/testutil"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), davescloudserver.PetNewParams{
		Pet: davescloudserver.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        davescloudserver.Int(10),
			Category: davescloudserver.CategoryParam{
				ID:   davescloudserver.Int(1),
				Name: davescloudserver.String("Dogs"),
			},
			Status: davescloudserver.PetStatusAvailable,
			Tags: []davescloudserver.PetTagParam{{
				ID:   davescloudserver.Int(0),
				Name: davescloudserver.String("name"),
			}},
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), davescloudserver.PetUpdateParams{
		Pet: davescloudserver.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        davescloudserver.Int(10),
			Category: davescloudserver.CategoryParam{
				ID:   davescloudserver.Int(1),
				Name: davescloudserver.String("Dogs"),
			},
			Status: davescloudserver.PetStatusAvailable,
			Tags: []davescloudserver.PetTagParam{{
				ID:   davescloudserver.Int(0),
				Name: davescloudserver.String("name"),
			}},
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), davescloudserver.PetFindByStatusParams{
		Status: davescloudserver.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), davescloudserver.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		0,
		davescloudserver.PetUpdateByIDParams{
			Name:   davescloudserver.String("name"),
			Status: davescloudserver.String("status"),
		},
	)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		davescloudserver.PetUploadImageParams{
			AdditionalMetadata: davescloudserver.String("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
