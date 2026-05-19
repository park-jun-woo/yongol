-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, 'draft')
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowListByOrg :many
-- +no-pagination
SELECT * FROM workflows WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WorkflowCountByOrg :one
SELECT COUNT(*) FROM workflows WHERE org_id = @org_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;

-- name: WorkflowCreateVersion :one
INSERT INTO workflows (org_id, root_workflow_id, title, trigger_event, status)
VALUES (@org_id, @root_workflow_id, @title, @trigger_event, 'draft')
RETURNING *;

-- name: WorkflowListVersions :many
-- +no-pagination
SELECT * FROM workflows WHERE root_workflow_id = @root_workflow_id ORDER BY created_at DESC;

-- name: WorkflowSetCronExpr :exec
UPDATE workflows SET cron_expr = @cron_expr WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;
