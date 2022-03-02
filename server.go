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
	"go.mongodb.org/mongo-driver/mongo"
)

const defaultPort = "8080"

// var db *mongo.Database

// func init() {
// 	db = startMongoDB()
// 	fmt.Println("db printed")
// }
func main() {
	helper.InitializeLog()

	db := startMongoDB()
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	resolver := graph.NewResolverHandler(userService)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
