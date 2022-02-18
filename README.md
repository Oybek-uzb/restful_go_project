# restful-go-project

# user-service

# REST API

GET /users -- a list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204/200
PATCH /users/:id -- partially update user -- 204/200

