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
	service.Route{
		"Insert",
		"POST",
		"/test/service/weburl",
		posterWeburl,
	},
	service.Route{
		"Delete",
		"DELETE",
		"/test/service/weburl/{id:[0-9]+}",
		deleterWeburl,
	},
	service.Route{
		"Update",
		"PUT",
		"/test/service/weburl/{id:[0-9]+}",
		putterWeburl,
	},
}
