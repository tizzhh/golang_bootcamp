package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	COOL_ESKIMO         string = "CE"
	APRICOT_AARDVARK    string = "AA"
	NATURAL_TIGER       string = "NT"
	DAZZLING_ELDERBERRY string = "DE"
	YELLOW_RAMBUTAN     string = "YR"
)

var shopPrices = map[string]float64{
	COOL_ESKIMO:         10.0,
	APRICOT_AARDVARK:    15.0,
	NATURAL_TIGER:       17.0,
	DAZZLING_ELDERBERRY: 21.0,
	YELLOW_RAMBUTAN:     23.0,
}

type apiFormat struct {
	Money      float64 `json:"money"`
	CandyType  string  `json:"candyType"`
	CandyCount int     `json:"candyCount"`
}

type successReponse struct {
	Message string  `json:"thanks"`
	Change  float64 `json:"change"`
}

type apiError struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(apiError{Error: err})
}

func CandyShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	var body apiFormat

	if r.Method != http.MethodPost {
		JSONError(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		JSONError(w, "wrong API format", http.StatusBadRequest)
		return
	}

	if _, ok := shopPrices[body.CandyType]; !ok {
		JSONError(w, "unknown candy Type", http.StatusBadRequest)
		return
	}

	reqMoney := shopPrices[body.CandyType] * float64(body.CandyCount)
	if body.Money < reqMoney {
		notEnoughMoney := fmt.Sprintf("You need %f more money!", reqMoney-body.Money)
		JSONError(w, notEnoughMoney, http.StatusPaymentRequired)
		return
	}

	json.NewEncoder(w).Encode(successReponse{Message: "Thank you!", Change: body.Money - reqMoney})
}
