# Login as zero-credit user
POST {{host}}/auth/login
Content-Type: application/json
{"email": "zero@zenflow.test", "password": "password123"}
HTTP 200
[Captures]
token_z: jsonpath "$.access_token"

# Create workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token_z}}
{"title": "Zero Credit Workflow", "trigger_event": "on_event"}
HTTP 201
[Captures]
wf_z_id: jsonpath "$.workflow.id"

# Try to activate -> 402 insufficient credits
POST {{host}}/workflows/{{wf_z_id}}/activate
Content-Type: application/json
Authorization: Bearer {{token_z}}
HTTP 402
