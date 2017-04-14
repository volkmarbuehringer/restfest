package main

import (
	"restfest/service"
)

//go:generate go run ../gen.go rester 0

func main() {
	service.Listen()
}
