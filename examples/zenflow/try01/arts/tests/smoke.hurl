# ZenFlow try01 — smoke scenario
# Register -> Login -> CreateWorkflow -> ListWorkflows -> ActivateWorkflow

POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "admin+{{newUuid}}@example.com",
  "password": "p@ssw0rd123"
}
HTTP 201
[Captures]
user_id: jsonpath "$.user.id"
[Asserts]
jsonpath "$.user.email" exists
jsonpath "$.user.role" == "admin"


POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "{{user_id}}@placeholder",
  "password": "wrongpassword"
}
HTTP 401


# Re-run login with the previously registered credentials is not possible
# because captures are per-file and the email uses newUuid. Instead we
# register a second account deterministically for the remaining flow.
POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "flow+{{newUuid}}@example.com",
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
  "title": "Daily digest",
  "trigger_event": "cron.daily"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.status" == "draft"
jsonpath "$.workflow.title" == "Daily digest"


GET {{host}}/workflows?page=1&per_page=20
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.total" isInteger
jsonpath "$.items" isCollection


POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"
