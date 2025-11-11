package controller

import (
	"fmt"

	"github.com/JoePereiraa/first-crud-golang/src/configuration/rest_err"
	"github.com/JoePereiraa/first-crud-golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError("Some fields are incorrect")

		context.JSON(restErr.Code, restErr)
		fmt.Println(err)
		return
	}

	fmt.Println(userRequest)
}
