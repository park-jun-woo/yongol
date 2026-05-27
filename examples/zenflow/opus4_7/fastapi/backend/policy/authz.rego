package authz

# @ownership organization: organizations.id
# @ownership workflow: workflows.org_id
# @ownership execution_log: execution_logs.org_id
# @ownership webhook: webhooks.org_id
# @ownership template: templates.org_id
# @ownership audit_log: audit_logs.org_id

default allow := false

# Admin can create workflows
allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
}

# Same org can list workflows
allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
}

# Same org can get workflow
allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin + same org can add action
allow if {
    input.action == "AddAction"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin + same org can activate workflow
allow if {
    input.action == "ActivateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Same org can pause workflow
allow if {
    input.action == "PauseWorkflow"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Same org can archive workflow
allow if {
    input.action == "ArchiveWorkflow"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Same org can execute workflow
allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Same org can list execution logs
allow if {
    input.action == "ListExecutionLogs"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Members can view their own workflows
allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
    input.claims.role == "member"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin + same org can create workflow version
allow if {
    input.action == "CreateWorkflowVersion"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Same org can list workflow versions
allow if {
    input.action == "ListWorkflowVersions"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any authenticated user can create webhooks
allow if {
    input.action == "CreateWebhook"
    input.resource == "webhook"
}

# Any authenticated user can list webhooks
allow if {
    input.action == "ListWebhooks"
    input.resource == "webhook"
}

# Same org can delete webhook
allow if {
    input.action == "DeleteWebhook"
    input.resource == "webhook"
    data.owners.webhook[input.resource_id] == input.claims.org_id
}

# Same org can manage schedules
allow if {
    input.action in {"SetSchedule", "GetSchedule", "DeleteSchedule"}
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any authenticated user can get execution report
allow if {
    input.action == "GetExecutionReport"
    input.resource == "execution_log"
}

# Any authenticated user can publish template
allow if {
    input.action == "PublishTemplate"
    input.resource == "template"
}

# Any authenticated user can list audit logs
allow if {
    input.action in {"ListAuditLogs", "GetRecentAuditLogs"}
    input.resource == "audit_log"
}

# Any authenticated user can clone template
allow if {
    input.action == "CloneTemplate"
    input.resource == "template"
}

# Admin + same org can auto-assign workflow
allow if {
    input.action == "AutoAssignWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Admin + same org can verify org address
allow if {
    input.action == "VerifyOrgAddress"
    input.resource == "organization"
    input.claims.role == "admin"
    data.owners.organization[input.resource_id] == input.claims.org_id
}

# Admin + same org can save workflow actions
allow if {
    input.action == "SaveWorkflowActions"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Any authenticated user can view dashboard
allow if {
    input.action == "GetDashboard"
    input.resource == "organization"
}

# Same org can get audit log detail
allow if {
    input.action == "GetAuditLog"
    input.resource == "audit_log"
    data.owners.audit_log[input.resource_id] == input.claims.org_id
}

# Any authenticated user can get execution detail
allow if {
    input.action == "GetExecutionDetail"
    input.resource == "execution_log"
    data.owners.execution_log[input.resource_id] == input.claims.org_id
}
