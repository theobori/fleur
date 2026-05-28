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
)

type Server struct {
	options   *Options
	evaluator *evaluator.Evaluator
}

func NewServer(options *Options, evaluator *evaluator.Evaluator) *Server {
	return &Server{
		options:   options,
		evaluator: evaluator,
	}
}

func (s *Server) preProcessPath(path string) string {
	path = strings.TrimLeft(path, "/")

	if s.options.EnablePersonalGopherspaces {
		path = RenderPersonalGopherspacePath(path, "/home")
	}

	path = SafePath(path)
	if !strings.HasPrefix(path, "/") {
		path = filepath.Join(s.options.DirectoryPath, path)
	}

	return path
}

func (s *Server) sendGophermap(conn net.Conn, itemType byte, message string) error {
	return gserver.SendGophermap(conn, itemType, message, s.options.Domain, s.options.Port)
}

func (s *Server) sendGophermapError(conn net.Conn, message string) error {
	// Absolute path leak prevention
	message = strings.ReplaceAll(message, s.options.DirectoryPath, "")
	// TODO: send only if a menuentry has been request

	return s.sendGophermap(conn, gophermap.ItemTypeErrorCode, message)
}

func (s *Server) handleGophermapFile(conn net.Conn, filePath string) error {
	message, err := s.evaluator.EvalFile(filePath)
	if err != nil {
		return err
	}

	err = gserver.SendMessage(conn, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) handleFile(conn net.Conn, filePath string) error {
	filePathExtension := filepath.Ext(filePath)
	filePathName := filepath.Base(filePath)
	if filePathExtension == ".gophermap" || filePathName == "gophermap" {
		return s.handleGophermapFile(conn, filePath)
	}

	source, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = gserver.SendBytes(conn, source)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) getDirectoryFilesMessage(absoluteDirectoryPath string, relativeDirectoryPath string) (string, error) {
	entries, err := os.ReadDir(absoluteDirectoryPath)
	if err != nil {
		return "", err
	}

	lines := make([]string, len(entries)+1)

	previousDirectoryItem := gophermap.Item{
		ItemType:    gophermap.ItemTypeGopherMenu,
		Description: "..",
		Selector:    filepath.Join(relativeDirectoryPath, ".."),
		Domain:      s.options.Domain,
		Port:        s.options.Port,
	}

	lines[0] = previousDirectoryItem.String()

	for i, entry := range entries {
		entryName := entry.Name()
		absoluteEntryPath := filepath.Join(absoluteDirectoryPath, entryName)

		var item *gophermap.Item
		if entry.IsDir() {
			item, err = gophermap.NewItemFromDirectoryPath(absoluteEntryPath, s.options.Domain, s.options.Port)
		} else {
			item, err = gophermap.NewItemFromFilePath(absoluteEntryPath, s.options.Domain, s.options.Port)
		}

		if err != nil {
			return "", err
		}

		relativeEntryPath := filepath.Join(relativeDirectoryPath, entryName)

		item.Selector = relativeEntryPath
		item.Description = entryName

		lines[i+1] = item.String()
	}

	message := strings.Join(lines, "\n")

	return message, nil
}

func (s *Server) handleDirectory(conn net.Conn, absoluteDirectoryPath string, relativeDirectoryPath string) error {
	var (
		err     error
		message string
	)

	indexFilePath := filepath.Join(absoluteDirectoryPath, gophermap.DefaultIndexFileName)
	indexFileContent, err := os.ReadFile(indexFilePath)
	if err == nil {
		message, err = s.evaluator.Eval(string(indexFileContent))
	} else {
		message, err = s.getDirectoryFilesMessage(absoluteDirectoryPath, relativeDirectoryPath)
	}

	if err != nil {
		return err
	}

	err = gserver.SendMessage(conn, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) handlePath(conn net.Conn, path string) error {
	absolutePath := s.preProcessPath(path)
	relativePath := strings.TrimPrefix(absolutePath, s.options.DirectoryPath)
	relativePath = "/" + strings.TrimPrefix(relativePath, "/")

	fileInfo, err := os.Stat(absolutePath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return s.handleDirectory(conn, absolutePath, relativePath)
	}

	return s.handleFile(conn, absolutePath)
}

func (s *Server) handleMessage(conn net.Conn, message string) error {
	if message == gopher.CRLF {
		return s.handleDirectory(conn, "/", "")
	}

	return s.handlePath(conn, message)
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
		s.sendGophermapError(conn, "Gopher messages must end with <CR><LF> (\\r\\n)")
		return
	}

	message = strings.TrimSuffix(message, gopher.CRLF)

	err = s.handleMessage(conn, message)
	if err != nil {
		s.sendGophermapError(conn, err.Error())
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
