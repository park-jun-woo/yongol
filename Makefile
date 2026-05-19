VERSION := v0.3.16

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/yongol
