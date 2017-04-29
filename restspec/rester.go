package main

import (
	"restfest/service"
)

//go:generate go run ../gen.go restspec 1

func main() {
	service.Listen(routes)
}
