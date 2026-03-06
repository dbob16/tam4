package baskets

import (
	"net/http"
	"encoding/json"
	"strconv"
)

func BasketRouter() *http.ServeMux {
	basketRouter := http.NewServeMux()

	basketRouter.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")

		rtnBaskets, err := GetAllBaskets(api_key)
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
		json.NewEncoder(w).Encode(rtnBaskets)
	})

	basketRouter.HandleFunc("GET /{prefix}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")

		rtnBaskets, err := GetPrefixBaskets(api_key, prefix)
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
		json.NewEncoder(w).Encode(rtnBaskets)
	})

	basketRouter.HandleFunc("GET /{prefix}/{basket_id}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")
		basket_id := r.PathValue("basket_id")

		i_basket_id, err := strconv.Atoi(basket_id)
		if err != nil {
			http.Error(w, "Please put valid integer in for ID", http.StatusBadRequest)
			return
		}

		Basket, err := GetSingleBasket(api_key, prefix, i_basket_id)
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
		json.NewEncoder(w).Encode(Basket)
	})

	basketRouter.HandleFunc("GET /{prefix}/{id_from}/{id_to}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")
		id_from := r.PathValue("id_from")
		id_to := r.PathValue("id_to")

		i_id_from, err := strconv.Atoi(id_from)
		if err != nil {
			http.Error(w, "Please put valid integer in for start ID", http.StatusBadRequest)
			return
		}

		i_id_to, err := strconv.Atoi(id_to)
		if err != nil {
			http.Error(w, "Please put valid integer in for end ID", http.StatusBadRequest)
			return
		}

		Basket, err := GetRangeBaskets(api_key, prefix, i_id_from, i_id_to)
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
		json.NewEncoder(w).Encode(Basket)
	})

	basketRouter.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		var baskets []Basket
		json.NewDecoder(r.Body).Decode(&baskets)

		rtnBaskets, err := PostBaskets(api_key, baskets)
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
		json.NewEncoder(w).Encode(rtnBaskets)
	})

	return basketRouter
}
