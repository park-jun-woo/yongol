POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "smoke+{{newUuid}}@example.com",
  "password": "smokepass1",
  "org_name": "smoke-org"
}
HTTP 201
[Asserts]
jsonpath "$.user.id" isInteger
[Captures]
email: jsonpath "$.user.email"

POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "{{email}}",
  "password": "smokepass1"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Smoke Workflow",
  "trigger_event": "smoke.event"
}
HTTP 201
[Asserts]
jsonpath "$.workflow.id" isInteger
jsonpath "$.workflow.status" == "draft"
[Captures]
workflow_id: jsonpath "$.workflow.id"

POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

POST {{host}}/workflows/{{workflow_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.log.status" == "success"
jsonpath "$.log.credits_spent" == 1

GET {{host}}/workflows?page=1&per_page=20&sort_by=created_at&sort_dir=desc
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection
jsonpath "$.total" isInteger
