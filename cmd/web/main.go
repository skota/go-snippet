package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"../../data/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//parameters

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "Mysql DSN")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")

	flag.Parse()

	db := connect(*dsn)

	defer db.Close()

	//inititaize App struct
	app := &App{
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on port %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		//cannot connect to database
		log.Fatal(err)
	}

	return db
}
