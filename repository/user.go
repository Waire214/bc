package repository

import (
	"coin/graph/model"
	"coin/helper"
	"coin/ports"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx    = context.TODO()
	source string
	target string
	val    bool
)

type User struct {
	ID                string `json:"id" bson:"id"`
	Name              string `json:"name" bson:"name"`
	IsVerified        *bool  `json:"is_Verified" bson:"isverified"`
	BankName          string `json:"bank_name" bson:"bankname"`
	BankCode          string `json:"bank_code" bson:"bankcode"`
	BankAccountNumber string `json:"bank_account_number" bson:"bankaccountnumber"`
}

type paystack struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankId        int    `json:"bank_id"`
}
type userInfra struct {
	UserCollection *mongo.Collection
}

func NewUserRepository(conn *mongo.Database) ports.UserRepository {
	return &userInfra{
		UserCollection: conn.Collection("user"),
	}
}

func (collection *userInfra) AddUser(user model.UserInput) (*model.User, error) {
	helper.LogEvent("INFO", "Persisting new user")

	user.Name, user.IsVerified = ValidateData(user)
	newUser := model.User(user)
	dbUser := User(user)
	_, err := collection.UserCollection.InsertOne(
		ctx,
		dbUser,
	)
	if err != nil {
		helper.LogEvent("ERROR", helper.MongoDBError)
		return &model.User{}, helper.ErrorMessage(helper.MongoDBError, err.Error())
	}

	helper.LogEvent("INFO", "Persisting new user successful")
	return &newUser, nil
}

func (collection *userInfra) GetUser(accountNumber string, bankCode string) (string, error) {
	helper.LogEvent("INFO", "Retrieving user info with query bank_code: "+bankCode+" and bank_account_number "+accountNumber)

	var user = User{}
	filter := bson.M{"bankcode": bankCode, "bankaccountnumber": accountNumber}
	err := collection.UserCollection.FindOne(ctx, filter).Decode(&user)

	if err != nil || user == (User{}) {
		helper.LogEvent("ERROR", helper.NoRecordFound)
		return "name record not found", helper.ErrorMessage(helper.NoRecordFound, helper.NoRecordFound)
	}

	helper.LogEvent("INFO", "Retrieving user info with query bank_code: "+bankCode+" and bank_account_number "+accountNumber+" completed successfully")
	return user.BankName, nil
}

func ValidateData(user model.UserInput) (string, *bool) {
	var payStackData paystack
	url := "https://api.paystack.co/bank/resolve?account_number=" + user.BankAccountNumber + "&bank_code=" + user.BankCode
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		log.Fatalln(err)
	}
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{helper.Config.BearerToken},
	}
	res, err := client.Do(req)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		log.Fatalln(err)
	}
	json.Unmarshal(body, &payStackData)

	if user.Name != payStackData.Data.AccountName {
		source = user.Name
		target = payStackData.Data.AccountName
		distance := levenshtein.DistanceForStrings([]rune(strings.ToUpper(source)), []rune(target), levenshtein.DefaultOptions)
		if distance <= 2 {
			val = true
			user.IsVerified = &val
			return source, user.IsVerified
		}
		if distance >= 3 {
			val = false
			user.IsVerified = &val
			return source, user.IsVerified
		}
	}
	return user.Name, user.IsVerified
}
