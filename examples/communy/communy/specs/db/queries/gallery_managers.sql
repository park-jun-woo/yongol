-- name: GalleryManagerCreate :one
INSERT INTO gallery_managers (gallery_id, user_id, manager_role)
VALUES (@gallery_id, @user_id, @manager_role)
RETURNING *;

-- name: GalleryManagerFindByGalleryAndUser :one
-- +allow-sensitive
SELECT * FROM gallery_managers WHERE gallery_id = @gallery_id AND user_id = @user_id;

-- name: GalleryManagerListByGallery :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM gallery_managers WHERE gallery_id = @gallery_id ORDER BY created_at ASC;

-- name: GalleryManagerDelete :exec
DELETE FROM gallery_managers WHERE gallery_id = @gallery_id AND user_id = @user_id;

-- name: GalleryManagerCountByGallery :one
SELECT COUNT(*) FROM gallery_managers WHERE gallery_id = @gallery_id;
