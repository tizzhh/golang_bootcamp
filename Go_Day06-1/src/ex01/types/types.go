package types

import (
	"html/template"
	"time"
)

type ErrorData struct {
	Code    int
	Message string
}

type ArticleData struct {
	Id       int           `json:"id"`
	PostDate time.Time     `json:"post_date"`
	Title    string        `json:"title"`
	Text     template.HTML `json:"article_text"`
}

type ArticlesPage struct {
	CurPage    int
	PrevPage   int
	NextPage   int
	TotalPages int
	Articles   []ArticleData
}

type ArticlePage struct {
	Article     ArticleData
	PrevPageNum string
}
