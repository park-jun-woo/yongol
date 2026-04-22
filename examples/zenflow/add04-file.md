# ZenFlow Add-on #04 — Execution Report Files

## Overview
Save workflow execution results as files. Exercises the `file` built-in backend.

## manifest.yaml
- Add `file.backend: local`.

## New Endpoints
- **POST /workflows/{id}/execute-with-report** (`ExecuteWithReport`) — execute + generate/upload report file.
- **GET /execution-logs/{id}/report** (`GetExecutionReport`) — download the file linked to a log.

## DDL Changes
- Add `report_key VARCHAR(255) NOT NULL DEFAULT ''` to `execution_logs` (file storage key).

## SSaC Design
- `ExecuteWithReport` — existing `ExecuteWorkflow` flow plus:
  - `@call report.GenerateReport({...})` → `@call file.Upload({Key: reportKey, Body: ...})` → store `report_key` on the log.
- `GetExecutionReport` — `@get ExecutionLog` → `@call file.Download({Key: log.ReportKey})`.

## Custom Functions
- `report.GenerateReport(WorkflowID, ActionCount, Status)` — format the execution result as text (purity-safe).

## Verification Points
- Built-in `file` package via `@call` (no package-prefix `@model` — use `@call file.Upload({...})`).
- `file.backend: local` configuration.
- Go-interface parameter matching (context.Context omitted; Key/Body names must match).

## E2E Scenario
Create/activate/execute-with-report workflow → verify `report_key` set on log → download file.
