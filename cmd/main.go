package main

import (
	"github.com/muratovdias/url-shortner/src/application"
	"log"
)

// @title						Url-Shortener API
// @version					1.0
// @host						localhost:8080
// @BasePath	/
func main() {
	app, err := application.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
