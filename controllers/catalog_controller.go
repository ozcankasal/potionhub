// controllers/catalog_controller.go

package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/services"
)

// CatalogController is the trusted guardian of the Potion Catalog,
// a tome of great power containing the secrets of countless magical concoctions.
// Only those deemed worthy may gaze upon its pages.
type CatalogController struct {
	potionService services.PotionService
}

// NewCatalogController is a spell that summons the CatalogController,
// binding it with the ancient knowledge of the potionService.
func NewCatalogController() *CatalogController {
	potionService := services.NewPotionService()

	return &CatalogController{
		potionService: potionService,
	}
}

// GetCatalog is a magical incantation that reveals the Potion Catalog
// to travelers who have successfully navigated the enchanted barriers.
func (cc *CatalogController) GetCatalog(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")

	// The mystical runes for filtering and sorting the catalog's contents
	filter := c.Query("filter")
	sort := c.Query("sort")

	// Unveil the hidden potions based on the traveler's desires
	potions := cc.potionService.FetchPotionCatalog(filter, sort)
	fmt.Println(potions)
	// Present the sacred tome to the worthy traveler
	c.HTML(200, "catalog.html", gin.H{
		"potions":  potions,
		"title":    "Potion Catalog",
		"username": username.(string),
	})
}
