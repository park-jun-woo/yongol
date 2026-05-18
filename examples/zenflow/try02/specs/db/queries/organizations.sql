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
