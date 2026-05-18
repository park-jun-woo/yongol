# ZenFlow Add-on #10 — Conditional Update without @if

## Overview

Implement conditional logic in SSaC without `@if` using the no-op value pattern. The func returns a "do nothing" value when the condition is not met, and `@put` always executes but effectively becomes a no-op.

## New Endpoint

### POST /workflows/{id}/auto-assign (`AutoAssignWorkflow`)

Automatically assign a workflow to the best-matching team member based on trigger event type. If no match is found, the assignment field stays empty.

1. `@auth` — admin only, same org.
2. `@get` the workflow.
3. `@get` list of org members.
4. `@call func` — pure matching logic: returns `{MemberID: <uuid>, Confidence: "high"}` on match, `{MemberID: "00000000-0000-0000-0000-000000000000", Confidence: "none"}` on no match.
5. `@put` always executes — if MemberID is zero UUID, the DB column stays unchanged via `COALESCE(NULLIF(@member_id, '00000000-...')::uuid, assigned_to)`.
6. `@response` the updated workflow.

## Key Pattern

```
func (pure): evaluate condition → return real value OR no-op sentinel
sqlc query: COALESCE(NULLIF(@value, sentinel), existing_column)
SSaC: @call matchFunc → @put alwaysExecute (no branching needed)
```

### sqlc query

```sql
-- name: WorkflowAutoAssign :exec
UPDATE workflows
SET assigned_to = COALESCE(NULLIF(@member_id::text, '00000000-0000-0000-0000-000000000000')::uuid, assigned_to),
    assignment_confidence = @confidence
WHERE id = @id;
```

### func

`func/workflow/match_member.go` — takes workflow trigger_event + member list, returns best match or zero-UUID sentinel.

## DDL Change

Add `assigned_to UUID` (`-- @nullable`) and `assignment_confidence VARCHAR(10) NOT NULL DEFAULT 'none'` to `workflows`.

## E2E Scenario

- Create workflow + members, auto-assign → verify member assigned.
- Create workflow with unmatched trigger → verify assigned_to stays null, confidence = "none".
