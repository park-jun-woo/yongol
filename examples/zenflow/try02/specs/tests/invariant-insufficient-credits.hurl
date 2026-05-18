# Invariant: Insufficient Credits — org with zero credits cannot activate

# Login as admin with zero credits org
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "zero@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token_zero: jsonpath "$.access_token"

# Create a workflow
POST {{host}}/workflows
Authorization: Bearer {{token_zero}}
Content-Type: application/json
{
  "title": "Zero Credits Workflow",
  "trigger_event": "test.event"
}

HTTP 201
[Captures]
zero_workflow_id: jsonpath "$.id"

# Try to activate — must be 402 Payment Required
POST {{host}}/workflows/{{zero_workflow_id}}/activate
Authorization: Bearer {{token_zero}}

HTTP 402
