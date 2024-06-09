package main

import (
	"fmt"
	"github/shubash/filemanger/router"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	r := router.Router()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8081",handler))
	fmt.Println("server is ready")
}