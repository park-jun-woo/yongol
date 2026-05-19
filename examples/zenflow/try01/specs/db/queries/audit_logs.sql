-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail)
RETURNING *;

-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;

-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action)
  AND (@filter_actor_id::text = '' OR actor_id::text = @filter_actor_id)
ORDER BY
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'asc'  THEN created_at END ASC,
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'desc' THEN created_at END DESC,
  CASE WHEN @sort_by = 'action' AND @sort_dir = 'asc'  THEN action END ASC,
  CASE WHEN @sort_by = 'action' AND @sort_dir = 'desc' THEN action END DESC
LIMIT sqlc.arg(per_page)::bigint OFFSET (sqlc.arg(page)::bigint - 1) * sqlc.arg(per_page)::bigint;

-- name: AuditLogCountByOrgIDFiltered :one
SELECT COUNT(*) FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action)
  AND (@filter_actor_id::text = '' OR actor_id::text = @filter_actor_id);

-- name: AuditLogListRecent :many
SELECT * FROM audit_logs WHERE org_id = @org_id ORDER BY created_at DESC LIMIT 10;

-- name: OwnerLookupAuditLog :one
SELECT org_id FROM audit_logs WHERE id = @id;
