package api

import (
	"encoding/json"
	"net/http"
	"searchRest/renderer"
	"searchRest/types"
	"strconv"
)

type APIResonse struct {
	Name   string        `json:"name"`
	Total  int           `json:"total"`
	Places []types.Place `json:"places"`
}

type APIError struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(APIError{Error: err})
}

func ApiResponse(w http.ResponseWriter, r *http.Request) {
	hasPage := r.URL.Query().Has("page")
	if !hasPage {
		// http.Error(w, "Missing 'page' param", http.StatusBadRequest)
		JSONError(w, "Missing 'page' param", http.StatusBadRequest)
		return
	}
	pageNum := r.URL.Query().Get("page")
	intPageNum, err := strconv.Atoi(pageNum)
	if err != nil {
		// http.Error(w, "'page' param is not an int", http.StatusBadRequest)
		JSONError(w, "'page' param is not an int: "+pageNum, http.StatusBadRequest)
		return
	}

	if intPageNum < 1 {
		// http.Error(w, "'page' param should be >= 1", http.StatusBadRequest)
		JSONError(w, "'page' param should be >= 1", http.StatusBadRequest)
		return
	}

	places, totalEntries, err := renderer.INDEXNAME.GetPlaces(renderer.PAGINATION_LIMIT, (intPageNum-1)*renderer.PAGINATION_LIMIT)
	if err != nil {
		JSONError(w, "Error during getting values "+err.Error(), http.StatusBadRequest)
		return
	}

	if intPageNum > totalEntries/renderer.PAGINATION_LIMIT+1 {
		JSONError(w, "'page' param exceeds max number of pages", http.StatusBadRequest)
		return
	}

	response := APIResonse{Name: string(renderer.INDEXNAME), Total: totalEntries, Places: places}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
