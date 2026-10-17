# go-skyvern

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-skyvern.svg)](https://pkg.go.dev/github.com/psyb0t/go-skyvern)
[![CI](https://github.com/psyb0t/go-skyvern/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-skyvern/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-skyvern/badges/version.svg)](https://github.com/psyb0t/go-skyvern/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-skyvern/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-skyvern/badges/importers.svg)](https://github.com/psyb0t/go-skyvern/blob/badges/importers.md)

Unofficial Go client for the [Skyvern API](https://api.skyvern.com/openapi.json), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Skyvern publishes. I am not affiliated with Skyvern.

The package is `skyvern`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-skyvern
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-skyvern"
)

func main() {
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
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Method names come from the spec's operation IDs. For example, `GetVersionV1VersionGetWithResponse` reads the API version and `RunTaskV1RunTasksPostWithResponse` starts a task.

The full API, with the spec's descriptions as doc comments, is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-skyvern) and in `go doc github.com/psyb0t/go-skyvern` from inside your module.

## Auth

The spec's `ApiKeyAuth` scheme sends your Skyvern API key in the `x-api-key` header. Get the key from [Skyvern settings](https://app.skyvern.com/settings), export it as `SKYVERN_API_KEY`, and attach it with `WithRequestEditorFn` as shown above.

The cloud base URL is `https://api.skyvern.com`. The spec also lists `https://api-staging.skyvern.com` and `http://localhost:8000` for staging and local deployments.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://api.skyvern.com/openapi.json` to `api/openapi.json`.
- `scripts/patch-spec.sh` writes `api/openapi.patched.json`. It converts nullable unions to `nullable`, constants to enums, and examples arrays to a single example, then declares OpenAPI 3.0.3 for oapi-codegen. It also normalizes description punctuation for generated comments.
- oapi-codegen runs with `config.yaml`, including the `number` to `float64` mapping.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release.

## Agent integrations

The [skill](.agents/skills/go-skyvern) teaches a coding agent how to use go-skyvern when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-skyvern@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-skyvern@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-skyvern
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Skyvern's published API specification. The specification itself is not redistributed here. It declares no license or terms of service URL.
