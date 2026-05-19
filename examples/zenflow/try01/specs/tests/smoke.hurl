# Step 1: Register seed admin in pre-seeded org
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "seed-admin@zenflow.test", "password": "Str0ngP@ss", "org_id": "00000000-0000-0000-0000-000000000001", "role": "admin" }
HTTP 201

# Step 2: Login as seed admin
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "seed-admin@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
seed_token: jsonpath "$.access_token"

# Step 3: Create Org A with credits (now authenticated)
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{seed_token}}
{ "name": "Org A", "plan_type": "pro", "credits_balance": 100 }
HTTP 201
[Captures]
org_a_id: jsonpath "$.organization.id"
[Asserts]
jsonpath "$.organization.name" == "Org A"

# Step 4: Create Org B with credits
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{seed_token}}
{ "name": "Org B", "plan_type": "free", "credits_balance": 50 }
HTTP 201
[Captures]
org_b_id: jsonpath "$.organization.id"

# Step 5: Create Zero-Credit Org
POST {{host}}/api/organizations
Content-Type: application/json
Authorization: Bearer {{seed_token}}
{ "name": "Zero Org", "plan_type": "free", "credits_balance": 0 }
HTTP 201
[Captures]
zero_org_id: jsonpath "$.organization.id"

# Step 6: Register admin user in Org A
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "admin-a@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{org_a_id}}", "role": "admin" }
HTTP 201
[Asserts]
jsonpath "$.user.email" == "admin-a@zenflow.test"

# Step 7: Register admin user in Org B
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "admin-b@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{org_b_id}}", "role": "admin" }
HTTP 201

# Step 8: Register user in zero-credit org
POST {{host}}/auth/register
Content-Type: application/json
{ "email": "zero@zenflow.test", "password": "Str0ngP@ss", "org_id": "{{zero_org_id}}", "role": "admin" }
HTTP 201

# Step 9: Login as admin A
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "admin-a@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
token_a: jsonpath "$.access_token"
[Asserts]
jsonpath "$.access_token" isString

# Step 10: Create a workflow
POST {{host}}/api/workflows
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "title": "My Workflow", "trigger_event": "on_signup" }
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"
[Asserts]
jsonpath "$.workflow.title" == "My Workflow"
jsonpath "$.workflow.status" == "draft"

# Step 11: List workflows
GET {{host}}/api/workflows
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 1

# Step 12: Get workflow
GET {{host}}/api/workflows/{{workflow_id}}
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.id" == "{{workflow_id}}"

# Step 13: Add action 1
POST {{host}}/api/workflows/{{workflow_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "action_type": "send_email", "config": "{}", "sequence_order": 1 }
HTTP 201
[Asserts]
jsonpath "$.action.action_type" == "send_email"

# Step 14: Add action 2
POST {{host}}/api/workflows/{{workflow_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "action_type": "webhook", "config": "{}", "sequence_order": 2 }
HTTP 201

# Step 15: List actions
GET {{host}}/api/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.actions" count == 2

# Step 16: Create new version
POST {{host}}/api/workflows/{{workflow_id}}/new-version
Authorization: Bearer {{token_a}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2

# Step 17: List versions
GET {{host}}/api/workflows/{{workflow_id}}/versions
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2

# Step 18: Create a webhook
POST {{host}}/api/webhooks
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "url": "https://example.com/hook", "event_type": "workflow.executed" }
HTTP 201
[Captures]
webhook_id: jsonpath "$.webhook.id"
[Asserts]
jsonpath "$.webhook.url" == "https://example.com/hook"

# Step 19: List webhooks
GET {{host}}/api/webhooks
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.webhooks" count >= 1

# Step 20: Delete webhook
DELETE {{host}}/api/webhooks/{{webhook_id}}
Authorization: Bearer {{token_a}}
HTTP 204

# Step 45: Auto-assign workflow
POST {{host}}/api/workflows/{{workflow_id}}/auto-assign
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.assignment_confidence" != "none"

# Step 44: Verify org address
POST {{host}}/api/organizations/{{org_a_id}}/verify-address
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "address": "123 Main St, Seoul" }
HTTP 200
[Asserts]
jsonpath "$.organization.address_verified" == true

# Step 42: Batch save actions (3 actions)
PUT {{host}}/api/workflows/{{workflow_id}}/actions/batch
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "items": "[{\"type\":\"email\",\"config\":\"{}\",\"sequence_order\":1},{\"type\":\"slack\",\"config\":\"{}\",\"sequence_order\":2},{\"type\":\"webhook\",\"config\":\"{}\",\"sequence_order\":3}]" }
HTTP 200
[Asserts]
jsonpath "$.actions" count == 3

# Step 43: Batch save again with 2 actions (replaces old ones)
PUT {{host}}/api/workflows/{{workflow_id}}/actions/batch
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "items": "[{\"type\":\"sms\",\"config\":\"{}\",\"sequence_order\":1},{\"type\":\"push\",\"config\":\"{}\",\"sequence_order\":2}]" }
HTTP 200
[Asserts]
jsonpath "$.actions" count == 2

# Step 22: Publish workflow as template
POST {{host}}/api/templates
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "workflow_id": "{{workflow_id}}", "description": "A test template", "category": "automation" }
HTTP 201
[Captures]
template_id: jsonpath "$.template.id"
[Asserts]
jsonpath "$.template.title" == "My Workflow"
jsonpath "$.template.category" == "automation"

# Step 23: Duplicate publish attempt should 409
POST {{host}}/api/templates
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "workflow_id": "{{workflow_id}}", "description": "Duplicate", "category": "automation" }
HTTP 409

# Step 24: List templates
GET {{host}}/api/templates
HTTP 200
[Asserts]
jsonpath "$.items" count >= 1

# Step 25: Get template detail
GET {{host}}/api/templates/{{template_id}}
HTTP 200
[Asserts]
jsonpath "$.template.id" == "{{template_id}}"

# Step 26: Login as admin B for cloning
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "admin-b@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
token_b: jsonpath "$.access_token"

# Step 27: Clone template (Org B clones from Org A)
POST {{host}}/api/templates/{{template_id}}/clone
Authorization: Bearer {{token_b}}
HTTP 201
[Asserts]
jsonpath "$.workflow.title" == "My Workflow"

# Step 28: Re-login as admin A
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "admin-a@zenflow.test", "password": "Str0ngP@ss" }
HTTP 200
[Captures]
token_a: jsonpath "$.access_token"

# Step 29: Activate the workflow
POST {{host}}/api/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 34: Set schedule
POST {{host}}/api/workflows/{{workflow_id}}/schedule
Content-Type: application/json
Authorization: Bearer {{token_a}}
{ "cron": "0 9 * * MON" }
HTTP 200
[Asserts]
jsonpath "$.cron" == "0 9 * * MON"

# Step 35: Get schedule
GET {{host}}/api/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.cron" contains "0 9 * * MON"

# Step 36: Delete schedule
DELETE {{host}}/api/workflows/{{workflow_id}}/schedule
Authorization: Bearer {{token_a}}
HTTP 204

# Step 30: Execute the workflow
POST {{host}}/api/workflows/{{workflow_id}}/execute
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.execution_log.status" == "completed"
jsonpath "$.execution_log.credits_spent" == 1

# Step 31: Execute with report
POST {{host}}/api/workflows/{{workflow_id}}/execute-with-report
Authorization: Bearer {{token_a}}
HTTP 200
[Captures]
exec_log_id: jsonpath "$.execution_log.id"
[Asserts]
jsonpath "$.execution_log.status" == "completed"
jsonpath "$.execution_log.report_key" != ""

# Step 32: Get execution report
GET {{host}}/api/execution-logs/{{exec_log_id}}/report
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.content" != ""

# Step 33: List execution logs
GET {{host}}/api/workflows/{{workflow_id}}/executions
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.execution_logs" count >= 1

# Step 19: Pause the workflow
POST {{host}}/api/workflows/{{workflow_id}}/pause
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# Step 20: Re-activate (paused -> active)
POST {{host}}/api/workflows/{{workflow_id}}/activate
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 21: Archive the workflow
POST {{host}}/api/workflows/{{workflow_id}}/archive
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "archived"

# Step 39: Get dashboard
GET {{host}}/api/dashboard
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.summary.org_name" == "Org A"

# Step 40: Get execution detail
GET {{host}}/api/execution-logs/{{exec_log_id}}/detail
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.detail.status" == "completed"

# Step 37: List audit logs
GET {{host}}/api/audit-logs
Authorization: Bearer {{token_a}}
HTTP 200
[Captures]
audit_log_id: jsonpath "$.items[0].id"
[Asserts]
jsonpath "$.items" count >= 3
jsonpath "$.total" >= 3

# Step 41: Get audit log detail
GET {{host}}/api/audit-logs/{{audit_log_id}}
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.audit_log.id" == "{{audit_log_id}}"

# Step 38: Get recent audit logs
GET {{host}}/api/audit-logs/recent
Authorization: Bearer {{token_a}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 1
