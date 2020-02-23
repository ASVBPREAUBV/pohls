package detox

import (
	"github.com/go-openapi/errors"
	"strconv"
)

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

