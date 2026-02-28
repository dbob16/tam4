package prefixes

import (
	"net/http"
	"encoding/json"
)

func PrefixRouter() (*http.ServeMux) {
	prefixRouter := http.NewServeMux()

	prefixRouter.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefixes, err := GetAllPrefixes(api_key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(prefixes)
	})

	prefixRouter.HandleFunc("GET /{prefix_str}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefixStr := r.PathValue("prefix_str")

		rtnPrefix, err := GetOnePrefix(api_key, prefixStr)
		if err != nil {
			errText := err.Error()
			if errText == "Invalid API Key." {
				http.Error(w, errText, http.StatusUnauthorized)
			} else {
				http.Error(w, errText, http.StatusNotFound)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rtnPrefix)
	})

	prefixRouter.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		var prefix Prefix

		err := json.NewDecoder(r.Body).Decode(&prefix)
		if err != nil {
			http.Error(w, "Issue decoding prefix", http.StatusBadRequest)
			return
		}

		rtnPrefix, err := PostOnePrefix(api_key, prefix)
		if err != nil {
			errStr := err.Error()
			if errStr == "Invalid API Key." {
				http.Error(w, errStr, http.StatusUnauthorized)
				return
			} else {
				http.Error(w, errStr, http.StatusInternalServerError)
				return
			}
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rtnPrefix)
	})

	prefixRouter.HandleFunc("DELETE /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.URL.Query().Get("prefix")

		rtnStr, err := DeleteOnePrefix(api_key, prefix)
		if err != nil {
			errStr := err.Error()
			if errStr == "Invalid API Key." {
				http.Error(w, errStr, http.StatusUnauthorized)
			} else {
				http.Error(w, errStr, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rtnStr)
	})

	return prefixRouter
}
