package api

import (
	"encoding/json"
	"net/http"
	"searchRest/db"
	"searchRest/renderer"
	"searchRest/types"
	"strconv"
)

type APIResonse struct {
	Name   string        `json:"name"`
	Total  int           `json:"total,omitempty"`
	Places []types.Place `json:"places"`
}

type APIError struct {
	Error string `json:"error"`
}

const (
	PageSearchName    string = "Places"
	ClosestSearchName string = "Recommendation"
)

func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(APIError{Error: err})
}

func ApiResponse(w http.ResponseWriter, r *http.Request) {
	hasPage := r.URL.Query().Has("page")
	if !hasPage {
		JSONError(w, "Missing 'page' param", http.StatusBadRequest)
		return
	}
	pageNum := r.URL.Query().Get("page")
	intPageNum, err := strconv.Atoi(pageNum)
	if err != nil {
		JSONError(w, "'page' param is not an int: "+pageNum, http.StatusBadRequest)
		return
	}

	if intPageNum < 1 {
		JSONError(w, "'page' param should be >= 1", http.StatusBadRequest)
		return
	}

	places, totalEntries, err := renderer.INDEXNAME.GetPlaces(renderer.PAGINATION_LIMIT, (intPageNum-1)*renderer.PAGINATION_LIMIT, 0, 0, db.MODE_PAGE)
	if err != nil {
		JSONError(w, "Error during getting values: "+err.Error(), http.StatusBadRequest)
		return
	}

	if intPageNum > totalEntries/renderer.PAGINATION_LIMIT+1 {
		JSONError(w, "'page' param exceeds max number of pages", http.StatusBadRequest)
		return
	}

	response := APIResonse{Name: PageSearchName, Total: totalEntries, Places: places}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func ApiClosestRests(w http.ResponseWriter, r *http.Request) {
	hasLat := r.URL.Query().Has("lat")
	hasLon := r.URL.Query().Has("lon")
	if !hasLat || !hasLon {
		JSONError(w, "Missing 'lat' or 'lon' param", http.StatusBadRequest)
		return
	}

	latVal := r.URL.Query().Get("lat")
	latValFloat, err := strconv.ParseFloat(latVal, 64)
	if err != nil {
		JSONError(w, "'lat' param is not float: "+latVal, http.StatusBadRequest)
		return
	}
	lonVal := r.URL.Query().Get("lat")
	lonValFloat, err := strconv.ParseFloat(lonVal, 64)
	if err != nil {
		JSONError(w, "'lon' param is not float: "+lonVal, http.StatusBadRequest)
		return
	}

	places, _, err := renderer.INDEXNAME.GetPlaces(renderer.CLOSESTS_LIMIT, 0, latValFloat, lonValFloat, db.MODE_CLOSEST)
	if err != nil {
		JSONError(w, "Error searching closest places: "+err.Error(), http.StatusBadRequest)
		return
	}

	response := APIResonse{Name: ClosestSearchName, Total: 0, Places: places}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
