package main

import (
	"restfest/service"
)

//go:generate go run ../gen.go rester 0

//example for generic service with mapping to structures
func main() {
	service.Listen(routing())
}
