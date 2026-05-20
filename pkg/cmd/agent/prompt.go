//ff:func feature=agent type=helper control=sequence
//ff:what buildPrompt — 레이어 판별 + 예시 + 프롬프트 조합

package agent

import (
	"fmt"
	"path/filepath"
	"strings"
)

// layer identifies the SSOT type of a file path.
type layer int

const (
	layerSSaC layer = iota
	layerDDL
	layerSQLcQuery
	layerOpenAPI
	layerManifest
	layerRego
	layerStateDiagram
	layerFuncSpec
	layerHurl
	layerUnknown
)

// layerPriority defines the fix order — lower index = fix first.
var layerPriority = []layer{
	layerDDL,
	layerSQLcQuery,
	layerOpenAPI,
	layerManifest,
	layerRego,
	layerStateDiagram,
	layerFuncSpec,
	layerHurl,
	layerSSaC,
}

// classifyFile returns the SSOT layer for a file path (relative to specs-dir).
func classifyFile(relPath string) layer {
	dir := filepath.Dir(relPath)
	base := filepath.Base(relPath)
	ext := filepath.Ext(base)

	switch {
	case strings.HasPrefix(dir, "service") && ext == ".ssac":
		return layerSSaC
	case dir == "db" && ext == ".sql" && !strings.HasPrefix(dir, "db/queries"):
		return layerDDL
	case strings.HasPrefix(dir, "db/queries") && ext == ".sql":
		return layerSQLcQuery
	case relPath == "api/openapi.yaml" || relPath == filepath.Join("api", "openapi.yaml"):
		return layerOpenAPI
	case base == "manifest.yaml" && (dir == "." || dir == ""):
		return layerManifest
	case strings.HasPrefix(dir, "policy") && ext == ".rego":
		return layerRego
	case strings.HasPrefix(dir, "states") && ext == ".md":
		return layerStateDiagram
	case strings.HasPrefix(dir, "func") && ext == ".go":
		return layerFuncSpec
	case strings.HasPrefix(dir, "tests") && ext == ".hurl":
		return layerHurl
	default:
		return layerUnknown
	}
}

// layerName returns a human-readable name for a layer.
func layerName(l layer) string {
	switch l {
	case layerSSaC:
		return "SSaC"
	case layerDDL:
		return "DDL"
	case layerSQLcQuery:
		return "sqlc query"
	case layerOpenAPI:
		return "OpenAPI"
	case layerManifest:
		return "manifest"
	case layerRego:
		return "Rego"
	case layerStateDiagram:
		return "stateDiagram"
	case layerFuncSpec:
		return "func spec"
	case layerHurl:
		return "Hurl"
	default:
		return "unknown"
	}
}

// buildSystemPrompt returns the system prompt with docs sections and layer example.
func buildSystemPrompt(l layer, diagMsgs []string) string {
	base := "You fix yongol SSOT files. Output ONLY the corrected file content. No explanations. No markdown.\n\n"

	docSection := searchDocs(l, diagMsgs)
	if docSection != "" {
		base += docSection + "\n\n"
	}

	base += "Example for " + layerName(l) + ":\n" + layerExample(l)
	return base
}

// buildUserPrompt assembles the user prompt from feature desc, file content, and diagnostics.
func buildUserPrompt(desc, path, filename, content string, messages []string) string {
	var b strings.Builder
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\nPath: %s\n\n", desc, path)
	}
	fmt.Fprintf(&b, "Current file (%s):\n%s\n\nValidate errors:\n", filename, content)
	for _, m := range messages {
		b.WriteString(m)
		b.WriteByte('\n')
	}
	b.WriteString("\nFix the file. Output ONLY the corrected file content.")
	return b.String()
}

// layerExample returns a minimal working example for the given SSOT layer.
func layerExample(l layer) string {
	switch l {
	case layerSSaC:
		return ssacExample
	case layerDDL:
		return ddlExample
	case layerSQLcQuery:
		return sqlcExample
	case layerOpenAPI:
		return openapiExample
	case layerManifest:
		return manifestExample
	case layerRego:
		return regoExample
	case layerStateDiagram:
		return stateDiagramExample
	case layerFuncSpec:
		return funcSpecExample
	case layerHurl:
		return hurlExample
	default:
		return ""
	}
}

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

const stateDiagramExample = `# workflow

` + "```mermaid" + `
stateDiagram-v2
    [*] --> draft
    draft --> active : ActivateWorkflow
    active --> paused : PauseWorkflow
    active --> archived : ArchiveWorkflow
    paused --> active : ActivateWorkflow
` + "```"

const funcSpecExample = `package billing

//ff:func feature=billing type=logic
//ff:what IsZeroBalance — checks if org credits are zero

type IsZeroBalanceRequest struct {
	Balance int64
}

func IsZeroBalance(req IsZeroBalanceRequest) bool {
	return req.Balance <= 0
}`

const hurlExample = `# Login
POST http://localhost:8080/auth/login
Content-Type: application/json
{"email": "admin@test.com", "password": "password123"}
HTTP 200
[Asserts]
jsonpath "$.access_token" isString`
