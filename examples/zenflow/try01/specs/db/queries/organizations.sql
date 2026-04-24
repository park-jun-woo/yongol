-- name: OrganizationCreate :one
INSERT INTO organizations (name, plan_type, credits_balance)
VALUES (@name, @plan_type, @credits_balance)
RETURNING *;

-- name: OrganizationFindByID :one
SELECT * FROM organizations WHERE id = @id;

-- name: OrganizationUpdateCredits :exec
UPDATE organizations SET credits_balance = @credits_balance WHERE id = @id;
