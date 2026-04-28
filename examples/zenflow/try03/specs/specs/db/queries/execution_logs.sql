-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent, executed_at)
VALUES (@workflow_id, @org_id, @status, @credits_spent, NOW())
RETURNING *;

-- name: ExecutionLogFindByID :one
SELECT * FROM execution_logs
WHERE id = @id;

-- name: ExecutionLogListByWorkflowID :many
-- +no-pagination
SELECT * FROM execution_logs
WHERE workflow_id = @workflow_id
ORDER BY executed_at DESC;
