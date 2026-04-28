-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING *;

-- name: ExecutionLogListByOrgIDPaged :many
SELECT * FROM execution_logs
WHERE org_id = @org_id
ORDER BY executed_at DESC, id DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page)::int - 1) * sqlc.arg(per_page);

-- name: ExecutionLogCountByOrgID :one
SELECT COUNT(*) FROM execution_logs WHERE org_id = @org_id;
