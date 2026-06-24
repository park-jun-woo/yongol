# =============================================================================
# [name] auth smoke — happy path (bearer mode)
# Harness-supplied variables: host, admin_email, admin_password, token_nonadmin
# An admin user (admin_email/admin_password) must be seeded before running.
# =============================================================================

# --- Login (admin) -> 200, capture tokens ---
POST {{host}}/api/v1/auth/login
Content-Type: application/json
{ "email": "{{admin_email}}", "password": "{{admin_password}}" }
HTTP 200
[Captures]
token: jsonpath "$.access_token"
refresh: jsonpath "$.refresh_token"
[Asserts]
jsonpath "$.access_token" isString
jsonpath "$.refresh_token" isString

# --- GetMe (admin) -> 200 ---
GET {{host}}/api/v1/auth/me
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.email" == "{{admin_email}}"
jsonpath "$.role" == "admin"

# --- Signup (new user) -> 200, returns profile ---
POST {{host}}/api/v1/auth/signup
Content-Type: application/json
{ "email": "smoke-user@[name].com", "password": "password123" }
HTTP 200
[Asserts]
jsonpath "$.email" == "smoke-user@[name].com"
jsonpath "$.role" == "member"

# --- RefreshToken (rotate) -> 200, capture rotated tokens ---
POST {{host}}/api/v1/auth/refresh
Content-Type: application/json
{ "refresh_token": "{{refresh}}" }
HTTP 200
[Captures]
refresh: jsonpath "$.refresh_token"
[Asserts]
jsonpath "$.access_token" isString

# --- Logout (revoke refresh token) -> 200 ---
POST {{host}}/api/v1/auth/logout
Content-Type: application/json
{ "refresh_token": "{{refresh}}" }
HTTP 200
[Asserts]
jsonpath "$.success" == true
