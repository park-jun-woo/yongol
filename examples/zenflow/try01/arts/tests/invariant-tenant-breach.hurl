# zenflow.md §7 — cross-tenant execute -> 403

POST http://localhost:8080/auth/signup
Content-Type: application/json
{
  "email": "orga-admin@zenflow.test",
  "password": "password123",
  "name": "Org A Admin",
  "org_name": "Org A",
  "plan_type": "pro",
  "credits_balance": 50
}
HTTP 200
[Captures]
a_token: jsonpath "$.access_token"

POST http://localhost:8080/workflows
Authorization: Bearer {{a_token}}
Content-Type: application/json
{
  "title": "Org A Wf",
  "trigger_event": "manual"
}
HTTP 200
[Captures]
a_workflow_id: jsonpath "$.workflow.id"

POST http://localhost:8080/auth/signup
Content-Type: application/json
{
  "email": "orgb-admin@zenflow.test",
  "password": "password123",
  "name": "Org B Admin",
  "org_name": "Org B",
  "plan_type": "pro",
  "credits_balance": 50
}
HTTP 200
[Captures]
b_token: jsonpath "$.access_token"

POST http://localhost:8080/workflows/{{a_workflow_id}}/execute
Authorization: Bearer {{b_token}}
HTTP 403
