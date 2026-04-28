-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, payload_template, sequence_order)
VALUES (@workflow_id, @action_type, @payload_template, @sequence_order)
RETURNING *;

-- name: ActionListByWorkflowID :many
-- +no-pagination
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC;
