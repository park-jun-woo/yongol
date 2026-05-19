-- name: TemplateCreate :one
INSERT INTO templates (org_id, source_workflow_id, title, description, category, is_public)
VALUES (@org_id, @source_workflow_id, @title, @description, @category, @is_public)
RETURNING *;

-- name: TemplateListWithCursor :many
-- +no-pagination
SELECT * FROM templates WHERE is_public = true ORDER BY id ASC;

-- name: TemplateFindByID :one
SELECT * FROM templates WHERE id = @id;

-- name: OwnerLookupTemplate :one
SELECT org_id FROM templates WHERE id = @id;
