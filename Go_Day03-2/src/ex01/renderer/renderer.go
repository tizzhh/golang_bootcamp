package renderer

import (
	"net/http"
	"searchRest/db"
	"searchRest/types"
	"strconv"
	"text/template"
)

const (
	PAGINATION_LIMIT int     = 10
	TEMPLATE_PATH    string  = "templates/index.html"
	INDEXNAME        db.Indx = "places"
)

func RenderPage(w http.ResponseWriter, r *http.Request) {
	hasPage := r.URL.Query().Has("page")
	if !hasPage {
		http.Error(w, "Missing 'page' param", http.StatusBadRequest)
		return
	}
	pageNum := r.URL.Query().Get("page")
	intPageNum, err := strconv.Atoi(pageNum)
	if err != nil {
		http.Error(w, "'page' param is not an int", http.StatusBadRequest)
		return
	}

	places, totalEntries, err := INDEXNAME.GetPlaces(PAGINATION_LIMIT, (intPageNum-1)*PAGINATION_LIMIT)
	if err != nil {
		http.Error(w, "Error during getting values "+err.Error(), http.StatusBadRequest)
		return
	}

	pageData := types.RestPage{
		TotalNumberOfRests: totalEntries,
		CurPage:            (intPageNum),
		PrevPage:           (intPageNum - 1),
		NextPage:           (intPageNum + 1),
		TotalPages:         totalEntries / PAGINATION_LIMIT,
		Rests:              places,
	}

	tmpl, err := template.ParseFiles(TEMPLATE_PATH)
	if err != nil {
		http.Error(w, "Error template creation "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Error template populating "+err.Error(), http.StatusInternalServerError)
		return
	}
}
