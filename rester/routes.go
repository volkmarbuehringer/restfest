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
		"read id",
		"GET",
		"/test/service/{tab}/{id:[0-9]+}",
		getByIDHandler,
	},

	service.Route{
		"read all",
		"GET",
		"/test/service/{tab}",
		getAllHandler,
	},

	service.Route{
		"Insert",
		"POST",
		"/test/service/{tab}",
		poster,
	},

	service.Route{
		"Update",
		"PUT",
		"/test/service/{tab}/{id:[0-9]+}",
		putter,
	},

	service.Route{
		"Delete",
		"DELETE",
		"/test/service/{tab}/{id:[0-9]+}",
		deleter,
	},
}
