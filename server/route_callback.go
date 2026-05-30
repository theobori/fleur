package server

type RouteCallback func(server *Server, ctx *RequestContext) error
