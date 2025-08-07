package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shariif-Mohsoo/ds-from-scratch-go/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/*
{
  "previewLimit": 50,
  "server": "localhost",
  "port": 5432,
  "driver": "PostgreSQL",
  "name": "localhost",
  "database": "rssagg",
  "username": "postgres",
  "password": "mohsoo"
}
postgres://username:password@localhost:5432/databasename?sslmode=disable
*/

type apiConfig struct {
	DB *database.Queries //it is going to hold the connection with db
}

func main() {
	// fmt.Println("Hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		//log.Fatal is to immediatelly exit from program.
		log.Fatal("Port is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
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

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	//path will be /v1/ready
	router.Mount("/v1", v1Router)

	//setting up the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	//start handling http requests.
	log.Printf("Server starting at port %v", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}
