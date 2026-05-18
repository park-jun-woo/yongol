package webhook

import "log"

// @func deliver
// @description Simulated webhook HTTP POST delivery (purity: no real network)

type DeliverRequest struct {
	URL     string
	Payload string
}

type DeliverResponse struct{}

func Deliver(req DeliverRequest) (DeliverResponse, error) {
	log.Printf("webhook deliver (simulated): url=%s payload=%s", req.URL, req.Payload)
	return DeliverResponse{}, nil
}
