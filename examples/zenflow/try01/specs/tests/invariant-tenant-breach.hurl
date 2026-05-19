# Login as admin in org B
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin-b@zenflow.test", "password": "password123"}
HTTP 200
[Captures]
token_b: jsonpath "$.access_token"

# Login as admin in org A
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin@zenflow.test", "password": "password123"}
HTTP 200
[Captures]
token_a: jsonpath "$.access_token"

# Org A creates a workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token_a}}
{"title": "Org A Workflow", "trigger_event": "on_event"}
HTTP 201
[Captures]
wf_a_id: jsonpath "$.workflow.id"

# Org B tries to access Org A's workflow -> 403
GET {{host}}/workflows/{{wf_a_id}}
Authorization: Bearer {{token_b}}
HTTP 403
