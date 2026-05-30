-- name: GalleryCreate :one
INSERT INTO galleries (slug, name, description, category, owner_id)
VALUES (@slug, @name, @description, @category, @owner_id)
RETURNING *;

-- name: GalleryFindBySlug :one
-- +allow-sensitive
SELECT * FROM galleries WHERE slug = @slug;

-- name: GalleryFindByID :one
-- +allow-sensitive
SELECT * FROM galleries WHERE id = @id;

-- name: GalleryList :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM galleries WHERE status = 'active' ORDER BY created_at DESC;

-- name: GalleryListByCategory :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM galleries WHERE category = @category AND status = 'active' ORDER BY created_at DESC;

-- name: GalleryCount :one
SELECT COUNT(*) FROM galleries WHERE status = 'active';

-- name: GalleryCountByCategory :one
SELECT COUNT(*) FROM galleries WHERE category = @category AND status = 'active';

-- name: GalleryUpdateStatus :exec
UPDATE galleries SET status = @status WHERE id = @id;

-- name: GalleryIncrementPostCount :exec
UPDATE galleries SET post_count = post_count + 1 WHERE id = @id;

-- name: OwnerLookupGallery :one
SELECT owner_id FROM galleries WHERE id = @id;
