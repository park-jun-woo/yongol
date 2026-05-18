package schedule

import "encoding/json"

// @func decodeValue
// @error 500
// @description JSON-decodes a session raw value string into a plain string.

type DecodeValueRequest struct {
	Raw string
}

type DecodeValueResponse struct {
	Value string
}

func DecodeValue(req DecodeValueRequest) (DecodeValueResponse, error) {
	var s string
	if err := json.Unmarshal([]byte(req.Raw), &s); err != nil {
		// If not valid JSON string, return raw as-is
		return DecodeValueResponse{Value: req.Raw}, nil
	}
	return DecodeValueResponse{Value: s}, nil
}
