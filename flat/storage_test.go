package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatingMemoryStorage(t *testing.T) {
	s, err := NewStorage(Memory)

	assert.Nil(t, err)
	assert.IsType(t, &StorageMemory{}, s)
}

func TestCreatingJSONStorage(t *testing.T) {
	s, err := NewStorage(JSON)

	assert.Nil(t, err)
	assert.IsType(t, &StorageJSON{}, s)
}

func TestSaveBeerReturnsErrorIfBeerAlreadyExists(t *testing.T) {
	var availableStorage []Storage
	availableStorage = append(availableStorage, new(StorageMemory))
	// TODO: this requires mocking of the interaction with to actual files
	//stgJSON, errS := NewStorageJSON(JSONDataLocation)
	//if errS != nil {
	//	t.Errorf("unexpected error while creating JSON storage: %s", errS.Error())
	//}
	//availableStorage = append(availableStorage, stgJSON)

	sampleBeer := Beer{
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
	}

	for _, storage := range availableStorage {

		err := storage.SaveBeer(sampleBeer)

		assert.Nil(t, err)

		errDupe := storage.SaveBeer(sampleBeer)

		assert.NotNil(t, errDupe)
		assert.Equal(t, "beer already exists", errDupe.Error())
	}
}