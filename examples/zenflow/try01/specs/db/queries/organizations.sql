-- name: OrganizationCreate :one
INSERT INTO organizations (name, plan_type, credits_balance)
VALUES (@name, @plan_type, @credits_balance)
RETURNING id, name, plan_type, credits_balance, address, address_verified, created_at;

-- name: OrganizationFindByID :one
SELECT id, name, plan_type, credits_balance, address, address_verified, created_at
FROM organizations
WHERE id = @id;

-- name: OrganizationVerifyAddress :exec
UPDATE organizations SET address = @address, address_verified = true
WHERE id = @id;
