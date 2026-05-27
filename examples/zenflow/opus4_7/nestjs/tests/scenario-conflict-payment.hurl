# Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd!"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Create workflow for state conflict tests
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Conflict Test WF",
  "trigger_event": "on_conflict"
}
HTTP 201
[Captures]
conflict_wf_id: jsonpath "$.workflow.id"

# Activate the workflow (draft -> active)
POST {{host}}/workflows/{{conflict_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 200

# ActivateWorkflow — 409 already active
POST {{host}}/workflows/{{conflict_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 409

# Archive the workflow (active -> archived)
POST {{host}}/workflows/{{conflict_wf_id}}/archive
Authorization: Bearer {{token}}
HTTP 200

# ArchiveWorkflow — 409 already archived
POST {{host}}/workflows/{{conflict_wf_id}}/archive
Authorization: Bearer {{token}}
HTTP 409

# PauseWorkflow — 409 cannot pause an archived workflow
POST {{host}}/workflows/{{conflict_wf_id}}/pause
Authorization: Bearer {{token}}
HTTP 409

# ExecuteWorkflow — 409 cannot execute an archived workflow
POST {{host}}/workflows/{{conflict_wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 409

# ExecuteWithReport — 409 cannot execute an archived workflow
POST {{host}}/workflows/{{conflict_wf_id}}/execute-with-report
Authorization: Bearer {{token}}
HTTP 409

# ActivateWorkflow — 402 insufficient credits (separate workflow)
POST {{host}}/workflows
Content-Type: application/json
Authorization: Bearer {{token}}
{
  "title": "Credits Test WF",
  "trigger_event": "on_credits"
}
HTTP 201
[Captures]
credits_wf_id: jsonpath "$.workflow.id"

POST {{host}}/workflows/{{credits_wf_id}}/activate
Authorization: Bearer {{token}}
HTTP 402

# ExecuteWorkflow — 402 insufficient credits
POST {{host}}/workflows/{{credits_wf_id}}/execute
Authorization: Bearer {{token}}
HTTP 402

# ExecuteWithReport — 402 insufficient credits
POST {{host}}/workflows/{{credits_wf_id}}/execute-with-report
Authorization: Bearer {{token}}
HTTP 402
