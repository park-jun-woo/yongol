-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, @status)
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows WHERE id = @id;

-- name: WorkflowListByOrg :many
SELECT * FROM workflows
WHERE org_id = @org_id
ORDER BY
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'asc'  THEN created_at END ASC,
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'desc' THEN created_at END DESC,
  CASE WHEN @sort_by = 'title'      AND @sort_dir = 'asc'  THEN title END ASC,
  CASE WHEN @sort_by = 'title'      AND @sort_dir = 'desc' THEN title END DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page)::int - 1) * sqlc.arg(per_page);

-- name: WorkflowCountByOrg :one
SELECT COUNT(*) FROM workflows WHERE org_id = @org_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;
