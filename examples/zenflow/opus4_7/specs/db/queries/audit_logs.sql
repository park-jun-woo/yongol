-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail)
RETURNING *;

-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;

-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
AND (@filter_action::text = '' OR action = @filter_action)
ORDER BY created_at DESC
LIMIT sqlc.arg(per_page)::bigint
OFFSET sqlc.arg(page_offset)::bigint;

-- name: AuditLogCountByOrgIDFiltered :one
SELECT COUNT(*) FROM audit_logs
WHERE org_id = @org_id
AND (@filter_action::text = '' OR action = @filter_action);

-- name: OwnerLookupAuditLog :one
SELECT org_id FROM audit_logs WHERE id = @id;
