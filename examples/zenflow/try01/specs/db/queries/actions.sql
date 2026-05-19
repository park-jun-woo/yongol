-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING id, workflow_id, action_type, config, sequence_order, created_at;

-- name: ActionListByWorkflowID :many
SELECT id, workflow_id, action_type, config, sequence_order, created_at
FROM actions
WHERE workflow_id = @workflow_id
ORDER BY sequence_order ASC
LIMIT 1000;

-- name: ActionFindByID :one
SELECT id, workflow_id, action_type, config, sequence_order, created_at
FROM actions
WHERE id = @id;

-- name: ActionDeleteByWorkflowID :exec
DELETE FROM actions
WHERE workflow_id = @workflow_id;

-- name: ActionBatchInsert :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @workflow_id, item->>'action_type', item->>'config', (item->>'sequence_order')::bigint
FROM jsonb_array_elements(@items::jsonb) AS item;
