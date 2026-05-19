package schedule

import (
	"encoding/json"
)

// @func decodeValue
// @description JSON-unquote raw session value

type DecodeValueRequest struct {
	Raw string
}

type DecodeValueResponse struct {
	Value string
}

func DecodeValue(req DecodeValueRequest) (DecodeValueResponse, error) {
	var unquoted string
	if err := json.Unmarshal([]byte(req.Raw), &unquoted); err != nil {
		return DecodeValueResponse{Value: req.Raw}, nil
	}
	return DecodeValueResponse{Value: unquoted}, nil
}
