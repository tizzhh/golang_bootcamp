package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type articleData struct {
	Id       int
	PostDate time.Time
	Title    string
	Text     string
}

type postgres struct {
	// db *pgxpool.Pool
	db *pgx.Conn
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connString string) (*postgres, error) {
	var err error = nil
	// pgOnce.Do(func() {
	// 	var db *pgxpool.Pool
	// 	db, err = pgxpool.New(ctx, connString)

	// pgInstance = &postgres{db}
	// })
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection: %w", err)
	}
	pgInstance = &postgres{conn}

	return pgInstance, nil
}

func (pg *postgres) CloseConn(ctx context.Context) {
	// pg.db.Close()
	pg.db.Close(ctx)
}

func (pg *postgres) GetArticles(ctx context.Context, limit, offset int) ([]articleData, error) {
	query := fmt.Sprintf(`SELECT * FROM article WHERE id > %d LIMIT %d`, offset, limit)
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query articles: %w", err)
	}
	defer rows.Close()

	articles := []articleData{}
	for rows.Next() {
		article := articleData{}
		err := rows.Scan(&article.Id, &article.PostDate, &article.Title, &article.Text)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}
