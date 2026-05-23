
package agent

const ssacExample = `package service

import "github.com/park-jun-woo/ssac/pkg/auth"

// @verify-password User.email=request.email User.password_hash vs request.password -> user 401 "Invalid credentials"
// @call IssueTokenResponse token = auth.IssueToken({ID: user.ID, Email: user.Email, Role: user.Role, OrgID: user.OrgID})
// @response { access_token: token.AccessToken }
func Login() {}`

const ddlExample = `CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, -- @sensitive
    role TEXT NOT NULL CHECK (role IN ('admin', 'member')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);`

const sqlcExample = `-- name: GetUserByEmail :one
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE email = $1;`

const openapiExample = `paths:
  /auth/login:
    post:
      operationId: Login
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [email, password]
              properties:
                email: { type: string, format: email }
                password: { type: string, minLength: 8 }
      responses:
        '200':
          description: JWT token
          content:
            application/json:
              schema:
                type: object
                required: [access_token]
                properties:
                  access_token: { type: string }
        '401':
          description: Invalid credentials
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ErrorResponse'`

const manifestExample = `apiVersion: yongol/v1
kind: Project
metadata: { name: myproject }
backend:
  lang: go
  framework: gin
  module: github.com/org/myproject
  auth:
    type: jwt
    secret_env: JWT_SECRET
    user_table: users
    claims:
      ID: user_id:int64
      Email: email:string
      Role: role:string
frontend: { lang: typescript, framework: react, bundler: vite, name: myproject }`

const regoExample = `package authz

# @ownership workflow: workflows.org_id

default allow = false

allow if {
    input.action == "ListWorkflows"
    input.claims.org_id == data.owners.workflow
}`

const stateDiagramExample = "# workflow\n\n```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> active : ActivateWorkflow\n    active --> paused : PauseWorkflow\n    active --> archived : ArchiveWorkflow\n    paused --> active : ActivateWorkflow\n```"

var funcSpecExample = "package billing\n\n" +
	"/" + "/ff:func feature=billing type=logic\n" +
	"/" + "/ff:what IsZeroBalance — checks if org credits are zero\n\n" +
	"type IsZeroBalanceRequest struct {\n\tBalance int64\n}\n\n" +
	"func IsZeroBalance(req IsZeroBalanceRequest) bool {\n\treturn req.Balance <= 0\n}"

const hurlExample = `# Login
POST http://localhost:8080/auth/login
Content-Type: application/json
{"email": "admin@test.com", "password": "password123"}
HTTP 200
[Asserts]
jsonpath "$.access_token" isString`
