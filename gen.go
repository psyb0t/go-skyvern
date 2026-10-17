// Package skyvern is a Go client for the Skyvern API, generated with
// oapi-codegen from the OpenAPI document Skyvern publishes.
package skyvern

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/patch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.patched.json
