package resolver

import (
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ryanbradynd05/go-tmdb"
	"path/filepath"
)

func FilePathToTmdbCollection(string string) (tmdb.Collection, error) {
	file := filepath.Base(string)
	media := FilePathToMedia(file)
	//TODO media health check

	collection, err := MediaToTmdbCollection(media)
	if err != nil {
		return collection, err
	}
	return collection, nil
}

func MediaToTmdbCollection(media Media) (tmdb.Collection, error) {
	fmt.Println(media)

	config := tmdb.Config{
		APIKey:   config.TmdbToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI := tmdb.Init(config)
	collectionSearchResults, err := tmdbAPI.SearchCollection(media.Name, nil)

	if err != nil {
		return tmdb.Collection{}, err
	}

	fmt.Println(collectionSearchResults.Results[0])


	firstResult := collectionSearchResults.Results[0]

	collection, err := tmdbAPI.GetCollectionInfo(firstResult.ID, nil)

	if err != nil {
		return *collection, err
	}

	return *collection, nil
}
