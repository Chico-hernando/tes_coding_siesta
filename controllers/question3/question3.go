package question3

import (
	"net/http"
	"tes_coding_siesta/init/question3"
	"tes_coding_siesta/init/response"
	"tes_coding_siesta/models/model_question3"

	"github.com/gin-gonic/gin"
)

func GetQuestion3(c *gin.Context) {

	var response response.Response
	var request question3.PhoneNumber
	if err := c.ShouldBindJSON(&request); err != nil {

		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := model_question3.Question3(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Success Response"
	response.Data = result

	c.JSON(http.StatusCreated, response)
}
