-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, user_id, action, resource_type, resource_id, details)
VALUES (@org_id, @user_id, @action, @resource_type, @resource_id, @details)
RETURNING *;

-- name: AuditLogListByOrg :many
-- +no-pagination
SELECT * FROM audit_logs WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: AuditLogCountByOrg :one
SELECT COUNT(*) FROM audit_logs WHERE org_id = @org_id;

-- name: AuditLogListRecent :many
-- +no-pagination
SELECT * FROM audit_logs WHERE org_id = @org_id ORDER BY created_at DESC LIMIT 10;

-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;
