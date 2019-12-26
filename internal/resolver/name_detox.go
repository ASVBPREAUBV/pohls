package resolver

import (
	"regexp"
)


func NameDetox(fileName string) string {

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

	// if everything is removed dont bother

	return NoSpecialCharacters.ReplaceAllString(fileName, " ")
}
