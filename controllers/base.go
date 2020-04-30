package controllers

import (
	"github.com/fabric-go-gateway/models"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  baseController
 * @Version: 1.0.0
 * @Date: 12/7/19 10:08 上午
 */

func ReturnOk(object interface{}) *models.Message {
	return &models.Message{
		Code:    "200",
		Message: "Success",
		Data:    object,
	}
}

func ReturnError(err error) *models.Message {
	return &models.Message{
		Code:    "500",
		Message: "Error",
		Data:    err.Error(),
	}
}

func ReturnUnAuthorizedMsg() *models.Message {
	return &models.Message{
		Code:    "400",
		Message: "User is not authorized",
	}
}
