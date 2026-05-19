# Invariant: Zero-credit org cannot activate a workflow

# Register seed user for bootstrap
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "broke-seed@zenflow.test", "password": "Str0ngP@ss", "org_id": "00000000-0000-0000-0000-000000000001", "role": "admin" }
HTTP 201

# Login for auth
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "broke-seed@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
broke_seed_token: jsonpath "$.access_token"

# Create zero-credit org
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{broke_seed_token}}
{ "name": "Broke Org", "plan_type": "free", "credits_balance": 0 }
HTTP 201
[Captures]
broke_org_id: jsonpath "$.organization.id"

# Register admin in broke org
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "broke-admin@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{broke_org_id}}", "role": "admin" }
HTTP 201

# Login as broke admin
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "broke-admin@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
broke_token: jsonpath "$.access_token"

# Create workflow
POST {{host}}/api/workflows
Content-Type: application/json
Authorization: Bearer {{broke_token}}
{ "title": "Broke Workflow", "trigger_event": "on_timer" }
HTTP 201
[Captures]
broke_wf_id: jsonpath "$.workflow.id"

# Attempt to activate with zero credits -> 402
POST {{host}}/api/workflows/{{broke_wf_id}}/activate
Authorization: Bearer {{broke_token}}
HTTP 402
