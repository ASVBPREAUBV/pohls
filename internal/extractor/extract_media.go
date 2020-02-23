package extractor

func FileNameToMedia(filenames []string) (media []Media) {
	for _, filename := range filenames {
		rawMedia := FilePathToMedia(filename)
		rawMedia.SourceFilePath = filename
		parsedMedia, err := ParseMediaThroughTmdb(rawMedia)
		if err == nil {
			media = append(media, rawMedia)
		} else {
			media = append(media, parsedMedia)
		}

	}
	return media
}
