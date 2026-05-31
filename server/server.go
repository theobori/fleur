package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/theobori/fleur/gopher"
	gserver "github.com/theobori/fleur/gopher/server"
	"github.com/theobori/fleur/gophermap"
	"github.com/theobori/fleur/gophermap/evaluator"
	"github.com/theobori/fleur/internal/common"
)

type Server struct {
	options   *Options
	evaluator *evaluator.Evaluator
	Router    *Router
}

func NewServerWithRouter(options *Options, evaluator *evaluator.Evaluator, router *Router) *Server {
	return &Server{
		options:   options,
		evaluator: evaluator,
		Router:    router,
	}
}

func NewServer(options *Options, evaluator *evaluator.Evaluator) *Server {
	return NewServerWithRouter(options, evaluator, NewRouter())
}

func (s *Server) Options() Options {
	return *s.options
}

func (s *Server) NewItem(itemType byte, description string, selector string) *gophermap.Item {
	return &gophermap.Item{
		ItemType:    itemType,
		Description: description,
		Selector:    selector,
		Domain:      s.options.Domain,
		Port:        s.options.Port,
	}
}

func (s *Server) SendGophermap(conn net.Conn, itemType byte, message string) error {
	return gserver.SendGophermap(conn, itemType, message, s.options.Domain, s.options.Port)
}

func (s *Server) SendError(conn net.Conn, message string) error {
	// Absolute path leak prevention
	message = strings.ReplaceAll(message, s.options.DirectoryPath, "")
	return gserver.SendString(conn, fmt.Sprintf("Error: %s", message))
}

func (s *Server) handleGophermapFilePath(ctx *RequestContext) error {
	res, err := s.evaluator.EvalFile(ctx.Path, ctx.VirtualPath)
	if err != nil {
		return err
	}

	err = gserver.SendString(ctx.Conn, res)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) HandleFile(ctx *RequestContext) error {
	filePathExtension := filepath.Ext(ctx.Path)
	filePathName := filepath.Base(ctx.Path)

	if filePathExtension == ".gophermap" || filePathName == "gophermap" {
		return s.handleGophermapFilePath(ctx)
	}

	source, err := os.ReadFile(ctx.Path)
	if err != nil {
		return err
	}

	err = gserver.SendBytes(ctx.Conn, source)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) HandleDirectory(ctx *RequestContext) error {
	var (
		err     error
		message string
	)

	indexFilePath := filepath.Join(ctx.Path, gophermap.DefaultIndexFileName)
	indexFileContent, err := os.ReadFile(indexFilePath)
	if err == nil {
		message, err = s.evaluator.EvalWithPathContext(
			string(indexFileContent),
			ctx.Path,
			ctx.VirtualPath,
		)
	} else {
		message, err = gophermap.GetDirectoryFilesText(
			ctx.Path,
			ctx.VirtualPath,
			s.options.Domain,
			s.options.Port,
		)
	}

	if err != nil {
		return err
	}

	err = gserver.SendString(ctx.Conn, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) HandlePath(ctx *RequestContext) error {
	fileInfo, err := os.Stat(ctx.Path)
	if err != nil {
		if s.options.Verbose {
			log.Printf("The path '%s' cannot be stat\n", ctx.VirtualPath)
		}
		return err
	}

	if fileInfo.IsDir() {
		if s.options.Verbose {
			log.Printf("The path '%s' is served as a directory\n", ctx.VirtualPath)
		}
		return s.HandleDirectory(ctx)
	}

	if s.options.Verbose {
		log.Printf("The path '%s' is served as a file\n", ctx.VirtualPath)
	}

	return s.HandleFile(ctx)
}

func (s *Server) HandleRequest(ctx *RequestContext) error {
	if s.options.Verbose {
		log.Printf("The path '%s' has been requested\n", ctx.VirtualPath)
	}

	ok, err := s.Router.Route(s, ctx)
	if err != nil {
		return err
	}

	if ok {
		if s.options.Verbose {
			log.Printf("The path '%s' matched a route\n", ctx.VirtualPath)
		}
		return nil
	}

	// If no route has matched it will try to serve a file/directory
	return s.HandlePath(ctx)
}

func (s *Server) getRequestContext(conn net.Conn, message string) *RequestContext {
	var (
		virtualPath     string
		searchParameter string
	)

	questionMarkIndex := strings.Index(message, string(gopher.HT))
	if questionMarkIndex != -1 {
		searchParameter = message[questionMarkIndex+1:]
		virtualPath = message[:questionMarkIndex]
	} else {
		searchParameter = ""
		virtualPath = message
	}

	virtualPath = common.SafePath(virtualPath)
	virtualPath = "/" + strings.TrimPrefix(virtualPath, "/")

	ctx := RequestContext{
		Conn:            conn,
		VirtualPath:     virtualPath,
		Path:            filepath.Join(s.options.DirectoryPath, virtualPath),
		SearchParameter: searchParameter,
	}

	return &ctx
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	if !strings.HasSuffix(message, gopher.CRLF) {
		s.SendError(conn, "Gopher messages must end with <CR><LF> (\\r\\n)")
		return
	}
	message = strings.TrimSuffix(message, gopher.CRLF)

	ctx := s.getRequestContext(conn, message)

	err = s.HandleRequest(ctx)
	if err != nil {
		s.SendError(conn, err.Error())
		return
	}
}

func (s *Server) Serve() error {
	// Expose a port to every network interface availables for the moment
	address := fmt.Sprintf(":%d", s.options.Port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleConnection(conn)
	}
}
