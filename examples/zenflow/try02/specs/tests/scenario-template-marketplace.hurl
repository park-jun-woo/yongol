# ZenFlow Scenario — Workflow Template Marketplace
# Org A: create workflow → publish template → duplicate publish attempt (409)
# Org B: list templates (cursor pagination) → get template → clone template → confirm own workflow created

# Step 1: Login as Org A admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token_a: jsonpath "$.access_token"

# Step 2: Org A — create a workflow to publish
POST {{host}}/workflows
Authorization: Bearer {{token_a}}
Content-Type: application/json
{
  "title": "Marketplace Workflow",
  "trigger_event": "user.signup"
}

HTTP 201
[Captures]
src_workflow_id: jsonpath "$.id"
[Asserts]
jsonpath "$.status" == "draft"

# Step 3: Org A — publish as template
POST {{host}}/templates
Authorization: Bearer {{token_a}}
Content-Type: application/json
{
  "workflow_id": "{{src_workflow_id}}",
  "title": "Email Onboarding Template",
  "description": "A workflow for onboarding new users",
  "category": "onboarding"
}

HTTP 201
[Captures]
template_id: jsonpath "$.id"
[Asserts]
jsonpath "$.title" == "Email Onboarding Template"
jsonpath "$.category" == "onboarding"
jsonpath "$.clone_count" == 0

# Step 4: Org A — duplicate publish attempt must be 409
POST {{host}}/templates
Authorization: Bearer {{token_a}}
Content-Type: application/json
{
  "workflow_id": "{{src_workflow_id}}",
  "title": "Duplicate Template",
  "description": "Should fail",
  "category": "onboarding"
}

HTTP 409

# Step 5: Get template detail
GET {{host}}/templates/{{template_id}}
Authorization: Bearer {{token_a}}

HTTP 200
[Asserts]
jsonpath "$.id" == "{{template_id}}"
jsonpath "$.title" == "Email Onboarding Template"

# Step 6: List templates — first page (no cursor)
GET {{host}}/templates?per_page=20
Authorization: Bearer {{token_a}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 7: List templates — with category filter
GET {{host}}/templates?category=onboarding&per_page=20
Authorization: Bearer {{token_a}}

HTTP 200
[Asserts]
jsonpath "$.items" count > 0

# Step 8: Login as Org B admin
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin-b@zenflow.test",
  "password": "testpassword"
}

HTTP 200
[Captures]
token_b: jsonpath "$.access_token"

# Step 9: Org B — list templates (cursor pagination)
GET {{host}}/templates?per_page=10
Authorization: Bearer {{token_b}}

HTTP 200
[Captures]
first_item_id: jsonpath "$.items[0].id"
[Asserts]
jsonpath "$.items" count > 0

# Step 10: Org B — list templates with cursor
GET {{host}}/templates?cursor={{first_item_id}}&per_page=10
Authorization: Bearer {{token_b}}

HTTP 200

# Step 11: Org B — clone the template
POST {{host}}/templates/{{template_id}}/clone
Authorization: Bearer {{token_b}}

HTTP 201
[Asserts]
jsonpath "$.title" == "Marketplace Workflow"
jsonpath "$.status" == "draft"
