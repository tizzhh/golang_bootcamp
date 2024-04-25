package renderer

import (
	"fmt"
	"html/template"
	"myArticles/types"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const (
	INDEX_TEMPLATE_PATH   string = "templates/index.html"
	ARTICLE_TEMPLATE_PATH string = "templates/article.html"
	BASE_TEMPLATE_PATH    string = "templates/base.html"
	ERROR_TEMPLATE_PATH   string = "templates/error.html"
	ADMIN_LOGIN_PATH      string = "templates/log_in_form.html"
	ADD_ARTICLE_FORM      string = "templates/create_article.html"
)

func RenderError(w http.ResponseWriter, code int, msg string) error {
	pageData := types.ErrorData{
		Code:    code,
		Message: msg,
	}
	tmpl, err := template.ParseFiles(BASE_TEMPLATE_PATH, ERROR_TEMPLATE_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func RenderIndexArticles(w http.ResponseWriter, articles []types.ArticleData, curPage, totalPages int) error {
	pageData := types.ArticlesPage{
		CurPage:    curPage,
		PrevPage:   curPage - 1,
		NextPage:   curPage + 1,
		TotalPages: totalPages,
		Articles:   articles,
	}
	tmpl, err := template.ParseFiles(BASE_TEMPLATE_PATH, INDEX_TEMPLATE_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func RenderArticle(w http.ResponseWriter, article types.ArticleData, prevPageNum int) error {
	article.Text = template.HTML(mdToHtml([]byte(article.Text)))

	pageData := types.ArticlePage{
		Article: types.ArticleData{
			Id:       article.Id,
			PostDate: article.PostDate,
			Title:    article.Title,
			Text:     article.Text,
		},
		PrevPageNum: prevPageNum,
	}
	tmpl, err := template.ParseFiles(BASE_TEMPLATE_PATH, ARTICLE_TEMPLATE_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func RenderAdminLogInForm(w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles(BASE_TEMPLATE_PATH, ADMIN_LOGIN_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, "")
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func RenderAddArticleForm(w http.ResponseWriter, id int) error {
	tmpl, err := template.ParseFiles(BASE_TEMPLATE_PATH, ADD_ARTICLE_FORM)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, id)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func mdToHtml(md []byte) []byte {
	extentions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extentions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
