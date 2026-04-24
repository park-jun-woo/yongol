# zenflow.md §7 — Happy Path
# Signup -> CreateWorkflow -> AddAction(2) -> Activate -> Execute

POST http://localhost:8080/auth/signup
Content-Type: application/json
{
  "email": "admin-happy@zenflow.test",
  "password": "password123",
  "name": "Happy Admin",
  "org_name": "Happy Org",
  "plan_type": "pro",
  "credits_balance": 100
}
HTTP 200
[Captures]
admin_token: jsonpath "$.access_token"

POST http://localhost:8080/workflows
Authorization: Bearer {{admin_token}}
Content-Type: application/json
{
  "title": "Daily Report",
  "trigger_event": "manual"
}
HTTP 200
[Captures]
workflow_id: jsonpath "$.workflow.id"

POST http://localhost:8080/workflows/{{workflow_id}}/actions
Authorization: Bearer {{admin_token}}
Content-Type: application/json
{
  "action_type": "http_get",
  "payload_template": "{}",
  "sequence_order": 1
}
HTTP 200

POST http://localhost:8080/workflows/{{workflow_id}}/actions
Authorization: Bearer {{admin_token}}
Content-Type: application/json
{
  "action_type": "send_email",
  "payload_template": "{}",
  "sequence_order": 2
}
HTTP 200

POST http://localhost:8080/workflows/{{workflow_id}}/activate
Authorization: Bearer {{admin_token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

POST http://localhost:8080/workflows/{{workflow_id}}/execute
Authorization: Bearer {{admin_token}}
HTTP 200
[Asserts]
jsonpath "$.log.status" == "completed"
jsonpath "$.log.credits_spent" == 1
