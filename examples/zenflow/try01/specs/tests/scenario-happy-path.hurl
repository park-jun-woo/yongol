# Happy path scenario — register, login, create workflow, list, activate.
# Same end-to-end flow as smoke.hurl, kept separate for naming compliance
# with KindScenario detection (`scenario-*.hurl`).

POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "scenario+{{newUuid}}@example.com",
  "password": "p@ssw0rd123"
}
HTTP 201
[Captures]
login_email: jsonpath "$.user.email"


POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "{{login_email}}",
  "password": "p@ssw0rd123"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"


POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Scenario workflow",
  "trigger_event": "manual.trigger"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"


GET {{host}}/workflows?page=1&per_page=20
Authorization: Bearer {{token}}
HTTP 200


POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"
