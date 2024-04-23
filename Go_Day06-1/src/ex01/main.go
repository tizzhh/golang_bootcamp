package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"myArticles/db"
	"os"
	"strings"
	"time"
)

const (
	PAGE_LIMIT       int    = 3
	CREDENTIALS_PATH string = "admin_credentials.txt"
)

type Config struct {
	adminUser, adminPass, dbUser, dbPass string
}

// This is worthless since the task makes me push the credentials for some reason. Should be in .gitignore
func getCfg() (Config, error) {
	var cfg Config
	file, err := os.Open(CREDENTIALS_PATH)
	if err != nil {
		return Config{}, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	fileScann := bufio.NewScanner(file)
	fileScann.Split(bufio.ScanLines)
	for fileScann.Scan() {
		line := fileScann.Text()
		parts := strings.Split(line, ":")
		switch parts[0] {
		case "adminUser":
			cfg.adminUser = parts[1]
		case "adminPass":
			cfg.adminPass = parts[1]
		case "dbUser":
			cfg.dbUser = parts[1]
		case "dbPass":
			cfg.dbPass = parts[1]
		}
	}

	return cfg, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()

	cfg, err := getCfg()
	URL := fmt.Sprintf(`postgres://%s:%s@localhost:5432/golang_day06`, cfg.dbUser, cfg.dbPass)
	if err != nil {
		log.Fatalf("Could not read credentials: %s\n", err.Error())
	}

	conn, err := db.NewPG(ctx, URL)
	if err != nil {
		log.Fatalf("Error during connection creation: %s\n", err.Error())
	}
	fmt.Println(conn.GetArticles(ctx, PAGE_LIMIT, 0))

	conn.CloseConn(ctx)
}
