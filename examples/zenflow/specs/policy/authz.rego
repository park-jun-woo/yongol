package authz

import rego.v1

# @ownership workflow: workflows.org_id
# @ownership organization: organizations.id

default allow := false

is_same_org if {
	data.owners[input.resource][input.resource_id] == input.claims.org_id
}

is_admin if {
	input.claims.role == "admin"
}

# Organization-scope mutations (admin only).
allow if {
	input.action == "CreateWorkflow"
	input.resource == "organization"
	is_admin
	is_same_org
}

# Organization-scope reads (same-org, any role).
allow if {
	input.action == "ListWorkflows"
	input.resource == "organization"
	is_same_org
}

# Workflow reads (same-org, any role).
allow if {
	input.action == "GetWorkflow"
	input.resource == "workflow"
	is_same_org
}

allow if {
	input.action == "ExecuteWorkflow"
	input.resource == "workflow"
	is_same_org
}

# Workflow mutations (admin + same-org).
allow if {
	input.action == "AddAction"
	input.resource == "workflow"
	is_admin
	is_same_org
}

allow if {
	input.action == "ActivateWorkflow"
	input.resource == "workflow"
	is_admin
	is_same_org
}

allow if {
	input.action == "PauseWorkflow"
	input.resource == "workflow"
	is_admin
	is_same_org
}

allow if {
	input.action == "ArchiveWorkflow"
	input.resource == "workflow"
	is_admin
	is_same_org
}
