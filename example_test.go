package skyvern_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-skyvern"
)

// This README example requires an API key and only compiles during tests.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := skyvern.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", os.Getenv("SKYVERN_API_KEY"))

			return nil
		},
	)

	client, err := skyvern.NewClientWithResponses(
		"https://api.skyvern.com",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.GetVersionV1VersionGetWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
