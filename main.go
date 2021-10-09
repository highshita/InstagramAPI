package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.client

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ere := client.Connect(ctx)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		log.Fatal("Connection to database failed", err)
	} else {
		log.Println("Connection to database successfull!")
	}

	r := NewRouter()
	r.Methods(http.MethodGet).Handler(`/`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This works!\n")
	}))

	//ADDING USER
	r.Methods(http.MethodPost).Handler(`/users`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		collection := client.Database("gramdb").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		doc := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}
		result, err := collection.InsertOne(ctx, doc)
		if err != nil {
			fmt.Fprint(w, "User creation failed!\n", result)
		} else {
			fmt.Fprint(w, "You are now an user!\n", result)
		}
	}))

	//ADDING POST
	r.Methods(http.MethodPost).Handler(`/posts`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		var post Posts
		json.NewDecoder(r.Body).Decode(&post)
		collection := client.Database("instadb").Collection("posts")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		doc := bson.M{"caption": post.Caption, "url": post.Url, "currentTime": post.CurrentTime, "userID": post.UserID}
		result, err := collection.InsertOne(ctx, doc)
		if err != nil {
			fmt.Fprint(w, "Post creation failed!\n", result)
		} else {
			fmt.Fprint(w, "Post Created successfully!\n", result)
		}
	}))

	//GET ALL POSTS UNDER OF AN USER
	r.Methods(http.MethodGet).Handler(`/posts/users/:id`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		findAllPosts(w, r, id)
	}))

	//SEARCH FOR AN USER 
	r.Methods(http.MethodGet).Handler(`/users/:id`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		findUser(w, r, id)
	}))

	//GET POST WITH ID
	r.Methods(http.MethodGet).Handler(`/posts/:id`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		findPost(w, r, id)
	}))

	http.ListenAndServe(":9999", r)
}

}