package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/theobori/fleur/gopher"
	gclient "github.com/theobori/fleur/gopher/client"
)

func Request(host string, port int, path string, searchParameter string, useTls bool, verifyTls bool) ([]byte, error) {
	var err error

	ip := net.ParseIP(host)

	var address string
	if ip != nil && strings.Count(host, ":") >= 2 {
		address = fmt.Sprintf("[%s]:%d", host, port)
	} else {
		address = fmt.Sprintf("%s:%d", host, port)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	var conn net.Conn
	if useTls {
		tlsConf := &tls.Config{
			InsecureSkipVerify: !verifyTls,
		}

		conn, err = tls.Dial("tcp", address, tlsConf)
		if err != nil {
			return nil, err
		}
	} else {
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return nil, err
		}
	}
	defer conn.Close()

	message := path
	if searchParameter != "" {
		message += string(gopher.HT) + searchParameter
	}

	err = gclient.SendString(conn, message)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	return b, nil
}
