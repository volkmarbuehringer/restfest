package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"read id",
		"GET",
		"/test/service/weburl/{id:[0-9]+}",
		getByIDHandlerWeburl,
	},
	Route{
		"read id",
		"GET",
		"/test/service/{tab}/{id:[0-9]+}",
		getByIDHandler,
	},

	Route{
		"read all",
		"GET",
		"/test/service/{tab}",
		getAllHandler,
	},

	Route{
		"Insert",
		"POST",
		"/test/service/{tab}",
		poster,
	},

	Route{
		"Update",
		"PUT",
		"/test/service/{tab}/{id:[0-9]+}",
		putter,
	},

	Route{
		"Delete",
		"DELETE",
		"/test/service/{tab}/{id:[0-9]+}",
		deleter,
	},
}
