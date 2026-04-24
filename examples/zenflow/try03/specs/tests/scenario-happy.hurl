# ZenFlow happy path — register -> login -> create workflow -> activate -> execute
POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "alice@example.com",
  "password": "password123",
  "org_name": "Alice Corp"
}

HTTP 200
[Captures]
user_id: jsonpath "$.user.id"


POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "alice@example.com",
  "password": "password123"
}

HTTP 200
[Captures]
access_token: jsonpath "$.access_token"


POST {{host}}/workflows
Authorization: Bearer {{access_token}}
Content-Type: application/json
{
  "title": "Nightly Report",
  "trigger_event": "cron"
}

HTTP 200
[Captures]
wf_id: jsonpath "$.workflow.id"


POST {{host}}/workflows/{{wf_id}}/activate
Authorization: Bearer {{access_token}}

HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"


POST {{host}}/workflows/{{wf_id}}/execute
Authorization: Bearer {{access_token}}

HTTP 200
[Asserts]
jsonpath "$.log.credits_spent" == 1
