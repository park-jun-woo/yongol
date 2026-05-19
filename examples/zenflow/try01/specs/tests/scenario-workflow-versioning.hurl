# Scenario: Workflow Versioning
# Create workflow -> add 2 actions -> create new version -> list versions -> v1 and v2 present

# Bootstrap
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "ver-admin@zenflow.test", "password": "Str0ngP@ss", "org_id": "00000000-0000-0000-0000-000000000001", "role": "admin" }
HTTP 201

POST {{host}}/auth/login
Content-Type: application/json
{ "email": "ver-admin@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
ver_token: jsonpath "$.access_token"

# Create workflow
POST {{host}}/api/workflows
Content-Type: application/json
Authorization: Bearer {{ver_token}}
{ "title": "Version Test WF", "trigger_event": "on_signup" }
HTTP 201
[Captures]
ver_wf_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.version" == 1

# Add action 1
POST {{host}}/api/workflows/{{ver_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{ver_token}}
{ "action_type": "send_email", "config": "{}", "sequence_order": 1 }
HTTP 201

# Add action 2
POST {{host}}/api/workflows/{{ver_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{ver_token}}
{ "action_type": "webhook", "config": "{}", "sequence_order": 2 }
HTTP 201

# Create new version
POST {{host}}/api/workflows/{{ver_wf_id}}/new-version
Authorization: Bearer {{ver_token}}
HTTP 201
[Captures]
ver_v2_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.version" == 2
jsonpath "$.workflow.status" == "draft"

# List versions using original workflow ID
GET {{host}}/api/workflows/{{ver_wf_id}}/versions
Authorization: Bearer {{ver_token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2

# Check that v2 has the copied actions
GET {{host}}/api/workflows/{{ver_v2_id}}/actions
Authorization: Bearer {{ver_token}}
HTTP 200
[Asserts]
jsonpath "$.actions" count == 2
