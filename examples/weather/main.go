package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	gserver "github.com/theobori/fleur/gopher/server"
	"github.com/theobori/fleur/gophermap"
	"github.com/theobori/fleur/gophermap/evaluator"
	"github.com/theobori/fleur/server"
)

func main() {
	serverOptions, err := server.NewOptions(
		7070,
		"./examples/weather",
		"localhost",
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
	em.Set(
		"^z$",
		func(e *evaluator.Evaluator, ctx *evaluator.ExtensionContext) (string, error) {
			options := e.Options()

			item, err := gophermap.NewItem(
				gophermap.ItemTypeInlineText,
				time.Now().String()[:19],
				"/",
				options.Domain,
				options.Port,
			)
			if err != nil {
				return "", err
			}

			return item.String(), nil
		},
	)
	em.SetWithWeight(
		1,
		".*{{current-time}}.*", // You could even create you own variable system by matching ".*{{.*}}.*"
		func(e *evaluator.Evaluator, ctx *evaluator.ExtensionContext) (string, error) {
			ok, err := gophermap.IsGophermapLine(ctx.Line)
			if err != nil || !ok {
				return "", err
			}

			ans := strings.ReplaceAll(ctx.Line, "{{current-time}}", time.Now().String()[:19])

			return ans, nil
		},
	)

	router := server.NewRouter()
	router.SetWithWeight(
		0,
		".*",
		func(server *server.Server, ctx *server.RequestContext) error {
			_, err := os.Stat(ctx.Path)
			if err != nil {
				return server.SendGophermap(ctx.Conn, gophermap.ItemTypeInlineText, "Default gophermap page")
			}

			return server.HandlePath(ctx)
		},
	)

	router.SetWithWeight(
		1,
		"^/weather/get$",
		func(server *server.Server, ctx *server.RequestContext) error {
			url := "https://wttr.in/"
			if len(ctx.SearchParameter) > 0 {
				url += "/" + ctx.SearchParameter
			}

			url += "?format=4"

			res, err := http.Get(url)
			if err != nil {
				return server.SendGophermap(ctx.Conn, gophermap.ItemTypeInlineText, "Unable to create the weather api request")
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return server.SendGophermap(ctx.Conn, gophermap.ItemTypeInlineText, "Unable to read the weather api response")
			}

			bodyString := strings.TrimSuffix(string(body), "\n")

			page := gophermap.RenderMenu(
				server.NewItem(gophermap.ItemTypeInlineText, "You have requested weather in "+ctx.SearchParameter, "/"),
				server.NewItem(gophermap.ItemTypeInlineText, bodyString, "/"),
			)

			return gserver.SendString(ctx.Conn, page)
		},
	)
	router.SetWithWeight(
		1,
		"^/weather$",
		func(server *server.Server, ctx *server.RequestContext) error {
			page := gophermap.RenderMenu(
				server.NewItem(gophermap.ItemTypeInlineText, "Welcome to the gopher weather", ""),
				server.NewItem(gophermap.ItemTypeInlineText, "", ""),
				server.NewItem(gophermap.ItemTypeGopherFullTextSearch, "See current weather", "/weather/get"),
			)

			return gserver.SendString(ctx.Conn, page)
		},
	)
	evaluator := evaluator.NewEvaluator(&evaluatorOptions, em)
	server := server.NewServerWithRouter(serverOptions, evaluator, router)

	err = server.Serve()
	if err != nil {
		log.Fatalln(err)
	}
}
