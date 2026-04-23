# ZenFlow happy-path: register -> create workflow -> add 2 actions -> activate -> execute -> log recorded, credits deducted.

# 1) Register admin + org (100 credits)
POST {{host}}/auth/register
Content-Type: application/json
{
  "org_name": "Acme",
  "email": "admin@acme.test",
  "password": "password123"
}
HTTP 201
[Captures]
access_token: jsonpath "$.access_token"
org_id: jsonpath "$.organization.id"
initial_credits: jsonpath "$.organization.credits_balance"

# 2) Create workflow
POST {{host}}/workflows
Authorization: Bearer {{access_token}}
Content-Type: application/json
{
  "title": "Order fulfillment",
  "trigger_event": "order.placed"
}
HTTP 201
[Captures]
workflow_id: jsonpath "$.id"
[Asserts]
jsonpath "$.status" == "draft"

# 3) Add first action
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{access_token}}
Content-Type: application/json
{
  "action_type": "send_email",
  "payload_template": {"template": "welcome"},
  "sequence_order": 1
}
HTTP 201

# 4) Add second action
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{access_token}}
Content-Type: application/json
{
  "action_type": "webhook",
  "payload_template": {"url": "https://example.test/hook"},
  "sequence_order": 2
}
HTTP 201

# 5) Activate
POST {{host}}/workflows/{{workflow_id}}/activate
Authorization: Bearer {{access_token}}
HTTP 200
[Asserts]
jsonpath "$.status" == "active"

# 6) Execute
POST {{host}}/workflows/{{workflow_id}}/execute
Authorization: Bearer {{access_token}}
HTTP 200
[Asserts]
jsonpath "$.log.status" == "completed"
jsonpath "$.log.credits_spent" == 1
jsonpath "$.log.workflow_id" == {{workflow_id}}
