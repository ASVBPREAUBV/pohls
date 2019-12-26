package resolver

import (
	"fmt"
	"github.com/ryanbradynd05/go-tmdb"
	"path/filepath"
)

func FilePathToTmdbCollection(string string) tmdb.Collection {
	file := filepath.Base(string)
	FilePathToMedia(file)
	bla := tmdb.Collection{}
	return bla
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

	fmt.Println(media)

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

