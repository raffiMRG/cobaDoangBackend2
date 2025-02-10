package controllers

import (
	"backend/model"
	"backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	var response model.BaseResponseModel
	root := "../new" // Mengatur direktori root, bisa diganti dengan direktori yang diinginkan

	allDataLocale, err := repositories.GetAllDataFromLocal(root)

	if err != nil {
		response = model.BaseResponseModel{
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.BaseResponseModel{
			Message: "Data retrieved successfully",
			Data:    allDataLocale,
		}
	}

	c.JSON(http.StatusOK, response)
}

func GetSelectedData(c *gin.Context) {
	var response model.BaseResponseModel
	selectedData, err := repositories.SelectedData()

	if err != nil {
		response = model.BaseResponseModel{
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		response = model.BaseResponseModel{
			Message: "Data retrieved successfully",
			Data:    selectedData,
		}
	}

	c.JSON(http.StatusOK, response)
}
