package main

import (
	"log"
	"net/http"

	"github.com/aayush1607/instagram_api/config"
	"github.com/aayush1607/instagram_api/controllers"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/users", controllers.CreateUser)
	http.HandleFunc("/users/", controllers.GetUser)
	http.HandleFunc("/posts", controllers.CreatePost)
	http.HandleFunc("/posts/", controllers.GetPost)
	http.HandleFunc("/posts/users/", controllers.GetPostsByUser)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
