# ZenFlow Add-on #08 — Batch Operations

## Overview

Save multiple items in one request using the `jsonb_array_elements` pattern. Exercises the func serialization + single-query batch INSERT pattern that replaces loops in SSaC.

## New Endpoint

### PUT /workflows/{id}/actions (`SaveWorkflowActions`)

Replace all actions of a workflow in one request (delete-then-insert).

1. `@auth` — admin only, same org.
2. `@get` the workflow (verify exists).
3. `@call` func to serialize the actions array to JSON.
4. `@delete` existing actions for the workflow.
5. `@put` batch insert via `jsonb_array_elements`.
6. `@response` confirmation.

## Key Pattern

```
func (pure): []Action → JSON string serialization
sqlc query (:exec): jsonb_array_elements(@items::jsonb) for batch INSERT
SSaC: @call serialize → @delete old → @put batchInsert
```

### sqlc query

```sql
-- name: ActionBatchInsert :exec
INSERT INTO actions (workflow_id, type, config, sequence_order)
SELECT @workflow_id, item->>'type', item->>'config', (item->>'sequence_order')::bigint
FROM jsonb_array_elements(@items::jsonb) AS item;
```

### func

`func/workflow/serialize_actions.go` — takes `[]ActionInput`, returns `{ItemsJSON string}`.

## E2E Scenario

- Create a workflow, batch-save 3 actions, verify count matches.
- Batch-save again with 2 actions, verify old ones replaced.
