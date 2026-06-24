-- name: PostCreate :one
INSERT INTO posts (gallery_id, user_id, title, body, is_anonymous)
VALUES (@gallery_id, @user_id, @title, @body, @is_anonymous)
RETURNING *;

-- name: PostFindByID :one
-- +allow-sensitive
SELECT * FROM posts WHERE id = @id;

-- name: PostListByGallery :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM posts WHERE gallery_id = @gallery_id AND status = 'published' ORDER BY created_at DESC;

-- name: PostCountByGallery :one
SELECT COUNT(*) FROM posts WHERE gallery_id = @gallery_id AND status = 'published';

-- name: PostUpdateStatus :exec
UPDATE posts SET status = @status WHERE id = @id;

-- name: PostUpdateUpvotes :exec
UPDATE posts SET upvotes = @upvotes WHERE id = @id;

-- name: PostUpdateDownvotes :exec
UPDATE posts SET downvotes = @downvotes WHERE id = @id;

-- name: PostUpdateIsConcept :exec
UPDATE posts SET is_concept = @is_concept WHERE id = @id;

-- name: PostIncrementViewCount :exec
UPDATE posts SET view_count = view_count + 1 WHERE id = @id;

-- name: PostIncrementCommentCount :exec
UPDATE posts SET comment_count = comment_count + 1 WHERE id = @id;

-- name: PostListConcept :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM posts WHERE is_concept = TRUE AND status = 'published' ORDER BY created_at DESC;

-- name: PostCountConcept :one
SELECT COUNT(*) FROM posts WHERE is_concept = TRUE AND status = 'published';

-- name: PostListRealtime :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM posts WHERE status = 'published' ORDER BY created_at DESC;

-- name: PostCountRealtime :one
SELECT COUNT(*) FROM posts WHERE status = 'published';

-- name: OwnerLookupPost :one
SELECT user_id FROM posts WHERE id = @id;
