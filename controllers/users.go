package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aayush1607/instagram_api/config"
	"github.com/aayush1607/instagram_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		usersCollection := config.MI.DB.Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User

		id, ok := r.URL.Query()["id"]
		if !ok || len(id[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid, err := primitive.ObjectIDFromHex(id[0])
		if err != nil {
			log.Println("Invalid Param id")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		findResult := usersCollection.FindOne(ctx, bson.M{"_id": oid})
		if err := findResult.Err(); err != nil {
			w.WriteHeader(404)
			return
		}

		err = findResult.Decode(&user)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		usersCollection := config.MI.DB.Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		user := models.User{}
		json.NewDecoder(r.Body).Decode(&user)
		pass := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.Password = string(hashedPassword)
		result, err := usersCollection.InsertOne(ctx, user)
		if err != nil {
			w.WriteHeader(500)
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}
