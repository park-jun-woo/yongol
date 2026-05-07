# Login
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd1"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"
[Asserts]
jsonpath "$.access_token" isString

# ListWorkflows
GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection
jsonpath "$.total" isInteger

# CreateWorkflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Smoke Test Workflow",
  "trigger_event": "user.signup"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.status" == "draft"

# ActivateWorkflow
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# PauseWorkflow
POST {{host}}/workflows/{{workflow_id}}/pause
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# ArchiveWorkflow (reactivate first, then archive)
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

POST {{host}}/workflows/{{workflow_id}}/archive
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "archived"

# ExecuteWorkflow (create a new one for execution)
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Smoke Exec Workflow",
  "trigger_event": "order.placed"
}
HTTP 201
[Captures]
exec_wf_id: jsonpath "$.workflow.id"

POST {{host}}/workflows/{{exec_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

POST {{host}}/workflows/{{exec_wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.log.credits_spent" isInteger
