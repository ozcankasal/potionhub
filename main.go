package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// Potion struct represents a magical concoction in PotionHub
type Potion struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Properties  []string
}

func main() {

	r := gin.Default()

	// Load templates
	r.SetFuncMap(template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"join":            strings.Join,
	})

	r.LoadHTMLGlob("templates/*")

	// Define routes
	r.GET("/", homeHandler)
	r.GET("/catalog", catalogHandler)
	r.GET("/potion/:id", potionDetailsHandler)

	r.Run() // Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func catalogHandler(c *gin.Context) {
	potions := fetchPotionsWithFilters(c.Query("filter"), c.Query("sort"))
	c.HTML(http.StatusOK, "catalog.html", gin.H{
		"potions": potions,
	})

}

func potionDetailsHandler(c *gin.Context) {

	potionID := c.Param("id")
	potion := fetchPotionByID(potionID)

	if potion != nil {
		c.HTML(http.StatusOK, "potion_details.html", gin.H{
			"potion": potion,
		})
	} else {
		c.HTML(http.StatusNotFound, "potion_not_found.html", gin.H{
			"potionID": potionID,
		})
	}
}

// Format float64 as a dollar amount
func formatAsDollars(value float64) string {
	return fmt.Sprintf("$%.2f", value)
}

// Check if a slice contains a given element
func contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

// Fetch the entire potion catalog
func fetchPotionCatalog() []Potion {
	// Simulate fetching data from a database or an API
	return []Potion{
		{
			ID:          "1",
			Name:        "Elixir of Endless Endurance",
			Description: "A rejuvenating brew that imbues the drinker with seemingly limitless stamina.",
			Price:       50.0,
			Properties:  []string{"Stamina", "Endurance"},
		},
		{
			ID:          "2",
			Name:        "Potion of the Wandering Mind",
			Description: "A concoction that enhances the imagination and creative thinking of the drinker.",
			Price:       25.0,
			Properties:  []string{"Creativity", "Imagination"},
		},
		{
			ID:          "3",
			Name:        "Brew of the Fierce Flame",
			Description: "A fiery brew that imbues the drinker with the power of flame.",
			Price:       75.0,
			Properties:  []string{"Fire", "Power"},
		},
		{
			ID:          "4",
			Name:        "Essence of the Silver Tongue",
			Description: "A smooth elixir that enhances the persuasive abilities of the drinker.",
			Price:       40.0,
			Properties:  []string{"Persuasion", "Charisma"},
		},
		{
			ID:          "5",
			Name:        "Potion of the Calm Seas",
			Description: "A calming brew that soothes the nerves of the drinker.",
			Price:       20.0,
			Properties:  []string{"Calm", "Relaxation"},
		},
		{
			ID:          "6",
			Name:        "Syrup of the Golden Glow",
			Description: "A sweet syrup that enhances the natural radiance of the drinker.",
			Price:       30.0,
			Properties:  []string{"Radiance", "Beauty"},
		},
		{
			ID:          "7",
			Name:        "Tonic of the Unbreakable Will",
			Description: "A potent tonic that strengthens the willpower of the drinker.",
			Price:       60.0,
			Properties:  []string{"Willpower", "Strength"},
		},
		{
			ID:          "8",
			Name:        "Philter of the Swift Feet",
			Description: "A swift potion that enhances the speed and agility of the drinker.",
			Price:       80.0,
			Properties:  []string{"Speed", "Agility"},
		},
		{
			ID:          "9",
			Name:        "Balm of the Soaring Spirit",
			Description: "A restorative balm that uplifts the spirit of the drinker.",
			Price:       35.0,
			Properties:  []string{"Uplift", "Spirit"},
		},
		{
			ID:          "10",
			Name:        "Elixir of the Shrouded Mind",
			Description: "A mysterious elixir that enhances the intuition and foresight of the drinker.",
			Price:       45.0,
			Properties:  []string{"Intuition", "Foresight"},
		},
		{
			ID:          "11",
			Name:        "Sap of the Mighty Oak",
			Description: "A thick sap that enhances the strength and endurance of the drinker.",
			Price:       70.0,
			Properties:  []string{"Strength", "Endurance"},
		},
	}
}

// Fetch potions filtered and sorted by their properties
func fetchPotionsWithFilters(filter, sortCriteria string) []Potion {
	// Simulate fetching data from a database or an API
	potions := fetchPotionCatalog()
	filteredPotions := []Potion{}

	for _, potion := range potions {
		if filter == "" || contains(potion.Properties, filter) {
			filteredPotions = append(filteredPotions, potion)
		}
	}

	if sortCriteria != "" {
		sortPotions(filteredPotions, sortCriteria)
	}

	return filteredPotions
}

// Fetch a single potion by its ID
func fetchPotionByID(id string) *Potion {
	potions := fetchPotionCatalog()
	for _, potion := range potions {
		if potion.ID == id {
			return &potion
		}
	}
	return nil
}

// Sort potions based on the given criteria
func sortPotions(potions []Potion, sortCriteria string) {
	switch sortCriteria {
	case "price_asc":
		sort.Slice(potions, func(i, j int) bool {
			return potions[i].Price < potions[j].Price
		})
	case "price_desc":
		sort.Slice(potions, func(i, j int) bool {
			return potions[i].Price > potions[j].Price
		})
	case "name_asc":
		sort.Slice(potions, func(i, j int) bool {
			return potions[i].Name < potions[j].Name
		})
	case "name_desc":
		sort.Slice(potions, func(i, j int) bool {
			return potions[i].Name > potions[j].Name
		})
	}
}
