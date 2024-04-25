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
)

func RenderIndexArticles(w http.ResponseWriter, articles []types.ArticleData, curPage, totalPages int) error {
	pageData := types.ArticlesPage{
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

func RenderArticle(w http.ResponseWriter, article types.ArticleData, prevPageNum string) error {
	article.Text = template.HTML(MdToHtml([]byte(article.Text)))

	pageData := types.ArticlePage{
		Article: types.ArticleData{
			Id:       article.Id,
			PostDate: article.PostDate,
			Title:    article.Title,
			Text:     article.Text,
		},
		PrevPageNum: prevPageNum,
	}
	tmpl, err := template.ParseFiles(ARTICLE_TEMPLATE_PATH)
	if err != nil {
		return fmt.Errorf("error during template creation: " + err.Error())
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		return fmt.Errorf("error during template populating: " + err.Error())
	}

	return nil
}

func MdToHtml(md []byte) []byte {
	extentions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extentions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
