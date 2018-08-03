package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSaveBeerInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)
	sampleBeer := Beer{
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
	}

	err := storage.SaveBeer(sampleBeer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.cellar))

	beer := storage.cellar[0]
	assert.NotNil(t, sampleBeer.ID)
	assert.Equal(t, sampleBeer.Name, beer.Name)
	assert.Equal(t, sampleBeer.Brewery, beer.Brewery)
	assert.Equal(t, sampleBeer.Abv, beer.Abv)
	assert.Equal(t, sampleBeer.ShortDesc, beer.ShortDesc)
	assert.NotNil(t, sampleBeer.Created)
	assert.True(t, sampleBeer.Created.Before(time.Now()))
}

func TestSaveBeerReturnsErrorIfBeerAlreadyExistsInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)
	sampleBeer := Beer{
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
	}

	err := storage.SaveBeer(sampleBeer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.cellar))

	errDupe := storage.SaveBeer(sampleBeer)

	assert.NotNil(t, errDupe)
	assert.Equal(t, "beer already exists", errDupe.Error())
}

func TestSaveMultipleBeersInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	sampleBeer1 := Beer{
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
	}

	sampleBeer2 := Beer{
		Name:    "Tecate",
		Brewery: "Cuahutemoc Moctezuma",
		Abv:     5,
		ShortDesc: "Very smooth, medium bodied brew. Malt sweetness is thin, and can be likened to diluted sugar water. " +
			"Touch of fructose-like sweetness. Light citric hop flavours gently prick the palate with tea-like notes that follow and fade quickly. " +
			"Finishes a bit dry with husk tannins and a pasty mouthfeel.",
	}

	err := storage.SaveBeer(sampleBeer1, sampleBeer2)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(storage.cellar))
}

func TestSaveReviewInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	sampleBeer := Beer{
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
	}

	err := storage.SaveBeer(sampleBeer)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.cellar))

	sampleReview := Review{
		BeerID:    1,
		FirstName: "Wonder",
		LastName:  "Woman",
		Score:     8,
		Text:      "Nice beer.",
	}

	errR := storage.SaveReview(sampleReview)

	assert.Nil(t, errR)
	assert.Equal(t, 1, len(storage.reviews))

	review := storage.reviews[0]
	assert.NotNil(t, review.ID)
	assert.Equal(t, sampleReview.FirstName, review.FirstName)
	assert.Equal(t, sampleReview.LastName, review.LastName)
	assert.Equal(t, sampleReview.Score, review.Score)
	assert.Equal(t, sampleReview.Text, review.Text)
	assert.NotNil(t, sampleReview.Created)
	assert.True(t, sampleReview.Created.Before(time.Now()))
}

func TestFindBeersReturnsExpectedResultInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	sampleBeer1 := Beer{
		ID:   1,
		Name: "Pliny the Elder",
	}

	sampleBeer2 := Beer{
		ID:   2,
		Name: "Bath Ale",
	}

	storage.cellar = append(storage.cellar, sampleBeer1)
	storage.cellar = append(storage.cellar, sampleBeer2)

	result, err := storage.FindBeers()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Contains(t, result, sampleBeer1)
	assert.Contains(t, result, sampleBeer2)
}

func TestFindBeersReturnsEmptyResultIfNoBeersFoundInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	result, err := storage.FindBeers()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestFindBeerReturnsExpectedResultInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	sampleBeer := Beer{
		ID:   1,
		Name: "Pliny the Elder",
	}

	storage.cellar = append(storage.cellar, sampleBeer)

	result, err := storage.FindBeer(Beer{ID: 1})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, sampleBeer.ID, result[0].ID)
	assert.Equal(t, sampleBeer.Name, result[0].Name)
}

func TestFindBeerReturnsEmptyResultIfNoMatchingBeersFoundInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	result, err := storage.FindBeer(Beer{ID: 1})

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestFindBeerReturnsErrorIfNoBeerIDGivenInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	_, err := storage.FindBeer(Beer{})

	assert.NotNil(t, err)
	assert.Equal(t, "no beer ID specified", err.Error())
}

func TestFindReviewReturnsExpectedResultInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	sampleReview := Review{
		BeerID: 2,
		FirstName: "John",
		LastName:  "Doe",
		Score: 3,
	}

	storage.reviews = append(storage.reviews, sampleReview)

	result, err := storage.FindReview(Review{BeerID: 2})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, sampleReview.BeerID, result[0].BeerID)
	assert.Equal(t, sampleReview.FirstName, result[0].FirstName)
	assert.Equal(t, sampleReview.LastName, result[0].LastName)
	assert.Equal(t, sampleReview.Score, result[0].Score)
}

func TestFindReviewReturnsEmptyResultIfNoMatchingReviewsFoundInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	result, err := storage.FindReview(Review{BeerID: 1})

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestFindReviewReturnsErrorIfNoBeerIDGivenInMemoryStorage(t *testing.T) {
	storage := new(StorageMemory)

	_, err := storage.FindReview(Review{})

	assert.NotNil(t, err)
	assert.Equal(t, "no beer ID specified", err.Error())
}