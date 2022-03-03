package main

import (
	"coin/graph"
	"coin/graph/generated"
	"coin/helper"
	"coin/repository"
	"coin/services"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
)

const defaultPort = "8080"

func main() {
	helper.InitializeLog()

	db := startMongoDB()
	router := chi.NewRouter()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	resolver := graph.NewResolverHandler(userService)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Use(
		// Middleware,
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log api request calls
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func startMongoDB() *mongo.Database {
	helper.LogEvent("INFO", "Initializing Mongo!")
	db, err := repository.ConnectToMongo(helper.Config.DbType, helper.Config.MongoDbUserName, helper.Config.MongoDbPassword, helper.Config.MongoDbHost, helper.Config.MongoDbPort, helper.Config.MongoDbAuthDb, helper.Config.MongoDbName)
	if err != nil {
		helper.LogEvent("ERROR", "MongoDB database connection error: "+err.Error())
		log.Fatal()
	}
	return db
}

