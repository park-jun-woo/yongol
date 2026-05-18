package authz

# @ownership workflow: workflows.org_id
# @ownership webhook: webhooks.org_id

default allow := false

# Admin can create workflows (no resource_id yet)
allow if {
	input.action == "CreateWorkflow"
	input.resource == "workflow"
	input.claims.role == "admin"
}

# Admin can list workflows (no resource_id for list)
allow if {
	input.action == "ListWorkflows"
	input.resource == "workflow"
	input.claims.role == "admin"
}

# Admin can do everything on specific workflows in their org
allow if {
	input.action in {"GetWorkflow", "ActivateWorkflow", "PauseWorkflow", "ArchiveWorkflow", "ExecuteWorkflow", "CreateAction", "ListActions", "ListExecutionLogs", "CreateWorkflowVersion", "ListWorkflowVersions"}
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Members can list workflows (no resource_id for list)
allow if {
	input.action == "ListWorkflows"
	input.resource == "workflow"
	input.claims.role == "member"
}

# Any role can view dashboard
allow if {
	input.action == "GetDashboard"
	input.resource == "dashboard"
	input.claims.role in {"admin", "member"}
}

# Admin can list audit logs
allow if {
	input.action == "ListAuditLogs"
	input.resource == "audit_log"
	input.claims.role == "admin"
}

# Admin can publish templates (uses workflow resource_id)
allow if {
	input.action == "PublishTemplate"
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any role can clone templates (no specific resource ownership check)
allow if {
	input.action == "CloneTemplate"
	input.resource == "template"
	input.claims.role in {"admin", "member"}
}

# Admin can manage webhooks (create, list — no resource_id)
allow if {
	input.action in {"CreateWebhook", "ListWebhooks"}
	input.resource == "webhook"
	input.claims.role == "admin"
}

# Admin can delete specific webhooks in their org
allow if {
	input.action == "DeleteWebhook"
	input.resource == "webhook"
	input.claims.role == "admin"
	data.owners.webhook[input.resource_id] == input.claims.org_id
}

# Members can read specific workflows and list actions/logs/versions in their org
allow if {
	input.action in {"GetWorkflow", "ListActions", "ListExecutionLogs", "ListWorkflowVersions"}
	input.resource == "workflow"
	input.claims.role == "member"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}
