package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/theobori/fleur/gopher"
	"github.com/theobori/fleur/gophermap"
	"github.com/theobori/fleur/gophermap/evaluator"
	"github.com/theobori/fleur/server"
)

func main() {
	var (
		err                        error
		directoryPath              string
		domain                     string
		port                       int
		enablePersonalGopherspaces bool
		enableAutoInlineText       bool
		verbose                    bool
	)

	flag.StringVar(
		&directoryPath,
		"directory",
		"./fleur",
		"It specifies an input directory path that will be the root of the Gopher server",
	)
	flag.StringVar(
		&domain,
		"domain",
		"",
		"Gopher domain",
	)
	flag.IntVar(
		&port,
		"port",
		gopher.DefaultPort,
		"Gopher port",
	)
	flag.BoolVar(
		&enablePersonalGopherspaces,
		"enable-personal-gopherspaces",
		false,
		"Enable personal Gopherspaces, it allows each user of the system to serve its own files",
	)
	flag.BoolVar(
		&enableAutoInlineText,
		"enable-auto-inline-text",
		false,
		"Relax non compliant text error and convert to gophermap inline text",
	)
	flag.BoolVar(
		&verbose,
		"verbose",
		false,
		"Enable verbose logs.",
	)

	flag.Parse()

	if port < 0 {
		log.Fatalln("The port should at least be a positive integer.")
	}

	serverOptions, err := server.NewOptions(
		port,
		directoryPath,
		domain,
		enablePersonalGopherspaces,
		verbose,
	)
	if err != nil {
		log.Fatalln(err)
	}

	evaluatorOptions := evaluator.Options{
		Port:                 serverOptions.Port,
		DirectoryPath:        serverOptions.DirectoryPath,
		Domain:               serverOptions.Domain,
		EnableAutoInlineText: enableAutoInlineText,
	}

	em := evaluator.RFC1436ItemsExtensionManager()
	em.Set(
		`^\*$`,
		func(e *evaluator.Evaluator, ctx *evaluator.ExtensionContext) (string, error) {
			fileInfo, err := os.Stat(ctx.Path)
			if err != nil {
				return "", err
			}

			path := ctx.Path
			virtualPath := ctx.VirtualPath
			if !fileInfo.IsDir() {
				path = filepath.Dir(path)
				virtualPath = filepath.Dir(virtualPath)
			}

			options := e.Options()
			text, err := gophermap.GetDirectoryFilesText(
				path,
				virtualPath,
				options.Domain,
				options.Port,
			)
			if err != nil {
				return ctx.Line, err
			}

			return text, nil
		},
	)

	router := server.NewRouter()
	if enablePersonalGopherspaces {
		router.Set(
			"^/~.*",
			func(server *server.Server, ctx *server.RequestContext) error {
				splittedVirtualPath := strings.Split(ctx.VirtualPath[2:], "/")
				username := splittedVirtualPath[0]
				rest := strings.Join(splittedVirtualPath[1:], "/")
				userPath := filepath.Join("/home", username, "public_gopher", rest)

				fileInfo, err := os.Stat(userPath)
				if err != nil {
					errMessage := strings.ReplaceAll(err.Error(), userPath, ctx.VirtualPath)
					return fmt.Errorf(errMessage)
				}

				ctx.Path = userPath

				if !fileInfo.IsDir() {
					return server.HandleFile(ctx)
				}

				return server.HandleDirectory(ctx)
			},
		)
	}
	evaluator := evaluator.NewEvaluator(&evaluatorOptions, em)
	server := server.NewServerWithRouter(serverOptions, evaluator, router)

	err = server.Serve()
	if err != nil {
		log.Fatalln(err)
	}
}
