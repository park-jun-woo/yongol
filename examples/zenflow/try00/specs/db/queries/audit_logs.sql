-- name: AuditLogCreate :exec
INSERT INTO audit_logs (org_id, actor_id, action, resource_type, resource_id, detail)
VALUES (@org_id, @actor_id, @action, @resource_type, @resource_id, @detail);

-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = @id;

-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action)
ORDER BY
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'asc'  THEN created_at END ASC,
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'desc' THEN created_at END DESC,
  CASE WHEN @sort_by = 'action' AND @sort_dir = 'asc'  THEN action END ASC,
  CASE WHEN @sort_by = 'action' AND @sort_dir = 'desc' THEN action END DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page) - 1) * sqlc.arg(per_page);

-- name: AuditLogCountByOrgIDFiltered :one
SELECT COUNT(*) FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action);
