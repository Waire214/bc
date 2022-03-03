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

	"github.com/texttheater/golang-levenshtein/levenshtein"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx    = context.TODO()
	source string
	target string
	val    bool
)

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

	_, err := collection.UserCollection.InsertOne(
		ctx,
		user,
	)
	if err != nil {
		return &model.User{}, helper.ErrorMessage(helper.MongoDBError, err.Error())
	}

	helper.LogEvent("INFO", "Persisting new user successful")
	newUser := model.User(user)
	return &newUser, nil
}

func ValidateData(user model.UserInput) (string, *bool) {
	var payStackData paystack

	req, err := http.Get("https://api.paystack.co/bank/resolve?account_number=" + user.BankAccountNumber + "&bank_code=" + user.BankCode)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer sk_test_87f0bed09dbdc1b48465b3835130184eeda71589")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &payStackData)

	if user.Name != payStackData.Data.AccountName {
		source = user.Name
		target = payStackData.Data.AccountName
		distance := levenshtein.DistanceForStrings([]rune(source), []rune(target), levenshtein.DefaultOptions)
		if distance <= 2 {
			val = true
			user.IsVerified = &val
			log.Println(target)
			log.Println(source)
			return source, user.IsVerified
		} else {
			val = false
			user.IsVerified = &val
			return source, user.IsVerified
		}
	}
	return user.Name, user.IsVerified
}
