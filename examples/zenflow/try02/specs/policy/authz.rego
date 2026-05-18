package authz

# @ownership workflow: workflows.org_id
# @ownership webhook: webhooks.org_id
# @ownership template: templates.org_id
# @ownership execution_log: execution_logs.org_id
# @ownership audit_log: audit_logs.org_id
# @ownership organization: organizations.id

import future.keywords.if

default allow = false

is_admin if {
    input.claims.role == "admin"
}

is_same_org if {
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

is_authenticated if {
    input.claims.org_id != ""
    input.claims.org_id != null
}

allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    is_admin
    is_authenticated
}

allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    input.claims.email != ""
    is_authenticated
}

allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
    is_same_org
}

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

allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    is_same_org
}

allow if {
    input.action == "ExecuteWithReport"
    input.resource == "workflow"
    is_same_org
}

allow if {
    input.action == "ListExecutionLogs"
    input.resource == "workflow"
    is_same_org
}

allow if {
    input.action == "CreateWorkflowVersion"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ListWorkflowVersions"
    input.resource == "workflow"
    is_same_org
}

is_same_org_execution_log if {
    data.owners.execution_log[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "GetExecutionReport"
    input.resource == "execution_log"
    is_same_org_execution_log
}

is_same_org_webhook if {
    data.owners.webhook[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "CreateWebhook"
    input.resource == "webhook"
    is_admin
    is_authenticated
}

allow if {
    input.action == "ListWebhooks"
    input.resource == "webhook"
    is_authenticated
}

allow if {
    input.action == "DeleteWebhook"
    input.resource == "webhook"
    is_admin
    is_same_org_webhook
}

is_same_org_template if {
    data.owners.template[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "PublishTemplate"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ListTemplates"
    input.resource == "template"
    is_authenticated
}

allow if {
    input.action == "GetTemplate"
    input.resource == "template"
    is_authenticated
}

allow if {
    input.action == "CloneTemplate"
    input.resource == "template"
    is_authenticated
}

allow if {
    input.action == "SetSchedule"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "GetSchedule"
    input.resource == "workflow"
    is_same_org
}

allow if {
    input.action == "DeleteSchedule"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ListAuditLogs"
    input.resource == "audit_log"
    is_admin
    is_authenticated
}

allow if {
    input.action == "GetRecentAuditLogs"
    input.resource == "audit_log"
    is_admin
    is_authenticated
}

is_same_org_audit_log if {
    data.owners.audit_log[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "GetAuditLog"
    input.resource == "audit_log"
    is_admin
    is_same_org_audit_log
}

allow if {
    input.action == "GetDashboard"
    input.resource == "dashboard"
    is_authenticated
}

allow if {
    input.action == "GetExecutionDetail"
    input.resource == "execution_log"
    is_same_org_execution_log
}

allow if {
    input.action == "SaveWorkflowActions"
    input.resource == "workflow"
    is_admin
    is_same_org
}

is_same_org_organization if {
    data.owners.organization[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "VerifyOrgAddress"
    input.resource == "organization"
    is_admin
    is_same_org_organization
}

allow if {
    input.action == "AutoAssignWorkflow"
    input.resource == "workflow"
    is_admin
    is_same_org
}
