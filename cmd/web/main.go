package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//parameters
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")

	flag.Parse()

	//inititaize App struct
	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on port %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)

}
