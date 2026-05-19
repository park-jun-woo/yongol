package versioning

// @func nextVersion
// @description Increment the version number

type NextVersionRequest struct {
	CurrentVersion int64
}

type NextVersionResponse struct {
	Version int64
}

func NextVersion(req NextVersionRequest) (NextVersionResponse, error) {
	return NextVersionResponse{Version: req.CurrentVersion + 1}, nil
}
