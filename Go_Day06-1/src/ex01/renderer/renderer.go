package renderer

import (
	"fmt"
	"html/template"
	"myArticles/types"
	"net/http"
)

const (
	INDEX_TEMPLATE_PATH string = "templates/index.html"
)

func GetIndexArticles(w http.ResponseWriter, articles []types.ArticleData, curPage, totalPages int) error {
	pageData := types.ArticlePage{
		CurPage:    curPage,
		PrevPage:   curPage - 1,
		NextPage:   curPage + 1,
		TotalPages: totalPages,
		Articles:   articles,
	}
	tmpl, err := template.ParseFiles(INDEX_TEMPLATE_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}
