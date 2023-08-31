package signup

import (
	"net/http"

	"github.com/amren1254/gin-docker/database"
	"github.com/amren1254/gin-docker/model"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func Signup(c *gin.Context) {
	var user model.UserDetails
	var db database.DatabaseRepository
	if err := c.BindJSON(&user); err != nil {
		return
	}
	db.InsertUser(user)
	c.IndentedJSON(http.StatusCreated, Response{Message: "User Created"})
}
