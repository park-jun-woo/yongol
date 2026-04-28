# ZenFlow Smoke Tests
# Bearer auth mode

# Register a new organization admin user
POST {{host}}/auth/register
Content-Type: application/json
{
  "org_id": 1,
  "email": "admin@zenflow-smoke.example.com",
  "password": "p@ssw0rd1",
  "role": "admin"
}
HTTP 201
[Asserts]
jsonpath "$.user.id" isInteger
jsonpath "$.user.email" == "admin@zenflow-smoke.example.com"

# Login
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow-smoke.example.com",
  "password": "p@ssw0rd1"
}
HTTP 200
[Asserts]
jsonpath "$.access_token" exists
[Captures]
token: jsonpath "$.access_token"

# Create a workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Smoke Test Workflow",
  "trigger_event": "manual"
}
HTTP 201
[Asserts]
jsonpath "$.workflow.id" isInteger
jsonpath "$.workflow.status" == "draft"
[Captures]
workflow_id: jsonpath "$.workflow.id"

# Add an action to the workflow
POST {{host}}/workflows/{{workflow_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "action_type": "http_call",
  "payload_template": {"url": "https://example.com"},
  "sequence_order": 1
}
HTTP 201
[Asserts]
jsonpath "$.action.id" isInteger

# Get workflow
GET {{host}}/workflows/{{workflow_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.id" == {{workflow_id}}

# List workflows
GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" exists
