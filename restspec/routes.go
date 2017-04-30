package main

import "restfest/service"

var routes = service.Routes{
	service.Route{
		"read id",
		"GET",
		"/test/service/los/{id:[0-9]+}",
		getByIDHandlerLos,
	},
	service.Route{
		"read all",
		"GET",
		"/test/service/los",
		getAllHandlerLos,
	},
	service.Route{
		"Insert",
		"POST",
		"/test/service/los",
		posterLos,
	},
	service.Route{
		"Delete",
		"DELETE",
		"/test/service/los/{id:[0-9]+}",
		deleterLos,
	},
	service.Route{
		"Update",
		"PUT",
		"/test/service/los/{id:[0-9]+}",
		putterLos,
	},
}
