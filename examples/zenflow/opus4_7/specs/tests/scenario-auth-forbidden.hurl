# Establish auth context first (satisfies XOH-06)
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd!"
}
HTTP 200

# Login — 401 bad credentials
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "wrong@example.com",
  "password": "wrongpassword1"
}
HTTP 401

# CreateWorkflow — 403 no token
POST {{host}}/workflows
Content-Type: application/json
{
  "title": "Unauthorized",
  "trigger_event": "on_test"
}
HTTP 403

# ListWorkflows — 403 no token
GET {{host}}/workflows
HTTP 403

# GetWorkflow — 403 no token
GET {{host}}/workflows/1
HTTP 403

# AddAction — 403 no token
POST {{host}}/workflows/1/actions
Content-Type: application/json
{
  "action_type": "send_email",
  "config": "{}",
  "sequence_order": 1
}
HTTP 403

# SaveWorkflowActions — 403 no token
PUT {{host}}/workflows/1/actions
Content-Type: application/json
{
  "actions_json": "[]"
}
HTTP 403

# ActivateWorkflow — 403 no token
POST {{host}}/workflows/1/activate
HTTP 403

# PauseWorkflow — 403 no token
POST {{host}}/workflows/1/pause
HTTP 403

# ArchiveWorkflow — 403 no token
POST {{host}}/workflows/1/archive
HTTP 403

# ExecuteWorkflow — 403 no token
POST {{host}}/workflows/1/execute
HTTP 403

# ExecuteWithReport — 403 no token
POST {{host}}/workflows/1/execute-with-report
HTTP 403

# ListExecutionLogs — 403 no token
GET {{host}}/workflows/1/execution-logs
HTTP 403

# CreateWorkflowVersion — 403 no token
POST {{host}}/workflows/1/new-version
HTTP 403

# ListWorkflowVersions — 403 no token
GET {{host}}/workflows/1/versions
HTTP 403

# AutoAssignWorkflow — 403 no token
POST {{host}}/workflows/1/auto-assign
HTTP 403

# CreateWebhook — 403 no token
POST {{host}}/webhooks
Content-Type: application/json
{
  "url": "https://example.com/hook",
  "event_type": "workflow.executed"
}
HTTP 403

# ListWebhooks — 403 no token
GET {{host}}/webhooks
HTTP 403

# DeleteWebhook — 403 no token
DELETE {{host}}/webhooks/1
HTTP 403

# PublishTemplate — 403 no token
POST {{host}}/templates
Content-Type: application/json
{
  "source_workflow_id": 1,
  "title": "Unauthorized Template",
  "description": "test",
  "category": "testing"
}
HTTP 403

# CloneTemplate — 403 no token
POST {{host}}/templates/1/clone
HTTP 403

# GetExecutionReport — 403 no token
GET {{host}}/execution-logs/1/report
HTTP 403

# GetExecutionDetail — 403 no token
GET {{host}}/execution-logs/1/detail
HTTP 403

# SetSchedule — 403 no token
POST {{host}}/workflows/1/schedule
Content-Type: application/json
{
  "cron": "0 9 * * 1"
}
HTTP 403

# GetSchedule — 403 no token
GET {{host}}/workflows/1/schedule
HTTP 403

# DeleteSchedule — 403 no token
DELETE {{host}}/workflows/1/schedule
HTTP 403

# ListAuditLogs — 403 no token
GET {{host}}/audit-logs?page=1&per_page=20
HTTP 403

# GetRecentAuditLogs — 403 no token
GET {{host}}/audit-logs/recent
HTTP 403

# GetAuditLog — 403 no token
GET {{host}}/audit-logs/1
HTTP 403

# GetDashboard — 403 no token
GET {{host}}/dashboard
HTTP 403

# VerifyOrgAddress — 403 no token
POST {{host}}/organizations/1/verify-address
Content-Type: application/json
{
  "address": "123 Test Street"
}
HTTP 403
