package main

import (
	"github.com/muratovdias/url-shortner/internal/application"
	"log"
)

func main() {
	app, err := application.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
