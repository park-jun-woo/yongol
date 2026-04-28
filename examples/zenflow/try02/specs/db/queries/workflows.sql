-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, 'draft')
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowListByOrgIDPaged :many
SELECT * FROM workflows
WHERE org_id = @org_id
ORDER BY
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'asc'  THEN created_at END ASC,
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'desc' THEN created_at END DESC,
  id DESC
LIMIT sqlc.arg(per_page)::bigint OFFSET (sqlc.arg(page)::bigint - 1) * sqlc.arg(per_page)::bigint;

-- name: WorkflowCountByOrgID :one
SELECT COUNT(*) FROM workflows WHERE org_id = @org_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows WHERE id = @id;
