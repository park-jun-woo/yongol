# Smoke: Login → CreateWorkflow → ActivateWorkflow → ExecuteWorkflow
# Seed (organizations id=1 with credits, users id=1 admin) is loaded by tests/seed.sql

POST {{host}}/auth/login
Content-Type: application/json
{ "email": "admin@zenflow.test", "password": "p@ssw0rd" }
HTTP 200
[Captures]
token: jsonpath "$.access_token"

GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.total" isInteger

POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{ "title": "Smoke Workflow", "trigger_event": "manual" }
HTTP 200
[Captures]
wfid: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.status" == "draft"

POST {{host}}/workflows/{{wfid}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

POST {{host}}/workflows/{{wfid}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.log.workflow_id" == {{wfid}}
jsonpath "$.log.credits_spent" == 1
