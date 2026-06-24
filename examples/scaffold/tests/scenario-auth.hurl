# =============================================================================
# [name] auth error-path scenarios (4xx coverage)
# Harness-supplied variables: host, token_nonadmin (a valid non-admin JWT)
# =============================================================================

# --- Signup: missing password -> 400 ---
POST {{host}}/api/v1/auth/signup
Content-Type: application/json
{ "email": "bad-signup@[name].com" }
HTTP 400

# --- Signup: first registration -> 200 ---
POST {{host}}/api/v1/auth/signup
Content-Type: application/json
{ "email": "dup@[name].com", "password": "password123" }
HTTP 200

# --- Signup: duplicate email -> 409 ---
POST {{host}}/api/v1/auth/signup
Content-Type: application/json
{ "email": "dup@[name].com", "password": "password123" }
HTTP 409

# --- Login: wrong password -> 401 ---
POST {{host}}/api/v1/auth/login
Content-Type: application/json
{ "email": "dup@[name].com", "password": "wrong-password" }
HTTP 401

# --- Login: malformed body -> 400 ---
POST {{host}}/api/v1/auth/login
Content-Type: application/json
{ "email": "dup@[name].com" }
HTTP 400

# --- GetMe: no token -> 401 ---
GET {{host}}/api/v1/auth/me
HTTP 401

# --- GetMe: non-admin token -> 403 ---
GET {{host}}/api/v1/auth/me
Authorization: Bearer {{token_nonadmin}}
HTTP 403
[Asserts]
jsonpath "$.code" exists

# --- RefreshToken: invalid token -> 401 ---
POST {{host}}/api/v1/auth/refresh
Content-Type: application/json
{ "refresh_token": "not-a-real-token" }
HTTP 401

# --- RefreshToken: malformed body -> 400 ---
POST {{host}}/api/v1/auth/refresh
Content-Type: application/json
{ }
HTTP 400

# --- Logout: malformed body -> 400 ---
POST {{host}}/api/v1/auth/logout
Content-Type: application/json
{ }
HTTP 400
