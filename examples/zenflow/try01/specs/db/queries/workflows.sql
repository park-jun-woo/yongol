-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, @status)
RETURNING id, org_id, title, trigger_event, status, version, assignment_confidence, cron, created_at;

-- name: WorkflowFindByID :one
SELECT id, org_id, title, trigger_event, status, version, assignment_confidence, cron, created_at
FROM workflows
WHERE id = @id;

-- name: WorkflowListByOrgID :many
SELECT id, org_id, title, trigger_event, status, version, assignment_confidence, cron, created_at
FROM workflows
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT sqlc.arg(per_page) OFFSET sqlc.arg(page);

-- name: WorkflowCountByOrgID :one
SELECT COUNT(*) as total
FROM workflows
WHERE org_id = @org_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows
SET status = @status
WHERE id = @id;

-- name: WorkflowSetSchedule :exec
UPDATE workflows
SET cron = @cron
WHERE id = @id;

-- name: WorkflowDeleteSchedule :exec
UPDATE workflows
SET cron = NULL
WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;
