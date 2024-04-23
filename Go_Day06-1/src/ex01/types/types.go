package types

import "time"

type ArticleData struct {
	Id       int       `json:"id"`
	PostDate time.Time `json:"post_date"`
	Title    string    `json:"title"`
	Text     string    `json:"article_text"`
}

type ArticlePage struct {
	CurPage    int
	PrevPage   int
	NextPage   int
	TotalPages int
	Articles   []ArticleData
}
