package login

import (
	"log"
	"net/http"

	"github.com/amren1254/gin-docker/auth"
	"github.com/amren1254/gin-docker/database"
	"github.com/amren1254/gin-docker/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//get user details
	// var user model.UserDetails
	user := model.UserDetails{}
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
	}
	isValidUserCredentials, err := database.VerifyUserCredential(user)
	if err != nil {
		log.Println(err)
	}
	if isValidUserCredentials {
		//generate and assign token
		token, err := auth.GenerateAuthToken(user.Username)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"SuccessMessage": "UserLoggedIn", "token": token})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"WarningMessage": "Wrong Credential"})
	}
}
