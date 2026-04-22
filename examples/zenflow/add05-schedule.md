# ZenFlow Add-on #05 — Workflow Scheduling (session + `@publish`)

## Overview
Register a cron schedule for a workflow; persist next-run time in `session`; on fire, `@publish` a trigger event that reuses `ExecuteWorkflow`.

## Verification Points
- Built-in `session` package via `@call session.Set/Get/Delete`.
- `session.backend: postgres` configuration.
- `@publish` reuse: schedule fire → `@publish "workflow.schedule.trigger"` → existing `ExecuteWorkflow` path.

## manifest.yaml
- Add `session.backend: postgres`.

## New Endpoints
- **POST /workflows/{id}/schedule** (`SetSchedule`) — register a cron schedule; store next-run in session.
- **GET /workflows/{id}/schedule** (`GetSchedule`) — read current schedule from session.
- **DELETE /workflows/{id}/schedule** (`DeleteSchedule`) — clear schedule from session.

## DDL
None. The session backend handles storage (via its own internal session table).

## SSaC Design
- `SetSchedule`: `@get Workflow` → enforce org isolation → `@call schedule.ParseCron({Expression: request.cron})` → `@call session.Set({Key: scheduleKey, Value: cronExpr, TTL: 0})` → `@response`.
- `GetSchedule`: `@call session.Get({Key: scheduleKey})` → `@response`.
- `DeleteSchedule`: `@call session.Delete({Key: scheduleKey})` → `@response`.

## Custom Functions
- `schedule.ParseCron(Expression)` — validate cron expression + compute next fire time (purity-safe).

## E2E Scenario
Create/activate workflow → set schedule → get schedule → delete schedule → get returns empty.
