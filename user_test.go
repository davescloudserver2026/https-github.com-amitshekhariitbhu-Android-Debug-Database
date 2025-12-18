// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package davescloudserver_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/internal/testutil"
	"github.com/davescloudserver2026/https-github.com-amitshekhariitbhu-Android-Debug-Database/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.New(context.TODO(), davescloudserver.UserNewParams{
		User: davescloudserver.UserParam{
			ID:         davescloudserver.Int(10),
			Email:      davescloudserver.String("john@email.com"),
			FirstName:  davescloudserver.String("John"),
			LastName:   davescloudserver.String("James"),
			Password:   davescloudserver.String("12345"),
			Phone:      davescloudserver.String("12345"),
			Username:   davescloudserver.String("theUser"),
			UserStatus: davescloudserver.Int(1),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "username")
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Users.Update(
		context.TODO(),
		"username",
		davescloudserver.UserUpdateParams{
			User: davescloudserver.UserParam{
				ID:         davescloudserver.Int(10),
				Email:      davescloudserver.String("john@email.com"),
				FirstName:  davescloudserver.String("John"),
				LastName:   davescloudserver.String("James"),
				Password:   davescloudserver.String("12345"),
				Phone:      davescloudserver.String("12345"),
				Username:   davescloudserver.String("theUser"),
				UserStatus: davescloudserver.Int(1),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.Users.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.NewWithList(context.TODO(), davescloudserver.UserNewWithListParams{
		Items: []davescloudserver.UserParam{{
			ID:         davescloudserver.Int(10),
			Email:      davescloudserver.String("john@email.com"),
			FirstName:  davescloudserver.String("John"),
			LastName:   davescloudserver.String("James"),
			Password:   davescloudserver.String("12345"),
			Phone:      davescloudserver.String("12345"),
			Username:   davescloudserver.String("theUser"),
			UserStatus: davescloudserver.Int(1),
		}},
	})
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), davescloudserver.UserLoginParams{
		Password: davescloudserver.String("password"),
		Username: davescloudserver.String("username"),
	})
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.Users.Logout(context.TODO())
	if err != nil {
		var apierr *davescloudserver.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
