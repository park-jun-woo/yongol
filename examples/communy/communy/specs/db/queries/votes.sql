-- name: VoteCreate :one
INSERT INTO votes (post_id, user_id, vote_type)
VALUES (@post_id, @user_id, @vote_type)
RETURNING *;

-- name: VoteFindByPostAndUser :one
-- +allow-sensitive
SELECT * FROM votes WHERE post_id = @post_id AND user_id = @user_id;
