package main

type Post struct {
	Caption     string `json:"caption" bson:"caption"`
	ImageUrl    string `json:"url" bson:"url"`
	TimeStamp   string `json:"timestamp" bson:"timestamp"`
}
type User struct {
	ID       string `json:id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Posts    []Post
}


