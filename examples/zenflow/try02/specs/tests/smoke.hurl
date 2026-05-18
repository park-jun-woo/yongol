# ZenFlow Smoke Test — Happy Path
# Admin creates a workflow, adds two actions, activates it, executes it, then verifies the log.

# Step 1: Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Step 2: Create a workflow
POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "My Test Workflow",
  "trigger_event": "user.signup"
}

HTTP 201
[Captures]
workflow_id: jsonpath "$.id"
[Asserts]
jsonpath "$.status" == "draft"
jsonpath "$.title" == "My Test Workflow"

# Step 3: List workflows
GET {{host}}/workflows
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 4: Get workflow
GET {{host}}/workflows/{{workflow_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.id" == "{{workflow_id}}"
jsonpath "$.status" == "draft"

# Step 5: Add first action
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "send_email",
  "config": "{}",
  "sequence_order": 1
}

HTTP 201
[Asserts]
jsonpath "$.action_type" == "send_email"
jsonpath "$.sequence_order" == 1

# Step 6: Add second action
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "notify_slack",
  "config": "{}",
  "sequence_order": 2
}

HTTP 201
[Asserts]
jsonpath "$.action_type" == "notify_slack"
jsonpath "$.sequence_order" == 2

# Step 7: Activate the workflow
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.status" == "active"

# Step 8: Execute the workflow
POST {{host}}/workflows/{{workflow_id}}/execute
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.credits_spent" == 1

# Step 9: Verify the execution log
GET {{host}}/workflows/{{workflow_id}}/logs
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 10: Pause the workflow
POST {{host}}/workflows/{{workflow_id}}/pause
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.status" == "paused"

# Step 11: Re-activate (resume)
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.status" == "active"

# Step 12a: ExecuteWithReport — execute workflow and generate report file
POST {{host}}/workflows/{{workflow_id}}/execute-with-report
Authorization: Bearer {{token}}

HTTP 200
[Captures]
report_log_id: jsonpath "$.id"
[Asserts]
jsonpath "$.status" == "success"
jsonpath "$.report_key" != ""

# Step 12b: GetExecutionReport — download the report file
GET {{host}}/execution-logs/{{report_log_id}}/report
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.body" != ""

# Step 12c: SetSchedule — register a cron schedule
POST {{host}}/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "cron": "0 9 * * 1"
}

HTTP 200
[Asserts]
jsonpath "$.expression" == "0 9 * * 1"

# Step 12d: GetSchedule — read current schedule
GET {{host}}/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.expression" == "0 9 * * 1"

# Step 12e: DeleteSchedule — clear the schedule
DELETE {{host}}/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token}}

HTTP 204

# Step 12f: GetSchedule after delete — returns empty expression
GET {{host}}/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.expression" == ""

# Step 12: Archive the workflow
POST {{host}}/workflows/{{workflow_id}}/archive
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.status" == "archived"

# Step 13: Create a new workflow for versioning test
POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Versioning Test Workflow",
  "trigger_event": "user.signup"
}

HTTP 201
[Captures]
ver_workflow_id: jsonpath "$.id"
[Asserts]
jsonpath "$.version" == 1

# Step 14: CreateWorkflowVersion — clone the workflow
POST {{host}}/workflows/{{ver_workflow_id}}/new-version
Authorization: Bearer {{token}}

HTTP 201
[Asserts]
jsonpath "$.version" == 2
jsonpath "$.status" == "draft"

# Step 15: ListWorkflowVersions — list all versions
GET {{host}}/workflows/{{ver_workflow_id}}/versions
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count == 2

# Step 16: CreateWebhook — register a webhook URL
POST {{host}}/webhooks
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "url": "https://example.com/webhook",
  "event_type": "completed"
}

HTTP 201
[Captures]
webhook_id: jsonpath "$.id"
[Asserts]
jsonpath "$.url" == "https://example.com/webhook"
jsonpath "$.event_type" == "completed"

# Step 17: ListWebhooks — list org webhooks
GET {{host}}/webhooks
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 18: DeleteWebhook — delete the webhook
DELETE {{host}}/webhooks/{{webhook_id}}
Authorization: Bearer {{token}}

HTTP 204

# Step 19: PublishTemplate — publish the versioning test workflow as a template
POST {{host}}/templates
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "workflow_id": "{{ver_workflow_id}}",
  "title": "Smoke Test Template",
  "description": "Template for smoke test",
  "category": "test"
}

HTTP 201
[Captures]
template_id: jsonpath "$.id"
[Asserts]
jsonpath "$.title" == "Smoke Test Template"
jsonpath "$.clone_count" == 0

# Step 20: ListTemplates — list templates
GET {{host}}/templates?per_page=20
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 21: GetTemplate — get template detail
GET {{host}}/templates/{{template_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.id" == "{{template_id}}"

# Step 22: CloneTemplate — clone the template
POST {{host}}/templates/{{template_id}}/clone
Authorization: Bearer {{token}}

HTTP 201
[Asserts]
jsonpath "$.status" == "draft"

# Step 23: ListAuditLogs — list audit entries with offset pagination
GET {{host}}/audit-logs?page=1&per_page=20&sort_by=created_at&sort_dir=desc
Authorization: Bearer {{token}}

HTTP 200
[Captures]
audit_log_id: jsonpath "$.items[0].id"
[Asserts]
jsonpath "$.items" count > 0
jsonpath "$.total" > 0

# Step 24: GetRecentAuditLogs — get recent entries from DB
GET {{host}}/audit-logs/recent
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 25: GetDashboard — org summary
GET {{host}}/dashboard
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.summary" != null

# Step 26: GetAuditLog — single audit log detail
GET {{host}}/audit-logs/{{audit_log_id}}
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.audit_log.id" == {{audit_log_id}}

# Step 27: GetExecutionDetail — enriched execution log detail
GET {{host}}/execution-logs/{{report_log_id}}/detail
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.detail.status" != null
