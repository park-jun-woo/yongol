package webhookdelivery

import "log"

// @func deliver
// @description Simulates delivering a webhook payload to a URL (no real network call)

type DeliverRequest struct {
	URL     string
	Payload string
}

type DeliverResponse struct {
}

func Deliver(req DeliverRequest) (DeliverResponse, error) {
	log.Printf("webhook deliver: url=%s payload=%s", req.URL, req.Payload)
	return DeliverResponse{}, nil
}
