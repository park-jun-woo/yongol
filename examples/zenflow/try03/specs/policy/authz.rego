package authz

# @ownership workflow: workflows.org_id

default allow := false

# Create does not carry a ResourceID — ownership check is N/A,
# admin role suffices. The OrgID used by the handler is taken from
# currentUser.OrgID in the SSaC sequence.
allow if {
    input.action == "CreateWorkflow"
    input.resource == "workflow"
    input.claims.role == "admin"
}

# List scopes by caller's org automatically (handler filters by
# currentUser.OrgID). No ownership check needed.
allow if {
    input.action == "ListWorkflows"
    input.resource == "workflow"
    input.claims.role in {"admin", "member"}
}

# Read of a single workflow is allowed to any member of the owning org.
allow if {
    input.action == "GetWorkflow"
    input.resource == "workflow"
    input.claims.role in {"admin", "member"}
    data.owners.workflow[input.resource_id] == input.claims.org_id
}

# Mutating a workflow requires admin AND owning org.
allow if {
    input.action in {"ActivateWorkflow", "ExecuteWorkflow"}
    input.resource == "workflow"
    input.claims.role == "admin"
    data.owners.workflow[input.resource_id] == input.claims.org_id
}
