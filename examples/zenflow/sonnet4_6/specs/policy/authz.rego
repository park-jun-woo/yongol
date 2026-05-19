package authz

# @ownership workflow: workflows.org_id
# @ownership webhook: webhooks.org_id
# @ownership template: templates.org_id

default allow := false

allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
}

allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "AddAction"
    input.resource == "workflow"
}

allow if {
    input.action == "ActivateWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "PauseWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "ArchiveWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
}

allow if {
    input.action == "ListExecutionLogs"
    input.resource == "workflow"
}

allow if {
    input.action == "CreateWorkflowVersion"
    input.resource == "workflow"
}

allow if {
    input.action == "ListWorkflowVersions"
    input.resource == "workflow"
}

allow if {
    input.action == "CreateWebhook"
    input.resource == "webhook"
}

allow if {
    input.action == "ListWebhooks"
    input.resource == "webhook"
}

allow if {
    input.action == "DeleteWebhook"
    input.resource == "webhook"
}

allow if {
    input.action == "PublishTemplate"
    input.resource == "template"
}

allow if {
    input.action == "ListTemplates"
    input.resource == "template"
}

allow if {
    input.action == "GetTemplate"
    input.resource == "template"
}

allow if {
    input.action == "CloneTemplate"
    input.resource == "template"
}

allow if {
    input.action == "ExecuteWithReport"
    input.resource == "workflow"
}

allow if {
    input.action == "GetExecutionReport"
    input.resource == "execution"
}

allow if {
    input.action == "SetSchedule"
    input.resource == "workflow"
}

allow if {
    input.action == "GetSchedule"
    input.resource == "workflow"
}

allow if {
    input.action == "DeleteSchedule"
    input.resource == "workflow"
}

allow if {
    input.action == "ListAuditLogs"
    input.resource == "audit"
}

allow if {
    input.action == "GetRecentAuditLogs"
    input.resource == "audit"
}

allow if {
    input.action == "GetDashboard"
    input.resource == "dashboard"
}

allow if {
    input.action == "GetAuditLog"
    input.resource == "audit"
}

allow if {
    input.action == "GetExecutionDetail"
    input.resource == "execution"
}

allow if {
    input.action == "SaveWorkflowActions"
    input.resource == "workflow"
}

allow if {
    input.action == "VerifyOrgAddress"
    input.resource == "organization"
}

allow if {
    input.action == "AutoAssignWorkflow"
    input.resource == "workflow"
}
