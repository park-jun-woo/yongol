# Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin@zenflow.test", "password": "password123"}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Create workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{"title": "Versioned Workflow", "trigger_event": "on_event"}
HTTP 201
[Captures]
wf_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.version" == 1

# Add action 1
POST {{host}}/workflows/{{wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{"action_type": "email", "config": "notify", "sequence_order": 1}
HTTP 201

# Add action 2
POST {{host}}/workflows/{{wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{"action_type": "webhook", "config": "call api", "sequence_order": 2}
HTTP 201

# Create new version
POST {{host}}/workflows/{{wf_id}}/new-version
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2
jsonpath "$.workflow.status" == "draft"

# List versions — should return v1 and v2
GET {{host}}/workflows/{{wf_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2
