package question1

import (
	"net/http"
	"tes_coding_siesta/init/response"
	"tes_coding_siesta/models/model_question1"

	"github.com/gin-gonic/gin"
)

func GetQuestion1(c *gin.Context) {

	var response response.Response

	result, err := model_question1.Question1()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success Response"
	response.Data = result

	c.JSON(http.StatusOK, response)
}
