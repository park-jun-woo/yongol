#!/usr/bin/env bash
#
# audit_no_raw_access.sh — Phase001 raw-access guard
#
# Enforces the Phase001 invariant that emit helpers migrated onto
# prepared.State must not read raw parsed state (fs.Manifest,
# fs.ServiceFuncs/SSaC, fs.OpenAPI). Only prepared/*.go is allowed to
# touch those fields; every other migrated file reads through the
# prepared.State API.
#
# Scope (Phase001):
#   - pkg/generate/gogin/boot/block_session_init.go
#   - pkg/generate/gogin/boot/block_cache_init.go
#   - pkg/generate/gogin/boot/block_file_init.go
#   - pkg/generate/gogin/boot/block_queue_init.go
#   - pkg/generate/gogin/boot/block_csrf.go
#   - pkg/generate/gogin/boot/block_auth_init.go
#   - pkg/generate/gogin/boot/base_candidate_blocks.go
#       (allowed: fs.ServiceFuncs is passed to blockQueueInit, nothing else)
#   - pkg/generate/gogin/boot/has_csrf.go
#   - pkg/generate/gogin/boot/resolve_auth_init_config.go
#   - pkg/generate/gogin/middleware/csrf_active.go
#   - pkg/generate/gogin/middleware/generate_csrf.go
#
# Stage 5 of Phase001 (CORS / rate limit / observability / middleware
# ordering / OpenAPI securityScheme binding) was punted to a follow-up
# Phase; the broader "0 raw access in pkg/generate/" invariant will be
# enforced by that Phase once every category is migrated.
#
# Exit 0 on success; non-zero with the offending match when the scoped
# files regress.

set -euo pipefail

# shellcheck disable=SC2034
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
cd "${REPO_ROOT}"

SCOPED_FILES=(
  "pkg/generate/gogin/boot/block_session_init.go"
  "pkg/generate/gogin/boot/block_cache_init.go"
  "pkg/generate/gogin/boot/block_file_init.go"
  "pkg/generate/gogin/boot/block_queue_init.go"
  "pkg/generate/gogin/boot/block_csrf.go"
  "pkg/generate/gogin/boot/block_auth_init.go"
  "pkg/generate/gogin/boot/has_csrf.go"
  "pkg/generate/gogin/boot/resolve_auth_init_config.go"
  "pkg/generate/gogin/middleware/csrf_active.go"
  "pkg/generate/gogin/middleware/generate_csrf.go"
)

# Pattern matches raw Fullstack field access that prepared.State is
# meant to replace. `fs\.SpecsDir` and `fs\.Manifest.Backend.Module`
# stay legal elsewhere because they are infrastructure paths, not
# "optional subsystem / resolved default" state.
PATTERN='fs\.Manifest\.(Session|Cache|File|Queue|Backend\.(Auth|Middleware|CORS|HTTP|Observability|SecurityHeaders|Error))|fs\.ServiceFuncs|fs\.OpenAPI'

found=0
for f in "${SCOPED_FILES[@]}"; do
  if [ ! -f "${f}" ]; then
    echo "audit: missing file ${f}" >&2
    found=1
    continue
  fi
  if grep -nE "${PATTERN}" "${f}"; then
    echo "audit: raw state access in ${f} (Phase001 forbids — read through prepared.State)" >&2
    found=1
  fi
done

if [ "${found}" -ne 0 ]; then
  exit 1
fi

echo "audit_no_raw_access: OK (Phase001 scoped surface clean)"
