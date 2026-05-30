CREATE TABLE posts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    gallery_id BIGINT NOT NULL REFERENCES galleries(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'published' CHECK (status IN ('published', 'hidden', 'deleted')),
    is_anonymous BOOLEAN NOT NULL DEFAULT FALSE,
    is_concept BOOLEAN NOT NULL DEFAULT FALSE,
    upvotes BIGINT NOT NULL DEFAULT 0,
    downvotes BIGINT NOT NULL DEFAULT 0,
    comment_count BIGINT NOT NULL DEFAULT 0,
    view_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW() -- @sensitive
);
