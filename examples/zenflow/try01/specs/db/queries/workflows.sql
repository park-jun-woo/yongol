-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, 'draft')
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowListByOrg :many
-- +no-pagination
SELECT * FROM workflows WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;

-- name: WorkflowCreateVersion :one
INSERT INTO workflows (org_id, title, trigger_event, status, version, root_workflow_id)
VALUES (@org_id, @title, @trigger_event, 'draft', @version, @root_workflow_id)
RETURNING *;

-- name: WorkflowListVersions :many
-- +no-pagination
SELECT * FROM workflows WHERE (root_workflow_id = @root_id OR id = @root_id) AND org_id = @org_id ORDER BY version ASC;

-- name: WorkflowAutoAssign :exec
UPDATE workflows
SET assigned_to = COALESCE(NULLIF(@member_id::text, '00000000-0000-0000-0000-000000000000')::uuid, assigned_to),
    assignment_confidence = @confidence
WHERE id = @id;
