-- name: OwnerLookupAuditLog :one
SELECT org_id FROM audit_logs WHERE id = @id;

-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;

-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail)
RETURNING *;

-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
  AND (sqlc.arg(filter_action)::text = '' OR action = sqlc.arg(filter_action)::text)
  AND (sqlc.arg(filter_actor_id)::text = '' OR actor_id = sqlc.arg(filter_actor_id)::uuid)
ORDER BY
  CASE WHEN sqlc.arg(sort_by)::text = 'action' AND sqlc.arg(sort_dir)::text = 'asc' THEN action END ASC,
  CASE WHEN sqlc.arg(sort_by)::text = 'action' AND sqlc.arg(sort_dir)::text = 'desc' THEN action END DESC,
  CASE WHEN sqlc.arg(sort_by)::text = 'created_at' AND sqlc.arg(sort_dir)::text = 'asc' THEN created_at END ASC,
  created_at DESC
LIMIT sqlc.arg(per_page)::bigint OFFSET (sqlc.arg(page)::bigint - 1) * sqlc.arg(per_page)::bigint;

-- name: AuditLogCountByOrgIDFiltered :one
SELECT COUNT(*) FROM audit_logs
WHERE org_id = @org_id
  AND (sqlc.arg(filter_action)::text = '' OR action = sqlc.arg(filter_action)::text)
  AND (sqlc.arg(filter_actor_id)::text = '' OR actor_id = sqlc.arg(filter_actor_id)::uuid);

-- name: AuditLogListRecent :many
-- +no-pagination
SELECT * FROM audit_logs
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_n)::bigint;
