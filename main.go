package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TonyGLL/url-shortener/api"
	db "github.com/TonyGLL/url-shortener/db/sql"
	"github.com/TonyGLL/url-shortener/util"
	_ "github.com/lib/pq"
)

// @contact.name Tony Gonzalez
// @contact.utl https://github.com/TonyGLL
// @contact.email tonygllambia@gmail.com
func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		log.Fatal("CONFIG_FILE environment variable not set")
	}

	config, err := util.LoadConfig(".", configFile)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, config)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
