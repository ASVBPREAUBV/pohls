package resolver_test

import (
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ryanbradynd05/go-tmdb"
	"testing"
)

func TestTmdbSearch(t *testing.T) {
	config.ReadConfig()
	testNames := []string{
		"The Artist",
	}

	config := tmdb.Config{
		APIKey:   config.TmdbToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI := tmdb.Init(config)
	collectionSearchResults, err := tmdbAPI.SearchMovie(testNames[0], nil)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(collectionSearchResults.Results[0])

	//for _, filename := range collectionSearchResults.Results {
	// fmt.Println(filename.Title) }

}
