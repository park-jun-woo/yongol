package authz

# @ownership workflow: workflows.org_id

default allow = false

is_admin if {
    input.claims.role == "admin"
}

is_same_org if {
    input.claims.org_id == data.owners.workflow
}

allow if {
    input.action == "CreateWorkflow"
    is_admin
}

allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    is_admin
}

allow if {
    input.action == "GetWorkflow"
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
    is_admin
    is_same_org
}

allow if {
    input.action == "CreateAction"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ListExecutionLogs"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "CreateOrganization"
}

allow if {
    input.action == "Register"
}

allow if {
    input.action == "Login"
}

allow if {
    input.action == "BatchSaveActions"
    input.resource == "workflow"
    is_admin
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
}

allow if {
    input.action == "AutoAssignWorkflow"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ScheduleWorkflow"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "DeleteWorkflowSchedule"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ExecuteWorkflowWithReport"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "GetExecutionReport"
    input.resource == "execution_log"
}

allow if {
    input.action == "GetExecutionDetail"
    input.resource == "execution_log"
}

allow if {
    input.action == "CreateWebhook"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "ListWebhooks"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "DeleteWebhook"
    input.resource == "webhook"
    is_admin
}

allow if {
    input.action == "VerifyOrganizationAddress"
    input.resource == "organization"
    is_admin
}

allow if {
    input.action == "PublishTemplate"
    input.resource == "workflow"
    is_admin
    is_same_org
}

allow if {
    input.action == "ListTemplates"
}

allow if {
    input.action == "GetTemplate"
}

allow if {
    input.action == "CloneTemplate"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "GetDashboard"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "ListAuditLogs"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "GetAuditLog"
    input.resource == "user"
    is_admin
}

allow if {
    input.action == "GetRecentAuditLogs"
    input.resource == "user"
    is_admin
}
