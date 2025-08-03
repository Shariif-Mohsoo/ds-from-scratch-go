package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		//log.Fatal is to immediatelly exit from program.
		log.Fatal("Port is not found in the environment")
	}
	/*
		At initial PORT willn't exist in current shell session to add this, in terminal run
		export PORT=Number(e.g; 100,2000,3000,etc)
		OR GET THE PACKAGE TO DO IT AUTOMATICALLY
		go get github.com/joho/godotenv
		go mod vendor.

		NOW LET'S INSTALL ROUTER:
		go get github.com/go-chi/chi
		go get github.com/go-chi/cors
		go mod vendor
		go mod tidy
		go mod vendor

		NOTE:
		IN THIS SERVER ALL THE REQUESTS WILL COME IN AND OUT IN JSON FORMAT.
	*/

	//creating the router
	router := chi.NewRouter()

	//people make requests to the server from browser.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	//to check whether the server is live or not.
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	//path will be /v1/ready
	router.Mount("/v1", v1Router)

	//setting up the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	//start handling http requests.
	log.Printf("Server starting at port %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}
