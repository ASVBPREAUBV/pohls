package resolver

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/ryanbradynd05/go-tmdb"
	"path/filepath"
	"regexp"
	"strconv"
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

	fmt.Println(fileName)

	mediaName := cleanMediaName(fileName)

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

func extractSeason(filename string) (int, error) {
	seasonStrings := seasonPattern.FindStringSubmatch(filename)
	if len(seasonStrings) == 0 {
		return 99, errors.TooFewItems("seasonStrings", "seasonStrings", 2)
	} else {
		seasonString := seasonStrings[1]
		seasonNumber, err := strconv.Atoi(seasonString)
		if err != nil {
			return 0, err
		}
		return seasonNumber, nil
	}
}

func extractEpisode(filename string) (int, error) {
	episodeStrings := episodePattern.FindStringSubmatch(filename)
	fmt.Println(episodeStrings)
	if len(episodeStrings) == 0 {
		return 99, errors.TooFewItems("episodeStrings", "episodeStrings", 2)
	} else {
		episodeString := episodeStrings[1]
		episodeNumber, err := strconv.Atoi(episodeString)
		if err != nil {
			return 0, err
		}
		return episodeNumber, nil
	}
}

func extractYear(filename string) (int, error) {
	yearStrings := yearPattern.FindStringSubmatch(filename)
	if len(yearStrings) == 0 {
		return 99, errors.TooFewItems("yearStrings", "yearStrings", 2)
	} else {
		yearString := yearStrings[2]
		yearNo, err := strconv.Atoi(yearString)
		if err != nil {
			return 0, err
		}
		return yearNo, nil
	}
}

func cleanMediaName(fileName string) string {

	fileName = seasonEpisodePattern.ReplaceAllString(fileName, "")
	fileName = yearPattern.ReplaceAllString(fileName, "")

	fileName = resolutionPattern.ReplaceAllString(fileName, "")
	fileName = qualityPattern.ReplaceAllString(fileName, "")
	fileName = codecPattern.ReplaceAllString(fileName, "")
	fileName = audioPattern.ReplaceAllString(fileName, "")
	fileName = groupPattern.ReplaceAllString(fileName, "")
	fileName = regionPattern.ReplaceAllString(fileName, "")
	fileName = extendedPattern.ReplaceAllString(fileName, "")
	fileName = hardcodedPattern.ReplaceAllString(fileName, "")
	fileName = properPattern.ReplaceAllString(fileName, "")
	fileName = repackPattern.ReplaceAllString(fileName, "")
	fileName = containerPattern.ReplaceAllString(fileName, "")
	fileName = widescreenPattern.ReplaceAllString(fileName, "")
	fileName = websitePattern.ReplaceAllString(fileName, "")
	fileName = languagePattern.ReplaceAllString(fileName, "")
	fileName = sbsPattern.ReplaceAllString(fileName, "")
	fileName = unratedPattern.ReplaceAllString(fileName, "")
	fileName = sizePattern.ReplaceAllString(fileName, "")
	fileName = is3dPattern.ReplaceAllString(fileName, "")
	fileName = otherPatterns.ReplaceAllString(fileName, "")

	NoSpecialCharacters := regexp.MustCompile("[^a-zA-Z0-9]+")

	return NoSpecialCharacters.ReplaceAllString(fileName, " ")
}
