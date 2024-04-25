package db

import (
	"context"
	"fmt"
	"myArticles/types"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Db            *pgx.Conn
	TotalArticles int
}

var (
	pgInstance *Postgres
)

const (
	MAX_LEN_OF_TEXT int = 50
)

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	var err error = nil
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection: %w", err)
	}
	pgInstance = &Postgres{conn, 0}

	return pgInstance, nil
}

func (pg *Postgres) GetArticle(id int) (types.ArticleData, error) {
	var article types.ArticleData
	rows, err := getRows(pg, fmt.Sprintf(`SELECT * FROM article WHERE id = %d`, id))
	if err != nil {
		return article, err
	}
	defer rows.Close()

	if !rows.Next() {
		return article, pgx.ErrNoRows
	}

	err = rows.Scan(&article.Id, &article.PostDate, &article.Title, &article.Text)
	if err != nil {
		return article, fmt.Errorf("unable to scan row: %w", err)
	}

	return article, nil
}

func (pg *Postgres) GetArticles(limit, offset int) ([]types.ArticleData, error) {
	rows, err := getRows(pg, fmt.Sprintf(`SELECT * FROM article OFFSET %d LIMIT %d`, offset, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []types.ArticleData{}
	for rows.Next() {
		article := types.ArticleData{}
		err := rows.Scan(&article.Id, &article.PostDate, &article.Title, &article.Text)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		if len(article.Text) >= MAX_LEN_OF_TEXT {
			article.Text = article.Text[:MAX_LEN_OF_TEXT] + "..."
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (pg *Postgres) CreateArticle(title, text string) (int, error) {
	var id int
	err := pg.Db.QueryRow(context.Background(), fmt.Sprintf(`INSERT INTO article (title, article_text) VALUES ('%s', '%s') RETURNING id`, title, text)).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to create an article: %w", err)
	}

	return id, nil
}

func (pg *Postgres) GetTotalNumOfArticles() error {
	rows, err := getRows(pg, `SELECT COUNT(*) FROM article`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var total int
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return fmt.Errorf("unable to scan row: %w", err)
		}
	}

	pg.TotalArticles = total
	return nil
}

func getRows(pg *Postgres, query string) (pgx.Rows, error) {
	rows, err := pg.Db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to scan row: %w", err)
	}

	return rows, nil
}
