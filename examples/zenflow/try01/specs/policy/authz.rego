package authz

# @ownership workflow: workflows.org_id
# @ownership execution_log: execution_logs.org_id
# @ownership webhook: webhooks.org_id
# @ownership audit_log: audit_logs.org_id
# @ownership template: templates.org_id
# @ownership organization: organizations.id

default allow := false

# Admin can do anything
allow if {
    input.action in {"CreateWorkflow"}
    input.resource == "workflow"
    input.claims.role == "admin"
}

# Audit log — any authenticated user in same org
allow if {
    input.action in {"ListAuditLogs", "GetRecentAuditLogs"}
    input.resource == "audit_log"
    input.claims.role in {"admin", "member"}
}

allow if {
    input.action == "GetAuditLog"
    input.resource == "audit_log"
    input.claims.role in {"admin", "member"}
    input.claims.org_id == data.owners.audit_log[input.resource_id]
}

# Execution log — same-org access
allow if {
    input.action in {"GetExecutionReport", "GetExecutionDetail"}
    input.resource == "execution_log"
    input.claims.org_id == data.owners.execution_log[input.resource_id]
}

# Webhook management — admin only
allow if {
    input.action in {"CreateWebhook", "ListWebhooks"}
    input.resource == "webhook"
    input.claims.role == "admin"
}

allow if {
    input.action == "DeleteWebhook"
    input.resource == "webhook"
    input.claims.role == "admin"
    input.claims.org_id == data.owners.webhook[input.resource_id]
}

# Template marketplace — any authenticated user
allow if {
    input.action in {"PublishTemplate", "ListTemplates", "CloneTemplate"}
    input.resource == "template"
    input.claims.role in {"admin", "member"}
}

allow if {
    input.action == "GetTemplate"
    input.resource == "template"
    input.claims.role in {"admin", "member"}
    input.claims.org_id == data.owners.template[input.resource_id]
}

# Organization management
allow if {
    input.action == "VerifyOrgAddress"
    input.resource == "organization"
    input.claims.role == "admin"
    input.claims.org_id == data.owners.organization[input.resource_id]
}

# Allow any authenticated user to view templates
allow if {
    input.action == "GetTemplate"
    input.resource == "template"
    input.claims.role in {"admin", "member"}
}

# Same-org access for listing and getting workflows
allow if {
    input.action in {"ListWorkflows", "GetWorkflow", "ListExecutionLogs", "ListWorkflowVersions"}
    input.resource == "workflow"
    input.claims.org_id == data.owners.workflow[input.resource_id]
}

# Collection-level list (no resource_id) — just check authenticated
allow if {
    input.action in {"ListWorkflows", "GetDashboard"}
    input.resource == "workflow"
    input.claims.org_id != ""
}

# Admin same-org for state transitions and execution
allow if {
    input.action in {"ActivateWorkflow", "PauseWorkflow", "ArchiveWorkflow", "ExecuteWorkflow", "ExecuteWithReport", "AddAction", "CreateWorkflowVersion", "SetSchedule", "GetSchedule", "DeleteSchedule", "SaveWorkflowActions", "AutoAssignWorkflow"}
    input.resource == "workflow"
    input.claims.role == "admin"
    input.claims.org_id == data.owners.workflow[input.resource_id]
}

# Member can view workflows and execution logs in their org
allow if {
    input.action in {"GetWorkflow", "ListWorkflows", "ListExecutionLogs", "ListWorkflowVersions"}
    input.resource == "workflow"
    input.claims.role == "member"
    input.claims.org_id == data.owners.workflow[input.resource_id]
}

# Collection-level list for members
allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    input.claims.role == "member"
    input.claims.org_id != ""
}
