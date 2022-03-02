package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"coin/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(dbType, dbUsername, dbPassword, dbHost, dbPort, authdb, dbname string) (*mongo.Database, error) {
	helper.LogEvent("INFO", "Establishing mongoDB connection with given credentials...")

	var mongoCredentials, authSource string

	if dbUsername != "" && dbPassword != "" {
		mongoCredentials = fmt.Sprint(dbUsername, ":", dbPassword, "@")
		authSource = fmt.Sprint("authSource=", authdb, "&")
	}

	mongoUrl := fmt.Sprint(dbType, "://", mongoCredentials, dbHost, ":", dbPort, "/?", authSource, "directConnection=true&serverSelectionTimeoutMS=2000")

	clientOptions := options.Client().ApplyURI(mongoUrl)

	helper.LogEvent("INFO", "Connecting to MongoDB...")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.MongoDBError, "TEST"+err.Error()))
		return &mongo.Database{}, err
	}

	// Check the connection
	helper.LogEvent("INFO", "Confirming MongoDB Connection...")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.MongoDBError, err.Error()))
		return &mongo.Database{}, err
	}

	//helper.LogEvent("Info", "Connected to MongoDB!")
	helper.LogEvent("INFO", "Establishing Database collections and indexes...")

	conn := client.Database(dbname)

	return conn, nil
}

func CreateIndex(collection *mongo.Collection, field string, unique bool) bool {

	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		fmt.Println(err.Error())

		return false
	}
	return true
}

func GetPage(page string) (*options.FindOptions, error) {
	if page == "all" {
		return nil, nil
	}
	var limit, e = strconv.ParseInt(helper.Config.PageLimit, 10, 64)
	var pageSize, ee = strconv.ParseInt(page, 10, 64)
	if e != nil || ee != nil {
		return nil, helper.ErrorMessage(helper.NoRecordFound, "Error in page-size or limit-size.")
	}
	findOptions := options.Find().SetLimit(limit).SetSkip(limit * (pageSize - 1))
	return findOptions, nil
}
