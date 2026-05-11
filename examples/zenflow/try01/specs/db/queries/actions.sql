-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING *;

-- name: ActionCopyToWorkflow :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @target_workflow_id, a.action_type, a.config, a.sequence_order FROM actions a WHERE a.workflow_id = @source_workflow_id;

-- name: ActionListByWorkflowID :many
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC
LIMIT 1000;
