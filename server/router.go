package server

import (
	"regexp"
	"strings"
)

type Router struct {
	routes map[string]RouteCallback
}

func NewRouter() *Router {
	return &Router{
		routes: map[string]RouteCallback{},
	}
}

func (t *Router) Set(routePattern string, callback RouteCallback) {
	t.routes[routePattern] = callback
}

func (t *Router) Delete(routePattern string) bool {
	_, ok := t.routes[routePattern]
	if !ok {
		return false
	}

	delete(t.routes, routePattern)

	return true
}

func (t *Router) Route(server *Server, ctx *RequestContext) (bool, error) {
	route := strings.TrimSuffix(ctx.VirtualPath, "/")

	for routePattern, callback := range t.routes {
		ok, _ := regexp.MatchString(routePattern, route)
		if !ok {
			continue
		}

		err := callback(server, ctx)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
