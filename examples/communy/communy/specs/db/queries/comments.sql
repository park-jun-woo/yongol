-- name: CommentCreate :one
INSERT INTO comments (post_id, user_id, body, is_anonymous)
VALUES (@post_id, @user_id, @body, @is_anonymous)
RETURNING *;

-- name: CommentFindByID :one
-- +allow-sensitive
SELECT * FROM comments WHERE id = @id;

-- name: CommentListByPost :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM comments WHERE post_id = @post_id AND status = 'published' ORDER BY created_at ASC;

-- name: CommentCountByPost :one
SELECT COUNT(*) FROM comments WHERE post_id = @post_id AND status = 'published';

-- name: CommentUpdateStatus :exec
UPDATE comments SET status = @status WHERE id = @id;

-- name: OwnerLookupComment :one
SELECT user_id FROM comments WHERE id = @id;
