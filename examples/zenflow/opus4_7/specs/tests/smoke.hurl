# Login
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin@zenflow.io", "password": "admin123!"}
HTTP 200
[Captures]
token: jsonpath "$.access_token"
[Asserts]
jsonpath "$.access_token" isString

# CreateWorkflow
POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{"title": "Smoke Test Workflow", "trigger_event": "on_signup"}
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
jsonpath "$.workflow.title" == "Smoke Test Workflow"

# AddAction 1
POST {{host}}/workflows/{{wf_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{"action_type": "send_email", "config": "{\"to\":\"user@example.com\"}", "sequence_order": 1}
HTTP 201
[Asserts]
jsonpath "$.action.action_type" == "send_email"
jsonpath "$.action.sequence_order" == 1

# AddAction 2
POST {{host}}/workflows/{{wf_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{"action_type": "webhook", "config": "{\"url\":\"https://example.com/hook\"}", "sequence_order": 2}
HTTP 201
[Asserts]
jsonpath "$.action.action_type" == "webhook"

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

# PauseWorkflow (first reactivate since Execute is self-loop on active)
POST {{host}}/workflows/{{wf_id}}/pause
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# Reactivate from paused
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

# CreateWorkflowVersion (use the archived workflow)
POST {{host}}/workflows/{{wf_id}}/new-version
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2

# ListWorkflowVersions
GET {{host}}/workflows/{{wf_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 1

# PublishTemplate
POST {{host}}/templates
Authorization: Bearer {{token}}
Content-Type: application/json
{"workflow_id": "{{wf_id}}", "title": "Smoke Template", "description": "test", "category": "general"}
HTTP 201
[Captures]
tmpl_id: jsonpath "$.template.id"
[Asserts]
jsonpath "$.template.title" == "Smoke Template"

# ListTemplates
GET {{host}}/templates
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 1

# GetTemplate
GET {{host}}/templates/{{tmpl_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.template.id" == {{tmpl_id}}

# CloneTemplate
POST {{host}}/templates/{{tmpl_id}}/clone
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.status" == "draft"
