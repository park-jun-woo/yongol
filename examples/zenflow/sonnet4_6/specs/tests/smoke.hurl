# Login
POST {{host}}/auth/login
{
  "email": "admin@example.com",
  "password": "testpassword123"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# CreateWorkflow
POST {{host}}/workflows
Authorization: Bearer {{token}}
{
  "title": "Test Workflow",
  "trigger_event": "manual"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.id"
org_id: jsonpath "$.org_id"

# ListWorkflows
GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200

# GetWorkflow
GET {{host}}/workflows/{{workflow_id}}
Authorization: Bearer {{token}}
HTTP 200

# AddAction
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token}}
{
  "action_type": "http",
  "config": "{}",
  "sequence_order": 1
}
HTTP 201

# ActivateWorkflow
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# PauseWorkflow
POST {{host}}/workflows/{{workflow_id}}/pause
Authorization: Bearer {{token}}
HTTP 200

# ActivateWorkflow again (from paused)
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# ArchiveWorkflow
POST {{host}}/workflows/{{workflow_id}}/archive
Authorization: Bearer {{token}}
HTTP 200

# Create another workflow for execution
POST {{host}}/workflows
Authorization: Bearer {{token}}
{
  "title": "Exec Workflow",
  "trigger_event": "manual"
}
HTTP 201
[Captures]
exec_workflow_id: jsonpath "$.id"

# Activate for execution
POST {{host}}/workflows/{{exec_workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# ExecuteWorkflow
POST {{host}}/workflows/{{exec_workflow_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Captures]
log_id: jsonpath "$.id"

# ListExecutionLogs
GET {{host}}/workflows/{{exec_workflow_id}}/execution-logs
Authorization: Bearer {{token}}
HTTP 200

# CreateWorkflowVersion
POST {{host}}/workflows/{{exec_workflow_id}}/new-version
Authorization: Bearer {{token}}
HTTP 201

# ListWorkflowVersions
GET {{host}}/workflows/{{exec_workflow_id}}/versions
Authorization: Bearer {{token}}
HTTP 200

# CreateWebhook
POST {{host}}/webhooks
Authorization: Bearer {{token}}
{
  "url": "https://example.com/webhook",
  "event_type": "workflow.executed"
}
HTTP 201
[Captures]
webhook_id: jsonpath "$.id"

# ListWebhooks
GET {{host}}/webhooks
Authorization: Bearer {{token}}
HTTP 200

# DeleteWebhook
DELETE {{host}}/webhooks/{{webhook_id}}
Authorization: Bearer {{token}}
HTTP 204

# PublishTemplate
POST {{host}}/templates
Authorization: Bearer {{token}}
{
  "source_workflow_id": "{{exec_workflow_id}}",
  "title": "My Template",
  "description": "A template",
  "category": "automation"
}
HTTP 201
[Captures]
template_id: jsonpath "$.id"

# ListTemplates
GET {{host}}/templates
Authorization: Bearer {{token}}
HTTP 200

# GetTemplate
GET {{host}}/templates/{{template_id}}
Authorization: Bearer {{token}}
HTTP 200

# CloneTemplate
POST {{host}}/templates/{{template_id}}/clone
Authorization: Bearer {{token}}
HTTP 201

# Create workflow for report execution
POST {{host}}/workflows
Authorization: Bearer {{token}}
{
  "title": "Report Workflow",
  "trigger_event": "manual"
}
HTTP 201
[Captures]
report_workflow_id: jsonpath "$.id"

# Activate for report execution
POST {{host}}/workflows/{{report_workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# ExecuteWithReport
POST {{host}}/workflows/{{report_workflow_id}}/execute-with-report
Authorization: Bearer {{token}}
HTTP 200
[Captures]
report_log_id: jsonpath "$.id"

# GetExecutionReport
GET {{host}}/execution-logs/{{report_log_id}}/report
Authorization: Bearer {{token}}
HTTP 200

# SetSchedule
POST {{host}}/workflows/{{exec_workflow_id}}/schedule
Authorization: Bearer {{token}}
{
  "cron_expr": "0 9 * * *"
}
HTTP 200

# GetSchedule
GET {{host}}/workflows/{{exec_workflow_id}}/schedule
Authorization: Bearer {{token}}
HTTP 200

# DeleteSchedule
DELETE {{host}}/workflows/{{exec_workflow_id}}/schedule
Authorization: Bearer {{token}}
HTTP 204

# ListAuditLogs
GET {{host}}/audit-logs
Authorization: Bearer {{token}}
HTTP 200
[Captures]
audit_log_id: jsonpath "$.items[0].id"

# GetRecentAuditLogs
GET {{host}}/audit-logs/recent
Authorization: Bearer {{token}}
HTTP 200

# GetDashboard
GET {{host}}/dashboard
Authorization: Bearer {{token}}
HTTP 200

# GetAuditLog
GET {{host}}/audit-logs/{{audit_log_id}}
Authorization: Bearer {{token}}
HTTP 200

# GetExecutionDetail
GET {{host}}/execution-logs/{{log_id}}/detail
Authorization: Bearer {{token}}
HTTP 200

# SaveWorkflowActions
PUT {{host}}/workflows/{{exec_workflow_id}}/actions
Authorization: Bearer {{token}}
HTTP 200

# VerifyOrgAddress
POST {{host}}/organizations/{{org_id}}/verify-address
Authorization: Bearer {{token}}
{
  "address": "123 Main St, Springfield"
}
HTTP 200

# AutoAssignWorkflow
POST {{host}}/workflows/{{exec_workflow_id}}/auto-assign
Authorization: Bearer {{token}}
HTTP 200
