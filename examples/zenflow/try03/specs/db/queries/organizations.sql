-- name: OrganizationCreate :one
INSERT INTO organizations (name, plan_type, credits_balance)
VALUES (@name, @plan_type, @credits_balance)
RETURNING id, name, plan_type, credits_balance, created_at;

-- name: OrganizationFindByID :one
SELECT id, name, plan_type, credits_balance, created_at
FROM organizations
WHERE id = @id;

-- name: OrganizationFindWithCredits :one
SELECT id, name, plan_type, credits_balance, created_at
FROM organizations
WHERE id = @id AND credits_balance > 0;

-- name: OrganizationDeductCredits :exec
UPDATE organizations
SET credits_balance = credits_balance - @amount
WHERE id = @id AND credits_balance >= @amount;
