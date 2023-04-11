package controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ctrl *HomeController) GetHome(c *gin.Context) {
	c.HTML(200, "base", nil)
}
