package resolver

import (
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ryanbradynd05/go-tmdb"
	"path/filepath"
)

func FilePathListToTmDbCollectionList(filenames []string) (collections []tmdb.Collection) {

	for _, filename := range filenames {
		collection, err := FilePathToTmdbCollection(filename)
		if err != nil {
			fmt.Print(err)
		} else {
			collections = append(collections, collection)
		}
	}

	return collections
}

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
	config := tmdb.Config{
		APIKey:   config.TmdbToken,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI := tmdb.Init(config)
	collectionSearchResults, err := tmdbAPI.SearchCollection(media.Name, nil)

	firstResult := collectionSearchResults.Results[0]

	collection, err := tmdbAPI.GetCollectionInfo(firstResult.ID, nil)

	if err != nil {
		return *collection, err
	}

	return *collection, nil
}

func FilePathToMedia(fileName string) Media {
	media := Media{}

	seasonNo, seasonError := extractSeason(fileName)
	if seasonError == nil {
		media.Season = seasonNo
	}

	eposodeNo, episodeError := extractEpisode(fileName)
	if episodeError == nil {
		media.Episode = eposodeNo
	}

	if episodeError == nil && seasonError == nil {
		media.MediaType = MediaTypeSeries
	} else {
		media.MediaType = MediaTypeMovie
	}

	year, yearError := extractYear(fileName)
	if yearError == nil {
		media.Year = year
	}

	//fmt.Println(fileName)

	mediaName := NameDetox(fileName)

	media.Name = mediaName

	//fmt.Println(media)

	/*episodeStrings := episodePattern.FindStringSubmatch(fileName)
	fmt.Println(episodeStrings)

	res := seasonPattern.ReplaceAllString(fileName, "")
	fmt.Println(res)*/

	/*	if len(seasonStrings) > 11 {
			media.MediaType = MediaTypeSeries
			seasonString := seasonStrings[1]
			seasonNumber, seasonError := strconv.Atoi(seasonString)
			if seasonError != nil {
				panic(seasonError)
			}
			media.Season = seasonNumber
		}*/

	//fmt.Print(seasonPattern.(fileName))
	return media
}
