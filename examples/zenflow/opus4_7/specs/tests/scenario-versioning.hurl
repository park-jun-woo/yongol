# Login
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin@zenflow.io", "password": "admin123!"}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Create source workflow
POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{"title": "Versioned Workflow", "trigger_event": "on_order"}
HTTP 201
[Captures]
src_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.version" == 1

# Add action to source
POST {{host}}/workflows/{{src_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{"action_type": "notify", "config": "{}", "sequence_order": 1}
HTTP 201

# Add second action
POST {{host}}/workflows/{{src_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{"action_type": "webhook", "config": "{}", "sequence_order": 2}
HTTP 201

# Create new version
POST {{host}}/workflows/{{src_id}}/new-version
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2
jsonpath "$.workflow.status" == "draft"

# List versions
GET {{host}}/workflows/{{src_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2
