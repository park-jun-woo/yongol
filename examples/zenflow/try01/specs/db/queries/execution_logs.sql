-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING *;

-- name: ExecutionLogFindByID :one
SELECT * FROM execution_logs WHERE id = @id;
