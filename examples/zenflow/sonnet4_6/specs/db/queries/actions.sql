-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING *;

-- name: ActionListByWorkflow :many
-- +no-pagination
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC;

-- name: ActionDeleteByWorkflow :exec
DELETE FROM actions WHERE workflow_id = @workflow_id;
