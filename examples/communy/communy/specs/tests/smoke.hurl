# Login
POST {{host}}/auth/login
{
  "email": "admin@example.com",
  "password": "testpassword123"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# CreateGallery
POST {{host}}/galleries
Authorization: Bearer {{token}}
{
  "slug": "programming",
  "name": "Programming Gallery",
  "description": "A gallery for programming topics",
  "category": "major"
}
HTTP 200
[Captures]
gallery_id: jsonpath "$.id"

# ListGalleries
GET {{host}}/galleries
HTTP 200

# GetGallery
GET {{host}}/galleries/{{gallery_id}}
HTTP 200

# GetGallery 404
GET {{host}}/galleries/999999
HTTP 404

# SuspendGallery
POST {{host}}/galleries/{{gallery_id}}/suspend
Authorization: Bearer {{token}}
HTTP 200

# SuspendGallery 409 (already suspended)
POST {{host}}/galleries/{{gallery_id}}/suspend
Authorization: Bearer {{token}}
HTTP 409

# UnsuspendGallery
POST {{host}}/galleries/{{gallery_id}}/unsuspend
Authorization: Bearer {{token}}
HTTP 200

# Register second user for manager
POST {{host}}/auth/register
{
  "email": "manager@example.com",
  "password": "testpassword123",
  "nickname": "manager1"
}
HTTP 200
[Captures]
manager_user_id: jsonpath "$.id"

# AddGalleryManager
POST {{host}}/galleries/{{gallery_id}}/managers
Authorization: Bearer {{token}}
{
  "user_id": {{manager_user_id}},
  "manager_role": "manager"
}
HTTP 200

# ListGalleryManagers
GET {{host}}/galleries/{{gallery_id}}/managers
HTTP 200

# RemoveGalleryManager
DELETE {{host}}/galleries/{{gallery_id}}/managers/{{manager_user_id}}
Authorization: Bearer {{token}}
HTTP 204

# CreatePost
POST {{host}}/galleries/{{gallery_id}}/posts
Authorization: Bearer {{token}}
{
  "title": "First Post",
  "body": "Hello world!",
  "is_anonymous": false
}
HTTP 200
[Captures]
post_id: jsonpath "$.id"

# ListPosts
GET {{host}}/galleries/{{gallery_id}}/posts
HTTP 200

# GetPost
GET {{host}}/posts/{{post_id}}
HTTP 200

# CreateComment
POST {{host}}/posts/{{post_id}}/comments
Authorization: Bearer {{token}}
{
  "body": "Nice post!",
  "is_anonymous": false
}
HTTP 200
[Captures]
comment_id: jsonpath "$.id"

# ListComments
GET {{host}}/posts/{{post_id}}/comments
HTTP 200

# VotePost
POST {{host}}/posts/{{post_id}}/vote
Authorization: Bearer {{token}}
{
  "vote_type": "up"
}
HTTP 200

# CreateReport
POST {{host}}/posts/{{post_id}}/report
Authorization: Bearer {{token}}
{
  "reason": "Inappropriate content"
}
HTTP 200
[Captures]
report_id: jsonpath "$.id"

# CreateCommentReport
POST {{host}}/comments/{{comment_id}}/report
Authorization: Bearer {{token}}
{
  "reason": "Spam comment"
}
HTTP 200
[Captures]
comment_report_id: jsonpath "$.id"

# ListReports
GET {{host}}/galleries/{{gallery_id}}/reports
Authorization: Bearer {{token}}
HTTP 200

# ResolveReport
POST {{host}}/reports/{{report_id}}/resolve
Authorization: Bearer {{token}}
HTTP 200

# DismissReport
POST {{host}}/reports/{{comment_report_id}}/dismiss
Authorization: Bearer {{token}}
HTTP 200

# BanUser
POST {{host}}/galleries/{{gallery_id}}/bans
Authorization: Bearer {{token}}
{
  "user_id": {{manager_user_id}},
  "reason": "Spamming"
}
HTTP 200
[Captures]
ban_id: jsonpath "$.id"

# ListBans
GET {{host}}/galleries/{{gallery_id}}/bans
Authorization: Bearer {{token}}
HTTP 200

# UnbanUser
DELETE {{host}}/galleries/{{gallery_id}}/bans/{{ban_id}}
Authorization: Bearer {{token}}
HTTP 204

# Create another post for hide/unhide test
POST {{host}}/galleries/{{gallery_id}}/posts
Authorization: Bearer {{token}}
{
  "title": "Post to Hide",
  "body": "This will be hidden",
  "is_anonymous": false
}
HTTP 200
[Captures]
hide_post_id: jsonpath "$.id"

# HidePost
POST {{host}}/posts/{{hide_post_id}}/hide
Authorization: Bearer {{token}}
HTTP 200

# UnhidePost
POST {{host}}/posts/{{hide_post_id}}/unhide
Authorization: Bearer {{token}}
HTTP 200

# DeletePost
DELETE {{host}}/posts/{{hide_post_id}}
Authorization: Bearer {{token}}
HTTP 200

# DeleteComment
DELETE {{host}}/comments/{{comment_id}}
Authorization: Bearer {{token}}
HTTP 200

# Create post for comment hide test
POST {{host}}/galleries/{{gallery_id}}/posts
Authorization: Bearer {{token}}
{
  "title": "Comment Hide Test",
  "body": "Testing comment hide",
  "is_anonymous": false
}
HTTP 200
[Captures]
hide_comment_post_id: jsonpath "$.id"

# Create comment to hide
POST {{host}}/posts/{{hide_comment_post_id}}/comments
Authorization: Bearer {{token}}
{
  "body": "Comment to hide",
  "is_anonymous": false
}
HTTP 200
[Captures]
hide_comment_id: jsonpath "$.id"

# HideComment
POST {{host}}/comments/{{hide_comment_id}}/hide
Authorization: Bearer {{token}}
HTTP 200

# ListConceptPosts
GET {{host}}/posts/concept
HTTP 200

# ListRealtimePosts
GET {{host}}/posts/realtime
HTTP 200
