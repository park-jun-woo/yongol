package authz

# @ownership workflow: workflows.org_id

default allow := false

allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
}

allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    input.claims.role == "admin"
}

allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    input.claims.role == "member"
}

allow if {
    input.action == "ActivateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

allow if {
    input.action == "ExecuteWorkflow"
    input.resource == "workflow"
    input.claims.role == "member"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}
