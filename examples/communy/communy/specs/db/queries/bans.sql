-- name: BanCreate :one
INSERT INTO bans (gallery_id, user_id, reason)
VALUES (@gallery_id, @user_id, @reason)
RETURNING *;

-- name: BanFindByGalleryAndUser :one
-- +allow-sensitive
SELECT * FROM bans WHERE gallery_id = @gallery_id AND user_id = @user_id;

-- name: BanListByGallery :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM bans WHERE gallery_id = @gallery_id ORDER BY created_at DESC;

-- name: BanCountByGallery :one
SELECT COUNT(*) FROM bans WHERE gallery_id = @gallery_id;

-- name: BanDelete :exec
DELETE FROM bans WHERE id = @id;

-- name: BanFindByID :one
-- +allow-sensitive
SELECT * FROM bans WHERE id = @id;
