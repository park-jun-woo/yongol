-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING *;

-- name: ActionCopyToWorkflow :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @new_workflow_id, src.action_type, src.config, src.sequence_order
FROM actions src WHERE src.workflow_id = @source_workflow_id;

-- name: ActionListByWorkflow :many
-- +no-pagination
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC;

-- name: ActionDeleteByWorkflow :exec
DELETE FROM actions WHERE workflow_id = @workflow_id;

-- name: ActionBatchInsert :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @workflow_id, item->>'type', item->>'config', (item->>'sequence_order')::bigint
FROM jsonb_array_elements(@items::jsonb) AS item;
