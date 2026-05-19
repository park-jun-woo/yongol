-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING id, workflow_id, org_id, status, credits_spent, report_key, executed_at;

-- name: ExecutionLogFindByID :one
SELECT id, workflow_id, org_id, status, credits_spent, report_key, executed_at
FROM execution_logs
WHERE id = @id;

-- name: ExecutionLogListByWorkflowID :many
SELECT id, workflow_id, org_id, status, credits_spent, report_key, executed_at
FROM execution_logs
WHERE workflow_id = @workflow_id
ORDER BY executed_at DESC
LIMIT sqlc.arg(per_page) OFFSET sqlc.arg(page);

-- name: ExecutionLogCountByWorkflowID :one
SELECT COUNT(*) as total
FROM execution_logs
WHERE workflow_id = @workflow_id;
