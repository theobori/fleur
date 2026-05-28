package server

import (
	"net"
	"strings"

	"github.com/theobori/fleur/gopher"
	"github.com/theobori/fleur/gophermap"
)

func SendBytes(conn net.Conn, bytes []byte) error {
	bytes = append(bytes, gopher.CR)
	bytes = append(bytes, gopher.LF)
	bytes = append(bytes, '.')
	bytes = append(bytes, gopher.CR)
	bytes = append(bytes, gopher.LF)

	_, err := conn.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func SendMessage(conn net.Conn, message string) error {
	lines := strings.Split(message, "\n")
	gopherMessage := strings.Join(lines, gopher.CRLF)
	bytes := []byte(gopherMessage)

	return SendBytes(conn, bytes)
}

func SendGophermap(conn net.Conn, itemType byte, message string, domain string, port int) error {
	lines := strings.Split(message, "\n")
	for i, line := range lines {
		item := gophermap.Item{
			ItemType:    itemType,
			Description: line,
			Selector:    "/",
			Domain:      domain,
			Port:        port,
		}

		lines[i] = item.String()
	}

	gophermapErrorMessage := strings.Join(lines, "\n")

	return SendMessage(conn, gophermapErrorMessage)
}
