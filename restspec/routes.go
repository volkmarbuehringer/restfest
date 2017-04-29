package main

import "restfest/service"

var routes = service.Routes{
	service.Route{
		"read id",
		"GET",
		"/test/service/weburl/{id:[0-9]+}",
		getByIDHandlerWeburl,
	},
	service.Route{
		"read all",
		"GET",
		"/test/service/weburl",
		getAllHandlerWeburl,
	},
}
