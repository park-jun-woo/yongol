# ZenFlow Add-on #01 — Workflow Versioning

## Overview
Clone an existing workflow into a new `draft` version, copying its actions. Support listing all versions of a workflow.

## New Endpoints
- **POST /workflows/{id}/new-version** (`CreateWorkflowVersion`) — clone workflow, bump version, copy actions.
- **GET /workflows/{id}/versions** (`ListWorkflowVersions`) — list all versions by `root_workflow_id`.

## DDL Changes
- Add `version BIGINT NOT NULL DEFAULT 1`, `root_workflow_id BIGINT NOT NULL DEFAULT 0` to `workflows`.

## sqlc Queries
- `WorkflowCreateVersion :one` — INSERT with explicit version / root.
- `WorkflowListVersions :many` — `WHERE (root_workflow_id = $1 OR id = $1) AND org_id = $2`.
- `ActionCopyToWorkflow :exec` — bulk `INSERT ... SELECT` of actions.

## Custom Functions
- `resolveRootID(WorkflowID, RootWorkflowID)` — if `root_workflow_id` is 0, return own ID; else return existing root. (Delegated to Func because SSaC has no `if`.)
- `nextVersion(CurrentVersion)` — returns `CurrentVersion + 1`.

## Design Decisions
1. **`root_workflow_id` pattern.** Use the ID of the original version rather than parent ID — enables a single OR query instead of recursive traversal.
2. **`INSERT ... SELECT` for action copy.** SSaC has no loops — bulk copy at the DB layer.
3. **`DEFAULT` values preserve existing flow.** `CreateWorkflow` unchanged (version=1, root_workflow_id=0).

## Authorization
- `CreateWorkflowVersion`: admin only.
- `ListWorkflowVersions`: any authenticated user; tenant isolation handled by the query `AND org_id = $2`.

## E2E Scenario
Create workflow → add 2 actions → create new version → list versions → v1 and v2 present.
