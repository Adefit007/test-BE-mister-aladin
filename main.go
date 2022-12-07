package main

import (
	"article/database"
	"article/pkg/mysql"
	"article/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// On http (API)
	r := mux.NewRouter()

	mysql.DatabaseInit()

	database.RunMigration()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}