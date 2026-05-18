# Invariant: Tenant Breach - Org B user cannot access Org A workflow

# Register seed user in pre-seeded org for bootstrap
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "inv-seed@zenflow.test", "password": "Str0ngP@ss", "org_id": "00000000-0000-0000-0000-000000000001", "role": "admin" }
HTTP 201

# Login for auth
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "inv-seed@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
inv_seed_token: jsonpath "$.access_token"

# Create Org A
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{inv_seed_token}}
{ "name": "Inv Org A", "plan_type": "pro", "credits_balance": 100 }
HTTP 201
[Captures]
inv_org_a_id: jsonpath "$.organization.id"

# Create Org B
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{inv_seed_token}}
{ "name": "Inv Org B", "plan_type": "pro", "credits_balance": 100 }
HTTP 201
[Captures]
inv_org_b_id: jsonpath "$.organization.id"

# Register admin in Org A
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "inv-admin-a@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{inv_org_a_id}}", "role": "admin" }
HTTP 201

# Register admin in Org B
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "inv-admin-b@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{inv_org_b_id}}", "role": "admin" }
HTTP 201

# Login as A
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "inv-admin-a@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
inv_token_a: jsonpath "$.access_token"

# Login as B
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "inv-admin-b@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
inv_token_b: jsonpath "$.access_token"

# A creates a workflow
POST {{host}}/api/workflows
Content-Type: application/json
Authorization: Bearer {{inv_token_a}}
{ "title": "A Private Workflow", "trigger_event": "on_event" }
HTTP 201
[Captures]
inv_wf_id: jsonpath "$.workflow.id"

# B tries to access A's workflow -> 403
GET {{host}}/api/workflows/{{inv_wf_id}}
Authorization: Bearer {{inv_token_b}}
HTTP 403

# B tries to activate A's workflow -> 403
POST {{host}}/api/workflows/{{inv_wf_id}}/activate
Authorization: Bearer {{inv_token_b}}
HTTP 403
