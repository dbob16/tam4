package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
	"git.dilangilluly.us/dbob16/tam4/apiserver/apikeys"
	"git.dilangilluly.us/dbob16/tam4/apiserver/prefixes"
	"git.dilangilluly.us/dbob16/tam4/apiserver/tickets"
)

func main() {
	log.Println(localdb.InitDB())
	app := http.NewServeMux()

	app.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		key_auth := apikeys.CheckAPIKey(api_key)

		response := mainpath_response{Status: "healthy", WhoAmI: "tam4_server", Authenticated: key_auth}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	})

	app.Handle("/api/api_keys/", http.StripPrefix("/api/api_keys", apikeys.ApiKeyRouter()))
	app.Handle("/api/prefixes/", http.StripPrefix("/api/prefixes", prefixes.PrefixRouter()))
	app.Handle("/api/tickets/", http.StripPrefix("/api/tickets", tickets.TicketRouter()))

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
