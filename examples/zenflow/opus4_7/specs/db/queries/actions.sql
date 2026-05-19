-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING *;

-- name: ActionListByWorkflowID :many
-- +no-pagination
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC;

-- name: ActionCopyToWorkflow :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @new_workflow_id, a.action_type, a.config, a.sequence_order FROM actions a WHERE a.workflow_id = @source_workflow_id;

-- name: ActionDeleteByWorkflowID :exec
DELETE FROM actions WHERE workflow_id = @workflow_id;

-- name: ActionBatchInsert :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @workflow_id, item->>'type', item->>'config', (item->>'sequence_order')::bigint
FROM jsonb_array_elements(sqlc.arg(items)::text::jsonb) AS item;
