package skyvern_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-skyvern"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "EXAMPLE-DO-NOT-USE"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *skyvern.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := skyvern.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", testAPIKey)

			return nil
		},
	)

	client, err := skyvern.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestGetVersion(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/version", r.URL.Path)
		assert.Equal(t, testAPIKey, r.Header.Get("x-api-key"))

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"version":"example-build"}`))
		assert.NoError(t, err)
	})

	res, err := client.GetVersionV1VersionGetWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	assert.Equal(t, "example-build", (*res.JSON200)["version"])
}

func TestCreateBrowserSession(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/browser_sessions", r.URL.Path)
		assert.Equal(t, testAPIKey, r.Header.Get("x-api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var body skyvern.CreateBrowserSessionRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, new(60), body.Timeout)

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{
			"browser_session_id":"pbs_example",
			"organization_id":"org_example",
			"created_at":"2000-01-01T00:00:00Z",
			"modified_at":"2000-01-01T00:00:00Z",
			"timeout":60,"browser_address":null,"warning":null
		}`))
		assert.NoError(t, err)
	})

	res, err := client.CreateBrowserSessionV1BrowserSessionsPostWithResponse(
		context.Background(),
		&skyvern.CreateBrowserSessionV1BrowserSessionsPostParams{},
		skyvern.CreateBrowserSessionRequest{Timeout: new(60)},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	assert.Equal(t, "pbs_example", res.JSON200.BrowserSessionId)
	assert.Equal(t, new(60), res.JSON200.Timeout)
	assert.Nil(t, res.JSON200.BrowserAddress)
	assert.Nil(t, res.JSON200.Warning)
}

func TestUnauthorized(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Empty(t, r.Header.Get("x-api-key"))
			w.WriteHeader(http.StatusForbidden)
		},
	))
	t.Cleanup(srv.Close)

	client, err := skyvern.NewClientWithResponses(srv.URL)
	require.NoError(t, err)

	res, err := client.GetBrowserSessionsV1BrowserSessionsGetWithResponse(
		context.Background(),
		&skyvern.GetBrowserSessionsV1BrowserSessionsGetParams{},
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, res.StatusCode())
	assert.Nil(t, res.JSON200)
}

func TestActionSchema(t *testing.T) {
	t.Parallel()

	var action skyvern.ClickAction

	err := json.Unmarshal([]byte(`{
		"type":"ai_click","selector":null,"timeout":1234.567890123
	}`), &action)
	require.NoError(t, err)
	require.NotNil(t, action.Type)
	require.NotNil(t, action.Timeout)
	assert.Equal(t, skyvern.ClickActionType("ai_click"), *action.Type)
	assert.Equal(t, 1234.567890123, *action.Timeout)
	assert.Nil(t, action.Selector)
}
