# Setup: Register and Login
POST {{host}}/auth/register
{
  "email": "tester@example.com",
  "password": "testpassword123",
  "nickname": "tester"
}
HTTP 200
[Captures]
tester_id: jsonpath "$.id"

POST {{host}}/auth/login
{
  "email": "tester@example.com",
  "password": "testpassword123"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# --- 401 Login ---
POST {{host}}/auth/login
{
  "email": "tester@example.com",
  "password": "wrongpassword99"
}
HTTP 401

# --- 409 Register duplicate ---
POST {{host}}/auth/register
{
  "email": "tester@example.com",
  "password": "testpassword123",
  "nickname": "tester2"
}
HTTP 409

# --- 403 CreateGallery without auth ---
POST {{host}}/galleries
{
  "slug": "test-gal",
  "name": "Test",
  "description": "Test",
  "category": "major"
}
HTTP 403

# Setup gallery
POST {{host}}/galleries
Authorization: Bearer {{token}}
{
  "slug": "err-test",
  "name": "Error Test Gallery",
  "description": "For error testing",
  "category": "major"
}
HTTP 200
[Captures]
gal_id: jsonpath "$.id"

# --- 409 CreateGallery duplicate slug ---
POST {{host}}/galleries
Authorization: Bearer {{token}}
{
  "slug": "err-test",
  "name": "Duplicate Slug",
  "description": "Should fail",
  "category": "major"
}
HTTP 409

# --- 403 SuspendGallery (non-admin) ---
POST {{host}}/galleries/{{gal_id}}/suspend
Authorization: Bearer {{token}}
HTTP 403

# --- 404 SuspendGallery ---
POST {{host}}/galleries/999999/suspend
Authorization: Bearer {{token}}
HTTP 404

# --- 409 SuspendGallery (already suspended test via unsuspend on active) ---
POST {{host}}/galleries/{{gal_id}}/unsuspend
Authorization: Bearer {{token}}
HTTP 409

# --- 403 UnsuspendGallery without auth ---
POST {{host}}/galleries/{{gal_id}}/unsuspend
HTTP 403

# --- 404 UnsuspendGallery ---
POST {{host}}/galleries/999999/unsuspend
Authorization: Bearer {{token}}
HTTP 404

# --- 403 AddGalleryManager without auth ---
POST {{host}}/galleries/{{gal_id}}/managers
{
  "user_id": 1,
  "manager_role": "manager"
}
HTTP 403

# --- 404 AddGalleryManager gallery not found ---
POST {{host}}/galleries/999999/managers
Authorization: Bearer {{token}}
{
  "user_id": {{tester_id}},
  "manager_role": "manager"
}
HTTP 404

# --- 403 RemoveGalleryManager without auth ---
DELETE {{host}}/galleries/{{gal_id}}/managers/1
HTTP 403

# --- 404 RemoveGalleryManager gallery not found ---
DELETE {{host}}/galleries/999999/managers/1
Authorization: Bearer {{token}}
HTTP 404

# --- 404 RemoveGalleryManager manager not found ---
DELETE {{host}}/galleries/{{gal_id}}/managers/999999
Authorization: Bearer {{token}}
HTTP 404

# --- 404 ListGalleryManagers ---
GET {{host}}/galleries/999999/managers
HTTP 404

# --- 403 CreatePost without auth ---
POST {{host}}/galleries/{{gal_id}}/posts
{
  "title": "No Auth",
  "body": "Fail",
  "is_anonymous": false
}
HTTP 403

# --- 404 CreatePost gallery not found ---
POST {{host}}/galleries/999999/posts
Authorization: Bearer {{token}}
{
  "title": "No Gallery",
  "body": "Fail",
  "is_anonymous": false
}
HTTP 404

# --- 404 ListPosts ---
GET {{host}}/galleries/999999/posts
HTTP 404

# Setup post
POST {{host}}/galleries/{{gal_id}}/posts
Authorization: Bearer {{token}}
{
  "title": "Test Post",
  "body": "Test body",
  "is_anonymous": false
}
HTTP 200
[Captures]
post_id: jsonpath "$.id"

# --- 404 GetPost ---
GET {{host}}/posts/999999
HTTP 404

# --- 403 DeletePost without auth ---
DELETE {{host}}/posts/{{post_id}}
HTTP 403

# --- 404 DeletePost ---
DELETE {{host}}/posts/999999
Authorization: Bearer {{token}}
HTTP 404

# --- 403 HidePost without auth ---
POST {{host}}/posts/{{post_id}}/hide
HTTP 403

# --- 404 HidePost ---
POST {{host}}/posts/999999/hide
Authorization: Bearer {{token}}
HTTP 404

# --- 403 UnhidePost without auth ---
POST {{host}}/posts/{{post_id}}/unhide
HTTP 403

# --- 404 UnhidePost ---
POST {{host}}/posts/999999/unhide
Authorization: Bearer {{token}}
HTTP 404

# --- 403 CreateComment without auth ---
POST {{host}}/posts/{{post_id}}/comments
{
  "body": "No Auth",
  "is_anonymous": false
}
HTTP 403

# --- 404 CreateComment post not found ---
POST {{host}}/posts/999999/comments
Authorization: Bearer {{token}}
{
  "body": "No Post",
  "is_anonymous": false
}
HTTP 404

# --- 404 ListComments ---
GET {{host}}/posts/999999/comments
HTTP 404

# Setup comment
POST {{host}}/posts/{{post_id}}/comments
Authorization: Bearer {{token}}
{
  "body": "Test comment",
  "is_anonymous": false
}
HTTP 200
[Captures]
comment_id: jsonpath "$.id"

# --- 403 DeleteComment without auth ---
DELETE {{host}}/comments/{{comment_id}}
HTTP 403

# --- 404 DeleteComment ---
DELETE {{host}}/comments/999999
Authorization: Bearer {{token}}
HTTP 404

# --- 403 HideComment without auth ---
POST {{host}}/comments/{{comment_id}}/hide
HTTP 403

# --- 404 HideComment ---
POST {{host}}/comments/999999/hide
Authorization: Bearer {{token}}
HTTP 404

# --- 403 VotePost without auth ---
POST {{host}}/posts/{{post_id}}/vote
{
  "vote_type": "up"
}
HTTP 403

# --- 404 VotePost ---
POST {{host}}/posts/999999/vote
Authorization: Bearer {{token}}
{
  "vote_type": "up"
}
HTTP 404

# Vote first
POST {{host}}/posts/{{post_id}}/vote
Authorization: Bearer {{token}}
{
  "vote_type": "up"
}
HTTP 200

# --- 409 VotePost duplicate ---
POST {{host}}/posts/{{post_id}}/vote
Authorization: Bearer {{token}}
{
  "vote_type": "up"
}
HTTP 409

# --- 403 CreateReport without auth ---
POST {{host}}/posts/{{post_id}}/report
{
  "reason": "Test"
}
HTTP 403

# --- 404 CreateReport ---
POST {{host}}/posts/999999/report
Authorization: Bearer {{token}}
{
  "reason": "Test"
}
HTTP 404

# --- 403 CreateCommentReport without auth ---
POST {{host}}/comments/{{comment_id}}/report
{
  "reason": "Test"
}
HTTP 403

# --- 404 CreateCommentReport ---
POST {{host}}/comments/999999/report
Authorization: Bearer {{token}}
{
  "reason": "Test"
}
HTTP 404

# --- 403 ListReports without auth ---
GET {{host}}/galleries/{{gal_id}}/reports
HTTP 403

# --- 404 ListReports ---
GET {{host}}/galleries/999999/reports
Authorization: Bearer {{token}}
HTTP 404

# Setup report
POST {{host}}/posts/{{post_id}}/report
Authorization: Bearer {{token}}
{
  "reason": "Bad content"
}
HTTP 200
[Captures]
report_id: jsonpath "$.id"

# --- 403 ResolveReport without auth ---
POST {{host}}/reports/{{report_id}}/resolve
HTTP 403

# --- 404 ResolveReport ---
POST {{host}}/reports/999999/resolve
Authorization: Bearer {{token}}
HTTP 404

# --- 403 DismissReport without auth ---
POST {{host}}/reports/{{report_id}}/dismiss
HTTP 403

# --- 404 DismissReport ---
POST {{host}}/reports/999999/dismiss
Authorization: Bearer {{token}}
HTTP 404

# --- 403 BanUser without auth ---
POST {{host}}/galleries/{{gal_id}}/bans
{
  "user_id": 1,
  "reason": "Test"
}
HTTP 403

# --- 404 BanUser ---
POST {{host}}/galleries/999999/bans
Authorization: Bearer {{token}}
{
  "user_id": {{tester_id}},
  "reason": "Test"
}
HTTP 404

# Ban user
POST {{host}}/galleries/{{gal_id}}/bans
Authorization: Bearer {{token}}
{
  "user_id": {{tester_id}},
  "reason": "Testing"
}
HTTP 200
[Captures]
ban_id: jsonpath "$.id"

# --- 409 BanUser duplicate ---
POST {{host}}/galleries/{{gal_id}}/bans
Authorization: Bearer {{token}}
{
  "user_id": {{tester_id}},
  "reason": "Duplicate"
}
HTTP 409

# --- 403 UnbanUser without auth ---
DELETE {{host}}/galleries/{{gal_id}}/bans/{{ban_id}}
HTTP 403

# --- 404 UnbanUser ---
DELETE {{host}}/galleries/{{gal_id}}/bans/999999
Authorization: Bearer {{token}}
HTTP 404

# --- 403 ListBans without auth ---
GET {{host}}/galleries/{{gal_id}}/bans
HTTP 403

# --- 404 ListBans ---
GET {{host}}/galleries/999999/bans
Authorization: Bearer {{token}}
HTTP 404

# --- 409 State transitions ---

# Setup: Create post, delete it, try to hide (409 - already deleted)
POST {{host}}/galleries/{{gal_id}}/posts
Authorization: Bearer {{token}}
{
  "title": "State Test",
  "body": "For state testing",
  "is_anonymous": false
}
HTTP 200
[Captures]
state_post_id: jsonpath "$.id"

DELETE {{host}}/posts/{{state_post_id}}
Authorization: Bearer {{token}}
HTTP 200

# HidePost on deleted post → 409
POST {{host}}/posts/{{state_post_id}}/hide
Authorization: Bearer {{token}}
HTTP 409

# UnhidePost on deleted post → 409
POST {{host}}/posts/{{state_post_id}}/unhide
Authorization: Bearer {{token}}
HTTP 409

# DeletePost on deleted post → 409
DELETE {{host}}/posts/{{state_post_id}}
Authorization: Bearer {{token}}
HTTP 409

# Setup: Create comment, delete it, try to hide (409)
POST {{host}}/posts/{{post_id}}/comments
Authorization: Bearer {{token}}
{
  "body": "State test comment",
  "is_anonymous": false
}
HTTP 200
[Captures]
state_comment_id: jsonpath "$.id"

DELETE {{host}}/comments/{{state_comment_id}}
Authorization: Bearer {{token}}
HTTP 200

# HideComment on deleted comment → 409
POST {{host}}/comments/{{state_comment_id}}/hide
Authorization: Bearer {{token}}
HTTP 409

# DeleteComment on deleted comment → 409
DELETE {{host}}/comments/{{state_comment_id}}
Authorization: Bearer {{token}}
HTTP 409

# ResolveReport then DismissReport → 409
POST {{host}}/reports/{{report_id}}/resolve
Authorization: Bearer {{token}}
HTTP 200

POST {{host}}/reports/{{report_id}}/dismiss
Authorization: Bearer {{token}}
HTTP 409

# ResolveReport on already resolved → 409
POST {{host}}/reports/{{report_id}}/resolve
Authorization: Bearer {{token}}
HTTP 409

# SuspendGallery → 409 (need admin, regular user gets 403)
# UnsuspendGallery on active gallery → 409
POST {{host}}/galleries/{{gal_id}}/unsuspend
Authorization: Bearer {{token}}
HTTP 409

# AddGalleryManager duplicate → 409
# Setup manager
POST {{host}}/auth/register
{
  "email": "mgr2@example.com",
  "password": "testpassword123",
  "nickname": "mgr2"
}
HTTP 200
[Captures]
mgr2_id: jsonpath "$.id"

POST {{host}}/galleries/{{gal_id}}/managers
Authorization: Bearer {{token}}
{
  "user_id": {{mgr2_id}},
  "manager_role": "manager"
}
HTTP 200

POST {{host}}/galleries/{{gal_id}}/managers
Authorization: Bearer {{token}}
{
  "user_id": {{mgr2_id}},
  "manager_role": "manager"
}
HTTP 409
