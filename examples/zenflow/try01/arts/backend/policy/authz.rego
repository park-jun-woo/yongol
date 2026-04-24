package authz

default allow := false

# Admins can manage workflows.
allow if {
    input.action in {"CreateWorkflow", "ListWorkflows", "ActivateWorkflow"}
    input.resource == "workflow"
    input.claims.role == "admin"
}
