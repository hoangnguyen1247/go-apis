package home

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
    g := e.Group("/home")
    {
        g.GET("/", handler.GetHome)
        g.GET("/hello", handler.GetHomeHello)
    }
}

func (handler *Controller) GetHome(ginContext *gin.Context) {
    ginContext.HTML(http.StatusOK, "home/get-home.tmpl", gin.H{})
}

func (handler *Controller) GetHomeHello(ginContext *gin.Context) {
    ginContext.HTML(http.StatusOK, "home/get-hello.tmpl", gin.H{})
}
