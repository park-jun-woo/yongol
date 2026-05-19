-- name: ExecutionLogCreate :one
INSERT INTO execution_logs (workflow_id, org_id, status, credits_spent)
VALUES (@workflow_id, @org_id, @status, @credits_spent)
RETURNING *;

-- name: ExecutionLogListByWorkflow :many
-- +no-pagination
SELECT * FROM execution_logs WHERE workflow_id = @workflow_id ORDER BY executed_at DESC;

-- name: ExecutionLogCountByWorkflow :one
SELECT COUNT(*) FROM execution_logs WHERE workflow_id = @workflow_id;

-- name: ExecutionLogFindByID :one
SELECT * FROM execution_logs WHERE id = @id;

-- name: ExecutionLogUpdateReportKey :exec
UPDATE execution_logs SET report_file_key = @report_file_key WHERE id = @id;

-- name: ExecutionLogListByOrg :many
-- +no-pagination
SELECT * FROM execution_logs WHERE org_id = @org_id ORDER BY executed_at DESC;

-- name: ExecutionLogCountByOrg :one
SELECT COUNT(*) FROM execution_logs WHERE org_id = @org_id;
