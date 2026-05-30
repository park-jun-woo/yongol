-- name: ReportCreate :one
INSERT INTO reports (gallery_id, target_type, target_id, reporter_id, reason)
VALUES (@gallery_id, @target_type, @target_id, @reporter_id, @reason)
RETURNING *;

-- name: ReportFindByID :one
-- +allow-sensitive
SELECT * FROM reports WHERE id = @id;

-- name: ReportListByGallery :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM reports WHERE gallery_id = @gallery_id ORDER BY created_at DESC;

-- name: ReportCountByGallery :one
SELECT COUNT(*) FROM reports WHERE gallery_id = @gallery_id;

-- name: ReportUpdateStatus :exec
UPDATE reports SET status = @status WHERE id = @id;
