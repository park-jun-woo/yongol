-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING id, workflow_id, org_id, status, credits_spent, executed_at;
