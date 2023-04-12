package services

import (
	"sort"
	"strings"

	"github.com/ozcankasal/potionhub/models"
)

type PotionService interface {
	FetchPotionCatalog(filter string, sortOption string) []models.Potion
}

type potionServiceImpl struct {
	potions []models.Potion
}

func NewPotionService() PotionService {
	return &potionServiceImpl{
		potions: []models.Potion{
			{Name: "Potion of Healing", Description: "Heals the drinker, restoring health.", Price: 10.0, Properties: []string{}},
			{Name: "Potion of Invisibility", Description: "Makes the drinker invisible for a short duration.", Price: 20.0, Properties: []string{}},
			{Name: "Elixir of Strength", Description: "Temporarily increases the drinker's strength.", Price: 15.0, Properties: []string{}},
			{Name: "Elixir of Wisdom", Description: "Temporarily increases the drinker's wisdom.", Price: 15.0, Properties: []string{}},
			{Name: "Potion of Swiftness", Description: "Increases the drinker's movement speed for a short duration.", Price: 12.0, Properties: []string{}},
			{Name: "Potion of Poison", Description: "Poisons the drinker, dealing damage over time.", Price: 5.0, Properties: []string{}},
		},
	}
}

func (ps *potionServiceImpl) FetchPotionCatalog(filter string, sortOption string) []models.Potion {
	filteredPotions := []models.Potion{}

	// Filter potions
	for _, potion := range ps.potions {
		if filter == "" || strings.Contains(strings.ToLower(potion.Name), strings.ToLower(filter)) {
			filteredPotions = append(filteredPotions, potion)
		}
	}

	// Sort potions
	switch sortOption {
	case "name_asc":
		sort.Slice(filteredPotions, func(i, j int) bool {
			return filteredPotions[i].Name < filteredPotions[j].Name
		})
	case "name_desc":
		sort.Slice(filteredPotions, func(i, j int) bool {
			return filteredPotions[i].Name > filteredPotions[j].Name
		})
	case "price_asc":
		sort.Slice(filteredPotions, func(i, j int) bool {
			return filteredPotions[i].Price < filteredPotions[j].Price
		})
	case "price_desc":
		sort.Slice(filteredPotions, func(i, j int) bool {
			return filteredPotions[i].Price > filteredPotions[j].Price
		})
	}

	return filteredPotions
}
