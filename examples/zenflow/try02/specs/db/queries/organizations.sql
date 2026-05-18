-- name: OrganizationFindByID :one
SELECT * FROM organizations WHERE id = @id;

-- name: OrganizationCreate :one
INSERT INTO organizations (name, plan_type, credits_balance)
VALUES (@name, @plan_type, @credits_balance)
RETURNING *;

-- name: OrganizationUpdateCredits :one
UPDATE organizations
SET credits_balance = credits_balance - @amount
WHERE id = @id AND credits_balance >= @amount
RETURNING *;

-- name: OrganizationUpdateGeocode :exec
UPDATE organizations
SET latitude = @latitude, longitude = @longitude, address_verified = true
WHERE id = @id;

-- name: OwnerLookupOrganization :one
SELECT id FROM organizations WHERE id = @id;
