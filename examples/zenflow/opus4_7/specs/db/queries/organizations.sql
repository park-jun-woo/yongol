-- name: OrganizationFindByID :one
SELECT * FROM organizations WHERE id = @id;

-- name: OrganizationDeductCredit :exec
UPDATE organizations SET credits_balance = credits_balance - @amount WHERE id = @id AND credits_balance >= @amount;
