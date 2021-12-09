package post

import (
	"friday/server/models"
	"friday/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	posts, err := models.GetAllPosts()
	utils.FatalError{Error: err}.Handle()

	c.JSON(http.StatusOK, posts)
}