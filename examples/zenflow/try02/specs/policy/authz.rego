package authz

# @ownership workflow: workflows.org_id
# @ownership webhook: webhooks.org_id

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
