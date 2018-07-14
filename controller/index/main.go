package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	name string
}

func New() (*IndexController, error) {
	return &IndexController{
	}, nil
}

func (handler *IndexController) SetRoutes(e *gin.Engine) {
	g := e.Group("/")
	{
		g.GET("/:params", handler.GetIndex)
	}
}

func (handler *IndexController) GetIndex(ginContext *gin.Context) {
	ginContext.HTML(http.StatusOK, "index/get-index.tmpl", gin.H{})
}