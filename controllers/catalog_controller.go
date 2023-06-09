package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/services"
)

type CatalogController struct {
	potionService services.PotionService
}

func NewCatalogController() *CatalogController {
	potionService := services.NewPotionService()

	return &CatalogController{
		potionService: potionService,
	}
}

func (cc *CatalogController) GetCatalog(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")

	potions := cc.potionService.FetchPotionCatalog(filter, sort)
	fmt.Println(potions)
	c.HTML(200, "catalog.html", gin.H{
		"potions": potions,
		"title":   "Potion Catalog",
	})
}
