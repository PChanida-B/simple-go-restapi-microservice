package main

import (
	"log"
	"os"

	"github.com/PChanida-B/simple-go-restapi-microservice/auth"
	"github.com/PChanida-B/simple-go-restapi-microservice/router"
	"github.com/PChanida-B/simple-go-restapi-microservice/service"
	"github.com/PChanida-B/simple-go-restapi-microservice/store"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {
	r := router.NewRouter()

	/*New Handler*/
	handler := service.NewHandler(store.NewMariaDBStore(os.Getenv("DSN")))

	authHandler := service.NewAuthHandler(auth.NewJWTAuth(os.Getenv("SIGN")))
	r.GET("/tokenz", authHandler.AccessToken)

	/*Routing*/
	auth := r.Group("", authHandler.Authorization)
	auth.POST("/resources", handler.CreateHandler)
	auth.GET("/resources", handler.ReadAllHandler)
	auth.GET("/resources/:id", handler.ReadHandler)
	auth.DELETE("/resources/:id", handler.DeleteHandler)
	auth.PUT("/resources/:id", handler.UpdateHandler)

	r.ListenAndServe()()
}
