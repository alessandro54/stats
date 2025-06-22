# Makefile
.PHONY: wire

wire:
	go run github.com/google/wire/cmd/wire@latest ./...