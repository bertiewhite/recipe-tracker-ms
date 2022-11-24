package datastore

import (
	"time"
)

type Recipe struct {
	ID           string
	Author       string
	Method       string
	Ingredients  string
	Rating       int
	CreationDate time.Time
}

type MadeLog struct {
	RecipeID string
	Date     time.Time
}

type DataStore interface {
	AddRecipe(Recipe)
	GetRecipe(string) Recipe
}
