-- name: AuditLogFindByID :one
SELECT id, org_id, user_id, action, created_at
FROM audit_logs
WHERE id = @id;

-- name: AuditLogListByOrgID :many
SELECT id, org_id, user_id, action, created_at
FROM audit_logs
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT 1000;

-- name: AuditLogListRecent :many
SELECT id, org_id, user_id, action, created_at
FROM audit_logs
ORDER BY created_at DESC
LIMIT 20;

-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, user_id, action)
VALUES (@org_id, @user_id, @action)
RETURNING id, org_id, user_id, action, created_at;
