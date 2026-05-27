# Login as admin to get token
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd!"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# GetWorkflow — 404 nonexistent
GET {{host}}/workflows/999999
Authorization: Bearer {{token}}
HTTP 404

# AddAction — 404 nonexistent workflow
POST {{host}}/workflows/999999/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "action_type": "send_email",
  "config": "{}",
  "sequence_order": 1
}
HTTP 404

# SaveWorkflowActions — 404 nonexistent workflow
PUT {{host}}/workflows/999999/actions
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "actions_json": "[]"
}
HTTP 404

# ActivateWorkflow — 404 nonexistent
POST {{host}}/workflows/999999/activate
Authorization: Bearer {{token}}
HTTP 404

# PauseWorkflow — 404 nonexistent
POST {{host}}/workflows/999999/pause
Authorization: Bearer {{token}}
HTTP 404

# ArchiveWorkflow — 404 nonexistent
POST {{host}}/workflows/999999/archive
Authorization: Bearer {{token}}
HTTP 404

# ExecuteWorkflow — 404 nonexistent
POST {{host}}/workflows/999999/execute
Authorization: Bearer {{token}}
HTTP 404

# ExecuteWithReport — 404 nonexistent
POST {{host}}/workflows/999999/execute-with-report
Authorization: Bearer {{token}}
HTTP 404

# ListExecutionLogs — 404 nonexistent workflow
GET {{host}}/workflows/999999/execution-logs
Authorization: Bearer {{token}}
HTTP 404

# CreateWorkflowVersion — 404 nonexistent
POST {{host}}/workflows/999999/new-version
Authorization: Bearer {{token}}
HTTP 404

# ListWorkflowVersions — 404 nonexistent
GET {{host}}/workflows/999999/versions
Authorization: Bearer {{token}}
HTTP 404

# AutoAssignWorkflow — 404 nonexistent
POST {{host}}/workflows/999999/auto-assign
Authorization: Bearer {{token}}
HTTP 404

# DeleteWebhook — 404 nonexistent
DELETE {{host}}/webhooks/999999
Authorization: Bearer {{token}}
HTTP 404

# GetTemplate — 404 nonexistent
GET {{host}}/templates/999999
Authorization: Bearer {{token}}
HTTP 404

# CloneTemplate — 404 nonexistent
POST {{host}}/templates/999999/clone
Authorization: Bearer {{token}}
HTTP 404

# GetExecutionReport — 404 nonexistent
GET {{host}}/execution-logs/999999/report
Authorization: Bearer {{token}}
HTTP 404

# GetExecutionDetail — 404 nonexistent
GET {{host}}/execution-logs/999999/detail
Authorization: Bearer {{token}}
HTTP 404

# SetSchedule — 404 nonexistent
POST {{host}}/workflows/999999/schedule
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "cron": "0 9 * * 1"
}
HTTP 404

# GetSchedule — 404 nonexistent
GET {{host}}/workflows/999999/schedule
Authorization: Bearer {{token}}
HTTP 404

# DeleteSchedule — 404 nonexistent
DELETE {{host}}/workflows/999999/schedule
Authorization: Bearer {{token}}
HTTP 404

# GetAuditLog — 404 nonexistent
GET {{host}}/audit-logs/999999
Authorization: Bearer {{token}}
HTTP 404

# GetDashboard — 404 nonexistent org
GET {{host}}/dashboard
Authorization: Bearer {{token}}
HTTP 404

# VerifyOrgAddress — 404 nonexistent org
POST {{host}}/organizations/999999/verify-address
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "address": "123 Test Street"
}
HTTP 404
