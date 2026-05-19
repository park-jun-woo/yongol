-- name: TemplateFindByID :one
SELECT id, workflow_id, title, description, category, created_at
FROM templates
WHERE id = @id;

-- name: TemplateListAll :many
SELECT id, workflow_id, title, description, category, created_at
FROM templates
ORDER BY created_at DESC
LIMIT 1000;

-- name: TemplateCreate :one
INSERT INTO templates (workflow_id, title, description, category)
VALUES (@workflow_id, @title, @description, @category)
RETURNING id, workflow_id, title, description, category, created_at;
