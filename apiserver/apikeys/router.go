package apikeys

import (
	"encoding/json"
	"net/http"
)

func ApiKeyRouter() (*http.ServeMux) {
	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		api_pw := r.URL.Query().Get("api_pw")

		results, err := ListApikeys(api_pw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(results)
	})

	apiRouter.HandleFunc("POST /", func (w http.ResponseWriter, r *http.Request) {
		var api_req ApiReq
		err := json.NewDecoder(r.Body).Decode(&api_req)
		if err != nil {
			http.Error(w, "Invalid JSON.", http.StatusBadRequest)
			return
		}

		new_key, err := CreateAPIkey(api_req.ApiPW, api_req.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		rtn_key := ApiKey{
			ApiKey: new_key,
			Description: api_req.Description,
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rtn_key)
	})

	apiRouter.HandleFunc("DELETE /", func(w http.ResponseWriter, r *http.Request) {
		api_pw, key_to_delete := r.URL.Query().Get("api_pw"), r.URL.Query().Get("key_to_delete")

		rtn_str, err := DeleteAPIKey(api_pw, key_to_delete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rtn_str)
	})

	return apiRouter
}
