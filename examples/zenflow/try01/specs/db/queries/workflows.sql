-- name: WorkflowCreate :one
INSERT INTO workflows (owner_id, title, trigger_event, status)
VALUES (@owner_id, @title, @trigger_event, 'draft')
RETURNING id, owner_id, title, trigger_event, status, created_at;

-- name: WorkflowFindByID :one
SELECT id, owner_id, title, trigger_event, status, created_at
FROM workflows
WHERE id = @id;

-- name: WorkflowListByOwnerID :many
SELECT id, owner_id, title, trigger_event, status, created_at
FROM workflows
WHERE owner_id = @owner_id
ORDER BY id DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page)::int - 1) * sqlc.arg(per_page);

-- name: WorkflowCountByOwnerID :one
SELECT COUNT(*) FROM workflows WHERE owner_id = @owner_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows SET status = @status WHERE id = @id;
