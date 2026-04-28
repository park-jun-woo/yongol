package authz

# @ownership workflow: workflows.org_id

default allow := false

# CreateWorkflow — admin only (no resource yet)
allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
}

# ListWorkflows — any user in the same org (collection scope)
allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
}

# ActivateWorkflow — admin and same org as workflow.org_id
allow if {
    input.action == "ActivateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ExecuteWorkflow — admin in the owning org
allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# ExecuteWorkflow — member in the owning org
allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "member"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}
