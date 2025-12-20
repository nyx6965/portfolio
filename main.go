package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	view "github.com/porfolio/templates/view"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB
var (
	TURSO_DATABASE_URL string
	TURSO_AUTH_TOKEN   string
)

func main() {

	env()     // load env variables if needed
	DB = db() // make sure DB is initialized

	fname := flag.String("word", "", "Markdown file to process")
	flag.Parse()

	if *fname != "" {
		mds := []byte(ReadFile(*fname))

		html := MDToHTML(mds)
		if DB == nil {
			log.Fatal("DB is nil")
		}

		fmt.Println(string(html))
		CreatePost(DB, "something", string(html))
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := view.Home()
		component.Render(r.Context(), w)
	})

	http.ListenAndServe(":3000", r)
}

func env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		return
	}

	TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
	TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")
}

func db() *sql.DB {

	url := TURSO_DATABASE_URL + "?authToken=" + TURSO_AUTH_TOKEN

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", TURSO_DATABASE_URL, err)
		os.Exit(1)
	}

	ctx := context.Background()

	schema := `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

	`

	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	defer db.Close()
	return db
}
