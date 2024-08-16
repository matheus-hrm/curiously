package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	homePage "gitub.com/matheus-hrm/curiously/templates/home"
	"gitub.com/matheus-hrm/curiously/utils"
)

func HomeHandler(c *gin.Context) {
	err := homePage.BaseLayout().Render(c, c.Writer)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, err)
	}
}
