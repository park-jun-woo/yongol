-- name: OwnerLookupTemplate :one
SELECT org_id FROM templates WHERE id = @id;

-- name: TemplateCreate :one
INSERT INTO templates (source_workflow_id, org_id, title, description, category)
VALUES (@source_workflow_id, @org_id, @title, @description, @category)
RETURNING *;

-- name: TemplateFindByID :one
SELECT * FROM templates WHERE id = @id;

-- name: TemplateFindBySourceWorkflow :one
SELECT * FROM templates WHERE source_workflow_id = @source_workflow_id;

-- name: TemplateListCursor :many
SELECT * FROM templates
WHERE (sqlc.arg(cursor)::text = '' OR id < sqlc.arg(cursor)::uuid)
  AND (sqlc.arg(category)::text = '' OR category = sqlc.arg(category)::text)
ORDER BY id DESC
LIMIT sqlc.arg(per_page)::bigint;

-- name: TemplateIncrementCloneCount :exec
UPDATE templates SET clone_count = clone_count + 1 WHERE id = @id;
