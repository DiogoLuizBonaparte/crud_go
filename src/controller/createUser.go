package controller

import (
	"fmt"
	"log"

	validation "github.com/DiogoLuizBonaparte/go_api.git/src/configuration/validation"

	"github.com/DiogoLuizBonaparte/go_api.git/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	log.Println("Init createUser controller")
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

}
