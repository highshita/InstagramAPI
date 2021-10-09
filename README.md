# InstagramAPI
A Instagram API made for submission for internship opportunity at Appointy.

Task  | Instagram Backend API

The task is to develop a basic version of aInstagram. You are only required to develop the API for the system. Below are the details.

You are required to Design and Develop an HTTP JSON API capable of the following operations,
Create an User
Should be a POST request
Use JSON request body
URL should be ‘/users'
Get a user using id
Should be a GET request
Id should be in the url parameter
URL should be ‘/users/<id here>’
Create a Post
Should be a POST request
Use JSON request body
URL should be ‘/posts'
Get a post using id
Should be a GET request
Id should be in the url parameter
URL should be ‘/posts/<id here>’
List all posts of a user
Should be a GET request
URL should be ‘/posts/users/<Id here>'

Additional Constraints/Requirements:
The API should be developed using Go.
MongoDB should be used for storage.
Only packages/libraries listed here and here can be used.

Scoring:
Completion Percentage
No. of working endpoints among the ones listed above.
Quality of Code
Reusability
Consistency in naming variables, methods, functions, types
Idiomatic i.e. in Go’s style
Passwords should be securely stored such they can't be reverse engineered
Make the server thread safe
Add pagination to the list endpoint
Add unit tests
Resources:
Completing the Golang tour should give one a good grip over the language. Do this well and you will complete the task with ease.
This book covers both the workings of web and Go based servers.

Users should have the following attributes
Id
Name
Email
Password

Posts should have the following Attributes. All fields are mandatory unless marked optional:
Id
Caption
Image URL
Posted Timestamp

