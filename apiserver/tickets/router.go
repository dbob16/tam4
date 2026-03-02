package tickets

import (
	"net/http"
	"encoding/json"
	"strconv"
)

func TicketRouter() *http.ServeMux {
	ticketRouter := http.NewServeMux()

	ticketRouter.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")

		tickets, err := GetAllTickets(api_key)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tickets)
	})

	ticketRouter.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		var tickets []Ticket
		json.NewDecoder(r.Body).Decode(&tickets)

		tickets, err := PostTickets(api_key, tickets)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tickets)
	})

	ticketRouter.HandleFunc("GET /{prefix}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")

		tickets, err := GetPrefixTickets(api_key, prefix)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tickets)
	})

	ticketRouter.HandleFunc("GET /{prefix}/{ticket_id}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")
		ticket_id := r.PathValue("ticket_id")

		t_id_i, err := strconv.Atoi(ticket_id)
		if err != nil {
			http.Error(w, "Invalid integer for ticket ID.", http.StatusBadRequest)
			return
		}

		ticket, err := GetSingleTicket(api_key, prefix, t_id_i)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(ticket)
	})

	ticketRouter.HandleFunc("GET /{prefix}/{id_from}/{id_to}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")
		id_from := r.PathValue("id_from")
		id_to := r.PathValue("id_to")

		id_from_num, err := strconv.Atoi(id_from)
		if err != nil {
			http.Error(w, "Error converting from number to int, please check that it's a valid integer.", http.StatusBadRequest)
			return
		}

		id_to_num, err := strconv.Atoi(id_to)
		if err != nil {
			http.Error(w, "Error converting to number to int, please check that it's a valid integer.", http.StatusBadRequest)
			return
		}

		tickets, err := GetRangeTickets(api_key, prefix, id_from_num, id_to_num)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tickets)
	})

	return ticketRouter
}
