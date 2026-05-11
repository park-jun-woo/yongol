-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING *;

-- name: ExecutionLogListByWorkflowID :many
SELECT * FROM execution_logs WHERE workflow_id = @workflow_id ORDER BY executed_at DESC
LIMIT 1000;
