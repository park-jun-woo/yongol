# Tenant isolation: Org B's user cannot execute Org A's workflow — 403 Forbidden.

# Register Org A admin
POST {{host}}/auth/register
Content-Type: application/json
{
  "org_name": "OrgA",
  "email": "a-admin@test.test",
  "password": "password123"
}
HTTP 201
[Captures]
a_token: jsonpath "$.access_token"

# Org A admin creates and activates a workflow
POST {{host}}/workflows
Authorization: Bearer {{a_token}}
Content-Type: application/json
{ "title": "A workflow", "trigger_event": "evt" }
HTTP 201
[Captures]
a_workflow_id: jsonpath "$.id"

POST {{host}}/workflows/{{a_workflow_id}}/activate
Authorization: Bearer {{a_token}}
HTTP 200

# Register Org B admin (different org)
POST {{host}}/auth/register
Content-Type: application/json
{
  "org_name": "OrgB",
  "email": "b-admin@test.test",
  "password": "password123"
}
HTTP 201
[Captures]
b_token: jsonpath "$.access_token"

# Org B admin attempts to execute Org A's workflow — must be rejected.
POST {{host}}/workflows/{{a_workflow_id}}/execute
Authorization: Bearer {{b_token}}
HTTP 403
