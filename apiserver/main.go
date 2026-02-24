package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"git.dilangilluly.us/dbob16/tam4/apiserver/settings"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		response := mainpath_response{Status: "healthy", WhoAmI: "tam4_server", Authenticated: false}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	})

	app.HandleFunc("GET /api/settings", func(w http.ResponseWriter, r *http.Request) {
		config := settings.ReadSettings()

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(config)
	})

	if len(os.Args) > 1 {
		if os.Args[1] == "dev" {
			log.Println("Listening on http://localhost:8000/")
			http.ListenAndServe("localhost:8000", app)
		} else {
			log.Fatal("Unrecognized argument.")
		}
	} else {
		log.Println("Listening on http://0.0.0.0:8000/")
		http.ListenAndServe(":8000", app)
	}
}

type mainpath_response struct {
	Status string `json:"healthy"`
	WhoAmI string `json:"whoami"`
	Authenticated bool `json:"authenticated"`
}
