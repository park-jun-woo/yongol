-- name: ActionListByWorkflowID :many
SELECT * FROM actions
WHERE workflow_id = @workflow_id
ORDER BY sequence_order ASC
LIMIT 1000;
