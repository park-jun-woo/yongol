VERSION := v0.7.15

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/yongol
