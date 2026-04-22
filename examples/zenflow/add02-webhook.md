# ZenFlow Add-on #02 — Webhook Notifications

## Overview
On workflow execution completion, publish an event with `@publish`; a `@subscribe` handler delivers the event to registered webhook URLs. Exercises yongol's pub/sub pipeline.

## manifest.yaml
- Add `queue.backend: postgres`.

## New Endpoints
- **POST /webhooks** (`CreateWebhook`) — register a webhook URL for the org.
- **GET /webhooks** (`ListWebhooks`) — list org's webhooks.
- **DELETE /webhooks/{id}** (`DeleteWebhook`) — delete a webhook.

## DDL
- `webhooks` table: `id, org_id (FK), url, event_type, created_at`.

## SSaC Changes
- `execute_workflow.ssac` — before the existing `@response`, add `@publish "workflow.executed" {WorkflowID: wf.ID, OrgID: wf.OrgID, Status: "completed"}`.
- `on_workflow_executed.ssac` (new) — `@subscribe "workflow.executed"` → load org's webhook URLs → `@call webhook.Deliver({URL: ..., Payload: ...})`.

## Custom Functions
- `webhook.Deliver(URL, Payload)` — simulated HTTP POST (Func purity forbids real network; simulation only).

## Verification Points
- `@publish` / `@subscribe` sequence types.
- `queue.backend: postgres` (uses the queue backend table).
- Asynchronous event handling pattern.
- Crosscheck: `@publish topic → @subscribe exists` (WARNING if missing).

## E2E Scenario
Create org → register webhook → create/activate/execute workflow → verify execution log + webhook delivery.

## Notes
- `@subscribe` functions are queue-triggered, not HTTP. `@response` is not allowed.
- The message struct must be declared inline in the same `.ssac` file.
