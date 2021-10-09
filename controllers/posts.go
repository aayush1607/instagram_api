package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/aayush1607/instagram_api/config"
	"github.com/aayush1607/instagram_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		postsCollection := config.MI.DB.Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var post models.Post

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

		findResult := postsCollection.FindOne(ctx, bson.M{"_id": oid})
		if err := findResult.Err(); err != nil {
			w.WriteHeader(404)
			return
		}

		err = findResult.Decode(&post)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		postsCollection := config.MI.DB.Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		post := models.Post{}

		json.NewDecoder(r.Body).Decode(&post)

		var empty_time time.Time
		if post.Timestamp == empty_time {
			post.Timestamp = time.Now()
		}

		result, err := postsCollection.InsertOne(ctx, post)
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

func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		postsCollection := config.MI.DB.Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var posts []models.Post

		id, ok := r.URL.Query()["id"]
		pg, page_ok := r.URL.Query()["page"]
		lm, limit_ok := r.URL.Query()["limit"]

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

		filter := bson.M{"user": oid}
		findOptions := options.Find()
		var page, limitVal int
		if !page_ok || len(pg[0]) < 1 {
			page = 1
		} else {
			page, _ = strconv.Atoi(pg[0])
		}
		if !limit_ok || len(lm[0]) < 1 {
			limitVal = 1
		} else {
			limitVal, _ = strconv.Atoi(lm[0])
		}

		var limit int64 = int64(limitVal)

		total, _ := postsCollection.CountDocuments(ctx, filter)

		findOptions.SetSkip((int64(page) - 1) * limit)
		findOptions.SetLimit(limit)

		cursor, err := postsCollection.Find(ctx, filter, findOptions)
		defer cursor.Close(ctx)

		if err != nil {
			w.WriteHeader(404)
			return
		}

		for cursor.Next(ctx) {
			var post models.Post
			cursor.Decode(&post)
			posts = append(posts, post)
		}

		last := math.Ceil(float64(total / limit))
		if last < 1 && total > 0 {
			last = 1
		}
		post_reponse := models.PostsResponse{
			Posts:     posts,
			Total:     total,
			Page:      int64(page),
			Last_page: int64(last),
			Limit:     limit,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post_reponse)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}
