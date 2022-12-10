package controllers

import (
	"fmt"
	"net/http"

	"api-server/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	doc := bson.D{{Key: "name", Value: "king"}, {Key: "age", Value: 20}, {Key: "role", Value: "admin"}}
	database.TestCollection.InsertOne(database.CTX, doc)
}

func Update(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	fmt.Println(productId)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	fmt.Println(productId)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
