package client

import (
	"net"

	"github.com/theobori/fleur/gopher"
)

func SendBytes(conn net.Conn, bytes []byte) error {
	bytes = append(bytes, gopher.CR)
	bytes = append(bytes, gopher.LF)

	_, err := conn.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func SendString(conn net.Conn, s string) error {
	bytes := []byte(s)

	return SendBytes(conn, bytes)
}
