package main

import (
	"fmt"
	"log"

	"github.com/theobori/fleur/client"
)

func main() {
	response, err := client.Request("bitreich.org", 70, "/billion-gophers", "", false, false)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(response))
}
