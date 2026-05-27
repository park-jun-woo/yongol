-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event)
VALUES (@org_id, @title, @trigger_event)
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowListByOrgID :many
-- +no-pagination
SELECT * FROM workflows WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;

-- name: WorkflowCreateVersion :one
INSERT INTO workflows (org_id, title, trigger_event, version, root_workflow_id)
VALUES (@org_id, @title, @trigger_event, @version, @root_workflow_id)
RETURNING *;

-- name: WorkflowListVersions :many
-- +no-pagination
SELECT * FROM workflows WHERE (root_workflow_id = @root_id OR id = @root_id) AND org_id = @org_id ORDER BY version ASC;

-- name: WorkflowAutoAssign :exec
UPDATE workflows
SET assigned_to = sqlc.arg(member_id)::bigint,
    assignment_confidence = @confidence
WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;
