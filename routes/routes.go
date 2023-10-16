package routes

import (
	"tes_coding_siesta/controllers/question1"
	"tes_coding_siesta/controllers/question2"
	"tes_coding_siesta/controllers/question3"
	"tes_coding_siesta/controllers/question4"
	"tes_coding_siesta/controllers/question5"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/question/1", question1.GetQuestion1)
	r.GET("/question/2", question2.GetQuestion2)
	r.POST("/question/3", question3.GetQuestion3)
	r.GET("/question/4", question4.GetQuestion4)
	r.GET("/question/5", question5.GetQuestion5)

	return r
}
