package drawing

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func DrawingRouter() *http.ServeMux {
	drawingRouter := http.NewServeMux()

	drawingRouter.HandleFunc("GET /{prefix}/{id_from}/{id_to}", func(w http.ResponseWriter, r *http.Request) {
		api_key := r.URL.Query().Get("api_key")
		prefix := r.PathValue("prefix")
		id_from := r.PathValue("id_from")
		id_to := r.PathValue("id_to")

		i_id_from, err := strconv.Atoi(id_from)
		if err != nil {
			http.Error(w, "Invalid integer for the first number.", http.StatusBadRequest)
			return
		}

		i_id_to, err := strconv.Atoi(id_to)
		if err != nil {
			http.Error(w, "Invalid integer for the last number.", http.StatusBadRequest)
			return
		}

		rtnDrawing, err := GetRangeDrawing(api_key, prefix, i_id_from, i_id_to)
		if err != nil{
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
		json.NewEncoder(w).Encode(rtnDrawing)
	})

	drawingRouter.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var drawings []DrawingEntry
		json.NewDecoder(r.Body).Decode(&drawings)
		api_key := r.URL.Query().Get("api_key")

		_, err := PostDrawingEntries(api_key, drawings)
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
		json.NewEncoder(w).Encode(drawings)
	})

	return drawingRouter
}
