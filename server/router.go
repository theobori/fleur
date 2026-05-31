package server

import (
	"fmt"
	"regexp"
	"strings"
)

const RouteMinimumMaxWeight = 10

type Router struct {
	routes    []map[string]RouteCallback
	maxWeight int
}

func NewRouterWithMaxWeight(maxWeight int) (*Router, error) {
	if maxWeight < RouteMinimumMaxWeight {
		return nil, fmt.Errorf("The minimum weight allowed is %d", RouteMinimumMaxWeight)
	}

	routes := []map[string]RouteCallback{}
	for range maxWeight + 1 {
		routes = append(routes, map[string]RouteCallback{})
	}

	return &Router{
		routes:    routes,
		maxWeight: maxWeight,
	}, nil
}

func NewRouter() *Router {
	router, _ := NewRouterWithMaxWeight(RouteMinimumMaxWeight)

	return router
}

func (r *Router) getIndexFromWeight(weight int) (int, error) {
	if weight < 0 || weight > r.maxWeight {
		return -1, fmt.Errorf("Weight must be between %d and %d", 0, r.maxWeight)
	}

	i := r.maxWeight - weight

	return i, nil
}

func (r *Router) SetWithWeight(weight int, pattern string, callback RouteCallback) error {
	i, err := r.getIndexFromWeight(weight)
	if err != nil {
		return err
	}

	r.routes[i][pattern] = callback

	return nil
}

func (r *Router) Set(pattern string, callback RouteCallback) {
	r.SetWithWeight(0, pattern, callback)
}

func (r *Router) DeleteWithWeight(weight int, pattern string) (bool, error) {
	i, err := r.getIndexFromWeight(weight)
	if err != nil {
		return false, err
	}

	_, ok := r.routes[i][pattern]
	if !ok {
		return false, nil
	}

	delete(r.routes[i], pattern)

	return true, nil
}

func (r *Router) Delete(pattern string) bool {
	ok, _ := r.DeleteWithWeight(0, pattern)

	return ok
}

func (r *Router) Route(server *Server, ctx *RequestContext) (bool, error) {
	route := strings.TrimSuffix(ctx.VirtualPath, "/")

	for _, routes := range r.routes {
		for pattern, callback := range routes {
			ok, _ := regexp.MatchString(pattern, route)
			if !ok {
				continue
			}

			err := callback(server, ctx)
			if err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return false, nil
}
