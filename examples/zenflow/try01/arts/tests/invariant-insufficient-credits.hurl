# zenflow.md §7 — insufficient credits -> 402

POST http://localhost:8080/auth/signup
Content-Type: application/json
{
  "email": "broke@zenflow.test",
  "password": "password123",
  "name": "Broke Admin",
  "org_name": "Broke Org",
  "plan_type": "free",
  "credits_balance": 0
}
HTTP 200
[Captures]
broke_token: jsonpath "$.access_token"

POST http://localhost:8080/workflows
Authorization: Bearer {{broke_token}}
Content-Type: application/json
{
  "title": "Never Runs",
  "trigger_event": "manual"
}
HTTP 200
[Captures]
workflow_id: jsonpath "$.workflow.id"

POST http://localhost:8080/workflows/{{workflow_id}}/activate
Authorization: Bearer {{broke_token}}
HTTP 402
