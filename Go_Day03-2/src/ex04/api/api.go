package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"searchRest/db"
	"searchRest/renderer"
	"searchRest/types"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type APIResonse struct {
	Name   string        `json:"name"`
	Total  int           `json:"total,omitempty"`
	Places []types.Place `json:"places"`
}

type APIError struct {
	Error string `json:"error"`
}

type Token struct {
	Token string `json:"token"`
}

const (
	PageSearchName    string = "Places"
	ClosestSearchName string = "Recommendation"
)

var secretKey = []byte("secret-key")

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
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing auth token", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		http.Error(w, "Something wrong with token: "+err.Error(), http.StatusUnauthorized)
		return
	}

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

func GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24).Unix()})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		JSONError(w, "error during token creation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Token{tokenString})
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
