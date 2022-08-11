package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
