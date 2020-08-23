package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"go-gin-crud-api/models"
	"go-gin-crud-api/routes"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {

	conn, err := connectDB() // connect to database
	if err != nil {
		return
	}

	router := gin.Default() // creates a gin router with default middleware

	router.Use(dbMiddleware(*conn)) // use middleware in dbMiddleware

	usersGroup := router.Group("users") // group of routers
	{
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("login", routes.UsersLogin)
	}

	postsGroup := router.Group("posts") // group of routers
	{
		postsGroup.GET("index", routes.PostsIndex)
		postsGroup.GET("index/:id", routes.PostByID)
		postsGroup.POST("create", authMiddleWare(), routes.PostsCreate)
		postsGroup.GET("myposts", authMiddleWare(), routes.PostsByCurrentUser)
		postsGroup.PUT("update", authMiddleWare(), routes.PostsUpdate)
		postsGroup.DELETE("delete/:id", authMiddleWare(), routes.PostsDelete)
	}

	router.Run(":3000")
}

/*
	Function to connect the database using pgx/v4
	returns: pgx.Conn, error
*/
func connectDB() (c *pgx.Conn, err error) {
	database_url, _ := os.LookupEnv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

/*
	Handler function to name the connection
*/
func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

/*
	Handler function to check Authorization Header
	Bearer token validation
*/
func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
