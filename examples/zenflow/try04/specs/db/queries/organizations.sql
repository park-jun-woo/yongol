-- name: OrganizationFindByID :one
SELECT * FROM organizations
WHERE id = @id;

-- name: OrganizationDeductCredit :exec
UPDATE organizations
SET credits_balance = credits_balance - @amount
WHERE id = @id;

-- name: OwnerLookupOrganization :one
SELECT org_id AS owner_id FROM users
WHERE id = @resource_id;
