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
		err           error
		directoryPath string
		domain        string
		port          int
	)

	flag.StringVar(
		&directoryPath,
		"directory",
		"./fleur",
		"Root directory of the Gopher server",
	)
	flag.StringVar(
		&domain,
		"domain",
		"localhost",
		"Gopher domain",
	)
	flag.IntVar(
		&port,
		"port",
		gopher.DefaultPort,
		"Gopher port",
	)

	flag.Parse()

	serverOptions, err := server.NewOptions(
		port,
		directoryPath,
		domain,
		true,
	)
	if err != nil {
		log.Fatalln(err)
	}

	evaluatorOptions := evaluator.Options{
		Port:                 serverOptions.Port,
		DirectoryPath:        serverOptions.DirectoryPath,
		Domain:               serverOptions.Domain,
		EnableAutoInlineText: true,
	}

	em := evaluator.RFC1436ItemsExtensionManager()
	evaluator := evaluator.NewEvaluator(&evaluatorOptions, em)
	server := server.NewServer(serverOptions, evaluator)

	err = server.Serve()
	if err != nil {
		log.Fatalln(err)
	}
}
