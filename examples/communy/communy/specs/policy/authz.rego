package authz

# @ownership gallery: galleries.owner_id
# @ownership post: posts.user_id
# @ownership comment: comments.user_id

default allow := false

# Gallery — any authenticated user can create
allow if {
    input.action == "CreateGallery"
    input.resource == "gallery"
}

# Gallery — admin only
allow if {
    input.action == "SuspendGallery"
    input.resource == "gallery"
    input.claims.role == "admin"
}

allow if {
    input.action == "UnsuspendGallery"
    input.resource == "gallery"
    input.claims.role == "admin"
}

# Gallery management — gallery owner only
allow if {
    input.action == "AddGalleryManager"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}

allow if {
    input.action == "RemoveGalleryManager"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}

# Post — any authenticated user can create
allow if {
    input.action == "CreatePost"
    input.resource == "post"
}

# Post — owner can delete
allow if {
    input.action == "DeletePost"
    input.resource == "post"
    data.owners.post[input.resource_id] == input.claims.id
}

# Post moderation — admin only
allow if {
    input.action == "HidePost"
    input.resource == "post"
    input.claims.role == "admin"
}

allow if {
    input.action == "UnhidePost"
    input.resource == "post"
    input.claims.role == "admin"
}

# Comment — any authenticated user can create
allow if {
    input.action == "CreateComment"
    input.resource == "comment"
}

# Comment — owner can delete
allow if {
    input.action == "DeleteComment"
    input.resource == "comment"
    data.owners.comment[input.resource_id] == input.claims.id
}

# Comment moderation — admin only
allow if {
    input.action == "HideComment"
    input.resource == "comment"
    input.claims.role == "admin"
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

allow if {
    input.action == "CreateCommentReport"
    input.resource == "comment"
}

# Report management — gallery owner
allow if {
    input.action == "ListReports"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}

allow if {
    input.action == "ResolveReport"
    input.resource == "report"
}

allow if {
    input.action == "DismissReport"
    input.resource == "report"
}

# Ban — gallery owner
allow if {
    input.action == "BanUser"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}

allow if {
    input.action == "UnbanUser"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}

allow if {
    input.action == "ListBans"
    input.resource == "gallery"
    data.owners.gallery[input.resource_id] == input.claims.id
}
