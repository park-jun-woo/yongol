# Step 1: Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{"email": "admin@zenflow.test", "password": "password123"}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Step 2: Create workflow
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{"title": "My Workflow", "trigger_event": "on_signup"}
HTTP 201
[Captures]
workflow_id: jsonpath "$.workflow.id"
org_id: jsonpath "$.workflow.org_id"
[Asserts]
jsonpath "$.workflow.title" == "My Workflow"
jsonpath "$.workflow.status" == "draft"

# Step 3: List workflows
GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 1

# Step 4: Get workflow
GET {{host}}/workflows/{{workflow_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.title" == "My Workflow"

# Step 5: Add action 1
POST {{host}}/workflows/{{workflow_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{"action_type": "email", "config": "send welcome", "sequence_order": 1}
HTTP 201
[Asserts]
jsonpath "$.action.action_type" == "email"

# Step 6: Add action 2
POST {{host}}/workflows/{{workflow_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{"action_type": "webhook", "config": "notify slack", "sequence_order": 2}
HTTP 201

# Step 7: Activate workflow
POST {{host}}/workflows/{{workflow_id}}/activate
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 8: Execute workflow
POST {{host}}/workflows/{{workflow_id}}/execute
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_log.status" == "completed"
jsonpath "$.execution_log.credits_spent" == 1

# Step 9: List execution logs
GET {{host}}/workflows/{{workflow_id}}/execution-logs
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.execution_logs" count >= 1

# Step 10: Pause workflow
POST {{host}}/workflows/{{workflow_id}}/pause
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "paused"

# Step 11: Re-activate workflow
POST {{host}}/workflows/{{workflow_id}}/activate
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "active"

# Step 12: Archive workflow
POST {{host}}/workflows/{{workflow_id}}/archive
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.status" == "archived"

# Step 13: Create new version (need a new workflow for this since archived can't be versioned)
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{"title": "Version Test WF", "trigger_event": "on_test"}
HTTP 201
[Captures]
ver_wf_id: jsonpath "$.workflow.id"

# Step 14: Create workflow version
POST {{host}}/workflows/{{ver_wf_id}}/new-version
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.version" == 2

# Step 15: List workflow versions
GET {{host}}/workflows/{{ver_wf_id}}/versions
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflows" count >= 2

# Step 16: Create webhook
POST {{host}}/webhooks
Content-Type: application/json
Authorization: Bearer {{token}}
{"url": "https://example.com/hook", "event_type": "workflow.executed"}
HTTP 201
[Captures]
hook_id: jsonpath "$.webhook.id"
[Asserts]
jsonpath "$.webhook.url" == "https://example.com/hook"

# Step 17: List webhooks
GET {{host}}/webhooks
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.webhooks" count >= 1

# Step 18: Delete webhook
DELETE {{host}}/webhooks/{{hook_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.ok" == true

# Step 19: Create workflow for template
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{"title": "Template Source WF", "trigger_event": "on_template"}
HTTP 201
[Captures]
tpl_wf_id: jsonpath "$.workflow.id"

# Step 20: Publish template
POST {{host}}/templates
Content-Type: application/json
Authorization: Bearer {{token}}
{"workflow_id": "{{tpl_wf_id}}", "title": "My Template", "description": "A workflow template", "category": "automation"}
HTTP 201
[Captures]
tpl_id: jsonpath "$.template.id"
[Asserts]
jsonpath "$.template.title" == "My Template"

# Step 21: List templates
GET {{host}}/templates
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 1

# Step 22: Get template
GET {{host}}/templates/{{tpl_id}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.template.title" == "My Template"

# Step 23: Clone template
POST {{host}}/templates/{{tpl_id}}/clone
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 201
[Asserts]
jsonpath "$.workflow.title" == "My Template"

# Step 24: Create workflow for execute-with-report
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{"title": "Report WF", "trigger_event": "on_report"}
HTTP 201
[Captures]
rpt_wf_id: jsonpath "$.workflow.id"

# Step 25: Activate for report test
POST {{host}}/workflows/{{rpt_wf_id}}/activate
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200

# Step 26: Execute with report
POST {{host}}/workflows/{{rpt_wf_id}}/execute-with-report
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Captures]
rpt_log_id: jsonpath "$.execution_log.id"
[Asserts]
jsonpath "$.execution_log.status" == "completed"

# Step 27: Get execution report
GET {{host}}/execution-logs/{{rpt_log_id}}/report
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.content" isString

# Step 28: Set schedule
POST {{host}}/workflows/{{rpt_wf_id}}/schedule
Content-Type: application/json
Authorization: Bearer {{token}}
{"cron": "0 9 * * 1"}
HTTP 200
[Asserts]
jsonpath "$.schedule.cron" == "0 9 * * 1"

# Step 29: Get schedule
GET {{host}}/workflows/{{rpt_wf_id}}/schedule
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.schedule.value" isString

# Step 30: Delete schedule
DELETE {{host}}/workflows/{{rpt_wf_id}}/schedule
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.ok" == true

# Step 31: List audit logs
GET {{host}}/audit-logs
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 0
jsonpath "$.total" >= 0

# Step 32: Get recent audit logs
GET {{host}}/audit-logs/recent
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 0

# Step 33: Get audit log (use rpt_log_id which is not an audit_log; expect 403 from ownership check)
GET {{host}}/audit-logs/{{rpt_log_id}}
Authorization: Bearer {{token}}
HTTP 403

# Step 34: Get dashboard
GET {{host}}/dashboard
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.summary.org_name" isString
jsonpath "$.summary.credits_balance" isInteger

# Step 35: Batch save actions on ver_wf_id
PUT {{host}}/workflows/{{ver_wf_id}}/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{"actions": [{"action_type": "email", "config": "batch1", "sequence_order": 1}, {"action_type": "sms", "config": "batch2", "sequence_order": 2}]}
HTTP 200
[Asserts]
jsonpath "$.ok" == true

# Step 36: Get execution detail
GET {{host}}/execution-logs/{{rpt_log_id}}/detail
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detail.status" isString

# Step 37: Verify org address
POST {{host}}/organizations/{{org_id}}/verify-address
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.organization.address_verified" == true

# Step 38: Auto-assign workflow
POST {{host}}/workflows/{{ver_wf_id}}/auto-assign
Content-Type: application/json
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.workflow.assignment_confidence" isString
