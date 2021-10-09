package main

import (
	"fmt"
	"net/http"
)

func findUser(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "Find user with id : ", id)
}

func findPost(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "Find post with id", id)
}

func findAllPosts(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "Find All post with user id", id)
}
