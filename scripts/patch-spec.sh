#!/bin/bash
# Converts 3.1 schema keywords to 3.0 equivalents for oapi-codegen:
# nullable unions become nullable schemas, const becomes a singleton enum,
# and schema examples arrays become a single example. Operations stay intact.
# Description punctuation is normalized for the generated Go comments.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

tmp="$(mktemp "${SPEC_DIR}/.patched.XXXXXX")"
trap 'rm -f -- "$tmp"' EXIT

# shellcheck disable=SC2016 # $key is a jq variable.
jq '
  walk(if type == "object" then
    (if has("const") then .enum = [.const] | del(.const) else . end)
    | (if (.description? | type) == "string" then
         .description |= gsub("\u2014|\u2013"; ",") else . end)
    | (if (.examples? | type) == "array" then
         .example = .examples[0] | del(.examples) else . end)
    | reduce ["anyOf", "oneOf"][] as $key (.;
        if (.[$key]? | type) == "array" and
           any(.[$key][]; .type? == "null") then
          .nullable = true
          | .[$key] |= map(select(.type? != "null"))
          | if (.[$key] | length) == 1 then
              . + .[$key][0] | del(.[$key])
            else . end
        else . end)
    else . end)
  | .openapi = "3.0.3"
' "${SPEC_DIR}/openapi.json" >"$tmp"
chmod 0644 "$tmp"
mv -- "$tmp" "${SPEC_DIR}/openapi.patched.json"
log INFO "converted OpenAPI 3.1 schemas to 3.0"
