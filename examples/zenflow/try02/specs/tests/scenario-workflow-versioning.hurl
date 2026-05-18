# ZenFlow Scenario — Workflow Versioning
# Admin creates a workflow, adds 2 actions, creates a new version, then lists versions.

# Step 1: Login as admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Step 2: Create a workflow (v1)
POST {{host}}/workflows
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Versioned Workflow",
  "trigger_event": "user.signup"
}

HTTP 201
[Captures]
workflow_id: jsonpath "$.id"
[Asserts]
jsonpath "$.status" == "draft"
jsonpath "$.version" == 1

# Step 3: Add first action to v1
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "send_email",
  "config": "{}",
  "sequence_order": 1
}

HTTP 201

# Step 4: Add second action to v1
POST {{host}}/workflows/{{workflow_id}}/actions
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "action_type": "notify_slack",
  "config": "{}",
  "sequence_order": 2
}

HTTP 201

# Step 5: Create new version (v2)
POST {{host}}/workflows/{{workflow_id}}/new-version
Authorization: Bearer {{token}}

HTTP 201
[Asserts]
jsonpath "$.status" == "draft"
jsonpath "$.version" == 2
jsonpath "$.title" == "Versioned Workflow"

# Step 6: List all versions from original workflow id
GET {{host}}/workflows/{{workflow_id}}/versions
Authorization: Bearer {{token}}

HTTP 200
[Asserts]
jsonpath "$.items" count == 2
