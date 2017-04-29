package main

import (
	"restfest/service"
)

//go:generate go run ../gen.go restspec 1

//example for service with fixed routes and without mapping
func main() {
	service.Listen(routes)
}
