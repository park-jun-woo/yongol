VERSION := v0.7.32

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/yongol
