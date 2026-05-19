VERSION := v0.5.0

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/yongol
