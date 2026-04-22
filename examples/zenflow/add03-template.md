# ZenFlow Add-on #03 — Workflow Template Marketplace

## Overview
Publish a workflow as a public template; other orgs search and clone it. Exercises cursor pagination, `@exists` guard, and `@dto` models.

## New Endpoints
- **POST /templates** (`PublishTemplate`) — publish a workflow as a template.
- **GET /templates** (`ListTemplates`) — cursor-paginated search with category filter.
- **POST /templates/{id}/clone** (`CloneTemplate`) — clone a template into caller's org as a new workflow.
- **GET /templates/{id}** (`GetTemplate`) — template detail.

## DDL
- `templates` table: `id, source_workflow_id (FK), org_id (FK), title, description, category, clone_count INT DEFAULT 0, created_at`.
- `CREATE UNIQUE INDEX idx_templates_source ON templates(source_workflow_id)` — prevent double-publish.

## `@dto` Models
- `TemplateDetail` — template + author org name + action count. Denormalized view, no backing table.

## Pagination (standard OpenAPI parameters — no `x-*` extensions)
```yaml
/templates:
  get:
    operationId: ListTemplates
    parameters:
      - { name: cursor,   in: query, schema: { type: string } }
      - { name: per_page, in: query, schema: { type: integer, default: 20, maximum: 100 } }
      - { name: category, in: query, schema: { type: string } }
    responses:
      '200':
        content:
          application/json:
            schema:
              properties:
                items: { type: array, items: { $ref: '#/components/schemas/Template' } }
              required: [items]
```
Cursor is `id DESC` (fixed). `category` is a standard query parameter.

## SSaC Design
- `PublishTemplate`: `@get Workflow` → `@exists Template "Already published" 409` → `@post Template`.
- `CloneTemplate`: `@get Template` → clone workflow into caller's org → copy actions → increment `clone_count`.

## Verification Points
- Cursor pagination: `@get []Template items = Template.ListCursor({...})` + explicit `@response { items: items }`.
- `@exists` guard for duplicate prevention (409 if non-nil).
- `@dto` model (no DDL-matching).
- Response shape: `{ items: [...] }` — last-page detection via `len(items) < per_page`.

## E2E Scenario
- Org A: create workflow → publish template → duplicate publish attempt → 409.
- Org B: list templates (cursor pagination) → clone template → confirm own workflow created.
