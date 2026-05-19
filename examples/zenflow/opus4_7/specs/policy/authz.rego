package authz

# @ownership workflow: workflows.org_id
# @ownership execution_log: execution_logs.org_id
# @ownership template: templates.org_id

default allow := false

# Login — no auth required (security: [] in OpenAPI)

# CreateWorkflow — admin only
allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
}

# ListWorkflows — any authenticated user in same org
allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
}

# GetWorkflow — same org (ownership check)
allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# AddAction — admin + same org
allow if {
    input.action == "AddAction"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ActivateWorkflow — admin + same org
allow if {
    input.action == "ActivateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# PauseWorkflow — admin + same org
allow if {
    input.action == "PauseWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ArchiveWorkflow — admin + same org
allow if {
    input.action == "ArchiveWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ExecuteWorkflow — admin, same org
allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ExecuteWorkflow — member, same org
allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "member"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ListExecutionLogs — same org (ownership check)
allow if {
    input.action == "ListExecutionLogs"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# CreateWorkflowVersion — admin + same org
allow if {
    input.action == "CreateWorkflowVersion"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ListWorkflowVersions — same org
allow if {
    input.action == "ListWorkflowVersions"
    input.resource == "workflow"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# PublishTemplate — admin only
allow if {
    input.action == "PublishTemplate"
    input.resource == "template"
    input.claims.role == "admin"
}

# ListTemplates — any authenticated user
allow if {
    input.action == "ListTemplates"
    input.resource == "template"
}

# GetTemplate — any authenticated user
allow if {
    input.action == "GetTemplate"
    input.resource == "template"
}

# CloneTemplate — any authenticated user
allow if {
    input.action == "CloneTemplate"
    input.resource == "template"
}
