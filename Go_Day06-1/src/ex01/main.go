package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"myArticles/db"
	"myArticles/renderer"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	PAGE_LIMIT       int    = 3
	CREDENTIALS_PATH string = "admin_credentials.txt"
)

type Config struct {
	adminUser, adminPass, dbUser, dbPass string
}

type App struct {
	Router *mux.Router
	DB     *db.Postgres
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()

	cfg, err := getCfg()
	URL := fmt.Sprintf(`postgres://%s:%s@localhost:5432/golang_day06`, cfg.dbUser, cfg.dbPass)
	if err != nil {
		log.Fatalf("Could not read credentials: %s\n", err.Error())
	}

	var app App
	err = app.Init(ctx, URL)
	if err != nil {
		log.Fatalf("Error during initialization: %s\n", err.Error())
	}
	defer app.DB.Db.Close(context.Background())
	app.Run()
}

func (a *App) Init(ctx context.Context, URL string) error {
	db, err := db.NewPG(ctx, URL)
	if err != nil {
		log.Fatalf("Error during connection creation: %s\n", err.Error())
	}
	a.DB = db
	a.DB.GetTotalNumOfArticles()

	a.Router = mux.NewRouter()
	a.InitRoutes()

	return nil
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8888", a.Router))
}

func (a *App) InitRoutes() {
	a.Router.HandleFunc("/", a.GetArticles).Methods("GET")
}

func (a *App) GetArticles(w http.ResponseWriter, r *http.Request) {
	pageNum := r.URL.Query().Get("page")
	intPageNum, err := strconv.Atoi(pageNum)
	if err != nil {
		if !r.URL.Query().Has("page") {
			intPageNum = 1
		} else {
			RespondWithError(w, "Page should be an int: "+pageNum, http.StatusBadRequest)
			return
		}
	}
	if intPageNum < 1 || intPageNum > (a.DB.TotalArticles/PAGE_LIMIT+1) {
		RespondWithError(w, "Page should be in [1; maxPages]", http.StatusBadRequest)
		return
	}
	offset := (intPageNum - 1) * PAGE_LIMIT
	articles, err := a.DB.GetArticles(PAGE_LIMIT, offset)
	if err != nil {
		log.Println("Error during db query", err.Error())
		RespondWithError(w, "Server error", http.StatusBadGateway)
		return
	}
	err = renderer.GetIndexArticles(w, articles, intPageNum, a.DB.TotalArticles/PAGE_LIMIT+1)
	if err != nil {
		log.Println("Error rendering index with GET: ", err.Error())
		RespondWithError(w, "Server error", http.StatusBadGateway)
		return
	}
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
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
