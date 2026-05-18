# Step 1: Seed organization (direct DB insert assumed via migration)
# The org with id=1 should exist with credits_balance > 0

# Step 2: Register admin user
POST {{host}}/auth/register
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "securepass123",
  "org_id": 1,
  "role": "admin"
}
HTTP 201
[Asserts]
jsonpath "$.user.id" isInteger
jsonpath "$.user.email" == "admin@zenflow.test"

# Step 3: Login
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "securepass123"
}
HTTP 200
[Asserts]
jsonpath "$.access_token" isString
[Captures]
token: jsonpath "$.access_token"

# Step 4: Create a workflow
POST {{host}}/api/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Onboarding Flow",
  "trigger_event": "user.signup"
}
HTTP 201
[Asserts]
jsonpath "$.workflow.id" isInteger
jsonpath "$.workflow.status" == "draft"
jsonpath "$.workflow.title" == "Onboarding Flow"
[Captures]
wf_id: jsonpath "$.workflow.id"

# Step 5: Add two actions
POST {{host}}/api/workflows/{{wf_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "send_email",
  "config": "{\"template\": \"welcome\"}",
  "sequence_order": 1
}
HTTP 201
[Asserts]
jsonpath "$.action.id" isInteger

POST {{host}}/api/workflows/{{wf_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "create_task",
  "config": "{\"task\": \"setup_profile\"}",
  "sequence_order": 2
}
HTTP 201

# Step 6: List actions
GET {{host}}/api/workflows/{{wf_id}}/actions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.actions" isCollection

# Step 7: Activate workflow
POST {{host}}/api/workflows/{{wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 8: Execute workflow
POST {{host}}/api/workflows/{{wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_log.status" == "success"
jsonpath "$.execution_log.credits_spent" == 1

# Step 9: Verify execution logs
GET {{host}}/api/workflows/{{wf_id}}/logs
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_logs" isCollection

# Step 10: List workflows
GET {{host}}/api/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" isCollection

# Step 11: Get workflow detail
GET {{host}}/api/workflows/{{wf_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 12: Pause workflow
POST {{host}}/api/workflows/{{wf_id}}/pause
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# Step 13: Re-activate workflow
POST {{host}}/api/workflows/{{wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 14: Archive workflow
POST {{host}}/api/workflows/{{wf_id}}/archive
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "archived"

# Step 15: Create new version of workflow
POST {{host}}/api/workflows/{{wf_id}}/new-version
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.status" == "draft"
jsonpath "$.workflow.version" == 2

# Step 15.5: Create webhook
POST {{host}}/api/webhooks
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "url": "https://example.com/hook",
  "event_type": "workflow.executed"
}
HTTP 201
[Asserts]
jsonpath "$.webhook.id" isInteger

# Step 15.6: List webhooks
GET {{host}}/api/webhooks
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.webhooks" isCollection

# Step 15.7: Delete webhook (create another first, then delete it)
POST {{host}}/api/webhooks
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "url": "https://example.com/hook-to-delete",
  "event_type": "workflow.executed"
}
HTTP 201
[Captures]
wh_id: jsonpath "$.webhook.id"

DELETE {{host}}/api/webhooks/{{wh_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.deleted" isBoolean

# Step 15.8: Publish template
POST {{host}}/api/templates
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "source_workflow_id": {{wf_id}},
  "title": "Onboarding Template",
  "description": "A template for onboarding workflows",
  "category": "onboarding"
}
HTTP 201
[Asserts]
jsonpath "$.template.id" isInteger
[Captures]
tmpl_id: jsonpath "$.template.id"

# Step 15.9: List templates
GET {{host}}/api/templates?per_page=10
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection

# Step 15.10: Get template
GET {{host}}/api/templates/{{tmpl_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.template.title" == "Onboarding Template"

# Step 15.11: Clone template
POST {{host}}/api/templates/{{tmpl_id}}/clone
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.title" == "Onboarding Template"
jsonpath "$.workflow.status" == "draft"

# Step 15.12: List audit logs
GET {{host}}/api/audit-logs?page=1&per_page=10&sort_by=created_at&sort_dir=desc
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection
jsonpath "$.total" isInteger

# Step 15.13: Dashboard
GET {{host}}/api/dashboard
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.summary.name" == "Test Org"
jsonpath "$.summary.plan_type" == "pro"

# Step 16: List workflow versions
GET {{host}}/api/workflows/{{wf_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" isCollection
