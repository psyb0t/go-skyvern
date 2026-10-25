# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.5.1, 2026-10-25

- `AuditEventExportRow` gained an `ImpersonatedBy` field, a `*string`.
- `FileDownloadBlock` and `FileDownloadBlockYAML` gained a `SftpSourcePath` field, a `*string` sent as `sftp_source_path`.

## v0.5.0, 2026-10-24

- Added the run activity export endpoint. `(*Client).ExportRunActivityV1AuditEventsRunsExportGet` and `(*ClientWithResponses).ExportRunActivityV1AuditEventsRunsExportGetWithResponse` send the request, built by `NewExportRunActivityV1AuditEventsRunsExportGetRequest` and parsed by `ParseExportRunActivityV1AuditEventsRunsExportGetResponse`. The result type is `ExportRunActivityV1AuditEventsRunsExportGetResponse`.
- Added `ExportRunActivityV1AuditEventsRunsExportGetParams` with a `Format` of type `ExportRunActivityV1AuditEventsRunsExportGetParamsFormat`, which has the values `ExportRunActivityV1AuditEventsRunsExportGetParamsFormatCsv` and `ExportRunActivityV1AuditEventsRunsExportGetParamsFormatJson`.
- Added `RunActivityExportResponse`, with `Rows` and `Next`, and `RunActivityExportRow`, one row per workflow run.
- `ClientInterface` gained `ExportRunActivityV1AuditEventsRunsExportGet` and `ClientWithResponsesInterface` gained `ExportRunActivityV1AuditEventsRunsExportGetWithResponse`. Custom types that implement either interface, such as test mocks, must add the new method.
- `ExportAuditEventsV1AuditEventsExportGetParams` gained `Action`, `ActorUserId` and `ResourceType` filter fields, all `*string`.
- `Action` gained a `DownloadUrl` field, a `*string` with the origin URL recorded by a download action.

## v0.4.0, 2026-10-23

- `WorkflowCreateYAMLRequest` lost its `RecordingId` field, which the Skyvern API no longer has. Code that sets `RecordingId` must drop it.
- `AuditExportDenialDetail` gained `Action` and `Role` fields, both `*string`.
- The Go pin moves from 1.26.9 to 1.26.10 in `go.mod`, `Dockerfile.dev` and the agent skill setup notes, because govulncheck flagged GO-2026-6809 in `crypto/x509`. The development image also moves to `ca-certificates` 20261020-r0.

## v0.3.1, 2026-10-20

- `GetWorkflowRunsV1WorkflowsRunsGetParams` gained `CreatedAtStart` and `CreatedAtEnd` fields, both `*time.Time`. `CreatedAtStart` includes runs created at or after the timestamp. `CreatedAtEnd` includes only runs created before it. `NewGetWorkflowRunsV1WorkflowsRunsGetRequest` sends them as the `created_at_start` and `created_at_end` query parameters.
- A free-text `SearchKey` on the workflow runs listing now covers only the 30 days before `CreatedAtEnd` (or before now) unless `CreatedAtStart` is set. A complete browser profile ID or browser session ID is not limited. Set `CreatedAtStart` to search older runs.
- The doc comments for `GetWorkflowRunsV1WorkflowsRunsGet` and its `WithResponse` variant describe the search window and the new fields.

## v0.3.0, 2026-10-19

- `(*Client).UpdateWorkflowV1AgentsWorkflowIdPost` and `(*ClientWithResponses).UpdateWorkflowV1AgentsWorkflowIdPostWithResponse` take a `WorkflowRequest` body again instead of `WorkflowUpdateRequest`. The alias `UpdateWorkflowV1AgentsWorkflowIdPostJSONRequestBody` points back to `WorkflowRequest`. The types `WorkflowUpdateRequest`, `WorkflowUpdateYAMLRequest`, `WorkflowUpdateYAMLRequestProxyLocation2` and `WorkflowUpdateYAMLRequest_ProxyLocation`, added in v0.2.0, are removed. Code that builds a `WorkflowUpdateRequest` for this call needs to switch back to `WorkflowRequest`.
- `BrowserProfile` gained a `LinkedCredentialId` field, the ID of the linked credential.
- `CredentialResponse` gained an `ItemId` field and a `VaultId` field, the 1Password item and vault IDs for a linked credential, and an `UpstreamMissingAt` field, the time the linked 1Password item was last found missing upstream.
- Deleting a workflow no longer parses a 404 response as no content. A 404 is unhandled again, as it was before v0.2.0.

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
