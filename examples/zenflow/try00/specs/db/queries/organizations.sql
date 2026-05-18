-- name: OrganizationFindByID :one
SELECT * FROM organizations WHERE id = @id;

-- name: OrganizationDeductCredits :exec
UPDATE organizations SET credits_balance = credits_balance - @amount WHERE id = @id;
