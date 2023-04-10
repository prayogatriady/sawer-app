package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/sawer-app/middleware"
	"github.com/prayogatriady/sawer-app/model"
	"github.com/prayogatriady/sawer-app/service"
)

type UserContInterface interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)

	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
	DeleteUser(c *gin.Context)
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

	userResponse, err := uc.UserService.Signup(ctx, &userRequest)
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

func (uc *UserController) Signin(c *gin.Context) {
	ctx := context.Background()

	// Get request body
	var userRequest *model.UserSigninRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	token, err := uc.UserService.Signin(ctx, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User signed in",
		"body":    token,
	})
}

func (uc *UserController) Profile(c *gin.Context) {
	ctx := context.Background()

	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.Profile(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User retrieved",
		"body":    userResponse,
	})
}

func (uc *UserController) EditProfile(c *gin.Context) {
	ctx := context.Background()

	// Get request body
	var userRequest *model.UserEditRequest
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	userResponse, err := uc.UserService.EditProfile(ctx, userId, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User updated",
		"body":    userResponse,
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	ctx := context.Background()

	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	if err := uc.UserService.DeleteUser(ctx, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User deleted",
	})
}
