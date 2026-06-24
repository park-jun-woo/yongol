package authz

# Authentication (valid bearer token) is handled by the bearerAuth middleware.
# Authorization below is role-based: the protected profile endpoint is granted
# to admins only. Non-admin callers receive 403 (exercised by smoke/scenario
# hurl with a non-admin token). Public endpoints (Signup/Login/RefreshToken/
# Logout) carry no @auth and are not listed here.
default allow := false

allow if {
    input.action == "GetMe"
    input.resource == "user"
    input.claims.role == "admin"
}
