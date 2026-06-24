package authz

# @ownership gallery: galleries.owner_id
# @ownership post: posts.user_id
# @ownership comment: comments.user_id

default allow := false

# Post — any authenticated user can create
allow if {
    input.action == "CreatePost"
    input.resource == "post"
}

# Comment — any authenticated user can create
allow if {
    input.action == "CreateComment"
    input.resource == "comment"
}

# Vote — any authenticated user
allow if {
    input.action == "VotePost"
    input.resource == "post"
}

# Report — any authenticated user can create
allow if {
    input.action == "CreateReport"
    input.resource == "post"
}

# Moderation queue — admin only
allow if {
    input.action == "ListReports"
    input.resource == "report"
    input.claims.role == "admin"
}

allow if {
    input.action == "ResolveReport"
    input.resource == "report"
    input.claims.role == "admin"
}

allow if {
    input.action == "DismissReport"
    input.resource == "report"
    input.claims.role == "admin"
}
