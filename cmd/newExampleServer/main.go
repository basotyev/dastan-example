package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newExampleServer/internal/app"
	"newExampleServer/internal/app/handler"
	"newExampleServer/pkg"
)

func main() {
	db, err := pkg.NewPostgresConnection("postgres://postgres-user:postgres-pass@localhost:5434/test-db?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return
	}
	di := app.NewDI(db)
	router := gin.New()
	handler.InitRoutes(router, di)
	http.ListenAndServe("localhost:8080", router)
}
