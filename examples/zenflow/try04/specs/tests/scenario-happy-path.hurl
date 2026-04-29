POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd1"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "My First Workflow",
  "trigger_event": "user.signup"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.status" == "draft"

POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

POST {{host}}/workflows/{{workflow_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.log.workflow_id" == "{{workflow_id}}"
jsonpath "$.log.credits_spent" isInteger
