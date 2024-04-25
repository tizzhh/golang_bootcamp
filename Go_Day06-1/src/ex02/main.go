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

	"github.com/go-chi/httprate"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

const (
	PAGE_LIMIT       int    = 3
	CREDENTIALS_PATH string = "admin_credentials.txt"
	FAVICON_PATH     string = "images/amazing_logo.png"
)

type Config struct {
	adminUser, adminPass, dbUser, dbPass string
}

type App struct {
	Router *mux.Router
	DB     *db.Postgres
	Cfg    Config
}

var jwtKey = []byte("my_secret_key")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()

	cfg, err := getCfg()
	URL := fmt.Sprintf(`postgres://%s:%s@localhost:5432/golang_day06`, cfg.dbUser, cfg.dbPass)
	if err != nil {
		log.Fatalf("Could not read credentials: %s\n", err.Error())
	}

	var app App
	err = app.Init(ctx, URL, cfg)
	if err != nil {
		log.Fatalf("Error during initialization: %s\n", err.Error())
	}
	defer app.DB.Db.Close(context.Background())
	app.Run()
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8888", a.Router))
}

func (a *App) InitRoutes() {
	a.Router.HandleFunc("/", a.getArticlesHandler).Methods("GET")
	a.Router.HandleFunc("/article", a.getArticlesHandler).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.getArticleHandler).Methods("GET")
	a.Router.HandleFunc("/"+FAVICON_PATH, faviconHandler).Methods("GET")
	a.Router.HandleFunc("/admin", a.adminLogInHandler).Methods("GET", "POST")
	a.Router.HandleFunc("/admin/add_article", a.adminAddArticleHandler).Methods("GET", "POST")
}

func (a *App) Init(ctx context.Context, URL string, cfg Config) error {
	db, err := db.NewPG(ctx, URL)
	if err != nil {
		log.Fatalf("Error during connection creation: %s\n", err.Error())
	}
	a.DB = db
	a.DB.GetTotalNumOfArticles()

	a.Router = mux.NewRouter()
	a.Router.Use(httprate.Limit(
		100,
		1*time.Second,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			RespondWithError(w, "429 Too Many Requests", http.StatusTooManyRequests)
		}),
	))
	a.InitRoutes()

	a.Cfg = cfg

	return nil
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, FAVICON_PATH)
}

func (a *App) adminAddArticleHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("No token")
			RespondWithError(w, "No token", http.StatusUnauthorized)
			return
		}
		log.Println(err.Error())
		RespondWithError(w, "", http.StatusBadRequest)
		return
	}

	token := c.Value
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Wrong token")
			RespondWithError(w, "Wrong token", http.StatusUnauthorized)
			return
		}
		log.Println(err.Error())
		RespondWithError(w, "Wrong token", http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		RespondWithError(w, "Wrong token", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		err = renderer.RenderAddArticleForm(w, 0, http.StatusOK)
		if err != nil {
			log.Println("Error rendering article form with GET: ", err.Error())
			RespondWithError(w, "Server error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		id, err := a.DB.CreateArticle(r.Form.Get("title"), r.Form.Get("content"))
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, "Failed to create a new article", http.StatusInternalServerError)
			return
		}
		a.DB.GetTotalNumOfArticles()
		err = renderer.RenderAddArticleForm(w, id, http.StatusCreated)
		if err != nil {
			log.Println("Error rendering article form with POST: ", err.Error())
			RespondWithError(w, "Server error", http.StatusInternalServerError)
			return
		}
	}
}

func (a *App) adminLogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderer.RenderAdminLogInForm(w)
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, "", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if a.Cfg.adminPass != password || a.Cfg.adminUser != username {
			RespondWithError(w, "Wrong username and/or password", http.StatusUnauthorized)
			return
		}

		expirationDate := time.Now().Add(24 * time.Hour)
		claims := jwt.MapClaims{
			"exp":      expirationDate.Unix(),
			"username": username,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString(jwtKey)
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, "Token error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenStr,
			Expires: expirationDate,
		})

		http.Redirect(w, r, "/admin/add_article", http.StatusSeeOther)
	}
}

func (a *App) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, "Ivalid article id: "+vars["id"], http.StatusBadRequest)
		return
	}
	article, err := a.DB.GetArticle(id)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			RespondWithError(w, "Article not found", http.StatusNotFound)
		default:
			log.Println(err.Error())
			RespondWithError(w, "", http.StatusBadRequest)
		}
		return
	}
	pageNum := 1
	refererPage := r.Header.Get("Referer")
	if strings.Contains(refererPage, "?page=") {
		pageNum, _ = strconv.Atoi(strings.Split(refererPage, "?page=")[1])
	}
	err = renderer.RenderArticle(w, article, pageNum)
	if err != nil {
		log.Println("Error rendering article with GET: ", err.Error())
		RespondWithError(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func (a *App) getArticlesHandler(w http.ResponseWriter, r *http.Request) {
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
	if intPageNum < 1 || intPageNum > ((a.DB.TotalArticles-1)/PAGE_LIMIT+1) {
		RespondWithError(w, "Page should be in [1; maxPages]", http.StatusBadRequest)
		return
	}
	offset := (intPageNum - 1) * PAGE_LIMIT
	articles, err := a.DB.GetArticles(PAGE_LIMIT, offset)
	if err != nil {
		log.Println("Error during db query", err.Error())
		RespondWithError(w, "Server error", http.StatusInternalServerError)
		return
	}
	err = renderer.RenderIndexArticles(w, articles, intPageNum, (a.DB.TotalArticles-1)/PAGE_LIMIT+1)
	if err != nil {
		log.Println("Error rendering index with GET: ", err.Error())
		RespondWithError(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	err := renderer.RenderError(w, code, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
