package migration

const richSQL = `
-- a leading line comment
CREATE TABLE "Order" (
    id BIGINT PRIMARY KEY,
    status VARCHAR(20) NOT NULL DEFAULT 'pending; not done', -- trailing comment
    note TEXT /* block ; comment */ DEFAULT 'a''b'
);
CREATE INDEX idx ON "Order" (status);
`
