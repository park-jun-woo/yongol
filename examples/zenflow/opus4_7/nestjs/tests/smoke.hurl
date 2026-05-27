# Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd!"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# CreateWorkflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Smoke Test Workflow",
  "trigger_event": "on_signup"
}
HTTP 201
[Captures]
wf_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.title" == "Smoke Test Workflow"
jsonpath "$.workflow.status" == "draft"

# ListWorkflows
GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 1

# GetWorkflow
GET {{host}}/workflows/{{wf_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.id" == {{wf_id}}

# AddAction 1
POST {{host}}/workflows/{{wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "action_type": "send_email",
  "config": "{\"to\": \"user@example.com\"}",
  "sequence_order": 1
}
HTTP 201
[Asserts]
jsonpath "$.action.action_type" == "send_email"

# AddAction 2
POST {{host}}/workflows/{{wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "action_type": "webhook_call",
  "config": "{\"url\": \"https://example.com/hook\"}",
  "sequence_order": 2
}
HTTP 201

# ActivateWorkflow
POST {{host}}/workflows/{{wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# ExecuteWorkflow
POST {{host}}/workflows/{{wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_log.status" == "completed"
jsonpath "$.execution_log.credits_spent" == 1

# ListExecutionLogs
GET {{host}}/workflows/{{wf_id}}/execution-logs
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_logs" count >= 1

# PauseWorkflow (reactivate first)
POST {{host}}/workflows/{{wf_id}}/pause
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# Re-activate
POST {{host}}/workflows/{{wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# ArchiveWorkflow
POST {{host}}/workflows/{{wf_id}}/archive
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "archived"

# CreateWorkflowVersion — create a new version from the original draft workflow
# First create a fresh workflow to version
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Versioned Workflow",
  "trigger_event": "on_update"
}
HTTP 201
[Captures]
ver_wf_id: jsonpath "$.workflow.id"

# Add action to the versionable workflow
POST {{host}}/workflows/{{ver_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "action_type": "log_event",
  "config": "{}",
  "sequence_order": 1
}
HTTP 201

# CreateWorkflowVersion
POST {{host}}/workflows/{{ver_wf_id}}/new-version
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2
jsonpath "$.workflow.status" == "draft"

# ListWorkflowVersions
GET {{host}}/workflows/{{ver_wf_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2

# CreateWebhook
POST {{host}}/webhooks
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "url": "https://example.com/webhook",
  "event_type": "workflow.executed"
}
HTTP 201
[Captures]
webhook_id: jsonpath "$.webhook.id"
[Asserts]
jsonpath "$.webhook.url" == "https://example.com/webhook"

# ListWebhooks
GET {{host}}/webhooks
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.webhooks" count >= 1

# DeleteWebhook
DELETE {{host}}/webhooks/{{webhook_id}}
Authorization: Bearer {{token}}
HTTP 204

# PublishTemplate
POST {{host}}/templates
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "source_workflow_id": "{{ver_wf_id}}",
  "title": "Test Template",
  "description": "A test template",
  "category": "testing"
}
HTTP 201
[Captures]
template_id: jsonpath "$.template.id"
[Asserts]
jsonpath "$.template.title" == "Test Template"

# PublishTemplate duplicate — 409
POST {{host}}/templates
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "source_workflow_id": "{{ver_wf_id}}",
  "title": "Duplicate",
  "description": "dup",
  "category": "testing"
}
HTTP 409

# ListTemplates
GET {{host}}/templates?per_page=10
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 1

# GetTemplate
GET {{host}}/templates/{{template_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.template.id" == {{template_id}}

# CloneTemplate
POST {{host}}/templates/{{template_id}}/clone
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.title" == "Versioned Workflow"

# ExecuteWithReport — create a new active workflow for this test
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Report Workflow",
  "trigger_event": "on_report"
}
HTTP 201
[Captures]
rpt_wf_id: jsonpath "$.workflow.id"

# Activate
POST {{host}}/workflows/{{rpt_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# ExecuteWithReport
POST {{host}}/workflows/{{rpt_wf_id}}/execute-with-report
Authorization: Bearer {{token}}
HTTP 200
[Captures]
rpt_log_id: jsonpath "$.execution_log.id"
[Asserts]
jsonpath "$.execution_log.status" == "completed"
jsonpath "$.report.report_key" isString

# GetExecutionReport
GET {{host}}/execution-logs/{{rpt_log_id}}/report
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.report_key" isString

# SetSchedule
POST {{host}}/workflows/{{rpt_wf_id}}/schedule
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "cron": "0 9 * * 1"
}
HTTP 200
[Asserts]
jsonpath "$.schedule.valid" == true

# GetSchedule
GET {{host}}/workflows/{{rpt_wf_id}}/schedule
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.cron" startsWith "\""

# DeleteSchedule
DELETE {{host}}/workflows/{{rpt_wf_id}}/schedule
Authorization: Bearer {{token}}
HTTP 204

# ListAuditLogs
GET {{host}}/audit-logs?page=1&per_page=20
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection
jsonpath "$.total" isInteger

# GetRecentAuditLogs
GET {{host}}/audit-logs/recent
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection

# SaveWorkflowActions — batch save 3 actions
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Batch Actions WF",
  "trigger_event": "on_batch"
}
HTTP 201
[Captures]
batch_wf_id: jsonpath "$.workflow.id"

PUT {{host}}/workflows/{{batch_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "actions_json": "[{\"type\":\"send_email\",\"config\":\"{}\",\"sequence_order\":1},{\"type\":\"webhook_call\",\"config\":\"{}\",\"sequence_order\":2},{\"type\":\"log_event\",\"config\":\"{}\",\"sequence_order\":3}]"
}
HTTP 200
[Asserts]
jsonpath "$.message" == "ok"

# SaveWorkflowActions — replace with 2 actions
PUT {{host}}/workflows/{{batch_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "actions_json": "[{\"type\":\"notify\",\"config\":\"{}\",\"sequence_order\":1},{\"type\":\"archive\",\"config\":\"{}\",\"sequence_order\":2}]"
}
HTTP 200
[Asserts]
jsonpath "$.message" == "ok"

# GetDashboard
GET {{host}}/dashboard
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.summary.org_name" isString
jsonpath "$.summary.plan_type" isString

# GetExecutionDetail — create a workflow, activate, execute to get exec log
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Detail Test WF",
  "trigger_event": "on_detail"
}
HTTP 201
[Captures]
detail_wf_id: jsonpath "$.workflow.id"

POST {{host}}/workflows/{{detail_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

POST {{host}}/workflows/{{detail_wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Captures]
detail_log_id: jsonpath "$.execution_log.id"

# GetExecutionDetail
GET {{host}}/execution-logs/{{detail_log_id}}/detail
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detail.status" == "completed"
jsonpath "$.detail.workflow_title" == "Detail Test WF"

# AutoAssignWorkflow — auto-assign a workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Auto Assign WF",
  "trigger_event": "on_signup"
}
HTTP 201
[Captures]
assign_wf_id: jsonpath "$.workflow.id"

POST {{host}}/workflows/{{assign_wf_id}}/auto-assign
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.assignment_confidence" == "high"

# VerifyOrgAddress — capture org_id from a workflow, then verify address
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Geo Test WF",
  "trigger_event": "on_geo"
}
HTTP 201
[Captures]
geo_org_id: jsonpath "$.workflow.org_id"

POST {{host}}/organizations/{{geo_org_id}}/verify-address
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "address": "123 Test Street, Seoul, Korea"
}
HTTP 200
[Asserts]
jsonpath "$.organization.latitude" == "37.5665"
jsonpath "$.organization.longitude" == "126.9780"
jsonpath "$.organization.address_verified" == true

# GetAuditLog — capture ID from list then fetch detail
GET {{host}}/audit-logs?page=0&per_page=100
Authorization: Bearer {{token}}
HTTP 200
[Captures]
audit_log_id: jsonpath "$.items[0].id"

GET {{host}}/audit-logs/{{audit_log_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.audit_log.id" == {{audit_log_id}}
