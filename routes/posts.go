package routes

import (
	"fmt"
	"go-gin-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func PostsIndex(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	posts, err := models.GetAllPosts(&conn)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostByID(c *gin.Context) {
	postID := c.Param("id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	posts, err := models.GetPostByID(postID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostsCreate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	post := models.Post{}
	c.ShouldBindJSON(&post)
	err := post.Create(&conn, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func PostsByCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	posts, err := models.GetPostsByCurrentUser(userID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostsUpdate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	postSent := models.Post{}
	err := c.ShouldBindJSON(&postSent)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form sent"})
		return
	}

	postBeingUpdated, err := models.GetPostByID(postSent.ID.String(), &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postBeingUpdated.AuthorID.String() != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this post"})
		return
	}

	postSent.AuthorID = postBeingUpdated.AuthorID
	err = postSent.Update(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": postSent})
}

func PostsDelete(c *gin.Context) {
	userID := c.GetString("user_id")
	postID := c.Param("id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	postBeingDeleted, err := models.GetPostByID(postID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postBeingDeleted.AuthorID.String() != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this post"})
		return
	}

	errDel := models.Delete(postID, &conn)
	if errDel != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": postBeingDeleted, "message": "The post has been deleted"})
}
