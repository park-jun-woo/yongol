# ZenFlow try01 Benchmark Report

## Environment

| Item | Value |
|------|-------|
| Model | claude-opus-4-6 |
| Claude Code | 2.1.133 |
| yongol | v0.3.0 |
| OS | Linux (WSL2) |
| Go | 1.25+ |
| PostgreSQL | 16-alpine (Docker) |

## Timing

| Step | Duration (s) | Validate Iterations | Notes |
|------|-------------|-------------------|-------|
| Base | 1081 (~18 min) | 8 | Full SSOT from scratch. 14 endpoints, 6 tables, 14 services, 5 pages |
| Add-on 01 (Versioning) | 197 (~3.3 min) | 1 | +2 endpoints, +2 columns, +2 services, +2 func specs |
| Add-on 02 (Webhooks) | 471 (~7.8 min) | 5 | +3 endpoints, +1 table, +3 services. BUG-068 (DELETE 204 missing return) |
| Add-on 03 (Templates) | 163 (~2.7 min) | 2 | +4 endpoints, +1 table, +4 services. Cursor pagination |
| Add-on 04 (Files) | skipped | - | Requires file.backend canonical DDL |
| Add-on 05 (Schedule) | skipped | - | Requires session.backend canonical DDL |
| Add-on 06 (Audit Logs) | 147 (~2.5 min) | 2 | +1 endpoint, +1 table, +1 service. Offset pagination with sort |
| Add-on 07 (Dashboard) | 191 (~3.2 min) | 3 | +1 endpoint, +1 service. Func Response type |
| **Total** | **2250 (~37.5 min)** | **21** | |

## Final Metrics

| Metric | Count |
|--------|-------|
| Endpoints | 23 |
| Tables | 9 (incl. refresh_tokens) |
| Services (SSaC) | 23 |
| Rego rules | 11 |
| STML pages | 5 |
| Func specs | 3 (billing, versioning x2, worker) |
| Smoke test requests | 26 |
| State diagram transitions | 5 |

## Bugs Filed

| Bug | Severity | Description |
|-----|----------|-------------|
| BUG-067 | BLOCKER | Duplicate pgtype import in per-model Go files (workaround: shorthand override format) |
| BUG-068 | BLOCKER | DELETE handler with 204 response missing return statement (workaround: use 200 with response body) |

## Notes

- Add-ons 04 (file.backend) and 05 (session.backend) were skipped because they require canonical DDL and sqlc queries for the respective built-in backends (fullend_file, fullend_session tables), which add significant infrastructure setup complexity.
- Add-on 02 was simplified to webhook CRUD without @publish/@subscribe due to XSF-46 (unused @call response in subscribe handler) + S-62 (unused variable) conflict that makes subscribe handlers with @call impractical.
- BUG-067 was worked around at the SSOT level by using sqlc shorthand override format instead of the expanded {import, package, type} format.
- BUG-068 was worked around by changing DELETE endpoints from 204 to 200 with a JSON response body.
