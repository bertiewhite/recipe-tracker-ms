package datastore

import (
	"fmt"
	"sync"
)

type InMemoryDataStore struct {
	store map[string]Recipe
	lock  sync.Mutex
}

const RecipeExistsError = "recipe with that ID already exists"

func (ds *InMemoryDataStore) AddRecipe(recipe Recipe) error {
	if _, ok := ds.store[recipe.ID]; ok {
		return fmt.Errorf(RecipeExistsError)
	}
	ds.lock.Lock()
	defer ds.lock.Unlock()
	ds.store[recipe.ID] = recipe
	return nil
}
