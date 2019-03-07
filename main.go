package main

import (
	"log"
	"net/http"

	"github.com/radean0909/jumpcloud-hasher/controllers"
	"github.com/radean0909/jumpcloud-hasher/routes"
	"github.com/radean0909/jumpcloud-hasher/utils/hasher"
	"github.com/radean0909/jumpcloud-hasher/utils/queue"

	"github.com/radean0909/jumpcloud-hasher/utils/database"
)

func main() {
	db, err := database.Connect("test") // This could be an environment variable
	if err != nil {
		log.Fatal(err)
	}

	queue, err := queue.NewQueue(100, 10, hasher.Process)
	queue.Start()
	hashController := controllers.NewHashController(db, queue)
	statsController := controllers.NewStatsController(db, queue)
	shutdownController := controllers.NewShutdownController(db, queue)

	mux := http.NewServeMux()
	router := routes.NewRouter(mux)

	router.AddRoute("/hash", hashController.Create)
	router.AddRoute("/hash/", hashController.Get)
	router.AddRoute("/stats", statsController.Get)
	router.AddRoute("/shutdown", shutdownController.Get)

	router.ParseRoutes()

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
