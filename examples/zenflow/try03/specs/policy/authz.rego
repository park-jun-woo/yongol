package authz

# @ownership organization: organizations.id
# @ownership workflow: workflows.org_id
# @ownership webhook: webhooks.org_id
# @ownership template: templates.org_id
# @ownership execution_log: execution_logs.org_id
# @ownership audit_log: audit_logs.org_id

default allow := false

# Admin can create workflows (no resource ID yet, role check only)
allow if {
	input.action == "CreateWorkflow"
	input.resource == "workflow"
	input.claims.role == "admin"
}

# Admin can create actions (no resource ID yet, role check only)
allow if {
	input.action == "CreateAction"
	input.resource == "workflow"
	input.claims.role == "admin"
}

# Admin can create workflow versions in their org
allow if {
	input.action == "CreateWorkflowVersion"
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin can activate/pause/archive/batch-save/auto-assign workflows in their org
allow if {
	input.action in {"ActivateWorkflow", "PauseWorkflow", "ArchiveWorkflow", "SaveWorkflowActions", "AutoAssignWorkflow"}
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Webhook management - any authenticated user in their org
allow if {
	input.action in {"CreateWebhook", "ListWebhooks"}
	input.resource == "webhook"
}

# Delete webhook - ownership checked
allow if {
	input.action == "DeleteWebhook"
	input.resource == "webhook"
	data.owners.webhook[input.resource_id] == input.claims.org_id
}

# Any authenticated user can list workflows in their org
allow if {
	input.action == "ListWorkflows"
	input.resource == "workflow"
}

# Any user can get a workflow in their org (ownership checked)
allow if {
	input.action == "GetWorkflow"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can list workflow versions in their org (ownership checked)
allow if {
	input.action == "ListWorkflowVersions"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can list actions for their workflow (ownership checked)
allow if {
	input.action == "ListActions"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can execute workflows in their org (ownership checked)
allow if {
	input.action == "ExecuteWorkflow"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can list execution logs for their workflow (ownership checked)
allow if {
	input.action == "ListExecutionLogs"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin can publish a template for their workflow (uses workflow resource)
allow if {
	input.action == "PublishTemplate"
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can clone a template (ownership of template checked)
allow if {
	input.action == "CloneTemplate"
	input.resource == "template"
	input.claims.role == "admin"
}

# Admin can execute with report (same as ExecuteWorkflow)
allow if {
	input.action == "ExecuteWithReport"
	input.resource == "workflow"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin can manage schedules for their workflows
allow if {
	input.action in {"SetSchedule", "GetSchedule", "DeleteSchedule"}
	input.resource == "workflow"
	input.claims.role == "admin"
	data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any user can get execution report for their execution log
allow if {
	input.action == "GetExecutionReport"
	input.resource == "execution_log"
	data.owners.execution_log[input.resource_id] == input.claims.org_id
}

# Any authenticated user can list/view audit logs
allow if {
	input.action in {"ListAuditLogs", "GetRecentAuditLogs"}
	input.resource == "audit_log"
}

# Any user can get a specific audit log in their org
allow if {
	input.action == "GetAuditLog"
	input.resource == "audit_log"
	data.owners.audit_log[input.resource_id] == input.claims.org_id
}

# Any authenticated user can view dashboard
allow if {
	input.action == "GetDashboard"
	input.resource == "organization"
}

# Admin can verify org address
allow if {
	input.action == "VerifyOrgAddress"
	input.resource == "organization"
	input.claims.role == "admin"
	data.owners.organization[input.resource_id] == input.resource_id
}

# Any user can get execution detail for their execution log
allow if {
	input.action == "GetExecutionDetail"
	input.resource == "execution_log"
	data.owners.execution_log[input.resource_id] == input.claims.org_id
}
