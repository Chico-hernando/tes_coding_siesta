package question2

import (
	"net/http"
	"tes_coding_siesta/init/response"
	"tes_coding_siesta/models/model_question2"

	"github.com/gin-gonic/gin"
)

func GetQuestion2(c *gin.Context) {

	var response response.Response

	lat := c.Query("lat")
	long := c.Query("long")
	result, err := model_question2.Question2(lat, long)
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
