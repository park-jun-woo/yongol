-- name: ActionCreate :one
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
VALUES (@workflow_id, @action_type, @config, @sequence_order)
RETURNING *;

-- name: ActionListByWorkflow :many
-- +no-pagination
SELECT * FROM actions WHERE workflow_id = @workflow_id ORDER BY sequence_order ASC;

-- name: OwnerLookupAction :one
SELECT c.org_id FROM workflows c JOIN actions l ON l.workflow_id = c.id WHERE l.id = @id;

-- name: ActionCopyToWorkflow :exec
INSERT INTO actions (workflow_id, action_type, config, sequence_order)
SELECT @target_workflow_id, a.action_type, a.config, a.sequence_order
FROM actions a WHERE a.workflow_id = @source_workflow_id;
