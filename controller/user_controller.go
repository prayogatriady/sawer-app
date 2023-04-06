package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/service"
)

type UserContInterface interface {
	Signup(c *gin.Context)
	// Profile(c *gin.Context)
}

type UserController struct {
	UserService service.UserServInterface
}

func NewUserController(userService service.UserServInterface) UserContInterface {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Signup(c *gin.Context) {
	ctx := context.Background()

	var userRequest model.UserSignupRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.CreateUser(ctx, &userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User created",
		"body":    userResponse,
	})
}

// func (uc *UserController) Profile(c *gin.Context) {

// 	var word []string
// 	if err = json.NewDecoder(resp.Body).Decode(&word); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  "500 - INTERNAL SERVER ERROR",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	uc.Users[i].Word = word[0]

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  "200 - STATUS OK",
// 		"message": "User retrieved",
// 		"body":    uc.Users,
// 	})
// }
