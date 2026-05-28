package main

import (
	"flag"
	"log"

	"github.com/theobori/fleur/gopher"
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

	flag.Parse()

	if port < 0 {
		log.Fatalln("The port should at least be a positive integer.")
	}

	serverOptions, err := server.NewOptions(port, directoryPath, domain, enablePersonalGopherspaces)
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

	evaluator := evaluator.NewEvaluator(&evaluatorOptions, em)
	server := server.NewServer(serverOptions, evaluator)

	err = server.Serve()
	if err != nil {
		log.Fatalln(err)
	}
}
