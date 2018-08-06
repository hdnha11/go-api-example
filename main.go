package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hdnha11/go-api-example/project"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	godotenv.Load()

	// Connect to MongoDB
	session, err := mgo.Dial(os.Getenv("MONGODB_HOST"))
	if err != nil {
		log.Fatal("Cannot connect to MongoDB: ", err)
	}
	defer session.Close()
	db := session.DB(os.Getenv("MONGODB_DATABASE"))

	router := httprouter.New()

	// Projects
	projectRepository := project.NewMongoRepository(db)
	projectService := project.NewService(projectRepository)
	project.InitRouters(router, projectService)

	// Ping
	router.GET("/status", func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		res.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	log.Printf("GoAPIExample is listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
