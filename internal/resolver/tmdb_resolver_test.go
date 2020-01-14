package resolver_test

import (
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ryanbradynd05/go-tmdb"
	"testing"
)

func TestTmdbSearch(t *testing.T) {
	config.ReadConfig()
	testNames := []string{
		"The Artist",
	}

	tmdbAPI := tmdb.Init(config.TmdbConfig)
	_, err := tmdbAPI.SearchMovie(testNames[0], nil)

	if err != nil {
		t.Error(err)
	}

	//fmt.Println(collectionSearchResults.Results[0])

	//for _, filename := range collectionSearchResults.Results {
	// fmt.Println(filename.Title) }

}
