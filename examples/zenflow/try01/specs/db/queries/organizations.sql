-- name: OrganizationFindByID :one
SELECT * FROM organizations WHERE id = @id;

-- name: OrganizationDeductCredit :exec
UPDATE organizations SET credits_balance = credits_balance - @amount WHERE id = @id;

-- name: OrganizationUpdateGeocode :exec
UPDATE organizations SET latitude = @latitude, longitude = @longitude, address_verified = true WHERE id = @id;

-- name: OwnerLookupOrganization :one
SELECT id FROM organizations WHERE id = @id;
