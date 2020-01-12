package resolver_test

import (
	"fmt"
	"github.com/ryanbradynd05/go-tmdb"
	"testing"
)

func TestTmdbSearch(t *testing.T) {
	testNames := []string{
		"The Artist",
	}

	config := tmdb.Config{
		APIKey:   "",
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI := tmdb.Init(config)
	collectionSearchResults, err := tmdbAPI.SearchCollection(testNames[0], nil)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(collectionSearchResults.Results[0].Name)

	//for _, filename := range collectionSearchResults.Results {
	// fmt.Println(filename.Title) }

}
