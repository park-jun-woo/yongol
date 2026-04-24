package authz

default allow := false

allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
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
    input.action in {"ActivateWorkflow", "ExecuteWorkflow"}
    input.resource == "workflow"
    input.claims.role == "admin"
}
