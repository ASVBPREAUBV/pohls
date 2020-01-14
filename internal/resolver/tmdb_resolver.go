package resolver

import (
	"errors"
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ryanbradynd05/go-tmdb"
)

func ParseMediaThroughTmdb(media Media) (Media, error) {
	if media.MediaType == MediaTypeMovie {
		return MediaToTmdbMovie(media)
	}
	if media.MediaType == MediaTypeSeries {
		return MediaToTmdbSeries(media)
	}
	return Media{}, errors.New("no media found")
}

func MediaToTmdbMovie(media Media) (Media, error) {
	fmt.Println(media)

	tmdbAPI := tmdb.Init(config.TmdbConfig)

	movieSearchResults, err := tmdbAPI.SearchMovie(media.Title, nil)

	if err != nil {
		return media, err
	}

	if len(movieSearchResults.Results) > 0 {
		firstResultTitle := movieSearchResults.Results[0].Title
		media.Title = firstResultTitle
		return media, err
	} else {
		return media, errors.New("no movie found")
	}
}

func MediaToTmdbSeries(media Media) (Media, error) {
	fmt.Println(media)

	tmdbAPI := tmdb.Init(config.TmdbConfig)

	movieSearchResults, err := tmdbAPI.SearchTv(media.Title, nil)

	if err != nil {
		return media, err
	}

	if len(movieSearchResults.Results) > 0 {
		firstResultTitle := movieSearchResults.Results[0].Name
		media.Title = firstResultTitle
		return media, err
	} else {
		return media, errors.New("no series found")
	}
}
