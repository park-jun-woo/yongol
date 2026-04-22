# ZenFlow Add-on #06 — Audit Logs (cache + offset pagination + sort)

## Overview
Record an audit entry on every significant operation (workflow create/activate/execute, template publish/clone, etc.). Recent entries served from `cache`; full history via DB with offset pagination and sort.

## Verification Points
- Built-in `cache` package via `@call cache.Set/Get`.
- `cache.backend: postgres`.
- Offset pagination with sort (multiple allowed columns, runtime switching).
- Propagating SSOT change: inserting one `@post AuditLog.Create(...)` line into several existing services — regression check.

## manifest.yaml
- Add `cache.backend: postgres`.

## New Endpoints
- **GET /audit-logs** (`ListAuditLogs`) — offset pagination, sortable by `created_at` / `action`, filterable by `action` / `actor_id`.
- **GET /audit-logs/recent** (`GetRecentAuditLogs`) — cache-backed recent N entries for dashboard.

## DDL
- `audit_logs` table: `id, org_id (FK), actor_id (FK users), action VARCHAR(100), resource_type VARCHAR(50), resource_id BIGINT, detail TEXT, created_at`.
- Indexes: `org_id`, `created_at`, `action`.

## Pagination & Sort (standard OpenAPI parameters)
```yaml
/audit-logs:
  get:
    operationId: ListAuditLogs
    parameters:
      - { name: page,      in: query, schema: { type: integer, default: 1 } }
      - { name: per_page,  in: query, schema: { type: integer, default: 20, maximum: 100 } }
      - { name: sort_by,   in: query, schema: { type: string, enum: [created_at, action], default: created_at } }
      - { name: sort_dir,  in: query, schema: { type: string, enum: [asc, desc], default: desc } }
      - { name: action,    in: query, schema: { type: string } }
      - { name: actor_id,  in: query, schema: { type: integer, format: int64 } }
    responses:
      '200':
        content:
          application/json:
            schema:
              properties:
                items: { type: array, items: { $ref: '#/components/schemas/AuditLog' } }
                total: { type: integer }
              required: [items, total]
```

## SSaC Design
- `ListAuditLogs`:
  - `@get []AuditLog items = AuditLog.ListByOrgIDPaged({OrgID: currentUser.OrgID, Page: request.page, PerPage: request.per_page, SortBy: request.sort_by, SortDir: request.sort_dir, FilterAction: request.action, FilterActorID: request.actor_id})`
  - `@get int64 total = AuditLog.CountByOrgIDFiltered({OrgID: currentUser.OrgID, FilterAction: request.action, FilterActorID: request.actor_id})`
  - `@response { items: items, total: total }`
- `GetRecentAuditLogs`:
  - `@call cache.Get({Key: cacheKey})` — on hit, return; on miss, fall back to DB query.
- Insert `@post AuditLog.Create({...})` one line into existing services: `CreateWorkflow`, `ActivateWorkflow`, `ExecuteWorkflow`, `PublishTemplate`, `CloneTemplate`.

## E2E Scenario
- Create → activate → execute workflow → list audit logs (≥3 entries) → verify sort-by-action → verify filter-by-action.
- Regression: existing tests for modified services still pass (response / status unchanged).
