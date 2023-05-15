package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thisdougb/practicalgo/config"
	"github.com/thisdougb/practicalgo/datastore"
	"github.com/thisdougb/practicalgo/webhandler"
)

var (
	version = "dev"     // populated by the compiler at build time
	commit  = "none"    // populated by the compiler at build time
	date    = "unknown" // populated by the compiler at build time
)

func main() {

	var cfg *config.Config

	log.Printf("Version %s, built at %s, commit %s \n", version, date, commit)

	ds := datastore.NewRedisDatastore(
		cfg.ValueAsStr("REDIS_HOST"),
		cfg.ValueAsStr("REDIS_PORT"),
		cfg.ValueAsStr("REDIS_USERNAME"),
		cfg.ValueAsStr("REDIS_PASSWORD"),
		cfg.ValueAsBool("REDIS_TLS"))

	for {
		log.Printf("Datastore connecting, host: %s:%s, username: %s\n",
			cfg.ValueAsStr("REDIS_HOST"),
			cfg.ValueAsStr("REDIS_PORT"),
			cfg.ValueAsStr("REDIS_USERNAME"))

		err := ds.Connect()
		if err == nil {
			log.Println("Datastore connected.")
			break
		}
		log.Println("Datastore connect failed:", err.Error())
		time.Sleep(5 * time.Second)
	}
	defer ds.Disconnect()

	// Setup web server
	handler := webhandler.New(ds)
	http.HandleFunc("/listusers/", handler.ListUsersEndpoint)

	log.Println("webserver.Start(): listening on port", cfg.ValueAsStr("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+cfg.ValueAsStr("API_PORT"), nil))

}
