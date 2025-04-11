package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thiagomelosz/CRUD---Go/src/configuration/logger"
	"github.com/Thiagomelosz/CRUD---Go/src/configuration/validation"
	"github.com/Thiagomelosz/CRUD---Go/src/model/request"
	"github.com/Thiagomelosz/CRUD---Go/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
