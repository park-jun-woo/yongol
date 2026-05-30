CREATE TABLE gallery_managers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    gallery_id BIGINT NOT NULL REFERENCES galleries(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    manager_role TEXT NOT NULL CHECK (manager_role IN ('sub_owner', 'manager')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- @sensitive
    UNIQUE (gallery_id, user_id)
);
