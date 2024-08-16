package views

import (
	"github.com/gin-gonic/gin"
	views "gitub.com/matheus-hrm/curiously/views/pages"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", views.HomeHandler)
}
