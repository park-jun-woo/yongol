-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent, report_key)
VALUES (@workflow_id, @org_id, @status, @credits_spent, @report_key)
RETURNING *;

-- name: ExecutionLogListByWorkflow :many
-- +no-pagination
SELECT * FROM execution_logs WHERE workflow_id = @workflow_id ORDER BY executed_at DESC;

-- name: ExecutionLogFindByID :one
SELECT * FROM execution_logs WHERE id = @id;

-- name: ExecutionLogSetReportKey :exec
UPDATE execution_logs SET report_key = @report_key WHERE id = @id;

-- name: OwnerLookupExecutionLog :one
SELECT org_id FROM execution_logs WHERE id = @id;
