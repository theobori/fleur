package server

import "net"

type RequestContext struct {
	Conn        net.Conn
	Path        string
	VirtualPath string
	// Parameters []string ??
}
