# Invariant: Tenant Breach — Org A user cannot access Org B's workflow

# Login as Org B admin to create a workflow
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin-b@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token_b: jsonpath "$.access_token"

# Create workflow in Org B
POST {{host}}/workflows
Authorization: Bearer {{token_b}}
Content-Type: application/json
{
  "title": "Org B Workflow",
  "trigger_event": "org.b.event"
}

HTTP 201
[Captures]
org_b_workflow_id: jsonpath "$.id"

# Login as Org A user
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token_a: jsonpath "$.access_token"

# Org A tries to get Org B's workflow — must be 403
GET {{host}}/workflows/{{org_b_workflow_id}}
Authorization: Bearer {{token_a}}

HTTP 403
