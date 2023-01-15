package main

import (
	"log"
	"os"

	"github.com/PChanida-B/simple-go-restapi-microservice/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	r := router.NewRouter()
	/*New Handler*/
	r.ListenAndServe()()
}
