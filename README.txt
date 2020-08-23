#### Run the project in cmd(in project folder): ./dev
     The app will run in port :3000

#### Routes:
_______________________________________________________________________________
POST   /users/register           --> go-gin-crud-api/routes.UsersRegister     | registers a new user 
POST   /users/login              --> go-gin-crud-api/routes.UsersLogin        | login
GET    /posts/index              --> go-gin-crud-api/routes.PostsIndex        | gets all posts
GET    /posts/index/:id          --> go-gin-crud-api/routes.PostByID          | gets a post using id in url
POST   /posts/create             --> go-gin-crud-api/routes.PostsCreate       | creates a post (with auth)
GET    /posts/myposts            --> go-gin-crud-api/routes.PostsByCurrentUser| gets all posts of the user (with auth)
PUT    /posts/update             --> go-gin-crud-api/routes.PostsUpdate       | updates the post using id in parameters 
                                                                              | (with auth)
DELETE /posts/delete/:id         --> go-gin-crud-api/routes.PostsDelete       | deletes the post using id in parameters ______________________________________________________________________________| (with auth)

Reigitster request body:                      password, password_confirm, email
Login request body:                           email, password
Create Post request body:                     title, body; Authorization: Bearer token
Get Posts of the current user Authorization: Bearer token
Update Post request body:                     id of the post, title, body
Delete Post Authorization:                   Bearer token

#### Other details

In project folder there there are 2 files provided by me: 
 - postman collection (go-gin-crud-api.postman_collection.json)
 - database file to import (exported_database.sql)

Import the postman collection in Postman. There are ready requests.
Import the database file in pgAdmin or shell to the database named goBlog.

* Required: to name the database as 'goBlog'
* In project username of PostgreSQL is default: 'postgres', and the password: 'password'