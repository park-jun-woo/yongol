-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event)
VALUES (@org_id, @title, @trigger_event)
RETURNING *;

-- name: WorkflowListByOrg :many
-- +no-pagination
SELECT * FROM workflows WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;

-- name: WorkflowCreateVersion :one
INSERT INTO workflows (org_id, title, trigger_event, version, root_workflow_id)
VALUES (@org_id, @title, @trigger_event, @version, @root_workflow_id)
RETURNING *;

-- name: WorkflowAutoAssign :exec
UPDATE workflows
SET assigned_to = CASE WHEN @member_id::text = '' THEN assigned_to ELSE @member_id END,
    assignment_confidence = @confidence
WHERE id = @id;

-- name: WorkflowListVersions :many
-- +no-pagination
SELECT * FROM workflows WHERE (root_workflow_id = @root_id OR id = @root_id) AND org_id = @org_id ORDER BY version ASC;
