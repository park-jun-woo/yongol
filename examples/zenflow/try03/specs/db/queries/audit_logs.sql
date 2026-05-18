-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;

-- name: AuditLogCreate :one
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail)
RETURNING *;

-- name: AuditLogInsert :exec
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail);

-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
AND (@filter_action::text = '' OR action = @filter_action)
AND (@filter_actor_id::text = '' OR actor_id = @filter_actor_id::uuid)
ORDER BY
  CASE WHEN @sort_by::text = 'action' AND @sort_dir::text = 'asc' THEN action END ASC,
  CASE WHEN @sort_by::text = 'action' AND @sort_dir::text = 'desc' THEN action END DESC,
  CASE WHEN @sort_by::text = 'created_at' AND @sort_dir::text = 'asc' THEN created_at END ASC,
  CASE WHEN @sort_by::text != 'action' AND @sort_dir::text != 'asc' THEN created_at END DESC
LIMIT sqlc.arg(per_page)::bigint OFFSET sqlc.arg(page_offset)::bigint;

-- name: AuditLogCountByOrgIDFiltered :one
SELECT count(*) FROM audit_logs
WHERE org_id = @org_id
AND (@filter_action::text = '' OR action = @filter_action)
AND (@filter_actor_id::text = '' OR actor_id = @filter_actor_id::uuid);

-- name: AuditLogListRecent :many
-- +no-pagination
SELECT * FROM audit_logs WHERE org_id = @org_id ORDER BY created_at DESC LIMIT 10;

-- name: OwnerLookupAuditLog :one
SELECT org_id FROM audit_logs WHERE id = @id;
