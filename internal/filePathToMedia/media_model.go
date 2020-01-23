package filePathToMedia

type Media struct {
	SourceFilePath      string
	DestinationFilePath string
	Title               string
	MediaType           MediaType
	Season              int
	Episode             int
	Year                int
	//	Extended   bool
	//	Hardcoded  bool
	//	Proper     bool
	//	Repack     bool
	//	Widescreen bool
	//	Unrated    bool
	//	Is3d       bool
}

type MediaType string

const (
	MediaTypeSeries MediaType = "series"
	MediaTypeMovie  MediaType = "movie"
)
