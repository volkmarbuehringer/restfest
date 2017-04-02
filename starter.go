package restfest

import "restfest/service"

//go:generate go run gen.go

func Starter() {
	service.Listen()
}
