package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	name string
}

func New() (*Controller, error) {
	return &Controller{
	}, nil
}

func (handler *Controller) BindGin(e *gin.Engine) {
	g := e.Group("/")
	{
		g.GET("/", handler.GetIndex)
		g.GET("/hello", handler.GetHello)
	}
}

func (handler *Controller) GetIndex(ginContext *gin.Context) {
	ginContext.HTML(http.StatusOK, "index/get-index.tmpl", gin.H{})
}

func (handler *Controller) GetHello(ginContext *gin.Context) {
	ginContext.HTML(http.StatusOK, "index/get-hello.tmpl", gin.H{})
}
