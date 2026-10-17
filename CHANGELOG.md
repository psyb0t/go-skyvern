# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.2.0, 2026-10-17

- `(*Client).UpdateWorkflowV1AgentsWorkflowIdPost` and `(*ClientWithResponses).UpdateWorkflowV1AgentsWorkflowIdPostWithResponse` now take a `WorkflowUpdateRequest` body instead of `WorkflowRequest`. The alias `UpdateWorkflowV1AgentsWorkflowIdPostJSONRequestBody` now points to `WorkflowUpdateRequest`, which accepts a `YamlDefinition` string or a `JsonDefinition` of the new `WorkflowUpdateYAMLRequest` type. Code that builds a `WorkflowRequest` for this call needs to switch to the new type.
- `Workflow` gained a `BuildSource` field of the new type `WorkflowBuildSource`, which records whether a workflow version was written by the UI, the workflow copilot, an MCP client, the CLI, another API caller, or Skyvern itself.
- `CredentialVaultType` gained a new value, `CredentialVaultTypeOnepassword`, for credentials stored in 1Password.
- `WorkflowSchedule` gained a `CreatedBy` field, the user the schedule runs as. It is null for schedules created before owners were recorded.
- `CreateBrowserSessionV1BrowserSessionsPostResponse` gained a `JSON429` field and a `GetJSON429` method for a new rate limit response, with error code `CreateBrowserSessionV1BrowserSessionsPost429JSONResponseBodyErrorCodeTooManyStartingBrowserSessions`.
- Deleting a workflow now parses a 404 response with no body instead of leaving it unhandled.
- The doc comments for workflow run trigger type `mcp` now mention the Skyvern CLI as well as first-party MCP clients.

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `skyvern`, a REST client generated with oapi-codegen from the Skyvern API specification, version 1.0.0, OpenAPI 3.1.0, at https://api.skyvern.com/openapi.json. It covers tasks, agents, runs, browser sessions and profiles, credentials, scripts, schedules, tags, artifacts and audit exports.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, converts unsupported schema forms for oapi-codegen, and regenerates the client. The spec is not committed.
